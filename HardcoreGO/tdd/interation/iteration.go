package iteration

func Repeat(character string, repititionTime int)string{
	var repeated string

	for i := 1; i<=repititionTime; i++ {
		repeated += character
	}
	return repeated
}

