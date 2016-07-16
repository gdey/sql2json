/* sql2json is a very simple package that provides a method for taking an
initialized db and an SQL and returns the json representation of the returned
rows.
*/
package sql2json

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// SQL2JSON takes a set of rows and returns a json representation of the rows.
// The main object is an array, wich each rows consisting of a map if column_name
// values pairs. All fields are returned as strings.
func SQL2JSON(rows *sql.Rows) ([]byte, error) {
	var rs []map[string]interface{}
	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Error getting column names: %v", err)
	}
	for rows.Next() {
		var vals []interface{}
		for i := 0; i < len(cols); i++ {
			// TODO
			// Since it's simpler to just make everything a string. However, it
			// Would be better to convert the values to actual values.
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
	}
	return json.Marshal(rs)
}

// Query takes a db, and a select statment and returns a JSON representation of the
// rows returned. As with SQL2JSON all fields values are returned as strings.
func Query(db *sql.DB, selectSQL string) ([]byte, error) {
	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, fmt.Errorf("Error Querying (%v) db: %v", selectSQL, err)
	}
	defer rows.Close()
	return SQL2JSON(rows)
}
