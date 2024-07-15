package ardent

import (
	"database/sql"
)

func GetGroups(db *sql.DB) ([]string) {
	var groups []string
	rows, err := db.Query("SELECT name FROM groups")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil
		}
		groups = append(groups, name)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return groups
}