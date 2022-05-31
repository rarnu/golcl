package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartAxisList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_Create").Call(obj)
	return ret
}

func ChartAxisList_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_Free").Call(obj)
}

func ChartAxisList_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartAxisList_StaticClassType").Call()
	return TClass(ret)
}

func ChartAxisList_GetBottomAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_GetBottomAxis").Call(obj)
	return ret
}

func ChartAxisList_SetBottomAxis(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_SetBottomAxis").Call(obj, value)
}

func ChartAxisList_GetLeftAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_GetLeftAxis").Call(obj)
	return ret
}

func ChartAxisList_SetLeftAxis(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_SetLeftAxis").Call(obj, value)
}

func ChartAxisList_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_Add").Call(obj)
	return ret
}

func ChartAxisList_Draw(obj uintptr, ACurrentZ int32, AIndex int32) {
	_, _, _ = getLazyProc("ChartAxisList_Draw").Call(obj, uintptr(ACurrentZ), uintptr(AIndex))
}

func ChartAxisList_GetAxisByAlign(obj uintptr, AAlign TChartAxisAlignment) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_GetAxisByAlign").Call(obj, uintptr(AAlign))
	return ret
}

func ChartAxisList_GetEnumerator(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_GetEnumerator").Call(obj)
	return ret
}

func ChartAxisList_Measure(obj uintptr, AExtent TDoubleRect, AClipRect TRect, ADepth int32) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_Measure").Call(obj, uintptr(unsafe.Pointer(&AExtent)), uintptr(unsafe.Pointer(&AClipRect)), uintptr(ADepth))
	return ret
}

func ChartAxisList_Prepare(obj uintptr, ARect TRect) {
	_, _, _ = getLazyProc("ChartAxisList_Prepare").Call(obj, uintptr(unsafe.Pointer(&ARect)))
}

func ChartAxisList_PrepareGroups(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_PrepareGroups").Call(obj)
}

func ChartAxisList_SetAxisByAlign(obj uintptr, AAlign TChartAxisAlignment, AValue uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_SetAxisByAlign").Call(obj, uintptr(AAlign), AValue)
}

func ChartAxisList_UpdateBiDiMode(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_UpdateBiDiMode").Call(obj)
}

func ChartAxisList_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_BeginUpdate").Call(obj)
}

func ChartAxisList_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_Clear").Call(obj)
}

func ChartAxisList_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ChartAxisList_Delete").Call(obj, uintptr(Index))
}

func ChartAxisList_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_EndUpdate").Call(obj)
}

func ChartAxisList_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ChartAxisList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisList_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_Insert").Call(obj, uintptr(Index))
	return ret
}

func ChartAxisList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartAxisList_ClassType").Call(obj)
	return TClass(ret)
}

func ChartAxisList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisList_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartAxisList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartAxisList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartAxisList_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisList_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartAxisList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisList_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartAxisList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisList_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisList_GetCount").Call(obj)
	return int32(ret)
}

func ChartAxisList_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartAxisList_GetItems").Call(obj, uintptr(Index))
	return ret
}

func ChartAxisList_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisList_SetItems").Call(obj, uintptr(Index), value)
}
