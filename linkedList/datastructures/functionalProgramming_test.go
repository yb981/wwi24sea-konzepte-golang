package datastructures

import (
	"testing"
)

func TestForEach(t *testing.T) {
	testfunc := func (a int) {
		a = a + 2
	}

	tests := []struct {
		input 		[]int
		expected 	[]int
	}{
		{[]int{1,2,3}, []int{3,4,5} },
		{[]int{-3,-2, -1}, []int{-1,0,1} },
	}

	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}
		testList.ForEach(testfunc)

		var resultList LinkedList[int]
		for _, v := range test.expected {
			resultList.Append(v)
		}
		if !testList.Equals(&resultList) {
			t.Errorf("Expected %v, but got %v", resultList.ToString(), testList.ToString())
		}
	}
}

func TestFilter(t *testing.T) {
	isEvenFunc := func (a int) bool {
		return a % 2 == 0
	}

	tests := []struct {
		input 		[]int
		expected 	[]int
	}{
		{[]int{1,2,3}, []int{2} },
		{[]int{-3,-2, -1}, []int{-2} },
		{[]int{50,-2, 20}, []int{50, -2, 20} },
	}

	for _, test := range tests {
		var testList LinkedList[int]
		for _, v := range test.input {
			testList.Append(v)
		}
		testList = testList.Filter(isEvenFunc)

		var resultList LinkedList[int]
		for _, v := range test.expected {
			resultList.Append(v)
		}
		if !testList.Equals(&resultList) {
			t.Errorf("Expected %v, but got %v", resultList.ToString(), testList.ToString())
		}
	}
}