package sql2json

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type DBConn *sql.DB

func SQL2JSON(rows *sql.Rows) ([]byte, error) {
	var rs []map[string]interface{}
	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Error getting column names: %v", err)
	}

	for rows.Next() {
		vals := make([]interface{}, len(cols), len(cols))
		var r = make(map[string]interface{})
		rows.Scan(vals...)
		// Run through the cols assigning the column name and the value gotten from the query.
		for i, c := range cols {
			r[c] = vals[i]
		}
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
