package main

import (
	"alltehalgo"
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

var files = []string{
	"a_example",
	"b_small",
	"c_medium",
	"d_quite_big",
	"e_also_big",
}

func readFile(r io.Reader) []int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	numbers := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("ermahgerd! %#v\n", err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func main() {
	var (
		current = 4
		input   = "input/" + files[current] + ".in"
		output  = "output/" + files[current] + ".out"
	)
	rfile, err := os.Open(input)
	if err != nil {
		log.Fatalf("omgah! %#v\n", err)
	}

	result := alltehalgo.LaBoeuf(readFile(rfile))

	wfile, err := os.Open(output)
	if err != nil {
		wfile, err = os.Create(output)
		if err != nil {
			log.Panicf("still ded! %#v\n", err)
		}
		log.Printf("nvm.\n")
	}

	writer := bufio.NewWriter(wfile)
	writer.WriteString(strconv.Itoa(len(result)) + "\n")
	for index, entry := range result {
		writer.WriteString(strconv.Itoa(entry))
		if index != len(result)-1 {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
}
