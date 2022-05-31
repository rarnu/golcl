package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartTitle_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_Create").Call(obj)
	return ret
}

func ChartTitle_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartTitle_Free").Call(obj)
}

func ChartTitle_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartTitle_StaticClassType").Call()
	return TClass(ret)
}

func ChartTitle_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetBrush").Call(obj)
	return ret
}

func ChartTitle_SetBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetBrush").Call(obj, value)
}

func ChartTitle_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetFont").Call(obj)
	return ret
}

func ChartTitle_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetFont").Call(obj, value)
}

func ChartTitle_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetFrame").Call(obj)
	return ret
}

func ChartTitle_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetFrame").Call(obj, value)
}

func ChartTitle_GetMargin(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartTitle_GetMargin").Call(obj)
	return TChartDistance(ret)
}

func ChartTitle_SetMargin(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartTitle_SetMargin").Call(obj, uintptr(value))
}

func ChartTitle_GetText(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetText").Call(obj)
	return ret
}

func ChartTitle_SetText(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetText").Call(obj, value)
}

func ChartTitle_GetWordwrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartTitle_GetWordwrap").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartTitle_SetWordwrap(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartTitle_SetWordwrap").Call(obj, GoBoolToDBool(value))
}

func ChartTitle_GetCalloutAngle(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartTitle_GetCalloutAngle").Call(obj)
	return Cardinal(ret)
}

func ChartTitle_SetCalloutAngle(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartTitle_SetCalloutAngle").Call(obj, uintptr(value))
}

func ChartTitle_GetClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartTitle_GetClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartTitle_SetClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartTitle_SetClipped").Call(obj, GoBoolToDBool(value))
}

func ChartTitle_GetOverlapPolicy(obj uintptr) TChartMarksOverlapPolicy {
	ret, _, _ := getLazyProc("ChartTitle_GetOverlapPolicy").Call(obj)
	return TChartMarksOverlapPolicy(ret)
}

func ChartTitle_SetOverlapPolicy(obj uintptr, value TChartMarksOverlapPolicy) {
	_, _, _ = getLazyProc("ChartTitle_SetOverlapPolicy").Call(obj, uintptr(value))
}

func ChartTitle_SetOnGetShape(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartTitle_SetOnGetShape").Call(obj, addEventToMap(obj, fn))
}

func ChartTitle_GetShape(obj uintptr) TChartLabelShape {
	ret, _, _ := getLazyProc("ChartTitle_GetShape").Call(obj)
	return TChartLabelShape(ret)
}

func ChartTitle_SetShape(obj uintptr, value TChartLabelShape) {
	_, _, _ = getLazyProc("ChartTitle_SetShape").Call(obj, uintptr(value))
}

func ChartTitle_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartTitle_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartTitle_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartTitle_SetTextFormat").Call(obj, uintptr(value))
}

func ChartTitle_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ChartTitle_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ChartTitle_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("ChartTitle_SetAlignment").Call(obj, uintptr(value))
}

func ChartTitle_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetMargins").Call(obj)
	return ret
}

func ChartTitle_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetMargins").Call(obj, value)
}

func ChartTitle_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartTitle_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartTitle_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartTitle_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartTitle_Draw(obj uintptr, ADrawer uintptr) {
	_, _, _ = getLazyProc("ChartTitle_Draw").Call(obj, ADrawer)
}

func ChartTitle_IsPointInBounds(obj uintptr, APoint TPoint) bool {
	ret, _, _ := getLazyProc("ChartTitle_IsPointInBounds").Call(obj, uintptr(unsafe.Pointer(&APoint)))
	return DBoolToGoBool(ret)
}

func ChartTitle_Measure(obj uintptr, ADrawer uintptr, ADir, ALeft, ARight int32, AY int32) {
	_, _, _ = getLazyProc("ChartTitle_Measure").Call(obj, ADrawer, uintptr(ADir), uintptr(ALeft), uintptr(ARight), uintptr(AY))
}

func ChartTitle_UpdateBidiMode(obj uintptr) {
	_, _, _ = getLazyProc("ChartTitle_UpdateBidiMode").Call(obj)
}

func ChartTitle_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartTitle_Assign").Call(obj, ASource)
}

func ChartTitle_DrawLabel(obj uintptr, ADrawer uintptr, ADataPoint, ALabelCenter TPoint, AText string, APrevLabelPoly uintptr) {
	_, _, _ = getLazyProc("ChartTitle_DrawLabel").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ADataPoint)), uintptr(unsafe.Pointer(&ALabelCenter)), GoStrToDStr(AText), APrevLabelPoly)
}

func ChartTitle_GetLabelPolygon(obj uintptr, ADrawer uintptr, ASize TPoint) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetLabelPolygon").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ASize)))
	return ret
}

func ChartTitle_GetTextRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartTitle_GetTextRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTitle_MeasureLabel(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartTitle_MeasureLabel").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTitle_MeasureLabelHeight(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartTitle_MeasureLabelHeight").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTitle_SetInsideDir(obj uintptr, dx, dy float64) {
	_, _, _ = getLazyProc("ChartTitle_SetInsideDir").Call(obj, uintptr(unsafe.Pointer(&dx)), uintptr(unsafe.Pointer(&dy)))
}

func ChartTitle_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTitle_GetOwner").Call(obj)
	return ret
}

func ChartTitle_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartTitle_SetOwner").Call(obj, AOwner)
}
