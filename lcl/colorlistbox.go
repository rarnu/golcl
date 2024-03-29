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

type TColorListBox struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewColorListBox(owner IComponent) *TColorListBox {
	c := new(TColorListBox)
	c.instance = ColorListBox_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsColorListBox(obj any) *TColorListBox {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TColorListBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsColorListBox.
func ColorListBoxFromInst(inst uintptr) *TColorListBox {
	return AsColorListBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsColorListBox.
func ColorListBoxFromObj(obj IObject) *TColorListBox {
	return AsColorListBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsColorListBox.
func ColorListBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorListBox {
	return AsColorListBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TColorListBox) Free() {
	if c.instance != 0 {
		ColorListBox_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TColorListBox) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TColorListBox) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TColorListBox) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TColorListBox) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TColorListBox) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TColorListBoxClass() TClass {
	return ColorListBox_StaticClassType()
}

func (c *TColorListBox) AddItem(Item string, AObject IObject) {
	ColorListBox_AddItem(c.instance, Item, CheckPtr(AObject))
}

// 清除。
func (c *TColorListBox) Clear() {
	ColorListBox_Clear(c.instance)
}

// 清除选择。
func (c *TColorListBox) ClearSelection() {
	ColorListBox_ClearSelection(c.instance)
}

// 删除选择的。
func (c *TColorListBox) DeleteSelected() {
	ColorListBox_DeleteSelected(c.instance)
}

func (c *TColorListBox) ItemAtPos(Pos TPoint, Existing bool) int32 {
	return ColorListBox_ItemAtPos(c.instance, Pos, Existing)
}

func (c *TColorListBox) ItemRect(Index int32) TRect {
	return ColorListBox_ItemRect(c.instance, Index)
}

// 全选。
func (c *TColorListBox) SelectAll() {
	ColorListBox_SelectAll(c.instance)
}

// 是否可以获得焦点。
func (c *TColorListBox) CanFocus() bool {
	return ColorListBox_CanFocus(c.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TColorListBox) ContainsControl(Control IControl) bool {
	return ColorListBox_ContainsControl(c.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TColorListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(ColorListBox_ControlAtPos(c.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TColorListBox) DisableAlign() {
	ColorListBox_DisableAlign(c.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (c *TColorListBox) EnableAlign() {
	ColorListBox_EnableAlign(c.instance)
}

// 查找子控件。
//
// Find sub controls.
func (c *TColorListBox) FindChildControl(ControlName string) *TControl {
	return AsControl(ColorListBox_FindChildControl(c.instance, ControlName))
}

func (c *TColorListBox) FlipChildren(AllLevels bool) {
	ColorListBox_FlipChildren(c.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (c *TColorListBox) Focused() bool {
	return ColorListBox_Focused(c.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TColorListBox) HandleAllocated() bool {
	return ColorListBox_HandleAllocated(c.instance)
}

// 插入一个控件。
//
// Insert a control.
func (c *TColorListBox) InsertControl(AControl IControl) {
	ColorListBox_InsertControl(c.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (c *TColorListBox) Invalidate() {
	ColorListBox_Invalidate(c.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (c *TColorListBox) PaintTo(DC HDC, X int32, Y int32) {
	ColorListBox_PaintTo(c.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (c *TColorListBox) RemoveControl(AControl IControl) {
	ColorListBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (c *TColorListBox) Realign() {
	ColorListBox_Realign(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TColorListBox) Repaint() {
	ColorListBox_Repaint(c.instance)
}

// 按比例缩放。
//
// Scale by.
func (c *TColorListBox) ScaleBy(M int32, D int32) {
	ColorListBox_ScaleBy(c.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (c *TColorListBox) ScrollBy(DeltaX int32, DeltaY int32) {
	ColorListBox_ScrollBy(c.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TColorListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ColorListBox_SetBounds(c.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (c *TColorListBox) SetFocus() {
	ColorListBox_SetFocus(c.instance)
}

// 控件更新。
//
// Update.
func (c *TColorListBox) Update() {
	ColorListBox_Update(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TColorListBox) BringToFront() {
	ColorListBox_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TColorListBox) ClientToScreen(Point TPoint) TPoint {
	return ColorListBox_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TColorListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ColorListBox_ClientToParent(c.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TColorListBox) Dragging() bool {
	return ColorListBox_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TColorListBox) HasParent() bool {
	return ColorListBox_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TColorListBox) Hide() {
	ColorListBox_Hide(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TColorListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ColorListBox_Perform(c.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TColorListBox) Refresh() {
	ColorListBox_Refresh(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TColorListBox) ScreenToClient(Point TPoint) TPoint {
	return ColorListBox_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TColorListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ColorListBox_ParentToClient(c.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TColorListBox) SendToBack() {
	ColorListBox_SendToBack(c.instance)
}

// 显示控件。
//
// Show control.
func (c *TColorListBox) Show() {
	ColorListBox_Show(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TColorListBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ColorListBox_GetTextBuf(c.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TColorListBox) GetTextLen() int32 {
	return ColorListBox_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TColorListBox) SetTextBuf(Buffer string) {
	ColorListBox_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TColorListBox) FindComponent(AName string) *TComponent {
	return AsComponent(ColorListBox_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TColorListBox) GetNamePath() string {
	return ColorListBox_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TColorListBox) Assign(Source IObject) {
	ColorListBox_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TColorListBox) ClassType() TClass {
	return ColorListBox_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TColorListBox) ClassName() string {
	return ColorListBox_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TColorListBox) InstanceSize() int32 {
	return ColorListBox_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TColorListBox) InheritsFrom(AClass TClass) bool {
	return ColorListBox_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TColorListBox) Equals(Obj IObject) bool {
	return ColorListBox_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TColorListBox) GetHashCode() int32 {
	return ColorListBox_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TColorListBox) ToString() string {
	return ColorListBox_ToString(c.instance)
}

func (c *TColorListBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ColorListBox_AnchorToNeighbour(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (c *TColorListBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ColorListBox_AnchorParallel(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TColorListBox) AnchorHorizontalCenterTo(ASibling IControl) {
	ColorListBox_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TColorListBox) AnchorVerticalCenterTo(ASibling IControl) {
	ColorListBox_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TColorListBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ColorListBox_AnchorSame(c.instance, ASide, CheckPtr(ASibling))
}

func (c *TColorListBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ColorListBox_AnchorAsAlign(c.instance, ATheAlign, ASpace)
}

func (c *TColorListBox) AnchorClient(ASpace int32) {
	ColorListBox_AnchorClient(c.instance, ASpace)
}

func (c *TColorListBox) ScaleDesignToForm(ASize int32) int32 {
	return ColorListBox_ScaleDesignToForm(c.instance, ASize)
}

func (c *TColorListBox) ScaleFormToDesign(ASize int32) int32 {
	return ColorListBox_ScaleFormToDesign(c.instance, ASize)
}

func (c *TColorListBox) Scale96ToForm(ASize int32) int32 {
	return ColorListBox_Scale96ToForm(c.instance, ASize)
}

func (c *TColorListBox) ScaleFormTo96(ASize int32) int32 {
	return ColorListBox_ScaleFormTo96(c.instance, ASize)
}

func (c *TColorListBox) Scale96ToFont(ASize int32) int32 {
	return ColorListBox_Scale96ToFont(c.instance, ASize)
}

func (c *TColorListBox) ScaleFontTo96(ASize int32) int32 {
	return ColorListBox_ScaleFontTo96(c.instance, ASize)
}

func (c *TColorListBox) ScaleScreenToFont(ASize int32) int32 {
	return ColorListBox_ScaleScreenToFont(c.instance, ASize)
}

func (c *TColorListBox) ScaleFontToScreen(ASize int32) int32 {
	return ColorListBox_ScaleFontToScreen(c.instance, ASize)
}

func (c *TColorListBox) Scale96ToScreen(ASize int32) int32 {
	return ColorListBox_Scale96ToScreen(c.instance, ASize)
}

func (c *TColorListBox) ScaleScreenTo96(ASize int32) int32 {
	return ColorListBox_ScaleScreenTo96(c.instance, ASize)
}

func (c *TColorListBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ColorListBox_AutoAdjustLayout(c.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (c *TColorListBox) FixDesignFontsPPI(ADesignTimePPI int32) {
	ColorListBox_FixDesignFontsPPI(c.instance, ADesignTimePPI)
}

func (c *TColorListBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ColorListBox_ScaleFontsPPI(c.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TColorListBox) Align() TAlign {
	return ColorListBox_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TColorListBox) SetAlign(value TAlign) {
	ColorListBox_SetAlign(c.instance, value)
}

func (c *TColorListBox) DefaultColorColor() TColor {
	return ColorListBox_GetDefaultColorColor(c.instance)
}

func (c *TColorListBox) SetDefaultColorColor(value TColor) {
	ColorListBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorListBox) NoneColorColor() TColor {
	return ColorListBox_GetNoneColorColor(c.instance)
}

func (c *TColorListBox) SetNoneColorColor(value TColor) {
	ColorListBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorListBox) Selected() TColor {
	return ColorListBox_GetSelected(c.instance)
}

func (c *TColorListBox) SetSelected(value TColor) {
	ColorListBox_SetSelected(c.instance, value)
}

func (c *TColorListBox) Style() TColorBoxStyle {
	return ColorListBox_GetStyle(c.instance)
}

func (c *TColorListBox) SetStyle(value TColorBoxStyle) {
	ColorListBox_SetStyle(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TColorListBox) Anchors() TAnchors {
	return ColorListBox_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TColorListBox) SetAnchors(value TAnchors) {
	ColorListBox_SetAnchors(c.instance, value)
}

func (c *TColorListBox) BiDiMode() TBiDiMode {
	return ColorListBox_GetBiDiMode(c.instance)
}

func (c *TColorListBox) SetBiDiMode(value TBiDiMode) {
	ColorListBox_SetBiDiMode(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TColorListBox) Color() TColor {
	return ColorListBox_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TColorListBox) SetColor(value TColor) {
	ColorListBox_SetColor(c.instance, value)
}

// 获取约束控件大小。
func (c *TColorListBox) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ColorListBox_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TColorListBox) SetConstraints(value *TSizeConstraints) {
	ColorListBox_SetConstraints(c.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TColorListBox) DoubleBuffered() bool {
	return ColorListBox_GetDoubleBuffered(c.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TColorListBox) SetDoubleBuffered(value bool) {
	ColorListBox_SetDoubleBuffered(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TColorListBox) Enabled() bool {
	return ColorListBox_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TColorListBox) SetEnabled(value bool) {
	ColorListBox_SetEnabled(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TColorListBox) Font() *TFont {
	return AsFont(ColorListBox_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TColorListBox) SetFont(value *TFont) {
	ColorListBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ItemHeight() int32 {
	return ColorListBox_GetItemHeight(c.instance)
}

func (c *TColorListBox) SetItemHeight(value int32) {
	ColorListBox_SetItemHeight(c.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (c *TColorListBox) ParentColor() bool {
	return ColorListBox_GetParentColor(c.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (c *TColorListBox) SetParentColor(value bool) {
	ColorListBox_SetParentColor(c.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TColorListBox) ParentDoubleBuffered() bool {
	return ColorListBox_GetParentDoubleBuffered(c.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TColorListBox) SetParentDoubleBuffered(value bool) {
	ColorListBox_SetParentDoubleBuffered(c.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TColorListBox) ParentFont() bool {
	return ColorListBox_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TColorListBox) SetParentFont(value bool) {
	ColorListBox_SetParentFont(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TColorListBox) ParentShowHint() bool {
	return ColorListBox_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TColorListBox) SetParentShowHint(value bool) {
	ColorListBox_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TColorListBox) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ColorListBox_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TColorListBox) SetPopupMenu(value IComponent) {
	ColorListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TColorListBox) ShowHint() bool {
	return ColorListBox_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TColorListBox) SetShowHint(value bool) {
	ColorListBox_SetShowHint(c.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TColorListBox) TabOrder() TTabOrder {
	return ColorListBox_GetTabOrder(c.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TColorListBox) SetTabOrder(value TTabOrder) {
	ColorListBox_SetTabOrder(c.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TColorListBox) TabStop() bool {
	return ColorListBox_GetTabStop(c.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TColorListBox) SetTabStop(value bool) {
	ColorListBox_SetTabStop(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TColorListBox) Visible() bool {
	return ColorListBox_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TColorListBox) SetVisible(value bool) {
	ColorListBox_SetVisible(c.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TColorListBox) SetOnClick(fn TNotifyEvent) {
	ColorListBox_SetOnClick(c.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (c *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
	ColorListBox_SetOnContextPopup(c.instance, fn)
}

// 设置双击事件。
func (c *TColorListBox) SetOnDblClick(fn TNotifyEvent) {
	ColorListBox_SetOnDblClick(c.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
	ColorListBox_SetOnDragDrop(c.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
	ColorListBox_SetOnDragOver(c.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (c *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
	ColorListBox_SetOnEndDrag(c.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (c *TColorListBox) SetOnEnter(fn TNotifyEvent) {
	ColorListBox_SetOnEnter(c.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (c *TColorListBox) SetOnExit(fn TNotifyEvent) {
	ColorListBox_SetOnExit(c.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TColorListBox) SetOnKeyDown(fn TKeyEvent) {
	ColorListBox_SetOnKeyDown(c.instance, fn)
}

// 设置键键下事件。
func (c *TColorListBox) SetOnKeyPress(fn TKeyPressEvent) {
	ColorListBox_SetOnKeyPress(c.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TColorListBox) SetOnKeyUp(fn TKeyEvent) {
	ColorListBox_SetOnKeyUp(c.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TColorListBox) SetOnMouseDown(fn TMouseEvent) {
	ColorListBox_SetOnMouseDown(c.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TColorListBox) SetOnMouseEnter(fn TNotifyEvent) {
	ColorListBox_SetOnMouseEnter(c.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TColorListBox) SetOnMouseLeave(fn TNotifyEvent) {
	ColorListBox_SetOnMouseLeave(c.instance, fn)
}

// 设置鼠标移动事件。
func (c *TColorListBox) SetOnMouseMove(fn TMouseMoveEvent) {
	ColorListBox_SetOnMouseMove(c.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TColorListBox) SetOnMouseUp(fn TMouseEvent) {
	ColorListBox_SetOnMouseUp(c.instance, fn)
}

// 获取画布。
func (c *TColorListBox) Canvas() *TCanvas {
	return AsCanvas(ColorListBox_GetCanvas(c.instance))
}

func (c *TColorListBox) Count() int32 {
	return ColorListBox_GetCount(c.instance)
}

func (c *TColorListBox) Items() *TStrings {
	return AsStrings(ColorListBox_GetItems(c.instance))
}

func (c *TColorListBox) SetItems(value IStrings) {
	ColorListBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TColorListBox) TopIndex() int32 {
	return ColorListBox_GetTopIndex(c.instance)
}

func (c *TColorListBox) SetTopIndex(value int32) {
	ColorListBox_SetTopIndex(c.instance, value)
}

func (c *TColorListBox) MultiSelect() bool {
	return ColorListBox_GetMultiSelect(c.instance)
}

func (c *TColorListBox) SetMultiSelect(value bool) {
	ColorListBox_SetMultiSelect(c.instance, value)
}

func (c *TColorListBox) SelCount() int32 {
	return ColorListBox_GetSelCount(c.instance)
}

func (c *TColorListBox) ItemIndex() int32 {
	return ColorListBox_GetItemIndex(c.instance)
}

func (c *TColorListBox) SetItemIndex(value int32) {
	ColorListBox_SetItemIndex(c.instance, value)
}

// 获取依靠客户端总数。
func (c *TColorListBox) DockClientCount() int32 {
	return ColorListBox_GetDockClientCount(c.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (c *TColorListBox) DockSite() bool {
	return ColorListBox_GetDockSite(c.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (c *TColorListBox) SetDockSite(value bool) {
	ColorListBox_SetDockSite(c.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TColorListBox) MouseInClient() bool {
	return ColorListBox_GetMouseInClient(c.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TColorListBox) VisibleDockClientCount() int32 {
	return ColorListBox_GetVisibleDockClientCount(c.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (c *TColorListBox) Brush() *TBrush {
	return AsBrush(ColorListBox_GetBrush(c.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (c *TColorListBox) ControlCount() int32 {
	return ColorListBox_GetControlCount(c.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TColorListBox) Handle() HWND {
	return ColorListBox_GetHandle(c.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TColorListBox) ParentWindow() HWND {
	return ColorListBox_GetParentWindow(c.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TColorListBox) SetParentWindow(value HWND) {
	ColorListBox_SetParentWindow(c.instance, value)
}

func (c *TColorListBox) Showing() bool {
	return ColorListBox_GetShowing(c.instance)
}

// 获取使用停靠管理。
func (c *TColorListBox) UseDockManager() bool {
	return ColorListBox_GetUseDockManager(c.instance)
}

// 设置使用停靠管理。
func (c *TColorListBox) SetUseDockManager(value bool) {
	ColorListBox_SetUseDockManager(c.instance, value)
}

func (c *TColorListBox) Action() *TAction {
	return AsAction(ColorListBox_GetAction(c.instance))
}

func (c *TColorListBox) SetAction(value IComponent) {
	ColorListBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorListBox) BoundsRect() TRect {
	return ColorListBox_GetBoundsRect(c.instance)
}

func (c *TColorListBox) SetBoundsRect(value TRect) {
	ColorListBox_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TColorListBox) ClientHeight() int32 {
	return ColorListBox_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TColorListBox) SetClientHeight(value int32) {
	ColorListBox_SetClientHeight(c.instance, value)
}

func (c *TColorListBox) ClientOrigin() TPoint {
	return ColorListBox_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TColorListBox) ClientRect() TRect {
	return ColorListBox_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TColorListBox) ClientWidth() int32 {
	return ColorListBox_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TColorListBox) SetClientWidth(value int32) {
	ColorListBox_SetClientWidth(c.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (c *TColorListBox) ControlState() TControlState {
	return ColorListBox_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TColorListBox) SetControlState(value TControlState) {
	ColorListBox_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TColorListBox) ControlStyle() TControlStyle {
	return ColorListBox_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TColorListBox) SetControlStyle(value TControlStyle) {
	ColorListBox_SetControlStyle(c.instance, value)
}

func (c *TColorListBox) Floating() bool {
	return ColorListBox_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TColorListBox) Parent() *TWinControl {
	return AsWinControl(ColorListBox_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TColorListBox) SetParent(value IWinControl) {
	ColorListBox_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TColorListBox) Left() int32 {
	return ColorListBox_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TColorListBox) SetLeft(value int32) {
	ColorListBox_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TColorListBox) Top() int32 {
	return ColorListBox_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TColorListBox) SetTop(value int32) {
	ColorListBox_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TColorListBox) Width() int32 {
	return ColorListBox_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TColorListBox) SetWidth(value int32) {
	ColorListBox_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TColorListBox) Height() int32 {
	return ColorListBox_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TColorListBox) SetHeight(value int32) {
	ColorListBox_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TColorListBox) Cursor() TCursor {
	return ColorListBox_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TColorListBox) SetCursor(value TCursor) {
	ColorListBox_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TColorListBox) Hint() string {
	return ColorListBox_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TColorListBox) SetHint(value string) {
	ColorListBox_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TColorListBox) ComponentCount() int32 {
	return ColorListBox_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TColorListBox) ComponentIndex() int32 {
	return ColorListBox_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TColorListBox) SetComponentIndex(value int32) {
	ColorListBox_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TColorListBox) Owner() *TComponent {
	return AsComponent(ColorListBox_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TColorListBox) Name() string {
	return ColorListBox_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TColorListBox) SetName(value string) {
	ColorListBox_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TColorListBox) Tag() int {
	return ColorListBox_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TColorListBox) SetTag(value int) {
	ColorListBox_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TColorListBox) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ColorListBox_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TColorListBox) SetAnchorSideLeft(value *TAnchorSide) {
	ColorListBox_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TColorListBox) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ColorListBox_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TColorListBox) SetAnchorSideTop(value *TAnchorSide) {
	ColorListBox_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TColorListBox) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ColorListBox_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TColorListBox) SetAnchorSideRight(value *TAnchorSide) {
	ColorListBox_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TColorListBox) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ColorListBox_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TColorListBox) SetAnchorSideBottom(value *TAnchorSide) {
	ColorListBox_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(ColorListBox_GetChildSizing(c.instance))
}

func (c *TColorListBox) SetChildSizing(value *TControlChildSizing) {
	ColorListBox_SetChildSizing(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TColorListBox) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ColorListBox_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TColorListBox) SetBorderSpacing(value *TControlBorderSpacing) {
	ColorListBox_SetBorderSpacing(c.instance, CheckPtr(value))
}

func (c *TColorListBox) Colors(Index int32) TColor {
	return ColorListBox_GetColors(c.instance, Index)
}

func (c *TColorListBox) ColorNames(Index int32) string {
	return ColorListBox_GetColorNames(c.instance, Index)
}

// 获取指定索引停靠客户端。
func (c *TColorListBox) DockClients(Index int32) *TControl {
	return AsControl(ColorListBox_GetDockClients(c.instance, Index))
}

// 获取指定索引子控件。
func (c *TColorListBox) Controls(Index int32) *TControl {
	return AsControl(ColorListBox_GetControls(c.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TColorListBox) Components(AIndex int32) *TComponent {
	return AsComponent(ColorListBox_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TColorListBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ColorListBox_GetAnchorSide(c.instance, AKind))
}
