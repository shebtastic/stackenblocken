package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c []chan []int
	data := []int{}
	res := []int{}

	rand.Seed(time.Now().UnixNano())
	randLen := 50000
	chunkSize := 333

	for i := 0; i < randLen; i++ {
		data = append(data, rand.Intn(300))
	}

	for i := 0; i < len(data); i += chunkSize {
		endIndex := i + chunkSize
		if endIndex > len(data) {
			endIndex = len(data)
		}
		c = append(c, async(data[i:endIndex]))
	}
	for _, v := range c {
		res = append(res, <-v...)
	}
	fmt.Printf("total:\n%#v\n", len(res))
}

func async(items []int) chan []int {
	r := make(chan []int)

	go func() {
		defer close(r)
		fmt.Printf("chunksize:%d\n", len(items))
		r <- items
	}()

	return r
}
