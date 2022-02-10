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

type TImageList struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewImageList(owner IComponent) *TImageList {
	i := new(TImageList)
	i.instance = ImageList_Create(CheckPtr(owner))
	i.ptr = unsafe.Pointer(i.instance)
	return i
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsImageList(obj interface{}) *TImageList {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TImageList{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsImageList.
func ImageListFromInst(inst uintptr) *TImageList {
	return AsImageList(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsImageList.
func ImageListFromObj(obj IObject) *TImageList {
	return AsImageList(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsImageList.
func ImageListFromUnsafePointer(ptr unsafe.Pointer) *TImageList {
	return AsImageList(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (i *TImageList) Free() {
	if i.instance != 0 {
		ImageList_Free(i.instance)
		i.instance, i.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (i *TImageList) Instance() uintptr {
	return i.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (i *TImageList) UnsafeAddr() unsafe.Pointer {
	return i.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (i *TImageList) IsValid() bool {
	return i.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (i *TImageList) Is() TIs {
	return TIs(i.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (i *TImageList) As() TAs {
//    return TAs(i.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TImageListClass() TClass {
	return ImageList_StaticClassType()
}

func (i *TImageList) StretchDraw(ACanvas *TCanvas, AIndex int32, ARect TRect, AEnabled bool) {
	ImageList_StretchDraw(i.instance, CheckPtr(ACanvas), AIndex, ARect, AEnabled)
}

func (i *TImageList) AddSliced(Image *TBitmap, AHorizontalCount int32, AVerticalCount int32) int32 {
	return ImageList_AddSliced(i.instance, CheckPtr(Image), AHorizontalCount, AVerticalCount)
}

func (i *TImageList) GetHotSpot() TPoint {
	return ImageList_GetHotSpot(i.instance)
}

func (i *TImageList) HideDragImage() {
	ImageList_HideDragImage(i.instance)
}

func (i *TImageList) ShowDragImage() {
	ImageList_ShowDragImage(i.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (i *TImageList) Assign(Source IObject) {
	ImageList_Assign(i.instance, CheckPtr(Source))
}

func (i *TImageList) Add(Image *TBitmap, Mask *TBitmap) int32 {
	return ImageList_Add(i.instance, CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) AddIcon(Image *TIcon) int32 {
	return ImageList_AddIcon(i.instance, CheckPtr(Image))
}

func (i *TImageList) AddImages(Value IComponent) {
	ImageList_AddImages(i.instance, CheckPtr(Value))
}

func (i *TImageList) AddMasked(Image *TBitmap, MaskColor TColor) int32 {
	return ImageList_AddMasked(i.instance, CheckPtr(Image), MaskColor)
}

// 清除。
func (i *TImageList) Clear() {
	ImageList_Clear(i.instance)
}

func (i *TImageList) Delete(Index int32) {
	ImageList_Delete(i.instance, Index)
}

func (i *TImageList) Insert(Index int32, Image *TBitmap, Mask *TBitmap) {
	ImageList_Insert(i.instance, Index, CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) InsertIcon(Index int32, Image *TIcon) {
	ImageList_InsertIcon(i.instance, Index, CheckPtr(Image))
}

func (i *TImageList) InsertMasked(Index int32, Image *TBitmap, MaskColor TColor) {
	ImageList_InsertMasked(i.instance, Index, CheckPtr(Image), MaskColor)
}

func (i *TImageList) Move(CurIndex int32, NewIndex int32) {
	ImageList_Move(i.instance, CurIndex, NewIndex)
}

func (i *TImageList) Replace(Index int32, Image *TBitmap, Mask *TBitmap) {
	ImageList_Replace(i.instance, Index, CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) ReplaceMasked(Index int32, NewImage *TBitmap, MaskColor TColor) {
	ImageList_ReplaceMasked(i.instance, Index, CheckPtr(NewImage), MaskColor)
}

func (i *TImageList) BeginUpdate() {
	ImageList_BeginUpdate(i.instance)
}

func (i *TImageList) EndUpdate() {
	ImageList_EndUpdate(i.instance)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (i *TImageList) FindComponent(AName string) *TComponent {
	return AsComponent(ImageList_FindComponent(i.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (i *TImageList) GetNamePath() string {
	return ImageList_GetNamePath(i.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (i *TImageList) HasParent() bool {
	return ImageList_HasParent(i.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (i *TImageList) ClassType() TClass {
	return ImageList_ClassType(i.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (i *TImageList) ClassName() string {
	return ImageList_ClassName(i.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (i *TImageList) InstanceSize() int32 {
	return ImageList_InstanceSize(i.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (i *TImageList) InheritsFrom(AClass TClass) bool {
	return ImageList_InheritsFrom(i.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (i *TImageList) Equals(Obj IObject) bool {
	return ImageList_Equals(i.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (i *TImageList) GetHashCode() int32 {
	return ImageList_GetHashCode(i.instance)
}

// 文本类信息。
//
// Text information.
func (i *TImageList) ToString() string {
	return ImageList_ToString(i.instance)
}

func (i *TImageList) Scaled() bool {
	return ImageList_GetScaled(i.instance)
}

func (i *TImageList) SetScaled(value bool) {
	ImageList_SetScaled(i.instance, value)
}

func (i *TImageList) ShareImages() bool {
	return ImageList_GetShareImages(i.instance)
}

func (i *TImageList) SetShareImages(value bool) {
	ImageList_SetShareImages(i.instance, value)
}

func (i *TImageList) Count() int32 {
	return ImageList_GetCount(i.instance)
}

func (i *TImageList) BlendColor() TColor {
	return ImageList_GetBlendColor(i.instance)
}

func (i *TImageList) SetBlendColor(value TColor) {
	ImageList_SetBlendColor(i.instance, value)
}

func (i *TImageList) BkColor() TColor {
	return ImageList_GetBkColor(i.instance)
}

func (i *TImageList) SetBkColor(value TColor) {
	ImageList_SetBkColor(i.instance, value)
}

func (i *TImageList) AllocBy() int32 {
	return ImageList_GetAllocBy(i.instance)
}

func (i *TImageList) SetAllocBy(value int32) {
	ImageList_SetAllocBy(i.instance, value)
}

func (i *TImageList) DrawingStyle() TDrawingStyle {
	return ImageList_GetDrawingStyle(i.instance)
}

func (i *TImageList) SetDrawingStyle(value TDrawingStyle) {
	ImageList_SetDrawingStyle(i.instance, value)
}

// 获取高度。
//
// Get height.
func (i *TImageList) Height() int32 {
	return ImageList_GetHeight(i.instance)
}

// 设置高度。
//
// Set height.
func (i *TImageList) SetHeight(value int32) {
	ImageList_SetHeight(i.instance, value)
}

func (i *TImageList) ImageType() TImageType {
	return ImageList_GetImageType(i.instance)
}

func (i *TImageList) SetImageType(value TImageType) {
	ImageList_SetImageType(i.instance, value)
}

func (i *TImageList) Masked() bool {
	return ImageList_GetMasked(i.instance)
}

func (i *TImageList) SetMasked(value bool) {
	ImageList_SetMasked(i.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (i *TImageList) SetOnChange(fn TNotifyEvent) {
	ImageList_SetOnChange(i.instance, fn)
}

// 获取宽度。
//
// Get width.
func (i *TImageList) Width() int32 {
	return ImageList_GetWidth(i.instance)
}

// 设置宽度。
//
// Set width.
func (i *TImageList) SetWidth(value int32) {
	ImageList_SetWidth(i.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (i *TImageList) DragCursor() TCursor {
	return ImageList_GetDragCursor(i.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (i *TImageList) SetDragCursor(value TCursor) {
	ImageList_SetDragCursor(i.instance, value)
}

// 获取是否在拖拽中。
//
// Get Is it in the middle of dragging.
func (i *TImageList) Dragging() bool {
	return ImageList_GetDragging(i.instance)
}

// 获取组件总数。
//
// Get the total number of components.
func (i *TImageList) ComponentCount() int32 {
	return ImageList_GetComponentCount(i.instance)
}

// 获取组件索引。
//
// Get component index.
func (i *TImageList) ComponentIndex() int32 {
	return ImageList_GetComponentIndex(i.instance)
}

// 设置组件索引。
//
// Set component index.
func (i *TImageList) SetComponentIndex(value int32) {
	ImageList_SetComponentIndex(i.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (i *TImageList) Owner() *TComponent {
	return AsComponent(ImageList_GetOwner(i.instance))
}

// 获取组件名称。
//
// Get the component name.
func (i *TImageList) Name() string {
	return ImageList_GetName(i.instance)
}

// 设置组件名称。
//
// Set the component name.
func (i *TImageList) SetName(value string) {
	ImageList_SetName(i.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (i *TImageList) Tag() int {
	return ImageList_GetTag(i.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (i *TImageList) SetTag(value int) {
	ImageList_SetTag(i.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (i *TImageList) Components(AIndex int32) *TComponent {
	return AsComponent(ImageList_GetComponents(i.instance, AIndex))
}
