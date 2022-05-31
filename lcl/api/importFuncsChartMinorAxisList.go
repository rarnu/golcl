package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartMinorAxisList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_Create").Call(obj)
	return ret
}

func ChartMinorAxisList_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_Free").Call(obj)
}

func ChartMinorAxisList_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_Owner").Call(obj)
	return ret
}

func ChartMinorAxisList_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_Add").Call(obj)
	return ret
}

func ChartMinorAxisList_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_Assign").Call(obj, Source)
}

func ChartMinorAxisList_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_BeginUpdate").Call(obj)
}

func ChartMinorAxisList_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_Clear").Call(obj)
}

func ChartMinorAxisList_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ChartMinorAxisList_Delete").Call(obj, uintptr(Index))
}

func ChartMinorAxisList_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_EndUpdate").Call(obj)
}

func ChartMinorAxisList_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ChartMinorAxisList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxisList_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_Insert").Call(obj, uintptr(Index))
	return ret
}

func ChartMinorAxisList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartMinorAxisList_ClassType").Call(obj)
	return TClass(ret)
}

func ChartMinorAxisList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxisList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxisList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisList_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartMinorAxisList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartMinorAxisList_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisList_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartMinorAxisList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxisList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxisList_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetCount").Call(obj)
	return int32(ret)
}

func ChartMinorAxisList_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetItems").Call(obj, uintptr(Index))
	return ret
}

func ChartMinorAxisList_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisList_SetItems").Call(obj, uintptr(Index), value)
}

func ChartMinorAxisList_StaticClassType() TClass {
	r, _, _ := getLazyProc("ChartMinorAxisList_StaticClassType").Call()
	return TClass(r)
}

func ChartMinorAxisList_GetChart(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetChart").Call(obj)
	return ret
}

func ChartMinorAxisList_GetAxes(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetAxes").Call(obj, uintptr(AIndex))
	return ret
}

func ChartMinorAxisList_GetParentAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisList_GetParentAxis").Call(obj)
	return ret
}
