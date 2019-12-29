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

func (p *point) move(a point) point {
	p.x += a.x
	p.y += a.y
	return point{p.x, p.y}
}

func (p *point) equal(a point) bool {
	return p.x == a.x && p.y == a.y
}

func (p *point) isZero() bool {
	return p.x == 0 && p.y == 0
}

func (p *point) distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func main() {
	var w1 point
	var w2 point

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	w1Actions := loadActions(strings.Split(scanner.Text(), ","))
	scanner.Scan()
	w2Actions := loadActions(strings.Split(scanner.Text(), ","))

	var w1Path []point
	for _, a := range w1Actions {
		old := point{w1.x, w1.y}
		new := w1.move(a)
		w1Path = append(w1Path, getPath(old, new)...)
	}

	// p1
	// minDistance := -1
	// for _, a := range w2Actions {
	// 	old := point{w2.x, w2.y}
	// 	new := w2.move(a)
	// 	path := getPath(old, new)
	// 	for _, p1 := range path {
	// 		for _, p2 := range w1Path {
	// 			if p1.equal(p2) && !p1.isZero() {
	// 				distance := p1.distance()
	// 				if minDistance == -1 || minDistance > distance {
	// 					minDistance = distance
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// fmt.Println(minDistance)

	// p2
	var w2Path []point
	for _, a := range w2Actions {
		old := point{w2.x, w2.y}
		new := w2.move(a)
		w2Path = append(w2Path, getPath(old, new)...)
	}

	minSteps := -1
	for i1, p1 := range w1Path {
		for i2, p2 := range w2Path {
			if p1.equal(p2) && !p1.isZero() {
				steps := i1 + i2
				if minSteps == -1 || minSteps > steps {
					minSteps = steps
				}
			}
		}
	}

	fmt.Println(minSteps)
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

func getPath(old, new point) []point {
	var path []point

	if old.x == new.x {
		if old.y < new.y {
			for i := old.y; i < new.y; i++ {
				path = append(path, point{old.x, i})
			}
		} else {
			for i := old.y; i > new.y; i-- {
				path = append(path, point{old.x, i})
			}
		}
	} else {
		if old.x < new.x {
			for i := old.x; i < new.x; i++ {
				path = append(path, point{i, old.y})
			}
		} else {
			for i := old.x; i > new.x; i-- {
				path = append(path, point{i, old.y})
			}
		}
	}

	return path
}
