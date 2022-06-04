package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

type TChromium struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewChromium(owner IComponent) *TChromium {
	a := new(TChromium)
	a.instance = Chromium_Create(CheckPtr(owner))
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsChromium(obj any) *TChromium {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChromium{instance: instance, ptr: ptr}
}

// ApplicationFromInst 新建一个对象来自已经存在的对象实例指针。
func ChromiumFromInst(inst uintptr) *TChromium {
	return AsChromium(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func ChromiumFromObj(obj IObject) *TChromium {
	return AsChromium(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func ChromiumFromUnsafePointer(ptr unsafe.Pointer) *TChromium {
	return AsChromium(ptr)
}

// Free 释放对象。
func (c *TChromium) Free() {
	if c.instance != 0 {
		Chromium_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TChromium) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TChromium) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TChromium) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TChromium) Is() TIs {
	return TIs(c.instance)
}

func TChromiumClass() TClass {
	return Chromium_StaticClassType()
}

func (c *TChromium) ShowDevTools(inspectElementAt TPoint) {
	Chromium_ShowDevTools(c.instance, inspectElementAt)
}

func (c *TChromium) CloseDevTools() {
	Chromium_CloseDevTools(c.instance)
}

func (c *TChromium) MoveFormTo(x, y int32) {
	Chromium_MoveFormTo(c.instance, x, y)
}

func (c *TChromium) MoveFormBy(x, y int32) {
	Chromium_MoveFormBy(c.instance, x, y)
}

func (c *TChromium) ResizeFormWidthTo(x int32) {
	Chromium_ResizeFormWidthTo(c.instance, x)
}

func (c *TChromium) ResizeFormHeightTo(y int32) {
	Chromium_ResizeFormHeightTo(c.instance, y)
}

func (c *TChromium) SetFormLeftTo(x int32) {
	Chromium_SetFormLeftTo(c.instance, x)
}

func (c *TChromium) SetFormTopTo(y int32) {
	Chromium_SetFormTopTo(c.instance, y)
}

func (c *TChromium) SaveAsBitmapStream(stream IStream) bool {
	return Chromium_SaveAsBitmapStream(c.instance, CheckPtr(stream))
}

func (c *TChromium) TakeSnapshot(bitmap *TBitmap) bool {
	return Chromium_TakeSnapshot(c.instance, CheckPtr(bitmap))
}

func (c *TChromium) CloseBrowser(force bool) {
	Chromium_CloseBrowser(c.instance, force)
}

func (c *TChromium) CloseAllBrowsers() {
	Chromium_CloseAllBrowsers(c.instance)
}

func (c *TChromium) TryCloseBrowser() bool {
	return Chromium_TryCloseBrowser(c.instance)
}

func (c *TChromium) LoadURL(url string) {
	Chromium_LoadURL(c.instance, url)
}

func (c *TChromium) LoadString(html string) {
	Chromium_LoadString(c.instance, html)
}

func (c *TChromium) GoBack() {
	Chromium_GoBack(c.instance)
}

func (c *TChromium) GoForward() {
	Chromium_GoForward(c.instance)
}

func (c *TChromium) Reload() {
	Chromium_Reload(c.instance)
}

func (c *TChromium) ReloadIgnoreCache() {
	Chromium_ReloadIgnoreCache(c.instance)
}

func (c *TChromium) StopLoad() {
	Chromium_StopLoad(c.instance)
}

func (c *TChromium) StartDownload(url string) {
	Chromium_StartDownload(c.instance, url)
}

func (c *TChromium) DownloadImage(imageUrl string, isFav bool, maxImqgeSize Cardinal, bypassCache bool) {
	Chromium_DownloadImage(c.instance, imageUrl, isFav, maxImqgeSize, bypassCache)
}

func (c *TChromium) ClearCertificateExceptions(aClearImmediately bool) bool {
	return Chromium_ClearCertificateExceptions(c.instance, aClearImmediately)
}

func (c *TChromium) ClearHttpAuthCredentials(aClearImmediately bool) bool {
	return Chromium_ClearHttpAuthCredentials(c.instance, aClearImmediately)
}

func (c *TChromium) CloseAllConnections(aCloseImmediately bool) bool {
	return Chromium_CloseAllConnections(c.instance, aCloseImmediately)
}

func (c *TChromium) ExecuteJavaScript(code string, scriptUrl string) {
	Chromium_ExecuteJavaScript(c.instance, code, scriptUrl)
}

func (c *TChromium) UpdatePreferences() {
	Chromium_UpdatePreferences(c.instance)
}

func (c *TChromium) SavePreferences(filename string) {
	Chromium_SavePreferences(c.instance, filename)
}

func (c *TChromium) ResolveHost(url string) {
	Chromium_ResolveHost(c.instance, url)
}

func (c *TChromium) ClearCache() {
	Chromium_ClearCache(c.instance)
}

func (c *TChromium) DeleteCookies(url string, cookieName string, aDeleteImmediately bool) bool {
	return Chromium_DeleteCookies(c.instance, url, cookieName, aDeleteImmediately)
}

func (c *TChromium) SetCookie(url, name, value, domain, path string, secure, httpOnly, hasExpires bool, creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, prioroty TCefCookiePriority, setImmediately bool) bool {
	return Chromium_SetCookie(c.instance, url, name, value, domain, path, secure, httpOnly, hasExpires, creation, lastAccess, expires,
		sameSite, prioroty, setImmediately)
}

func (c *TChromium) FlushCookieStore(aFlushImmediately bool) bool {
	return Chromium_FlushCookieStore(c.instance, aFlushImmediately)
}

func (c *TChromium) SendDevToolsMessage(message string) bool {
	return Chromium_SendDevToolsMessage(c.instance, message)
}

func (c *TChromium) Find(text string, forward, matchCase, findNext bool) {
	Chromium_Find(c.instance, text, forward, matchCase, findNext)
}

func (c *TChromium) StopFinding(clearSelection bool) {
	Chromium_StopFinding(c.instance, clearSelection)
}

func (c *TChromium) Print() {
	Chromium_Print(c.instance)
}

func (c *TChromium) PrintToPDF(filepath string, title string, url string) {
	Chromium_PrintToPDF(c.instance, filepath, title, url)
}

func (c *TChromium) ClipboardCopy() {
	Chromium_ClipboardCopy(c.instance)
}

func (c *TChromium) ClipboardPaste() {
	Chromium_ClipboardPaste(c.instance)
}

func (c *TChromium) ClipboardCut() {
	Chromium_ClipboardCut(c.instance)
}

func (c *TChromium) ClipboardUndo() {
	Chromium_ClipboardUndo(c.instance)
}

func (c *TChromium) ClipboardRedo() {
	Chromium_ClipboardRedo(c.instance)
}

func (c *TChromium) ClipboardDel() {
	Chromium_ClipboardDel(c.instance)
}

func (c *TChromium) SelectAll() {
	Chromium_SelectAll(c.instance)
}

func (c *TChromium) IncZoomStep() {
	Chromium_IncZoomStep(c.instance)
}

func (c *TChromium) DecZoomStep() {
	Chromium_DecZoomStep(c.instance)
}

func (c *TChromium) IncZoomPct() {
	Chromium_IncZoomPct(c.instance)
}

func (c *TChromium) DecZoomPct() {
	Chromium_DecZoomPct(c.instance)
}

func (c *TChromium) ResetZoomStep() {
	Chromium_ResetZoomStep(c.instance)
}

func (c *TChromium) ResetZoomLevel() {
	Chromium_ResetZoomLevel(c.instance)
}

func (c *TChromium) ResetZoomPct() {
	Chromium_ResetZoomPct(c.instance)
}

func (c *TChromium) ReadZoom() {
	Chromium_ReadZoom(c.instance)
}

func (c *TChromium) MultithreadApp() bool {
	return Chromium_GetMultithreadApp(c.instance)
}

func (c *TChromium) IsLoading() bool {
	return Chromium_GetIsLoading(c.instance)
}

func (c *TChromium) HasDocument() bool {
	return Chromium_GetHasDocument(c.instance)
}

func (c *TChromium) HasView() bool {
	return Chromium_GetHasView(c.instance)
}

func (c *TChromium) HasDevTools() bool {
	return Chromium_GetHasDevTools(c.instance)
}

func (c *TChromium) CanGoBack() bool {
	return Chromium_GetCanGoBack(c.instance)
}

func (c *TChromium) CanGoForward() bool {
	return Chromium_GetCanGoForward(c.instance)
}

func (c *TChromium) IsPopUp() bool {
	return Chromium_GetIsPopUp(c.instance)
}

func (c *TChromium) Initialized() bool {
	return Chromium_GetInitialized(c.instance)
}

func (c *TChromium) DocumentURL() string {
	return Chromium_GetDocumentURL(c.instance)
}

func (c *TChromium) DoNotTrack() bool {
	return Chromium_GetDoNotTrack(c.instance)
}

func (c *TChromium) SetDoNotTrack(value bool) {
	Chromium_SetDoNotTrack(c.instance, value)
}

func (c *TChromium) Block3rdPartyCookies() bool {
	return Chromium_GetBlock3rdPartyCookies(c.instance)
}

func (c *TChromium) SetBlock3rdPartyCookies(value bool) {
	Chromium_SetBlock3rdPartyCookies(c.instance, value)
}

func (c *TChromium) JavascriptEnabled() bool {
	return Chromium_GetJavascriptEnabled(c.instance)
}

func (c *TChromium) SetJavascriptEnabled(value bool) {
	Chromium_SetJavascriptEnabled(c.instance, value)
}

func (c *TChromium) LoadImagesAutomatically() bool {
	return Chromium_GetLoadImagesAutomatically(c.instance)
}

func (c *TChromium) SetLoadImagesAutomatically(value bool) {
	Chromium_SetLoadImagesAutomatically(c.instance, value)
}

func (c *TChromium) DefaultUrl() string {
	return Chromium_GetDefaultUrl(c.instance)
}

func (c *TChromium) SetDefaultUrl(value string) {
	Chromium_SetDefaultUrl(c.instance, value)
}

func (c *TChromium) DefaultEncoding() string {
	return Chromium_GetDefaultEncoding(c.instance)
}

func (c *TChromium) SetDefaultEncoding(value string) {
	Chromium_SetDefaultEncoding(c.instance, value)
}

func (c *TChromium) AcceptLanguageList() string {
	return Chromium_GetAcceptLanguageList(c.instance)
}

func (c *TChromium) SetAcceptLanguageList(value string) {
	Chromium_SetAcceptLanguageList(c.instance, value)
}

func (c *TChromium) ZoomLevel() float64 {
	return Chromium_GetZoomLevel(c.instance)
}

func (c *TChromium) SetZoomLevel(value float64) {
	Chromium_SetZoomLevel(c.instance, value)
}

func (c *TChromium) ZoomPct() float64 {
	return Chromium_GetZoomPct(c.instance)
}

func (c *TChromium) SetZoomPct(value float64) {
	Chromium_SetZoomPct(c.instance, value)
}

func (c *TChromium) ZoomStep() int32 {
	return Chromium_GetZoomStep(c.instance)
}

func (c *TChromium) SetZoomStep(value int32) {
	Chromium_SetZoomStep(c.instance, value)
}

func (c *TChromium) Options() *TChromiumOptions {
	return AsChromiumOptions(Chromium_GetOptions(c.instance))
}

func (c *TChromium) SetOptions(value *TChromiumOptions) {
	Chromium_SetOptions(c.instance, CheckPtr(value))
}

func (c *TChromium) FontOptions() *TChromiumFontOptions {
	return AsChromiumFontOptions(Chromium_GetFontOptions(c.instance))
}

func (c *TChromium) SetFontOptions(value *TChromiumFontOptions) {
	Chromium_SetFontOptions(c.instance, CheckPtr(value))
}

func (c *TChromium) PDFPrintOptions() *TPDFPrintOptions {
	return AsPDFPrintOptions(Chromium_GetPDFPrintOptions(c.instance))
}

func (c *TChromium) SetPDFPrintOptions(value *TPDFPrintOptions) {
	Chromium_SetPDFPrintOptions(c.instance, CheckPtr(value))
}

func (c *TChromium) AcceptCookies() TCefCookiePref {
	return Chromium_GetAcceptCookies(c.instance)
}

func (c *TChromium) SetAcceptCookies(value TCefCookiePref) {
	Chromium_SetAcceptCookies(c.instance, value)
}

func (c *TChromium) SetOnPrefsUpdated(event TNotifyEvent) {
	Chromium_SetOnPrefsUpdated(c.instance, event)
}

func (c *TChromium) SetOnCookiesFlushed(event TNotifyEvent) {
	Chromium_SetOnCookiesFlushed(c.instance, event)
}

func (c *TChromium) SetOnCertificateExceptionsCleared(event TNotifyEvent) {
	Chromium_SetOnCertificateExceptionsCleared(c.instance, event)
}

func (c *TChromium) SetOnHttpAuthCredentialsCleared(event TNotifyEvent) {
	Chromium_SetOnHttpAuthCredentialsCleared(c.instance, event)
}

func (c *TChromium) SetOnAllConnectionsClosed(event TNotifyEvent) {
	Chromium_SetOnAllConnectionsClosed(c.instance, event)
}

func (c *TChromium) SetOnCookieSet(event TOnCookieSet) {
	Chromium_SetOnCookieSet(c.instance, event)
}

func (c *TChromium) SetOnPdfPrintFinished(event TOnOKNotifyEvent) {
	Chromium_SetOnPdfPrintFinished(c.instance, event)
}

func (c *TChromium) SetOnPrefsAvailable(event TOnOKNotifyEvent) {
	Chromium_SetOnPrefsAvailable(c.instance, event)
}

func (c *TChromium) SetOnCookiesDeleted(event TCheckItemChange) {
	Chromium_SetOnCookiesDeleted(c.instance, event)
}
