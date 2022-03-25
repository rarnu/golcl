package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynAutoComplete struct {
	IComponent
	instance uintptr
	ptr      unsafe.Pointer
}

func NewSynAutoComplete(owner IComponent) *TSynAutoComplete {
	s := new(TSynAutoComplete)
	s.instance = SynAutoComplete_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func AsSynAutoComplete(obj any) *TSynAutoComplete {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynAutoComplete{instance: instance, ptr: ptr}
}

func SynAutoCompleteFromInst(inst uintptr) *TSynAutoComplete {
	return AsSynAutoComplete(inst)
}

func SynAutoCompleteFromObj(obj IObject) *TSynAutoComplete {
	return AsSynAutoComplete(obj)
}

func SynAutoCompleteFromUnsafePointer(ptr unsafe.Pointer) *TSynAutoComplete {
	return AsSynAutoComplete(ptr)
}

func (s *TSynAutoComplete) Free() {
	if s.instance != 0 {
		SynAutoComplete_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynAutoComplete) Instance() uintptr {
	return s.instance
}

func (s *TSynAutoComplete) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynAutoComplete) IsValid() bool {
	return s.instance != 0
}

func (s *TSynAutoComplete) Is() TIs {
	return TIs(s.instance)
}

func TSynAutoCompleteClass() TClass {
	return SynAutoComplete_StaticClassType()
}

func (s *TSynAutoComplete) ClassName() string {
	return SynAutoComplete_ClassName(s.instance)
}

func (s *TSynAutoComplete) ClassType() TClass {
	return SynAutoComplete_ClassType(s.instance)
}

func (s *TSynAutoComplete) InstanceSize() int32 {
	return SynAutoComplete_InstanceSize(s.instance)
}

func (s *TSynAutoComplete) InheritsFrom(aClass TClass) bool {
	return SynAutoComplete_InheritsFrom(s.instance, aClass)
}

func (s *TSynAutoComplete) Equals(Obj IObject) bool {
	return SynAutoComplete_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSynAutoComplete) GetHashCode() int32 {
	return SynAutoComplete_GetHashCode(s.instance)
}

func (s *TSynAutoComplete) FindComponent(AName string) *TComponent {
	return AsComponent(SynAutoComplete_FindComponent(s.instance, AName))
}

func (s *TSynAutoComplete) GetNamePath() string {
	return SynAutoComplete_GetNamePath(s.instance)
}

func (s *TSynAutoComplete) HasParent() bool {
	return SynAutoComplete_HasParent(s.instance)
}

func (s *TSynAutoComplete) Assign(Source IObject) {
	SynAutoComplete_Assign(s.instance, CheckPtr(Source))
}

func (s *TSynAutoComplete) ComponentCount() int32 {
	return SynAutoComplete_ComponentCount(s.instance)
}

func (s *TSynAutoComplete) ComponentIndex() int32 {
	return SynAutoComplete_ComponentIndex(s.instance)
}

func (s *TSynAutoComplete) SetComponentIndex(value int32) {
	SynAutoComplete_SetComponentIndex(s.instance, value)
}

func (s *TSynAutoComplete) Owner() *TComponent {
	return AsComponent(SynAutoComplete_Owner(s.instance))
}

func (s *TSynAutoComplete) Components(AIndex int32) *TComponent {
	return AsComponent(SynAutoComplete_Components(s.instance, AIndex))
}

func (s *TSynAutoComplete) GetAutoCompleteList() *TStrings {
	return AsStrings(SynAutoComplete_GetAutoCompleteList(s.instance))
}

func (s *TSynAutoComplete) SetAutoCompleteList(value *IStrings) {
	SynAutoComplete_SetAutoCompleteList(s.instance, CheckPtr(value))
}

func (s *TSynAutoComplete) GetEditor() *TSynEdit {
	return AsSynEdit(SynAutoComplete_GetEditor(s.instance))
}

func (s *TSynAutoComplete) SetEditor(value *TSynEdit) {
	SynAutoComplete_SetEditor(s.instance, CheckPtr(value))
}

func (s *TSynAutoComplete) GetEndOfTokenChr() string {
	return SynAutoComplete_GetEndOfTokenChr(s.instance)
}

func (s *TSynAutoComplete) SetEndOfTokenChr(value string) {
	SynAutoComplete_SetEndOfTokenChr(s.instance, value)
}

func (s *TSynAutoComplete) GetExecCommandID() int32 {
	return SynAutoComplete_GetExecCommandID(s.instance)
}

func (s *TSynAutoComplete) SetExecCommandID(value int32) {
	SynAutoComplete_SetExecCommandID(s.instance, value)
}

func (s *TSynAutoComplete) Name() string {
	return SynAutoComplete_GetName(s.instance)
}

func (s *TSynAutoComplete) SetName(value string) {
	SynAutoComplete_SetName(s.instance, value)
}

func (s *TSynAutoComplete) GetShortCut() TShortCut {
	return SynAutoComplete_GetShortCut(s.instance)
}

func (s *TSynAutoComplete) SetShortCut(value TShortCut) {
	SynAutoComplete_SetShortCut(s.instance, value)
}

func (s *TSynAutoComplete) Tag() int {
	return SynAutoComplete_GetTag(s.instance)
}

func (s *TSynAutoComplete) SetTag(value int) {
	SynAutoComplete_SetTag(s.instance, value)
}
