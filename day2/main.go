package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	direction string
	distance  int
}

func main() {
	part1()
	part2()
}

func part1() {
	input := make(chan step, 0)
	go getinput("input.txt", input)
	cnt := 0
	pos, depth := 0, 0
	for i := range input {
		fmt.Printf("%4d: %s %d\n", cnt, i.direction, i.distance)
		switch i.direction {
		case "forward":
			pos += i.distance
		case "down":
			depth += i.distance
		case "up":
			depth -= i.distance
		}
		cnt++
	}
	fmt.Printf("read %d steps, position=%d depth=%d x=%d\n", cnt, pos, depth, pos*depth)
	fmt.Println("done")
}

func part2() {
	input := make(chan step, 0)
	go getinput("input.txt", input)
	cnt := 0
	pos, depth, aim := 0, 0, 0
	for i := range input {
		fmt.Printf("%4d: %s %d\n", cnt, i.direction, i.distance)
		switch i.direction {
		case "forward":
			pos += i.distance
			depth += aim * i.distance
		case "down":
			aim += i.distance
		case "up":
			aim -= i.distance
		}
		cnt++
	}
	fmt.Printf("read %d steps, position=%d depth=%d x=%d\n", cnt, pos, depth, pos*depth)
	fmt.Println("done")
}

func getinput(filename string, output chan step) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		linenum int
	)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		d, _ := strconv.Atoi(parts[1])
		output <- step{direction: strings.ToLower(parts[0]), distance: d}
		linenum++
	}
	close(output)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
