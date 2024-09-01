package util

import (
	"fmt"
	"strconv"
	"strings"
)

var sizeUnits = map[string]uint64{
	"kb":  1000,                      // kilobyte (decimal)
	"kib": 1024,                      // kibibyte (binary)
	"mb":  1000 * 1000,               // megabyte (decimal)
	"mib": 1024 * 1024,               // mebibyte (binary)
	"gb":  1000 * 1000 * 1000,        // gigabyte (decimal)
	"gib": 1024 * 1024 * 1024,        // gibibyte (binary)
	"tb":  1000 * 1000 * 1000 * 1000, // terabyte (decimal)
	"tib": 1024 * 1024 * 1024 * 1024, // tebibyte (binary)
}

// ParseSizeUnit parses a size string into a number of bytes.
func ParseSizeUnit(sizeStr string) (uint64, error) {
	// Trim any surrounding whitespace and convert to lowercase.
	sizeStr = strings.TrimSpace(strings.ToLower(sizeStr))

	// Find the position where the letters start.
	for unit := range sizeUnits {
		if strings.HasSuffix(sizeStr, unit) {
			// Extract the numeric part of the string.
			numberStr := strings.TrimSuffix(sizeStr, unit)
			number, err := strconv.ParseUint(numberStr, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number format: %w", err)
			}

			// Convert the number to bytes using the corresponding unit.
			return number * sizeUnits[unit], nil
		}
	}

	number, err := strconv.ParseUint(sizeStr, 10, 64)
	return number, err
}

var sizeUnitsS = []struct {
	Unit   string
	Factor uint64
}{
	{"B", 1},
	{"KiB", 1024},
	{"MiB", 1024 * 1024},
	{"GiB", 1024 * 1024 * 1024},
	{"TiB", 1024 * 1024 * 1024 * 1024},
	{"PiB", 1024 * 1024 * 1024 * 1024 * 1024},
}

// FormatSizeUnit formats a number of bytes into a human-readable string.
func FormatSizeUnit(bytes uint64) string {
	if bytes == 0 {
		return "0 B"
	}

	var result strings.Builder
	for i := len(sizeUnitsS) - 1; i >= 0; i-- {
		unit := sizeUnitsS[i]
		if bytes >= unit.Factor {
			value := float64(bytes) / float64(unit.Factor)
			result.WriteString(fmt.Sprintf("%.2f %s", value, unit.Unit))
			return result.String()
		}
	}

	// This should never be reached since we handle all sizes
	return "0 B"
}
