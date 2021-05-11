package main

import (
	"log"

	"github.com/barelyhuman/mobile-version-sync/cmd"
)

func main() {
	done := cmd.Setup()
	if !done {
		return
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
