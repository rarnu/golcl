package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func PieSeries_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_Create").Call(obj)
	return ret
}

func PieSeries_Free(obj uintptr) {
	_, _, _ = getLazyProc("PieSeries_Free").Call(obj)
}

func PieSeries_StaticClassType() TClass {
	ret, _, _ := getLazyProc("PieSeries_StaticClassType").Call()
	return TClass(ret)
}

func PieSeries_GetEdgePen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetEdgePen").Call(obj)
	return ret
}

func PieSeries_SetEdgePen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PieSeries_SetEdgePen").Call(obj, value)
}

func PieSeries_GetExploded(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_GetExploded").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_SetExploded(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PieSeries_SetExploded").Call(obj, GoBoolToDBool(value))
}

func PieSeries_GetFixedRadius(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("PieSeries_GetFixedRadius").Call(obj)
	return TChartDistance(ret)
}

func PieSeries_SetFixedRadius(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("PieSeries_SetFixedRadius").Call(obj, uintptr(value))
}

func PieSeries_GetMarkDistancePercent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_GetMarkDistancePercent").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_SetMarkDistancePercent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PieSeries_SetMarkDistancePercent").Call(obj, GoBoolToDBool(value))
}

func PieSeries_GetMarkPositions(obj uintptr) TPieMarkPositions {
	ret, _, _ := getLazyProc("PieSeries_GetMarkPositions").Call(obj)
	return TPieMarkPositions(ret)
}

func PieSeries_SetMarkPositions(obj uintptr, value TPieMarkPositions) {
	_, _, _ = getLazyProc("PieSeries_SetMarkPositions").Call(obj, uintptr(value))
}

func PieSeries_GetRotateLabels(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_GetRotateLabels").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_SetRotateLabels(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PieSeries_SetRotateLabels").Call(obj, GoBoolToDBool(value))
}

func PieSeries_GetMarks(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetMarks").Call(obj)
	return ret
}

func PieSeries_SetMarks(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PieSeries_SetMarks").Call(obj, value)
}

func PieSeries_SetOnGetMark(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PieSeries_SetOnGetMark").Call(obj, addEventToMap(obj, fn))
}

func PieSeries_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("PieSeries_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func PieSeries_SetTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("PieSeries_SetTitle").Call(obj, GoStrToDStr(value))
}

func PieSeries_GetLegend(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetLegend").Call(obj)
	return ret
}

func PieSeries_SetLegend(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PieSeries_SetLegend").Call(obj, value)
}

func PieSeries_GetActive(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_GetActive").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_SetActive(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PieSeries_SetActive").Call(obj, GoBoolToDBool(value))
}

func PieSeries_GetDepth(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("PieSeries_GetDepth").Call(obj)
	return TChartDistance(ret)
}

func PieSeries_SetDepth(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("PieSeries_SetDepth").Call(obj, uintptr(value))
}

func PieSeries_GetDragOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PieSeries_GetDragOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_SetDragOrigin(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("PieSeries_SetDragOrigin").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PieSeries_GetShadow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetShadow").Call(obj)
	return ret
}

func PieSeries_SetShadow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PieSeries_SetShadow").Call(obj, value)
}

func PieSeries_GetTransparency(obj uintptr) TChartTransparency {
	ret, _, _ := getLazyProc("PieSeries_GetTransparency").Call(obj)
	return TChartTransparency(ret)
}

func PieSeries_SetTransparency(obj uintptr, value TChartTransparency) {
	_, _, _ = getLazyProc("PieSeries_SetTransparency").Call(obj, uintptr(value))
}

func PieSeries_GetZPosition(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("PieSeries_GetZPosition").Call(obj)
	return TChartDistance(ret)
}

func PieSeries_SetZPosition(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("PieSeries_SetZPosition").Call(obj, uintptr(value))
}

func PieSeries_AddPie(obj uintptr, AValue float64, AText string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_AddPie").Call(obj, uintptr(unsafe.Pointer(&AValue)), GoStrToDStr(AText), uintptr(AColor))
	return int32(ret)
}

func PieSeries_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("PieSeries_Assign").Call(obj, ASource)
}

func PieSeries_CalcBorderPoint(obj uintptr, ASlice uintptr, ARadius, AAngle float64) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PieSeries_CalcBorderPoint").Call(obj, ASlice, uintptr(unsafe.Pointer(&ARadius)), uintptr(unsafe.Pointer(&AAngle)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_CalcInnerRadius(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PieSeries_CalcInnerRadius").Call(obj)
	return int32(ret)
}

