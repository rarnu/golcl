package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- THeaderControl ---------------------------

func HeaderControl_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_Create").Call(obj)
	return ret
}

func HeaderControl_Free(obj uintptr) {
	getLazyProc("HeaderControl_Free").Call(obj)
}

func HeaderControl_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("HeaderControl_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func HeaderControl_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func HeaderControl_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func HeaderControl_DisableAlign(obj uintptr) {
	getLazyProc("HeaderControl_DisableAlign").Call(obj)
}

func HeaderControl_EnableAlign(obj uintptr) {
	getLazyProc("HeaderControl_EnableAlign").Call(obj)
}

func HeaderControl_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func HeaderControl_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("HeaderControl_InsertControl").Call(obj, AControl)
}

func HeaderControl_Invalidate(obj uintptr) {
	getLazyProc("HeaderControl_Invalidate").Call(obj)
}

func HeaderControl_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("HeaderControl_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func HeaderControl_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("HeaderControl_RemoveControl").Call(obj, AControl)
}

func HeaderControl_Realign(obj uintptr) {
	getLazyProc("HeaderControl_Realign").Call(obj)
}

func HeaderControl_Repaint(obj uintptr) {
	getLazyProc("HeaderControl_Repaint").Call(obj)
}

func HeaderControl_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("HeaderControl_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func HeaderControl_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("HeaderControl_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func HeaderControl_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("HeaderControl_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func HeaderControl_SetFocus(obj uintptr) {
	getLazyProc("HeaderControl_SetFocus").Call(obj)
}

func HeaderControl_Update(obj uintptr) {
	getLazyProc("HeaderControl_Update").Call(obj)
}

func HeaderControl_BringToFront(obj uintptr) {
	getLazyProc("HeaderControl_BringToFront").Call(obj)
}

func HeaderControl_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("HeaderControl_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("HeaderControl_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_Hide(obj uintptr) {
	getLazyProc("HeaderControl_Hide").Call(obj)
}

func HeaderControl_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("HeaderControl_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func HeaderControl_Refresh(obj uintptr) {
	getLazyProc("HeaderControl_Refresh").Call(obj)
}

func HeaderControl_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("HeaderControl_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("HeaderControl_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_SendToBack(obj uintptr) {
	getLazyProc("HeaderControl_SendToBack").Call(obj)
}

func HeaderControl_Show(obj uintptr) {
	getLazyProc("HeaderControl_Show").Call(obj)
}

func HeaderControl_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("HeaderControl_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func HeaderControl_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetTextLen").Call(obj)
	return int32(ret)
}

func HeaderControl_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("HeaderControl_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func HeaderControl_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func HeaderControl_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderControl_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderControl_Assign(obj uintptr, Source uintptr) {
	getLazyProc("HeaderControl_Assign").Call(obj, Source)
}

func HeaderControl_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("HeaderControl_ClassType").Call(obj)
	return TClass(ret)
}

func HeaderControl_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderControl_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderControl_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_InstanceSize").Call(obj)
	return int32(ret)
}

func HeaderControl_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("HeaderControl_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func HeaderControl_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetHashCode").Call(obj)
	return int32(ret)
}

func HeaderControl_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderControl_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderControl_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("HeaderControl_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func HeaderControl_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("HeaderControl_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func HeaderControl_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("HeaderControl_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func HeaderControl_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("HeaderControl_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func HeaderControl_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("HeaderControl_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func HeaderControl_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("HeaderControl_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func HeaderControl_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("HeaderControl_AnchorClient").Call(obj, uintptr(ASpace))
}

func HeaderControl_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("HeaderControl_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func HeaderControl_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("HeaderControl_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func HeaderControl_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("HeaderControl_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func HeaderControl_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("HeaderControl_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func HeaderControl_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("HeaderControl_GetAlign").Call(obj)
	return TAlign(ret)
}

func HeaderControl_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("HeaderControl_SetAlign").Call(obj, uintptr(value))
}

func HeaderControl_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func HeaderControl_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("HeaderControl_SetAnchors").Call(obj, uintptr(value))
}

func HeaderControl_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetBorderWidth").Call(obj)
	return int32(ret)
}

func HeaderControl_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetBorderWidth").Call(obj, uintptr(value))
}

func HeaderControl_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("HeaderControl_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func HeaderControl_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("HeaderControl_SetDragCursor").Call(obj, uintptr(value))
}

func HeaderControl_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("HeaderControl_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func HeaderControl_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("HeaderControl_SetDragKind").Call(obj, uintptr(value))
}

func HeaderControl_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("HeaderControl_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func HeaderControl_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("HeaderControl_SetDragMode").Call(obj, uintptr(value))
}

func HeaderControl_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetEnabled(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetFont").Call(obj)
	return ret
}

func HeaderControl_SetFont(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetFont").Call(obj, value)
}

func HeaderControl_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetImages").Call(obj)
	return ret
}

func HeaderControl_SetImages(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetImages").Call(obj, value)
}

func HeaderControl_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetConstraints").Call(obj)
	return ret
}

func HeaderControl_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetConstraints").Call(obj, value)
}

func HeaderControl_GetSections(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetSections").Call(obj)
	return ret
}

func HeaderControl_SetSections(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetSections").Call(obj, value)
}

func HeaderControl_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetShowHint(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetParentFont(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetPopupMenu").Call(obj)
	return ret
}

func HeaderControl_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetPopupMenu").Call(obj, value)
}

