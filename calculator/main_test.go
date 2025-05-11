// main.go
//
// Test zum starten das RPN Calculator Programm.
//
// test-coverage:
//
// Author: Ajun Anpalakan
// Date: 22.02.2025
package main

import (
	"os"
	"testing"
)

func TestGetInput(t *testing.T) {
	// Erstellt temporäre Datei testinput für simulierten Input
	tmpFile, err := os.CreateTemp("", "testinput")
	if err != nil {
		t.Fatalf("Fehler beim Erstellen der Datei: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Datei wir NACH dem Test gelöscht

	_, err = tmpFile.WriteString("testeingabe\n")
	if err != nil {
		t.Fatalf("Fehler beim Schreiben in die Testdatei: %v", err)
	}
	tmpFile.Seek(0, 0)

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = tmpFile

	// Funktion testen
	result := getInput()
	expected := "testeingabe"

	if result != expected {
		t.Errorf("Erwartet %q, aber erhalten %q", expected, result)
	}
}
