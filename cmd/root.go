package cmd

import (
	cmd "github.com/elastic/beats/libbeat/cmd"
	"github.com/rfding/connextametricbeat/beater"
)

// Name of this beat
var Name = "connextametricbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
