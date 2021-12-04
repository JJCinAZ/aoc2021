package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type grid [5][5]int // [row][col]

type puzzle struct {
	Calls []int
	Grids []grid
}

func main() {
	part2(os.Args[1])
}

func part1(filename string) {
	p := readpuzzle(filename)
	for _, c := range p.Calls {
		p.mark(c)
		if sum, found := p.findwinner(); found {
			fmt.Printf("part1: found winner: %d\n", sum*c)
			break
		}
	}
}

func part2(filename string) {
	p := readpuzzle(filename)
	for _, c := range p.Calls {
		p.mark(c)
		for {
			if sum, found := p.findwinner(); found {
				fmt.Printf("part2: found winner: (%d, %d) %d\n", sum, c, sum*c)
			} else {
				break
			}
		}
	}
}

func (p *puzzle) findwinner() (int, bool) {
	for i := range p.Grids {
		if sum, win := p.Grids[i].winner(); win {
			fmt.Printf("Grid %d winner:\n%s\n", i+1, p.Grids[i].string())
			// Remove this winning grid from the list of grids
			p.Grids = append(p.Grids[:i], p.Grids[i+1:]...)
			return sum, true
		}
	}
	return 0, false
}

func (g *grid) string() string {
	a := make([]string, 0, 5)
	for row := 0; row < len(g); row++ {
		s := fmt.Sprintf("%3d %3d %3d %3d %3d",
			g[row][0], g[row][1], g[row][2], g[row][3], g[row][4])
		a = append(a, s)
	}
	return strings.Join(a, "\n")
}

func (g *grid) winner() (int, bool) {
	for row := 0; row < len(g); row++ {
		s := 0
		for col := 0; col < len(g[row]); col++ {
			s += g[row][col]
		}
		if s == -5 {
			return g.sumUnMarked(), true
		}
	}
	for col := 0; col < len(g[0]); col++ {
		s := 0
		for row := 0; row < len(g); row++ {
			s += g[row][col]
		}
		if s == -5 {
			return g.sumUnMarked(), true
		}
	}
	return 0, false
}

func (g *grid) sumUnMarked() int {
	s := 0
	for _, r := range g {
		for _, i := range r {
			if i != -1 {
				s += i
			}
		}
	}
	return s
}

func (p *puzzle) mark(n int) {
	for i := range p.Grids {
		p.Grids[i].mark(n)
	}
}

func (g *grid) mark(n int) {
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == n {
				g[row][col] = -1
			}
		}
	}
}

func readpuzzle(filename string) puzzle {
	var (
		p   puzzle
		f   *os.File
		s   *bufio.Scanner
		err error
	)
	if f, err = os.Open(filename); err != nil {
		log.Fatal(err)
	}
	s = bufio.NewScanner(f)
	p.Grids = make([]grid, 0)
	state, linecount, row, gridIdx := 0, 0, 0, 0
	curgrid := new(grid)
	for s.Scan() {
		line := s.Text()
		linecount++
		switch state {
		case 0:
			if p.Calls, err = stringsToInts(strings.Split(line, ",")); err != nil {
				log.Fatalf("expected comma separated numbers at line %d (%s)", linecount, err.Error())
			}
			state = 1
		case 1:
			if len(line) != 0 {
				log.Fatalf("expected blank line at line %d", linecount)
			}
			state = 2
			row = 0
		case 2:
			a, err := stringsToInts(strings.Split(line, " "))
			if len(a) != 5 || err != nil {
				log.Fatalf("expected 5 numbers at line %d (%s)", linecount, err.Error())
			}
			for col := 0; col < 5; col++ {
				curgrid[row][col] = a[col]
			}
			if row == 4 {
				p.Grids = append(p.Grids, *curgrid)
				curgrid = new(grid)
				gridIdx++
				state = 1
			} else {
				row++
			}
		}
	}
	if state != 1 {
		log.Fatalf("incomplete grid at line %d", linecount)
	}
	fmt.Printf("read %d grids and %d calls\n", len(p.Grids), len(p.Calls))
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	return p
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