func PieSeries_FindContainingSlice(obj uintptr, APoint TPoint) int32 {
	ret, _, _ := getLazyProc("PieSeries_FindContainingSlice").Call(obj, uintptr(unsafe.Pointer(&APoint)))
	return int32(ret)
}

func PieSeries_SliceColor(obj uintptr, AIndex int32) TColor {
	ret, _, _ := getLazyProc("PieSeries_SliceColor").Call(obj, uintptr(AIndex))
	return TColor(ret)
}

func PieSeries_GetColor(obj uintptr, AIndex int32) TColor {
	ret, _, _ := getLazyProc("PieSeries_GetColor").Call(obj, uintptr(AIndex))
	return TColor(ret)
}

func PieSeries_GetXImgValue(obj uintptr, AIndex int32) int32 {
	ret, _, _ := getLazyProc("PieSeries_GetXImgValue").Call(obj, uintptr(AIndex))
	return int32(ret)
}

func PieSeries_GetXMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetXMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetXMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetXMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetXValue(obj uintptr, AIndex int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetXValue").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetXValues(obj uintptr, AIndex, AXIndex int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetXValues").Call(obj, uintptr(AIndex), uintptr(AXIndex), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetYImgValue(obj uintptr, AIndex int32) int32 {
	ret, _, _ := getLazyProc("PieSeries_GetYImgValue").Call(obj, uintptr(AIndex))
	return int32(ret)
}

func PieSeries_GetYMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetYMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetYMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetYMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetYValue(obj uintptr, AIndex int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetYValue").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetYValues(obj uintptr, AIndex, AYIndex int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GetYValues").Call(obj, uintptr(AIndex), uintptr(AYIndex), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_SetColor(obj uintptr, AIndex int32, AColor TColor) {
	_, _, _ = getLazyProc("PieSeries_SetColor").Call(obj, uintptr(AIndex), uintptr(AColor))
}

func PieSeries_SetText(obj uintptr, AIndex int32, AValue string) {
	_, _, _ = getLazyProc("PieSeries_SetText").Call(obj, uintptr(AIndex), GoStrToDStr(AValue))
}

func PieSeries_SetXValue(obj uintptr, AIndex int32, AValue float64) {
	_, _, _ = getLazyProc("PieSeries_SetXValue").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&AValue)))
}

func PieSeries_SetXValues(obj uintptr, AIndex, AXIndex int32, AValue float64) {
	_, _, _ = getLazyProc("PieSeries_SetXValues").Call(obj, uintptr(AIndex), uintptr(AXIndex), uintptr(unsafe.Pointer(&AValue)))
}

func PieSeries_SetYValue(obj uintptr, AIndex int32, AValue float64) {
	_, _, _ = getLazyProc("PieSeries_SetYValue").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&AValue)))
}

func PieSeries_SetYValues(obj uintptr, AIndex, AYIndex int32, AValue float64) {
	_, _, _ = getLazyProc("PieSeries_SetYValues").Call(obj, uintptr(AIndex), uintptr(AYIndex), uintptr(unsafe.Pointer(&AValue)))
}

func PieSeries_Add(obj uintptr, AValue float64, AXLabel string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_Add").Call(obj, uintptr(AValue), GoStrToDStr(AXLabel), uintptr(AColor))
	return int32(ret)
}

func PieSeries_AddNull(obj uintptr, ALabel string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_AddNull").Call(obj, GoStrToDStr(ALabel), uintptr(AColor))
	return int32(ret)
}

func PieSeries_AddX(obj uintptr, AX float64, ALabel string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_AddX").Call(obj, uintptr(unsafe.Pointer(&AX)), GoStrToDStr(ALabel), uintptr(AColor))
	return int32(ret)
}

func PieSeries_AddXY(obj uintptr, AX, AY float64, AXLabel string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_AddXY").Call(obj, uintptr(unsafe.Pointer(&AX)), uintptr(unsafe.Pointer(&AY)), GoStrToDStr(AXLabel), uintptr(AColor))
	return int32(ret)
}

func PieSeries_AddY(obj uintptr, AY float64, ALabel string, AColor TColor) int32 {
	ret, _, _ := getLazyProc("PieSeries_AddY").Call(obj, uintptr(unsafe.Pointer(&AY)), GoStrToDStr(ALabel), uintptr(AColor))
	return int32(ret)
}

func PieSeries_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("PieSeries_BeginUpdate").Call(obj)
}

func PieSeries_Clear(obj uintptr) {
	_, _, _ = getLazyProc("PieSeries_Clear").Call(obj)
}

func PieSeries_Count(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PieSeries_Count").Call(obj)
	return int32(ret)
}

