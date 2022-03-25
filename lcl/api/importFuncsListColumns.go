package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TListColumns ---------------------------

func ListColumns_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListColumns_Create").Call(obj)
	return ret
}

func ListColumns_Free(obj uintptr) {
	_, _, _ = getLazyProc("ListColumns_Free").Call(obj)
}

func ListColumns_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListColumns_Add").Call(obj)
	return ret
}

func ListColumns_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListColumns_Owner").Call(obj)
	return ret
}

func ListColumns_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ListColumns_Assign").Call(obj, Source)
}

func ListColumns_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ListColumns_BeginUpdate").Call(obj)
}

func ListColumns_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ListColumns_Clear").Call(obj)
}

func ListColumns_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ListColumns_Delete").Call(obj, uintptr(Index))
}

func ListColumns_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ListColumns_EndUpdate").Call(obj)
}

func ListColumns_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ListColumns_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ListColumns_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumns_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumns_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListColumns_Insert").Call(obj, uintptr(Index))
	return ret
}

func ListColumns_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListColumns_ClassType").Call(obj)
	return TClass(ret)
}

func ListColumns_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumns_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumns_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumns_InstanceSize").Call(obj)
	return int32(ret)
}

func ListColumns_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListColumns_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListColumns_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListColumns_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListColumns_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumns_GetHashCode").Call(obj)
	return int32(ret)
}

func ListColumns_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListColumns_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListColumns_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListColumns_GetCount").Call(obj)
	return int32(ret)
}

func ListColumns_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListColumns_GetItems").Call(obj, uintptr(Index))
	return ret
}

func ListColumns_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ListColumns_SetItems").Call(obj, uintptr(Index), value)
}

func ListColumns_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListColumns_StaticClassType").Call()
	return TClass(r)
}
