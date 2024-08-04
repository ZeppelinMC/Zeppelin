package block

type MossCarpet struct {
}

func (b MossCarpet) Encode() (string, BlockProperties) {
	return "minecraft:moss_carpet", BlockProperties{}
}

func (b MossCarpet) New(props BlockProperties) Block {
	return MossCarpet{}
}