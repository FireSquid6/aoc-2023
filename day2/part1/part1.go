package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Round struct {
	red   int
	blue  int
	green int
}

var bag = Round{
	red:   12,
	green: 13,
	blue:  14,
}

func Solution() {
	file, err := os.Open("./part1/example.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 1
	total := 0
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		idLength := strconv.Itoa(id)
		prefix := len(idLength) + 5

		line = line[prefix:]

		valid := IsValid(bag, GetRounds(line))

		if valid {
			total += id
		}

		id += 1
	}

	fmt.Println(total)
}

func GetRounds(line string) []Round {
	roundStrings := strings.Split(line, ";")
	rounds := []Round{}

	for _, roundString := range roundStrings {
		round := Round{}

		balls := strings.Split(roundString, ",")
		for _, ball := range balls {
			digits := ""

			for _, c := range ball {
				if unicode.IsDigit(c) {
					digits = strings.Join([]string{digits, string(c)}, "")
				} else {
					num, err := strconv.Atoi(digits)

					if err != nil {
						fmt.Println(line)
						log.Fatal("strconv failed. See line 67")
					}

					switch string(c) {
					case "r":
						round.red = num
					case "b":
						round.blue = num
					case "g":
						round.green = num
					}
				}
			}
		}

		rounds = append(rounds, round)
	}

	return rounds
}

func IsValid(bag Round, rounds []Round) bool {
	for _, round := range rounds {
		bag.red -= round.red
		bag.blue -= round.blue
		bag.green -= round.green
	}

	return (bag.red >= 0 && bag.blue >= 0 && bag.green >= 0)
}
