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
		mode1 := digit(instruction, 3)
		mode2 := digit(instruction, 4)
		switch opcode {
		case 1, 2:
			val1 := getVal(arr, i+1, mode1)
			val2 := getVal(arr, i+2, mode2)

			switch opcode {
			case 1:
				arr[arr[i+3]] = val1 + val2
			case 2:
				arr[arr[i+3]] = val1 * val2
			}

			i += 4
		case 3:
			var input int
			fmt.Scanf("%d", &input)
			arr[arr[i+1]] = input

			i += 2
		case 4:
			fmt.Println(getVal(arr, i+1, mode1))

			i += 2
		case 5:
			val := getVal(arr, i+1, mode1)
			if val != 0 {
				i = getVal(arr, i+2, mode2)
			} else {
				i += 3
			}
		case 6:
			val := getVal(arr, i+1, mode1)
			if val == 0 {
				i = getVal(arr, i+2, mode2)
			} else {
				i += 3
			}
		case 7:
			val1 := getVal(arr, i+1, mode1)
			val2 := getVal(arr, i+2, mode2)
			if val1 < val2 {
				arr[arr[i+3]] = 1
			} else {
				arr[arr[i+3]] = 0
			}

			i += 4
		case 8:
			val1 := getVal(arr, i+1, mode1)
			val2 := getVal(arr, i+2, mode2)
			if val1 == val2 {
				arr[arr[i+3]] = 1
			} else {
				arr[arr[i+3]] = 0
			}

			i += 4
		case 99:
			return
		default:
			log.Panicln("unkown opcode", opcode, i)
		}
	}
}

func getVal(arr []int, idx, mode int) int {
	if mode == 0 {
		return arr[arr[idx]]
	}
	return arr[idx]
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
