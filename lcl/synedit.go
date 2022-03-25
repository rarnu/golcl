package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynEdit struct {
	IWinControl
	instance uintptr
	ptr      unsafe.Pointer
}

func NewSynEdit(owner IComponent) *TSynEdit {
	s := new(TSynEdit)
	s.instance = SynEdit_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func AsSynEdit(obj any) *TSynEdit {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynEdit{instance: instance, ptr: ptr}
}

func SynEditFromInst(inst uintptr) *TSynEdit {
	return AsSynEdit(inst)
}

func SynEditFromObj(obj IObject) *TSynEdit {
	return AsSynEdit(obj)
}

func SynEditFromUnsafePointer(ptr unsafe.Pointer) *TSynEdit {
	return AsSynEdit(ptr)
}

func (s *TSynEdit) Free() {
	if s.instance != 0 {
		SynEdit_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynEdit) Instance() uintptr {
	return s.instance
}

func (s *TSynEdit) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynEdit) IsValid() bool {
	return s.instance != 0
}

func (s *TSynEdit) Is() TIs {
	return TIs(s.instance)
}

func TSynEditClass() TClass {
	return SynEdit_StaticClassType()
}

func (s *TSynEdit) Clear() {
	SynEdit_Clear(s.instance)
}

func (s *TSynEdit) ClearSelection() {
	SynEdit_ClearSelection(s.instance)
}

func (s *TSynEdit) CopyToClipboard() {
	SynEdit_CopyToClipboard(s.instance)
}

func (s *TSynEdit) CutToClipboard() {
	SynEdit_CutToClipboard(s.instance)
}

func (s *TSynEdit) PasteFromClipboard() {
	SynEdit_PasteFromClipboard(s.instance)
}

func (s *TSynEdit) Undo() {
	SynEdit_Undo(s.instance)
}

func (s *TSynEdit) Redo() {
	SynEdit_Redo(s.instance)
}

func (s *TSynEdit) SelectAll() {
	SynEdit_SelectAll(s.instance)
}

func (s *TSynEdit) CanFocus() bool {
	return SynEdit_CanFocus(s.instance)
}

func (s *TSynEdit) ContainsControl(Control IControl) bool {
	return SynEdit_ContainsControl(s.instance, CheckPtr(Control))
}

func (s *TSynEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool) *TControl {
	return AsControl(SynEdit_ControlAtPos(s.instance, Pos, AllowDisabled, AllowWinControls))
}

func (s *TSynEdit) DisableAlign() {
	SynEdit_DisableAlign(s.instance)
}

func (s *TSynEdit) EnableAlign() {
	SynEdit_EnableAlign(s.instance)
}

func (s *TSynEdit) FindChildControl(ControlName string) *TControl {
	return AsControl(SynEdit_FindChildControl(s.instance, ControlName))
}

func (s *TSynEdit) FlipChildren(AllLevels bool) {
	SynEdit_FlipChildren(s.instance, AllLevels)
}

func (s *TSynEdit) Focused() bool {
	return SynEdit_Focused(s.instance)
}

func (s *TSynEdit) HandleAllocated() bool {
	return SynEdit_HandleAllocated(s.instance)
}

func (s *TSynEdit) InsertControl(AControl IControl) {
	SynEdit_InsertControl(s.instance, CheckPtr(AControl))
}

func (s *TSynEdit) Invalidate() {
	SynEdit_Invalidate(s.instance)
}

func (s *TSynEdit) PaintTo(hDC HDC, X int32, Y int32) {
	SynEdit_PaintTo(s.instance, hDC, X, Y)
}

func (s *TSynEdit) RemoveControl(AControl IControl) {
	SynEdit_RemoveControl(s.instance, CheckPtr(AControl))
}

func (s *TSynEdit) Realign() {
	SynEdit_Realign(s.instance)
}

func (s *TSynEdit) Repaint() {
	SynEdit_Repaint(s.instance)
}

func (s *TSynEdit) ScaleBy(M int32, D int32) {
	SynEdit_ScaleBy(s.instance, M, D)
}

func (s *TSynEdit) ScrollBy(DeltaX int32, DeltaY int32) {
	SynEdit_ScrollBy(s.instance, DeltaX, DeltaY)
}

func (s *TSynEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	SynEdit_SetBounds(s.instance, ALeft, ATop, AWidth, AHeight)
}

func (s *TSynEdit) SetFocus() {
	SynEdit_SetFocus(s.instance)
}

func (s *TSynEdit) Update() {
	SynEdit_Update(s.instance)
}

func (s *TSynEdit) BringToFront() {
	SynEdit_BringToFront(s.instance)
}

func (s *TSynEdit) ClientToScreen(Point TPoint) TPoint {
	return SynEdit_ClientToScreen(s.instance, Point)
}

func (s *TSynEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return SynEdit_ClientToParent(s.instance, Point, CheckPtr(AParent))
}

func (s *TSynEdit) Dragging() bool {
	return SynEdit_Dragging(s.instance)
}

func (s *TSynEdit) HasParent() bool {
	return SynEdit_HasParent(s.instance)
}

func (s *TSynEdit) Hide() {
	SynEdit_Hide(s.instance)
}

func (s *TSynEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return SynEdit_Perform(s.instance, Msg, WParam, LParam)
}

func (s *TSynEdit) Refresh() {
	SynEdit_Refresh(s.instance)
}

func (s *TSynEdit) ScreenToClient(Point TPoint) TPoint {
	return SynEdit_ScreenToClient(s.instance, Point)
}

func (s *TSynEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return SynEdit_ParentToClient(s.instance, Point, CheckPtr(AParent))
}

func (s *TSynEdit) SendToBack() {
	SynEdit_SendToBack(s.instance)
}

func (s *TSynEdit) Show() {
	SynEdit_Show(s.instance)
}

func (s *TSynEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return SynEdit_GetTextBuf(s.instance, Buffer, BufSize)
}

func (s *TSynEdit) GetTextLen() int32 {
	return SynEdit_GetTextLen(s.instance)
}

func (s *TSynEdit) SetTextBuf(Buffer string) {
	SynEdit_SetTextBuf(s.instance, Buffer)
}

func (s *TSynEdit) FindComponent(AName string) *TComponent {
	return AsComponent(SynEdit_FindComponent(s.instance, AName))
}

func (s *TSynEdit) GetNamePath() string {
	return SynEdit_GetNamePath(s.instance)
}

func (s *TSynEdit) Assign(Source IObject) {
	SynEdit_Assign(s.instance, CheckPtr(Source))
}

func (s *TSynEdit) ClassType() TClass {
	return SynEdit_ClassType(s.instance)
}

func (s *TSynEdit) ClassName() string {
	return SynEdit_ClassName(s.instance)
}

func (s *TSynEdit) InstanceSize() int32 {
	return SynEdit_InstanceSize(s.instance)
}

func (s *TSynEdit) InheritsFrom(AClass TClass) bool {
	return SynEdit_InheritsFrom(s.instance, AClass)
}

func (s *TSynEdit) Equals(Obj IObject) bool {
	return SynEdit_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSynEdit) GetHashCode() int32 {
	return SynEdit_GetHashCode(s.instance)
}

func (s *TSynEdit) ToString() string {
	return SynEdit_ToString(s.instance)
}

func (s *TSynEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	SynEdit_AnchorToNeighbour(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (s *TSynEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	SynEdit_AnchorParallel(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (s *TSynEdit) AnchorHorizontalCenterTo(ASibling IControl) {
	SynEdit_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TSynEdit) AnchorVerticalCenterTo(ASibling IControl) {
	SynEdit_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TSynEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	SynEdit_AnchorSame(s.instance, ASide, CheckPtr(ASibling))
}

func (s *TSynEdit) AnchorAsAlign(AAlign TAlign, ASpace int32) {
	SynEdit_AnchorAsAlign(s.instance, AAlign, ASpace)
}

func (s *TSynEdit) AnchorClient(ASpace int32) {
	SynEdit_AnchorClient(s.instance, ASpace)
}

func (s *TSynEdit) ScaleDesignToForm(ASize int32) int32 {
	return SynEdit_ScaleDesignToForm(s.instance, ASize)
}

func (s *TSynEdit) ScaleFormToDesign(ASize int32) int32 {
	return SynEdit_ScaleFormToDesign(s.instance, ASize)
}

func (s *TSynEdit) Scale96ToForm(ASize int32) int32 {
	return SynEdit_Scale96ToForm(s.instance, ASize)
}

func (s *TSynEdit) ScaleFormTo96(ASize int32) int32 {
	return SynEdit_ScaleFormTo96(s.instance, ASize)
}

func (s *TSynEdit) Scale96ToFont(ASize int32) int32 {
	return SynEdit_Scale96ToFont(s.instance, ASize)
}

func (s *TSynEdit) ScaleFontTo96(ASize int32) int32 {
	return SynEdit_ScaleFontTo96(s.instance, ASize)
}

func (s *TSynEdit) ScaleScreenToFont(ASize int32) int32 {
	return SynEdit_ScaleScreenToFont(s.instance, ASize)
}

func (s *TSynEdit) ScaleFontToScreen(ASize int32) int32 {
	return SynEdit_ScaleFontToScreen(s.instance, ASize)
}

func (s *TSynEdit) Scale96ToScreen(ASize int32) int32 {
	return SynEdit_Scale96ToScreen(s.instance, ASize)
}

func (s *TSynEdit) ScaleScreenTo96(ASize int32) int32 {
	return SynEdit_ScaleScreenTo96(s.instance, ASize)
}

func (s *TSynEdit) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, AOldFormHeight int32) {
	SynEdit_AutoAdjustLayout(s.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, AOldFormHeight)
}

func (s *TSynEdit) FixDesignFontsPPI(APPI int32) {
	SynEdit_FixDesignFontsPPI(s.instance, APPI)
}

func (s *TSynEdit) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	SynEdit_ScaleFontsPPI(s.instance, AToPPI, AProportion)
}

func (s *TSynEdit) Align() TAlign {
	return SynEdit_GetAlign(s.instance)
}

func (s *TSynEdit) SetAlign(value TAlign) {
	SynEdit_SetAlign(s.instance, value)
}

func (s *TSynEdit) Anchors() TAnchors {
	return SynEdit_GetAnchors(s.instance)
}

func (s *TSynEdit) SetAnchors(value TAnchors) {
	SynEdit_SetAnchors(s.instance, value)
}

func (s *TSynEdit) BiDiMode() TBiDiMode {
	return SynEdit_GetBiDiMode(s.instance)
}

func (s *TSynEdit) SetBiDiMode(value TBiDiMode) {
	SynEdit_SetBiDiMode(s.instance, value)
}

func (s *TSynEdit) BorderStyle() TBorderStyle {
	return SynEdit_GetBorderStyle(s.instance)
}

func (s *TSynEdit) SetBorderStyle(value TBorderStyle) {
	SynEdit_SetBorderStyle(s.instance, value)
}

func (s *TSynEdit) Color() TColor {
	return SynEdit_GetColor(s.instance)
}

func (s *TSynEdit) SetColor(value TColor) {
	SynEdit_SetColor(s.instance, value)
}

func (s *TSynEdit) Constraints() *TSizeConstraints {
	return AsSizeConstraints(SynEdit_GetConstraints(s.instance))
}

func (s *TSynEdit) SetConstraints(value *TSizeConstraints) {
	SynEdit_SetConstraints(s.instance, CheckPtr(value))
}

func (s *TSynEdit) DoubleBuffered() bool {
	return SynEdit_GetDoubleBuffered(s.instance)
}

func (s *TSynEdit) SetDoubleBuffered(value bool) {
	SynEdit_SetDoubleBuffered(s.instance, value)
}

func (s *TSynEdit) Enabled() bool {
	return SynEdit_GetEnabled(s.instance)
}

func (s *TSynEdit) SetEnabled(value bool) {
	SynEdit_SetEnabled(s.instance, value)
}

func (s *TSynEdit) Font() *TFont {
	return AsFont(SynEdit_GetFont(s.instance))
}

func (s *TSynEdit) SetFont(value *TFont) {
	SynEdit_SetFont(s.instance, CheckPtr(value))
}

func (s *TSynEdit) HideSelection() bool {
	return SynEdit_GetHideSelection(s.instance)
}

func (s *TSynEdit) SetHideSelection(value bool) {
	SynEdit_SetHideSelection(s.instance, value)
}

func (s *TSynEdit) Lines() *TStrings {
	return AsStrings(SynEdit_GetLines(s.instance))
}

func (s *TSynEdit) SetLines(value *IStrings) {
	SynEdit_SetLines(s.instance, CheckPtr(value))
}

func (s *TSynEdit) ParentColor() bool {
	return SynEdit_GetParentColor(s.instance)
}

func (s *TSynEdit) SetParentColor(value bool) {
	SynEdit_SetParentColor(s.instance, value)
}

func (s *TSynEdit) ParentDoubleBuffered() bool {
	return SynEdit_GetParentDoubleBuffered(s.instance)
}

func (s *TSynEdit) SetParentDoubleBuffered(value bool) {
	SynEdit_SetParentDoubleBuffered(s.instance, value)
}

func (s *TSynEdit) ParentFont() bool {
	return SynEdit_GetParentFont(s.instance)
}

func (s *TSynEdit) SetParentFont(value bool) {
	SynEdit_SetParentFont(s.instance, value)
}

func (s *TSynEdit) ParentShowHint() bool {
	return SynEdit_GetParentShowHint(s.instance)
}

func (s *TSynEdit) SetParentShowHint(value bool) {
	SynEdit_SetParentShowHint(s.instance, value)
}

func (s *TSynEdit) PopupMenu() *TPopupMenu {
	return AsPopupMenu(SynEdit_GetPopupMenu(s.instance))
}

func (s *TSynEdit) SetPopupMenu(value IComponent) {
	SynEdit_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TSynEdit) ReadOnly() bool {
	return SynEdit_GetReadOnly(s.instance)
}

func (s *TSynEdit) SetReadOnly(value bool) {
	SynEdit_SetReadOnly(s.instance, value)
}

func (s *TSynEdit) ScrollBars() TScrollStyle {
	return SynEdit_GetScrollBars(s.instance)
}

func (s *TSynEdit) SetScrollBars(value TScrollStyle) {
	SynEdit_SetScrollBars(s.instance, value)
}

func (s *TSynEdit) ShowHint() bool {
	return SynEdit_GetShowHint(s.instance)
}

func (s *TSynEdit) SetShowHint(value bool) {
	SynEdit_SetShowHint(s.instance, value)
}

func (s *TSynEdit) TabOrder() TTabOrder {
	return SynEdit_GetTabOrder(s.instance)
}

func (s *TSynEdit) SetTabOrder(value TTabOrder) {
	SynEdit_SetTabOrder(s.instance, value)
}

func (s *TSynEdit) TabStop() bool {
	return SynEdit_GetTabStop(s.instance)
}

func (s *TSynEdit) SetTabStop(value bool) {
	SynEdit_SetTabStop(s.instance, value)
}

func (s *TSynEdit) Visible() bool {
	return SynEdit_GetVisible(s.instance)
}

func (s *TSynEdit) SetVisible(value bool) {
	SynEdit_SetVisible(s.instance, value)
}

func (s *TSynEdit) WantTabs() bool {
	return SynEdit_GetWantTabs(s.instance)
}

func (s *TSynEdit) SetWantTabs(value bool) {
	SynEdit_SetWantTabs(s.instance, value)
}

func (s *TSynEdit) SetOnChange(fn TNotifyEvent) {
	SynEdit_SetOnChange(s.instance, fn)
}

func (s *TSynEdit) SetOnClick(fn TNotifyEvent) {
	SynEdit_SetOnClick(s.instance, fn)
}

func (s *TSynEdit) SetOnDblClick(fn TNotifyEvent) {
	SynEdit_SetOnDblClick(s.instance, fn)
}

func (s *TSynEdit) SetOnContextPopup(fn TContextPopupEvent) {
	SynEdit_SetOnContextPopup(s.instance, fn)
}

func (s *TSynEdit) SetOnDragDrop(fn TDragDropEvent) {
	SynEdit_SetOnDragDrop(s.instance, fn)
}

func (s *TSynEdit) SetOnDragOver(fn TDragOverEvent) {
	SynEdit_SetOnDragOver(s.instance, fn)
}

func (s *TSynEdit) SetOnEndDrag(fn TEndDragEvent) {
	SynEdit_SetOnEndDrag(s.instance, fn)
}

func (s *TSynEdit) SetOnEnter(fn TNotifyEvent) {
	SynEdit_SetOnEnter(s.instance, fn)
}

func (s *TSynEdit) SetOnExit(fn TNotifyEvent) {
	SynEdit_SetOnExit(s.instance, fn)
}

func (s *TSynEdit) SetOnKeyDown(fn TKeyEvent) {
	SynEdit_SetOnKeyDown(s.instance, fn)
}

func (s *TSynEdit) SetOnKeyPress(fn TKeyPressEvent) {
	SynEdit_SetOnKeyPress(s.instance, fn)
}

func (s *TSynEdit) SetOnKeyUp(fn TKeyEvent) {
	SynEdit_SetOnKeyUp(s.instance, fn)
}

func (s *TSynEdit) SetOnMouseDown(fn TMouseEvent) {
	SynEdit_SetOnMouseDown(s.instance, fn)
}

func (s *TSynEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	SynEdit_SetOnMouseMove(s.instance, fn)
}

func (s *TSynEdit) SetOnMouseUp(fn TMouseEvent) {
	SynEdit_SetOnMouseUp(s.instance, fn)
}

func (s *TSynEdit) SetOnMouseEnter(fn TNotifyEvent) {
	SynEdit_SetOnMouseEnter(s.instance, fn)
}

func (s *TSynEdit) SetOnMouseLeave(fn TNotifyEvent) {
	SynEdit_SetOnMouseLeave(s.instance, fn)
}

func (s *TSynEdit) CanUndo() bool {
	return SynEdit_GetCanUndo(s.instance)
}

func (s *TSynEdit) CanRedo() bool {
	return SynEdit_GetCanRedo(s.instance)
}

func (s *TSynEdit) Modified() bool {
	return SynEdit_GetModified(s.instance)
}

func (s *TSynEdit) SetModified(value bool) {
	SynEdit_SetModified(s.instance, value)
}

func (s *TSynEdit) SelStart() int32 {
	return SynEdit_GetSelStart(s.instance)
}

func (s *TSynEdit) SetSelStart(value int32) {
	SynEdit_SetSelStart(s.instance, value)
}

func (s *TSynEdit) SelText() string {
	return SynEdit_GetSelText(s.instance)
}

func (s *TSynEdit) SetSelText(value string) {
	SynEdit_SetSelText(s.instance, value)
}

func (s *TSynEdit) Text() string {
	return SynEdit_GetText(s.instance)
}

func (s *TSynEdit) SetText(value string) {
	SynEdit_SetText(s.instance, value)
}

func (s *TSynEdit) DockClientCount() int32 {
	return SynEdit_GetDockClientCount(s.instance)
}

func (s *TSynEdit) DockSite() bool {
	return SynEdit_GetDockSite(s.instance)
}

func (s *TSynEdit) SetDockSite(value bool) {
	SynEdit_SetDockSite(s.instance, value)
}

func (s *TSynEdit) MouseInClient() bool {
	return SynEdit_GetMouseInClient(s.instance)
}

func (s *TSynEdit) VisibleDockClientCount() int32 {
	return SynEdit_GetVisibleDockClientCount(s.instance)
}

func (s *TSynEdit) Brush() *TBrush {
	return AsBrush(SynEdit_GetBrush(s.instance))
}

func (s *TSynEdit) ControlCount() int32 {
	return SynEdit_GetControlCount(s.instance)
}

func (s *TSynEdit) Handle() HWND {
	return SynEdit_GetHandle(s.instance)
}

func (s *TSynEdit) ParentWindow() HWND {
	return SynEdit_GetParentWindow(s.instance)
}

func (s *TSynEdit) SetParentWindow(value HWND) {
	SynEdit_SetParentWindow(s.instance, value)
}

func (s *TSynEdit) Showing() bool {
	return SynEdit_GetShowing(s.instance)
}

func (s *TSynEdit) UseDockManager() bool {
	return SynEdit_GetUseDockManager(s.instance)
}

func (s *TSynEdit) SetUseDockManager(value bool) {
	SynEdit_SetUseDockManager(s.instance, value)
}

func (s *TSynEdit) Action() *TAction {
	return AsAction(SynEdit_GetAction(s.instance))
}

func (s *TSynEdit) SetAction(value IComponent) {
	SynEdit_SetAction(s.instance, CheckPtr(value))
}

func (s *TSynEdit) BoundsRect() TRect {
	return SynEdit_GetBoundsRect(s.instance)
}

func (s *TSynEdit) SetBoundsRect(value TRect) {
	SynEdit_SetBoundsRect(s.instance, value)
}

func (s *TSynEdit) ClientHeight() int32 {
	return SynEdit_GetClientHeight(s.instance)
}

func (s *TSynEdit) SetClientHeight(value int32) {
	SynEdit_SetClientHeight(s.instance, value)
}

func (s *TSynEdit) ClientOrigin() TPoint {
	return SynEdit_GetClientOrigin(s.instance)
}

func (s *TSynEdit) ClientRect() TRect {
	return SynEdit_GetClientRect(s.instance)
}

func (s *TSynEdit) ClientWidth() int32 {
	return SynEdit_GetClientWidth(s.instance)
}

func (s *TSynEdit) SetClientWidth(value int32) {
	SynEdit_SetClientWidth(s.instance, value)
}

func (s *TSynEdit) ControlState() TControlState {
	return SynEdit_GetControlState(s.instance)
}

func (s *TSynEdit) SetControlState(value TControlState) {
	SynEdit_SetControlState(s.instance, value)
}

func (s *TSynEdit) ControlStyle() TControlStyle {
	return SynEdit_GetControlStyle(s.instance)
}

func (s *TSynEdit) SetControlStyle(value TControlStyle) {
	SynEdit_SetControlStyle(s.instance, value)
}

func (s *TSynEdit) Floating() bool {
	return SynEdit_GetFloating(s.instance)
}

func (s *TSynEdit) Parent() *TWinControl {
	return AsWinControl(SynEdit_GetParent(s.instance))
}

func (s *TSynEdit) SetParent(value IWinControl) {
	SynEdit_SetParent(s.instance, CheckPtr(value))
}

func (s *TSynEdit) Left() int32 {
	return SynEdit_GetLeft(s.instance)
}

func (s *TSynEdit) SetLeft(value int32) {
	SynEdit_SetLeft(s.instance, value)
}

func (s *TSynEdit) Top() int32 {
	return SynEdit_GetTop(s.instance)
}

func (s *TSynEdit) SetTop(value int32) {
	SynEdit_SetTop(s.instance, value)
}

func (s *TSynEdit) Width() int32 {
	return SynEdit_GetWidth(s.instance)
}

func (s *TSynEdit) SetWidth(value int32) {
	SynEdit_SetWidth(s.instance, value)
}

func (s *TSynEdit) Height() int32 {
	return SynEdit_GetHeight(s.instance)
}

func (s *TSynEdit) SetHeight(value int32) {
	SynEdit_SetHeight(s.instance, value)
}

func (s *TSynEdit) Cursor() TCursor {
	return SynEdit_GetCursor(s.instance)
}

func (s *TSynEdit) SetCursor(value TCursor) {
	SynEdit_SetCursor(s.instance, value)
}

func (s *TSynEdit) Hint() string {
	return SynEdit_GetHint(s.instance)
}

func (s *TSynEdit) SetHint(value string) {
	SynEdit_SetHint(s.instance, value)
}

func (s *TSynEdit) ComponentCount() int32 {
	return SynEdit_GetComponentCount(s.instance)
}

func (s *TSynEdit) ComponentIndex() int32 {
	return SynEdit_GetComponentIndex(s.instance)
}

func (s *TSynEdit) SetComponentIndex(value int32) {
	SynEdit_SetComponentIndex(s.instance, value)
}

func (s *TSynEdit) Owner() *TComponent {
	return AsComponent(SynEdit_GetOwner(s.instance))
}

func (s *TSynEdit) Name() string {
	return SynEdit_GetName(s.instance)
}

func (s *TSynEdit) SetName(value string) {
	SynEdit_SetName(s.instance, value)
}

func (s *TSynEdit) Tag() int {
	return SynEdit_GetTag(s.instance)
}

func (s *TSynEdit) SetTag(value int) {
	SynEdit_SetTag(s.instance, value)
}

func (s *TSynEdit) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(SynEdit_GetAnchorSideLeft(s.instance))
}

func (s *TSynEdit) SetAnchorSideLeft(value *TAnchorSide) {
	SynEdit_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

func (s *TSynEdit) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(SynEdit_GetAnchorSideTop(s.instance))
}

func (s *TSynEdit) SetAnchorSideTop(value *TAnchorSide) {
	SynEdit_SetAnchorSideTop(s.instance, CheckPtr(value))
}

func (s *TSynEdit) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(SynEdit_GetAnchorSideRight(s.instance))
}

func (s *TSynEdit) SetAnchorSideRight(value *TAnchorSide) {
	SynEdit_SetAnchorSideRight(s.instance, CheckPtr(value))
}

func (s *TSynEdit) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(SynEdit_GetAnchorSideBottom(s.instance))
}

