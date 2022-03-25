package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPngImage ---------------------------

func PngImage_Create() uintptr {
	ret, _, _ := getLazyProc("PngImage_Create").Call()
	return ret
}

func PngImage_Free(obj uintptr) {
	_, _, _ = getLazyProc("PngImage_Free").Call(obj)
}

func PngImage_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("PngImage_Assign").Call(obj, Source)
}

func PngImage_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("PngImage_LoadFromStream").Call(obj, Stream)
}

func PngImage_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("PngImage_SaveToStream").Call(obj, Stream)
}

func PngImage_LoadFromResourceName(obj uintptr, Instance uintptr, Name string) {
	_, _, _ = getLazyProc("PngImage_LoadFromResourceName").Call(obj, Instance, GoStrToDStr(Name))
}

func PngImage_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32) {
	_, _, _ = getLazyProc("PngImage_LoadFromResourceID").Call(obj, Instance, uintptr(ResID))
}

func PngImage_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PngImage_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PngImage_LoadFromFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("PngImage_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func PngImage_SaveToFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("PngImage_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func PngImage_SetSize(obj uintptr, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("PngImage_SetSize").Call(obj, uintptr(AWidth), uintptr(AHeight))
}

func PngImage_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PngImage_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PngImage_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PngImage_ClassType").Call(obj)
	return TClass(ret)
}

func PngImage_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PngImage_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PngImage_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PngImage_InstanceSize").Call(obj)
	return int32(ret)
}

func PngImage_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PngImage_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PngImage_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PngImage_GetHashCode").Call(obj)
	return int32(ret)
}

func PngImage_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PngImage_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PngImage_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PngImage_GetCanvas").Call(obj)
	return ret
}

func PngImage_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PngImage_GetWidth").Call(obj)
	return int32(ret)
}

func PngImage_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PngImage_GetHeight").Call(obj)
	return int32(ret)
}

func PngImage_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("PngImage_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func PngImage_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("PngImage_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func PngImage_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PngImage_SetModified").Call(obj, GoBoolToDBool(value))
}

func PngImage_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("PngImage_GetPalette").Call(obj)
	return ret
}

func PngImage_SetPalette(obj uintptr, value HPALETTE) {
	_, _, _ = getLazyProc("PngImage_SetPalette").Call(obj, value)
}

func PngImage_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("PngImage_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func PngImage_SetPaletteModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PngImage_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func PngImage_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PngImage_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func PngImage_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PngImage_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func PngImage_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PngImage_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func PngImage_StaticClassType() TClass {
	r, _, _ := getLazyProc("PngImage_StaticClassType").Call()
	return TClass(r)
}
