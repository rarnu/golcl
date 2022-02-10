package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TOpenDialog ---------------------------

func OpenDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenDialog_Create").Call(obj)
	return ret
}

func OpenDialog_Free(obj uintptr) {
	getLazyProc("OpenDialog_Free").Call(obj)
}

func OpenDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func OpenDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("OpenDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func OpenDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func OpenDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("OpenDialog_Assign").Call(obj, Source)
}

func OpenDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("OpenDialog_ClassType").Call(obj)
	return TClass(ret)
}

func OpenDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func OpenDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("OpenDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func OpenDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func OpenDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func OpenDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_GetFiles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenDialog_GetFiles").Call(obj)
	return ret
}

func OpenDialog_GetDefaultExt(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetDefaultExt").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetDefaultExt(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetDefaultExt").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetFileName(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetFileName").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilter(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetFilter").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetFilter(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetFilter").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenDialog_GetFilterIndex").Call(obj)
	return int32(ret)
}

func OpenDialog_SetFilterIndex(obj uintptr, value int32) {
	getLazyProc("OpenDialog_SetFilterIndex").Call(obj, uintptr(value))
}

func OpenDialog_GetInitialDir(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetInitialDir").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetInitialDir(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetInitialDir").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("OpenDialog_GetOptions").Call(obj)
	return TOpenOptions(ret)
}

func OpenDialog_SetOptions(obj uintptr, value TOpenOptions) {
	getLazyProc("OpenDialog_SetOptions").Call(obj, uintptr(value))
}

func OpenDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetTitle(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("OpenDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func OpenDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("OpenDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func OpenDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("OpenDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func OpenDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func OpenDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func OpenDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("OpenDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func OpenDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenDialog_GetOwner").Call(obj)
	return ret
}

func OpenDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenDialog_SetName(obj uintptr, value string) {
	getLazyProc("OpenDialog_SetName").Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("OpenDialog_GetTag").Call(obj)
	return int(ret)
}

func OpenDialog_SetTag(obj uintptr, value int) {
	getLazyProc("OpenDialog_SetTag").Call(obj, uintptr(value))
}

func OpenDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("OpenDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func OpenDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("OpenDialog_StaticClassType").Call()
	return TClass(r)
}
