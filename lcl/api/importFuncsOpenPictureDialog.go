package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TOpenPictureDialog ---------------------------

func OpenPictureDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenPictureDialog_Create").Call(obj)
	return ret
}

func OpenPictureDialog_Free(obj uintptr) {
	_, _, _ = getLazyProc("OpenPictureDialog_Free").Call(obj)
}

func OpenPictureDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenPictureDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func OpenPictureDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("OpenPictureDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func OpenPictureDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenPictureDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func OpenPictureDialog_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("OpenPictureDialog_Assign").Call(obj, Source)
}

func OpenPictureDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("OpenPictureDialog_ClassType").Call(obj)
	return TClass(ret)
}

func OpenPictureDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenPictureDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func OpenPictureDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("OpenPictureDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func OpenPictureDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("OpenPictureDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func OpenPictureDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func OpenPictureDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_GetFilter(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetFilter").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFilter(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetFilter").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFiles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetFiles").Call(obj)
	return ret
}

func OpenPictureDialog_GetDefaultExt(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetDefaultExt").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetDefaultExt(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetDefaultExt").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFileName(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetFileName").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFilterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetFilterIndex").Call(obj)
	return int32(ret)
}

func OpenPictureDialog_SetFilterIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetFilterIndex").Call(obj, uintptr(value))
}

func OpenPictureDialog_GetInitialDir(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetInitialDir").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetInitialDir(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetInitialDir").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetOptions").Call(obj)
	return TOpenOptions(ret)
}

func OpenPictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetOptions").Call(obj, uintptr(value))
}

func OpenPictureDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetHandle").Call(obj)
	return ret
}

func OpenPictureDialog_SetOnClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func OpenPictureDialog_SetOnShow(obj uintptr, fn any) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func OpenPictureDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func OpenPictureDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func OpenPictureDialog_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func OpenPictureDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetOwner").Call(obj)
	return ret
}

func OpenPictureDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func OpenPictureDialog_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetName").Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetTag").Call(obj)
	return int(ret)
}

func OpenPictureDialog_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("OpenPictureDialog_SetTag").Call(obj, uintptr(value))
}

func OpenPictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("OpenPictureDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func OpenPictureDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("OpenPictureDialog_StaticClassType").Call()
	return TClass(r)
}
