package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solution() {
	file, err := os.Open("./part-1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := []int{}

		for _, c := range line {
			n, err := strconv.Atoi(string(c))

			if err == nil {
				nums = append(nums, n)
			}
		}
		fmt.Println(nums)
		fmt.Println(nums[0] + nums[len(nums)-1])

		answer += (10 * nums[0]) + nums[len(nums)-1]
	}

	fmt.Println("\nThe final answer for part 1 is:")
	fmt.Println(answer)
}
