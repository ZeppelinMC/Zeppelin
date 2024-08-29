// qNBT is a very efficient NBT decoder!
// Caveats: does not work with []any, map[string]any or maps in general (except for map[string]string)
// Specification:
// Uses very lite, custom reflection (mostly unsafe)

package qnbt
