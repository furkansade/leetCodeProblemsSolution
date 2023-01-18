func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sChar := strings.Split(s, "")
	tChar := strings.Split(t, "")

	sort.Strings(sChar)
	sort.Strings(tChar)

	for i := 0; i < len(s); i++ {
		if sChar[i] != tChar[i] {
			return false
		}
	}
	return true
}