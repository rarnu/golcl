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

type TCollection struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewCollection(AOwner *TCollectionItem) *TCollection {
	c := new(TCollection)
	c.instance = Collection_Create(CheckPtr(AOwner))
	c.ptr = unsafe.Pointer(c.instance)
	setFinalizer(c, (*TCollection).Free)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsCollection(obj interface{}) *TCollection {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TCollection{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsCollection.
func CollectionFromInst(inst uintptr) *TCollection {
	return AsCollection(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsCollection.
func CollectionFromObj(obj IObject) *TCollection {
	return AsCollection(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCollection.
func CollectionFromUnsafePointer(ptr unsafe.Pointer) *TCollection {
	return AsCollection(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TCollection) Free() {
	if c.instance != 0 {
		Collection_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TCollection) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TCollection) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TCollection) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TCollection) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TCollection) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TCollectionClass() TClass {
	return Collection_StaticClassType()
}

// 组件所有者。
//
// component owner.
func (c *TCollection) Owner() *TObject {
	return AsObject(Collection_Owner(c.instance))
}

func (c *TCollection) Add() *TCollectionItem {
	return AsCollectionItem(Collection_Add(c.instance))
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCollection) Assign(Source IObject) {
	Collection_Assign(c.instance, CheckPtr(Source))
}

func (c *TCollection) BeginUpdate() {
	Collection_BeginUpdate(c.instance)
}

// 清除。
func (c *TCollection) Clear() {
	Collection_Clear(c.instance)
}

func (c *TCollection) Delete(Index int32) {
	Collection_Delete(c.instance, Index)
}

func (c *TCollection) EndUpdate() {
	Collection_EndUpdate(c.instance)
}

func (c *TCollection) FindItemID(ID int32) *TCollectionItem {
	return AsCollectionItem(Collection_FindItemID(c.instance, ID))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TCollection) GetNamePath() string {
	return Collection_GetNamePath(c.instance)
}

func (c *TCollection) Insert(Index int32) *TCollectionItem {
	return AsCollectionItem(Collection_Insert(c.instance, Index))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TCollection) ClassType() TClass {
	return Collection_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCollection) ClassName() string {
	return Collection_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCollection) InstanceSize() int32 {
	return Collection_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCollection) InheritsFrom(AClass TClass) bool {
	return Collection_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCollection) Equals(Obj IObject) bool {
	return Collection_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCollection) GetHashCode() int32 {
	return Collection_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TCollection) ToString() string {
	return Collection_ToString(c.instance)
}

func (c *TCollection) Count() int32 {
	return Collection_GetCount(c.instance)
}

func (c *TCollection) Items(Index int32) *TCollectionItem {
	return AsCollectionItem(Collection_GetItems(c.instance, Index))
}

func (c *TCollection) SetItems(Index int32, value *TCollectionItem) {
	Collection_SetItems(c.instance, Index, CheckPtr(value))
}
