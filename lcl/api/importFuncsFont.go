package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TFont ---------------------------

func Font_Create() uintptr {
	ret, _, _ := getLazyProc("Font_Create").Call()
	return ret
}

func Font_Free(obj uintptr) {
	_, _, _ = getLazyProc("Font_Free").Call(obj)
}

func Font_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Font_Assign").Call(obj, Source)
}

func Font_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Font_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Font_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Font_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Font_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Font_ClassType").Call(obj)
	return TClass(ret)
}

func Font_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Font_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Font_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_InstanceSize").Call(obj)
	return int32(ret)
}

func Font_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Font_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Font_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Font_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Font_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_GetHashCode").Call(obj)
	return int32(ret)
}

func Font_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Font_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Font_GetHandle(obj uintptr) HFONT {
	ret, _, _ := getLazyProc("Font_GetHandle").Call(obj)
	return ret
}

func Font_SetHandle(obj uintptr, value HFONT) {
	_, _, _ = getLazyProc("Font_SetHandle").Call(obj, value)
}

func Font_GetPixelsPerInch(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_GetPixelsPerInch").Call(obj)
	return int32(ret)
}

func Font_SetPixelsPerInch(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Font_SetPixelsPerInch").Call(obj, uintptr(value))
}

func Font_GetCharset(obj uintptr) TFontCharset {
	ret, _, _ := getLazyProc("Font_GetCharset").Call(obj)
	return TFontCharset(ret)
}

func Font_SetCharset(obj uintptr, value TFontCharset) {
	_, _, _ = getLazyProc("Font_SetCharset").Call(obj, uintptr(value))
}

func Font_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Font_GetColor").Call(obj)
	return TColor(ret)
}

func Font_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Font_SetColor").Call(obj, uintptr(value))
}

func Font_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_GetHeight").Call(obj)
	return int32(ret)
}

func Font_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Font_SetHeight").Call(obj, uintptr(value))
}

func Font_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Font_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Font_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Font_SetName").Call(obj, GoStrToDStr(value))
}

func Font_GetOrientation(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_GetOrientation").Call(obj)
	return int32(ret)
}

func Font_SetOrientation(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Font_SetOrientation").Call(obj, uintptr(value))
}

func Font_GetPitch(obj uintptr) TFontPitch {
	ret, _, _ := getLazyProc("Font_GetPitch").Call(obj)
	return TFontPitch(ret)
}

func Font_SetPitch(obj uintptr, value TFontPitch) {
	_, _, _ = getLazyProc("Font_SetPitch").Call(obj, uintptr(value))
}

func Font_GetSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Font_GetSize").Call(obj)
	return int32(ret)
}

func Font_SetSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Font_SetSize").Call(obj, uintptr(value))
}

func Font_GetStyle(obj uintptr) TFontStyles {
	ret, _, _ := getLazyProc("Font_GetStyle").Call(obj)
	return TFontStyles(ret)
}

func Font_SetStyle(obj uintptr, value TFontStyles) {
	_, _, _ = getLazyProc("Font_SetStyle").Call(obj, uintptr(value))
}

func Font_GetQuality(obj uintptr) TFontQuality {
	ret, _, _ := getLazyProc("Font_GetQuality").Call(obj)
	return TFontQuality(ret)
}

func Font_SetQuality(obj uintptr, value TFontQuality) {
	_, _, _ = getLazyProc("Font_SetQuality").Call(obj, uintptr(value))
}

func Font_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Font_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Font_StaticClassType() TClass {
	r, _, _ := getLazyProc("Font_StaticClassType").Call()
	return TClass(r)
}
