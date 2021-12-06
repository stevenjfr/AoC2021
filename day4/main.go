package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func buildBingoCards(lines []string) [][][]int {
	var cards [][][]int

	var lineCounter = 0

	var newCard [][]int
	for _, line := range lines {

		re := regexp.MustCompile("[0-9]+")
		matches := re.FindAllString(line, -1)

		if len(matches) > 0 {

			var numbers []int

			for _, number := range matches {
				value, _ := strconv.Atoi(number)
				numbers = append(numbers, value)

			}
			newCard = append(newCard, numbers)
			if lineCounter > 3 {
				cards = append(cards, newCard)
				lineCounter = 0
				newCard = [][]int{}
			} else {
				lineCounter++
			}

		} else {
			//cards = append(cards, newCard)
		}

	}
	return cards

}

func AnyWinners(cards [][][]int, cardsWon []int) []int {
	var minusone = -1
	var winners []int
	for i, card := range cards {

		if contains(cardsWon, i) {
			//return nil
		} else {
			for _, line := range card {
				if line[0] == minusone && line[1] == minusone && line[2] == minusone && line[3] == minusone && line[4] == minusone {
					winners = append(winners, i)
				}
			}

			var v = 0
			for v < 5 {
				if card[0][v] == minusone && card[1][v] == minusone && card[2][v] == minusone && card[3][v] == minusone && card[4][v] == minusone {
					winners = append(winners, i)
				}
				v++
			}
		}
	}
	return winners
}

func markCards(cards [][][]int, called int) [][][]int {

	for i, card := range cards {

		for v, line := range card {

			for j, number := range line {
				if number == called {
					cards[i][v][j] = -1
				}
			}

		}
	}

	return cards
}

func calcuateAnswer(cards [][][]int, cardsCopy [][][]int, card int) int {

	winningCard := cards[card]

	var total = 0
	for i, line := range winningCard {

		for v, number := range line {
			if cardsCopy[card][i][v] == -1 {
			} else {
				total = (total + number)
			}

		}
	}
	return total
}

func code1() {

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	numbers, lines := lines[0], lines[1:]

	numbersArray := strings.Split(numbers, ",")

	//fmt.Println(numbers)

	cards := buildBingoCards(lines)

	cardsCopy := buildBingoCards(lines)

	var cardsWon []int
	for _, number := range numbersArray {
		//println(i)
		num, _ := strconv.Atoi(number)

		cardsCopy = markCards(cardsCopy, num)
		card := AnyWinners(cardsCopy, cardsWon)

		if card == nil {
		} else {
			for _, val := range card {
				if contains(cardsWon, val) {
				} else {
					cardsWon = append(cardsWon, val)
					fmt.Println("winner is ", card)
					total := calcuateAnswer(cards, cardsCopy, val)
					totalll := total * num
					fmt.Println("total is ", totalll)
					// if total == 0 {
					// 	os.Exit(0)
					// }
					//os.Exit(0)
				}
			}
		}
		//	println("end")
	}
	//println("end")
	//fmt.Println(cards)

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	code1()
	//code2()

}
