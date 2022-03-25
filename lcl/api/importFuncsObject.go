package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TObject ---------------------------

func Object_Create() uintptr {
	ret, _, _ := getLazyProc("Object_Create").Call()
	return ret
}

func Object_Free(obj uintptr) {
	_, _, _ = getLazyProc("Object_Free").Call(obj)
}

func Object_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Object_ClassType").Call(obj)
	return TClass(ret)
}

func Object_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Object_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Object_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Object_InstanceSize").Call(obj)
	return int32(ret)
}

func Object_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Object_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Object_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Object_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Object_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Object_GetHashCode").Call(obj)
	return int32(ret)
}

func Object_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Object_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Object_StaticClassType() TClass {
	r, _, _ := getLazyProc("Object_StaticClassType").Call()
	return TClass(r)
}
