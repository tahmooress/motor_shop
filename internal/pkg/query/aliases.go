package query

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(ns.String)
	if err != nil {
		return nil, fmt.Errorf("nullString MarshalJSON >> %w", err)
	}

	return data, nil
}

//func (ns *NullString) UnmarshalJSON(b []byte) error{
//	if len
//}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(nf.Float64)
	if err != nil {
		return nil, fmt.Errorf("nullFloat64 MarshalJSON >> %w", err)
	}

	return data, nil
}

type NullTime struct {
	sql.NullTime
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(nt.Time)
	if err != nil {
		return nil, fmt.Errorf("nullTime MarshalJSON >> %w", err)
	}

	return data, nil
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if len(b) > 0 || b != nil {
		nt.Valid = true

		return json.Unmarshal(b, &nt.Time)
	}

	nt.Valid = false

	return nil
}

type NullBool struct {
	sql.NullBool
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(nb.Bool)
	if err != nil {
		return nil, fmt.Errorf("nullBool MarshalJSON >> %w", err)
	}

	return data, nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(ni.NullInt64)
	if err != nil {
		return nil, fmt.Errorf("nullInt64 MarshalJSON >> %w", err)
	}

	return data, nil
}

type NullInt32 struct {
	sql.NullInt32
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	data, err := json.Marshal(ni.NullInt32)
	if err != nil {
		return nil, fmt.Errorf("nullInt64 MarshalJSON >> %w", err)
	}

	return data, nil
}
