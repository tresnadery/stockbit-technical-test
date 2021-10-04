package domain

import "time"

type Log struct {
	ID        int64      `json:"id"`
	Request   string     `json:"request"`
	Response  string     `json:"response"`
	CreatedAt *time.Time `json:"created_at"`
}
