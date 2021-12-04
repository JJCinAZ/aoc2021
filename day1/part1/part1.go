package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run(filename string) {
	depths, err := getinput(filename)
	if err != nil {
		log.Fatal(err)
	}
	prev, inc := 0, 0
	for i, d := range depths {
		if i > 0 {
			if d > prev {
				inc++
			}
		}
		prev = d
	}
	fmt.Printf("read %d depth measurements, %d were increased values\n", len(depths), inc)
}

func getinput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	a := make([]int, 0, 256)
	scanner := bufio.NewScanner(f)
	linenum := 0
	for scanner.Scan() {
		line := scanner.Text()
		linenum++
		if i, err := strconv.Atoi(line); err != nil {
			return nil, fmt.Errorf("error in line #%d: %s", linenum, err.Error())
		} else {
			a = append(a, i)
		}
	}
	return a, scanner.Err()
}
