package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartPen_Create() uintptr {
	ret, _, _ := getLazyProc("ChartPen_Create").Call()
	return ret
}

func ChartPen_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartPen_Free").Call(obj)
}

func ChartPen_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartPen_StaticClassType").Call()
	return TClass(ret)
}

func ChartPen_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartPen_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartPen_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChartPen_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ChartPen_Assign(obj uintptr, ASource uintptr) {
	_, _, _ = getLazyProc("ChartPen_Assign").Call(obj, ASource)
}

func ChartPen_EffVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartPen_EffVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChartPen_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartPen_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ChartPen_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ChartPen_ClassType").Call(obj)
	return TClass(ret)
}

func ChartPen_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Pen_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ChartPen_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartPen_InstanceSize").Call(obj)
	return int32(ret)
}

func ChartPen_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ChartPen_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ChartPen_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartPen_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ChartPen_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartPen_GetHashCode").Call(obj)
	return int32(ret)
}

func ChartPen_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ChartPen_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ChartPen_GetHandle(obj uintptr) HPEN {
	ret, _, _ := getLazyProc("ChartPen_GetHandle").Call(obj)
	return ret
}

func ChartPen_SetHandle(obj uintptr, value HPEN) {
	_, _, _ = getLazyProc("ChartPen_SetHandle").Call(obj, value)
}

func ChartPen_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ChartPen_GetColor").Call(obj)
	return TColor(ret)
}

func ChartPen_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ChartPen_SetColor").Call(obj, uintptr(value))
}

func ChartPen_GetMode(obj uintptr) TPenMode {
	ret, _, _ := getLazyProc("ChartPen_GetMode").Call(obj)
	return TPenMode(ret)
}

func ChartPen_SetMode(obj uintptr, value TPenMode) {
	_, _, _ = getLazyProc("ChartPen_SetMode").Call(obj, uintptr(value))
}

func ChartPen_GetStyle(obj uintptr) TPenStyle {
	ret, _, _ := getLazyProc("ChartPen_GetStyle").Call(obj)
	return TPenStyle(ret)
}

func ChartPen_SetStyle(obj uintptr, value TPenStyle) {
	_, _, _ = getLazyProc("ChartPen_SetStyle").Call(obj, uintptr(value))
}

func ChartPen_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChartPen_GetWidth").Call(obj)
	return int32(ret)
}

func ChartPen_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ChartPen_SetWidth").Call(obj, uintptr(value))
}

func ChartPen_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChartPen_SetOnChange").Call(obj, addEventToMap(obj, fn))
}
