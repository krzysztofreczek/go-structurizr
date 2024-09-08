package model

import (
	"fmt"
	"sort"

	"github.com/cnf/structhash"
)

const version = 1

// Component is an open structure representing the details of a scraped component.
//
// ID is a unique identifier for the component.
// Kind represents the component's level or type in C4 diagrams.
// Name is the name of the component.
// Description provides a brief explanation of the component's responsibility.
// Technology describes the technology the component is based on.
// Tags is a set of generic strings used to group and reference components.
type Component struct {
	ID          string
	Kind        string
	Name        string
	Description string
	Technology  string
	Tags        []string
}

// Structure is an open structure representing the entire scraped system.
//
// Components contains all the scraped components, indexed by their IDs.
// Relations contains all the connections between components, indexed by their IDs.
type Structure struct {
	Components map[string]Component
	Relations  map[string]map[string]struct{}
}

// NewStructure creates and returns an empty Structure.
func NewStructure() Structure {
	return Structure{
		Components: make(map[string]Component),
		Relations:  make(map[string]map[string]struct{}),
	}
}

// AddComponent adds a component and creates a corresponding relation to its parent.
//
// If a parent with the given ID does not exist, the relation will not be created.
func (s Structure) AddComponent(c Component, parentID string) {
	s.Components[c.ID] = c
	if parentID != "" {
		_, ok := s.Relations[parentID]
		if !ok {
			s.Relations[parentID] = make(map[string]struct{})
		}
		s.Relations[parentID][c.ID] = struct{}{}
	}
}

// Checksum returns a hash of the Structure.
//
// The checksum can be used to track changes between different structures.
func (s Structure) Checksum() (string, error) {
	cIDs := make([]string, 0)
	for id := range s.Components {
		cIDs = append(cIDs, id)
	}

	sort.Strings(cIDs)

	accu := make([]string, 0)
	for _, cID := range cIDs {
		c := s.Components[cID]
		sort.Strings(c.Tags)

		cHash, err := structhash.Hash(c, version)
		if err != nil {
			return "", err
		}
		accu = append(accu, cHash)

		r := s.Relations[cID]
		for _, rID := range cIDs {
			if _, exists := r[rID]; exists {
				rel := fmt.Sprintf("%s-%s", cID, rID)
				accu = append(accu, rel)
			}
		}
	}

	return structhash.Hash(accu, version)
}
