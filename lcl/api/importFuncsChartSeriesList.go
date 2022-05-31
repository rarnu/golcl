package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartSeriesList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartSeriesList_Create").Call(obj)
	return ret
}

func ChartSeriesList_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartSeriesList_Free").Call(obj)
}

func ChartSeriesList_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartSeriesList_StaticClassType").Call()
	return TClass(ret)
}

func ChartSeriesList_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ChartSeriesList_Clear").Call(obj)
}

func ChartSeriesList_Count(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartSeriesList_Count").Call(obj)
	return int32(ret)
}

func ChartSeriesList_GetEnumerator(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartSeriesList_GetEnumerator").Call(obj)
	return ret
}

func ChartSeriesList_UpdateBiDiMode(obj uintptr) {
	_, _, _ = getLazyProc("ChartSeriesList_UpdateBiDiMode").Call(obj)
}

func ChartSeriesList_GetItems(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ChartSeriesList_GetItems").Call(obj, uintptr(AIndex))
	return ret
}
