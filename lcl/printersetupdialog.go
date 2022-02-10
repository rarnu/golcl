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

type TPrinterSetupDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewPrinterSetupDialog(owner IComponent) *TPrinterSetupDialog {
	p := new(TPrinterSetupDialog)
	p.instance = PrinterSetupDialog_Create(CheckPtr(owner))
	p.ptr = unsafe.Pointer(p.instance)
	return p
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsPrinterSetupDialog(obj interface{}) *TPrinterSetupDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPrinterSetupDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPrinterSetupDialog.
func PrinterSetupDialogFromInst(inst uintptr) *TPrinterSetupDialog {
	return AsPrinterSetupDialog(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsPrinterSetupDialog.
func PrinterSetupDialogFromObj(obj IObject) *TPrinterSetupDialog {
	return AsPrinterSetupDialog(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPrinterSetupDialog.
func PrinterSetupDialogFromUnsafePointer(ptr unsafe.Pointer) *TPrinterSetupDialog {
	return AsPrinterSetupDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (p *TPrinterSetupDialog) Free() {
	if p.instance != 0 {
		PrinterSetupDialog_Free(p.instance)
		p.instance, p.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (p *TPrinterSetupDialog) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (p *TPrinterSetupDialog) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (p *TPrinterSetupDialog) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (p *TPrinterSetupDialog) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (p *TPrinterSetupDialog) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TPrinterSetupDialogClass() TClass {
	return PrinterSetupDialog_StaticClassType()
}

// 执行。
func (p *TPrinterSetupDialog) Execute() bool {
	return PrinterSetupDialog_Execute(p.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (p *TPrinterSetupDialog) FindComponent(AName string) *TComponent {
	return AsComponent(PrinterSetupDialog_FindComponent(p.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPrinterSetupDialog) GetNamePath() string {
	return PrinterSetupDialog_GetNamePath(p.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (p *TPrinterSetupDialog) HasParent() bool {
	return PrinterSetupDialog_HasParent(p.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPrinterSetupDialog) Assign(Source IObject) {
	PrinterSetupDialog_Assign(p.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPrinterSetupDialog) ClassType() TClass {
	return PrinterSetupDialog_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPrinterSetupDialog) ClassName() string {
	return PrinterSetupDialog_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPrinterSetupDialog) InstanceSize() int32 {
	return PrinterSetupDialog_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPrinterSetupDialog) InheritsFrom(AClass TClass) bool {
	return PrinterSetupDialog_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPrinterSetupDialog) Equals(Obj IObject) bool {
	return PrinterSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPrinterSetupDialog) GetHashCode() int32 {
	return PrinterSetupDialog_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPrinterSetupDialog) ToString() string {
	return PrinterSetupDialog_ToString(p.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (p *TPrinterSetupDialog) Handle() HWND {
	return PrinterSetupDialog_GetHandle(p.instance)
}

func (p *TPrinterSetupDialog) SetOnClose(fn TNotifyEvent) {
	PrinterSetupDialog_SetOnClose(p.instance, fn)
}

// 设置显示事件。
func (p *TPrinterSetupDialog) SetOnShow(fn TNotifyEvent) {
	PrinterSetupDialog_SetOnShow(p.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (p *TPrinterSetupDialog) ComponentCount() int32 {
	return PrinterSetupDialog_GetComponentCount(p.instance)
}

// 获取组件索引。
//
// Get component index.
func (p *TPrinterSetupDialog) ComponentIndex() int32 {
	return PrinterSetupDialog_GetComponentIndex(p.instance)
}

// 设置组件索引。
//
// Set component index.
func (p *TPrinterSetupDialog) SetComponentIndex(value int32) {
	PrinterSetupDialog_SetComponentIndex(p.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (p *TPrinterSetupDialog) Owner() *TComponent {
	return AsComponent(PrinterSetupDialog_GetOwner(p.instance))
}

// 获取组件名称。
//
// Get the component name.
func (p *TPrinterSetupDialog) Name() string {
	return PrinterSetupDialog_GetName(p.instance)
}

// 设置组件名称。
//
// Set the component name.
func (p *TPrinterSetupDialog) SetName(value string) {
	PrinterSetupDialog_SetName(p.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (p *TPrinterSetupDialog) Tag() int {
	return PrinterSetupDialog_GetTag(p.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (p *TPrinterSetupDialog) SetTag(value int) {
	PrinterSetupDialog_SetTag(p.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (p *TPrinterSetupDialog) Components(AIndex int32) *TComponent {
	return AsComponent(PrinterSetupDialog_GetComponents(p.instance, AIndex))
}
