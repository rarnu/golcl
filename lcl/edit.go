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

type TEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(CheckPtr(owner))
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsEdit(obj interface{}) *TEdit {
    instance, ptr := getInstance(obj)
    if instance == 0 {
        return nil
    }
    return &TEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsEdit.
func EditFromInst(inst uintptr) *TEdit {
    return AsEdit(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsEdit.
func EditFromObj(obj IObject) *TEdit {
    return AsEdit(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsEdit.
func EditFromUnsafePointer(ptr unsafe.Pointer) *TEdit {
    return AsEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
        e.instance, e.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (e *TEdit) Instance() uintptr {
    return e.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (e *TEdit) UnsafeAddr() unsafe.Pointer {
    return e.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (e *TEdit) Is() TIs {
    return TIs(e.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (e *TEdit) As() TAs {
//    return TAs(e.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TEditClass() TClass {
    return Edit_StaticClassType()
}

// 清除。
func (e *TEdit) Clear() {
    Edit_Clear(e.instance)
}

// 清除选择。
func (e *TEdit) ClearSelection() {
    Edit_ClearSelection(e.instance)
}

// 复制到粘贴板。
func (e *TEdit) CopyToClipboard() {
    Edit_CopyToClipboard(e.instance)
}

// 剪切到粘贴板。
func (e *TEdit) CutToClipboard() {
    Edit_CutToClipboard(e.instance)
}

// 从剪切板粘贴。
func (e *TEdit) PasteFromClipboard() {
    Edit_PasteFromClipboard(e.instance)
}

// 撤销上一次操作。
func (e *TEdit) Undo() {
    Edit_Undo(e.instance)
}

// 全选。
func (e *TEdit) SelectAll() {
    Edit_SelectAll(e.instance)
}

// 是否可以获得焦点。
func (e *TEdit) CanFocus() bool {
    return Edit_CanFocus(e.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (e *TEdit) ContainsControl(Control IControl) bool {
    return Edit_ContainsControl(e.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (e *TEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Edit_ControlAtPos(e.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (e *TEdit) DisableAlign() {
    Edit_DisableAlign(e.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (e *TEdit) EnableAlign() {
    Edit_EnableAlign(e.instance)
}

// 查找子控件。
//
// Find sub controls.
func (e *TEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(Edit_FindChildControl(e.instance, ControlName))
}

func (e *TEdit) FlipChildren(AllLevels bool) {
    Edit_FlipChildren(e.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (e *TEdit) Focused() bool {
    return Edit_Focused(e.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (e *TEdit) HandleAllocated() bool {
    return Edit_HandleAllocated(e.instance)
}

// 插入一个控件。
//
// Insert a control.
func (e *TEdit) InsertControl(AControl IControl) {
    Edit_InsertControl(e.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (e *TEdit) Invalidate() {
    Edit_Invalidate(e.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (e *TEdit) PaintTo(DC HDC, X int32, Y int32) {
    Edit_PaintTo(e.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (e *TEdit) RemoveControl(AControl IControl) {
    Edit_RemoveControl(e.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (e *TEdit) Realign() {
    Edit_Realign(e.instance)
}

// 重绘。
//
// Repaint.
func (e *TEdit) Repaint() {
    Edit_Repaint(e.instance)
}

// 按比例缩放。
//
// Scale by.
func (e *TEdit) ScaleBy(M int32, D int32) {
    Edit_ScaleBy(e.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (e *TEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    Edit_ScrollBy(e.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Edit_SetBounds(e.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (e *TEdit) SetFocus() {
    Edit_SetFocus(e.instance)
}

// 控件更新。
//
// Update.
func (e *TEdit) Update() {
    Edit_Update(e.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (e *TEdit) BringToFront() {
    Edit_BringToFront(e.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (e *TEdit) ClientToScreen(Point TPoint) TPoint {
    return Edit_ClientToScreen(e.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (e *TEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ClientToParent(e.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (e *TEdit) Dragging() bool {
    return Edit_Dragging(e.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (e *TEdit) HasParent() bool {
    return Edit_HasParent(e.instance)
}

// 隐藏控件。
//
// Hidden control.
func (e *TEdit) Hide() {
    Edit_Hide(e.instance)
}

// 发送一个消息。
//
// Send a message.
func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Edit_Perform(e.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (e *TEdit) Refresh() {
    Edit_Refresh(e.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (e *TEdit) ScreenToClient(Point TPoint) TPoint {
    return Edit_ScreenToClient(e.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (e *TEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ParentToClient(e.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (e *TEdit) SendToBack() {
    Edit_SendToBack(e.instance)
}

// 显示控件。
//
// Show control.
func (e *TEdit) Show() {
    Edit_Show(e.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (e *TEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Edit_GetTextBuf(e.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (e *TEdit) GetTextLen() int32 {
    return Edit_GetTextLen(e.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (e *TEdit) SetTextBuf(Buffer string) {
    Edit_SetTextBuf(e.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (e *TEdit) FindComponent(AName string) *TComponent {
    return AsComponent(Edit_FindComponent(e.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (e *TEdit) GetNamePath() string {
    return Edit_GetNamePath(e.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (e *TEdit) Assign(Source IObject) {
    Edit_Assign(e.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (e *TEdit) ClassType() TClass {
    return Edit_ClassType(e.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (e *TEdit) ClassName() string {
    return Edit_ClassName(e.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (e *TEdit) InstanceSize() int32 {
    return Edit_InstanceSize(e.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (e *TEdit) InheritsFrom(AClass TClass) bool {
    return Edit_InheritsFrom(e.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (e *TEdit) Equals(Obj IObject) bool {
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (e *TEdit) GetHashCode() int32 {
    return Edit_GetHashCode(e.instance)
}

// 文本类信息。
//
// Text information.
func (e *TEdit) ToString() string {
    return Edit_ToString(e.instance)
}

func (e *TEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Edit_AnchorToNeighbour(e.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (e *TEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Edit_AnchorParallel(e.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (e *TEdit) AnchorHorizontalCenterTo(ASibling IControl) {
    Edit_AnchorHorizontalCenterTo(e.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (e *TEdit) AnchorVerticalCenterTo(ASibling IControl) {
    Edit_AnchorVerticalCenterTo(e.instance, CheckPtr(ASibling))
}

func (e *TEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Edit_AnchorSame(e.instance, ASide, CheckPtr(ASibling))
}

func (e *TEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Edit_AnchorAsAlign(e.instance, ATheAlign, ASpace)
}

func (e *TEdit) AnchorClient(ASpace int32) {
    Edit_AnchorClient(e.instance, ASpace)
}

func (e *TEdit) ScaleDesignToForm(ASize int32) int32 {
    return Edit_ScaleDesignToForm(e.instance, ASize)
}

func (e *TEdit) ScaleFormToDesign(ASize int32) int32 {
    return Edit_ScaleFormToDesign(e.instance, ASize)
}

func (e *TEdit) Scale96ToForm(ASize int32) int32 {
    return Edit_Scale96ToForm(e.instance, ASize)
}

func (e *TEdit) ScaleFormTo96(ASize int32) int32 {
    return Edit_ScaleFormTo96(e.instance, ASize)
}

func (e *TEdit) Scale96ToFont(ASize int32) int32 {
    return Edit_Scale96ToFont(e.instance, ASize)
}

func (e *TEdit) ScaleFontTo96(ASize int32) int32 {
    return Edit_ScaleFontTo96(e.instance, ASize)
}

func (e *TEdit) ScaleScreenToFont(ASize int32) int32 {
    return Edit_ScaleScreenToFont(e.instance, ASize)
}

func (e *TEdit) ScaleFontToScreen(ASize int32) int32 {
    return Edit_ScaleFontToScreen(e.instance, ASize)
}

func (e *TEdit) Scale96ToScreen(ASize int32) int32 {
    return Edit_Scale96ToScreen(e.instance, ASize)
}

func (e *TEdit) ScaleScreenTo96(ASize int32) int32 {
    return Edit_ScaleScreenTo96(e.instance, ASize)
}

func (e *TEdit) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Edit_AutoAdjustLayout(e.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (e *TEdit) FixDesignFontsPPI(ADesignTimePPI int32) {
    Edit_FixDesignFontsPPI(e.instance, ADesignTimePPI)
}

func (e *TEdit) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Edit_ScaleFontsPPI(e.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (e *TEdit) Align() TAlign {
    return Edit_GetAlign(e.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (e *TEdit) SetAlign(value TAlign) {
    Edit_SetAlign(e.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (e *TEdit) Alignment() TAlignment {
    return Edit_GetAlignment(e.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (e *TEdit) SetAlignment(value TAlignment) {
    Edit_SetAlignment(e.instance, value)
}

// 获取四个角位置的锚点。
func (e *TEdit) Anchors() TAnchors {
    return Edit_GetAnchors(e.instance)
}

// 设置四个角位置的锚点。
func (e *TEdit) SetAnchors(value TAnchors) {
    Edit_SetAnchors(e.instance, value)
}

// 获取自动选择。
func (e *TEdit) AutoSelect() bool {
    return Edit_GetAutoSelect(e.instance)
}

// 设置自动选择。
func (e *TEdit) SetAutoSelect(value bool) {
    Edit_SetAutoSelect(e.instance, value)
}

// 获取自动调整大小。
func (e *TEdit) AutoSize() bool {
    return Edit_GetAutoSize(e.instance)
}

// 设置自动调整大小。
func (e *TEdit) SetAutoSize(value bool) {
    Edit_SetAutoSize(e.instance, value)
}

func (e *TEdit) BiDiMode() TBiDiMode {
    return Edit_GetBiDiMode(e.instance)
}

func (e *TEdit) SetBiDiMode(value TBiDiMode) {
    Edit_SetBiDiMode(e.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (e *TEdit) BorderStyle() TBorderStyle {
    return Edit_GetBorderStyle(e.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    Edit_SetBorderStyle(e.instance, value)
}

func (e *TEdit) CharCase() TEditCharCase {
    return Edit_GetCharCase(e.instance)
}

func (e *TEdit) SetCharCase(value TEditCharCase) {
    Edit_SetCharCase(e.instance, value)
}

// 获取颜色。
//
// Get color.
func (e *TEdit) Color() TColor {
    return Edit_GetColor(e.instance)
}

// 设置颜色。
//
// Set color.
func (e *TEdit) SetColor(value TColor) {
    Edit_SetColor(e.instance, value)
}

// 获取约束控件大小。
func (e *TEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Edit_GetConstraints(e.instance))
}

// 设置约束控件大小。
func (e *TEdit) SetConstraints(value *TSizeConstraints) {
    Edit_SetConstraints(e.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (e *TEdit) DoubleBuffered() bool {
    return Edit_GetDoubleBuffered(e.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (e *TEdit) SetDoubleBuffered(value bool) {
    Edit_SetDoubleBuffered(e.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (e *TEdit) DragCursor() TCursor {
    return Edit_GetDragCursor(e.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (e *TEdit) SetDragCursor(value TCursor) {
    Edit_SetDragCursor(e.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (e *TEdit) DragKind() TDragKind {
    return Edit_GetDragKind(e.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (e *TEdit) SetDragKind(value TDragKind) {
    Edit_SetDragKind(e.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (e *TEdit) DragMode() TDragMode {
    return Edit_GetDragMode(e.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (e *TEdit) SetDragMode(value TDragMode) {
    Edit_SetDragMode(e.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (e *TEdit) Enabled() bool {
    return Edit_GetEnabled(e.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (e *TEdit) SetEnabled(value bool) {
    Edit_SetEnabled(e.instance, value)
}

// 获取字体。
//
// Get Font.
func (e *TEdit) Font() *TFont {
    return AsFont(Edit_GetFont(e.instance))
}

// 设置字体。
//
// Set Font.
func (e *TEdit) SetFont(value *TFont) {
    Edit_SetFont(e.instance, CheckPtr(value))
}

// 获取隐藏选择。
func (e *TEdit) HideSelection() bool {
    return Edit_GetHideSelection(e.instance)
}

// 设置隐藏选择。
func (e *TEdit) SetHideSelection(value bool) {
    Edit_SetHideSelection(e.instance, value)
}

// 获取最大长度。
func (e *TEdit) MaxLength() int32 {
    return Edit_GetMaxLength(e.instance)
}

// 设置最大长度。
func (e *TEdit) SetMaxLength(value int32) {
    Edit_SetMaxLength(e.instance, value)
}

// 获取只能输入数字。
func (e *TEdit) NumbersOnly() bool {
    return Edit_GetNumbersOnly(e.instance)
}

// 设置只能输入数字。
func (e *TEdit) SetNumbersOnly(value bool) {
    Edit_SetNumbersOnly(e.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (e *TEdit) ParentColor() bool {
    return Edit_GetParentColor(e.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (e *TEdit) SetParentColor(value bool) {
    Edit_SetParentColor(e.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (e *TEdit) ParentDoubleBuffered() bool {
    return Edit_GetParentDoubleBuffered(e.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (e *TEdit) SetParentDoubleBuffered(value bool) {
    Edit_SetParentDoubleBuffered(e.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (e *TEdit) ParentFont() bool {
    return Edit_GetParentFont(e.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (e *TEdit) SetParentFont(value bool) {
    Edit_SetParentFont(e.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (e *TEdit) ParentShowHint() bool {
    return Edit_GetParentShowHint(e.instance)
}

// 设置以父容器的ShowHint属性为准。
func (e *TEdit) SetParentShowHint(value bool) {
    Edit_SetParentShowHint(e.instance, value)
}

// 获取密码掩码字符。
func (e *TEdit) PasswordChar() uint16 {
    return Edit_GetPasswordChar(e.instance)
}

// 设置密码掩码字符。
func (e *TEdit) SetPasswordChar(value uint16) {
    Edit_SetPasswordChar(e.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (e *TEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Edit_GetPopupMenu(e.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (e *TEdit) SetPopupMenu(value IComponent) {
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

// 获取只读。
func (e *TEdit) ReadOnly() bool {
    return Edit_GetReadOnly(e.instance)
}

// 设置只读。
func (e *TEdit) SetReadOnly(value bool) {
    Edit_SetReadOnly(e.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (e *TEdit) ShowHint() bool {
    return Edit_GetShowHint(e.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (e *TEdit) SetShowHint(value bool) {
    Edit_SetShowHint(e.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (e *TEdit) TabOrder() TTabOrder {
    return Edit_GetTabOrder(e.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (e *TEdit) SetTabOrder(value TTabOrder) {
    Edit_SetTabOrder(e.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (e *TEdit) TabStop() bool {
    return Edit_GetTabStop(e.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (e *TEdit) SetTabStop(value bool) {
    Edit_SetTabStop(e.instance, value)
}

// 获取文本。
func (e *TEdit) Text() string {
    strLen := e.GetTextLen()
    if strLen != 0 {
        var buffStr string
        e.GetTextBuf(&buffStr, strLen+1)
        return buffStr
    }
    return ""
}

// 设置文本。
func (e *TEdit) SetText(value string) {
    Edit_SetText(e.instance, value)
}

// 获取提示文本。
func (e *TEdit) TextHint() string {
    return Edit_GetTextHint(e.instance)
}

// 设置提示文本。
func (e *TEdit) SetTextHint(value string) {
    Edit_SetTextHint(e.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (e *TEdit) Visible() bool {
    return Edit_GetVisible(e.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (e *TEdit) SetVisible(value bool) {
    Edit_SetVisible(e.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (e *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
    Edit_SetOnContextPopup(e.instance, fn)
}

// 设置双击事件。
func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (e *TEdit) SetOnDragDrop(fn TDragDropEvent) {
    Edit_SetOnDragDrop(e.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (e *TEdit) SetOnDragOver(fn TDragOverEvent) {
    Edit_SetOnDragOver(e.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (e *TEdit) SetOnEndDrag(fn TEndDragEvent) {
    Edit_SetOnEndDrag(e.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

// 设置键键下事件。
func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

// 设置鼠标移动事件。
func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

// 获取能否撤销。
func (e *TEdit) CanUndo() bool {
    return Edit_GetCanUndo(e.instance)
}

// 获取修改。
//
// Get modified.
func (e *TEdit) Modified() bool {
    return Edit_GetModified(e.instance)
}

// 设置修改。
//
// Set modified.
func (e *TEdit) SetModified(value bool) {
    Edit_SetModified(e.instance, value)
}

// 获取选择的长度。
func (e *TEdit) SelLength() int32 {
    return Edit_GetSelLength(e.instance)
}

// 设置选择的长度。
func (e *TEdit) SetSelLength(value int32) {
    Edit_SetSelLength(e.instance, value)
}

// 获取选择的启始位置。
func (e *TEdit) SelStart() int32 {
    return Edit_GetSelStart(e.instance)
}

// 设置选择的启始位置。
func (e *TEdit) SetSelStart(value int32) {
    Edit_SetSelStart(e.instance, value)
}

// 获取选择的文本。
func (e *TEdit) SelText() string {
    return Edit_GetSelText(e.instance)
}

// 设置选择的文本。
func (e *TEdit) SetSelText(value string) {
    Edit_SetSelText(e.instance, value)
}

// 获取依靠客户端总数。
func (e *TEdit) DockClientCount() int32 {
    return Edit_GetDockClientCount(e.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (e *TEdit) DockSite() bool {
    return Edit_GetDockSite(e.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (e *TEdit) SetDockSite(value bool) {
    Edit_SetDockSite(e.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (e *TEdit) MouseInClient() bool {
    return Edit_GetMouseInClient(e.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (e *TEdit) VisibleDockClientCount() int32 {
    return Edit_GetVisibleDockClientCount(e.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (e *TEdit) Brush() *TBrush {
    return AsBrush(Edit_GetBrush(e.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (e *TEdit) ControlCount() int32 {
    return Edit_GetControlCount(e.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (e *TEdit) Handle() HWND {
    return Edit_GetHandle(e.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (e *TEdit) ParentWindow() HWND {
    return Edit_GetParentWindow(e.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (e *TEdit) SetParentWindow(value HWND) {
    Edit_SetParentWindow(e.instance, value)
}

func (e *TEdit) Showing() bool {
    return Edit_GetShowing(e.instance)
}

// 获取使用停靠管理。
func (e *TEdit) UseDockManager() bool {
    return Edit_GetUseDockManager(e.instance)
}

// 设置使用停靠管理。
func (e *TEdit) SetUseDockManager(value bool) {
    Edit_SetUseDockManager(e.instance, value)
}

func (e *TEdit) Action() *TAction {
    return AsAction(Edit_GetAction(e.instance))
}

func (e *TEdit) SetAction(value IComponent) {
    Edit_SetAction(e.instance, CheckPtr(value))
}

func (e *TEdit) BoundsRect() TRect {
    return Edit_GetBoundsRect(e.instance)
}

func (e *TEdit) SetBoundsRect(value TRect) {
    Edit_SetBoundsRect(e.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (e *TEdit) ClientHeight() int32 {
    return Edit_GetClientHeight(e.instance)
}

// 设置客户区高度。
//
// Set client height.
func (e *TEdit) SetClientHeight(value int32) {
    Edit_SetClientHeight(e.instance, value)
}

func (e *TEdit) ClientOrigin() TPoint {
    return Edit_GetClientOrigin(e.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (e *TEdit) ClientRect() TRect {
    return Edit_GetClientRect(e.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (e *TEdit) ClientWidth() int32 {
    return Edit_GetClientWidth(e.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (e *TEdit) SetClientWidth(value int32) {
    Edit_SetClientWidth(e.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (e *TEdit) ControlState() TControlState {
    return Edit_GetControlState(e.instance)
}

// 设置控件状态。
//
// Set control state.
func (e *TEdit) SetControlState(value TControlState) {
    Edit_SetControlState(e.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (e *TEdit) ControlStyle() TControlStyle {
    return Edit_GetControlStyle(e.instance)
}

// 设置控件样式。
//
// Set control style.
func (e *TEdit) SetControlStyle(value TControlStyle) {
    Edit_SetControlStyle(e.instance, value)
}

func (e *TEdit) Floating() bool {
    return Edit_GetFloating(e.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (e *TEdit) Parent() *TWinControl {
    return AsWinControl(Edit_GetParent(e.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (e *TEdit) SetParent(value IWinControl) {
    Edit_SetParent(e.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (e *TEdit) Left() int32 {
    return Edit_GetLeft(e.instance)
}

// 设置左边位置。
//
// Set Left position.
func (e *TEdit) SetLeft(value int32) {
    Edit_SetLeft(e.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (e *TEdit) Top() int32 {
    return Edit_GetTop(e.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (e *TEdit) SetTop(value int32) {
    Edit_SetTop(e.instance, value)
}

// 获取宽度。
//
// Get width.
func (e *TEdit) Width() int32 {
    return Edit_GetWidth(e.instance)
}

// 设置宽度。
//
// Set width.
func (e *TEdit) SetWidth(value int32) {
    Edit_SetWidth(e.instance, value)
}

// 获取高度。
//
// Get height.
func (e *TEdit) Height() int32 {
    return Edit_GetHeight(e.instance)
}

// 设置高度。
//
// Set height.
func (e *TEdit) SetHeight(value int32) {
    Edit_SetHeight(e.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (e *TEdit) Cursor() TCursor {
    return Edit_GetCursor(e.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (e *TEdit) SetCursor(value TCursor) {
    Edit_SetCursor(e.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (e *TEdit) Hint() string {
    return Edit_GetHint(e.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (e *TEdit) SetHint(value string) {
    Edit_SetHint(e.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (e *TEdit) ComponentCount() int32 {
    return Edit_GetComponentCount(e.instance)
}

// 获取组件索引。
//
// Get component index.
func (e *TEdit) ComponentIndex() int32 {
    return Edit_GetComponentIndex(e.instance)
}

// 设置组件索引。
//
// Set component index.
func (e *TEdit) SetComponentIndex(value int32) {
    Edit_SetComponentIndex(e.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (e *TEdit) Owner() *TComponent {
    return AsComponent(Edit_GetOwner(e.instance))
}

// 获取组件名称。
//
// Get the component name.
func (e *TEdit) Name() string {
    return Edit_GetName(e.instance)
}

// 设置组件名称。
//
// Set the component name.
func (e *TEdit) SetName(value string) {
    Edit_SetName(e.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (e *TEdit) Tag() int {
    return Edit_GetTag(e.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (e *TEdit) SetTag(value int) {
    Edit_SetTag(e.instance, value)
}

// 获取左边锚点。
func (e *TEdit) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Edit_GetAnchorSideLeft(e.instance))
}

// 设置左边锚点。
func (e *TEdit) SetAnchorSideLeft(value *TAnchorSide) {
    Edit_SetAnchorSideLeft(e.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (e *TEdit) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Edit_GetAnchorSideTop(e.instance))
}

// 设置顶边锚点。
func (e *TEdit) SetAnchorSideTop(value *TAnchorSide) {
    Edit_SetAnchorSideTop(e.instance, CheckPtr(value))
}

// 获取右边锚点。
func (e *TEdit) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Edit_GetAnchorSideRight(e.instance))
}

// 设置右边锚点。
func (e *TEdit) SetAnchorSideRight(value *TAnchorSide) {
    Edit_SetAnchorSideRight(e.instance, CheckPtr(value))
}

// 获取底边锚点。
func (e *TEdit) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Edit_GetAnchorSideBottom(e.instance))
}

// 设置底边锚点。
func (e *TEdit) SetAnchorSideBottom(value *TAnchorSide) {
    Edit_SetAnchorSideBottom(e.instance, CheckPtr(value))
}

func (e *TEdit) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(Edit_GetChildSizing(e.instance))
}

func (e *TEdit) SetChildSizing(value *TControlChildSizing) {
    Edit_SetChildSizing(e.instance, CheckPtr(value))
}

// 获取边框间距。
func (e *TEdit) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Edit_GetBorderSpacing(e.instance))
}

// 设置边框间距。
func (e *TEdit) SetBorderSpacing(value *TControlBorderSpacing) {
    Edit_SetBorderSpacing(e.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (e *TEdit) DockClients(Index int32) *TControl {
    return AsControl(Edit_GetDockClients(e.instance, Index))
}

// 获取指定索引子控件。
func (e *TEdit) Controls(Index int32) *TControl {
    return AsControl(Edit_GetControls(e.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (e *TEdit) Components(AIndex int32) *TComponent {
    return AsComponent(Edit_GetComponents(e.instance, AIndex))
}

// 获取锚侧面。
func (e *TEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Edit_GetAnchorSide(e.instance, AKind))
}
