package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

func ChartTextElement_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTextElement_Create").Call(obj)
	return ret
}

func ChartTextElement_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartTextElement_Free").Call(obj)
}

func ChartTextElement_StaticClassType() TClass {
	ret, _, _ := getLazyProc("CharTextElement_StaticClassType").Call()
	return TClass(ret)
}

func ChartTextElement_GetCalloutAngle(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartTextElement_GetCalloutAngle").Call(obj)
	return Cardinal(ret)
}

func ChartTextElement_SetCalloutAngle(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartTextElement_SetCalloutAngle").Call(obj, uintptr(value))
}

func ChartTextElement_GetClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartTextElement_GetClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartTextElement_SetClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartTextElement_SetClipped").Call(obj, GoBoolToDBool(value))
}

func ChartTextElement_GetOverlapPolicy(obj uintptr) TChartMarksOverlapPolicy {
	ret, _, _ := getLazyProc("ChartTextElement_GetOverlapPolicy").Call(obj)
	return TChartMarksOverlapPolicy(ret)
}

func ChartTextElement_SetOverlapPolicy(obj uintptr, value TChartMarksOverlapPolicy) {
	_, _, _ = getLazyProc("ChartTextElement_SetOverlapPolicy").Call(obj, uintptr(value))
}

func ChartTextElement_SetOnGetShape(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartTextElement_SetOnGetShape").Call(obj, addEventToMap(obj, fn))
}

func ChartTextElement_GetShape(obj uintptr) TChartLabelShape {
	ret, _, _ := getLazyProc("ChartTextElement_GetShape").Call(obj)
	return TChartLabelShape(ret)
}

func ChartTextElement_SetShape(obj uintptr, value TChartLabelShape) {
	_, _, _ = getLazyProc("ChartTextElement_SetShape").Call(obj, uintptr(value))
}

func ChartTextElement_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartTextElement_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartTextElement_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartTextElement_SetTextFormat").Call(obj, uintptr(value))
}

func ChartTextElement_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ChartTextElement_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ChartTextElement_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("ChartTextElement_SetAlignment").Call(obj, uintptr(value))
}

func ChartTextElement_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTextElement_GetMargins").Call(obj)
	return ret
}

func ChartTextElement_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartTextElement_SetMargins").Call(obj, value)
}

func ChartTextElement_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartTextElement_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartTextElement_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartTextElement_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartTextElement_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartTextElement_Assign").Call(obj, ASource)
}

func ChartTextElement_DrawLabel(obj uintptr, ADrawer uintptr, ADataPoint, ALabelCenter TPoint, AText string, APrevLabelPoly uintptr) {
	_, _, _ = getLazyProc("ChartTextElement_DrawLabel").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ADataPoint)), uintptr(unsafe.Pointer(&ALabelCenter)), GoStrToDStr(AText), APrevLabelPoly)
}

func ChartTextElement_GetLabelPolygon(obj uintptr, ADrawer uintptr, ASize TPoint) uintptr {
	ret, _, _ := getLazyProc("ChartTextElement_GetLabelPolygon").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ASize)))
	return ret
}

func ChartTextElement_GetTextRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ChartTextElement_GetTextRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTextElement_MeasureLabel(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartTextElement_MeasureLabel").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTextElement_MeasureLabelHeight(obj uintptr, ADrawer uintptr, AText string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("ChartTextElement_MeasureLabelHeight").Call(obj, ADrawer, GoStrToDStr(AText), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ChartTextElement_SetInsideDir(obj uintptr, dx, dy float64) {
	_, _, _ = getLazyProc("ChartTextElement_SetInsideDir").Call(obj, uintptr(unsafe.Pointer(&dx)), uintptr(unsafe.Pointer(&dy)))
}

func ChartTextElement_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartTextElement_GetOwner").Call(obj)
	return ret
}

func ChartTextElement_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartTextElement_SetOwner").Call(obj, AOwner)
}
