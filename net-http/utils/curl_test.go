package utils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	const url = "https://golang.org"
	r, err := GetContent(url)
	if err != nil {
		t.Error("Get returns error")
		fmt.Println(err)
	}
	if r.Code != 200 {
		t.Error("response code error.")
	}
	if r.ContentLength <= 0 {
		t.Error("response content length error.")
	}
	if r.ContentType == "" {
		t.Error("response content type error.")
	}
	if r.Title == "" {
		t.Error("response title error.")
	}
}
