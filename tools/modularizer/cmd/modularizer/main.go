package main

import (
	"flag"
	"log/slog"
	"os"

	"go.yaml.in/yaml/v2"

	"github.com/cooperlutz/go-full/tools/modularizer"
	"github.com/cooperlutz/go-full/tools/modularizer/module"
)

func help() {
	println("Usage:")
	println("  modularizer modularize --config=path/to/modularizer.yaml")
	println("  modularizer help")
}

func main() {
	modularize := flag.NewFlagSet("modularize", flag.ExitOnError)
	modularizerConfig := modularize.String("config", "tools/modularizer/modularizer.yaml", "Path to the modularizer configuration file")

	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {
	case "modularize":
		mod := module.ModuleConfig{}
		f, err := os.ReadFile(*modularizerConfig)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(f, &mod)
		if err != nil {
			panic(err)
		}
		slog.Info("Loaded modularizer configuration:", "config", mod)
		mods := modularizer.FromConfig(mod)
		for _, m := range mods {
			if err := m.CreateModule(); err != nil {
				panic(err)
			}
		}
		completionMessage := `
Module created successfully!

Next steps:
1. Review and update the generated code as needed
2. Implement the module in the app/app.go
3. Implement the frontend app in the api/frontend/src/app/router
`
		slog.Info(completionMessage)
	case "help":
		help()
	default:
		println("Unknown command:", os.Args[1])
	}
}
