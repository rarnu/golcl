package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartArrow_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartArrow_Create").Call(obj)
	return ret
}

func ChartArrow_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartArrow_Free").Call(obj)
}

func ChartArrow_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartArrow_StaticClassType").Call()
	return TClass(ret)
}

func ChartArrow_GetBaseLength(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartArrow_GetBaseLength").Call(obj)
	return TChartDistance(ret)
}

func ChartArrow_SetBaseLength(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartArrow_SetBaseLength").Call(obj, uintptr(value))
}

func ChartArrow_GetInverted(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartArrow_GetInverted").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartArrow_SetInverted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartArrow_SetInverted").Call(obj, GoBoolToDBool(value))
}

func ChartArrow_GetLength(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartArrow_GetLength").Call(obj)
	return TChartDistance(ret)
}

func ChartArrow_SetLength(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartArrow_SetLength").Call(obj, uintptr(value))
}

func ChartArrow_GetWidth(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartArrow_GetWidth").Call(obj)
	return TChartDistance(ret)
}

func ChartArrow_SetWidth(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartArrow_SetWidth").Call(obj, uintptr(value))
}

func ChartArrow_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartArrow_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartArrow_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartArrow_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartArrow_Draw(obj uintptr, ADrawer uintptr, AEndPos TPoint, AAngle float64, APen uintptr) {
	_, _, _ = getLazyProc("ChartArrow_Draw").Call(obj, ADrawer, uintptr(unsafe.Pointer(&AEndPos)), uintptr(unsafe.Pointer(&AAngle)), APen)
}

func ChartArrow_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartArrow_GetOwner").Call(obj)
	return ret
}

func ChartArrow_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartArrow_SetOwner").Call(obj, AOwner)
}
