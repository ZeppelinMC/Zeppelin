// Package list provides parsing of playerlist files (whitelist.json, ops.json etc)

package list

type WhitelistPlayer struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type OperatorPlayer struct {
	UUID                string `json:"uuid"`
	Name                string `json:"name"`
	Level               int32  `json:"level"`
	BypassesPlayerLimit bool   `json:"bypassesPlayerLimit"`
}
