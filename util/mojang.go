package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchUUID returns the uuid and case-sensitive name of the username passed.
// The username passed can be case-insensitive.
// The first value in the slice returned is the UUID and the following one is the case-sensitive username.
func FetchUUID(username string) ([]string, error) {
	data, err := httpGet(fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username))
	if err != nil {
		return nil, err
	}

	if data["errorMessage"] != "" {
		return nil, fmt.Errorf("%v", data["errorMessage"])
	}

	return []string{data["id"], data[data["name"]]}, nil
}

func httpGet(s string) (map[string]string, error) {
	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
