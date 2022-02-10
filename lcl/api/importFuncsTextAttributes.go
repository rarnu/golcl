package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTextAttributes ---------------------------

func TextAttributes_Assign(obj uintptr, Source uintptr) {
	getLazyProc("TextAttributes_Assign").Call(obj, Source)
}

func TextAttributes_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TextAttributes_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TextAttributes_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TextAttributes_ClassType").Call(obj)
	return TClass(ret)
}

func TextAttributes_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TextAttributes_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TextAttributes_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TextAttributes_InstanceSize").Call(obj)
	return int32(ret)
}

func TextAttributes_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TextAttributes_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TextAttributes_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TextAttributes_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TextAttributes_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TextAttributes_GetHashCode").Call(obj)
	return int32(ret)
}

func TextAttributes_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TextAttributes_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TextAttributes_GetCharset(obj uintptr) TFontCharset {
	ret, _, _ := getLazyProc("TextAttributes_GetCharset").Call(obj)
	return TFontCharset(ret)
}

func TextAttributes_SetCharset(obj uintptr, value TFontCharset) {
	getLazyProc("TextAttributes_SetCharset").Call(obj, uintptr(value))
}

func TextAttributes_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TextAttributes_GetColor").Call(obj)
	return TColor(ret)
}

func TextAttributes_SetColor(obj uintptr, value TColor) {
	getLazyProc("TextAttributes_SetColor").Call(obj, uintptr(value))
}

func TextAttributes_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TextAttributes_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TextAttributes_SetName(obj uintptr, value string) {
	getLazyProc("TextAttributes_SetName").Call(obj, GoStrToDStr(value))
}

func TextAttributes_GetPitch(obj uintptr) TFontPitch {
	ret, _, _ := getLazyProc("TextAttributes_GetPitch").Call(obj)
	return TFontPitch(ret)
}

func TextAttributes_SetPitch(obj uintptr, value TFontPitch) {
	getLazyProc("TextAttributes_SetPitch").Call(obj, uintptr(value))
}

func TextAttributes_GetSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TextAttributes_GetSize").Call(obj)
	return int32(ret)
}

func TextAttributes_SetSize(obj uintptr, value int32) {
	getLazyProc("TextAttributes_SetSize").Call(obj, uintptr(value))
}

func TextAttributes_GetStyle(obj uintptr) TFontStyles {
	ret, _, _ := getLazyProc("TextAttributes_GetStyle").Call(obj)
	return TFontStyles(ret)
}

func TextAttributes_SetStyle(obj uintptr, value TFontStyles) {
	getLazyProc("TextAttributes_SetStyle").Call(obj, uintptr(value))
}

func TextAttributes_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TextAttributes_GetHeight").Call(obj)
	return int32(ret)
}

func TextAttributes_SetHeight(obj uintptr, value int32) {
	getLazyProc("TextAttributes_SetHeight").Call(obj, uintptr(value))
}

func TextAttributes_StaticClassType() TClass {
	r, _, _ := getLazyProc("TextAttributes_StaticClassType").Call()
	return TClass(r)
}
