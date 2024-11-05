package http

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestEasyDemo(t *testing.T) {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("pong"))
		if err != nil {
			t.Fatal(err)
		}
	})
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	time.Sleep(time.Second)
	//request := httptest.NewRequest("GET", "http://localhost:8080/ping", nil)
	request, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://127.1:8080/ping")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(all) != "pong" {
		t.Fatalf("Expected 'pong', got '%s'", string(all))
	}
}
