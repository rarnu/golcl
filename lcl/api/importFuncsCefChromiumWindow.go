package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChromiumWindow_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_Create").Call(obj)
	return ret
}

func ChromiumWindow_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Free").Call(obj)
}

func ChromiumWindow_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChromiumWindow_StaticClassType").Call()
	return TClass(ret)
}

func ChromiumWindow_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetParent").Call(obj)
	return ret
}

func ChromiumWindow_SetParent(obj uintptr, Value uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SetParent").Call(obj, Value)
}

func ChromiumWindow_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ChromiumWindow_GetAlign").Call(obj)
	return TAlign(ret)
}

func ChromiumWindow_SetAlign(obj uintptr, Value TAlign) {
	_, _, _ = getLazyProc("ChromiumWindow_SetAlign").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ChromiumWindow_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ChromiumWindow_SetAnchors(obj uintptr, Value TAnchors) {
	_, _, _ = getLazyProc("ChromiumWindow_SetAnchors").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetConstraints").Call(obj)
	return ret
}

func ChromiumWindow_SetConstraints(obj uintptr, Value uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SetConstraints").Call(obj, Value)
}

func ChromiumWindow_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_SetDoubleBuffered(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_SetDoubleBuffered").Call(obj, GoBoolToDBool(Value))
}

func ChromiumWindow_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_SetEnabled(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_SetEnabled").Call(obj, GoBoolToDBool(Value))
}

func ChromiumWindow_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_SetShowHint(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_SetShowHint").Call(obj, GoBoolToDBool(Value))
}

func ChromiumWindow_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ChromiumWindow_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ChromiumWindow_SetTabOrder(obj uintptr, Value TTabOrder) {
	_, _, _ = getLazyProc("ChromiumWindow_SetTabOrder").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_SetTabStop(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_SetTabStop").Call(obj, GoBoolToDBool(Value))
}

func ChromiumWindow_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_SetVisible(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_SetVisible").Call(obj, GoBoolToDBool(Value))
}

func ChromiumWindow_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumWindow_GetLeft").Call(obj)
	return int32(ret)
}

func ChromiumWindow_SetLeft(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumWindow_SetLeft").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumWindow_GetTop").Call(obj)
	return int32(ret)
}

func ChromiumWindow_SetTop(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumWindow_SetTop").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumWindow_GetWidth").Call(obj)
	return int32(ret)
}

func ChromiumWindow_SetWidth(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumWindow_SetWidth").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumWindow_GetHeight").Call(obj)
	return int32(ret)
}

func ChromiumWindow_SetHeight(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumWindow_SetHeight").Call(obj, uintptr(Value))
}

func ChromiumWindow_GetTag(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetTag").Call(obj)
	return ret
}

func ChromiumWindow_SetTag(obj uintptr, Value uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SetTag").Call(obj, Value)
}

func ChromiumWindow_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumWindow_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumWindow_SetName(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumWindow_SetName").Call(obj, GoStrToDStr(Value))
}

func ChromiumWindow_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetBorderSpacing").Call(obj)
	return ret
}

func ChromiumWindow_SetBorderSpacing(obj uintptr, Value uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SetBorderSpacing").Call(obj, Value)
}

func ChromiumWindow_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ChromiumWindow_GetHandle").Call(obj)
	return ret
}

func ChromiumWindow_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetOwner").Call(obj)
	return ret
}

func ChromiumWindow_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Invalidate").Call(obj)
}

func ChromiumWindow_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SetFocus").Call(obj)
}

func ChromiumWindow_Update(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Update").Call(obj)
}

func ChromiumWindow_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_BringToFront").Call(obj)
}

func ChromiumWindow_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Hide").Call(obj)
}

func ChromiumWindow_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Refresh").Call(obj)
}

func ChromiumWindow_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_SendToBack").Call(obj)
}

func ChromiumWindow_Show(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_Show").Call(obj)
}

func ChromiumWindow_SetOnAfterCreated(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnAfterCreated").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_SetOnBeforeClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnBeforeClose").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_SetOnClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ChromiumWindow_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func ChromiumWindow_CreateBrowser(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_CreateBrowser").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_CloseBrowser(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ChromiumWindow_CloseBrowser").Call(obj, GoBoolToDBool(value))
}

func ChromiumWindow_LoadURL(obj uintptr, value string) {
	_, _, _ = getLazyProc("ChromiumWindow_LoadURL").Call(obj, GoStrToDStr(value))
}

func ChromiumWindow_NotifyMoveOrResizeStarted(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_NotifyMoveOrResizeStarted").Call(obj)
}

func ChromiumWindow_GetChromiumBrowser(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChromiumWindow_GetChromiumBrowser").Call(obj)
	return ret
}

func ChromiumWindow_GetInitialized(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChromiumWindow_GetInitialized").Call(obj)
	return DBoolToGoBool(ret)
}

func ChromiumWindow_UpdateSize(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumWindow_UpdateSize").Call(obj)
}
