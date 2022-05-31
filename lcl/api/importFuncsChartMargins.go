package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TChart ---------------------------

func ChartMargins_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMargins_Create").Call(obj)
	return ret
}

func ChartMargins_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartMargins_Free").Call(obj)
}

func ChartMargins_GetValue(obj uintptr, AIndex int32) int32 {
	ret, _, _ := getLazyProc("ChartMargins_GetValue").Call(obj, uintptr(AIndex))
	return int32(ret)
}

func ChartMargins_SetValue(obj uintptr, AIndex int32, AValue TChartDistance) {
	_, _, _ = getLazyProc("ChartMargins_SetValue").Call(obj, uintptr(AIndex), uintptr(AValue))
}

func ChartMargins_Assign(obj uintptr, source uintptr) {
	_, _, _ = getLazyProc("ChartMargins_Assign").Call(obj, source)
}

func ChartMargins_ExpandRectScaled(obj uintptr, ADrawer uintptr, ARect TRect) {
	_, _, _ = getLazyProc("ChartMargins_ExpandRectScaled").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ARect)))
}

func ChartMargins_GetData(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartMargins_GetData").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMargins_GetLeft(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMargins_GetLeft").Call(obj)
	return TChartDistance(ret)
}

func ChartMargins_SetLeft(obj uintptr, AValue TChartDistance) {
	_, _, _ = getLazyProc("ChartMargins_SetLeft").Call(obj, uintptr(AValue))
}

func ChartMargins_GetTop(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMargins_GetTop").Call(obj)
	return TChartDistance(ret)
}

func ChartMargins_SetTop(obj uintptr, AValue TChartDistance) {
	_, _, _ = getLazyProc("ChartMargins_SetTop").Call(obj, uintptr(AValue))
}

func ChartMargins_GetRight(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMargins_GetRight").Call(obj)
	return TChartDistance(ret)
}

func ChartMargins_SetRight(obj uintptr, AValue TChartDistance) {
	_, _, _ = getLazyProc("ChartMargins_SetRight").Call(obj, uintptr(AValue))
}

func ChartMargins_GetBottom(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMargins_GetBottom").Call(obj)
	return TChartDistance(ret)
}

func ChartMargins_SetBottom(obj uintptr, AValue TChartDistance) {
	_, _, _ = getLazyProc("ChartMargins_SetBottom").Call(obj, uintptr(AValue))
}

func ChartMargins_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartMargins_StaticClassType").Call()
	return TClass(ret)
}

func ChartMargins_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMargins_GetOwner").Call(obj)
	return ret
}

func ChartMargins_SetOwner(obj uintptr, AValue uintptr) {
	_, _, _ = getLazyProc("ChartMargins_SetOwner").Call(obj, AValue)
}

func ChartMargins_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMargins_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMargins_SetVisible(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("ChartMargins_SetVisible").Call(obj, GoBoolToDBool(AValue))
}
