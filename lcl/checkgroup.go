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

type TCheckGroup struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewCheckGroup(owner IComponent) *TCheckGroup {
	c := new(TCheckGroup)
	c.instance = CheckGroup_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsCheckGroup(obj any) *TCheckGroup {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TCheckGroup{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsCheckGroup.
func CheckGroupFromInst(inst uintptr) *TCheckGroup {
	return AsCheckGroup(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsCheckGroup.
func CheckGroupFromObj(obj IObject) *TCheckGroup {
	return AsCheckGroup(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCheckGroup.
func CheckGroupFromUnsafePointer(ptr unsafe.Pointer) *TCheckGroup {
	return AsCheckGroup(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TCheckGroup) Free() {
	if c.instance != 0 {
		CheckGroup_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TCheckGroup) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TCheckGroup) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TCheckGroup) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TCheckGroup) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TCheckGroup) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TCheckGroupClass() TClass {
	return CheckGroup_StaticClassType()
}

func (c *TCheckGroup) FlipChildren(AllLevels bool) {
	CheckGroup_FlipChildren(c.instance, AllLevels)
}

func (c *TCheckGroup) Rows() int32 {
	return CheckGroup_Rows(c.instance)
}

// 是否可以获得焦点。
func (c *TCheckGroup) CanFocus() bool {
	return CheckGroup_CanFocus(c.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TCheckGroup) ContainsControl(Control IControl) bool {
	return CheckGroup_ContainsControl(c.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TCheckGroup) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(CheckGroup_ControlAtPos(c.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TCheckGroup) DisableAlign() {
	CheckGroup_DisableAlign(c.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (c *TCheckGroup) EnableAlign() {
	CheckGroup_EnableAlign(c.instance)
}

// 查找子控件。
//
// Find sub controls.
func (c *TCheckGroup) FindChildControl(ControlName string) *TControl {
	return AsControl(CheckGroup_FindChildControl(c.instance, ControlName))
}

// 返回是否获取焦点。
//
// Return to get focus.
func (c *TCheckGroup) Focused() bool {
	return CheckGroup_Focused(c.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TCheckGroup) HandleAllocated() bool {
	return CheckGroup_HandleAllocated(c.instance)
}

// 插入一个控件。
//
// Insert a control.
func (c *TCheckGroup) InsertControl(AControl IControl) {
	CheckGroup_InsertControl(c.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (c *TCheckGroup) Invalidate() {
	CheckGroup_Invalidate(c.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (c *TCheckGroup) PaintTo(DC HDC, X int32, Y int32) {
	CheckGroup_PaintTo(c.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (c *TCheckGroup) RemoveControl(AControl IControl) {
	CheckGroup_RemoveControl(c.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (c *TCheckGroup) Realign() {
	CheckGroup_Realign(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TCheckGroup) Repaint() {
	CheckGroup_Repaint(c.instance)
}

// 按比例缩放。
//
// Scale by.
func (c *TCheckGroup) ScaleBy(M int32, D int32) {
	CheckGroup_ScaleBy(c.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (c *TCheckGroup) ScrollBy(DeltaX int32, DeltaY int32) {
	CheckGroup_ScrollBy(c.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TCheckGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	CheckGroup_SetBounds(c.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (c *TCheckGroup) SetFocus() {
	CheckGroup_SetFocus(c.instance)
}

// 控件更新。
//
// Update.
func (c *TCheckGroup) Update() {
	CheckGroup_Update(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TCheckGroup) BringToFront() {
	CheckGroup_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TCheckGroup) ClientToScreen(Point TPoint) TPoint {
	return CheckGroup_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TCheckGroup) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return CheckGroup_ClientToParent(c.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TCheckGroup) Dragging() bool {
	return CheckGroup_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TCheckGroup) HasParent() bool {
	return CheckGroup_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TCheckGroup) Hide() {
	CheckGroup_Hide(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TCheckGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return CheckGroup_Perform(c.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TCheckGroup) Refresh() {
	CheckGroup_Refresh(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TCheckGroup) ScreenToClient(Point TPoint) TPoint {
	return CheckGroup_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TCheckGroup) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return CheckGroup_ParentToClient(c.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TCheckGroup) SendToBack() {
	CheckGroup_SendToBack(c.instance)
}

// 显示控件。
//
// Show control.
func (c *TCheckGroup) Show() {
	CheckGroup_Show(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TCheckGroup) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return CheckGroup_GetTextBuf(c.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TCheckGroup) GetTextLen() int32 {
	return CheckGroup_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TCheckGroup) SetTextBuf(Buffer string) {
	CheckGroup_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TCheckGroup) FindComponent(AName string) *TComponent {
	return AsComponent(CheckGroup_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TCheckGroup) GetNamePath() string {
	return CheckGroup_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCheckGroup) Assign(Source IObject) {
	CheckGroup_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TCheckGroup) ClassType() TClass {
	return CheckGroup_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCheckGroup) ClassName() string {
	return CheckGroup_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCheckGroup) InstanceSize() int32 {
	return CheckGroup_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCheckGroup) InheritsFrom(AClass TClass) bool {
	return CheckGroup_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCheckGroup) Equals(Obj IObject) bool {
	return CheckGroup_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCheckGroup) GetHashCode() int32 {
	return CheckGroup_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TCheckGroup) ToString() string {
	return CheckGroup_ToString(c.instance)
}

func (c *TCheckGroup) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	CheckGroup_AnchorToNeighbour(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (c *TCheckGroup) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	CheckGroup_AnchorParallel(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TCheckGroup) AnchorHorizontalCenterTo(ASibling IControl) {
	CheckGroup_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TCheckGroup) AnchorVerticalCenterTo(ASibling IControl) {
	CheckGroup_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TCheckGroup) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	CheckGroup_AnchorSame(c.instance, ASide, CheckPtr(ASibling))
}

func (c *TCheckGroup) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	CheckGroup_AnchorAsAlign(c.instance, ATheAlign, ASpace)
}

func (c *TCheckGroup) AnchorClient(ASpace int32) {
	CheckGroup_AnchorClient(c.instance, ASpace)
}

func (c *TCheckGroup) ScaleDesignToForm(ASize int32) int32 {
	return CheckGroup_ScaleDesignToForm(c.instance, ASize)
}

func (c *TCheckGroup) ScaleFormToDesign(ASize int32) int32 {
	return CheckGroup_ScaleFormToDesign(c.instance, ASize)
}

func (c *TCheckGroup) Scale96ToForm(ASize int32) int32 {
	return CheckGroup_Scale96ToForm(c.instance, ASize)
}

func (c *TCheckGroup) ScaleFormTo96(ASize int32) int32 {
	return CheckGroup_ScaleFormTo96(c.instance, ASize)
}

func (c *TCheckGroup) Scale96ToFont(ASize int32) int32 {
	return CheckGroup_Scale96ToFont(c.instance, ASize)
}

func (c *TCheckGroup) ScaleFontTo96(ASize int32) int32 {
	return CheckGroup_ScaleFontTo96(c.instance, ASize)
}

func (c *TCheckGroup) ScaleScreenToFont(ASize int32) int32 {
	return CheckGroup_ScaleScreenToFont(c.instance, ASize)
}

func (c *TCheckGroup) ScaleFontToScreen(ASize int32) int32 {
	return CheckGroup_ScaleFontToScreen(c.instance, ASize)
}

func (c *TCheckGroup) Scale96ToScreen(ASize int32) int32 {
	return CheckGroup_Scale96ToScreen(c.instance, ASize)
}

func (c *TCheckGroup) ScaleScreenTo96(ASize int32) int32 {
	return CheckGroup_ScaleScreenTo96(c.instance, ASize)
}

func (c *TCheckGroup) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	CheckGroup_AutoAdjustLayout(c.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (c *TCheckGroup) FixDesignFontsPPI(ADesignTimePPI int32) {
	CheckGroup_FixDesignFontsPPI(c.instance, ADesignTimePPI)
}

func (c *TCheckGroup) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	CheckGroup_ScaleFontsPPI(c.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TCheckGroup) Align() TAlign {
	return CheckGroup_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TCheckGroup) SetAlign(value TAlign) {
	CheckGroup_SetAlign(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TCheckGroup) Anchors() TAnchors {
	return CheckGroup_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TCheckGroup) SetAnchors(value TAnchors) {
	CheckGroup_SetAnchors(c.instance, value)
}

func (c *TCheckGroup) AutoFill() bool {
	return CheckGroup_GetAutoFill(c.instance)
}

func (c *TCheckGroup) SetAutoFill(value bool) {
	CheckGroup_SetAutoFill(c.instance, value)
}

// 获取自动调整大小。
func (c *TCheckGroup) AutoSize() bool {
	return CheckGroup_GetAutoSize(c.instance)
}

// 设置自动调整大小。
func (c *TCheckGroup) SetAutoSize(value bool) {
	CheckGroup_SetAutoSize(c.instance, value)
}

func (c *TCheckGroup) BiDiMode() TBiDiMode {
	return CheckGroup_GetBiDiMode(c.instance)
}

func (c *TCheckGroup) SetBiDiMode(value TBiDiMode) {
	CheckGroup_SetBiDiMode(c.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (c *TCheckGroup) Caption() string {
	return CheckGroup_GetCaption(c.instance)
}

// 设置控件标题。
//
// Set the control title.
func (c *TCheckGroup) SetCaption(value string) {
	CheckGroup_SetCaption(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TCheckGroup) ClientHeight() int32 {
	return CheckGroup_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TCheckGroup) SetClientHeight(value int32) {
	CheckGroup_SetClientHeight(c.instance, value)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TCheckGroup) ClientWidth() int32 {
	return CheckGroup_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TCheckGroup) SetClientWidth(value int32) {
	CheckGroup_SetClientWidth(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TCheckGroup) Color() TColor {
	return CheckGroup_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TCheckGroup) SetColor(value TColor) {
	CheckGroup_SetColor(c.instance, value)
}

func (c *TCheckGroup) ColumnLayout() TColumnLayout {
	return CheckGroup_GetColumnLayout(c.instance)
}

func (c *TCheckGroup) SetColumnLayout(value TColumnLayout) {
	CheckGroup_SetColumnLayout(c.instance, value)
}

func (c *TCheckGroup) Columns() int32 {
	return CheckGroup_GetColumns(c.instance)
}

func (c *TCheckGroup) SetColumns(value int32) {
	CheckGroup_SetColumns(c.instance, value)
}

// 获取约束控件大小。
func (c *TCheckGroup) Constraints() *TSizeConstraints {
	return AsSizeConstraints(CheckGroup_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TCheckGroup) SetConstraints(value *TSizeConstraints) {
	CheckGroup_SetConstraints(c.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TCheckGroup) DoubleBuffered() bool {
	return CheckGroup_GetDoubleBuffered(c.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TCheckGroup) SetDoubleBuffered(value bool) {
	CheckGroup_SetDoubleBuffered(c.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (c *TCheckGroup) DragCursor() TCursor {
	return CheckGroup_GetDragCursor(c.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (c *TCheckGroup) SetDragCursor(value TCursor) {
	CheckGroup_SetDragCursor(c.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (c *TCheckGroup) DragMode() TDragMode {
	return CheckGroup_GetDragMode(c.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (c *TCheckGroup) SetDragMode(value TDragMode) {
	CheckGroup_SetDragMode(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TCheckGroup) Enabled() bool {
	return CheckGroup_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TCheckGroup) SetEnabled(value bool) {
	CheckGroup_SetEnabled(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TCheckGroup) Font() *TFont {
	return AsFont(CheckGroup_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TCheckGroup) SetFont(value *TFont) {
	CheckGroup_SetFont(c.instance, CheckPtr(value))
}

func (c *TCheckGroup) Items() *TStrings {
	return AsStrings(CheckGroup_GetItems(c.instance))
}

func (c *TCheckGroup) SetItems(value IStrings) {
	CheckGroup_SetItems(c.instance, CheckPtr(value))
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TCheckGroup) SetOnClick(fn TNotifyEvent) {
	CheckGroup_SetOnClick(c.instance, fn)
}

// 设置双击事件。
func (c *TCheckGroup) SetOnDblClick(fn TNotifyEvent) {
	CheckGroup_SetOnDblClick(c.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TCheckGroup) SetOnDragDrop(fn TDragDropEvent) {
	CheckGroup_SetOnDragDrop(c.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TCheckGroup) SetOnDragOver(fn TDragOverEvent) {
	CheckGroup_SetOnDragOver(c.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (c *TCheckGroup) SetOnEndDrag(fn TEndDragEvent) {
	CheckGroup_SetOnEndDrag(c.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (c *TCheckGroup) SetOnEnter(fn TNotifyEvent) {
	CheckGroup_SetOnEnter(c.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (c *TCheckGroup) SetOnExit(fn TNotifyEvent) {
	CheckGroup_SetOnExit(c.instance, fn)
}

func (c *TCheckGroup) SetOnItemClick(fn TCheckGroupClicked) {
	CheckGroup_SetOnItemClick(c.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TCheckGroup) SetOnKeyDown(fn TKeyEvent) {
	CheckGroup_SetOnKeyDown(c.instance, fn)
}

// 设置键键下事件。
func (c *TCheckGroup) SetOnKeyPress(fn TKeyPressEvent) {
	CheckGroup_SetOnKeyPress(c.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TCheckGroup) SetOnKeyUp(fn TKeyEvent) {
	CheckGroup_SetOnKeyUp(c.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TCheckGroup) SetOnMouseDown(fn TMouseEvent) {
	CheckGroup_SetOnMouseDown(c.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TCheckGroup) SetOnMouseEnter(fn TNotifyEvent) {
	CheckGroup_SetOnMouseEnter(c.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TCheckGroup) SetOnMouseLeave(fn TNotifyEvent) {
	CheckGroup_SetOnMouseLeave(c.instance, fn)
}

// 设置鼠标移动事件。
func (c *TCheckGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	CheckGroup_SetOnMouseMove(c.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TCheckGroup) SetOnMouseUp(fn TMouseEvent) {
	CheckGroup_SetOnMouseUp(c.instance, fn)
}

// 设置鼠标滚轮事件。
func (c *TCheckGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	CheckGroup_SetOnMouseWheel(c.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (c *TCheckGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	CheckGroup_SetOnMouseWheelDown(c.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (c *TCheckGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	CheckGroup_SetOnMouseWheelUp(c.instance, fn)
}

// 设置大小被改变事件。
func (c *TCheckGroup) SetOnResize(fn TNotifyEvent) {
	CheckGroup_SetOnResize(c.instance, fn)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TCheckGroup) ParentFont() bool {
	return CheckGroup_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TCheckGroup) SetParentFont(value bool) {
	CheckGroup_SetParentFont(c.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (c *TCheckGroup) ParentColor() bool {
	return CheckGroup_GetParentColor(c.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (c *TCheckGroup) SetParentColor(value bool) {
	CheckGroup_SetParentColor(c.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TCheckGroup) ParentDoubleBuffered() bool {
	return CheckGroup_GetParentDoubleBuffered(c.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TCheckGroup) SetParentDoubleBuffered(value bool) {
	CheckGroup_SetParentDoubleBuffered(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TCheckGroup) ParentShowHint() bool {
	return CheckGroup_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TCheckGroup) SetParentShowHint(value bool) {
	CheckGroup_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TCheckGroup) PopupMenu() *TPopupMenu {
	return AsPopupMenu(CheckGroup_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TCheckGroup) SetPopupMenu(value IComponent) {
	CheckGroup_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TCheckGroup) ShowHint() bool {
	return CheckGroup_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TCheckGroup) SetShowHint(value bool) {
	CheckGroup_SetShowHint(c.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TCheckGroup) TabOrder() TTabOrder {
	return CheckGroup_GetTabOrder(c.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TCheckGroup) SetTabOrder(value TTabOrder) {
	CheckGroup_SetTabOrder(c.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TCheckGroup) TabStop() bool {
	return CheckGroup_GetTabStop(c.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TCheckGroup) SetTabStop(value bool) {
	CheckGroup_SetTabStop(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TCheckGroup) Visible() bool {
	return CheckGroup_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TCheckGroup) SetVisible(value bool) {
	CheckGroup_SetVisible(c.instance, value)
}

func (c *TCheckGroup) ParentBackground() bool {
	return CheckGroup_GetParentBackground(c.instance)
}

func (c *TCheckGroup) SetParentBackground(value bool) {
	CheckGroup_SetParentBackground(c.instance, value)
}

// 获取依靠客户端总数。
func (c *TCheckGroup) DockClientCount() int32 {
	return CheckGroup_GetDockClientCount(c.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (c *TCheckGroup) DockSite() bool {
	return CheckGroup_GetDockSite(c.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (c *TCheckGroup) SetDockSite(value bool) {
	CheckGroup_SetDockSite(c.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TCheckGroup) MouseInClient() bool {
	return CheckGroup_GetMouseInClient(c.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TCheckGroup) VisibleDockClientCount() int32 {
	return CheckGroup_GetVisibleDockClientCount(c.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (c *TCheckGroup) Brush() *TBrush {
	return AsBrush(CheckGroup_GetBrush(c.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (c *TCheckGroup) ControlCount() int32 {
	return CheckGroup_GetControlCount(c.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TCheckGroup) Handle() HWND {
	return CheckGroup_GetHandle(c.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TCheckGroup) ParentWindow() HWND {
	return CheckGroup_GetParentWindow(c.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TCheckGroup) SetParentWindow(value HWND) {
	CheckGroup_SetParentWindow(c.instance, value)
}

func (c *TCheckGroup) Showing() bool {
	return CheckGroup_GetShowing(c.instance)
}

// 获取使用停靠管理。
func (c *TCheckGroup) UseDockManager() bool {
	return CheckGroup_GetUseDockManager(c.instance)
}

// 设置使用停靠管理。
func (c *TCheckGroup) SetUseDockManager(value bool) {
	CheckGroup_SetUseDockManager(c.instance, value)
}

func (c *TCheckGroup) Action() *TAction {
	return AsAction(CheckGroup_GetAction(c.instance))
}

func (c *TCheckGroup) SetAction(value IComponent) {
	CheckGroup_SetAction(c.instance, CheckPtr(value))
}

func (c *TCheckGroup) BoundsRect() TRect {
	return CheckGroup_GetBoundsRect(c.instance)
}

func (c *TCheckGroup) SetBoundsRect(value TRect) {
	CheckGroup_SetBoundsRect(c.instance, value)
}

func (c *TCheckGroup) ClientOrigin() TPoint {
	return CheckGroup_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TCheckGroup) ClientRect() TRect {
	return CheckGroup_GetClientRect(c.instance)
}

// 获取控件状态。
//
// Get control state.
func (c *TCheckGroup) ControlState() TControlState {
	return CheckGroup_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TCheckGroup) SetControlState(value TControlState) {
	CheckGroup_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TCheckGroup) ControlStyle() TControlStyle {
	return CheckGroup_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TCheckGroup) SetControlStyle(value TControlStyle) {
	CheckGroup_SetControlStyle(c.instance, value)
}

func (c *TCheckGroup) Floating() bool {
	return CheckGroup_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TCheckGroup) Parent() *TWinControl {
	return AsWinControl(CheckGroup_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TCheckGroup) SetParent(value IWinControl) {
	CheckGroup_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TCheckGroup) Left() int32 {
	return CheckGroup_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TCheckGroup) SetLeft(value int32) {
	CheckGroup_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TCheckGroup) Top() int32 {
	return CheckGroup_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TCheckGroup) SetTop(value int32) {
	CheckGroup_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TCheckGroup) Width() int32 {
	return CheckGroup_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TCheckGroup) SetWidth(value int32) {
	CheckGroup_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TCheckGroup) Height() int32 {
	return CheckGroup_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TCheckGroup) SetHeight(value int32) {
	CheckGroup_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TCheckGroup) Cursor() TCursor {
	return CheckGroup_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TCheckGroup) SetCursor(value TCursor) {
	CheckGroup_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TCheckGroup) Hint() string {
	return CheckGroup_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TCheckGroup) SetHint(value string) {
	CheckGroup_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TCheckGroup) ComponentCount() int32 {
	return CheckGroup_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TCheckGroup) ComponentIndex() int32 {
	return CheckGroup_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TCheckGroup) SetComponentIndex(value int32) {
	CheckGroup_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TCheckGroup) Owner() *TComponent {
	return AsComponent(CheckGroup_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TCheckGroup) Name() string {
	return CheckGroup_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TCheckGroup) SetName(value string) {
	CheckGroup_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TCheckGroup) Tag() int {
	return CheckGroup_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TCheckGroup) SetTag(value int) {
	CheckGroup_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TCheckGroup) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(CheckGroup_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TCheckGroup) SetAnchorSideLeft(value *TAnchorSide) {
	CheckGroup_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TCheckGroup) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(CheckGroup_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TCheckGroup) SetAnchorSideTop(value *TAnchorSide) {
	CheckGroup_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TCheckGroup) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(CheckGroup_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TCheckGroup) SetAnchorSideRight(value *TAnchorSide) {
	CheckGroup_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TCheckGroup) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(CheckGroup_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TCheckGroup) SetAnchorSideBottom(value *TAnchorSide) {
	CheckGroup_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TCheckGroup) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(CheckGroup_GetChildSizing(c.instance))
}

func (c *TCheckGroup) SetChildSizing(value *TControlChildSizing) {
	CheckGroup_SetChildSizing(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TCheckGroup) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(CheckGroup_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TCheckGroup) SetBorderSpacing(value *TControlBorderSpacing) {
	CheckGroup_SetBorderSpacing(c.instance, CheckPtr(value))
}

// 获取是否选中。
func (c *TCheckGroup) Checked(Index int32) bool {
	return CheckGroup_GetChecked(c.instance, Index)
}

// 设置是否选中。
func (c *TCheckGroup) SetChecked(Index int32, value bool) {
	CheckGroup_SetChecked(c.instance, Index, value)
}

func (c *TCheckGroup) CheckEnabled(Index int32) bool {
	return CheckGroup_GetCheckEnabled(c.instance, Index)
}

func (c *TCheckGroup) SetCheckEnabled(Index int32, value bool) {
	CheckGroup_SetCheckEnabled(c.instance, Index, value)
}

// 获取指定索引停靠客户端。
func (c *TCheckGroup) DockClients(Index int32) *TControl {
	return AsControl(CheckGroup_GetDockClients(c.instance, Index))
}

// 获取指定索引子控件。
func (c *TCheckGroup) Controls(Index int32) *TControl {
	return AsControl(CheckGroup_GetControls(c.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TCheckGroup) Components(AIndex int32) *TComponent {
	return AsComponent(CheckGroup_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TCheckGroup) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(CheckGroup_GetAnchorSide(c.instance, AKind))
}
