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

type TParaAttributes struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsParaAttributes(obj any) *TParaAttributes {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TParaAttributes{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsParaAttributes.
func ParaAttributesFromInst(inst uintptr) *TParaAttributes {
	return AsParaAttributes(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsParaAttributes.
func ParaAttributesFromObj(obj IObject) *TParaAttributes {
	return AsParaAttributes(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsParaAttributes.
func ParaAttributesFromUnsafePointer(ptr unsafe.Pointer) *TParaAttributes {
	return AsParaAttributes(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TParaAttributes) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TParaAttributes) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TParaAttributes) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TParaAttributes) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TParaAttributes) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TParaAttributesClass() TClass {
	return ParaAttributes_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TParaAttributes) Assign(Source IObject) {
	ParaAttributes_Assign(p.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TParaAttributes) GetNamePath() string {
	return ParaAttributes_GetNamePath(p.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TParaAttributes) ClassType() TClass {
	return ParaAttributes_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TParaAttributes) ClassName() string {
	return ParaAttributes_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TParaAttributes) InstanceSize() int32 {
	return ParaAttributes_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TParaAttributes) InheritsFrom(AClass TClass) bool {
	return ParaAttributes_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TParaAttributes) Equals(Obj IObject) bool {
	return ParaAttributes_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TParaAttributes) GetHashCode() int32 {
	return ParaAttributes_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TParaAttributes) ToString() string {
	return ParaAttributes_ToString(p.instance)
}

// 获取文字对齐。
//
// Get Text alignment.
func (p *TParaAttributes) Alignment() TAlignment {
	return ParaAttributes_GetAlignment(p.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (p *TParaAttributes) SetAlignment(value TAlignment) {
	ParaAttributes_SetAlignment(p.instance, value)
}

func (p *TParaAttributes) FirstIndent() int32 {
	return ParaAttributes_GetFirstIndent(p.instance)
}

func (p *TParaAttributes) SetFirstIndent(value int32) {
	ParaAttributes_SetFirstIndent(p.instance, value)
}

func (p *TParaAttributes) LeftIndent() int32 {
	return ParaAttributes_GetLeftIndent(p.instance)
}

func (p *TParaAttributes) SetLeftIndent(value int32) {
	ParaAttributes_SetLeftIndent(p.instance, value)
}

func (p *TParaAttributes) Numbering() TNumberingStyle {
	return ParaAttributes_GetNumbering(p.instance)
}

func (p *TParaAttributes) SetNumbering(value TNumberingStyle) {
	ParaAttributes_SetNumbering(p.instance, value)
}

func (p *TParaAttributes) RightIndent() int32 {
	return ParaAttributes_GetRightIndent(p.instance)
}

func (p *TParaAttributes) SetRightIndent(value int32) {
	ParaAttributes_SetRightIndent(p.instance, value)
}

func (p *TParaAttributes) TabCount() int32 {
	return ParaAttributes_GetTabCount(p.instance)
}

func (p *TParaAttributes) SetTabCount(value int32) {
	ParaAttributes_SetTabCount(p.instance, value)
}

func (p *TParaAttributes) Tab(Index uint8) int32 {
	return ParaAttributes_GetTab(p.instance, Index)
}

func (p *TParaAttributes) SetTab(Index uint8, value int32) {
	ParaAttributes_SetTab(p.instance, Index, value)
}
