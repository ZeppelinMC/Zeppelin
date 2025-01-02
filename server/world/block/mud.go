package block

type Mud struct {
}

func (b Mud) Encode() (string, BlockProperties) {
	return "minecraft:mud", BlockProperties{}
}

func (b Mud) New(props BlockProperties) Block {
	return Mud{}
}