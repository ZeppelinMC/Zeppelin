package util

func MapEqual(m1, m2 map[string]string) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if len(m1) != len(m2) {
		return false
	}
	for key, value1 := range m1 {
		if m2[key] != value1 {
			return false
		}
	}

	return true
}
