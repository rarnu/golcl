package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TChart struct {
	IChart
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewChart(owner IComponent) *TChart {
	c := new(TChart)
	c.instance = Chart_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsChart(obj any) *TChart {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChart{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsChart.
func ChartFromInst(inst uintptr) *TChart {
	return AsChart(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsChart.
func ChartFromObj(obj IObject) *TChart {
	return AsChart(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsChart.
func ChartFromUnsafePointer(ptr unsafe.Pointer) *TChart {
	return AsChart(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (w *TChart) Free() {
	if w.instance != 0 {
		Chart_Free(w.instance)
		w.instance, w.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (w *TChart) Instance() uintptr {
	return w.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (w *TChart) UnsafeAddr() unsafe.Pointer {
	return w.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (w *TChart) IsValid() bool {
	return w.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (w *TChart) Is() TIs {
	return TIs(w.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (w *TChart) As() TAs {
//    return TAs(w.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TChartClass() TClass {
	return Chart_StaticClassType()
}

// 是否可以获得焦点。
func (w *TChart) CanFocus() bool {
	return Chart_CanFocus(w.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (w *TChart) ContainsControl(Control IControl) bool {
	return Chart_ContainsControl(w.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (w *TChart) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(Chart_ControlAtPos(w.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (w *TChart) DisableAlign() {
	Chart_DisableAlign(w.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (w *TChart) EnableAlign() {
	Chart_EnableAlign(w.instance)
}

// 查找子控件。
//
// Find sub controls.
func (w *TChart) FindChildControl(ControlName string) *TControl {
	return AsControl(Chart_FindChildControl(w.instance, ControlName))
}

func (w *TChart) FlipChildren(AllLevels bool) {
	Chart_FlipChildren(w.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (w *TChart) Focused() bool {
	return Chart_Focused(w.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (w *TChart) HandleAllocated() bool {
	return Chart_HandleAllocated(w.instance)
}

// 插入一个控件。
//
// Insert a control.
func (w *TChart) InsertControl(AControl IControl) {
	Chart_InsertControl(w.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (w *TChart) Invalidate() {
	Chart_Invalidate(w.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (w *TChart) PaintTo(DC HDC, X int32, Y int32) {
	Chart_PaintTo(w.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (w *TChart) RemoveControl(AControl IControl) {
	Chart_RemoveControl(w.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (w *TChart) Realign() {
	Chart_Realign(w.instance)
}

// 重绘。
//
// Repaint.
func (w *TChart) Repaint() {
	Chart_Repaint(w.instance)
}

// 按比例缩放。
//
// Scale by.
func (w *TChart) ScaleBy(M int32, D int32) {
	Chart_ScaleBy(w.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (w *TChart) ScrollBy(DeltaX int32, DeltaY int32) {
	Chart_ScrollBy(w.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (w *TChart) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	Chart_SetBounds(w.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (w *TChart) SetFocus() {
	Chart_SetFocus(w.instance)
}

// 控件更新。
//
// Update.
func (w *TChart) Update() {
	Chart_Update(w.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (w *TChart) BringToFront() {
	Chart_BringToFront(w.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (w *TChart) ClientToScreen(Point TPoint) TPoint {
	return Chart_ClientToScreen(w.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (w *TChart) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return Chart_ClientToParent(w.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (w *TChart) Dragging() bool {
	return Chart_Dragging(w.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (w *TChart) HasParent() bool {
	return Chart_HasParent(w.instance)
}

// 隐藏控件。
//
// Hidden control.
func (w *TChart) Hide() {
	Chart_Hide(w.instance)
}

// 发送一个消息。
//
// Send a message.
func (w *TChart) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return Chart_Perform(w.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (w *TChart) Refresh() {
	Chart_Refresh(w.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (w *TChart) ScreenToClient(Point TPoint) TPoint {
	return Chart_ScreenToClient(w.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (w *TChart) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return Chart_ParentToClient(w.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (w *TChart) SendToBack() {
	Chart_SendToBack(w.instance)
}

// 显示控件。
//
// Show control.
func (w *TChart) Show() {
	Chart_Show(w.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (w *TChart) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Chart_GetTextBuf(w.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (w *TChart) GetTextLen() int32 {
	return Chart_GetTextLen(w.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (w *TChart) SetTextBuf(Buffer string) {
	Chart_SetTextBuf(w.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (w *TChart) FindComponent(AName string) *TComponent {
	return AsComponent(Chart_FindComponent(w.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (w *TChart) GetNamePath() string {
	return Chart_GetNamePath(w.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (w *TChart) Assign(Source IObject) {
	Chart_Assign(w.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (w *TChart) ClassType() TClass {
	return Chart_ClassType(w.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (w *TChart) ClassName() string {
	return Chart_ClassName(w.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (w *TChart) InstanceSize() int32 {
	return Chart_InstanceSize(w.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (w *TChart) InheritsFrom(AClass TClass) bool {
	return Chart_InheritsFrom(w.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (w *TChart) Equals(Obj IObject) bool {
	return Chart_Equals(w.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (w *TChart) GetHashCode() int32 {
	return Chart_GetHashCode(w.instance)
}

// 文本类信息。
//
// Text information.
func (w *TChart) ToString() string {
	return Chart_ToString(w.instance)
}

func (w *TChart) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Chart_AnchorToNeighbour(w.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (w *TChart) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Chart_AnchorParallel(w.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (w *TChart) AnchorHorizontalCenterTo(ASibling IControl) {
	Chart_AnchorHorizontalCenterTo(w.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (w *TChart) AnchorVerticalCenterTo(ASibling IControl) {
	Chart_AnchorVerticalCenterTo(w.instance, CheckPtr(ASibling))
}

func (w *TChart) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	Chart_AnchorSame(w.instance, ASide, CheckPtr(ASibling))
}

func (w *TChart) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	Chart_AnchorAsAlign(w.instance, ATheAlign, ASpace)
}

func (w *TChart) AnchorClient(ASpace int32) {
	Chart_AnchorClient(w.instance, ASpace)
}

func (w *TChart) ScaleDesignToForm(ASize int32) int32 {
	return Chart_ScaleDesignToForm(w.instance, ASize)
}

func (w *TChart) ScaleFormToDesign(ASize int32) int32 {
	return Chart_ScaleFormToDesign(w.instance, ASize)
}

func (w *TChart) Scale96ToForm(ASize int32) int32 {
	return Chart_Scale96ToForm(w.instance, ASize)
}

func (w *TChart) ScaleFormTo96(ASize int32) int32 {
	return Chart_ScaleFormTo96(w.instance, ASize)
}

func (w *TChart) Scale96ToFont(ASize int32) int32 {
	return Chart_Scale96ToFont(w.instance, ASize)
}

func (w *TChart) ScaleFontTo96(ASize int32) int32 {
	return Chart_ScaleFontTo96(w.instance, ASize)
}

func (w *TChart) ScaleScreenToFont(ASize int32) int32 {
	return Chart_ScaleScreenToFont(w.instance, ASize)
}

func (w *TChart) ScaleFontToScreen(ASize int32) int32 {
	return Chart_ScaleFontToScreen(w.instance, ASize)
}

func (w *TChart) Scale96ToScreen(ASize int32) int32 {
	return Chart_Scale96ToScreen(w.instance, ASize)
}

func (w *TChart) ScaleScreenTo96(ASize int32) int32 {
	return Chart_ScaleScreenTo96(w.instance, ASize)
}

func (w *TChart) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	Chart_AutoAdjustLayout(w.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (w *TChart) FixDesignFontsPPI(ADesignTimePPI int32) {
	Chart_FixDesignFontsPPI(w.instance, ADesignTimePPI)
}

func (w *TChart) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	Chart_ScaleFontsPPI(w.instance, AToPPI, AProportion)
}

// 获取依靠客户端总数。
func (w *TChart) DockClientCount() int32 {
	return Chart_GetDockClientCount(w.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (w *TChart) DockSite() bool {
	return Chart_GetDockSite(w.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (w *TChart) SetDockSite(value bool) {
	Chart_SetDockSite(w.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (w *TChart) DoubleBuffered() bool {
	return Chart_GetDoubleBuffered(w.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (w *TChart) SetDoubleBuffered(value bool) {
	Chart_SetDoubleBuffered(w.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (w *TChart) MouseInClient() bool {
	return Chart_GetMouseInClient(w.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (w *TChart) VisibleDockClientCount() int32 {
	return Chart_GetVisibleDockClientCount(w.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (w *TChart) Brush() *TBrush {
	return AsBrush(Chart_GetBrush(w.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (w *TChart) ControlCount() int32 {
	return Chart_GetControlCount(w.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (w *TChart) Handle() HWND {
	return Chart_GetHandle(w.instance)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (w *TChart) ParentDoubleBuffered() bool {
	return Chart_GetParentDoubleBuffered(w.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (w *TChart) SetParentDoubleBuffered(value bool) {
	Chart_SetParentDoubleBuffered(w.instance, value)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (w *TChart) ParentWindow() HWND {
	return Chart_GetParentWindow(w.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (w *TChart) SetParentWindow(value HWND) {
	Chart_SetParentWindow(w.instance, value)
}

func (w *TChart) Showing() bool {
	return Chart_GetShowing(w.instance)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (w *TChart) TabOrder() TTabOrder {
	return Chart_GetTabOrder(w.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (w *TChart) SetTabOrder(value TTabOrder) {
	Chart_SetTabOrder(w.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (w *TChart) TabStop() bool {
	return Chart_GetTabStop(w.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (w *TChart) SetTabStop(value bool) {
	Chart_SetTabStop(w.instance, value)
}

// 获取使用停靠管理。
func (w *TChart) UseDockManager() bool {
	return Chart_GetUseDockManager(w.instance)
}

// 设置使用停靠管理。
func (w *TChart) SetUseDockManager(value bool) {
	Chart_SetUseDockManager(w.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (w *TChart) Enabled() bool {
	return Chart_GetEnabled(w.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (w *TChart) SetEnabled(value bool) {
	Chart_SetEnabled(w.instance, value)
}

func (w *TChart) Action() *TAction {
	return AsAction(Chart_GetAction(w.instance))
}

func (w *TChart) SetAction(value IComponent) {
	Chart_SetAction(w.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (w *TChart) Align() TAlign {
	return Chart_GetAlign(w.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (w *TChart) SetAlign(value TAlign) {
	Chart_SetAlign(w.instance, value)
}

// 获取四个角位置的锚点。
func (w *TChart) Anchors() TAnchors {
	return Chart_GetAnchors(w.instance)
}

// 设置四个角位置的锚点。
func (w *TChart) SetAnchors(value TAnchors) {
	Chart_SetAnchors(w.instance, value)
}

func (w *TChart) BiDiMode() TBiDiMode {
	return Chart_GetBiDiMode(w.instance)
}

func (w *TChart) SetBiDiMode(value TBiDiMode) {
	Chart_SetBiDiMode(w.instance, value)
}

func (w *TChart) BoundsRect() TRect {
	return Chart_GetBoundsRect(w.instance)
}

func (w *TChart) SetBoundsRect(value TRect) {
	Chart_SetBoundsRect(w.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (w *TChart) ClientHeight() int32 {
	return Chart_GetClientHeight(w.instance)
}

// 设置客户区高度。
//
// Set client height.
func (w *TChart) SetClientHeight(value int32) {
	Chart_SetClientHeight(w.instance, value)
}

func (w *TChart) ClientOrigin() TPoint {
	return Chart_GetClientOrigin(w.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (w *TChart) ClientRect() TRect {
	return Chart_GetClientRect(w.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (w *TChart) ClientWidth() int32 {
	return Chart_GetClientWidth(w.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (w *TChart) SetClientWidth(value int32) {
	Chart_SetClientWidth(w.instance, value)
}

// 获取约束控件大小。
func (w *TChart) Constraints() *TSizeConstraints {
	return AsSizeConstraints(Chart_GetConstraints(w.instance))
}

// 设置约束控件大小。
func (w *TChart) SetConstraints(value *TSizeConstraints) {
	Chart_SetConstraints(w.instance, CheckPtr(value))
}

// 获取控件状态。
//
// Get control state.
func (w *TChart) ControlState() TControlState {
	return Chart_GetControlState(w.instance)
}

// 设置控件状态。
//
// Set control state.
func (w *TChart) SetControlState(value TControlState) {
	Chart_SetControlState(w.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (w *TChart) ControlStyle() TControlStyle {
	return Chart_GetControlStyle(w.instance)
}

// 设置控件样式。
//
// Set control style.
func (w *TChart) SetControlStyle(value TControlStyle) {
	Chart_SetControlStyle(w.instance, value)
}

func (w *TChart) Floating() bool {
	return Chart_GetFloating(w.instance)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (w *TChart) ShowHint() bool {
	return Chart_GetShowHint(w.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (w *TChart) SetShowHint(value bool) {
	Chart_SetShowHint(w.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (w *TChart) Visible() bool {
	return Chart_GetVisible(w.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (w *TChart) SetVisible(value bool) {
	Chart_SetVisible(w.instance, value)
}

// 获取控件父容器。
//
// Get control parent container.
func (w *TChart) Parent() *TWinControl {
	return AsWinControl(Chart_GetParent(w.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (w *TChart) SetParent(value IWinControl) {
	Chart_SetParent(w.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (w *TChart) Left() int32 {
	return Chart_GetLeft(w.instance)
}

// 设置左边位置。
//
// Set Left position.
func (w *TChart) SetLeft(value int32) {
	Chart_SetLeft(w.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (w *TChart) Top() int32 {
	return Chart_GetTop(w.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (w *TChart) SetTop(value int32) {
	Chart_SetTop(w.instance, value)
}

// 获取宽度。
//
// Get width.
func (w *TChart) Width() int32 {
	return Chart_GetWidth(w.instance)
}

// 设置宽度。
//
// Set width.
func (w *TChart) SetWidth(value int32) {
	Chart_SetWidth(w.instance, value)
}

// 获取高度。
//
// Get height.
func (w *TChart) Height() int32 {
	return Chart_GetHeight(w.instance)
}

// 设置高度。
//
// Set height.
func (w *TChart) SetHeight(value int32) {
	Chart_SetHeight(w.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (w *TChart) Cursor() TCursor {
	return Chart_GetCursor(w.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (w *TChart) SetCursor(value TCursor) {
	Chart_SetCursor(w.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (w *TChart) Hint() string {
	return Chart_GetHint(w.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (w *TChart) SetHint(value string) {
	Chart_SetHint(w.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (w *TChart) ComponentCount() int32 {
	return Chart_GetComponentCount(w.instance)
}

// 获取组件索引。
//
// Get component index.
func (w *TChart) ComponentIndex() int32 {
	return Chart_GetComponentIndex(w.instance)
}

// 设置组件索引。
//
// Set component index.
func (w *TChart) SetComponentIndex(value int32) {
	Chart_SetComponentIndex(w.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (w *TChart) Owner() *TComponent {
	return AsComponent(Chart_GetOwner(w.instance))
}

// 获取组件名称。
//
// Get the component name.
func (w *TChart) Name() string {
	return Chart_GetName(w.instance)
}

// 设置组件名称。
//
// Set the component name.
func (w *TChart) SetName(value string) {
	Chart_SetName(w.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (w *TChart) Tag() int {
	return Chart_GetTag(w.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (w *TChart) SetTag(value int) {
	Chart_SetTag(w.instance, value)
}

// 获取左边锚点。
func (w *TChart) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(Chart_GetAnchorSideLeft(w.instance))
}

// 设置左边锚点。
func (w *TChart) SetAnchorSideLeft(value *TAnchorSide) {
	Chart_SetAnchorSideLeft(w.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (w *TChart) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(Chart_GetAnchorSideTop(w.instance))
}

// 设置顶边锚点。
func (w *TChart) SetAnchorSideTop(value *TAnchorSide) {
	Chart_SetAnchorSideTop(w.instance, CheckPtr(value))
}

// 获取右边锚点。
func (w *TChart) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(Chart_GetAnchorSideRight(w.instance))
}

// 设置右边锚点。
func (w *TChart) SetAnchorSideRight(value *TAnchorSide) {
	Chart_SetAnchorSideRight(w.instance, CheckPtr(value))
}

// 获取底边锚点。
func (w *TChart) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(Chart_GetAnchorSideBottom(w.instance))
}

// 设置底边锚点。
func (w *TChart) SetAnchorSideBottom(value *TAnchorSide) {
	Chart_SetAnchorSideBottom(w.instance, CheckPtr(value))
}

func (w *TChart) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(Chart_GetChildSizing(w.instance))
}

func (w *TChart) SetChildSizing(value *TControlChildSizing) {
	Chart_SetChildSizing(w.instance, CheckPtr(value))
}

// 获取边框间距。
func (w *TChart) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(Chart_GetBorderSpacing(w.instance))
}

// 设置边框间距。
func (w *TChart) SetBorderSpacing(value *TControlBorderSpacing) {
	Chart_SetBorderSpacing(w.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (w *TChart) DockClients(Index int32) *TControl {
	return AsControl(Chart_GetDockClients(w.instance, Index))
}

// 获取指定索引子控件。
func (w *TChart) Controls(Index int32) *TControl {
	return AsControl(Chart_GetControls(w.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (w *TChart) Components(AIndex int32) *TComponent {
	return AsComponent(Chart_GetComponents(w.instance, AIndex))
}

// 获取锚侧面。
func (w *TChart) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(Chart_GetAnchorSide(w.instance, AKind))
}

func (w *TChart) Canvas() *TCanvas {
	return AsCanvas(Chart_GetCanvas(w.instance))
}

func (w *TChart) SetCanvas(ACanvas *TCanvas) {
	Chart_SetCanvas(w.instance, CheckPtr(ACanvas))
}

func (w *TChart) BorderStyle() TBorderStyle {
	return Chart_GetBorderStyle(w.instance)
}

func (w *TChart) SetBorderStyle(value TBorderStyle) {
	Chart_SetBorderStyle(w.instance, value)
}

func (w *TChart) SetOnPaint(event TNotifyEvent) {
	Chart_SetOnPaint(w.instance, event)
}

func (w *TChart) SetOnClick(event TNotifyEvent) {
	Chart_SetOnClick(w.instance, event)
}

func (w *TChart) SetOnContextPopup(event TContextPopupEvent) {
	Chart_SetOnContextPopup(w.instance, event)
}

func (w *TChart) SetOnDblClick(event TNotifyEvent) {
	Chart_SetOnDblClick(w.instance, event)
}

func (w *TChart) SetOnDragDrop(event TDragDropEvent) {
	Chart_SetOnDragDrop(w.instance, event)
}

func (w *TChart) SetOnDragOver(event TDragOverEvent) {
	Chart_SetOnDragOver(w.instance, event)
}

func (w *TChart) SetOnEndDrag(event TEndDragEvent) {
	Chart_SetOnEndDrag(w.instance, event)
}

func (w *TChart) SetOnMouseDown(event TMouseEvent) {
	Chart_SetOnMouseDown(w.instance, event)
}

func (w *TChart) SetOnMouseEnter(event TNotifyEvent) {
	Chart_SetOnMouseEnter(w.instance, event)
}

func (w *TChart) SetOnMouseLeave(event TNotifyEvent) {
	Chart_SetOnMouseLeave(w.instance, event)
}

func (w *TChart) SetOnMouseMove(event TMouseMoveEvent) {
	Chart_SetOnMouseMove(w.instance, event)
}

func (w *TChart) SetOnMouseUp(event TMouseEvent) {
	Chart_SetOnMouseUp(w.instance, event)
}

func (w *TChart) SetOnResize(event TNotifyEvent) {
	Chart_SetOnResize(w.instance, event)
}

func (w *TChart) AutoFocus() bool {
	return Chart_GetAutoFocus(w.instance)
}

func (w *TChart) SetAutoFocus(value bool) {
	Chart_SetAutoFocus(w.instance, value)
}

func (w *TChart) AllowPanning() bool {
	return Chart_GetAllowPanning(w.instance)
}

func (w *TChart) SetAllowPanning(value bool) {
	Chart_SetAllowPanning(w.instance, value)
}

func (w *TChart) AllowZoom() bool {
	return Chart_GetAllowZoom(w.instance)
}

func (w *TChart) SetAllowZoom(value bool) {
	Chart_SetAllowZoom(w.instance, value)
}

func (w *TChart) AntialiasingMode() TChartAntialiasingMode {
	return Chart_GetAntialiasingMode(w.instance)
}

func (w *TChart) SetAntialiasingMode(value TChartAntialiasingMode) {
	Chart_SetAntialiasingMode(w.instance, value)
}

func (w *TChart) AxisList() *TChartAxisList {
	return AsChartAxisList(Chart_GetAxisList(w.instance))
}

func (w *TChart) SetAxisList(value *TChartAxisList) {
	Chart_SetAxisList(w.instance, CheckPtr(value))
}

func (w *TChart) AxisVisible() bool {
	return Chart_GetAxisVisible(w.instance)
}

func (w *TChart) SetAxisVisible(value bool) {
	Chart_SetAxisVisible(w.instance, value)
}

func (w *TChart) BackColor() TColor {
	return Chart_GetBackColor(w.instance)
}

func (w *TChart) SetBackColor(value TColor) {
	Chart_SetBackColor(w.instance, value)
}

func (w *TChart) BottomAxis() *TChartAxis {
	return AsChartAxis(Chart_GetBottomAxis(w.instance))
}

func (w *TChart) SetBottomAxis(value *TChartAxis) {
	Chart_SetBottomAxis(w.instance, CheckPtr(value))
}

func (w *TChart) Depth() TChartDistance {
	return Chart_GetDepth(w.instance)
}

func (w *TChart) SetDepth(value TChartDistance) {
	Chart_SetDepth(w.instance, value)
}

func (w *TChart) ExpandPercentage() int32 {
	return Chart_GetExpandPercentage(w.instance)
}

func (w *TChart) SetExpandPercentage(value int32) {
	Chart_SetExpandPercentage(w.instance, value)
}

func (w *TChart) Extent() *TChartExtent {
	return AsChartExtent(Chart_GetExtent(w.instance))
}

func (w *TChart) SetExtent(value *TChartExtent) {
	Chart_SetExtent(w.instance, CheckPtr(value))
}

func (w *TChart) ExtentSizeLimit() *TChartExtent {
	return AsChartExtent(Chart_GetExtentSizeLimit(w.instance))
}

func (w *TChart) SetExtentSizeLimit(value *TChartExtent) {
	Chart_SetExtentSizeLimit(w.instance, CheckPtr(value))
}

func (w *TChart) Foot() *TChartTitle {
	return AsChartTitle(Chart_GetFoot(w.instance))
}

func (w *TChart) SetFoot(value *TChartTitle) {
	Chart_SetFoot(w.instance, CheckPtr(value))
}

func (w *TChart) Frame() *TChartPen {
	return AsChartPen(Chart_GetFrame(w.instance))
}

func (w *TChart) SetFrame(value *TChartPen) {
	Chart_SetFrame(w.instance, CheckPtr(value))
}

func (w *TChart) LeftAxis() *TChartAxis {
	return AsChartAxis(Chart_GetLeftAxis(w.instance))
}

func (w *TChart) SetLeftAxis(value *TChartAxis) {
	Chart_SetLeftAxis(w.instance, CheckPtr(value))
}

func (w *TChart) Legend() *TChartLegend {
	return AsChartLegend(Chart_GetLegend(w.instance))
}

func (w *TChart) SetLegend(value *TChartLegend) {
	Chart_SetLegend(w.instance, CheckPtr(value))
}

func (w *TChart) Margins() *TChartMargins {
	return AsChartMargins(Chart_GetMargins(w.instance))
}

func (w *TChart) SetMargins(value *TChartMargins) {
	Chart_SetMargins(w.instance, CheckPtr(value))
}

func (w *TChart) MarginsExternal() *TChartMargins {
	return AsChartMargins(Chart_GetMarginsExternal(w.instance))
}

func (w *TChart) SetMarginsExternal(value *TChartMargins) {
	Chart_SetMarginsExternal(w.instance, CheckPtr(value))
}

func (w *TChart) Proportional() bool {
	return Chart_GetProportional(w.instance)
}

func (w *TChart) SetProportional(value bool) {
	Chart_SetProportional(w.instance, value)
}

func (w *TChart) Series() *TChartSeriesList {
	return AsChartSeriesList(Chart_GetSeries(w.instance))
}

func (w *TChart) Title() *TChartTitle {
	return AsChartTitle(Chart_GetTitle(w.instance))
}

func (w *TChart) SetTitle(value *TChartTitle) {
	Chart_SetTitle(w.instance, CheckPtr(value))
}

func (w *TChart) Toolset() *TBasicChartToolset {
	return AsBasicChartToolset(Chart_GetToolset(w.instance))
}

func (w *TChart) SetToolset(value *TBasicChartToolset) {
	Chart_SetToolset(w.instance, CheckPtr(value))
}

func (w *TChart) SetOnAfterCustomDrawBackground(event TChartAfterCustomDrawEvent) {
	Chart_SetOnAfterCustomDrawBackground(w.instance, event)
}

func (w *TChart) SetOnAfterCustomDrawBackWall(event TChartAfterCustomDrawEvent) {
	Chart_SetOnAfterCustomDrawBackWall(w.instance, event)
}

func (w *TChart) SetOnAfterDraw(event TChartDrawEvent) {
	Chart_SetOnAfterDraw(w.instance, event)
}

func (w *TChart) SetOnAfterDrawBackground(event TChartAfterDrawEvent) {
	Chart_SetOnAfterDrawBackground(w.instance, event)
}

func (w *TChart) SetOnAfterDrawBackWall(event TChartAfterDrawEvent) {
	Chart_SetOnAfterDrawBackWall(w.instance, event)
}

func (w *TChart) SetOnAfterPaint(event TChartEvent) {
	Chart_SetOnAfterPaint(w.instance, event)
}

func (w *TChart) SetOnBeforeCustomDrawBackground(event TChartBeforeCustomDrawEvent) {
	Chart_SetOnBeforeCustomDrawBackground(w.instance, event)
}

func (w *TChart) SetOnBeforeDrawBackground(event TChartBeforeDrawEvent) {
	Chart_SetOnBeforeDrawBackground(w.instance, event)
}

func (w *TChart) SetOnBeforeCustomDrawBackWall(event TChartBeforeCustomDrawEvent) {
	Chart_SetOnBeforeCustomDrawBackWall(w.instance, event)
}

func (w *TChart) SetOnBeforeDrawBackWall(event TChartBeforeDrawEvent) {
	Chart_SetOnBeforeDrawBackWall(w.instance, event)
}

func (w *TChart) SetOnDrawLegend(event TChartDrawLegendEvent) {
	Chart_SetOnDrawLegend(w.instance, event)
}

func (w *TChart) SetOnExtentChanged(event TChartEvent) {
	Chart_SetOnExtentChanged(w.instance, event)
}

func (w *TChart) SetOnExtentChanging(event TChartEvent) {
	Chart_SetOnExtentChanging(w.instance, event)
}

func (w *TChart) SetOnExtentValidate(event TChartExtentValidateEvent) {
	Chart_SetOnExtentValidate(w.instance, event)
}

func (w *TChart) SetOnFullExtentChanged(event TChartEvent) {
	Chart_SetOnFullExtentChanged(w.instance, event)
}

func (w *TChart) ActiveToolIndex() int32 {
	return Chart_GetActiveToolIndex(w.instance)
}

func (w *TChart) ChartHeight() int32 {
	return Chart_GetChartHeight(w.instance)
}

func (w *TChart) ChartWidth() int32 {
	return Chart_GetChartWidth(w.instance)
}

func (w *TChart) ClipRect() TRect {
	return Chart_GetClipRect(w.instance)
}

func (w *TChart) HorAxis() *TChartAxis {
	return AsChartAxis(Chart_GetHorAxis(w.instance))
}

func (w *TChart) IsZoomed() bool {
	return Chart_GetIsZoomed(w.instance)
}

func (w *TChart) MinDataSpace() int32 {
	return Chart_GetMinDataSpace(w.instance)
}

func (w *TChart) SetMinDataSpace(value int32) {
	Chart_SetMinDataSpace(w.instance, value)
}

func (w *TChart) SetOnChartPaint(event TChartPaintEvent) {
	Chart_SetOnChartPaint(w.instance, event)
}

func (w *TChart) RenderingParams() *TChartRenderingParams {
	return AsChartRenderingParams(Chart_GetRenderingParams(w.instance))
}

func (w *TChart) SetRenderingParams(value *TChartRenderingParams) {
	Chart_SetRenderingParams(w.instance, CheckPtr(value))
}

func (w *TChart) SeriesCount() int32 {
	return Chart_GetSeriesCount(w.instance)
}

func (w *TChart) VertAxis() *TChartAxis {
	return AsChartAxis(Chart_GetVertAxis(w.instance))
}

func (w *TChart) XGraphMax() float64 {
	return Chart_GetXGraphMax(w.instance)
}

func (w *TChart) XGraphMin() float64 {
	return Chart_GetXGraphMin(w.instance)
}

func (w *TChart) YGraphMax() float64 {
	return Chart_GetYGraphMax(w.instance)
}

func (w *TChart) YGraphMin() float64 {
	return Chart_GetYGraphMin(w.instance)
}

func (w *TChart) ScalingValid() bool {
	return Chart_GetScalingValid(w.instance)
}

func (w *TChart) LockClipRect() {
	Chart_LockClipRect(w.instance)
}

func (w *TChart) UnlockClipRect() {
	Chart_UnlockClipRect(w.instance)
}

func (w *TChart) EraseBackground(DC HDC) {
	Chart_EraseBackground(w.instance, DC)
}

func (w *TChart) Paint() {
	Chart_Paint(w.instance)
}

func (w *TChart) SetChildOrder(child IComponent, order int32) {
	Chart_SetChildOrder(w.instance, CheckPtr(child), order)
}

func (w *TChart) DrawLineHoriz(drawer IChartDrawer, AY int32) {
	Chart_DrawLineHoriz(w.instance, CheckPtr(drawer), AY)
}

func (w *TChart) DrawLineVert(drawer IChartDrawer, AX int32) {
	Chart_DrawLineVert(w.instance, CheckPtr(drawer), AX)
}

func (w *TChart) AddSeries(ASeries IBasicChartSeries) {
	Chart_AddSeries(w.instance, CheckPtr(ASeries))
}

func (w *TChart) ClearSeries() {
	Chart_ClearSeries(w.instance)
}

func (w *TChart) Clone() *TChart {
	return AsChart(Chart_Clone(w.instance))
}

func (w *TChart) CopyToClipboard(AClass TRasterImageClass) {
	Chart_CopyToClipboard(w.instance, AClass)
}

func (w *TChart) CopyToClipboardBitmap() {
	Chart_CopyToClipboardBitmap(w.instance)
}

func (w *TChart) DeleteSeries(ASeries IBasicChartSeries) {
	Chart_DeleteSeries(w.instance, CheckPtr(ASeries))
}

func (w *TChart) DisableRedrawing() {
	Chart_DisableRedrawing(w.instance)
}

func (w *TChart) Draw(drawer IChartDrawer, rect TRect) {
	Chart_Draw(w.instance, CheckPtr(drawer), rect)
}

func (w *TChart) DrawLegendOn(canvas *TCanvas, rect *TRect) {
	Chart_DrawLegendOn(w.instance, CheckPtr(canvas), rect)
}

func (w *TChart) EnableRedrawing() {
	Chart_EnableRedrawing(w.instance)
}

func (w *TChart) GetAllSeriesAxisLimits(axis *TChartAxis, amin *float64, amax *float64) {
	Chart_GetAllSeriesAxisLimits(w.instance, CheckPtr(axis), amin, amax)
}

func (w *TChart) GetLegendItems(AIncludeHidden bool) *TChartLegendItems {
	return AsChartLegendItems(Chart_GetLegendItems(w.instance, AIncludeHidden))
}

func (w *TChart) PaintOnAuxCanvas(canvas *TCanvas, rect TRect) {
	Chart_PaintOnAuxCanvas(w.instance, CheckPtr(canvas), rect)
}

func (w *TChart) PaintOnCanvas(canvas *TCanvas, rect TRect) {
	Chart_PaintOnCanvas(w.instance, CheckPtr(canvas), rect)
}

func (w *TChart) Prepare() {
	Chart_Prepare(w.instance)
}

func (w *TChart) RemoveSeries(ASeries IBasicChartSeries) {
	Chart_RemoveSeries(w.instance, CheckPtr(ASeries))
}

func (w *TChart) SaveToBitmapFile(AFileName string) {
	Chart_SaveToBitmapFile(w.instance, AFileName)
}

func (w *TChart) SaveToFile(AClass TRasterImageClass, AFileName string) {
	Chart_SaveToFile(w.instance, AClass, AFileName)
}

func (w *TChart) StyleChanged(sender IObject) {
	Chart_StyleChanged(w.instance, CheckPtr(sender))
}

func (w *TChart) UsesBuiltinToolset() bool {
	return Chart_UsesBuiltinToolset(w.instance)
}

func (w *TChart) ZoomFull(AImmediateRecalc bool) {
	Chart_ZoomFull(w.instance, AImmediateRecalc)
}

func (w *TChart) XGraphToImage(AX float64) int32 {
	return Chart_XGraphToImage(w.instance, AX)
}

func (w *TChart) XImageToGraph(AX int32) float64 {
	return Chart_XImageToGraph(w.instance, AX)
}

func (w *TChart) YGraphToImage(AY float64) int32 {
	return Chart_YGraphToImage(w.instance, AY)
}

func (w *TChart) YImageToGraph(AY int32) float64 {
	return Chart_YImageToGraph(w.instance, AY)
}
