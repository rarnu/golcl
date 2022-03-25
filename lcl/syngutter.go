package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynGutter struct {
	IObject
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynGutter(obj any) *TSynGutter {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynGutter{instance: instance, ptr: ptr}
}

func SynGutterFromInst(inst uintptr) *TSynGutter {
	return AsSynGutter(inst)
}

func SynGutterFromObj(obj IObject) *TSynGutter {
	return AsSynGutter(obj)
}

func SynGutterFromUnsafePointer(ptr unsafe.Pointer) *TSynGutter {
	return AsSynGutter(ptr)
}

func (s *TSynGutter) Instance() uintptr {
	return s.instance
}

func (s *TSynGutter) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynGutter) IsValid() bool {
	return s.instance != 0
}

func (s *TSynGutter) Is() TIs {
	return TIs(s.instance)
}

func (s *TSynGutter) Free() {
	if s.instance != 0 {
		SynGutter_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynGutter) ClassName() string {
	return SynGutter_ClassName(s.instance)
}

func (s *TSynGutter) GetHashCode() int32 {
	return SynGutter_GetHashCode(s.instance)
}

func (s *TSynGutter) Equals(obj IObject) bool {
	return SynGutter_Equals(s.instance, CheckPtr(obj))
}

func (s *TSynGutter) ClassType() TClass {
	return SynGutter_ClassType(s.instance)
}

func (s *TSynGutter) InstanceSize() int32 {
	return SynGutter_InstanceSize(s.instance)
}

func (s *TSynGutter) InheritsFrom(aClass TClass) bool {
	return SynGutter_InheritsFrom(s.instance, aClass)
}

func TSynGutterClass() TClass {
	return SynGutter_StaticClassType()
}

func (s *TSynGutter) GetAutoSize() bool {
	return SynGutter_GetAutoSize(s.instance)
}

func (s *TSynGutter) SetAutoSize(value bool) {
	SynGutter_SetAutoSize(s.instance, value)
}

func (s *TSynGutter) GetColor() TColor {
	return SynGutter_GetColor(s.instance)
}

func (s *TSynGutter) SetColor(value TColor) {
	SynGutter_SetColor(s.instance, value)
}

func (s *TSynGutter) GetVisible() bool {
	return SynGutter_GetVisible(s.instance)
}

func (s *TSynGutter) SetVisible(value bool) {
	SynGutter_SetVisible(s.instance, value)
}

func (s *TSynGutter) GetWidth() int32 {
	return SynGutter_GetWidth(s.instance)
}

func (s *TSynGutter) SetWidth(value int32) {
	SynGutter_SetWidth(s.instance, value)
}
