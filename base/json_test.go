package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type jsonA struct {
	CacheKeyQueryParams []string `json:"cacheKeyQueryParams"`
}

func TestSliceToObject(t *testing.T) {
	str, err := json.Marshal(jsonA{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(str))
}

func TestNilVal(t *testing.T) {
	var a *jsonA
	err := json.Unmarshal([]byte("{\"cacheKeyQueryParams\":[\"1\"]}"), a)
	if err == nil {
		t.Fatal(err)
	}

	a = &jsonA{}
	err = json.Unmarshal([]byte("{\"cacheKeyQueryParams\":[\"1\"]}"), a)
	if err != nil {
		t.Fatal(err)
	}
}

type Config struct {
	Config3 Config2
}

type Config2 struct {
	Config3 Config3 `json:"config3"`
}

type Config3 struct {
	CacheKeyQueryParams []string `json:"cacheKeyQueryParams"`
}

func TestInline(t *testing.T) {
	var config Config
	err := json.Unmarshal([]byte(`{"config3":{"cacheKeyQueryParams":["1"]}}`), &config)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.Config3.Config3.CacheKeyQueryParams)
}

func TestCompositeStructure(t *testing.T) {

	type Info struct {
		Name string `json:"name"`
	}

	type User struct {
		Info        // 匿名组合
		Name string `json:"-"` // 忽略 User.Name
	}
	jsonData := `{"name": "aaa"}`
	var user User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Info.Name: %s\n", user.Info.Name) // 输出: Info.Name: aaa
	fmt.Printf("User.Name: %s\n", user.Name)      // 输出: User.Name:
}
