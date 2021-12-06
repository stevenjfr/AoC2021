package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	present bool
	age     int
}

type fishCount struct {
	age   int
	count int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func code2() {

	lines, _ := readLines("input.txt")

	fishes := []int{}
	// create a array representation of the text input
	for _, line := range strings.Split(lines[0], ",") {
		num, _ := strconv.Atoi(line)
		fishes = append(fishes, num)
	}

	fmt.Println(fishes)

	var days int
	for days < 80 {
		newfish := []int{}
		for i, fishie := range fishes {

			if fishie == 0 {
				fishes[i] = 6
				newfish = append(newfish, 8)
			}

			if fishie < 9 && fishie > 0 {
				fishes[i] = fishie - 1
			}
		}

		fishes = append(fishes, newfish...)

		//	fmt.Println(fishes)
		fmt.Println("day count: ", days)
		fmt.Println(len(fishes))
		days++
	}

}

func code3() {
	fish := [9]int{}

	lines, _ := readLines("input.txt")

	// create a array representation of the text input
	for _, line := range strings.Split(lines[0], ",") {
		num, _ := strconv.Atoi(line)

		fish[num] += +1

	}

	var total = makeFish(fish, 256)
	fmt.Println(total)

}

func makeFish(counts [9]int, days int) int {
	var total int
	for d := 1; d < days; d++ {

		pointer := (d + 7) % 9

		counts[pointer] += counts[d%9]

	}
	for _, ages := range counts {
		total += ages
	}
	return total
}

func main() {
	code3()
	//code2()

}
