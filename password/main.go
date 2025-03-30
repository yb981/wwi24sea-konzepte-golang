package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const maxLen = 255
var targetHashes = map[string]bool{
	"32c5c26e20908ebd80269d32f51cb5bb": true,
	"648d5d9cc7cafe536fdbc6331f00c6a0": true,
	"d31daf6579548a2a1bf5a9bd57b5bb89": true,
}
var affixes = "!#+0123456789"

type Result struct {
	Password string
	Duration time.Duration
}

func hashMD5(input string) string {
	h := md5.Sum([]byte(input))
	return hex.EncodeToString(h[:])
}

func processLine(line string, found *int32, results map[string]*Result, start time.Time, mu *sync.Mutex) {
	if len(line) == 0 || len(line) > maxLen || atomic.LoadInt32(found) >= int32(len(targetHashes)) {
		return
	}
	originalHash := hashMD5(line)
	mu.Lock()
	if targetHashes[originalHash] && results[originalHash] == nil {
		results[originalHash] = &Result{line, time.Since(start)}
		atomic.AddInt32(found, 1)
	}
	mu.Unlock()
	for _, aff := range affixes {
		if atomic.LoadInt32(found) >= int32(len(targetHashes)) {
			return
		}
		suffixed := line + string(aff)
		if h := hashMD5(suffixed); targetHashes[h] {
			mu.Lock()
			if results[h] == nil {
				results[h] = &Result{suffixed, time.Since(start)}
				atomic.AddInt32(found, 1)
			}
			mu.Unlock()
		}
		prefixed := string(aff) + line
		if h := hashMD5(prefixed); targetHashes[h] {
			mu.Lock()
			if results[h] == nil {
				results[h] = &Result{prefixed, time.Since(start)}
				atomic.AddInt32(found, 1)
			}
			mu.Unlock()
		}
	}
}

func main() {
	start := time.Now()
	found := int32(0)
	results := make(map[string]*Result)

	file, err := os.Open("rockyou.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	var mu sync.Mutex
	workerPool := make(chan struct{}, 16)

	for scanner.Scan() {
		line := scanner.Text()
		workerPool <- struct{}{}
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			processLine(l, &found, results, start, &mu)
			<-workerPool
		}(line)
	}
	wg.Wait()
	fmt.Println("Ergebnisse:")
	for hash, res := range results {
		if res != nil {
			fmt.Printf("Hash %s => Passwort: %s (in %v gefunden)\n", hash, res.Password, res.Duration)
		}
	}
	fmt.Printf("Gesamtdauer: %v\n", time.Since(start))
}
