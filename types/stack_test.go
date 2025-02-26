// stack_test.go
// Testing for stack.go file
//
// Author: Lukas Gröning
// Date: 26.02.2025
//
// This file contains tests.

package types

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	q := NewStack[float64]()
	if q == nil {
		t.Error("Expected a new Stack, instead recieved nil")
	}
}
