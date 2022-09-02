package test

import "github.com/krzysztofreczek/go-structurizr/pkg/model"

type PublicInterface interface {
	DoSomethingPublic()
}

type privateInterface interface {
	DoSomethingPrivate()
}

type PublicComponent struct{}

type PublicComponentHasInfo struct{}

func (r PublicComponentHasInfo) Info() model.Info {
	return model.ComponentInfo(
		"test.PublicComponentHasInfo",
		"public",
	)
}

func (r PublicComponentHasInfo) DoSomethingPublic() {
	panic("implement me")
}

func (r PublicComponentHasInfo) DoSomethingPrivate() {
	panic("implement me")
}

type privateComponent struct{}

type privateComponentHasInfo struct{}

func (r privateComponentHasInfo) Info() model.Info {
	return model.ComponentInfo(
		"test.privateComponentHasInfo",
		"private",
	)
}

func (r privateComponentHasInfo) DoSomethingPublic() {
	panic("implement me")
}

func (r privateComponentHasInfo) DoSomethingPrivate() {
	panic("implement me")
}

type RootEmpty struct{}

func NewRootEmpty() RootEmpty {
	return RootEmpty{}
}

type RootWithSimpleTypes struct {
	i int
	s string
	f float32
	b bool
}

func NewRootWithSimpleTypes() RootWithSimpleTypes {
	return RootWithSimpleTypes{
		i: 0,
		s: "",
		f: 0,
		b: false,
	}
}

type RootWithCircularDependencies struct {
	nested []RootWithCircularDependencies
}

func NewRootWithCircularDependencies() RootWithCircularDependencies {
	r := RootWithCircularDependencies{}
	r.nested = []RootWithCircularDependencies{r}
	return r
}

type RootHasInfoWithCircularDependencies struct {
	nested []RootHasInfoWithCircularDependencies
}

func NewRootHasInfoWithCircularDependencies() RootHasInfoWithCircularDependencies {
	r := RootHasInfoWithCircularDependencies{}
	r.nested = []RootHasInfoWithCircularDependencies{r}
	return r
}

func (r RootHasInfoWithCircularDependencies) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithCircularDependencies",
	)
}

type RootHasInfoWithCircularPointerDependencies struct {
	nested []*RootHasInfoWithCircularPointerDependencies
}

func NewRootHasInfoWithCircularPointerDependencies() RootHasInfoWithCircularPointerDependencies {
	r := RootHasInfoWithCircularPointerDependencies{}
	r.nested = []*RootHasInfoWithCircularPointerDependencies{&r}
	return r
}

func (r RootHasInfoWithCircularPointerDependencies) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithCircularPointerDependencies",
	)
}

type RootEmptyHasInfo struct{}

func NewRootEmptyHasInfo() RootEmptyHasInfo {
	return *NewRootEmptyHasInfoPtr()
}

func NewRootEmptyHasInfoPtr() *RootEmptyHasInfo {
	return &RootEmptyHasInfo{}
}

func (r RootEmptyHasInfo) Info() model.Info {
	return model.ComponentInfo(
		"test.RootEmptyHasInfo",
		"root description",
		"root technology",
		"root tag 1",
		"root tag 2",
	)
}

type RootEmptyPtrHasInfo struct{}

func NewRootEmptyPtrHasInfo() RootEmptyPtrHasInfo {
	return *NewRootEmptyPtrHasInfoPtr()
}

func NewRootEmptyPtrHasInfoPtr() *RootEmptyPtrHasInfo {
	return &RootEmptyPtrHasInfo{}
}

func (r *RootEmptyPtrHasInfo) Info() model.Info {
	return model.ComponentInfo(
		"test.RootEmptyPtrHasInfo",
		"root description",
		"root technology",
		"root tag 1",
		"root tag 2",
	)
}

type RootWithPublicPointerToPublicComponent struct {
	Ptr *PublicComponent
}

func NewRootWithPublicPointerToPublicComponent() RootWithPublicPointerToPublicComponent {
	return RootWithPublicPointerToPublicComponent{
		Ptr: &PublicComponent{},
	}
}

type RootWithPublicPointerToPrivateComponent struct {
	Ptr *privateComponent
}

