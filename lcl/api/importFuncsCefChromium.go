package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

func Chromium_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chromium_Create").Call(obj)
	return ret
}

func Chromium_Free(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_Free").Call(obj)
}

func Chromium_StaticClassType() TClass {
	ret, _, _ := getLazyProc("Chromium_StaticClassType").Call()
	return TClass(ret)
}

func Chromium_ShowDevTools(obj uintptr, inspectElementAt TPoint) {
	_, _, _ = getLazyProc("Chromium_ShowDevTools").Call(obj, uintptr(unsafe.Pointer(&inspectElementAt)))
}

func Chromium_CloseDevTools(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_CloseDevTools").Call(obj)
}

func Chromium_MoveFormTo(obj uintptr, x int32, y int32) {
	_, _, _ = getLazyProc("Chromium_MoveFormTo").Call(obj, uintptr(x), uintptr(y))
}

func Chromium_MoveFormBy(obj uintptr, x int32, y int32) {
	_, _, _ = getLazyProc("Chromium_MoveFormBy").Call(obj, uintptr(x), uintptr(y))
}

func Chromium_ResizeFormWidthTo(obj uintptr, x int32) {
	_, _, _ = getLazyProc("Chromium_ResizeFormWidthTo").Call(obj, uintptr(x))
}

func Chromium_ResizeFormHeightTo(obj uintptr, y int32) {
	_, _, _ = getLazyProc("Chromium_ResizeFormHeightTo").Call(obj, uintptr(y))
}

func Chromium_SetFormLeftTo(obj uintptr, x int32) {
	_, _, _ = getLazyProc("Chromium_SetFormLeftTo").Call(obj, uintptr(x))
}

func Chromium_SetFormTopTo(obj uintptr, y int32) {
	_, _, _ = getLazyProc("Chromium_SetFormTopTo").Call(obj, uintptr(y))
}

func Chromium_SaveAsBitmapStream(obj uintptr, stream uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_SaveAsBitmapStream").Call(obj, stream)
	return DBoolToGoBool(ret)
}

func Chromium_TakeSnapshot(obj uintptr, bitmap uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_TakeSnapshot").Call(obj, bitmap)
	return DBoolToGoBool(ret)
}

func Chromium_CloseBrowser(obj uintptr, force bool) {
	_, _, _ = getLazyProc("Chromium_CloseBrowser").Call(obj, GoBoolToDBool(force))
}

func Chromium_CloseAllBrowsers(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_CloseAllBrowsers").Call(obj)
}

func Chromium_TryCloseBrowser(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_TryCloseBrowser").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_LoadURL(obj uintptr, url string) {
	_, _, _ = getLazyProc("Chromium_LoadURL").Call(obj, GoStrToDStr(url))
}

func Chromium_LoadString(obj uintptr, html string) {
	_, _, _ = getLazyProc("Chromium_LoadString").Call(obj, GoStrToDStr(html))
}

func Chromium_GoBack(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_GoBack").Call(obj)
}

func Chromium_GoForward(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_GoForward").Call(obj)
}

func Chromium_Reload(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_Reload").Call(obj)
}

func Chromium_ReloadIgnoreCache(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ReloadIgnoreCache").Call(obj)
}

func Chromium_StopLoad(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_StopLoad").Call(obj)
}

func Chromium_StartDownload(obj uintptr, url string) {
	_, _, _ = getLazyProc("Chromium_StartDownload").Call(obj, GoStrToDStr(url))
}

func Chromium_DownloadImage(obj uintptr, imageUrl string, isFav bool, maxImageSize Cardinal, bypassCache bool) {
	_, _, _ = getLazyProc("Chromium_DownloadImage").Call(obj, GoStrToDStr(imageUrl), GoBoolToDBool(isFav), uintptr(maxImageSize), GoBoolToDBool(bypassCache))
}

