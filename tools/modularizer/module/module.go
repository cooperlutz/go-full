package module

import (
	"github.com/cooperlutz/go-full/tools/modularizer/utils"
)

// must be in snake_case during initialization
type StringOfVaryingCases string

// Flat returns the string in flat case (e.g. "useraccount")
func (s StringOfVaryingCases) Flat() string {
	return utils.SnakeToFlat(string(s))
}

// Pascal returns the string in PascalCase (e.g. "UserAccount")
func (s StringOfVaryingCases) Pascal() string {
	return utils.SnakeToPascal(string(s))
}

// Kebab returns the string in kebab-case (e.g. "user-account")
func (s StringOfVaryingCases) Kebab() string {
	return utils.SnakeToKebab(string(s))
}

// Camel returns the string in camelCase (e.g. "userAccount")
func (s StringOfVaryingCases) Camel() string {
	return utils.SnakeToCamel(string(s))
}

// Snake returns the string in snake_case (e.g. "user_account") -
// this is just the original string, but we include this method for consistency and to make it clear that the original string should be in snake_case
func (s StringOfVaryingCases) Snake() string {
	return string(s)
}

// FirstLetterLower returns the first letter of the string in lowercase (e.g. "u")
func (s StringOfVaryingCases) FirstLetterLower() string {
	return utils.FirstLetter(string(s))
}

type ModuleConfig struct {
	Modules []Module `yaml:"modules"`
}

// Module encompasses all the information about a module, including its name, description, entities, commands, events, and queries. It also includes helper methods to get the aggregate root entity, non-aggregate entities, emitted events, and consumed events for the module.
type Module struct {
	Name           StringOfVaryingCases `yaml:"name"`
	Description    string               `yaml:"description"`
	DefaultQueries bool                 `yaml:"defaultQueries"`
	Entities       []Entity             `yaml:"entities"`
	Commands       []Command            `yaml:"commands"`
	Events         []Event              `yaml:"events"`
	Queries        []Query              `yaml:"queries"`
}

// AggregateRoot returns the aggregate root entity for the module,
// which is the entity that has IsAggregateRoot set to true.
// If no entity is marked as the aggregate root, it returns an empty Entity struct.
func (m Module) AggregateRoot() Entity {
	for _, entity := range m.Entities {
		if entity.IsAggregateRoot {
			return entity
		}
	}
	return Entity{}
}

// NonAggregateEntities returns a slice of all the entities in the module that are not marked as the aggregate root (i.e. IsAggregateRoot is false).
func (m Module) NonAggregateEntities() []Entity {
	var nonAggregateEntities []Entity
	for _, entity := range m.Entities {
		if !entity.IsAggregateRoot {
			nonAggregateEntities = append(nonAggregateEntities, entity)
		}
	}
	return nonAggregateEntities
}

// EmittedEvents returns a slice of all the events in the module that are marked as emitted (i.e. Kind is "emitted").
func (m Module) EmittedEvents() []Event {
	var emittedEvents []Event
	for _, event := range m.Events {
		if event.Kind == "emitted" {
			emittedEvents = append(emittedEvents, event)
		}
	}
	return emittedEvents
}

// ConsumedEvents returns a slice of all the events in the module that are marked as consumed (i.e. Kind is "consumed").
func (m Module) ConsumedEvents() []Event {
	var consumedEvents []Event
	for _, event := range m.Events {
		if event.Kind == "consumed" {
			consumedEvents = append(consumedEvents, event)
		}
	}
	return consumedEvents
}

// Entity represents an entity in the module, which has a name, description, fields,
// and a boolean indicating whether it is the aggregate root entity for the module.
type Entity struct {
	Name            StringOfVaryingCases `yaml:"name"`
	Description     string               `yaml:"description"`
	Fields          []Field              `yaml:"fields"`
	IsAggregateRoot bool                 `yaml:"isAggregateRoot"`
}

// Command represents a command in the module, which has a name, description, fields, and a list of events that it emits when handled.
type Command struct {
	Name          StringOfVaryingCases   `yaml:"name"`
	Description   string                 `yaml:"description"`
	EventsEmitted []StringOfVaryingCases `yaml:"events_emitted"`
	Params        []Field                `yaml:"params"`
}

