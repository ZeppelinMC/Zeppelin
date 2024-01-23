package block

type Stone struct{}

func (s Stone) EncodedName() string {
	return "minecraft:stone"
}

func (s Stone) New(m map[string]string) Block {
	return s
}

func (s Stone) Properties() map[string]string {
	return nil
}
