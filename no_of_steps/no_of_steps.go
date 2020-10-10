package no_of_steps

import "fmt"

func numberOfSteps(num int) int {
	res := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		res++
	}

	fmt.Println(res)
	return res
}
