package status

import (
	"encoding/base64"
	"encoding/json"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
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
	data := make([]byte, len(dimgpnb64txt)+base64.StdEncoding.EncodedLen(len(f))+1)
	copy(data, []byte(dimgpnb64txt))
	base64.StdEncoding.Encode(data[len(dimgpnb64txt):], f)
	data[len(data)-1] = '"'

	return data, nil
}

const dimgpnb64txt = `"data:image/png;base64,`

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

func (s StatusResponse) Encode(w encoding.Writer) error {
	data, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	return w.ByteArray(data)
}

func (s *StatusResponse) Decode(r encoding.Reader) error {
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

func (StatusRequest) Encode(encoding.Writer) error {
	return nil
}

func (StatusRequest) Decode(r encoding.Reader) error {
	return nil
}
