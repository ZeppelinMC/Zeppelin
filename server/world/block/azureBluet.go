package block

type AzureBluet struct {
}

func (b AzureBluet) Encode() (string, BlockProperties) {
	return "minecraft:azure_bluet", BlockProperties{}
}

func (b AzureBluet) New(props BlockProperties) Block {
	return AzureBluet{}
}