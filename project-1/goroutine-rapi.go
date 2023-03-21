package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		mtx.Lock()
		go printLoop2(i, &wg, &mtx)
		go printLoop(i, &wg)
	}
	wg.Wait()
}
func printLoop(angka int, wg *sync.WaitGroup) {
	fmt.Println("[coba1 coba2 coba3]", angka)
	wg.Done()
}
func printLoop2(angka int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	fmt.Println("[bisa1 bisa2 bisa3]", angka)
	wg.Done()
	mtx.Unlock()
}
