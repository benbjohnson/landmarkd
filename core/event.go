package core

import (
    "time"
)

type Event struct {
	Timestamp *time.Time             `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}
