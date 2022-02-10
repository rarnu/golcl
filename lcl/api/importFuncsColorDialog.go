package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TColorDialog ---------------------------

func ColorDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorDialog_Create").Call(obj)
	return ret
}

func ColorDialog_Free(obj uintptr) {
	getLazyProc("ColorDialog_Free").Call(obj)
}

func ColorDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ColorDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ColorDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ColorDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ColorDialog_Assign").Call(obj, Source)
}

func ColorDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ColorDialog_ClassType").Call(obj)
	return TClass(ret)
}

func ColorDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func ColorDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ColorDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ColorDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ColorDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func ColorDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ColorDialog_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorDialog_GetColor").Call(obj)
	return TColor(ret)
}

func ColorDialog_SetColor(obj uintptr, value TColor) {
	getLazyProc("ColorDialog_SetColor").Call(obj, uintptr(value))
}

func ColorDialog_GetCustomColors(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorDialog_GetCustomColors").Call(obj)
	return ret
}

func ColorDialog_SetCustomColors(obj uintptr, value uintptr) {
	getLazyProc("ColorDialog_SetCustomColors").Call(obj, value)
}

func ColorDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ColorDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func ColorDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("ColorDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func ColorDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("ColorDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func ColorDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func ColorDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ColorDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ColorDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func ColorDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorDialog_GetOwner").Call(obj)
	return ret
}

func ColorDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorDialog_SetName(obj uintptr, value string) {
	getLazyProc("ColorDialog_SetName").Call(obj, GoStrToDStr(value))
}

func ColorDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ColorDialog_GetTag").Call(obj)
	return int(ret)
}

func ColorDialog_SetTag(obj uintptr, value int) {
	getLazyProc("ColorDialog_SetTag").Call(obj, uintptr(value))
}

func ColorDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ColorDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ColorDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("ColorDialog_StaticClassType").Call()
	return TClass(r)
}
