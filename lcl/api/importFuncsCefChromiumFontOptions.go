package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChromiumFontOptions_Create() uintptr {
	ret, _, _ := getLazyProc("ChromiumFontOptions_Create").Call()
	return ret
}

func ChromiumFontOptions_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumFontOptions_Free").Call(obj)
}

func ChromiumFontOptions_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChromiumFontOptions_StaticClassType").Call()
	return TClass(ret)
}

func ChromiumFontOptions_GetStandardFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetStandardFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetStandardFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetStandardFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetFixedFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetFixedFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetFixedFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetFixedFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetSerifFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetSerifFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetSerifFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetSerifFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetSansSerifFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetSansSerifFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetSansSerifFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetSansSerifFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetCursiveFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetCursiveFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetCursiveFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetCursiveFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetFantasyFontFamily(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetFantasyFontFamily").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumFontOptions_SetFantasyFontFamily(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetFantasyFontFamily").Call(obj, GoStrToDStr(Value))
}

func ChromiumFontOptions_GetDefaultFontSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetDefaultFontSize").Call(obj)
	return int32(ret)
}

func ChromiumFontOptions_SetDefaultFontSize(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetDefaultFontSize").Call(obj, uintptr(Value))
}

func ChromiumFontOptions_GetDefaultFixedFontSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetDefaultFixedFontSize").Call(obj)
	return int32(ret)
}

func ChromiumFontOptions_SetDefaultFixedFontSize(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetDefaultFixedFontSize").Call(obj, uintptr(Value))
}

func ChromiumFontOptions_GetMinimumFontSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetMinimumFontSize").Call(obj)
	return int32(ret)
}

func ChromiumFontOptions_SetMinimumFontSize(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetMinimumFontSize").Call(obj, uintptr(Value))
}

func ChromiumFontOptions_GetMinimumLogicalFontSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetMinimumLogicalFontSize").Call(obj)
	return int32(ret)
}

func ChromiumFontOptions_SetMinimumLogicalFontSize(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetMinimumLogicalFontSize").Call(obj, uintptr(Value))
}

func ChromiumFontOptions_GetRemoteFonts(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumFontOptions_GetRemoteFonts").Call(obj)
	return TCefState(ret)
}

func ChromiumFontOptions_SetRemoteFonts(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumFontOptions_SetRemoteFonts").Call(obj, uintptr(Value))
}
