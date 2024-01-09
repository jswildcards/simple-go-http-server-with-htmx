package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description sql.NullString
	DoneAt      sql.NullInt64
}
