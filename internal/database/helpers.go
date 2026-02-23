package database

import (
	"encoding/json"
)

// JSON is a simple wrapper for MySQL JSON columns.
type JSON []byte

func (j JSON) Value() (any, error) {
	if len(j) == 0 {
		return "[]", nil
	}
	return string(j), nil
}

func (j *JSON) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		*j = append((*j)[0:0], v...)
	case string:
		*j = []byte(v)
	case nil:
		*j = nil
	}
	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("[]"), nil
	}
	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	*j = append((*j)[0:0], data...)
	return nil
}

func NewJSON(v any) JSON {
	b, _ := json.Marshal(v)
	return b
}
