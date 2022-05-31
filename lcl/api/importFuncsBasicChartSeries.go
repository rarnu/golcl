package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func BasicChartSeries_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BasicChartSeries_Create").Call(obj)
	return ret
}

func BasicChartSeries_Free(obj uintptr) {
	_, _, _ = getLazyProc("BasicChartSeries_Free").Call(obj)
}

func BasicChartSeries_StaticClassType() TClass {
	ret, _, _ := getLazyProc("BasicChartSeries_StaticClassType").Call()
	return TClass(ret)
}

func BasicChartSeries_GetActive(obj uintptr) bool {
	ret, _, _ := getLazyProc("BasicChartSeries_GetActive").Call(obj)
	return DBoolToGoBool(ret)
}

func BasicChartSeries_SetActive(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BasicChartSeries_SetActive").Call(obj, GoBoolToDBool(value))
}

func BasicChartSeries_GetDepth(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("BasicChartSeries_GetDepth").Call(obj)
	return TChartDistance(ret)
}

func BasicChartSeries_SetDepth(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("BasicChartSeries_SetDepth").Call(obj, uintptr(value))
}

func BasicChartSeries_GetDragOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BasicChartSeries_GetDragOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_SetDragOrigin(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("BasicChartSeries_SetDragOrigin").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func BasicChartSeries_GetShadow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BasicChartSeries_GetShadow").Call(obj)
	return ret
}

func BasicChartSeries_SetShadow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BasicChartSeries_SetShadow").Call(obj, value)
}

func BasicChartSeries_GetTransparency(obj uintptr) TChartTransparency {
	ret, _, _ := getLazyProc("BasicChartSeries_GetTransparency").Call(obj)
	return TChartTransparency(ret)
}

func BasicChartSeries_SetTransparency(obj uintptr, value TChartTransparency) {
	_, _, _ = getLazyProc("BasicChartSeries_SetTransparency").Call(obj, uintptr(value))
}

func BasicChartSeries_GetZPosition(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("BasicChartSeries_GetZPosition").Call(obj)
	return TChartDistance(ret)
}

func BasicChartSeries_SetZPosition(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("BasicChartSeries_SetZPosition").Call(obj, uintptr(value))
}

func BasicChartSeries_AxisToGraphX(obj uintptr, AX float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("BasicChartSeries_AxisToGraphX").Call(obj, uintptr(unsafe.Pointer(&AX)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_AxisToGraphY(obj uintptr, AY float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("BasicChartSeries_AxisToGraphY").Call(obj, uintptr(unsafe.Pointer(&AY)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_GraphToAxisX(obj uintptr, AX float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("BasicChartSeries_GraphToAxisX").Call(obj, uintptr(unsafe.Pointer(&AX)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_GraphToAxisY(obj uintptr, AY float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("BasicChartSeries_GraphToAxisY").Call(obj, uintptr(unsafe.Pointer(&AY)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_Draw(obj uintptr, ADrawer uintptr) {
	_, _, _ = getLazyProc("BasicChartSeries_Draw").Call(obj, ADrawer)
}

func BasicChartSeries_GetAxisBounds(obj uintptr, AAxis uintptr, AMin, AMax float64) bool {
	ret, _, _ := getLazyProc("BasicChartSeries_GetAxisBounds").Call(obj, AAxis, uintptr(unsafe.Pointer(&AMin)), uintptr(unsafe.Pointer(&AMax)))
	return DBoolToGoBool(ret)
}

func BasicChartSeries_GetGraphBounds(obj uintptr) TDoubleRect {
	var ret TDoubleRect
	_, _, _ = getLazyProc("BasicChartSeries_GetGraphBounds").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BasicChartSeries_IsEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("BasicChartSeries_IsEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func BasicChartSeries_MovePoint(obj uintptr, AIndex int32, ANewPos TPoint) {
	_, _, _ = getLazyProc("BasicChartSeries_MovePoint").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&ANewPos)))
}

func BasicChartSeries_MovePointEx(obj uintptr, AIndex int32, AXIndex, AYIndex int32, ANewPos TDoublePoint) {
	_, _, _ = getLazyProc("BasicChartSeries_MovePointEx").Call(obj, uintptr(AIndex), uintptr(AXIndex), uintptr(AYIndex), uintptr(unsafe.Pointer(&ANewPos)))
}

func BasicChartSeries_UpdateBiDiMode(obj uintptr) {
	_, _, _ = getLazyProc("BasicChartSeries_UpdateBiDiMode").Call(obj)
}
