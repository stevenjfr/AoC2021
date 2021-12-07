package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func code1() {
	lines, _ := readLines("input.txt")

	numbers := strings.Split(lines[0], ",")

	var slice = make([]int, len(numbers))

	for i, _ := range slice {

		slice[i] = runStrategy(i, numbers)

	}

	smallest := slice[0]
	var pointer int
	for i, num := range slice[1:] {
		if num < smallest {
			smallest = num
			pointer = i + 1
		}
	}
	fmt.Println(smallest, pointer)
	fmt.Println(triangle(3))

}

func triangle(k int) int {
	var total int
	for i := 1; i <= k; i++ {
		total += i

	}
	return total
}

func runStrategy(i int, numbers []string) int {

	var total int
	for _, num := range numbers {
		numActual, _ := strconv.Atoi(num)

		//part 1
		//diff(i, numActual)

		//part2
		total += triangle(diff(i, numActual))
	}
	return total
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	code1()
	//code2()

}
