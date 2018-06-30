package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
	title    string
	finished bool
	created  time.Time
}

var db, err = sql.Open("mysql", "root:root@/todo?charset=utf8")

func get_all(query *sql.Rows) []map[string]string {
	column, _ := query.Columns()
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}
	s := make(type, 0)
	results := make([]map[string]string, len(column))
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			fmt.Println(err)
			return nil
		}
		row := make(map[string]string)
		for k, v := range values {
			key := column[k]
			row[key] = string(v)
		}
		fmt.Println(row)
		results = append(results, row)
	}
	return results
}
