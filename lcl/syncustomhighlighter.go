package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynCustomHighlighter struct {
	IComponent
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynCustomHighlighter(obj any) *TSynCustomHighlighter {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynCustomHighlighter{instance: instance, ptr: ptr}
}

func SynCustomHighlighterFromInst(inst uintptr) *TSynCustomHighlighter {
	return AsSynCustomHighlighter(inst)
}

func SynCustomHighlighterFromObj(obj IObject) *TSynCustomHighlighter {
	return AsSynCustomHighlighter(obj)
}

func SynCustomHighlighterFromUnsafePointer(ptr unsafe.Pointer) *TSynCustomHighlighter {
	return AsSynCustomHighlighter(ptr)
}

func (s *TSynCustomHighlighter) Instance() uintptr {
	return s.instance
}

func (s *TSynCustomHighlighter) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynCustomHighlighter) IsValid() bool {
	return s.instance != 0
}

func (s *TSynCustomHighlighter) Is() TIs {
	return TIs(s.instance)
}

func TSynCustomHighlighterClass() TClass {
	return SynCustomHighlighter_StaticClassType()
}

func (s *TSynCustomHighlighter) GetDefaultFilter() string {
	return SynCustomHighlighter_GetDefaultFilter(s.instance)
}

func (s *TSynCustomHighlighter) SetDefaultFilter(value string) {
	SynCustomHighlighter_SetDefaultFilter(s.instance, value)
}

func (s *TSynCustomHighlighter) GetEnabled() bool {
	return SynCustomHighlighter_GetEnabled(s.instance)
}

func (s *TSynCustomHighlighter) SetEnabled(value bool) {
	SynCustomHighlighter_SetEnabled(s.instance, value)
}

func (s *TSynCustomHighlighter) ClassName() string {
	return SynCustomHighlighter_ClassName(s.instance)
}

func (s *TSynCustomHighlighter) ClassType() TClass {
	return SynCustomHighlighter_ClassType(s.instance)
}

func (s *TSynCustomHighlighter) InstanceSize() int32 {
	return SynCustomHighlighter_InstanceSize(s.instance)
}

func (s *TSynCustomHighlighter) InheritsFrom(aClass TClass) bool {
	return SynCustomHighlighter_InheritsFrom(s.instance, aClass)
}

func (s *TSynCustomHighlighter) Equals(Obj IObject) bool {
	return SynCustomHighlighter_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSynCustomHighlighter) GetHashCode() int32 {
	return SynCustomHighlighter_GetHashCode(s.instance)
}

func (s *TSynCustomHighlighter) FindComponent(AName string) *TComponent {
	return AsComponent(SynCustomHighlighter_FindComponent(s.instance, AName))
}

func (s *TSynCustomHighlighter) GetNamePath() string {
	return SynCustomHighlighter_GetNamePath(s.instance)
}

func (s *TSynCustomHighlighter) HasParent() bool {
	return SynCustomHighlighter_HasParent(s.instance)
}

func (s *TSynCustomHighlighter) Assign(Source IObject) {
	SynCustomHighlighter_Assign(s.instance, CheckPtr(Source))
}

func (s *TSynCustomHighlighter) ComponentCount() int32 {
	return SynCustomHighlighter_ComponentCount(s.instance)
}

func (s *TSynCustomHighlighter) ComponentIndex() int32 {
	return SynCustomHighlighter_ComponentIndex(s.instance)
}

func (s *TSynCustomHighlighter) SetComponentIndex(value int32) {
	SynCustomHighlighter_SetComponentIndex(s.instance, value)
}

func (s *TSynCustomHighlighter) Owner() *TComponent {
	return AsComponent(SynCustomHighlighter_Owner(s.instance))
}

func (s *TSynCustomHighlighter) Components(AIndex int32) *TComponent {
	return AsComponent(SynCustomHighlighter_Components(s.instance, AIndex))
}

func (s *TSynCustomHighlighter) Name() string {
	return SynCustomHighlighter_Name(s.instance)
}

func (s *TSynCustomHighlighter) SetName(value string) {
	SynCustomHighlighter_SetName(s.instance, value)
}

func (s *TSynCustomHighlighter) Tag() int {
	return SynCustomHighlighter_Tag(s.instance)
}

func (s *TSynCustomHighlighter) SetTag(value int) {
	SynCustomHighlighter_SetTag(s.instance, value)
}
