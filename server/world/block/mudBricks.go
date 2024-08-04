package block

type MudBricks struct {
}

func (b MudBricks) Encode() (string, BlockProperties) {
	return "minecraft:mud_bricks", BlockProperties{}
}

func (b MudBricks) New(props BlockProperties) Block {
	return MudBricks{}
}