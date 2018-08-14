package psi

// Decode treats s as text encoded according to EN 300 468 Annex A. It uses
// appropriate conversion according to selection byte
func DecodeText(s []byte) string {
	if len(s) == 0 {
		return ""
	}
	sel := s[0]
	if sel >= 0x20 {
		return textenc.DecodeISO6937(s)
	}
	s = s[1:]
	switch sel {
	case 1:
		return textenc.DecodeISO8859_5(s)
	case 2:
		return textenc.DecodeISO8859_6(s)
	case 3:
		return textenc.DecodeISO8859_7(s)
	case 4:
		return textenc.DecodeISO8859_8(s)
	case 5:
		return textenc.DecodeISO8859_9(s)
	case 0x10:
		if len(s) < 2 {
			break
		}
		n := int(uint16(s[0])<<8 | uint16(s[1]))
		if n > 0 && n != 12 && n < 16 {
			return textenc.DecodeISO8859(n, s[2:])
		}
	}
	// Assume UTF8
	return string(s)
}

func EncodeText(s string) []byte {
	// BUG: this assumes that s is always ISO6937 compatible.
	return []byte(s)
}
