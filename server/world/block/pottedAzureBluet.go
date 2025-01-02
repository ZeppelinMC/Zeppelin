package block

type PottedAzureBluet struct {
}

func (b PottedAzureBluet) Encode() (string, BlockProperties) {
	return "minecraft:potted_azure_bluet", BlockProperties{}
}

func (b PottedAzureBluet) New(props BlockProperties) Block {
	return PottedAzureBluet{}
}