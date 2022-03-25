package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMemoryStream ---------------------------

func MemoryStream_Create() uintptr {
	ret, _, _ := getLazyProc("MemoryStream_Create").Call()
	return ret
}

func MemoryStream_Free(obj uintptr) {
	_, _, _ = getLazyProc("MemoryStream_Free").Call(obj)
}

func MemoryStream_Clear(obj uintptr) {
	_, _, _ = getLazyProc("MemoryStream_Clear").Call(obj)
}

func MemoryStream_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("MemoryStream_LoadFromStream").Call(obj, Stream)
}

func MemoryStream_LoadFromFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("MemoryStream_LoadFromFile").Call(obj, GoStrToDStr(FileName))
}

func MemoryStream_Seek(obj uintptr, Offset int64, Origin TSeekOrigin) int64 {
	var ret int64
	_, _, _ = getLazyProc("MemoryStream_Seek").Call(obj, uintptr(unsafe.Pointer(&Offset)), uintptr(Origin), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MemoryStream_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("MemoryStream_SaveToStream").Call(obj, Stream)
}

func MemoryStream_SaveToFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("MemoryStream_SaveToFile").Call(obj, GoStrToDStr(FileName))
}

func MemoryStream_CopyFrom(obj uintptr, Source uintptr, Count int64) int64 {
	var ret int64
	_, _, _ = getLazyProc("MemoryStream_CopyFrom").Call(obj, Source, uintptr(unsafe.Pointer(&Count)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MemoryStream_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MemoryStream_ClassType").Call(obj)
	return TClass(ret)
}

func MemoryStream_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MemoryStream_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MemoryStream_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MemoryStream_InstanceSize").Call(obj)
	return int32(ret)
}

func MemoryStream_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MemoryStream_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MemoryStream_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MemoryStream_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MemoryStream_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MemoryStream_GetHashCode").Call(obj)
	return int32(ret)
}

func MemoryStream_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MemoryStream_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MemoryStream_GetMemory(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MemoryStream_GetMemory").Call(obj)
	return ret
}

func MemoryStream_GetPosition(obj uintptr) int64 {
	var ret int64
	_, _, _ = getLazyProc("MemoryStream_GetPosition").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MemoryStream_SetPosition(obj uintptr, value int64) {
	_, _, _ = getLazyProc("MemoryStream_SetPosition").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MemoryStream_GetSize(obj uintptr) int64 {
	var ret int64
	_, _, _ = getLazyProc("MemoryStream_GetSize").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MemoryStream_SetSize(obj uintptr, value int64) {
	_, _, _ = getLazyProc("MemoryStream_SetSize").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MemoryStream_StaticClassType() TClass {
	r, _, _ := getLazyProc("MemoryStream_StaticClassType").Call()
	return TClass(r)
}
