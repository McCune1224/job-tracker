package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Job struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Company  string    `json:"company"`
	Location string    `json:"location"`
	Salary   float64   `json:"salary"`
}

type JobModel struct {
	DB *sql.DB
}
