package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSelectDirectoryDialog ---------------------------

func SelectDirectoryDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_Create").Call(obj)
	return ret
}

func SelectDirectoryDialog_Free(obj uintptr) {
	getLazyProc("SelectDirectoryDialog_Free").Call(obj)
}

func SelectDirectoryDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func SelectDirectoryDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SelectDirectoryDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SelectDirectoryDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("SelectDirectoryDialog_Assign").Call(obj, Source)
}

func SelectDirectoryDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_ClassType").Call(obj)
	return TClass(ret)
}

func SelectDirectoryDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func SelectDirectoryDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SelectDirectoryDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SelectDirectoryDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func SelectDirectoryDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_GetFiles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetFiles").Call(obj)
	return ret
}

func SelectDirectoryDialog_GetDefaultExt(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetDefaultExt").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetDefaultExt(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetDefaultExt").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetFileName(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetFileName").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetFilter(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetFilter").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetFilter(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetFilter").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetFilterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetFilterIndex").Call(obj)
	return int32(ret)
}

func SelectDirectoryDialog_SetFilterIndex(obj uintptr, value int32) {
	getLazyProc("SelectDirectoryDialog_SetFilterIndex").Call(obj, uintptr(value))
}

func SelectDirectoryDialog_GetInitialDir(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetInitialDir").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetInitialDir(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetInitialDir").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetOptions").Call(obj)
	return TOpenOptions(ret)
}

func SelectDirectoryDialog_SetOptions(obj uintptr, value TOpenOptions) {
	getLazyProc("SelectDirectoryDialog_SetOptions").Call(obj, uintptr(value))
}

func SelectDirectoryDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetTitle(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func SelectDirectoryDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("SelectDirectoryDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func SelectDirectoryDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("SelectDirectoryDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func SelectDirectoryDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func SelectDirectoryDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SelectDirectoryDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("SelectDirectoryDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func SelectDirectoryDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetOwner").Call(obj)
	return ret
}

func SelectDirectoryDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SelectDirectoryDialog_SetName(obj uintptr, value string) {
	getLazyProc("SelectDirectoryDialog_SetName").Call(obj, GoStrToDStr(value))
}

func SelectDirectoryDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetTag").Call(obj)
	return int(ret)
}

func SelectDirectoryDialog_SetTag(obj uintptr, value int) {
	getLazyProc("SelectDirectoryDialog_SetTag").Call(obj, uintptr(value))
}

func SelectDirectoryDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SelectDirectoryDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SelectDirectoryDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("SelectDirectoryDialog_StaticClassType").Call()
	return TClass(r)
}
