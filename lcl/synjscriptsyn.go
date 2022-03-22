package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynJScriptSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynJScriptSyn(obj interface{}) *TSynJScriptSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynJScriptSyn{instance: instance, ptr: ptr}
}

func SynJScriptSynFromInst(inst uintptr) *TSynJScriptSyn {
	return AsSynJScriptSyn(inst)
}

func SynJScriptSynFromObj(obj IObject) *TSynJScriptSyn {
	return AsSynJScriptSyn(obj)
}

func SynJScriptSynFromUnsafePointer(ptr unsafe.Pointer) *TSynJScriptSyn {
	return AsSynJScriptSyn(ptr)
}

func (s *TSynJScriptSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynJScriptSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynJScriptSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynJScriptSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynJScriptSynClass() TClass {
	return SynJScriptSyn_StaticClassType()
}

func NewSynJScriptSyn(owner IComponent) *TSynJScriptSyn {
	s := new(TSynJScriptSyn)
	s.instance = SynJScriptSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynJScriptSyn) Free() {
	if s.instance != 0 {
		SynJScriptSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynJScriptSyn) GetEnabled() bool {
	return SynJScriptSyn_GetEnabled(s.instance)
}

func (s *TSynJScriptSyn) SetEnabled(value bool) {
	SynJScriptSyn_SetEnabled(s.instance, value)
}

func (s *TSynJScriptSyn) GetBracketAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetBracketAttri(s.instance))
}

func (s *TSynJScriptSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetCommentAttri(s.instance))
}

func (s *TSynJScriptSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynJScriptSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetKeyAttri(s.instance))
}

func (s *TSynJScriptSyn) GetNonReservedKeyAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetNonReservedKeyAttri(s.instance))
}

func (s *TSynJScriptSyn) GetEventAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetEventAttri(s.instance))
}

func (s *TSynJScriptSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetNumberAttri(s.instance))
}

func (s *TSynJScriptSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetSpaceAttri(s.instance))
}

func (s *TSynJScriptSyn) GetStringAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetStringAttri(s.instance))
}

func (s *TSynJScriptSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return SynHighlighterAttributesFromInst(SynJScriptSyn_GetSymbolAttri(s.instance))
}
