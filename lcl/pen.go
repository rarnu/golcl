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

type TPen struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewPen() *TPen {
	p := new(TPen)
	p.instance = Pen_Create()
	p.ptr = unsafe.Pointer(p.instance)
	setFinalizer(p, (*TPen).Free)
	return p
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPen(obj any) *TPen {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPen{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPen.
func PenFromInst(inst uintptr) *TPen {
	return AsPen(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsPen.
func PenFromObj(obj IObject) *TPen {
	return AsPen(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPen.
func PenFromUnsafePointer(ptr unsafe.Pointer) *TPen {
	return AsPen(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (p *TPen) Free() {
	if p.instance != 0 {
		Pen_Free(p.instance)
		p.instance, p.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPen) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPen) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPen) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPen) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPen) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TPenClass() TClass {
	return Pen_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPen) Assign(Source IObject) {
	Pen_Assign(p.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPen) GetNamePath() string {
	return Pen_GetNamePath(p.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPen) ClassType() TClass {
	return Pen_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPen) ClassName() string {
	return Pen_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPen) InstanceSize() int32 {
	return Pen_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPen) InheritsFrom(AClass TClass) bool {
	return Pen_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPen) Equals(Obj IObject) bool {
	return Pen_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPen) GetHashCode() int32 {
	return Pen_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPen) ToString() string {
	return Pen_ToString(p.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (p *TPen) Handle() HPEN {
	return Pen_GetHandle(p.instance)
}

// 设置控件句柄。
//
// Set Control handle.
func (p *TPen) SetHandle(value HPEN) {
	Pen_SetHandle(p.instance, value)
}

// 获取颜色。
//
// Get color.
func (p *TPen) Color() TColor {
	return Pen_GetColor(p.instance)
}

// 设置颜色。
//
// Set color.
func (p *TPen) SetColor(value TColor) {
	Pen_SetColor(p.instance, value)
}

func (p *TPen) Mode() TPenMode {
	return Pen_GetMode(p.instance)
}

func (p *TPen) SetMode(value TPenMode) {
	Pen_SetMode(p.instance, value)
}

func (p *TPen) Style() TPenStyle {
	return Pen_GetStyle(p.instance)
}

func (p *TPen) SetStyle(value TPenStyle) {
	Pen_SetStyle(p.instance, value)
}

// 获取宽度。
//
// Get width.
func (p *TPen) Width() int32 {
	return Pen_GetWidth(p.instance)
}

// 设置宽度。
//
// Set width.
func (p *TPen) SetWidth(value int32) {
	Pen_SetWidth(p.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (p *TPen) SetOnChange(fn TNotifyEvent) {
	Pen_SetOnChange(p.instance, fn)
}
