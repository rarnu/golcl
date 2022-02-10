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

type TColorDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewColorDialog(owner IComponent) *TColorDialog {
	c := new(TColorDialog)
	c.instance = ColorDialog_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsColorDialog(obj interface{}) *TColorDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TColorDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsColorDialog.
func ColorDialogFromInst(inst uintptr) *TColorDialog {
	return AsColorDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsColorDialog.
func ColorDialogFromObj(obj IObject) *TColorDialog {
	return AsColorDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsColorDialog.
func ColorDialogFromUnsafePointer(ptr unsafe.Pointer) *TColorDialog {
	return AsColorDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TColorDialog) Free() {
	if c.instance != 0 {
		ColorDialog_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TColorDialog) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TColorDialog) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TColorDialog) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TColorDialog) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TColorDialog) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TColorDialogClass() TClass {
	return ColorDialog_StaticClassType()
}

// 执行。
func (c *TColorDialog) Execute() bool {
	return ColorDialog_Execute(c.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TColorDialog) FindComponent(AName string) *TComponent {
	return AsComponent(ColorDialog_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TColorDialog) GetNamePath() string {
	return ColorDialog_GetNamePath(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TColorDialog) HasParent() bool {
	return ColorDialog_HasParent(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TColorDialog) Assign(Source IObject) {
	ColorDialog_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TColorDialog) ClassType() TClass {
	return ColorDialog_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TColorDialog) ClassName() string {
	return ColorDialog_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TColorDialog) InstanceSize() int32 {
	return ColorDialog_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TColorDialog) InheritsFrom(AClass TClass) bool {
	return ColorDialog_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TColorDialog) Equals(Obj IObject) bool {
	return ColorDialog_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TColorDialog) GetHashCode() int32 {
	return ColorDialog_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TColorDialog) ToString() string {
	return ColorDialog_ToString(c.instance)
}

// 获取颜色。
//
// Get color.
func (c *TColorDialog) Color() TColor {
	return ColorDialog_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TColorDialog) SetColor(value TColor) {
	ColorDialog_SetColor(c.instance, value)
}

func (c *TColorDialog) CustomColors() *TStrings {
	return AsStrings(ColorDialog_GetCustomColors(c.instance))
}

func (c *TColorDialog) SetCustomColors(value IStrings) {
	ColorDialog_SetCustomColors(c.instance, CheckPtr(value))
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TColorDialog) Handle() HWND {
	return ColorDialog_GetHandle(c.instance)
}

func (c *TColorDialog) SetOnClose(fn TNotifyEvent) {
	ColorDialog_SetOnClose(c.instance, fn)
}

// 设置显示事件。
func (c *TColorDialog) SetOnShow(fn TNotifyEvent) {
	ColorDialog_SetOnShow(c.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TColorDialog) ComponentCount() int32 {
	return ColorDialog_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TColorDialog) ComponentIndex() int32 {
	return ColorDialog_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TColorDialog) SetComponentIndex(value int32) {
	ColorDialog_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TColorDialog) Owner() *TComponent {
	return AsComponent(ColorDialog_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TColorDialog) Name() string {
	return ColorDialog_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TColorDialog) SetName(value string) {
	ColorDialog_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TColorDialog) Tag() int {
	return ColorDialog_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TColorDialog) SetTag(value int) {
	ColorDialog_SetTag(c.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TColorDialog) Components(AIndex int32) *TComponent {
	return AsComponent(ColorDialog_GetComponents(c.instance, AIndex))
}
