package sql2json

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func SQL2JSON(rows *sql.Rows) ([]byte, error) {
	var rs []map[string]interface{}
	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Error getting column names: %v", err)
	}
	var count int

	for rows.Next() {
		var vals []interface{}
		for i := 0; i < len(cols); i++ {
			var val string
			vals = append(vals, &val)
		}
		rows.Scan(vals...)
		var r = make(map[string]interface{})
		// Run through the cols assigning the column name and the value gotten from the query.
		for i, c := range cols {
			r[c] = vals[i]
		}
		rs = append(rs, r)
		count++
	}
	return json.Marshal(rs)
}

func Query(db *sql.DB, selectSQL string) ([]byte, error) {
	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, fmt.Errorf("Error Querying (%v) db: %v", selectSQL, err)
	}
	defer rows.Close()
	return SQL2JSON(rows)
}
