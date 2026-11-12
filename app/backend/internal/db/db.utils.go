package db

import (
	"database/sql"
	"strings"
	"time"
)

type LikeString interface {
	~string
}

func MakePlaceHolder(len int) []string {
	placeholders := make([]string, len)
	for i := 0; i < len; i++ {
		placeholders[i] = "?"
	}
	return placeholders
}

func MakeColm[T ~string](rows ...T) []string {
	res := make([]string, len(rows))
	for i, row := range rows {
		res[i] = string(row)
	}
	return res
}

func ToArgs[T any](items []T) []interface{} {
	args := make([]interface{}, len(items))
	for i, v := range items {
		args[i] = v
	}
	return args
}

func MakeSet(rows []Collom) []string {
	set := make([]string, len(rows))
	for i, row := range rows {
		set[i] = string(row) + " = ?"
	}
	return set
}

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NullTimePtr(t *time.Time) sql.NullTime {
	if t == nil || t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// escapeLike meng-escape karakter spesial LIKE agar input dicari apa adanya.
func EscapeLike(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)
	return r.Replace(s)
}

// isDuplicateEntry mendeteksi error MySQL 1062 (duplicate key)
// tanpa menambah dependency ke driver tertentu.
func IsDuplicateEntry(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "Error 1062") || strings.Contains(msg, "Duplicate entry")
}
