package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartStyle_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyle_Create").Call(obj)
	return ret
}

func ChartStyle_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyle_Free").Call(obj)
}

func ChartStyle_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartStyle_StaticClassType").Call()
	return TClass(ret)
}

func ChartStyle_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyle_GetBrush").Call(obj)
	return ret
}

func ChartStyle_SetBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartStyle_SetBrush").Call(obj, value)
}

func ChartStyle_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyle_GetFont").Call(obj)
	return ret
}

func ChartStyle_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartStyle_SetFont").Call(obj, value)
}

func ChartStyle_GetPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyle_GetPen").Call(obj)
	return ret
}

func ChartStyle_SetPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartStyle_SetPen").Call(obj, value)
}

func ChartStyle_GetRepeatCount(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChartStyle_GetRepeatCount").Call(obj)
	return Cardinal(ret)
}

func ChartStyle_SetRepeatCount(obj uintptr, value Cardinal) {
	_, _, _ = getLazyProc("ChartStyle_SetRepeatCount").Call(obj, uintptr(value))
}

func ChartStyle_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyle_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func ChartStyle_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartStyle_SetText").Call(obj, GoStrToDStr(value))
}

func ChartStyle_GetUseBrush(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartStyle_GetUseBrush").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartStyle_SetUseBrush(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartStyle_SetUseBrush").Call(obj, GoBoolToDBool(value))
}

func ChartStyle_GetUseFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartStyle_GetUseFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartStyle_SetUseFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartStyle_SetUseFont").Call(obj, GoBoolToDBool(value))
}

func ChartStyle_GetUsePen(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartStyle_GetUsePen").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartStyle_SetUsePen(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartStyle_SetUsePen").Call(obj, GoBoolToDBool(value))
}

func ChartStyle_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyle_GetCollection").Call(obj)
	return ret
}

func ChartStyle_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartStyle_SetCollection").Call(obj, value)
}

func ChartStyle_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartStyle_GetIndex").Call(obj)
	return int32(ret)
}

func ChartStyle_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartStyle_SetIndex").Call(obj, uintptr(value))
}

func ChartStyle_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyle_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartStyle_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartStyle_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func ChartStyle_Apply(obj uintptr, ADrawer uintptr, IgnoreBrush bool) {
	_, _, _ = getLazyProc("ChartStyle_Apply").Call(obj, ADrawer, GoBoolToDBool(IgnoreBrush))
}

func ChartStyle_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartStyle_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}
