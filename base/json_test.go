package main

import (
	"encoding/json"
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
