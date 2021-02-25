package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type meta struct {
	simulationDuration int
	numOfIntersections int
	numOfStreets       int
	numOfCars          int
	bonusPoints        int
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

type parsedData struct {
	meta    meta
	streets []street
	cars    []car
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
		meta: meta{
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

func main() {

	files := getFiles()
	data := readFile(files[0])

	fmt.Printf("\n\nfinalResult:\n%#v\n", data)
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
