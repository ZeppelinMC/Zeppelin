package block

import "strconv"

type GrassBlock struct {
	Snowy bool
}

func (g GrassBlock) EncodedName() string {
	return "minecraft:grass_block"
}

func (g GrassBlock) New(m map[string]string) Block {
	if s, ok := m["snowy"]; ok {
		g.Snowy, _ = strconv.ParseBool(s)
	}
	return g
}

func (g GrassBlock) Properties() map[string]string {
	return map[string]string{
		"snowy": strconv.FormatBool(g.Snowy),
	}
}
