package alltehalgo

import (
	"log"
)

func Better(x []int) []int {
	log.Printf("%#v", x)
	maxSlices := x[0]
	countTypes := x[1]
	pizzas := x[2:]
	log.Printf("maxSlices: %#v", maxSlices)
	log.Printf("count types: %#v", countTypes)

	sum := 0
	count := 0
	pizzasToOrder := make([]int, 0, len(pizzas))
	for i := len(pizzas)-1; i >= 0; i-- {
		if sum+pizzas[i] <= maxSlices {
			sum += pizzas[i]
			count++
			pizzasToOrder = append([]int{i}, pizzasToOrder...)
		}
	}

	log.Printf("maxSlices: %#v", sum )
	log.Printf("count: %#v", count)
	log.Printf("pizzasToOrder: %#v", pizzasToOrder)

	return pizzasToOrder
}
