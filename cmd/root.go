package cmd

import (
	cmd "github.com/elastic/beats/libbeat/cmd"
	"github.com/connexta/ddfbeat/beater"
)

// Name of this beat
var Name = "ddfbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
