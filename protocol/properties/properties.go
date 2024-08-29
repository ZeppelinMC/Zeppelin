// Package properties provides encoding and decoding of .properties files

package properties

type formatter string

func (f formatter) Rune() rune {
	return rune(f[0])
}

type ServerProperties struct {
	AcceptTransfers       bool   `properties:"accept-transfers"`
	AllowFlight           bool   `properties:"allow-flight"`
	AllowNether           bool   `properties:"allow-nether"`
	BroadcastConsoleToOps bool   `properties:"broadcast-console-to-ops"`
	BroadcastRconToOps    bool   `properties:"broadcast-rcon-to-ops"`
	Difficulty            string `properties:"difficulty"`

	EnableCommandBlock bool      `properties:"enable-command-block"`
	EnableRCON         bool      `properties:"enable-rcon"`
	EnableStatus       bool      `properties:"enable-status"`
	EnableQuery        bool      `properties:"enable-query"`
	EnableChat         bool      `properties:"enable-chat"`
	EnableEncryption   bool      `properties:"enable-encryption"`
	ChatFormatter      formatter `properties:"chat-formatter"`

	SystemChatFormat string `properties:"system-chat-format"`

	EnforceSecureProfile           bool   `properties:"enforce-secure-profile"`
	EnforceWhitelist               bool   `properties:"enforce-whitelist"`
	EntityBroadcastRangePrecentage int    `properties:"entity-broadcast-range-precentage"`
	ForceGamemode                  bool   `properties:"force-gamemode"`
	FunctionPermissionLevel        int    `properties:"function-permission-level"`
	Gamemode                       string `properties:"gamemode"`
	GenerateStructures             bool   `properties:"generate-structures"`
	GeneratorSettings              string `properties:"generator-settings"`
	Hardcore                       bool   `properties:"hardcore"`
	HideOnlinePlayers              bool   `properties:"hide-online-players"`
	InitialDisabledPacks           string `properties:"initial-disabled-packs"`
	InitialEnabledPacks            string `properties:"initial-enabled-packs"`
	LevelName                      string `properties:"level-name"`
	LevelSeed                      string `properties:"level-seed"`
	LevelType                      string `properties:"level-type"`
	LogIPs                         bool   `properties:"log-ips"`
	MaxPlayers                     int    `properties:"max-players"`
	MOTD                           string `properties:"motd"`
	NetworkCompressionThreshold    int    `properties:"network-compression-threshold"`

	OnlineMode              bool   `properties:"online-mode"`
	OpPermissionLevel       int    `properties:"op-permission-level"`
	PlayerIdleTimeout       int    `properties:"player-idle-timeout"`
	PreventProxyConnections bool   `properties:"prevent-proxy-connections"`
	PvP                     bool   `properties:"pvp"`
	QueryPort               uint16 `properties:"query.port"`
	RCONPassword            string `properties:"rcon.password"`
	RCONPort                uint16 `properties:"rcon.port"`

	// NOTE: unlike java edition, this accepts deflate, lz4, none AND gzip.
	RegionFileCompression string `properties:"region-file-compression"`
	ResourcePack          string `properties:"resource-pack"`
	ResourcePackId        string `properties:"resource-pack-id"`
	ResourcePackSha1      string `properties:"resource-pack-sha1"`
	RequireResourcePack   string `properties:"require-resource-pack"`
	ServerIp              string `properties:"server-ip"`
	ServerPort            uint16 `properties:"server-port"`

	SimulationDistance int `properties:"simulation-distance"`

	SpawnAnimals    bool `properties:"spawn-animals"`
	SpawnMonsters   bool `properties:"spawn-monsters"`
	SpawnNPCs       bool `properties:"spawn-npcs"`
	SpawnProtection int  `properties:"spawn-protection"`

	ViewDistance int  `properties:"view-distance"`
	WhiteList    bool `properties:"white-list"`
}

var Default = ServerProperties{
	AllowNether:                    true,
	BroadcastConsoleToOps:          true,
	BroadcastRconToOps:             true,
	Difficulty:                     "easy",
	EnableStatus:                   true,
	EnableChat:                     true,
	ChatFormatter:                  "§",
	EntityBroadcastRangePrecentage: 100,
	FunctionPermissionLevel:        2,
	Gamemode:                       "survival",
	GenerateStructures:             true,
	InitialEnabledPacks:            "vanilla",
	LevelName:                      "world",
	LevelType:                      "minecraft:normal",
	LogIPs:                         true,
	MaxPlayers:                     20,
	MOTD:                           "A Minecraft Server",
	NetworkCompressionThreshold:    256,
	OnlineMode:                     true,
	OpPermissionLevel:              4,
	PvP:                            true,
	QueryPort:                      25565,
	RCONPort:                       25575,
	RegionFileCompression:          "deflate",
	ServerPort:                     25565,
	SimulationDistance:             10,
	SpawnAnimals:                   true,
	SpawnMonsters:                  true,
	SpawnNPCs:                      true,
	SpawnProtection:                16,
	ViewDistance:                   10,
}
