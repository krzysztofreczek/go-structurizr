package model

type Component struct {
	ID          string
	Kind        string
	Name        string
	Description string
	Technology  string
	Tags        []string
}

type Structure struct {
	Components map[string]Component
	Relations  map[string]map[string]struct{}
}

func NewStructure() Structure {
	return Structure{
		Components: make(map[string]Component),
		Relations:  make(map[string]map[string]struct{}),
	}
}

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
