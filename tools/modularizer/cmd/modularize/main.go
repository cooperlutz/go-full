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
	println("  modularizer modularize -name=<module_name> -aggRoot=<aggregate_root_entity> -command=<command1> -command=<command2> -event=<event1> -event=<event2> -query=<query1> -query=<query2>")
	println("  modularizer help")
}

// // Define a custom type that is a slice of strings
// type stringSlice []string

// // Implement the flag.Value interface's Set method
// func (s *stringSlice) Set(val string) error {
// 	*s = append(*s, val)
// 	return nil
// }

// // Implement the flag.Value interface's String method
// func (s *stringSlice) String() string {
// 	return fmt.Sprintf("%v", *s)
// }

// func main() {
// 	moduleCmd := flag.NewFlagSet("modularize", flag.ExitOnError)

// 	moduleName := moduleCmd.String("name", "", "Name of the module to create")
// 	moduleEntities := &stringSlice{}
// 	moduleCmd.Var(moduleEntities, "entity", "Entities for the module (can specify multiple)")
// 	// moduleAggregateRootEntity := moduleCmd.String("aggRoot", "", "Name of the aggregate root entity")

// 	moduleCommands := &stringSlice{}
// 	moduleCmd.Var(moduleCommands, "command", "Commands for the module (can specify multiple)")
// 	moduleEvents := &stringSlice{}
// 	moduleCmd.Var(moduleEvents, "event", "Events for the module (can specify multiple)")
// 	moduleQueries := &stringSlice{}
// 	moduleCmd.Var(moduleQueries, "query", "Queries for the module (can specify multiple)")

// 	if len(os.Args) < 2 {
// 		help()
// 		return
// 	}

// 	switch os.Args[1] {
// 	case "modularize":

// 		moduleCmd.Parse(os.Args[2:])

// 		if *moduleName == "" {
// 			panic("Module name is required")
// 		}

// 		modularize := modularizer.NewModularizer(
// 			*moduleName,
// 			// *moduleAggregateRootEntity,
// 			*moduleEntities,
// 			*moduleCommands,
// 			*moduleEvents,
// 			*moduleQueries,
// 		)
// 		if err := modularize.CreateModule(); err != nil {
// 			panic(err)
// 		}

// 	case "help":
// 		help()
// 	default:
// 		println("Unknown command:", os.Args[1])
// 	}
// }

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
