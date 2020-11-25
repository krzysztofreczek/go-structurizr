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

// todo: components matching rule
// todo: scraped info from interface
// todo: scraped info from matching rule
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
			name:      "root with private value of public component that implements HasInfo interface",
			structure: test.NewRootWithPrivatePublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private value of private component that implements HasInfo interface",
			structure: test.NewRootWithPrivatePrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public map of pointers to public components",
			structure: test.NewRootWithPublicMapOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public map of pointers to private components",
			structure: test.NewRootWithPublicMapOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public map of public components",
			structure: test.NewRootWithPublicMapOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public map of private components",
			structure: test.NewRootWithPublicMapOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private map of pointers to public components",
			structure: test.NewRootWithPrivateMapOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private map of pointers to private components",
			structure: test.NewRootWithPrivateMapOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private map of public components",
			structure: test.NewRootWithPrivateMapOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private map of private components",
			structure: test.NewRootWithPrivateMapOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public slice of pointers to public components",
			structure: test.NewRootWithPublicSliceOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public slice of pointers to private components",
			structure: test.NewRootWithPublicSliceOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public slice of public components",
			structure: test.NewRootWithPublicSliceOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public slice of private components",
			structure: test.NewRootWithPublicSliceOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private slice of pointers to public components",
			structure: test.NewRootWithPrivateSliceOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private slice of pointers to private components",
			structure: test.NewRootWithPrivateSliceOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private slice of public components",
			structure: test.NewRootWithPrivateSliceOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private slice of private components",
			structure: test.NewRootWithPrivateSliceOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public array of pointers to public components",
			structure: test.NewRootWithPublicArrayOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public array of pointers to private components",
			structure: test.NewRootWithPublicArrayOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public array of public components",
			structure: test.NewRootWithPublicArrayOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public array of private components",
			structure: test.NewRootWithPublicArrayOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private array of pointers to public components",
			structure: test.NewRootWithPrivateArrayOfPointersToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private array of pointers to private components",
			structure: test.NewRootWithPrivateArrayOfPointersToPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private array of public components",
			structure: test.NewRootWithPrivateArrayOfPublicComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private array of private components",
			structure: test.NewRootWithPrivateArrayOfPrivateComponentHasInfoValue(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public public interface implemented with public component",
			structure: test.NewRootWithPublicPublicInterfaceImplementedWithPublicComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public public interface implemented with private component",
			structure: test.NewRootWithPublicPublicInterfaceImplementedWithPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public private interface implemented with public component",
			structure: test.NewRootWithPublicPrivateInterfaceImplementedWithPublicComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public private interface implemented with private component",
			structure: test.NewRootWithPublicPrivateInterfaceImplementedWithPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private public interface implemented with public component",
			structure: test.NewRootWithPrivatePublicInterfaceImplementedWithPublicComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private public interface implemented with private component",
			structure: test.NewRootWithPrivatePublicInterfaceImplementedWithPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private private interface implemented with public component",
			structure: test.NewRootWithPrivatePrivateInterfaceImplementedWithPublicComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private private interface implemented with private component",
			structure: test.NewRootWithPrivatePrivateInterfaceImplementedWithPrivateComponent(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
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
		{
			name:      "root that implements HasInfo interface with nested components that implement HasInfo interface",
			structure: test.NewRootHasInfoWithNestedComponents(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithNestedComponents"):        {},
				componentID("RootHasInfoWithComponentHasInfoPointer"): {},
				componentID("PublicComponentHasInfo"):                 {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithNestedComponents"): {
					componentID("RootHasInfoWithComponentHasInfoPointer"),
				},
				componentID("RootHasInfoWithComponentHasInfoPointer"): {
					componentID("PublicComponentHasInfo"),
				},
			},
		},
		{
			name:      "root that implements HasInfo interface with nested private components that implement HasInfo interface",
			structure: test.NewRootHasInfoWithNestedPrivateComponents(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithNestedPrivateComponents"): {},
				componentID("PublicComponentHasInfo"):                 {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithNestedPrivateComponents"): {
					componentID("PublicComponentHasInfo"),
				},
			},
		},
		{
			name:      "root with public map of HasInfo interfaces",
			structure: test.NewRootWithPublicMapOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private map of HasInfo interfaces",
			structure: test.NewRootWithPrivateMapOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
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
