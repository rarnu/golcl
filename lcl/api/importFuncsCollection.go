package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TCollection ---------------------------

func Collection_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Collection_Create").Call(obj)
	return ret
}

func Collection_Free(obj uintptr) {
	getLazyProc("Collection_Free").Call(obj)
}

func Collection_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Collection_Owner").Call(obj)
	return ret
}

func Collection_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Collection_Add").Call(obj)
	return ret
}

func Collection_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Collection_Assign").Call(obj, Source)
}

func Collection_BeginUpdate(obj uintptr) {
	getLazyProc("Collection_BeginUpdate").Call(obj)
}

func Collection_Clear(obj uintptr) {
	getLazyProc("Collection_Clear").Call(obj)
}

func Collection_Delete(obj uintptr, Index int32) {
	getLazyProc("Collection_Delete").Call(obj, uintptr(Index))
}

func Collection_EndUpdate(obj uintptr) {
	getLazyProc("Collection_EndUpdate").Call(obj)
}

func Collection_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("Collection_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func Collection_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Collection_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Collection_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Collection_Insert").Call(obj, uintptr(Index))
	return ret
}

func Collection_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Collection_ClassType").Call(obj)
	return TClass(ret)
}

func Collection_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Collection_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Collection_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Collection_InstanceSize").Call(obj)
	return int32(ret)
}

func Collection_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Collection_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Collection_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Collection_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Collection_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Collection_GetHashCode").Call(obj)
	return int32(ret)
}

func Collection_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Collection_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Collection_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Collection_GetCount").Call(obj)
	return int32(ret)
}

func Collection_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Collection_GetItems").Call(obj, uintptr(Index))
	return ret
}

func Collection_SetItems(obj uintptr, Index int32, value uintptr) {
	getLazyProc("Collection_SetItems").Call(obj, uintptr(Index), value)
}

func Collection_StaticClassType() TClass {
	r, _, _ := getLazyProc("Collection_StaticClassType").Call()
	return TClass(r)
}
