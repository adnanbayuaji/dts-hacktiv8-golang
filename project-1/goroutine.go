package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printLoop(i, &wg)
		go printLoop2(i, &wg)
	}
	wg.Wait()
}
func printLoop2(angka int, wg *sync.WaitGroup) {
	fmt.Println("[bisa1 bisa2 bisa3]", angka)
	wg.Done()
}
func printLoop(angka int, wg *sync.WaitGroup) {
	fmt.Println("[coba1 coba2 coba3]", angka)
	wg.Done()
}
