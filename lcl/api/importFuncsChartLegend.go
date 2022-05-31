package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartLegend_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_Create").Call(obj)
	return ret
}

func ChartLegend_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartLegend_Free").Call(obj)
}

func ChartLegend_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartLegend_StaticClassType").Call()
	return TClass(ret)
}

func ChartLegend_GetAlignment(obj uintptr) TLegendAlignment {
	ret, _, _ := getLazyProc("ChartLegend_GetAlignment").Call(obj)
	return TLegendAlignment(ret)
}

func ChartLegend_SetAlignment(obj uintptr, value TLegendAlignment) {
	_, _, _ = getLazyProc("ChartLegend_SetAlignment").Call(obj, uintptr(value))
}

func ChartLegend_GetBackgroundBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetBackgroundBrush").Call(obj)
	return ret
}

func ChartLegend_SetBackgroundBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetBackgroundBrush").Call(obj, value)
}

func ChartLegend_GetColumnCount(obj uintptr) TLegendColumnCount {
	ret, _, _ := getLazyProc("ChartLegend_GetColumnCount").Call(obj)
	return TLegendColumnCount(ret)
}

func ChartLegend_SetColumnCount(obj uintptr, value TLegendColumnCount) {
	_, _, _ = getLazyProc("ChartLegend_SetColumnCount").Call(obj, uintptr(value))
}

func ChartLegend_GetFixedItemWidth(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartLegend_GetFixedItemWidth").Call(obj)
	return Cardinal(ret)
}

func ChartLegend_SetFixedItemWidth(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartLegend_SetFixedItemWidth").Call(obj, uintptr(value))
}

func ChartLegend_GetFixedItemHeight(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartLegend_GetFixedItemHeight").Call(obj)
	return Cardinal(ret)
}

func ChartLegend_SetFixedItemHeight(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartLegend_SetFixedItemHeight").Call(obj, uintptr(value))
}

func ChartLegend_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetFont").Call(obj)
	return ret
}

func ChartLegend_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetFont").Call(obj, value)
}

func ChartLegend_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetFrame").Call(obj)
	return ret
}

func ChartLegend_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetFrame").Call(obj, value)
}

func ChartLegend_GetGridHorizontal(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetGridHorizontal").Call(obj)
	return ret
}

func ChartLegend_SetGridHorizontal(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetGridHorizontal").Call(obj, value)
}

func ChartLegend_GetGridVertical(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetGridVertical").Call(obj)
	return ret
}

func ChartLegend_SetGridVertical(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetGridVertical").Call(obj, value)
}

func ChartLegend_GetGroupFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetGroupFont").Call(obj)
	return ret
}

func ChartLegend_SetGroupFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetGroupFont").Call(obj, value)
}

func ChartLegend_GetGroupTitles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetGroupTitles").Call(obj)
	return ret
}

func ChartLegend_SetGroupTitles(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetGroupTitles").Call(obj, value)
}

func ChartLegend_GetInverted(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartLegend_GetInverted").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartLegend_SetInverted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartLegend_SetInverted").Call(obj, GoBoolToDBool(value))
}

func ChartLegend_GetItemFillOrder(obj uintptr) TLegendItemFillOrder {
	ret, _, _ := getLazyProc("ChartLegend_GetItemFillOrder").Call(obj)
	return TLegendItemFillOrder(ret)
}

func ChartLegend_SetItemFillOrder(obj uintptr, value TLegendItemFillOrder) {
	_, _, _ = getLazyProc("ChartLegend_SetItemFillOrder").Call(obj, uintptr(value))
}

func ChartLegend_GetMarginX(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartLegend_GetMarginX").Call(obj)
	return TChartDistance(ret)
}

func ChartLegend_SetMarginX(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartLegend_SetMarginX").Call(obj, uintptr(value))
}

func ChartLegend_GetMarginY(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartLegend_GetMarginY").Call(obj)
	return TChartDistance(ret)
}

func ChartLegend_SetMarginY(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartLegend_SetMarginY").Call(obj, uintptr(value))
}

func ChartLegend_GetSpacing(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartLegend_GetSpacing").Call(obj)
	return TChartDistance(ret)
}

func ChartLegend_SetSpacing(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartLegend_SetSpacing").Call(obj, uintptr(value))
}

func ChartLegend_GetSymbolFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetSymbolFrame").Call(obj)
	return ret
}

func ChartLegend_SetSymbolFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetSymbolFrame").Call(obj, value)
}

func ChartLegend_GetSymbolWidth(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("ChartLegend_GetSymbolWidth").Call(obj)
	return TChartDistance(ret)
}

func ChartLegend_SetSymbolWidth(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("ChartLegend_SetSymbolWidth").Call(obj, uintptr(value))
}

func ChartLegend_GetTextFormat(obj uintptr) TChartTextFormat {
	ret, _, _ := getLazyProc("ChartLegend_GetTextFormat").Call(obj)
	return TChartTextFormat(ret)
}

func ChartLegend_SetTextFormat(obj uintptr, value TChartTextFormat) {
	_, _, _ = getLazyProc("ChartLegend_SetTextFormat").Call(obj, uintptr(value))
}

func ChartLegend_GetTransparency(obj uintptr) TChartTransparency {
	ret, _, _ := getLazyProc("ChartLegend_GetTransparency").Call(obj)
	return TChartTransparency(ret)
}

func ChartLegend_SetTransparency(obj uintptr, value TChartTransparency) {
	_, _, _ = getLazyProc("ChartLegend_SetTransparency").Call(obj, uintptr(value))
}

func ChartLegend_GetUseSidebar(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartLegend_GetUseSidebar").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartLegend_SetUseSidebar(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartLegend_SetUseSidebar").Call(obj, GoBoolToDBool(value))
}

func ChartLegend_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartLegend_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartLegend_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartLegend_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartLegend_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartLegend_GetOwner").Call(obj)
	return ret
}

func ChartLegend_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("ChartLegend_SetOwner").Call(obj, AOwner)
}
