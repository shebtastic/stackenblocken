package alltehalgo

import "log"

type relativePizza struct {
	Position int
	Slices   int
	Prev     int
	Next     int
}

func bakeRelativePizzas(list []int) []relativePizza {
	pizzas := []relativePizza{}

	pizzas = append(pizzas, relativePizza{
		Position: 0,
		Prev:     0,
		Slices:   list[0],
		Next:     list[0] - list[1],
	})

	for i := 1; i < len(list)-1; i++ {
		pizzas = append(pizzas, relativePizza{
			Position: i,
			Prev:     list[i] - list[i-1],
			Slices:   list[i],
			Next:     list[i] - list[i+1],
		})
	}

	pizzas = append(pizzas, relativePizza{
		Position: len(list) - 1,
		Prev:     list[len(list)-1] - list[len(list)-1-1],
		Slices:   list[len(list)-1],
		Next:     0,
	})

	return pizzas
}

func remainingPizzas(all []relativePizza, selected []relativePizza) []relativePizza {
	pizzas := all[:]

	for _, pizza := range selected {
		pizzas[pizza.Position] = relativePizza{
			Position: -1,
		}
	}

	remainingPizzas := []relativePizza{}
	for _, pizza := range pizzas {
		if pizza.Position != -1 {
			remainingPizzas = append(remainingPizzas, pizza)
		}
	}

	return remainingPizzas
}

func findClosestPizza(pizzas []relativePizza, slices int) relativePizza {
	closest := pizzas[len(pizzas)-1]
	distance := closest.Slices
	for _, pizza := range pizzas {
		newDistance := pizza.Slices - slices
		if newDistance < 0 {
			newDistance = -newDistance
		}
		if newDistance < distance {
			distance = newDistance
			closest = pizza
		}
	}

	return closest
}

func checkBudget(pizzas []relativePizza, slices int) bool {
	positiveBudget, negativeBudget := 0, 0

	for _, pizza := range pizzas {
		if pizza.Next > 0 {
			positiveBudget += pizza.Next
		} else {
			negativeBudget += pizza.Next
		}
	}

	log.Printf("positiveBudget: %d, negativeBudget: %d", positiveBudget, negativeBudget)
	if slices > 0 {
		return negativeBudget > slices
	}
	return positiveBudget > slices
}

func LaBoeuf(list []int) []int {
	result := []int{}
	remainingSlices := list[0]

	relativePizzas := bakeRelativePizzas(list[2:])

	selectedPizzas := []relativePizza{}

	for i := list[1] - 1; i >= 0; i-- {
		if relativePizzas[i].Slices <= remainingSlices {
			selectedPizzas = append(selectedPizzas, relativePizzas[i])
			remainingSlices -= relativePizzas[i].Slices
		}
	}

	for _, pizza := range selectedPizzas {
		result = append(result, pizza.Position)
	}

	remainingPizzas := remainingPizzas(relativePizzas, selectedPizzas)
	closestPizza := findClosestPizza(remainingPizzas, remainingSlices)

	log.Printf("remaining slices: %d, closest pizza: %#v, budget: %t",
		remainingSlices,
		closestPizza,
		checkBudget(
			remainingPizzas,
			closestPizza.Slices-remainingSlices,
		),
	)

	return result
}
