package main

import (
	"os"

	"github.com/connexta/ddfbeat/cmd"

	_ "github.com/connexta/ddfbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
