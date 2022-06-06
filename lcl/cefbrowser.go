package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TBrowser struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewBrowser(owner IWinControl, align TAlign) *TBrowser {
	a := new(TBrowser)
	a.instance = Browser_Create(CheckPtr(owner), align)
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsBrowser(obj any) *TBrowser {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TBrowser{instance: instance, ptr: ptr}
}

// ApplicationFromInst 新建一个对象来自已经存在的对象实例指针。
func BrowserFromInst(inst uintptr) *TBrowser {
	return AsBrowser(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func BrowserFromObj(obj IObject) *TBrowser {
	return AsBrowser(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func BrowserFromUnsafePointer(ptr unsafe.Pointer) *TBrowser {
	return AsBrowser(ptr)
}

// Free 释放对象。
func (c *TBrowser) Free() {
	if c.instance != 0 {
		Browser_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TBrowser) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TBrowser) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TBrowser) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TBrowser) Is() TIs {
	return TIs(c.instance)
}

func (c *TBrowser) ChromiumWindow() *TChromiumWindow {
	return AsChromiumWindow(Browser_GetChromiumWindow(c.instance))
}

func (c *TBrowser) Chromium() *TChromium {
	return AsChromium(Browser_GetChromium(c.instance))
}

func (c *TBrowser) OpenLinkPopup() bool {
	return Browser_GetOpenLinkPopup(c.instance)
}

func (c *TBrowser) SetOpenLinkPopup(value bool) {
	Browser_SetOpenLinkPopup(c.instance, value)
}

func (c *TBrowser) SetOnInitComplete(event TNotifyEvent) {
	Browser_SetOnInitComplete(c.instance, event)
}

func (c *TBrowser) SetOnAfterCreated(event TNotifyEvent) {
	Browser_SetOnAfterCreated(c.instance, event)
}

func (c *TBrowser) SetOnBeforeClose(event TNotifyEvent) {
	Browser_SetOnBeforeClose(c.instance, event)
}
