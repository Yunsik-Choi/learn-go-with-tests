package iterator

func Repeat(s string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result += s
	}
	return result
}

func Repeat2(text string, repeatNum int) interface{} {
	result := ""
	for i := 0; i < repeatNum; i++ {
		result += text
	}
	return result
}
