package item

type Profile struct {
	Name       string       `nbt:"name"`
	ID         []int32      `nbt:"id"`
	Properties []Properties `nbt:"properties"`
}

type Properties struct {
	Name      string `nbt:"name"`
	Value     string `nbt:"value"`
	Signature string `nbt:"signature"` // optional
}
