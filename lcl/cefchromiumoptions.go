package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TChromiumOptions struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewChromiumOptions() *TChromiumOptions {
	a := new(TChromiumOptions)
	a.instance = ChromiumOptions_Create()
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsChromiumOptions(obj any) *TChromiumOptions {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChromiumOptions{instance: instance, ptr: ptr}
}

func ChromiumOptionsFromInst(inst uintptr) *TChromiumOptions {
	return AsChromiumOptions(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func ChromiumOptionsFromObj(obj IObject) *TChromiumOptions {
	return AsChromiumOptions(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func ChromiumOptionsFromUnsafePointer(ptr unsafe.Pointer) *TChromiumOptions {
	return AsChromiumOptions(ptr)
}

// Free 释放对象。
func (c *TChromiumOptions) Free() {
	if c.instance != 0 {
		ChromiumOptions_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TChromiumOptions) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TChromiumOptions) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TChromiumOptions) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TChromiumOptions) Is() TIs {
	return TIs(c.instance)
}

func TChromiumOptionsClass() TClass {
	return ChromiumOptions_StaticClassType()
}

func (c *TChromiumOptions) Javascript() TCefState {
	return ChromiumOptions_GetJavascript(c.instance)
}

func (c *TChromiumOptions) SetJavascript(value TCefState) {
	ChromiumOptions_SetJavascript(c.instance, value)
}

func (c *TChromiumOptions) JavascriptCloseWindows() TCefState {
	return ChromiumOptions_GetJavascriptCloseWindows(c.instance)
}

func (c *TChromiumOptions) SetJavascriptCloseWindows(value TCefState) {
	ChromiumOptions_SetJavascriptCloseWindows(c.instance, value)
}

func (c *TChromiumOptions) JavascriptAccessClipboard() TCefState {
	return ChromiumOptions_GetJavascriptAccessClipboard(c.instance)
}

func (c *TChromiumOptions) SetJavascriptAccessClipboard(value TCefState) {
	ChromiumOptions_SetJavascriptAccessClipboard(c.instance, value)
}

func (c *TChromiumOptions) JavascriptDomPaste() TCefState {
	return ChromiumOptions_GetJavascriptDomPaste(c.instance)
}

func (c *TChromiumOptions) SetJavascriptDomPaste(value TCefState) {
	ChromiumOptions_SetJavascriptDomPaste(c.instance, value)
}

func (c *TChromiumOptions) ImageLoading() TCefState {
	return ChromiumOptions_GetImageLoading(c.instance)
}

func (c *TChromiumOptions) SetImageLoading(value TCefState) {
	ChromiumOptions_SetImageLoading(c.instance, value)
}

func (c *TChromiumOptions) ImageShrinkStandaloneToFit() TCefState {
	return ChromiumOptions_GetImageShrinkStandaloneToFit(c.instance)
}

func (c *TChromiumOptions) SetImageShrinkStandaloneToFit(value TCefState) {
	ChromiumOptions_SetImageShrinkStandaloneToFit(c.instance, value)
}

func (c *TChromiumOptions) TextAreaResize() TCefState {
	return ChromiumOptions_GetTextAreaResize(c.instance)
}

func (c *TChromiumOptions) SetTextAreaResize(value TCefState) {
	ChromiumOptions_SetTextAreaResize(c.instance, value)
}

func (c *TChromiumOptions) TabToLinks() TCefState {
	return ChromiumOptions_GetTabToLinks(c.instance)
}

func (c *TChromiumOptions) SetTabToLinks(value TCefState) {
	ChromiumOptions_SetTabToLinks(c.instance, value)
}

func (c *TChromiumOptions) LocalStorage() TCefState {
	return ChromiumOptions_GetLocalStorage(c.instance)
}

func (c *TChromiumOptions) SetLocalStorage(value TCefState) {
	ChromiumOptions_SetLocalStorage(c.instance, value)
}

func (c *TChromiumOptions) Databases() TCefState {
	return ChromiumOptions_GetDatabases(c.instance)
}

func (c *TChromiumOptions) SetDatabases(value TCefState) {
	ChromiumOptions_SetDatabases(c.instance, value)
}

func (c *TChromiumOptions) Webgl() TCefState {
	return ChromiumOptions_GetWebgl(c.instance)
}

func (c *TChromiumOptions) SetWebgl(value TCefState) {
	ChromiumOptions_SetWebgl(c.instance, value)
}

func (c *TChromiumOptions) BackgroundColor() Cardinal {
	return ChromiumOptions_GetBackgroundColor(c.instance)
}

func (c *TChromiumOptions) SetBackgroundColor(value Cardinal) {
	ChromiumOptions_SetBackgroundColor(c.instance, value)
}

func (c *TChromiumOptions) AcceptLanguageList() string {
	return ChromiumOptions_GetAcceptLanguageList(c.instance)
}

func (c *TChromiumOptions) SetAcceptLanguageList(value string) {
	ChromiumOptions_SetAcceptLanguageList(c.instance, value)
}

func (c *TChromiumOptions) WindowlessFrameRate() int32 {
	return ChromiumOptions_GetWindowlessFrameRate(c.instance)
}

func (c *TChromiumOptions) SetWindowlessFrameRate(value int32) {
	ChromiumOptions_SetWindowlessFrameRate(c.instance, value)
}

func (c *TChromiumOptions) ChromeStatusBubble() TCefState {
	return ChromiumOptions_GetChromeStatusBubble(c.instance)
}

func (c *TChromiumOptions) SetChromeStatusBubble(value TCefState) {
	ChromiumOptions_SetChromeStatusBubble(c.instance, value)
}