func PieSeries_Delete(obj uintptr, AIndex int32) {
	_, _, _ = getLazyProc("PieSeries_Delete").Call(obj, uintptr(AIndex))
}

func PieSeries_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("PieSeries_EndUpdate").Call(obj)
}

func PieSeries_Extent(obj uintptr) TDoubleRect {
	var ret TDoubleRect
	_, _, _ = getLazyProc("PieSeries_Extent").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_FindYRange(obj uintptr, AXMin, AXMax float64, AYMin, AYMax float64) {
	_, _, _ = getLazyProc("PieSeries_FindYRange").Call(obj, uintptr(unsafe.Pointer(&AXMin)), uintptr(unsafe.Pointer(&AXMax)), uintptr(unsafe.Pointer(&AYMin)), uintptr(unsafe.Pointer(&AYMax)))
}

func PieSeries_FormattedMark(obj uintptr, AIndex int32, AFormat string, AYIndex int32) string {
	ret, _, _ := getLazyProc("PieSeries_FormattedMark").Call(obj, uintptr(AIndex), GoStrToDStr(AFormat), uintptr(AYIndex))
	return DStrToGoStr(ret)
}

func PieSeries_AxisToGraph(obj uintptr, APoint TDoublePoint) TDoublePoint {
	var ret TDoublePoint
	_, _, _ = getLazyProc("PieSeries_AxisToGraph").Call(obj, uintptr(unsafe.Pointer(&APoint)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_AxisToGraphX(obj uintptr, AX float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_AxisToGraphX").Call(obj, uintptr(unsafe.Pointer(&AX)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_AxisToGraphY(obj uintptr, AY float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_AxisToGraphY").Call(obj, uintptr(unsafe.Pointer(&AY)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GetAxisX(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetAxisX").Call(obj)
	return ret
}

func PieSeries_GetAxisY(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PieSeries_GetAxisY").Call(obj)
	return ret
}

func PieSeries_GetDepthColor(obj uintptr, AColor int32, Opposite bool) int32 {
	ret, _, _ := getLazyProc("PieSeries_GetDepthColor").Call(obj, uintptr(AColor), GoBoolToDBool(Opposite))
	return int32(ret)
}

func PieSeries_GraphToAxis(obj uintptr, APoint TDoublePoint) TDoublePoint {
	var ret TDoublePoint
	_, _, _ = getLazyProc("PieSeries_GraphToAxis").Call(obj, uintptr(unsafe.Pointer(&APoint)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GraphToAxisX(obj uintptr, AX float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GraphToAxisX").Call(obj, uintptr(unsafe.Pointer(&AX)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_GraphToAxisY(obj uintptr, AY float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("PieSeries_GraphToAxisY").Call(obj, uintptr(unsafe.Pointer(&AY)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_IsRotated(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_IsRotated").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_Draw(obj uintptr, ADrawer uintptr) {
	_, _, _ = getLazyProc("PieSeries_Draw").Call(obj, ADrawer)
}

func PieSeries_GetAxisBounds(obj uintptr, AAxis uintptr, AMin, AMax float64) bool {
	ret, _, _ := getLazyProc("PieSeries_GetAxisBounds").Call(obj, AAxis, uintptr(unsafe.Pointer(&AMin)), uintptr(unsafe.Pointer(&AMax)))
	return DBoolToGoBool(ret)
}

func PieSeries_GetGraphBounds(obj uintptr) TDoubleRect {
	var ret TDoubleRect
	_, _, _ = getLazyProc("PieSeries_GetGraphBounds").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PieSeries_IsEmpty(obj uintptr) bool {
	ret, _, _ := getLazyProc("PieSeries_IsEmpty").Call(obj)
	return DBoolToGoBool(ret)
}

func PieSeries_MovePoint(obj uintptr, AIndex int32, ANewPos TPoint) {
	_, _, _ = getLazyProc("PieSeries_MovePoint").Call(obj, uintptr(AIndex), uintptr(unsafe.Pointer(&ANewPos)))
}

func PieSeries_MovePointEx(obj uintptr, AIndex int32, AXIndex, AYIndex int32, ANewPos TDoublePoint) {
	_, _, _ = getLazyProc("PieSeries_MovePointEx").Call(obj, uintptr(AIndex), uintptr(AXIndex), uintptr(AYIndex), uintptr(unsafe.Pointer(&ANewPos)))
}

func PieSeries_UpdateBiDiMode(obj uintptr) {
	_, _, _ = getLazyProc("PieSeries_UpdateBiDiMode").Call(obj)
}
