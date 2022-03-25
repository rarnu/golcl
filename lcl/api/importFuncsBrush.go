package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TBrush ---------------------------

func Brush_Create() uintptr {
	ret, _, _ := getLazyProc("Brush_Create").Call()
	return ret
}

func Brush_Free(obj uintptr) {
	_, _, _ = getLazyProc("Brush_Free").Call(obj)
}

func Brush_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Brush_Assign").Call(obj, Source)
}

func Brush_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Brush_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Brush_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Brush_ClassType").Call(obj)
	return TClass(ret)
}

func Brush_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Brush_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Brush_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Brush_InstanceSize").Call(obj)
	return int32(ret)
}

func Brush_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Brush_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Brush_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Brush_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Brush_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Brush_GetHashCode").Call(obj)
	return int32(ret)
}

func Brush_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Brush_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Brush_GetBitmap(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Brush_GetBitmap").Call(obj)
	return ret
}

func Brush_SetBitmap(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Brush_SetBitmap").Call(obj, value)
}

func Brush_GetHandle(obj uintptr) HBRUSH {
	ret, _, _ := getLazyProc("Brush_GetHandle").Call(obj)
	return ret
}

func Brush_SetHandle(obj uintptr, value HBRUSH) {
	_, _, _ = getLazyProc("Brush_SetHandle").Call(obj, value)
}

func Brush_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Brush_GetColor").Call(obj)
	return TColor(ret)
}

func Brush_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Brush_SetColor").Call(obj, uintptr(value))
}

func Brush_GetStyle(obj uintptr) TBrushStyle {
	ret, _, _ := getLazyProc("Brush_GetStyle").Call(obj)
	return TBrushStyle(ret)
}

func Brush_SetStyle(obj uintptr, value TBrushStyle) {
	_, _, _ = getLazyProc("Brush_SetStyle").Call(obj, uintptr(value))
}

func Brush_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Brush_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Brush_StaticClassType() TClass {
	r, _, _ := getLazyProc("Brush_StaticClassType").Call()
	return TClass(r)
}
