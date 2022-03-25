package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TComponent ---------------------------

func Component_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Component_Create").Call(obj)
	return ret
}

func Component_Free(obj uintptr) {
	_, _, _ = getLazyProc("Component_Free").Call(obj)
}

func Component_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Component_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Component_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Component_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Component_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Component_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Component_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Component_Assign").Call(obj, Source)
}

func Component_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Component_ClassType").Call(obj)
	return TClass(ret)
}

func Component_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Component_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Component_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Component_InstanceSize").Call(obj)
	return int32(ret)
}

func Component_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Component_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Component_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Component_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Component_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Component_GetHashCode").Call(obj)
	return int32(ret)
}

func Component_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Component_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Component_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Component_GetComponentCount").Call(obj)
	return int32(ret)
}

func Component_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Component_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Component_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Component_SetComponentIndex").Call(obj, uintptr(value))
}

func Component_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Component_GetOwner").Call(obj)
	return ret
}

func Component_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Component_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Component_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Component_SetName").Call(obj, GoStrToDStr(value))
}

func Component_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Component_GetTag").Call(obj)
	return int(ret)
}

func Component_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Component_SetTag").Call(obj, uintptr(value))
}

func Component_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Component_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Component_StaticClassType() TClass {
	r, _, _ := getLazyProc("Component_StaticClassType").Call()
	return TClass(r)
}
