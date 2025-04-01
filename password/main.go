package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const maxLen = 255
const affixes = "!#+0123456789"

var targetHashes = [][16]byte{
	parseHash("32c5c26e20908ebd80269d32f51cb5bb"),
	parseHash("648d5d9cc7cafe536fdbc6331f00c6a0"),
	parseHash("d31daf6579548a2a1bf5a9bd57b5bb89"),
}

var results = [3]string{}
var resultTimes = [3]time.Duration{}
var resultMask uint32 // Bitmaske für gefundene Hashes

func parseHash(hexStr string) [16]byte {
	b, _ := hex.DecodeString(hexStr)
	var h [16]byte
	copy(h[:], b)
	return h
}

func hash(b []byte) [16]byte {
	return md5.Sum(b)
}

func main() {
	start := time.Now()
	file, err := os.Open("rockyou.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numWorkers := runtime.NumCPU() * 4
	lines := make(chan string, 1000)
	var wg sync.WaitGroup

	// Starte Worker
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(lines, start)
		}()
	}

	// Datei zeilenweise einlesen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || len(line) > maxLen {
			continue
		}
		// Abbruch sobald alle Hashes gefunden sind
		if atomic.LoadUint32(&resultMask) == 0b111 {
			break
		}
		lines <- line
	}
	close(lines)
	wg.Wait()

	fmt.Println("Ergebnisse:")
	for i, res := range results {
		if res != "" {
			fmt.Printf("Hash %x => Passwort: %q (in %v)\n", targetHashes[i], res, resultTimes[i])
		}
	}
	fmt.Printf("Gesamtdauer: %v\n", time.Since(start))
}

func worker(lines <-chan string, start time.Time) {
	buf := make([]byte, maxLen+10)

	for line := range lines {
		if atomic.LoadUint32(&resultMask) == 0b111 {
			return // abbrechen wenn alle gefunden
		}

		// Original testen
		check([]byte(line), start)

		// affix vorne
		copy(buf[1:], line)
		for i := 0; i < len(affixes); i++ {
			buf[0] = affixes[i]
			check(buf[:len(line)+1], start)
		}

		// affix hinten
		copy(buf, line)
		for i := 0; i < len(affixes); i++ {
			buf[len(line)] = affixes[i]
			check(buf[:len(line)+1], start)
		}
	}
}

func check(word []byte, start time.Time) {
	h := hash(word)

	for i, target := range targetHashes {
		mask := uint32(1 << i)
		if atomic.LoadUint32(&resultMask)&mask != 0 {
			continue
		}
		if h == target {
			if atomic.CompareAndSwapUint32(&resultMask, resultMask, resultMask|mask) {
				results[i] = string(word)
				resultTimes[i] = time.Since(start)
				fmt.Printf("✅ MATCH: %q => Hash %x\n", word, h)
			}
		}
	}
}
