package main

import (
	"fmt"
	"math"
)

func main() {
	total := 0

	for i := 111122; i <= 111122; i++ {
		// p1
		// hasDouble := false
		// for p := 6; p >= 1; p-- {
		// 	if p == 6 {
		// 		continue
		// 	}

		// 	prev := digit(i, p+1)
		// 	cur := digit(i, p)
		// 	if cur == prev {
		// 		hasDouble = true
		// 	}
		// 	if prev > cur {
		// 		break
		// 	}
		// 	if p == 1 && hasDouble {
		// 		total++
		// 	}
		// }

		// p2
		hasDouble := false
		matched := make(map[int]bool)
		for p := 6; p >= 1; p-- {
			if p == 6 {
				continue
			}

			prev := digit(i, p+1)
			cur := digit(i, p)

			if matched[cur] {
				hasDouble = false
			} else if cur == prev {
				hasDouble = true
				matched[cur] = true
			}
			if prev > cur {
				break
			}
			if p == 1 && hasDouble {
				total++
			}
		}
	}

	fmt.Println(total)
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
