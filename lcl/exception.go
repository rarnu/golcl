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

type Exception struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsException(obj interface{}) *Exception {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &Exception{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsException.
func ExceptionFromInst(inst uintptr) *Exception {
	return AsException(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsException.
func ExceptionFromObj(obj IObject) *Exception {
	return AsException(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsException.
func ExceptionFromUnsafePointer(ptr unsafe.Pointer) *Exception {
	return AsException(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
// 
// Return object instance pointer.
func (e *Exception) Instance() uintptr {
	return e.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (e *Exception) UnsafeAddr() unsafe.Pointer {
	return e.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (e *Exception) IsValid() bool {
	return e.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (e *Exception) Is() TIs {
	return TIs(e.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (e *Exception) As() TAs {
//    return TAs(e.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func ExceptionClass() TClass {
	return Exception_StaticClassType()
}

// 文本类信息。
//
// Text information.
func (e *Exception) ToString() string {
	return Exception_ToString(e.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (e *Exception) ClassType() TClass {
	return Exception_ClassType(e.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (e *Exception) ClassName() string {
	return Exception_ClassName(e.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (e *Exception) InstanceSize() int32 {
	return Exception_InstanceSize(e.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (e *Exception) InheritsFrom(AClass TClass) bool {
	return Exception_InheritsFrom(e.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (e *Exception) Equals(Obj IObject) bool {
	return Exception_Equals(e.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (e *Exception) GetHashCode() int32 {
	return Exception_GetHashCode(e.instance)
}

// 获取异常消息。
func (e *Exception) Message() string {
	return Exception_GetMessage(e.instance)
}

// 设置异常消息。
func (e *Exception) SetMessage(value string) {
	Exception_SetMessage(e.instance, value)
}

