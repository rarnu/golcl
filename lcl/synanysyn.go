package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynAnySyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynAnySyn(obj interface{}) *TSynAnySyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynAnySyn{instance: instance, ptr: ptr}
}

func SynAnySynFromInst(inst uintptr) *TSynAnySyn {
	return AsSynAnySyn(inst)
}

func SynAnySynFromObj(obj IObject) *TSynAnySyn {
	return AsSynAnySyn(obj)
}

func SynAnySynFromUnsafePointer(ptr unsafe.Pointer) *TSynAnySyn {
	return AsSynAnySyn(ptr)
}

func (s *TSynAnySyn) Instance() uintptr {
	return s.instance
}

func (s *TSynAnySyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynAnySyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynAnySyn) Is() TIs {
	return TIs(s.instance)
}

func TSynAnySynClass() TClass {
	return SynAnySyn_StaticClassType()
}

func NewSynAnySyn(owner IComponent) *TSynAnySyn {
	s := new(TSynAnySyn)
	s.instance = SynAnySyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynAnySyn) Free() {
	if s.instance != 0 {
		SynAnySyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynAnySyn) GetActiveDot() bool {
	return SynAnySyn_GetActiveDot(s.instance)
}

func (s *TSynAnySyn) SetActiveDot(value bool) {
	SynAnySyn_SetActiveDot(s.instance, value)
}

func (s *TSynAnySyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetCommentAttri(s.instance))
}

func (s *TSynAnySyn) GetComments() CommentStyles {
	return SynAnySyn_GetComments(s.instance)
}

func (s *TSynAnySyn) SetComments(value CommentStyles) {
	SynAnySyn_SetComments(s.instance, value)
}

func (s *TSynAnySyn) GetConstantAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetConstantAttri(s.instance))
}

func (s *TSynAnySyn) GetConstants() *TStrings {
	return AsStrings(SynAnySyn_GetConstants(s.instance))
}

func (s *TSynAnySyn) SetConstants(value *IStrings) {
	SynAnySyn_SetConstants(s.instance, CheckPtr(value))
}

func (s *TSynAnySyn) GetDetectPreprocessor() bool {
	return SynAnySyn_GetDetectPreprocessor(s.instance)
}

func (s *TSynAnySyn) SetDetectPreprocessor(value bool) {
	SynAnySyn_SetDetectPreprocessor(s.instance, value)
}

func (s *TSynAnySyn) GetDollarVariables() bool {
	return SynAnySyn_GetDollarVariables(s.instance)
}

func (s *TSynAnySyn) SetDollarVariables(value bool) {
	SynAnySyn_SetDollarVariables(s.instance, value)
}

func (s *TSynAnySyn) GetDotAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetDotAttri(s.instance))
}

func (s *TSynAnySyn) GetEnabled() bool {
	return SynAnySyn_GetEnabled(s.instance)
}

func (s *TSynAnySyn) SetEnabled(value bool) {
	SynAnySyn_SetEnabled(s.instance, value)
}

func (s *TSynAnySyn) GetEntity() bool {
	return SynAnySyn_GetEntity(s.instance)
}

func (s *TSynAnySyn) SetEntity(value bool) {
	SynAnySyn_SetEntity(s.instance, value)
}

func (s *TSynAnySyn) GetEntityAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetEntityAttri(s.instance))
}

func (s *TSynAnySyn) GetIdentifierAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetIdentifierAttri(s.instance))
}

func (s *TSynAnySyn) GetIdentifierChars() string {
	return SynAnySyn_GetIdentifierChars(s.instance)
}

func (s *TSynAnySyn) SetIdentifierChars(value string) {
	SynAnySyn_SetIdentifierChars(s.instance, value)
}

func (s *TSynAnySyn) GetKeyAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetKeyAttri(s.instance))
}

func (s *TSynAnySyn) GetKeyWords() *TStrings {
	return AsStrings(SynAnySyn_GetKeyWords(s.instance))
}

func (s *TSynAnySyn) SetKeyWords(value *IStrings) {
	SynAnySyn_SetKeyWords(s.instance, CheckPtr(value))
}

func (s *TSynAnySyn) GetMarkup() bool {
	return SynAnySyn_GetMarkup(s.instance)
}

func (s *TSynAnySyn) SetMarkup(value bool) {
	SynAnySyn_SetMarkup(s.instance, value)
}

func (s *TSynAnySyn) Name() string {
	return SynAnySyn_GetName(s.instance)
}

func (s *TSynAnySyn) SetName(value string) {
	SynAnySyn_SetName(s.instance, value)
}

func (s *TSynAnySyn) GetNumberAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetNumberAttri(s.instance))
}

func (s *TSynAnySyn) GetObjectAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetObjectAttri(s.instance))
}

func (s *TSynAnySyn) GetObjects() *TStrings {
	return AsStrings(SynAnySyn_GetObjects(s.instance))
}

func (s *TSynAnySyn) SetObjects(value *IStrings) {
	SynAnySyn_SetObjects(s.instance, CheckPtr(value))
}

func (s *TSynAnySyn) GetPreprocessorAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetPreprocessorAttri(s.instance))
}

func (s *TSynAnySyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetSpaceAttri(s.instance))
}

func (s *TSynAnySyn) GetStringAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetStringAttri(s.instance))
}

func (s *TSynAnySyn) GetStringDelim() TStringDelim {
	return SynAnySyn_GetStringDelim(s.instance)
}

func (s *TSynAnySyn) SetStringDelim(value TStringDelim) {
	SynAnySyn_SetStringDelim(s.instance, value)
}

func (s *TSynAnySyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetSymbolAttri(s.instance))
}

func (s *TSynAnySyn) GetVariableAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynAnySyn_GetVariableAttri(s.instance))
}
