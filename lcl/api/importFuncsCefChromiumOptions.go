package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChromiumOptions_Create() uintptr {
	ret, _, _ := getLazyProc("ChromiumOptions_Create").Call()
	return ret
}

func ChromiumOptions_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChromiumOptions_Free").Call(obj)
}

func ChromiumOptions_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChromiumOptions_StaticClassType").Call()
	return TClass(ret)
}

func ChromiumOptions_GetJavascript(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetJavascript").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetJavascript(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetJavascript").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetJavascriptCloseWindows(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetJavascriptCloseWindows").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetJavascriptCloseWindows(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetJavascriptCloseWindows").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetJavascriptAccessClipboard(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetJavascriptAccessClipboard").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetJavascriptAccessClipboard(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetJavascriptAccessClipboard").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetJavascriptDomPaste(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetJavascriptDomPaste").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetJavascriptDomPaste(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetJavascriptDomPaste").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetImageLoading(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetImageLoading").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetImageLoading(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetImageLoading").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetImageShrinkStandaloneToFit(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetImageShrinkStandaloneToFit").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetImageShrinkStandaloneToFit(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetImageShrinkStandaloneToFit").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetTextAreaResize(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetTextAreaResize").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetTextAreaResize(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetTextAreaResize").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetTabToLinks(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetTabToLinks").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetTabToLinks(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetTabToLinks").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetLocalStorage(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetLocalStorage").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetLocalStorage(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetLocalStorage").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetDatabases(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetDatabases").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetDatabases(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetDatabases").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetWebgl(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetWebgl").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetWebgl(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetWebgl").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetBackgroundColor(obj uintptr) Cardinal {
	ret, _, _ := getLazyProc("ChromiumOptions_GetBackgroundColor").Call(obj)
	return Cardinal(ret)
}

func ChromiumOptions_SetBackgroundColor(obj uintptr, Value Cardinal) {
	_, _, _ = getLazyProc("ChromiumOptions_SetBackgroundColor").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetAcceptLanguageList(obj uintptr) string {
	ret, _, _ := getLazyProc("ChromiumOptions_GetAcceptLanguageList").Call(obj)
	return DStrToGoStr(ret)
}

func ChromiumOptions_SetAcceptLanguageList(obj uintptr, Value string) {
	_, _, _ = getLazyProc("ChromiumOptions_SetAcceptLanguageList").Call(obj, GoStrToDStr(Value))
}

func ChromiumOptions_GetWindowlessFrameRate(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ChromiumOptions_GetWindowlessFrameRate").Call(obj)
	return int32(ret)
}

func ChromiumOptions_SetWindowlessFrameRate(obj uintptr, Value int32) {
	_, _, _ = getLazyProc("ChromiumOptions_SetWindowlessFrameRate").Call(obj, uintptr(Value))
}

func ChromiumOptions_GetChromeStatusBubble(obj uintptr) TCefState {
	ret, _, _ := getLazyProc("ChromiumOptions_GetChromeStatusBubble").Call(obj)
	return TCefState(ret)
}

func ChromiumOptions_SetChromeStatusBubble(obj uintptr, Value TCefState) {
	_, _, _ = getLazyProc("ChromiumOptions_SetChromeStatusBubble").Call(obj, uintptr(Value))
}
