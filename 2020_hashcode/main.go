package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var files = []string{
	"a_example.txt",
	"b_read_on.txt",
	"c_incunabula.txt",
	"d_tough_choices.txt",
	"e_so_many_books.txt",
	"f_libraries_of_the_world.txt",
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
type ByLibrarySignIn []Library

func (s ByBookScore) Len() int {
	return len(s)
}
func (s ByBookScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByBookScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s ByLibrarySignIn) Len() int {
	return len(s)
}
func (s ByLibrarySignIn) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLibrarySignIn) Less(i, j int) bool {
	return s[i].SignUpTime < s[j].SignUpTime
}

func readFile(r io.Reader) (int, int, int, []Book, []Library) {
	var (
		libraries = []Library{}
		books     = []Book{}

		libraryId = 0
	)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	numberOfBooks, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	numberOfLibraries, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	numberOfDays, _ := strconv.Atoi(scanner.Text())

	for bookIndex := 0; bookIndex < numberOfBooks; bookIndex++{
		_ = scanner.Scan()
		score, _ := strconv.Atoi(scanner.Text())
		books = append(books, Book{
			Id:    bookIndex,
			Score: score,
		})
	}

	for scanner.Scan() {
		bookCount, _ := strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		signUpTime, _ := strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		shippingSize, _ := strconv.Atoi(scanner.Text())

		tmpBooks := []Book{}
		for bookIndex := 0; bookIndex < bookCount; bookIndex++ {
			_ = scanner.Scan()
			bookIndex, _ := strconv.Atoi(scanner.Text())
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
	for _, file := range files {
		input := "input/" + file
		output := "output/" + file + ".out"

		rfile, err := os.Open(input)

		_, _, _, books, libraries := readFile(rfile)
		fmt.Printf("books %#v\nlibraries %#v\n\n", books, libraries)

		selectedLibraries := fastestSignUpTime(libraries)
		_ = os.Remove(output)

		wfile, err := os.Open(output)
		if err != nil {
			wfile, err = os.Create(output)
		}

		writer := bufio.NewWriter(wfile)
		writer.WriteString(strconv.Itoa(len(selectedLibraries)) + "\n")
		for _, library := range selectedLibraries {
			writer.WriteString(strconv.Itoa(library.Id) + " " + strconv.Itoa(len(library.Books)) + "\n")
			for index, book := range library.Books {
				writer.WriteString(strconv.Itoa(book.Id))
				if index == len(library.Books) - 1 {
					writer.WriteString("\n")
				} else {
					writer.WriteString(" ")
				}
			}
		}
		writer.Flush()

		//fmt.Printf("%#v\n", fastestSignUpTime(libraries))
		//fmt.Printf("%#v\n", highestBookScore(books))
	}
}

func fastestSignUpTime(libraries []Library) []Library {
	sort.Sort(ByLibrarySignIn(libraries))

	return libraries
}

func highestBookScore(books []Book) []Book {
	sort.Sort(ByBookScore(books))

	return books
}
