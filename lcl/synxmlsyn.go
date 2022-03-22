package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynXmlSyn struct {
	TSynCustomHighlighter
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynXmlSyn(obj interface{}) *TSynXmlSyn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynXmlSyn{instance: instance, ptr: ptr}
}

func SynXmlSynFromInst(inst uintptr) *TSynXmlSyn {
	return AsSynXmlSyn(inst)
}

func SynXmlSynFromObj(obj IObject) *TSynXmlSyn {
	return AsSynXmlSyn(obj)
}

func SynXmlSynFromUnsafePointer(ptr unsafe.Pointer) *TSynXmlSyn {
	return AsSynXmlSyn(ptr)
}

func (s *TSynXmlSyn) Instance() uintptr {
	return s.instance
}

func (s *TSynXmlSyn) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynXmlSyn) IsValid() bool {
	return s.instance != 0
}

func (s *TSynXmlSyn) Is() TIs {
	return TIs(s.instance)
}

func TSynXmlSynClass() TClass {
	return SynXMLSyn_StaticClassType()
}

func NewSynXmlSyn(owner IComponent) *TSynXmlSyn {
	s := new(TSynXmlSyn)
	s.instance = SynXMLSyn_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func (s *TSynXmlSyn) Free() {
	if s.instance != 0 {
		SynXMLSyn_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynXmlSyn) GetEnabled() bool {
	return SynXMLSyn_GetEnabled(s.instance)
}

func (s *TSynXmlSyn) SetEnabled(value bool) {
	SynXMLSyn_SetEnabled(s.instance, value)
}

func (s *TSynXmlSyn) GetElementAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetElementAttri(s.instance))
}

func (s *TSynXmlSyn) GetAttributeAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetAttributeAttri(s.instance))
}

func (s *TSynXmlSyn) GetNamespaceAttributeAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetNamespaceAttributeAttri(s.instance))
}

func (s *TSynXmlSyn) GetAttributeValueAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetAttributeValueAttri(s.instance))
}

func (s *TSynXmlSyn) GetNamespaceAttributeValueAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetNamespaceAttributeValueAttri(s.instance))
}

func (s *TSynXmlSyn) GetTextAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetTextAttri(s.instance))
}

func (s *TSynXmlSyn) GetCDATAAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetCDATAAttri(s.instance))
}

func (s *TSynXmlSyn) GetEntityRefAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetEntityRefAttri(s.instance))
}

func (s *TSynXmlSyn) GetProcessingInstructionAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetProcessingInstructionAttri(s.instance))
}

func (s *TSynXmlSyn) GetCommentAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetCommentAttri(s.instance))
}

func (s *TSynXmlSyn) GetDocTypeAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetDocTypeAttri(s.instance))
}

func (s *TSynXmlSyn) GetSpaceAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetSpaceAttri(s.instance))
}

func (s *TSynXmlSyn) GetSymbolAttri() *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(SynXMLSyn_GetSymbolAttri(s.instance))
}

func (s *TSynXmlSyn) GetWantBracesParsed() bool {
	return SynXMLSyn_GetWantBracesParsed(s.instance)
}

func (s *TSynXmlSyn) SetWantBracesParsed(value bool) {
	SynXMLSyn_SetWantBracesParsed(s.instance, value)
}
