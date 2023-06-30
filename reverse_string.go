package reverse_string

func ReverseString(str string) string {
	revStr := ""
	for _, i := range str {
		revStr = string(i) + revStr
	}
	return revStr
}
