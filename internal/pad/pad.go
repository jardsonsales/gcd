package pad

func Both(str string, padChar string, length int) string {
	if length <= len(str) {
		return str
	}
	miss := length - len(str)
	halfMiss := miss / 2
	res := ""
	for i := 0; i < halfMiss; i++ {
		res += padChar
	}
	res += str
	for i := 0; i < halfMiss; i++ {
		res += padChar
	}
	return res
}
