package alltehalgo

func LaBoeuf(list []int) []int {
	remainingSlices := list[0]
	numberOfPizzas := list[1]
	pizzas := list[2:]

	result := []int{}

	for i := numberOfPizzas - 1; i >= 0; i-- {
		if pizzas[i] < remainingSlices {
			result = append(result, i)
			remainingSlices -= pizzas[i]
		}
	}

	return result
}
