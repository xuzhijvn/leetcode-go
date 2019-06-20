package main

import (
	"fmt"
	"strings"
	"time"
)

/**
时间复杂度O(n^2)
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	//step1 : abcbd -> #a#b#c#b#d#
	s = getHelpString(s)
	//step2 : for loop
	longest := 1
	res := ""
	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1
		length := 1
		for left >= 0 && right < len(s) {
			if string(s[left]) == string(s[right]) {
				length += 2
				if length > longest {
					longest = length
					res = s[left : right+1]
				}
				left--
				right++
			} else {
				break
			}
		}
	}
	return strings.Replace(res, "#", "", -1)
}

func getHelpString(s string) string {
	if s == "" {
		return ""
	}
	res := "#"
	for i := 0; i < len(s); i++ {
		res += string(s[i]) + "#"
	}
	return res
}

func main() {
	start := time.Now()
	s := "aaabcaaaaaa"
	fmt.Println(longestPalindrome(s))
	fmt.Println(time.Now().Sub(start))
}
