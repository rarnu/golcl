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

type TPageControl struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewPageControl(owner IComponent) *TPageControl {
	p := new(TPageControl)
	p.instance = PageControl_Create(CheckPtr(owner))
	p.ptr = unsafe.Pointer(p.instance)
	return p
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPageControl(obj interface{}) *TPageControl {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPageControl{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPageControl.
func PageControlFromInst(inst uintptr) *TPageControl {
	return AsPageControl(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsPageControl.
func PageControlFromObj(obj IObject) *TPageControl {
	return AsPageControl(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPageControl.
func PageControlFromUnsafePointer(ptr unsafe.Pointer) *TPageControl {
	return AsPageControl(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (p *TPageControl) Free() {
	if p.instance != 0 {
		PageControl_Free(p.instance)
		p.instance, p.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPageControl) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPageControl) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPageControl) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPageControl) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPageControl) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TPageControlClass() TClass {
	return PageControl_StaticClassType()
}

func (p *TPageControl) SelectNextPage(GoForward bool, CheckTabVisible bool) {
	PageControl_SelectNextPage(p.instance, GoForward, CheckTabVisible)
}

func (p *TPageControl) TabRect(Index int32) TRect {
	return PageControl_TabRect(p.instance, Index)
}

// 是否可以获得焦点。
func (p *TPageControl) CanFocus() bool {
	return PageControl_CanFocus(p.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (p *TPageControl) ContainsControl(Control IControl) bool {
	return PageControl_ContainsControl(p.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (p *TPageControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(PageControl_ControlAtPos(p.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (p *TPageControl) DisableAlign() {
	PageControl_DisableAlign(p.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (p *TPageControl) EnableAlign() {
	PageControl_EnableAlign(p.instance)
}

// 查找子控件。
//
// Find sub controls.
func (p *TPageControl) FindChildControl(ControlName string) *TControl {
	return AsControl(PageControl_FindChildControl(p.instance, ControlName))
}

func (p *TPageControl) FlipChildren(AllLevels bool) {
	PageControl_FlipChildren(p.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (p *TPageControl) Focused() bool {
	return PageControl_Focused(p.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (p *TPageControl) HandleAllocated() bool {
	return PageControl_HandleAllocated(p.instance)
}

// 插入一个控件。
//
// Insert a control.
func (p *TPageControl) InsertControl(AControl IControl) {
	PageControl_InsertControl(p.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (p *TPageControl) Invalidate() {
	PageControl_Invalidate(p.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (p *TPageControl) PaintTo(DC HDC, X int32, Y int32) {
	PageControl_PaintTo(p.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (p *TPageControl) RemoveControl(AControl IControl) {
	PageControl_RemoveControl(p.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (p *TPageControl) Realign() {
	PageControl_Realign(p.instance)
}

// 重绘。
//
// Repaint.
func (p *TPageControl) Repaint() {
	PageControl_Repaint(p.instance)
}

// 按比例缩放。
//
// Scale by.
func (p *TPageControl) ScaleBy(M int32, D int32) {
	PageControl_ScaleBy(p.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (p *TPageControl) ScrollBy(DeltaX int32, DeltaY int32) {
	PageControl_ScrollBy(p.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (p *TPageControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	PageControl_SetBounds(p.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (p *TPageControl) SetFocus() {
	PageControl_SetFocus(p.instance)
}

// 控件更新。
//
// Update.
func (p *TPageControl) Update() {
	PageControl_Update(p.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (p *TPageControl) BringToFront() {
	PageControl_BringToFront(p.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (p *TPageControl) ClientToScreen(Point TPoint) TPoint {
	return PageControl_ClientToScreen(p.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (p *TPageControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return PageControl_ClientToParent(p.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (p *TPageControl) Dragging() bool {
	return PageControl_Dragging(p.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (p *TPageControl) HasParent() bool {
	return PageControl_HasParent(p.instance)
}

// 隐藏控件。
//
// Hidden control.
func (p *TPageControl) Hide() {
	PageControl_Hide(p.instance)
}

// 发送一个消息。
//
// Send a message.
func (p *TPageControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return PageControl_Perform(p.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (p *TPageControl) Refresh() {
	PageControl_Refresh(p.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (p *TPageControl) ScreenToClient(Point TPoint) TPoint {
	return PageControl_ScreenToClient(p.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (p *TPageControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return PageControl_ParentToClient(p.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (p *TPageControl) SendToBack() {
	PageControl_SendToBack(p.instance)
}

// 显示控件。
//
// Show control.
func (p *TPageControl) Show() {
	PageControl_Show(p.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (p *TPageControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return PageControl_GetTextBuf(p.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (p *TPageControl) GetTextLen() int32 {
	return PageControl_GetTextLen(p.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (p *TPageControl) SetTextBuf(Buffer string) {
	PageControl_SetTextBuf(p.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (p *TPageControl) FindComponent(AName string) *TComponent {
	return AsComponent(PageControl_FindComponent(p.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPageControl) GetNamePath() string {
	return PageControl_GetNamePath(p.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPageControl) Assign(Source IObject) {
	PageControl_Assign(p.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPageControl) ClassType() TClass {
	return PageControl_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPageControl) ClassName() string {
	return PageControl_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPageControl) InstanceSize() int32 {
	return PageControl_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPageControl) InheritsFrom(AClass TClass) bool {
	return PageControl_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPageControl) Equals(Obj IObject) bool {
	return PageControl_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPageControl) GetHashCode() int32 {
	return PageControl_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPageControl) ToString() string {
	return PageControl_ToString(p.instance)
}

func (p *TPageControl) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	PageControl_AnchorToNeighbour(p.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (p *TPageControl) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	PageControl_AnchorParallel(p.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (p *TPageControl) AnchorHorizontalCenterTo(ASibling IControl) {
	PageControl_AnchorHorizontalCenterTo(p.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (p *TPageControl) AnchorVerticalCenterTo(ASibling IControl) {
	PageControl_AnchorVerticalCenterTo(p.instance, CheckPtr(ASibling))
}

func (p *TPageControl) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	PageControl_AnchorSame(p.instance, ASide, CheckPtr(ASibling))
}

func (p *TPageControl) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	PageControl_AnchorAsAlign(p.instance, ATheAlign, ASpace)
}

func (p *TPageControl) AnchorClient(ASpace int32) {
	PageControl_AnchorClient(p.instance, ASpace)
}

func (p *TPageControl) ScaleDesignToForm(ASize int32) int32 {
	return PageControl_ScaleDesignToForm(p.instance, ASize)
}

func (p *TPageControl) ScaleFormToDesign(ASize int32) int32 {
	return PageControl_ScaleFormToDesign(p.instance, ASize)
}

func (p *TPageControl) Scale96ToForm(ASize int32) int32 {
	return PageControl_Scale96ToForm(p.instance, ASize)
}

func (p *TPageControl) ScaleFormTo96(ASize int32) int32 {
	return PageControl_ScaleFormTo96(p.instance, ASize)
}

func (p *TPageControl) Scale96ToFont(ASize int32) int32 {
	return PageControl_Scale96ToFont(p.instance, ASize)
}

func (p *TPageControl) ScaleFontTo96(ASize int32) int32 {
	return PageControl_ScaleFontTo96(p.instance, ASize)
}

func (p *TPageControl) ScaleScreenToFont(ASize int32) int32 {
	return PageControl_ScaleScreenToFont(p.instance, ASize)
}

func (p *TPageControl) ScaleFontToScreen(ASize int32) int32 {
	return PageControl_ScaleFontToScreen(p.instance, ASize)
}

func (p *TPageControl) Scale96ToScreen(ASize int32) int32 {
	return PageControl_Scale96ToScreen(p.instance, ASize)
}

func (p *TPageControl) ScaleScreenTo96(ASize int32) int32 {
	return PageControl_ScaleScreenTo96(p.instance, ASize)
}

func (p *TPageControl) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	PageControl_AutoAdjustLayout(p.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (p *TPageControl) FixDesignFontsPPI(ADesignTimePPI int32) {
	PageControl_FixDesignFontsPPI(p.instance, ADesignTimePPI)
}

func (p *TPageControl) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	PageControl_ScaleFontsPPI(p.instance, AToPPI, AProportion)
}

func (p *TPageControl) Options() TCTabControlOptions {
	return PageControl_GetOptions(p.instance)
}

func (p *TPageControl) SetOptions(value TCTabControlOptions) {
	PageControl_SetOptions(p.instance, value)
}

func (p *TPageControl) ActivePageIndex() int32 {
	return PageControl_GetActivePageIndex(p.instance)
}

func (p *TPageControl) SetActivePageIndex(value int32) {
	PageControl_SetActivePageIndex(p.instance, value)
}

func (p *TPageControl) PageCount() int32 {
	return PageControl_GetPageCount(p.instance)
}

func (p *TPageControl) ActivePage() *TTabSheet {
	return AsTabSheet(PageControl_GetActivePage(p.instance))
}

func (p *TPageControl) SetActivePage(value IWinControl) {
	PageControl_SetActivePage(p.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (p *TPageControl) Align() TAlign {
	return PageControl_GetAlign(p.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (p *TPageControl) SetAlign(value TAlign) {
	PageControl_SetAlign(p.instance, value)
}

// 获取四个角位置的锚点。
func (p *TPageControl) Anchors() TAnchors {
	return PageControl_GetAnchors(p.instance)
}

// 设置四个角位置的锚点。
func (p *TPageControl) SetAnchors(value TAnchors) {
	PageControl_SetAnchors(p.instance, value)
}

func (p *TPageControl) BiDiMode() TBiDiMode {
	return PageControl_GetBiDiMode(p.instance)
}

func (p *TPageControl) SetBiDiMode(value TBiDiMode) {
	PageControl_SetBiDiMode(p.instance, value)
}

// 获取约束控件大小。
func (p *TPageControl) Constraints() *TSizeConstraints {
	return AsSizeConstraints(PageControl_GetConstraints(p.instance))
}

// 设置约束控件大小。
func (p *TPageControl) SetConstraints(value *TSizeConstraints) {
	PageControl_SetConstraints(p.instance, CheckPtr(value))
}

// 获取停靠站点。
//
// Get Docking site.
func (p *TPageControl) DockSite() bool {
	return PageControl_GetDockSite(p.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (p *TPageControl) SetDockSite(value bool) {
	PageControl_SetDockSite(p.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (p *TPageControl) DoubleBuffered() bool {
	return PageControl_GetDoubleBuffered(p.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (p *TPageControl) SetDoubleBuffered(value bool) {
	PageControl_SetDoubleBuffered(p.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (p *TPageControl) DragCursor() TCursor {
	return PageControl_GetDragCursor(p.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (p *TPageControl) SetDragCursor(value TCursor) {
	PageControl_SetDragCursor(p.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (p *TPageControl) DragKind() TDragKind {
	return PageControl_GetDragKind(p.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (p *TPageControl) SetDragKind(value TDragKind) {
	PageControl_SetDragKind(p.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (p *TPageControl) DragMode() TDragMode {
	return PageControl_GetDragMode(p.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (p *TPageControl) SetDragMode(value TDragMode) {
	PageControl_SetDragMode(p.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (p *TPageControl) Enabled() bool {
	return PageControl_GetEnabled(p.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (p *TPageControl) SetEnabled(value bool) {
	PageControl_SetEnabled(p.instance, value)
}

// 获取字体。
//
// Get Font.
func (p *TPageControl) Font() *TFont {
	return AsFont(PageControl_GetFont(p.instance))
}

// 设置字体。
//
// Set Font.
func (p *TPageControl) SetFont(value *TFont) {
	PageControl_SetFont(p.instance, CheckPtr(value))
}

// 获取图标索引列表对象。
func (p *TPageControl) Images() *TImageList {
	return AsImageList(PageControl_GetImages(p.instance))
}

// 设置图标索引列表对象。
func (p *TPageControl) SetImages(value IComponent) {
	PageControl_SetImages(p.instance, CheckPtr(value))
}

func (p *TPageControl) MultiLine() bool {
	return PageControl_GetMultiLine(p.instance)
}

func (p *TPageControl) SetMultiLine(value bool) {
	PageControl_SetMultiLine(p.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (p *TPageControl) ParentDoubleBuffered() bool {
	return PageControl_GetParentDoubleBuffered(p.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (p *TPageControl) SetParentDoubleBuffered(value bool) {
	PageControl_SetParentDoubleBuffered(p.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (p *TPageControl) ParentFont() bool {
	return PageControl_GetParentFont(p.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (p *TPageControl) SetParentFont(value bool) {
	PageControl_SetParentFont(p.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (p *TPageControl) ParentShowHint() bool {
	return PageControl_GetParentShowHint(p.instance)
}

// 设置以父容器的ShowHint属性为准。
func (p *TPageControl) SetParentShowHint(value bool) {
	PageControl_SetParentShowHint(p.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (p *TPageControl) PopupMenu() *TPopupMenu {
	return AsPopupMenu(PageControl_GetPopupMenu(p.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (p *TPageControl) SetPopupMenu(value IComponent) {
	PageControl_SetPopupMenu(p.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (p *TPageControl) ShowHint() bool {
	return PageControl_GetShowHint(p.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (p *TPageControl) SetShowHint(value bool) {
	PageControl_SetShowHint(p.instance, value)
}

func (p *TPageControl) TabHeight() int16 {
	return PageControl_GetTabHeight(p.instance)
}

func (p *TPageControl) SetTabHeight(value int16) {
	PageControl_SetTabHeight(p.instance, value)
}

func (p *TPageControl) TabIndex() int32 {
	return PageControl_GetTabIndex(p.instance)
}

func (p *TPageControl) SetTabIndex(value int32) {
	PageControl_SetTabIndex(p.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (p *TPageControl) TabOrder() TTabOrder {
	return PageControl_GetTabOrder(p.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (p *TPageControl) SetTabOrder(value TTabOrder) {
	PageControl_SetTabOrder(p.instance, value)
}

func (p *TPageControl) TabPosition() TTabPosition {
	return PageControl_GetTabPosition(p.instance)
}

func (p *TPageControl) SetTabPosition(value TTabPosition) {
	PageControl_SetTabPosition(p.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (p *TPageControl) TabStop() bool {
	return PageControl_GetTabStop(p.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (p *TPageControl) SetTabStop(value bool) {
	PageControl_SetTabStop(p.instance, value)
}

func (p *TPageControl) TabWidth() int16 {
	return PageControl_GetTabWidth(p.instance)
}

func (p *TPageControl) SetTabWidth(value int16) {
	PageControl_SetTabWidth(p.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (p *TPageControl) Visible() bool {
	return PageControl_GetVisible(p.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (p *TPageControl) SetVisible(value bool) {
	PageControl_SetVisible(p.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (p *TPageControl) SetOnChange(fn TNotifyEvent) {
	PageControl_SetOnChange(p.instance, fn)
}

func (p *TPageControl) SetOnChanging(fn TTabChangingEvent) {
	PageControl_SetOnChanging(p.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (p *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
	PageControl_SetOnContextPopup(p.instance, fn)
}

func (p *TPageControl) SetOnDockDrop(fn TDockDropEvent) {
	PageControl_SetOnDockDrop(p.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (p *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
	PageControl_SetOnDragDrop(p.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (p *TPageControl) SetOnDragOver(fn TDragOverEvent) {
	PageControl_SetOnDragOver(p.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (p *TPageControl) SetOnEndDock(fn TEndDragEvent) {
	PageControl_SetOnEndDock(p.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (p *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
	PageControl_SetOnEndDrag(p.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (p *TPageControl) SetOnEnter(fn TNotifyEvent) {
	PageControl_SetOnEnter(p.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (p *TPageControl) SetOnExit(fn TNotifyEvent) {
	PageControl_SetOnExit(p.instance, fn)
}

func (p *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	PageControl_SetOnGetSiteInfo(p.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (p *TPageControl) SetOnMouseDown(fn TMouseEvent) {
	PageControl_SetOnMouseDown(p.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (p *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
	PageControl_SetOnMouseEnter(p.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (p *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
	PageControl_SetOnMouseLeave(p.instance, fn)
}

// 设置鼠标移动事件。
func (p *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
	PageControl_SetOnMouseMove(p.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (p *TPageControl) SetOnMouseUp(fn TMouseEvent) {
	PageControl_SetOnMouseUp(p.instance, fn)
}

// 设置大小被改变事件。
func (p *TPageControl) SetOnResize(fn TNotifyEvent) {
	PageControl_SetOnResize(p.instance, fn)
}

// 设置启动停靠。
func (p *TPageControl) SetOnStartDock(fn TStartDockEvent) {
	PageControl_SetOnStartDock(p.instance, fn)
}

func (p *TPageControl) SetOnUnDock(fn TUnDockEvent) {
	PageControl_SetOnUnDock(p.instance, fn)
}

// 获取依靠客户端总数。
func (p *TPageControl) DockClientCount() int32 {
	return PageControl_GetDockClientCount(p.instance)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (p *TPageControl) MouseInClient() bool {
	return PageControl_GetMouseInClient(p.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (p *TPageControl) VisibleDockClientCount() int32 {
	return PageControl_GetVisibleDockClientCount(p.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (p *TPageControl) Brush() *TBrush {
	return AsBrush(PageControl_GetBrush(p.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (p *TPageControl) ControlCount() int32 {
	return PageControl_GetControlCount(p.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (p *TPageControl) Handle() HWND {
	return PageControl_GetHandle(p.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (p *TPageControl) ParentWindow() HWND {
	return PageControl_GetParentWindow(p.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (p *TPageControl) SetParentWindow(value HWND) {
	PageControl_SetParentWindow(p.instance, value)
}

func (p *TPageControl) Showing() bool {
	return PageControl_GetShowing(p.instance)
}

// 获取使用停靠管理。
func (p *TPageControl) UseDockManager() bool {
	return PageControl_GetUseDockManager(p.instance)
}

// 设置使用停靠管理。
func (p *TPageControl) SetUseDockManager(value bool) {
	PageControl_SetUseDockManager(p.instance, value)
}

func (p *TPageControl) Action() *TAction {
	return AsAction(PageControl_GetAction(p.instance))
}

func (p *TPageControl) SetAction(value IComponent) {
	PageControl_SetAction(p.instance, CheckPtr(value))
}

func (p *TPageControl) BoundsRect() TRect {
	return PageControl_GetBoundsRect(p.instance)
}

func (p *TPageControl) SetBoundsRect(value TRect) {
	PageControl_SetBoundsRect(p.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (p *TPageControl) ClientHeight() int32 {
	return PageControl_GetClientHeight(p.instance)
}

// 设置客户区高度。
//
// Set client height.
func (p *TPageControl) SetClientHeight(value int32) {
	PageControl_SetClientHeight(p.instance, value)
}

func (p *TPageControl) ClientOrigin() TPoint {
	return PageControl_GetClientOrigin(p.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (p *TPageControl) ClientRect() TRect {
	return PageControl_GetClientRect(p.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (p *TPageControl) ClientWidth() int32 {
	return PageControl_GetClientWidth(p.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (p *TPageControl) SetClientWidth(value int32) {
	PageControl_SetClientWidth(p.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (p *TPageControl) ControlState() TControlState {
	return PageControl_GetControlState(p.instance)
}

// 设置控件状态。
//
// Set control state.
func (p *TPageControl) SetControlState(value TControlState) {
	PageControl_SetControlState(p.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (p *TPageControl) ControlStyle() TControlStyle {
	return PageControl_GetControlStyle(p.instance)
}

// 设置控件样式。
//
// Set control style.
func (p *TPageControl) SetControlStyle(value TControlStyle) {
	PageControl_SetControlStyle(p.instance, value)
}

func (p *TPageControl) Floating() bool {
	return PageControl_GetFloating(p.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (p *TPageControl) Parent() *TWinControl {
	return AsWinControl(PageControl_GetParent(p.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (p *TPageControl) SetParent(value IWinControl) {
	PageControl_SetParent(p.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (p *TPageControl) Left() int32 {
	return PageControl_GetLeft(p.instance)
}

// 设置左边位置。
//
// Set Left position.
func (p *TPageControl) SetLeft(value int32) {
	PageControl_SetLeft(p.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (p *TPageControl) Top() int32 {
	return PageControl_GetTop(p.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (p *TPageControl) SetTop(value int32) {
	PageControl_SetTop(p.instance, value)
}

// 获取宽度。
//
// Get width.
func (p *TPageControl) Width() int32 {
	return PageControl_GetWidth(p.instance)
}

// 设置宽度。
//
// Set width.
func (p *TPageControl) SetWidth(value int32) {
	PageControl_SetWidth(p.instance, value)
}

// 获取高度。
//
// Get height.
func (p *TPageControl) Height() int32 {
	return PageControl_GetHeight(p.instance)
}

// 设置高度。
//
// Set height.
func (p *TPageControl) SetHeight(value int32) {
	PageControl_SetHeight(p.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (p *TPageControl) Cursor() TCursor {
	return PageControl_GetCursor(p.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (p *TPageControl) SetCursor(value TCursor) {
	PageControl_SetCursor(p.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (p *TPageControl) Hint() string {
	return PageControl_GetHint(p.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (p *TPageControl) SetHint(value string) {
	PageControl_SetHint(p.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (p *TPageControl) ComponentCount() int32 {
	return PageControl_GetComponentCount(p.instance)
}

// 获取组件索引。
//
// Get component index.
func (p *TPageControl) ComponentIndex() int32 {
	return PageControl_GetComponentIndex(p.instance)
}

// 设置组件索引。
//
// Set component index.
func (p *TPageControl) SetComponentIndex(value int32) {
	PageControl_SetComponentIndex(p.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (p *TPageControl) Owner() *TComponent {
	return AsComponent(PageControl_GetOwner(p.instance))
}

// 获取组件名称。
//
// Get the component name.
func (p *TPageControl) Name() string {
	return PageControl_GetName(p.instance)
}

// 设置组件名称。
//
// Set the component name.
func (p *TPageControl) SetName(value string) {
	PageControl_SetName(p.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (p *TPageControl) Tag() int {
	return PageControl_GetTag(p.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (p *TPageControl) SetTag(value int) {
	PageControl_SetTag(p.instance, value)
}

// 获取左边锚点。
func (p *TPageControl) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(PageControl_GetAnchorSideLeft(p.instance))
}

// 设置左边锚点。
func (p *TPageControl) SetAnchorSideLeft(value *TAnchorSide) {
	PageControl_SetAnchorSideLeft(p.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (p *TPageControl) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(PageControl_GetAnchorSideTop(p.instance))
}

// 设置顶边锚点。
func (p *TPageControl) SetAnchorSideTop(value *TAnchorSide) {
	PageControl_SetAnchorSideTop(p.instance, CheckPtr(value))
}

// 获取右边锚点。
func (p *TPageControl) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(PageControl_GetAnchorSideRight(p.instance))
}

// 设置右边锚点。
func (p *TPageControl) SetAnchorSideRight(value *TAnchorSide) {
	PageControl_SetAnchorSideRight(p.instance, CheckPtr(value))
}

// 获取底边锚点。
func (p *TPageControl) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(PageControl_GetAnchorSideBottom(p.instance))
}

// 设置底边锚点。
func (p *TPageControl) SetAnchorSideBottom(value *TAnchorSide) {
	PageControl_SetAnchorSideBottom(p.instance, CheckPtr(value))
}

func (p *TPageControl) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(PageControl_GetChildSizing(p.instance))
}

func (p *TPageControl) SetChildSizing(value *TControlChildSizing) {
	PageControl_SetChildSizing(p.instance, CheckPtr(value))
}

// 获取边框间距。
func (p *TPageControl) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(PageControl_GetBorderSpacing(p.instance))
}

// 设置边框间距。
func (p *TPageControl) SetBorderSpacing(value *TControlBorderSpacing) {
	PageControl_SetBorderSpacing(p.instance, CheckPtr(value))
}

func (p *TPageControl) Pages(Index int32) *TTabSheet {
	return AsTabSheet(PageControl_GetPages(p.instance, Index))
}

// 获取指定索引停靠客户端。
func (p *TPageControl) DockClients(Index int32) *TControl {
	return AsControl(PageControl_GetDockClients(p.instance, Index))
}

// 获取指定索引子控件。
func (p *TPageControl) Controls(Index int32) *TControl {
	return AsControl(PageControl_GetControls(p.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (p *TPageControl) Components(AIndex int32) *TComponent {
	return AsComponent(PageControl_GetComponents(p.instance, AIndex))
}

// 获取锚侧面。
func (p *TPageControl) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(PageControl_GetAnchorSide(p.instance, AKind))
}
