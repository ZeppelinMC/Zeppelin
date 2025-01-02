package block

type BambooMosaic struct {
}

func (b BambooMosaic) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_mosaic", BlockProperties{}
}

func (b BambooMosaic) New(props BlockProperties) Block {
	return BambooMosaic{}
}