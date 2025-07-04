package query_postgres

// Where adds a condition
func (qb *QueryBuilder[T]) Where(field, op string, val interface{}) *QueryBuilder[T] {
	qb.conditions = append(qb.conditions, condition{field, op, val})
	return qb
}

// Limit sets a limit
func (qb *QueryBuilder[T]) Limit(n int) *QueryBuilder[T] {
	qb.limit = n
	return qb
}

// Offset sets an offset
func (qb *QueryBuilder[T]) Offset(n int) *QueryBuilder[T] {
	qb.offset = n
	return qb
}

// OrderBy sets order
func (qb *QueryBuilder[T]) OrderBy(clause string) *QueryBuilder[T] {
	qb.orderBy = clause
	return qb
}
