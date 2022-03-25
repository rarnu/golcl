package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TList ---------------------------

func List_Create() uintptr {
	ret, _, _ := getLazyProc("List_Create").Call()
	return ret
}

func List_Free(obj uintptr) {
	_, _, _ = getLazyProc("List_Free").Call(obj)
}

func List_Add(obj uintptr, Item uintptr) int32 {
	ret, _, _ := getLazyProc("List_Add").Call(obj, Item)
	return int32(ret)
}

func List_Clear(obj uintptr) {
	_, _, _ = getLazyProc("List_Clear").Call(obj)
}

func List_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("List_Delete").Call(obj, uintptr(Index))
}

func List_Expand(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("List_Expand").Call(obj)
	return ret
}

func List_IndexOf(obj uintptr, Item uintptr) int32 {
	ret, _, _ := getLazyProc("List_IndexOf").Call(obj, Item)
	return int32(ret)
}

func List_Insert(obj uintptr, Index int32, Item uintptr) {
	_, _, _ = getLazyProc("List_Insert").Call(obj, uintptr(Index), Item)
}

func List_Move(obj uintptr, CurIndex int32, NewIndex int32) {
	_, _, _ = getLazyProc("List_Move").Call(obj, uintptr(CurIndex), uintptr(NewIndex))
}

func List_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("List_ClassType").Call(obj)
	return TClass(ret)
}

func List_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("List_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func List_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("List_InstanceSize").Call(obj)
	return int32(ret)
}

func List_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("List_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func List_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("List_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func List_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("List_GetHashCode").Call(obj)
	return int32(ret)
}

func List_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("List_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func List_GetCapacity(obj uintptr) int32 {
	ret, _, _ := getLazyProc("List_GetCapacity").Call(obj)
	return int32(ret)
}

func List_SetCapacity(obj uintptr, value int32) {
	_, _, _ = getLazyProc("List_SetCapacity").Call(obj, uintptr(value))
}

func List_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("List_GetCount").Call(obj)
	return int32(ret)
}

func List_SetCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("List_SetCount").Call(obj, uintptr(value))
}

func List_GetList(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("List_GetList").Call(obj)
	return ret
}

func List_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("List_GetItems").Call(obj, uintptr(Index))
	return ret
}

func List_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("List_SetItems").Call(obj, uintptr(Index), value)
}

func List_StaticClassType() TClass {
	r, _, _ := getLazyProc("List_StaticClassType").Call()
	return TClass(r)
}
