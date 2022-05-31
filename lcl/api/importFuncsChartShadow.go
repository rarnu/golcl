package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartShadow_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartShadow_Create").Call(obj)
	return ret
}

func ChartShadow_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartShadow_Free").Call(obj)
}

func ChartShadow_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartShadow_StaticClassType").Call()
	return TClass(ret)
}

func ChartShadow_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ChartShadow_GetColor").Call(obj)
	return TColor(ret)
}

func ChartShadow_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ChartShadow_SetColor").Call(obj, uintptr(value))
}

func ChartShadow_GetOffsetX(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartShadow_GetOffsetX").Call(obj)
	return int32(ret)
}

func ChartShadow_SetOffsetX(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartShadow_SetOffsetX").Call(obj, uintptr(value))
}

func ChartShadow_GetOffsetY(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartShadow_GetOffsetY").Call(obj)
	return int32(ret)
}

func ChartShadow_SetOffsetY(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartShadow_SetOffsetY").Call(obj, uintptr(value))
}

func ChartShadow_GetTransparency(obj uintptr) TChartTransparency {
	ret, _, _ := getLazyProc("ChartShadow_GetTransparency").Call(obj)
	return TChartTransparency(ret)
}

func ChartShadow_SetTransparency(obj uintptr, value TChartTransparency) {
	_, _, _ = getLazyProc("ChartShadow_SetTransparency").Call(obj, uintptr(value))
}

func ChartShadow_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartShadow_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartShadow_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartShadow_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartShadow_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartShadow_GetOwner").Call(obj)
	return ret
}

func ChartShadow_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartShadow_SetOwner").Call(obj, AOwner)
}
