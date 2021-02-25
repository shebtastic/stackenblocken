package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputMeta struct {
	simulationDuration int
	numOfIntersections int
	numOfStreets       int
	numOfCars          int
	bonusPoints        int
}

type outputMeta struct {
	numberIntersections int
}

type street struct {
	startIntersection int
	streetName        string
	endIntersection   int
	time              int
}

type car struct {
	numOfStreetsToTravel int
	streets              []string
}

type greenlight struct {
	streetName string
	duration   int
}

type schedule struct {
	id                      int
	numberOfIncomingStreets int
	greenlights             []greenlight
}

type parsedData struct {
	meta    inputMeta
	streets []street
	cars    []car
}

type outputData struct {
	meta     outputMeta
	schedule schedule
}

const dataInputFolder string = "./data/"
const dataOutputFolder string = "./out/"

func async(items []int) chan []int {
	r := make(chan []int)

	go func() {
		defer close(r)
		fmt.Printf("chunksize:%d\n", len(items))
		r <- items
	}()

	return r
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

	simulationDuration, _ := strconv.Atoi(header[0])
	numOfIntersections, _ := strconv.Atoi(header[1])
	numOfStreets, _ := strconv.Atoi(header[2])
	numOfCars, _ := strconv.Atoi(header[3])
	bonusPoints, _ := strconv.Atoi(header[4])

	streets := []street{}
	for s := 0; s < numOfStreets; s++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		startIntersection, _ := strconv.Atoi(line[0])
		endIntersection, _ := strconv.Atoi(line[1])
		streetName := line[2]
		time, _ := strconv.Atoi(line[3])

		streets = append(streets, street{
			startIntersection: startIntersection,
			endIntersection:   endIntersection,
			streetName:        streetName,
			time:              time,
		})
	}

	cars := []car{}
	for c := 0; c < numOfCars; c++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		numOfStreetsToTravel, _ := strconv.Atoi(line[0])
		streets := line[1:]

		cars = append(cars, car{
			numOfStreetsToTravel: numOfStreetsToTravel,
			streets:              streets,
		})

	}

	return parsedData{
		meta: inputMeta{
			simulationDuration: simulationDuration,
			numOfIntersections: numOfIntersections,
			numOfStreets:       numOfStreets,
			numOfCars:          numOfCars,
			bonusPoints:        bonusPoints,
		},
		streets: streets,
		cars:    cars,
	}
}

func writeFile(file os.DirEntry, outputData outputData) {
	output := file.Name() + ".out"
	_ = os.Remove(output)

	wfile, err := os.Open(output)
	if err != nil {
		wfile, err = os.Create(output)
	}

	writer := bufio.NewWriter(wfile)
	// writer.WriteString(strconv.Itoa(len(selectedLibraries)) + "\n")
	// for _, library := range selectedLibraries {
	// 	writer.WriteString(strconv.Itoa(library.Id) + " " + strconv.Itoa(len(library.Books)) + "\n")
	// 	for index, book := range library.Books {
	// 		writer.WriteString(strconv.Itoa(book.Id))
	// 		if index == len(library.Books) - 1 {
	// 			writer.WriteString("\n")
	// 		} else {
	// 			writer.WriteString(" ")
	// 		}
	// 	}
	// }
	writer.Flush()
}

func main() {

	files := getFiles()
	selectedFile := files[0]
	data := readFile(selectedFile)

	fmt.Printf("\n\nfinalResult:\n%#v\n", data)

	writeFile(selectedFile, outputData{
		meta: outputMeta{
			numberIntersections: 0,
		},
	})
	// var c []chan []int
	// data := []int{}
	// res := []int{}

	// rand.Seed(time.Now().UnixNano())
	// randLen := 50000
	// chunkSize := 333

	// for i := 0; i < randLen; i++ {
	// 	data = append(data, rand.Intn(300))
	// }

	// for i := 0; i < len(data); i += chunkSize {
	// 	endIndex := i + chunkSize
	// 	if endIndex > len(data) {
	// 		endIndex = len(data)
	// 	}
	// 	c = append(c, async(data[i:endIndex]))
	// }
	// for _, v := range c {
	// 	res = append(res, <-v...)
	// }
	// fmt.Printf("total:\n%#v\n", len(res))
}
