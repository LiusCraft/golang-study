package main

import (
	"hash/fnv"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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

// ffff331
func TestUrlHashKey(t *testing.T) {
	baseURL := "htt://wxytestdnode.qbox.net/"
	counter := 0

	for {
		h := fnv.New32a()
		testURL := baseURL + strconv.Itoa(counter)
		h.Write([]byte(testURL))
		sum := h.Sum32()

		if sum%100 == 99 {
			log.Printf("找到了！URL: %s, hash值: %d", testURL, sum)
			break
		}
		counter++

		// 每1000次打印一次进度
		if counter%1000 == 0 {
			log.Printf("已尝试 %d 次...", counter)
		}
	}
}
