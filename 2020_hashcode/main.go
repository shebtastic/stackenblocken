package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	Id int
	BookCount    int
	SignUpTime   int
	ShippingSize int
	Books        []Book
}

type Book struct {
	Id    int
	Score int
}

type ByBookScore []Book

func (s ByBookScore) Len() int {
	return len(s)
}
func (s ByBookScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByBookScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func readFile(r io.Reader) (int, int, int, []Book, []Library) {
	var (
		libraries = []Library{}
		books     = []Book{}

		libraryId = 0
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
			Id:    bookIndex,
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
			Id: libraryId,
			BookCount:    bookCount,
			SignUpTime:   signUpTime,
			ShippingSize: shippingSize,
			Books:        tmpBooks,
		})
		libraryId++
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

		fmt.Printf("%#v\n", fastestSignUpTime(libraries))
		fmt.Printf("%#v\n", highestBookScore(books))
	}
}

func doStuff() int {
	return 0
}

func fastestSignUpTime(LiD []Library) Library {
	temp_signUpTime := LiD[0]

	for _, Library := range LiD {
		if Library.SignUpTime < temp_signUpTime.SignUpTime {
			temp_signUpTime = Library
		}
	}
	return temp_signUpTime
}

func highestBookScore(bookScore []Book) Book {
	temp_highestBookScore := bookScore[0]

	sort.Sort(ByBookScore(bookScore))

	return temp_highestBookScore
}