func Chromium_ClearCertificateExceptions(obj uintptr, aClearImmediately bool) bool {
	ret, _, _ := getLazyProc("Chromium_ClearCertificateExceptions").Call(obj, GoBoolToDBool(aClearImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_ClearHttpAuthCredentials(obj uintptr, aClearImmediately bool) bool {
	ret, _, _ := getLazyProc("Chromium_ClearHttpAuthCredentials").Call(obj, GoBoolToDBool(aClearImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_CloseAllConnections(obj uintptr, aCloseImmediately bool) bool {
	ret, _, _ := getLazyProc("Chromium_CloseAllConnections").Call(obj, GoBoolToDBool(aCloseImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_ExecuteJavaScript(obj uintptr, code string, scriptUrl string) {
	_, _, _ = getLazyProc("Chromium_ExecuteJavaScript").Call(obj, GoStrToDStr(code), GoStrToDStr(scriptUrl))
}

func Chromium_UpdatePreferences(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_UpdatePreferences").Call(obj)
}

func Chromium_SavePreferences(obj uintptr, filename string) {
	_, _, _ = getLazyProc("Chromium_SavePreferences").Call(obj, GoStrToDStr(filename))
}

func Chromium_ResolveHost(obj uintptr, url string) {
	_, _, _ = getLazyProc("Chromium_ResolveHost").Call(obj, GoStrToDStr(url))
}

func Chromium_ClearCache(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClearCache").Call(obj)
}

func Chromium_DeleteCookies(obj uintptr, aUrl string, cookieName string, aDeleteImmediately bool) bool {
	ret, _, _ := getLazyProc("Chromium_DeleteCookies").Call(obj, GoStrToDStr(aUrl), GoStrToDStr(cookieName), GoBoolToDBool(aDeleteImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_SetCookie(obj uintptr,
	aUrl string, cookieName string, cookieValue string, aDomain string, aPath string,
	aSecure bool, aHttpOnly, aHasExpires bool,
	aCreation, aLastAccess, aExpires time.Time,
	aSameSite TCefCookieSameSite, aPriority TCefCookiePriority, aSetImmediately bool) bool {
	v1 := aCreation.Unix()
	v2 := aLastAccess.Unix()
	v3 := aExpires.Unix()
	ret, _, _ := getLazyProc("Chromium_SetCookie").Call(obj, GoStrToDStr(aUrl), GoStrToDStr(cookieName), GoStrToDStr(cookieValue), GoStrToDStr(aDomain), GoStrToDStr(aPath),
		GoBoolToDBool(aSecure), GoBoolToDBool(aHttpOnly), GoBoolToDBool(aHasExpires),
		uintptr(unsafe.Pointer(&v1)), uintptr(unsafe.Pointer(&v2)), uintptr(unsafe.Pointer(&v3)),
		uintptr(aSameSite), uintptr(aPriority), GoBoolToDBool(aSetImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_FlushCookieStore(obj uintptr, aFlushImmediately bool) bool {
	ret, _, _ := getLazyProc("Chromium_FlushCookieStore").Call(obj, GoBoolToDBool(aFlushImmediately))
	return DBoolToGoBool(ret)
}

func Chromium_SendDevToolsMessage(obj uintptr, message string) bool {
	ret, _, _ := getLazyProc("Chromium_SendDevToolsMessage").Call(obj, GoStrToDStr(message))
	return DBoolToGoBool(ret)
}

func Chromium_Find(obj uintptr, text string, forward bool, matchCase bool, findNext bool) {
	_, _, _ = getLazyProc("Chromium_Find").Call(obj, GoStrToDStr(text), GoBoolToDBool(forward), GoBoolToDBool(matchCase), GoBoolToDBool(findNext))
}

func Chromium_StopFinding(obj uintptr, clearSelection bool) {
	_, _, _ = getLazyProc("Chromium_StopFinding").Call(obj, GoBoolToDBool(clearSelection))
}

func Chromium_Print(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_Print").Call(obj)
}

func Chromium_PrintToPDF(obj uintptr, filepath string, title string, url string) {
	_, _, _ = getLazyProc("Chromium_PrintToPDF").Call(obj, GoStrToDStr(filepath), GoStrToDStr(title), GoStrToDStr(url))
}

func Chromium_ClipboardCopy(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardCopy").Call(obj)
}

func Chromium_ClipboardPaste(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardPaste").Call(obj)
}

func Chromium_ClipboardCut(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardCut").Call(obj)
}

func Chromium_ClipboardUndo(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardUndo").Call(obj)
}

func Chromium_ClipboardRedo(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardRedo").Call(obj)
}

func Chromium_ClipboardDel(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ClipboardDel").Call(obj)
}

func Chromium_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_SelectAll").Call(obj)
}

func Chromium_IncZoomStep(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_IncZoomStep").Call(obj)
}

func Chromium_DecZoomStep(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_DecZoomStep").Call(obj)
}

func Chromium_IncZoomPct(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_IncZoomPct").Call(obj)
}

func Chromium_DecZoomPct(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_DecZoomPct").Call(obj)
}

func Chromium_ResetZoomStep(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ResetZoomStep").Call(obj)
}

func Chromium_ResetZoomLevel(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ResetZoomLevel").Call(obj)
}

func Chromium_ResetZoomPct(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ResetZoomPct").Call(obj)
}

func Chromium_ReadZoom(obj uintptr) {
	_, _, _ = getLazyProc("Chromium_ReadZoom").Call(obj)
}

func Chromium_GetMultithreadApp(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetMultithreadApp").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetIsLoading(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetIsLoading").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetHasDocument(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetHasDocument").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetHasView(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetHasView").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetHasDevTools(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetHasDevTools").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetCanGoBack(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetCanGoBack").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetCanGoForward(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetCanGoForward").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetIsPopUp(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetIsPopUp").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetInitialized(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetInitialized").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_GetDocumentURL(obj uintptr) string {
	ret, _, _ := getLazyProc("Chromium_GetDocumentURL").Call(obj)
	return DStrToGoStr(ret)
}

func Chromium_GetDoNotTrack(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetDoNotTrack").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_SetDoNotTrack(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chromium_SetDoNotTrack").Call(obj, GoBoolToDBool(value))
}

func Chromium_GetBlock3rdPartyCookies(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetBlock3rdPartyCookies").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_SetBlock3rdPartyCookies(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chromium_SetBlock3rdPartyCookies").Call(obj, GoBoolToDBool(value))
}

func Chromium_GetJavascriptEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetJavascriptEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_SetJavascriptEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chromium_SetJavascriptEnabled").Call(obj, GoBoolToDBool(value))
}

func Chromium_GetLoadImagesAutomatically(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chromium_GetLoadImagesAutomatically").Call(obj)
	return DBoolToGoBool(ret)
}

func Chromium_SetLoadImagesAutomatically(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chromium_SetLoadImagesAutomatically").Call(obj, GoBoolToDBool(value))
}

func Chromium_GetDefaultUrl(obj uintptr) string {
	ret, _, _ := getLazyProc("Chromium_GetDefaultUrl").Call(obj)
	return DStrToGoStr(ret)
}

func Chromium_SetDefaultUrl(obj uintptr, value string) {
	_, _, _ = getLazyProc("Chromium_SetDefaultUrl").Call(obj, GoStrToDStr(value))
}

func Chromium_GetDefaultEncoding(obj uintptr) string {
	ret, _, _ := getLazyProc("Chromium_GetDefaultEncoding").Call(obj)
	return DStrToGoStr(ret)
}

func Chromium_SetDefaultEncoding(obj uintptr, value string) {
	_, _, _ = getLazyProc("Chromium_SetDefaultEncoding").Call(obj, GoStrToDStr(value))
}

func Chromium_GetAcceptLanguageList(obj uintptr) string {
	ret, _, _ := getLazyProc("Chromium_GetAcceptLanguageList").Call(obj)
	return DStrToGoStr(ret)
}

func Chromium_SetAcceptLanguageList(obj uintptr, value string) {
	_, _, _ = getLazyProc("Chromium_SetAcceptLanguageList").Call(obj, GoStrToDStr(value))
}

func Chromium_GetZoomLevel(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chromium_GetZoomLevel").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chromium_SetZoomLevel(obj uintptr, value float64) {
	_, _, _ = getLazyProc("Chromium_SetZoomLevel").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Chromium_GetZoomPct(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chromium_GetZoomPct").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chromium_SetZoomPct(obj uintptr, value float64) {
	_, _, _ = getLazyProc("Chromium_SetZoomPct").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Chromium_GetZoomStep(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chromium_GetZoomStep").Call(obj)
	return int32(ret)
}

func Chromium_SetZoomStep(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chromium_SetZoomStep").Call(obj, uintptr(value))
}

func Chromium_GetOptions(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chromium_GetOptions").Call(obj)
	return ret
}

func Chromium_SetOptions(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chromium_SetOptions").Call(obj, value)
}

func Chromium_GetFontOptions(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chromium_GetFontOptions").Call(obj)
	return ret
}

func Chromium_SetFontOptions(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chromium_SetFontOptions").Call(obj, value)
}

func Chromium_GetPDFPrintOptions(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chromium_GetPDFPrintOptions").Call(obj)
	return ret
}

func Chromium_SetPDFPrintOptions(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chromium_SetPDFPrintOptions").Call(obj, value)
}

func Chromium_GetAcceptCookies(obj uintptr) TCefCookiePref {
	ret, _, _ := getLazyProc("Chromium_GetAcceptCookies").Call(obj)
	return TCefCookiePref(ret)
}

func Chromium_SetAcceptCookies(obj uintptr, value TCefCookiePref) {
	_, _, _ = getLazyProc("Chromium_SetAcceptCookies").Call(obj, uintptr(value))
}

func Chromium_SetOnPrefsUpdated(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnPrefsUpdated").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnCookiesFlushed(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnCookiesFlushed").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnCertificateExceptionsCleared(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnCertificateExceptionsCleared").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnHttpAuthCredentialsCleared(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnHttpAuthCredentialsCleared").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnAllConnectionsClosed(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnAllConnectionsClosed").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnCookieSet(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnCookieSet").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnPdfPrintFinished(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnPdfPrintFinished").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnPrefsAvailable(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnPrefsAvailable").Call(obj, addEventToMap(obj, fn))
}

func Chromium_SetOnCookiesDeleted(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chromium_SetOnCookiesDeleted").Call(obj, addEventToMap(obj, fn))
}
