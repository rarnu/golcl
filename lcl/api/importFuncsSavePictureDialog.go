package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSavePictureDialog ---------------------------

func SavePictureDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SavePictureDialog_Create").Call(obj)
	return ret
}

func SavePictureDialog_Free(obj uintptr) {
	getLazyProc("SavePictureDialog_Free").Call(obj)
}

func SavePictureDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("SavePictureDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func SavePictureDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SavePictureDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SavePictureDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SavePictureDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SavePictureDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("SavePictureDialog_Assign").Call(obj, Source)
}

func SavePictureDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SavePictureDialog_ClassType").Call(obj)
	return TClass(ret)
}

func SavePictureDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SavePictureDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func SavePictureDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SavePictureDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SavePictureDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SavePictureDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SavePictureDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SavePictureDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func SavePictureDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_GetFilter(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetFilter").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetFilter(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetFilter").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFiles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SavePictureDialog_GetFiles").Call(obj)
	return ret
}

func SavePictureDialog_GetDefaultExt(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetDefaultExt").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetDefaultExt(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetDefaultExt").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetFileName(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetFileName").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFilterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SavePictureDialog_GetFilterIndex").Call(obj)
	return int32(ret)
}

func SavePictureDialog_SetFilterIndex(obj uintptr, value int32) {
	getLazyProc("SavePictureDialog_SetFilterIndex").Call(obj, uintptr(value))
}

func SavePictureDialog_GetInitialDir(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetInitialDir").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetInitialDir(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetInitialDir").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("SavePictureDialog_GetOptions").Call(obj)
	return TOpenOptions(ret)
}

func SavePictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
	getLazyProc("SavePictureDialog_SetOptions").Call(obj, uintptr(value))
}

func SavePictureDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetTitle(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SavePictureDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func SavePictureDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("SavePictureDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func SavePictureDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("SavePictureDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func SavePictureDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SavePictureDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func SavePictureDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SavePictureDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SavePictureDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("SavePictureDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func SavePictureDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SavePictureDialog_GetOwner").Call(obj)
	return ret
}

func SavePictureDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SavePictureDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SavePictureDialog_SetName(obj uintptr, value string) {
	getLazyProc("SavePictureDialog_SetName").Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SavePictureDialog_GetTag").Call(obj)
	return int(ret)
}

func SavePictureDialog_SetTag(obj uintptr, value int) {
	getLazyProc("SavePictureDialog_SetTag").Call(obj, uintptr(value))
}

func SavePictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SavePictureDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SavePictureDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("SavePictureDialog_StaticClassType").Call()
	return TClass(r)
}
