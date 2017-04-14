package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	str := strings.Fields(s)
	for i := range str {
		m[str[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
