package block

import (
	"strconv"
)

type TorchflowerCrop struct {
	Age int
}

func (b TorchflowerCrop) Encode() (string, BlockProperties) {
	return "minecraft:torchflower_crop", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b TorchflowerCrop) New(props BlockProperties) Block {
	return TorchflowerCrop{
		Age: atoi(props["age"]),
	}
}