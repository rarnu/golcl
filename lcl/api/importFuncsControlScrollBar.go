package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TControlScrollBar ---------------------------

func ControlScrollBar_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ControlScrollBar_Assign").Call(obj, Source)
}

func ControlScrollBar_IsScrollBarVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_IsScrollBarVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ControlScrollBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlScrollBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ControlScrollBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ControlScrollBar_ClassType").Call(obj)
	return TClass(ret)
}

func ControlScrollBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlScrollBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ControlScrollBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlScrollBar_InstanceSize").Call(obj)
	return int32(ret)
}

func ControlScrollBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ControlScrollBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ControlScrollBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlScrollBar_GetHashCode").Call(obj)
	return int32(ret)
}

func ControlScrollBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlScrollBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ControlScrollBar_GetKind(obj uintptr) TScrollBarKind {
	ret, _, _ := getLazyProc("ControlScrollBar_GetKind").Call(obj)
	return TScrollBarKind(ret)
}

func ControlScrollBar_GetScrollPos(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlScrollBar_GetScrollPos").Call(obj)
	return int32(ret)
}

func ControlScrollBar_GetIncrement(obj uintptr) TScrollBarInc {
	ret, _, _ := getLazyProc("ControlScrollBar_GetIncrement").Call(obj)
	return TScrollBarInc(ret)
}

func ControlScrollBar_SetIncrement(obj uintptr, value TScrollBarInc) {
	_, _, _ = getLazyProc("ControlScrollBar_SetIncrement").Call(obj, uintptr(value))
}

func ControlScrollBar_GetPosition(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlScrollBar_GetPosition").Call(obj)
	return int32(ret)
}

func ControlScrollBar_SetPosition(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlScrollBar_SetPosition").Call(obj, uintptr(value))
}

func ControlScrollBar_GetRange(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlScrollBar_GetRange").Call(obj)
	return int32(ret)
}

func ControlScrollBar_SetRange(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlScrollBar_SetRange").Call(obj, uintptr(value))
}

func ControlScrollBar_GetSmooth(obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_GetSmooth").Call(obj)
	return DBoolToGoBool(ret)
}

func ControlScrollBar_SetSmooth(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ControlScrollBar_SetSmooth").Call(obj, GoBoolToDBool(value))
}

func ControlScrollBar_GetTracking(obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_GetTracking").Call(obj)
	return DBoolToGoBool(ret)
}

func ControlScrollBar_SetTracking(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ControlScrollBar_SetTracking").Call(obj, GoBoolToDBool(value))
}

func ControlScrollBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlScrollBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ControlScrollBar_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ControlScrollBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ControlScrollBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("ControlScrollBar_StaticClassType").Call()
	return TClass(r)
}
