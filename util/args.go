package util

import "os"

// checks if the named string is in os.Args
func HasArgument(name string) bool {
	for _, arg := range os.Args {
		if arg == name {
			return true
		}
	}
	return false
}
