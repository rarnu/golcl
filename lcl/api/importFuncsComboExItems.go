package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TComboExItems ---------------------------

func ComboExItems_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_Add").Call(obj)
	return ret
}

func ComboExItems_AddItem(obj uintptr, Caption string, ImageIndex int32, SelectedImageIndex int32, OverlayImageIndex int32, Indent int32, Data uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_AddItem").Call(obj, GoStrToDStr(Caption), uintptr(ImageIndex), uintptr(SelectedImageIndex), uintptr(OverlayImageIndex), uintptr(Indent), Data)
	return ret
}

func ComboExItems_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_Insert").Call(obj, uintptr(Index))
	return ret
}

func ComboExItems_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_Owner").Call(obj)
	return ret
}

func ComboExItems_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ComboExItems_Assign").Call(obj, Source)
}

func ComboExItems_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ComboExItems_BeginUpdate").Call(obj)
}

func ComboExItems_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ComboExItems_Clear").Call(obj)
}

func ComboExItems_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ComboExItems_Delete").Call(obj, uintptr(Index))
}

func ComboExItems_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ComboExItems_EndUpdate").Call(obj)
}

func ComboExItems_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ComboExItems_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItems_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItems_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ComboExItems_ClassType").Call(obj)
	return TClass(ret)
}

func ComboExItems_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItems_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItems_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItems_InstanceSize").Call(obj)
	return int32(ret)
}

func ComboExItems_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ComboExItems_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ComboExItems_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboExItems_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ComboExItems_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItems_GetHashCode").Call(obj)
	return int32(ret)
}

func ComboExItems_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItems_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItems_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItems_GetCount").Call(obj)
	return int32(ret)
}

func ComboExItems_GetComboItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboExItems_GetComboItems").Call(obj, uintptr(Index))
	return ret
}

func ComboExItems_StaticClassType() TClass {
	r, _, _ := getLazyProc("ComboExItems_StaticClassType").Call()
	return TClass(r)
}
