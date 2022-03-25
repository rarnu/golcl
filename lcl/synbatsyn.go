package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynBatSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynBatSyn(obj any) *TSynBatSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynBatSyn{instance: instance, ptr: ptr}
}

func SynBatSynFromInst(inst uintptr) *TSynBatSyn {
	return AsSynBatSyn(inst)
}

func SynBatSynFromObj(obj IObject) *TSynBatSyn {
	return AsSynBatSyn(obj)
}

func SynBatSynFromUnsafePointer(ptr unsafe.Pointer) *TSynBatSyn {
	return AsSynBatSyn(ptr)
}

func (s *TSynBatSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynBatSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynBatSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynBatSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynBatSynClass() TClass {
	return SynBatSyn_StaticClassType()
}

func NewSynBatSyn(owner IComponent) *TSynBatSyn {
	s := new(TSynBatSyn)
	s.instance = SynBatSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynBatSyn) Free() {
	if s.instance != 0 {
		SynBatSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynBatSyn) GetEnabled() bool {
	return SynBatSyn_GetEnabled(s.instance)
}

func (s *TSynBatSyn) SetEnabled(value bool) {
	SynBatSyn_SetEnabled(s.instance, value)
}

func (s *TSynBatSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetCommentAttri(s.instance))
}

func (s *TSynBatSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynBatSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetKeyAttri(s.instance))
}

func (s *TSynBatSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetNumberAttri(s.instance))
}

func (s *TSynBatSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetSpaceAttri(s.instance))
}

func (s *TSynBatSyn) GetVariableAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynBatSyn_GetVariableAttri(s.instance))
}
