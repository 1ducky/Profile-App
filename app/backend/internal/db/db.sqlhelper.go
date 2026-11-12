package db

import "strings"

type Command string

const (
	SELECT Command = "SELECT"
	UPDATE Command = "UPDATE"
	INSERT Command = "INSERT"
	DELETE Command = "DELETE"
	FROM   Command = "FROM"
)

func CreateRow(rows ...Collom) string {
	if len(rows) == 0 {
		return "*"
	}
	if len(rows) == 1 {
		return string(rows[0])
	}
	var strRows []string
	for _, row := range rows {
		strRows = append(strRows, string(row))
	}
	return strings.Join(strRows, ",")
}

func BuildQuery(command Command, table Table, rows ...Collom) string {
	strRows := CreateRow(rows...)
	query := string(command) + strRows + string(FROM) + string(table)
	return query
}
