package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartRange_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartRange_Create").Call(obj)
	return ret
}

func ChartRange_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartRange_Free").Call(obj)
}

func ChartRange_Assign(obj uintptr, source uintptr) {
	_, _, _ = getLazyProc("ChartRange_Assign").Call(obj, source)
}

func ChartRange_CheckBoundsOrder(obj uintptr) {
	_, _, _ = getLazyProc("ChartRange_CheckBoundsOrder").Call(obj)
}

func ChartRange_Intersect(obj uintptr, AMin float64, AMax float64) {
	_, _, _ = getLazyProc("ChartRange_Intersect").Call(obj, uintptr(unsafe.Pointer(&AMin)), uintptr(unsafe.Pointer(&AMax)))
}

func ChartRange_GetMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartRange_GetMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartRange_SetMax(obj uintptr, AValue float64) {
	_, _, _ = getLazyProc("ChartRange_SetMax").Call(obj, uintptr(unsafe.Pointer(&AValue)))
}

func ChartRange_GetMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartRange_GetMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartRange_SetMin(obj uintptr, AValue float64) {
	_, _, _ = getLazyProc("ChartRange_SetMin").Call(obj, uintptr(unsafe.Pointer(&AValue)))
}

func ChartRange_GetUseMax(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartRange_GetUseMax").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartRange_SetUseMax(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("ChartRange_SetUseMax").Call(obj, GoBoolToDBool(AValue))
}

func ChartRange_GetUseMin(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartRange_GetUseMin").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartRange_SetUseMin(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("ChartRange_SetUseMin").Call(obj, GoBoolToDBool(AValue))
}

func ChartRange_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartRange_StaticClassType").Call()
	return TClass(ret)
}

func ChartRange_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartRange_GetOwner").Call(obj)
	return ret
}

func ChartRange_SetOwner(obj uintptr, AValue uintptr) {
	_, _, _ = getLazyProc("ChartRange_SetOwner").Call(obj, AValue)
}

func ChartRange_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartRange_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartRange_SetVisible(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("ChartRange_SetVisible").Call(obj, GoBoolToDBool(AValue))
}
