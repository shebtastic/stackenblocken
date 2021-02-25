package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

const dataInputFolder string = "./data/"
const dataOutputFolder string = "./out/"

type pizza struct {
	index       int
	ingredients []string
	delivered   bool
}

type parsedData struct {
	numTeams2       int
	numTeams3       int
	numTeams4       int
	availablePizzas []*pizza
}

type result struct {
	teams2Deliveries [][]pizza
	teams3Deliveries [][]pizza
	teams4Deliveries [][]pizza
}

func main() {
	res := result{
		teams2Deliveries: [][]pizza{},
		teams3Deliveries: [][]pizza{},
		teams4Deliveries: [][]pizza{},
	}

	files := getFiles()
	data := readFile(files[4])

	res.teams2Deliveries = append(res.teams2Deliveries, <-assign(data.availablePizzas, 2, data.numTeams2)...)
	res.teams3Deliveries = append(res.teams3Deliveries, <-assign(data.availablePizzas, 3, data.numTeams3)...)
	res.teams4Deliveries = append(res.teams4Deliveries, <-assign(data.availablePizzas, 4, data.numTeams4)...)

	fmt.Printf("\n\nfinalResult:\n%#v\n", res)
}

func assign(pizzas []*pizza, teamSize int, limit int) chan [][]pizza {
	r := make(chan [][]pizza)

	go func() {
		defer close(r)

		result := [][]pizza{}
		// fmt.Printf("teamsize: %d, limit: %d, len(result): %d\n", teamSize, limit, len(result))
		for len(result) < limit {
			order := []pizza{}
			for len(order) < teamSize {
				undelivered := filterDelivered(pizzas)
				if len(undelivered) == 0 {
					break
				}
				available := len(undelivered)
				selected := undelivered[rand.Intn(available)]
				if !(*selected).delivered {
					selected.delivered = true
					order = append(order, *selected)
				}
			}
			result = append(result, order)
		}
		// fmt.Printf("partialResult:%#v\n", result)

		r <- result
	}()

	return r
}

func filterDelivered(pizzas []*pizza) []*pizza {
	filtered := []*pizza{}
	for _, pizza := range pizzas {
		if pizza.delivered == false {
			filtered = append(filtered, pizza)
		}
	}
	return filtered
}

func getFiles() []os.DirEntry {
	files, _ := os.ReadDir(dataInputFolder)
	return files
}

func readFile(file os.DirEntry) parsedData {
	reader, _ := os.Open(dataInputFolder + file.Name())
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	header := strings.Split(scanner.Text(), " ")

	numPizzas, _ := strconv.Atoi(header[0])
	availablePizzas := make([]*pizza, numPizzas)

	numTeams2, _ := strconv.Atoi(header[1])
	numTeams3, _ := strconv.Atoi(header[2])
	numTeams4, _ := strconv.Atoi(header[3])
	index := 0

	for {
		done := !scanner.Scan()
		if done {
			break
		}

		ingredients := strings.Split(scanner.Text(), " ")[1:]
		sort.Strings(ingredients)

		availablePizzas[index] = &pizza{
			index:       index,
			ingredients: ingredients,
			delivered:   false,
		}

		index++
	}

	return parsedData{
		availablePizzas: availablePizzas,
		numTeams2:       numTeams2,
		numTeams3:       numTeams3,
		numTeams4:       numTeams4,
	}
}
