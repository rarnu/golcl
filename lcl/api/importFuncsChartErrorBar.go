package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartErrorBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartErrorBar_Create").Call(obj)
	return ret
}

func ChartErrorBar_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartErrorBar_Free").Call(obj)
}

func ChartErrorBar_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartErrorBar_StaticClassType").Call()
	return TClass(ret)
}

func ChartErrorBar_GetPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartErrorBar_GetPen").Call(obj)
	return ret
}

func ChartErrorBar_SetPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartErrorBar_SetPen").Call(obj, value)
}

func ChartErrorBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartErrorBar_GetWidth").Call(obj)
	return int32(ret)
}

func ChartErrorBar_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartErrorBar_SetWidth").Call(obj, uintptr(value))
}

func ChartErrorBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartErrorBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartErrorBar_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartErrorBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartErrorBar_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartErrorBar_SetOwner").Call(obj, AOwner)
}
