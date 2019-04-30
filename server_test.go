package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}


func TestHandlerGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response status is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Connot retrieve JSON post")
	}
}

func TestHandlerPut(t *testing.T) {
	json := strings.NewReader(`{"content":"Updated post","author":"h-hiroki"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response status is %v", writer.Code)
	}
}