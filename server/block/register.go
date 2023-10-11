package block

import (
	_ "embed"
	"encoding/json"
)

//go:embed blocks.json
var jsonBytes []byte

type jsonBlock struct {
	Properties map[string][]string `json:"properties"`

	States []struct {
		Id         int  `json:"id"`
		Default    bool `json:"default"`
		Properties map[string]string
	}
}

var jsonBlocks map[string]jsonBlock

func init() {
	if err := json.Unmarshal(jsonBytes, &jsonBlocks); err != nil {
		panic(err)
	}
}
