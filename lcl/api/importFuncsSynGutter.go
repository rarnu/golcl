package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSynGutter ---------------------------

func SynGutter_GetAutoSize(obj uintptr) bool {
	r, _, _ := getLazyProc("SynGutter_GetAutoSize").Call(obj)
	return DBoolToGoBool(r)
}

func SynGutter_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynGutter_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func SynGutter_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynGutter_StaticClassType").Call()
	return TClass(r)
}

func SynGutter_GetColor(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynGutter_GetColor").Call(obj)
	return TColor(r)
}

func SynGutter_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SynGutter_SetColor").Call(obj, uintptr(value))
}

func SynGutter_GetVisible(obj uintptr) bool {
	r, _, _ := getLazyProc("SynGutter_GetVisible").Call(obj)
	return DBoolToGoBool(r)
}

func SynGutter_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynGutter_SetVisible").Call(obj, GoBoolToDBool(value))
}

func SynGutter_GetWidth(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynGutter_GetWidth").Call(obj)
	return int32(r)
}

func SynGutter_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynGutter_SetWidth").Call(obj, uintptr(value))
}
