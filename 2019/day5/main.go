package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	defer f.Close()

	// load input
	arr := []int{}
	for scanner.Scan() {
		l := scanner.Text()

		strs := strings.Split(l, ",")

		for _, s := range strs {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Panic(err)
			}
			arr = append(arr, i)
		}
	}

	// p1
	i := 0
	for {
		instruction := arr[i]
		opcode := digit(instruction, 1) + digit(instruction, 2)*10
		switch opcode {
		case 1, 2:
			mode1 := digit(instruction, 3)
			mode2 := digit(instruction, 4)

			var o1, o2 int
			if mode1 == 0 {
				o1 = arr[arr[i+1]]
			} else {
				o1 = arr[i+1]
			}

			if mode2 == 0 {
				o2 = arr[arr[i+2]]
			} else {
				o2 = arr[i+2]
			}

			switch opcode {
			case 1:
				arr[arr[i+3]] = o1 + o2
			case 2:
				arr[arr[i+3]] = o1 * o2
			}

			i += 4
		case 3:
			var input int
			fmt.Scanf("%d", &input)
			arr[arr[i+1]] = input

			i += 2
		case 4:
			mode := digit(instruction, 3)

			if mode == 0 {
				fmt.Println(arr[arr[i+1]])
			} else {
				fmt.Println(arr[i+1])
			}

			i += 2
		case 99:
			return
		default:
			log.Panicln("unkown opcode", opcode, i)
		}
	}
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
