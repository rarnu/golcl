package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartExtent_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExtent_Create").Call(obj)
	return ret
}

func ChartExtent_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartExtent_Free").Call(obj)
}

func ChartExtent_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartExtent_StaticClassType").Call()
	return TClass(ret)
}

func ChartExtent_GetUseXMax(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExtent_GetUseXMax").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartExtent_SetUseXMax(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartExtent_SetUseXMax").Call(obj, GoBoolToDBool(value))
}

func ChartExtent_GetUseXMin(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExtent_GetUseXMin").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartExtent_SetUseXMin(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartExtent_SetUseXMin").Call(obj, GoBoolToDBool(value))
}

func ChartExtent_GetUseYMax(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExtent_GetUseYMax").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartExtent_SetUseYMax(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartExtent_SetUseYMax").Call(obj, GoBoolToDBool(value))
}

func ChartExtent_GetUseYMin(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExtent_GetUseYMin").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartExtent_SetUseYMin(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartExtent_SetUseYMin").Call(obj, GoBoolToDBool(value))
}

func ChartExtent_GetXMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExtent_GetXMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExtent_SetXMax(obj uintptr, value float64) {
	_, _, _ = getLazyProc("ChartExtent_SetXMax").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ChartExtent_GetXMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExtent_GetXMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExtent_SetXMin(obj uintptr, value float64) {
	_, _, _ = getLazyProc("ChartExtent_SetXMin").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ChartExtent_GetYMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExtent_GetYMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExtent_SetYMax(obj uintptr, value float64) {
	_, _, _ = getLazyProc("ChartExtent_SetYMax").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ChartExtent_GetYMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartExtent_GetYMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartExtent_SetYMin(obj uintptr, value float64) {
	_, _, _ = getLazyProc("ChartExtent_SetYMin").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ChartExtent_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartExtent_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartExtent_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartExtent_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartExtent_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartExtent_Assign").Call(obj, ASource)
}

func ChartExtent_CheckBoundsOrder(obj uintptr) {
	_, _, _ = getLazyProc("ChartExtent_CheckBoundsOrder").Call(obj)
}

func ChartExtent_FixTo(obj uintptr, ABounds TDoubleRect) {
	_, _, _ = getLazyProc("ChartExtent_FixTo").Call(obj, uintptr(unsafe.Pointer(&ABounds)))
}

func ChartExtent_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartExtent_GetOwner").Call(obj)
	return ret
}

func ChartExtent_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartExtent_SetOwner").Call(obj, AOwner)
}
