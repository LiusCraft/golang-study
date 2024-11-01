package main

import (
	"testing"
)

func TestNumber(t *testing.T) {
	a := 0
	for a >= 0 {
		a += 100000000000000000
	}
	a = 0
	t.Log("结束", a)
}
