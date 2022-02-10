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

type TValueListEditor struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewValueListEditor(owner IComponent) *TValueListEditor {
	v := new(TValueListEditor)
	v.instance = ValueListEditor_Create(CheckPtr(owner))
	v.ptr = unsafe.Pointer(v.instance)
	return v
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsValueListEditor(obj interface{}) *TValueListEditor {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TValueListEditor{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsValueListEditor.
func ValueListEditorFromInst(inst uintptr) *TValueListEditor {
	return AsValueListEditor(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsValueListEditor.
func ValueListEditorFromObj(obj IObject) *TValueListEditor {
	return AsValueListEditor(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsValueListEditor.
func ValueListEditorFromUnsafePointer(ptr unsafe.Pointer) *TValueListEditor {
	return AsValueListEditor(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (v *TValueListEditor) Free() {
	if v.instance != 0 {
		ValueListEditor_Free(v.instance)
		v.instance, v.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (v *TValueListEditor) Instance() uintptr {
	return v.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (v *TValueListEditor) UnsafeAddr() unsafe.Pointer {
	return v.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (v *TValueListEditor) IsValid() bool {
	return v.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (v *TValueListEditor) Is() TIs {
	return TIs(v.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (v *TValueListEditor) As() TAs {
//    return TAs(v.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TValueListEditorClass() TClass {
	return ValueListEditor_StaticClassType()
}

func (v *TValueListEditor) DeleteRow(ARow int32) {
	ValueListEditor_DeleteRow(v.instance, ARow)
}

// 刷新控件。
//
// Refresh control.
func (v *TValueListEditor) Refresh() {
	ValueListEditor_Refresh(v.instance)
}

func (v *TValueListEditor) CellRect(ACol int32, ARow int32) TRect {
	return ValueListEditor_CellRect(v.instance, ACol, ARow)
}

func (v *TValueListEditor) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
	ValueListEditor_MouseToCell(v.instance, X, Y, ACol, ARow)
}

func (v *TValueListEditor) MouseCoord(X int32, Y int32) TGridCoord {
	return ValueListEditor_MouseCoord(v.instance, X, Y)
}

// 是否可以获得焦点。
func (v *TValueListEditor) CanFocus() bool {
	return ValueListEditor_CanFocus(v.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (v *TValueListEditor) ContainsControl(Control IControl) bool {
	return ValueListEditor_ContainsControl(v.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (v *TValueListEditor) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(ValueListEditor_ControlAtPos(v.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (v *TValueListEditor) DisableAlign() {
	ValueListEditor_DisableAlign(v.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (v *TValueListEditor) EnableAlign() {
	ValueListEditor_EnableAlign(v.instance)
}

// 查找子控件。
//
// Find sub controls.
func (v *TValueListEditor) FindChildControl(ControlName string) *TControl {
	return AsControl(ValueListEditor_FindChildControl(v.instance, ControlName))
}

func (v *TValueListEditor) FlipChildren(AllLevels bool) {
	ValueListEditor_FlipChildren(v.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (v *TValueListEditor) Focused() bool {
	return ValueListEditor_Focused(v.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (v *TValueListEditor) HandleAllocated() bool {
	return ValueListEditor_HandleAllocated(v.instance)
}

// 插入一个控件。
//
// Insert a control.
func (v *TValueListEditor) InsertControl(AControl IControl) {
	ValueListEditor_InsertControl(v.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (v *TValueListEditor) Invalidate() {
	ValueListEditor_Invalidate(v.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (v *TValueListEditor) PaintTo(DC HDC, X int32, Y int32) {
	ValueListEditor_PaintTo(v.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (v *TValueListEditor) RemoveControl(AControl IControl) {
	ValueListEditor_RemoveControl(v.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (v *TValueListEditor) Realign() {
	ValueListEditor_Realign(v.instance)
}

// 重绘。
//
// Repaint.
func (v *TValueListEditor) Repaint() {
	ValueListEditor_Repaint(v.instance)
}

// 按比例缩放。
//
// Scale by.
func (v *TValueListEditor) ScaleBy(M int32, D int32) {
	ValueListEditor_ScaleBy(v.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (v *TValueListEditor) ScrollBy(DeltaX int32, DeltaY int32) {
	ValueListEditor_ScrollBy(v.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (v *TValueListEditor) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ValueListEditor_SetBounds(v.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (v *TValueListEditor) SetFocus() {
	ValueListEditor_SetFocus(v.instance)
}

// 控件更新。
//
// Update.
func (v *TValueListEditor) Update() {
	ValueListEditor_Update(v.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (v *TValueListEditor) BringToFront() {
	ValueListEditor_BringToFront(v.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (v *TValueListEditor) ClientToScreen(Point TPoint) TPoint {
	return ValueListEditor_ClientToScreen(v.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (v *TValueListEditor) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ValueListEditor_ClientToParent(v.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (v *TValueListEditor) Dragging() bool {
	return ValueListEditor_Dragging(v.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (v *TValueListEditor) HasParent() bool {
	return ValueListEditor_HasParent(v.instance)
}

// 隐藏控件。
//
// Hidden control.
func (v *TValueListEditor) Hide() {
	ValueListEditor_Hide(v.instance)
}

// 发送一个消息。
//
// Send a message.
func (v *TValueListEditor) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ValueListEditor_Perform(v.instance, Msg, WParam, LParam)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (v *TValueListEditor) ScreenToClient(Point TPoint) TPoint {
	return ValueListEditor_ScreenToClient(v.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (v *TValueListEditor) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ValueListEditor_ParentToClient(v.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (v *TValueListEditor) SendToBack() {
	ValueListEditor_SendToBack(v.instance)
}

// 显示控件。
//
// Show control.
func (v *TValueListEditor) Show() {
	ValueListEditor_Show(v.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (v *TValueListEditor) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ValueListEditor_GetTextBuf(v.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (v *TValueListEditor) GetTextLen() int32 {
	return ValueListEditor_GetTextLen(v.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (v *TValueListEditor) SetTextBuf(Buffer string) {
	ValueListEditor_SetTextBuf(v.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (v *TValueListEditor) FindComponent(AName string) *TComponent {
	return AsComponent(ValueListEditor_FindComponent(v.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (v *TValueListEditor) GetNamePath() string {
	return ValueListEditor_GetNamePath(v.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (v *TValueListEditor) Assign(Source IObject) {
	ValueListEditor_Assign(v.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (v *TValueListEditor) ClassType() TClass {
	return ValueListEditor_ClassType(v.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (v *TValueListEditor) ClassName() string {
	return ValueListEditor_ClassName(v.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (v *TValueListEditor) InstanceSize() int32 {
	return ValueListEditor_InstanceSize(v.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (v *TValueListEditor) InheritsFrom(AClass TClass) bool {
	return ValueListEditor_InheritsFrom(v.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (v *TValueListEditor) Equals(Obj IObject) bool {
	return ValueListEditor_Equals(v.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (v *TValueListEditor) GetHashCode() int32 {
	return ValueListEditor_GetHashCode(v.instance)
}

// 文本类信息。
//
// Text information.
func (v *TValueListEditor) ToString() string {
	return ValueListEditor_ToString(v.instance)
}

func (v *TValueListEditor) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ValueListEditor_AnchorToNeighbour(v.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (v *TValueListEditor) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ValueListEditor_AnchorParallel(v.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (v *TValueListEditor) AnchorHorizontalCenterTo(ASibling IControl) {
	ValueListEditor_AnchorHorizontalCenterTo(v.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (v *TValueListEditor) AnchorVerticalCenterTo(ASibling IControl) {
	ValueListEditor_AnchorVerticalCenterTo(v.instance, CheckPtr(ASibling))
}

func (v *TValueListEditor) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ValueListEditor_AnchorSame(v.instance, ASide, CheckPtr(ASibling))
}

func (v *TValueListEditor) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ValueListEditor_AnchorAsAlign(v.instance, ATheAlign, ASpace)
}

func (v *TValueListEditor) AnchorClient(ASpace int32) {
	ValueListEditor_AnchorClient(v.instance, ASpace)
}

func (v *TValueListEditor) ScaleDesignToForm(ASize int32) int32 {
	return ValueListEditor_ScaleDesignToForm(v.instance, ASize)
}

func (v *TValueListEditor) ScaleFormToDesign(ASize int32) int32 {
	return ValueListEditor_ScaleFormToDesign(v.instance, ASize)
}

func (v *TValueListEditor) Scale96ToForm(ASize int32) int32 {
	return ValueListEditor_Scale96ToForm(v.instance, ASize)
}

func (v *TValueListEditor) ScaleFormTo96(ASize int32) int32 {
	return ValueListEditor_ScaleFormTo96(v.instance, ASize)
}

func (v *TValueListEditor) Scale96ToFont(ASize int32) int32 {
	return ValueListEditor_Scale96ToFont(v.instance, ASize)
}

func (v *TValueListEditor) ScaleFontTo96(ASize int32) int32 {
	return ValueListEditor_ScaleFontTo96(v.instance, ASize)
}

func (v *TValueListEditor) ScaleScreenToFont(ASize int32) int32 {
	return ValueListEditor_ScaleScreenToFont(v.instance, ASize)
}

func (v *TValueListEditor) ScaleFontToScreen(ASize int32) int32 {
	return ValueListEditor_ScaleFontToScreen(v.instance, ASize)
}

func (v *TValueListEditor) Scale96ToScreen(ASize int32) int32 {
	return ValueListEditor_Scale96ToScreen(v.instance, ASize)
}

func (v *TValueListEditor) ScaleScreenTo96(ASize int32) int32 {
	return ValueListEditor_ScaleScreenTo96(v.instance, ASize)
}

func (v *TValueListEditor) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ValueListEditor_AutoAdjustLayout(v.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (v *TValueListEditor) FixDesignFontsPPI(ADesignTimePPI int32) {
	ValueListEditor_FixDesignFontsPPI(v.instance, ADesignTimePPI)
}

func (v *TValueListEditor) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ValueListEditor_ScaleFontsPPI(v.instance, AToPPI, AProportion)
}

func (v *TValueListEditor) ColCount() int32 {
	return ValueListEditor_GetColCount(v.instance)
}

func (v *TValueListEditor) SetColCount(value int32) {
	ValueListEditor_SetColCount(v.instance, value)
}

func (v *TValueListEditor) RowCount() int32 {
	return ValueListEditor_GetRowCount(v.instance)
}

func (v *TValueListEditor) VisibleColCount() int32 {
	return ValueListEditor_GetVisibleColCount(v.instance)
}

func (v *TValueListEditor) VisibleRowCount() int32 {
	return ValueListEditor_GetVisibleRowCount(v.instance)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (v *TValueListEditor) Align() TAlign {
	return ValueListEditor_GetAlign(v.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (v *TValueListEditor) SetAlign(value TAlign) {
	ValueListEditor_SetAlign(v.instance, value)
}

// 获取四个角位置的锚点。
func (v *TValueListEditor) Anchors() TAnchors {
	return ValueListEditor_GetAnchors(v.instance)
}

// 设置四个角位置的锚点。
func (v *TValueListEditor) SetAnchors(value TAnchors) {
	ValueListEditor_SetAnchors(v.instance, value)
}

func (v *TValueListEditor) BiDiMode() TBiDiMode {
	return ValueListEditor_GetBiDiMode(v.instance)
}

func (v *TValueListEditor) SetBiDiMode(value TBiDiMode) {
	ValueListEditor_SetBiDiMode(v.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (v *TValueListEditor) BorderStyle() TBorderStyle {
	return ValueListEditor_GetBorderStyle(v.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (v *TValueListEditor) SetBorderStyle(value TBorderStyle) {
	ValueListEditor_SetBorderStyle(v.instance, value)
}

// 获取颜色。
//
// Get color.
func (v *TValueListEditor) Color() TColor {
	return ValueListEditor_GetColor(v.instance)
}

// 设置颜色。
//
// Set color.
func (v *TValueListEditor) SetColor(value TColor) {
	ValueListEditor_SetColor(v.instance, value)
}

// 获取约束控件大小。
func (v *TValueListEditor) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ValueListEditor_GetConstraints(v.instance))
}

// 设置约束控件大小。
func (v *TValueListEditor) SetConstraints(value *TSizeConstraints) {
	ValueListEditor_SetConstraints(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) DefaultColWidth() int32 {
	return ValueListEditor_GetDefaultColWidth(v.instance)
}

func (v *TValueListEditor) SetDefaultColWidth(value int32) {
	ValueListEditor_SetDefaultColWidth(v.instance, value)
}

func (v *TValueListEditor) DefaultDrawing() bool {
	return ValueListEditor_GetDefaultDrawing(v.instance)
}

func (v *TValueListEditor) SetDefaultDrawing(value bool) {
	ValueListEditor_SetDefaultDrawing(v.instance, value)
}

func (v *TValueListEditor) DefaultRowHeight() int32 {
	return ValueListEditor_GetDefaultRowHeight(v.instance)
}

func (v *TValueListEditor) SetDefaultRowHeight(value int32) {
	ValueListEditor_SetDefaultRowHeight(v.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (v *TValueListEditor) DoubleBuffered() bool {
	return ValueListEditor_GetDoubleBuffered(v.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (v *TValueListEditor) SetDoubleBuffered(value bool) {
	ValueListEditor_SetDoubleBuffered(v.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (v *TValueListEditor) DragCursor() TCursor {
	return ValueListEditor_GetDragCursor(v.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (v *TValueListEditor) SetDragCursor(value TCursor) {
	ValueListEditor_SetDragCursor(v.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (v *TValueListEditor) DragKind() TDragKind {
	return ValueListEditor_GetDragKind(v.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (v *TValueListEditor) SetDragKind(value TDragKind) {
	ValueListEditor_SetDragKind(v.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (v *TValueListEditor) DragMode() TDragMode {
	return ValueListEditor_GetDragMode(v.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (v *TValueListEditor) SetDragMode(value TDragMode) {
	ValueListEditor_SetDragMode(v.instance, value)
}

func (v *TValueListEditor) DropDownRows() int32 {
	return ValueListEditor_GetDropDownRows(v.instance)
}

func (v *TValueListEditor) SetDropDownRows(value int32) {
	ValueListEditor_SetDropDownRows(v.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (v *TValueListEditor) Enabled() bool {
	return ValueListEditor_GetEnabled(v.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (v *TValueListEditor) SetEnabled(value bool) {
	ValueListEditor_SetEnabled(v.instance, value)
}

func (v *TValueListEditor) FixedColor() TColor {
	return ValueListEditor_GetFixedColor(v.instance)
}

func (v *TValueListEditor) SetFixedColor(value TColor) {
	ValueListEditor_SetFixedColor(v.instance, value)
}

func (v *TValueListEditor) FixedCols() int32 {
	return ValueListEditor_GetFixedCols(v.instance)
}

func (v *TValueListEditor) SetFixedCols(value int32) {
	ValueListEditor_SetFixedCols(v.instance, value)
}

// 获取字体。
//
// Get Font.
func (v *TValueListEditor) Font() *TFont {
	return AsFont(ValueListEditor_GetFont(v.instance))
}

// 设置字体。
//
// Set Font.
func (v *TValueListEditor) SetFont(value *TFont) {
	ValueListEditor_SetFont(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) GridLineWidth() int32 {
	return ValueListEditor_GetGridLineWidth(v.instance)
}

func (v *TValueListEditor) SetGridLineWidth(value int32) {
	ValueListEditor_SetGridLineWidth(v.instance, value)
}

func (v *TValueListEditor) Options() TGridOptions {
	return ValueListEditor_GetOptions(v.instance)
}

func (v *TValueListEditor) SetOptions(value TGridOptions) {
	ValueListEditor_SetOptions(v.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (v *TValueListEditor) ParentColor() bool {
	return ValueListEditor_GetParentColor(v.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (v *TValueListEditor) SetParentColor(value bool) {
	ValueListEditor_SetParentColor(v.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (v *TValueListEditor) ParentDoubleBuffered() bool {
	return ValueListEditor_GetParentDoubleBuffered(v.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (v *TValueListEditor) SetParentDoubleBuffered(value bool) {
	ValueListEditor_SetParentDoubleBuffered(v.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (v *TValueListEditor) ParentFont() bool {
	return ValueListEditor_GetParentFont(v.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (v *TValueListEditor) SetParentFont(value bool) {
	ValueListEditor_SetParentFont(v.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (v *TValueListEditor) ParentShowHint() bool {
	return ValueListEditor_GetParentShowHint(v.instance)
}

// 设置以父容器的ShowHint属性为准。
func (v *TValueListEditor) SetParentShowHint(value bool) {
	ValueListEditor_SetParentShowHint(v.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (v *TValueListEditor) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ValueListEditor_GetPopupMenu(v.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (v *TValueListEditor) SetPopupMenu(value IComponent) {
	ValueListEditor_SetPopupMenu(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) ScrollBars() TScrollStyle {
	return ValueListEditor_GetScrollBars(v.instance)
}

func (v *TValueListEditor) SetScrollBars(value TScrollStyle) {
	ValueListEditor_SetScrollBars(v.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (v *TValueListEditor) ShowHint() bool {
	return ValueListEditor_GetShowHint(v.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (v *TValueListEditor) SetShowHint(value bool) {
	ValueListEditor_SetShowHint(v.instance, value)
}

func (v *TValueListEditor) Strings() *TStrings {
	return AsStrings(ValueListEditor_GetStrings(v.instance))
}

func (v *TValueListEditor) SetStrings(value IStrings) {
	ValueListEditor_SetStrings(v.instance, CheckPtr(value))
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (v *TValueListEditor) TabOrder() TTabOrder {
	return ValueListEditor_GetTabOrder(v.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (v *TValueListEditor) SetTabOrder(value TTabOrder) {
	ValueListEditor_SetTabOrder(v.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (v *TValueListEditor) Visible() bool {
	return ValueListEditor_GetVisible(v.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (v *TValueListEditor) SetVisible(value bool) {
	ValueListEditor_SetVisible(v.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (v *TValueListEditor) SetOnClick(fn TNotifyEvent) {
	ValueListEditor_SetOnClick(v.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (v *TValueListEditor) SetOnContextPopup(fn TContextPopupEvent) {
	ValueListEditor_SetOnContextPopup(v.instance, fn)
}

// 设置双击事件。
func (v *TValueListEditor) SetOnDblClick(fn TNotifyEvent) {
	ValueListEditor_SetOnDblClick(v.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (v *TValueListEditor) SetOnDragDrop(fn TDragDropEvent) {
	ValueListEditor_SetOnDragDrop(v.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (v *TValueListEditor) SetOnDragOver(fn TDragOverEvent) {
	ValueListEditor_SetOnDragOver(v.instance, fn)
}

func (v *TValueListEditor) SetOnDrawCell(fn TDrawCellEvent) {
	ValueListEditor_SetOnDrawCell(v.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (v *TValueListEditor) SetOnEndDock(fn TEndDragEvent) {
	ValueListEditor_SetOnEndDock(v.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (v *TValueListEditor) SetOnEndDrag(fn TEndDragEvent) {
	ValueListEditor_SetOnEndDrag(v.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (v *TValueListEditor) SetOnEnter(fn TNotifyEvent) {
	ValueListEditor_SetOnEnter(v.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (v *TValueListEditor) SetOnExit(fn TNotifyEvent) {
	ValueListEditor_SetOnExit(v.instance, fn)
}

func (v *TValueListEditor) SetOnGetEditMask(fn TGetEditEvent) {
	ValueListEditor_SetOnGetEditMask(v.instance, fn)
}

func (v *TValueListEditor) SetOnGetEditText(fn TGetEditEvent) {
	ValueListEditor_SetOnGetEditText(v.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (v *TValueListEditor) SetOnKeyDown(fn TKeyEvent) {
	ValueListEditor_SetOnKeyDown(v.instance, fn)
}

// 设置键键下事件。
func (v *TValueListEditor) SetOnKeyPress(fn TKeyPressEvent) {
	ValueListEditor_SetOnKeyPress(v.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (v *TValueListEditor) SetOnKeyUp(fn TKeyEvent) {
	ValueListEditor_SetOnKeyUp(v.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (v *TValueListEditor) SetOnMouseDown(fn TMouseEvent) {
	ValueListEditor_SetOnMouseDown(v.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (v *TValueListEditor) SetOnMouseEnter(fn TNotifyEvent) {
	ValueListEditor_SetOnMouseEnter(v.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (v *TValueListEditor) SetOnMouseLeave(fn TNotifyEvent) {
	ValueListEditor_SetOnMouseLeave(v.instance, fn)
}

// 设置鼠标移动事件。
func (v *TValueListEditor) SetOnMouseMove(fn TMouseMoveEvent) {
	ValueListEditor_SetOnMouseMove(v.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (v *TValueListEditor) SetOnMouseUp(fn TMouseEvent) {
	ValueListEditor_SetOnMouseUp(v.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (v *TValueListEditor) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	ValueListEditor_SetOnMouseWheelDown(v.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (v *TValueListEditor) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	ValueListEditor_SetOnMouseWheelUp(v.instance, fn)
}

func (v *TValueListEditor) SetOnSelectCell(fn TSelectCellEvent) {
	ValueListEditor_SetOnSelectCell(v.instance, fn)
}

func (v *TValueListEditor) SetOnSetEditText(fn TSetEditEvent) {
	ValueListEditor_SetOnSetEditText(v.instance, fn)
}

// 设置启动停靠。
func (v *TValueListEditor) SetOnStartDock(fn TStartDockEvent) {
	ValueListEditor_SetOnStartDock(v.instance, fn)
}

func (v *TValueListEditor) SetOnTopLeftChanged(fn TNotifyEvent) {
	ValueListEditor_SetOnTopLeftChanged(v.instance, fn)
}

// 获取画布。
func (v *TValueListEditor) Canvas() *TCanvas {
	return AsCanvas(ValueListEditor_GetCanvas(v.instance))
}

func (v *TValueListEditor) Col() int32 {
	return ValueListEditor_GetCol(v.instance)
}

func (v *TValueListEditor) SetCol(value int32) {
	ValueListEditor_SetCol(v.instance, value)
}

func (v *TValueListEditor) EditorMode() bool {
	return ValueListEditor_GetEditorMode(v.instance)
}

func (v *TValueListEditor) SetEditorMode(value bool) {
	ValueListEditor_SetEditorMode(v.instance, value)
}

func (v *TValueListEditor) GridHeight() int32 {
	return ValueListEditor_GetGridHeight(v.instance)
}

func (v *TValueListEditor) GridWidth() int32 {
	return ValueListEditor_GetGridWidth(v.instance)
}

func (v *TValueListEditor) LeftCol() int32 {
	return ValueListEditor_GetLeftCol(v.instance)
}

func (v *TValueListEditor) SetLeftCol(value int32) {
	ValueListEditor_SetLeftCol(v.instance, value)
}

func (v *TValueListEditor) Selection() TGridRect {
	return ValueListEditor_GetSelection(v.instance)
}

func (v *TValueListEditor) SetSelection(value TGridRect) {
	ValueListEditor_SetSelection(v.instance, value)
}

func (v *TValueListEditor) Row() int32 {
	return ValueListEditor_GetRow(v.instance)
}

func (v *TValueListEditor) SetRow(value int32) {
	ValueListEditor_SetRow(v.instance, value)
}

func (v *TValueListEditor) TopRow() int32 {
	return ValueListEditor_GetTopRow(v.instance)
}

func (v *TValueListEditor) SetTopRow(value int32) {
	ValueListEditor_SetTopRow(v.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (v *TValueListEditor) TabStop() bool {
	return ValueListEditor_GetTabStop(v.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (v *TValueListEditor) SetTabStop(value bool) {
	ValueListEditor_SetTabStop(v.instance, value)
}

// 获取依靠客户端总数。
func (v *TValueListEditor) DockClientCount() int32 {
	return ValueListEditor_GetDockClientCount(v.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (v *TValueListEditor) DockSite() bool {
	return ValueListEditor_GetDockSite(v.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (v *TValueListEditor) SetDockSite(value bool) {
	ValueListEditor_SetDockSite(v.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (v *TValueListEditor) MouseInClient() bool {
	return ValueListEditor_GetMouseInClient(v.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (v *TValueListEditor) VisibleDockClientCount() int32 {
	return ValueListEditor_GetVisibleDockClientCount(v.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (v *TValueListEditor) Brush() *TBrush {
	return AsBrush(ValueListEditor_GetBrush(v.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (v *TValueListEditor) ControlCount() int32 {
	return ValueListEditor_GetControlCount(v.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (v *TValueListEditor) Handle() HWND {
	return ValueListEditor_GetHandle(v.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (v *TValueListEditor) ParentWindow() HWND {
	return ValueListEditor_GetParentWindow(v.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (v *TValueListEditor) SetParentWindow(value HWND) {
	ValueListEditor_SetParentWindow(v.instance, value)
}

func (v *TValueListEditor) Showing() bool {
	return ValueListEditor_GetShowing(v.instance)
}

// 获取使用停靠管理。
func (v *TValueListEditor) UseDockManager() bool {
	return ValueListEditor_GetUseDockManager(v.instance)
}

// 设置使用停靠管理。
func (v *TValueListEditor) SetUseDockManager(value bool) {
	ValueListEditor_SetUseDockManager(v.instance, value)
}

func (v *TValueListEditor) Action() *TAction {
	return AsAction(ValueListEditor_GetAction(v.instance))
}

func (v *TValueListEditor) SetAction(value IComponent) {
	ValueListEditor_SetAction(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) BoundsRect() TRect {
	return ValueListEditor_GetBoundsRect(v.instance)
}

func (v *TValueListEditor) SetBoundsRect(value TRect) {
	ValueListEditor_SetBoundsRect(v.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (v *TValueListEditor) ClientHeight() int32 {
	return ValueListEditor_GetClientHeight(v.instance)
}

// 设置客户区高度。
//
// Set client height.
func (v *TValueListEditor) SetClientHeight(value int32) {
	ValueListEditor_SetClientHeight(v.instance, value)
}

func (v *TValueListEditor) ClientOrigin() TPoint {
	return ValueListEditor_GetClientOrigin(v.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (v *TValueListEditor) ClientRect() TRect {
	return ValueListEditor_GetClientRect(v.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (v *TValueListEditor) ClientWidth() int32 {
	return ValueListEditor_GetClientWidth(v.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (v *TValueListEditor) SetClientWidth(value int32) {
	ValueListEditor_SetClientWidth(v.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (v *TValueListEditor) ControlState() TControlState {
	return ValueListEditor_GetControlState(v.instance)
}

// 设置控件状态。
//
// Set control state.
func (v *TValueListEditor) SetControlState(value TControlState) {
	ValueListEditor_SetControlState(v.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (v *TValueListEditor) ControlStyle() TControlStyle {
	return ValueListEditor_GetControlStyle(v.instance)
}

// 设置控件样式。
//
// Set control style.
func (v *TValueListEditor) SetControlStyle(value TControlStyle) {
	ValueListEditor_SetControlStyle(v.instance, value)
}

func (v *TValueListEditor) Floating() bool {
	return ValueListEditor_GetFloating(v.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (v *TValueListEditor) Parent() *TWinControl {
	return AsWinControl(ValueListEditor_GetParent(v.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (v *TValueListEditor) SetParent(value IWinControl) {
	ValueListEditor_SetParent(v.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (v *TValueListEditor) Left() int32 {
	return ValueListEditor_GetLeft(v.instance)
}

// 设置左边位置。
//
// Set Left position.
func (v *TValueListEditor) SetLeft(value int32) {
	ValueListEditor_SetLeft(v.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (v *TValueListEditor) Top() int32 {
	return ValueListEditor_GetTop(v.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (v *TValueListEditor) SetTop(value int32) {
	ValueListEditor_SetTop(v.instance, value)
}

// 获取宽度。
//
// Get width.
func (v *TValueListEditor) Width() int32 {
	return ValueListEditor_GetWidth(v.instance)
}

// 设置宽度。
//
// Set width.
func (v *TValueListEditor) SetWidth(value int32) {
	ValueListEditor_SetWidth(v.instance, value)
}

// 获取高度。
//
// Get height.
func (v *TValueListEditor) Height() int32 {
	return ValueListEditor_GetHeight(v.instance)
}

// 设置高度。
//
// Set height.
func (v *TValueListEditor) SetHeight(value int32) {
	ValueListEditor_SetHeight(v.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (v *TValueListEditor) Cursor() TCursor {
	return ValueListEditor_GetCursor(v.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (v *TValueListEditor) SetCursor(value TCursor) {
	ValueListEditor_SetCursor(v.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (v *TValueListEditor) Hint() string {
	return ValueListEditor_GetHint(v.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (v *TValueListEditor) SetHint(value string) {
	ValueListEditor_SetHint(v.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (v *TValueListEditor) ComponentCount() int32 {
	return ValueListEditor_GetComponentCount(v.instance)
}

// 获取组件索引。
//
// Get component index.
func (v *TValueListEditor) ComponentIndex() int32 {
	return ValueListEditor_GetComponentIndex(v.instance)
}

// 设置组件索引。
//
// Set component index.
func (v *TValueListEditor) SetComponentIndex(value int32) {
	ValueListEditor_SetComponentIndex(v.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (v *TValueListEditor) Owner() *TComponent {
	return AsComponent(ValueListEditor_GetOwner(v.instance))
}

// 获取组件名称。
//
// Get the component name.
func (v *TValueListEditor) Name() string {
	return ValueListEditor_GetName(v.instance)
}

// 设置组件名称。
//
// Set the component name.
func (v *TValueListEditor) SetName(value string) {
	ValueListEditor_SetName(v.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (v *TValueListEditor) Tag() int {
	return ValueListEditor_GetTag(v.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (v *TValueListEditor) SetTag(value int) {
	ValueListEditor_SetTag(v.instance, value)
}

// 获取左边锚点。
func (v *TValueListEditor) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ValueListEditor_GetAnchorSideLeft(v.instance))
}

// 设置左边锚点。
func (v *TValueListEditor) SetAnchorSideLeft(value *TAnchorSide) {
	ValueListEditor_SetAnchorSideLeft(v.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (v *TValueListEditor) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ValueListEditor_GetAnchorSideTop(v.instance))
}

// 设置顶边锚点。
func (v *TValueListEditor) SetAnchorSideTop(value *TAnchorSide) {
	ValueListEditor_SetAnchorSideTop(v.instance, CheckPtr(value))
}

// 获取右边锚点。
func (v *TValueListEditor) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ValueListEditor_GetAnchorSideRight(v.instance))
}

// 设置右边锚点。
func (v *TValueListEditor) SetAnchorSideRight(value *TAnchorSide) {
	ValueListEditor_SetAnchorSideRight(v.instance, CheckPtr(value))
}

// 获取底边锚点。
func (v *TValueListEditor) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ValueListEditor_GetAnchorSideBottom(v.instance))
}

// 设置底边锚点。
func (v *TValueListEditor) SetAnchorSideBottom(value *TAnchorSide) {
	ValueListEditor_SetAnchorSideBottom(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(ValueListEditor_GetChildSizing(v.instance))
}

func (v *TValueListEditor) SetChildSizing(value *TControlChildSizing) {
	ValueListEditor_SetChildSizing(v.instance, CheckPtr(value))
}

// 获取边框间距。
func (v *TValueListEditor) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ValueListEditor_GetBorderSpacing(v.instance))
}

// 设置边框间距。
func (v *TValueListEditor) SetBorderSpacing(value *TControlBorderSpacing) {
	ValueListEditor_SetBorderSpacing(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) Cells(ACol int32, ARow int32) string {
	return ValueListEditor_GetCells(v.instance, ACol, ARow)
}

func (v *TValueListEditor) SetCells(ACol int32, ARow int32, value string) {
	ValueListEditor_SetCells(v.instance, ACol, ARow, value)
}

func (v *TValueListEditor) Values(Key string) string {
	return ValueListEditor_GetValues(v.instance, Key)
}

func (v *TValueListEditor) SetValues(Key string, value string) {
	ValueListEditor_SetValues(v.instance, Key, value)
}

func (v *TValueListEditor) ColWidths(Index int32) int32 {
	return ValueListEditor_GetColWidths(v.instance, Index)
}

func (v *TValueListEditor) SetColWidths(Index int32, value int32) {
	ValueListEditor_SetColWidths(v.instance, Index, value)
}

func (v *TValueListEditor) RowHeights(Index int32) int32 {
	return ValueListEditor_GetRowHeights(v.instance, Index)
}

func (v *TValueListEditor) SetRowHeights(Index int32, value int32) {
	ValueListEditor_SetRowHeights(v.instance, Index, value)
}

// 获取指定索引停靠客户端。
func (v *TValueListEditor) DockClients(Index int32) *TControl {
	return AsControl(ValueListEditor_GetDockClients(v.instance, Index))
}

// 获取指定索引子控件。
func (v *TValueListEditor) Controls(Index int32) *TControl {
	return AsControl(ValueListEditor_GetControls(v.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (v *TValueListEditor) Components(AIndex int32) *TComponent {
	return AsComponent(ValueListEditor_GetComponents(v.instance, AIndex))
}

// 获取锚侧面。
func (v *TValueListEditor) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ValueListEditor_GetAnchorSide(v.instance, AKind))
}
