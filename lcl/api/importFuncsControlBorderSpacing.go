package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TControlBorderSpacing ---------------------------

func ControlBorderSpacing_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ControlBorderSpacing_Assign").Call(obj, Source)
}

func ControlBorderSpacing_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ControlBorderSpacing_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ControlBorderSpacing_ClassType").Call(obj)
	return TClass(ret)
}

func ControlBorderSpacing_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlBorderSpacing_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ControlBorderSpacing_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_InstanceSize").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ControlBorderSpacing_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ControlBorderSpacing_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlBorderSpacing_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ControlBorderSpacing_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetHashCode").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlBorderSpacing_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ControlBorderSpacing_GetControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControl").Call(obj)
	return ret
}

func ControlBorderSpacing_GetAroundLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetAroundLeft").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetAroundTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetAroundTop").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetAroundRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetAroundRight").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetAroundBottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetAroundBottom").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlLeft").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlTop").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlWidth").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlHeight").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlRight").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_GetControlBottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetControlBottom").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ControlBorderSpacing_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ControlBorderSpacing_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetLeft").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetLeft(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetLeft").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetTop").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetTop(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetTop").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetRight").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetRight(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetRight").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetBottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetBottom").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetBottom(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetBottom").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetAround(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetAround").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetAround(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetAround").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetInnerBorder(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetInnerBorder").Call(obj)
	return int32(ret)
}

func ControlBorderSpacing_SetInnerBorder(obj uintptr, value int32) {
	getLazyProc("ControlBorderSpacing_SetInnerBorder").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetCellAlignHorizontal(obj uintptr) TControlCellAlign {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetCellAlignHorizontal").Call(obj)
	return TControlCellAlign(ret)
}

func ControlBorderSpacing_SetCellAlignHorizontal(obj uintptr, value TControlCellAlign) {
	getLazyProc("ControlBorderSpacing_SetCellAlignHorizontal").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetCellAlignVertical(obj uintptr) TControlCellAlign {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetCellAlignVertical").Call(obj)
	return TControlCellAlign(ret)
}

func ControlBorderSpacing_SetCellAlignVertical(obj uintptr, value TControlCellAlign) {
	getLazyProc("ControlBorderSpacing_SetCellAlignVertical").Call(obj, uintptr(value))
}

func ControlBorderSpacing_GetSpace(obj uintptr, Kind TAnchorKind) int32 {
	ret, _, _ := getLazyProc("ControlBorderSpacing_GetSpace").Call(obj, uintptr(Kind))
	return int32(ret)
}

func ControlBorderSpacing_SetSpace(obj uintptr, Kind TAnchorKind, value int32) {
	getLazyProc("ControlBorderSpacing_SetSpace").Call(obj, uintptr(Kind), uintptr(value))
}

func ControlBorderSpacing_StaticClassType() TClass {
	r, _, _ := getLazyProc("ControlBorderSpacing_StaticClassType").Call()
	return TClass(r)
}
