package model

import (
	"github.com/rmanzoku/friction"
)

func WarmUpInnoDB() ([]string, error) {
	tables, _ := friction.ShowTables(db)

	for _, t := range tables {
		columns, err := friction.GetIndexColumns(db, t)
		if err != nil {
			panic(err)
		}

		for _, c := range columns {
			friction.WarmUp(db, t, c, 1)
		}
	}

	return tables, nil
}
