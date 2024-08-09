package item

// This is for passwords to chests.
// You would assign items ('keys') a password tag which would unlock a container item with the same lock key.
type Lock struct {
	Lock string `nbt:"lock"`
}
