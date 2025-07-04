package query_postgres

import (
	"github.com/jmoiron/sqlx"
)

type condition struct {
	Field string
	Op    string
	Val   interface{}
}

type QueryBuilder[T any] struct {
	table      string
	db         *sqlx.DB
	conditions []condition
	orderBy    string
	limit      int
	offset     int
}

type inser struct {
}

type delete struct {
}

type update struct {
}
