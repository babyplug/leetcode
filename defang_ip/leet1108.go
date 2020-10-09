package defang_ip

import (
	"fmt"
	"strings"
)

var (
	test1 = "1.1.1.1"
	test2 = "255.100.50.0"
)

func Solved() {
	fmt.Println("Example 1: ")
	fmt.Println("Input: ", test1)
	fmt.Println("Output: ", defangIPaddr(test1))

	fmt.Println("Example 2: ")
	fmt.Println("Input: ", test2)
	fmt.Println("Output: ", defangIPaddr(test2))
}

func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}
