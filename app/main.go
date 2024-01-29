package main

import (
	"bytes"
	"os"

	"github.com/dynamitemc/dynamite/server/world/anvil"
)

func main() {
	buf := bytes.NewBuffer(nil)

	if err := anvil.CopyChunk(buf, 0, 0, "world/region/"); err != nil {
		panic(err)
	}

	os.WriteFile("chunk-0.0.nbt", buf.Bytes(), 0666)

}
