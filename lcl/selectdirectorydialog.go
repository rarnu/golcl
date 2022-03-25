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

type TSelectDirectoryDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewSelectDirectoryDialog(owner IComponent) *TSelectDirectoryDialog {
	s := new(TSelectDirectoryDialog)
	s.instance = SelectDirectoryDialog_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsSelectDirectoryDialog(obj any) *TSelectDirectoryDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSelectDirectoryDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsSelectDirectoryDialog.
func SelectDirectoryDialogFromInst(inst uintptr) *TSelectDirectoryDialog {
	return AsSelectDirectoryDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsSelectDirectoryDialog.
func SelectDirectoryDialogFromObj(obj IObject) *TSelectDirectoryDialog {
	return AsSelectDirectoryDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsSelectDirectoryDialog.
func SelectDirectoryDialogFromUnsafePointer(ptr unsafe.Pointer) *TSelectDirectoryDialog {
	return AsSelectDirectoryDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (s *TSelectDirectoryDialog) Free() {
	if s.instance != 0 {
		SelectDirectoryDialog_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TSelectDirectoryDialog) Instance() uintptr {
	return s.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TSelectDirectoryDialog) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TSelectDirectoryDialog) IsValid() bool {
	return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TSelectDirectoryDialog) Is() TIs {
	return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TSelectDirectoryDialog) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TSelectDirectoryDialogClass() TClass {
	return SelectDirectoryDialog_StaticClassType()
}

// 执行。
func (s *TSelectDirectoryDialog) Execute() bool {
	return SelectDirectoryDialog_Execute(s.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TSelectDirectoryDialog) FindComponent(AName string) *TComponent {
	return AsComponent(SelectDirectoryDialog_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TSelectDirectoryDialog) GetNamePath() string {
	return SelectDirectoryDialog_GetNamePath(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TSelectDirectoryDialog) HasParent() bool {
	return SelectDirectoryDialog_HasParent(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TSelectDirectoryDialog) Assign(Source IObject) {
	SelectDirectoryDialog_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TSelectDirectoryDialog) ClassType() TClass {
	return SelectDirectoryDialog_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TSelectDirectoryDialog) ClassName() string {
	return SelectDirectoryDialog_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TSelectDirectoryDialog) InstanceSize() int32 {
	return SelectDirectoryDialog_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TSelectDirectoryDialog) InheritsFrom(AClass TClass) bool {
	return SelectDirectoryDialog_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TSelectDirectoryDialog) Equals(Obj IObject) bool {
	return SelectDirectoryDialog_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TSelectDirectoryDialog) GetHashCode() int32 {
	return SelectDirectoryDialog_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TSelectDirectoryDialog) ToString() string {
	return SelectDirectoryDialog_ToString(s.instance)
}

func (s *TSelectDirectoryDialog) Files() *TStrings {
	return AsStrings(SelectDirectoryDialog_GetFiles(s.instance))
}

func (s *TSelectDirectoryDialog) DefaultExt() string {
	return SelectDirectoryDialog_GetDefaultExt(s.instance)
}

func (s *TSelectDirectoryDialog) SetDefaultExt(value string) {
	SelectDirectoryDialog_SetDefaultExt(s.instance, value)
}

func (s *TSelectDirectoryDialog) FileName() string {
	return SelectDirectoryDialog_GetFileName(s.instance)
}

func (s *TSelectDirectoryDialog) SetFileName(value string) {
	SelectDirectoryDialog_SetFileName(s.instance, value)
}

func (s *TSelectDirectoryDialog) Filter() string {
	return SelectDirectoryDialog_GetFilter(s.instance)
}

func (s *TSelectDirectoryDialog) SetFilter(value string) {
	SelectDirectoryDialog_SetFilter(s.instance, value)
}

func (s *TSelectDirectoryDialog) FilterIndex() int32 {
	return SelectDirectoryDialog_GetFilterIndex(s.instance)
}

func (s *TSelectDirectoryDialog) SetFilterIndex(value int32) {
	SelectDirectoryDialog_SetFilterIndex(s.instance, value)
}

func (s *TSelectDirectoryDialog) InitialDir() string {
	return SelectDirectoryDialog_GetInitialDir(s.instance)
}

func (s *TSelectDirectoryDialog) SetInitialDir(value string) {
	SelectDirectoryDialog_SetInitialDir(s.instance, value)
}

func (s *TSelectDirectoryDialog) Options() TOpenOptions {
	return SelectDirectoryDialog_GetOptions(s.instance)
}

func (s *TSelectDirectoryDialog) SetOptions(value TOpenOptions) {
	SelectDirectoryDialog_SetOptions(s.instance, value)
}

func (s *TSelectDirectoryDialog) Title() string {
	return SelectDirectoryDialog_GetTitle(s.instance)
}

func (s *TSelectDirectoryDialog) SetTitle(value string) {
	SelectDirectoryDialog_SetTitle(s.instance, value)
}

// 获取控件句柄。
//
// Get Control handle.
func (s *TSelectDirectoryDialog) Handle() HWND {
	return SelectDirectoryDialog_GetHandle(s.instance)
}

func (s *TSelectDirectoryDialog) SetOnClose(fn TNotifyEvent) {
	SelectDirectoryDialog_SetOnClose(s.instance, fn)
}

// 设置显示事件。
func (s *TSelectDirectoryDialog) SetOnShow(fn TNotifyEvent) {
	SelectDirectoryDialog_SetOnShow(s.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TSelectDirectoryDialog) ComponentCount() int32 {
	return SelectDirectoryDialog_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TSelectDirectoryDialog) ComponentIndex() int32 {
	return SelectDirectoryDialog_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TSelectDirectoryDialog) SetComponentIndex(value int32) {
	SelectDirectoryDialog_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TSelectDirectoryDialog) Owner() *TComponent {
	return AsComponent(SelectDirectoryDialog_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TSelectDirectoryDialog) Name() string {
	return SelectDirectoryDialog_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TSelectDirectoryDialog) SetName(value string) {
	SelectDirectoryDialog_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TSelectDirectoryDialog) Tag() int {
	return SelectDirectoryDialog_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TSelectDirectoryDialog) SetTag(value int) {
	SelectDirectoryDialog_SetTag(s.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TSelectDirectoryDialog) Components(AIndex int32) *TComponent {
	return AsComponent(SelectDirectoryDialog_GetComponents(s.instance, AIndex))
}
