package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := getinput(os.Args[1])
	sort.Ints(positions)

	fmt.Printf("PART1: ")
	findshortest(positions, func(x int) int {
		return x
	})

	fmt.Printf("PART2: ")
	findshortest(positions, func(x int) int {
		sum := 0
		for x > 0 {
			sum += x
			x--
		}
		return sum
	})
}

// Positions must be sorted smallest to largest
func findshortest(positions []int, dist func(int) int) {
	lastpos, cheapestPosition, cheapest := 0, 0, math.MaxInt64
	for _, p := range positions {
		if p == lastpos {
			continue
		}
		totaldist := 0
		for _, x := range positions {
			totaldist += dist(abs(x - p))
		}
		if totaldist < cheapest {
			cheapest = totaldist
			cheapestPosition = p
		}
	}
	fmt.Printf("Cheapest position is %d with a cost of %d\n", cheapestPosition, cheapest)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getinput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	a := make([]int, 0)
	if scanner.Scan() {
		if a, err = stringsToInts(strings.Split(scanner.Text(), ",")); err != nil {
			log.Fatal(err)
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return a
}

func stringsToInts(s []string) ([]int, error) {
	a := make([]int, 0, len(s))
	for _, x := range s {
		if len(x) > 0 {
			if i, err := strconv.Atoi(x); err != nil {
				return nil, err
			} else {
				a = append(a, i)
			}
		}
	}
	return a, nil
}
