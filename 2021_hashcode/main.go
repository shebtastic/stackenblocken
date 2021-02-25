package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputMeta struct {
	SimulationDuration int
	NumOfIntersections int
	NumOfStreets       int
	NumOfCars          int
	BonusPoints        int
}

type outputMeta struct {
	NumberIntersections int
}

type street struct {
	StartIntersection int
	StreetName        string
	EndIntersection   int
	Time              int
}

type car struct {
	NumOfStreetsToTravel int
	Streets              []string
}

type greenlight struct {
	StreetName string
	Duration   int
}

type schedule struct {
	ID                      int
	NumberOfIncomingStreets int
	Greenlights             []greenlight
}

type parsedData struct {
	Meta    inputMeta
	Streets []street
	Cars    []car
}

type outputData struct {
	Meta     outputMeta
	Schedule []schedule
}

const dataInputFolder string = "./data/"
const dataOutputFolder string = "./output/"

func async(items []int) chan []int {
	r := make(chan []int)

	go func() {
		defer close(r)
		// fmt.Printf("chunksize:%d\n", len(items))
		r <- items
	}()

	return r
}

func getFiles() []os.DirEntry {
	files, _ := os.ReadDir(dataInputFolder)
	return files
}

func readFile(file string) parsedData {
	reader, _ := os.Open(file)
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
			StartIntersection: startIntersection,
			EndIntersection:   endIntersection,
			StreetName:        streetName,
			Time:              time,
		})
	}

	cars := []car{}
	for c := 0; c < numOfCars; c++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		numOfStreetsToTravel, _ := strconv.Atoi(line[0])
		streets := line[1:]

		cars = append(cars, car{
			NumOfStreetsToTravel: numOfStreetsToTravel,
			Streets:              streets,
		})

	}

	return parsedData{
		Meta: inputMeta{
			SimulationDuration: simulationDuration,
			NumOfIntersections: numOfIntersections,
			NumOfStreets:       numOfStreets,
			NumOfCars:          numOfCars,
			BonusPoints:        bonusPoints,
		},
		Streets: streets,
		Cars:    cars,
	}
}

func writeFile(file string, outputData outputData) {
	_ = os.Remove(file)

	wfile, err := os.Open(file)
	if err != nil {
		wfile, err = os.Create(file)
	}

	writer := bufio.NewWriter(wfile)
	writer.WriteString(strconv.Itoa(outputData.Meta.NumberIntersections) + "\n")
	for _, s := range outputData.Schedule {
		writer.WriteString(strconv.Itoa(s.ID) + "\n")
		writer.WriteString(strconv.Itoa(s.NumberOfIncomingStreets) + "\n")
		for _, g := range s.Greenlights {
			writer.WriteString(g.StreetName + " " + strconv.Itoa(g.Duration) + "\n")
		}
	}
	writer.Flush()
}

type intersection struct {
	ID      int
	Streets []street
	Cars    []car
}

func buildGraph(data parsedData) map[int]intersection {
	graph := make(map[int]intersection)

	for _, s := range data.Streets {
		id := s.EndIntersection
		if _, ok := graph[id]; !ok {
			graph[id] = intersection{
				ID:      id,
				Streets: []street{},
				Cars:    []car{},
			}
		}
		i := graph[id]
		graph[id] = intersection{
			ID:      i.ID,
			Streets: append(i.Streets, s),
			Cars:    i.Cars,
		}
	}

	for _, car := range data.Cars {
		startPosition := car.Streets[0]
		for index, i := range graph {
			for _, s := range i.Streets {
				if startPosition == s.StreetName {
					graph[index] = intersection{
						ID:      i.ID,
						Streets: i.Streets,
						Cars:    append(i.Cars, car),
					}
				}
			}
		}
	}

	return graph
}

func greenlightGraph(data parsedData) map[int]bool {
	graph := make(map[int]bool)
	for _, s := range data.Streets {
		id := s.EndIntersection
		graph[id] = false
	}

	return graph
}

func copyGraph(graph map[int]intersection) map[int]intersection {
	newGraph := map[int]intersection{}

	for _, i := range graph {
		streets := []street{}
		cars := []car{}
		for _, s := range i.Streets {
			streets = append(streets, s)
		}
		for _, c := range i.Cars {
			cs := []string{}
			for _, st := range c.Streets {
				cs = append(cs, st)
			}
			cars = append(cars, car{
				NumOfStreetsToTravel: c.NumOfStreetsToTravel,
				Streets:              cs,
			})
		}

		newGraph[i.ID] = intersection{
			ID:      i.ID,
			Streets: streets,
			Cars:    cars,
		}
	}

	return newGraph
}

