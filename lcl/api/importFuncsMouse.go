package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMouse ---------------------------

func Mouse_Create() uintptr {
	ret, _, _ := getLazyProc("Mouse_Create").Call()
	return ret
}

func Mouse_Free(obj uintptr) {
	getLazyProc("Mouse_Free").Call(obj)
}

func Mouse_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Mouse_ClassType").Call(obj)
	return TClass(ret)
}

func Mouse_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Mouse_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Mouse_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Mouse_InstanceSize").Call(obj)
	return int32(ret)
}

func Mouse_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Mouse_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Mouse_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Mouse_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Mouse_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Mouse_GetHashCode").Call(obj)
	return int32(ret)
}

func Mouse_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Mouse_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Mouse_GetCapture(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Mouse_GetCapture").Call(obj)
	return HWND(ret)
}

func Mouse_SetCapture(obj uintptr, value HWND) {
	getLazyProc("Mouse_SetCapture").Call(obj, uintptr(value))
}

func Mouse_GetCursorPos(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Mouse_GetCursorPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Mouse_SetCursorPos(obj uintptr, value TPoint) {
	getLazyProc("Mouse_SetCursorPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Mouse_GetIsDragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Mouse_GetIsDragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Mouse_GetWheelScrollLines(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Mouse_GetWheelScrollLines").Call(obj)
	return int32(ret)
}

func Mouse_StaticClassType() TClass {
	r, _, _ := getLazyProc("Mouse_StaticClassType").Call()
	return TClass(r)
}
