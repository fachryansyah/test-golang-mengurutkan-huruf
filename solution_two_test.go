package main

import "testing"

func TestSortWord(test *testing.T) {
	result := SortWord("osama")
	if result != "aaoms" {
		t.Errorf("SortWord function was error, want: %s got: %s", "aaoms", result)
	}
}
