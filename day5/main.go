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

	//var newlines []string
	for i, line := range lines {
		lines[i] = strings.Replace(line, " -> ", ",", -1)
	}

	//fmt.Println(lines)

	var newlines [][]string
	var diagonallines [][]string
	for _, line := range lines {
		split := strings.Split(line, ",")

		if split[0] == split[2] || split[1] == split[3] {
			newlines = append(newlines, split)
		} else {
			diagonallines = append(diagonallines, split)
		}
	}

	fmt.Println(newlines)
	fmt.Println(diagonallines)

	grid := [10][10]int{}

	for _, coord := range newlines {
		//y, _ := strconv.Atoi(coord[1])

		xstart, _ := strconv.Atoi(coord[0])
		xend, _ := strconv.Atoi(coord[2])
		ystart, _ := strconv.Atoi(coord[1])
		yend, _ := strconv.Atoi(coord[3])

		if xstart > xend {
			xend, xstart = xstart, xend
		}

		if ystart > yend {
			yend, ystart = ystart, yend
		}

		if yend == ystart {
			for v := xstart; v <= xend; v++ {

				grid[yend][v] += 1
			}

		} else if xend == xstart {

			for v := ystart; v <= yend; v++ {

				grid[v][xend] += 1
			}
		} else {
			var k = 0
			for v := ystart; v <= yend; v++ {

				grid[v][xstart+k] += 1
				k++
			}

		}

	}

	// SPLIT OUT DIAGONAL FOR NOW ITS HURTING MY HEAD

	//for _, coord := range diagonallines {

	// //y, _ := strconv.Atoi(coord[1])

	// xstart, _ := strconv.Atoi(coord[0])
	// xend, _ := strconv.Atoi(coord[2])
	// ystart, _ := strconv.Atoi(coord[1])
	// yend, _ := strconv.Atoi(coord[3])

	// if xstart > xend {
	// 	xend, xstart = xstart, xend
	// 	yend, ystart = ystart, yend
	// 	fmt.Println("start x ", xstart, "end x ", xend)
	// 	fmt.Println("start y ", xstart, "end y ", xend)
	// }

	// if ystart > yend {
	// 	yend, ystart = ystart, yend
	// 	xend, xstart = xstart, xend

	// }
	// var k = 0
	// for v := xstart; v <= xend; v++ {
	// 	grid[v][ystart+k] += 1
	// 	k++
	// }

	//	}

	var total = 0
	for _, line := range grid {
		fmt.Println(line)
		for _, entry := range line {
			if entry > 1 {
				total += 1
			}
		}
	}

	fmt.Println(total)

}

func main() {
	code1()
	//code2()

}
