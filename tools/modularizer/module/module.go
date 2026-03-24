package module

import (
	"errors"
	"slices"
	"strings"

	"github.com/cooperlutz/go-full/tools/modularizer/utils"
)

// must be in snake_case during initialization
type StringOfVaryingCases string

func (s StringOfVaryingCases) validate() error {
	if !utils.IsSnakeCase(string(s)) {
		return errors.New("string must be in snake_case")
	}
	return nil
}

// Flat returns the string in flat case (e.g. "useraccount")
func (s StringOfVaryingCases) Flat() string {
	return utils.SnakeToFlat(string(s))
}

// Pascal returns the string in PascalCase (e.g. "UserAccount")
func (s StringOfVaryingCases) Pascal() string {
	return utils.SnakeToPascal(string(s))
}

// PascalCapitalizeID returns the string in PascalCase with "ID" capitalized (e.g. "UserID")
func (s StringOfVaryingCases) PascalCapitalizeID() string {
	pascal := s.Pascal()
	return strings.ReplaceAll(pascal, "Id", "ID")
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

// returns the string in title case (e.g. "Useraccount")
func (s StringOfVaryingCases) Title() string {
	return utils.TitleCase(string(s))
}

// returns the string in title case with spaces (e.g. "User Account")
func (s StringOfVaryingCases) TitleWithSpaces() string {
	spaced := strings.ReplaceAll(string(s), "_", " ")
	return utils.TitleCase(spaced)
}

// returns the plural form of string in title case with spaces (e.g. "User Accounts")
func (s StringOfVaryingCases) PluralTitleWithSpaces() string {
	return utils.Pluralize(s.TitleWithSpaces())
}

// PluralSnake returns the plural form of the string in snake_case (e.g. "user_accounts")
func (s StringOfVaryingCases) PluralSnake() string {
	return utils.Pluralize(string(s))
}

// PluralPascal returns the plural form of the string in PascalCase (e.g. "UserAccounts")
func (s StringOfVaryingCases) PluralPascal() string {
	return utils.Pluralize(s.Pascal())
}

// PluralKebab returns the plural form of the string in kebab-case (e.g. "user-accounts")
func (s StringOfVaryingCases) PluralKebab() string {
	return utils.Pluralize(s.Kebab())
}

// PluralCamel returns the plural form of the string in camelCase (e.g. "userAccounts")
func (s StringOfVaryingCases) PluralCamel() string {
	return utils.Pluralize(s.Camel())
}

// PluralFlat returns the plural form of the string in flat case (e.g. "useraccounts")
func (s StringOfVaryingCases) PluralFlat() string {
	return utils.Pluralize(s.Flat())
}

type ModuleConfig struct {
	Modules []Module `yaml:"modules"`
}

// Module encompasses all the information about a module, including its name, description, entities, commands, events, and queries. It also includes helper methods to get the aggregate root entity, non-aggregate entities, emitted events, and consumed events for the module.
type Module struct {
	Name           StringOfVaryingCases `yaml:"name"`
	Metadata       map[string]any       `yaml:"metadata"`
	Icon           string               `yaml:"icon"`
	Description    string               `yaml:"description"`
	DefaultQueries bool                 `yaml:"defaultQueries"`
	Aggregates     []Entity             `yaml:"aggregates"`
	Commands       []Command            `yaml:"commands"`
	Events         []Event              `yaml:"events"`
	Queries        []Query              `yaml:"queries"`
}

func (m Module) Validate() error {
	if err := m.Name.validate(); err != nil {
		return err
	}
	for _, entity := range m.Aggregates {
		if err := entity.validate(); err != nil {
			return err
		}
	}
	for _, command := range m.Commands {
		if err := command.validate(); err != nil {
			return err
		}
	}
	for _, event := range m.Events {
		if err := event.validate(); err != nil {
			return err
		}
	}
	for _, query := range m.Queries {
		if err := query.validate(); err != nil {
			return err
		}
	}
	return nil
}

func (m Module) NestedObjects() []Field {
	var nestedObjects []Field
	for _, entity := range m.Aggregates {
		for _, field := range entity.Fields {
			if field.Type == "object" {
				nestedObjects = append(nestedObjects, field)
			}
		}
	}
	return nestedObjects
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
	Name        StringOfVaryingCases `yaml:"name"`
	Description string               `yaml:"description"`
	Icon        string               `yaml:"icon"`
	Fields      []Field              `yaml:"fields"`
}

func (e Entity) RequiredFields() []Field {
	var requiredFields []Field
	for _, field := range e.Fields {
		if !field.Optional {
			requiredFields = append(requiredFields, field)
		}
	}
	return requiredFields
}

func (e Entity) NestedObjects() []Field {
	var nestedObjects []Field
	for _, field := range e.Fields {
		if field.Type == "object" {
			nestedObjects = append(nestedObjects, field)
		}
	}
	return nestedObjects
}

func (e Entity) validate() error {
	if err := e.Name.validate(); err != nil {
		return err
	}
	for _, field := range e.Fields {
		if err := field.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Command represents a command in the module, which has a name, description, fields, and a list of events that it emits when handled.
type Command struct {
	Name          StringOfVaryingCases   `yaml:"name"`
	Description   string                 `yaml:"description"`
	EventsEmitted []StringOfVaryingCases `yaml:"events_emitted"`
	Params        []Field                `yaml:"params"`
}

func (c Command) validate() error {
	if err := c.Name.validate(); err != nil {
		return err
	}
	for _, param := range c.Params {
		if err := param.validate(); err != nil {
			return err
		}
	}
	for _, eventName := range c.EventsEmitted {
		if err := eventName.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Event represents an event in the module, which has a name, description, fields, and a kind (either "emitted" or "consumed") to indicate whether the event is emitted by the module or consumed by the module.
type Event struct {
	Name        StringOfVaryingCases `yaml:"name"`
	Description string               `yaml:"description"`
	Kind        string               `yaml:"kind"` // emitted or consumed
	Fields      []Field              `yaml:"fields"`
}

func (e Event) validate() error {
	if err := e.Name.validate(); err != nil {
		return err
	}
	if e.Kind != "emitted" && e.Kind != "consumed" {
		return errors.New("event kind must be either 'emitted' or 'consumed'")
	}
	for _, field := range e.Fields {
		if err := field.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Query represents a query in the module, which has a name
type Query struct {
	Name   StringOfVaryingCases `yaml:"name"`
	Params []Field              `yaml:"params"`
}

func (q Query) validate() error {
	if err := q.Name.validate(); err != nil {
		return err
	}
	for _, param := range q.Params {
		if err := param.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Field represents a field in an entity, command, or event, which has a name, type, and a boolean indicating whether the field is optional. It also includes helper methods to get the Go type, PostgreSQL type, and OpenAPI type for the field based on its type and whether it is optional.
type Field struct {
	Name        StringOfVaryingCases `yaml:"name"`
	Type        string               `yaml:"type"`
	Description string               `yaml:"description"`
	List        bool                 `yaml:"list"`
	Optional    bool                 `yaml:"optional"`
	Fields      []Field              `yaml:"fields,omitempty"` // for nested fields (e.g. in a struct or a list of structs)
}

func (f Field) RequiredFields() []Field {
	var requiredFields []Field
	for _, field := range f.Fields {
		if !field.Optional {
			requiredFields = append(requiredFields, field)
		}
	}
	return requiredFields
}

func (f Field) validate() error {
	if err := f.Name.validate(); err != nil {
		return err
	}

	allowedTypes := []string{
		"string", // supported
		"int16",  // supported
		"int32",  // supported
		"int64",  // supported
		// "uint",
		// "uint16",
		// "uint32",
		// "uint64",
		"float32",   // supported
		"float64",   // supported
		"bool",      // supported
		"time",      // supported
		"timestamp", // supported
		"date",
		"uuid",   // supported
		"object", // supported
	}
	if !slices.Contains(allowedTypes, f.Type) {
		return errors.New("invalid field type: " + f.Type)
	}

	return nil
}

func (f Field) FieldsToGoStructFields() string {
	var structFields strings.Builder
	for i, field := range f.Fields {
		structFields.WriteString(field.Name.Pascal() + " " + field.GoType())
		if i < len(f.Fields)-1 {
			structFields.WriteString("; ")
		}
	}
	return structFields.String()
}

// GoType returns the Go type for the field based on its type and whether it is optional.
// It supports basic types like string, int32, int64, float32, float64, bool, and time, and returns a pointer to the type if the field is optional. For any other types, it returns the type as-is.
func (f Field) GoType() string {
	switch f.Type {
	case "string":
		if f.Optional && !f.List {
			return "*string"
		}
		if f.Optional && f.List {
			return "[]*string"
		}
		if !f.Optional && f.List {
			return "[]string"
		}
		return "string"
	case "int8":
		if f.Optional && !f.List {
			return "*int8"
		}
		if f.Optional && f.List {
			return "[]*int8"
		}
		if !f.Optional && f.List {
			return "[]int8"
		}
		return "int8"
	case "int16":
		if f.Optional && !f.List {
			return "*int16"
		}
		if f.Optional && f.List {
			return "[]*int16"
		}
		if !f.Optional && f.List {
			return "[]int16"
		}
		return "int16"
	case "int32":
		if f.Optional && !f.List {
			return "*int32"
		}
		if f.Optional && f.List {
			return "[]*int32"
		}
		if !f.Optional && f.List {
			return "[]int32"
		}
		return "int32"
	case "int64":
		if f.Optional && !f.List {
			return "*int64"
		}
		if f.Optional && f.List {
			return "[]*int64"
		}
		if !f.Optional && f.List {
			return "[]int64"
		}
		return "int64"
	case "uint16":
		if f.Optional && !f.List {
			return "*uint16"
		}
		if f.Optional && f.List {
			return "[]*uint16"
		}
		if !f.Optional && f.List {
			return "[]uint16"
		}
		return "uint16"
	case "uint32":
		if f.Optional && !f.List {
			return "*uint32"
		}
		if f.Optional && f.List {
			return "[]*uint32"
		}
		if !f.Optional && f.List {
			return "[]uint32"
		}
		return "uint32"
	case "uint64":
		if f.Optional && !f.List {
			return "*uint64"
		}
		if f.Optional && f.List {
			return "[]*uint64"
		}
		if !f.Optional && f.List {
			return "[]uint64"
		}
		return "uint64"
	case "float32":
		if f.Optional && !f.List {
			return "*float32"
		}
		if f.Optional && f.List {
			return "[]*float32"
		}
		if !f.Optional && f.List {
			return "[]float32"
		}
		return "float32"
	case "float64":
		if f.Optional && !f.List {
			return "*float64"
		}
		if f.Optional && f.List {
			return "[]*float64"
		}
		if !f.Optional && f.List {
			return "[]float64"
		}
		return "float64"
	case "bool":
		if f.Optional && !f.List {
			return "*bool"
		}
		if f.Optional && f.List {
			return "[]*bool"
		}
		if !f.Optional && f.List {
			return "[]bool"
		}
		return "bool"
	case "time":
		if f.Optional && !f.List {
			return "*time.Time"
		}
		if f.Optional && f.List {
			return "[]*time.Time"
		}
		if !f.Optional && f.List {
			return "[]time.Time"
		}
		return "time.Time"
	case "timestamp":
		if f.Optional && !f.List {
			return "*time.Time"
		}
		if f.Optional && f.List {
			return "[]*time.Time"
		}
		if !f.Optional && f.List {
			return "[]time.Time"
		}
		return "time.Time"
	case "date":
		if f.Optional && !f.List {
			return "*time.Time"
		}
		if f.Optional && f.List {
			return "[]*time.Time"
		}
		if !f.Optional && f.List {
			return "[]time.Time"
		}
		return "time.Time"
	case "uuid":
		if f.Optional && !f.List {
			return "*uuid.UUID"
		}
		if f.Optional && f.List {
			return "[]*uuid.UUID"
		}
		if !f.Optional && f.List {
			return "[]uuid.UUID"
		}
		return "uuid.UUID"
	case "object":
		if f.Optional && !f.List {
			return "*" + f.Name.Pascal()
		}
		if f.Optional && f.List {
			return "[]*" + f.Name.Pascal()
		}
		if !f.Optional && f.List {
			return "[]" + f.Name.Pascal()
		}
		return f.Name.Pascal()
	default:
		return f.Type
	}
}

func (f Field) MapFnFromPgToDomain() *string {
	switch f.Type {
	case "float32":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeFloat4ToFloat32Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "float64":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeFloat8ToFloat64Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "bool":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeBoolToBoolPtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "time":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeTimeNullToTimePtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "timestamp":
		if f.Optional && !f.List {
			result := "pgxutil.TimestampzToTimePtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.PgtypeTimestampzSliceToTimeSlicePtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.PgtypeTimestampzSliceToTimeSlice(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.PgtypeTimestampzToTime(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "string":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeTextToStringPtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "uuid":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeUUIDToUUIDPtr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.PgtypeUUIDSliceToUUIDSlice(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.PgtypeUUIDToUUID(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "int16":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeInt2ToInt16Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "int32":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeInt4ToInt32Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "int64":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeInt8ToInt64Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "utilitee.SliceOfValuesToSliceOfPointers(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "uint16":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeInt4ToUint16Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.PgtypeInt4SliceToUint16SliceOfPtrs(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.PgtypeInt4SliceToUint16Slice(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.PgtypeInt4ToUint16(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "uint32":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeInt8ToUint32Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.PgtypeInt8SliceToUint32SliceOfPtrs(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.PgtypeInt8SliceToUint32Slice(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.PgtypeInt8ToUint32(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "uint64":
		if f.Optional && !f.List {
			result := "pgxutil.PgtypeNumericToUint64Ptr(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.PgtypeNumericSliceToUint64SliceOfPtrs(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.PgtypeNumericSliceToUint64Slice(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.PgtypeNumericToUint64(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	case "object":
		if f.Optional && !f.List {
			result := "map" + f.Name.Pascal() + "FromPgToDomain(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if f.Optional && f.List {
			result := "mapSliceOf" + f.Name.Pascal() + "FromPgToDomain(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && f.List {
			result := "mapSliceOf" + f.Name.Pascal() + "FromPgToDomain(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		if !f.Optional && !f.List {
			result := "map" + f.Name.Pascal() + "FromPgToDomain(e." + f.Name.PascalCapitalizeID() + "),"
			return &result
		}
		return nil
	default:
		return nil
	}
}

func (f Field) MapFnFromDomainToPg() *string {
	switch f.Type {
	case "float32":
		if f.Optional && !f.List {
			result := "pgxutil.Float32ToPgtypeFloat4"
			return &result
		}
	case "float64":
		if f.Optional && !f.List {
			result := "pgxutil.Float64ToPgtypeFloat8"
			return &result
		}
	case "bool":
		if f.Optional && !f.List {
			result := "pgxutil.BoolToPgtypeBool"
			return &result
		}
	case "time":
		if f.Optional && !f.List {
			result := "pgxutil.TimeToPgtypeTimeNull"
			return &result
		}
	case "timestamp":
		if f.Optional && !f.List {
			result := "pgxutil.TimePtrToPgtypeTimestampz"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.TimeSliceToPgtypeTimestampzSlice"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.TimeToPgtypeTimestampz"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.TimePtrSliceToPgtypeTimestampzSlice"
			return &result
		}
	case "string":
		if f.Optional && !f.List {
			result := "pgxutil.StringToPgtypeText"
			return &result
		}
	case "uuid":
		if f.Optional && !f.List {
			result := "pgxutil.UUIDPtrToPgtypeUUID"
			return &result
		}
		if f.Optional && f.List {
			result := "pgxutil.UUIDSliceOfPtrsToPgtypeUUIDSlice"
			return &result
		}
		if !f.Optional && f.List {
			result := "pgxutil.UUIDSliceToPgtypeUUIDSlice"
			return &result
		}
		if !f.Optional && !f.List {
			result := "pgxutil.UUIDToPgtypeUUID"
			return &result
		}
	case "int16":
		if f.Optional && !f.List {
			result := "pgxutil.Int16ToPgtypeInt2"
			return &result
		}
	case "int32":
		if f.Optional && !f.List {
			result := "pgxutil.Int32ToPgtypeInt4"
			return &result
		}
	case "int64":
		if f.Optional && !f.List {
			result := "pgxutil.Int64ToPgtypeInt8"
			return &result
		}
	case "uint16":
		if f.Optional && !f.List {
			result := "pgxutil.Uint16ToPgtypeInt4"
			return &result
		}
	case "uint32":
		if f.Optional && !f.List {
			result := "pgxutil.Uint32ToPgtypeInt8"
			return &result
		}
	case "uint64":
		if f.Optional && !f.List {
			result := "pgxutil.Uint64ToPgtypeNumeric"
			return &result
		}
	default:
		if f.Optional && f.List && f.Type != "object" {
			result := "pgxutil.SliceOfPtrsToPgtype"
			return &result
		}
		return nil
	}
	if f.Optional && f.List {
		result := "pgxutil.SliceOfPtrsToPgtype"
		return &result
	}

	return nil
}

// PgSqlType returns the PostgreSQL type for the field based on its type and whether it is optional.
func (f Field) PgSqlType() string {
	switch f.Type {
	case "string":
		if f.Optional && !f.List {
			return "TEXT"
		}
		if f.Optional && f.List {
			return "TEXT[]"
		}
		if !f.Optional && f.List {
			return "TEXT[] NOT NULL"
		}
		return "TEXT NOT NULL"
	case "int16":
		if f.Optional && !f.List {
			return "SMALLINT"
		}
		if f.Optional && f.List {
			return "SMALLINT[]"
		}
		if !f.Optional && f.List {
			return "SMALLINT[] NOT NULL"
		}
		return "SMALLINT NOT NULL"
	case "int32":
		if f.Optional && !f.List {
			return "INTEGER"
		}
		if f.Optional && f.List {
			return "INTEGER[]"
		}
		if !f.Optional && f.List {
			return "INTEGER[] NOT NULL"
		}
		return "INTEGER NOT NULL"
	case "int64":
		if f.Optional && !f.List {
			return "BIGINT"
		}
		if f.Optional && f.List {
			return "BIGINT[]"
		}
		if !f.Optional && f.List {
			return "BIGINT[] NOT NULL"
		}
		return "BIGINT NOT NULL"
	case "uint16":
		if f.Optional && !f.List {
			return "INTEGER"
		}
		if f.Optional && f.List {
			return "INTEGER[]"
		}
		if !f.Optional && f.List {
			return "INTEGER[] NOT NULL"
		}
		return "INTEGER NOT NULL"
	case "uint32":
		if f.Optional && !f.List {
			return "BIGINT"
		}
		if f.Optional && f.List {
			return "BIGINT[]"
		}
		if !f.Optional && f.List {
			return "BIGINT[] NOT NULL"
		}
		return "BIGINT NOT NULL"
	case "uint64":
		if f.Optional && !f.List {
			return "NUMERIC(20)"
		}
		if f.Optional && f.List {
			return "NUMERIC(20)[]"
		}
		if !f.Optional && f.List {
			return "NUMERIC(20)[] NOT NULL"
		}
		return "NUMERIC(20) NOT NULL"
	case "float32":
		if f.Optional && !f.List {
			return "REAL"
		}
		if f.Optional && f.List {
			return "REAL[]"
		}
		if !f.Optional && f.List {
			return "REAL[] NOT NULL"
		}
		return "REAL NOT NULL"
	case "float64":
		if f.Optional && !f.List {
			return "DOUBLE PRECISION"
		}
		if f.Optional && f.List {
			return "DOUBLE PRECISION[]"
		}
		if !f.Optional && f.List {
			return "DOUBLE PRECISION[] NOT NULL"
		}
		return "DOUBLE PRECISION NOT NULL"
	case "bool":
		if f.Optional && !f.List {
			return "BOOLEAN"
		}
		if f.Optional && f.List {
			return "BOOLEAN[]"
		}
		if !f.Optional && f.List {
			return "BOOLEAN[] NOT NULL"
		}
		return "BOOLEAN NOT NULL"
	case "timestamp":
		if f.Optional && !f.List {
			return "TIMESTAMP WITH TIME ZONE"
		}
		if f.Optional && f.List {
			return "TIMESTAMP WITH TIME ZONE[]"
		}
		if !f.Optional && f.List {
			return "TIMESTAMP WITH TIME ZONE[] NOT NULL"
		}
		return "TIMESTAMP WITH TIME ZONE NOT NULL"
	case "time":
		if f.Optional && !f.List {
			return "TIME WITH TIME ZONE"
		}
		if f.Optional && f.List {
			return "TIME WITH TIME ZONE[]"
		}
		if !f.Optional && f.List {
			return "TIME WITH TIME ZONE[] NOT NULL"
		}
		return "TIME WITH TIME ZONE NOT NULL"
	case "date":
		if f.Optional && !f.List {
			return "DATE"
		}
		if f.Optional && f.List {
			return "DATE[]"
		}
		if !f.Optional && f.List {
			return "DATE[] NOT NULL"
		}
		return "DATE NOT NULL"
	case "uuid":
		if f.Optional && !f.List {
			return "UUID"
		}
		if f.Optional && f.List {
			return "UUID[]"
		}
		if !f.Optional && f.List {
			return "UUID[] NOT NULL"
		}
		return "UUID NOT NULL"
	case "object":
		return "JSONB"
	default:
		return f.Type
	}
}

// OpenApiType returns the OpenAPI type for the field based on its type and whether it is optional.
func (f Field) OpenApiType() string {
	switch f.Type {
	case "string":
		if f.List {
			return "type: array\n          items:\n            type: string"
		}
		return "type: string"
	case "int16":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int16"
		}
		return "type: integer\n          format: int16"
	case "int32":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int32"
		}
		return "type: integer\n          format: int32"
	case "int64":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int64"
		}
		return "type: integer\n          format: int64"
	case "uint16":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int16"
		}
		return "type: integer\n          format: int16"
	case "uint32":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int32"
		}
		return "type: integer\n          format: int32"
	case "uint64":
		if f.List {
			return "type: array\n          items:\n            type: integer\n            format: int64"
		}
		return "type: integer\n          format: int64"
	case "float32":
		if f.List {
			return "type: array\n          items:\n            type: number\n            format: float"
		}
		return "type: number\n          format: float"
	case "float64":
		if f.List {
			return "type: array\n          items:\n            type: number\n            format: double"
		}
		return "type: number\n          format: double"
	case "bool":
		if f.List {
			return "type: array\n          items:\n            type: boolean"
		}
		return "type: boolean"
	case "time":
		if f.List {
			return "type: array\n          items:\n            type: string\n            format: date-time"
		}
		return "type: string\n          format: date-time"
	case "timestamp":
		if f.List {
			return "type: array\n          items:\n            type: string\n            format: date-time"
		}
		return "type: string\n          format: date-time"
	case "date":
		if f.List {
			return "type: array\n          items:\n            type: string\n            format: date"
		}
		return "type: string\n          format: date"
	case "uuid":
		if f.List {
			return "type: array\n          items:\n            type: string\n            format: uuid"
		}
		return "type: string\n          format: uuid"
	case "object":
		if f.List {
			return "type: array\n          items:\n            $ref: '#/components/schemas/" + f.Name.Pascal() + "'"
		}
		return "$ref: '#/components/schemas/" + f.Name.Pascal() + "'"
	default:
		return f.Type
	}
}
