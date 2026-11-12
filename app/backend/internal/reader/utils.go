package reader

import (
	"bytes"
	"io"
	"net/http"
)

func DetectMime(r io.Reader) (string, io.Reader, error) {
	buf := make([]byte, 512)
	n, err := io.ReadFull(r, buf)
	if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
		return "", nil, err
	}
	buf = buf[:n]

	mimeType := http.DetectContentType(buf)

	// Gabungkan lagi bytes yang sudah dibaca dengan sisa reader
	// supaya reader tetap utuh dari awal untuk dipakai selanjutnya
	fullReader := io.MultiReader(bytes.NewReader(buf), r)

	return mimeType, fullReader, nil
}
