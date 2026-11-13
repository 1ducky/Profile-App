package main

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

func NewFileKey(raw string) (string, error) {
	if raw == "" {
		return "", errors.New("file key kosong")
	}

	k := strings.ReplaceAll(raw, `\`, "/") // "\" -> "/" di semua OS
	k = path.Clean(k)                      // rapikan "//", "./", dan "a/../b"
	k = strings.TrimPrefix(k, "/")         // pastikan relatif

	// cegah path traversal
	if k == "." || k == ".." || strings.HasPrefix(k, "../") {
		return "", errors.New("file key tidak valid")
	}

	return k, nil
}

func main() {

	constent := []string{"uploads/f19a4e7b-a323-4a13-aaa1-742338ba5353_274dc8aa-3fe5-48a5-8dc7-0d85fb2d0105.png", "uploads\\f19a4e7b-a323-4a13-aaa1-742338ba5353_274dc8aa-3fe5-48a5-8dc7-0d85fb2d0105.png"}
	for _, item := range constent {
		fmt.Println(NewFileKey(item))
	}
}
