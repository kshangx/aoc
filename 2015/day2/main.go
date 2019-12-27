package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total1 := 0
	total2 := 0

	for scanner.Scan() {
		dimInts := []int{}

		dims := strings.Split(scanner.Text(), "x")
		for _, dim := range dims {
			d, err := strconv.Atoi(dim)
			if err != nil {
				log.Panic(err)
			}
			dimInts = append(dimInts, d)
		}

		sort.Ints(dimInts)

		// p1
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 3; j++ {
				total1 += dimInts[i] * dimInts[j] * 2
			}
		}

		total1 += dimInts[0] * dimInts[1]

		// p2
		total2 += dimInts[0]*2 + dimInts[1]*2 + dimInts[0]*dimInts[1]*dimInts[2]
	}

	fmt.Println(total1, total2)

}
