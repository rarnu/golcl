package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TParaAttributes ---------------------------

func ParaAttributes_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ParaAttributes_Assign").Call(obj, Source)
}

func ParaAttributes_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ParaAttributes_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ParaAttributes_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ParaAttributes_ClassType").Call(obj)
	return TClass(ret)
}

func ParaAttributes_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ParaAttributes_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ParaAttributes_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_InstanceSize").Call(obj)
	return int32(ret)
}

func ParaAttributes_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ParaAttributes_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ParaAttributes_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ParaAttributes_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ParaAttributes_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetHashCode").Call(obj)
	return int32(ret)
}

func ParaAttributes_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ParaAttributes_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ParaAttributes_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("ParaAttributes_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func ParaAttributes_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("ParaAttributes_SetAlignment").Call(obj, uintptr(value))
}

func ParaAttributes_GetFirstIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetFirstIndent").Call(obj)
	return int32(ret)
}

func ParaAttributes_SetFirstIndent(obj uintptr, value int32) {
	getLazyProc("ParaAttributes_SetFirstIndent").Call(obj, uintptr(value))
}

func ParaAttributes_GetLeftIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetLeftIndent").Call(obj)
	return int32(ret)
}

func ParaAttributes_SetLeftIndent(obj uintptr, value int32) {
	getLazyProc("ParaAttributes_SetLeftIndent").Call(obj, uintptr(value))
}

func ParaAttributes_GetNumbering(obj uintptr) TNumberingStyle {
	ret, _, _ := getLazyProc("ParaAttributes_GetNumbering").Call(obj)
	return TNumberingStyle(ret)
}

func ParaAttributes_SetNumbering(obj uintptr, value TNumberingStyle) {
	getLazyProc("ParaAttributes_SetNumbering").Call(obj, uintptr(value))
}

func ParaAttributes_GetRightIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetRightIndent").Call(obj)
	return int32(ret)
}

func ParaAttributes_SetRightIndent(obj uintptr, value int32) {
	getLazyProc("ParaAttributes_SetRightIndent").Call(obj, uintptr(value))
}

func ParaAttributes_GetTabCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetTabCount").Call(obj)
	return int32(ret)
}

func ParaAttributes_SetTabCount(obj uintptr, value int32) {
	getLazyProc("ParaAttributes_SetTabCount").Call(obj, uintptr(value))
}

func ParaAttributes_GetTab(obj uintptr, Index uint8) int32 {
	ret, _, _ := getLazyProc("ParaAttributes_GetTab").Call(obj, uintptr(Index))
	return int32(ret)
}

func ParaAttributes_SetTab(obj uintptr, Index uint8, value int32) {
	getLazyProc("ParaAttributes_SetTab").Call(obj, uintptr(Index), uintptr(value))
}

func ParaAttributes_StaticClassType() TClass {
	r, _, _ := getLazyProc("ParaAttributes_StaticClassType").Call()
	return TClass(r)
}
