package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartMinorAxisMarks_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_Create").Call(obj)
	return ret
}

func ChartMinorAxisMarks_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_Free").Call(obj)
}

func ChartMinorAxisMarks_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_StaticClassType").Call()
	return TClass(ret)
}

func ChartMinorAxisMarks_GetSourceExchangeXY(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetSourceExchangeXY").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetSourceExchangeXY(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetSourceExchangeXY").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxisMarks_GetStripes(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetStripes").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetStripes(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetStripes").Call(obj, value)
}

func ChartMinorAxisMarks_GetArrow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetArrow").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetArrow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetArrow").Call(obj, value)
}

func ChartMinorAxisMarks_GetAutoMargins(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetAutoMargins").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetAutoMargins(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetAutoMargins").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxisMarks_GetDistanceToCenter(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetDistanceToCenter").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetDistanceToCenter(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetDistanceToCenter").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxisMarks_GetFormat(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetFormat").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxisMarks_SetFormat(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetFormat").Call(obj, GoStrToDStr(value))
}

func ChartMinorAxisMarks_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetFrame").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetFrame").Call(obj, value)
}

func ChartMinorAxisMarks_GetLabelBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetLabelBrush").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetLabelBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetLabelBrush").Call(obj, value)
}

func ChartMinorAxisMarks_GetLinkDistance(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetLinkDistance").Call(obj)
	return int32(ret)
}

func ChartMinorAxisMarks_SetLinkDistance(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetLinkDistance").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetLinkPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetLinkPen").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetLinkPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetLinkPen").Call(obj, value)
}

func ChartMinorAxisMarks_GetStyle(obj uintptr) TSeriesMarksStyle {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetStyle").Call(obj)
	return TSeriesMarksStyle(ret)
}

func ChartMinorAxisMarks_SetStyle(obj uintptr, value TSeriesMarksStyle) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetStyle").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetYIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetYIndex").Call(obj)
	return int32(ret)
}

func ChartMinorAxisMarks_SetYIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetYIndex").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetAttachment(obj uintptr) TChartMarkAttachment {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetAttachment").Call(obj)
	return TChartMarkAttachment(ret)
}

func ChartMinorAxisMarks_SetAttachment(obj uintptr, value TChartMarkAttachment) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetAttachment").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetDistance(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetDistance").Call(obj)
	return TChartDistance(ret)
}

func ChartMinorAxisMarks_SetDistance(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetDistance").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetLabelFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetLabelFont").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetLabelFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetLabelFont").Call(obj, value)
}

func ChartMinorAxisMarks_GetCalloutAngle(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetCalloutAngle").Call(obj)
	return Cardinal(ret)
}

func ChartMinorAxisMarks_SetCalloutAngle(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetCalloutAngle").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetClipped").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxisMarks_GetOverlapPolicy(obj uintptr) TChartMarksOverlapPolicy {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetOverlapPolicy").Call(obj)
	return TChartMarksOverlapPolicy(ret)
}

func ChartMinorAxisMarks_SetOverlapPolicy(obj uintptr, value TChartMarksOverlapPolicy) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetOverlapPolicy").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_SetOnGetShape(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetOnGetShape").Call(obj, addEventToMap(obj, fn))
}

func ChartMinorAxisMarks_GetShape(obj uintptr) TChartLabelShape {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetShape").Call(obj)
	return TChartLabelShape(ret)
}

func ChartMinorAxisMarks_SetShape(obj uintptr, value TChartLabelShape) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetShape").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartMinorAxisMarks_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetTextFormat").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ChartMinorAxisMarks_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetAlignment").Call(obj, uintptr(value))
}

func ChartMinorAxisMarks_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetMargins").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetMargins").Call(obj, value)
}

func ChartMinorAxisMarks_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxisMarks_Measure(obj uintptr, ADrawer uintptr, AIsVertical bool, ATickLength int32, AValues uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_Measure").Call(obj, ADrawer, GoBoolToDBool(AIsVertical), uintptr(ATickLength), AValues)
	return int32(ret)
}

func ChartMinorAxisMarks_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_Assign").Call(obj, ASource)
}

func ChartMinorAxisMarks_CenterHeightOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMinorAxisMarks_CenterHeightOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMinorAxisMarks_CenterOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMinorAxisMarks_CenterOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMinorAxisMarks_IsMarkLabelsVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_IsMarkLabelsVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxisMarks_SetAdditionalAngle(obj uintptr, AAngle float64) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetAdditionalAngle").Call(obj, uintptr(unsafe.Pointer(&AAngle)))
}

func ChartMinorAxisMarks_DrawLabel(obj uintptr, ADrawer uintptr, ADataPoint, ALabelCenter TPoint, AText string, APrevLabelPoly uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_DrawLabel").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ADataPoint)), uintptr(unsafe.Pointer(&ALabelCenter)), GoStrToDStr(AText), APrevLabelPoly)
}

func ChartMinorAxisMarks_GetLabelPolygon(obj uintptr, ADrawer uintptr, ASize TPoint) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetLabelPolygon").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ASize)))
	return ret
}

func ChartMinorAxisMarks_GetTextRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartMinorAxisMarks_GetTextRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMinorAxisMarks_MeasureLabel(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMinorAxisMarks_MeasureLabel").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMinorAxisMarks_MeasureLabelHeight(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMinorAxisMarks_MeasureLabelHeight").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMinorAxisMarks_SetInsideDir(obj uintptr, dx, dy float64) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetInsideDir").Call(obj, uintptr(unsafe.Pointer(&dx)), uintptr(unsafe.Pointer(&dy)))
}

func ChartMinorAxisMarks_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxisMarks_GetOwner").Call(obj)
	return ret
}

func ChartMinorAxisMarks_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxisMarks_SetOwner").Call(obj, AOwner)
}