func (s *TSynEdit) SetAnchorSideBottom(value *TAnchorSide) {
	SynEdit_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TSynEdit) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(SynEdit_GetChildSizing(s.instance))
}

func (s *TSynEdit) SetChildSizing(value *TControlChildSizing) {
	SynEdit_SetChildSizing(s.instance, CheckPtr(value))
}

func (s *TSynEdit) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(SynEdit_GetBorderSpacing(s.instance))
}

func (s *TSynEdit) SetBorderSpacing(value *TControlBorderSpacing) {
	SynEdit_SetBorderSpacing(s.instance, CheckPtr(value))
}

func (s *TSynEdit) DockClients(index int32) *TControl {
	return AsControl(SynEdit_GetDockClients(s.instance, index))
}

func (s *TSynEdit) Controls(index int32) *TControl {
	return AsControl(SynEdit_GetControls(s.instance, index))
}

func (s *TSynEdit) Components(index int32) *TComponent {
	return AsComponent(SynEdit_GetComponents(s.instance, index))
}

func (s *TSynEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(SynEdit_GetAnchorSide(s.instance, AKind))
}

func (s *TSynEdit) BlockIndent() int32 {
	return SynEdit_GetBlockIndent(s.instance)
}

func (s *TSynEdit) SetBlockIndent(value int32) {
	SynEdit_SetBlockIndent(s.instance, value)
}

