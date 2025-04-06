package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

// Bitmaske fuer welche Hashes schon gefunden
var resultMask uint32

// Hex-String zu [16]byte-Array für MD5-Hash
func parseHash(hexStr string) [16]byte {
	b, _ := hex.DecodeString(hexStr)
	var h [16]byte
	copy(h[:], b) // kopiert in [16]byte
	return h
}

func hash(b []byte) [16]byte {
	return md5.Sum(b)
}

func splitChunks(data []byte, parts int) [][]byte {
	total := len(data)
	chunks := make([][]byte, 0, parts)

	start := 0
	for i := 0; i < parts; i++ {
		end := start + (total-start)/(parts-i)
		if i < parts-1 {
			// Bis zum nächsten \n (Zeilenende)
			for end < total && data[end] != '\n' {
				end++
			}
			if end < total {
				end++ // das \n mitnehmen
			}
		} else {
			end = total // letzter Chunk bekommt den Rest
		}
		chunks = append(chunks, data[start:end])
		start = end
	}
	return chunks
}

func process(chunk []byte, start time.Time) {
	// Puffer fuer Varianten mit affix-Zeichen
	buf := make([]byte, maxLen+10)

	// läuft weiter wenn nicht alle Hashes gefunden wurden und noch Daten im Chunk sind --> sonst abruch
	for atomic.LoadUint32(&resultMask) != 0b111 && len(chunk) > 0 {
		idx := 0
		for idx < len(chunk) && chunk[idx] != '\n' {
			idx++
		}
		line := chunk[:idx]
		if idx < len(chunk) {
			chunk = chunk[idx+1:]
		} else {
			chunk = nil
		}

		if len(line) == 0 || len(line) > maxLen {
			continue
		}

		// original
		check(line, start)

		//zeichen vorne
		copy(buf[1:], line)
		for i := 0; i < len(affixes); i++ {
			buf[0] = affixes[i]
			check(buf[:len(line)+1], start)
		}

		//zeichen hinten
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

		//wenn hash schon gefunden skip
		if atomic.LoadUint32(&resultMask)&mask != 0 {
			continue
		}

		if h == target {
			// exklusiv sichern (thread-sicher)
			if atomic.CompareAndSwapUint32(&resultMask, resultMask, resultMask|mask) {
				results[i] = string(word)
				resultTimes[i] = time.Since(start)
				fmt.Printf("MATCH: %q => Hash %x\n", word, h)
			}
		}
	}
}
