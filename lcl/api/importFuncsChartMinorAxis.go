package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartMinorAxis_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_Create").Call(obj)
	return ret
}

func ChartMinorAxis_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_Free").Call(obj)
}

func ChartMinorAxis_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartMinorAxis_StaticClassType").Call()
	return TClass(ret)
}

func ChartMinorAxis_GetMarks(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetMarks").Call(obj)
	return ret
}

func ChartMinorAxis_SetMarks(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetMarks").Call(obj, value)
}

func ChartMinorAxis_GetAlignment(obj uintptr) TChartAxisAlignment {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetAlignment").Call(obj)
	return TChartAxisAlignment(ret)
}

func ChartMinorAxis_SetAlignment(obj uintptr, value TChartAxisAlignment) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetAlignment").Call(obj, uintptr(value))
}

func ChartMinorAxis_GetArrow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetArrow").Call(obj)
	return ret
}

func ChartMinorAxis_SetArrow(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetArrow").Call(obj, value)
}

func ChartMinorAxis_GetGrid(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetGrid").Call(obj)
	return ret
}

func ChartMinorAxis_SetGrid(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetGrid").Call(obj, value)
}

func ChartMinorAxis_GetIntervals(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetIntervals").Call(obj)
	return ret
}

func ChartMinorAxis_SetIntervals(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetIntervals").Call(obj, value)
}

func ChartMinorAxis_GetTickColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetTickColor").Call(obj)
	return TColor(ret)
}

func ChartMinorAxis_SetTickColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetTickColor").Call(obj, uintptr(value))
}

func ChartMinorAxis_GetTickInnerLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetTickInnerLength").Call(obj)
	return int32(ret)
}

func ChartMinorAxis_SetTickInnerLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetTickInnerLength").Call(obj, uintptr(value))
}

func ChartMinorAxis_GetTickLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetTickLength").Call(obj)
	return int32(ret)
}

func ChartMinorAxis_SetTickLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetTickLength").Call(obj, uintptr(value))
}

func ChartMinorAxis_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxis_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartMinorAxis_IsFlipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxis_IsFlipped").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxis_TryApplyStripes(obj uintptr, ADrawer uintptr, AIndex Cardinal) bool {
	ret, _, _ := getLazyProc("ChartMinorAxis_TryApplyStripes").Call(obj, ADrawer, uintptr(AIndex))
	return DBoolToGoBool(ret)
}

func ChartMinorAxis_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxis_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_Assign").Call(obj, Source)
}

func ChartMinorAxis_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartMinorAxis_ClassType").Call(obj)
	return TClass(ret)
}

func ChartMinorAxis_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxis_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxis_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxis_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartMinorAxis_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartMinorAxis_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartMinorAxis_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartMinorAxis_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartMinorAxis_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartMinorAxis_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxis_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxis_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetCollection").Call(obj)
	return ret
}

func ChartMinorAxis_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetCollection").Call(obj, value)
}

func ChartMinorAxis_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetIndex").Call(obj)
	return int32(ret)
}

func ChartMinorAxis_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetIndex").Call(obj, uintptr(value))
}

func ChartMinorAxis_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartMinorAxis_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartMinorAxis_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChartMinorAxis_SetDisplayName").Call(obj, GoStrToDStr(value))
}
