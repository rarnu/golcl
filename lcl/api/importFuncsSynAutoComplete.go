package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynAutoComplete_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_Create").Call(owner)
	return r
}

func SynAutoComplete_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynAutoComplete_Free").Call(obj)
}

func SynAutoComplete_GetAutoCompleteList(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_GetAutoCompleteList").Call(obj)
	return r
}

func SynAutoComplete_SetAutoCompleteList(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynAutoComplete_SetAutoCompleteList").Call(obj, value)
}

func SynAutoComplete_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynAutoComplete_StaticClassType").Call()
	return TClass(r)
}

func SynAutoComplete_ClassName(obj uintptr) string {
	r, _, _ := getLazyProc("SynAutoComplete_ClassName").Call(obj)
	return DStrToGoStr(r)
}

func SynAutoComplete_GetHashCode(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynAutoComplete_GetHashCode").Call(obj)
	return int32(r)
}

func SynAutoComplete_Equals(obj uintptr, other uintptr) bool {
	r, _, _ := getLazyProc("SynAutoComplete_Equals").Call(obj, other)
	return DBoolToGoBool(r)
}

func SynAutoComplete_ClassType(obj uintptr) TClass {
	r, _, _ := getLazyProc("SynAutoComplete_ClassType").Call(obj)
	return TClass(r)
}

func SynAutoComplete_InstanceSize(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynAutoComplete_InstanceSize").Call(obj)
	return int32(r)
}

func SynAutoComplete_InheritsFrom(obj uintptr, classType TClass) bool {
	r, _, _ := getLazyProc("SynAutoComplete_InheritsFrom").Call(obj, uintptr(classType))
	return DBoolToGoBool(r)
}

func SynAutoComplete_FindComponent(obj uintptr, AName string) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_FindComponent").Call(obj, GoStrToDStr(AName))
	return r
}

func SynAutoComplete_GetNamePath(obj uintptr) string {
	r, _, _ := getLazyProc("SynAutoComplete_GetNamePath").Call(obj)
	return DStrToGoStr(r)
}

func SynAutoComplete_HasParent(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAutoComplete_HasParent").Call(obj)
	return DBoolToGoBool(r)
}

func SynAutoComplete_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SynAutoComplete_Assign").Call(obj, Source)
}

func SynAutoComplete_ComponentCount(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynAutoComplete_GetComponentCount").Call(obj)
	return int32(r)
}

func SynAutoComplete_ComponentIndex(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynAutoComplete_GetComponentIndex").Call(obj)
	return int32(r)
}

func SynAutoComplete_SetComponentIndex(obj uintptr, index int32) {
	_, _, _ = getLazyProc("SynAutoComplete_SetComponentIndex").Call(obj, uintptr(index))
}

func SynAutoComplete_Components(obj uintptr, index int32) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_GetComponents").Call(obj, uintptr(index))
	return r
}

func SynAutoComplete_Owner(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_GetOwner").Call(obj)
	return r
}

func SynAutoComplete_GetEditor(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAutoComplete_GetEditor").Call(obj)
	return r
}

func SynAutoComplete_SetEditor(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynAutoComplete_SetEditor").Call(obj, value)
}

func SynAutoComplete_GetEndOfTokenChr(obj uintptr) string {
	r, _, _ := getLazyProc("SynAutoComplete_GetEndOfTokenChr").Call(obj)
	return DStrToGoStr(r)
}

func SynAutoComplete_SetEndOfTokenChr(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynAutoComplete_SetEndOfTokenChr").Call(obj, GoStrToDStr(value))
}

func SynAutoComplete_GetExecCommandID(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynAutoComplete_GetExecCommandID").Call(obj)
	return int32(r)
}

func SynAutoComplete_SetExecCommandID(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynAutoComplete_SetExecCommandID").Call(obj, uintptr(value))
}

func SynAutoComplete_GetName(obj uintptr) string {
	r, _, _ := getLazyProc("SynAutoComplete_GetName").Call(obj)
	return DStrToGoStr(r)
}

func SynAutoComplete_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynAutoComplete_SetName").Call(obj, GoStrToDStr(value))
}

func SynAutoComplete_GetShortCut(obj uintptr) TShortCut {
	r, _, _ := getLazyProc("SynAutoComplete_GetShortCut").Call(obj)
	return TShortCut(r)
}

func SynAutoComplete_SetShortCut(obj uintptr, value TShortCut) {
	_, _, _ = getLazyProc("SynAutoComplete_SetShortCut").Call(obj, uintptr(value))
}

func SynAutoComplete_GetTag(obj uintptr) int {
	r, _, _ := getLazyProc("SynAutoComplete_GetTag").Call(obj)
	return int(r)
}

func SynAutoComplete_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("SynAutoComplete_SetTag").Call(obj, uintptr(value))
}
