package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTaskDialogButtonItem ---------------------------

func TaskDialogButtonItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_Create").Call(obj)
	return ret
}

func TaskDialogButtonItem_Free(obj uintptr) {
	getLazyProc("TaskDialogButtonItem_Free").Call(obj)
}

func TaskDialogButtonItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtonItem_Assign(obj uintptr, Source uintptr) {
	getLazyProc("TaskDialogButtonItem_Assign").Call(obj, Source)
}

func TaskDialogButtonItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_ClassType").Call(obj)
	return TClass(ret)
}

func TaskDialogButtonItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtonItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_InstanceSize").Call(obj)
	return int32(ret)
}

func TaskDialogButtonItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TaskDialogButtonItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TaskDialogButtonItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetHashCode").Call(obj)
	return int32(ret)
}

func TaskDialogButtonItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtonItem_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func TaskDialogButtonItem_SetModalResult(obj uintptr, value TModalResult) {
	getLazyProc("TaskDialogButtonItem_SetModalResult").Call(obj, uintptr(value))
}

func TaskDialogButtonItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtonItem_SetCaption(obj uintptr, value string) {
	getLazyProc("TaskDialogButtonItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func TaskDialogButtonItem_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func TaskDialogButtonItem_SetDefault(obj uintptr, value bool) {
	getLazyProc("TaskDialogButtonItem_SetDefault").Call(obj, GoBoolToDBool(value))
}

func TaskDialogButtonItem_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetCollection").Call(obj)
	return ret
}

func TaskDialogButtonItem_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("TaskDialogButtonItem_SetCollection").Call(obj, value)
}

func TaskDialogButtonItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetIndex").Call(obj)
	return int32(ret)
}

func TaskDialogButtonItem_SetIndex(obj uintptr, value int32) {
	getLazyProc("TaskDialogButtonItem_SetIndex").Call(obj, uintptr(value))
}

func TaskDialogButtonItem_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtonItem_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtonItem_SetDisplayName(obj uintptr, value string) {
	getLazyProc("TaskDialogButtonItem_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func TaskDialogButtonItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("TaskDialogButtonItem_StaticClassType").Call()
	return TClass(r)
}
