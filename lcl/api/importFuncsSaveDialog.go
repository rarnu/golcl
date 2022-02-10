package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSaveDialog ---------------------------

func SaveDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SaveDialog_Create").Call(obj)
	return ret
}

func SaveDialog_Free(obj uintptr) {
	getLazyProc("SaveDialog_Free").Call(obj)
}

func SaveDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("SaveDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func SaveDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SaveDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SaveDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SaveDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SaveDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("SaveDialog_Assign").Call(obj, Source)
}

func SaveDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SaveDialog_ClassType").Call(obj)
	return TClass(ret)
}

func SaveDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SaveDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func SaveDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SaveDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SaveDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SaveDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SaveDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SaveDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func SaveDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_GetFiles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SaveDialog_GetFiles").Call(obj)
	return ret
}

func SaveDialog_GetDefaultExt(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetDefaultExt").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetDefaultExt(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetDefaultExt").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetFileName(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetFileName").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilter(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetFilter").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetFilter(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetFilter").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SaveDialog_GetFilterIndex").Call(obj)
	return int32(ret)
}

func SaveDialog_SetFilterIndex(obj uintptr, value int32) {
	getLazyProc("SaveDialog_SetFilterIndex").Call(obj, uintptr(value))
}

func SaveDialog_GetInitialDir(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetInitialDir").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetInitialDir(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetInitialDir").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("SaveDialog_GetOptions").Call(obj)
	return TOpenOptions(ret)
}

func SaveDialog_SetOptions(obj uintptr, value TOpenOptions) {
	getLazyProc("SaveDialog_SetOptions").Call(obj, uintptr(value))
}

func SaveDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetTitle(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SaveDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func SaveDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("SaveDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func SaveDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("SaveDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func SaveDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SaveDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func SaveDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SaveDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SaveDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("SaveDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func SaveDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SaveDialog_GetOwner").Call(obj)
	return ret
}

func SaveDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SaveDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SaveDialog_SetName(obj uintptr, value string) {
	getLazyProc("SaveDialog_SetName").Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SaveDialog_GetTag").Call(obj)
	return int(ret)
}

func SaveDialog_SetTag(obj uintptr, value int) {
	getLazyProc("SaveDialog_SetTag").Call(obj, uintptr(value))
}

func SaveDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SaveDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SaveDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("SaveDialog_StaticClassType").Call()
	return TClass(r)
}
