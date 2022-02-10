package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTaskDialogRadioButtonItem ---------------------------

func TaskDialogRadioButtonItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_Create").Call(obj)
	return ret
}

func TaskDialogRadioButtonItem_Free(obj uintptr) {
	getLazyProc("TaskDialogRadioButtonItem_Free").Call(obj)
}

func TaskDialogRadioButtonItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogRadioButtonItem_Assign(obj uintptr, Source uintptr) {
	getLazyProc("TaskDialogRadioButtonItem_Assign").Call(obj, Source)
}

func TaskDialogRadioButtonItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_ClassType").Call(obj)
	return TClass(ret)
}

func TaskDialogRadioButtonItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogRadioButtonItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_InstanceSize").Call(obj)
	return int32(ret)
}

func TaskDialogRadioButtonItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TaskDialogRadioButtonItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TaskDialogRadioButtonItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetHashCode").Call(obj)
	return int32(ret)
}

func TaskDialogRadioButtonItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogRadioButtonItem_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func TaskDialogRadioButtonItem_SetModalResult(obj uintptr, value TModalResult) {
	getLazyProc("TaskDialogRadioButtonItem_SetModalResult").Call(obj, uintptr(value))
}

func TaskDialogRadioButtonItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogRadioButtonItem_SetCaption(obj uintptr, value string) {
	getLazyProc("TaskDialogRadioButtonItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func TaskDialogRadioButtonItem_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func TaskDialogRadioButtonItem_SetDefault(obj uintptr, value bool) {
	getLazyProc("TaskDialogRadioButtonItem_SetDefault").Call(obj, GoBoolToDBool(value))
}

func TaskDialogRadioButtonItem_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetCollection").Call(obj)
	return ret
}

func TaskDialogRadioButtonItem_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("TaskDialogRadioButtonItem_SetCollection").Call(obj, value)
}

func TaskDialogRadioButtonItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetIndex").Call(obj)
	return int32(ret)
}

func TaskDialogRadioButtonItem_SetIndex(obj uintptr, value int32) {
	getLazyProc("TaskDialogRadioButtonItem_SetIndex").Call(obj, uintptr(value))
}

func TaskDialogRadioButtonItem_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogRadioButtonItem_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogRadioButtonItem_SetDisplayName(obj uintptr, value string) {
	getLazyProc("TaskDialogRadioButtonItem_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func TaskDialogRadioButtonItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("TaskDialogRadioButtonItem_StaticClassType").Call()
	return TClass(r)
}
