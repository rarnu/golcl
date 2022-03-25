package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynCppSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynCppSyn(obj any) *TSynCppSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynCppSyn{instance: instance, ptr: ptr}
}

func SynCppSynFromInst(inst uintptr) *TSynCppSyn {
	return AsSynCppSyn(inst)
}

func SynCppSynFromObj(obj IObject) *TSynCppSyn {
	return AsSynCppSyn(obj)
}

func SynCppSynFromUnsafePointer(ptr unsafe.Pointer) *TSynCppSyn {
	return AsSynCppSyn(ptr)
}

func (s *TSynCppSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynCppSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynCppSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynCppSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynCppSynClass() TClass {
	return SynCppSyn_StaticClassType()
}

func NewSynCppSyn(owner IComponent) *TSynCppSyn {
	s := new(TSynCppSyn)
	s.instance = SynCppSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynCppSyn) Free() {
	if s.instance != 0 {
		SynCppSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynCppSyn) GetEnabled() bool {
	return SynCppSyn_GetEnabled(s.instance)
}

func (s *TSynCppSyn) SetEnabled(value bool) {
	SynCppSyn_SetEnabled(s.instance, value)
}

func (s *TSynCppSyn) GetAsmAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetAsmAttri(s.instance))
}

func (s *TSynCppSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetCommentAttri(s.instance))
}

func (s *TSynCppSyn) GetDirecAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetDirecAttri(s.instance))
}

func (s *TSynCppSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynCppSyn) GetInvalidAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetInvalidAttri(s.instance))
}

func (s *TSynCppSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetKeyAttri(s.instance))
}

func (s *TSynCppSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetNumberAttri(s.instance))
}

func (s *TSynCppSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetSpaceAttri(s.instance))
}

func (s *TSynCppSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetStringAttri(s.instance))
}

func (s *TSynCppSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynCppSyn_GetSymbolAttri(s.instance))
}
