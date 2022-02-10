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

type TStatusPanel struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewStatusPanel(AOwner *TCollection) *TStatusPanel {
	s := new(TStatusPanel)
	s.instance = StatusPanel_Create(CheckPtr(AOwner))
	s.ptr = unsafe.Pointer(s.instance)
	setFinalizer(s, (*TStatusPanel).Free)
	return s
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsStatusPanel(obj interface{}) *TStatusPanel {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TStatusPanel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsStatusPanel.
func StatusPanelFromInst(inst uintptr) *TStatusPanel {
	return AsStatusPanel(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsStatusPanel.
func StatusPanelFromObj(obj IObject) *TStatusPanel {
	return AsStatusPanel(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsStatusPanel.
func StatusPanelFromUnsafePointer(ptr unsafe.Pointer) *TStatusPanel {
	return AsStatusPanel(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (s *TStatusPanel) Free() {
	if s.instance != 0 {
		StatusPanel_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TStatusPanel) Instance() uintptr {
	return s.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TStatusPanel) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TStatusPanel) IsValid() bool {
	return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TStatusPanel) Is() TIs {
	return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TStatusPanel) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TStatusPanelClass() TClass {
	return StatusPanel_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TStatusPanel) Assign(Source IObject) {
	StatusPanel_Assign(s.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TStatusPanel) GetNamePath() string {
	return StatusPanel_GetNamePath(s.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TStatusPanel) ClassType() TClass {
	return StatusPanel_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TStatusPanel) ClassName() string {
	return StatusPanel_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TStatusPanel) InstanceSize() int32 {
	return StatusPanel_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TStatusPanel) InheritsFrom(AClass TClass) bool {
	return StatusPanel_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TStatusPanel) Equals(Obj IObject) bool {
	return StatusPanel_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TStatusPanel) GetHashCode() int32 {
	return StatusPanel_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TStatusPanel) ToString() string {
	return StatusPanel_ToString(s.instance)
}

// 获取文字对齐。
//
// Get Text alignment.
func (s *TStatusPanel) Alignment() TAlignment {
	return StatusPanel_GetAlignment(s.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (s *TStatusPanel) SetAlignment(value TAlignment) {
	StatusPanel_SetAlignment(s.instance, value)
}

func (s *TStatusPanel) BiDiMode() TBiDiMode {
	return StatusPanel_GetBiDiMode(s.instance)
}

func (s *TStatusPanel) SetBiDiMode(value TBiDiMode) {
	StatusPanel_SetBiDiMode(s.instance, value)
}

func (s *TStatusPanel) Style() TStatusPanelStyle {
	return StatusPanel_GetStyle(s.instance)
}

func (s *TStatusPanel) SetStyle(value TStatusPanelStyle) {
	StatusPanel_SetStyle(s.instance, value)
}

// 获取文本。
func (s *TStatusPanel) Text() string {
	return StatusPanel_GetText(s.instance)
}

// 设置文本。
func (s *TStatusPanel) SetText(value string) {
	StatusPanel_SetText(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TStatusPanel) Width() int32 {
	return StatusPanel_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TStatusPanel) SetWidth(value int32) {
	StatusPanel_SetWidth(s.instance, value)
}

func (s *TStatusPanel) Collection() *TCollection {
	return AsCollection(StatusPanel_GetCollection(s.instance))
}

func (s *TStatusPanel) SetCollection(value *TCollection) {
	StatusPanel_SetCollection(s.instance, CheckPtr(value))
}

func (s *TStatusPanel) Index() int32 {
	return StatusPanel_GetIndex(s.instance)
}

func (s *TStatusPanel) SetIndex(value int32) {
	StatusPanel_SetIndex(s.instance, value)
}

func (s *TStatusPanel) DisplayName() string {
	return StatusPanel_GetDisplayName(s.instance)
}

func (s *TStatusPanel) SetDisplayName(value string) {
	StatusPanel_SetDisplayName(s.instance, value)
}