// Event represents an event in the module, which has a name, description, fields, and a kind (either "emitted" or "consumed") to indicate whether the event is emitted by the module or consumed by the module.
type Event struct {
	Name        StringOfVaryingCases `yaml:"name"`
	Description string               `yaml:"description"`
	Kind        string               `yaml:"kind"` // emitted or consumed
	Fields      []Field              `yaml:"fields"`
}

// Query represents a query in the module, which has a name
type Query struct {
	Name   StringOfVaryingCases `yaml:"name"`
	Params []Field              `yaml:"params"`
}

// Field represents a field in an entity, command, or event, which has a name, type, and a boolean indicating whether the field is optional. It also includes helper methods to get the Go type, PostgreSQL type, and OpenAPI type for the field based on its type and whether it is optional.
type Field struct {
	Name     StringOfVaryingCases `yaml:"name"`
	Type     string               `yaml:"type"`
	Optional bool                 `yaml:"optional"`
}

// GoType returns the Go type for the field based on its type and whether it is optional.
// It supports basic types like string, int32, int64, float32, float64, bool, and time, and returns a pointer to the type if the field is optional. For any other types, it returns the type as-is.
func (f Field) GoType() string {
	switch f.Type {
	case "string":
		if f.Optional {
			return "*string"
		}
		return "string"
	case "int":
		if f.Optional {
			return "*int"
		}
		return "int"
	case "int8":
		if f.Optional {
			return "*int8"
		}
		return "int8"
	case "int16":
		if f.Optional {
			return "*int16"
		}
		return "int16"
	case "int32":
		if f.Optional {
			return "*int32"
		}
		return "int32"
	case "int64":
		if f.Optional {
			return "*int64"
		}
		return "int64"
	case "uint":
		if f.Optional {
			return "*uint"
		}
		return "uint"
	case "uint8":
		if f.Optional {
			return "*uint8"
		}
		return "uint8"
	case "uint16":
		if f.Optional {
			return "*uint16"
		}
		return "uint16"
	case "uint32":
		if f.Optional {
			return "*uint32"
		}
		return "uint32"
	case "uint64":
		if f.Optional {
			return "*uint64"
		}
		return "uint64"
	case "float32":
		if f.Optional {
			return "*float32"
		}
		return "float32"
	case "float64":
		if f.Optional {
			return "*float64"
		}
		return "float64"
	case "bool":
		if f.Optional {
			return "*bool"
		}
		return "bool"
	case "time":
		if f.Optional {
			return "*time.Time"
		}
		return "time.Time"
	default:
		return f.Type
	}
}

// PgSqlType returns the PostgreSQL type for the field based on its type and whether it is optional.
func (f Field) PgSqlType() string {
	switch f.Type {
	case "string":
		if f.Optional {
			return "TEXT"
		}
		return "TEXT NOT NULL"
	case "int32":
		if f.Optional {
			return "SMALLINT"
		}
		return "INT NOT NULL"
	case "int64":
		if f.Optional {
			return "BIGINT"
		}
		return "BIGINT NOT NULL"
	case "float32":
		if f.Optional {
			return "REAL"
		}
		return "REAL NOT NULL"
	case "float64":
		if f.Optional {
			return "DOUBLE PRECISION"
		}
		return "DOUBLE PRECISION NOT NULL"
	case "bool":
		if f.Optional {
			return "BOOLEAN"
		}
		return "BOOLEAN NOT NULL"
	case "time":
		if f.Optional {
			return "TIMESTAMP"
		}
		return "TIMESTAMP NOT NULL"
	default:
		return f.Type
	}
}

// OpenApiType returns the OpenAPI type for the field based on its type and whether it is optional.
func (f Field) OpenApiType() string {
	switch f.Type {
	case "string":
		return "string"
	case "int32":
		return "integer\nformat: int32"
	case "int64":
		return "integer\nformat: int64"
	case "float32":
		return "number\nformat: float"
	case "float64":
		return "number\nformat: double"
	case "bool":
		return "boolean"
	case "time":
		return "string\nformat: date-time"
	default:
		return f.Type
	}
}
