package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TSynGutter ---------------------------

func SynGutter_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynGutter_Free").Call(obj)
}

func SynGutter_ClassName(obj uintptr) string {
	r, _, _ := getLazyProc("SynGutter_ClassName").Call(obj)
	return DStrToGoStr(r)
}

func SynGutter_GetHashCode(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynGutter_GetHashCode").Call(obj)
	return int32(r)
}

func SynGutter_Equals(obj, obj2 uintptr) bool {
	r, _, _ := getLazyProc("SynGutter_Equals").Call(obj, obj2)
	return DBoolToGoBool(r)
}

func SynGutter_ClassType(obj uintptr) TClass {
	r, _, _ := getLazyProc("SynGutter_ClassType").Call(obj)
	return TClass(r)
}

func SynGutter_InstanceSize(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynGutter_InstanceSize").Call(obj)
	return int32(r)
}

func SynGutter_InheritsFrom(obj uintptr, AClass TClass) bool {
	r, _, _ := getLazyProc("SynGutter_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(r)
}

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
