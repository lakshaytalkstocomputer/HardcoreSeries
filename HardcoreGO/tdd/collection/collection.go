package collection

func ArraySum(numbers []int) int{
	var sum int
	for _ , data := range numbers{
		sum += data
	}
	return sum
}