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

type TButton struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewButton(owner IComponent) *TButton {
	b := new(TButton)
	b.instance = Button_Create(CheckPtr(owner))
	b.ptr = unsafe.Pointer(b.instance)
	return b
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsButton(obj any) *TButton {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsButton.
func ButtonFromInst(inst uintptr) *TButton {
	return AsButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsButton.
func ButtonFromObj(obj IObject) *TButton {
	return AsButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsButton.
func ButtonFromUnsafePointer(ptr unsafe.Pointer) *TButton {
	return AsButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (b *TButton) Free() {
	if b.instance != 0 {
		Button_Free(b.instance)
		b.instance, b.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TButton) Instance() uintptr {
	return b.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TButton) UnsafeAddr() unsafe.Pointer {
	return b.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TButton) IsValid() bool {
	return b.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TButton) Is() TIs {
	return TIs(b.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TButton) As() TAs {
//    return TAs(b.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TButtonClass() TClass {
	return Button_StaticClassType()
}

// 单击。
func (b *TButton) Click() {
	Button_Click(b.instance)
}

// 是否可以获得焦点。
func (b *TButton) CanFocus() bool {
	return Button_CanFocus(b.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (b *TButton) ContainsControl(Control IControl) bool {
	return Button_ContainsControl(b.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (b *TButton) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(Button_ControlAtPos(b.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (b *TButton) DisableAlign() {
	Button_DisableAlign(b.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (b *TButton) EnableAlign() {
	Button_EnableAlign(b.instance)
}

// 查找子控件。
//
// Find sub controls.
func (b *TButton) FindChildControl(ControlName string) *TControl {
	return AsControl(Button_FindChildControl(b.instance, ControlName))
}

func (b *TButton) FlipChildren(AllLevels bool) {
	Button_FlipChildren(b.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (b *TButton) Focused() bool {
	return Button_Focused(b.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (b *TButton) HandleAllocated() bool {
	return Button_HandleAllocated(b.instance)
}

// 插入一个控件。
//
// Insert a control.
func (b *TButton) InsertControl(AControl IControl) {
	Button_InsertControl(b.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (b *TButton) Invalidate() {
	Button_Invalidate(b.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (b *TButton) PaintTo(DC HDC, X int32, Y int32) {
	Button_PaintTo(b.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (b *TButton) RemoveControl(AControl IControl) {
	Button_RemoveControl(b.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (b *TButton) Realign() {
	Button_Realign(b.instance)
}

// 重绘。
//
// Repaint.
func (b *TButton) Repaint() {
	Button_Repaint(b.instance)
}

// 按比例缩放。
//
// Scale by.
func (b *TButton) ScaleBy(M int32, D int32) {
	Button_ScaleBy(b.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (b *TButton) ScrollBy(DeltaX int32, DeltaY int32) {
	Button_ScrollBy(b.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	Button_SetBounds(b.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (b *TButton) SetFocus() {
	Button_SetFocus(b.instance)
}

// 控件更新。
//
// Update.
func (b *TButton) Update() {
	Button_Update(b.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (b *TButton) BringToFront() {
	Button_BringToFront(b.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TButton) ClientToScreen(Point TPoint) TPoint {
	return Button_ClientToScreen(b.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return Button_ClientToParent(b.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TButton) Dragging() bool {
	return Button_Dragging(b.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (b *TButton) HasParent() bool {
	return Button_HasParent(b.instance)
}

// 隐藏控件。
//
// Hidden control.
func (b *TButton) Hide() {
	Button_Hide(b.instance)
}

// 发送一个消息。
//
// Send a message.
func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return Button_Perform(b.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (b *TButton) Refresh() {
	Button_Refresh(b.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TButton) ScreenToClient(Point TPoint) TPoint {
	return Button_ScreenToClient(b.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return Button_ParentToClient(b.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (b *TButton) SendToBack() {
	Button_SendToBack(b.instance)
}

// 显示控件。
//
// Show control.
func (b *TButton) Show() {
	Button_Show(b.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Button_GetTextBuf(b.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TButton) GetTextLen() int32 {
	return Button_GetTextLen(b.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TButton) SetTextBuf(Buffer string) {
	Button_SetTextBuf(b.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TButton) FindComponent(AName string) *TComponent {
	return AsComponent(Button_FindComponent(b.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (b *TButton) GetNamePath() string {
	return Button_GetNamePath(b.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TButton) Assign(Source IObject) {
	Button_Assign(b.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (b *TButton) ClassType() TClass {
	return Button_ClassType(b.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TButton) ClassName() string {
	return Button_ClassName(b.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TButton) InstanceSize() int32 {
	return Button_InstanceSize(b.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TButton) InheritsFrom(AClass TClass) bool {
	return Button_InheritsFrom(b.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (b *TButton) Equals(Obj IObject) bool {
	return Button_Equals(b.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TButton) GetHashCode() int32 {
	return Button_GetHashCode(b.instance)
}

// 文本类信息。
//
// Text information.
func (b *TButton) ToString() string {
	return Button_ToString(b.instance)
}

func (b *TButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Button_AnchorToNeighbour(b.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (b *TButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Button_AnchorParallel(b.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (b *TButton) AnchorHorizontalCenterTo(ASibling IControl) {
	Button_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (b *TButton) AnchorVerticalCenterTo(ASibling IControl) {
	Button_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	Button_AnchorSame(b.instance, ASide, CheckPtr(ASibling))
}

func (b *TButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	Button_AnchorAsAlign(b.instance, ATheAlign, ASpace)
}

func (b *TButton) AnchorClient(ASpace int32) {
	Button_AnchorClient(b.instance, ASpace)
}

func (b *TButton) ScaleDesignToForm(ASize int32) int32 {
	return Button_ScaleDesignToForm(b.instance, ASize)
}

func (b *TButton) ScaleFormToDesign(ASize int32) int32 {
	return Button_ScaleFormToDesign(b.instance, ASize)
}

func (b *TButton) Scale96ToForm(ASize int32) int32 {
	return Button_Scale96ToForm(b.instance, ASize)
}

func (b *TButton) ScaleFormTo96(ASize int32) int32 {
	return Button_ScaleFormTo96(b.instance, ASize)
}

func (b *TButton) Scale96ToFont(ASize int32) int32 {
	return Button_Scale96ToFont(b.instance, ASize)
}

func (b *TButton) ScaleFontTo96(ASize int32) int32 {
	return Button_ScaleFontTo96(b.instance, ASize)
}

func (b *TButton) ScaleScreenToFont(ASize int32) int32 {
	return Button_ScaleScreenToFont(b.instance, ASize)
}

func (b *TButton) ScaleFontToScreen(ASize int32) int32 {
	return Button_ScaleFontToScreen(b.instance, ASize)
}

func (b *TButton) Scale96ToScreen(ASize int32) int32 {
	return Button_Scale96ToScreen(b.instance, ASize)
}

func (b *TButton) ScaleScreenTo96(ASize int32) int32 {
	return Button_ScaleScreenTo96(b.instance, ASize)
}

func (b *TButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	Button_AutoAdjustLayout(b.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (b *TButton) FixDesignFontsPPI(ADesignTimePPI int32) {
	Button_FixDesignFontsPPI(b.instance, ADesignTimePPI)
}

func (b *TButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	Button_ScaleFontsPPI(b.instance, AToPPI, AProportion)
}

func (b *TButton) Action() *TAction {
	return AsAction(Button_GetAction(b.instance))
}

func (b *TButton) SetAction(value IComponent) {
	Button_SetAction(b.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TButton) Align() TAlign {
	return Button_GetAlign(b.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TButton) SetAlign(value TAlign) {
	Button_SetAlign(b.instance, value)
}

// 获取四个角位置的锚点。
func (b *TButton) Anchors() TAnchors {
	return Button_GetAnchors(b.instance)
}

// 设置四个角位置的锚点。
func (b *TButton) SetAnchors(value TAnchors) {
	Button_SetAnchors(b.instance, value)
}

func (b *TButton) BiDiMode() TBiDiMode {
	return Button_GetBiDiMode(b.instance)
}

func (b *TButton) SetBiDiMode(value TBiDiMode) {
	Button_SetBiDiMode(b.instance, value)
}

func (b *TButton) Cancel() bool {
	return Button_GetCancel(b.instance)
}

func (b *TButton) SetCancel(value bool) {
	Button_SetCancel(b.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (b *TButton) Caption() string {
	return Button_GetCaption(b.instance)
}

// 设置控件标题。
//
// Set the control title.
func (b *TButton) SetCaption(value string) {
	Button_SetCaption(b.instance, value)
}

// 获取约束控件大小。
func (b *TButton) Constraints() *TSizeConstraints {
	return AsSizeConstraints(Button_GetConstraints(b.instance))
}

// 设置约束控件大小。
func (b *TButton) SetConstraints(value *TSizeConstraints) {
	Button_SetConstraints(b.instance, CheckPtr(value))
}

func (b *TButton) Default() bool {
	return Button_GetDefault(b.instance)
}

func (b *TButton) SetDefault(value bool) {
	Button_SetDefault(b.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (b *TButton) DoubleBuffered() bool {
	return Button_GetDoubleBuffered(b.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (b *TButton) SetDoubleBuffered(value bool) {
	Button_SetDoubleBuffered(b.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (b *TButton) DragCursor() TCursor {
	return Button_GetDragCursor(b.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (b *TButton) SetDragCursor(value TCursor) {
	Button_SetDragCursor(b.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (b *TButton) DragKind() TDragKind {
	return Button_GetDragKind(b.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (b *TButton) SetDragKind(value TDragKind) {
	Button_SetDragKind(b.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (b *TButton) DragMode() TDragMode {
	return Button_GetDragMode(b.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (b *TButton) SetDragMode(value TDragMode) {
	Button_SetDragMode(b.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (b *TButton) Enabled() bool {
	return Button_GetEnabled(b.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (b *TButton) SetEnabled(value bool) {
	Button_SetEnabled(b.instance, value)
}

// 获取字体。
//
// Get Font.
func (b *TButton) Font() *TFont {
	return AsFont(Button_GetFont(b.instance))
}

// 设置字体。
//
// Set Font.
func (b *TButton) SetFont(value *TFont) {
	Button_SetFont(b.instance, CheckPtr(value))
}

// 获取模态对话框显示结果。
func (b *TButton) ModalResult() TModalResult {
	return Button_GetModalResult(b.instance)
}

// 设置模态对话框显示结果。
func (b *TButton) SetModalResult(value TModalResult) {
	Button_SetModalResult(b.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (b *TButton) ParentDoubleBuffered() bool {
	return Button_GetParentDoubleBuffered(b.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (b *TButton) SetParentDoubleBuffered(value bool) {
	Button_SetParentDoubleBuffered(b.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (b *TButton) ParentFont() bool {
	return Button_GetParentFont(b.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (b *TButton) SetParentFont(value bool) {
	Button_SetParentFont(b.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (b *TButton) ParentShowHint() bool {
	return Button_GetParentShowHint(b.instance)
}

// 设置以父容器的ShowHint属性为准。
func (b *TButton) SetParentShowHint(value bool) {
	Button_SetParentShowHint(b.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (b *TButton) PopupMenu() *TPopupMenu {
	return AsPopupMenu(Button_GetPopupMenu(b.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (b *TButton) SetPopupMenu(value IComponent) {
	Button_SetPopupMenu(b.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TButton) ShowHint() bool {
	return Button_GetShowHint(b.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TButton) SetShowHint(value bool) {
	Button_SetShowHint(b.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (b *TButton) TabOrder() TTabOrder {
	return Button_GetTabOrder(b.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (b *TButton) SetTabOrder(value TTabOrder) {
	Button_SetTabOrder(b.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (b *TButton) TabStop() bool {
	return Button_GetTabStop(b.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (b *TButton) SetTabStop(value bool) {
	Button_SetTabStop(b.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (b *TButton) Visible() bool {
	return Button_GetVisible(b.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (b *TButton) SetVisible(value bool) {
	Button_SetVisible(b.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (b *TButton) SetOnClick(fn TNotifyEvent) {
	Button_SetOnClick(b.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (b *TButton) SetOnContextPopup(fn TContextPopupEvent) {
	Button_SetOnContextPopup(b.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (b *TButton) SetOnDragDrop(fn TDragDropEvent) {
	Button_SetOnDragDrop(b.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (b *TButton) SetOnDragOver(fn TDragOverEvent) {
	Button_SetOnDragOver(b.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (b *TButton) SetOnEndDrag(fn TEndDragEvent) {
	Button_SetOnEndDrag(b.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (b *TButton) SetOnEnter(fn TNotifyEvent) {
	Button_SetOnEnter(b.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (b *TButton) SetOnExit(fn TNotifyEvent) {
	Button_SetOnExit(b.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
	Button_SetOnKeyDown(b.instance, fn)
}

// 设置键键下事件。
func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
	Button_SetOnKeyPress(b.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
	Button_SetOnKeyUp(b.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
	Button_SetOnMouseDown(b.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
	Button_SetOnMouseEnter(b.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
	Button_SetOnMouseLeave(b.instance, fn)
}

// 设置鼠标移动事件。
func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
	Button_SetOnMouseMove(b.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
	Button_SetOnMouseUp(b.instance, fn)
}

// 获取依靠客户端总数。
func (b *TButton) DockClientCount() int32 {
	return Button_GetDockClientCount(b.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (b *TButton) DockSite() bool {
	return Button_GetDockSite(b.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (b *TButton) SetDockSite(value bool) {
	Button_SetDockSite(b.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (b *TButton) MouseInClient() bool {
	return Button_GetMouseInClient(b.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (b *TButton) VisibleDockClientCount() int32 {
	return Button_GetVisibleDockClientCount(b.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (b *TButton) Brush() *TBrush {
	return AsBrush(Button_GetBrush(b.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (b *TButton) ControlCount() int32 {
	return Button_GetControlCount(b.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (b *TButton) Handle() HWND {
	return Button_GetHandle(b.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (b *TButton) ParentWindow() HWND {
	return Button_GetParentWindow(b.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (b *TButton) SetParentWindow(value HWND) {
	Button_SetParentWindow(b.instance, value)
}

func (b *TButton) Showing() bool {
	return Button_GetShowing(b.instance)
}

// 获取使用停靠管理。
func (b *TButton) UseDockManager() bool {
	return Button_GetUseDockManager(b.instance)
}

// 设置使用停靠管理。
func (b *TButton) SetUseDockManager(value bool) {
	Button_SetUseDockManager(b.instance, value)
}

func (b *TButton) BoundsRect() TRect {
	return Button_GetBoundsRect(b.instance)
}

func (b *TButton) SetBoundsRect(value TRect) {
	Button_SetBoundsRect(b.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (b *TButton) ClientHeight() int32 {
	return Button_GetClientHeight(b.instance)
}

// 设置客户区高度。
//
// Set client height.
func (b *TButton) SetClientHeight(value int32) {
	Button_SetClientHeight(b.instance, value)
}

func (b *TButton) ClientOrigin() TPoint {
	return Button_GetClientOrigin(b.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (b *TButton) ClientRect() TRect {
	return Button_GetClientRect(b.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (b *TButton) ClientWidth() int32 {
	return Button_GetClientWidth(b.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (b *TButton) SetClientWidth(value int32) {
	Button_SetClientWidth(b.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (b *TButton) ControlState() TControlState {
	return Button_GetControlState(b.instance)
}

// 设置控件状态。
//
// Set control state.
func (b *TButton) SetControlState(value TControlState) {
	Button_SetControlState(b.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (b *TButton) ControlStyle() TControlStyle {
	return Button_GetControlStyle(b.instance)
}

// 设置控件样式。
//
// Set control style.
func (b *TButton) SetControlStyle(value TControlStyle) {
	Button_SetControlStyle(b.instance, value)
}

func (b *TButton) Floating() bool {
	return Button_GetFloating(b.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (b *TButton) Parent() *TWinControl {
	return AsWinControl(Button_GetParent(b.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (b *TButton) SetParent(value IWinControl) {
	Button_SetParent(b.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (b *TButton) Left() int32 {
	return Button_GetLeft(b.instance)
}

// 设置左边位置。
//
// Set Left position.
func (b *TButton) SetLeft(value int32) {
	Button_SetLeft(b.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (b *TButton) Top() int32 {
	return Button_GetTop(b.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (b *TButton) SetTop(value int32) {
	Button_SetTop(b.instance, value)
}

// 获取宽度。
//
// Get width.
func (b *TButton) Width() int32 {
	return Button_GetWidth(b.instance)
}

// 设置宽度。
//
// Set width.
func (b *TButton) SetWidth(value int32) {
	Button_SetWidth(b.instance, value)
}

// 获取高度。
//
// Get height.
func (b *TButton) Height() int32 {
	return Button_GetHeight(b.instance)
}

// 设置高度。
//
// Set height.
func (b *TButton) SetHeight(value int32) {
	Button_SetHeight(b.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (b *TButton) Cursor() TCursor {
	return Button_GetCursor(b.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (b *TButton) SetCursor(value TCursor) {
	Button_SetCursor(b.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TButton) Hint() string {
	return Button_GetHint(b.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TButton) SetHint(value string) {
	Button_SetHint(b.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (b *TButton) ComponentCount() int32 {
	return Button_GetComponentCount(b.instance)
}

// 获取组件索引。
//
// Get component index.
func (b *TButton) ComponentIndex() int32 {
	return Button_GetComponentIndex(b.instance)
}

// 设置组件索引。
//
// Set component index.
func (b *TButton) SetComponentIndex(value int32) {
	Button_SetComponentIndex(b.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (b *TButton) Owner() *TComponent {
	return AsComponent(Button_GetOwner(b.instance))
}

// 获取组件名称。
//
// Get the component name.
func (b *TButton) Name() string {
	return Button_GetName(b.instance)
}

// 设置组件名称。
//
// Set the component name.
func (b *TButton) SetName(value string) {
	Button_SetName(b.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (b *TButton) Tag() int {
	return Button_GetTag(b.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (b *TButton) SetTag(value int) {
	Button_SetTag(b.instance, value)
}

// 获取左边锚点。
func (b *TButton) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(Button_GetAnchorSideLeft(b.instance))
}

// 设置左边锚点。
func (b *TButton) SetAnchorSideLeft(value *TAnchorSide) {
	Button_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (b *TButton) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(Button_GetAnchorSideTop(b.instance))
}

// 设置顶边锚点。
func (b *TButton) SetAnchorSideTop(value *TAnchorSide) {
	Button_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// 获取右边锚点。
func (b *TButton) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(Button_GetAnchorSideRight(b.instance))
}

// 设置右边锚点。
func (b *TButton) SetAnchorSideRight(value *TAnchorSide) {
	Button_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// 获取底边锚点。
func (b *TButton) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(Button_GetAnchorSideBottom(b.instance))
}

// 设置底边锚点。
func (b *TButton) SetAnchorSideBottom(value *TAnchorSide) {
	Button_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

func (b *TButton) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(Button_GetChildSizing(b.instance))
}

func (b *TButton) SetChildSizing(value *TControlChildSizing) {
	Button_SetChildSizing(b.instance, CheckPtr(value))
}

// 获取边框间距。
func (b *TButton) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(Button_GetBorderSpacing(b.instance))
}

// 设置边框间距。
func (b *TButton) SetBorderSpacing(value *TControlBorderSpacing) {
	Button_SetBorderSpacing(b.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (b *TButton) DockClients(Index int32) *TControl {
	return AsControl(Button_GetDockClients(b.instance, Index))
}

// 获取指定索引子控件。
func (b *TButton) Controls(Index int32) *TControl {
	return AsControl(Button_GetControls(b.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (b *TButton) Components(AIndex int32) *TComponent {
	return AsComponent(Button_GetComponents(b.instance, AIndex))
}

// 获取锚侧面。
func (b *TButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(Button_GetAnchorSide(b.instance, AKind))
}
