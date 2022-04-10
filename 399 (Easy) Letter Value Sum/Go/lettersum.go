// Daily Programmer #399 (Easy)
// https://www.reddit.com/r/dailyprogrammer/comments/onfehl/20210719_challenge_399_easy_letter_value_sum/

package main

import (
	"fmt"
)

func letterSum(s string) int {
	if s == "" {
		return 0
	}

	var sum int
	for i := 0; i < len(s); i++ {
		sum += int(s[i]) - 96
	}

	return sum
}

func main() {
	fmt.Printf("> ")
	var input string
	fmt.Scanln(&input)

	fmt.Printf("%v\n", letterSum(input))
}
