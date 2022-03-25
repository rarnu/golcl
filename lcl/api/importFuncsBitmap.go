package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TBitmap ---------------------------

func Bitmap_Create() uintptr {
	ret, _, _ := getLazyProc("Bitmap_Create").Call()
	return ret
}

func Bitmap_Free(obj uintptr) {
	_, _, _ = getLazyProc("Bitmap_Free").Call(obj)
}

func Bitmap_LoadFromDevice(obj uintptr, ADc HDC) {
	_, _, _ = getLazyProc("Bitmap_LoadFromDevice").Call(obj, ADc)
}

func Bitmap_EndUpdate(obj uintptr, AStreamIsValid bool) {
	_, _, _ = getLazyProc("Bitmap_EndUpdate").Call(obj, GoBoolToDBool(AStreamIsValid))
}

func Bitmap_BeginUpdate(obj uintptr, ACanvasOnly bool) {
	_, _, _ = getLazyProc("Bitmap_BeginUpdate").Call(obj, GoBoolToDBool(ACanvasOnly))
}

func Bitmap_Clear(obj uintptr) {
	_, _, _ = getLazyProc("Bitmap_Clear").Call(obj)
}

func Bitmap_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Bitmap_Assign").Call(obj, Source)
}

func Bitmap_FreeImage(obj uintptr) {
	_, _, _ = getLazyProc("Bitmap_FreeImage").Call(obj)
}

func Bitmap_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Bitmap_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Bitmap_LoadFromStream").Call(obj, Stream)
}

func Bitmap_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Bitmap_SaveToStream").Call(obj, Stream)
}

func Bitmap_SetSize(obj uintptr, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Bitmap_SetSize").Call(obj, uintptr(AWidth), uintptr(AHeight))
}

func Bitmap_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string) {
	_, _, _ = getLazyProc("Bitmap_LoadFromResourceName").Call(obj, Instance, GoStrToDStr(ResName))
}

func Bitmap_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32) {
	_, _, _ = getLazyProc("Bitmap_LoadFromResourceID").Call(obj, Instance, uintptr(ResID))
}

func Bitmap_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Bitmap_LoadFromFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Bitmap_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func Bitmap_SaveToFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Bitmap_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func Bitmap_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Bitmap_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Bitmap_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Bitmap_ClassType").Call(obj)
	return TClass(ret)
}

func Bitmap_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Bitmap_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Bitmap_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bitmap_InstanceSize").Call(obj)
	return int32(ret)
}

func Bitmap_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Bitmap_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Bitmap_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bitmap_GetHashCode").Call(obj)
	return int32(ret)
}

func Bitmap_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Bitmap_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Bitmap_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bitmap_GetCanvas").Call(obj)
	return ret
}

func Bitmap_GetHandle(obj uintptr) HBITMAP {
	ret, _, _ := getLazyProc("Bitmap_GetHandle").Call(obj)
	return ret
}

func Bitmap_SetHandle(obj uintptr, value HBITMAP) {
	_, _, _ = getLazyProc("Bitmap_SetHandle").Call(obj, value)
}

func Bitmap_GetHandleType(obj uintptr) TBitmapHandleType {
	ret, _, _ := getLazyProc("Bitmap_GetHandleType").Call(obj)
	return TBitmapHandleType(ret)
}

func Bitmap_SetHandleType(obj uintptr, value TBitmapHandleType) {
	_, _, _ = getLazyProc("Bitmap_SetHandleType").Call(obj, uintptr(value))
}

func Bitmap_GetMaskHandle(obj uintptr) HBITMAP {
	ret, _, _ := getLazyProc("Bitmap_GetMaskHandle").Call(obj)
	return ret
}

func Bitmap_SetMaskHandle(obj uintptr, value HBITMAP) {
	_, _, _ = getLazyProc("Bitmap_SetMaskHandle").Call(obj, value)
}

func Bitmap_GetPixelFormat(obj uintptr) TPixelFormat {
	ret, _, _ := getLazyProc("Bitmap_GetPixelFormat").Call(obj)
	return TPixelFormat(ret)
}

func Bitmap_SetPixelFormat(obj uintptr, value TPixelFormat) {
	_, _, _ = getLazyProc("Bitmap_SetPixelFormat").Call(obj, uintptr(value))
}

func Bitmap_GetTransparentMode(obj uintptr) TTransparentMode {
	ret, _, _ := getLazyProc("Bitmap_GetTransparentMode").Call(obj)
	return TTransparentMode(ret)
}

func Bitmap_SetTransparentMode(obj uintptr, value TTransparentMode) {
	_, _, _ = getLazyProc("Bitmap_SetTransparentMode").Call(obj, uintptr(value))
}

func Bitmap_GetEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_GetEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func Bitmap_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bitmap_GetHeight").Call(obj)
	return int32(ret)
}

func Bitmap_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bitmap_SetHeight").Call(obj, uintptr(value))
}

func Bitmap_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Bitmap_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bitmap_SetModified").Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetPalette(obj uintptr) HPALETTE {
	ret, _, _ := getLazyProc("Bitmap_GetPalette").Call(obj)
	return ret
}

func Bitmap_SetPalette(obj uintptr, value HPALETTE) {
	_, _, _ = getLazyProc("Bitmap_SetPalette").Call(obj, value)
}

func Bitmap_GetPaletteModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_GetPaletteModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Bitmap_SetPaletteModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bitmap_SetPaletteModified").Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bitmap_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func Bitmap_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bitmap_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bitmap_GetWidth").Call(obj)
	return int32(ret)
}

func Bitmap_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bitmap_SetWidth").Call(obj, uintptr(value))
}

func Bitmap_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Bitmap_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Bitmap_GetScanLine(obj uintptr, Row int32) uintptr {
	ret, _, _ := getLazyProc("Bitmap_GetScanLine").Call(obj, uintptr(Row))
	return ret
}

func Bitmap_StaticClassType() TClass {
	r, _, _ := getLazyProc("Bitmap_StaticClassType").Call()
	return TClass(r)
}