func NewRootWithPublicPointerToPrivateComponent() RootWithPublicPointerToPrivateComponent {
	return RootWithPublicPointerToPrivateComponent{
		Ptr: &privateComponent{},
	}
}

type RootWithPublicPointerToPublicComponentHasInfo struct {
	Ptr *PublicComponentHasInfo
}

func NewRootWithPublicPointerToPublicComponentHasInfo() RootWithPublicPointerToPublicComponentHasInfo {
	return RootWithPublicPointerToPublicComponentHasInfo{
		Ptr: &PublicComponentHasInfo{},
	}
}

type RootWithPublicPointerToPrivateComponentHasInfo struct {
	Ptr *privateComponentHasInfo
}

func NewRootWithPublicPointerToPrivateComponentHasInfo() RootWithPublicPointerToPrivateComponentHasInfo {
	return RootWithPublicPointerToPrivateComponentHasInfo{
		Ptr: &privateComponentHasInfo{},
	}
}

type RootWithPrivatePointerToPublicComponent struct {
	ptr *PublicComponent
}

func NewRootWithPrivatePointerToPublicComponent() RootWithPrivatePointerToPublicComponent {
	return RootWithPrivatePointerToPublicComponent{
		ptr: &PublicComponent{},
	}
}

type RootWithPrivatePointerToPrivateComponent struct {
	ptr *privateComponent
}

func NewRootWithPrivatePointerToPrivateComponent() RootWithPrivatePointerToPrivateComponent {
	return RootWithPrivatePointerToPrivateComponent{
		ptr: &privateComponent{},
	}
}

type RootWithPrivatePointerToPublicComponentHasInfo struct {
	ptr *PublicComponentHasInfo
}

func NewRootWithPrivatePointerToPublicComponentHasInfo() RootWithPrivatePointerToPublicComponentHasInfo {
	return RootWithPrivatePointerToPublicComponentHasInfo{
		ptr: &PublicComponentHasInfo{},
	}
}

func NewRootWithNilPrivatePointerToPublicComponentHasInfo() RootWithPrivatePointerToPublicComponentHasInfo {
	return RootWithPrivatePointerToPublicComponentHasInfo{}
}

type RootWithPrivatePointerToPrivateComponentHasInfo struct {
	ptr *privateComponentHasInfo
}

func NewRootWithPrivatePointerToPrivateComponentHasInfo() RootWithPrivatePointerToPrivateComponentHasInfo {
	return RootWithPrivatePointerToPrivateComponentHasInfo{
		ptr: &privateComponentHasInfo{},
	}
}

func NewRootWithNilPrivatePointerToPrivateComponentHasInfo() RootWithPrivatePointerToPrivateComponentHasInfo {
	return RootWithPrivatePointerToPrivateComponentHasInfo{}
}

type RootWithPublicPublicComponentValue struct {
	Value PublicComponent
}

func NewRootWithPublicPublicComponentValue() RootWithPublicPublicComponentValue {
	return RootWithPublicPublicComponentValue{
		Value: PublicComponent{},
	}
}

type RootWithPublicPrivateComponentValue struct {
	Value privateComponent
}

func NewRootWithPublicPrivateComponentValue() RootWithPublicPrivateComponentValue {
	return RootWithPublicPrivateComponentValue{
		Value: privateComponent{},
	}
}

type RootWithPublicPublicComponentHasInfoValue struct {
	Value PublicComponentHasInfo
}

func NewRootWithPublicPublicComponentHasInfoValue() RootWithPublicPublicComponentHasInfoValue {
	return RootWithPublicPublicComponentHasInfoValue{
		Value: PublicComponentHasInfo{},
	}
}

type RootWithPublicPrivateComponentHasInfoValue struct {
	Value privateComponentHasInfo
}

func NewRootWithPublicPrivateComponentHasInfoValue() RootWithPublicPrivateComponentHasInfoValue {
	return RootWithPublicPrivateComponentHasInfoValue{
		Value: privateComponentHasInfo{},
	}
}

type RootWithPrivatePublicComponentValue struct {
	value PublicComponent
}

func NewRootWithPrivatePublicComponentValue() RootWithPrivatePublicComponentValue {
	return RootWithPrivatePublicComponentValue{
		value: PublicComponent{},
	}
}

