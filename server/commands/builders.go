package commands

const (
	StringSingleWord = iota
	StringQuotablePhrase
	StringGreedyPhrase
)

func NewCommand(name string, arguments ...Argument) *Command {
	return &Command{Name: name, Arguments: arguments}
}

func NewBoolArgument(name string) Argument {
	return Argument{
		Name:     name,
		ParserID: 0,
	}
}

func NewFloatArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name:       name,
		ParserID:   1,
		Properties: props,
	}
}

func NewDoubleArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name:       name,
		ParserID:   2,
		Properties: props,
	}
}

func NewIntegerArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name:       name,
		ParserID:   3,
		Properties: props,
	}
}

func NewLongArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name:       name,
		ParserID:   4,
		Properties: props,
	}
}

func NewStringArgument(name string, properties byte) Argument {
	props := Properties{Flags: properties}
	return Argument{
		Name:       name,
		ParserID:   5,
		Properties: props,
	}
}
