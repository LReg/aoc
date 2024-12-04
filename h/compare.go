package h

// SafeCompareRuneSl I guess you can just change []rune to string instead of using this function
func SafeCompRuneSl(runes []rune, i int, s string) bool {
	return len(runes) >= i+len(s) && string(runes[i:i+len(s)]) == s
}
