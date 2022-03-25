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

type TGraphic struct {
	IGraphic
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewGraphic() *TGraphic {
	g := new(TGraphic)
	g.instance = Graphic_Create()
	g.ptr = unsafe.Pointer(g.instance)
	setFinalizer(g, (*TGraphic).Free)
	return g
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsGraphic(obj any) *TGraphic {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TGraphic{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsGraphic.
func GraphicFromInst(inst uintptr) *TGraphic {
	return AsGraphic(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsGraphic.
func GraphicFromObj(obj IObject) *TGraphic {
	return AsGraphic(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGraphic.
func GraphicFromUnsafePointer(ptr unsafe.Pointer) *TGraphic {
	return AsGraphic(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (g *TGraphic) Free() {
	if g.instance != 0 {
		Graphic_Free(g.instance)
		g.instance, g.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (g *TGraphic) Instance() uintptr {
	return g.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (g *TGraphic) UnsafeAddr() unsafe.Pointer {
	return g.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (g *TGraphic) IsValid() bool {
	return g.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (g *TGraphic) Is() TIs {
	return TIs(g.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (g *TGraphic) As() TAs {
//    return TAs(g.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TGraphicClass() TClass {
	return Graphic_StaticClassType()
}

// 与一个对象进行比较。
//
// Compare with an object.
func (g *TGraphic) Equals(Obj IObject) bool {
	return Graphic_Equals(g.instance, CheckPtr(Obj))
}

// 从文件加载。
func (g *TGraphic) LoadFromFile(Filename string) {
	Graphic_LoadFromFile(g.instance, Filename)
}

// 保存至文件。
func (g *TGraphic) SaveToFile(Filename string) {
	Graphic_SaveToFile(g.instance, Filename)
}

// 文件流加载。
func (g *TGraphic) LoadFromStream(Stream IStream) {
	Graphic_LoadFromStream(g.instance, CheckPtr(Stream))
}

// 保存至流。
func (g *TGraphic) SaveToStream(Stream IStream) {
	Graphic_SaveToStream(g.instance, CheckPtr(Stream))
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (g *TGraphic) Assign(Source IObject) {
	Graphic_Assign(g.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (g *TGraphic) GetNamePath() string {
	return Graphic_GetNamePath(g.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (g *TGraphic) ClassType() TClass {
	return Graphic_ClassType(g.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (g *TGraphic) ClassName() string {
	return Graphic_ClassName(g.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (g *TGraphic) InstanceSize() int32 {
	return Graphic_InstanceSize(g.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (g *TGraphic) InheritsFrom(AClass TClass) bool {
	return Graphic_InheritsFrom(g.instance, AClass)
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (g *TGraphic) GetHashCode() int32 {
	return Graphic_GetHashCode(g.instance)
}

// 文本类信息。
//
// Text information.
func (g *TGraphic) ToString() string {
	return Graphic_ToString(g.instance)
}

// 获取对象是否为空。
//
// Get object is empty.
func (g *TGraphic) Empty() bool {
	return Graphic_GetEmpty(g.instance)
}

// 获取高度。
//
// Get height.
func (g *TGraphic) Height() int32 {
	return Graphic_GetHeight(g.instance)
}

// 设置高度。
//
// Set height.
func (g *TGraphic) SetHeight(value int32) {
	Graphic_SetHeight(g.instance, value)
}

// 获取修改。
//
// Get modified.
func (g *TGraphic) Modified() bool {
	return Graphic_GetModified(g.instance)
}

// 设置修改。
//
// Set modified.
func (g *TGraphic) SetModified(value bool) {
	Graphic_SetModified(g.instance, value)
}

func (g *TGraphic) Palette() HPALETTE {
	return Graphic_GetPalette(g.instance)
}

func (g *TGraphic) SetPalette(value HPALETTE) {
	Graphic_SetPalette(g.instance, value)
}

func (g *TGraphic) PaletteModified() bool {
	return Graphic_GetPaletteModified(g.instance)
}

func (g *TGraphic) SetPaletteModified(value bool) {
	Graphic_SetPaletteModified(g.instance, value)
}

// 获取透明。
//
// Get transparent.
func (g *TGraphic) Transparent() bool {
	return Graphic_GetTransparent(g.instance)
}

// 设置透明。
//
// Set transparent.
func (g *TGraphic) SetTransparent(value bool) {
	Graphic_SetTransparent(g.instance, value)
}

// 获取宽度。
//
// Get width.
func (g *TGraphic) Width() int32 {
	return Graphic_GetWidth(g.instance)
}

// 设置宽度。
//
// Set width.
func (g *TGraphic) SetWidth(value int32) {
	Graphic_SetWidth(g.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (g *TGraphic) SetOnChange(fn TNotifyEvent) {
	Graphic_SetOnChange(g.instance, fn)
}
