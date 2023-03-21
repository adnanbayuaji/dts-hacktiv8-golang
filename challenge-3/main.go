package main

import (
	"fmt"
)

func main() {
	var str string = "selamat malam"

	chars := []rune(str)
	fmt.Println(loopStr(chars...))
}

func loopStr(str ...rune) map[string]int {
	result := make(map[string]int)
	for _, v := range str {
		fmt.Println(string(v))
		if val, ok := result[string(v)]; ok {
			result[string(v)] = val + 1
		} else {
			result[string(v)] = 1
		}
	}
	return result
}
