package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTrayIcon ---------------------------

func TrayIcon_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_Create").Call(obj)
	return ret
}

func TrayIcon_Free(obj uintptr) {
	_, _, _ = getLazyProc("TrayIcon_Free").Call(obj)
}

func TrayIcon_ShowBalloonHint(obj uintptr) {
	_, _, _ = getLazyProc("TrayIcon_ShowBalloonHint").Call(obj)
}

func TrayIcon_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func TrayIcon_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrayIcon_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func TrayIcon_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TrayIcon_Assign").Call(obj, Source)
}

func TrayIcon_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TrayIcon_ClassType").Call(obj)
	return TClass(ret)
}

func TrayIcon_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrayIcon_InstanceSize").Call(obj)
	return int32(ret)
}

func TrayIcon_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TrayIcon_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TrayIcon_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TrayIcon_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TrayIcon_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrayIcon_GetHashCode").Call(obj)
	return int32(ret)
}

func TrayIcon_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_GetAnimateInterval(obj uintptr) uint32 {
	ret, _, _ := getLazyProc("TrayIcon_GetAnimateInterval").Call(obj)
	return uint32(ret)
}

func TrayIcon_SetAnimateInterval(obj uintptr, value uint32) {
	_, _, _ = getLazyProc("TrayIcon_SetAnimateInterval").Call(obj, uintptr(value))
}

func TrayIcon_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("TrayIcon_SetHint").Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonHint(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_GetBalloonHint").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("TrayIcon_SetBalloonHint").Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_GetBalloonTitle").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("TrayIcon_SetBalloonTitle").Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTimeout(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrayIcon_GetBalloonTimeout").Call(obj)
	return int32(ret)
}

func TrayIcon_SetBalloonTimeout(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TrayIcon_SetBalloonTimeout").Call(obj, uintptr(value))
}

func TrayIcon_GetBalloonFlags(obj uintptr) TBalloonFlags {
	ret, _, _ := getLazyProc("TrayIcon_GetBalloonFlags").Call(obj)
	return TBalloonFlags(ret)
}

func TrayIcon_SetBalloonFlags(obj uintptr, value TBalloonFlags) {
	_, _, _ = getLazyProc("TrayIcon_SetBalloonFlags").Call(obj, uintptr(value))
}

func TrayIcon_GetIcon(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_GetIcon").Call(obj)
	return ret
}

func TrayIcon_SetIcon(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TrayIcon_SetIcon").Call(obj, value)
}

func TrayIcon_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_GetPopupMenu").Call(obj)
	return ret
}

func TrayIcon_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TrayIcon_SetPopupMenu").Call(obj, value)
}

func TrayIcon_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrayIcon_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TrayIcon_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TrayIcon_SetVisible").Call(obj, GoBoolToDBool(value))
}

func TrayIcon_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TrayIcon_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func TrayIcon_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TrayIcon_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func TrayIcon_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TrayIcon_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func TrayIcon_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TrayIcon_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func TrayIcon_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TrayIcon_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func TrayIcon_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrayIcon_GetComponentCount").Call(obj)
	return int32(ret)
}

func TrayIcon_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrayIcon_GetComponentIndex").Call(obj)
	return int32(ret)
}

func TrayIcon_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TrayIcon_SetComponentIndex").Call(obj, uintptr(value))
}

func TrayIcon_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_GetOwner").Call(obj)
	return ret
}

func TrayIcon_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TrayIcon_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TrayIcon_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("TrayIcon_SetName").Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("TrayIcon_GetTag").Call(obj)
	return int(ret)
}

func TrayIcon_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("TrayIcon_SetTag").Call(obj, uintptr(value))
}

func TrayIcon_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("TrayIcon_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func TrayIcon_StaticClassType() TClass {
	r, _, _ := getLazyProc("TrayIcon_StaticClassType").Call()
	return TClass(r)
}
