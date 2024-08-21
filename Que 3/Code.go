package main

import (
	"fmt"
	"reflect"
)

func findAnagrams(s string, p string) []int {
	var result []int
	pCount := make(map[rune]int)
	sCount := make(map[rune]int)

	for _, char := range p {
		pCount[char]++
	}

	for i, char := range s {
		sCount[char]++
		if i >= len(p) {
			removeChar := rune(s[i-len(p)])
			if sCount[removeChar] == 1 {
				delete(sCount, removeChar)
			} else {
				sCount[removeChar]--
			}
		}
		if reflect.DeepEqual(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	fmt.Println(result) // Output: [0 6]
}
