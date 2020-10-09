package collection

func ArraySum(numbers []int) int{
	var sum int
	for _ , data := range numbers{
		sum += data
	}
	return sum
}

func SumAll(slices ...[]int) []int{
	var result []int

	for _, s := range slices{
		result = append(result, ArraySum(s))
	}
	return result
}

func SumAllTail(slices ...[]int)[]int{
	var result []int

	for _,s := range slices{
		tail := s[1:]
		result = append(result, ArraySum(tail))
	}

	return result
}