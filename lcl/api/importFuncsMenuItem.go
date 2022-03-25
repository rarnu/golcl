package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TMenuItem ---------------------------

func MenuItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MenuItem_Create").Call(obj)
	return ret
}

func MenuItem_Free(obj uintptr) {
	_, _, _ = getLazyProc("MenuItem_Free").Call(obj)
}

func MenuItem_Insert(obj uintptr, Index int32, Item uintptr) {
	_, _, _ = getLazyProc("MenuItem_Insert").Call(obj, uintptr(Index), Item)
}

func MenuItem_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("MenuItem_Delete").Call(obj, uintptr(Index))
}

func MenuItem_Clear(obj uintptr) {
	_, _, _ = getLazyProc("MenuItem_Clear").Call(obj)
}

func MenuItem_Click(obj uintptr) {
	_, _, _ = getLazyProc("MenuItem_Click").Call(obj)
}

func MenuItem_IndexOf(obj uintptr, Item uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_IndexOf").Call(obj, Item)
	return int32(ret)
}

func MenuItem_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_Add(obj uintptr, Item uintptr) {
	_, _, _ = getLazyProc("MenuItem_Add").Call(obj, Item)
}

func MenuItem_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("MenuItem_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func MenuItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("MenuItem_Assign").Call(obj, Source)
}

func MenuItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MenuItem_ClassType").Call(obj)
	return TClass(ret)
}

func MenuItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_InstanceSize").Call(obj)
	return int32(ret)
}

func MenuItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MenuItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MenuItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MenuItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_GetHashCode").Call(obj)
	return int32(ret)
}

func MenuItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_GetHandle(obj uintptr) HMENU {
	ret, _, _ := getLazyProc("MenuItem_GetHandle").Call(obj)
	return ret
}

func MenuItem_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_GetCount").Call(obj)
	return int32(ret)
}

func MenuItem_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetParent").Call(obj)
	return ret
}

func MenuItem_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetAction").Call(obj)
	return ret
}

func MenuItem_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MenuItem_SetAction").Call(obj, value)
}

func MenuItem_GetAutoCheck(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetAutoCheck").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetAutoCheck(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetAutoCheck").Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetBitmap(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetBitmap").Call(obj)
	return ret
}

func MenuItem_SetBitmap(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MenuItem_SetBitmap").Call(obj, value)
}

func MenuItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("MenuItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func MenuItem_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetChecked(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetChecked").Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetDefault(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetDefault").Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetGroupIndex(obj uintptr) uint8 {
	ret, _, _ := getLazyProc("MenuItem_GetGroupIndex").Call(obj)
	return uint8(ret)
}

func MenuItem_SetGroupIndex(obj uintptr, value uint8) {
	_, _, _ = getLazyProc("MenuItem_SetGroupIndex").Call(obj, uintptr(value))
}

func MenuItem_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("MenuItem_SetHint").Call(obj, GoStrToDStr(value))
}

func MenuItem_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_GetImageIndex").Call(obj)
	return int32(ret)
}

func MenuItem_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MenuItem_SetImageIndex").Call(obj, uintptr(value))
}

func MenuItem_GetRadioItem(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetRadioItem").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetRadioItem(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetRadioItem").Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetShortCut(obj uintptr) TShortCut {
	ret, _, _ := getLazyProc("MenuItem_GetShortCut").Call(obj)
	return TShortCut(ret)
}

func MenuItem_SetShortCut(obj uintptr, value TShortCut) {
	_, _, _ = getLazyProc("MenuItem_SetShortCut").Call(obj, uintptr(value))
}

func MenuItem_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("MenuItem_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func MenuItem_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MenuItem_SetVisible").Call(obj, GoBoolToDBool(value))
}

func MenuItem_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("MenuItem_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func MenuItem_SetOnMeasureItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("MenuItem_SetOnMeasureItem").Call(obj, addEventToMap(obj, fn))
}

func MenuItem_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_GetComponentCount").Call(obj)
	return int32(ret)
}

func MenuItem_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MenuItem_GetComponentIndex").Call(obj)
	return int32(ret)
}

func MenuItem_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MenuItem_SetComponentIndex").Call(obj, uintptr(value))
}

func MenuItem_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetOwner").Call(obj)
	return ret
}

func MenuItem_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("MenuItem_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func MenuItem_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("MenuItem_SetName").Call(obj, GoStrToDStr(value))
}

func MenuItem_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("MenuItem_GetTag").Call(obj)
	return int(ret)
}

func MenuItem_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("MenuItem_SetTag").Call(obj, uintptr(value))
}

func MenuItem_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetItems").Call(obj, uintptr(Index))
	return ret
}

func MenuItem_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("MenuItem_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func MenuItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("MenuItem_StaticClassType").Call()
	return TClass(r)
}
