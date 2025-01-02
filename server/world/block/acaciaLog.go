package block

type AcaciaLog struct {
	Axis string
}

func (b AcaciaLog) Encode() (string, BlockProperties) {
	return "minecraft:acacia_log", BlockProperties{
		"axis": b.Axis,
	}
}

func (b AcaciaLog) New(props BlockProperties) Block {
	return AcaciaLog{
		Axis: props["axis"],
	}
}