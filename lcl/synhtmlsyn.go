package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynHtmlSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynHtmlSyn(obj any) *TSynHtmlSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynHtmlSyn{instance: instance, ptr: ptr}
}

func SynHtmlSynFromInst(inst uintptr) *TSynHtmlSyn {
	return AsSynHtmlSyn(inst)
}

func SynHtmlSynFromObj(obj IObject) *TSynHtmlSyn {
	return AsSynHtmlSyn(obj)
}

func SynHtmlSynFromUnsafePointer(ptr unsafe.Pointer) *TSynHtmlSyn {
	return AsSynHtmlSyn(ptr)
}

func (s *TSynHtmlSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynHtmlSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynHtmlSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynHtmlSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynHtmlSynClass() TClass {
	return SynHTMLSyn_StaticClassType()
}

func NewSynHtmlSyn(owner IComponent) *TSynHtmlSyn {
	s := new(TSynHtmlSyn)
	s.instance = SynHTMLSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynHtmlSyn) Free() {
	if s.instance != 0 {
		SynHTMLSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynHtmlSyn) GetEnabled() bool {
	return SynHTMLSyn_GetEnabled(s.instance)
}

func (s *TSynHtmlSyn) SetEnabled(value bool) {
	SynHTMLSyn_SetEnabled(s.instance, value)
}

func (s *TSynHtmlSyn) GetAndAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetAndAttri(s.instance))
}

func (s *TSynHtmlSyn) GetASPAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetASPAttri(s.instance))
}

func (s *TSynHtmlSyn) GetCDATAAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetCDATAAttri(s.instance))
}

func (s *TSynHtmlSyn) GetDOCTYPEAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetDOCTYPEAttri(s.instance))
}

func (s *TSynHtmlSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetCommentAttri(s.instance))
}

func (s *TSynHtmlSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynHtmlSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetKeyAttri(s.instance))
}

func (s *TSynHtmlSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetSpaceAttri(s.instance))
}

func (s *TSynHtmlSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetSymbolAttri(s.instance))
}

func (s *TSynHtmlSyn) GetTextAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetTextAttri(s.instance))
}

func (s *TSynHtmlSyn) GetUndefKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetUndefKeyAttri(s.instance))
}

func (s *TSynHtmlSyn) GetValueAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynHTMLSyn_GetValueAttri(s.instance))
}

func (s *TSynHtmlSyn) GetMode() TSynHTMLSynMode {
	return SynHTMLSyn_GetMode(s.instance)
}

func (s *TSynHtmlSyn) SetMode(value TSynHTMLSynMode) {
	SynHTMLSyn_SetMode(s.instance, value)
}
