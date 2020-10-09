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
// Calculate sum of tail of each slice adn store it in result
//  We define tail as all the elements without the first element
func SumAllTail(slices ...[]int)[]int{
	var result []int

	// We need to check whether slice is empty or not
	//  if empty then append zero
	for _,s := range slices{

		if len(s) == 0{

			result = append(result, 0)

		} else{

			tail := s[1:]
			result = append(result, ArraySum(tail))

		}
	}

	return result
}