package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var filename string

func main() {
	defaults := pflag.NewFlagSet("defaults for all commands", pflag.ExitOnError)

	cmdGen := pflag.NewFlagSet("gen", pflag.ExitOnError)
	cmdGen.StringVarP(&filename, "filename", "f", "spec.yaml", "spec filename path")
	cmdGen.AddFlagSet(defaults)

	if len(os.Args) == 1 {
		fmt.Println("no subcommand given")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gen":
		cmdGen.Parse(os.Args[2:])
		fmt.Println(filename)
	default:
		fmt.Println("no define subcommand")
		os.Exit(1)
	}
}
