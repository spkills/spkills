package model

func WarmUpInnoDB() ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]string, 0)
	for rows.Next() {
		var ret string
		err := rows.Scan(&ret)
		if err != nil {
			return nil, err
		}
		tables = append(tables, ret)
	}

	return tables, nil
}
