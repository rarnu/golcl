package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TGIFImage ---------------------------

func GIFImage_Create() uintptr {
	ret, _, _ := getLazyProc("GIFImage_Create").Call()
	return ret
}

func GIFImage_Free(obj uintptr) {
	_, _, _ = getLazyProc("GIFImage_Free").Call(obj)
}

func GIFImage_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("GIFImage_SaveToStream").Call(obj, Stream)
}

func GIFImage_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("GIFImage_LoadFromStream").Call(obj, Stream)
}

func GIFImage_Clear(obj uintptr) {
	_, _, _ = getLazyProc("GIFImage_Clear").Call(obj)
}

func GIFImage_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("GIFImage_Assign").Call(obj, Source)
}

func GIFImage_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("GIFImage_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func GIFImage_LoadFromFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("GIFImage_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func GIFImage_SaveToFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("GIFImage_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func GIFImage_SetSize(obj uintptr, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("GIFImage_SetSize").Call(obj, uintptr(AWidth), uintptr(AHeight))
}

func GIFImage_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("GIFImage_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func GIFImage_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("GIFImage_ClassType").Call(obj)
	return TClass(ret)
}

func GIFImage_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("GIFImage_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func GIFImage_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GIFImage_InstanceSize").Call(obj)
	return int32(ret)
}

func GIFImage_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("GIFImage_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func GIFImage_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GIFImage_GetHashCode").Call(obj)
	return int32(ret)
}

func GIFImage_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("GIFImage_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func GIFImage_GetBitsPerPixel(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GIFImage_GetBitsPerPixel").Call(obj)
	return int32(ret)
}

func GIFImage_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("GIFImage_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func GIFImage_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GIFImage_GetHeight").Call(obj)
	return int32(ret)
}

func GIFImage_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GIFImage_SetHeight").Call(obj, uintptr(value))
}

func GIFImage_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("GIFImage_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func GIFImage_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GIFImage_SetModified").Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("GIFImage_GetPalette").Call(obj)
	return ret
}

func GIFImage_SetPalette(obj uintptr, value HPALETTE) {
	_, _, _ = getLazyProc("GIFImage_SetPalette").Call(obj, value)
}

func GIFImage_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("GIFImage_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func GIFImage_SetPaletteModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GIFImage_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("GIFImage_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func GIFImage_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GIFImage_GetWidth").Call(obj)
	return int32(ret)
}

func GIFImage_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GIFImage_SetWidth").Call(obj, uintptr(value))
}

func GIFImage_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GIFImage_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func GIFImage_StaticClassType() TClass {
	r, _, _ := getLazyProc("GIFImage_StaticClassType").Call()
	return TClass(r)
}
