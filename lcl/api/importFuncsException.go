package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- Exception ---------------------------

func Exception_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Exception_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Exception_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Exception_ClassType").Call(obj)
	return TClass(ret)
}

func Exception_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Exception_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Exception_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Exception_InstanceSize").Call(obj)
	return int32(ret)
}

func Exception_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Exception_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Exception_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Exception_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Exception_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Exception_GetHashCode").Call(obj)
	return int32(ret)
}

func Exception_GetMessage(obj uintptr) string {
	ret, _, _ := getLazyProc("Exception_GetMessage").Call(obj)
	return DStrToGoStr(ret)
}

func Exception_SetMessage(obj uintptr, value string) {
	_, _, _ = getLazyProc("Exception_SetMessage").Call(obj, GoStrToDStr(value))
}

func Exception_StaticClassType() TClass {
	r, _, _ := getLazyProc("Exception_StaticClassType").Call()
	return TClass(r)
}
