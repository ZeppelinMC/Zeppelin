package item

type AttributeModifiers struct {
	Modifiers     []Modifier `nbt:"modifier"`
	ShowInTooltip bool       `nbt:"show_in_tooltip"`
}

type ModifierOperation string

const (
	AddValue           ModifierOperation = "add_value"
	AddMultipliedBase  ModifierOperation = "add_multiplied_base"
	AddMultipliedTotal ModifierOperation = "add_multiplied_total"
)

type Modifier struct {
	Type     string            `nbt:"type"`
	Slot     string            `nbt:"slot"`
	Id       string            `nbt:"id"`
	Amount   float64           `nbt:"amount"`
	Operator ModifierOperation `nbt:"operation"`
}
