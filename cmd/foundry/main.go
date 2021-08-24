package main

import (
	"fmt"
	"os"

	"github.com/smiyaguchi/foundry/internal/gen"
	"github.com/smiyaguchi/foundry/internal/spec"
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
		spec, err := spec.Load(filename)
		if err != nil {
			fmt.Printf("failed to load spec file: %v\n", err)
			os.Exit(1)
		}
		s, err := gen.Convert(spec)
		if err != nil {
			fmt.Printf("failed to generate data: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(s)
	default:
		fmt.Println("no define subcommand")
		os.Exit(1)
	}
}
