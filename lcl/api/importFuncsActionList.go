package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TActionList ---------------------------

func ActionList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ActionList_Create").Call(obj)
	return ret
}

func ActionList_Free(obj uintptr) {
	getLazyProc("ActionList_Free").Call(obj)
}

func ActionList_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ActionList_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ActionList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ActionList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ActionList_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ActionList_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ActionList_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ActionList_Assign").Call(obj, Source)
}

func ActionList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ActionList_ClassType").Call(obj)
	return TClass(ret)
}

func ActionList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ActionList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ActionList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ActionList_InstanceSize").Call(obj)
	return int32(ret)
}

func ActionList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ActionList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ActionList_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ActionList_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ActionList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ActionList_GetHashCode").Call(obj)
	return int32(ret)
}

func ActionList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ActionList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ActionList_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ActionList_GetImages").Call(obj)
	return ret
}

func ActionList_SetImages(obj uintptr, value uintptr) {
	getLazyProc("ActionList_SetImages").Call(obj, value)
}

func ActionList_GetState(obj uintptr) TActionListState {
	ret, _, _ := getLazyProc("ActionList_GetState").Call(obj)
	return TActionListState(ret)
}

func ActionList_SetState(obj uintptr, value TActionListState) {
	getLazyProc("ActionList_SetState").Call(obj, uintptr(value))
}

func ActionList_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ActionList_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ActionList_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ActionList_GetComponentCount").Call(obj)
	return int32(ret)
}

func ActionList_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ActionList_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ActionList_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ActionList_SetComponentIndex").Call(obj, uintptr(value))
}

func ActionList_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ActionList_GetOwner").Call(obj)
	return ret
}

func ActionList_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ActionList_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ActionList_SetName(obj uintptr, value string) {
	getLazyProc("ActionList_SetName").Call(obj, GoStrToDStr(value))
}

func ActionList_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ActionList_GetTag").Call(obj)
	return int(ret)
}

func ActionList_SetTag(obj uintptr, value int) {
	getLazyProc("ActionList_SetTag").Call(obj, uintptr(value))
}

func ActionList_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ActionList_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ActionList_StaticClassType() TClass {
	r, _, _ := getLazyProc("ActionList_StaticClassType").Call()
	return TClass(r)
}
