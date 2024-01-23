package block

type Dirt struct{}

func (d Dirt) EncodedName() string {
	return "minecraft:dirt"
}

func (d Dirt) New(m map[string]string) Block {
	return d
}

func (d Dirt) Properties() map[string]string {
	return nil
}
