package no2

// Contains 包含
func Contains(a, b string) bool {
	la, lb := len(a), len(b)
	if lb >= la {
		panic("lb >= la")
	}

	m := make(map[byte]struct{})
	for _, sa := range []byte(a) {
		m[sa] = struct{}{}
	}

	for _, sb := range []byte(b) {
		if _, ok := m[sb]; !ok {
			return false
		}
	}

	return true
}
