package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TGridColumn ---------------------------

func GridColumn_Assign(obj uintptr, Source uintptr) {
	getLazyProc("GridColumn_Assign").Call(obj, Source)
}

func GridColumn_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("GridColumn_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func GridColumn_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("GridColumn_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func GridColumn_IsDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_IsDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumn_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("GridColumn_ClassType").Call(obj)
	return TClass(ret)
}

func GridColumn_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_InstanceSize").Call(obj)
	return int32(ret)
}

func GridColumn_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("GridColumn_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func GridColumn_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func GridColumn_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetHashCode").Call(obj)
	return int32(ret)
}

func GridColumn_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_GetGrid(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumn_GetGrid").Call(obj)
	return ret
}

func GridColumn_GetDefaultWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetDefaultWidth").Call(obj)
	return int32(ret)
}

func GridColumn_GetStoredWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetStoredWidth").Call(obj)
	return int32(ret)
}

func GridColumn_GetWidthChanged(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_GetWidthChanged").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumn_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("GridColumn_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func GridColumn_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("GridColumn_SetAlignment").Call(obj, uintptr(value))
}

func GridColumn_GetButtonStyle(obj uintptr) TColumnButtonStyle {
	ret, _, _ := getLazyProc("GridColumn_GetButtonStyle").Call(obj)
	return TColumnButtonStyle(ret)
}

func GridColumn_SetButtonStyle(obj uintptr, value TColumnButtonStyle) {
	getLazyProc("GridColumn_SetButtonStyle").Call(obj, uintptr(value))
}

func GridColumn_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("GridColumn_GetColor").Call(obj)
	return TColor(ret)
}

func GridColumn_SetColor(obj uintptr, value TColor) {
	getLazyProc("GridColumn_SetColor").Call(obj, uintptr(value))
}

func GridColumn_GetDropDownRows(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetDropDownRows").Call(obj)
	return int32(ret)
}

func GridColumn_SetDropDownRows(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetDropDownRows").Call(obj, uintptr(value))
}

func GridColumn_GetExpanded(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_GetExpanded").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumn_SetExpanded(obj uintptr, value bool) {
	getLazyProc("GridColumn_SetExpanded").Call(obj, GoBoolToDBool(value))
}

func GridColumn_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumn_GetFont").Call(obj)
	return ret
}

func GridColumn_SetFont(obj uintptr, value uintptr) {
	getLazyProc("GridColumn_SetFont").Call(obj, value)
}

func GridColumn_GetLayout(obj uintptr) TTextLayout {
	ret, _, _ := getLazyProc("GridColumn_GetLayout").Call(obj)
	return TTextLayout(ret)
}

func GridColumn_SetLayout(obj uintptr, value TTextLayout) {
	getLazyProc("GridColumn_SetLayout").Call(obj, uintptr(value))
}

func GridColumn_GetMinSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetMinSize").Call(obj)
	return int32(ret)
}

func GridColumn_SetMinSize(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetMinSize").Call(obj, uintptr(value))
}

func GridColumn_GetMaxSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetMaxSize").Call(obj)
	return int32(ret)
}

func GridColumn_SetMaxSize(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetMaxSize").Call(obj, uintptr(value))
}

func GridColumn_GetPickList(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumn_GetPickList").Call(obj)
	return ret
}

func GridColumn_SetPickList(obj uintptr, value uintptr) {
	getLazyProc("GridColumn_SetPickList").Call(obj, value)
}

func GridColumn_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumn_SetReadOnly(obj uintptr, value bool) {
	getLazyProc("GridColumn_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func GridColumn_GetSizePriority(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetSizePriority").Call(obj)
	return int32(ret)
}

func GridColumn_SetSizePriority(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetSizePriority").Call(obj, uintptr(value))
}

func GridColumn_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("GridColumn_GetTag").Call(obj)
	return int(ret)
}

func GridColumn_SetTag(obj uintptr, value int) {
	getLazyProc("GridColumn_SetTag").Call(obj, uintptr(value))
}

func GridColumn_GetTitle(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumn_GetTitle").Call(obj)
	return ret
}

func GridColumn_SetTitle(obj uintptr, value uintptr) {
	getLazyProc("GridColumn_SetTitle").Call(obj, value)
}

func GridColumn_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetWidth").Call(obj)
	return int32(ret)
}

func GridColumn_SetWidth(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetWidth").Call(obj, uintptr(value))
}

func GridColumn_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("GridColumn_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func GridColumn_SetVisible(obj uintptr, value bool) {
	getLazyProc("GridColumn_SetVisible").Call(obj, GoBoolToDBool(value))
}

func GridColumn_GetValueChecked(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_GetValueChecked").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_SetValueChecked(obj uintptr, value string) {
	getLazyProc("GridColumn_SetValueChecked").Call(obj, GoStrToDStr(value))
}

func GridColumn_GetValueUnchecked(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_GetValueUnchecked").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_SetValueUnchecked(obj uintptr, value string) {
	getLazyProc("GridColumn_SetValueUnchecked").Call(obj, GoStrToDStr(value))
}

func GridColumn_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GridColumn_GetCollection").Call(obj)
	return ret
}

func GridColumn_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("GridColumn_SetCollection").Call(obj, value)
}

func GridColumn_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GridColumn_GetIndex").Call(obj)
	return int32(ret)
}

func GridColumn_SetIndex(obj uintptr, value int32) {
	getLazyProc("GridColumn_SetIndex").Call(obj, uintptr(value))
}

func GridColumn_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("GridColumn_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func GridColumn_SetDisplayName(obj uintptr, value string) {
	getLazyProc("GridColumn_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func GridColumn_StaticClassType() TClass {
	r, _, _ := getLazyProc("GridColumn_StaticClassType").Call()
	return TClass(r)
}
