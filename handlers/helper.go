package handlers

import (
	"io/ioutil"
	re "regexp"
	"strings"
)

const OPEN_BRACES = "{"
const CLOSE_BRACES = "}"
const COMMENT_SINGLE_LINE = "//"

type ASTNode struct {
	Type     string
	Name     string
	Depth    int
	Children []ASTItem
}

// ASTItem - the actual item
type ASTItem struct {
	ASTNode
	Parent *ASTItem
}

type AST struct {
	Root ASTItem
}

type UsableAST struct {
	Type     string
	Name     string
	Depth    int
	Children []UsableAST
}

func (astInstance *UsableAST) Find(name string) *UsableAST {
	if astInstance.Name == name {
		return astInstance
	}

	for _, node := range astInstance.Children {
		if node.Name == name {
			return &node
		}
		finding := node.Find(name)
		if finding != nil && finding.Name == name {
			return finding
		}
	}

	return nil
}

func (ast AST) Usable() UsableAST {
	root := ast.Root
	var usableAST = UsableAST{
		Name:     root.Name,
		Type:     root.Type,
		Depth:    0,
		Children: []UsableAST{},
	}

	for _, item := range root.Children {
		usableAST.Children = append(usableAST.Children, item.Usable())
	}

	return usableAST
}

func (node ASTNode) Usable() UsableAST {
	var usable = UsableAST{
		Name:     node.Name,
		Type:     node.Type,
		Depth:    node.Depth,
		Children: []UsableAST{},
	}

	for _, item := range node.Children {
		usable.Children = append(usable.Children, item.Usable())
	}

	return usable
}

func readGradleFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func parseToken(stringData string) []([]string) {
	rows := strings.Split(stringData, "\n")
	var tokens [][]string

	for _, item := range rows {
		if len(item) <= 0 {
			continue
		}

		whitespaceCheck := re.MustCompile(`\s`)
		lineTokens := whitespaceCheck.Split(item, -1)
		tokens = append(tokens, lineTokens)
	}

	return tokens
}

func contructAST(tokenSet [][]string) UsableAST {
	var objectRoot *ASTItem
	// var bracketStack []string
	ast := &AST{
		Root: ASTItem{
			ASTNode: ASTNode{
				Type:     "dict",
				Name:     "root",
				Depth:    0,
				Children: []ASTItem{},
			},
			Parent: &ASTItem{},
		},
	}

	objectRoot = &ast.Root

	for _, line := range tokenSet {

		lineString := strings.TrimSpace(strings.Join(line, ""))

		if len(lineString) <= 0 {
			continue
		}

		if len(lineString) >= 2 && lineString[0:len(COMMENT_SINGLE_LINE)] == COMMENT_SINGLE_LINE {
			continue
		}

		var values []string
		var key string
		for index, token := range line {

			if len(token) <= 0 {
				continue
			}

			switch token {
			case OPEN_BRACES:
				{
					// bracketStack = append(bracketStack, OPEN_BRACES)
					key = line[index-1]
					objectRoot.Children = append(objectRoot.Children, ASTItem{
						ASTNode: ASTNode{
							Type:     "dict",
							Name:     key,
							Depth:    objectRoot.Depth + 1,
							Children: []ASTItem{},
						},
						Parent: objectRoot,
					})

					objectRoot = &objectRoot.Children[len(objectRoot.Children)-1]

					break
				}
			case CLOSE_BRACES:
				{
					// bracketStack = append(bracketStack, CLOSE_BRACES)
					objectRoot = objectRoot.Parent
					break
				}
			default:
				{
					values = append(values, token)
				}
			}
		}

		if len(values) <= 0 {
			continue
		}

		if values[0] == "" || values[0] == objectRoot.Name {
			continue
		}

		keyItem := ASTItem{
			ASTNode: ASTNode{
				Type:     "key",
				Name:     values[0],
				Children: []ASTItem{},
				Depth:    objectRoot.Depth + 1,
			},
			Parent: objectRoot,
		}

		keyItem.ASTNode.Children = append(keyItem.ASTNode.Children, ASTItem{
			ASTNode{
				Type:     "value",
				Name:     strings.Join(values[1:], " "),
				Depth:    keyItem.Depth + 1,
				Children: []ASTItem{},
			},
			&ASTItem{},
		},
		)

		objectRoot.Children = append(objectRoot.Children, keyItem)
	}

	return ast.Usable()
}
