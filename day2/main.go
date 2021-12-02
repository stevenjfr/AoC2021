package main

import (
	"bufio"
	"fmt"
	"log"
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
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	m := map[string]int{"forward": 0, "up": 0, "down": 0}

	for _, line := range lines {
		entries := strings.Split(line, " ")

		num, _ := strconv.Atoi(entries[1])
		m[entries[0]] += num

	}

	var down = m["down"] - m["up"]

	fmt.Println("depth ", down)
	fmt.Println("forward ", m["forward"])

	fmt.Println("total ", (m["forward"] * down))

}

func code2() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var forward = 0
	var aim = 0
	var deep = 0

	for _, line := range lines {
		entries := strings.Split(line, " ")

		num, _ := strconv.Atoi(entries[1])

		if entries[0] == "forward" {
			forward += num
			deep = (deep + (num * aim))
		} else if entries[0] == "up" {
			aim -= num
		} else if entries[0] == "down" {
			aim += num
		}

	}
	fmt.Println("total ", forward*deep)

}
func main() {
	//code1()
	code2()
}
