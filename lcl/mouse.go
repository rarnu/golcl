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

type TMouse struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewMouse() *TMouse {
	m := new(TMouse)
	m.instance = Mouse_Create()
	m.ptr = unsafe.Pointer(m.instance)
	setFinalizer(m, (*TMouse).Free)
	return m
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsMouse(obj any) *TMouse {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TMouse{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsMouse.
func MouseFromInst(inst uintptr) *TMouse {
	return AsMouse(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsMouse.
func MouseFromObj(obj IObject) *TMouse {
	return AsMouse(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMouse.
func MouseFromUnsafePointer(ptr unsafe.Pointer) *TMouse {
	return AsMouse(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (m *TMouse) Free() {
	if m.instance != 0 {
		Mouse_Free(m.instance)
		m.instance, m.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (m *TMouse) Instance() uintptr {
	return m.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (m *TMouse) UnsafeAddr() unsafe.Pointer {
	return m.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (m *TMouse) IsValid() bool {
	return m.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (m *TMouse) Is() TIs {
	return TIs(m.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (m *TMouse) As() TAs {
//    return TAs(m.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TMouseClass() TClass {
	return Mouse_StaticClassType()
}

// 获取类的类型信息。
//
// Get class type information.
func (m *TMouse) ClassType() TClass {
	return Mouse_ClassType(m.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TMouse) ClassName() string {
	return Mouse_ClassName(m.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TMouse) InstanceSize() int32 {
	return Mouse_InstanceSize(m.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TMouse) InheritsFrom(AClass TClass) bool {
	return Mouse_InheritsFrom(m.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (m *TMouse) Equals(Obj IObject) bool {
	return Mouse_Equals(m.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TMouse) GetHashCode() int32 {
	return Mouse_GetHashCode(m.instance)
}

// 文本类信息。
//
// Text information.
func (m *TMouse) ToString() string {
	return Mouse_ToString(m.instance)
}

func (m *TMouse) Capture() HWND {
	return Mouse_GetCapture(m.instance)
}

func (m *TMouse) SetCapture(value HWND) {
	Mouse_SetCapture(m.instance, value)
}

func (m *TMouse) CursorPos() TPoint {
	return Mouse_GetCursorPos(m.instance)
}

func (m *TMouse) SetCursorPos(value TPoint) {
	Mouse_SetCursorPos(m.instance, value)
}

func (m *TMouse) IsDragging() bool {
	return Mouse_GetIsDragging(m.instance)
}

func (m *TMouse) WheelScrollLines() int32 {
	return Mouse_GetWheelScrollLines(m.instance)
}
