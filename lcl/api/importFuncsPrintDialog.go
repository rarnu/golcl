package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPrintDialog ---------------------------

func PrintDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PrintDialog_Create").Call(obj)
	return ret
}

func PrintDialog_Free(obj uintptr) {
	getLazyProc("PrintDialog_Free").Call(obj)
}

func PrintDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrintDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func PrintDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PrintDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PrintDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PrintDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PrintDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrintDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PrintDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("PrintDialog_Assign").Call(obj, Source)
}

func PrintDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PrintDialog_ClassType").Call(obj)
	return TClass(ret)
}

func PrintDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PrintDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PrintDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func PrintDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PrintDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PrintDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PrintDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PrintDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func PrintDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PrintDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PrintDialog_GetCollate(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrintDialog_GetCollate").Call(obj)
	return DBoolToGoBool(ret)
}

func PrintDialog_SetCollate(obj uintptr, value bool) {
	getLazyProc("PrintDialog_SetCollate").Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetCopies(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetCopies").Call(obj)
	return int32(ret)
}

func PrintDialog_SetCopies(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetCopies").Call(obj, uintptr(value))
}

func PrintDialog_GetFromPage(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetFromPage").Call(obj)
	return int32(ret)
}

func PrintDialog_SetFromPage(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetFromPage").Call(obj, uintptr(value))
}

func PrintDialog_GetMinPage(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetMinPage").Call(obj)
	return int32(ret)
}

func PrintDialog_SetMinPage(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetMinPage").Call(obj, uintptr(value))
}

func PrintDialog_GetMaxPage(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetMaxPage").Call(obj)
	return int32(ret)
}

func PrintDialog_SetMaxPage(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetMaxPage").Call(obj, uintptr(value))
}

func PrintDialog_GetOptions(obj uintptr) TPrintDialogOptions {
	ret, _, _ := getLazyProc("PrintDialog_GetOptions").Call(obj)
	return TPrintDialogOptions(ret)
}

func PrintDialog_SetOptions(obj uintptr, value TPrintDialogOptions) {
	getLazyProc("PrintDialog_SetOptions").Call(obj, uintptr(value))
}

func PrintDialog_GetPrintToFile(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrintDialog_GetPrintToFile").Call(obj)
	return DBoolToGoBool(ret)
}

func PrintDialog_SetPrintToFile(obj uintptr, value bool) {
	getLazyProc("PrintDialog_SetPrintToFile").Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetPrintRange(obj uintptr) TPrintRange {
	ret, _, _ := getLazyProc("PrintDialog_GetPrintRange").Call(obj)
	return TPrintRange(ret)
}

func PrintDialog_SetPrintRange(obj uintptr, value TPrintRange) {
	getLazyProc("PrintDialog_SetPrintRange").Call(obj, uintptr(value))
}

func PrintDialog_GetToPage(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetToPage").Call(obj)
	return int32(ret)
}

func PrintDialog_SetToPage(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetToPage").Call(obj, uintptr(value))
}

func PrintDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("PrintDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func PrintDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("PrintDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func PrintDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("PrintDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func PrintDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func PrintDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrintDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PrintDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("PrintDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func PrintDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PrintDialog_GetOwner").Call(obj)
	return ret
}

func PrintDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PrintDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PrintDialog_SetName(obj uintptr, value string) {
	getLazyProc("PrintDialog_SetName").Call(obj, GoStrToDStr(value))
}

func PrintDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PrintDialog_GetTag").Call(obj)
	return int(ret)
}

func PrintDialog_SetTag(obj uintptr, value int) {
	getLazyProc("PrintDialog_SetTag").Call(obj, uintptr(value))
}

func PrintDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PrintDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PrintDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("PrintDialog_StaticClassType").Call()
	return TClass(r)
}
