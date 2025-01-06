package main

import (
	"fmt"
	"testing"
)

type A struct {
	*B   `json:"b"`
	name string
}
type B struct {
	list []string `json:"list"`
}

func (a *A) Change() {
	a.name = "change"
	a.B = &B{list: []string{"a", "b"}}
	a.B.list[0] = "c"
}

func TestJsonUnmarshal(t *testing.T) {
	a := A{}
	a.Change()
	t.Logf("%#+v, %#+v", a, a.B)
}

func testfff() {
	fmt.Println("testfff successed")
}

func TestPointToId(t *testing.T) {
	funs := map[string]func(){}
	funs[fmt.Sprintf("%p", testfff)] = testfff
	funs[fmt.Sprintf("%p", testfff)]()
}
	