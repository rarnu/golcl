package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TIcon ---------------------------

func Icon_Create() uintptr {
	ret, _, _ := getLazyProc("Icon_Create").Call()
	return ret
}

func Icon_Free(obj uintptr) {
	getLazyProc("Icon_Free").Call(obj)
}

func Icon_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Icon_Assign").Call(obj, Source)
}

func Icon_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Icon_LoadFromStream(obj uintptr, Stream uintptr) {
	getLazyProc("Icon_LoadFromStream").Call(obj, Stream)
}

func Icon_SaveToStream(obj uintptr, Stream uintptr) {
	getLazyProc("Icon_SaveToStream").Call(obj, Stream)
}

func Icon_SetSize(obj uintptr, AWidth int32, AHeight int32) {
	getLazyProc("Icon_SetSize").Call(obj, uintptr(AWidth), uintptr(AHeight))
}

func Icon_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string) {
	getLazyProc("Icon_LoadFromResourceName").Call(obj, Instance, GoStrToDStr(ResName))
}

func Icon_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32) {
	getLazyProc("Icon_LoadFromResourceID").Call(obj, Instance, uintptr(ResID))
}

func Icon_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Icon_LoadFromFile(obj uintptr, Filename string) {
	getLazyProc("Icon_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func Icon_SaveToFile(obj uintptr, Filename string) {
	getLazyProc("Icon_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func Icon_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Icon_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Icon_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Icon_ClassType").Call(obj)
	return TClass(ret)
}

func Icon_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Icon_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Icon_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Icon_InstanceSize").Call(obj)
	return int32(ret)
}

func Icon_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Icon_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Icon_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Icon_GetHashCode").Call(obj)
	return int32(ret)
}

func Icon_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Icon_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Icon_GetHandle(obj uintptr) HICON {
	ret, _, _ := getLazyProc("Icon_GetHandle").Call(obj)
	return HICON(ret)
}

func Icon_SetHandle(obj uintptr, value HICON) {
	getLazyProc("Icon_SetHandle").Call(obj, uintptr(value))
}

func Icon_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func Icon_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Icon_GetHeight").Call(obj)
	return int32(ret)
}

func Icon_SetHeight(obj uintptr, value int32) {
	getLazyProc("Icon_SetHeight").Call(obj, uintptr(value))
}

func Icon_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Icon_SetModified(obj uintptr, value bool) {
	getLazyProc("Icon_SetModified").Call(obj, GoBoolToDBool(value))
}

func Icon_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("Icon_GetPalette").Call(obj)
	return HPALETTE(ret)
}

func Icon_SetPalette(obj uintptr, value HPALETTE) {
	getLazyProc("Icon_SetPalette").Call(obj, uintptr(value))
}

func Icon_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Icon_SetPaletteModified(obj uintptr, value bool) {
	getLazyProc("Icon_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func Icon_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Icon_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func Icon_SetTransparent(obj uintptr, value bool) {
	getLazyProc("Icon_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func Icon_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Icon_GetWidth").Call(obj)
	return int32(ret)
}

func Icon_SetWidth(obj uintptr, value int32) {
	getLazyProc("Icon_SetWidth").Call(obj, uintptr(value))
}

func Icon_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("Icon_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Icon_StaticClassType() TClass {
	r, _, _ := getLazyProc("Icon_StaticClassType").Call()
	return TClass(r)
}
