package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TChromiumWindow struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewChromiumWindow(owner IComponent) *TChromiumWindow {
	a := new(TChromiumWindow)
	a.instance = ChromiumWindow_Create(CheckPtr(owner))
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsChromiumWindow(obj any) *TChromiumWindow {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChromiumWindow{instance: instance, ptr: ptr}
}

// ApplicationFromInst 新建一个对象来自已经存在的对象实例指针。
func ChromiumWindowFromInst(inst uintptr) *TChromiumWindow {
	return AsChromiumWindow(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func ChromiumWindowFromObj(obj IObject) *TChromiumWindow {
	return AsChromiumWindow(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func ChromiumWindowFromUnsafePointer(ptr unsafe.Pointer) *TChromiumWindow {
	return AsChromiumWindow(ptr)
}

// Free 释放对象。
func (c *TChromiumWindow) Free() {
	if c.instance != 0 {
		ChromiumWindow_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TChromiumWindow) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TChromiumWindow) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TChromiumWindow) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TChromiumWindow) Is() TIs {
	return TIs(c.instance)
}

func TChromiumWindowClass() TClass {
	return ChromiumWindow_StaticClassType()
}

func (c *TChromiumWindow) Parent() *TWinControl {
	return AsWinControl(ChromiumWindow_GetParent(c.instance))
}

func (c *TChromiumWindow) SetParent(value IWinControl) {
	ChromiumWindow_SetParent(c.instance, CheckPtr(value))
}

func (c *TChromiumWindow) Align() TAlign {
	return ChromiumWindow_GetAlign(c.instance)
}

func (c *TChromiumWindow) SetAlign(value TAlign) {
	ChromiumWindow_SetAlign(c.instance, value)
}

func (c *TChromiumWindow) Anchors() TAnchors {
	return ChromiumWindow_GetAnchors(c.instance)
}

func (c *TChromiumWindow) SetAnchors(value TAnchors) {
	ChromiumWindow_SetAnchors(c.instance, value)
}

func (c *TChromiumWindow) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ChromiumWindow_GetConstraints(c.instance))
}

func (c *TChromiumWindow) SetConstraints(value *TSizeConstraints) {
	ChromiumWindow_SetConstraints(c.instance, CheckPtr(value))
}

func (c *TChromiumWindow) DoubleBuffered() bool {
	return ChromiumWindow_GetDoubleBuffered(c.instance)
}

func (c *TChromiumWindow) SetDoubleBuffered(value bool) {
	ChromiumWindow_SetDoubleBuffered(c.instance, value)
}

func (c *TChromiumWindow) Enabled() bool {
	return ChromiumWindow_GetEnabled(c.instance)
}

func (c *TChromiumWindow) SetEnabled(value bool) {
	ChromiumWindow_SetEnabled(c.instance, value)
}

func (c *TChromiumWindow) ShowHint() bool {
	return ChromiumWindow_GetShowHint(c.instance)
}

func (c *TChromiumWindow) SetShowHint(value bool) {
	ChromiumWindow_SetShowHint(c.instance, value)
}

func (c *TChromiumWindow) TabOrder() TTabOrder {
	return ChromiumWindow_GetTabOrder(c.instance)
}

func (c *TChromiumWindow) SetTabOrder(value TTabOrder) {
	ChromiumWindow_SetTabOrder(c.instance, value)
}

func (c *TChromiumWindow) TabStop() bool {
	return ChromiumWindow_GetTabStop(c.instance)
}

func (c *TChromiumWindow) SetTabStop(value bool) {
	ChromiumWindow_SetTabStop(c.instance, value)
}

func (c *TChromiumWindow) Visible() bool {
	return ChromiumWindow_GetVisible(c.instance)
}

func (c *TChromiumWindow) SetVisible(value bool) {
	ChromiumWindow_SetVisible(c.instance, value)
}

func (c *TChromiumWindow) Left() int32 {
	return ChromiumWindow_GetLeft(c.instance)
}

func (c *TChromiumWindow) SetLeft(value int32) {
	ChromiumWindow_SetLeft(c.instance, value)
}

func (c *TChromiumWindow) Top() int32 {
	return ChromiumWindow_GetTop(c.instance)
}

func (c *TChromiumWindow) SetTop(value int32) {
	ChromiumWindow_SetTop(c.instance, value)
}

func (c *TChromiumWindow) Width() int32 {
	return ChromiumWindow_GetWidth(c.instance)
}

func (c *TChromiumWindow) SetWidth(value int32) {
	ChromiumWindow_SetWidth(c.instance, value)
}

func (c *TChromiumWindow) Height() int32 {
	return ChromiumWindow_GetHeight(c.instance)
}

func (c *TChromiumWindow) SetHeight(value int32) {
	ChromiumWindow_SetHeight(c.instance, value)
}

func (c *TChromiumWindow) Tag() *TObject {
	return AsObject(ChromiumWindow_GetTag(c.instance))
}

func (c *TChromiumWindow) SetTag(value IObject) {
	ChromiumWindow_SetTag(c.instance, CheckPtr(value))
}

func (c *TChromiumWindow) Name() string {
	return ChromiumWindow_GetName(c.instance)
}

func (c *TChromiumWindow) SetName(value string) {
	ChromiumWindow_SetName(c.instance, value)
}

func (c *TChromiumWindow) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ChromiumWindow_GetBorderSpacing(c.instance))
}

