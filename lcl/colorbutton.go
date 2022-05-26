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

type TColorButton struct {
	IControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewColorButton(owner IComponent) *TColorButton {
	c := new(TColorButton)
	c.instance = ColorButton_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsColorButton(obj interface{}) *TColorButton {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TColorButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsColorButton.
func ColorButtonFromInst(inst uintptr) *TColorButton {
	return AsColorButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsColorButton.
func ColorButtonFromObj(obj IObject) *TColorButton {
	return AsColorButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsColorButton.
func ColorButtonFromUnsafePointer(ptr unsafe.Pointer) *TColorButton {
	return AsColorButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TColorButton) Free() {
	if c.instance != 0 {
		ColorButton_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TColorButton) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TColorButton) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TColorButton) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TColorButton) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TColorButton) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TColorButtonClass() TClass {
	return ColorButton_StaticClassType()
}

// 单击。
func (c *TColorButton) Click() {
	ColorButton_Click(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TColorButton) BringToFront() {
	ColorButton_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TColorButton) ClientToScreen(Point TPoint) TPoint {
	return ColorButton_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TColorButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ColorButton_ClientToParent(c.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TColorButton) Dragging() bool {
	return ColorButton_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TColorButton) HasParent() bool {
	return ColorButton_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TColorButton) Hide() {
	ColorButton_Hide(c.instance)
}

// 要求重绘。
//
// Redraw.
func (c *TColorButton) Invalidate() {
	ColorButton_Invalidate(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TColorButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ColorButton_Perform(c.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TColorButton) Refresh() {
	ColorButton_Refresh(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TColorButton) Repaint() {
	ColorButton_Repaint(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TColorButton) ScreenToClient(Point TPoint) TPoint {
	return ColorButton_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TColorButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ColorButton_ParentToClient(c.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TColorButton) SendToBack() {
	ColorButton_SendToBack(c.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TColorButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ColorButton_SetBounds(c.instance, ALeft, ATop, AWidth, AHeight)
}

// 显示控件。
//
// Show control.
func (c *TColorButton) Show() {
	ColorButton_Show(c.instance)
}

// 控件更新。
//
// Update.
func (c *TColorButton) Update() {
	ColorButton_Update(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TColorButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ColorButton_GetTextBuf(c.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TColorButton) GetTextLen() int32 {
	return ColorButton_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TColorButton) SetTextBuf(Buffer string) {
	ColorButton_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TColorButton) FindComponent(AName string) *TComponent {
	return AsComponent(ColorButton_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TColorButton) GetNamePath() string {
	return ColorButton_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TColorButton) Assign(Source IObject) {
	ColorButton_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TColorButton) ClassType() TClass {
	return ColorButton_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TColorButton) ClassName() string {
	return ColorButton_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TColorButton) InstanceSize() int32 {
	return ColorButton_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TColorButton) InheritsFrom(AClass TClass) bool {
	return ColorButton_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TColorButton) Equals(Obj IObject) bool {
	return ColorButton_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TColorButton) GetHashCode() int32 {
	return ColorButton_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TColorButton) ToString() string {
	return ColorButton_ToString(c.instance)
}

func (c *TColorButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ColorButton_AnchorToNeighbour(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (c *TColorButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ColorButton_AnchorParallel(c.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TColorButton) AnchorHorizontalCenterTo(ASibling IControl) {
	ColorButton_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TColorButton) AnchorVerticalCenterTo(ASibling IControl) {
	ColorButton_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TColorButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ColorButton_AnchorSame(c.instance, ASide, CheckPtr(ASibling))
}

func (c *TColorButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ColorButton_AnchorAsAlign(c.instance, ATheAlign, ASpace)
}

func (c *TColorButton) AnchorClient(ASpace int32) {
	ColorButton_AnchorClient(c.instance, ASpace)
}

func (c *TColorButton) ScaleDesignToForm(ASize int32) int32 {
	return ColorButton_ScaleDesignToForm(c.instance, ASize)
}

func (c *TColorButton) ScaleFormToDesign(ASize int32) int32 {
	return ColorButton_ScaleFormToDesign(c.instance, ASize)
}

func (c *TColorButton) Scale96ToForm(ASize int32) int32 {
	return ColorButton_Scale96ToForm(c.instance, ASize)
}

func (c *TColorButton) ScaleFormTo96(ASize int32) int32 {
	return ColorButton_ScaleFormTo96(c.instance, ASize)
}

func (c *TColorButton) Scale96ToFont(ASize int32) int32 {
	return ColorButton_Scale96ToFont(c.instance, ASize)
}

func (c *TColorButton) ScaleFontTo96(ASize int32) int32 {
	return ColorButton_ScaleFontTo96(c.instance, ASize)
}

func (c *TColorButton) ScaleScreenToFont(ASize int32) int32 {
	return ColorButton_ScaleScreenToFont(c.instance, ASize)
}

func (c *TColorButton) ScaleFontToScreen(ASize int32) int32 {
	return ColorButton_ScaleFontToScreen(c.instance, ASize)
}

func (c *TColorButton) Scale96ToScreen(ASize int32) int32 {
	return ColorButton_Scale96ToScreen(c.instance, ASize)
}

func (c *TColorButton) ScaleScreenTo96(ASize int32) int32 {
	return ColorButton_ScaleScreenTo96(c.instance, ASize)
}

func (c *TColorButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ColorButton_AutoAdjustLayout(c.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (c *TColorButton) FixDesignFontsPPI(ADesignTimePPI int32) {
	ColorButton_FixDesignFontsPPI(c.instance, ADesignTimePPI)
}

func (c *TColorButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ColorButton_ScaleFontsPPI(c.instance, AToPPI, AProportion)
}

func (c *TColorButton) Action() *TAction {
	return AsAction(ColorButton_GetAction(c.instance))
}

func (c *TColorButton) SetAction(value IComponent) {
	ColorButton_SetAction(c.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TColorButton) Align() TAlign {
	return ColorButton_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TColorButton) SetAlign(value TAlign) {
	ColorButton_SetAlign(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TColorButton) Anchors() TAnchors {
	return ColorButton_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TColorButton) SetAnchors(value TAnchors) {
	ColorButton_SetAnchors(c.instance, value)
}

func (c *TColorButton) AllowAllUp() bool {
	return ColorButton_GetAllowAllUp(c.instance)
}

func (c *TColorButton) SetAllowAllUp(value bool) {
	ColorButton_SetAllowAllUp(c.instance, value)
}

// 获取边框的宽度。
func (c *TColorButton) BorderWidth() int32 {
	return ColorButton_GetBorderWidth(c.instance)
}

// 设置边框的宽度。
func (c *TColorButton) SetBorderWidth(value int32) {
	ColorButton_SetBorderWidth(c.instance, value)
}

func (c *TColorButton) ButtonColorAutoSize() bool {
	return ColorButton_GetButtonColorAutoSize(c.instance)
}

func (c *TColorButton) SetButtonColorAutoSize(value bool) {
	ColorButton_SetButtonColorAutoSize(c.instance, value)
}

func (c *TColorButton) ButtonColorSize() int32 {
	return ColorButton_GetButtonColorSize(c.instance)
}

func (c *TColorButton) SetButtonColorSize(value int32) {
	ColorButton_SetButtonColorSize(c.instance, value)
}

func (c *TColorButton) ButtonColor() TColor {
	return ColorButton_GetButtonColor(c.instance)
}

func (c *TColorButton) SetButtonColor(value TColor) {
	ColorButton_SetButtonColor(c.instance, value)
}

func (c *TColorButton) ColorDialog() *TColorDialog {
	return AsColorDialog(ColorButton_GetColorDialog(c.instance))
}

func (c *TColorButton) SetColorDialog(value IComponent) {
	ColorButton_SetColorDialog(c.instance, CheckPtr(value))
}

// 获取约束控件大小。
func (c *TColorButton) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ColorButton_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TColorButton) SetConstraints(value *TSizeConstraints) {
	ColorButton_SetConstraints(c.instance, CheckPtr(value))
}

// 获取控件标题。
//
// Get the control title.
func (c *TColorButton) Caption() string {
	return ColorButton_GetCaption(c.instance)
}

// 设置控件标题。
//
// Set the control title.
func (c *TColorButton) SetCaption(value string) {
	ColorButton_SetCaption(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TColorButton) Color() TColor {
	return ColorButton_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TColorButton) SetColor(value TColor) {
	ColorButton_SetColor(c.instance, value)
}

func (c *TColorButton) Down() bool {
	return ColorButton_GetDown(c.instance)
}

func (c *TColorButton) SetDown(value bool) {
	ColorButton_SetDown(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TColorButton) Enabled() bool {
	return ColorButton_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TColorButton) SetEnabled(value bool) {
	ColorButton_SetEnabled(c.instance, value)
}

// 获取平面样式。
func (c *TColorButton) Flat() bool {
	return ColorButton_GetFlat(c.instance)
}

// 设置平面样式。
func (c *TColorButton) SetFlat(value bool) {
	ColorButton_SetFlat(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TColorButton) Font() *TFont {
	return AsFont(ColorButton_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TColorButton) SetFont(value *TFont) {
	ColorButton_SetFont(c.instance, CheckPtr(value))
}

// 获取团组索引。
func (c *TColorButton) GroupIndex() int32 {
	return ColorButton_GetGroupIndex(c.instance)
}

// 设置团组索引。
func (c *TColorButton) SetGroupIndex(value int32) {
	ColorButton_SetGroupIndex(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TColorButton) Hint() string {
	return ColorButton_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TColorButton) SetHint(value string) {
	ColorButton_SetHint(c.instance, value)
}

func (c *TColorButton) Layout() TButtonLayout {
	return ColorButton_GetLayout(c.instance)
}

func (c *TColorButton) SetLayout(value TButtonLayout) {
	ColorButton_SetLayout(c.instance, value)
}

func (c *TColorButton) Spacing() int32 {
	return ColorButton_GetSpacing(c.instance)
}

func (c *TColorButton) SetSpacing(value int32) {
	ColorButton_SetSpacing(c.instance, value)
}

// 获取透明。
//
// Get transparent.
func (c *TColorButton) Transparent() bool {
	return ColorButton_GetTransparent(c.instance)
}

// 设置透明。
//
// Set transparent.
func (c *TColorButton) SetTransparent(value bool) {
	ColorButton_SetTransparent(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TColorButton) Visible() bool {
	return ColorButton_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TColorButton) SetVisible(value bool) {
	ColorButton_SetVisible(c.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TColorButton) SetOnClick(fn TNotifyEvent) {
	ColorButton_SetOnClick(c.instance, fn)
}

func (c *TColorButton) SetOnColorChanged(fn TNotifyEvent) {
	ColorButton_SetOnColorChanged(c.instance, fn)
}

// 设置双击事件。
func (c *TColorButton) SetOnDblClick(fn TNotifyEvent) {
	ColorButton_SetOnDblClick(c.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TColorButton) SetOnMouseDown(fn TMouseEvent) {
	ColorButton_SetOnMouseDown(c.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TColorButton) SetOnMouseEnter(fn TNotifyEvent) {
	ColorButton_SetOnMouseEnter(c.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TColorButton) SetOnMouseLeave(fn TNotifyEvent) {
	ColorButton_SetOnMouseLeave(c.instance, fn)
}

// 设置鼠标移动事件。
func (c *TColorButton) SetOnMouseMove(fn TMouseMoveEvent) {
	ColorButton_SetOnMouseMove(c.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TColorButton) SetOnMouseUp(fn TMouseEvent) {
	ColorButton_SetOnMouseUp(c.instance, fn)
}

// 设置鼠标滚轮事件。
func (c *TColorButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	ColorButton_SetOnMouseWheel(c.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (c *TColorButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	ColorButton_SetOnMouseWheelDown(c.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (c *TColorButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	ColorButton_SetOnMouseWheelUp(c.instance, fn)
}

// 设置绘画事件。
func (c *TColorButton) SetOnPaint(fn TNotifyEvent) {
	ColorButton_SetOnPaint(c.instance, fn)
}

// 设置大小被改变事件。
func (c *TColorButton) SetOnResize(fn TNotifyEvent) {
	ColorButton_SetOnResize(c.instance, fn)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TColorButton) ShowHint() bool {
	return ColorButton_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TColorButton) SetShowHint(value bool) {
	ColorButton_SetShowHint(c.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TColorButton) ParentFont() bool {
	return ColorButton_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TColorButton) SetParentFont(value bool) {
	ColorButton_SetParentFont(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TColorButton) ParentShowHint() bool {
	return ColorButton_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TColorButton) SetParentShowHint(value bool) {
	ColorButton_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TColorButton) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ColorButton_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TColorButton) SetPopupMenu(value IComponent) {
	ColorButton_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取图像在images中的索引。
func (c *TColorButton) ImageIndex() int32 {
	return ColorButton_GetImageIndex(c.instance)
}

// 设置图像在images中的索引。
func (c *TColorButton) SetImageIndex(value int32) {
	ColorButton_SetImageIndex(c.instance, value)
}

// 获取图标索引列表对象。
func (c *TColorButton) Images() *TImageList {
	return AsImageList(ColorButton_GetImages(c.instance))
}

// 设置图标索引列表对象。
func (c *TColorButton) SetImages(value IComponent) {
	ColorButton_SetImages(c.instance, CheckPtr(value))
}

func (c *TColorButton) ImageWidth() int32 {
	return ColorButton_GetImageWidth(c.instance)
}

func (c *TColorButton) SetImageWidth(value int32) {
	ColorButton_SetImageWidth(c.instance, value)
}

func (c *TColorButton) ShowCaption() bool {
	return ColorButton_GetShowCaption(c.instance)
}

func (c *TColorButton) SetShowCaption(value bool) {
	ColorButton_SetShowCaption(c.instance, value)
}

func (c *TColorButton) BiDiMode() TBiDiMode {
	return ColorButton_GetBiDiMode(c.instance)
}

func (c *TColorButton) SetBiDiMode(value TBiDiMode) {
	ColorButton_SetBiDiMode(c.instance, value)
}

func (c *TColorButton) Glyph() *TBitmap {
	return AsBitmap(ColorButton_GetGlyph(c.instance))
}

func (c *TColorButton) SetGlyph(value *TBitmap) {
	ColorButton_SetGlyph(c.instance, CheckPtr(value))
}

func (c *TColorButton) NumGlyphs() TNumGlyphs {
	return ColorButton_GetNumGlyphs(c.instance)
}

func (c *TColorButton) SetNumGlyphs(value TNumGlyphs) {
	ColorButton_SetNumGlyphs(c.instance, value)
}

func (c *TColorButton) BoundsRect() TRect {
	return ColorButton_GetBoundsRect(c.instance)
}

func (c *TColorButton) SetBoundsRect(value TRect) {
	ColorButton_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TColorButton) ClientHeight() int32 {
	return ColorButton_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TColorButton) SetClientHeight(value int32) {
	ColorButton_SetClientHeight(c.instance, value)
}

func (c *TColorButton) ClientOrigin() TPoint {
	return ColorButton_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TColorButton) ClientRect() TRect {
	return ColorButton_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TColorButton) ClientWidth() int32 {
	return ColorButton_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TColorButton) SetClientWidth(value int32) {
	ColorButton_SetClientWidth(c.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (c *TColorButton) ControlState() TControlState {
	return ColorButton_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TColorButton) SetControlState(value TControlState) {
	ColorButton_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TColorButton) ControlStyle() TControlStyle {
	return ColorButton_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TColorButton) SetControlStyle(value TControlStyle) {
	ColorButton_SetControlStyle(c.instance, value)
}

func (c *TColorButton) Floating() bool {
	return ColorButton_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TColorButton) Parent() *TWinControl {
	return AsWinControl(ColorButton_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TColorButton) SetParent(value IWinControl) {
	ColorButton_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TColorButton) Left() int32 {
	return ColorButton_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TColorButton) SetLeft(value int32) {
	ColorButton_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TColorButton) Top() int32 {
	return ColorButton_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TColorButton) SetTop(value int32) {
	ColorButton_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TColorButton) Width() int32 {
	return ColorButton_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TColorButton) SetWidth(value int32) {
	ColorButton_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TColorButton) Height() int32 {
	return ColorButton_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TColorButton) SetHeight(value int32) {
	ColorButton_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TColorButton) Cursor() TCursor {
	return ColorButton_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TColorButton) SetCursor(value TCursor) {
	ColorButton_SetCursor(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TColorButton) ComponentCount() int32 {
	return ColorButton_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TColorButton) ComponentIndex() int32 {
	return ColorButton_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TColorButton) SetComponentIndex(value int32) {
	ColorButton_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TColorButton) Owner() *TComponent {
	return AsComponent(ColorButton_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TColorButton) Name() string {
	return ColorButton_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TColorButton) SetName(value string) {
	ColorButton_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TColorButton) Tag() int {
	return ColorButton_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TColorButton) SetTag(value int) {
	ColorButton_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TColorButton) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ColorButton_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TColorButton) SetAnchorSideLeft(value *TAnchorSide) {
	ColorButton_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TColorButton) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ColorButton_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TColorButton) SetAnchorSideTop(value *TAnchorSide) {
	ColorButton_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TColorButton) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ColorButton_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TColorButton) SetAnchorSideRight(value *TAnchorSide) {
	ColorButton_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TColorButton) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ColorButton_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TColorButton) SetAnchorSideBottom(value *TAnchorSide) {
	ColorButton_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TColorButton) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ColorButton_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TColorButton) SetBorderSpacing(value *TControlBorderSpacing) {
	ColorButton_SetBorderSpacing(c.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TColorButton) Components(AIndex int32) *TComponent {
	return AsComponent(ColorButton_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TColorButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ColorButton_GetAnchorSide(c.instance, AKind))
}
