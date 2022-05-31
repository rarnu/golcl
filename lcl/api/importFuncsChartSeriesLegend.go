package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartSeriesLegend_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartSeriesLegend_Create").Call(obj)
	return ret
}

func ChartSeriesLegend_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartSeriesLegend_Free").Call(obj)
}

func ChartSeriesLegend_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartSeriesLegend_StaticClassType").Call()
	return TClass(ret)
}

func ChartSeriesLegend_GetFormat(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetFormat").Call(obj)
	return DStrToGoStr(ret)
}

func ChartSeriesLegend_SetFormat(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetFormat").Call(obj, GoStrToDStr(value))
}

func ChartSeriesLegend_GetGroupIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetGroupIndex").Call(obj)
	return int32(ret)
}

func ChartSeriesLegend_SetGroupIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetGroupIndex").Call(obj, uintptr(value))
}

func ChartSeriesLegend_GetMultiplicity(obj uintptr) TLegendMultiplicity {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetMultiplicity").Call(obj)
	return TLegendMultiplicity(ret)
}

func ChartSeriesLegend_SetMultiplicity(obj uintptr, value TLegendMultiplicity) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetMultiplicity").Call(obj, uintptr(value))
}

func ChartSeriesLegend_GetOrder(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetOrder").Call(obj)
	return int32(ret)
}

func ChartSeriesLegend_SetOrder(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetOrder").Call(obj, uintptr(value))
}

func ChartSeriesLegend_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartSeriesLegend_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetTextFormat").Call(obj, uintptr(value))
}

func ChartSeriesLegend_GetUserItemsCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetUserItemsCount").Call(obj)
	return int32(ret)
}

func ChartSeriesLegend_SetUserItemsCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetUserItemsCount").Call(obj, uintptr(value))
}

func ChartSeriesLegend_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartSeriesLegend_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartSeriesLegend_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartSeriesLegend_GetOwner").Call(obj)
	return ret
}

func ChartSeriesLegend_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartSeriesLegend_SetOwner").Call(obj, AOwner)
}