func (s *TSynEdit) SetOnMouseLink(fn TOnMouseLinkEvent) {
	SynEdit_SetOnMouseLink(s.instance, fn)
}

func (s *TSynEdit) SetOnClickLink(fn TMouseEvent) {
	SynEdit_SetOnClickLink(s.instance, fn)
}

func (s *TSynEdit) BlockTabIndent() int32 {
	return SynEdit_GetBlockTabIndent(s.instance)
}

func (s *TSynEdit) SetBlockTabIndent(value int32) {
	SynEdit_SetBlockTabIndent(s.instance, value)
}

func (s *TSynEdit) RightEdge() int32 {
	return SynEdit_GetRightEdge(s.instance)
}

func (s *TSynEdit) SetRightEdge(value int32) {
	SynEdit_SetRightEdge(s.instance, value)
}

func (s *TSynEdit) RightEdgeColor() TColor {
	return SynEdit_GetRightEdgeColor(s.instance)
}

func (s *TSynEdit) SetRightEdgeColor(value TColor) {
	SynEdit_SetRightEdgeColor(s.instance, value)
}

func (s *TSynEdit) TabWidth() int32 {
	return SynEdit_GetTabWidth(s.instance)
}

func (s *TSynEdit) SetTabWidth(value int32) {
	SynEdit_SetTabWidth(s.instance, value)
}

