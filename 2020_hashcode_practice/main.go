package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

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
		log.Println(number)
	}

	return numbers
}

func main() {
	rfile, err := os.Open("input/a_example.in")
	if err != nil {
		log.Fatalf("omgah! %#v\n", err)
	}

	log.Printf("%#v", readFile(rfile))

	wfile, err := os.Open("output/a.out")
	if err != nil {
		wfile, err = os.Create("output/a.out")
		if err != nil {
			log.Panicf("still ded! %#v\n", err)
		}
		log.Printf("nvm.\n")
	}

	writer := bufio.NewWriter(wfile)
	writer.WriteString("diz be output")
	writer.Flush()
}
