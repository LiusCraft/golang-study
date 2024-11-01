package testcover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	ass := assert.New(t)
	ass.Equal("", toCamelCase(""))
	ass.Equal("Helloworld", toCamelCase("hello_world"))
	ass.Equal("Helloworld", toCamelCase("HelloWorld"))
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("defer 1")
	}()
	a := 1
aaa:
	defer func() {
		fmt.Println("defer 2")
	}()
	a++
	if a < 2 {
		goto aaa
	}

	for i := 0; i < 10; i++ {
		defer func(a int) {
			fmt.Println(i, a)
		}(i)
	}
}
