package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TMainMenu ---------------------------

func MainMenu_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MainMenu_Create").Call(obj)
	return ret
}

func MainMenu_Free(obj uintptr) {
	getLazyProc("MainMenu_Free").Call(obj)
}

func MainMenu_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("MainMenu_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func MainMenu_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("MainMenu_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func MainMenu_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("MainMenu_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func MainMenu_Assign(obj uintptr, Source uintptr) {
	getLazyProc("MainMenu_Assign").Call(obj, Source)
}

func MainMenu_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MainMenu_ClassType").Call(obj)
	return TClass(ret)
}

func MainMenu_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MainMenu_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MainMenu_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MainMenu_InstanceSize").Call(obj)
	return int32(ret)
}

func MainMenu_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MainMenu_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MainMenu_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MainMenu_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MainMenu_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MainMenu_GetHashCode").Call(obj)
	return int32(ret)
}

func MainMenu_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MainMenu_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MainMenu_GetImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MainMenu_GetImagesWidth").Call(obj)
	return int32(ret)
}

func MainMenu_SetImagesWidth(obj uintptr, value int32) {
	getLazyProc("MainMenu_SetImagesWidth").Call(obj, uintptr(value))
}

func MainMenu_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("MainMenu_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func MainMenu_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("MainMenu_SetBiDiMode").Call(obj, uintptr(value))
}

func MainMenu_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MainMenu_GetImages").Call(obj)
	return ret
}

func MainMenu_SetImages(obj uintptr, value uintptr) {
	getLazyProc("MainMenu_SetImages").Call(obj, value)
}

func MainMenu_GetOwnerDraw(obj uintptr) bool {
	ret, _, _ := getLazyProc("MainMenu_GetOwnerDraw").Call(obj)
	return DBoolToGoBool(ret)
}

func MainMenu_SetOwnerDraw(obj uintptr, value bool) {
	getLazyProc("MainMenu_SetOwnerDraw").Call(obj, GoBoolToDBool(value))
}

func MainMenu_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("MainMenu_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func MainMenu_GetHandle(obj uintptr) HMENU {
	ret, _, _ := getLazyProc("MainMenu_GetHandle").Call(obj)
	return HMENU(ret)
}

func MainMenu_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MainMenu_GetItems").Call(obj)
	return ret
}

func MainMenu_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MainMenu_GetComponentCount").Call(obj)
	return int32(ret)
}

func MainMenu_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MainMenu_GetComponentIndex").Call(obj)
	return int32(ret)
}

func MainMenu_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("MainMenu_SetComponentIndex").Call(obj, uintptr(value))
}

func MainMenu_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MainMenu_GetOwner").Call(obj)
	return ret
}

func MainMenu_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("MainMenu_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func MainMenu_SetName(obj uintptr, value string) {
	getLazyProc("MainMenu_SetName").Call(obj, GoStrToDStr(value))
}

func MainMenu_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("MainMenu_GetTag").Call(obj)
	return int(ret)
}

func MainMenu_SetTag(obj uintptr, value int) {
	getLazyProc("MainMenu_SetTag").Call(obj, uintptr(value))
}

func MainMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("MainMenu_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func MainMenu_StaticClassType() TClass {
	r, _, _ := getLazyProc("MainMenu_StaticClassType").Call()
	return TClass(r)
}
