package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPrinterSetupDialog ---------------------------

func PrinterSetupDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PrinterSetupDialog_Create").Call(obj)
	return ret
}

func PrinterSetupDialog_Free(obj uintptr) {
	_, _, _ = getLazyProc("PrinterSetupDialog_Free").Call(obj)
}

func PrinterSetupDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrinterSetupDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func PrinterSetupDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PrinterSetupDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PrinterSetupDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PrinterSetupDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PrinterSetupDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PrinterSetupDialog_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("PrinterSetupDialog_Assign").Call(obj, Source)
}

func PrinterSetupDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PrinterSetupDialog_ClassType").Call(obj)
	return TClass(ret)
}

func PrinterSetupDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PrinterSetupDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PrinterSetupDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrinterSetupDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func PrinterSetupDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PrinterSetupDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PrinterSetupDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PrinterSetupDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PrinterSetupDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func PrinterSetupDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PrinterSetupDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PrinterSetupDialog_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetHandle").Call(obj)
	return ret
}

func PrinterSetupDialog_SetOnClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PrinterSetupDialog_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func PrinterSetupDialog_SetOnShow(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PrinterSetupDialog_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func PrinterSetupDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func PrinterSetupDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PrinterSetupDialog_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PrinterSetupDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func PrinterSetupDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetOwner").Call(obj)
	return ret
}

func PrinterSetupDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PrinterSetupDialog_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("PrinterSetupDialog_SetName").Call(obj, GoStrToDStr(value))
}

func PrinterSetupDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetTag").Call(obj)
	return int(ret)
}

func PrinterSetupDialog_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("PrinterSetupDialog_SetTag").Call(obj, uintptr(value))
}

func PrinterSetupDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PrinterSetupDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PrinterSetupDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("PrinterSetupDialog_StaticClassType").Call()
	return TClass(r)
}
