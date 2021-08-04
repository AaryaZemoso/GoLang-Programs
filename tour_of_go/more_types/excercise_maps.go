package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	count := make(map[string]int)
	for i := range str{
		_, ok := count[str[i]]
		if ok{
			count[str[i]]++;
		} else {
			count[str[i]] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
