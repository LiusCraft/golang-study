package main

import (
	"testing"
	"time"
)

func TestTimeUnixToTime(t *testing.T) {
	timestamp := time.Now().Unix()
	for i := 0; i < 24; i++ {
		tm := time.Unix(timestamp+int64(i*3600), 0)
		t.Log(tm.Format("2006-01-02 15:04:05"), tm.Unix()%86400, tm.UnixMilli()%86400000)
	}

}
