package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi((scanner.Text()))
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}
func code1() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println("Starting script")
	var largerThan int = 0
	var pointer int = lines[0]
	for i, line := range lines {
		if line > pointer {
			fmt.Println(line, "increased", i)
			pointer = line
			largerThan++
		} else {
			fmt.Println(line, "decreased")
			pointer = line
		}
	}

	fmt.Println("number of increases: ", largerThan)

}
func code2() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println("Starting script")
	var largerThan int = 0
	var valueToCompare int = (lines[0] + lines[1] + lines[2])
	for i, line := range lines {
		if i > 1997 {
			fmt.Println("WE ARE AT THE END", line)
			fmt.Println("number of increases: ", largerThan)
			os.Exit(0)
		} else {
			var average = (lines[i] + lines[i+1] + lines[i+2])

			if average > valueToCompare {
				fmt.Println(i, average, ">", valueToCompare, " increased")
				valueToCompare = average
				largerThan++
			} else {
				fmt.Println(i, average, "< ", valueToCompare, " decreased or equal")
				valueToCompare = average
				//pointer = line
			}
		}
	}

}
func main() {
	code1()
	code2()
}
