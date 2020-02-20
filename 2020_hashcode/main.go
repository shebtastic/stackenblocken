package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var files = []string{
	"a_example.txt",
	/*
	"c_incunabula.txt",
	"e_so_many_books.txt",
	"b_read_on.txt",
	"d_tough_choices.txt",
	"f_libraries_of_the_world.txt",
	*/
}

type Library struct {
	BookCount int
	SignUpTime int
	ShippingSize int
	Books []Book
}

type Book struct {
	Id int
	Score int
}

func readFile(r io.Reader) (int, int, int, []Book, []Library) {
	var (
		libraries = []Library{}
		books = []Book{}
	)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	
	_ = scanner.Scan()
	scannedLine := strings.Split(scanner.Text(), " ")

	numberOfBooks, _ := strconv.Atoi(scannedLine[0])
	numberOfLibraries, _ := strconv.Atoi(scannedLine[1])
	numberOfDays, _ := strconv.Atoi(scannedLine[2])

	scanner.Scan()
	scannedLine = strings.Split(scanner.Text(), " ")
	for bookIndex, bookScore := range scannedLine {
		score, _ := strconv.Atoi(bookScore)
		books = append(books, Book{
			Id: bookIndex,
			Score: score,
		})
	}

	for scanner.Scan() {
		scannedLine = strings.Split(scanner.Text(), " ")
		bookCount, _ := strconv.Atoi(scannedLine[0])
		signUpTime, _ := strconv.Atoi(scannedLine[1])
		shippingSize, _ := strconv.Atoi(scannedLine[2])

		scanner.Scan()
		scannedLine = strings.Split(scanner.Text(), " ")
		tmpBooks := []Book{}
		for _, bookIndexAsString := range scannedLine {
			bookIndex, _ := strconv.Atoi(bookIndexAsString)
			tmpBooks = append(tmpBooks, books[bookIndex])
		}

		libraries = append(libraries, Library{
			BookCount: bookCount,
			SignUpTime: signUpTime,
			ShippingSize: shippingSize,
			Books: tmpBooks,
		})
	}

	return numberOfBooks, numberOfLibraries, numberOfDays, books, libraries

}

func main() {
	fmt.Println("hashcode 2020")
	for _, file := range files {
		input := "input/" + file
		output := "output/" + file + ".out"

		rfile, err := os.Open(input)

		numberOfBooks, numberOfLibraries, numberOfDays, books, libraries := readFile(rfile)
		numberOfUsedLibraries := doStuff()

		fmt.Println(numberOfBooks, numberOfLibraries, numberOfDays, books, libraries, numberOfUsedLibraries)

		_ = os.Remove(output)

		wfile, err := os.Open(output)
		if err != nil {
			wfile, err = os.Create(output)
		}

		writer := bufio.NewWriter(wfile)
		writer.WriteString("\n")
		writer.Flush()
	}
}

func doStuff() int {
	return 0
}
