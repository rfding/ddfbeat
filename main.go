package main

import (
	"os"

	"github.com/rfding/connextametricbeat/cmd"

	_ "github.com/rfding/connextametricbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
