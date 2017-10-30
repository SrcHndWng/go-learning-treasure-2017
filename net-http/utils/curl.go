package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Result contains curl response contents.
type Result struct {
	Code          int    `json:"code"`
	ContentType   string `json:"contentType"`
	ContentLength int64  `json:"contentLength"`
	Title         string `json:"title"`
}

var result = Result{}

// GetContent calls net/http methods and returns Result.
func GetContent(url string) (*Result, error) {
	if err := getHead(url); err != nil {
		return nil, err
	}
	if err := getTitle(url); err != nil {
		return nil, err
	}
	return &result, nil
}

func getHead(url string) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	result.Code = resp.StatusCode
	result.ContentType = resp.Header.Get("Content-Type")
	result.ContentLength = resp.ContentLength
	return nil
}

func getTitle(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	html := string(byteArray)
	getTitleStr := func(html string) string {
		afterTag := strings.SplitAfter(html, "<title>")[1]
		beforeTag := strings.SplitAfter(afterTag, "</title>")[0]
		title := strings.Replace(beforeTag, "</title>", "", -1)
		return title
	}
	result.Title = getTitleStr(html)
	return nil
}
