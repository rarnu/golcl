package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TPDFPrintOptions struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func NewPDFPrintOptions() *TPDFPrintOptions {
	a := new(TPDFPrintOptions)
	a.instance = PDFPrintOptions_Create()
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

func AsPDFPrintOptions(obj any) *TPDFPrintOptions {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPDFPrintOptions{instance: instance, ptr: ptr}
}

func PDFPrintOptionsFromInst(inst uintptr) *TPDFPrintOptions {
	return AsPDFPrintOptions(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func PDFPrintOptionsFromObj(obj IObject) *TPDFPrintOptions {
	return AsPDFPrintOptions(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func PDFPrintOptionsFromUnsafePointer(ptr unsafe.Pointer) *TPDFPrintOptions {
	return AsPDFPrintOptions(ptr)
}

// Free 释放对象。
func (c *TPDFPrintOptions) Free() {
	if c.instance != 0 {
		PDFPrintOptions_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (c *TPDFPrintOptions) Instance() uintptr {
	return c.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (c *TPDFPrintOptions) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// IsValid 检测地址是否为空。
func (c *TPDFPrintOptions) IsValid() bool {
	return c.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (c *TPDFPrintOptions) Is() TIs {
	return TIs(c.instance)
}

func TPDFPrintOptionsClass() TClass {
	return PDFPrintOptions_StaticClassType()
}

func (c *TPDFPrintOptions) PageWidth() int32 {
	return PDFPrintOptions_GetPageWidth(c.instance)
}

func (c *TPDFPrintOptions) SetPageWidth(value int32) {
	PDFPrintOptions_SetPageWidth(c.instance, value)
}

func (c *TPDFPrintOptions) PageHeight() int32 {
	return PDFPrintOptions_GetPageHeight(c.instance)
}

func (c *TPDFPrintOptions) SetPageHeight(value int32) {
	PDFPrintOptions_SetPageHeight(c.instance, value)
}

func (c *TPDFPrintOptions) ScaleFactor() int32 {
	return PDFPrintOptions_GetScaleFactor(c.instance)
}

func (c *TPDFPrintOptions) SetScaleFactor(value int32) {
	PDFPrintOptions_SetScaleFactor(c.instance, value)
}

func (c *TPDFPrintOptions) MarginTop() int32 {
	return PDFPrintOptions_GetMarginTop(c.instance)
}

func (c *TPDFPrintOptions) SetMarginTop(value int32) {
	PDFPrintOptions_SetMarginTop(c.instance, value)
}

func (c *TPDFPrintOptions) MarginBottom() int32 {
	return PDFPrintOptions_GetMarginBottom(c.instance)
}

func (c *TPDFPrintOptions) SetMarginBottom(value int32) {
	PDFPrintOptions_SetMarginBottom(c.instance, value)
}

func (c *TPDFPrintOptions) MarginLeft() int32 {
	return PDFPrintOptions_GetMarginLeft(c.instance)
}

func (c *TPDFPrintOptions) SetMarginLeft(value int32) {
	PDFPrintOptions_SetMarginLeft(c.instance, value)
}

func (c *TPDFPrintOptions) MarginRight() int32 {
	return PDFPrintOptions_GetMarginRight(c.instance)
}

func (c *TPDFPrintOptions) SetMarginRight(value int32) {
	PDFPrintOptions_SetMarginRight(c.instance, value)
}

func (c *TPDFPrintOptions) MarginType() TCefPdfPrintMarginType {
	return PDFPrintOptions_GetMarginType(c.instance)
}

func (c *TPDFPrintOptions) SetMarginType(value TCefPdfPrintMarginType) {
	PDFPrintOptions_SetMarginType(c.instance, value)
}

func (c *TPDFPrintOptions) HeaderFooterEnabled() bool {
	return PDFPrintOptions_GetHeaderFooterEnabled(c.instance)
}

func (c *TPDFPrintOptions) SetHeaderFooterEnabled(value bool) {
	PDFPrintOptions_SetHeaderFooterEnabled(c.instance, value)
}

func (c *TPDFPrintOptions) SelectionOnly() bool {
	return PDFPrintOptions_GetSelectionOnly(c.instance)
}

func (c *TPDFPrintOptions) SetSelectionOnly(value bool) {
	PDFPrintOptions_SetSelectionOnly(c.instance, value)
}

func (c *TPDFPrintOptions) Landscape() bool {
	return PDFPrintOptions_GetLandscape(c.instance)
}

func (c *TPDFPrintOptions) SetLandscape(value bool) {
	PDFPrintOptions_SetLandscape(c.instance, value)
}

func (c *TPDFPrintOptions) BackgroundsEnabled() bool {
	return PDFPrintOptions_GetBackgroundsEnabled(c.instance)
}

func (c *TPDFPrintOptions) SetBackgroundsEnabled(value bool) {
	PDFPrintOptions_SetBackgroundsEnabled(c.instance, value)
}
