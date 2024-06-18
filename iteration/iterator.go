package iteration

func Repeat(character string, count int) string {
	var output string
	if count > 0 {
		for i := 0; i < count; i++ {
			output += character
		}
	}
	return output
}
