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

type TXButton struct {
	IControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewXButton(owner IComponent) *TXButton {
	x := new(TXButton)
	x.instance = XButton_Create(CheckPtr(owner))
	x.ptr = unsafe.Pointer(x.instance)
	return x
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsXButton(obj any) *TXButton {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TXButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsXButton.
func XButtonFromInst(inst uintptr) *TXButton {
	return AsXButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsXButton.
func XButtonFromObj(obj IObject) *TXButton {
	return AsXButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsXButton.
func XButtonFromUnsafePointer(ptr unsafe.Pointer) *TXButton {
	return AsXButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (x *TXButton) Free() {
	if x.instance != 0 {
		XButton_Free(x.instance)
		x.instance, x.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (x *TXButton) Instance() uintptr {
	return x.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (x *TXButton) UnsafeAddr() unsafe.Pointer {
	return x.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (x *TXButton) IsValid() bool {
	return x.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (x *TXButton) Is() TIs {
	return TIs(x.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (x *TXButton) As() TAs {
//    return TAs(x.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TXButtonClass() TClass {
	return XButton_StaticClassType()
}

// 将控件置于最前。
//
// Bring the control to the front.
func (x *TXButton) BringToFront() {
	XButton_BringToFront(x.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (x *TXButton) ClientToScreen(Point TPoint) TPoint {
	return XButton_ClientToScreen(x.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (x *TXButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return XButton_ClientToParent(x.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (x *TXButton) Dragging() bool {
	return XButton_Dragging(x.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (x *TXButton) HasParent() bool {
	return XButton_HasParent(x.instance)
}

// 隐藏控件。
//
// Hidden control.
func (x *TXButton) Hide() {
	XButton_Hide(x.instance)
}

// 要求重绘。
//
// Redraw.
func (x *TXButton) Invalidate() {
	XButton_Invalidate(x.instance)
}

// 发送一个消息。
//
// Send a message.
func (x *TXButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return XButton_Perform(x.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (x *TXButton) Refresh() {
	XButton_Refresh(x.instance)
}

// 重绘。
//
// Repaint.
func (x *TXButton) Repaint() {
	XButton_Repaint(x.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (x *TXButton) ScreenToClient(Point TPoint) TPoint {
	return XButton_ScreenToClient(x.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (x *TXButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return XButton_ParentToClient(x.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (x *TXButton) SendToBack() {
	XButton_SendToBack(x.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (x *TXButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	XButton_SetBounds(x.instance, ALeft, ATop, AWidth, AHeight)
}

// 显示控件。
//
// Show control.
func (x *TXButton) Show() {
	XButton_Show(x.instance)
}

// 控件更新。
//
// Update.
func (x *TXButton) Update() {
	XButton_Update(x.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (x *TXButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return XButton_GetTextBuf(x.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (x *TXButton) GetTextLen() int32 {
	return XButton_GetTextLen(x.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (x *TXButton) SetTextBuf(Buffer string) {
	XButton_SetTextBuf(x.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (x *TXButton) FindComponent(AName string) *TComponent {
	return AsComponent(XButton_FindComponent(x.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (x *TXButton) GetNamePath() string {
	return XButton_GetNamePath(x.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (x *TXButton) Assign(Source IObject) {
	XButton_Assign(x.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (x *TXButton) ClassType() TClass {
	return XButton_ClassType(x.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (x *TXButton) ClassName() string {
	return XButton_ClassName(x.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (x *TXButton) InstanceSize() int32 {
	return XButton_InstanceSize(x.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (x *TXButton) InheritsFrom(AClass TClass) bool {
	return XButton_InheritsFrom(x.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (x *TXButton) Equals(Obj IObject) bool {
	return XButton_Equals(x.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (x *TXButton) GetHashCode() int32 {
	return XButton_GetHashCode(x.instance)
}

// 文本类信息。
//
// Text information.
func (x *TXButton) ToString() string {
	return XButton_ToString(x.instance)
}

func (x *TXButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	XButton_AnchorToNeighbour(x.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (x *TXButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	XButton_AnchorParallel(x.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (x *TXButton) AnchorHorizontalCenterTo(ASibling IControl) {
	XButton_AnchorHorizontalCenterTo(x.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (x *TXButton) AnchorVerticalCenterTo(ASibling IControl) {
	XButton_AnchorVerticalCenterTo(x.instance, CheckPtr(ASibling))
}

func (x *TXButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	XButton_AnchorSame(x.instance, ASide, CheckPtr(ASibling))
}

func (x *TXButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	XButton_AnchorAsAlign(x.instance, ATheAlign, ASpace)
}

func (x *TXButton) AnchorClient(ASpace int32) {
	XButton_AnchorClient(x.instance, ASpace)
}

func (x *TXButton) ScaleDesignToForm(ASize int32) int32 {
	return XButton_ScaleDesignToForm(x.instance, ASize)
}

func (x *TXButton) ScaleFormToDesign(ASize int32) int32 {
	return XButton_ScaleFormToDesign(x.instance, ASize)
}

func (x *TXButton) Scale96ToForm(ASize int32) int32 {
	return XButton_Scale96ToForm(x.instance, ASize)
}

func (x *TXButton) ScaleFormTo96(ASize int32) int32 {
	return XButton_ScaleFormTo96(x.instance, ASize)
}

func (x *TXButton) Scale96ToFont(ASize int32) int32 {
	return XButton_Scale96ToFont(x.instance, ASize)
}

func (x *TXButton) ScaleFontTo96(ASize int32) int32 {
	return XButton_ScaleFontTo96(x.instance, ASize)
}

func (x *TXButton) ScaleScreenToFont(ASize int32) int32 {
	return XButton_ScaleScreenToFont(x.instance, ASize)
}

func (x *TXButton) ScaleFontToScreen(ASize int32) int32 {
	return XButton_ScaleFontToScreen(x.instance, ASize)
}

func (x *TXButton) Scale96ToScreen(ASize int32) int32 {
	return XButton_Scale96ToScreen(x.instance, ASize)
}

func (x *TXButton) ScaleScreenTo96(ASize int32) int32 {
	return XButton_ScaleScreenTo96(x.instance, ASize)
}

func (x *TXButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	XButton_AutoAdjustLayout(x.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (x *TXButton) FixDesignFontsPPI(ADesignTimePPI int32) {
	XButton_FixDesignFontsPPI(x.instance, ADesignTimePPI)
}

func (x *TXButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	XButton_ScaleFontsPPI(x.instance, AToPPI, AProportion)
}

// 获取控件标题。
//
// Get the control title.
func (x *TXButton) Caption() string {
	return XButton_GetCaption(x.instance)
}

// 设置控件标题。
//
// Set the control title.
func (x *TXButton) SetCaption(value string) {
	XButton_SetCaption(x.instance, value)
}

func (x *TXButton) ShowCaption() bool {
	return XButton_GetShowCaption(x.instance)
}

func (x *TXButton) SetShowCaption(value bool) {
	XButton_SetShowCaption(x.instance, value)
}

func (x *TXButton) BackColor() TColor {
	return XButton_GetBackColor(x.instance)
}

func (x *TXButton) SetBackColor(value TColor) {
	XButton_SetBackColor(x.instance, value)
}

func (x *TXButton) HoverColor() TColor {
	return XButton_GetHoverColor(x.instance)
}

func (x *TXButton) SetHoverColor(value TColor) {
	XButton_SetHoverColor(x.instance, value)
}

func (x *TXButton) DownColor() TColor {
	return XButton_GetDownColor(x.instance)
}

func (x *TXButton) SetDownColor(value TColor) {
	XButton_SetDownColor(x.instance, value)
}

// 获取边框的宽度。
func (x *TXButton) BorderWidth() int32 {
	return XButton_GetBorderWidth(x.instance)
}

// 设置边框的宽度。
func (x *TXButton) SetBorderWidth(value int32) {
	XButton_SetBorderWidth(x.instance, value)
}

func (x *TXButton) BorderColor() TColor {
	return XButton_GetBorderColor(x.instance)
}

func (x *TXButton) SetBorderColor(value TColor) {
	XButton_SetBorderColor(x.instance, value)
}

// 获取图片。
func (x *TXButton) Picture() *TPicture {
	return AsPicture(XButton_GetPicture(x.instance))
}

// 设置图片。
func (x *TXButton) SetPicture(value *TPicture) {
	XButton_SetPicture(x.instance, CheckPtr(value))
}

func (x *TXButton) DrawMode() TDrawImageMode {
	return XButton_GetDrawMode(x.instance)
}

func (x *TXButton) SetDrawMode(value TDrawImageMode) {
	XButton_SetDrawMode(x.instance, value)
}

func (x *TXButton) NormalFontColor() TColor {
	return XButton_GetNormalFontColor(x.instance)
}

func (x *TXButton) SetNormalFontColor(value TColor) {
	XButton_SetNormalFontColor(x.instance, value)
}

func (x *TXButton) DownFontColor() TColor {
	return XButton_GetDownFontColor(x.instance)
}

func (x *TXButton) SetDownFontColor(value TColor) {
	XButton_SetDownFontColor(x.instance, value)
}

func (x *TXButton) HoverFontColor() TColor {
	return XButton_GetHoverFontColor(x.instance)
}

func (x *TXButton) SetHoverFontColor(value TColor) {
	XButton_SetHoverFontColor(x.instance, value)
}

func (x *TXButton) Action() *TAction {
	return AsAction(XButton_GetAction(x.instance))
}

func (x *TXButton) SetAction(value IComponent) {
	XButton_SetAction(x.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (x *TXButton) Align() TAlign {
	return XButton_GetAlign(x.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (x *TXButton) SetAlign(value TAlign) {
	XButton_SetAlign(x.instance, value)
}

// 获取四个角位置的锚点。
func (x *TXButton) Anchors() TAnchors {
	return XButton_GetAnchors(x.instance)
}

// 设置四个角位置的锚点。
func (x *TXButton) SetAnchors(value TAnchors) {
	XButton_SetAnchors(x.instance, value)
}

func (x *TXButton) BiDiMode() TBiDiMode {
	return XButton_GetBiDiMode(x.instance)
}

func (x *TXButton) SetBiDiMode(value TBiDiMode) {
	XButton_SetBiDiMode(x.instance, value)
}

// 获取约束控件大小。
func (x *TXButton) Constraints() *TSizeConstraints {
	return AsSizeConstraints(XButton_GetConstraints(x.instance))
}

// 设置约束控件大小。
func (x *TXButton) SetConstraints(value *TSizeConstraints) {
	XButton_SetConstraints(x.instance, CheckPtr(value))
}

// 获取控件启用。
//
// Get the control enabled.
func (x *TXButton) Enabled() bool {
	return XButton_GetEnabled(x.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (x *TXButton) SetEnabled(value bool) {
	XButton_SetEnabled(x.instance, value)
}

// 获取字体。
//
// Get Font.
func (x *TXButton) Font() *TFont {
	return AsFont(XButton_GetFont(x.instance))
}

// 设置字体。
//
// Set Font.
func (x *TXButton) SetFont(value *TFont) {
	XButton_SetFont(x.instance, CheckPtr(value))
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (x *TXButton) ParentFont() bool {
	return XButton_GetParentFont(x.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (x *TXButton) SetParentFont(value bool) {
	XButton_SetParentFont(x.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (x *TXButton) ParentShowHint() bool {
	return XButton_GetParentShowHint(x.instance)
}

// 设置以父容器的ShowHint属性为准。
func (x *TXButton) SetParentShowHint(value bool) {
	XButton_SetParentShowHint(x.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (x *TXButton) PopupMenu() *TPopupMenu {
	return AsPopupMenu(XButton_GetPopupMenu(x.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (x *TXButton) SetPopupMenu(value IComponent) {
	XButton_SetPopupMenu(x.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (x *TXButton) ShowHint() bool {
	return XButton_GetShowHint(x.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (x *TXButton) SetShowHint(value bool) {
	XButton_SetShowHint(x.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (x *TXButton) Visible() bool {
	return XButton_GetVisible(x.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (x *TXButton) SetVisible(value bool) {
	XButton_SetVisible(x.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (x *TXButton) SetOnClick(fn TNotifyEvent) {
	XButton_SetOnClick(x.instance, fn)
}

// 设置双击事件。
func (x *TXButton) SetOnDblClick(fn TNotifyEvent) {
	XButton_SetOnDblClick(x.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (x *TXButton) SetOnMouseDown(fn TMouseEvent) {
	XButton_SetOnMouseDown(x.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (x *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	XButton_SetOnMouseEnter(x.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (x *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	XButton_SetOnMouseLeave(x.instance, fn)
}

// 设置鼠标移动事件。
func (x *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	XButton_SetOnMouseMove(x.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (x *TXButton) SetOnMouseUp(fn TMouseEvent) {
	XButton_SetOnMouseUp(x.instance, fn)
}

func (x *TXButton) BoundsRect() TRect {
	return XButton_GetBoundsRect(x.instance)
}

func (x *TXButton) SetBoundsRect(value TRect) {
	XButton_SetBoundsRect(x.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (x *TXButton) ClientHeight() int32 {
	return XButton_GetClientHeight(x.instance)
}

// 设置客户区高度。
//
// Set client height.
func (x *TXButton) SetClientHeight(value int32) {
	XButton_SetClientHeight(x.instance, value)
}

func (x *TXButton) ClientOrigin() TPoint {
	return XButton_GetClientOrigin(x.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (x *TXButton) ClientRect() TRect {
	return XButton_GetClientRect(x.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (x *TXButton) ClientWidth() int32 {
	return XButton_GetClientWidth(x.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (x *TXButton) SetClientWidth(value int32) {
	XButton_SetClientWidth(x.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (x *TXButton) ControlState() TControlState {
	return XButton_GetControlState(x.instance)
}

// 设置控件状态。
//
// Set control state.
func (x *TXButton) SetControlState(value TControlState) {
	XButton_SetControlState(x.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (x *TXButton) ControlStyle() TControlStyle {
	return XButton_GetControlStyle(x.instance)
}

// 设置控件样式。
//
// Set control style.
func (x *TXButton) SetControlStyle(value TControlStyle) {
	XButton_SetControlStyle(x.instance, value)
}

func (x *TXButton) Floating() bool {
	return XButton_GetFloating(x.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (x *TXButton) Parent() *TWinControl {
	return AsWinControl(XButton_GetParent(x.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (x *TXButton) SetParent(value IWinControl) {
	XButton_SetParent(x.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (x *TXButton) Left() int32 {
	return XButton_GetLeft(x.instance)
}

// 设置左边位置。
//
// Set Left position.
func (x *TXButton) SetLeft(value int32) {
	XButton_SetLeft(x.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (x *TXButton) Top() int32 {
	return XButton_GetTop(x.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (x *TXButton) SetTop(value int32) {
	XButton_SetTop(x.instance, value)
}

// 获取宽度。
//
// Get width.
func (x *TXButton) Width() int32 {
	return XButton_GetWidth(x.instance)
}

// 设置宽度。
//
// Set width.
func (x *TXButton) SetWidth(value int32) {
	XButton_SetWidth(x.instance, value)
}

// 获取高度。
//
// Get height.
func (x *TXButton) Height() int32 {
	return XButton_GetHeight(x.instance)
}

// 设置高度。
//
// Set height.
func (x *TXButton) SetHeight(value int32) {
	XButton_SetHeight(x.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (x *TXButton) Cursor() TCursor {
	return XButton_GetCursor(x.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (x *TXButton) SetCursor(value TCursor) {
	XButton_SetCursor(x.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (x *TXButton) Hint() string {
	return XButton_GetHint(x.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (x *TXButton) SetHint(value string) {
	XButton_SetHint(x.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (x *TXButton) ComponentCount() int32 {
	return XButton_GetComponentCount(x.instance)
}

// 获取组件索引。
//
// Get component index.
func (x *TXButton) ComponentIndex() int32 {
	return XButton_GetComponentIndex(x.instance)
}

// 设置组件索引。
//
// Set component index.
func (x *TXButton) SetComponentIndex(value int32) {
	XButton_SetComponentIndex(x.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (x *TXButton) Owner() *TComponent {
	return AsComponent(XButton_GetOwner(x.instance))
}

// 获取组件名称。
//
// Get the component name.
func (x *TXButton) Name() string {
	return XButton_GetName(x.instance)
}

// 设置组件名称。
//
// Set the component name.
func (x *TXButton) SetName(value string) {
	XButton_SetName(x.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (x *TXButton) Tag() int {
	return XButton_GetTag(x.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (x *TXButton) SetTag(value int) {
	XButton_SetTag(x.instance, value)
}

// 获取左边锚点。
func (x *TXButton) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(XButton_GetAnchorSideLeft(x.instance))
}

// 设置左边锚点。
func (x *TXButton) SetAnchorSideLeft(value *TAnchorSide) {
	XButton_SetAnchorSideLeft(x.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (x *TXButton) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(XButton_GetAnchorSideTop(x.instance))
}

// 设置顶边锚点。
func (x *TXButton) SetAnchorSideTop(value *TAnchorSide) {
	XButton_SetAnchorSideTop(x.instance, CheckPtr(value))
}

// 获取右边锚点。
func (x *TXButton) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(XButton_GetAnchorSideRight(x.instance))
}

// 设置右边锚点。
func (x *TXButton) SetAnchorSideRight(value *TAnchorSide) {
	XButton_SetAnchorSideRight(x.instance, CheckPtr(value))
}

// 获取底边锚点。
func (x *TXButton) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(XButton_GetAnchorSideBottom(x.instance))
}

// 设置底边锚点。
func (x *TXButton) SetAnchorSideBottom(value *TAnchorSide) {
	XButton_SetAnchorSideBottom(x.instance, CheckPtr(value))
}

// 获取边框间距。
func (x *TXButton) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(XButton_GetBorderSpacing(x.instance))
}

// 设置边框间距。
func (x *TXButton) SetBorderSpacing(value *TControlBorderSpacing) {
	XButton_SetBorderSpacing(x.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (x *TXButton) Components(AIndex int32) *TComponent {
	return AsComponent(XButton_GetComponents(x.instance, AIndex))
}

// 获取锚侧面。
func (x *TXButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(XButton_GetAnchorSide(x.instance, AKind))
}
