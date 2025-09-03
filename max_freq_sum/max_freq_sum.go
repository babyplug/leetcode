package main

import "fmt"

func MaxFreqSum(s string) int {
	vowels := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	maxVowel := 0

	consonant := make(map[rune]int, 'z'-'a'+1)

	maxConsonant := 0
	for _, r := range s {
		if c, ok := vowels[r]; ok {
			c = c + 1
			vowels[r] = c
			if c > maxVowel {
				maxVowel = c
			}
			continue
		}
		c, ok := consonant[r]
		if !ok {
			c = 0
		}
		c = c + 1
		consonant[r] = c
		if c > maxConsonant {
			maxConsonant = c
		}
	}

	return maxVowel + maxConsonant
}

func main() {
	fmt.Println(MaxFreqSum("successes"))
}
