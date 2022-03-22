package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynCssSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynCssSyn(obj interface{}) *TSynCssSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynCssSyn{instance: instance, ptr: ptr}
}

func SynCssSynFromInst(inst uintptr) *TSynCssSyn {
	return AsSynCssSyn(inst)
}

func SynCssSynFromObj(obj IObject) *TSynCssSyn {
	return AsSynCssSyn(obj)
}

func SynCssSynFromUnsafePointer(ptr unsafe.Pointer) *TSynCssSyn {
	return AsSynCssSyn(ptr)
}

func (s *TSynCssSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynCssSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynCssSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynCssSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynCssSynClass() TClass {
	return SynCssSyn_StaticClassType()
}

func NewSynCssSyn(owner IComponent) *TSynCssSyn {
	s := new(TSynCssSyn)
	s.instance = SynCssSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynCssSyn) Free() {
	if s.instance != 0 {
		SynCssSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynCssSyn) GetEnabled() bool {
	return SynCssSyn_GetEnabled(s.instance)
}

func (s *TSynCssSyn) SetEnabled(value bool) {
	SynCssSyn_SetEnabled(s.instance, value)
}

func (s *TSynCssSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetCommentAttri(s.instance))
}

func (s *TSynCssSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynCssSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetKeyAttri(s.instance))
}

func (s *TSynCssSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetNumberAttri(s.instance))
}

func (s *TSynCssSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetSpaceAttri(s.instance))
}

func (s *TSynCssSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetStringAttri(s.instance))
}

func (s *TSynCssSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetSymbolAttri(s.instance))
}

func (s *TSynCssSyn) GetMeasurementUnitAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetMeasurementUnitAttri(s.instance))
}

func (s *TSynCssSyn) GetSelectorAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCssSyn_GetSelectorAttri(s.instance))
}
