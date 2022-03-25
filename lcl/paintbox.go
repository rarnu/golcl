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

type TPaintBox struct {
	IControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewPaintBox(owner IComponent) *TPaintBox {
	p := new(TPaintBox)
	p.instance = PaintBox_Create(CheckPtr(owner))
	p.ptr = unsafe.Pointer(p.instance)
	return p
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPaintBox(obj any) *TPaintBox {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPaintBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPaintBox.
func PaintBoxFromInst(inst uintptr) *TPaintBox {
	return AsPaintBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsPaintBox.
func PaintBoxFromObj(obj IObject) *TPaintBox {
	return AsPaintBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPaintBox.
func PaintBoxFromUnsafePointer(ptr unsafe.Pointer) *TPaintBox {
	return AsPaintBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (p *TPaintBox) Free() {
	if p.instance != 0 {
		PaintBox_Free(p.instance)
		p.instance, p.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPaintBox) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPaintBox) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPaintBox) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPaintBox) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPaintBox) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TPaintBoxClass() TClass {
	return PaintBox_StaticClassType()
}

// 将控件置于最前。
//
// Bring the control to the front.
func (p *TPaintBox) BringToFront() {
	PaintBox_BringToFront(p.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (p *TPaintBox) ClientToScreen(Point TPoint) TPoint {
	return PaintBox_ClientToScreen(p.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (p *TPaintBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return PaintBox_ClientToParent(p.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (p *TPaintBox) Dragging() bool {
	return PaintBox_Dragging(p.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (p *TPaintBox) HasParent() bool {
	return PaintBox_HasParent(p.instance)
}

// 隐藏控件。
//
// Hidden control.
func (p *TPaintBox) Hide() {
	PaintBox_Hide(p.instance)
}

// 要求重绘。
//
// Redraw.
func (p *TPaintBox) Invalidate() {
	PaintBox_Invalidate(p.instance)
}

// 发送一个消息。
//
// Send a message.
func (p *TPaintBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return PaintBox_Perform(p.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (p *TPaintBox) Refresh() {
	PaintBox_Refresh(p.instance)
}

// 重绘。
//
// Repaint.
func (p *TPaintBox) Repaint() {
	PaintBox_Repaint(p.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (p *TPaintBox) ScreenToClient(Point TPoint) TPoint {
	return PaintBox_ScreenToClient(p.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (p *TPaintBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return PaintBox_ParentToClient(p.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (p *TPaintBox) SendToBack() {
	PaintBox_SendToBack(p.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (p *TPaintBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	PaintBox_SetBounds(p.instance, ALeft, ATop, AWidth, AHeight)
}

// 显示控件。
//
// Show control.
func (p *TPaintBox) Show() {
	PaintBox_Show(p.instance)
}

// 控件更新。
//
// Update.
func (p *TPaintBox) Update() {
	PaintBox_Update(p.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (p *TPaintBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return PaintBox_GetTextBuf(p.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (p *TPaintBox) GetTextLen() int32 {
	return PaintBox_GetTextLen(p.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (p *TPaintBox) SetTextBuf(Buffer string) {
	PaintBox_SetTextBuf(p.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (p *TPaintBox) FindComponent(AName string) *TComponent {
	return AsComponent(PaintBox_FindComponent(p.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPaintBox) GetNamePath() string {
	return PaintBox_GetNamePath(p.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPaintBox) Assign(Source IObject) {
	PaintBox_Assign(p.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPaintBox) ClassType() TClass {
	return PaintBox_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPaintBox) ClassName() string {
	return PaintBox_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPaintBox) InstanceSize() int32 {
	return PaintBox_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPaintBox) InheritsFrom(AClass TClass) bool {
	return PaintBox_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPaintBox) Equals(Obj IObject) bool {
	return PaintBox_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPaintBox) GetHashCode() int32 {
	return PaintBox_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPaintBox) ToString() string {
	return PaintBox_ToString(p.instance)
}

func (p *TPaintBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	PaintBox_AnchorToNeighbour(p.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (p *TPaintBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	PaintBox_AnchorParallel(p.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (p *TPaintBox) AnchorHorizontalCenterTo(ASibling IControl) {
	PaintBox_AnchorHorizontalCenterTo(p.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (p *TPaintBox) AnchorVerticalCenterTo(ASibling IControl) {
	PaintBox_AnchorVerticalCenterTo(p.instance, CheckPtr(ASibling))
}

func (p *TPaintBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	PaintBox_AnchorSame(p.instance, ASide, CheckPtr(ASibling))
}

func (p *TPaintBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	PaintBox_AnchorAsAlign(p.instance, ATheAlign, ASpace)
}

func (p *TPaintBox) AnchorClient(ASpace int32) {
	PaintBox_AnchorClient(p.instance, ASpace)
}

func (p *TPaintBox) ScaleDesignToForm(ASize int32) int32 {
	return PaintBox_ScaleDesignToForm(p.instance, ASize)
}

func (p *TPaintBox) ScaleFormToDesign(ASize int32) int32 {
	return PaintBox_ScaleFormToDesign(p.instance, ASize)
}

func (p *TPaintBox) Scale96ToForm(ASize int32) int32 {
	return PaintBox_Scale96ToForm(p.instance, ASize)
}

func (p *TPaintBox) ScaleFormTo96(ASize int32) int32 {
	return PaintBox_ScaleFormTo96(p.instance, ASize)
}

func (p *TPaintBox) Scale96ToFont(ASize int32) int32 {
	return PaintBox_Scale96ToFont(p.instance, ASize)
}

func (p *TPaintBox) ScaleFontTo96(ASize int32) int32 {
	return PaintBox_ScaleFontTo96(p.instance, ASize)
}

func (p *TPaintBox) ScaleScreenToFont(ASize int32) int32 {
	return PaintBox_ScaleScreenToFont(p.instance, ASize)
}

func (p *TPaintBox) ScaleFontToScreen(ASize int32) int32 {
	return PaintBox_ScaleFontToScreen(p.instance, ASize)
}

func (p *TPaintBox) Scale96ToScreen(ASize int32) int32 {
	return PaintBox_Scale96ToScreen(p.instance, ASize)
}

func (p *TPaintBox) ScaleScreenTo96(ASize int32) int32 {
	return PaintBox_ScaleScreenTo96(p.instance, ASize)
}

func (p *TPaintBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	PaintBox_AutoAdjustLayout(p.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (p *TPaintBox) FixDesignFontsPPI(ADesignTimePPI int32) {
	PaintBox_FixDesignFontsPPI(p.instance, ADesignTimePPI)
}

func (p *TPaintBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	PaintBox_ScaleFontsPPI(p.instance, AToPPI, AProportion)
}

// 获取画布。
func (p *TPaintBox) Canvas() *TCanvas {
	return AsCanvas(PaintBox_GetCanvas(p.instance))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (p *TPaintBox) Align() TAlign {
	return PaintBox_GetAlign(p.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (p *TPaintBox) SetAlign(value TAlign) {
	PaintBox_SetAlign(p.instance, value)
}

// 获取四个角位置的锚点。
func (p *TPaintBox) Anchors() TAnchors {
	return PaintBox_GetAnchors(p.instance)
}

// 设置四个角位置的锚点。
func (p *TPaintBox) SetAnchors(value TAnchors) {
	PaintBox_SetAnchors(p.instance, value)
}

// 获取颜色。
//
// Get color.
func (p *TPaintBox) Color() TColor {
	return PaintBox_GetColor(p.instance)
}

// 设置颜色。
//
// Set color.
func (p *TPaintBox) SetColor(value TColor) {
	PaintBox_SetColor(p.instance, value)
}

// 获取约束控件大小。
func (p *TPaintBox) Constraints() *TSizeConstraints {
	return AsSizeConstraints(PaintBox_GetConstraints(p.instance))
}

// 设置约束控件大小。
func (p *TPaintBox) SetConstraints(value *TSizeConstraints) {
	PaintBox_SetConstraints(p.instance, CheckPtr(value))
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (p *TPaintBox) DragCursor() TCursor {
	return PaintBox_GetDragCursor(p.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (p *TPaintBox) SetDragCursor(value TCursor) {
	PaintBox_SetDragCursor(p.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (p *TPaintBox) DragMode() TDragMode {
	return PaintBox_GetDragMode(p.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (p *TPaintBox) SetDragMode(value TDragMode) {
	PaintBox_SetDragMode(p.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (p *TPaintBox) Enabled() bool {
	return PaintBox_GetEnabled(p.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (p *TPaintBox) SetEnabled(value bool) {
	PaintBox_SetEnabled(p.instance, value)
}

// 获取字体。
//
// Get Font.
func (p *TPaintBox) Font() *TFont {
	return AsFont(PaintBox_GetFont(p.instance))
}

// 设置字体。
//
// Set Font.
func (p *TPaintBox) SetFont(value *TFont) {
	PaintBox_SetFont(p.instance, CheckPtr(value))
}

// 获取使用父容器颜色。
//
// Get parent color.
func (p *TPaintBox) ParentColor() bool {
	return PaintBox_GetParentColor(p.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (p *TPaintBox) SetParentColor(value bool) {
	PaintBox_SetParentColor(p.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (p *TPaintBox) ParentFont() bool {
	return PaintBox_GetParentFont(p.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (p *TPaintBox) SetParentFont(value bool) {
	PaintBox_SetParentFont(p.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (p *TPaintBox) ParentShowHint() bool {
	return PaintBox_GetParentShowHint(p.instance)
}

// 设置以父容器的ShowHint属性为准。
func (p *TPaintBox) SetParentShowHint(value bool) {
	PaintBox_SetParentShowHint(p.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (p *TPaintBox) PopupMenu() *TPopupMenu {
	return AsPopupMenu(PaintBox_GetPopupMenu(p.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (p *TPaintBox) SetPopupMenu(value IComponent) {
	PaintBox_SetPopupMenu(p.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (p *TPaintBox) ShowHint() bool {
	return PaintBox_GetShowHint(p.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (p *TPaintBox) SetShowHint(value bool) {
	PaintBox_SetShowHint(p.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (p *TPaintBox) Visible() bool {
	return PaintBox_GetVisible(p.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (p *TPaintBox) SetVisible(value bool) {
	PaintBox_SetVisible(p.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (p *TPaintBox) SetOnClick(fn TNotifyEvent) {
	PaintBox_SetOnClick(p.instance, fn)
}

// 设置双击事件。
func (p *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
	PaintBox_SetOnDblClick(p.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (p *TPaintBox) SetOnDragDrop(fn TDragDropEvent) {
	PaintBox_SetOnDragDrop(p.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (p *TPaintBox) SetOnDragOver(fn TDragOverEvent) {
	PaintBox_SetOnDragOver(p.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (p *TPaintBox) SetOnEndDrag(fn TEndDragEvent) {
	PaintBox_SetOnEndDrag(p.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (p *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
	PaintBox_SetOnMouseDown(p.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (p *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
	PaintBox_SetOnMouseEnter(p.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (p *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
	PaintBox_SetOnMouseLeave(p.instance, fn)
}

// 设置鼠标移动事件。
func (p *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
	PaintBox_SetOnMouseMove(p.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (p *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
	PaintBox_SetOnMouseUp(p.instance, fn)
}

// 设置绘画事件。
func (p *TPaintBox) SetOnPaint(fn TNotifyEvent) {
	PaintBox_SetOnPaint(p.instance, fn)
}

func (p *TPaintBox) Action() *TAction {
	return AsAction(PaintBox_GetAction(p.instance))
}

func (p *TPaintBox) SetAction(value IComponent) {
	PaintBox_SetAction(p.instance, CheckPtr(value))
}

func (p *TPaintBox) BiDiMode() TBiDiMode {
	return PaintBox_GetBiDiMode(p.instance)
}

func (p *TPaintBox) SetBiDiMode(value TBiDiMode) {
	PaintBox_SetBiDiMode(p.instance, value)
}

func (p *TPaintBox) BoundsRect() TRect {
	return PaintBox_GetBoundsRect(p.instance)
}

func (p *TPaintBox) SetBoundsRect(value TRect) {
	PaintBox_SetBoundsRect(p.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (p *TPaintBox) ClientHeight() int32 {
	return PaintBox_GetClientHeight(p.instance)
}

// 设置客户区高度。
//
// Set client height.
func (p *TPaintBox) SetClientHeight(value int32) {
	PaintBox_SetClientHeight(p.instance, value)
}

func (p *TPaintBox) ClientOrigin() TPoint {
	return PaintBox_GetClientOrigin(p.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (p *TPaintBox) ClientRect() TRect {
	return PaintBox_GetClientRect(p.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (p *TPaintBox) ClientWidth() int32 {
	return PaintBox_GetClientWidth(p.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (p *TPaintBox) SetClientWidth(value int32) {
	PaintBox_SetClientWidth(p.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (p *TPaintBox) ControlState() TControlState {
	return PaintBox_GetControlState(p.instance)
}

// 设置控件状态。
//
// Set control state.
func (p *TPaintBox) SetControlState(value TControlState) {
	PaintBox_SetControlState(p.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (p *TPaintBox) ControlStyle() TControlStyle {
	return PaintBox_GetControlStyle(p.instance)
}

// 设置控件样式。
//
// Set control style.
func (p *TPaintBox) SetControlStyle(value TControlStyle) {
	PaintBox_SetControlStyle(p.instance, value)
}

func (p *TPaintBox) Floating() bool {
	return PaintBox_GetFloating(p.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (p *TPaintBox) Parent() *TWinControl {
	return AsWinControl(PaintBox_GetParent(p.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (p *TPaintBox) SetParent(value IWinControl) {
	PaintBox_SetParent(p.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (p *TPaintBox) Left() int32 {
	return PaintBox_GetLeft(p.instance)
}

// 设置左边位置。
//
// Set Left position.
func (p *TPaintBox) SetLeft(value int32) {
	PaintBox_SetLeft(p.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (p *TPaintBox) Top() int32 {
	return PaintBox_GetTop(p.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (p *TPaintBox) SetTop(value int32) {
	PaintBox_SetTop(p.instance, value)
}

// 获取宽度。
//
// Get width.
func (p *TPaintBox) Width() int32 {
	return PaintBox_GetWidth(p.instance)
}

// 设置宽度。
//
// Set width.
func (p *TPaintBox) SetWidth(value int32) {
	PaintBox_SetWidth(p.instance, value)
}

// 获取高度。
//
// Get height.
func (p *TPaintBox) Height() int32 {
	return PaintBox_GetHeight(p.instance)
}

// 设置高度。
//
// Set height.
func (p *TPaintBox) SetHeight(value int32) {
	PaintBox_SetHeight(p.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (p *TPaintBox) Cursor() TCursor {
	return PaintBox_GetCursor(p.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (p *TPaintBox) SetCursor(value TCursor) {
	PaintBox_SetCursor(p.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (p *TPaintBox) Hint() string {
	return PaintBox_GetHint(p.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (p *TPaintBox) SetHint(value string) {
	PaintBox_SetHint(p.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (p *TPaintBox) ComponentCount() int32 {
	return PaintBox_GetComponentCount(p.instance)
}

// 获取组件索引。
//
// Get component index.
func (p *TPaintBox) ComponentIndex() int32 {
	return PaintBox_GetComponentIndex(p.instance)
}

// 设置组件索引。
//
// Set component index.
func (p *TPaintBox) SetComponentIndex(value int32) {
	PaintBox_SetComponentIndex(p.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (p *TPaintBox) Owner() *TComponent {
	return AsComponent(PaintBox_GetOwner(p.instance))
}

// 获取组件名称。
//
// Get the component name.
func (p *TPaintBox) Name() string {
	return PaintBox_GetName(p.instance)
}

// 设置组件名称。
//
// Set the component name.
func (p *TPaintBox) SetName(value string) {
	PaintBox_SetName(p.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (p *TPaintBox) Tag() int {
	return PaintBox_GetTag(p.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (p *TPaintBox) SetTag(value int) {
	PaintBox_SetTag(p.instance, value)
}

// 获取左边锚点。
func (p *TPaintBox) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(PaintBox_GetAnchorSideLeft(p.instance))
}

// 设置左边锚点。
func (p *TPaintBox) SetAnchorSideLeft(value *TAnchorSide) {
	PaintBox_SetAnchorSideLeft(p.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (p *TPaintBox) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(PaintBox_GetAnchorSideTop(p.instance))
}

// 设置顶边锚点。
func (p *TPaintBox) SetAnchorSideTop(value *TAnchorSide) {
	PaintBox_SetAnchorSideTop(p.instance, CheckPtr(value))
}

// 获取右边锚点。
func (p *TPaintBox) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(PaintBox_GetAnchorSideRight(p.instance))
}

// 设置右边锚点。
func (p *TPaintBox) SetAnchorSideRight(value *TAnchorSide) {
	PaintBox_SetAnchorSideRight(p.instance, CheckPtr(value))
}

// 获取底边锚点。
func (p *TPaintBox) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(PaintBox_GetAnchorSideBottom(p.instance))
}

// 设置底边锚点。
func (p *TPaintBox) SetAnchorSideBottom(value *TAnchorSide) {
	PaintBox_SetAnchorSideBottom(p.instance, CheckPtr(value))
}

// 获取边框间距。
func (p *TPaintBox) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(PaintBox_GetBorderSpacing(p.instance))
}

// 设置边框间距。
func (p *TPaintBox) SetBorderSpacing(value *TControlBorderSpacing) {
	PaintBox_SetBorderSpacing(p.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (p *TPaintBox) Components(AIndex int32) *TComponent {
	return AsComponent(PaintBox_GetComponents(p.instance, AIndex))
}

// 获取锚侧面。
func (p *TPaintBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(PaintBox_GetAnchorSide(p.instance, AKind))
}
