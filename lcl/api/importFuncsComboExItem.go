package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TComboExItem ---------------------------

func ComboExItem_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ComboExItem_Assign").Call(obj, Source)
}

func ComboExItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ComboExItem_ClassType").Call(obj)
	return TClass(ret)
}

func ComboExItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_InstanceSize").Call(obj)
	return int32(ret)
}

func ComboExItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ComboExItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ComboExItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboExItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ComboExItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_GetHashCode").Call(obj)
	return int32(ret)
}

func ComboExItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItem_GetIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_GetIndent").Call(obj)
	return int32(ret)
}

func ComboExItem_SetIndent(obj uintptr, value int32) {
	getLazyProc("ComboExItem_SetIndent").Call(obj, uintptr(value))
}

func ComboExItem_GetOverlayImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_GetOverlayImageIndex").Call(obj)
	return int32(ret)
}

func ComboExItem_SetOverlayImageIndex(obj uintptr, value int32) {
	getLazyProc("ComboExItem_SetOverlayImageIndex").Call(obj, uintptr(value))
}

func ComboExItem_GetData(obj uintptr) unsafe.Pointer {
	ret, _, _ := getLazyProc("ComboExItem_GetData").Call(obj)
	return unsafe.Pointer(ret)
}

func ComboExItem_SetData(obj uintptr, value unsafe.Pointer) {
	getLazyProc("ComboExItem_SetData").Call(obj, uintptr(value))
}

func ComboExItem_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItem_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItem_SetCaption(obj uintptr, value string) {
	getLazyProc("ComboExItem_SetCaption").Call(obj, GoStrToDStr(value))
}

func ComboExItem_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_GetImageIndex").Call(obj)
	return int32(ret)
}

func ComboExItem_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("ComboExItem_SetImageIndex").Call(obj, uintptr(value))
}

func ComboExItem_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboExItem_GetCollection").Call(obj)
	return ret
}

func ComboExItem_SetCollection(obj uintptr, value uintptr) {
	getLazyProc("ComboExItem_SetCollection").Call(obj, value)
}

func ComboExItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboExItem_GetIndex").Call(obj)
	return int32(ret)
}

func ComboExItem_SetIndex(obj uintptr, value int32) {
	getLazyProc("ComboExItem_SetIndex").Call(obj, uintptr(value))
}

func ComboExItem_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboExItem_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboExItem_SetDisplayName(obj uintptr, value string) {
	getLazyProc("ComboExItem_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func ComboExItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("ComboExItem_StaticClassType").Call()
	return TClass(r)
}
