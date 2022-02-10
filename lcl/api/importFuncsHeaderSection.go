package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- THeaderSection ---------------------------

func HeaderSection_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderSection_Create").Call(obj)
	return ret
}

func HeaderSection_Free(obj uintptr) {
	getLazyProc("HeaderSection_Free").Call(obj)
}

func HeaderSection_Assign(obj uintptr, Source uintptr) {
	getLazyProc("HeaderSection_Assign").Call(obj, Source)
}

func HeaderSection_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSection_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSection_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("HeaderSection_ClassType").Call(obj)
	return TClass(ret)
}

func HeaderSection_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSection_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSection_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_InstanceSize").Call(obj)
	return int32(ret)
}

func HeaderSection_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("HeaderSection_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func HeaderSection_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderSection_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func HeaderSection_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetHashCode").Call(obj)
	return int32(ret)
}

func HeaderSection_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSection_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSection_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetLeft").Call(obj)
	return int32(ret)
}

func HeaderSection_GetRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetRight").Call(obj)
	return int32(ret)
}

func HeaderSection_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("HeaderSection_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func HeaderSection_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("HeaderSection_SetAlignment").Call(obj, uintptr(value))
}

func HeaderSection_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetImageIndex").Call(obj)
	return int32(ret)
}

func HeaderSection_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("HeaderSection_SetImageIndex").Call(obj, uintptr(value))
}

func HeaderSection_GetMaxWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetMaxWidth").Call(obj)
	return int32(ret)
}

func HeaderSection_SetMaxWidth(obj uintptr, value int32) {
	getLazyProc("HeaderSection_SetMaxWidth").Call(obj, uintptr(value))
}

func HeaderSection_GetMinWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetMinWidth").Call(obj)
	return int32(ret)
}

func HeaderSection_SetMinWidth(obj uintptr, value int32) {
	getLazyProc("HeaderSection_SetMinWidth").Call(obj, uintptr(value))
}

func HeaderSection_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSection_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSection_SetText(obj uintptr, value string) {
	getLazyProc("HeaderSection_SetText").Call(obj, GoStrToDStr(value))
}

func HeaderSection_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetWidth").Call(obj)
	return int32(ret)
}

func HeaderSection_SetWidth(obj uintptr, value int32) {
	getLazyProc("HeaderSection_SetWidth").Call(obj, uintptr(value))
}

func HeaderSection_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderSection_GetCollection").Call(obj)
	return ret
}

func HeaderSection_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("HeaderSection_SetCollection").Call(obj, value)
}

func HeaderSection_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSection_GetIndex").Call(obj)
	return int32(ret)
}

func HeaderSection_SetIndex(obj uintptr, value int32) {
	getLazyProc("HeaderSection_SetIndex").Call(obj, uintptr(value))
}

func HeaderSection_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSection_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSection_SetDisplayName(obj uintptr, value string) {
	getLazyProc("HeaderSection_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func HeaderSection_StaticClassType() TClass {
	r, _, _ := getLazyProc("HeaderSection_StaticClassType").Call()
	return TClass(r)
}
