package utils

import "database/sql"

func NullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func NullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: f,
		Valid:   true,
	}
}

func Float64FromNull(n sql.NullFloat64) float64 {
	if !n.Valid {
		return 0
	}

	return n.Float64
}
