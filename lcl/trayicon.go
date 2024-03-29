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

type TTrayIcon struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewTrayIcon(owner IComponent) *TTrayIcon {
	t := new(TTrayIcon)
	t.instance = TrayIcon_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTrayIcon(obj any) *TTrayIcon {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTrayIcon{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTrayIcon.
func TrayIconFromInst(inst uintptr) *TTrayIcon {
	return AsTrayIcon(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsTrayIcon.
func TrayIconFromObj(obj IObject) *TTrayIcon {
	return AsTrayIcon(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTrayIcon.
func TrayIconFromUnsafePointer(ptr unsafe.Pointer) *TTrayIcon {
	return AsTrayIcon(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TTrayIcon) Free() {
	if t.instance != 0 {
		TrayIcon_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTrayIcon) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTrayIcon) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTrayIcon) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTrayIcon) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTrayIcon) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TTrayIconClass() TClass {
	return TrayIcon_StaticClassType()
}

func (t *TTrayIcon) ShowBalloonHint() {
	TrayIcon_ShowBalloonHint(t.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TTrayIcon) FindComponent(AName string) *TComponent {
	return AsComponent(TrayIcon_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTrayIcon) GetNamePath() string {
	return TrayIcon_GetNamePath(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TTrayIcon) HasParent() bool {
	return TrayIcon_HasParent(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTrayIcon) Assign(Source IObject) {
	TrayIcon_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTrayIcon) ClassType() TClass {
	return TrayIcon_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTrayIcon) ClassName() string {
	return TrayIcon_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTrayIcon) InstanceSize() int32 {
	return TrayIcon_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTrayIcon) InheritsFrom(AClass TClass) bool {
	return TrayIcon_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTrayIcon) Equals(Obj IObject) bool {
	return TrayIcon_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTrayIcon) GetHashCode() int32 {
	return TrayIcon_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTrayIcon) ToString() string {
	return TrayIcon_ToString(t.instance)
}

func (t *TTrayIcon) AnimateInterval() uint32 {
	return TrayIcon_GetAnimateInterval(t.instance)
}

func (t *TTrayIcon) SetAnimateInterval(value uint32) {
	TrayIcon_SetAnimateInterval(t.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TTrayIcon) Hint() string {
	return TrayIcon_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TTrayIcon) SetHint(value string) {
	TrayIcon_SetHint(t.instance, value)
}

func (t *TTrayIcon) BalloonHint() string {
	return TrayIcon_GetBalloonHint(t.instance)
}

func (t *TTrayIcon) SetBalloonHint(value string) {
	TrayIcon_SetBalloonHint(t.instance, value)
}

func (t *TTrayIcon) BalloonTitle() string {
	return TrayIcon_GetBalloonTitle(t.instance)
}

func (t *TTrayIcon) SetBalloonTitle(value string) {
	TrayIcon_SetBalloonTitle(t.instance, value)
}

func (t *TTrayIcon) BalloonTimeout() int32 {
	return TrayIcon_GetBalloonTimeout(t.instance)
}

func (t *TTrayIcon) SetBalloonTimeout(value int32) {
	TrayIcon_SetBalloonTimeout(t.instance, value)
}

func (t *TTrayIcon) BalloonFlags() TBalloonFlags {
	return TrayIcon_GetBalloonFlags(t.instance)
}

func (t *TTrayIcon) SetBalloonFlags(value TBalloonFlags) {
	TrayIcon_SetBalloonFlags(t.instance, value)
}

// 获取图标。
//
// Get icon.
func (t *TTrayIcon) Icon() *TIcon {
	return AsIcon(TrayIcon_GetIcon(t.instance))
}

// 设置图标。
//
// Set icon.
func (t *TTrayIcon) SetIcon(value *TIcon) {
	TrayIcon_SetIcon(t.instance, CheckPtr(value))
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TTrayIcon) PopupMenu() *TPopupMenu {
	return AsPopupMenu(TrayIcon_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TTrayIcon) SetPopupMenu(value IComponent) {
	TrayIcon_SetPopupMenu(t.instance, CheckPtr(value))
}

// 获取控件可视。
//
// Get the control visible.
func (t *TTrayIcon) Visible() bool {
	return TrayIcon_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TTrayIcon) SetVisible(value bool) {
	TrayIcon_SetVisible(t.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (t *TTrayIcon) SetOnClick(fn TNotifyEvent) {
	TrayIcon_SetOnClick(t.instance, fn)
}

// 设置双击事件。
func (t *TTrayIcon) SetOnDblClick(fn TNotifyEvent) {
	TrayIcon_SetOnDblClick(t.instance, fn)
}

// 设置鼠标移动事件。
func (t *TTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
	TrayIcon_SetOnMouseMove(t.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TTrayIcon) SetOnMouseUp(fn TMouseEvent) {
	TrayIcon_SetOnMouseUp(t.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TTrayIcon) SetOnMouseDown(fn TMouseEvent) {
	TrayIcon_SetOnMouseDown(t.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TTrayIcon) ComponentCount() int32 {
	return TrayIcon_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TTrayIcon) ComponentIndex() int32 {
	return TrayIcon_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TTrayIcon) SetComponentIndex(value int32) {
	TrayIcon_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTrayIcon) Owner() *TComponent {
	return AsComponent(TrayIcon_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TTrayIcon) Name() string {
	return TrayIcon_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTrayIcon) SetName(value string) {
	TrayIcon_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TTrayIcon) Tag() int {
	return TrayIcon_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TTrayIcon) SetTag(value int) {
	TrayIcon_SetTag(t.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TTrayIcon) Components(AIndex int32) *TComponent {
	return AsComponent(TrayIcon_GetComponents(t.instance, AIndex))
}
