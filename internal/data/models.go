package data

import "database/sql"

type Models struct {
	Authors AuthorModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Authors: AuthorModel{DB: db},
	}
}
