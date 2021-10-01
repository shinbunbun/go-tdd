package iteration

func Repeat(str string, cnt int) string {
	res := ""
	for i := 0; i < cnt; i++ {
		res += str
	}
	return res
}
