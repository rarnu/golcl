package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TIconOptions ---------------------------

func IconOptions_Assign(obj uintptr, Source uintptr) {
	getLazyProc("IconOptions_Assign").Call(obj, Source)
}

func IconOptions_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("IconOptions_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func IconOptions_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("IconOptions_ClassType").Call(obj)
	return TClass(ret)
}

func IconOptions_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("IconOptions_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func IconOptions_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("IconOptions_InstanceSize").Call(obj)
	return int32(ret)
}

func IconOptions_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("IconOptions_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func IconOptions_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("IconOptions_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func IconOptions_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("IconOptions_GetHashCode").Call(obj)
	return int32(ret)
}

func IconOptions_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("IconOptions_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func IconOptions_GetArrangement(obj uintptr) TIconArrangement {
	ret, _, _ := getLazyProc("IconOptions_GetArrangement").Call(obj)
	return TIconArrangement(ret)
}

func IconOptions_SetArrangement(obj uintptr, value TIconArrangement) {
	getLazyProc("IconOptions_SetArrangement").Call(obj, uintptr(value))
}

func IconOptions_GetAutoArrange(obj uintptr) bool {
	ret, _, _ := getLazyProc("IconOptions_GetAutoArrange").Call(obj)
	return DBoolToGoBool(ret)
}

func IconOptions_SetAutoArrange(obj uintptr, value bool) {
	getLazyProc("IconOptions_SetAutoArrange").Call(obj, GoBoolToDBool(value))
}

func IconOptions_StaticClassType() TClass {
	r, _, _ := getLazyProc("IconOptions_StaticClassType").Call()
	return TClass(r)
}
