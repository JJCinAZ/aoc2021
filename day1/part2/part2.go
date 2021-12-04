package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type measurement struct {
	depth  int
	window rune
}

func Run(filename string) {
	input := make(chan measurement, 0)
	go getinput(filename, input)
	prev, inc, cnt := 0, 0, 0
	for i := range input {
		fmt.Printf("%4d: %c %d\n", cnt, i.window, i.depth)
		if cnt > 0 {
			if i.depth > prev {
				inc++
			}
		}
		prev = i.depth
		cnt++
	}
	fmt.Printf("read %d depth measurements, %d were increased values\n", cnt, inc)
	fmt.Println("done")
}

func getinput(filename string, output chan measurement) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		linenum int
		m       [4]measurement
	)
	/*
		0  A
		1  A B
		2  A B C
		3    B C D
		4  A   C D
		5  A B   D
		6  A B C
		7    B C D
		8      C D
		9        D
	*/
	for i, c := 0, 'A'; i < 4; {
		m[i].window = c
		i++
		c++
	}
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i, err := strconv.Atoi(line); err == nil {
			b := (idx - 2) % 4
			switch {
			case idx == 0:
				m[0].depth = i
			case idx == 1:
				m[0].depth += i
				m[1].depth = i
			default:
				m[b].depth += i
				output <- m[b]
				m[b].depth = 0
				m[(b+1)%4].depth += i
				m[(b+2)%4].depth = i
			}
			linenum++
			idx++
		}
	}
	close(output)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
