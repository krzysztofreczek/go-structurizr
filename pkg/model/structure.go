package model

// Component is an open structure that represents details of scraped component.
//
// ID is a unique identifier of the component
// Kind is a type that reflects component level in terms of C4 diagrams
// Name is a component name
// Description explains the responsibility of the component
// Technology describes technology that the component is based on
// Tags is a set of generic string tags that may be used as reference to a group of components
type Component struct {
	ID          string
	Kind        string
	Name        string
	Description string
	Technology  string
	Tags        []string
}

// Structure is an open stricture that represents whole scraped structure.
//
// Components contains all the scraped components by its IDs
// Relations contains all the connections between components by its IDs
type Structure struct {
	Components map[string]Component
	Relations  map[string]map[string]struct{}
}

// NewStructure instantiates an empty structure.
func NewStructure() Structure {
	return Structure{
		Components: make(map[string]Component),
		Relations:  make(map[string]map[string]struct{}),
	}
}

// AddComponent adds component and corresponding relation to its parent.
//
// In case a parent of given ID does not exist relation will not be created.
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
