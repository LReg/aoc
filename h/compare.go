package h

func SafeCompRuneSl(runes []rune, i int, s string) bool {
	return len(runes) >= i+len(s) && string(runes[i:i+len(s)]) == s
}
