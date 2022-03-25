package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TAction ---------------------------

func Action_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Action_Create").Call(obj)
	return ret
}

func Action_Free(obj uintptr) {
	_, _, _ = getLazyProc("Action_Free").Call(obj)
}

func Action_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_Update(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_Update").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Action_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Action_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Action_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Action_Assign").Call(obj, Source)
}

func Action_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Action_ClassType").Call(obj)
	return TClass(ret)
}

func Action_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Action_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_InstanceSize").Call(obj)
	return int32(ret)
}

func Action_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Action_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Action_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Action_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetHashCode").Call(obj)
	return int32(ret)
}

func Action_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Action_GetAutoCheck(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_GetAutoCheck").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_SetAutoCheck(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Action_SetAutoCheck").Call(obj, GoBoolToDBool(value))
}

func Action_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func Action_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("Action_SetCaption").Call(obj, GoStrToDStr(value))
}

func Action_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_SetChecked(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Action_SetChecked").Call(obj, GoBoolToDBool(value))
}

func Action_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Action_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Action_GetGroupIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetGroupIndex").Call(obj)
	return int32(ret)
}

func Action_SetGroupIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Action_SetGroupIndex").Call(obj, uintptr(value))
}

func Action_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Action_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Action_SetHint").Call(obj, GoStrToDStr(value))
}

func Action_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetImageIndex").Call(obj)
	return int32(ret)
}

func Action_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Action_SetImageIndex").Call(obj, uintptr(value))
}

func Action_GetShortCut(obj uintptr) TShortCut {
	ret, _, _ := getLazyProc("Action_GetShortCut").Call(obj)
	return TShortCut(ret)
}

func Action_SetShortCut(obj uintptr, value TShortCut) {
	_, _, _ = getLazyProc("Action_SetShortCut").Call(obj, uintptr(value))
}

func Action_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Action_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Action_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Action_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Action_SetOnExecute(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Action_SetOnExecute").Call(obj, addEventToMap(obj, fn))
}

func Action_SetOnUpdate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Action_SetOnUpdate").Call(obj, addEventToMap(obj, fn))
}

func Action_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetIndex").Call(obj)
	return int32(ret)
}

func Action_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Action_SetIndex").Call(obj, uintptr(value))
}

func Action_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetComponentCount").Call(obj)
	return int32(ret)
}

func Action_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Action_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Action_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Action_SetComponentIndex").Call(obj, uintptr(value))
}

func Action_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Action_GetOwner").Call(obj)
	return ret
}

func Action_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Action_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Action_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Action_SetName").Call(obj, GoStrToDStr(value))
}

func Action_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Action_GetTag").Call(obj)
	return int(ret)
}

func Action_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Action_SetTag").Call(obj, uintptr(value))
}

func Action_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Action_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Action_StaticClassType() TClass {
	r, _, _ := getLazyProc("Action_StaticClassType").Call()
	return TClass(r)
}
