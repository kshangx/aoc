package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var p1, p2 float64

	for scanner.Scan() {
		l := scanner.Text()
		i, err := strconv.Atoi(l)
		if err != nil {
			log.Panic(err)
		}
		f := math.Floor(float64(i)/3) - 2
		p1 += f

		for f >= 0 {
			p2 += f
			f = math.Floor(f/3) - 2
		}
	}

	fmt.Println(int(p1), int(p2))
}
