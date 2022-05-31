package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartExprParam_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExprParam_Create").Call(obj)
	return ret
}

func ChartExprParam_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartExprParam_Free").Call(obj)
}

func ChartExprParam_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParam_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParam_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ChartExprParam_Assign").Call(obj, Source)
}

func ChartExprParam_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartExprParam_ClassType").Call(obj)
	return TClass(ret)
}

func ChartExprParam_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParam_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParam_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParam_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartExprParam_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartExprParam_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartExprParam_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExprParam_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartExprParam_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParam_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartExprParam_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParam_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParam_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExprParam_GetCollection").Call(obj)
	return ret
}

func ChartExprParam_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartExprParam_SetCollection").Call(obj, value)
}

func ChartExprParam_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParam_GetIndex").Call(obj)
	return int32(ret)
}

func ChartExprParam_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartExprParam_SetIndex").Call(obj, uintptr(value))
}

func ChartExprParam_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParam_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParam_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartExprParam_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func ChartExprParam_StaticClassType() TClass {
	r, _, _ := getLazyProc("ChartExprParam_StaticClassType").Call()
	return TClass(r)
}

func ChartExprParam_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParam_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParam_SetName(obj uintptr, AValue string) {
	_, _, _ = getLazyProc("ChartExprParam_SetName").Call(obj, GoStrToDStr(AValue))
}

func ChartExprParam_GetValue(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExprParam_GetValue").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExprParam_SetValue(obj uintptr, AValue float64) {
	_, _, _ = getLazyProc("ChartExprParam_SetValue").Call(obj, uintptr(unsafe.Pointer(&AValue)))
}
