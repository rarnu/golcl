package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartAxisIntervalParams_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_Create").Call(obj)
	return ret
}

func ChartAxisIntervalParams_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_Free").Call(obj)
}

func ChartAxisIntervalParams_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_StaticClassType").Call()
	return TClass(ret)
}

func ChartAxisIntervalParams_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetCount").Call(obj)
	return int32(ret)
}

func ChartAxisIntervalParams_SetCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetCount").Call(obj, uintptr(value))
}

func ChartAxisIntervalParams_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetMaxLength").Call(obj)
	return int32(ret)
}

func ChartAxisIntervalParams_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetMaxLength").Call(obj, uintptr(value))
}

func ChartAxisIntervalParams_GetMinLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetMinLength").Call(obj)
	return int32(ret)
}

func ChartAxisIntervalParams_SetMinLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetMinLength").Call(obj, uintptr(value))
}

func ChartAxisIntervalParams_GetNiceSteps(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetNiceSteps").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisIntervalParams_SetNiceSteps(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetNiceSteps").Call(obj, GoStrToDStr(value))
}

func ChartAxisIntervalParams_GetOptions(obj uintptr) TAxisIntervalParamOptions {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetOptions").Call(obj)
	return TAxisIntervalParamOptions(ret)
}

func ChartAxisIntervalParams_SetOptions(obj uintptr, value TAxisIntervalParamOptions) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetOptions").Call(obj, uintptr(value))
}

func ChartAxisIntervalParams_GetTolerance(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartAxisIntervalParams_GetTolerance").Call(obj)
	return Cardinal(ret)
}

func ChartAxisIntervalParams_SetTolerance(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartAxisIntervalParams_SetTolerance").Call(obj, uintptr(value))
}
