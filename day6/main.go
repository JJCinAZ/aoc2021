package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type population struct {
	fish [9]uint // count of fish at each day stage (0...8)
}

func main() {
	simulate(os.Args[1], 80)
	simulate(os.Args[1], 256)
}

func simulate(filename string, days int) {
	pop := getinput(filename)
	for day := 1; day <= days; day++ {
		pop = processday(pop)
		fmt.Printf("There are %d fish after %d days\n", pop.count(), day)
	}
}

func (p *population) count() uint {
	var c uint
	for i := 0; i <= 8; i++ {
		c += p.fish[i]
	}
	return c
}

func processday(p population) population {
	var newp population
	// 0 <- 1, 1 <- 2, 2 <- 3, 3 <- 4, 4 <- 5, 5 <- 6, 6 <- 7, 7 <- 8
	for i := 0; i <= 7; i++ {
		newp.fish[i] = p.fish[i+1]
	}
	newp.fish[6] += p.fish[0] // cycle those at 0 days left
	newp.fish[8] = p.fish[0]  // spawn from those at 0 days left
	return newp
}

func getinput(filename string) population {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var pop population
	if scanner.Scan() {
		if a, err := stringsToUint8(strings.Split(scanner.Text(), ",")); err != nil {
			log.Fatal(err)
		} else {
			for _, x := range a {
				if x < 1 || x > 8 {
					log.Fatal("out of range value")
				}
				pop.fish[x]++
			}
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return pop
}

func stringsToUint8(s []string) ([]byte, error) {
	a := make([]byte, 0, len(s))
	for _, x := range s {
		if len(x) > 0 {
			if i, err := strconv.Atoi(x); err != nil {
				return nil, err
			} else {
				a = append(a, byte(i))
			}
		}
	}
	return a, nil
}
