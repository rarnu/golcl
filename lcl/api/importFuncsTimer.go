package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTimer ---------------------------

func Timer_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Timer_Create").Call(obj)
	return ret
}

func Timer_Free(obj uintptr) {
	_, _, _ = getLazyProc("Timer_Free").Call(obj)
}

func Timer_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Timer_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Timer_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Timer_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Timer_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Timer_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Timer_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Timer_Assign").Call(obj, Source)
}

func Timer_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Timer_ClassType").Call(obj)
	return TClass(ret)
}

func Timer_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Timer_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Timer_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Timer_InstanceSize").Call(obj)
	return int32(ret)
}

func Timer_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Timer_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Timer_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Timer_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Timer_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Timer_GetHashCode").Call(obj)
	return int32(ret)
}

func Timer_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Timer_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Timer_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Timer_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Timer_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Timer_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Timer_GetInterval(obj uintptr) uint32 {
	ret, _, _ := getLazyProc("Timer_GetInterval").Call(obj)
	return uint32(ret)
}

func Timer_SetInterval(obj uintptr, value uint32) {
	_, _, _ = getLazyProc("Timer_SetInterval").Call(obj, uintptr(value))
}

func Timer_SetOnTimer(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Timer_SetOnTimer").Call(obj, addEventToMap(obj, fn))
}

func Timer_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Timer_GetComponentCount").Call(obj)
	return int32(ret)
}

func Timer_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Timer_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Timer_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Timer_SetComponentIndex").Call(obj, uintptr(value))
}

func Timer_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Timer_GetOwner").Call(obj)
	return ret
}

func Timer_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Timer_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Timer_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Timer_SetName").Call(obj, GoStrToDStr(value))
}

func Timer_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Timer_GetTag").Call(obj)
	return int(ret)
}

func Timer_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Timer_SetTag").Call(obj, uintptr(value))
}

func Timer_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Timer_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Timer_StaticClassType() TClass {
	r, _, _ := getLazyProc("Timer_StaticClassType").Call()
	return TClass(r)
}
