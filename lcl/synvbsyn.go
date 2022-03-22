package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynVBSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynVBSyn(obj interface{}) *TSynVBSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynVBSyn{instance: instance, ptr: ptr}
}

func SynVBSynFromInst(inst uintptr) *TSynVBSyn {
	return AsSynVBSyn(inst)
}

func SynVBSynFromObj(obj IObject) *TSynVBSyn {
	return AsSynVBSyn(obj)
}

func SynVBSynFromUnsafePointer(ptr unsafe.Pointer) *TSynVBSyn {
	return AsSynVBSyn(ptr)
}

func (s *TSynVBSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynVBSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynVBSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynVBSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynVBSynClass() TClass {
	return SynVBSyn_StaticClassType()
}

func NewSynVBSyn(owner IComponent) *TSynVBSyn {
	s := new(TSynVBSyn)
	s.instance = SynVBSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynVBSyn) Free() {
	if s.instance != 0 {
		SynVBSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynVBSyn) GetEnabled() bool {
	return SynVBSyn_GetEnabled(s.instance)
}

func (s *TSynVBSyn) SetEnabled(value bool) {
	SynVBSyn_SetEnabled(s.instance, value)
}

func (s *TSynVBSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetCommentAttri(s.instance))
}

func (s *TSynVBSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynVBSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetKeyAttri(s.instance))
}

func (s *TSynVBSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetNumberAttri(s.instance))
}

func (s *TSynVBSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetSpaceAttri(s.instance))
}

func (s *TSynVBSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetStringAttri(s.instance))
}

func (s *TSynVBSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynVBSyn_GetSymbolAttri(s.instance))
}
