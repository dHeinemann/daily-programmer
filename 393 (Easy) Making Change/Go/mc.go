// Daily Programmer #393
// https://www.reddit.com/r/dailyprogrammer/comments/nucsik/20210607_challenge_393_easy_making_change/
package main

import (
	"fmt"
)

func getNumCoins(value int) int {
	currentValue := value
	coinCount := 0

	denominations := []int{
		500,
		100,
		25,
		10,
		5,
		1,
	}

	for i := 0; i < len(denominations); i++ {
		for currentValue >= denominations[i] {
			currentValue -= denominations[i]
			coinCount++
		}
	}

	return coinCount
}

func main() {
	fmt.Print("> ")

	var value int
	fmt.Scanf("%d", &value)

	fmt.Printf("%d coins\n", getNumCoins(value))
}
