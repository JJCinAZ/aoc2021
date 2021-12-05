package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type line struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	part1(os.Args[1])
	part2(os.Args[1])
}

func part1(filename string) {
	grid := makegrid(1000, 1000)
	for _, line := range getinput(filename) {
		if line.isHorizontal() {
			x1, x2 := orderpair(line.x1, line.x2)
			for x := x1; x <= x2; x++ {
				grid[line.y1][x]++
			}
		} else if line.isVertical() {
			y1, y2 := orderpair(line.y1, line.y2)
			for y := y1; y <= y2; y++ {
				grid[y][line.x1]++
			}
		}
	}
	fmt.Printf("part1: points with two or more overlapping lines: %d\n", pointcount(grid, 2))
}

func part2(filename string) {
	grid := makegrid(1000, 1000)
	for _, line := range getinput(filename) {
		if line.isVertical() {
			y1, y2 := orderpair(line.y1, line.y2)
			for y := y1; y <= y2; y++ {
				grid[y][line.x1]++
			}
		} else {
			y, yd := line.y1, ordertostep(line.y1, line.y2)
			if line.x1 <= line.x2 {
				for x := line.x1; x <= line.x2; x++ {
					grid[y][x]++
					y += yd
				}
			} else {
				for x := line.x1; x >= line.x2; x-- {
					grid[y][x]++
					y += yd
				}
			}
		}
	}
	fmt.Printf("part2: points with two or more overlapping lines: %d\n", pointcount(grid, 2))
}

func ordertostep(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return -1
}

func pointcount(grid [][]int, c int) int {
	total := 0
	for _, row := range grid {
		for _, p := range row {
			if p >= 2 {
				total++
			}
		}
	}
	return total
}

func makegrid(w, h int) [][]int {
	g := make([][]int, h)
	for i := range g {
		g[i] = make([]int, w)
	}
	return g
}

func orderpair(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func (l *line) isDiagonal() bool {
	if l.x1 != l.x2 && l.y1 != l.y2 {
		return true
	}
	return false
}

func (l *line) isHorizontal() bool {
	if l.y1 == l.y2 {
		return true
	}
	return false
}

func (l *line) isVertical() bool {
	if l.x1 == l.x2 {
		return true
	}
	return false
}

func getinput(filename string) []line {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	lines := make([]line, 0, 256)
	for scanner.Scan() {
		a := re.FindStringSubmatch(scanner.Text())
		if a == nil {
			log.Fatal("error on input")
		}
		lines = append(lines, parse2line(a))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return lines
}

func parse2line(a []string) line {
	var l line
	l.x1, _ = strconv.Atoi(a[1])
	l.y1, _ = strconv.Atoi(a[2])
	l.x2, _ = strconv.Atoi(a[3])
	l.y2, _ = strconv.Atoi(a[4])
	return l
}
