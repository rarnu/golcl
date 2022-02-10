package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TGridColumns ---------------------------

func GridColumns_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumns_Add").Call(obj)
	return ret
}

func GridColumns_Clear(obj uintptr) {
	getLazyProc("GridColumns_Clear").Call(obj)
}

func GridColumns_RealIndex(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("GridColumns_RealIndex").Call(obj, uintptr(Index))
	return int32(ret)
}

func GridColumns_IndexOf(obj uintptr, Column uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumns_IndexOf").Call(obj, Column)
	return int32(ret)
}

func GridColumns_IsDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumns_IsDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumns_HasIndex(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("GridColumns_HasIndex").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func GridColumns_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumns_Owner").Call(obj)
	return ret
}

func GridColumns_Assign(obj uintptr, Source uintptr) {
	getLazyProc("GridColumns_Assign").Call(obj, Source)
}

func GridColumns_BeginUpdate(obj uintptr) {
	getLazyProc("GridColumns_BeginUpdate").Call(obj)
}

func GridColumns_Delete(obj uintptr, Index int32) {
	getLazyProc("GridColumns_Delete").Call(obj, uintptr(Index))
}

func GridColumns_EndUpdate(obj uintptr) {
	getLazyProc("GridColumns_EndUpdate").Call(obj)
}

func GridColumns_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("GridColumns_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func GridColumns_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumns_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumns_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("GridColumns_Insert").Call(obj, uintptr(Index))
	return ret
}

func GridColumns_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("GridColumns_ClassType").Call(obj)
	return TClass(ret)
}

func GridColumns_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumns_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumns_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumns_InstanceSize").Call(obj)
	return int32(ret)
}

func GridColumns_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("GridColumns_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func GridColumns_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumns_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func GridColumns_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumns_GetHashCode").Call(obj)
	return int32(ret)
}

func GridColumns_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumns_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumns_GetGrid(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumns_GetGrid").Call(obj)
	return ret
}

func GridColumns_GetVisibleCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumns_GetVisibleCount").Call(obj)
	return int32(ret)
}

func GridColumns_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumns_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumns_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumns_GetCount").Call(obj)
	return int32(ret)
}

func GridColumns_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("GridColumns_GetItems").Call(obj, uintptr(Index))
	return ret
}

func GridColumns_SetItems(obj uintptr, Index int32, value uintptr) {
	getLazyProc("GridColumns_SetItems").Call(obj, uintptr(Index), value)
}

func GridColumns_StaticClassType() TClass {
	r, _, _ := getLazyProc("GridColumns_StaticClassType").Call()
	return TClass(r)
}
