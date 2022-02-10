package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TDragDockObject ---------------------------

func DragDockObject_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragDockObject_Create").Call(obj)
	return ret
}

func DragDockObject_Free(obj uintptr) {
	getLazyProc("DragDockObject_Free").Call(obj)
}

func DragDockObject_Assign(obj uintptr, Source uintptr) {
	getLazyProc("DragDockObject_Assign").Call(obj, Source)
}

func DragDockObject_HideDragImage(obj uintptr) {
	getLazyProc("DragDockObject_HideDragImage").Call(obj)
}

func DragDockObject_ShowDragImage(obj uintptr) {
	getLazyProc("DragDockObject_ShowDragImage").Call(obj)
}

func DragDockObject_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("DragDockObject_ClassType").Call(obj)
	return TClass(ret)
}

func DragDockObject_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("DragDockObject_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func DragDockObject_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DragDockObject_InstanceSize").Call(obj)
	return int32(ret)
}

func DragDockObject_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("DragDockObject_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func DragDockObject_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("DragDockObject_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func DragDockObject_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DragDockObject_GetHashCode").Call(obj)
	return int32(ret)
}

func DragDockObject_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("DragDockObject_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func DragDockObject_GetDockRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("DragDockObject_GetDockRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragDockObject_SetDockRect(obj uintptr, value TRect) {
	getLazyProc("DragDockObject_SetDockRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragDockObject_GetDropAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("DragDockObject_GetDropAlign").Call(obj)
	return TAlign(ret)
}

func DragDockObject_GetDropOnControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragDockObject_GetDropOnControl").Call(obj)
	return ret
}

func DragDockObject_GetEraseDockRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("DragDockObject_GetEraseDockRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragDockObject_SetEraseDockRect(obj uintptr, value TRect) {
	getLazyProc("DragDockObject_SetEraseDockRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragDockObject_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("DragDockObject_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func DragDockObject_SetFloating(obj uintptr, value bool) {
	getLazyProc("DragDockObject_SetFloating").Call(obj, GoBoolToDBool(value))
}

func DragDockObject_GetControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragDockObject_GetControl").Call(obj)
	return ret
}

func DragDockObject_SetControl(obj uintptr, value uintptr) {
	getLazyProc("DragDockObject_SetControl").Call(obj, value)
}

func DragDockObject_GetAlwaysShowDragImages(obj uintptr) bool {
	ret, _, _ := getLazyProc("DragDockObject_GetAlwaysShowDragImages").Call(obj)
	return DBoolToGoBool(ret)
}

func DragDockObject_SetAlwaysShowDragImages(obj uintptr, value bool) {
	getLazyProc("DragDockObject_SetAlwaysShowDragImages").Call(obj, GoBoolToDBool(value))
}

func DragDockObject_GetDragPos(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("DragDockObject_GetDragPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragDockObject_SetDragPos(obj uintptr, value TPoint) {
	getLazyProc("DragDockObject_SetDragPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragDockObject_GetDragTarget(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragDockObject_GetDragTarget").Call(obj)
	return ret
}

func DragDockObject_SetDragTarget(obj uintptr, value uintptr) {
	getLazyProc("DragDockObject_SetDragTarget").Call(obj, value)
}

func DragDockObject_GetDragTargetPos(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("DragDockObject_GetDragTargetPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragDockObject_SetDragTargetPos(obj uintptr, value TPoint) {
	getLazyProc("DragDockObject_SetDragTargetPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragDockObject_GetDropped(obj uintptr) bool {
	ret, _, _ := getLazyProc("DragDockObject_GetDropped").Call(obj)
	return DBoolToGoBool(ret)
}

func DragDockObject_StaticClassType() TClass {
	r, _, _ := getLazyProc("DragDockObject_StaticClassType").Call()
	return TClass(r)
}
