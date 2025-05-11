// passwordRecov_test.go
// Test für Passwortwiederherstellung in Go
//
// Author: Ajun Anpalakan
// Date: 03.04.2025
package main

import (
	"bytes"
	"crypto/md5"
	"testing"
	"time"
)

/*
Commands fuer Test

go test -v
go test -cover

go test -coverprofile=coverage
go tool cover -html=coverage -o coverage.html

*/

func TestParseHash(t *testing.T) {
	expected := md5.Sum([]byte("password"))
	actual := parseHash("5f4dcc3b5aa765d61d8327deb882cf99")
	if actual != expected {
		t.Errorf("parseHash failed: expected %x, got %x", expected, actual)
	}
}

func TestHash(t *testing.T) {
	input := []byte("test123")
	expected := md5.Sum(input)
	actual := hash(input)
	if actual != expected {
		t.Errorf("hash failed: expected %x, got %x", expected, actual)
	}
}

func TestSplitChunks(t *testing.T) {
	data := []byte("one\ntwo\nthree\nfour\nfive\nsix\nseven\neight\nnine\nten\n")
	parts := 3
	chunks := splitChunks(data, parts)

	if len(chunks) != parts {
		t.Errorf("Expected %d chunks, got %d", parts, len(chunks))
	}

	combined := bytes.Join(chunks, []byte{})
	if !bytes.Equal(data, combined) {
		t.Errorf("Recombined chunks do not match original data")
	}
}

func TestProcess(t *testing.T) {
	originalResults := results
	originalTimes := resultTimes
	originalMask := resultMask
	defer func() {
		results = originalResults
		resultTimes = originalTimes
		resultMask = originalMask
	}()

	cases := []struct {
		name, input, expected string
	}{
		{"Match with affix", "foo\npasswort\nbar\n", "!passwort"},
		{"Last line no newline", "foo\nbar\nfinalpass", "#finalpass"},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			results = [3]string{}
			resultTimes = [3]time.Duration{}
			resultMask = 0
			targetHashes[0] = md5.Sum([]byte(tc.expected))

			process([]byte(tc.input), time.Now())

			if got := results[0]; got != tc.expected {
				t.Errorf("Test %d: expected %q, got %q", i, tc.expected, got)
			}
			if resultMask&1 == 0 {
				t.Errorf("Test %d: resultMask bit 0 not set", i)
			}
		})
	}
}

func TestProcess_IgnoreInvalidLines(t *testing.T) {
	input := bytes.Join([][]byte{
		[]byte(""),                          // leer
		bytes.Repeat([]byte("a"), maxLen+1), // zu lang
		[]byte("meinpasswort\r"),            // mit CR
	}, []byte("\n"))

	results = [3]string{}
	resultMask = 0
	targetHashes[0] = md5.Sum([]byte("!meinpasswort"))

	process(input, time.Now())

	if results[0] != "" {
		t.Errorf("Expected no match, got %q", results[0])
	}
}