func (s *TSynEdit) Options() TSynEditorOptions {
	return SynEdit_GetOptions(s.instance)
}

func (s *TSynEdit) SetOptions(value TSynEditorOptions) {
	SynEdit_SetOptions(s.instance, value)
}

func (s *TSynEdit) Options2() TSynEditorOptions2 {
	return SynEdit_GetOptions2(s.instance)
}

func (s *TSynEdit) SetOptions2(value TSynEditorOptions2) {
	SynEdit_SetOptions2(s.instance, value)
}

func (s *TSynEdit) MouseOptions() TSynEditorMouseOptions {
	return SynEdit_GetMouseOptions(s.instance)
}

func (s *TSynEdit) SetMouseOptions(value TSynEditorMouseOptions) {
	SynEdit_SetMouseOptions(s.instance, value)
}

func (s *TSynEdit) VisibleSpecialChars() TSynVisibleSpecialChars {
	return SynEdit_GetVisibleSpecialChars(s.instance)
}

func (s *TSynEdit) SetVisibleSpecialChars(value TSynVisibleSpecialChars) {
	SynEdit_SetVisibleSpecialChars(s.instance, value)
}

func (s *TSynEdit) GetHighlighter() *TSynCustomHighlighter {
	return AsSynCustomHighlighter(SynEdit_GetHighlighter(s.instance))
}

func (s *TSynEdit) SetHighlighter(value *TSynCustomHighlighter) {
	SynEdit_SetHighlighter(s.instance, CheckPtr(value))
}

func (s *TSynEdit) GetGutter() *TSynGutter {
	return AsSynGutter(SynEdit_GetGutter(s.instance))
}

func (s *TSynEdit) GetRightGutter() *TSynGutter {
	return AsSynGutter(SynEdit_GetRightGutter(s.instance))
}
