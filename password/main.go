// main.go
// Startet das Password Recovery Programm.
// Davor muss die Datei "rockyou.txt" im gleichen Verzeichnis liegen.
//
// Author: Ajun Anpalakan
// Date: 03.04.2025

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

	//Datei einlesen
	data, err := os.ReadFile("rockyou.txt")
	if err != nil {
		panic(err)
	}

	numWorkers := runtime.NumCPU() * 8

	//aufteilen in Chunks
	chunks := splitChunks(data, numWorkers)

	var wg sync.WaitGroup

	//Jeder Chunk in einer Goroutine verarbeiten
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

	//Speicherverbrauch
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Verbrauchter Speicher: %.4f MB\n", float64(m.Alloc)/1024/1024)

}
