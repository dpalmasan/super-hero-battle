package handlers

import (
	"testing"
)

func TestGenerateRandomIds(t *testing.T) {

	ids, err := GenerateRandomIds(10, 3, 9)
	if err == nil || len(ids) != 0 {
		t.Fatal(`Expected error when size < n`)
	}

	ids, err = GenerateRandomIds(10, 3, 100)
	if len(ids) != 10 {
		t.Fatal("Expected an array of ids of length 10!")
	}
}
