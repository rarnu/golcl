package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynHighlighterAttributes_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynHighlighterAttributes_StaticClassType").Call()
	return TClass(r)
}

func SynHighlighterAttributes_GetBackground(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetBackground").Call(obj)
	return TColor(r)
}

func SynHighlighterAttributes_SetBackground(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetBackground").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetForeground(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetForeground").Call(obj)
	return TColor(r)
}

func SynHighlighterAttributes_SetForeground(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetForeground").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetFrameColor(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetFrameColor").Call(obj)
	return TColor(r)
}

func SynHighlighterAttributes_SetFrameColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetFrameColor").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetFrameStyle(obj uintptr) TSynLineStyle {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetFrameStyle").Call(obj)
	return TSynLineStyle(r)
}

func SynHighlighterAttributes_SetFrameStyle(obj uintptr, value TSynLineStyle) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetFrameStyle").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetFrameEdges(obj uintptr) TSynFrameEdges {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetFrameEdges").Call(obj)
	return TSynFrameEdges(r)
}

func SynHighlighterAttributes_SetFrameEdges(obj uintptr, value TSynFrameEdges) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetFrameEdges").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetStyle(obj uintptr) TFontStyles {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetStyle").Call(obj)
	return TFontStyles(r)
}

func SynHighlighterAttributes_SetStyle(obj uintptr, value TFontStyles) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetStyle").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetStyleMask(obj uintptr) TFontStyles {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetStyleMask").Call(obj)
	return TFontStyles(r)
}

func SynHighlighterAttributes_SetStyleMask(obj uintptr, value TFontStyles) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetStyleMask").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetBackPriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetBackPriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetBackPriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetBackPriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetForePriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetForePriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetForePriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetForePriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetFramePriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetFramePriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetFramePriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetFramePriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetBoldPriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetBoldPriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetBoldPriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetBoldPriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetItalicPriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetItalicPriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetItalicPriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetItalicPriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetUnderlinePriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetUnderlinePriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetUnderlinePriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetUnderlinePriority").Call(obj, uintptr(value))
}

func SynHighlighterAttributes_GetStrikeOutPriority(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynHighlighterAttributes_GetStrikeOutPriority").Call(obj)
	return int32(r)
}

func SynHighlighterAttributes_SetStrikeOutPriority(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynHighlighterAttributes_SetStrikeOutPriority").Call(obj, uintptr(value))
}
