package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynUNIXShellScriptSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynUNIXShellScriptSyn(obj any) *TSynUNIXShellScriptSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynUNIXShellScriptSyn{instance: instance, ptr: ptr}
}

func SynUNIXShellScriptSynFromInst(inst uintptr) *TSynUNIXShellScriptSyn {
	return AsSynUNIXShellScriptSyn(inst)
}

func SynUNIXShellScriptSynFromObj(obj IObject) *TSynUNIXShellScriptSyn {
	return AsSynUNIXShellScriptSyn(obj)
}

func SynUNIXShellScriptSynFromUnsafePointer(ptr unsafe.Pointer) *TSynUNIXShellScriptSyn {
	return AsSynUNIXShellScriptSyn(ptr)
}

func (s *TSynUNIXShellScriptSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynUNIXShellScriptSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynUNIXShellScriptSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynUNIXShellScriptSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynUNIXShellScriptSynClass() TClass {
	return SynUNIXShellScriptSyn_StaticClassType()
}

func NewSynUNIXShellScriptSyn(owner IComponent) *TSynUNIXShellScriptSyn {
	s := new(TSynUNIXShellScriptSyn)
	s.instance = SynUNIXShellScriptSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynUNIXShellScriptSyn) Free() {
	if s.instance != 0 {
		SynUNIXShellScriptSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynUNIXShellScriptSyn) GetEnabled() bool {
	return SynUNIXShellScriptSyn_GetEnabled(s.instance)
}

func (s *TSynUNIXShellScriptSyn) SetEnabled(value bool) {
	SynUNIXShellScriptSyn_SetEnabled(s.instance, value)
}

func (s *TSynUNIXShellScriptSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetCommentAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetKeyAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetSecondKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetSecondKeyAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetNumberAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetSpaceAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetStringAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetSymbolAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetVarAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynUNIXShellScriptSyn_GetVarAttri(s.instance))
}

func (s *TSynUNIXShellScriptSyn) GetSecondKeyWords() *TStrings {
	return AsStrings(SynUNIXShellScriptSyn_GetSecondKeyWords(s.instance))
}

func (s *TSynUNIXShellScriptSyn) SetSecondKeyWords(value *IStrings) {
	SynUNIXShellScriptSyn_SetSecondKeyWords(s.instance, CheckPtr(value))
}
