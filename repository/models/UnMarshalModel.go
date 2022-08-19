package models

import "encoding/json"

// UnmarshalModel ...
type UnmarshalModel struct {
	To interface{} // can be anything
}

// Scan implement sql.Scanner interface
func (a *UnmarshalModel) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(b, &a.To)
}
