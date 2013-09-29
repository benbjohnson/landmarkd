package core

import (
	"github.com/skydb/sky.go"
)

type Project struct {
	ApiKey string
	table  *sky.Table
}

// Tracks an event.
func (p *Project) Track(user *User, device *Device, event *Event) error {
	// TODO: Create a Sky event.
	// TODO: Save to Sky.
	return nil
}
