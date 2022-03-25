package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPicture ---------------------------

func Picture_Create() uintptr {
	ret, _, _ := getLazyProc("Picture_Create").Call()
	return ret
}

func Picture_Free(obj uintptr) {
	_, _, _ = getLazyProc("Picture_Free").Call(obj)
}

func Picture_LoadFromFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Picture_LoadFromFile").Call(obj, GoStrToDStr(Filename))
}

func Picture_SaveToFile(obj uintptr, Filename string) {
	_, _, _ = getLazyProc("Picture_SaveToFile").Call(obj, GoStrToDStr(Filename))
}

func Picture_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Picture_LoadFromStream").Call(obj, Stream)
}

func Picture_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Picture_SaveToStream").Call(obj, Stream)
}

func Picture_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Picture_Assign").Call(obj, Source)
}

func Picture_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Picture_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Picture_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Picture_ClassType").Call(obj)
	return TClass(ret)
}

func Picture_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Picture_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Picture_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Picture_InstanceSize").Call(obj)
	return int32(ret)
}

func Picture_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Picture_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Picture_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Picture_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Picture_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Picture_GetHashCode").Call(obj)
	return int32(ret)
}

func Picture_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Picture_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Picture_GetBitmap(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Picture_GetBitmap").Call(obj)
	return ret
}

func Picture_SetBitmap(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Picture_SetBitmap").Call(obj, value)
}

func Picture_GetGraphic(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Picture_GetGraphic").Call(obj)
	return ret
}

func Picture_SetGraphic(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Picture_SetGraphic").Call(obj, value)
}

func Picture_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Picture_GetHeight").Call(obj)
	return int32(ret)
}

func Picture_GetIcon(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Picture_GetIcon").Call(obj)
	return ret
}

func Picture_SetIcon(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Picture_SetIcon").Call(obj, value)
}

func Picture_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Picture_GetWidth").Call(obj)
	return int32(ret)
}

func Picture_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Picture_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Picture_StaticClassType() TClass {
	r, _, _ := getLazyProc("Picture_StaticClassType").Call()
	return TClass(r)
}
