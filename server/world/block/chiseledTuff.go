package block

type ChiseledTuff struct {
}

func (b ChiseledTuff) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_tuff", BlockProperties{}
}

func (b ChiseledTuff) New(props BlockProperties) Block {
	return ChiseledTuff{}
}