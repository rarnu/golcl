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

type TLinkLabel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewLinkLabel(owner IComponent) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = LinkLabel_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsLinkLabel(obj interface{}) *TLinkLabel {
    instance, ptr := getInstance(obj)
    if instance == 0 {
        return nil
    }
    return &TLinkLabel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsLinkLabel.
func LinkLabelFromInst(inst uintptr) *TLinkLabel {
    return AsLinkLabel(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsLinkLabel.
func LinkLabelFromObj(obj IObject) *TLinkLabel {
    return AsLinkLabel(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsLinkLabel.
func LinkLabelFromUnsafePointer(ptr unsafe.Pointer) *TLinkLabel {
    return AsLinkLabel(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (l *TLinkLabel) Free() {
    if l.instance != 0 {
        LinkLabel_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TLinkLabel) Instance() uintptr {
    return l.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TLinkLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TLinkLabel) IsValid() bool {
    return l.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TLinkLabel) Is() TIs {
    return TIs(l.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TLinkLabel) As() TAs {
//    return TAs(l.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TLinkLabelClass() TClass {
    return LinkLabel_StaticClassType()
}

// 要求重绘。
//
// Redraw.
func (l *TLinkLabel) Invalidate() {
    LinkLabel_Invalidate(l.instance)
}

// 重绘。
//
// Repaint.
func (l *TLinkLabel) Repaint() {
    LinkLabel_Repaint(l.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (l *TLinkLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LinkLabel_SetBounds(l.instance, ALeft, ATop, AWidth, AHeight)
}

// 控件更新。
//
// Update.
func (l *TLinkLabel) Update() {
    LinkLabel_Update(l.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (l *TLinkLabel) BringToFront() {
    LinkLabel_BringToFront(l.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (l *TLinkLabel) ClientToScreen(Point TPoint) TPoint {
    return LinkLabel_ClientToScreen(l.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (l *TLinkLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ClientToParent(l.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (l *TLinkLabel) Dragging() bool {
    return LinkLabel_Dragging(l.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (l *TLinkLabel) HasParent() bool {
    return LinkLabel_HasParent(l.instance)
}

// 隐藏控件。
//
// Hidden control.
func (l *TLinkLabel) Hide() {
    LinkLabel_Hide(l.instance)
}

// 发送一个消息。
//
// Send a message.
func (l *TLinkLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LinkLabel_Perform(l.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (l *TLinkLabel) Refresh() {
    LinkLabel_Refresh(l.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (l *TLinkLabel) ScreenToClient(Point TPoint) TPoint {
    return LinkLabel_ScreenToClient(l.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (l *TLinkLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ParentToClient(l.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (l *TLinkLabel) SendToBack() {
    LinkLabel_SendToBack(l.instance)
}

// 显示控件。
//
// Show control.
func (l *TLinkLabel) Show() {
    LinkLabel_Show(l.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (l *TLinkLabel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return LinkLabel_GetTextBuf(l.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (l *TLinkLabel) GetTextLen() int32 {
    return LinkLabel_GetTextLen(l.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (l *TLinkLabel) SetTextBuf(Buffer string) {
    LinkLabel_SetTextBuf(l.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (l *TLinkLabel) FindComponent(AName string) *TComponent {
    return AsComponent(LinkLabel_FindComponent(l.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (l *TLinkLabel) GetNamePath() string {
    return LinkLabel_GetNamePath(l.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TLinkLabel) Assign(Source IObject) {
    LinkLabel_Assign(l.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (l *TLinkLabel) ClassType() TClass {
    return LinkLabel_ClassType(l.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TLinkLabel) ClassName() string {
    return LinkLabel_ClassName(l.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TLinkLabel) InstanceSize() int32 {
    return LinkLabel_InstanceSize(l.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TLinkLabel) InheritsFrom(AClass TClass) bool {
    return LinkLabel_InheritsFrom(l.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (l *TLinkLabel) Equals(Obj IObject) bool {
    return LinkLabel_Equals(l.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TLinkLabel) GetHashCode() int32 {
    return LinkLabel_GetHashCode(l.instance)
}

// 文本类信息。
//
// Text information.
func (l *TLinkLabel) ToString() string {
    return LinkLabel_ToString(l.instance)
}

func (l *TLinkLabel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    LinkLabel_AnchorToNeighbour(l.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (l *TLinkLabel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    LinkLabel_AnchorParallel(l.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (l *TLinkLabel) AnchorHorizontalCenterTo(ASibling IControl) {
    LinkLabel_AnchorHorizontalCenterTo(l.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (l *TLinkLabel) AnchorVerticalCenterTo(ASibling IControl) {
    LinkLabel_AnchorVerticalCenterTo(l.instance, CheckPtr(ASibling))
}

func (l *TLinkLabel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    LinkLabel_AnchorSame(l.instance, ASide, CheckPtr(ASibling))
}

func (l *TLinkLabel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    LinkLabel_AnchorAsAlign(l.instance, ATheAlign, ASpace)
}

func (l *TLinkLabel) AnchorClient(ASpace int32) {
    LinkLabel_AnchorClient(l.instance, ASpace)
}

func (l *TLinkLabel) ScaleDesignToForm(ASize int32) int32 {
    return LinkLabel_ScaleDesignToForm(l.instance, ASize)
}

func (l *TLinkLabel) ScaleFormToDesign(ASize int32) int32 {
    return LinkLabel_ScaleFormToDesign(l.instance, ASize)
}

func (l *TLinkLabel) Scale96ToForm(ASize int32) int32 {
    return LinkLabel_Scale96ToForm(l.instance, ASize)
}

func (l *TLinkLabel) ScaleFormTo96(ASize int32) int32 {
    return LinkLabel_ScaleFormTo96(l.instance, ASize)
}

func (l *TLinkLabel) Scale96ToFont(ASize int32) int32 {
    return LinkLabel_Scale96ToFont(l.instance, ASize)
}

func (l *TLinkLabel) ScaleFontTo96(ASize int32) int32 {
    return LinkLabel_ScaleFontTo96(l.instance, ASize)
}

func (l *TLinkLabel) ScaleScreenToFont(ASize int32) int32 {
    return LinkLabel_ScaleScreenToFont(l.instance, ASize)
}

func (l *TLinkLabel) ScaleFontToScreen(ASize int32) int32 {
    return LinkLabel_ScaleFontToScreen(l.instance, ASize)
}

func (l *TLinkLabel) Scale96ToScreen(ASize int32) int32 {
    return LinkLabel_Scale96ToScreen(l.instance, ASize)
}

func (l *TLinkLabel) ScaleScreenTo96(ASize int32) int32 {
    return LinkLabel_ScaleScreenTo96(l.instance, ASize)
}

func (l *TLinkLabel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    LinkLabel_AutoAdjustLayout(l.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (l *TLinkLabel) FixDesignFontsPPI(ADesignTimePPI int32) {
    LinkLabel_FixDesignFontsPPI(l.instance, ADesignTimePPI)
}

func (l *TLinkLabel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    LinkLabel_ScaleFontsPPI(l.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (l *TLinkLabel) Align() TAlign {
    return LinkLabel_GetAlign(l.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (l *TLinkLabel) SetAlign(value TAlign) {
    LinkLabel_SetAlign(l.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (l *TLinkLabel) Alignment() TLinkAlignment {
    return LinkLabel_GetAlignment(l.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (l *TLinkLabel) SetAlignment(value TLinkAlignment) {
    LinkLabel_SetAlignment(l.instance, value)
}

// 获取四个角位置的锚点。
func (l *TLinkLabel) Anchors() TAnchors {
    return LinkLabel_GetAnchors(l.instance)
}

// 设置四个角位置的锚点。
func (l *TLinkLabel) SetAnchors(value TAnchors) {
    LinkLabel_SetAnchors(l.instance, value)
}

// 获取自动调整大小。
func (l *TLinkLabel) AutoSize() bool {
    return LinkLabel_GetAutoSize(l.instance)
}

// 设置自动调整大小。
func (l *TLinkLabel) SetAutoSize(value bool) {
    LinkLabel_SetAutoSize(l.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (l *TLinkLabel) Caption() string {
    return LinkLabel_GetCaption(l.instance)
}

// 设置控件标题。
//
// Set the control title.
func (l *TLinkLabel) SetCaption(value string) {
    LinkLabel_SetCaption(l.instance, value)
}

// 获取颜色。
//
// Get color.
func (l *TLinkLabel) Color() TColor {
    return LinkLabel_GetColor(l.instance)
}

// 设置颜色。
//
// Set color.
func (l *TLinkLabel) SetColor(value TColor) {
    LinkLabel_SetColor(l.instance, value)
}

// 获取约束控件大小。
func (l *TLinkLabel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(LinkLabel_GetConstraints(l.instance))
}

// 设置约束控件大小。
func (l *TLinkLabel) SetConstraints(value *TSizeConstraints) {
    LinkLabel_SetConstraints(l.instance, CheckPtr(value))
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (l *TLinkLabel) DragCursor() TCursor {
    return LinkLabel_GetDragCursor(l.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (l *TLinkLabel) SetDragCursor(value TCursor) {
    LinkLabel_SetDragCursor(l.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (l *TLinkLabel) DragKind() TDragKind {
    return LinkLabel_GetDragKind(l.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (l *TLinkLabel) SetDragKind(value TDragKind) {
    LinkLabel_SetDragKind(l.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (l *TLinkLabel) DragMode() TDragMode {
    return LinkLabel_GetDragMode(l.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (l *TLinkLabel) SetDragMode(value TDragMode) {
    LinkLabel_SetDragMode(l.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (l *TLinkLabel) Enabled() bool {
    return LinkLabel_GetEnabled(l.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (l *TLinkLabel) SetEnabled(value bool) {
    LinkLabel_SetEnabled(l.instance, value)
}

// 获取字体。
//
// Get Font.
func (l *TLinkLabel) Font() *TFont {
    return AsFont(LinkLabel_GetFont(l.instance))
}

// 设置字体。
//
// Set Font.
func (l *TLinkLabel) SetFont(value *TFont) {
    LinkLabel_SetFont(l.instance, CheckPtr(value))
}

// 获取使用父容器颜色。
//
// Get parent color.
func (l *TLinkLabel) ParentColor() bool {
    return LinkLabel_GetParentColor(l.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (l *TLinkLabel) SetParentColor(value bool) {
    LinkLabel_SetParentColor(l.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (l *TLinkLabel) ParentFont() bool {
    return LinkLabel_GetParentFont(l.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (l *TLinkLabel) SetParentFont(value bool) {
    LinkLabel_SetParentFont(l.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (l *TLinkLabel) ParentShowHint() bool {
    return LinkLabel_GetParentShowHint(l.instance)
}

// 设置以父容器的ShowHint属性为准。
func (l *TLinkLabel) SetParentShowHint(value bool) {
    LinkLabel_SetParentShowHint(l.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (l *TLinkLabel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(LinkLabel_GetPopupMenu(l.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (l *TLinkLabel) SetPopupMenu(value IComponent) {
    LinkLabel_SetPopupMenu(l.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (l *TLinkLabel) ShowHint() bool {
    return LinkLabel_GetShowHint(l.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (l *TLinkLabel) SetShowHint(value bool) {
    LinkLabel_SetShowHint(l.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (l *TLinkLabel) Visible() bool {
    return LinkLabel_GetVisible(l.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (l *TLinkLabel) SetVisible(value bool) {
    LinkLabel_SetVisible(l.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (l *TLinkLabel) SetOnClick(fn TNotifyEvent) {
    LinkLabel_SetOnClick(l.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (l *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
    LinkLabel_SetOnContextPopup(l.instance, fn)
}

// 设置双击事件。
func (l *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
    LinkLabel_SetOnDblClick(l.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (l *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
    LinkLabel_SetOnDragDrop(l.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (l *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
    LinkLabel_SetOnDragOver(l.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (l *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
    LinkLabel_SetOnEndDrag(l.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (l *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
    LinkLabel_SetOnMouseDown(l.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (l *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
    LinkLabel_SetOnMouseEnter(l.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (l *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
    LinkLabel_SetOnMouseLeave(l.instance, fn)
}

// 设置鼠标移动事件。
func (l *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    LinkLabel_SetOnMouseMove(l.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (l *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
    LinkLabel_SetOnMouseUp(l.instance, fn)
}

func (l *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
    LinkLabel_SetOnLinkClick(l.instance, fn)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (l *TLinkLabel) MouseInClient() bool {
    return LinkLabel_GetMouseInClient(l.instance)
}

func (l *TLinkLabel) Action() *TAction {
    return AsAction(LinkLabel_GetAction(l.instance))
}

func (l *TLinkLabel) SetAction(value IComponent) {
    LinkLabel_SetAction(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) BiDiMode() TBiDiMode {
    return LinkLabel_GetBiDiMode(l.instance)
}

func (l *TLinkLabel) SetBiDiMode(value TBiDiMode) {
    LinkLabel_SetBiDiMode(l.instance, value)
}

func (l *TLinkLabel) BoundsRect() TRect {
    return LinkLabel_GetBoundsRect(l.instance)
}

func (l *TLinkLabel) SetBoundsRect(value TRect) {
    LinkLabel_SetBoundsRect(l.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (l *TLinkLabel) ClientHeight() int32 {
    return LinkLabel_GetClientHeight(l.instance)
}

// 设置客户区高度。
//
// Set client height.
func (l *TLinkLabel) SetClientHeight(value int32) {
    LinkLabel_SetClientHeight(l.instance, value)
}

func (l *TLinkLabel) ClientOrigin() TPoint {
    return LinkLabel_GetClientOrigin(l.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (l *TLinkLabel) ClientRect() TRect {
    return LinkLabel_GetClientRect(l.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (l *TLinkLabel) ClientWidth() int32 {
    return LinkLabel_GetClientWidth(l.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (l *TLinkLabel) SetClientWidth(value int32) {
    LinkLabel_SetClientWidth(l.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (l *TLinkLabel) ControlState() TControlState {
    return LinkLabel_GetControlState(l.instance)
}

// 设置控件状态。
//
// Set control state.
func (l *TLinkLabel) SetControlState(value TControlState) {
    LinkLabel_SetControlState(l.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (l *TLinkLabel) ControlStyle() TControlStyle {
    return LinkLabel_GetControlStyle(l.instance)
}

// 设置控件样式。
//
// Set control style.
func (l *TLinkLabel) SetControlStyle(value TControlStyle) {
    LinkLabel_SetControlStyle(l.instance, value)
}

func (l *TLinkLabel) Floating() bool {
    return LinkLabel_GetFloating(l.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (l *TLinkLabel) Parent() *TWinControl {
    return AsWinControl(LinkLabel_GetParent(l.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (l *TLinkLabel) SetParent(value IWinControl) {
    LinkLabel_SetParent(l.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (l *TLinkLabel) Left() int32 {
    return LinkLabel_GetLeft(l.instance)
}

// 设置左边位置。
//
// Set Left position.
func (l *TLinkLabel) SetLeft(value int32) {
    LinkLabel_SetLeft(l.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (l *TLinkLabel) Top() int32 {
    return LinkLabel_GetTop(l.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (l *TLinkLabel) SetTop(value int32) {
    LinkLabel_SetTop(l.instance, value)
}

// 获取宽度。
//
// Get width.
func (l *TLinkLabel) Width() int32 {
    return LinkLabel_GetWidth(l.instance)
}

// 设置宽度。
//
// Set width.
func (l *TLinkLabel) SetWidth(value int32) {
    LinkLabel_SetWidth(l.instance, value)
}

// 获取高度。
//
// Get height.
func (l *TLinkLabel) Height() int32 {
    return LinkLabel_GetHeight(l.instance)
}

// 设置高度。
//
// Set height.
func (l *TLinkLabel) SetHeight(value int32) {
    LinkLabel_SetHeight(l.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (l *TLinkLabel) Hint() string {
    return LinkLabel_GetHint(l.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (l *TLinkLabel) SetHint(value string) {
    LinkLabel_SetHint(l.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (l *TLinkLabel) ComponentCount() int32 {
    return LinkLabel_GetComponentCount(l.instance)
}

// 获取组件索引。
//
// Get component index.
func (l *TLinkLabel) ComponentIndex() int32 {
    return LinkLabel_GetComponentIndex(l.instance)
}

// 设置组件索引。
//
// Set component index.
func (l *TLinkLabel) SetComponentIndex(value int32) {
    LinkLabel_SetComponentIndex(l.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (l *TLinkLabel) Owner() *TComponent {
    return AsComponent(LinkLabel_GetOwner(l.instance))
}

// 获取组件名称。
//
// Get the component name.
func (l *TLinkLabel) Name() string {
    return LinkLabel_GetName(l.instance)
}

// 设置组件名称。
//
// Set the component name.
func (l *TLinkLabel) SetName(value string) {
    LinkLabel_SetName(l.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (l *TLinkLabel) Tag() int {
    return LinkLabel_GetTag(l.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (l *TLinkLabel) SetTag(value int) {
    LinkLabel_SetTag(l.instance, value)
}

// 获取左边锚点。
func (l *TLinkLabel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(LinkLabel_GetAnchorSideLeft(l.instance))
}

// 设置左边锚点。
func (l *TLinkLabel) SetAnchorSideLeft(value *TAnchorSide) {
    LinkLabel_SetAnchorSideLeft(l.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (l *TLinkLabel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(LinkLabel_GetAnchorSideTop(l.instance))
}

// 设置顶边锚点。
func (l *TLinkLabel) SetAnchorSideTop(value *TAnchorSide) {
    LinkLabel_SetAnchorSideTop(l.instance, CheckPtr(value))
}

// 获取右边锚点。
func (l *TLinkLabel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(LinkLabel_GetAnchorSideRight(l.instance))
}

// 设置右边锚点。
func (l *TLinkLabel) SetAnchorSideRight(value *TAnchorSide) {
    LinkLabel_SetAnchorSideRight(l.instance, CheckPtr(value))
}

// 获取底边锚点。
func (l *TLinkLabel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(LinkLabel_GetAnchorSideBottom(l.instance))
}

// 设置底边锚点。
func (l *TLinkLabel) SetAnchorSideBottom(value *TAnchorSide) {
    LinkLabel_SetAnchorSideBottom(l.instance, CheckPtr(value))
}

// 获取边框间距。
func (l *TLinkLabel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(LinkLabel_GetBorderSpacing(l.instance))
}

// 设置边框间距。
func (l *TLinkLabel) SetBorderSpacing(value *TControlBorderSpacing) {
    LinkLabel_SetBorderSpacing(l.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (l *TLinkLabel) Components(AIndex int32) *TComponent {
    return AsComponent(LinkLabel_GetComponents(l.instance, AIndex))
}

// 获取锚侧面。
func (l *TLinkLabel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(LinkLabel_GetAnchorSide(l.instance, AKind))
}

