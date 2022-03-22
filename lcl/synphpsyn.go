package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynPHPSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynPHPSyn(obj interface{}) *TSynPHPSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynPHPSyn{instance: instance, ptr: ptr}
}

func SynPHPSynFromInst(inst uintptr) *TSynPHPSyn {
	return AsSynPHPSyn(inst)
}

func SynPHPSynFromObj(obj IObject) *TSynPHPSyn {
	return AsSynPHPSyn(obj)
}

func SynPHPSynFromUnsafePointer(ptr unsafe.Pointer) *TSynPHPSyn {
	return AsSynPHPSyn(ptr)
}

func (s *TSynPHPSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynPHPSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynPHPSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynPHPSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynPHPSynClass() TClass {
	return SynPHPSyn_StaticClassType()
}

func NewSynPHPSyn(owner IComponent) *TSynPHPSyn {
	s := new(TSynPHPSyn)
	s.instance = SynPHPSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynPHPSyn) Free() {
	if s.instance != 0 {
		SynPHPSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynPHPSyn) GetEnabled() bool {
	return SynPHPSyn_GetEnabled(s.instance)
}

func (s *TSynPHPSyn) SetEnabled(value bool) {
	SynPHPSyn_SetEnabled(s.instance, value)
}

func (s *TSynPHPSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetCommentAttri(s.instance))
}

func (s *TSynPHPSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynPHPSyn) GetInvalidSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetInvalidSymbolAttri(s.instance))
}

func (s *TSynPHPSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetKeyAttri(s.instance))
}

func (s *TSynPHPSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetNumberAttri(s.instance))
}

func (s *TSynPHPSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetSpaceAttri(s.instance))
}

func (s *TSynPHPSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetStringAttri(s.instance))
}

func (s *TSynPHPSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetSymbolAttri(s.instance))
}

func (s *TSynPHPSyn) GetVariableAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPHPSyn_GetVariableAttri(s.instance))
}
