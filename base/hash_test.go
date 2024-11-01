package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Gf struct {
	Name string
	Age  int
	Sub  struct {
		Name string
		f    *int
	}
}

func TestStruct(t *testing.T) {
	ast := assert.New(t)

	a := 1

	Sub := struct {
		Name string
		f    *int
	}{"1", &a}

	key1 := Gf{"1", 1, Sub}
	key2 := Gf{"1", 1, Sub}

	log.Println(key1, key2)
	ast.Equal(key1, key2)

	tesMap := map[Gf]bool{}

	tesMap[key1] = true
	ast.Equal(true, tesMap[key2])
	log.Println(tesMap[key2])
}
