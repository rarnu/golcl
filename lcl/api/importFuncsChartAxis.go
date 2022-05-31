package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartAxis_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_Create").Call(obj)
	return ret
}

func ChartAxis_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxis_Free").Call(obj)
}

func ChartAxis_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartAxis_StaticClassType").Call()
	return TClass(ret)
}

func ChartAxis_GetAtDataOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_GetAtDataOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_SetAtDataOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxis_SetAtDataOnly").Call(obj, GoBoolToDBool(value))
}

func ChartAxis_GetAxisPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetAxisPen").Call(obj)
	return ret
}

func ChartAxis_SetAxisPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetAxisPen").Call(obj, value)
}

func ChartAxis_GetGroup(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_GetGroup").Call(obj)
	return int32(ret)
}

func ChartAxis_SetGroup(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxis_SetGroup").Call(obj, uintptr(value))
}

func ChartAxis_GetInverted(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_GetInverted").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_SetInverted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxis_SetInverted").Call(obj, GoBoolToDBool(value))
}

func ChartAxis_GetLabelSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_GetLabelSize").Call(obj)
	return int32(ret)
}

func ChartAxis_SetLabelSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxis_SetLabelSize").Call(obj, uintptr(value))
}

func ChartAxis_GetMargin(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartAxis_GetMargin").Call(obj)
	return TChartDistance(ret)
}

func ChartAxis_SetMargin(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartAxis_SetMargin").Call(obj, uintptr(value))
}

func ChartAxis_GetMarginsForMarks(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_GetMarginsForMarks").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_SetMarginsForMarks(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxis_SetMarginsForMarks").Call(obj, GoBoolToDBool(value))
}

func ChartAxis_GetMinors(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetMinors").Call(obj)
	return ret
}

func ChartAxis_SetMinors(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetMinors").Call(obj, value)
}

func ChartAxis_GetPosition(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("ChartAxis_GetPosition").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxis_SetPosition(obj uintptr, value float64) {
	_, _, _ = getLazyProc("ChartAxis_SetPosition").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ChartAxis_GetPositionUnits(obj uintptr) TChartUnits {
	ret, _, _ := getLazyProc("ChartAxis_GetPositionUnits").Call(obj)
	return TChartUnits(ret)
}

func ChartAxis_SetPositionUnits(obj uintptr, value TChartUnits) {
	_, _, _ = getLazyProc("ChartAxis_SetPositionUnits").Call(obj, uintptr(value))
}

func ChartAxis_GetRange(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetRange").Call(obj)
	return ret
}

func ChartAxis_SetRange(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetRange").Call(obj, value)
}

func ChartAxis_GetTitle(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetTitle").Call(obj)
	return ret
}

func ChartAxis_SetTitle(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetTitle").Call(obj, value)
}

func ChartAxis_GetTransformations(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetTransformations").Call(obj)
	return ret
}

func ChartAxis_SetTransformations(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetTransformations").Call(obj, value)
}

func ChartAxis_GetZPosition(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartAxis_GetZPosition").Call(obj)
	return TChartDistance(ret)
}

func ChartAxis_SetZPosition(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartAxis_SetZPosition").Call(obj, uintptr(value))
}

func ChartAxis_SetOnGetMarkText(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartAxis_SetOnGetMarkText").Call(obj, addEventToMap(obj, fn))
}

func ChartAxis_GetAlignment(obj uintptr) TChartAxisAlignment {
	ret, _, _ := getLazyProc("ChartAxis_GetAlignment").Call(obj)
	return TChartAxisAlignment(ret)
}

func ChartAxis_SetAlignment(obj uintptr, value TChartAxisAlignment) {
	_, _, _ = getLazyProc("ChartAxis_SetAlignment").Call(obj, uintptr(value))
}

func ChartAxis_GetArrow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetArrow").Call(obj)
	return ret
}

func ChartAxis_SetArrow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetArrow").Call(obj, value)
}

func ChartAxis_GetMarks(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetMarks").Call(obj)
	return ret
}

func ChartAxis_SetMarks(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetMarks").Call(obj, value)
}

func ChartAxis_GetGrid(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetGrid").Call(obj)
	return ret
}

func ChartAxis_SetGrid(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetGrid").Call(obj, value)
}

