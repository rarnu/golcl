package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TListItem ---------------------------

func ListItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItem_Create").Call(obj)
	return ret
}

func ListItem_Free(obj uintptr) {
	getLazyProc("ListItem_Free").Call(obj)
}

func ListItem_DisplayRectSubItem(obj uintptr, subItem int32, Code TDisplayCode) TRect {
	var ret TRect
	getLazyProc("ListItem_DisplayRectSubItem").Call(obj, uintptr(subItem), uintptr(Code), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListItem_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ListItem_Assign").Call(obj, Source)
}

func ListItem_Delete(obj uintptr) {
	getLazyProc("ListItem_Delete").Call(obj)
}

func ListItem_DisplayRect(obj uintptr, Code TDisplayCode) TRect {
	var ret TRect
	getLazyProc("ListItem_DisplayRect").Call(obj, uintptr(Code), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListItem_EditCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_EditCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_MakeVisible(obj uintptr, PartialOK bool) {
	getLazyProc("ListItem_MakeVisible").Call(obj, GoBoolToDBool(PartialOK))
}

func ListItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListItem_ClassType").Call(obj)
	return TClass(ret)
}

func ListItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_InstanceSize").Call(obj)
	return int32(ret)
}

func ListItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetHashCode").Call(obj)
	return int32(ret)
}

func ListItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListItem_GetDropTarget(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_GetDropTarget").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_SetDropTarget(obj uintptr, value bool) {
	getLazyProc("ListItem_SetDropTarget").Call(obj, GoBoolToDBool(value))
}

func ListItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ListItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ListItem_SetCaption(obj uintptr, value string) {
	getLazyProc("ListItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func ListItem_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_SetChecked(obj uintptr, value bool) {
	getLazyProc("ListItem_SetChecked").Call(obj, GoBoolToDBool(value))
}

func ListItem_GetCut(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_GetCut").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_SetCut(obj uintptr, value bool) {
	getLazyProc("ListItem_SetCut").Call(obj, GoBoolToDBool(value))
}

func ListItem_GetData(obj uintptr) unsafe.Pointer {
	ret, _, _ := getLazyProc("ListItem_GetData").Call(obj)
	return unsafe.Pointer(ret)
}

func ListItem_SetData(obj uintptr, value unsafe.Pointer) {
	getLazyProc("ListItem_SetData").Call(obj, uintptr(value))
}

func ListItem_GetFocused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_GetFocused").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_SetFocused(obj uintptr, value bool) {
	getLazyProc("ListItem_SetFocused").Call(obj, GoBoolToDBool(value))
}

func ListItem_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetImageIndex").Call(obj)
	return int32(ret)
}

func ListItem_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("ListItem_SetImageIndex").Call(obj, uintptr(value))
}

func ListItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetIndex").Call(obj)
	return int32(ret)
}

func ListItem_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetLeft").Call(obj)
	return int32(ret)
}

func ListItem_SetLeft(obj uintptr, value int32) {
	getLazyProc("ListItem_SetLeft").Call(obj, uintptr(value))
}

func ListItem_GetListView(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItem_GetListView").Call(obj)
	return ret
}

func ListItem_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItem_GetOwner").Call(obj)
	return ret
}

func ListItem_GetPosition(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ListItem_GetPosition").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListItem_SetPosition(obj uintptr, value TPoint) {
	getLazyProc("ListItem_SetPosition").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListItem_GetSelected(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListItem_GetSelected").Call(obj)
	return DBoolToGoBool(ret)
}

func ListItem_SetSelected(obj uintptr, value bool) {
	getLazyProc("ListItem_SetSelected").Call(obj, GoBoolToDBool(value))
}

func ListItem_GetStateIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetStateIndex").Call(obj)
	return int32(ret)
}

func ListItem_SetStateIndex(obj uintptr, value int32) {
	getLazyProc("ListItem_SetStateIndex").Call(obj, uintptr(value))
}

func ListItem_GetSubItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListItem_GetSubItems").Call(obj)
	return ret
}

func ListItem_SetSubItems(obj uintptr, value uintptr) {
	getLazyProc("ListItem_SetSubItems").Call(obj, value)
}

func ListItem_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListItem_GetTop").Call(obj)
	return int32(ret)
}

func ListItem_SetTop(obj uintptr, value int32) {
	getLazyProc("ListItem_SetTop").Call(obj, uintptr(value))
}

func ListItem_GetSubItemImages(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("ListItem_GetSubItemImages").Call(obj, uintptr(Index))
	return int32(ret)
}

func ListItem_SetSubItemImages(obj uintptr, Index int32, value int32) {
	getLazyProc("ListItem_SetSubItemImages").Call(obj, uintptr(Index), uintptr(value))
}

func ListItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListItem_StaticClassType").Call()
	return TClass(r)
}
