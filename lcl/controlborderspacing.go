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

type TControlBorderSpacing struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsControlBorderSpacing(obj any) *TControlBorderSpacing {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TControlBorderSpacing{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsControlBorderSpacing.
func ControlBorderSpacingFromInst(inst uintptr) *TControlBorderSpacing {
	return AsControlBorderSpacing(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsControlBorderSpacing.
func ControlBorderSpacingFromObj(obj IObject) *TControlBorderSpacing {
	return AsControlBorderSpacing(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsControlBorderSpacing.
func ControlBorderSpacingFromUnsafePointer(ptr unsafe.Pointer) *TControlBorderSpacing {
	return AsControlBorderSpacing(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TControlBorderSpacing) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TControlBorderSpacing) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TControlBorderSpacing) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TControlBorderSpacing) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TControlBorderSpacing) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TControlBorderSpacingClass() TClass {
	return ControlBorderSpacing_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TControlBorderSpacing) Assign(Source IObject) {
	ControlBorderSpacing_Assign(c.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TControlBorderSpacing) GetNamePath() string {
	return ControlBorderSpacing_GetNamePath(c.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TControlBorderSpacing) ClassType() TClass {
	return ControlBorderSpacing_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TControlBorderSpacing) ClassName() string {
	return ControlBorderSpacing_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TControlBorderSpacing) InstanceSize() int32 {
	return ControlBorderSpacing_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TControlBorderSpacing) InheritsFrom(AClass TClass) bool {
	return ControlBorderSpacing_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TControlBorderSpacing) Equals(Obj IObject) bool {
	return ControlBorderSpacing_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TControlBorderSpacing) GetHashCode() int32 {
	return ControlBorderSpacing_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TControlBorderSpacing) ToString() string {
	return ControlBorderSpacing_ToString(c.instance)
}

func (c *TControlBorderSpacing) Control() *TControl {
	return AsControl(ControlBorderSpacing_GetControl(c.instance))
}

func (c *TControlBorderSpacing) AroundLeft() int32 {
	return ControlBorderSpacing_GetAroundLeft(c.instance)
}

func (c *TControlBorderSpacing) AroundTop() int32 {
	return ControlBorderSpacing_GetAroundTop(c.instance)
}

func (c *TControlBorderSpacing) AroundRight() int32 {
	return ControlBorderSpacing_GetAroundRight(c.instance)
}

func (c *TControlBorderSpacing) AroundBottom() int32 {
	return ControlBorderSpacing_GetAroundBottom(c.instance)
}

func (c *TControlBorderSpacing) ControlLeft() int32 {
	return ControlBorderSpacing_GetControlLeft(c.instance)
}

func (c *TControlBorderSpacing) ControlTop() int32 {
	return ControlBorderSpacing_GetControlTop(c.instance)
}

func (c *TControlBorderSpacing) ControlWidth() int32 {
	return ControlBorderSpacing_GetControlWidth(c.instance)
}

func (c *TControlBorderSpacing) ControlHeight() int32 {
	return ControlBorderSpacing_GetControlHeight(c.instance)
}

func (c *TControlBorderSpacing) ControlRight() int32 {
	return ControlBorderSpacing_GetControlRight(c.instance)
}

func (c *TControlBorderSpacing) ControlBottom() int32 {
	return ControlBorderSpacing_GetControlBottom(c.instance)
}

// 设置改变事件。
//
// Set changed event.
func (c *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
	ControlBorderSpacing_SetOnChange(c.instance, fn)
}

// 获取左边位置。
//
// Get Left position.
func (c *TControlBorderSpacing) Left() int32 {
	return ControlBorderSpacing_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TControlBorderSpacing) SetLeft(value int32) {
	ControlBorderSpacing_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TControlBorderSpacing) Top() int32 {
	return ControlBorderSpacing_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TControlBorderSpacing) SetTop(value int32) {
	ControlBorderSpacing_SetTop(c.instance, value)
}

func (c *TControlBorderSpacing) Right() int32 {
	return ControlBorderSpacing_GetRight(c.instance)
}

func (c *TControlBorderSpacing) SetRight(value int32) {
	ControlBorderSpacing_SetRight(c.instance, value)
}

func (c *TControlBorderSpacing) Bottom() int32 {
	return ControlBorderSpacing_GetBottom(c.instance)
}

func (c *TControlBorderSpacing) SetBottom(value int32) {
	ControlBorderSpacing_SetBottom(c.instance, value)
}

func (c *TControlBorderSpacing) Around() int32 {
	return ControlBorderSpacing_GetAround(c.instance)
}

func (c *TControlBorderSpacing) SetAround(value int32) {
	ControlBorderSpacing_SetAround(c.instance, value)
}

func (c *TControlBorderSpacing) InnerBorder() int32 {
	return ControlBorderSpacing_GetInnerBorder(c.instance)
}

func (c *TControlBorderSpacing) SetInnerBorder(value int32) {
	ControlBorderSpacing_SetInnerBorder(c.instance, value)
}

func (c *TControlBorderSpacing) CellAlignHorizontal() TControlCellAlign {
	return ControlBorderSpacing_GetCellAlignHorizontal(c.instance)
}

func (c *TControlBorderSpacing) SetCellAlignHorizontal(value TControlCellAlign) {
	ControlBorderSpacing_SetCellAlignHorizontal(c.instance, value)
}

func (c *TControlBorderSpacing) CellAlignVertical() TControlCellAlign {
	return ControlBorderSpacing_GetCellAlignVertical(c.instance)
}

func (c *TControlBorderSpacing) SetCellAlignVertical(value TControlCellAlign) {
	ControlBorderSpacing_SetCellAlignVertical(c.instance, value)
}

func (c *TControlBorderSpacing) Space(Kind TAnchorKind) int32 {
	return ControlBorderSpacing_GetSpace(c.instance, Kind)
}

func (c *TControlBorderSpacing) SetSpace(Kind TAnchorKind, value int32) {
	ControlBorderSpacing_SetSpace(c.instance, Kind, value)
}
