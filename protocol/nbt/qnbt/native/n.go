package native

func Convert16(s []byte) {
	s[0], s[1] = s[1], s[0]
}

func Convert32(s []byte) {
	s[0], s[1], s[2], s[3] = s[3], s[2], s[1], s[0]
}

func Convert64(s []byte) {
	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7] = s[7], s[6], s[5], s[4], s[3], s[2], s[1], s[0]
}
