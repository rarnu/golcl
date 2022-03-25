package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynJavaSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynJavaSyn(obj any) *TSynJavaSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynJavaSyn{instance: instance, ptr: ptr}
}

func SynJavaSynFromInst(inst uintptr) *TSynJavaSyn {
	return AsSynJavaSyn(inst)
}

func SynJavaSynFromObj(obj IObject) *TSynJavaSyn {
	return AsSynJavaSyn(obj)
}

func SynJavaSynFromUnsafePointer(ptr unsafe.Pointer) *TSynJavaSyn {
	return AsSynJavaSyn(ptr)
}

func (s *TSynJavaSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynJavaSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynJavaSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynJavaSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynJavaSynClass() TClass {
	return SynJavaSyn_StaticClassType()
}

func NewSynJavaSyn(owner IComponent) *TSynJavaSyn {
	s := new(TSynJavaSyn)
	s.instance = SynJavaSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynJavaSyn) Free() {
	if s.instance != 0 {
		SynJavaSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynJavaSyn) GetEnabled() bool {
	return SynJavaSyn_GetEnabled(s.instance)
}

func (s *TSynJavaSyn) SetEnabled(value bool) {
	SynJavaSyn_SetEnabled(s.instance, value)
}

func (s *TSynJavaSyn) GetAnnotationAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetAnnotationAttri(s.instance))
}

func (s *TSynJavaSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetCommentAttri(s.instance))
}

func (s *TSynJavaSyn) GetDocumentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetDocumentAttri(s.instance))
}

func (s *TSynJavaSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynJavaSyn) GetInvalidAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetInvalidAttri(s.instance))
}

func (s *TSynJavaSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetKeyAttri(s.instance))
}

func (s *TSynJavaSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetNumberAttri(s.instance))
}

func (s *TSynJavaSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetSpaceAttri(s.instance))
}

func (s *TSynJavaSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetStringAttri(s.instance))
}

func (s *TSynJavaSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynJavaSyn_GetSymbolAttri(s.instance))
}
