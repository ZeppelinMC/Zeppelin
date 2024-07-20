package status

import (
	"encoding/base64"
	"encoding/json"

	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/text"
)

type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}
type StatusSample struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type StatusPlayers struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []StatusSample `json:"sample,omitempty"`
}

type Favicon []byte

func (f Favicon) MarshalJSON() ([]byte, error) {
	if len(f) == 0 {
		return []byte{'"', '"'}, nil
	}
	data := []byte(`"data:image/png;base64,`)
	l0 := len(data)
	l1 := base64.StdEncoding.EncodedLen(len(f))
	data = append(data, make([]byte, l1)...)
	base64.StdEncoding.Encode(data[l0:], f)
	data = append(data, '"')

	return data, nil
}

type StatusResponseData struct {
	Version            StatusVersion      `json:"version"`
	Players            StatusPlayers      `json:"players"`
	Description        text.TextComponent `json:"description"`
	Favicon            Favicon            `json:"favicon,omitempty"`
	EnforcesSecureChat bool               `json:"enforcesSecureChat"`
}

type StatusResponse struct {
	Data StatusResponseData
}

func (StatusResponse) ID() int32 {
	return 0x00
}

func (s StatusResponse) Encode(w io.Writer) error {
	data, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	return w.ByteArray(data)
}

func (s *StatusResponse) Decode(r io.Reader) error {
	var data []byte
	if err := r.ByteArray(&data); err != nil {
		return err
	}
	return json.Unmarshal(data, &s.Data)
}

type StatusRequest struct {
}

func (StatusRequest) ID() int32 {
	return 0x00
}

func (StatusRequest) Encode(io.Writer) error {
	return nil
}

func (StatusRequest) Decode(r io.Reader) error {
	return nil
}
