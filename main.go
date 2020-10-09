package main

import (
	"fmt"

	"github.com/babyplug/leetcode/candies"
	"github.com/babyplug/leetcode/good_pairs"
	"github.com/babyplug/leetcode/running_sum"
)

func main() {
	fmt.Println("Running 1d Sum")
	running_sum.Solved()

	fmt.Println("Kids With the Greatest Number of candies")
	candies.Solved()

	fmt.Println("1512. Number of Good Pairs")
	good_pairs.Solved()
}
