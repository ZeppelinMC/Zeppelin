package commands

import (
	"github.com/aimjel/minecraft/protocol/types"
)

type stringArgType byte

const (
	SingleWord stringArgType = iota
	QuotablePhrase
	GreedyPhrase
)

const (
	EntitySingle = iota + 1
	EntityPlayerOnly
)

func (a Argument) Min(min uint64) Argument {
	//todo add checks, dont allow arguments that arent numbers to access this
	a.Parser.Properties.Flags |= 1
	a.Parser.Properties.Min = min
	return a
}

func (a Argument) Max(max uint64) Argument {
	//todo add checks, dont allow arguments that arent numbers to access this
	a.Parser.Properties.Flags |= 2
	a.Parser.Properties.Max = max
	return a
}

func (a Argument) MinMax(min, max uint64) Argument {
	a.Parser.Properties.Flags |= 0x03
	a.Parser.Properties.Min = min
	a.Parser.Properties.Max = max
	return a
}

func NewBoolArg(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 0,
		},
	}
}

func NewFloatArg(name string) Argument {
	props := types.CommandProperties{Flags: 0}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         1,
			Properties: props,
		},
	}
}

func NewIntArg(name string) Argument {
	props := types.CommandProperties{Flags: 0}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         3,
			Properties: props,
		},
	}
}

func (a Argument) SetSuggest(s func(ctx SuggestionsContext)) Argument {
	a.Suggest = s
	return a
}

func (a Argument) SetAlternative(arg Argument) Argument {
	a.Alternative = &arg
	return a
}

func NewStrArg(name string, properties stringArgType) Argument {
	props := types.CommandProperties{Flags: uint8(properties)}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         5,
			Properties: props,
		},
	}
}

func NewEntityArgument(name string, properties byte) Argument {
	props := types.CommandProperties{Flags: properties}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         6,
			Properties: props,
		},
	}
}

func NewGamemodeArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 39,
		},
	}
}

func NewChatComponentArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 17,
		},
	}
}

func NewDimensionArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 38,
		},
	}
}

func NewVector3Argument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 10,
		},
	}
}

func NewVector2Argument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 11,
		},
	}
}
