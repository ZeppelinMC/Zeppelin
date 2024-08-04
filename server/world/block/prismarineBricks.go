package block

type PrismarineBricks struct {
}

func (b PrismarineBricks) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_bricks", BlockProperties{}
}

func (b PrismarineBricks) New(props BlockProperties) Block {
	return PrismarineBricks{}
}