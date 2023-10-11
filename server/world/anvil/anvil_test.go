package anvil

import (
	"testing"
)

func TestReader_ReadChunk(t *testing.T) {
	rd := NewReader("testdata/")

	_, err := rd.ReadChunk(0, 0)
	if err != nil {
		t.Fatal(err)
	}
}
