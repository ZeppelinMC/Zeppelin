package block

type Sand struct{}

func (s Sand) EncodedName() string {
	return "minecraft:sand"
}

func (s Sand) New(m map[string]string) Block {
	return s
}

func (s Sand) Properties() map[string]string {
	return nil
}
