package main

import (
	"fmt"
)

func main() {
	var s = "abcdedcbafghijk"
	
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var c, r, maxLength int
	var longestP string
	ss := preProcess(s)
	length := len(ss)
	p := make([]int, length)
	for i := 0; i < length; i++ {
		if i < r {
			mirror := 2*c - i
			p[i] = min(r-i, p[mirror])
		}
		if nr := expand(ss, i, i+p[i]+1); nr > r {
			p[i] = nr - i
			r = nr
			c = i
		}
		if p[i] > maxLength {
			maxLength = p[i]
			b := (i - maxLength) / 2
			longestP = subString(s, b, b+maxLength-1)
		}
	}
	return longestP
}
func expand(s string, c, b int) int {
	for i := b; i < len(s); i++ {
		if 2*c-i < 0 || s[i] != s[2*c-i] {
			return i - 1
		}
	}
	return len(s) - 1
}
func subString(s string, b, e int) string {
	if e >= len(s)-1 {
		return s[b:]
	}
	return s[b : e+1]
}
func preProcess(s string) string {
	ss := make([]rune, len(s)*2+1)
	for i, val := range s {
		ss[2*i] = rune('*')
		ss[2*i+1] = val
	}
	ss[2*len(s)] = rune('*')
	return string(ss)
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

