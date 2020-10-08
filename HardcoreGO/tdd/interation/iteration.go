package iteration

func Repeat(character string)string{
	var repeated string

	for i := 1; i<=6; i++ {
		repeated = repeated + character
	}
	return repeated
}