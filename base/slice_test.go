package main

import (
	"testing"
)

func TestRange(t *testing.T) {
	list := []int{1, 2, 3, 4}
	for i, e := range list {
		t.Log(i, e)
	}
}
