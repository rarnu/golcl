package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TStatusPanels ---------------------------

func StatusPanels_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_Create").Call(obj)
	return ret
}

func StatusPanels_Free(obj uintptr) {
	getLazyProc("StatusPanels_Free").Call(obj)
}

func StatusPanels_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_Add").Call(obj)
	return ret
}

func StatusPanels_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_Insert").Call(obj, uintptr(Index))
	return ret
}

func StatusPanels_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_Owner").Call(obj)
	return ret
}

func StatusPanels_Assign(obj uintptr, Source uintptr) {
	getLazyProc("StatusPanels_Assign").Call(obj, Source)
}

func StatusPanels_BeginUpdate(obj uintptr) {
	getLazyProc("StatusPanels_BeginUpdate").Call(obj)
}

func StatusPanels_Clear(obj uintptr) {
	getLazyProc("StatusPanels_Clear").Call(obj)
}

func StatusPanels_Delete(obj uintptr, Index int32) {
	getLazyProc("StatusPanels_Delete").Call(obj, uintptr(Index))
}

func StatusPanels_EndUpdate(obj uintptr) {
	getLazyProc("StatusPanels_EndUpdate").Call(obj)
}

func StatusPanels_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func StatusPanels_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanels_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanels_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StatusPanels_ClassType").Call(obj)
	return TClass(ret)
}

func StatusPanels_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanels_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanels_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanels_InstanceSize").Call(obj)
	return int32(ret)
}

func StatusPanels_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StatusPanels_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StatusPanels_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusPanels_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func StatusPanels_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanels_GetHashCode").Call(obj)
	return int32(ret)
}

func StatusPanels_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanels_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanels_GetCapacity(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanels_GetCapacity").Call(obj)
	return int32(ret)
}

func StatusPanels_SetCapacity(obj uintptr, value int32) {
	getLazyProc("StatusPanels_SetCapacity").Call(obj, uintptr(value))
}

func StatusPanels_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanels_GetCount").Call(obj)
	return int32(ret)
}

func StatusPanels_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StatusPanels_GetItems").Call(obj, uintptr(Index))
	return ret
}

func StatusPanels_SetItems(obj uintptr, Index int32, value uintptr) {
	getLazyProc("StatusPanels_SetItems").Call(obj, uintptr(Index), value)
}

func StatusPanels_StaticClassType() TClass {
	r, _, _ := getLazyProc("StatusPanels_StaticClassType").Call()
	return TClass(r)
}

//--------------------------- TStatusPanel ---------------------------

func StatusPanel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusPanel_Create").Call(obj)
	return ret
}

func StatusPanel_Free(obj uintptr) {
	getLazyProc("StatusPanel_Free").Call(obj)
}

func StatusPanel_Assign(obj uintptr, Source uintptr) {
	getLazyProc("StatusPanel_Assign").Call(obj, Source)
}

func StatusPanel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StatusPanel_ClassType").Call(obj)
	return TClass(ret)
}

func StatusPanel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanel_InstanceSize").Call(obj)
	return int32(ret)
}

func StatusPanel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StatusPanel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StatusPanel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusPanel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func StatusPanel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanel_GetHashCode").Call(obj)
	return int32(ret)
}

func StatusPanel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanel_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("StatusPanel_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func StatusPanel_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("StatusPanel_SetAlignment").Call(obj, uintptr(value))
}

func StatusPanel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("StatusPanel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func StatusPanel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("StatusPanel_SetBiDiMode").Call(obj, uintptr(value))
}

func StatusPanel_GetStyle(obj uintptr) TStatusPanelStyle {
	ret, _, _ := getLazyProc("StatusPanel_GetStyle").Call(obj)
	return TStatusPanelStyle(ret)
}

func StatusPanel_SetStyle(obj uintptr, value TStatusPanelStyle) {
	getLazyProc("StatusPanel_SetStyle").Call(obj, uintptr(value))
}

func StatusPanel_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanel_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanel_SetText(obj uintptr, value string) {
	getLazyProc("StatusPanel_SetText").Call(obj, GoStrToDStr(value))
}

func StatusPanel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanel_GetWidth").Call(obj)
	return int32(ret)
}

func StatusPanel_SetWidth(obj uintptr, value int32) {
	getLazyProc("StatusPanel_SetWidth").Call(obj, uintptr(value))
}

func StatusPanel_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusPanel_GetCollection").Call(obj)
	return ret
}

func StatusPanel_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("StatusPanel_SetCollection").Call(obj, value)
}

func StatusPanel_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusPanel_GetIndex").Call(obj)
	return int32(ret)
}

func StatusPanel_SetIndex(obj uintptr, value int32) {
	getLazyProc("StatusPanel_SetIndex").Call(obj, uintptr(value))
}

func StatusPanel_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusPanel_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func StatusPanel_SetDisplayName(obj uintptr, value string) {
	getLazyProc("StatusPanel_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func StatusPanel_StaticClassType() TClass {
	r, _, _ := getLazyProc("StatusPanel_StaticClassType").Call()
	return TClass(r)
}
