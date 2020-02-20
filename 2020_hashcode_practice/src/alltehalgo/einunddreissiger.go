package alltehalgo

func HelpSlices(list []int) []int {
	remainingSlices := list[0]
	numberOfPizzas := list[1]
	pizzas := list[2:]

	result := []int{}
	new_result := result
	best_result := []int{}

	for i := numberOfPizzas - 1; i >= 0; i-- {
		if pizzas[i] < remainingSlices {
			result = append(result, i)
			remainingSlices -= pizzas[i]
		}

		if new_result > best_result {
			best_result = new_result
		} else {
			HelpSlices(best_result)
		}
	}

	return best_result
}

// Einen Wert rausschmeißen und dann wieder kontrollieren, dass mit allen Werten