func main() {

	files := getFiles()
	selectedFile := files[0]
	data := readFile(dataInputFolder + selectedFile.Name())

	graph := buildGraph(data)

	res := outputData{
		Schedule: []schedule{},
	}
	for tick := data.Meta.SimulationDuration; tick > 0; tick-- {
		j, _ := json.Marshal(res)
		fmt.Println(string(j))

		graphCopy := copyGraph(graph)
		for _, i := range graph {
			if len(i.Cars) > 0 {
				c := i.Cars[0]
				currentStreet := c.Streets[0]
				nextStreet := ""
				if len(c.Streets) > 2 {
					nextStreet = c.Streets[1]
				}
				alreadySet := false
				news := []schedule{}
				for sid, s := range res.Schedule {
					if s.ID == i.ID {
						for gid, g := range s.Greenlights {
							if g.StreetName == currentStreet {
								res.Schedule[sid].Greenlights[gid].Duration++
								alreadySet = true
							}
						}
						if !alreadySet {
							res.Schedule[sid] = schedule{
								ID: s.ID,
								Greenlights: append(s.Greenlights, greenlight{
									StreetName: currentStreet,
									Duration:   1,
								}),
							}
							alreadySet = true
						}
					}
					news = append(news, s)
				}

				if !alreadySet {
					news = append(res.Schedule, schedule{
						ID: i.ID,
						Greenlights: []greenlight{
							{
								StreetName: currentStreet,
								Duration:   1,
							},
						},
					})
				}
				res = outputData{
					Schedule: news,
				}

				graphCopy[i.ID] = intersection{
					ID:      i.ID,
					Streets: i.Streets,
					Cars:    i.Cars[1:],
				}

				if c.NumOfStreetsToTravel > 1 {
					c = car{
						NumOfStreetsToTravel: c.NumOfStreetsToTravel - 1,
						Streets:              c.Streets[1:],
					}
					for _, nextIntersection := range graph {
						for _, nextIntersectionStreet := range nextIntersection.Streets {
							if nextStreet == nextIntersectionStreet.StreetName {
								graphCopy[nextIntersection.ID] = intersection{
									ID:      nextIntersection.ID,
									Streets: nextIntersection.Streets,
									Cars:    append(nextIntersection.Cars, c),
								}
							}
						}
					}
				}
			}
		}
		graph = graphCopy
	}

	schedules := []schedule{}
	//schedule duplicate
	// - merge greenlights on intersections
	for _, s := range res.Schedule {
		greens := map[string]int{}
		for _, g := range s.Greenlights {
			if _, ok := greens[g.StreetName]; !ok {
				greens[g.StreetName] = 1
			} else {
				greens[g.StreetName]++
			}
		}

		sgreens := []greenlight{}
		for k, d := range greens {
			sgreens = append(sgreens, greenlight{
				StreetName: k,
				Duration:   d,
			})
		}
		schedules = append(schedules, schedule{
			ID:                      s.ID,
			NumberOfIncomingStreets: len(s.Greenlights),
			Greenlights:             sgreens,
		})
	}
	res = outputData{
		Meta: outputMeta{
			NumberIntersections: len(res.Schedule),
		},
		Schedule: schedules,
	}

	// j, _ := json.Marshal(res)
	// fmt.Println(string(j))

	fmt.Printf("\n\nfinalResult:\n%#v\n", res)
	writeFile(dataOutputFolder+selectedFile.Name()+".out", res)

	/*
		map[int]main.intersection{
			0: main.intersection{
				ID: 0,
				Streets: []main.street{
					main.street{
						StartIntersection: 2,
						StreetName:        "rue-de-londres",
						EndIntersection:   0,
						Time:              1}},
				Cars: []main.car{
					main.car{
						NumOfStreetsToTravel: 4,
						Streets: []string{
							"rue-de-londres",
							"rue-d-amsterdam",
							"rue-de-moscou",
							"rue-de-rome"}}}},
			1: main.intersection{
				ID: 1,
				Streets: []main.street{
					main.street{
						StartIntersection: 0,
						StreetName:        "rue-d-amsterdam",
						EndIntersection:   1,
						Time:              1},
					main.street{
						StartIntersection: 3,
						StreetName:        "rue-d-athenes",
						EndIntersection:   1,
						Time:              1}},
				Cars: []main.car{
					main.car{
						NumOfStreetsToTravel: 3,
						Streets: []string{
							"rue-d-athenes",
							"rue-de-moscou",
							"rue-de-londres"}}}},
			2: main.intersection{
				ID: 2,
				Streets: []main.street{
					main.street{
						StartIntersection: 1,
						StreetName:        "rue-de-moscou",
						EndIntersection:   2,
						Time:              3}},
				Cars: []main.car{}},
			3: main.intersection{
				ID: 3,
				Streets: []main.street{
					main.street{
						StartIntersection: 2,
						StreetName:        "rue-de-rome",
						EndIntersection:   3,
						Time:              2}},
				Cars: []main.car{}}}

		/*
			 outputData{
					Meta: outputMeta{
						NumberIntersections: 5,
					},
					Schedule: []schedule{
						{
							ID:                      5,
							NumberOfIncomingStreets: 3,
							Greenlights: []greenlight{
								{
									StreetName: "rue-de-eins",
									Duration:   10,
								},
								{
									StreetName: "rue-de-zwo",
									Duration:   6,
								},
								{
									StreetName: "rue-de-tres",
									Duration:   9,
								},
							},
						},
					},
				}
	*/
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
