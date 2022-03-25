//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TBitmap struct {
	IGraphic
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewBitmap() *TBitmap {
	b := new(TBitmap)
	b.instance = Bitmap_Create()
	b.ptr = unsafe.Pointer(b.instance)
	setFinalizer(b, (*TBitmap).Free)
	return b
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBitmap(obj any) *TBitmap {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TBitmap{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsBitmap.
func BitmapFromInst(inst uintptr) *TBitmap {
	return AsBitmap(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsBitmap.
func BitmapFromObj(obj IObject) *TBitmap {
	return AsBitmap(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBitmap.
func BitmapFromUnsafePointer(ptr unsafe.Pointer) *TBitmap {
	return AsBitmap(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (b *TBitmap) Free() {
	if b.instance != 0 {
		Bitmap_Free(b.instance)
		b.instance, b.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TBitmap) Instance() uintptr {
	return b.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TBitmap) UnsafeAddr() unsafe.Pointer {
	return b.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TBitmap) IsValid() bool {
	return b.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TBitmap) Is() TIs {
	return TIs(b.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TBitmap) As() TAs {
//    return TAs(b.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TBitmapClass() TClass {
	return Bitmap_StaticClassType()
}

// 从设备驱动中加载Bitmap。
//
// Load the Bitmap from the device driver.
func (b *TBitmap) LoadFromDevice(ADc HDC) {
	Bitmap_LoadFromDevice(b.instance, ADc)
}

// 用于ScanLine属性，aStreamIsValid 默认为 false。
//
// Used for ScanLine property, aStreamIsValid defaults to false.
func (b *TBitmap) EndUpdate(AStreamIsValid bool) {
	Bitmap_EndUpdate(b.instance, AStreamIsValid)
}

// 用于ScanLine属性，aCanvasOnly 默认为 false。
//
// Used for ScanLine properties, aCanvasOnly defaults to false.
func (b *TBitmap) BeginUpdate(ACanvasOnly bool) {
	Bitmap_BeginUpdate(b.instance, ACanvasOnly)
}

// 清除bitmap数据。
//
// Clear bitmap data.
func (b *TBitmap) Clear() {
	Bitmap_Clear(b.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBitmap) Assign(Source IObject) {
	Bitmap_Assign(b.instance, CheckPtr(Source))
}

func (b *TBitmap) FreeImage() {
	Bitmap_FreeImage(b.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (b *TBitmap) HandleAllocated() bool {
	return Bitmap_HandleAllocated(b.instance)
}

// 文件流加载。
func (b *TBitmap) LoadFromStream(Stream IStream) {
	Bitmap_LoadFromStream(b.instance, CheckPtr(Stream))
}

// 保存至流。
func (b *TBitmap) SaveToStream(Stream IStream) {
	Bitmap_SaveToStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
	Bitmap_SetSize(b.instance, AWidth, AHeight)
}

func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
	Bitmap_LoadFromResourceName(b.instance, Instance, ResName)
}

func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
	Bitmap_LoadFromResourceID(b.instance, Instance, ResID)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBitmap) Equals(Obj IObject) bool {
	return Bitmap_Equals(b.instance, CheckPtr(Obj))
}

// 从文件加载。
func (b *TBitmap) LoadFromFile(Filename string) {
	Bitmap_LoadFromFile(b.instance, Filename)
}

// 保存至文件。
func (b *TBitmap) SaveToFile(Filename string) {
	Bitmap_SaveToFile(b.instance, Filename)
}

// 获取类名路径。
//
// Get the class name path.
func (b *TBitmap) GetNamePath() string {
	return Bitmap_GetNamePath(b.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (b *TBitmap) ClassType() TClass {
	return Bitmap_ClassType(b.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBitmap) ClassName() string {
	return Bitmap_ClassName(b.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBitmap) InstanceSize() int32 {
	return Bitmap_InstanceSize(b.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBitmap) InheritsFrom(AClass TClass) bool {
	return Bitmap_InheritsFrom(b.instance, AClass)
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBitmap) GetHashCode() int32 {
	return Bitmap_GetHashCode(b.instance)
}

// 文本类信息。
//
// Text information.
func (b *TBitmap) ToString() string {
	return Bitmap_ToString(b.instance)
}

// 获取画布。
func (b *TBitmap) Canvas() *TCanvas {
	return AsCanvas(Bitmap_GetCanvas(b.instance))
}

// 获取控件句柄。
//
// Get Control handle.
func (b *TBitmap) Handle() HBITMAP {
	return Bitmap_GetHandle(b.instance)
}

// 设置控件句柄。
//
// Set Control handle.
func (b *TBitmap) SetHandle(value HBITMAP) {
	Bitmap_SetHandle(b.instance, value)
}

func (b *TBitmap) HandleType() TBitmapHandleType {
	return Bitmap_GetHandleType(b.instance)
}

func (b *TBitmap) SetHandleType(value TBitmapHandleType) {
	Bitmap_SetHandleType(b.instance, value)
}

func (b *TBitmap) MaskHandle() HBITMAP {
	return Bitmap_GetMaskHandle(b.instance)
}

func (b *TBitmap) SetMaskHandle(value HBITMAP) {
	Bitmap_SetMaskHandle(b.instance, value)
}

func (b *TBitmap) PixelFormat() TPixelFormat {
	return Bitmap_GetPixelFormat(b.instance)
}

func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
	Bitmap_SetPixelFormat(b.instance, value)
}

func (b *TBitmap) TransparentMode() TTransparentMode {
	return Bitmap_GetTransparentMode(b.instance)
}

func (b *TBitmap) SetTransparentMode(value TTransparentMode) {
	Bitmap_SetTransparentMode(b.instance, value)
}

func (b *TBitmap) Empty() bool {
	return Bitmap_GetEmpty(b.instance)
}

// 获取高度。
//
// Get height.
func (b *TBitmap) Height() int32 {
	return Bitmap_GetHeight(b.instance)
}

// 设置高度。
//
// Set height.
func (b *TBitmap) SetHeight(value int32) {
	Bitmap_SetHeight(b.instance, value)
}

// 获取修改。
//
// Get modified.
func (b *TBitmap) Modified() bool {
	return Bitmap_GetModified(b.instance)
}

// 设置修改。
//
// Set modified.
func (b *TBitmap) SetModified(value bool) {
	Bitmap_SetModified(b.instance, value)
}

func (b *TBitmap) Palette() HPALETTE {
	return Bitmap_GetPalette(b.instance)
}

func (b *TBitmap) SetPalette(value HPALETTE) {
	Bitmap_SetPalette(b.instance, value)
}

func (b *TBitmap) PaletteModified() bool {
	return Bitmap_GetPaletteModified(b.instance)
}

func (b *TBitmap) SetPaletteModified(value bool) {
	Bitmap_SetPaletteModified(b.instance, value)
}

// 获取透明。
//
// Get transparent.
func (b *TBitmap) Transparent() bool {
	return Bitmap_GetTransparent(b.instance)
}

// 设置透明。
//
// Set transparent.
func (b *TBitmap) SetTransparent(value bool) {
	Bitmap_SetTransparent(b.instance, value)
}

// 获取宽度。
//
// Get width.
func (b *TBitmap) Width() int32 {
	return Bitmap_GetWidth(b.instance)
}

// 设置宽度。
//
// Set width.
func (b *TBitmap) SetWidth(value int32) {
	Bitmap_SetWidth(b.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
	Bitmap_SetOnChange(b.instance, fn)
}

func (b *TBitmap) ScanLine(Row int32) uintptr {
	return Bitmap_GetScanLine(b.instance, Row)
}
