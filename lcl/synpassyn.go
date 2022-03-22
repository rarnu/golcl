package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynPasSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynPasSyn(obj interface{}) *TSynPasSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynPasSyn{instance: instance, ptr: ptr}
}

func SynPasSynFromInst(inst uintptr) *TSynPasSyn {
	return AsSynPasSyn(inst)
}

func SynPasSynFromObj(obj IObject) *TSynPasSyn {
	return AsSynPasSyn(obj)
}

func SynPasSynFromUnsafePointer(ptr unsafe.Pointer) *TSynPasSyn {
	return AsSynPasSyn(ptr)
}

func (s *TSynPasSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynPasSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynPasSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynPasSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynPasSynClass() TClass {
	return SynPasSyn_StaticClassType()
}

func NewSynPasSyn(owner IComponent) *TSynPasSyn {
	s := new(TSynPasSyn)
	s.instance = SynPasSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynPasSyn) Free() {
	if s.instance != 0 {
		SynPasSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynPasSyn) GetEnabled() bool {
	return SynPasSyn_GetEnabled(s.instance)
}

func (s *TSynPasSyn) SetEnabled(value bool) {
	SynPasSyn_SetEnabled(s.instance, value)
}

func (s *TSynPasSyn) GetAsmAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetAsmAttri(s.instance))
}

func (s *TSynPasSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetCommentAttri(s.instance))
}

func (s *TSynPasSyn) GetIDEDirectiveAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetIDEDirectiveAttri(s.instance))
}

func (s *TSynPasSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynPasSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetKeyAttri(s.instance))
}

func (s *TSynPasSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetNumberAttri(s.instance))
}

func (s *TSynPasSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetSpaceAttri(s.instance))
}

func (s *TSynPasSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetStringAttri(s.instance))
}

func (s *TSynPasSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetSymbolAttri(s.instance))
}

func (s *TSynPasSyn) GetProcedureHeaderName() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetProcedureHeaderName(s.instance))
}

func (s *TSynPasSyn) GetCaseLabelAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetCaseLabelAttri(s.instance))
}

func (s *TSynPasSyn) GetDirectiveAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPasSyn_GetDirectiveAttri(s.instance))
}

func (s *TSynPasSyn) GetCompilerMode() TPascalCompilerMode {
	return SynPasSyn_GetCompilerMode(s.instance)
}

func (s *TSynPasSyn) SetCompilerMode(value TPascalCompilerMode) {
	SynPasSyn_SetCompilerMode(s.instance, value)
}

func (s *TSynPasSyn) GetNestedComments() bool {
	return SynPasSyn_GetNestedComments(s.instance)
}

func (s *TSynPasSyn) SetNestedComments(value bool) {
	SynPasSyn_SetNestedComments(s.instance, value)
}

func (s *TSynPasSyn) GetTypeHelpers() bool {
	return SynPasSyn_GetTypeHelpers(s.instance)
}

func (s *TSynPasSyn) SetTypeHelpers(value bool) {
	SynPasSyn_SetTypeHelpers(s.instance, value)
}

func (s *TSynPasSyn) GetD4syntax() bool {
	return SynPasSyn_GetD4syntax(s.instance)
}

func (s *TSynPasSyn) SetD4syntax(value bool) {
	SynPasSyn_SetD4syntax(s.instance, value)
}

func (s *TSynPasSyn) GetExtendedKeywordsMode() bool {
	return SynPasSyn_GetExtendedKeywordsMode(s.instance)
}

func (s *TSynPasSyn) SetExtendedKeywordsMode(value bool) {
	SynPasSyn_SetExtendedKeywordsMode(s.instance, value)
}

func (s *TSynPasSyn) GetStringKeywordMode() TSynPasStringMode {
	return SynPasSyn_GetStringKeywordMode(s.instance)
}

func (s *TSynPasSyn) SetStringKeywordMode(value TSynPasStringMode) {
	SynPasSyn_SetStringKeywordMode(s.instance, value)
}