func ChartAxis_GetIntervals(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetIntervals").Call(obj)
	return ret
}

func ChartAxis_SetIntervals(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetIntervals").Call(obj, value)
}

func ChartAxis_GetTickColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ChartAxis_GetTickColor").Call(obj)
	return TColor(ret)
}

func ChartAxis_SetTickColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ChartAxis_SetTickColor").Call(obj, uintptr(value))
}

func ChartAxis_GetTickInnerLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_GetTickInnerLength").Call(obj)
	return int32(ret)
}

func ChartAxis_SetTickInnerLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxis_SetTickInnerLength").Call(obj, uintptr(value))
}

func ChartAxis_GetTickLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_GetTickLength").Call(obj)
	return int32(ret)
}

func ChartAxis_SetTickLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxis_SetTickLength").Call(obj, uintptr(value))
}

func ChartAxis_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxis_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartAxis_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetCollection").Call(obj)
	return ret
}

func ChartAxis_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxis_SetCollection").Call(obj, value)
}

func ChartAxis_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_GetIndex").Call(obj)
	return int32(ret)
}

func ChartAxis_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxis_SetIndex").Call(obj, uintptr(value))
}

func ChartAxis_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxis_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxis_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartAxis_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func ChartAxis_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartAxis_Assign").Call(obj, ASource)
}

func ChartAxis_GetHitTestInfoAt(obj uintptr, APoint TPoint, ADelta int32) TChartAxisHitTests {
	ret, _, _ := getLazyProc("ChartAxis_GetHitTestInfoAt").Call(obj, uintptr(unsafe.Pointer(&APoint)), uintptr(ADelta))
	return TChartAxisHitTests(ret)
}

func ChartAxis_Draw(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxis_Draw").Call(obj)
}

func ChartAxis_DrawTitle(obj uintptr, ASize int32) {
	_, _, _ = getLazyProc("ChartAxis_DrawTitle").Call(obj, uintptr(ASize))
}

func ChartAxis_GetChart(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetChart").Call(obj)
	return ret
}

func ChartAxis_GetTransform(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxis_GetTransform").Call(obj)
	return ret
}

func ChartAxis_IsDefaultPosition(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_IsDefaultPosition").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_IsPointInside(obj uintptr, APoint TPoint) bool {
	ret, _, _ := getLazyProc("ChartAxis_IsPointInside").Call(obj, uintptr(unsafe.Pointer(&APoint)))
	return DBoolToGoBool(ret)
}

func ChartAxis_IsVertical(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_IsVertical").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_Measure(obj uintptr, AExtent TDoubleRect, AClipRect TRect, AMeasureData TChartAxisGroup) {
	_, _, _ = getLazyProc("ChartAxis_Measure").Call(obj, uintptr(unsafe.Pointer(&AExtent)), uintptr(unsafe.Pointer(&AClipRect)), uintptr(unsafe.Pointer(&AMeasureData)))
}

func ChartAxis_MeasureLabelSize(obj uintptr, ADrawer uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_MeasureLabelSize").Call(obj, ADrawer)
	return int32(ret)
}

func ChartAxis_MeasureTitleSize(obj uintptr, ADrawer uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxis_MeasureTitleSize").Call(obj, ADrawer)
	return int32(ret)
}

func ChartAxis_PositionToCoord(obj uintptr, ARect TRect) int32 {
	ret, _, _ := getLazyProc("ChartAxis_PositionToCoord").Call(obj, uintptr(unsafe.Pointer(&ARect)))
	return int32(ret)
}

func ChartAxis_UpdateBidiMode(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxis_UpdateBidiMode").Call(obj)
}

func ChartAxis_UpdateBounds(obj uintptr, AMin, AMax float64) {
	_, _, _ = getLazyProc("ChartAxis_UpdateBounds").Call(obj, uintptr(unsafe.Pointer(&AMin)), uintptr(unsafe.Pointer(&AMax)))
}

func ChartAxis_IsFlipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxis_IsFlipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxis_TryApplyStripes(obj uintptr, ADrawer uintptr, AIndex Cardinal) bool {
	ret, _, _ := getLazyProc("ChartAxis_TryApplyStripes").Call(obj, ADrawer, uintptr(AIndex))
	return DBoolToGoBool(ret)
}

func ChartAxis_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxis_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}
