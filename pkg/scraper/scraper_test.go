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

func TestScraper_Scrape_package_matching(t *testing.T) {
	var tests = []struct {
		name                       string
		structure                  interface{}
		packages                   []string
		expectedNumberOfComponents int
	}{
		{
			name:      "structure within given package",
			structure: test.NewRootEmptyHasInfo(),
			packages: []string{
				"github.com/krzysztofreczek/go-structurizr/pkg/internal/test",
			},
			expectedNumberOfComponents: 1,
		},
		{
			name:      "structure out of given package",
			structure: test.NewRootEmptyHasInfo(),
			packages: []string{
				"github.com/krzysztofreczek/go-structurizr/pkg/foo",
			},
			expectedNumberOfComponents: 0,
		},
		{
			name:      "structure within one of given packages",
			structure: test.NewRootEmptyHasInfo(),
			packages: []string{
				"github.com/krzysztofreczek/go-structurizr/pkg/foo",
				"github.com/krzysztofreczek/go-structurizr/pkg/internal/test",
			},
			expectedNumberOfComponents: 1,
		},
		{
			name:      "structure within given package prefix",
			structure: test.NewRootEmptyHasInfo(),
			packages: []string{
				"github.com/krzysztofreczek/go-structurizr/pkg",
			},
			expectedNumberOfComponents: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := scraper.NewConfiguration(tt.packages...)
			s := scraper.NewScraper(c)
			result := s.Scrape(tt.structure)
			require.Len(t, result.Components, tt.expectedNumberOfComponents)
		})
	}
}

