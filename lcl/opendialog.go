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

type TOpenDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewOpenDialog(owner IComponent) *TOpenDialog {
	o := new(TOpenDialog)
	o.instance = OpenDialog_Create(CheckPtr(owner))
	o.ptr = unsafe.Pointer(o.instance)
	return o
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsOpenDialog(obj any) *TOpenDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TOpenDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsOpenDialog.
func OpenDialogFromInst(inst uintptr) *TOpenDialog {
	return AsOpenDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsOpenDialog.
func OpenDialogFromObj(obj IObject) *TOpenDialog {
	return AsOpenDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsOpenDialog.
func OpenDialogFromUnsafePointer(ptr unsafe.Pointer) *TOpenDialog {
	return AsOpenDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (o *TOpenDialog) Free() {
	if o.instance != 0 {
		OpenDialog_Free(o.instance)
		o.instance, o.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (o *TOpenDialog) Instance() uintptr {
	return o.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (o *TOpenDialog) UnsafeAddr() unsafe.Pointer {
	return o.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (o *TOpenDialog) IsValid() bool {
	return o.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (o *TOpenDialog) Is() TIs {
	return TIs(o.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (o *TOpenDialog) As() TAs {
//    return TAs(o.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TOpenDialogClass() TClass {
	return OpenDialog_StaticClassType()
}

// 执行。
func (o *TOpenDialog) Execute() bool {
	return OpenDialog_Execute(o.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (o *TOpenDialog) FindComponent(AName string) *TComponent {
	return AsComponent(OpenDialog_FindComponent(o.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (o *TOpenDialog) GetNamePath() string {
	return OpenDialog_GetNamePath(o.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (o *TOpenDialog) HasParent() bool {
	return OpenDialog_HasParent(o.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (o *TOpenDialog) Assign(Source IObject) {
	OpenDialog_Assign(o.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (o *TOpenDialog) ClassType() TClass {
	return OpenDialog_ClassType(o.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (o *TOpenDialog) ClassName() string {
	return OpenDialog_ClassName(o.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (o *TOpenDialog) InstanceSize() int32 {
	return OpenDialog_InstanceSize(o.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (o *TOpenDialog) InheritsFrom(AClass TClass) bool {
	return OpenDialog_InheritsFrom(o.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (o *TOpenDialog) Equals(Obj IObject) bool {
	return OpenDialog_Equals(o.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (o *TOpenDialog) GetHashCode() int32 {
	return OpenDialog_GetHashCode(o.instance)
}

// 文本类信息。
//
// Text information.
func (o *TOpenDialog) ToString() string {
	return OpenDialog_ToString(o.instance)
}

func (o *TOpenDialog) Files() *TStrings {
	return AsStrings(OpenDialog_GetFiles(o.instance))
}

func (o *TOpenDialog) DefaultExt() string {
	return OpenDialog_GetDefaultExt(o.instance)
}

func (o *TOpenDialog) SetDefaultExt(value string) {
	OpenDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenDialog) FileName() string {
	return OpenDialog_GetFileName(o.instance)
}

func (o *TOpenDialog) SetFileName(value string) {
	OpenDialog_SetFileName(o.instance, value)
}

func (o *TOpenDialog) Filter() string {
	return OpenDialog_GetFilter(o.instance)
}

func (o *TOpenDialog) SetFilter(value string) {
	OpenDialog_SetFilter(o.instance, value)
}

func (o *TOpenDialog) FilterIndex() int32 {
	return OpenDialog_GetFilterIndex(o.instance)
}

func (o *TOpenDialog) SetFilterIndex(value int32) {
	OpenDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenDialog) InitialDir() string {
	return OpenDialog_GetInitialDir(o.instance)
}

func (o *TOpenDialog) SetInitialDir(value string) {
	OpenDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenDialog) Options() TOpenOptions {
	return OpenDialog_GetOptions(o.instance)
}

func (o *TOpenDialog) SetOptions(value TOpenOptions) {
	OpenDialog_SetOptions(o.instance, value)
}

func (o *TOpenDialog) Title() string {
	return OpenDialog_GetTitle(o.instance)
}

func (o *TOpenDialog) SetTitle(value string) {
	OpenDialog_SetTitle(o.instance, value)
}

// 获取控件句柄。
//
// Get Control handle.
func (o *TOpenDialog) Handle() HWND {
	return OpenDialog_GetHandle(o.instance)
}

func (o *TOpenDialog) SetOnClose(fn TNotifyEvent) {
	OpenDialog_SetOnClose(o.instance, fn)
}

// 设置显示事件。
func (o *TOpenDialog) SetOnShow(fn TNotifyEvent) {
	OpenDialog_SetOnShow(o.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (o *TOpenDialog) ComponentCount() int32 {
	return OpenDialog_GetComponentCount(o.instance)
}

// 获取组件索引。
//
// Get component index.
func (o *TOpenDialog) ComponentIndex() int32 {
	return OpenDialog_GetComponentIndex(o.instance)
}

// 设置组件索引。
//
// Set component index.
func (o *TOpenDialog) SetComponentIndex(value int32) {
	OpenDialog_SetComponentIndex(o.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (o *TOpenDialog) Owner() *TComponent {
	return AsComponent(OpenDialog_GetOwner(o.instance))
}

// 获取组件名称。
//
// Get the component name.
func (o *TOpenDialog) Name() string {
	return OpenDialog_GetName(o.instance)
}

// 设置组件名称。
//
// Set the component name.
func (o *TOpenDialog) SetName(value string) {
	OpenDialog_SetName(o.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (o *TOpenDialog) Tag() int {
	return OpenDialog_GetTag(o.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (o *TOpenDialog) SetTag(value int) {
	OpenDialog_SetTag(o.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (o *TOpenDialog) Components(AIndex int32) *TComponent {
	return AsComponent(OpenDialog_GetComponents(o.instance, AIndex))
}
