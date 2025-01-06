package trappingwater

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAA(t *testing.T) {
	str, _ := json.Marshal(map[string]interface{}{
		"test": "aaa",
		"bbb":  1,
		"hkh":  true,
	})
	log.Println(string(str))
}
