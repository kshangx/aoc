package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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

	p2Arr := make([]int, len(arr))
	copy(p2Arr, arr)

	// p1
	// for i := 0; i < len(arr); i += 4 {
	// 	switch arr[i] {
	// 	case 1:
	// 		arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
	// 	case 2:
	// 		arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
	// 	case 99:
	// 		fmt.Println(arr[0])
	// 		return
	// 	default:
	// 		log.Panic("unknown opcode", arr[i])
	// 	}
	// }

	// p2
	for i1 := 0; i1 < 100; i1++ {
		for i2 := 0; i2 < 100; i2++ {
			copy(arr, p2Arr)
			arr[1] = i1
			arr[2] = i2
			for i := 0; i < len(arr); i += 4 {
				switch arr[i] {
				case 1:
					arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
				case 2:
					arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
				case 99:
					if arr[0] == 19690720 {
						fmt.Println(i1, i2)
						return
					}
				}
			}
		}
	}
}
