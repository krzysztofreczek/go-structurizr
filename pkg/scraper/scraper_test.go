package scraper_test

import (
	"fmt"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/internal"
	"github.com/krzysztofreczek/go-structurizr/pkg/internal/test"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/stretchr/testify/require"
)

const (
	testPKG = "github.com/krzysztofreczek/go-structurizr/pkg/internal/test"
)

// todo: different code structures
// todo: components matching rule
// todo: component that implements interface

// todo: scraped info from matching rule
// todo: scraped info from implements interface

// todo: package matching

func TestScraper_Scrap(t *testing.T) {
	c := scraper.NewConfiguration(
		testPKG,
	)
	var tests = []struct {
		name                 string
		structure            interface{}
		expectedComponentIDs map[string]struct{}
		expectedRelations    map[string][]string
	}{
		{
			name:                 "empty root",
			structure:            test.NewRootEmpty(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "empty root that implements HasInfo interface",
			structure: test.NewRootEmptyHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootEmptyHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "pointer to empty root that implements HasInfo interface",
			structure: test.NewRootEmptyHasInfoPtr(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootEmptyHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:                 "empty root that pointer implements HasInfo interface",
			structure:            test.NewRootEmptyPtrHasInfo(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "pointer to empty root that pointer implements HasInfo interface",
			structure: test.NewRootEmptyPtrHasInfoPtr(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootEmptyPtrHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:                 "root with public pointer to public component",
			structure:            test.NewRootWithPublicPointerToPublicComponent(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:                 "root with public pointer to private component",
			structure:            test.NewRootWithPublicPointerToPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "root with public pointer to public component that implements HasInfo interface",
			structure: test.NewRootWithPublicPointerToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public pointer to private component that implements HasInfo interface",
			structure: test.NewRootWithPublicPointerToPrivateComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:                 "root with private pointer to public component",
			structure:            test.NewRootWithPrivatePointerToPublicComponent(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:                 "root with private pointer to private component",
			structure:            test.NewRootWithPrivatePointerToPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "root with private pointer to public component that implements HasInfo interface",
			structure: test.NewRootWithPrivatePointerToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private pointer to private component that implements HasInfo interface",
			structure: test.NewRootWithPrivatePointerToPrivateComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:                 "root with public value of public component",
			structure:            test.NewRootWithPublicPublicComponentValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:                 "root with public value of private component",
			structure:            test.NewRootWithPublicPrivateComponentValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "root with public value of public component that implements HasInfo interface",
			structure: test.NewRootWithPublicPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public value of private component that implements HasInfo interface",
			structure: test.NewRootWithPublicPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:                 "root with private value of public component",
			structure:            test.NewRootWithPrivatePublicComponentValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:                 "root with private value of private component",
			structure:            test.NewRootWithPrivatePrivateComponentValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			// this is limitation of golang reflection capabilities
			name:                 "root with private value of public component that implements HasInfo interface",
			structure:            test.NewRootWithPrivatePublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			// this is limitation of golang reflection capabilities
			name:                 "root with private value of private component that implements HasInfo interface",
			structure:            test.NewRootWithPrivatePrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},

		{
			name:      "root that implements HasInfo interface with value component that implements HasInfo interface",
			structure: test.NewRootHasInfoWithComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithComponentHasInfoValue"): {},
				componentID("PublicComponentHasInfo"):               {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithComponentHasInfoValue"): {
					componentID("PublicComponentHasInfo"),
				},
			},
		},
		{
			name:      "root that implements HasInfo interface with pointer to component that implements HasInfo interface",
			structure: test.NewRootHasInfoWithComponentHasInfoPointer(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithComponentHasInfoPointer"): {},
				componentID("PublicComponentHasInfo"):                 {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithComponentHasInfoPointer"): {
					componentID("PublicComponentHasInfo"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scraper.NewScraper(c)
			result := s.Scrap(tt.structure)
			requireEqualComponents(t, tt.expectedComponentIDs, result.Components)
			requireEqualRelations(t, tt.expectedRelations, result.Relations)
		})
	}
}

func requireEqualComponents(
	t *testing.T,
	expectedComponentIDs map[string]struct{},
	actualComponents map[string]model.Component,
) {
	require.Len(t, actualComponents, len(expectedComponentIDs))
	for id := range expectedComponentIDs {
		_, contains := actualComponents[id]
		require.True(t, contains, "actual components: %+v", actualComponents)
	}
}

func requireEqualRelations(
	t *testing.T,
	expectedRelations map[string][]string,
	actualRelations map[string]map[string]struct{},
) {
	require.Len(t, actualRelations, len(expectedRelations))
	for id, expectedRelationIDs := range expectedRelations {
		actualRelationIDs, contains := actualRelations[id]
		require.True(t, contains, "actual relations: %+v", actualRelations)
		require.Len(t, actualRelationIDs, len(expectedRelationIDs))
		for _, id = range expectedRelationIDs {
			_, contains := actualRelationIDs[id]
			require.True(t, contains, "actual relation IDs: %+v", actualRelationIDs)
		}
	}
}

func componentID(name string) string {
	id := fmt.Sprintf("%s.%s", testPKG, name)
	return internal.Hash(id)
}
