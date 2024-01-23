package command

type Command struct {
	name []string

	execute func()

	args []argument
}

type argument struct {
}
