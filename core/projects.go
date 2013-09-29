package core

import (
	"sync"
)

type Projects struct {
	sync.Lock
	itemsByApiKey map[string]*Project
}

// Creates a new collection of projects.
func NewProjects() *Projects {
	return &Projects{}
}

// Retrieves a project associated with a given API key.
func (s *Projects) FindByApiKey(key string) (*Project, error) {
	s.Lock()
	defer s.Unlock()

	// Find local reference to project.
	if project := s.itemsByApiKey[key]; project != nil {
		return project
	}

	// TODO: Find project from Redis.

	return nil
}
