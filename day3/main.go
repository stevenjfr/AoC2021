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

// func code1() {

// 	lines, err := readLines("input.txt")
// 	if err != nil {
// 		log.Fatalf("readLines: %s", err)
// 	}

// 	var slices [5][12]int
// 	var remaining [12][5]int

// 	for i, line := range lines {
// 		binary := strings.Split(line, "")

// 		for v, number := range binary {
// 			num, _ := strconv.Atoi(number)
// 			slices[v][i] = num
// 			remaining[i][v] = num
// 		}

// 	}

// 	var final []int
// 	for _, bitstream := range slices {

// 		var oneCount int
// 		var zeroCount int
// 		for _, bit := range bitstream {

// 			if bit > 0 {
// 				oneCount++
// 			} else {
// 				zeroCount++
// 			}
// 		}
// 		if oneCount > zeroCount {
// 			final = append(final, 1)
// 		} else {
// 			final = append(final, 0)
// 		}

// 	}

// 	var gamma = strings.Trim(strings.Join(strings.Split(fmt.Sprint(final), " "), ""), "[]")
// 	i, _ := strconv.ParseInt(gamma, 2, 64)
// 	fmt.Println(i)
// 	//var gammagammaDecimal = i

// 	var finalreverse []int
// 	for i := len(slices) - 1; i >= 0; i-- {
// 		//	println(i)
// 		bitstream := slices[i]
// 		//for _, bitstream := range slices {

// 		var oneCount int
// 		var zeroCount int
// 		for _, bit := range bitstream {

// 			if bit > 0 {
// 				oneCount++
// 			} else {
// 				zeroCount++
// 			}
// 		}
// 		if oneCount < zeroCount {
// 			finalreverse = prependInt(finalreverse, 1)
// 		} else {
// 			finalreverse = prependInt(finalreverse, 0)
// 		}

// 	}
// 	//fmt.Println(finalreverse)
// 	var epsilon = strings.Trim(strings.Join(strings.Split(fmt.Sprint(finalreverse), " "), ""), "[]")
// 	e, _ := strconv.ParseInt(epsilon, 2, 64)

// 	var total = i * e
// 	fmt.Println(total)

// }

// func prependInt(x []int, y int) []int {
// 	x = append(x, 0)
// 	copy(x[1:], x)
// 	x[0] = y
// 	return x
// }

func pruneRemaining(input [][]int, keep int, position int) [][]int {
	var remaining [][]int

	for _, v := range input {
		if v[position] == keep {
			remaining = append(remaining, v)
		} else {
			//	fmt.Println(v[position], "didnt have ", keep, " as pisition ", position)
		}
	}
	//	fmt.Println(remaining)
	return remaining

}

func calculateKeep(input [][]int, position int) int {
	var bits []int
	for _, line := range input {
		bits = append(bits, line[position])
	}
	var oneCount int
	var zeroCount int

	for _, bit := range bits {

		if bit > 0 {
			oneCount++
		} else {
			zeroCount++
		}
	}

	if oneCount > zeroCount {
		return 1

	} else if oneCount == zeroCount {
		return 1

	} else {
		return 0
	}
}

func calculateKeepLeast(input [][]int, position int) int {
	var bits []int
	for _, line := range input {
		bits = append(bits, line[position])
	}
	var oneCount int
	var zeroCount int

	for _, bit := range bits {

		if bit > 0 {
			zeroCount++
		} else {
			oneCount++
		}
	}

	if oneCount > zeroCount {
		return 1

	} else {
		return 0

	}
}

func makeArrayFromText(lines []string) [][]int {
	var remaining [][]int
	// create a array representation of the text input
	for _, line := range lines {
		binary := strings.Split(line, "")
		var newArray []int
		for _, number := range binary {
			num, _ := strconv.Atoi(number)
			newArray = append(newArray, num)

		}
		remaining = append(remaining, newArray)

	}
	return remaining

}
func code2() {

	lines, _ := readLines("input.txt")

	var remaining = makeArrayFromText(lines)

	//for every bit position
	for i := 0; i < len(remaining[0]); i++ {
		if len(remaining) > 1 {
			toKeep := calculateKeep(remaining, i)
			//	fmt.Printf("keep any values with %s at position %s", toKeep, i)
			//fmt.Println(toKeep)
			remaining = pruneRemaining(remaining, toKeep, i)
			//fmt.Println(remaining)
		} else {
			fmt.Println("loop ended")
			break
		}

	}
	var gamma = strings.Trim(strings.Join(strings.Split(fmt.Sprint(remaining), " "), ""), "[]")
	g, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println(g)

	var remainingLeast = makeArrayFromText(lines)

	//for every bit position
	for i := 0; i < len(remainingLeast[0]); i++ {
		if len(remainingLeast) > 1 {
			toKeep := calculateKeepLeast(remainingLeast, i)
			//fmt.Printf("keep any values with %s at position %s", toKeep, i)
			//fmt.Println(toKeep)
			remainingLeast = pruneRemaining(remainingLeast, toKeep, i)
			//fmt.Println(remainingLeast)
		} else {
			fmt.Println("loop ended")
			break
		}

	}
	var epsilon = strings.Trim(strings.Join(strings.Split(fmt.Sprint(remainingLeast), " "), ""), "[]")
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(e)

}
func main() {
	//code1()
	code2()

}
