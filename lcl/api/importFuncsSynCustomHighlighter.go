package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynCustomHighlighter_GetDefaultFilter(obj uintptr) string {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetDefaultFilter").Call(obj)
	return DStrToGoStr(r)
}

func SynCustomHighlighter_SetDefaultFilter(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetDefaultFilter").Call(obj, GoStrToDStr(value))
}

func SynCustomHighlighter_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynCustomHighlighter_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynCustomHighlighter_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynCustomHighlighter_StaticClassType").Call()
	return TClass(r)
}

func SynCustomHighlighter_ClassName(obj uintptr) string {
	r, _, _ := getLazyProc("SynCustomHighlighter_ClassName").Call(obj)
	return DStrToGoStr(r)
}

func SynCustomHighlighter_GetHashCode(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetHashCode").Call(obj)
	return int32(r)
}

func SynCustomHighlighter_Equals(obj uintptr, other uintptr) bool {
	r, _, _ := getLazyProc("SynCustomHighlighter_Equals").Call(obj, other)
	return DBoolToGoBool(r)
}

func SynCustomHighlighter_ClassType(obj uintptr) TClass {
	r, _, _ := getLazyProc("SynCustomHighlighter_ClassType").Call(obj)
	return TClass(r)
}

func SynCustomHighlighter_InstanceSize(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCustomHighlighter_InstanceSize").Call(obj)
	return int32(r)
}

func SynCustomHighlighter_InheritsFrom(obj uintptr, classType TClass) bool {
	r, _, _ := getLazyProc("SynCustomHighlighter_InheritsFrom").Call(obj, uintptr(classType))
	return DBoolToGoBool(r)
}

func SynCustomHighlighter_FindComponent(obj uintptr, AName string) uintptr {
	r, _, _ := getLazyProc("SynCustomHighlighter_FindComponent").Call(obj, GoStrToDStr(AName))
	return r
}

func SynCustomHighlighter_GetNamePath(obj uintptr) string {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetNamePath").Call(obj)
	return DStrToGoStr(r)
}

func SynCustomHighlighter_HasParent(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCustomHighlighter_HasParent").Call(obj)
	return DBoolToGoBool(r)
}

func SynCustomHighlighter_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SynCustomHighlighter_Assign").Call(obj, Source)
}

func SynCustomHighlighter_ComponentCount(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetComponentCount").Call(obj)
	return int32(r)
}

func SynCustomHighlighter_ComponentIndex(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetComponentIndex").Call(obj)
	return int32(r)
}

func SynCustomHighlighter_SetComponentIndex(obj uintptr, index int32) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetComponentIndex").Call(obj, uintptr(index))
}

func SynCustomHighlighter_Components(obj uintptr, index int32) uintptr {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetComponents").Call(obj, uintptr(index))
	return r
}

func SynCustomHighlighter_Owner(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetOwner").Call(obj)
	return r
}

func SynCustomHighlighter_Name(obj uintptr) string {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetName").Call(obj)
	return DStrToGoStr(r)
}

func SynCustomHighlighter_SetName(obj uintptr, name string) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetName").Call(obj, GoStrToDStr(name))
}

func SynCustomHighlighter_Tag(obj uintptr) int {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetTag").Call(obj)
	return int(r)
}

func SynCustomHighlighter_SetTag(obj uintptr, tag int) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetTag").Call(obj, uintptr(tag))
}
