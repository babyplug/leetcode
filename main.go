package main

import (
	"fmt"

	"github.com/babyplug/leetcode/candies"
	"github.com/babyplug/leetcode/defang_ip"
	"github.com/babyplug/leetcode/good_pairs"
	"github.com/babyplug/leetcode/running_sum"
)

func main() {
	fmt.Println("Running 1d Sum")
	running_sum.Solved()

	fmt.Println("Kids With the Greatest Number of candies")
	candies.Solved()

	fmt.Println("#1512. Number of Good Pairs")
	good_pairs.Solved()

	fmt.Println("#1108. Defanging an IP Address")
	defang_ip.Solved()

}
