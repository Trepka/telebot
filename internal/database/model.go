package database

import (
	"database/sql"
)

// WasteType provides type of waste
type TelebotMessage struct {
	ID  sql.NullInt32  `json:"id"`
	RUS sql.NullString `json:"rus"`
	EN  sql.NullString `json:"en"`
}

// WasteTypeList is a slice of WasteType structs
type MessageList []TelebotMessage
