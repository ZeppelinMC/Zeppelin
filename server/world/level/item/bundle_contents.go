package item

type BundleContent *struct {
	ID         string `nbt:"id"`
	Count      int32  `nbt:"count"`
	Components int32  `nbt:"components"` // optional map of data components
}
