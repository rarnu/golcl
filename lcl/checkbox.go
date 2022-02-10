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

type TCheckBox struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewCheckBox(owner IComponent) *TCheckBox {
	c := new(TCheckBox)
	c.instance = CheckBox_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsCheckBox(obj interface{}) *TCheckBox {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TCheckBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsCheckBox.
func CheckBoxFromInst(inst uintptr) *TCheckBox {
	return AsCheckBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsCheckBox.
func CheckBoxFromObj(obj IObject) *TCheckBox {
	return AsCheckBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCheckBox.
func CheckBoxFromUnsafePointer(ptr unsafe.Pointer) *TCheckBox {
	return AsCheckBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (c *TCheckBox) Free() {
	if c.instance != 0 {
		CheckBox_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TCheckBox) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TCheckBox) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TCheckBox) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TCheckBox) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TCheckBox) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TCheckBoxClass() TClass {
	return CheckBox_StaticClassType()
}

// 是否可以获得焦点。
func (c *TCheckBox) CanFocus() bool {
	return CheckBox_CanFocus(c.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TCheckBox) ContainsControl(Control IControl) bool {
	return CheckBox_ContainsControl(c.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TCheckBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(CheckBox_ControlAtPos(c.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TCheckBox) DisableAlign() {
	CheckBox_DisableAlign(c.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (c *TCheckBox) EnableAlign() {
	CheckBox_EnableAlign(c.instance)
}

// 查找子控件。
//
// Find sub controls.
func (c *TCheckBox) FindChildControl(ControlName string) *TControl {
	return AsControl(CheckBox_FindChildControl(c.instance, ControlName))
}

func (c *TCheckBox) FlipChildren(AllLevels bool) {
	CheckBox_FlipChildren(c.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (c *TCheckBox) Focused() bool {
	return CheckBox_Focused(c.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TCheckBox) HandleAllocated() bool {
	return CheckBox_HandleAllocated(c.instance)
}

// 插入一个控件。
//
// Insert a control.
func (c *TCheckBox) InsertControl(AControl IControl) {
	CheckBox_InsertControl(c.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (c *TCheckBox) Invalidate() {
	CheckBox_Invalidate(c.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (c *TCheckBox) PaintTo(DC HDC, X int32, Y int32) {
	CheckBox_PaintTo(c.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (c *TCheckBox) RemoveControl(AControl IControl) {
	CheckBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (c *TCheckBox) Realign() {
	CheckBox_Realign(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TCheckBox) Repaint() {
	CheckBox_Repaint(c.instance)
}

// 按比例缩放。
//
// Scale by.
func (c *TCheckBox) ScaleBy(M int32, D int32) {
	CheckBox_ScaleBy(c.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (c *TCheckBox) ScrollBy(DeltaX int32, DeltaY int32) {
	CheckBox_ScrollBy(c.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TCheckBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	CheckBox_SetBounds(c.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (c *TCheckBox) SetFocus() {
	CheckBox_SetFocus(c.instance)
}

// 控件更新。
//
// Update.
func (c *TCheckBox) Update() {
	CheckBox_Update(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TCheckBox) BringToFront() {
	CheckBox_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TCheckBox) ClientToScreen(Point TPoint) TPoint {
	return CheckBox_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TCheckBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return CheckBox_ClientToParent(c.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TCheckBox) Dragging() bool {
	return CheckBox_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TCheckBox) HasParent() bool {
	return CheckBox_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TCheckBox) Hide() {
	CheckBox_Hide(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TCheckBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return CheckBox_Perform(c.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TCheckBox) Refresh() {
	CheckBox_Refresh(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TCheckBox) ScreenToClient(Point TPoint) TPoint {
	return CheckBox_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TCheckBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return CheckBox_ParentToClient(c.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TCheckBox) SendToBack() {
	CheckBox_SendToBack(c.instance)
}

// 显示控件。
//
// Show control.
func (c *TCheckBox) Show() {
	CheckBox_Show(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TCheckBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return CheckBox_GetTextBuf(c.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TCheckBox) GetTextLen() int32 {
	return CheckBox_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TCheckBox) SetTextBuf(Buffer string) {
	CheckBox_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TCheckBox) FindComponent(AName string) *TComponent {
	return AsComponent(CheckBox_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TCheckBox) GetNamePath() string {
	return CheckBox_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCheckBox) Assign(Source IObject) {
	CheckBox_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TCheckBox) ClassType() TClass {
	return CheckBox_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCheckBox) ClassName() string {
	return CheckBox_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCheckBox) InstanceSize() int32 {
	return CheckBox_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCheckBox) InheritsFrom(AClass TClass) bool {
	return CheckBox_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCheckBox) Equals(Obj IObject) bool {
	return CheckBox_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCheckBox) GetHashCode() int32 {
	return CheckBox_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TCheckBox) ToString() string {
	return CheckBox_ToString(c.instance)
}

func (c *TCheckBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	CheckBox_AnchorToNeighbour(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (c *TCheckBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	CheckBox_AnchorParallel(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TCheckBox) AnchorHorizontalCenterTo(ASibling IControl) {
	CheckBox_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TCheckBox) AnchorVerticalCenterTo(ASibling IControl) {
	CheckBox_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TCheckBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	CheckBox_AnchorSame(c.instance, ASide, CheckPtr(ASibling))
}

func (c *TCheckBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	CheckBox_AnchorAsAlign(c.instance, ATheAlign, ASpace)
}

func (c *TCheckBox) AnchorClient(ASpace int32) {
	CheckBox_AnchorClient(c.instance, ASpace)
}

func (c *TCheckBox) ScaleDesignToForm(ASize int32) int32 {
	return CheckBox_ScaleDesignToForm(c.instance, ASize)
}

func (c *TCheckBox) ScaleFormToDesign(ASize int32) int32 {
	return CheckBox_ScaleFormToDesign(c.instance, ASize)
}

func (c *TCheckBox) Scale96ToForm(ASize int32) int32 {
	return CheckBox_Scale96ToForm(c.instance, ASize)
}

func (c *TCheckBox) ScaleFormTo96(ASize int32) int32 {
	return CheckBox_ScaleFormTo96(c.instance, ASize)
}

func (c *TCheckBox) Scale96ToFont(ASize int32) int32 {
	return CheckBox_Scale96ToFont(c.instance, ASize)
}

func (c *TCheckBox) ScaleFontTo96(ASize int32) int32 {
	return CheckBox_ScaleFontTo96(c.instance, ASize)
}

func (c *TCheckBox) ScaleScreenToFont(ASize int32) int32 {
	return CheckBox_ScaleScreenToFont(c.instance, ASize)
}

func (c *TCheckBox) ScaleFontToScreen(ASize int32) int32 {
	return CheckBox_ScaleFontToScreen(c.instance, ASize)
}

func (c *TCheckBox) Scale96ToScreen(ASize int32) int32 {
	return CheckBox_Scale96ToScreen(c.instance, ASize)
}

func (c *TCheckBox) ScaleScreenTo96(ASize int32) int32 {
	return CheckBox_ScaleScreenTo96(c.instance, ASize)
}

func (c *TCheckBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	CheckBox_AutoAdjustLayout(c.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (c *TCheckBox) FixDesignFontsPPI(ADesignTimePPI int32) {
	CheckBox_FixDesignFontsPPI(c.instance, ADesignTimePPI)
}

func (c *TCheckBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	CheckBox_ScaleFontsPPI(c.instance, AToPPI, AProportion)
}

// 设置改变事件。
//
// Set changed event.
func (c *TCheckBox) SetOnChange(fn TNotifyEvent) {
	CheckBox_SetOnChange(c.instance, fn)
}

func (c *TCheckBox) Action() *TAction {
	return AsAction(CheckBox_GetAction(c.instance))
}

func (c *TCheckBox) SetAction(value IComponent) {
	CheckBox_SetAction(c.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TCheckBox) Align() TAlign {
	return CheckBox_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TCheckBox) SetAlign(value TAlign) {
	CheckBox_SetAlign(c.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (c *TCheckBox) Alignment() TLeftRight {
	return CheckBox_GetAlignment(c.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (c *TCheckBox) SetAlignment(value TLeftRight) {
	CheckBox_SetAlignment(c.instance, value)
}

func (c *TCheckBox) AllowGrayed() bool {
	return CheckBox_GetAllowGrayed(c.instance)
}

func (c *TCheckBox) SetAllowGrayed(value bool) {
	CheckBox_SetAllowGrayed(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TCheckBox) Anchors() TAnchors {
	return CheckBox_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TCheckBox) SetAnchors(value TAnchors) {
	CheckBox_SetAnchors(c.instance, value)
}

func (c *TCheckBox) BiDiMode() TBiDiMode {
	return CheckBox_GetBiDiMode(c.instance)
}

func (c *TCheckBox) SetBiDiMode(value TBiDiMode) {
	CheckBox_SetBiDiMode(c.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (c *TCheckBox) Caption() string {
	return CheckBox_GetCaption(c.instance)
}

// 设置控件标题。
//
// Set the control title.
func (c *TCheckBox) SetCaption(value string) {
	CheckBox_SetCaption(c.instance, value)
}

// 获取是否选中。
func (c *TCheckBox) Checked() bool {
	return CheckBox_GetChecked(c.instance)
}

// 设置是否选中。
func (c *TCheckBox) SetChecked(value bool) {
	CheckBox_SetChecked(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TCheckBox) Color() TColor {
	return CheckBox_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TCheckBox) SetColor(value TColor) {
	CheckBox_SetColor(c.instance, value)
}

// 获取约束控件大小。
func (c *TCheckBox) Constraints() *TSizeConstraints {
	return AsSizeConstraints(CheckBox_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TCheckBox) SetConstraints(value *TSizeConstraints) {
	CheckBox_SetConstraints(c.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TCheckBox) DoubleBuffered() bool {
	return CheckBox_GetDoubleBuffered(c.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TCheckBox) SetDoubleBuffered(value bool) {
	CheckBox_SetDoubleBuffered(c.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (c *TCheckBox) DragCursor() TCursor {
	return CheckBox_GetDragCursor(c.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (c *TCheckBox) SetDragCursor(value TCursor) {
	CheckBox_SetDragCursor(c.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (c *TCheckBox) DragKind() TDragKind {
	return CheckBox_GetDragKind(c.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (c *TCheckBox) SetDragKind(value TDragKind) {
	CheckBox_SetDragKind(c.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (c *TCheckBox) DragMode() TDragMode {
	return CheckBox_GetDragMode(c.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (c *TCheckBox) SetDragMode(value TDragMode) {
	CheckBox_SetDragMode(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TCheckBox) Enabled() bool {
	return CheckBox_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TCheckBox) SetEnabled(value bool) {
	CheckBox_SetEnabled(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TCheckBox) Font() *TFont {
	return AsFont(CheckBox_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TCheckBox) SetFont(value *TFont) {
	CheckBox_SetFont(c.instance, CheckPtr(value))
}

// 获取使用父容器颜色。
//
// Get parent color.
func (c *TCheckBox) ParentColor() bool {
	return CheckBox_GetParentColor(c.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (c *TCheckBox) SetParentColor(value bool) {
	CheckBox_SetParentColor(c.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TCheckBox) ParentDoubleBuffered() bool {
	return CheckBox_GetParentDoubleBuffered(c.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TCheckBox) SetParentDoubleBuffered(value bool) {
	CheckBox_SetParentDoubleBuffered(c.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TCheckBox) ParentFont() bool {
	return CheckBox_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TCheckBox) SetParentFont(value bool) {
	CheckBox_SetParentFont(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TCheckBox) ParentShowHint() bool {
	return CheckBox_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TCheckBox) SetParentShowHint(value bool) {
	CheckBox_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TCheckBox) PopupMenu() *TPopupMenu {
	return AsPopupMenu(CheckBox_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TCheckBox) SetPopupMenu(value IComponent) {
	CheckBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TCheckBox) ShowHint() bool {
	return CheckBox_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TCheckBox) SetShowHint(value bool) {
	CheckBox_SetShowHint(c.instance, value)
}

func (c *TCheckBox) State() TCheckBoxState {
	return CheckBox_GetState(c.instance)
}

func (c *TCheckBox) SetState(value TCheckBoxState) {
	CheckBox_SetState(c.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TCheckBox) TabOrder() TTabOrder {
	return CheckBox_GetTabOrder(c.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TCheckBox) SetTabOrder(value TTabOrder) {
	CheckBox_SetTabOrder(c.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TCheckBox) TabStop() bool {
	return CheckBox_GetTabStop(c.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TCheckBox) SetTabStop(value bool) {
	CheckBox_SetTabStop(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TCheckBox) Visible() bool {
	return CheckBox_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TCheckBox) SetVisible(value bool) {
	CheckBox_SetVisible(c.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TCheckBox) SetOnClick(fn TNotifyEvent) {
	CheckBox_SetOnClick(c.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (c *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
	CheckBox_SetOnContextPopup(c.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
	CheckBox_SetOnDragDrop(c.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
	CheckBox_SetOnDragOver(c.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (c *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
	CheckBox_SetOnEndDrag(c.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (c *TCheckBox) SetOnEnter(fn TNotifyEvent) {
	CheckBox_SetOnEnter(c.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (c *TCheckBox) SetOnExit(fn TNotifyEvent) {
	CheckBox_SetOnExit(c.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TCheckBox) SetOnKeyDown(fn TKeyEvent) {
	CheckBox_SetOnKeyDown(c.instance, fn)
}

// 设置键键下事件。
func (c *TCheckBox) SetOnKeyPress(fn TKeyPressEvent) {
	CheckBox_SetOnKeyPress(c.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TCheckBox) SetOnKeyUp(fn TKeyEvent) {
	CheckBox_SetOnKeyUp(c.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
	CheckBox_SetOnMouseDown(c.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
	CheckBox_SetOnMouseEnter(c.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
	CheckBox_SetOnMouseLeave(c.instance, fn)
}

// 设置鼠标移动事件。
func (c *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
	CheckBox_SetOnMouseMove(c.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
	CheckBox_SetOnMouseUp(c.instance, fn)
}

// 获取依靠客户端总数。
func (c *TCheckBox) DockClientCount() int32 {
	return CheckBox_GetDockClientCount(c.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (c *TCheckBox) DockSite() bool {
	return CheckBox_GetDockSite(c.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (c *TCheckBox) SetDockSite(value bool) {
	CheckBox_SetDockSite(c.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TCheckBox) MouseInClient() bool {
	return CheckBox_GetMouseInClient(c.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TCheckBox) VisibleDockClientCount() int32 {
	return CheckBox_GetVisibleDockClientCount(c.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (c *TCheckBox) Brush() *TBrush {
	return AsBrush(CheckBox_GetBrush(c.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (c *TCheckBox) ControlCount() int32 {
	return CheckBox_GetControlCount(c.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TCheckBox) Handle() HWND {
	return CheckBox_GetHandle(c.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TCheckBox) ParentWindow() HWND {
	return CheckBox_GetParentWindow(c.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TCheckBox) SetParentWindow(value HWND) {
	CheckBox_SetParentWindow(c.instance, value)
}

func (c *TCheckBox) Showing() bool {
	return CheckBox_GetShowing(c.instance)
}

// 获取使用停靠管理。
func (c *TCheckBox) UseDockManager() bool {
	return CheckBox_GetUseDockManager(c.instance)
}

// 设置使用停靠管理。
func (c *TCheckBox) SetUseDockManager(value bool) {
	CheckBox_SetUseDockManager(c.instance, value)
}

func (c *TCheckBox) BoundsRect() TRect {
	return CheckBox_GetBoundsRect(c.instance)
}

func (c *TCheckBox) SetBoundsRect(value TRect) {
	CheckBox_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TCheckBox) ClientHeight() int32 {
	return CheckBox_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TCheckBox) SetClientHeight(value int32) {
	CheckBox_SetClientHeight(c.instance, value)
}

func (c *TCheckBox) ClientOrigin() TPoint {
	return CheckBox_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TCheckBox) ClientRect() TRect {
	return CheckBox_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TCheckBox) ClientWidth() int32 {
	return CheckBox_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TCheckBox) SetClientWidth(value int32) {
	CheckBox_SetClientWidth(c.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (c *TCheckBox) ControlState() TControlState {
	return CheckBox_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TCheckBox) SetControlState(value TControlState) {
	CheckBox_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TCheckBox) ControlStyle() TControlStyle {
	return CheckBox_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TCheckBox) SetControlStyle(value TControlStyle) {
	CheckBox_SetControlStyle(c.instance, value)
}

func (c *TCheckBox) Floating() bool {
	return CheckBox_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TCheckBox) Parent() *TWinControl {
	return AsWinControl(CheckBox_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TCheckBox) SetParent(value IWinControl) {
	CheckBox_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TCheckBox) Left() int32 {
	return CheckBox_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TCheckBox) SetLeft(value int32) {
	CheckBox_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TCheckBox) Top() int32 {
	return CheckBox_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TCheckBox) SetTop(value int32) {
	CheckBox_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TCheckBox) Width() int32 {
	return CheckBox_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TCheckBox) SetWidth(value int32) {
	CheckBox_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TCheckBox) Height() int32 {
	return CheckBox_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TCheckBox) SetHeight(value int32) {
	CheckBox_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TCheckBox) Cursor() TCursor {
	return CheckBox_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TCheckBox) SetCursor(value TCursor) {
	CheckBox_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TCheckBox) Hint() string {
	return CheckBox_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TCheckBox) SetHint(value string) {
	CheckBox_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TCheckBox) ComponentCount() int32 {
	return CheckBox_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TCheckBox) ComponentIndex() int32 {
	return CheckBox_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TCheckBox) SetComponentIndex(value int32) {
	CheckBox_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TCheckBox) Owner() *TComponent {
	return AsComponent(CheckBox_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TCheckBox) Name() string {
	return CheckBox_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TCheckBox) SetName(value string) {
	CheckBox_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TCheckBox) Tag() int {
	return CheckBox_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TCheckBox) SetTag(value int) {
	CheckBox_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TCheckBox) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(CheckBox_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TCheckBox) SetAnchorSideLeft(value *TAnchorSide) {
	CheckBox_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TCheckBox) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(CheckBox_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TCheckBox) SetAnchorSideTop(value *TAnchorSide) {
	CheckBox_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TCheckBox) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(CheckBox_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TCheckBox) SetAnchorSideRight(value *TAnchorSide) {
	CheckBox_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TCheckBox) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(CheckBox_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TCheckBox) SetAnchorSideBottom(value *TAnchorSide) {
	CheckBox_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(CheckBox_GetChildSizing(c.instance))
}

func (c *TCheckBox) SetChildSizing(value *TControlChildSizing) {
	CheckBox_SetChildSizing(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TCheckBox) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(CheckBox_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TCheckBox) SetBorderSpacing(value *TControlBorderSpacing) {
	CheckBox_SetBorderSpacing(c.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (c *TCheckBox) DockClients(Index int32) *TControl {
	return AsControl(CheckBox_GetDockClients(c.instance, Index))
}

// 获取指定索引子控件。
func (c *TCheckBox) Controls(Index int32) *TControl {
	return AsControl(CheckBox_GetControls(c.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TCheckBox) Components(AIndex int32) *TComponent {
	return AsComponent(CheckBox_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TCheckBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(CheckBox_GetAnchorSide(c.instance, AKind))
}

