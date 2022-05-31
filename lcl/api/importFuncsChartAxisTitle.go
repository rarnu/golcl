package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartAxisTitle_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_Create").Call(obj)
	return ret
}

func ChartAxisTitle_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_Free").Call(obj)
}

func ChartAxisTitle_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartAxisTitle_StaticClassType").Call()
	return TClass(ret)
}

func ChartAxisTitle_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisTitle_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetCaption").Call(obj, GoStrToDStr(value))
}

func ChartAxisTitle_GetPositionOnMarks(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetPositionOnMarks").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetPositionOnMarks(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetPositionOnMarks").Call(obj, GoBoolToDBool(value))
}

func ChartAxisTitle_GetWordwrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetWordwrap").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetWordwrap(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetWordwrap").Call(obj, GoBoolToDBool(value))
}

func ChartAxisTitle_GetArrow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetArrow").Call(obj)
	return ret
}

func ChartAxisTitle_SetArrow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetArrow").Call(obj, value)
}

func ChartAxisTitle_GetAutoMargins(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetAutoMargins").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetAutoMargins(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetAutoMargins").Call(obj, GoBoolToDBool(value))
}

func ChartAxisTitle_GetDistanceToCenter(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetDistanceToCenter").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetDistanceToCenter(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetDistanceToCenter").Call(obj, GoBoolToDBool(value))
}

func ChartAxisTitle_GetFormat(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetFormat").Call(obj)
	return DStrToGoStr(ret)
}

func ChartAxisTitle_SetFormat(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetFormat").Call(obj, GoStrToDStr(value))
}

func ChartAxisTitle_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetFrame").Call(obj)
	return ret
}

func ChartAxisTitle_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetFrame").Call(obj, value)
}

func ChartAxisTitle_GetLabelBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetLabelBrush").Call(obj)
	return ret
}

func ChartAxisTitle_SetLabelBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetLabelBrush").Call(obj, value)
}

func ChartAxisTitle_GetLinkDistance(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetLinkDistance").Call(obj)
	return int32(ret)
}

func ChartAxisTitle_SetLinkDistance(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetLinkDistance").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetLinkPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetLinkPen").Call(obj)
	return ret
}

func ChartAxisTitle_SetLinkPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetLinkPen").Call(obj, value)
}

func ChartAxisTitle_GetStyle(obj uintptr) TSeriesMarksStyle {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetStyle").Call(obj)
	return TSeriesMarksStyle(ret)
}

func ChartAxisTitle_SetStyle(obj uintptr, value TSeriesMarksStyle) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetStyle").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetYIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetYIndex").Call(obj)
	return int32(ret)
}

func ChartAxisTitle_SetYIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetYIndex").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetAttachment(obj uintptr) TChartMarkAttachment {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetAttachment").Call(obj)
	return TChartMarkAttachment(ret)
}

func ChartAxisTitle_SetAttachment(obj uintptr, value TChartMarkAttachment) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetAttachment").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetDistance(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetDistance").Call(obj)
	return TChartDistance(ret)
}

func ChartAxisTitle_SetDistance(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetDistance").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetLabelFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetLabelFont").Call(obj)
	return ret
}

func ChartAxisTitle_SetLabelFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetLabelFont").Call(obj, value)
}

func ChartAxisTitle_GetCalloutAngle(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetCalloutAngle").Call(obj)
	return Cardinal(ret)
}

func ChartAxisTitle_SetCalloutAngle(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetCalloutAngle").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetClipped").Call(obj, GoBoolToDBool(value))
}

func ChartAxisTitle_GetOverlapPolicy(obj uintptr) TChartMarksOverlapPolicy {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetOverlapPolicy").Call(obj)
	return TChartMarksOverlapPolicy(ret)
}

func ChartAxisTitle_SetOverlapPolicy(obj uintptr, value TChartMarksOverlapPolicy) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetOverlapPolicy").Call(obj, uintptr(value))
}

func ChartAxisTitle_SetOnGetShape(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetOnGetShape").Call(obj, addEventToMap(obj, fn))
}

func ChartAxisTitle_GetShape(obj uintptr) TChartLabelShape {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetShape").Call(obj)
	return TChartLabelShape(ret)
}

func ChartAxisTitle_SetShape(obj uintptr, value TChartLabelShape) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetShape").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartAxisTitle_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetTextFormat").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ChartAxisTitle_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetAlignment").Call(obj, uintptr(value))
}

func ChartAxisTitle_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetMargins").Call(obj)
	return ret
}

func ChartAxisTitle_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetMargins").Call(obj, value)
}

func ChartAxisTitle_CenterHeightOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartAxisTitle_CenterHeightOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxisTitle_CenterOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartAxisTitle_CenterOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxisTitle_IsMarkLabelsVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartAxisTitle_IsMarkLabelsVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartAxisTitle_SetAdditionalAngle(obj uintptr, AAngle float64) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetAdditionalAngle").Call(obj, uintptr(unsafe.Pointer(&AAngle)))
}

func ChartAxisTitle_DrawLabel(obj uintptr, ADrawer uintptr, ADataPoint, ALabelCenter TPoint, AText string, APrevLabelPoly uintptr) {
	_, _, _ = getLazyProc("ChartAxisTitle_DrawLabel").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ADataPoint)), uintptr(unsafe.Pointer(&ALabelCenter)), GoStrToDStr(AText), APrevLabelPoly)
}

func ChartAxisTitle_GetLabelPolygon(obj uintptr, ADrawer uintptr, ASize TPoint) uintptr {
	ret, _, _ := getLazyProc("ChartAxisTitle_GetLabelPolygon").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ASize)))
	return ret
}

func ChartAxisTitle_GetTextRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartAxisTitle_GetTextRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxisTitle_MeasureLabel(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartAxisTitle_MeasureLabel").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxisTitle_MeasureLabelHeight(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartAxisTitle_MeasureLabelHeight").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartAxisTitle_SetInsideDir(obj uintptr, dx, dy float64) {
	_, _, _ = getLazyProc("ChartAxisTitle_SetInsideDir").Call(obj, uintptr(unsafe.Pointer(&dx)), uintptr(unsafe.Pointer(&dy)))
}
