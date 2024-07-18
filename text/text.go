package text

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
	Action string `json:"action" nbt:"action"`
	Value  string `json:"value" nbt:"value"`
}

type HoverEventContents struct {
	ID    string `json:"id,omitempty" nbt:"id,omitempty"`
	Count int    `json:"count,omitempty" nbt:"count,omitempty"`
	Tag   string `json:"tag,omitempty" nbt:"tag,omitempty"`
	Type  string `json:"type,omitempty" nbt:"type,omitempty"`
	Name  string `json:"name,omitempty" nbt:"name,omitempty"`
}

type HoverEvent struct {
	Action   string             `json:"action" nbt:"action"`
	Contents HoverEventContents `json:"contents" nbt:"contents"`
}

type TextComponent struct {
	Type  string          `json:"type,omitempty" nbt:"type,omitempty"`
	Extra []TextComponent `json:"extra,omitempty" nbt:"extra,omitempty"`

	Color         string     `json:"color,omitempty" nbt:"color,omitempty"`
	Bold          bool       `json:"bold,omitempty" nbt:"bold,omitempty"`
	Italic        bool       `json:"italic,omitempty" nbt:"italic,omitempty"`
	Underlined    bool       `json:"underlined,omitempty" nbt:"underlined,omitempty"`
	Strikethrough bool       `json:"strikethrough,omitempty" nbt:"strikethrough,omitempty"`
	Obfuscated    bool       `json:"obfuscated,omitempty" nbt:"obfuscated,omitempty"`
	Font          string     `json:"font,omitempty" nbt:"font,omitempty"`
	Insertion     string     `json:"insertion,omitempty" nbt:"insertion,omitempty"`
	ClickEvent    ClickEvent `json:"click_event,omitempty" nbt:"click_event,omitempty"`
	HoverEvent    HoverEvent `json:"hover_event,omitempty" nbt:"hover_event,omitempty"`

	Text string `json:"text" nbt:"text"`
}
