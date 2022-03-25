package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynPythonSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynPythonSyn(obj any) *TSynPythonSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynPythonSyn{instance: instance, ptr: ptr}
}

func SynPythonSynFromInst(inst uintptr) *TSynPythonSyn {
	return AsSynPythonSyn(inst)
}

func SynPythonSynFromObj(obj IObject) *TSynPythonSyn {
	return AsSynPythonSyn(obj)
}

func SynPythonSynFromUnsafePointer(ptr unsafe.Pointer) *TSynPythonSyn {
	return AsSynPythonSyn(ptr)
}

func (s *TSynPythonSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynPythonSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynPythonSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynPythonSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynPythonSynClass() TClass {
	return SynPythonSyn_StaticClassType()
}

func NewSynPythonSyn(owner IComponent) *TSynPythonSyn {
	s := new(TSynPythonSyn)
	s.instance = SynPythonSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynPythonSyn) Free() {
	if s.instance != 0 {
		SynPythonSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynPythonSyn) GetEnabled() bool {
	return SynPythonSyn_GetEnabled(s.instance)
}

func (s *TSynPythonSyn) SetEnabled(value bool) {
	SynPythonSyn_SetEnabled(s.instance, value)
}

func (s *TSynPythonSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetCommentAttri(s.instance))
}

func (s *TSynPythonSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynPythonSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetKeyAttri(s.instance))
}

func (s *TSynPythonSyn) GetNonKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetNonKeyAttri(s.instance))
}

func (s *TSynPythonSyn) GetSystemAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetSystemAttri(s.instance))
}

func (s *TSynPythonSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetNumberAttri(s.instance))
}

func (s *TSynPythonSyn) GetHexAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetHexAttri(s.instance))
}

func (s *TSynPythonSyn) GetOctalAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetOctalAttri(s.instance))
}

func (s *TSynPythonSyn) GetFloatAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetFloatAttri(s.instance))
}

func (s *TSynPythonSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetSpaceAttri(s.instance))
}

func (s *TSynPythonSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetStringAttri(s.instance))
}

func (s *TSynPythonSyn) GetDocStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetDocStringAttri(s.instance))
}

func (s *TSynPythonSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetSymbolAttri(s.instance))
}

func (s *TSynPythonSyn) GetErrorAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynPythonSyn_GetErrorAttri(s.instance))
}
