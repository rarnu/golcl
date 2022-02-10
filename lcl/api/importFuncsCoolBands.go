package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TCoolBands ---------------------------

func CoolBands_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBands_Create").Call(obj)
	return ret
}

func CoolBands_Free(obj uintptr) {
	getLazyProc("CoolBands_Free").Call(obj)
}

func CoolBands_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBands_Add").Call(obj)
	return ret
}

func CoolBands_FindBand(obj uintptr, AControl uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBands_FindBand").Call(obj, AControl)
	return ret
}

func CoolBands_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBands_Owner").Call(obj)
	return ret
}

func CoolBands_Assign(obj uintptr, Source uintptr) {
	getLazyProc("CoolBands_Assign").Call(obj, Source)
}

func CoolBands_BeginUpdate(obj uintptr) {
	getLazyProc("CoolBands_BeginUpdate").Call(obj)
}

func CoolBands_Clear(obj uintptr) {
	getLazyProc("CoolBands_Clear").Call(obj)
}

func CoolBands_Delete(obj uintptr, Index int32) {
	getLazyProc("CoolBands_Delete").Call(obj, uintptr(Index))
}

func CoolBands_EndUpdate(obj uintptr) {
	getLazyProc("CoolBands_EndUpdate").Call(obj)
}

func CoolBands_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("CoolBands_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func CoolBands_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBands_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBands_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CoolBands_Insert").Call(obj, uintptr(Index))
	return ret
}

func CoolBands_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CoolBands_ClassType").Call(obj)
	return TClass(ret)
}

func CoolBands_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBands_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBands_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBands_InstanceSize").Call(obj)
	return int32(ret)
}

func CoolBands_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CoolBands_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CoolBands_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBands_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CoolBands_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBands_GetHashCode").Call(obj)
	return int32(ret)
}

func CoolBands_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBands_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBands_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBands_GetCount").Call(obj)
	return int32(ret)
}

func CoolBands_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CoolBands_GetItems").Call(obj, uintptr(Index))
	return ret
}

func CoolBands_SetItems(obj uintptr, Index int32, value uintptr) {
	getLazyProc("CoolBands_SetItems").Call(obj, uintptr(Index), value)
}

func CoolBands_StaticClassType() TClass {
	r, _, _ := getLazyProc("CoolBands_StaticClassType").Call()
	return TClass(r)
}
