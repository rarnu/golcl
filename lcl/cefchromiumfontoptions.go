package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TChromiumFontOptions struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewChromiumFontOptions() *TChromiumFontOptions {
	a := new(TChromiumFontOptions)
	a.instance = ChromiumFontOptions_Create()
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsChromiumFontOptions(obj any) *TChromiumFontOptions {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChromiumFontOptions{instance: instance, ptr: ptr}
}

func ChromiumFontOptionsFromInst(inst uintptr) *TChromiumFontOptions {
	return AsChromiumFontOptions(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func ChromiumFontOptionsFromObj(obj IObject) *TChromiumFontOptions {
	return AsChromiumFontOptions(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func ChromiumFontOptionsFromUnsafePointer(ptr unsafe.Pointer) *TChromiumFontOptions {
	return AsChromiumFontOptions(ptr)
}

// Free 释放对象。
func (c *TChromiumFontOptions) Free() {
	if c.instance != 0 {
		ChromiumFontOptions_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TChromiumFontOptions) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TChromiumFontOptions) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TChromiumFontOptions) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TChromiumFontOptions) Is() TIs {
	return TIs(c.instance)
}

func TChromiumFontOptionsClass() TClass {
	return ChromiumFontOptions_StaticClassType()
}

func (c *TChromiumFontOptions) StandardFontFamily() string {
	return ChromiumFontOptions_GetStandardFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetStandardFontFamily(value string) {
	ChromiumFontOptions_SetStandardFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) FixedFontFamily() string {
	return ChromiumFontOptions_GetFixedFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetFixedFontFamily(value string) {
	ChromiumFontOptions_SetFixedFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) SerifFontFamily() string {
	return ChromiumFontOptions_GetSerifFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetSerifFontFamily(value string) {
	ChromiumFontOptions_SetSerifFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) SansSerifFontFamily() string {
	return ChromiumFontOptions_GetSansSerifFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetSansSerifFontFamily(value string) {
	ChromiumFontOptions_SetSansSerifFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) CursiveFontFamily() string {
	return ChromiumFontOptions_GetCursiveFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetCursiveFontFamily(value string) {
	ChromiumFontOptions_SetCursiveFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) FantasyFontFamily() string {
	return ChromiumFontOptions_GetFantasyFontFamily(c.instance)
}

func (c *TChromiumFontOptions) SetFantasyFontFamily(value string) {
	ChromiumFontOptions_SetFantasyFontFamily(c.instance, value)
}

func (c *TChromiumFontOptions) DefaultFontSize() int32 {
	return ChromiumFontOptions_GetDefaultFontSize(c.instance)
}

func (c *TChromiumFontOptions) SetDefaultFontSize(value int32) {
	ChromiumFontOptions_SetDefaultFontSize(c.instance, value)
}

func (c *TChromiumFontOptions) DefaultFixedFontSize() int32 {
	return ChromiumFontOptions_GetDefaultFixedFontSize(c.instance)
}

func (c *TChromiumFontOptions) SetDefaultFixedFontSize(value int32) {
	ChromiumFontOptions_SetDefaultFixedFontSize(c.instance, value)
}

func (c *TChromiumFontOptions) MinimumFontSize() int32 {
	return ChromiumFontOptions_GetMinimumFontSize(c.instance)
}

func (c *TChromiumFontOptions) SetMinimumFontSize(value int32) {
	ChromiumFontOptions_SetMinimumFontSize(c.instance, value)
}

func (c *TChromiumFontOptions) MinimumLogicalFontSize() int32 {
	return ChromiumFontOptions_GetMinimumLogicalFontSize(c.instance)
}

func (c *TChromiumFontOptions) SetMinimumLogicalFontSize(value int32) {
	ChromiumFontOptions_SetMinimumLogicalFontSize(c.instance, value)
}

func (c *TChromiumFontOptions) RemoteFonts() TCefState {
	return ChromiumFontOptions_GetRemoteFonts(c.instance)
}

func (c *TChromiumFontOptions) SetRemoteFonts(value TCefState) {
	ChromiumFontOptions_SetRemoteFonts(c.instance, value)
}
