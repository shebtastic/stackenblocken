package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strconv"
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
	var (
		libraries = []Library{}
		books = []Book{}
		
		numberOfBooks = 0
		numberOfLibraries = 0
		numberOfDays = 0
	)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	
	numberOfBooks, _ = strconv.Atoi(scanner.Text())
	numberOfLibraries, _ = strconv.Atoi(scanner.Text())
	numberOfDays, _ = strconv.Atoi(scanner.Text())

	scanner.Split(bufio.ScanLines)
	tmpBooks := scanner.Text()
	fmt.Println(string(tmpBooks))

	for scanner.Scan() {
		tmpLibrary := scanner.Text()
		fmt.Println(string(tmpLibrary))
		tmpLibraryBooks := scanner.Text()
		fmt.Println(string(tmpLibraryBooks))
	}

	return numberOfBooks,
		numberOfLibraries,
		numberOfDays,
		books,
		libraries

}

func main() {
	fmt.Println("hashcode 2020")
	for _, file := range files {
		input := "input/" + file + ".in"
		output := "output/" + file + ".out"

		rfile, err := os.Open(input)
		if err != nil {
			fmt.Printf("couldn't open file! %#v\n", err)
		}

		_, _, _, _, _ = readFile(rfile)

		_ = os.Remove(output)

		wfile, err := os.Open(output)
		if err != nil {
			wfile, err = os.Create(output)
			if err != nil {
				fmt.Printf("failed to create file %#v\n", err)
			}
			fmt.Printf("created file %s\n", output)
		}

		writer := bufio.NewWriter(wfile)
		writer.WriteString("result")
		writer.Flush()
	}
}
