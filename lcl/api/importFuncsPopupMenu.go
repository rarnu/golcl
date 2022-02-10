package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TPopupMenu ---------------------------

func PopupMenu_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_Create").Call(obj)
	return ret
}

func PopupMenu_Free(obj uintptr) {
	getLazyProc("PopupMenu_Free").Call(obj)
}

func PopupMenu_CloseMenu(obj uintptr) {
	getLazyProc("PopupMenu_CloseMenu").Call(obj)
}

func PopupMenu_Popup(obj uintptr, X int32, Y int32) {
	getLazyProc("PopupMenu_Popup").Call(obj, uintptr(X), uintptr(Y))
}

func PopupMenu_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PopupMenu_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PopupMenu_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PopupMenu_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PopupMenu_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PopupMenu_Assign(obj uintptr, Source uintptr) {
	getLazyProc("PopupMenu_Assign").Call(obj, Source)
}

func PopupMenu_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PopupMenu_ClassType").Call(obj)
	return TClass(ret)
}

func PopupMenu_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PopupMenu_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PopupMenu_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PopupMenu_InstanceSize").Call(obj)
	return int32(ret)
}

func PopupMenu_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PopupMenu_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PopupMenu_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PopupMenu_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PopupMenu_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PopupMenu_GetHashCode").Call(obj)
	return int32(ret)
}

func PopupMenu_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PopupMenu_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PopupMenu_GetImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PopupMenu_GetImagesWidth").Call(obj)
	return int32(ret)
}

func PopupMenu_SetImagesWidth(obj uintptr, value int32) {
	getLazyProc("PopupMenu_SetImagesWidth").Call(obj, uintptr(value))
}

func PopupMenu_GetPopupComponent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_GetPopupComponent").Call(obj)
	return ret
}

func PopupMenu_SetPopupComponent(obj uintptr, value uintptr) {
	getLazyProc("PopupMenu_SetPopupComponent").Call(obj, value)
}

func PopupMenu_GetPopupPoint(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("PopupMenu_GetPopupPoint").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PopupMenu_GetAlignment(obj uintptr) TPopupAlignment {
	ret, _, _ := getLazyProc("PopupMenu_GetAlignment").Call(obj)
	return TPopupAlignment(ret)
}

func PopupMenu_SetAlignment(obj uintptr, value TPopupAlignment) {
	getLazyProc("PopupMenu_SetAlignment").Call(obj, uintptr(value))
}

func PopupMenu_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("PopupMenu_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func PopupMenu_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("PopupMenu_SetBiDiMode").Call(obj, uintptr(value))
}

func PopupMenu_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_GetImages").Call(obj)
	return ret
}

func PopupMenu_SetImages(obj uintptr, value uintptr) {
	getLazyProc("PopupMenu_SetImages").Call(obj, value)
}

func PopupMenu_GetOwnerDraw(obj uintptr) bool {
	ret, _, _ := getLazyProc("PopupMenu_GetOwnerDraw").Call(obj)
	return DBoolToGoBool(ret)
}

func PopupMenu_SetOwnerDraw(obj uintptr, value bool) {
	getLazyProc("PopupMenu_SetOwnerDraw").Call(obj, GoBoolToDBool(value))
}

func PopupMenu_SetOnPopup(obj uintptr, fn interface{}) {
	getLazyProc("PopupMenu_SetOnPopup").Call(obj, addEventToMap(obj, fn))
}

func PopupMenu_GetHandle(obj uintptr) HMENU {
	ret, _, _ := getLazyProc("PopupMenu_GetHandle").Call(obj)
	return HMENU(ret)
}

func PopupMenu_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_GetItems").Call(obj)
	return ret
}

func PopupMenu_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PopupMenu_GetComponentCount").Call(obj)
	return int32(ret)
}

func PopupMenu_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PopupMenu_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PopupMenu_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("PopupMenu_SetComponentIndex").Call(obj, uintptr(value))
}

func PopupMenu_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_GetOwner").Call(obj)
	return ret
}

func PopupMenu_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PopupMenu_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PopupMenu_SetName(obj uintptr, value string) {
	getLazyProc("PopupMenu_SetName").Call(obj, GoStrToDStr(value))
}

func PopupMenu_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PopupMenu_GetTag").Call(obj)
	return int(ret)
}

func PopupMenu_SetTag(obj uintptr, value int) {
	getLazyProc("PopupMenu_SetTag").Call(obj, uintptr(value))
}

func PopupMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PopupMenu_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PopupMenu_StaticClassType() TClass {
	r, _, _ := getLazyProc("PopupMenu_StaticClassType").Call()
	return TClass(r)
}
