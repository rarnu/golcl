package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPageSetupDialog ---------------------------

func PageSetupDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageSetupDialog_Create").Call(obj)
	return ret
}

func PageSetupDialog_Free(obj uintptr) {
	getLazyProc("PageSetupDialog_Free").Call(obj)
}

func PageSetupDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageSetupDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func PageSetupDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PageSetupDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PageSetupDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PageSetupDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PageSetupDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageSetupDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PageSetupDialog_Assign(obj uintptr, Source uintptr) {
	getLazyProc("PageSetupDialog_Assign").Call(obj, Source)
}

func PageSetupDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PageSetupDialog_ClassType").Call(obj)
	return TClass(ret)
}

func PageSetupDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PageSetupDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PageSetupDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func PageSetupDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PageSetupDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PageSetupDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PageSetupDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PageSetupDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func PageSetupDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PageSetupDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PageSetupDialog_GetMarginLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetMarginLeft").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetMarginLeft(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetMarginLeft").Call(obj, uintptr(value))
}

func PageSetupDialog_GetMarginTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetMarginTop").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetMarginTop(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetMarginTop").Call(obj, uintptr(value))
}

func PageSetupDialog_GetMarginRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetMarginRight").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetMarginRight(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetMarginRight").Call(obj, uintptr(value))
}

func PageSetupDialog_GetMarginBottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetMarginBottom").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetMarginBottom(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetMarginBottom").Call(obj, uintptr(value))
}

func PageSetupDialog_GetOptions(obj uintptr) TPageSetupDialogOptions {
	ret, _, _ := getLazyProc("PageSetupDialog_GetOptions").Call(obj)
	return TPageSetupDialogOptions(ret)
}

func PageSetupDialog_SetOptions(obj uintptr, value TPageSetupDialogOptions) {
	getLazyProc("PageSetupDialog_SetOptions").Call(obj, uintptr(value))
}

func PageSetupDialog_GetPageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetPageWidth").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetPageWidth(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetPageWidth").Call(obj, uintptr(value))
}

func PageSetupDialog_GetPageHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetPageHeight").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetPageHeight(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetPageHeight").Call(obj, uintptr(value))
}

func PageSetupDialog_GetUnits(obj uintptr) TPageMeasureUnits {
	ret, _, _ := getLazyProc("PageSetupDialog_GetUnits").Call(obj)
	return TPageMeasureUnits(ret)
}

func PageSetupDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("PageSetupDialog_GetHandle").Call(obj)
	return HWND(ret)
}

func PageSetupDialog_SetOnClose(obj uintptr, fn interface{}) {
	getLazyProc("PageSetupDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func PageSetupDialog_SetOnShow(obj uintptr, fn interface{}) {
	getLazyProc("PageSetupDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func PageSetupDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func PageSetupDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageSetupDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PageSetupDialog_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("PageSetupDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func PageSetupDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageSetupDialog_GetOwner").Call(obj)
	return ret
}

func PageSetupDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PageSetupDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PageSetupDialog_SetName(obj uintptr, value string) {
	getLazyProc("PageSetupDialog_SetName").Call(obj, GoStrToDStr(value))
}

func PageSetupDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PageSetupDialog_GetTag").Call(obj)
	return int(ret)
}

func PageSetupDialog_SetTag(obj uintptr, value int) {
	getLazyProc("PageSetupDialog_SetTag").Call(obj, uintptr(value))
}

func PageSetupDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PageSetupDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PageSetupDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("PageSetupDialog_StaticClassType").Call()
	return TClass(r)
}
