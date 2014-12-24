package net

import (
	"files"
	"fmt"
	"io"
	"net/http"
	"os"
)

type DownloadListener interface {
	OnStart()
	OnDone()
}

func DownloadFile(url, path string, listener DownloadListener) error {
	if listener != nil {
		listener.OnStart()
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get url content")
		return err
	}

	defer res.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create file:", path)
		return err
	}

	io.Copy(file, res.Body)

	if listener != nil {
		listener.OnDone()
	}

	return nil
}

func DownloadFileIfNotExist(url, path string, listener DownloadListener) error {
	if files.IsFileExists(path) {
		return nil
	}

	return DownloadFile(url, path, listener)
}
