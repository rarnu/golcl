package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TCollectionItem ---------------------------

func CollectionItem_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CollectionItem_Create").Call(obj)
	return ret
}

func CollectionItem_Free(obj uintptr) {
	_, _, _ = getLazyProc("CollectionItem_Free").Call(obj)
}

func CollectionItem_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CollectionItem_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CollectionItem_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("CollectionItem_Assign").Call(obj, Source)
}

func CollectionItem_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CollectionItem_ClassType").Call(obj)
	return TClass(ret)
}

func CollectionItem_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CollectionItem_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CollectionItem_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CollectionItem_InstanceSize").Call(obj)
	return int32(ret)
}

func CollectionItem_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CollectionItem_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CollectionItem_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CollectionItem_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CollectionItem_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CollectionItem_GetHashCode").Call(obj)
	return int32(ret)
}

func CollectionItem_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CollectionItem_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CollectionItem_GetCollection(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CollectionItem_GetCollection").Call(obj)
	return ret
}

func CollectionItem_SetCollection(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CollectionItem_SetCollection").Call(obj, value)
}

func CollectionItem_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CollectionItem_GetIndex").Call(obj)
	return int32(ret)
}

func CollectionItem_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CollectionItem_SetIndex").Call(obj, uintptr(value))
}

func CollectionItem_GetDisplayName(obj uintptr) string {
	ret, _, _ := getLazyProc("CollectionItem_GetDisplayName").Call(obj)
	return DStrToGoStr(ret)
}

func CollectionItem_SetDisplayName(obj uintptr, value string) {
	_, _, _ = getLazyProc("CollectionItem_SetDisplayName").Call(obj, GoStrToDStr(value))
}

func CollectionItem_StaticClassType() TClass {
	r, _, _ := getLazyProc("CollectionItem_StaticClassType").Call()
	return TClass(r)
}
