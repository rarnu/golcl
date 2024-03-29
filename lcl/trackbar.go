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

type TTrackBar struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewTrackBar(owner IComponent) *TTrackBar {
	t := new(TTrackBar)
	t.instance = TrackBar_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTrackBar(obj any) *TTrackBar {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTrackBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTrackBar.
func TrackBarFromInst(inst uintptr) *TTrackBar {
	return AsTrackBar(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsTrackBar.
func TrackBarFromObj(obj IObject) *TTrackBar {
	return AsTrackBar(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTrackBar.
func TrackBarFromUnsafePointer(ptr unsafe.Pointer) *TTrackBar {
	return AsTrackBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TTrackBar) Free() {
	if t.instance != 0 {
		TrackBar_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTrackBar) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTrackBar) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTrackBar) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTrackBar) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTrackBar) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TTrackBarClass() TClass {
	return TrackBar_StaticClassType()
}

func (t *TTrackBar) SetTick(Value int32) {
	TrackBar_SetTick(t.instance, Value)
}

// 是否可以获得焦点。
func (t *TTrackBar) CanFocus() bool {
	return TrackBar_CanFocus(t.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (t *TTrackBar) ContainsControl(Control IControl) bool {
	return TrackBar_ContainsControl(t.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (t *TTrackBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(TrackBar_ControlAtPos(t.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (t *TTrackBar) DisableAlign() {
	TrackBar_DisableAlign(t.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (t *TTrackBar) EnableAlign() {
	TrackBar_EnableAlign(t.instance)
}

// 查找子控件。
//
// Find sub controls.
func (t *TTrackBar) FindChildControl(ControlName string) *TControl {
	return AsControl(TrackBar_FindChildControl(t.instance, ControlName))
}

func (t *TTrackBar) FlipChildren(AllLevels bool) {
	TrackBar_FlipChildren(t.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (t *TTrackBar) Focused() bool {
	return TrackBar_Focused(t.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (t *TTrackBar) HandleAllocated() bool {
	return TrackBar_HandleAllocated(t.instance)
}

// 插入一个控件。
//
// Insert a control.
func (t *TTrackBar) InsertControl(AControl IControl) {
	TrackBar_InsertControl(t.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (t *TTrackBar) Invalidate() {
	TrackBar_Invalidate(t.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (t *TTrackBar) PaintTo(DC HDC, X int32, Y int32) {
	TrackBar_PaintTo(t.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (t *TTrackBar) RemoveControl(AControl IControl) {
	TrackBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (t *TTrackBar) Realign() {
	TrackBar_Realign(t.instance)
}

// 重绘。
//
// Repaint.
func (t *TTrackBar) Repaint() {
	TrackBar_Repaint(t.instance)
}

// 按比例缩放。
//
// Scale by.
func (t *TTrackBar) ScaleBy(M int32, D int32) {
	TrackBar_ScaleBy(t.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (t *TTrackBar) ScrollBy(DeltaX int32, DeltaY int32) {
	TrackBar_ScrollBy(t.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (t *TTrackBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	TrackBar_SetBounds(t.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (t *TTrackBar) SetFocus() {
	TrackBar_SetFocus(t.instance)
}

// 控件更新。
//
// Update.
func (t *TTrackBar) Update() {
	TrackBar_Update(t.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (t *TTrackBar) BringToFront() {
	TrackBar_BringToFront(t.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TTrackBar) ClientToScreen(Point TPoint) TPoint {
	return TrackBar_ClientToScreen(t.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TTrackBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return TrackBar_ClientToParent(t.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TTrackBar) Dragging() bool {
	return TrackBar_Dragging(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TTrackBar) HasParent() bool {
	return TrackBar_HasParent(t.instance)
}

// 隐藏控件。
//
// Hidden control.
func (t *TTrackBar) Hide() {
	TrackBar_Hide(t.instance)
}

// 发送一个消息。
//
// Send a message.
func (t *TTrackBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return TrackBar_Perform(t.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (t *TTrackBar) Refresh() {
	TrackBar_Refresh(t.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TTrackBar) ScreenToClient(Point TPoint) TPoint {
	return TrackBar_ScreenToClient(t.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TTrackBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return TrackBar_ParentToClient(t.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (t *TTrackBar) SendToBack() {
	TrackBar_SendToBack(t.instance)
}

// 显示控件。
//
// Show control.
func (t *TTrackBar) Show() {
	TrackBar_Show(t.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TTrackBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return TrackBar_GetTextBuf(t.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TTrackBar) GetTextLen() int32 {
	return TrackBar_GetTextLen(t.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TTrackBar) SetTextBuf(Buffer string) {
	TrackBar_SetTextBuf(t.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TTrackBar) FindComponent(AName string) *TComponent {
	return AsComponent(TrackBar_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTrackBar) GetNamePath() string {
	return TrackBar_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTrackBar) Assign(Source IObject) {
	TrackBar_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTrackBar) ClassType() TClass {
	return TrackBar_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTrackBar) ClassName() string {
	return TrackBar_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTrackBar) InstanceSize() int32 {
	return TrackBar_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTrackBar) InheritsFrom(AClass TClass) bool {
	return TrackBar_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTrackBar) Equals(Obj IObject) bool {
	return TrackBar_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTrackBar) GetHashCode() int32 {
	return TrackBar_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTrackBar) ToString() string {
	return TrackBar_ToString(t.instance)
}

func (t *TTrackBar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	TrackBar_AnchorToNeighbour(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (t *TTrackBar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	TrackBar_AnchorParallel(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (t *TTrackBar) AnchorHorizontalCenterTo(ASibling IControl) {
	TrackBar_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (t *TTrackBar) AnchorVerticalCenterTo(ASibling IControl) {
	TrackBar_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TTrackBar) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	TrackBar_AnchorSame(t.instance, ASide, CheckPtr(ASibling))
}

func (t *TTrackBar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	TrackBar_AnchorAsAlign(t.instance, ATheAlign, ASpace)
}

func (t *TTrackBar) AnchorClient(ASpace int32) {
	TrackBar_AnchorClient(t.instance, ASpace)
}

func (t *TTrackBar) ScaleDesignToForm(ASize int32) int32 {
	return TrackBar_ScaleDesignToForm(t.instance, ASize)
}

func (t *TTrackBar) ScaleFormToDesign(ASize int32) int32 {
	return TrackBar_ScaleFormToDesign(t.instance, ASize)
}

func (t *TTrackBar) Scale96ToForm(ASize int32) int32 {
	return TrackBar_Scale96ToForm(t.instance, ASize)
}

func (t *TTrackBar) ScaleFormTo96(ASize int32) int32 {
	return TrackBar_ScaleFormTo96(t.instance, ASize)
}

func (t *TTrackBar) Scale96ToFont(ASize int32) int32 {
	return TrackBar_Scale96ToFont(t.instance, ASize)
}

func (t *TTrackBar) ScaleFontTo96(ASize int32) int32 {
	return TrackBar_ScaleFontTo96(t.instance, ASize)
}

func (t *TTrackBar) ScaleScreenToFont(ASize int32) int32 {
	return TrackBar_ScaleScreenToFont(t.instance, ASize)
}

func (t *TTrackBar) ScaleFontToScreen(ASize int32) int32 {
	return TrackBar_ScaleFontToScreen(t.instance, ASize)
}

func (t *TTrackBar) Scale96ToScreen(ASize int32) int32 {
	return TrackBar_Scale96ToScreen(t.instance, ASize)
}

func (t *TTrackBar) ScaleScreenTo96(ASize int32) int32 {
	return TrackBar_ScaleScreenTo96(t.instance, ASize)
}

func (t *TTrackBar) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	TrackBar_AutoAdjustLayout(t.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (t *TTrackBar) FixDesignFontsPPI(ADesignTimePPI int32) {
	TrackBar_FixDesignFontsPPI(t.instance, ADesignTimePPI)
}

func (t *TTrackBar) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	TrackBar_ScaleFontsPPI(t.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TTrackBar) Align() TAlign {
	return TrackBar_GetAlign(t.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TTrackBar) SetAlign(value TAlign) {
	TrackBar_SetAlign(t.instance, value)
}

// 获取四个角位置的锚点。
func (t *TTrackBar) Anchors() TAnchors {
	return TrackBar_GetAnchors(t.instance)
}

// 设置四个角位置的锚点。
func (t *TTrackBar) SetAnchors(value TAnchors) {
	TrackBar_SetAnchors(t.instance, value)
}

// 获取边框的宽度。
func (t *TTrackBar) BorderWidth() int32 {
	return TrackBar_GetBorderWidth(t.instance)
}

// 设置边框的宽度。
func (t *TTrackBar) SetBorderWidth(value int32) {
	TrackBar_SetBorderWidth(t.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (t *TTrackBar) DoubleBuffered() bool {
	return TrackBar_GetDoubleBuffered(t.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (t *TTrackBar) SetDoubleBuffered(value bool) {
	TrackBar_SetDoubleBuffered(t.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (t *TTrackBar) DragCursor() TCursor {
	return TrackBar_GetDragCursor(t.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (t *TTrackBar) SetDragCursor(value TCursor) {
	TrackBar_SetDragCursor(t.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (t *TTrackBar) DragMode() TDragMode {
	return TrackBar_GetDragMode(t.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (t *TTrackBar) SetDragMode(value TDragMode) {
	TrackBar_SetDragMode(t.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TTrackBar) Enabled() bool {
	return TrackBar_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TTrackBar) SetEnabled(value bool) {
	TrackBar_SetEnabled(t.instance, value)
}

// 获取约束控件大小。
func (t *TTrackBar) Constraints() *TSizeConstraints {
	return AsSizeConstraints(TrackBar_GetConstraints(t.instance))
}

// 设置约束控件大小。
func (t *TTrackBar) SetConstraints(value *TSizeConstraints) {
	TrackBar_SetConstraints(t.instance, CheckPtr(value))
}

func (t *TTrackBar) LineSize() int32 {
	return TrackBar_GetLineSize(t.instance)
}

func (t *TTrackBar) SetLineSize(value int32) {
	TrackBar_SetLineSize(t.instance, value)
}

func (t *TTrackBar) Max() int32 {
	return TrackBar_GetMax(t.instance)
}

func (t *TTrackBar) SetMax(value int32) {
	TrackBar_SetMax(t.instance, value)
}

func (t *TTrackBar) Min() int32 {
	return TrackBar_GetMin(t.instance)
}

func (t *TTrackBar) SetMin(value int32) {
	TrackBar_SetMin(t.instance, value)
}

func (t *TTrackBar) Orientation() TTrackBarOrientation {
	return TrackBar_GetOrientation(t.instance)
}

func (t *TTrackBar) SetOrientation(value TTrackBarOrientation) {
	TrackBar_SetOrientation(t.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (t *TTrackBar) ParentDoubleBuffered() bool {
	return TrackBar_GetParentDoubleBuffered(t.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (t *TTrackBar) SetParentDoubleBuffered(value bool) {
	TrackBar_SetParentDoubleBuffered(t.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (t *TTrackBar) ParentShowHint() bool {
	return TrackBar_GetParentShowHint(t.instance)
}

// 设置以父容器的ShowHint属性为准。
func (t *TTrackBar) SetParentShowHint(value bool) {
	TrackBar_SetParentShowHint(t.instance, value)
}

func (t *TTrackBar) PageSize() int32 {
	return TrackBar_GetPageSize(t.instance)
}

func (t *TTrackBar) SetPageSize(value int32) {
	TrackBar_SetPageSize(t.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TTrackBar) PopupMenu() *TPopupMenu {
	return AsPopupMenu(TrackBar_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TTrackBar) SetPopupMenu(value IComponent) {
	TrackBar_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTrackBar) Frequency() int32 {
	return TrackBar_GetFrequency(t.instance)
}

func (t *TTrackBar) SetFrequency(value int32) {
	TrackBar_SetFrequency(t.instance, value)
}

func (t *TTrackBar) Position() int32 {
	return TrackBar_GetPosition(t.instance)
}

func (t *TTrackBar) SetPosition(value int32) {
	TrackBar_SetPosition(t.instance, value)
}

func (t *TTrackBar) SelEnd() int32 {
	return TrackBar_GetSelEnd(t.instance)
}

func (t *TTrackBar) SetSelEnd(value int32) {
	TrackBar_SetSelEnd(t.instance, value)
}

// 获取选择的启始位置。
func (t *TTrackBar) SelStart() int32 {
	return TrackBar_GetSelStart(t.instance)
}

// 设置选择的启始位置。
func (t *TTrackBar) SetSelStart(value int32) {
	TrackBar_SetSelStart(t.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TTrackBar) ShowHint() bool {
	return TrackBar_GetShowHint(t.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TTrackBar) SetShowHint(value bool) {
	TrackBar_SetShowHint(t.instance, value)
}

func (t *TTrackBar) ShowSelRange() bool {
	return TrackBar_GetShowSelRange(t.instance)
}

func (t *TTrackBar) SetShowSelRange(value bool) {
	TrackBar_SetShowSelRange(t.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (t *TTrackBar) TabOrder() TTabOrder {
	return TrackBar_GetTabOrder(t.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (t *TTrackBar) SetTabOrder(value TTabOrder) {
	TrackBar_SetTabOrder(t.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (t *TTrackBar) TabStop() bool {
	return TrackBar_GetTabStop(t.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (t *TTrackBar) SetTabStop(value bool) {
	TrackBar_SetTabStop(t.instance, value)
}

func (t *TTrackBar) TickMarks() TTickMark {
	return TrackBar_GetTickMarks(t.instance)
}

func (t *TTrackBar) SetTickMarks(value TTickMark) {
	TrackBar_SetTickMarks(t.instance, value)
}

func (t *TTrackBar) TickStyle() TTickStyle {
	return TrackBar_GetTickStyle(t.instance)
}

func (t *TTrackBar) SetTickStyle(value TTickStyle) {
	TrackBar_SetTickStyle(t.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (t *TTrackBar) Visible() bool {
	return TrackBar_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TTrackBar) SetVisible(value bool) {
	TrackBar_SetVisible(t.instance, value)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (t *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
	TrackBar_SetOnContextPopup(t.instance, fn)
}

// 设置改变事件。
//
// Set changed event.
func (t *TTrackBar) SetOnChange(fn TNotifyEvent) {
	TrackBar_SetOnChange(t.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
	TrackBar_SetOnDragDrop(t.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
	TrackBar_SetOnDragOver(t.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (t *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
	TrackBar_SetOnEndDrag(t.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (t *TTrackBar) SetOnEnter(fn TNotifyEvent) {
	TrackBar_SetOnEnter(t.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (t *TTrackBar) SetOnExit(fn TNotifyEvent) {
	TrackBar_SetOnExit(t.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (t *TTrackBar) SetOnKeyDown(fn TKeyEvent) {
	TrackBar_SetOnKeyDown(t.instance, fn)
}

// 设置键键下事件。
func (t *TTrackBar) SetOnKeyPress(fn TKeyPressEvent) {
	TrackBar_SetOnKeyPress(t.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (t *TTrackBar) SetOnKeyUp(fn TKeyEvent) {
	TrackBar_SetOnKeyUp(t.instance, fn)
}

// 获取依靠客户端总数。
func (t *TTrackBar) DockClientCount() int32 {
	return TrackBar_GetDockClientCount(t.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (t *TTrackBar) DockSite() bool {
	return TrackBar_GetDockSite(t.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (t *TTrackBar) SetDockSite(value bool) {
	TrackBar_SetDockSite(t.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (t *TTrackBar) MouseInClient() bool {
	return TrackBar_GetMouseInClient(t.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (t *TTrackBar) VisibleDockClientCount() int32 {
	return TrackBar_GetVisibleDockClientCount(t.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (t *TTrackBar) Brush() *TBrush {
	return AsBrush(TrackBar_GetBrush(t.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (t *TTrackBar) ControlCount() int32 {
	return TrackBar_GetControlCount(t.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (t *TTrackBar) Handle() HWND {
	return TrackBar_GetHandle(t.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (t *TTrackBar) ParentWindow() HWND {
	return TrackBar_GetParentWindow(t.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (t *TTrackBar) SetParentWindow(value HWND) {
	TrackBar_SetParentWindow(t.instance, value)
}

func (t *TTrackBar) Showing() bool {
	return TrackBar_GetShowing(t.instance)
}

// 获取使用停靠管理。
func (t *TTrackBar) UseDockManager() bool {
	return TrackBar_GetUseDockManager(t.instance)
}

// 设置使用停靠管理。
func (t *TTrackBar) SetUseDockManager(value bool) {
	TrackBar_SetUseDockManager(t.instance, value)
}

func (t *TTrackBar) Action() *TAction {
	return AsAction(TrackBar_GetAction(t.instance))
}

func (t *TTrackBar) SetAction(value IComponent) {
	TrackBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TTrackBar) BiDiMode() TBiDiMode {
	return TrackBar_GetBiDiMode(t.instance)
}

func (t *TTrackBar) SetBiDiMode(value TBiDiMode) {
	TrackBar_SetBiDiMode(t.instance, value)
}

func (t *TTrackBar) BoundsRect() TRect {
	return TrackBar_GetBoundsRect(t.instance)
}

func (t *TTrackBar) SetBoundsRect(value TRect) {
	TrackBar_SetBoundsRect(t.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (t *TTrackBar) ClientHeight() int32 {
	return TrackBar_GetClientHeight(t.instance)
}

// 设置客户区高度。
//
// Set client height.
func (t *TTrackBar) SetClientHeight(value int32) {
	TrackBar_SetClientHeight(t.instance, value)
}

func (t *TTrackBar) ClientOrigin() TPoint {
	return TrackBar_GetClientOrigin(t.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (t *TTrackBar) ClientRect() TRect {
	return TrackBar_GetClientRect(t.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (t *TTrackBar) ClientWidth() int32 {
	return TrackBar_GetClientWidth(t.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (t *TTrackBar) SetClientWidth(value int32) {
	TrackBar_SetClientWidth(t.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (t *TTrackBar) ControlState() TControlState {
	return TrackBar_GetControlState(t.instance)
}

// 设置控件状态。
//
// Set control state.
func (t *TTrackBar) SetControlState(value TControlState) {
	TrackBar_SetControlState(t.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (t *TTrackBar) ControlStyle() TControlStyle {
	return TrackBar_GetControlStyle(t.instance)
}

// 设置控件样式。
//
// Set control style.
func (t *TTrackBar) SetControlStyle(value TControlStyle) {
	TrackBar_SetControlStyle(t.instance, value)
}

func (t *TTrackBar) Floating() bool {
	return TrackBar_GetFloating(t.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (t *TTrackBar) Parent() *TWinControl {
	return AsWinControl(TrackBar_GetParent(t.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (t *TTrackBar) SetParent(value IWinControl) {
	TrackBar_SetParent(t.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (t *TTrackBar) Left() int32 {
	return TrackBar_GetLeft(t.instance)
}

// 设置左边位置。
//
// Set Left position.
func (t *TTrackBar) SetLeft(value int32) {
	TrackBar_SetLeft(t.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (t *TTrackBar) Top() int32 {
	return TrackBar_GetTop(t.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (t *TTrackBar) SetTop(value int32) {
	TrackBar_SetTop(t.instance, value)
}

// 获取宽度。
//
// Get width.
func (t *TTrackBar) Width() int32 {
	return TrackBar_GetWidth(t.instance)
}

// 设置宽度。
//
// Set width.
func (t *TTrackBar) SetWidth(value int32) {
	TrackBar_SetWidth(t.instance, value)
}

// 获取高度。
//
// Get height.
func (t *TTrackBar) Height() int32 {
	return TrackBar_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TTrackBar) SetHeight(value int32) {
	TrackBar_SetHeight(t.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (t *TTrackBar) Cursor() TCursor {
	return TrackBar_GetCursor(t.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (t *TTrackBar) SetCursor(value TCursor) {
	TrackBar_SetCursor(t.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TTrackBar) Hint() string {
	return TrackBar_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TTrackBar) SetHint(value string) {
	TrackBar_SetHint(t.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TTrackBar) ComponentCount() int32 {
	return TrackBar_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TTrackBar) ComponentIndex() int32 {
	return TrackBar_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TTrackBar) SetComponentIndex(value int32) {
	TrackBar_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTrackBar) Owner() *TComponent {
	return AsComponent(TrackBar_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TTrackBar) Name() string {
	return TrackBar_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTrackBar) SetName(value string) {
	TrackBar_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TTrackBar) Tag() int {
	return TrackBar_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TTrackBar) SetTag(value int) {
	TrackBar_SetTag(t.instance, value)
}

// 获取左边锚点。
func (t *TTrackBar) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(TrackBar_GetAnchorSideLeft(t.instance))
}

// 设置左边锚点。
func (t *TTrackBar) SetAnchorSideLeft(value *TAnchorSide) {
	TrackBar_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (t *TTrackBar) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(TrackBar_GetAnchorSideTop(t.instance))
}

// 设置顶边锚点。
func (t *TTrackBar) SetAnchorSideTop(value *TAnchorSide) {
	TrackBar_SetAnchorSideTop(t.instance, CheckPtr(value))
}

// 获取右边锚点。
func (t *TTrackBar) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(TrackBar_GetAnchorSideRight(t.instance))
}

// 设置右边锚点。
func (t *TTrackBar) SetAnchorSideRight(value *TAnchorSide) {
	TrackBar_SetAnchorSideRight(t.instance, CheckPtr(value))
}

// 获取底边锚点。
func (t *TTrackBar) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(TrackBar_GetAnchorSideBottom(t.instance))
}

// 设置底边锚点。
func (t *TTrackBar) SetAnchorSideBottom(value *TAnchorSide) {
	TrackBar_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

func (t *TTrackBar) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(TrackBar_GetChildSizing(t.instance))
}

func (t *TTrackBar) SetChildSizing(value *TControlChildSizing) {
	TrackBar_SetChildSizing(t.instance, CheckPtr(value))
}

// 获取边框间距。
func (t *TTrackBar) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(TrackBar_GetBorderSpacing(t.instance))
}

// 设置边框间距。
func (t *TTrackBar) SetBorderSpacing(value *TControlBorderSpacing) {
	TrackBar_SetBorderSpacing(t.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (t *TTrackBar) DockClients(Index int32) *TControl {
	return AsControl(TrackBar_GetDockClients(t.instance, Index))
}

// 获取指定索引子控件。
func (t *TTrackBar) Controls(Index int32) *TControl {
	return AsControl(TrackBar_GetControls(t.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TTrackBar) Components(AIndex int32) *TComponent {
	return AsComponent(TrackBar_GetComponents(t.instance, AIndex))
}

// 获取锚侧面。
func (t *TTrackBar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(TrackBar_GetAnchorSide(t.instance, AKind))
}
