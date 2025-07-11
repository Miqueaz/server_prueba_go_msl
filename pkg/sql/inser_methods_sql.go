package orm_sql

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

// Insert inserts a single record of type T into the table
func (qb *QueryBuilder[T]) Insert(ctx context.Context, entity T) (T, error) {
	// Use reflection to get field names and values
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	t := v.Type()
	var cols []string
	var placeholders []string
	var data = map[string]interface{}{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" || tag == "-" {
			continue
		}
		cols = append(cols, tag)
		placeholders = append(placeholders, fmt.Sprintf("%s", ":"+tag))
		data[tag] = v.Field(i).Interface()
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		qb.table,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := qb.db.NamedExecContext(ctx, query, data)

	//Si la estructura T tiene un campo ID, obtenemos el ID insertado

	return entity, err
}

// InsertMany inserts multiple records of type T into the table in a transaction
func (qb *QueryBuilder[T]) InsertMany(ctx context.Context, entities []T) error {
	tx, err := qb.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, entity := range entities {
		if _, err := qb.Insert(ctx, entity); err != nil {
			return err
		}
	}

	return tx.Commit()
}
