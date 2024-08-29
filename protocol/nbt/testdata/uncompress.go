package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"io"
	"os"
)

//go:embed bigTest.nbt
var bigTest []byte

func main() {
	gr, err := gzip.NewReader(bytes.NewReader(bigTest))
	if err != nil {
		panic(err)
	}

	out := make([]byte, 1024)

	n, err := gr.Read(out)
	if err != nil && err != io.EOF {
		panic(err)
	}

	err = os.WriteFile("bigTest.nbt", out[:n], 0666)
	if err != nil {
		panic(err)
	}
}
