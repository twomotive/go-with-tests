package iterations

func Repeat(c string, iternum int) string {

	var repeatedString string

	for i := 1; i <= iternum; i++ {
		repeatedString += c
	}

	return repeatedString
}
