package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TCoolBand ---------------------------

func CoolBand_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBand_Create").Call(obj)
	return ret
}

func CoolBand_Free(obj uintptr) {
	getLazyProc("CoolBand_Free").Call(obj)
}

func CoolBand_Assign(obj uintptr, Source uintptr) {
	getLazyProc("CoolBand_Assign").Call(obj, Source)
}

func CoolBand_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBand_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBand_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CoolBand_ClassType").Call(obj)
	return TClass(ret)
}

func CoolBand_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBand_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBand_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_InstanceSize").Call(obj)
	return int32(ret)
}

func CoolBand_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CoolBand_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CoolBand_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CoolBand_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetHashCode").Call(obj)
	return int32(ret)
}

func CoolBand_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBand_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBand_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetHeight").Call(obj)
	return int32(ret)
}

func CoolBand_GetBitmap(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBand_GetBitmap").Call(obj)
	return ret
}

func CoolBand_SetBitmap(obj uintptr, value uintptr) {
	getLazyProc("CoolBand_SetBitmap").Call(obj, value)
}

func CoolBand_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("CoolBand_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func CoolBand_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("CoolBand_SetBorderStyle").Call(obj, uintptr(value))
}

func CoolBand_GetBreak(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetBreak").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetBreak(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetBreak").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CoolBand_GetColor").Call(obj)
	return TColor(ret)
}

func CoolBand_SetColor(obj uintptr, value TColor) {
	getLazyProc("CoolBand_SetColor").Call(obj, uintptr(value))
}

func CoolBand_GetControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBand_GetControl").Call(obj)
	return ret
}

func CoolBand_SetControl(obj uintptr, value uintptr) {
	getLazyProc("CoolBand_SetControl").Call(obj, value)
}

func CoolBand_GetFixedBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetFixedBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetFixedBackground(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetFixedBackground").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetFixedSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetFixedSize").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetFixedSize(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetFixedSize").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetHorizontalOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetHorizontalOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetHorizontalOnly(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetHorizontalOnly").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetImageIndex").Call(obj)
	return int32(ret)
}

func CoolBand_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("CoolBand_SetImageIndex").Call(obj, uintptr(value))
}

func CoolBand_GetMinHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetMinHeight").Call(obj)
	return int32(ret)
}

func CoolBand_SetMinHeight(obj uintptr, value int32) {
	getLazyProc("CoolBand_SetMinHeight").Call(obj, uintptr(value))
}

func CoolBand_GetMinWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetMinWidth").Call(obj)
	return int32(ret)
}

func CoolBand_SetMinWidth(obj uintptr, value int32) {
	getLazyProc("CoolBand_SetMinWidth").Call(obj, uintptr(value))
}

func CoolBand_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetParentColor(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetParentBitmap(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetParentBitmap").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetParentBitmap(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetParentBitmap").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBand_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBand_SetText(obj uintptr, value string) {
	getLazyProc("CoolBand_SetText").Call(obj, GoStrToDStr(value))
}

func CoolBand_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBand_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBand_SetVisible(obj uintptr, value bool) {
	getLazyProc("CoolBand_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CoolBand_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetWidth").Call(obj)
	return int32(ret)
}

func CoolBand_SetWidth(obj uintptr, value int32) {
	getLazyProc("CoolBand_SetWidth").Call(obj, uintptr(value))
}

func CoolBand_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBand_GetCollection").Call(obj)
	return ret
}

func CoolBand_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("CoolBand_SetCollection").Call(obj, value)
}

func CoolBand_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBand_GetIndex").Call(obj)
	return int32(ret)
}

func CoolBand_SetIndex(obj uintptr, value int32) {
	getLazyProc("CoolBand_SetIndex").Call(obj, uintptr(value))
}

func CoolBand_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBand_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBand_SetDisplayName(obj uintptr, value string) {
	getLazyProc("CoolBand_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func CoolBand_StaticClassType() TClass {
	r, _, _ := getLazyProc("CoolBand_StaticClassType").Call()
	return TClass(r)
}
