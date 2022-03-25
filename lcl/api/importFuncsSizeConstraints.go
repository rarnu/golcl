package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSizeConstraints ---------------------------

func SizeConstraints_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SizeConstraints_Assign").Call(obj, Source)
}

func SizeConstraints_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SizeConstraints_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SizeConstraints_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SizeConstraints_ClassType").Call(obj)
	return TClass(ret)
}

func SizeConstraints_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SizeConstraints_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SizeConstraints_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SizeConstraints_InstanceSize").Call(obj)
	return int32(ret)
}

func SizeConstraints_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SizeConstraints_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SizeConstraints_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SizeConstraints_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SizeConstraints_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SizeConstraints_GetHashCode").Call(obj)
	return int32(ret)
}

func SizeConstraints_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SizeConstraints_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SizeConstraints_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SizeConstraints_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func SizeConstraints_GetMaxHeight(obj uintptr) TConstraintSize {
	ret, _, _ := getLazyProc("SizeConstraints_GetMaxHeight").Call(obj)
	return TConstraintSize(ret)
}

func SizeConstraints_SetMaxHeight(obj uintptr, value TConstraintSize) {
	_, _, _ = getLazyProc("SizeConstraints_SetMaxHeight").Call(obj, uintptr(value))
}

func SizeConstraints_GetMaxWidth(obj uintptr) TConstraintSize {
	ret, _, _ := getLazyProc("SizeConstraints_GetMaxWidth").Call(obj)
	return TConstraintSize(ret)
}

func SizeConstraints_SetMaxWidth(obj uintptr, value TConstraintSize) {
	_, _, _ = getLazyProc("SizeConstraints_SetMaxWidth").Call(obj, uintptr(value))
}

func SizeConstraints_GetMinHeight(obj uintptr) TConstraintSize {
	ret, _, _ := getLazyProc("SizeConstraints_GetMinHeight").Call(obj)
	return TConstraintSize(ret)
}

func SizeConstraints_SetMinHeight(obj uintptr, value TConstraintSize) {
	_, _, _ = getLazyProc("SizeConstraints_SetMinHeight").Call(obj, uintptr(value))
}

func SizeConstraints_GetMinWidth(obj uintptr) TConstraintSize {
	ret, _, _ := getLazyProc("SizeConstraints_GetMinWidth").Call(obj)
	return TConstraintSize(ret)
}

func SizeConstraints_SetMinWidth(obj uintptr, value TConstraintSize) {
	_, _, _ = getLazyProc("SizeConstraints_SetMinWidth").Call(obj, uintptr(value))
}

func SizeConstraints_StaticClassType() TClass {
	r, _, _ := getLazyProc("SizeConstraints_StaticClassType").Call()
	return TClass(r)
}
