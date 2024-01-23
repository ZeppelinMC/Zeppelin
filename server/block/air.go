package block

type Air struct {
}

func (a Air) EncodedName() string {
	return "minecraft:air"
}

func (a Air) New(m map[string]string) Block {
	return nil
}

func (a Air) Properties() map[string]string {
	return nil
}
