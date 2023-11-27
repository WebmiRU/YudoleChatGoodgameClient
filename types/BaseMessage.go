package types

import "encoding/json"

type Base struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
