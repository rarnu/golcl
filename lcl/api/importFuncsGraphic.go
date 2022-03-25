package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TGraphic ---------------------------

func Graphic_Create() uintptr {
	ret, _, _ := getLazyProc("Graphic_Create").Call()
	return ret
}

func Graphic_Free(obj uintptr) {
	_, _, _ = getLazyProc("Graphic_Free").Call(obj)
}

func Graphic_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Graphic_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Graphic_LoadFromFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Graphic_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func Graphic_SaveToFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Graphic_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func Graphic_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Graphic_LoadFromStream").Call(obj, Stream)
}

func Graphic_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Graphic_SaveToStream").Call(obj, Stream)
}

func Graphic_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Graphic_Assign").Call(obj, Source)
}

func Graphic_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Graphic_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Graphic_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Graphic_ClassType").Call(obj)
	return TClass(ret)
}

func Graphic_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Graphic_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Graphic_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Graphic_InstanceSize").Call(obj)
	return int32(ret)
}

func Graphic_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Graphic_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Graphic_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Graphic_GetHashCode").Call(obj)
	return int32(ret)
}

func Graphic_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Graphic_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Graphic_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("Graphic_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func Graphic_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Graphic_GetHeight").Call(obj)
	return int32(ret)
}

func Graphic_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Graphic_SetHeight").Call(obj, uintptr(value))
}

func Graphic_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Graphic_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Graphic_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Graphic_SetModified").Call(obj, GoBoolToDBool(value))
}

func Graphic_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("Graphic_GetPalette").Call(obj)
	return ret
}

func Graphic_SetPalette(obj uintptr, value HPALETTE) {
	_, _, _ = getLazyProc("Graphic_SetPalette").Call(obj, value)
}

func Graphic_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Graphic_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Graphic_SetPaletteModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Graphic_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func Graphic_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Graphic_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func Graphic_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Graphic_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func Graphic_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Graphic_GetWidth").Call(obj)
	return int32(ret)
}

func Graphic_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Graphic_SetWidth").Call(obj, uintptr(value))
}

func Graphic_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Graphic_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Graphic_StaticClassType() TClass {
	r, _, _ := getLazyProc("Graphic_StaticClassType").Call()
	return TClass(r)
}
