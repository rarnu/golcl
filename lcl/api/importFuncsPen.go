package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPen ---------------------------

func Pen_Create() uintptr {
	ret, _, _ := getLazyProc("Pen_Create").Call()
	return ret
}

func Pen_Free(obj uintptr) {
	_, _, _ = getLazyProc("Pen_Free").Call(obj)
}

func Pen_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Pen_Assign").Call(obj, Source)
}

func Pen_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Pen_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Pen_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Pen_ClassType").Call(obj)
	return TClass(ret)
}

func Pen_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Pen_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Pen_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Pen_InstanceSize").Call(obj)
	return int32(ret)
}

func Pen_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Pen_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Pen_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Pen_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Pen_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Pen_GetHashCode").Call(obj)
	return int32(ret)
}

func Pen_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Pen_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Pen_GetHandle(obj uintptr) HPEN {
	ret, _, _ := getLazyProc("Pen_GetHandle").Call(obj)
	return ret
}

func Pen_SetHandle(obj uintptr, value HPEN) {
	_, _, _ = getLazyProc("Pen_SetHandle").Call(obj, value)
}

func Pen_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Pen_GetColor").Call(obj)
	return TColor(ret)
}

func Pen_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Pen_SetColor").Call(obj, uintptr(value))
}

func Pen_GetMode(obj uintptr) TPenMode {
	ret, _, _ := getLazyProc("Pen_GetMode").Call(obj)
	return TPenMode(ret)
}

func Pen_SetMode(obj uintptr, value TPenMode) {
	_, _, _ = getLazyProc("Pen_SetMode").Call(obj, uintptr(value))
}

func Pen_GetStyle(obj uintptr) TPenStyle {
	ret, _, _ := getLazyProc("Pen_GetStyle").Call(obj)
	return TPenStyle(ret)
}

func Pen_SetStyle(obj uintptr, value TPenStyle) {
	_, _, _ = getLazyProc("Pen_SetStyle").Call(obj, uintptr(value))
}

func Pen_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Pen_GetWidth").Call(obj)
	return int32(ret)
}

func Pen_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Pen_SetWidth").Call(obj, uintptr(value))
}

func Pen_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Pen_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Pen_StaticClassType() TClass {
	r, _, _ := getLazyProc("Pen_StaticClassType").Call()
	return TClass(r)
}
