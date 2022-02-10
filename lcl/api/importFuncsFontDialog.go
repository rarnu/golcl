package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TFontDialog ---------------------------

func FontDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FontDialog_Create").Call(obj)
	return ret
}

func FontDialog_Free(obj uintptr) {
	getLazyProc("FontDialog_Free").Call(obj)
}

func FontDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("FontDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func FontDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("FontDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func FontDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("FontDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func FontDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("FontDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func FontDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("FontDialog_Assign").Call(obj, Source)
}

func FontDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("FontDialog_ClassType").Call(obj)
	return TClass(ret)
}

func FontDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("FontDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func FontDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FontDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func FontDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("FontDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func FontDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("FontDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func FontDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FontDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func FontDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("FontDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func FontDialog_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FontDialog_GetFont").Call(obj)
	return ret
}

func FontDialog_SetFont(obj uintptr, value uintptr) {
	getLazyProc("FontDialog_SetFont").Call(obj, value)
}

func FontDialog_GetOptions(obj uintptr) TFontDialogOptions {
	ret, _, _ := getLazyProc("FontDialog_GetOptions").Call(obj)
	return TFontDialogOptions(ret)
}

func FontDialog_SetOptions(obj uintptr, value TFontDialogOptions) {
	getLazyProc("FontDialog_SetOptions").Call(obj, uintptr(value))
}

func FontDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FontDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func FontDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("FontDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func FontDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("FontDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func FontDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FontDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func FontDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FontDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func FontDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("FontDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func FontDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FontDialog_GetOwner").Call(obj)
	return ret
}

func FontDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("FontDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func FontDialog_SetName(obj uintptr, value string) {
	getLazyProc("FontDialog_SetName").Call(obj, GoStrToDStr(value))
}

func FontDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("FontDialog_GetTag").Call(obj)
	return int(ret)
}

func FontDialog_SetTag(obj uintptr, value int) {
	getLazyProc("FontDialog_SetTag").Call(obj, uintptr(value))
}

func FontDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("FontDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func FontDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("FontDialog_StaticClassType").Call()
	return TClass(r)
}
