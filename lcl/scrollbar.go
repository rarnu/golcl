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

type TScrollBar struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewScrollBar(owner IComponent) *TScrollBar {
	s := new(TScrollBar)
	s.instance = ScrollBar_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsScrollBar(obj any) *TScrollBar {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TScrollBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsScrollBar.
func ScrollBarFromInst(inst uintptr) *TScrollBar {
	return AsScrollBar(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsScrollBar.
func ScrollBarFromObj(obj IObject) *TScrollBar {
	return AsScrollBar(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsScrollBar.
func ScrollBarFromUnsafePointer(ptr unsafe.Pointer) *TScrollBar {
	return AsScrollBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (s *TScrollBar) Free() {
	if s.instance != 0 {
		ScrollBar_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TScrollBar) Instance() uintptr {
	return s.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TScrollBar) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TScrollBar) IsValid() bool {
	return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TScrollBar) Is() TIs {
	return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TScrollBar) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TScrollBarClass() TClass {
	return ScrollBar_StaticClassType()
}

func (s *TScrollBar) SetParams(APosition int32, AMin int32, AMax int32) {
	ScrollBar_SetParams(s.instance, APosition, AMin, AMax)
}

// 是否可以获得焦点。
func (s *TScrollBar) CanFocus() bool {
	return ScrollBar_CanFocus(s.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (s *TScrollBar) ContainsControl(Control IControl) bool {
	return ScrollBar_ContainsControl(s.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (s *TScrollBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(ScrollBar_ControlAtPos(s.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (s *TScrollBar) DisableAlign() {
	ScrollBar_DisableAlign(s.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (s *TScrollBar) EnableAlign() {
	ScrollBar_EnableAlign(s.instance)
}

// 查找子控件。
//
// Find sub controls.
func (s *TScrollBar) FindChildControl(ControlName string) *TControl {
	return AsControl(ScrollBar_FindChildControl(s.instance, ControlName))
}

func (s *TScrollBar) FlipChildren(AllLevels bool) {
	ScrollBar_FlipChildren(s.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (s *TScrollBar) Focused() bool {
	return ScrollBar_Focused(s.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (s *TScrollBar) HandleAllocated() bool {
	return ScrollBar_HandleAllocated(s.instance)
}

// 插入一个控件。
//
// Insert a control.
func (s *TScrollBar) InsertControl(AControl IControl) {
	ScrollBar_InsertControl(s.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (s *TScrollBar) Invalidate() {
	ScrollBar_Invalidate(s.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (s *TScrollBar) PaintTo(DC HDC, X int32, Y int32) {
	ScrollBar_PaintTo(s.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (s *TScrollBar) RemoveControl(AControl IControl) {
	ScrollBar_RemoveControl(s.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (s *TScrollBar) Realign() {
	ScrollBar_Realign(s.instance)
}

// 重绘。
//
// Repaint.
func (s *TScrollBar) Repaint() {
	ScrollBar_Repaint(s.instance)
}

// 按比例缩放。
//
// Scale by.
func (s *TScrollBar) ScaleBy(M int32, D int32) {
	ScrollBar_ScaleBy(s.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (s *TScrollBar) ScrollBy(DeltaX int32, DeltaY int32) {
	ScrollBar_ScrollBy(s.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (s *TScrollBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ScrollBar_SetBounds(s.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (s *TScrollBar) SetFocus() {
	ScrollBar_SetFocus(s.instance)
}

// 控件更新。
//
// Update.
func (s *TScrollBar) Update() {
	ScrollBar_Update(s.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (s *TScrollBar) BringToFront() {
	ScrollBar_BringToFront(s.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TScrollBar) ClientToScreen(Point TPoint) TPoint {
	return ScrollBar_ClientToScreen(s.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TScrollBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ScrollBar_ClientToParent(s.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TScrollBar) Dragging() bool {
	return ScrollBar_Dragging(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TScrollBar) HasParent() bool {
	return ScrollBar_HasParent(s.instance)
}

// 隐藏控件。
//
// Hidden control.
func (s *TScrollBar) Hide() {
	ScrollBar_Hide(s.instance)
}

// 发送一个消息。
//
// Send a message.
func (s *TScrollBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ScrollBar_Perform(s.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (s *TScrollBar) Refresh() {
	ScrollBar_Refresh(s.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TScrollBar) ScreenToClient(Point TPoint) TPoint {
	return ScrollBar_ScreenToClient(s.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TScrollBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ScrollBar_ParentToClient(s.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (s *TScrollBar) SendToBack() {
	ScrollBar_SendToBack(s.instance)
}

// 显示控件。
//
// Show control.
func (s *TScrollBar) Show() {
	ScrollBar_Show(s.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TScrollBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ScrollBar_GetTextBuf(s.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TScrollBar) GetTextLen() int32 {
	return ScrollBar_GetTextLen(s.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TScrollBar) SetTextBuf(Buffer string) {
	ScrollBar_SetTextBuf(s.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TScrollBar) FindComponent(AName string) *TComponent {
	return AsComponent(ScrollBar_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TScrollBar) GetNamePath() string {
	return ScrollBar_GetNamePath(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TScrollBar) Assign(Source IObject) {
	ScrollBar_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TScrollBar) ClassType() TClass {
	return ScrollBar_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TScrollBar) ClassName() string {
	return ScrollBar_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TScrollBar) InstanceSize() int32 {
	return ScrollBar_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TScrollBar) InheritsFrom(AClass TClass) bool {
	return ScrollBar_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TScrollBar) Equals(Obj IObject) bool {
	return ScrollBar_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TScrollBar) GetHashCode() int32 {
	return ScrollBar_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TScrollBar) ToString() string {
	return ScrollBar_ToString(s.instance)
}

func (s *TScrollBar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ScrollBar_AnchorToNeighbour(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (s *TScrollBar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ScrollBar_AnchorParallel(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (s *TScrollBar) AnchorHorizontalCenterTo(ASibling IControl) {
	ScrollBar_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (s *TScrollBar) AnchorVerticalCenterTo(ASibling IControl) {
	ScrollBar_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TScrollBar) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ScrollBar_AnchorSame(s.instance, ASide, CheckPtr(ASibling))
}

func (s *TScrollBar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ScrollBar_AnchorAsAlign(s.instance, ATheAlign, ASpace)
}

func (s *TScrollBar) AnchorClient(ASpace int32) {
	ScrollBar_AnchorClient(s.instance, ASpace)
}

func (s *TScrollBar) ScaleDesignToForm(ASize int32) int32 {
	return ScrollBar_ScaleDesignToForm(s.instance, ASize)
}

func (s *TScrollBar) ScaleFormToDesign(ASize int32) int32 {
	return ScrollBar_ScaleFormToDesign(s.instance, ASize)
}

func (s *TScrollBar) Scale96ToForm(ASize int32) int32 {
	return ScrollBar_Scale96ToForm(s.instance, ASize)
}

func (s *TScrollBar) ScaleFormTo96(ASize int32) int32 {
	return ScrollBar_ScaleFormTo96(s.instance, ASize)
}

func (s *TScrollBar) Scale96ToFont(ASize int32) int32 {
	return ScrollBar_Scale96ToFont(s.instance, ASize)
}

func (s *TScrollBar) ScaleFontTo96(ASize int32) int32 {
	return ScrollBar_ScaleFontTo96(s.instance, ASize)
}

func (s *TScrollBar) ScaleScreenToFont(ASize int32) int32 {
	return ScrollBar_ScaleScreenToFont(s.instance, ASize)
}

func (s *TScrollBar) ScaleFontToScreen(ASize int32) int32 {
	return ScrollBar_ScaleFontToScreen(s.instance, ASize)
}

func (s *TScrollBar) Scale96ToScreen(ASize int32) int32 {
	return ScrollBar_Scale96ToScreen(s.instance, ASize)
}

func (s *TScrollBar) ScaleScreenTo96(ASize int32) int32 {
	return ScrollBar_ScaleScreenTo96(s.instance, ASize)
}

func (s *TScrollBar) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ScrollBar_AutoAdjustLayout(s.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (s *TScrollBar) FixDesignFontsPPI(ADesignTimePPI int32) {
	ScrollBar_FixDesignFontsPPI(s.instance, ADesignTimePPI)
}

func (s *TScrollBar) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ScrollBar_ScaleFontsPPI(s.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TScrollBar) Align() TAlign {
	return ScrollBar_GetAlign(s.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TScrollBar) SetAlign(value TAlign) {
	ScrollBar_SetAlign(s.instance, value)
}

// 获取四个角位置的锚点。
func (s *TScrollBar) Anchors() TAnchors {
	return ScrollBar_GetAnchors(s.instance)
}

// 设置四个角位置的锚点。
func (s *TScrollBar) SetAnchors(value TAnchors) {
	ScrollBar_SetAnchors(s.instance, value)
}

func (s *TScrollBar) BiDiMode() TBiDiMode {
	return ScrollBar_GetBiDiMode(s.instance)
}

func (s *TScrollBar) SetBiDiMode(value TBiDiMode) {
	ScrollBar_SetBiDiMode(s.instance, value)
}

// 获取约束控件大小。
func (s *TScrollBar) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ScrollBar_GetConstraints(s.instance))
}

// 设置约束控件大小。
func (s *TScrollBar) SetConstraints(value *TSizeConstraints) {
	ScrollBar_SetConstraints(s.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (s *TScrollBar) DoubleBuffered() bool {
	return ScrollBar_GetDoubleBuffered(s.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (s *TScrollBar) SetDoubleBuffered(value bool) {
	ScrollBar_SetDoubleBuffered(s.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (s *TScrollBar) DragCursor() TCursor {
	return ScrollBar_GetDragCursor(s.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (s *TScrollBar) SetDragCursor(value TCursor) {
	ScrollBar_SetDragCursor(s.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (s *TScrollBar) DragKind() TDragKind {
	return ScrollBar_GetDragKind(s.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (s *TScrollBar) SetDragKind(value TDragKind) {
	ScrollBar_SetDragKind(s.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (s *TScrollBar) DragMode() TDragMode {
	return ScrollBar_GetDragMode(s.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (s *TScrollBar) SetDragMode(value TDragMode) {
	ScrollBar_SetDragMode(s.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (s *TScrollBar) Enabled() bool {
	return ScrollBar_GetEnabled(s.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (s *TScrollBar) SetEnabled(value bool) {
	ScrollBar_SetEnabled(s.instance, value)
}

func (s *TScrollBar) Kind() TScrollBarKind {
	return ScrollBar_GetKind(s.instance)
}

func (s *TScrollBar) SetKind(value TScrollBarKind) {
	ScrollBar_SetKind(s.instance, value)
}

func (s *TScrollBar) LargeChange() TScrollBarInc {
	return ScrollBar_GetLargeChange(s.instance)
}

func (s *TScrollBar) SetLargeChange(value TScrollBarInc) {
	ScrollBar_SetLargeChange(s.instance, value)
}

func (s *TScrollBar) Max() int32 {
	return ScrollBar_GetMax(s.instance)
}

func (s *TScrollBar) SetMax(value int32) {
	ScrollBar_SetMax(s.instance, value)
}

func (s *TScrollBar) Min() int32 {
	return ScrollBar_GetMin(s.instance)
}

func (s *TScrollBar) SetMin(value int32) {
	ScrollBar_SetMin(s.instance, value)
}

func (s *TScrollBar) PageSize() int32 {
	return ScrollBar_GetPageSize(s.instance)
}

func (s *TScrollBar) SetPageSize(value int32) {
	ScrollBar_SetPageSize(s.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (s *TScrollBar) ParentDoubleBuffered() bool {
	return ScrollBar_GetParentDoubleBuffered(s.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (s *TScrollBar) SetParentDoubleBuffered(value bool) {
	ScrollBar_SetParentDoubleBuffered(s.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (s *TScrollBar) ParentShowHint() bool {
	return ScrollBar_GetParentShowHint(s.instance)
}

// 设置以父容器的ShowHint属性为准。
func (s *TScrollBar) SetParentShowHint(value bool) {
	ScrollBar_SetParentShowHint(s.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (s *TScrollBar) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ScrollBar_GetPopupMenu(s.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (s *TScrollBar) SetPopupMenu(value IComponent) {
	ScrollBar_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TScrollBar) Position() int32 {
	return ScrollBar_GetPosition(s.instance)
}

func (s *TScrollBar) SetPosition(value int32) {
	ScrollBar_SetPosition(s.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TScrollBar) ShowHint() bool {
	return ScrollBar_GetShowHint(s.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TScrollBar) SetShowHint(value bool) {
	ScrollBar_SetShowHint(s.instance, value)
}

func (s *TScrollBar) SmallChange() TScrollBarInc {
	return ScrollBar_GetSmallChange(s.instance)
}

func (s *TScrollBar) SetSmallChange(value TScrollBarInc) {
	ScrollBar_SetSmallChange(s.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (s *TScrollBar) TabOrder() TTabOrder {
	return ScrollBar_GetTabOrder(s.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (s *TScrollBar) SetTabOrder(value TTabOrder) {
	ScrollBar_SetTabOrder(s.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (s *TScrollBar) TabStop() bool {
	return ScrollBar_GetTabStop(s.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (s *TScrollBar) SetTabStop(value bool) {
	ScrollBar_SetTabStop(s.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (s *TScrollBar) Visible() bool {
	return ScrollBar_GetVisible(s.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (s *TScrollBar) SetVisible(value bool) {
	ScrollBar_SetVisible(s.instance, value)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (s *TScrollBar) SetOnContextPopup(fn TContextPopupEvent) {
	ScrollBar_SetOnContextPopup(s.instance, fn)
}

// 设置改变事件。
//
// Set changed event.
func (s *TScrollBar) SetOnChange(fn TNotifyEvent) {
	ScrollBar_SetOnChange(s.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (s *TScrollBar) SetOnDragDrop(fn TDragDropEvent) {
	ScrollBar_SetOnDragDrop(s.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (s *TScrollBar) SetOnDragOver(fn TDragOverEvent) {
	ScrollBar_SetOnDragOver(s.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (s *TScrollBar) SetOnEndDrag(fn TEndDragEvent) {
	ScrollBar_SetOnEndDrag(s.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (s *TScrollBar) SetOnEnter(fn TNotifyEvent) {
	ScrollBar_SetOnEnter(s.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (s *TScrollBar) SetOnExit(fn TNotifyEvent) {
	ScrollBar_SetOnExit(s.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (s *TScrollBar) SetOnKeyDown(fn TKeyEvent) {
	ScrollBar_SetOnKeyDown(s.instance, fn)
}

// 设置键键下事件。
func (s *TScrollBar) SetOnKeyPress(fn TKeyPressEvent) {
	ScrollBar_SetOnKeyPress(s.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (s *TScrollBar) SetOnKeyUp(fn TKeyEvent) {
	ScrollBar_SetOnKeyUp(s.instance, fn)
}

// 获取依靠客户端总数。
func (s *TScrollBar) DockClientCount() int32 {
	return ScrollBar_GetDockClientCount(s.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (s *TScrollBar) DockSite() bool {
	return ScrollBar_GetDockSite(s.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (s *TScrollBar) SetDockSite(value bool) {
	ScrollBar_SetDockSite(s.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (s *TScrollBar) MouseInClient() bool {
	return ScrollBar_GetMouseInClient(s.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (s *TScrollBar) VisibleDockClientCount() int32 {
	return ScrollBar_GetVisibleDockClientCount(s.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (s *TScrollBar) Brush() *TBrush {
	return AsBrush(ScrollBar_GetBrush(s.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (s *TScrollBar) ControlCount() int32 {
	return ScrollBar_GetControlCount(s.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (s *TScrollBar) Handle() HWND {
	return ScrollBar_GetHandle(s.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (s *TScrollBar) ParentWindow() HWND {
	return ScrollBar_GetParentWindow(s.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (s *TScrollBar) SetParentWindow(value HWND) {
	ScrollBar_SetParentWindow(s.instance, value)
}

func (s *TScrollBar) Showing() bool {
	return ScrollBar_GetShowing(s.instance)
}

// 获取使用停靠管理。
func (s *TScrollBar) UseDockManager() bool {
	return ScrollBar_GetUseDockManager(s.instance)
}

// 设置使用停靠管理。
func (s *TScrollBar) SetUseDockManager(value bool) {
	ScrollBar_SetUseDockManager(s.instance, value)
}

func (s *TScrollBar) Action() *TAction {
	return AsAction(ScrollBar_GetAction(s.instance))
}

func (s *TScrollBar) SetAction(value IComponent) {
	ScrollBar_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBar) BoundsRect() TRect {
	return ScrollBar_GetBoundsRect(s.instance)
}

func (s *TScrollBar) SetBoundsRect(value TRect) {
	ScrollBar_SetBoundsRect(s.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (s *TScrollBar) ClientHeight() int32 {
	return ScrollBar_GetClientHeight(s.instance)
}

// 设置客户区高度。
//
// Set client height.
func (s *TScrollBar) SetClientHeight(value int32) {
	ScrollBar_SetClientHeight(s.instance, value)
}

func (s *TScrollBar) ClientOrigin() TPoint {
	return ScrollBar_GetClientOrigin(s.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (s *TScrollBar) ClientRect() TRect {
	return ScrollBar_GetClientRect(s.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (s *TScrollBar) ClientWidth() int32 {
	return ScrollBar_GetClientWidth(s.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (s *TScrollBar) SetClientWidth(value int32) {
	ScrollBar_SetClientWidth(s.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (s *TScrollBar) ControlState() TControlState {
	return ScrollBar_GetControlState(s.instance)
}

// 设置控件状态。
//
// Set control state.
func (s *TScrollBar) SetControlState(value TControlState) {
	ScrollBar_SetControlState(s.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (s *TScrollBar) ControlStyle() TControlStyle {
	return ScrollBar_GetControlStyle(s.instance)
}

// 设置控件样式。
//
// Set control style.
func (s *TScrollBar) SetControlStyle(value TControlStyle) {
	ScrollBar_SetControlStyle(s.instance, value)
}

func (s *TScrollBar) Floating() bool {
	return ScrollBar_GetFloating(s.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (s *TScrollBar) Parent() *TWinControl {
	return AsWinControl(ScrollBar_GetParent(s.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (s *TScrollBar) SetParent(value IWinControl) {
	ScrollBar_SetParent(s.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (s *TScrollBar) Left() int32 {
	return ScrollBar_GetLeft(s.instance)
}

// 设置左边位置。
//
// Set Left position.
func (s *TScrollBar) SetLeft(value int32) {
	ScrollBar_SetLeft(s.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (s *TScrollBar) Top() int32 {
	return ScrollBar_GetTop(s.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (s *TScrollBar) SetTop(value int32) {
	ScrollBar_SetTop(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TScrollBar) Width() int32 {
	return ScrollBar_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TScrollBar) SetWidth(value int32) {
	ScrollBar_SetWidth(s.instance, value)
}

// 获取高度。
//
// Get height.
func (s *TScrollBar) Height() int32 {
	return ScrollBar_GetHeight(s.instance)
}

// 设置高度。
//
// Set height.
func (s *TScrollBar) SetHeight(value int32) {
	ScrollBar_SetHeight(s.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (s *TScrollBar) Cursor() TCursor {
	return ScrollBar_GetCursor(s.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (s *TScrollBar) SetCursor(value TCursor) {
	ScrollBar_SetCursor(s.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TScrollBar) Hint() string {
	return ScrollBar_GetHint(s.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TScrollBar) SetHint(value string) {
	ScrollBar_SetHint(s.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TScrollBar) ComponentCount() int32 {
	return ScrollBar_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TScrollBar) ComponentIndex() int32 {
	return ScrollBar_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TScrollBar) SetComponentIndex(value int32) {
	ScrollBar_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TScrollBar) Owner() *TComponent {
	return AsComponent(ScrollBar_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TScrollBar) Name() string {
	return ScrollBar_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TScrollBar) SetName(value string) {
	ScrollBar_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TScrollBar) Tag() int {
	return ScrollBar_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TScrollBar) SetTag(value int) {
	ScrollBar_SetTag(s.instance, value)
}

// 获取左边锚点。
func (s *TScrollBar) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ScrollBar_GetAnchorSideLeft(s.instance))
}

// 设置左边锚点。
func (s *TScrollBar) SetAnchorSideLeft(value *TAnchorSide) {
	ScrollBar_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (s *TScrollBar) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ScrollBar_GetAnchorSideTop(s.instance))
}

// 设置顶边锚点。
func (s *TScrollBar) SetAnchorSideTop(value *TAnchorSide) {
	ScrollBar_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// 获取右边锚点。
func (s *TScrollBar) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ScrollBar_GetAnchorSideRight(s.instance))
}

// 设置右边锚点。
func (s *TScrollBar) SetAnchorSideRight(value *TAnchorSide) {
	ScrollBar_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// 获取底边锚点。
func (s *TScrollBar) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ScrollBar_GetAnchorSideBottom(s.instance))
}

// 设置底边锚点。
func (s *TScrollBar) SetAnchorSideBottom(value *TAnchorSide) {
	ScrollBar_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TScrollBar) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(ScrollBar_GetChildSizing(s.instance))
}

func (s *TScrollBar) SetChildSizing(value *TControlChildSizing) {
	ScrollBar_SetChildSizing(s.instance, CheckPtr(value))
}

// 获取边框间距。
func (s *TScrollBar) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ScrollBar_GetBorderSpacing(s.instance))
}

// 设置边框间距。
func (s *TScrollBar) SetBorderSpacing(value *TControlBorderSpacing) {
	ScrollBar_SetBorderSpacing(s.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (s *TScrollBar) DockClients(Index int32) *TControl {
	return AsControl(ScrollBar_GetDockClients(s.instance, Index))
}

// 获取指定索引子控件。
func (s *TScrollBar) Controls(Index int32) *TControl {
	return AsControl(ScrollBar_GetControls(s.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TScrollBar) Components(AIndex int32) *TComponent {
	return AsComponent(ScrollBar_GetComponents(s.instance, AIndex))
}

// 获取锚侧面。
func (s *TScrollBar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ScrollBar_GetAnchorSide(s.instance, AKind))
}
