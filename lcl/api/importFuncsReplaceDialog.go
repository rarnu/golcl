package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TReplaceDialog ---------------------------

func ReplaceDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ReplaceDialog_Create").Call(obj)
	return ret
}

func ReplaceDialog_Free(obj uintptr) {
	getLazyProc("ReplaceDialog_Free").Call(obj)
}

func ReplaceDialog_CloseDialog(obj uintptr) {
	getLazyProc("ReplaceDialog_CloseDialog").Call(obj)
}

func ReplaceDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("ReplaceDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func ReplaceDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ReplaceDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ReplaceDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ReplaceDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ReplaceDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ReplaceDialog_Assign").Call(obj, Source)
}

func ReplaceDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ReplaceDialog_ClassType").Call(obj)
	return TClass(ret)
}

func ReplaceDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func ReplaceDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ReplaceDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ReplaceDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ReplaceDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ReplaceDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func ReplaceDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_GetReplaceText(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_GetReplaceText").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_SetReplaceText(obj uintptr, value string) {
	getLazyProc("ReplaceDialog_SetReplaceText").Call(obj, GoStrToDStr(value))
}

func ReplaceDialog_SetOnReplace(obj uintptr, fn interface{}) {
	getLazyProc("ReplaceDialog_SetOnReplace").Call(obj, addEventToMap(obj, fn))
}

func ReplaceDialog_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_GetLeft").Call(obj)
	return int32(ret)
}

func ReplaceDialog_SetLeft(obj uintptr, value int32) {
	getLazyProc("ReplaceDialog_SetLeft").Call(obj, uintptr(value))
}

func ReplaceDialog_GetPosition(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ReplaceDialog_GetPosition").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ReplaceDialog_SetPosition(obj uintptr, value TPoint) {
	getLazyProc("ReplaceDialog_SetPosition").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ReplaceDialog_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_GetTop").Call(obj)
	return int32(ret)
}

func ReplaceDialog_SetTop(obj uintptr, value int32) {
	getLazyProc("ReplaceDialog_SetTop").Call(obj, uintptr(value))
}

func ReplaceDialog_GetFindText(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_GetFindText").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_SetFindText(obj uintptr, value string) {
	getLazyProc("ReplaceDialog_SetFindText").Call(obj, GoStrToDStr(value))
}

func ReplaceDialog_GetOptions(obj uintptr) TFindOptions {
	ret, _, _ := getLazyProc("ReplaceDialog_GetOptions").Call(obj)
	return TFindOptions(ret)
}

func ReplaceDialog_SetOptions(obj uintptr, value TFindOptions) {
	getLazyProc("ReplaceDialog_SetOptions").Call(obj, uintptr(value))
}

func ReplaceDialog_SetOnFind(obj uintptr, fn interface{}) {
	getLazyProc("ReplaceDialog_SetOnFind").Call(obj, addEventToMap(obj, fn))
}

func ReplaceDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ReplaceDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func ReplaceDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("ReplaceDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func ReplaceDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("ReplaceDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func ReplaceDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func ReplaceDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ReplaceDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ReplaceDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ReplaceDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func ReplaceDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ReplaceDialog_GetOwner").Call(obj)
	return ret
}

func ReplaceDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ReplaceDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ReplaceDialog_SetName(obj uintptr, value string) {
	getLazyProc("ReplaceDialog_SetName").Call(obj, GoStrToDStr(value))
}

func ReplaceDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ReplaceDialog_GetTag").Call(obj)
	return int(ret)
}

func ReplaceDialog_SetTag(obj uintptr, value int) {
	getLazyProc("ReplaceDialog_SetTag").Call(obj, uintptr(value))
}

func ReplaceDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ReplaceDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ReplaceDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("ReplaceDialog_StaticClassType").Call()
	return TClass(r)
}
