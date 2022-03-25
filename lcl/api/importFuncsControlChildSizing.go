package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TControlChildSizing ---------------------------

func ControlChildSizing_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ControlChildSizing_Assign").Call(obj, Source)
}

func ControlChildSizing_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlChildSizing_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ControlChildSizing_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ControlChildSizing_ClassType").Call(obj)
	return TClass(ret)
}

func ControlChildSizing_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlChildSizing_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ControlChildSizing_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_InstanceSize").Call(obj)
	return int32(ret)
}

func ControlChildSizing_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ControlChildSizing_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ControlChildSizing_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ControlChildSizing_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ControlChildSizing_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetHashCode").Call(obj)
	return int32(ret)
}

func ControlChildSizing_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ControlChildSizing_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ControlChildSizing_GetControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ControlChildSizing_GetControl").Call(obj)
	return ret
}

func ControlChildSizing_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ControlChildSizing_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ControlChildSizing_GetLeftRightSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetLeftRightSpacing").Call(obj)
	return int32(ret)
}

func ControlChildSizing_SetLeftRightSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlChildSizing_SetLeftRightSpacing").Call(obj, uintptr(value))
}

func ControlChildSizing_GetTopBottomSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetTopBottomSpacing").Call(obj)
	return int32(ret)
}

func ControlChildSizing_SetTopBottomSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlChildSizing_SetTopBottomSpacing").Call(obj, uintptr(value))
}

func ControlChildSizing_GetHorizontalSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetHorizontalSpacing").Call(obj)
	return int32(ret)
}

func ControlChildSizing_SetHorizontalSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlChildSizing_SetHorizontalSpacing").Call(obj, uintptr(value))
}

func ControlChildSizing_GetVerticalSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetVerticalSpacing").Call(obj)
	return int32(ret)
}

func ControlChildSizing_SetVerticalSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlChildSizing_SetVerticalSpacing").Call(obj, uintptr(value))
}

func ControlChildSizing_GetEnlargeHorizontal(obj uintptr) TChildControlResizeStyle {
	ret, _, _ := getLazyProc("ControlChildSizing_GetEnlargeHorizontal").Call(obj)
	return TChildControlResizeStyle(ret)
}

func ControlChildSizing_SetEnlargeHorizontal(obj uintptr, value TChildControlResizeStyle) {
	_, _, _ = getLazyProc("ControlChildSizing_SetEnlargeHorizontal").Call(obj, uintptr(value))
}

func ControlChildSizing_GetEnlargeVertical(obj uintptr) TChildControlResizeStyle {
	ret, _, _ := getLazyProc("ControlChildSizing_GetEnlargeVertical").Call(obj)
	return TChildControlResizeStyle(ret)
}

func ControlChildSizing_SetEnlargeVertical(obj uintptr, value TChildControlResizeStyle) {
	_, _, _ = getLazyProc("ControlChildSizing_SetEnlargeVertical").Call(obj, uintptr(value))
}

func ControlChildSizing_GetShrinkHorizontal(obj uintptr) TChildControlResizeStyle {
	ret, _, _ := getLazyProc("ControlChildSizing_GetShrinkHorizontal").Call(obj)
	return TChildControlResizeStyle(ret)
}

func ControlChildSizing_SetShrinkHorizontal(obj uintptr, value TChildControlResizeStyle) {
	_, _, _ = getLazyProc("ControlChildSizing_SetShrinkHorizontal").Call(obj, uintptr(value))
}

func ControlChildSizing_GetShrinkVertical(obj uintptr) TChildControlResizeStyle {
	ret, _, _ := getLazyProc("ControlChildSizing_GetShrinkVertical").Call(obj)
	return TChildControlResizeStyle(ret)
}

func ControlChildSizing_SetShrinkVertical(obj uintptr, value TChildControlResizeStyle) {
	_, _, _ = getLazyProc("ControlChildSizing_SetShrinkVertical").Call(obj, uintptr(value))
}

func ControlChildSizing_GetLayout(obj uintptr) TControlChildrenLayout {
	ret, _, _ := getLazyProc("ControlChildSizing_GetLayout").Call(obj)
	return TControlChildrenLayout(ret)
}

func ControlChildSizing_SetLayout(obj uintptr, value TControlChildrenLayout) {
	_, _, _ = getLazyProc("ControlChildSizing_SetLayout").Call(obj, uintptr(value))
}

func ControlChildSizing_GetControlsPerLine(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ControlChildSizing_GetControlsPerLine").Call(obj)
	return int32(ret)
}

func ControlChildSizing_SetControlsPerLine(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ControlChildSizing_SetControlsPerLine").Call(obj, uintptr(value))
}

func ControlChildSizing_StaticClassType() TClass {
	r, _, _ := getLazyProc("ControlChildSizing_StaticClassType").Call()
	return TClass(r)
}
