package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynTeXSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynTeXSyn(obj interface{}) *TSynTeXSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynTeXSyn{instance: instance, ptr: ptr}
}

func SynTeXSynFromInst(inst uintptr) *TSynTeXSyn {
	return AsSynTeXSyn(inst)
}

func SynTeXSynFromObj(obj IObject) *TSynTeXSyn {
	return AsSynTeXSyn(obj)
}

func SynTeXSynFromUnsafePointer(ptr unsafe.Pointer) *TSynTeXSyn {
	return AsSynTeXSyn(ptr)
}

func (s *TSynTeXSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynTeXSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynTeXSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynTeXSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynTeXSynClass() TClass {
	return SynTeXSyn_StaticClassType()
}

func NewSynTeXSyn(owner IComponent) *TSynTeXSyn {
	s := new(TSynTeXSyn)
	s.instance = SynTeXSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynTeXSyn) Free() {
	if s.instance != 0 {
		SynTeXSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynTeXSyn) GetEnabled() bool {
	return SynTeXSyn_GetEnabled(s.instance)
}

func (s *TSynTeXSyn) SetEnabled(value bool) {
	SynTeXSyn_SetEnabled(s.instance, value)
}

func (s *TSynTeXSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetCommentAttri(s.instance))
}

func (s *TSynTeXSyn) GetTextAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetTextAttri(s.instance))
}

func (s *TSynTeXSyn) GetControlSequenceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetControlSequenceAttri(s.instance))
}

func (s *TSynTeXSyn) GetMathmodeAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetMathmodeAttri(s.instance))
}

func (s *TSynTeXSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetSpaceAttri(s.instance))
}

func (s *TSynTeXSyn) GetBraceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetBraceAttri(s.instance))
}

func (s *TSynTeXSyn) GetBracketAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynTeXSyn_GetBracketAttri(s.instance))
}
