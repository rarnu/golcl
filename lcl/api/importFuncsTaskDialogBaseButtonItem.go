package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTaskDialogBaseButtonItem ---------------------------

func TaskDialogBaseButtonItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_Create").Call(obj)
	return ret
}

func TaskDialogBaseButtonItem_Free(obj uintptr) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_Free").Call(obj)
}

func TaskDialogBaseButtonItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogBaseButtonItem_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_Assign").Call(obj, Source)
}

func TaskDialogBaseButtonItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_ClassType").Call(obj)
	return TClass(ret)
}

func TaskDialogBaseButtonItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogBaseButtonItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_InstanceSize").Call(obj)
	return int32(ret)
}

func TaskDialogBaseButtonItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TaskDialogBaseButtonItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TaskDialogBaseButtonItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetHashCode").Call(obj)
	return int32(ret)
}

func TaskDialogBaseButtonItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogBaseButtonItem_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func TaskDialogBaseButtonItem_SetModalResult(obj uintptr, value TModalResult) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetModalResult").Call(obj, uintptr(value))
}

func TaskDialogBaseButtonItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogBaseButtonItem_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func TaskDialogBaseButtonItem_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func TaskDialogBaseButtonItem_SetDefault(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetDefault").Call(obj, GoBoolToDBool(value))
}

func TaskDialogBaseButtonItem_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetCollection").Call(obj)
	return ret
}

func TaskDialogBaseButtonItem_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetCollection").Call(obj, value)
}

func TaskDialogBaseButtonItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetIndex").Call(obj)
	return int32(ret)
}

func TaskDialogBaseButtonItem_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetIndex").Call(obj, uintptr(value))
}

func TaskDialogBaseButtonItem_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogBaseButtonItem_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogBaseButtonItem_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialogBaseButtonItem_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func TaskDialogBaseButtonItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("TaskDialogBaseButtonItem_StaticClassType").Call()
	return TClass(r)
}
