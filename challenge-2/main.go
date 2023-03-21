package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
		if i == 4 {
			for j := 0; j < 11; j++ {
				if j != 5 {
					fmt.Println("Nilai j =", j)
					if j == 4 {
						const rusia = "САШАРВО"
						for index, runeValue := range rusia {
							fmt.Printf("character %#U starts at byte position %d\n", runeValue, index)
						}
					}
				}
			}
		}
	}
}
