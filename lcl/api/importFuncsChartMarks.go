package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartMarks_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_Create").Call(obj)
	return ret
}

func ChartMarks_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartMarks_Free").Call(obj)
}

func ChartMarks_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartMarks_StaticClassType").Call()
	return TClass(ret)
}

func ChartMarks_GetArrow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetArrow").Call(obj)
	return ret
}

func ChartMarks_SetArrow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetArrow").Call(obj, value)
}

func ChartMarks_GetAutoMargins(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMarks_GetAutoMargins").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMarks_SetAutoMargins(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMarks_SetAutoMargins").Call(obj, GoBoolToDBool(value))
}

func ChartMarks_GetDistanceToCenter(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMarks_GetDistanceToCenter").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMarks_SetDistanceToCenter(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMarks_SetDistanceToCenter").Call(obj, GoBoolToDBool(value))
}

func ChartMarks_GetFormat(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMarks_GetFormat").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMarks_SetFormat(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartMarks_SetFormat").Call(obj, GoStrToDStr(value))
}

func ChartMarks_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetFrame").Call(obj)
	return ret
}

func ChartMarks_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetFrame").Call(obj, value)
}

func ChartMarks_GetLabelBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetLabelBrush").Call(obj)
	return ret
}

func ChartMarks_SetLabelBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetLabelBrush").Call(obj, value)
}

func ChartMarks_GetLinkDistance(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMarks_GetLinkDistance").Call(obj)
	return int32(ret)
}

func ChartMarks_SetLinkDistance(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMarks_SetLinkDistance").Call(obj, uintptr(value))
}

func ChartMarks_GetLinkPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetLinkPen").Call(obj)
	return ret
}

func ChartMarks_SetLinkPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetLinkPen").Call(obj, value)
}

func ChartMarks_GetStyle(obj uintptr) TSeriesMarksStyle {
	ret, _, _ := getLazyProc("ChartMarks_GetStyle").Call(obj)
	return TSeriesMarksStyle(ret)
}

func ChartMarks_SetStyle(obj uintptr, value TSeriesMarksStyle) {
	_, _, _ = getLazyProc("ChartMarks_SetStyle").Call(obj, uintptr(value))
}

func ChartMarks_GetYIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMarks_GetYIndex").Call(obj)
	return int32(ret)
}

func ChartMarks_SetYIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMarks_SetYIndex").Call(obj, uintptr(value))
}

func ChartMarks_GetAttachment(obj uintptr) TChartMarkAttachment {
	ret, _, _ := getLazyProc("ChartMarks_GetAttachment").Call(obj)
	return TChartMarkAttachment(ret)
}

func ChartMarks_SetAttachment(obj uintptr, value TChartMarkAttachment) {
	_, _, _ = getLazyProc("ChartMarks_SetAttachment").Call(obj, uintptr(value))
}

func ChartMarks_GetDistance(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartMarks_GetDistance").Call(obj)
	return TChartDistance(ret)
}

func ChartMarks_SetDistance(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartMarks_SetDistance").Call(obj, uintptr(value))
}

func ChartMarks_GetLabelFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetLabelFont").Call(obj)
	return ret
}

func ChartMarks_SetLabelFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetLabelFont").Call(obj, value)
}

func ChartMarks_GetCalloutAngle(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartMarks_GetCalloutAngle").Call(obj)
	return Cardinal(ret)
}

func ChartMarks_SetCalloutAngle(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartMarks_SetCalloutAngle").Call(obj, uintptr(value))
}

func ChartMarks_GetClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMarks_GetClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMarks_SetClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMarks_SetClipped").Call(obj, GoBoolToDBool(value))
}

func ChartMarks_GetOverlapPolicy(obj uintptr) TChartMarksOverlapPolicy {
	ret, _, _ := getLazyProc("ChartMarks_GetOverlapPolicy").Call(obj)
	return TChartMarksOverlapPolicy(ret)
}

func ChartMarks_SetOverlapPolicy(obj uintptr, value TChartMarksOverlapPolicy) {
	_, _, _ = getLazyProc("ChartMarks_SetOverlapPolicy").Call(obj, uintptr(value))
}

func ChartMarks_SetOnGetShape(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartMarks_SetOnGetShape").Call(obj, addEventToMap(obj, fn))
}

func ChartMarks_GetShape(obj uintptr) TChartLabelShape {
	ret, _, _ := getLazyProc("ChartMarks_GetShape").Call(obj)
	return TChartLabelShape(ret)
}

func ChartMarks_SetShape(obj uintptr, value TChartLabelShape) {
	_, _, _ = getLazyProc("ChartMarks_SetShape").Call(obj, uintptr(value))
}

func ChartMarks_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartMarks_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartMarks_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartMarks_SetTextFormat").Call(obj, uintptr(value))
}

func ChartMarks_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ChartMarks_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ChartMarks_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("ChartMarks_SetAlignment").Call(obj, uintptr(value))
}

func ChartMarks_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetMargins").Call(obj)
	return ret
}

func ChartMarks_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetMargins").Call(obj, value)
}

func ChartMarks_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMarks_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMarks_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMarks_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartMarks_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartMarks_Assign").Call(obj, ASource)
}

func ChartMarks_CenterHeightOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMarks_CenterHeightOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMarks_CenterOffset(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMarks_CenterOffset").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMarks_IsMarkLabelsVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMarks_IsMarkLabelsVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMarks_SetAdditionalAngle(obj uintptr, AAngle float64) {
	_, _, _ = getLazyProc("ChartMarks_SetAdditionalAngle").Call(obj, uintptr(unsafe.Pointer(&AAngle)))
}

func ChartMarks_DrawLabel(obj uintptr, ADrawer uintptr, ADataPoint, ALabelCenter TPoint, AText string, APrevLabelPoly uintptr) {
	_, _, _ = getLazyProc("ChartMarks_DrawLabel").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ADataPoint)), uintptr(unsafe.Pointer(&ALabelCenter)), GoStrToDStr(AText), APrevLabelPoly)
}

func ChartMarks_GetLabelPolygon(obj uintptr, ADrawer uintptr, ASize TPoint) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetLabelPolygon").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ASize)))
	return ret
}

func ChartMarks_GetTextRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartMarks_GetTextRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMarks_MeasureLabel(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMarks_MeasureLabel").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMarks_MeasureLabelHeight(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartMarks_MeasureLabelHeight").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartMarks_SetInsideDir(obj uintptr, dx, dy float64) {
	_, _, _ = getLazyProc("ChartMarks_SetInsideDir").Call(obj, uintptr(unsafe.Pointer(&dx)), uintptr(unsafe.Pointer(&dy)))
}

func ChartMarks_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMarks_GetOwner").Call(obj)
	return ret
}

func ChartMarks_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartMarks_SetOwner").Call(obj, AOwner)
}
