package problems

// NC1 大数加法
func solve(s string, t string) string {
	if len(s) > len(t) {
		s, t = t, s
	}

	var res = []byte(t)
	var carry, sum uint8 = 0, 0
	sLen, tLen := len(s), len(t)

	for i := 1; i <= tLen; i++ {
		carry, sum = 0, t[tLen-i]-'0'+carry
		if i <= sLen {
			sum += s[sLen-i] - '0'
		}

		if sum > 9 {
			carry, sum = 1, sum%10
		}

		res[tLen-i] = sum + '0'

		if carry == 0 && i > sLen {
			break
		}
	}

	if carry == 1 {
		return `1` + string(res)
	}

	return string(res)
}