type RootWithPrivatePrivateComponentValue struct {
	value privateComponent
}

func NewRootWithPrivatePrivateComponentValue() RootWithPrivatePrivateComponentValue {
	return RootWithPrivatePrivateComponentValue{
		value: privateComponent{},
	}
}

type RootWithPrivatePublicComponentHasInfoValue struct {
	value PublicComponentHasInfo
}

func NewRootWithPrivatePublicComponentHasInfoValue() RootWithPrivatePublicComponentHasInfoValue {
	return RootWithPrivatePublicComponentHasInfoValue{
		value: PublicComponentHasInfo{},
	}
}

type RootWithPrivatePrivateComponentHasInfoValue struct {
	value privateComponentHasInfo
}

func NewRootWithPrivatePrivateComponentHasInfoValue() RootWithPrivatePrivateComponentHasInfoValue {
	return RootWithPrivatePrivateComponentHasInfoValue{
		value: privateComponentHasInfo{},
	}
}

type RootWithPublicMapOfPointersToPublicComponentHasInfo struct {
	Pointers map[string]*PublicComponentHasInfo
}

func NewRootWithPublicMapOfPointersToPublicComponentHasInfo() RootWithPublicMapOfPointersToPublicComponentHasInfo {
	return RootWithPublicMapOfPointersToPublicComponentHasInfo{
		Pointers: map[string]*PublicComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPublicMapOfPointersToPrivateComponentHasInfoValue struct {
	Pointers map[string]*privateComponentHasInfo
}

func NewRootWithPublicMapOfPointersToPrivateComponentHasInfoValue() RootWithPublicMapOfPointersToPrivateComponentHasInfoValue {
	return RootWithPublicMapOfPointersToPrivateComponentHasInfoValue{
		Pointers: map[string]*privateComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPublicMapOfPublicComponentHasInfoValue struct {
	Values map[string]PublicComponentHasInfo
}

func NewRootWithPublicMapOfPublicComponentHasInfoValue() RootWithPublicMapOfPublicComponentHasInfoValue {
	return RootWithPublicMapOfPublicComponentHasInfoValue{
		Values: map[string]PublicComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPublicMapOfPrivateComponentHasInfoValue struct {
	Values map[string]privateComponentHasInfo
}

func NewRootWithPublicMapOfPrivateComponentHasInfoValue() RootWithPublicMapOfPrivateComponentHasInfoValue {
	return RootWithPublicMapOfPrivateComponentHasInfoValue{
		Values: map[string]privateComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPrivateMapOfPointersToPublicComponentHasInfo struct {
	pointers map[string]*PublicComponentHasInfo
}

func NewRootWithPrivateMapOfPointersToPublicComponentHasInfo() RootWithPrivateMapOfPointersToPublicComponentHasInfo {
	return RootWithPrivateMapOfPointersToPublicComponentHasInfo{
		pointers: map[string]*PublicComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPrivateMapOfPointersToPrivateComponentHasInfoValue struct {
	pointers map[string]*privateComponentHasInfo
}

func NewRootWithPrivateMapOfPointersToPrivateComponentHasInfoValue() RootWithPrivateMapOfPointersToPrivateComponentHasInfoValue {
	return RootWithPrivateMapOfPointersToPrivateComponentHasInfoValue{
		pointers: map[string]*privateComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPrivateMapOfPublicComponentHasInfoValue struct {
	values map[string]PublicComponentHasInfo
}

func NewRootWithPrivateMapOfPublicComponentHasInfoValue() RootWithPrivateMapOfPublicComponentHasInfoValue {
	return RootWithPrivateMapOfPublicComponentHasInfoValue{
		values: map[string]PublicComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPrivateMapOfPrivateComponentHasInfoValue struct {
	values map[string]privateComponentHasInfo
}

func NewRootWithPrivateMapOfPrivateComponentHasInfoValue() RootWithPrivateMapOfPrivateComponentHasInfoValue {
	return RootWithPrivateMapOfPrivateComponentHasInfoValue{
		values: map[string]privateComponentHasInfo{
			"ID": {},
		},
	}
}

type RootWithPublicSliceOfPointersToPublicComponentHasInfo struct {
	Pointers []*PublicComponentHasInfo
}

func NewRootWithPublicSliceOfPointersToPublicComponentHasInfo() RootWithPublicSliceOfPointersToPublicComponentHasInfo {
	return RootWithPublicSliceOfPointersToPublicComponentHasInfo{
		Pointers: []*PublicComponentHasInfo{{}},
	}
}

type RootWithPublicSliceOfPointersToPrivateComponentHasInfoValue struct {
	Pointers []*privateComponentHasInfo
}

func NewRootWithPublicSliceOfPointersToPrivateComponentHasInfoValue() RootWithPublicSliceOfPointersToPrivateComponentHasInfoValue {
	return RootWithPublicSliceOfPointersToPrivateComponentHasInfoValue{
		Pointers: []*privateComponentHasInfo{{}},
	}
}

type RootWithPublicSliceOfPublicComponentHasInfoValue struct {
	Values []PublicComponentHasInfo
}

func NewRootWithPublicSliceOfPublicComponentHasInfoValue() RootWithPublicSliceOfPublicComponentHasInfoValue {
	return RootWithPublicSliceOfPublicComponentHasInfoValue{
		Values: []PublicComponentHasInfo{{}},
	}
}

type RootWithPublicSliceOfPrivateComponentHasInfoValue struct {
	Values []privateComponentHasInfo
}

func NewRootWithPublicSliceOfPrivateComponentHasInfoValue() RootWithPublicSliceOfPrivateComponentHasInfoValue {
	return RootWithPublicSliceOfPrivateComponentHasInfoValue{
		Values: []privateComponentHasInfo{{}},
	}
}

type RootWithPrivateSliceOfPointersToPublicComponentHasInfo struct {
	pointers []*PublicComponentHasInfo
}

func NewRootWithPrivateSliceOfPointersToPublicComponentHasInfo() RootWithPrivateSliceOfPointersToPublicComponentHasInfo {
	return RootWithPrivateSliceOfPointersToPublicComponentHasInfo{
		pointers: []*PublicComponentHasInfo{{}},
	}
}

type RootWithPrivateSliceOfPointersToPrivateComponentHasInfoValue struct {
	pointers []*privateComponentHasInfo
}

func NewRootWithPrivateSliceOfPointersToPrivateComponentHasInfoValue() RootWithPrivateSliceOfPointersToPrivateComponentHasInfoValue {
	return RootWithPrivateSliceOfPointersToPrivateComponentHasInfoValue{
		pointers: []*privateComponentHasInfo{{}},
	}
}

type RootWithPrivateSliceOfPublicComponentHasInfoValue struct {
	values []PublicComponentHasInfo
}

func NewRootWithPrivateSliceOfPublicComponentHasInfoValue() RootWithPrivateSliceOfPublicComponentHasInfoValue {
	return RootWithPrivateSliceOfPublicComponentHasInfoValue{
		values: []PublicComponentHasInfo{{}},
	}
}

type RootWithPrivateSliceOfPrivateComponentHasInfoValue struct {
	values []privateComponentHasInfo
}

func NewRootWithPrivateSliceOfPrivateComponentHasInfoValue() RootWithPrivateSliceOfPrivateComponentHasInfoValue {
	return RootWithPrivateSliceOfPrivateComponentHasInfoValue{
		values: []privateComponentHasInfo{{}},
	}
}

type RootWithPublicArrayOfPointersToPublicComponentHasInfo struct {
	Pointers [1]*PublicComponentHasInfo
}

func NewRootWithPublicArrayOfPointersToPublicComponentHasInfo() RootWithPublicArrayOfPointersToPublicComponentHasInfo {
	return RootWithPublicArrayOfPointersToPublicComponentHasInfo{
		Pointers: [1]*PublicComponentHasInfo{{}},
	}
}

type RootWithPublicArrayOfPointersToPrivateComponentHasInfoValue struct {
	Pointers [1]*privateComponentHasInfo
}

func NewRootWithPublicArrayOfPointersToPrivateComponentHasInfoValue() RootWithPublicArrayOfPointersToPrivateComponentHasInfoValue {
	return RootWithPublicArrayOfPointersToPrivateComponentHasInfoValue{
		Pointers: [1]*privateComponentHasInfo{{}},
	}
}

type RootWithPublicArrayOfPublicComponentHasInfoValue struct {
	Values [1]PublicComponentHasInfo
}

func NewRootWithPublicArrayOfPublicComponentHasInfoValue() RootWithPublicArrayOfPublicComponentHasInfoValue {
	return RootWithPublicArrayOfPublicComponentHasInfoValue{
		Values: [1]PublicComponentHasInfo{{}},
	}
}

type RootWithPublicArrayOfPrivateComponentHasInfoValue struct {
	Values [1]privateComponentHasInfo
}

func NewRootWithPublicArrayOfPrivateComponentHasInfoValue() RootWithPublicArrayOfPrivateComponentHasInfoValue {
	return RootWithPublicArrayOfPrivateComponentHasInfoValue{
		Values: [1]privateComponentHasInfo{{}},
	}
}

type RootWithPrivateArrayOfPointersToPublicComponentHasInfo struct {
	pointers [1]*PublicComponentHasInfo
}

func NewRootWithPrivateArrayOfPointersToPublicComponentHasInfo() RootWithPrivateArrayOfPointersToPublicComponentHasInfo {
	return RootWithPrivateArrayOfPointersToPublicComponentHasInfo{
		pointers: [1]*PublicComponentHasInfo{{}},
	}
}

type RootWithPrivateArrayOfPointersToPrivateComponentHasInfoValue struct {
	pointers [1]*privateComponentHasInfo
}

func NewRootWithPrivateArrayOfPointersToPrivateComponentHasInfoValue() RootWithPrivateArrayOfPointersToPrivateComponentHasInfoValue {
	return RootWithPrivateArrayOfPointersToPrivateComponentHasInfoValue{
		pointers: [1]*privateComponentHasInfo{{}},
	}
}

type RootWithPrivateArrayOfPublicComponentHasInfoValue struct {
	values [1]PublicComponentHasInfo
}

func NewRootWithPrivateArrayOfPublicComponentHasInfoValue() RootWithPrivateArrayOfPublicComponentHasInfoValue {
	return RootWithPrivateArrayOfPublicComponentHasInfoValue{
		values: [1]PublicComponentHasInfo{{}},
	}
}

type RootWithPrivateArrayOfPrivateComponentHasInfoValue struct {
	values [1]privateComponentHasInfo
}

func NewRootWithPrivateArrayOfPrivateComponentHasInfoValue() RootWithPrivateArrayOfPrivateComponentHasInfoValue {
	return RootWithPrivateArrayOfPrivateComponentHasInfoValue{
		values: [1]privateComponentHasInfo{{}},
	}
}

type RootWithPublicPublicInterface struct {
	I PublicInterface
}

func NewRootWithPublicPublicInterfaceImplementedWithPublicComponent() RootWithPublicPublicInterface {
	return RootWithPublicPublicInterface{
		I: PublicComponentHasInfo{},
	}
}

func NewRootWithPublicPublicInterfaceImplementedWithPrivateComponent() RootWithPublicPublicInterface {
	return RootWithPublicPublicInterface{
		I: privateComponentHasInfo{},
	}
}

type RootWithPublicPrivateInterface struct {
	I privateInterface
}

func NewRootWithPublicPrivateInterfaceImplementedWithPublicComponent() RootWithPublicPrivateInterface {
	return RootWithPublicPrivateInterface{
		I: PublicComponentHasInfo{},
	}
}

func NewRootWithPublicPrivateInterfaceImplementedWithPrivateComponent() RootWithPublicPrivateInterface {
	return RootWithPublicPrivateInterface{
		I: privateComponentHasInfo{},
	}
}

func NewRootWithPublicPrivateInterfaceWithNil() RootWithPublicPrivateInterface {
	return RootWithPublicPrivateInterface{}
}

type RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels struct {
	Value  PublicComponentHasInfo
	Nested RootHasInfoWithComponentHasInfoValue
}

func NewRootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels() RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels {
	return RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels{
		Value:  PublicComponentHasInfo{},
		Nested: NewRootHasInfoWithComponentHasInfoValue(),
	}
}

func (r RootHasInfoWithPublicPublicComponentHasInfoValueAtMultipleLevels) Info() model.Info {
	return model.ComponentInfo(
		"test.RootWithPublicPublicComponentHasInfoValueAtMultipleLevels",
		"root has info",
	)
}

type RootWithPrivatePublicInterface struct {
	i PublicInterface
}

func NewRootWithPrivatePublicInterfaceImplementedWithPublicComponent() RootWithPrivatePublicInterface {
	return RootWithPrivatePublicInterface{
		i: PublicComponentHasInfo{},
	}
}

func NewRootWithPrivatePublicInterfaceImplementedWithPrivateComponent() RootWithPrivatePublicInterface {
	return RootWithPrivatePublicInterface{
		i: privateComponentHasInfo{},
	}
}

func NewRootWithPrivatePublicInterfaceWithNil() RootWithPrivatePublicInterface {
	return RootWithPrivatePublicInterface{}
}

type RootWithPrivatePrivateInterface struct {
	i privateInterface
}

func NewRootWithPrivatePrivateInterfaceImplementedWithPublicComponent() RootWithPrivatePrivateInterface {
	return RootWithPrivatePrivateInterface{
		i: PublicComponentHasInfo{},
	}
}

func NewRootWithPrivatePrivateInterfaceImplementedWithPrivateComponent() RootWithPrivatePrivateInterface {
	return RootWithPrivatePrivateInterface{
		i: privateComponentHasInfo{},
	}
}

type RootHasInfoWithComponentHasInfoValue struct {
	Value PublicComponentHasInfo
}

func NewRootHasInfoWithComponentHasInfoValue() RootHasInfoWithComponentHasInfoValue {
	return RootHasInfoWithComponentHasInfoValue{
		Value: PublicComponentHasInfo{},
	}
}

func (r RootHasInfoWithComponentHasInfoValue) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithComponentHasInfoValue",
		"root has info",
	)
}

type RootHasInfoWithComponentHasInfoPointer struct {
	Ptr *PublicComponentHasInfo
}

func NewRootHasInfoWithComponentHasInfoPointer() RootHasInfoWithComponentHasInfoPointer {
	return RootHasInfoWithComponentHasInfoPointer{
		Ptr: &PublicComponentHasInfo{},
	}
}

func (r RootHasInfoWithComponentHasInfoPointer) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithComponentHasInfoPointer",
		"root has info",
	)
}

type RootHasInfoWithNestedComponents struct {
	SubRoot RootHasInfoWithComponentHasInfoPointer
}

func NewRootHasInfoWithNestedComponents() RootHasInfoWithNestedComponents {
	return RootHasInfoWithNestedComponents{
		SubRoot: NewRootHasInfoWithComponentHasInfoPointer(),
	}
}

func (r RootHasInfoWithNestedComponents) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithNestedComponents",
		"root with nested public components",
	)
}

type RootHasInfoWithNestedPrivateComponents struct {
	SubRoot RootWithPrivatePublicComponentHasInfoValue
}

func NewRootHasInfoWithNestedPrivateComponents() RootHasInfoWithNestedPrivateComponents {
	return RootHasInfoWithNestedPrivateComponents{
		SubRoot: NewRootWithPrivatePublicComponentHasInfoValue(),
	}
}

func (r RootHasInfoWithNestedPrivateComponents) Info() model.Info {
	return model.ComponentInfo(
		"test.RootHasInfoWithNestedPrivateComponents",
		"root with nested private components",
	)
}

type RootWithPublicMapOfHasInfoInterfaces struct {
	Infos map[string]model.HasInfo
}

func NewRootWithPublicMapOfHasInfoInterfaces() RootWithPublicMapOfHasInfoInterfaces {
	return RootWithPublicMapOfHasInfoInterfaces{
		Infos: map[string]model.HasInfo{
			"PUBLIC":  PublicComponentHasInfo{},
			"private": privateComponentHasInfo{},
		},
	}
}

type RootWithPrivateMapOfHasInfoInterfaces struct {
	infos map[string]model.HasInfo
}

func NewRootWithPrivateMapOfHasInfoInterfaces() RootWithPrivateMapOfHasInfoInterfaces {
	return RootWithPrivateMapOfHasInfoInterfaces{
		infos: map[string]model.HasInfo{
			"PUBLIC":  PublicComponentHasInfo{},
			"private": privateComponentHasInfo{},
		},
	}
}

type RootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces struct {
	F func() (PublicComponentHasInfo, privateComponentHasInfo)
}

func NewRootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces() RootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces {
	return RootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces{
		F: func() (PublicComponentHasInfo, privateComponentHasInfo) {
			return PublicComponentHasInfo{}, privateComponentHasInfo{}
		},
	}
}

func NewRootWithPublicFunctionNil() RootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces {
	return RootWithPublicFunctionReturningComponentsImplementingOfHasInfoInterfaces{}
}

type RootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces struct {
	f func() (PublicComponentHasInfo, privateComponentHasInfo)
}

func NewRootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces() RootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces {
	return RootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces{
		f: func() (PublicComponentHasInfo, privateComponentHasInfo) {
			return PublicComponentHasInfo{}, privateComponentHasInfo{}
		},
	}
}

func NewRootWithPrivateFunctionNil() RootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces {
	return RootWithPrivateFunctionReturningComponentsImplementingOfHasInfoInterfaces{}
}

type RootWithPublicFunctionReturningPointersToComponentsImplementingOfHasInfoInterfaces struct {
	F func() (*PublicComponentHasInfo, *privateComponentHasInfo)
}

func NewRootWithPublicFunctionReturningPointersToComponentsImplementingOfHasInfoInterfaces() RootWithPublicFunctionReturningPointersToComponentsImplementingOfHasInfoInterfaces {
	return RootWithPublicFunctionReturningPointersToComponentsImplementingOfHasInfoInterfaces{
		F: func() (*PublicComponentHasInfo, *privateComponentHasInfo) {
			return nil, nil
		},
	}
}

type RootWithPrivateFunctionReturningPointersComponentsImplementingOfHasInfoInterfaces struct {
	f func() (*PublicComponentHasInfo, *privateComponentHasInfo)
}

func NewRootWithPrivateFunctionReturningPointersComponentsImplementingOfHasInfoInterfaces() *RootWithPrivateFunctionReturningPointersComponentsImplementingOfHasInfoInterfaces {
	return &RootWithPrivateFunctionReturningPointersComponentsImplementingOfHasInfoInterfaces{
		f: func() (*PublicComponentHasInfo, *privateComponentHasInfo) {
			return nil, nil
		},
	}
}

type PublicInterfaceImplA struct{}

func (a PublicInterfaceImplA) Info() model.Info {
	return model.ComponentInfo(
		"test.PublicInterfaceImplA",
		"public",
	)
}

func (a PublicInterfaceImplA) DoSomethingPublic() {
	panic("implement me")
}

type PublicInterfaceImplB struct{}

func (b PublicInterfaceImplB) Info() model.Info {
	return model.ComponentInfo(
		"test.PublicInterfaceImplB",
		"public",
	)
}

func (b PublicInterfaceImplB) DoSomethingPublic() {
	panic("implement me")
}

type RootWithPublicInterfaceProperty struct {
	PI PublicInterface
}

func NewRootWithMultipleInterfaceImplementations() []RootWithPublicInterfaceProperty {
	return []RootWithPublicInterfaceProperty{
		{PI: PublicInterfaceImplA{}},
		{PI: PublicInterfaceImplB{}},
	}
}

type RootGenericHasInfo[T any] struct {
	t T
}

func NewRootGenericHasInfo[T any](t T) RootGenericHasInfo[T] {
	return RootGenericHasInfo[T]{t: t}
}

func (r RootGenericHasInfo[T]) Info() model.Info {
	return model.ComponentInfo(
		"test.RootGenericHasInfo",
		"public",
	)
}

type RootEmptyGenericHasInfo[T any] struct{}

func NewRootEmptyGenericHasInfo[T any](_ T) RootEmptyGenericHasInfo[T] {
	return RootEmptyGenericHasInfo[T]{}
}

func (r RootEmptyGenericHasInfo[T]) Info() model.Info {
	return model.ComponentInfo(
		"test.RootEmptyGenericHasInfo",
		"public",
	)
}
