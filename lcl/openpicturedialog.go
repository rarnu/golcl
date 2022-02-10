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

type TOpenPictureDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewOpenPictureDialog(owner IComponent) *TOpenPictureDialog {
	o := new(TOpenPictureDialog)
	o.instance = OpenPictureDialog_Create(CheckPtr(owner))
	o.ptr = unsafe.Pointer(o.instance)
	return o
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsOpenPictureDialog(obj interface{}) *TOpenPictureDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TOpenPictureDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsOpenPictureDialog.
func OpenPictureDialogFromInst(inst uintptr) *TOpenPictureDialog {
	return AsOpenPictureDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsOpenPictureDialog.
func OpenPictureDialogFromObj(obj IObject) *TOpenPictureDialog {
	return AsOpenPictureDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsOpenPictureDialog.
func OpenPictureDialogFromUnsafePointer(ptr unsafe.Pointer) *TOpenPictureDialog {
	return AsOpenPictureDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (o *TOpenPictureDialog) Free() {
	if o.instance != 0 {
		OpenPictureDialog_Free(o.instance)
		o.instance, o.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (o *TOpenPictureDialog) Instance() uintptr {
	return o.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (o *TOpenPictureDialog) UnsafeAddr() unsafe.Pointer {
	return o.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (o *TOpenPictureDialog) IsValid() bool {
	return o.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (o *TOpenPictureDialog) Is() TIs {
	return TIs(o.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (o *TOpenPictureDialog) As() TAs {
//    return TAs(o.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TOpenPictureDialogClass() TClass {
	return OpenPictureDialog_StaticClassType()
}

// 执行。
func (o *TOpenPictureDialog) Execute() bool {
	return OpenPictureDialog_Execute(o.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (o *TOpenPictureDialog) FindComponent(AName string) *TComponent {
	return AsComponent(OpenPictureDialog_FindComponent(o.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (o *TOpenPictureDialog) GetNamePath() string {
	return OpenPictureDialog_GetNamePath(o.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (o *TOpenPictureDialog) HasParent() bool {
	return OpenPictureDialog_HasParent(o.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (o *TOpenPictureDialog) Assign(Source IObject) {
	OpenPictureDialog_Assign(o.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (o *TOpenPictureDialog) ClassType() TClass {
	return OpenPictureDialog_ClassType(o.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (o *TOpenPictureDialog) ClassName() string {
	return OpenPictureDialog_ClassName(o.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (o *TOpenPictureDialog) InstanceSize() int32 {
	return OpenPictureDialog_InstanceSize(o.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (o *TOpenPictureDialog) InheritsFrom(AClass TClass) bool {
	return OpenPictureDialog_InheritsFrom(o.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (o *TOpenPictureDialog) Equals(Obj IObject) bool {
	return OpenPictureDialog_Equals(o.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (o *TOpenPictureDialog) GetHashCode() int32 {
	return OpenPictureDialog_GetHashCode(o.instance)
}

// 文本类信息。
//
// Text information.
func (o *TOpenPictureDialog) ToString() string {
	return OpenPictureDialog_ToString(o.instance)
}

func (o *TOpenPictureDialog) Filter() string {
	return OpenPictureDialog_GetFilter(o.instance)
}

func (o *TOpenPictureDialog) SetFilter(value string) {
	OpenPictureDialog_SetFilter(o.instance, value)
}

func (o *TOpenPictureDialog) Files() *TStrings {
	return AsStrings(OpenPictureDialog_GetFiles(o.instance))
}

func (o *TOpenPictureDialog) DefaultExt() string {
	return OpenPictureDialog_GetDefaultExt(o.instance)
}

func (o *TOpenPictureDialog) SetDefaultExt(value string) {
	OpenPictureDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenPictureDialog) FileName() string {
	return OpenPictureDialog_GetFileName(o.instance)
}

func (o *TOpenPictureDialog) SetFileName(value string) {
	OpenPictureDialog_SetFileName(o.instance, value)
}

func (o *TOpenPictureDialog) FilterIndex() int32 {
	return OpenPictureDialog_GetFilterIndex(o.instance)
}

func (o *TOpenPictureDialog) SetFilterIndex(value int32) {
	OpenPictureDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenPictureDialog) InitialDir() string {
	return OpenPictureDialog_GetInitialDir(o.instance)
}

func (o *TOpenPictureDialog) SetInitialDir(value string) {
	OpenPictureDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenPictureDialog) Options() TOpenOptions {
	return OpenPictureDialog_GetOptions(o.instance)
}

func (o *TOpenPictureDialog) SetOptions(value TOpenOptions) {
	OpenPictureDialog_SetOptions(o.instance, value)
}

func (o *TOpenPictureDialog) Title() string {
	return OpenPictureDialog_GetTitle(o.instance)
}

func (o *TOpenPictureDialog) SetTitle(value string) {
	OpenPictureDialog_SetTitle(o.instance, value)
}

// 获取控件句柄。
//
// Get Control handle.
func (o *TOpenPictureDialog) Handle() HWND {
	return OpenPictureDialog_GetHandle(o.instance)
}

func (o *TOpenPictureDialog) SetOnClose(fn TNotifyEvent) {
	OpenPictureDialog_SetOnClose(o.instance, fn)
}

// 设置显示事件。
func (o *TOpenPictureDialog) SetOnShow(fn TNotifyEvent) {
	OpenPictureDialog_SetOnShow(o.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (o *TOpenPictureDialog) ComponentCount() int32 {
	return OpenPictureDialog_GetComponentCount(o.instance)
}

// 获取组件索引。
//
// Get component index.
func (o *TOpenPictureDialog) ComponentIndex() int32 {
	return OpenPictureDialog_GetComponentIndex(o.instance)
}

// 设置组件索引。
//
// Set component index.
func (o *TOpenPictureDialog) SetComponentIndex(value int32) {
	OpenPictureDialog_SetComponentIndex(o.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (o *TOpenPictureDialog) Owner() *TComponent {
	return AsComponent(OpenPictureDialog_GetOwner(o.instance))
}

// 获取组件名称。
//
// Get the component name.
func (o *TOpenPictureDialog) Name() string {
	return OpenPictureDialog_GetName(o.instance)
}

// 设置组件名称。
//
// Set the component name.
func (o *TOpenPictureDialog) SetName(value string) {
	OpenPictureDialog_SetName(o.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (o *TOpenPictureDialog) Tag() int {
	return OpenPictureDialog_GetTag(o.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (o *TOpenPictureDialog) SetTag(value int) {
	OpenPictureDialog_SetTag(o.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (o *TOpenPictureDialog) Components(AIndex int32) *TComponent {
	return AsComponent(OpenPictureDialog_GetComponents(o.instance, AIndex))
}
