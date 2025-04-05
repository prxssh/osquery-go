package utils

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MapStringField(data map[string]any, key string, target *sql.NullString) {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case string:
			target.String = v
			target.Valid = true
		case nil:
			target.Valid = false
		default:
			target.String = fmt.Sprintf("%v", v)
			target.Valid = true
		}
	}
}

func MapInt32Field(data map[string]any, key string, target *sql.NullInt32) {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case int32:
			target.Int32 = v
			target.Valid = true
		case int:
			target.Int32 = int32(v)
			target.Valid = true
		case int64:
			target.Int32 = int32(v)
			target.Valid = true
		case float64:
			target.Int32 = int32(v)
			target.Valid = true
		case string:
			if v != "" {
				if i, err := strconv.ParseInt(v, 10, 32); err == nil {
					target.Int32 = int32(i)
					target.Valid = true
				}
			}
		case nil:
			target.Valid = false
		}
	}
}

func MapInt64Field(data map[string]any, key string, target *sql.NullInt64) {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case int64:
			target.Int64 = v
			target.Valid = true
		case int:
			target.Int64 = int64(v)
			target.Valid = true
		case int32:
			target.Int64 = int64(v)
			target.Valid = true
		case float64:
			target.Int64 = int64(v)
			target.Valid = true
		case string:
			if v != "" {
				if i, err := strconv.ParseInt(v, 10, 64); err == nil {
					target.Int64 = i
					target.Valid = true
				}
			}
		case nil:
			target.Valid = false
		}
	}
}

func MapFloat64Field(data map[string]any, key string, target *sql.NullFloat64) {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case float64:
			target.Float64 = v
			target.Valid = true
		case int:
			target.Float64 = float64(v)
			target.Valid = true
		case int32:
			target.Float64 = float64(v)
			target.Valid = true
		case int64:
			target.Float64 = float64(v)
			target.Valid = true
		case string:
			if v != "" {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					target.Float64 = f
					target.Valid = true
				}
			}
		case nil:
			target.Valid = false
		}
	}
}

func MapBoolField(data map[string]any, key string, target *sql.NullBool) {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case bool:
			target.Bool = v
			target.Valid = true
		case string:
			if v != "" {
				if b, err := strconv.ParseBool(v); err == nil {
					target.Bool = b
					target.Valid = true
				}
			}
		case int:
			target.Bool = v != 0
			target.Valid = true
		case float64:
			target.Bool = v != 0
			target.Valid = true
		case nil:
			target.Valid = false
		}
	}
}
