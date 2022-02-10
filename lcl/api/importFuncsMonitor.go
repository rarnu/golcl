package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMonitor ---------------------------

func Monitor_Create() uintptr {
	ret, _, _ := getLazyProc("Monitor_Create").Call()
	return ret
}

func Monitor_Free(obj uintptr) {
	getLazyProc("Monitor_Free").Call(obj)
}

func Monitor_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Monitor_ClassType").Call(obj)
	return TClass(ret)
}

func Monitor_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Monitor_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Monitor_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_InstanceSize").Call(obj)
	return int32(ret)
}

func Monitor_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Monitor_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Monitor_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Monitor_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Monitor_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetHashCode").Call(obj)
	return int32(ret)
}

func Monitor_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Monitor_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Monitor_GetHandle(obj uintptr) HMONITOR {
	ret, _, _ := getLazyProc("Monitor_GetHandle").Call(obj)
	return HMONITOR(ret)
}

func Monitor_GetMonitorNum(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetMonitorNum").Call(obj)
	return int32(ret)
}

func Monitor_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetLeft").Call(obj)
	return int32(ret)
}

func Monitor_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetHeight").Call(obj)
	return int32(ret)
}

func Monitor_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetTop").Call(obj)
	return int32(ret)
}

func Monitor_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetWidth").Call(obj)
	return int32(ret)
}

func Monitor_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Monitor_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Monitor_GetWorkareaRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Monitor_GetWorkareaRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Monitor_GetPrimary(obj uintptr) bool {
	ret, _, _ := getLazyProc("Monitor_GetPrimary").Call(obj)
	return DBoolToGoBool(ret)
}

func Monitor_GetPixelsPerInch(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Monitor_GetPixelsPerInch").Call(obj)
	return int32(ret)
}

func Monitor_StaticClassType() TClass {
	r, _, _ := getLazyProc("Monitor_StaticClassType").Call()
	return TClass(r)
}
