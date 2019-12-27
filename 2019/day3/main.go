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

type point struct {
	x int
	y int
}

func (p *point) move(a point) {
	p.x += a.x
	p.y += a.y
}

func (p *point) distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func main() {
	var w1 point
	var w2 point

	minDistance := -1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	w1Actions := loadActions(strings.Split(scanner.Text(), ","))
	scanner.Scan()
	w2Actions := loadActions(strings.Split(scanner.Text(), ","))

	var w1Path []point
	for _, a := range w1Actions {
		w1.move(a)
		w1Path = append(w1Path, point{w1.x, w1.y})
	}

	var w2Path []point
	for _, a := range w2Actions {
		w2.move(a)
		w2Path = append(w2Path, point{w2.x, w2.y})
		for _, p := range w1Path {
			if w2.x == p.x && w2.y == p.y {
				distance := w2.distance()
				if minDistance == -1 || minDistance > distance {
					minDistance = distance
				}
			}
		}
	}

	fmt.Println(w1Actions)
	fmt.Println(w2Actions)
	fmt.Println(w1Path)
	fmt.Println(w2Path)
	fmt.Println(minDistance)
}

func loadActions(pathStr []string) []point {
	var path []point

	for _, s := range pathStr {
		var p point
		d, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Panic(err)
		}
		switch s[0] {
		case 'R':
			p.x += d
		case 'L':
			p.x -= d
		case 'U':
			p.y += d
		case 'D':
			p.y -= d
		default:
			log.Panic(s[0])
		}
		path = append(path, p)
	}

	return path
}
