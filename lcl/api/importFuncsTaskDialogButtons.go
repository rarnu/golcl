package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTaskDialogButtons ---------------------------

func TaskDialogButtons_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_Add").Call(obj)
	return ret
}

func TaskDialogButtons_FindButton(obj uintptr, AModalResult TModalResult) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_FindButton").Call(obj, uintptr(AModalResult))
	return ret
}

func TaskDialogButtons_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_Owner").Call(obj)
	return ret
}

func TaskDialogButtons_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_Assign").Call(obj, Source)
}

func TaskDialogButtons_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_BeginUpdate").Call(obj)
}

func TaskDialogButtons_Clear(obj uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_Clear").Call(obj)
}

func TaskDialogButtons_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("TaskDialogButtons_Delete").Call(obj, uintptr(Index))
}

func TaskDialogButtons_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_EndUpdate").Call(obj)
}

func TaskDialogButtons_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func TaskDialogButtons_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtons_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtons_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_Insert").Call(obj, uintptr(Index))
	return ret
}

func TaskDialogButtons_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TaskDialogButtons_ClassType").Call(obj)
	return TClass(ret)
}

func TaskDialogButtons_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtons_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtons_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtons_InstanceSize").Call(obj)
	return int32(ret)
}

func TaskDialogButtons_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TaskDialogButtons_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TaskDialogButtons_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialogButtons_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TaskDialogButtons_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtons_GetHashCode").Call(obj)
	return int32(ret)
}

func TaskDialogButtons_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialogButtons_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialogButtons_GetDefaultButton(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_GetDefaultButton").Call(obj)
	return ret
}

func TaskDialogButtons_SetDefaultButton(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_SetDefaultButton").Call(obj, value)
}

func TaskDialogButtons_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialogButtons_GetCount").Call(obj)
	return int32(ret)
}

func TaskDialogButtons_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TaskDialogButtons_GetItems").Call(obj, uintptr(Index))
	return ret
}

func TaskDialogButtons_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("TaskDialogButtons_SetItems").Call(obj, uintptr(Index), value)
}

func TaskDialogButtons_StaticClassType() TClass {
	r, _, _ := getLazyProc("TaskDialogButtons_StaticClassType").Call()
	return TClass(r)
}
