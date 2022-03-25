package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TApplication ---------------------------

func Application_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Application_Create").Call(obj)
	return ret
}

func Application_Free(obj uintptr) {
	_, _, _ = getLazyProc("Application_Free").Call(obj)
}

func Application_ActivateHint(obj uintptr, CursorPos TPoint) {
	_, _, _ = getLazyProc("Application_ActivateHint").Call(obj, uintptr(unsafe.Pointer(&CursorPos)))
}

func Application_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Application_BringToFront").Call(obj)
}

func Application_CancelHint(obj uintptr) {
	_, _, _ = getLazyProc("Application_CancelHint").Call(obj)
}

func Application_HandleMessage(obj uintptr) {
	_, _, _ = getLazyProc("Application_HandleMessage").Call(obj)
}

func Application_HideHint(obj uintptr) {
	_, _, _ = getLazyProc("Application_HideHint").Call(obj)
}

func Application_Minimize(obj uintptr) {
	_, _, _ = getLazyProc("Application_Minimize").Call(obj)
}

func Application_ModalStarted(obj uintptr) {
	_, _, _ = getLazyProc("Application_ModalStarted").Call(obj)
}

func Application_ModalFinished(obj uintptr) {
	_, _, _ = getLazyProc("Application_ModalFinished").Call(obj)
}

func Application_ProcessMessages(obj uintptr) {
	_, _, _ = getLazyProc("Application_ProcessMessages").Call(obj)
}

func Application_Restore(obj uintptr) {
	_, _, _ = getLazyProc("Application_Restore").Call(obj)
}

func Application_RestoreTopMosts(obj uintptr) {
	_, _, _ = getLazyProc("Application_RestoreTopMosts").Call(obj)
}

func Application_Terminate(obj uintptr) {
	_, _, _ = getLazyProc("Application_Terminate").Call(obj)
}

func Application_MessageBox(obj uintptr, Text string, Caption string, Flags int32) int32 {
	ret, _, _ := getLazyProc("Application_MessageBox").Call(obj, GoStrToDStr(Text), GoStrToDStr(Caption), uintptr(Flags))
	return int32(ret)
}

func Application_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Application_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Application_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Application_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Application_Assign").Call(obj, Source)
}

func Application_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Application_ClassType").Call(obj)
	return TClass(ret)
}

func Application_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Application_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_InstanceSize").Call(obj)
	return int32(ret)
}

func Application_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Application_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Application_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Application_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetHashCode").Call(obj)
	return int32(ret)
}

func Application_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Application_GetScaled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetScaled").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetScaled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetScaled").Call(obj, GoBoolToDBool(value))
}

func Application_GetSingleInstanceEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetSingleInstanceEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetSingleInstanceEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetSingleInstanceEnabled").Call(obj, GoBoolToDBool(value))
}

func Application_GetLocation(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetLocation").Call(obj)
	return DStrToGoStr(ret)
}

func Application_GetStopOnException(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetStopOnException").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetStopOnException(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetStopOnException").Call(obj, GoBoolToDBool(value))
}

func Application_GetExceptionExitCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetExceptionExitCode").Call(obj)
	return int32(ret)
}

func Application_SetExceptionExitCode(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Application_SetExceptionExitCode").Call(obj, uintptr(value))
}

func Application_GetExeName(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetExeName").Call(obj)
	return DStrToGoStr(ret)
}

func Application_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Application_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Application_SetHint").Call(obj, GoStrToDStr(value))
}

func Application_GetHintColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Application_GetHintColor").Call(obj)
	return TColor(ret)
}

func Application_SetHintColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Application_SetHintColor").Call(obj, uintptr(value))
}

func Application_GetHintHidePause(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetHintHidePause").Call(obj)
	return int32(ret)
}

func Application_SetHintHidePause(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Application_SetHintHidePause").Call(obj, uintptr(value))
}

func Application_GetHintPause(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetHintPause").Call(obj)
	return int32(ret)
}

func Application_SetHintPause(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Application_SetHintPause").Call(obj, uintptr(value))
}

func Application_GetHintShortCuts(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetHintShortCuts").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetHintShortCuts(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetHintShortCuts").Call(obj, GoBoolToDBool(value))
}

func Application_GetHintShortPause(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetHintShortPause").Call(obj)
	return int32(ret)
}

func Application_SetHintShortPause(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Application_SetHintShortPause").Call(obj, uintptr(value))
}

func Application_GetIcon(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Application_GetIcon").Call(obj)
	return ret
}

func Application_SetIcon(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Application_SetIcon").Call(obj, value)
}

func Application_GetMainForm(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Application_GetMainForm").Call(obj)
	return ret
}

func Application_GetMainFormHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Application_GetMainFormHandle").Call(obj)
	return ret
}

func Application_GetMainFormOnTaskBar(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetMainFormOnTaskBar").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetMainFormOnTaskBar(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetMainFormOnTaskBar").Call(obj, GoBoolToDBool(value))
}

func Application_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Application_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Application_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Application_SetBiDiMode").Call(obj, uintptr(value))
}

func Application_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Application_GetShowMainForm(obj uintptr) bool {
	ret, _, _ := getLazyProc("Application_GetShowMainForm").Call(obj)
	return DBoolToGoBool(ret)
}

func Application_SetShowMainForm(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Application_SetShowMainForm").Call(obj, GoBoolToDBool(value))
}

func Application_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func Application_SetTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("Application_SetTitle").Call(obj, GoStrToDStr(value))
}

func Application_SetOnActivate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnActivate").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnDeactivate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnDeactivate").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnException(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnException").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnHelp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnHelp").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnHint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnHint").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnMinimize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnMinimize").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnRestore(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnRestore").Call(obj, addEventToMap(obj, fn))
}

func Application_SetOnShortCut(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Application_SetOnShortCut").Call(obj, addEventToMap(obj, fn))
}

func Application_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Application_GetHandle").Call(obj)
	return ret
}

func Application_SetHandle(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Application_SetHandle").Call(obj, value)
}

func Application_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetComponentCount").Call(obj)
	return int32(ret)
}

func Application_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Application_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Application_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Application_SetComponentIndex").Call(obj, uintptr(value))
}

func Application_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Application_GetOwner").Call(obj)
	return ret
}

func Application_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Application_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Application_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Application_SetName").Call(obj, GoStrToDStr(value))
}

func Application_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Application_GetTag").Call(obj)
	return int(ret)
}

func Application_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Application_SetTag").Call(obj, uintptr(value))
}

func Application_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Application_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Application_StaticClassType() TClass {
	r, _, _ := getLazyProc("Application_StaticClassType").Call()
	return TClass(r)
}
