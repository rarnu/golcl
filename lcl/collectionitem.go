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

type TCollectionItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewCollectionItem(AOwner *TCollection) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = CollectionItem_Create(CheckPtr(AOwner))
    c.ptr = unsafe.Pointer(c.instance)
    setFinalizer(c, (*TCollectionItem).Free)
    return c
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsCollectionItem(obj interface{}) *TCollectionItem {
    instance, ptr := getInstance(obj)
    if instance == 0 {
        return nil
    }
    return &TCollectionItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsCollectionItem.
func CollectionItemFromInst(inst uintptr) *TCollectionItem {
    return AsCollectionItem(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsCollectionItem.
func CollectionItemFromObj(obj IObject) *TCollectionItem {
    return AsCollectionItem(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCollectionItem.
func CollectionItemFromUnsafePointer(ptr unsafe.Pointer) *TCollectionItem {
    return AsCollectionItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (c *TCollectionItem) Free() {
    if c.instance != 0 {
        CollectionItem_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TCollectionItem) Instance() uintptr {
    return c.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TCollectionItem) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TCollectionItem) IsValid() bool {
    return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TCollectionItem) Is() TIs {
    return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TCollectionItem) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TCollectionItemClass() TClass {
    return CollectionItem_StaticClassType()
}

// 获取类名路径。
//
// Get the class name path.
func (c *TCollectionItem) GetNamePath() string {
    return CollectionItem_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCollectionItem) Assign(Source IObject) {
    CollectionItem_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TCollectionItem) ClassType() TClass {
    return CollectionItem_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCollectionItem) ClassName() string {
    return CollectionItem_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCollectionItem) InstanceSize() int32 {
    return CollectionItem_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCollectionItem) InheritsFrom(AClass TClass) bool {
    return CollectionItem_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCollectionItem) Equals(Obj IObject) bool {
    return CollectionItem_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCollectionItem) GetHashCode() int32 {
    return CollectionItem_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TCollectionItem) ToString() string {
    return CollectionItem_ToString(c.instance)
}

func (c *TCollectionItem) Collection() *TCollection {
    return AsCollection(CollectionItem_GetCollection(c.instance))
}

func (c *TCollectionItem) SetCollection(value *TCollection) {
    CollectionItem_SetCollection(c.instance, CheckPtr(value))
}

func (c *TCollectionItem) Index() int32 {
    return CollectionItem_GetIndex(c.instance)
}

func (c *TCollectionItem) SetIndex(value int32) {
    CollectionItem_SetIndex(c.instance, value)
}

func (c *TCollectionItem) DisplayName() string {
    return CollectionItem_GetDisplayName(c.instance)
}

func (c *TCollectionItem) SetDisplayName(value string) {
    CollectionItem_SetDisplayName(c.instance, value)
}

