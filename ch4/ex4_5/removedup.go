package removedup

func RemoveDup(s []string) {
	if len(s) == 0 {
		return
	}
	index := 0

	for i := 1; i < len(s); i++ {
		if s[index] != s[i] {
			index++
			s[index] = s[i]
		}
	}
}
