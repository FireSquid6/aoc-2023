package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	numberStrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := map[int]int{}

		// find all of the digits
		for i, c := range line {

			n, err := strconv.Atoi(string(c))

			if err == nil {
				nums[i] = n
			}
		}

		// find all of the words
		for word := range numberStrings {
			i := strings.Index(line, word)
			n := strings.Count(line, word)

			// need to find the index of every occurence of the word
			if n > 1 {
				fmt.Println(n)
			}

			if i != -1 {
				nums[i] = numberStrings[word]
			}
		}

		keys := make([]int, 0, len(nums))
		for k := range nums {
			keys = append(keys, k)
		}

		minIndex, maxIndex := minMax(keys)

		value := (10 * nums[maxIndex]) + nums[minIndex]

		// fmt.Println(line)
		// fmt.Println(nums)
		// fmt.Println(value)
		// fmt.Println()

		answer += value
	}

	fmt.Println("\nThe final answer is:")
	fmt.Println(answer)
}

func minMax(arr []int) (int, int) {
	// Initialize the variables to hold the maximum and minimum values to draw comparisons.
	max := arr[0]
	min := arr[0]
	// Iterate over the array
	for i := 0; i < len(arr); i++ {
		// if the current element is greater than the present maximum
		if arr[i] > max {
			max = arr[i]
		}
		// if the current element is smaller than the present minimum
		if arr[i] < min {
			min = arr[i]
		}
	}

	return max, min
}
