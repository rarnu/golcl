package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartStyleList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_Create").Call(obj)
	return ret
}

func ChartStyleList_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_Free").Call(obj)
}

func ChartStyleList_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_Owner").Call(obj)
	return ret
}

func ChartStyleList_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_Add").Call(obj)
	return ret
}

func ChartStyleList_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_Assign").Call(obj, Source)
}

func ChartStyleList_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_BeginUpdate").Call(obj)
}

func ChartStyleList_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_Clear").Call(obj)
}

func ChartStyleList_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ChartStyleList_Delete").Call(obj, uintptr(Index))
}

func ChartStyleList_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_EndUpdate").Call(obj)
}

func ChartStyleList_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ChartStyleList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyleList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartStyleList_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_Insert").Call(obj, uintptr(Index))
	return ret
}

func ChartStyleList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartStyleList_ClassType").Call(obj)
	return TClass(ret)
}

func ChartStyleList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyleList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartStyleList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartStyleList_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartStyleList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartStyleList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartStyleList_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartStyleList_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartStyleList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartStyleList_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartStyleList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyleList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartStyleList_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartStyleList_GetCount").Call(obj)
	return int32(ret)
}

func ChartStyleList_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_GetItems").Call(obj, uintptr(Index))
	return ret
}

func ChartStyleList_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ChartStyleList_SetItems").Call(obj, uintptr(Index), value)
}

func ChartStyleList_StaticClassType() TClass {
	r, _, _ := getLazyProc("ChartStyleList_StaticClassType").Call()
	return TClass(r)
}

func ChartStyleList_GetEnumerator(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_GetEnumerator").Call(obj)
	return ret
}

func ChartStyleList_GetStyle(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ChartStyleList_GetStyle").Call(obj, uintptr(AIndex))
	return ret
}