func HeaderControl_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetVisible(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetVisible").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnResize(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnSectionClick(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnSectionClick").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnSectionResize(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnSectionResize").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnSectionTrack(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnSectionTrack").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnSectionDrag(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnSectionDrag").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_SetOnSectionEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("HeaderControl_SetOnSectionEndDrag").Call(obj, addEventToMap(obj, fn))
}

func HeaderControl_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetCanvas").Call(obj)
	return ret
}

func HeaderControl_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetDockClientCount").Call(obj)
	return int32(ret)
}

func HeaderControl_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetDockSite(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func HeaderControl_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetBrush").Call(obj)
	return ret
}

func HeaderControl_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetControlCount").Call(obj)
	return int32(ret)
}

func HeaderControl_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("HeaderControl_GetHandle").Call(obj)
	return HWND(ret)
}

func HeaderControl_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("HeaderControl_GetParentWindow").Call(obj)
	return HWND(ret)
}

func HeaderControl_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("HeaderControl_SetParentWindow").Call(obj, uintptr(value))
}

func HeaderControl_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("HeaderControl_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func HeaderControl_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("HeaderControl_SetTabOrder").Call(obj, uintptr(value))
}

func HeaderControl_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetTabStop(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("HeaderControl_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func HeaderControl_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAction").Call(obj)
	return ret
}

func HeaderControl_SetAction(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetAction").Call(obj, value)
}

func HeaderControl_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("HeaderControl_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("HeaderControl_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func HeaderControl_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetClientHeight").Call(obj)
	return int32(ret)
}

func HeaderControl_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetClientHeight").Call(obj, uintptr(value))
}

func HeaderControl_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("HeaderControl_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("HeaderControl_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func HeaderControl_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetClientWidth").Call(obj)
	return int32(ret)
}

func HeaderControl_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetClientWidth").Call(obj, uintptr(value))
}

func HeaderControl_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("HeaderControl_GetControlState").Call(obj)
	return TControlState(ret)
}

func HeaderControl_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("HeaderControl_SetControlState").Call(obj, uintptr(value))
}

func HeaderControl_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("HeaderControl_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func HeaderControl_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("HeaderControl_SetControlStyle").Call(obj, uintptr(value))
}

func HeaderControl_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("HeaderControl_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func HeaderControl_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetParent").Call(obj)
	return ret
}

func HeaderControl_SetParent(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetParent").Call(obj, value)
}

func HeaderControl_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetLeft").Call(obj)
	return int32(ret)
}

func HeaderControl_SetLeft(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetLeft").Call(obj, uintptr(value))
}

func HeaderControl_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetTop").Call(obj)
	return int32(ret)
}

func HeaderControl_SetTop(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetTop").Call(obj, uintptr(value))
}

func HeaderControl_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetWidth").Call(obj)
	return int32(ret)
}

func HeaderControl_SetWidth(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetWidth").Call(obj, uintptr(value))
}

func HeaderControl_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetHeight").Call(obj)
	return int32(ret)
}

func HeaderControl_SetHeight(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetHeight").Call(obj, uintptr(value))
}

func HeaderControl_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("HeaderControl_GetCursor").Call(obj)
	return TCursor(ret)
}

func HeaderControl_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("HeaderControl_SetCursor").Call(obj, uintptr(value))
}

func HeaderControl_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderControl_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderControl_SetHint(obj uintptr, value string) {
	getLazyProc("HeaderControl_SetHint").Call(obj, GoStrToDStr(value))
}

func HeaderControl_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetComponentCount").Call(obj)
	return int32(ret)
}

func HeaderControl_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("HeaderControl_GetComponentIndex").Call(obj)
	return int32(ret)
}

func HeaderControl_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("HeaderControl_SetComponentIndex").Call(obj, uintptr(value))
}

func HeaderControl_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetOwner").Call(obj)
	return ret
}

func HeaderControl_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("HeaderControl_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func HeaderControl_SetName(obj uintptr, value string) {
	getLazyProc("HeaderControl_SetName").Call(obj, GoStrToDStr(value))
}

func HeaderControl_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("HeaderControl_GetTag").Call(obj)
	return int(ret)
}

func HeaderControl_SetTag(obj uintptr, value int) {
	getLazyProc("HeaderControl_SetTag").Call(obj, uintptr(value))
}

func HeaderControl_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchorSideLeft").Call(obj)
	return ret
}

func HeaderControl_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetAnchorSideLeft").Call(obj, value)
}

func HeaderControl_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchorSideTop").Call(obj)
	return ret
}

func HeaderControl_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetAnchorSideTop").Call(obj, value)
}

func HeaderControl_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchorSideRight").Call(obj)
	return ret
}

func HeaderControl_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetAnchorSideRight").Call(obj, value)
}

func HeaderControl_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchorSideBottom").Call(obj)
	return ret
}

func HeaderControl_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetAnchorSideBottom").Call(obj, value)
}

func HeaderControl_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetChildSizing").Call(obj)
	return ret
}

func HeaderControl_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetChildSizing").Call(obj, value)
}

func HeaderControl_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetBorderSpacing").Call(obj)
	return ret
}

func HeaderControl_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("HeaderControl_SetBorderSpacing").Call(obj, value)
}

func HeaderControl_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func HeaderControl_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetControls").Call(obj, uintptr(Index))
	return ret
}

func HeaderControl_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func HeaderControl_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("HeaderControl_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func HeaderControl_StaticClassType() TClass {
	r, _, _ := getLazyProc("HeaderControl_StaticClassType").Call()
	return TClass(r)
}
