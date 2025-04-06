package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	data, err := os.ReadFile("rockyou.txt")
	if err != nil {
		panic(err)
	}

	numWorkers := runtime.NumCPU() * 8

	chunks := splitChunks(data, numWorkers)

	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []byte) {
			defer wg.Done()
			process(chunk, start)
		}(chunk)
	}

	wg.Wait()

	fmt.Println("Ergebnisse:")
	for i, res := range results {
		if res != "" {
			fmt.Printf("Hash %x => Passwort: %q (in %v)\n", targetHashes[i], res, resultTimes[i])
		}
	}
	fmt.Printf("Gesamtdauer: %v\n", time.Since(start))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Verbrauchter Speicher: %.4f MB\n", float64(m.Alloc)/1024/1024)

}
