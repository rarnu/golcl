package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TJPEGImage ---------------------------

func JPEGImage_Create() uintptr {
	ret, _, _ := getLazyProc("JPEGImage_Create").Call()
	return ret
}

func JPEGImage_Free(obj uintptr) {
	getLazyProc("JPEGImage_Free").Call(obj)
}

func JPEGImage_Assign(obj uintptr, Source uintptr) {
	getLazyProc("JPEGImage_Assign").Call(obj, Source)
}

func JPEGImage_LoadFromStream(obj uintptr, Stream uintptr) {
	getLazyProc("JPEGImage_LoadFromStream").Call(obj, Stream)
}

func JPEGImage_SaveToStream(obj uintptr, Stream uintptr) {
	getLazyProc("JPEGImage_SaveToStream").Call(obj, Stream)
}

func JPEGImage_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("JPEGImage_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func JPEGImage_LoadFromFile(obj uintptr, Filename string) {
	getLazyProc("JPEGImage_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func JPEGImage_SaveToFile(obj uintptr, Filename string) {
	getLazyProc("JPEGImage_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func JPEGImage_SetSize(obj uintptr, AWidth int32, AHeight int32) {
	getLazyProc("JPEGImage_SetSize").Call(obj, uintptr(AWidth), uintptr(AHeight))
}

func JPEGImage_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("JPEGImage_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func JPEGImage_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("JPEGImage_ClassType").Call(obj)
	return TClass(ret)
}

func JPEGImage_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("JPEGImage_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func JPEGImage_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("JPEGImage_InstanceSize").Call(obj)
	return int32(ret)
}

func JPEGImage_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("JPEGImage_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func JPEGImage_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("JPEGImage_GetHashCode").Call(obj)
	return int32(ret)
}

func JPEGImage_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("JPEGImage_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func JPEGImage_GetPixelFormat(obj uintptr) TJPEGPixelFormat {
	ret, _, _ := getLazyProc("JPEGImage_GetPixelFormat").Call(obj)
	return TJPEGPixelFormat(ret)
}

func JPEGImage_SetPixelFormat(obj uintptr, value TJPEGPixelFormat) {
	getLazyProc("JPEGImage_SetPixelFormat").Call(obj, uintptr(value))
}

func JPEGImage_GetPerformance(obj uintptr) TJPEGPerformance {
	ret, _, _ := getLazyProc("JPEGImage_GetPerformance").Call(obj)
	return TJPEGPerformance(ret)
}

func JPEGImage_SetPerformance(obj uintptr, value TJPEGPerformance) {
	getLazyProc("JPEGImage_SetPerformance").Call(obj, uintptr(value))
}

func JPEGImage_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("JPEGImage_GetCanvas").Call(obj)
	return ret
}

func JPEGImage_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("JPEGImage_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func JPEGImage_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("JPEGImage_GetHeight").Call(obj)
	return int32(ret)
}

func JPEGImage_SetHeight(obj uintptr, value int32) {
	getLazyProc("JPEGImage_SetHeight").Call(obj, uintptr(value))
}

func JPEGImage_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("JPEGImage_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func JPEGImage_SetModified(obj uintptr, value bool) {
	getLazyProc("JPEGImage_SetModified").Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("JPEGImage_GetPalette").Call(obj)
	return HPALETTE(ret)
}

func JPEGImage_SetPalette(obj uintptr, value HPALETTE) {
	getLazyProc("JPEGImage_SetPalette").Call(obj, uintptr(value))
}

func JPEGImage_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("JPEGImage_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func JPEGImage_SetPaletteModified(obj uintptr, value bool) {
	getLazyProc("JPEGImage_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("JPEGImage_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func JPEGImage_SetTransparent(obj uintptr, value bool) {
	getLazyProc("JPEGImage_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("JPEGImage_GetWidth").Call(obj)
	return int32(ret)
}

func JPEGImage_SetWidth(obj uintptr, value int32) {
	getLazyProc("JPEGImage_SetWidth").Call(obj, uintptr(value))
}

func JPEGImage_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("JPEGImage_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func JPEGImage_StaticClassType() TClass {
	r, _, _ := getLazyProc("JPEGImage_StaticClassType").Call()
	return TClass(r)
}
