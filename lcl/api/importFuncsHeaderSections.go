package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- THeaderSections ---------------------------

func HeaderSections_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_Create").Call(obj)
	return ret
}

func HeaderSections_Free(obj uintptr) {
	getLazyProc("HeaderSections_Free").Call(obj)
}

func HeaderSections_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_Add").Call(obj)
	return ret
}

func HeaderSections_AddItem(obj uintptr, Item uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_AddItem").Call(obj, Item, uintptr(Index))
	return ret
}

func HeaderSections_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_Insert").Call(obj, uintptr(Index))
	return ret
}

func HeaderSections_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_Owner").Call(obj)
	return ret
}

func HeaderSections_Assign(obj uintptr, Source uintptr) {
	getLazyProc("HeaderSections_Assign").Call(obj, Source)
}

func HeaderSections_BeginUpdate(obj uintptr) {
	getLazyProc("HeaderSections_BeginUpdate").Call(obj)
}

func HeaderSections_Clear(obj uintptr) {
	getLazyProc("HeaderSections_Clear").Call(obj)
}

func HeaderSections_Delete(obj uintptr, Index int32) {
	getLazyProc("HeaderSections_Delete").Call(obj, uintptr(Index))
}

func HeaderSections_EndUpdate(obj uintptr) {
	getLazyProc("HeaderSections_EndUpdate").Call(obj)
}

func HeaderSections_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func HeaderSections_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSections_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSections_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("HeaderSections_ClassType").Call(obj)
	return TClass(ret)
}

func HeaderSections_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSections_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSections_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSections_InstanceSize").Call(obj)
	return int32(ret)
}

func HeaderSections_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("HeaderSections_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func HeaderSections_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderSections_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func HeaderSections_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSections_GetHashCode").Call(obj)
	return int32(ret)
}

func HeaderSections_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderSections_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderSections_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderSections_GetCount").Call(obj)
	return int32(ret)
}

func HeaderSections_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("HeaderSections_GetItems").Call(obj, uintptr(Index))
	return ret
}

func HeaderSections_SetItems(obj uintptr, Index int32, value uintptr) {
	getLazyProc("HeaderSections_SetItems").Call(obj, uintptr(Index), value)
}

func HeaderSections_StaticClassType() TClass {
	r, _, _ := getLazyProc("HeaderSections_StaticClassType").Call()
	return TClass(r)
}
