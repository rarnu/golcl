package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TListColumn ---------------------------

func ListColumn_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListColumn_Create").Call(obj)
	return ret
}

func ListColumn_Free(obj uintptr) {
	getLazyProc("ListColumn_Free").Call(obj)
}

func ListColumn_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ListColumn_Assign").Call(obj, Source)
}

func ListColumn_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumn_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumn_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListColumn_ClassType").Call(obj)
	return TClass(ret)
}

func ListColumn_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumn_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumn_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_InstanceSize").Call(obj)
	return int32(ret)
}

func ListColumn_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListColumn_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListColumn_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListColumn_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListColumn_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetHashCode").Call(obj)
	return int32(ret)
}

func ListColumn_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumn_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumn_GetSortIndicator(obj uintptr) TSortIndicator {
	ret, _, _ := getLazyProc("ListColumn_GetSortIndicator").Call(obj)
	return TSortIndicator(ret)
}

func ListColumn_SetSortIndicator(obj uintptr, value TSortIndicator) {
	getLazyProc("ListColumn_SetSortIndicator").Call(obj, uintptr(value))
}

func ListColumn_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ListColumn_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ListColumn_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("ListColumn_SetAlignment").Call(obj, uintptr(value))
}

func ListColumn_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListColumn_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ListColumn_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("ListColumn_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ListColumn_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumn_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumn_SetCaption(obj uintptr, value string) {
	getLazyProc("ListColumn_SetCaption").Call(obj, GoStrToDStr(value))
}

func ListColumn_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetImageIndex").Call(obj)
	return int32(ret)
}

func ListColumn_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetImageIndex").Call(obj, uintptr(value))
}

func ListColumn_GetMaxWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetMaxWidth").Call(obj)
	return int32(ret)
}

func ListColumn_SetMaxWidth(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetMaxWidth").Call(obj, uintptr(value))
}

func ListColumn_GetMinWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetMinWidth").Call(obj)
	return int32(ret)
}

func ListColumn_SetMinWidth(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetMinWidth").Call(obj, uintptr(value))
}

func ListColumn_GetTag(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetTag").Call(obj)
	return int32(ret)
}

func ListColumn_SetTag(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetTag").Call(obj, uintptr(value))
}

func ListColumn_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetWidth").Call(obj)
	return int32(ret)
}

func ListColumn_SetWidth(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetWidth").Call(obj, uintptr(value))
}

func ListColumn_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListColumn_GetCollection").Call(obj)
	return ret
}

func ListColumn_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("ListColumn_SetCollection").Call(obj, value)
}

func ListColumn_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumn_GetIndex").Call(obj)
	return int32(ret)
}

func ListColumn_SetIndex(obj uintptr, value int32) {
	getLazyProc("ListColumn_SetIndex").Call(obj, uintptr(value))
}

func ListColumn_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumn_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumn_SetDisplayName(obj uintptr, value string) {
	getLazyProc("ListColumn_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func ListColumn_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListColumn_StaticClassType").Call()
	return TClass(r)
}
