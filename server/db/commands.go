package db

import (
	"database/sql"
)

// GetEntries Get multiple entries from the database based on the query
func GetEntries(db *sql.DB, query string) []map[string]any {
	results := []map[string]any{}

	rows, err := db.Query(query)
	if err != nil {
		return results
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return results
	}

	for rows.Next() {
		// Hold the colum and a pointer to the item in the column
		columns := make([]any, len(cols))
		colPointers := make([]any, len(cols))

		for i := range cols {
			colPointers[i] = &columns[i]
		}

		// Scan the result
		if err := rows.Scan(colPointers...); err != nil {
			return results
		}

		// Create the map and retrieve the value for each colum from the pointer slice
		m := make(map[string]any)
		for i, colName := range cols {
			val := colPointers[i].(*any)
			m[colName] = *val
		}

		results = append(results, m)
	}

	return results
}

// GetEntry Get one single entry from the database
func GetEntry(db *sql.DB, query string) map[string]any {
	var result map[string]any

	results := getEntries(db, query)

	if len(results) > 0 {
		result = results[0]
	} else {
		result = make(map[string]any)
	}

	return result
}
