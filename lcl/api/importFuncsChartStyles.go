package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartStyles_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyles_Create").Call(obj)
	return ret
}

func ChartStyles_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyles_Free").Call(obj)
}

func ChartStyles_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartStyles_StaticClassType").Call()
	return TClass(ret)
}

func ChartStyles_GetStyles(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyles_GetStyles").Call(obj)
	return ret
}

func ChartStyles_SetStyles(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ChartStyles_SetStyles").Call(obj, value)
}

func ChartStyles_Add(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyles_Add").Call(obj)
	return ret
}

func ChartStyles_Apply(obj uintptr, ADrawer uintptr, AIndex Cardinal, IgnoreBrush bool) {
	_, _, _ = getLazyProc("ChartStyles_Apply").Call(obj, ADrawer, uintptr(AIndex), GoBoolToDBool(IgnoreBrush))
}

func ChartStyles_StyleByIndex(obj uintptr, AIndex Cardinal) uintptr {
	ret, _, _ := getLazyProc("ChartStyles_StyleByIndex").Call(obj, uintptr(AIndex))
	return ret
}
