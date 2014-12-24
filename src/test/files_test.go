package test

import (
	"files"
	"os"
	"testing"
)

func TestIsFileExistFalse(t *testing.T) {
	if ret := files.IsFileExist("notfound"); ret {
		t.Fail()
	}
}

func TestIsFileExistTrue(t *testing.T) {
	path := "notexistatall"
	if file, err := os.Create(path); err != nil {
		t.Fail()
	} else {
		file.Close()
		defer os.Remove(path)
		if ret := files.IsFileExist(path); !ret {
			t.Fail()
		}
	}
}
