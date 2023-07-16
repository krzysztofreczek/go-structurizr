package model_test

import (
	"testing"

	"github.com/marianoceneri/go-structurizr/pkg/model"
	"github.com/stretchr/testify/require"
)

const (
	emptyStructChecksum  = "v1_d751713988987e9331980363e24189ce"
	simpleStructChecksum = "v1_4cd1ab42ba6c8a15bc80d7918b7f3ec7"
)

func TestStructure_Checksum(t *testing.T) {
	tests := []struct {
		name      string
		structure model.Structure
		expected  string
	}{
		{
			name:      "empty",
			structure: model.Structure{},
			expected:  emptyStructChecksum,
		},
		{
			name:      "new structure",
			structure: model.NewStructure(),
			expected:  emptyStructChecksum,
		},
		{
			name:      "simple structure",
			structure: simpleStructure(),
			expected:  simpleStructChecksum,
		},
		{
			name:      "simple structure with different orders",
			structure: simpleStructureWithDifferentOrders(),
			expected:  simpleStructChecksum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.structure.Checksum()
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func simpleStructure() model.Structure {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_1"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_2"},
		},
		"ID_3": {
			ID:          "ID_3",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_1", "TAG_2"},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
			"ID_3": {},
		},
		"ID_2": {
			"ID_3": {},
		},
	}
	return s
}

func simpleStructureWithDifferentOrders() model.Structure {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_3": {
			ID:          "ID_3",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_2", "TAG_1"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_2"},
		},
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"TAG_1"},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_2": {
			"ID_3": {},
		},
		"ID_1": {
			"ID_3": {},
			"ID_2": {},
		},
	}
	return s
}
