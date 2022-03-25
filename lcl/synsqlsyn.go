package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynSQLSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynSQLSyn(obj any) *TSynSQLSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynSQLSyn{instance: instance, ptr: ptr}
}

func SynSQLSynFromInst(inst uintptr) *TSynSQLSyn {
	return AsSynSQLSyn(inst)
}

func SynSQLSynFromObj(obj IObject) *TSynSQLSyn {
	return AsSynSQLSyn(obj)
}

func SynSQLSynFromUnsafePointer(ptr unsafe.Pointer) *TSynSQLSyn {
	return AsSynSQLSyn(ptr)
}

func (s *TSynSQLSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynSQLSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynSQLSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynSQLSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynSQLSynClass() TClass {
	return SynSQLSyn_StaticClassType()
}

func NewSynSQLSyn(owner IComponent) *TSynSQLSyn {
	s := new(TSynSQLSyn)
	s.instance = SynSQLSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynSQLSyn) Free() {
	if s.instance != 0 {
		SynSQLSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynSQLSyn) GetEnabled() bool {
	return SynSQLSyn_GetEnabled(s.instance)
}

func (s *TSynSQLSyn) SetEnabled(value bool) {
	SynSQLSyn_SetEnabled(s.instance, value)
}

func (s *TSynSQLSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetCommentAttri(s.instance))
}

func (s *TSynSQLSyn) GetDataTypeAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetDataTypeAttri(s.instance))
}

func (s *TSynSQLSyn) GetDefaultPackageAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetDefaultPackageAttri(s.instance))
}

func (s *TSynSQLSyn) GetExceptionAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetExceptionAttri(s.instance))
}

func (s *TSynSQLSyn) GetFunctionAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetFunctionAttri(s.instance))
}

func (s *TSynSQLSyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetIdentifierAttri(s.instance))
}

func (s *TSynSQLSyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetKeyAttri(s.instance))
}

func (s *TSynSQLSyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetNumberAttri(s.instance))
}

func (s *TSynSQLSyn) GetPLSQLAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetPLSQLAttri(s.instance))
}

func (s *TSynSQLSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetSpaceAttri(s.instance))
}

func (s *TSynSQLSyn) GetSQLPlusAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetSQLPlusAttri(s.instance))
}

func (s *TSynSQLSyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetStringAttri(s.instance))
}

func (s *TSynSQLSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetSymbolAttri(s.instance))
}

func (s *TSynSQLSyn) GetTableNameAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetTableNameAttri(s.instance))
}

func (s *TSynSQLSyn) GetVariableAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynSQLSyn_GetVariableAttri(s.instance))
}

func (s *TSynSQLSyn) GetTableNames() *TStrings {
	return AsStrings(SynSQLSyn_GetTableNames(s.instance))
}

func (s *TSynSQLSyn) SetTableNames(value *IStrings) {
	SynSQLSyn_SetTableNames(s.instance, CheckPtr(value))
}

func (s *TSynSQLSyn) GetSQLDialect() TSQLDialect {
	return SynSQLSyn_GetSQLDialect(s.instance)
}

func (s *TSynSQLSyn) SetSQLDialect(value TSQLDialect) {
	SynSQLSyn_SetSQLDialect(s.instance, value)
}
