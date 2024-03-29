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

type TSavePictureDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewSavePictureDialog(owner IComponent) *TSavePictureDialog {
	s := new(TSavePictureDialog)
	s.instance = SavePictureDialog_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsSavePictureDialog(obj any) *TSavePictureDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSavePictureDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsSavePictureDialog.
func SavePictureDialogFromInst(inst uintptr) *TSavePictureDialog {
	return AsSavePictureDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsSavePictureDialog.
func SavePictureDialogFromObj(obj IObject) *TSavePictureDialog {
	return AsSavePictureDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsSavePictureDialog.
func SavePictureDialogFromUnsafePointer(ptr unsafe.Pointer) *TSavePictureDialog {
	return AsSavePictureDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (s *TSavePictureDialog) Free() {
	if s.instance != 0 {
		SavePictureDialog_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TSavePictureDialog) Instance() uintptr {
	return s.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TSavePictureDialog) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TSavePictureDialog) IsValid() bool {
	return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TSavePictureDialog) Is() TIs {
	return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TSavePictureDialog) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TSavePictureDialogClass() TClass {
	return SavePictureDialog_StaticClassType()
}

// 执行。
func (s *TSavePictureDialog) Execute() bool {
	return SavePictureDialog_Execute(s.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TSavePictureDialog) FindComponent(AName string) *TComponent {
	return AsComponent(SavePictureDialog_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TSavePictureDialog) GetNamePath() string {
	return SavePictureDialog_GetNamePath(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TSavePictureDialog) HasParent() bool {
	return SavePictureDialog_HasParent(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TSavePictureDialog) Assign(Source IObject) {
	SavePictureDialog_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TSavePictureDialog) ClassType() TClass {
	return SavePictureDialog_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TSavePictureDialog) ClassName() string {
	return SavePictureDialog_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TSavePictureDialog) InstanceSize() int32 {
	return SavePictureDialog_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TSavePictureDialog) InheritsFrom(AClass TClass) bool {
	return SavePictureDialog_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TSavePictureDialog) Equals(Obj IObject) bool {
	return SavePictureDialog_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TSavePictureDialog) GetHashCode() int32 {
	return SavePictureDialog_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TSavePictureDialog) ToString() string {
	return SavePictureDialog_ToString(s.instance)
}

func (s *TSavePictureDialog) Filter() string {
	return SavePictureDialog_GetFilter(s.instance)
}

func (s *TSavePictureDialog) SetFilter(value string) {
	SavePictureDialog_SetFilter(s.instance, value)
}

func (s *TSavePictureDialog) Files() *TStrings {
	return AsStrings(SavePictureDialog_GetFiles(s.instance))
}

func (s *TSavePictureDialog) DefaultExt() string {
	return SavePictureDialog_GetDefaultExt(s.instance)
}

func (s *TSavePictureDialog) SetDefaultExt(value string) {
	SavePictureDialog_SetDefaultExt(s.instance, value)
}

func (s *TSavePictureDialog) FileName() string {
	return SavePictureDialog_GetFileName(s.instance)
}

func (s *TSavePictureDialog) SetFileName(value string) {
	SavePictureDialog_SetFileName(s.instance, value)
}

func (s *TSavePictureDialog) FilterIndex() int32 {
	return SavePictureDialog_GetFilterIndex(s.instance)
}

func (s *TSavePictureDialog) SetFilterIndex(value int32) {
	SavePictureDialog_SetFilterIndex(s.instance, value)
}

func (s *TSavePictureDialog) InitialDir() string {
	return SavePictureDialog_GetInitialDir(s.instance)
}

func (s *TSavePictureDialog) SetInitialDir(value string) {
	SavePictureDialog_SetInitialDir(s.instance, value)
}

func (s *TSavePictureDialog) Options() TOpenOptions {
	return SavePictureDialog_GetOptions(s.instance)
}

func (s *TSavePictureDialog) SetOptions(value TOpenOptions) {
	SavePictureDialog_SetOptions(s.instance, value)
}

func (s *TSavePictureDialog) Title() string {
	return SavePictureDialog_GetTitle(s.instance)
}

func (s *TSavePictureDialog) SetTitle(value string) {
	SavePictureDialog_SetTitle(s.instance, value)
}

// 获取控件句柄。
//
// Get Control handle.
func (s *TSavePictureDialog) Handle() HWND {
	return SavePictureDialog_GetHandle(s.instance)
}

func (s *TSavePictureDialog) SetOnClose(fn TNotifyEvent) {
	SavePictureDialog_SetOnClose(s.instance, fn)
}

// 设置显示事件。
func (s *TSavePictureDialog) SetOnShow(fn TNotifyEvent) {
	SavePictureDialog_SetOnShow(s.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TSavePictureDialog) ComponentCount() int32 {
	return SavePictureDialog_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TSavePictureDialog) ComponentIndex() int32 {
	return SavePictureDialog_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TSavePictureDialog) SetComponentIndex(value int32) {
	SavePictureDialog_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TSavePictureDialog) Owner() *TComponent {
	return AsComponent(SavePictureDialog_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TSavePictureDialog) Name() string {
	return SavePictureDialog_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TSavePictureDialog) SetName(value string) {
	SavePictureDialog_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TSavePictureDialog) Tag() int {
	return SavePictureDialog_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TSavePictureDialog) SetTag(value int) {
	SavePictureDialog_SetTag(s.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TSavePictureDialog) Components(AIndex int32) *TComponent {
	return AsComponent(SavePictureDialog_GetComponents(s.instance, AIndex))
}
