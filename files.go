package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func AppendLine(path, line string) error {
	f, err := CreateOrOpenToAppend(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", line))
	return err
}

func CreateOrOpenToAppend(path string) (*os.File, error) {
	if IsFileExist(path) {
		return os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		return os.Create(path)
	}
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func CopyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := CreateOrOpenToAppend(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return nil
}
