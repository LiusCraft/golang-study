package main

import (
	"log"
	"testing"
)

func test() func() func() {
	log.Println("hello1")
	return func() func() {
		log.Println("hello2")
		return func() {
			log.Println("hello3")
		}
	}
}

func TestPrefixAndSuffix(t *testing.T) {
	defer test()()()
	//defer test()
	log.Println("mid")

}
