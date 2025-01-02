package item

type BucketEntityData struct {
	NoAI             bool    `nbt:"NoAI"`
	Silent           bool    `nbt:"Silent"`
	NoGravity        bool    `nbt:"NoGravity"`
	Glowing          bool    `nbt:"Glowing"`
	Invulnerable     bool    `nbt:"Invulnerable"`
	Health           float32 `nbt:"Health"`
	Age              int32   `nbt:"Age"`
	Variant          int32   `nbt:"Variant"`
	HuntingCooldown  int64   `nbt:"HuntingCooldown"`
	BucketVariantTag int32   `nbt:"BucketVariantTag"`
}
