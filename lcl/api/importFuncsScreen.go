package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TScreen ---------------------------

func Screen_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_Create").Call(obj)
	return ret
}

func Screen_Free(obj uintptr) {
	_, _, _ = getLazyProc("Screen_Free").Call(obj)
}

func Screen_BeginTempCursor(obj uintptr, aCursor TCursor) {
	_, _, _ = getLazyProc("Screen_BeginTempCursor").Call(obj, uintptr(aCursor))
}

func Screen_EndTempCursor(obj uintptr, aCursor TCursor) {
	_, _, _ = getLazyProc("Screen_EndTempCursor").Call(obj, uintptr(aCursor))
}

func Screen_BeginWaitCursor(obj uintptr) {
	_, _, _ = getLazyProc("Screen_BeginWaitCursor").Call(obj)
}

func Screen_EndWaitCursor(obj uintptr) {
	_, _, _ = getLazyProc("Screen_EndWaitCursor").Call(obj)
}

func Screen_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Screen_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Screen_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Screen_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Screen_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Screen_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Screen_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Screen_Assign").Call(obj, Source)
}

func Screen_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Screen_ClassType").Call(obj)
	return TClass(ret)
}

func Screen_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Screen_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Screen_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_InstanceSize").Call(obj)
	return int32(ret)
}

func Screen_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Screen_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Screen_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Screen_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Screen_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetHashCode").Call(obj)
	return int32(ret)
}

func Screen_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Screen_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Screen_GetRealCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Screen_GetRealCursor").Call(obj)
	return TCursor(ret)
}

func Screen_GetFocusedForm(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetFocusedForm").Call(obj)
	return ret
}

func Screen_GetActiveControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetActiveControl").Call(obj)
	return ret
}

func Screen_GetActiveForm(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetActiveForm").Call(obj)
	return ret
}

func Screen_GetCustomFormCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetCustomFormCount").Call(obj)
	return int32(ret)
}

func Screen_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Screen_GetCursor").Call(obj)
	return TCursor(ret)
}

func Screen_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Screen_SetCursor").Call(obj, uintptr(value))
}

func Screen_GetMonitorCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetMonitorCount").Call(obj)
	return int32(ret)
}

func Screen_GetDesktopRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Screen_GetDesktopRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Screen_GetDesktopHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetDesktopHeight").Call(obj)
	return int32(ret)
}

func Screen_GetDesktopLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetDesktopLeft").Call(obj)
	return int32(ret)
}

func Screen_GetDesktopTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetDesktopTop").Call(obj)
	return int32(ret)
}

func Screen_GetDesktopWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetDesktopWidth").Call(obj)
	return int32(ret)
}

func Screen_GetWorkAreaRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Screen_GetWorkAreaRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Screen_GetWorkAreaHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetWorkAreaHeight").Call(obj)
	return int32(ret)
}

func Screen_GetWorkAreaLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetWorkAreaLeft").Call(obj)
	return int32(ret)
}

func Screen_GetWorkAreaTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetWorkAreaTop").Call(obj)
	return int32(ret)
}

func Screen_GetWorkAreaWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetWorkAreaWidth").Call(obj)
	return int32(ret)
}

func Screen_GetFonts(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetFonts").Call(obj)
	return ret
}

func Screen_GetFormCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetFormCount").Call(obj)
	return int32(ret)
}

func Screen_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetHeight").Call(obj)
	return int32(ret)
}

func Screen_GetPixelsPerInch(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetPixelsPerInch").Call(obj)
	return int32(ret)
}

func Screen_GetPrimaryMonitor(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetPrimaryMonitor").Call(obj)
	return ret
}

func Screen_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetWidth").Call(obj)
	return int32(ret)
}

func Screen_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetComponentCount").Call(obj)
	return int32(ret)
}

func Screen_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Screen_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Screen_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Screen_SetComponentIndex").Call(obj, uintptr(value))
}

func Screen_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Screen_GetOwner").Call(obj)
	return ret
}

func Screen_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Screen_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Screen_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Screen_SetName").Call(obj, GoStrToDStr(value))
}

func Screen_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Screen_GetTag").Call(obj)
	return int(ret)
}

func Screen_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Screen_SetTag").Call(obj, uintptr(value))
}

func Screen_GetCursors(obj uintptr, Index int32) HICON {
	ret, _, _ := getLazyProc("Screen_GetCursors").Call(obj, uintptr(Index))
	return ret
}

func Screen_SetCursors(obj uintptr, Index int32, value HICON) {
	_, _, _ = getLazyProc("Screen_SetCursors").Call(obj, uintptr(Index), value)
}

func Screen_GetMonitors(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Screen_GetMonitors").Call(obj, uintptr(Index))
	return ret
}

func Screen_GetForms(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Screen_GetForms").Call(obj, uintptr(Index))
	return ret
}

func Screen_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Screen_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Screen_StaticClassType() TClass {
	r, _, _ := getLazyProc("Screen_StaticClassType").Call()
	return TClass(r)
}
