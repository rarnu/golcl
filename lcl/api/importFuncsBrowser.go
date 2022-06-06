package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func Browser_Create(obj uintptr, align TAlign) uintptr {
	ret, _, _ := getLazyProc("Browser_Create").Call(obj, uintptr(align))
	return ret
}

func Browser_Free(obj uintptr) {
	_, _, _ = getLazyProc("Browser_Free").Call(obj)
}

func Browser_GetChromiumWindow(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Browser_GetChromiumWindow").Call(obj)
	return ret
}

func Browser_GetChromium(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Browser_GetChromium").Call(obj)
	return ret
}

func Browser_GetOpenLinkPopup(obj uintptr) bool {
	ret, _, _ := getLazyProc("Browser_GetOpenLinkPopup").Call(obj)
	return DBoolToGoBool(ret)
}

func Browser_SetOpenLinkPopup(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Browser_SetOpenLinkPopup").Call(obj, GoBoolToDBool(value))
}

func Browser_SetOnInitComplete(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Browser_SetOnInitComplete").Call(obj, addEventToMap(obj, fn))
}

func Browser_SetOnAfterCreated(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Browser_SetOnAfterCreated").Call(obj, addEventToMap(obj, fn))
}

func Browser_SetOnBeforeClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Browser_SetOnBeforeClose").Call(obj, addEventToMap(obj, fn))
}
