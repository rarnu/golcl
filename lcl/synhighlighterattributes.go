package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynHighlighterAttributes struct {
	IObject
	instance uintptr
	ptr      unsafe.Pointer
}

func AsSynHighlighterAttributes(obj interface{}) *TSynHighlighterAttributes {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynHighlighterAttributes{instance: instance, ptr: ptr}
}

func SynHighlighterAttributesFromInst(inst uintptr) *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(inst)
}

func SynHighlighterAttributesFromObj(obj IObject) *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(obj)
}

func SynHighlighterAttributesFromUnsafePointer(ptr unsafe.Pointer) *TSynHighlighterAttributes {
	return AsSynHighlighterAttributes(ptr)
}

func (s *TSynHighlighterAttributes) Instance() uintptr {
	return s.instance
}

func (s *TSynHighlighterAttributes) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynHighlighterAttributes) IsValid() bool {
	return s.instance != 0
}

func (s *TSynHighlighterAttributes) Is() TIs {
	return TIs(s.instance)
}

func TSynHighlighterAttributesClass() TClass {
	return SynHighlighterAttributes_StaticClassType()
}

func (s *TSynHighlighterAttributes) ClassName() string {
	return SynHighlighterAttributes_ClassName(s.instance)
}

func (s *TSynHighlighterAttributes) Free() {
	if s.instance != 0 {
		SynHighlighterAttributes_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynHighlighterAttributes) GetHashCode() int32 {
	return SynHighlighterAttributes_GetHashCode(s.instance)
}

func (s *TSynHighlighterAttributes) Equals(obj IObject) bool {
	return SynHighlighterAttributes_Equals(s.instance, CheckPtr(obj))
}

func (s *TSynHighlighterAttributes) ClassType() TClass {
	return SynHighlighterAttributes_ClassType(s.instance)
}

func (s *TSynHighlighterAttributes) InstanceSize() int32 {
	return SynHighlighterAttributes_InstanceSize(s.instance)
}

func (s *TSynHighlighterAttributes) InheritsFrom(aClass TClass) bool {
	return SynHighlighterAttributes_InheritsFrom(s.instance, aClass)
}

func (s *TSynHighlighterAttributes) GetBackground() TColor {
	return SynHighlighterAttributes_GetBackground(s.instance)
}

func (s *TSynHighlighterAttributes) SetBackground(value TColor) {
	SynHighlighterAttributes_SetBackground(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetForeground() TColor {
	return SynHighlighterAttributes_GetForeground(s.instance)
}

func (s *TSynHighlighterAttributes) SetForeground(value TColor) {
	SynHighlighterAttributes_SetForeground(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetFrameColor() TColor {
	return SynHighlighterAttributes_GetFrameColor(s.instance)
}

func (s *TSynHighlighterAttributes) SetFrameColor(value TColor) {
	SynHighlighterAttributes_SetFrameColor(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetFrameStyle() TSynLineStyle {
	return SynHighlighterAttributes_GetFrameStyle(s.instance)
}

func (s *TSynHighlighterAttributes) SetFrameStyle(value TSynLineStyle) {
	SynHighlighterAttributes_SetFrameStyle(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetFrameEdges() TSynFrameEdges {
	return SynHighlighterAttributes_GetFrameEdges(s.instance)
}

func (s *TSynHighlighterAttributes) SetFrameEdges(value TSynFrameEdges) {
	SynHighlighterAttributes_SetFrameEdges(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetStyle() TFontStyles {
	return SynHighlighterAttributes_GetStyle(s.instance)
}

func (s *TSynHighlighterAttributes) SetStyle(value TFontStyles) {
	SynHighlighterAttributes_SetStyle(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetStyleMask() TFontStyles {
	return SynHighlighterAttributes_GetStyleMask(s.instance)
}

func (s *TSynHighlighterAttributes) SetStyleMask(value TFontStyles) {
	SynHighlighterAttributes_SetStyleMask(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetBackPriority() int32 {
	return SynHighlighterAttributes_GetBackPriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetBackPriority(value int32) {
	SynHighlighterAttributes_SetBackPriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetForePriority() int32 {
	return SynHighlighterAttributes_GetForePriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetForePriority(value int32) {
	SynHighlighterAttributes_SetForePriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetFramePriority() int32 {
	return SynHighlighterAttributes_GetFramePriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetFramePriority(value int32) {
	SynHighlighterAttributes_SetFramePriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetBoldPriority() int32 {
	return SynHighlighterAttributes_GetBoldPriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetBoldPriority(value int32) {
	SynHighlighterAttributes_SetBoldPriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetItalicPriority() int32 {
	return SynHighlighterAttributes_GetItalicPriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetItalicPriority(value int32) {
	SynHighlighterAttributes_SetItalicPriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetUnderlinePriority() int32 {
	return SynHighlighterAttributes_GetUnderlinePriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetUnderlinePriority(value int32) {
	SynHighlighterAttributes_SetUnderlinePriority(s.instance, value)
}

func (s *TSynHighlighterAttributes) GetStrikeOutPriority() int32 {
	return SynHighlighterAttributes_GetStrikeOutPriority(s.instance)
}

func (s *TSynHighlighterAttributes) SetStrikeOutPriority(value int32) {
	SynHighlighterAttributes_SetStrikeOutPriority(s.instance, value)
}
