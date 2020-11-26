package main

import (
	"github.com/go-ex/console"
	"log"
)

func main() {
	cmd := console.New()
	cmd.Init()

	has := cmd.HasCommand()
	if !has {
		cmd.RunHelp()
	} else {
		if err := cmd.RunCommand(); err != nil {
			log.Println(err)
		}
	}
}
