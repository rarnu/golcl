package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCollection ---------------------------

func ChartExprParams_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_Create").Call(obj)
	return ret
}

func ChartExprParams_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_Free").Call(obj)
}

func ChartExprParams_Owner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_Owner").Call(obj)
	return ret
}

func ChartExprParams_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_Add").Call(obj)
	return ret
}

func ChartExprParams_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_Assign").Call(obj, Source)
}

func ChartExprParams_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_BeginUpdate").Call(obj)
}

func ChartExprParams_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_Clear").Call(obj)
}

func ChartExprParams_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ChartExprParams_Delete").Call(obj, uintptr(Index))
}

func ChartExprParams_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_EndUpdate").Call(obj)
}

func ChartExprParams_FindItemID(obj uintptr, ID int32) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_FindItemID").Call(obj, uintptr(ID))
	return ret
}

func ChartExprParams_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParams_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParams_Insert(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_Insert").Call(obj, uintptr(Index))
	return ret
}

func ChartExprParams_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartExprParams_ClassType").Call(obj)
	return TClass(ret)
}

func ChartExprParams_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParams_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParams_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParams_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartExprParams_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartExprParams_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartExprParams_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExprParams_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartExprParams_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParams_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartExprParams_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartExprParams_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartExprParams_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartExprParams_GetCount").Call(obj)
	return int32(ret)
}

func ChartExprParams_GetItems(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_GetItems").Call(obj, uintptr(Index))
	return ret
}

func ChartExprParams_SetItems(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_SetItems").Call(obj, uintptr(Index), value)
}

func ChartExprParams_StaticClassType() TClass {
	r, _, _ := getLazyProc("ChartExprParams_StaticClassType").Call()
	return TClass(r)
}

func ChartExprParams_AddParam(obj uintptr, AName string, AValue float64) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_AddParam").Call(obj, GoStrToDStr(AName), uintptr(unsafe.Pointer(&AValue)))
	return ret
}

func ChartExprParams_FindParamByName(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_FindParamByName").Call(obj, GoStrToDStr(AName))
	return ret
}

func ChartExprParams_GetParams(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ChartExprParams_GetParams").Call(obj, uintptr(AIndex))
	return ret
}

func ChartExprParams_SetParams(obj uintptr, AIndex int32, AValue uintptr) {
	_, _, _ = getLazyProc("ChartExprParams_SetParams").Call(obj, uintptr(AIndex), AValue)
}

func ChartExprParams_GetValueByName(obj uintptr, AName string) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExprParams_GetValueByName").Call(obj, GoStrToDStr(AName), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExprParams_SetValueByName(obj uintptr, AName string, AValue float64) {
	_, _, _ = getLazyProc("ChartExprParams_SetValueByName").Call(obj, GoStrToDStr(AName), uintptr(unsafe.Pointer(&AValue)))
}
