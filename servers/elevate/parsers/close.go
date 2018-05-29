package parsers

import (
	"database/sql"
	"fmt"
)

func close(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		fmt.Printf("Error closing rows: %v", err)
	}
}
