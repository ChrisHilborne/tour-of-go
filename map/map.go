package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Split(s, " ")
	for _,value := range words {
		if _, present := m[value]; present {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
