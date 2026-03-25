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
	modularizerConfig := modularize.String("config", "", "Path to the modularizer configuration file")

	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {
	case "modularize":
		modularize.Parse(os.Args[2:])
		mod := module.ModuleConfig{}
		wd, err := os.Getwd()
		if err != nil {
			slog.Error("Error getting current working directory:", "error", err)
			return
		}
		f, err := os.ReadFile(wd + "/" + *modularizerConfig)
		if err != nil {
			slog.Error("Error reading modularizer configuration file:", "error", err)
			return
		}
		err = yaml.Unmarshal(f, &mod)
		if err != nil {
			slog.Error("Error unmarshalling modularizer configuration file:", "error", err)
			return
		}
		slog.Info("Loaded modularizer configuration:", "config", mod)
		mods, err := modularizer.FromConfig(mod)
		if err != nil {
			slog.Error("Error creating modules from configuration:", "error", err)
			return
		}
		for _, m := range mods {
			if err := m.CreateModule(); err != nil {
				slog.Error("Error creating module:", "error", err)
				return
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
		slog.Error("Unknown command:", "command", os.Args[1])
	}
}
