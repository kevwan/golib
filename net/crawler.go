package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Creator func(map[string]string) interface{}

func Extract(url, pattern string, fn Creator) []interface{} {
	result := []interface{}{}
	body := FetchUrlContent(url)
	items := ExtractItems(body, pattern)
	for _, item := range items {
		result = append(result, fn(item))
	}
	return result
}

func FetchUrlContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occurred while accessing url: %s, error: %v", url, err))
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occurred while paring response body of url: %s, error: %v", url, err))
		return ""
	}
	return string(body)
}

func ExtractItems(str, pattern string) []map[string]string {
	result := []map[string]string{}
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, item := range matches {
		entry := make(map[string]string)
		for i, name := range re.SubexpNames() {
			entry[name] = item[i]
		}
		result = append(result, entry)
	}
	return result
}
