package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any, maxRead int64) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxRead)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		var (
			syntaxErr *json.SyntaxError
			typeErr   *json.UnmarshalTypeError
			maxErr    *http.MaxBytesError
		)
		switch {
		case errors.Is(err, io.EOF):
			return errors.New("request body must not be empty")
		case errors.As(err, &maxErr):
			return fmt.Errorf("request body must not exceed %d bytes", maxRead)
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("malformed JSON at position %d", syntaxErr.Offset)
		case errors.As(err, &typeErr):
			return fmt.Errorf("field %q must be of type %s", typeErr.Field, typeErr.Type)
		default:
			// termasuk "unknown field ..." dan format waktu yang salah
			return errors.New(err.Error())
		}
	}

	// Tolak jika ada JSON kedua setelah objek pertama.
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("request body must contain a single JSON object")
	}
	return nil
}
