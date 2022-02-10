package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TAnchorSide ---------------------------

func AnchorSide_Assign(obj uintptr, Source uintptr) {
	getLazyProc("AnchorSide_Assign").Call(obj, Source)
}

func AnchorSide_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("AnchorSide_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func AnchorSide_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("AnchorSide_ClassType").Call(obj)
	return TClass(ret)
}

func AnchorSide_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("AnchorSide_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func AnchorSide_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("AnchorSide_InstanceSize").Call(obj)
	return int32(ret)
}

func AnchorSide_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("AnchorSide_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func AnchorSide_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("AnchorSide_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func AnchorSide_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("AnchorSide_GetHashCode").Call(obj)
	return int32(ret)
}

func AnchorSide_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("AnchorSide_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func AnchorSide_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("AnchorSide_GetOwner").Call(obj)
	return ret
}

func AnchorSide_GetKind(obj uintptr) TAnchorKind {
	ret, _, _ := getLazyProc("AnchorSide_GetKind").Call(obj)
	return TAnchorKind(ret)
}

func AnchorSide_GetControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("AnchorSide_GetControl").Call(obj)
	return ret
}

func AnchorSide_SetControl(obj uintptr, value uintptr) {
	getLazyProc("AnchorSide_SetControl").Call(obj, value)
}

func AnchorSide_GetSide(obj uintptr) TAnchorSideReference {
	ret, _, _ := getLazyProc("AnchorSide_GetSide").Call(obj)
	return TAnchorSideReference(ret)
}

func AnchorSide_SetSide(obj uintptr, value TAnchorSideReference) {
	getLazyProc("AnchorSide_SetSide").Call(obj, uintptr(value))
}

func AnchorSide_StaticClassType() TClass {
	r, _, _ := getLazyProc("AnchorSide_StaticClassType").Call()
	return TClass(r)
}
