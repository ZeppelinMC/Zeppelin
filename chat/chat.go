package chat

const (
	Text         = "text"
	Translatable = "translatable"
	Keybind      = "keybind"
	Score        = "score"
	Selector     = "selector"
	NBT          = "nbt"
)

const (
	Default    = "minecraft:default"
	Uniform    = "minecraft:uniform"
	Alt        = "minecraft:alt"
	Illageralt = "minecraft:illageralt"
)

const (
	OpenURL         = "open_url"
	RunCommand      = "run_command"
	SuggestCommand  = "suggest_command"
	ChangePage      = "change_page"
	CopyToClipboard = "copy_to_clipboard"
)

const (
	ShowText   = "show_text"
	ShowItem   = "show_item"
	ShowEntity = "show_entity"
)

type ClickEvent struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

type HoverEventContents struct {
	ID    string `json:"id,omitempty"`
	Count int    `json:"count,omitempty"`
	Tag   string `json:"tag,omitempty"`
	Type  string `json:"type,omitempty"`
	Name  string `json:"name,omitempty"`
}

type HoverEvent struct {
	Action   string             `json:"action"`
	Contents HoverEventContents `json:"contents"`
}

type TextComponent struct {
	Type  string          `json:"type,omitempty"`
	Extra []TextComponent `json:"extra,omitempty"`

	Color         string     `json:"color,omitempty"`
	Bold          bool       `json:"bold,omitempty"`
	Italic        bool       `json:"italic,omitempty"`
	Underlined    bool       `json:"underlined,omitempty"`
	Strikethrough bool       `json:"strikethrough,omitempty"`
	Obfuscated    bool       `json:"obfuscated,omitempty"`
	Font          string     `json:"font,omitempty"`
	Insertion     string     `json:"insertion,omitempty"`
	ClickEvent    ClickEvent `json:"click_event,omitempty"`
	HoverEvent    HoverEvent `json:"hover_event,omitempty"`

	Text string `json:"text"`
}
