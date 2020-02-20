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
	"c_incunabula.txt",
	"e_so_many_books.txt",
	"b_read_on.txt",
	"d_tough_choices.txt",
	"f_libraries_of_the_world.txt",
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
	fmt.Println("was called")
	var (
		libraries = []Library{}
		books = []Book{}
	)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	
	_ = scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")

	numberOfBooks, _ := strconv.Atoi(firstLine[0])
	numberOfLibraries, _ := strconv.Atoi(firstLine[1])
	numberOfDays, _ := strconv.Atoi(firstLine[2])

	fmt.Printf("%d, %d, %d\n", numberOfBooks, numberOfLibraries, numberOfDays)

	scanner.Scan()
	tmpBooks := strings.Split(scanner.Text(), " ")
	for bookIndex, bookScore := range tmpBooks {
		score, _ := strconv.Atoi(bookScore)
		books = append(books, Book{
			Id: bookIndex,
			Score: score,
		})
	}

	for scanner.Scan() {
		tmpLibrary := scanner.Text()
		scanner.Scan()
		fmt.Println(string(tmpLibrary))
		tmpLibraryBooks := scanner.Text()
		fmt.Println(string(tmpLibraryBooks))
	}

	return numberOfBooks, numberOfLibraries, numberOfDays, books, libraries

}

func main() {
	fmt.Println("hashcode 2020")
	for _, file := range files {
		input := "input/" + file
		output := "output/" + file + ".out"

		rfile, err := os.Open(input)

		numberOfBooks, numberOfLibraries, numberOfDays, books, _ := readFile(rfile)
		fmt.Printf("%#v", books)

		_ = os.Remove(output)

		wfile, err := os.Open(output)
		if err != nil {
			wfile, err = os.Create(output)
		}

		writer := bufio.NewWriter(wfile)
		writer.WriteString("result" + strconv.Itoa(numberOfBooks) + " " + strconv.Itoa(numberOfLibraries) + " "  + strconv.Itoa(numberOfDays) + "\n")
		writer.Flush()
	}
}
