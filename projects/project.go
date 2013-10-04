package projects

import (
	"github.com/skydb/sky.go"
    "github.com/skylandlabs/landmarkd/core"
)

// A Project is a collection of users and their events. In the landmarkd
// process, a Project simply links API keys to the table the project is
// stored on.
type Project struct {
	ApiKey string
	table  sky.Table
}

// Creates a new Project.
func New(apiKey string, table sky.Table) *Project {
    return &Project{
        ApiKey: apiKey,
        table: table,
    }
}

// Tracks an event.
func (p *Project) Track(user *core.User, device *core.Device, event *core.Event) error {
	// TODO: Create a Sky event.
	// TODO: Save to Sky.
	return nil
}
