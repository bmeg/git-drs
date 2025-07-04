package main

import (
	"log"
	"os"

	"github.com/bmeg/git-drs/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println("Error:", err.Error())
		os.Exit(1)
	}
}
