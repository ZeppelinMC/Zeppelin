package block

type InfestedDeepslate struct {
	Axis string
}

func (b InfestedDeepslate) Encode() (string, BlockProperties) {
	return "minecraft:infested_deepslate", BlockProperties{
		"axis": b.Axis,
	}
}

func (b InfestedDeepslate) New(props BlockProperties) Block {
	return InfestedDeepslate{
		Axis: props["axis"],
	}
}