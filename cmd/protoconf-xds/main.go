package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/protoconf/protoconf-xds/cmd/serve"
)

func main() {
	ui := cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	cmd := cli.NewCLI("protoconf-xds", "1")
	cmd.Args = os.Args[1:]
	cmd.Commands = map[string]cli.CommandFactory{
		"serve": serve.NewCommand,
	}

	code, err := cmd.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("error: %v", err))

	}
	os.Exit(code)

}
