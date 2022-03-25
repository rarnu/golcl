package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TListItems ---------------------------

func ListItems_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItems_Create").Call(obj)
	return ret
}

func ListItems_Free(obj uintptr) {
	_, _, _ = getLazyProc("ListItems_Free").Call(obj)
}

func ListItems_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItems_Add").Call(obj)
	return ret
}

func ListItems_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ListItems_Assign").Call(obj, Source)
}

func ListItems_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ListItems_BeginUpdate").Call(obj)
}

func ListItems_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ListItems_Clear").Call(obj)
}

func ListItems_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ListItems_Delete").Call(obj, uintptr(Index))
}

func ListItems_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ListItems_EndUpdate").Call(obj)
}

func ListItems_IndexOf(obj uintptr, Value uintptr) int32 {
	ret, _, _ := getLazyProc("ListItems_IndexOf").Call(obj, Value)
	return int32(ret)
}

func ListItems_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListItems_Insert").Call(obj, uintptr(Index))
	return ret
}

func ListItems_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItems_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListItems_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListItems_ClassType").Call(obj)
	return TClass(ret)
}

func ListItems_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItems_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListItems_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItems_InstanceSize").Call(obj)
	return int32(ret)
}

func ListItems_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListItems_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListItems_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItems_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListItems_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItems_GetHashCode").Call(obj)
	return int32(ret)
}

func ListItems_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItems_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListItems_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItems_GetCount").Call(obj)
	return int32(ret)
}

func ListItems_SetCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListItems_SetCount").Call(obj, uintptr(value))
}

func ListItems_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItems_GetOwner").Call(obj)
	return ret
}

func ListItems_GetItem(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListItems_GetItem").Call(obj, uintptr(Index))
	return ret
}

func ListItems_SetItem(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ListItems_SetItem").Call(obj, uintptr(Index), value)
}

func ListItems_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListItems_StaticClassType").Call()
	return TClass(r)
}