func TestScraper_Scrape_has_info_interface(t *testing.T) {
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
			name:                 "empty root with simple types",
			structure:            test.NewRootWithSimpleTypes(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:                 "root with circular dependencies",
			structure:            test.NewRootWithCircularDependencies(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "root has info with circular dependencies",
			structure: test.NewRootHasInfoWithCircularDependencies(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithCircularDependencies"): {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithCircularDependencies"): {
					componentID("RootHasInfoWithCircularDependencies"),
				},
			},
		},
		{
			name:      "root has info with circular pointer dependencies",
			structure: test.NewRootHasInfoWithCircularPointerDependencies(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithCircularPointerDependencies"): {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithCircularPointerDependencies"): {
					componentID("RootHasInfoWithCircularPointerDependencies"),
				},
			},
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
			name:                 "root with public private interface with nil",
			structure:            test.NewRootWithPublicPrivateInterfaceWithNil(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
		},
		{
			name:      "root with public value of public component that implements HasInfo interface at multiple levels",
			structure: test.NewRootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels(),
			expectedComponentIDs: map[string]struct{}{
				componentID("RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels"): {},
				componentID("RootHasInfoWithComponentHasInfoValue"):                             {},
				componentID("PublicComponentHasInfo"):                                           {},
			},
			expectedRelations: map[string][]string{
				componentID("RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels"): {
					componentID("RootHasInfoWithComponentHasInfoValue"),
					componentID("PublicComponentHasInfo"),
				},
				componentID("RootHasInfoWithComponentHasInfoValue"): {
					componentID("PublicComponentHasInfo"),
				},
			},
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
			name:                 "root with private public interface with nil",
			structure:            test.NewRootWithPrivatePublicInterfaceWithNil(),
			expectedComponentIDs: map[string]struct{}{},
			expectedRelations:    map[string][]string{},
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
		{
			name:      "root with nil private pointer to public component that implements HasInfo interface",
			structure: test.NewRootWithNilPrivatePointerToPublicComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with nil private pointer to private component that implements HasInfo interface",
			structure: test.NewRootWithNilPrivatePointerToPrivateComponentHasInfo(),
			expectedComponentIDs: map[string]struct{}{
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public function returning components that implement HasInfo interface",
			structure: test.NewRootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private function returning components that implement HasInfo interface",
			structure: test.NewRootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public function returning pointers to components that implement HasInfo interface",
			structure: test.NewRootWithPublicFunctionReturningPointersToComponentsImplementingOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private function returning pointers to components that implement HasInfo interface",
			structure: test.NewRootWithPrivateFunctionReturningPointersComponentsImplementingOfHasInfoInterfaces(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with public function nil",
			structure: test.NewRootWithPublicFunctionNil(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with private function nil",
			structure: test.NewRootWithPrivateFunctionNil(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicComponentHasInfo"):  {},
				componentID("privateComponentHasInfo"): {},
			},
			expectedRelations: map[string][]string{},
		},
		{
			name:      "root with multiple interface implementation",
			structure: test.NewRootWithMultipleInterfaceImplementations(),
			expectedComponentIDs: map[string]struct{}{
				componentID("PublicInterfaceImplA"): {},
				componentID("PublicInterfaceImplB"): {},
			},
			expectedRelations: map[string][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scraper.NewScraper(c)
			result := s.Scrape(tt.structure)
			requireEqualComponentIDs(t, tt.expectedComponentIDs, result.Components)
			requireEqualRelations(t, tt.expectedRelations, result.Relations)
		})
	}
}

func TestScraper_Scrape_has_info_interface_component_info(t *testing.T) {
	c := scraper.NewConfiguration(
		testPKG,
	)
	var tests = []struct {
		name               string
		structure          interface{}
		expectedComponents map[string]model.Component
	}{
		{
			name:      "pointer to empty root that implements HasInfo interface",
			structure: test.NewRootEmptyHasInfoPtr(),
			expectedComponents: map[string]model.Component{
				componentID("RootEmptyHasInfo"): {
					ID:          componentID("RootEmptyHasInfo"),
					Kind:        "component",
					Name:        "test.RootEmptyHasInfo",
					Description: "root description",
					Technology:  "root technology",
					Tags:        []string{"root tag 1", "root tag 2"},
				},
			},
		},
		{
			name:      "pointer to empty root that pointer implements HasInfo interface",
			structure: test.NewRootEmptyPtrHasInfoPtr(),
			expectedComponents: map[string]model.Component{
				componentID("RootEmptyPtrHasInfo"): {
					ID:          componentID("RootEmptyPtrHasInfo"),
					Kind:        "component",
					Name:        "test.RootEmptyPtrHasInfo",
					Description: "root description",
					Technology:  "root technology",
					Tags:        []string{"root tag 1", "root tag 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scraper.NewScraper(c)
			result := s.Scrape(tt.structure)
			requireEqualComponents(t, tt.expectedComponents, result.Components)
		})
	}
}

func TestScraper_Scrape_rules(t *testing.T) {
	c := scraper.NewConfiguration(
		testPKG,
	)

	ruleDefaultMatchAll, err := scraper.NewRule().
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(
				name,
				"match-all description",
				"match-all technology",
				"match-all tag 1",
				"match-all tag 2",
			)
		}).
		Build()
	require.NoError(t, err)

	ruleMatchPublicComponent, err := scraper.NewRule().
		WithNameRegexp("^test.PublicComponent$").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(
				name,
				"match-pc description",
				"match-pc technology",
				"match-pc tag 1",
				"match-pc tag 2",
			)
		}).
		Build()
	require.NoError(t, err)

	ruleMatchPublicComponentInAnotherPackage, err := scraper.NewRule().
		WithPkgRegexps("^github.com/krzysztofreczek/go-structurizr/pkg/foo$").
		WithNameRegexp("^test.PublicComponent$").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo()
		}).
		Build()
	require.NoError(t, err)

	ruleMatchPublicComponentWithNameAlias, err := scraper.NewRule().
		WithNameRegexp("^test.PublicComponent$").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			n := fmt.Sprintf("%sAlias", name)
			return model.ComponentInfo(n)
		}).
		Build()
	require.NoError(t, err)

	ruleMatchComponentWithNameGroups, err := scraper.NewRule().
		WithNameRegexp(`test\.(\w*)With(\w*)To(\w*)`).
		WithApplyFunc(func(name string, groups ...string) model.Info {
			n := fmt.Sprintf("test.%sWith%sTo%s", groups[0], groups[1], groups[2])
			return model.ComponentInfo(n)
		}).
		Build()
	require.NoError(t, err)

	var tests = []struct {
		name               string
		structure          interface{}
		rules              []scraper.Rule
		expectedComponents map[string]model.Component
	}{
		{
			name:               "no rules",
			structure:          test.NewRootEmpty(),
			rules:              make([]scraper.Rule, 0),
			expectedComponents: map[string]model.Component{},
		},
		{
			name:      "default match-all rule",
			structure: test.NewRootWithPublicPointerToPublicComponent(),
			rules:     []scraper.Rule{ruleDefaultMatchAll},
			expectedComponents: map[string]model.Component{
				componentID("RootWithPublicPointerToPublicComponent"): {
					ID:          componentID("RootWithPublicPointerToPublicComponent"),
					Kind:        "component",
					Name:        "test.RootWithPublicPointerToPublicComponent",
					Description: "match-all description",
					Technology:  "match-all technology",
					Tags:        []string{"match-all tag 1", "match-all tag 2"},
				},
				componentID("PublicComponent"): {
					ID:          componentID("PublicComponent"),
					Kind:        "component",
					Name:        "test.PublicComponent",
					Description: "match-all description",
					Technology:  "match-all technology",
					Tags:        []string{"match-all tag 1", "match-all tag 2"},
				},
			},
		},
		{
			name:      "default match-all rule with nil interfaces",
			structure: test.NewRootWithPrivatePublicInterfaceWithNil(),
			rules:     []scraper.Rule{ruleDefaultMatchAll},
			expectedComponents: map[string]model.Component{
				componentID("RootWithPrivatePublicInterface"): {
					ID:          componentID("RootWithPrivatePublicInterface"),
					Kind:        "component",
					Name:        "test.RootWithPrivatePublicInterface",
					Description: "match-all description",
					Technology:  "match-all technology",
					Tags:        []string{"match-all tag 1", "match-all tag 2"},
				},
				componentID("PublicInterface"): {
					ID:          componentID("PublicInterface"),
					Kind:        "component",
					Name:        "test.PublicInterface",
					Description: "match-all description",
					Technology:  "match-all technology",
					Tags:        []string{"match-all tag 1", "match-all tag 2"},
				},
			},
		},
		{
			name:      "match-public-component rule",
			structure: test.NewRootWithPublicPointerToPublicComponent(),
			rules:     []scraper.Rule{ruleMatchPublicComponent},
			expectedComponents: map[string]model.Component{
				componentID("PublicComponent"): {
					ID:          componentID("PublicComponent"),
					Kind:        "component",
					Name:        "test.PublicComponent",
					Description: "match-pc description",
					Technology:  "match-pc technology",
					Tags:        []string{"match-pc tag 1", "match-pc tag 2"},
				},
			},
		},
		{
			name:               "match-public-component-in-another-package rule",
			structure:          test.NewRootWithPublicPointerToPublicComponent(),
			rules:              []scraper.Rule{ruleMatchPublicComponentInAnotherPackage},
			expectedComponents: map[string]model.Component{},
		},
		{
			name:      "name with alias rule",
			structure: test.NewRootWithPublicPointerToPublicComponent(),
			rules:     []scraper.Rule{ruleMatchPublicComponentWithNameAlias},
			expectedComponents: map[string]model.Component{
				componentID("PublicComponent"): {
					ID:   componentID("PublicComponent"),
					Kind: "component",
					Name: "test.PublicComponentAlias",
					Tags: []string{},
				},
			},
		},
		{
			name:      "name recreated from groups rule",
			structure: test.NewRootWithPublicPointerToPublicComponent(),
			rules:     []scraper.Rule{ruleMatchComponentWithNameGroups},
			expectedComponents: map[string]model.Component{
				componentID("RootWithPublicPointerToPublicComponent"): {
					ID:   componentID("RootWithPublicPointerToPublicComponent"),
					Kind: "component",
					Name: "test.RootWithPublicPointerToPublicComponent",
					Tags: []string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := scraper.NewScraper(c)
			for _, r := range tt.rules {
				err := s.RegisterRule(r)
				require.NoError(t, err)
			}

			result := s.Scrape(tt.structure)
			requireEqualComponents(t, tt.expectedComponents, result.Components)
		})
	}
}

func requireEqualComponentIDs(
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

func requireEqualComponents(
	t *testing.T,
	expectedComponents map[string]model.Component,
	actualComponents map[string]model.Component,
) {
	require.Len(t, actualComponents, len(expectedComponents))
	for id, expectedComponent := range expectedComponents {
		actualComponent, contains := actualComponents[id]
		require.True(t, contains, "actual components: %+v'; expected components: %+v", actualComponents, expectedComponents)
		require.Equal(t, expectedComponent, actualComponent)
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
