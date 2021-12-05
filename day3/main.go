package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part2(os.Args[1])
}

func part2(filename string) {
	data, bits := getinput(filename)
	for len(data) > 1 && bits > 0 {
		onecount := getonescount(data, bits)
		fmt.Printf("onecount=%d, len=%d\n", onecount, len(data))
		if onecount >= len(data)-onecount {
			data = keep(data, bits, 1)
		} else {
			data = keep(data, bits, 0)
		}
		bits -= 1
	}
	fmt.Println(bits)
	fmt.Println(data)
	oxy := data[0]

	data, bits = getinput(filename)
	for len(data) > 1 && bits > 0 {
		onecount := getonescount(data, bits)
		fmt.Printf("onecount=%d, len=%d\n", onecount, len(data))
		if len(data)-onecount <= onecount {
			data = keep(data, bits, 0)
		} else {
			data = keep(data, bits, 1)
		}
		bits -= 1
	}
	fmt.Println(bits)
	fmt.Println(data)
	co2 := data[0]

	fmt.Printf("oxy: %d  co2: %d   total=%d\n", oxy, co2, oxy*co2)
}

func keep(data []int64, bit int, value int64) []int64 {
	list := make([]int64, 0, len(data))
	testbit := int64(1) << (bit - 1)
	value = value << (bit - 1)
	for _, x := range data {
		if x&testbit == value {
			list = append(list, x)
		}
	}
	return list
}

func getonescount(data []int64, bit int) int {
	count := 0
	testbit := int64(1) << (bit - 1)
	for _, x := range data {
		if x&testbit > 0 {
			count++
		}
	}
	return count
}

func getinput(filename string) ([]int64, int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make([]int64, 0, 256)
	bitlength := 0
	for scanner.Scan() {
		line := scanner.Text()
		if bitlength == 0 {
			bitlength = len(line)
		}
		i, _ := strconv.ParseInt(line, 2, 64)
		lines = append(lines, i)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return lines, bitlength
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		onecounts []int
		linecount int
	)
	for scanner.Scan() {
		line := scanner.Text()
		if linecount == 0 {
			onecounts = make([]int, len(line))
		}
		for i, c := range line {
			if c == '1' {
				onecounts[i]++
			}
		}
		linecount++
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	gamma, epsilon := 0, 0
	for i := 0; i < len(onecounts); i++ {
		if onecounts[i] > linecount-onecounts[i] {
			gamma |= 1
		} else {
			epsilon |= 1
		}
		gamma <<= 1
		epsilon <<= 1
	}
	gamma >>= 1
	epsilon >>= 1
	fmt.Printf("read %d samples, gamma = %b, epsilon = %b, power consumption=%d\n",
		linecount, gamma, epsilon, gamma*epsilon)
}
