package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SeriesPointer_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SeriesPointer_Create").Call(obj)
	return ret
}

func SeriesPointer_Free(obj uintptr) {
	_, _, _ = getLazyProc("SeriesPointer_Free").Call(obj)
}

func SeriesPointer_StaticClassType() TClass {
	ret, _, _ := getLazyProc("SeriesPointer_StaticClassType").Call()
	return TClass(ret)
}

func SeriesPointer_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SeriesPointer_GetBrush").Call(obj)
	return ret
}

func SeriesPointer_SetBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SeriesPointer_SetBrush").Call(obj, value)
}

func SeriesPointer_GetHorizSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SeriesPointer_GetHorizSize").Call(obj)
	return int32(ret)
}

func SeriesPointer_SetHorizSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SeriesPointer_SetHorizSize").Call(obj, uintptr(value))
}

func SeriesPointer_GetOverrideColor(obj uintptr) TOverrideColors {
	ret, _, _ := getLazyProc("SeriesPointer_GetOverrideColor").Call(obj)
	return TOverrideColors(ret)
}

func SeriesPointer_SetOverrideColor(obj uintptr, value TOverrideColors) {
	_, _, _ = getLazyProc("SeriesPointer_SetOverrideColor").Call(obj, uintptr(value))
}

func SeriesPointer_GetPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SeriesPointer_GetPen").Call(obj)
	return ret
}

func SeriesPointer_SetPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SeriesPointer_SetPen").Call(obj, value)
}

func SeriesPointer_GetStyle(obj uintptr) TSeriesPointerStyle {
	ret, _, _ := getLazyProc("SeriesPointer_GetStyle").Call(obj)
	return TSeriesPointerStyle(ret) // convert from TSeriesPointerStyle
}

func SeriesPointer_SetStyle(obj uintptr, value TSeriesPointerStyle) {
	_, _, _ = getLazyProc("SeriesPointer_SetStyle").Call(obj, uintptr(value))
}

func SeriesPointer_GetVertSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SeriesPointer_GetVertSize").Call(obj)
	return int32(ret)
}

func SeriesPointer_SetVertSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SeriesPointer_SetVertSize").Call(obj, uintptr(value))
}

func SeriesPointer_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("SeriesPointer_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func SeriesPointer_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SeriesPointer_SetVisible").Call(obj, GoBoolToDBool(value))
}

func SeriesPointer_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SeriesPointer_GetOwner").Call(obj)
	return ret
}

func SeriesPointer_SetOwner(obj uintptr, AOwner uintptr) {
	_, _, _ = getLazyProc("SeriesPointer_SetOwner").Call(obj, AOwner)
}
