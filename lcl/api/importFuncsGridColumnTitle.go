package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TGridColumnTitle ---------------------------

func GridColumnTitle_Assign(obj uintptr, Source uintptr) {
	getLazyProc("GridColumnTitle_Assign").Call(obj, Source)
}

func GridColumnTitle_FillTitleDefaultFont(obj uintptr) {
	getLazyProc("GridColumnTitle_FillTitleDefaultFont").Call(obj)
}

func GridColumnTitle_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("GridColumnTitle_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func GridColumnTitle_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("GridColumnTitle_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func GridColumnTitle_IsDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumnTitle_IsDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumnTitle_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumnTitle_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumnTitle_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("GridColumnTitle_ClassType").Call(obj)
	return TClass(ret)
}

func GridColumnTitle_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumnTitle_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumnTitle_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumnTitle_InstanceSize").Call(obj)
	return int32(ret)
}

func GridColumnTitle_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("GridColumnTitle_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func GridColumnTitle_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumnTitle_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func GridColumnTitle_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumnTitle_GetHashCode").Call(obj)
	return int32(ret)
}

func GridColumnTitle_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumnTitle_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumnTitle_GetColumn(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumnTitle_GetColumn").Call(obj)
	return ret
}

func GridColumnTitle_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("GridColumnTitle_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func GridColumnTitle_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("GridColumnTitle_SetAlignment").Call(obj, uintptr(value))
}

func GridColumnTitle_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumnTitle_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumnTitle_SetCaption(obj uintptr, value string) {
	getLazyProc("GridColumnTitle_SetCaption").Call(obj, GoStrToDStr(value))
}

func GridColumnTitle_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("GridColumnTitle_GetColor").Call(obj)
	return TColor(ret)
}

func GridColumnTitle_SetColor(obj uintptr, value TColor) {
	getLazyProc("GridColumnTitle_SetColor").Call(obj, uintptr(value))
}

func GridColumnTitle_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumnTitle_GetFont").Call(obj)
	return ret
}

func GridColumnTitle_SetFont(obj uintptr, value uintptr) {
	getLazyProc("GridColumnTitle_SetFont").Call(obj, value)
}

func GridColumnTitle_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumnTitle_GetImageIndex").Call(obj)
	return int32(ret)
}

func GridColumnTitle_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("GridColumnTitle_SetImageIndex").Call(obj, uintptr(value))
}

func GridColumnTitle_GetImageLayout(obj uintptr) TButtonLayout {
	ret, _, _ := getLazyProc("GridColumnTitle_GetImageLayout").Call(obj)
	return TButtonLayout(ret)
}

func GridColumnTitle_SetImageLayout(obj uintptr, value TButtonLayout) {
	getLazyProc("GridColumnTitle_SetImageLayout").Call(obj, uintptr(value))
}

func GridColumnTitle_GetLayout(obj uintptr) TTextLayout {
	ret, _, _ := getLazyProc("GridColumnTitle_GetLayout").Call(obj)
	return TTextLayout(ret)
}

func GridColumnTitle_SetLayout(obj uintptr, value TTextLayout) {
	getLazyProc("GridColumnTitle_SetLayout").Call(obj, uintptr(value))
}

func GridColumnTitle_GetMultiLine(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumnTitle_GetMultiLine").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumnTitle_SetMultiLine(obj uintptr, value bool) {
	getLazyProc("GridColumnTitle_SetMultiLine").Call(obj, GoBoolToDBool(value))
}

func GridColumnTitle_GetPrefixOption(obj uintptr) TPrefixOption {
	ret, _, _ := getLazyProc("GridColumnTitle_GetPrefixOption").Call(obj)
	return TPrefixOption(ret)
}

func GridColumnTitle_SetPrefixOption(obj uintptr, value TPrefixOption) {
	getLazyProc("GridColumnTitle_SetPrefixOption").Call(obj, uintptr(value))
}

func GridColumnTitle_StaticClassType() TClass {
	r, _, _ := getLazyProc("GridColumnTitle_StaticClassType").Call()
	return TClass(r)
}
