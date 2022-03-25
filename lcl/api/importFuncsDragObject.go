package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TDragObject ---------------------------

func DragObject_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragObject_Create").Call(obj)
	return ret
}

func DragObject_Free(obj uintptr) {
	_, _, _ = getLazyProc("DragObject_Free").Call(obj)
}

func DragObject_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("DragObject_Assign").Call(obj, Source)
}

func DragObject_HideDragImage(obj uintptr) {
	_, _, _ = getLazyProc("DragObject_HideDragImage").Call(obj)
}

func DragObject_ShowDragImage(obj uintptr) {
	_, _, _ = getLazyProc("DragObject_ShowDragImage").Call(obj)
}

func DragObject_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("DragObject_ClassType").Call(obj)
	return TClass(ret)
}

func DragObject_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("DragObject_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func DragObject_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DragObject_InstanceSize").Call(obj)
	return int32(ret)
}

func DragObject_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("DragObject_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func DragObject_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("DragObject_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func DragObject_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DragObject_GetHashCode").Call(obj)
	return int32(ret)
}

func DragObject_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("DragObject_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func DragObject_GetAlwaysShowDragImages(obj uintptr) bool {
	ret, _, _ := getLazyProc("DragObject_GetAlwaysShowDragImages").Call(obj)
	return DBoolToGoBool(ret)
}

func DragObject_SetAlwaysShowDragImages(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DragObject_SetAlwaysShowDragImages").Call(obj, GoBoolToDBool(value))
}

func DragObject_GetDragPos(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DragObject_GetDragPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragObject_SetDragPos(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("DragObject_SetDragPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragObject_GetDragTarget(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DragObject_GetDragTarget").Call(obj)
	return ret
}

func DragObject_SetDragTarget(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DragObject_SetDragTarget").Call(obj, value)
}

func DragObject_GetDragTargetPos(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DragObject_GetDragTargetPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DragObject_SetDragTargetPos(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("DragObject_SetDragTargetPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DragObject_GetDropped(obj uintptr) bool {
	ret, _, _ := getLazyProc("DragObject_GetDropped").Call(obj)
	return DBoolToGoBool(ret)
}

func DragObject_StaticClassType() TClass {
	r, _, _ := getLazyProc("DragObject_StaticClassType").Call()
	return TClass(r)
}