func (c *TChromiumWindow) SetBorderSpacing(value *TControlBorderSpacing) {
	ChromiumWindow_SetBorderSpacing(c.instance, CheckPtr(value))
}

func (c *TChromiumWindow) Handle() HWND {
	return ChromiumWindow_GetHandle(c.instance)
}

func (c *TChromiumWindow) Owner() *TWinControl {
	return AsWinControl(ChromiumWindow_GetOwner(c.instance))
}

func (c *TChromiumWindow) Focused() bool {
	return ChromiumWindow_Focused(c.instance)
}

func (c *TChromiumWindow) Invalidate() {
	ChromiumWindow_Invalidate(c.instance)
}

func (c *TChromiumWindow) SetFocus() {
	ChromiumWindow_SetFocus(c.instance)
}

func (c *TChromiumWindow) Update() {
	ChromiumWindow_Update(c.instance)
}

func (c *TChromiumWindow) BringToFront() {
	ChromiumWindow_BringToFront(c.instance)
}

func (c *TChromiumWindow) Hide() {
	ChromiumWindow_Hide(c.instance)
}

func (c *TChromiumWindow) Show() {
	ChromiumWindow_Show(c.instance)
}

func (c *TChromiumWindow) Refresh() {
	ChromiumWindow_Refresh(c.instance)
}

func (c *TChromiumWindow) SendToBack() {
	ChromiumWindow_SendToBack(c.instance)
}

func (c *TChromiumWindow) SetOnAfterCreated(value TNotifyEvent) {
	ChromiumWindow_SetOnAfterCreated(c.instance, value)
}

func (c *TChromiumWindow) SetOnBeforeClose(value TNotifyEvent) {
	ChromiumWindow_SetOnBeforeClose(c.instance, value)
}

func (c *TChromiumWindow) SetOnClose(value TNotifyEvent) {
	ChromiumWindow_SetOnClose(c.instance, value)
}

func (c *TChromiumWindow) SetOnEnter(value TNotifyEvent) {
	ChromiumWindow_SetOnEnter(c.instance, value)
}

func (c *TChromiumWindow) SetOnExit(value TNotifyEvent) {
	ChromiumWindow_SetOnExit(c.instance, value)
}

func (c *TChromiumWindow) SetOnResize(value TNotifyEvent) {
	ChromiumWindow_SetOnResize(c.instance, value)
}

func (c *TChromiumWindow) CreateBrowser() bool {
	return ChromiumWindow_CreateBrowser(c.instance)
}

func (c *TChromiumWindow) CloseBrowser(force bool) {
	ChromiumWindow_CloseBrowser(c.instance, force)
}

func (c *TChromiumWindow) LoadURL(url string) {
	ChromiumWindow_LoadURL(c.instance, url)
}

func (c *TChromiumWindow) NotifyMoveOrResizeStarted() {
	ChromiumWindow_NotifyMoveOrResizeStarted(c.instance)
}

func (c *TChromiumWindow) ChromiumBrowser() *TChromium {
	return AsChromium(ChromiumWindow_GetChromiumBrowser(c.instance))
}

func (c *TChromiumWindow) Initialized() bool {
	return ChromiumWindow_GetInitialized(c.instance)
}

func (c *TChromiumWindow) UpdateSize() {
	ChromiumWindow_UpdateSize(c.instance)
}
