package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TFindDialog ---------------------------

func FindDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FindDialog_Create").Call(obj)
	return ret
}

func FindDialog_Free(obj uintptr) {
	_, _, _ = getLazyProc("FindDialog_Free").Call(obj)
}

func FindDialog_CloseDialog(obj uintptr) {
	_, _, _ = getLazyProc("FindDialog_CloseDialog").Call(obj)
}

func FindDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("FindDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func FindDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("FindDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func FindDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("FindDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func FindDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("FindDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func FindDialog_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("FindDialog_Assign").Call(obj, Source)
}

func FindDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("FindDialog_ClassType").Call(obj)
	return TClass(ret)
}

func FindDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("FindDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func FindDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func FindDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("FindDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func FindDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("FindDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func FindDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func FindDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("FindDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func FindDialog_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_GetLeft").Call(obj)
	return int32(ret)
}

func FindDialog_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FindDialog_SetLeft").Call(obj, uintptr(value))
}

func FindDialog_GetPosition(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FindDialog_GetPosition").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FindDialog_SetPosition(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("FindDialog_SetPosition").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FindDialog_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_GetTop").Call(obj)
	return int32(ret)
}

func FindDialog_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FindDialog_SetTop").Call(obj, uintptr(value))
}

func FindDialog_GetFindText(obj uintptr) string {
	ret, _, _ := getLazyProc("FindDialog_GetFindText").Call(obj)
	return DStrToGoStr(ret)
}

func FindDialog_SetFindText(obj uintptr, value string) {
	_, _, _ = getLazyProc("FindDialog_SetFindText").Call(obj, GoStrToDStr(value))
}

func FindDialog_GetOptions(obj uintptr) TFindOptions {
	ret, _, _ := getLazyProc("FindDialog_GetOptions").Call(obj)
	return TFindOptions(ret)
}

func FindDialog_SetOptions(obj uintptr, value TFindOptions) {
	_, _, _ = getLazyProc("FindDialog_SetOptions").Call(obj, uintptr(value))
}

func FindDialog_SetOnFind(obj uintptr, fn any) {
	_, _, _ = getLazyProc("FindDialog_SetOnFind").Call(obj, addEventToMap(obj, fn))
}

func FindDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FindDialog_GetHandle").Call(obj)
	return ret
}

func FindDialog_SetOnClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("FindDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func FindDialog_SetOnShow(obj uintptr, fn any) {
	_, _, _ = getLazyProc("FindDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func FindDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func FindDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FindDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func FindDialog_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FindDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func FindDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FindDialog_GetOwner").Call(obj)
	return ret
}

func FindDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("FindDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func FindDialog_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("FindDialog_SetName").Call(obj, GoStrToDStr(value))
}

func FindDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("FindDialog_GetTag").Call(obj)
	return int(ret)
}

func FindDialog_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("FindDialog_SetTag").Call(obj, uintptr(value))
}

func FindDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("FindDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func FindDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("FindDialog_StaticClassType").Call()
	return TClass(r)
}
