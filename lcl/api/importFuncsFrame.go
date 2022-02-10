package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TFrame ---------------------------

func Frame_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_Create").Call(obj)
	return ret
}

func Frame_Free(obj uintptr) {
	getLazyProc("Frame_Free").Call(obj)
}

func Frame_ScrollInView(obj uintptr, AControl uintptr) {
	getLazyProc("Frame_ScrollInView").Call(obj, AControl)
}

func Frame_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Frame_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Frame_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Frame_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Frame_DisableAlign(obj uintptr) {
	getLazyProc("Frame_DisableAlign").Call(obj)
}

func Frame_EnableAlign(obj uintptr) {
	getLazyProc("Frame_EnableAlign").Call(obj)
}

func Frame_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Frame_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Frame_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("Frame_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Frame_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("Frame_InsertControl").Call(obj, AControl)
}

func Frame_Invalidate(obj uintptr) {
	getLazyProc("Frame_Invalidate").Call(obj)
}

func Frame_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("Frame_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func Frame_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("Frame_RemoveControl").Call(obj, AControl)
}

func Frame_Realign(obj uintptr) {
	getLazyProc("Frame_Realign").Call(obj)
}

func Frame_Repaint(obj uintptr) {
	getLazyProc("Frame_Repaint").Call(obj)
}

func Frame_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("Frame_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Frame_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("Frame_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Frame_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("Frame_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Frame_SetFocus(obj uintptr) {
	getLazyProc("Frame_SetFocus").Call(obj)
}

func Frame_Update(obj uintptr) {
	getLazyProc("Frame_Update").Call(obj)
}

func Frame_BringToFront(obj uintptr) {
	getLazyProc("Frame_BringToFront").Call(obj)
}

func Frame_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Frame_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Frame_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_Hide(obj uintptr) {
	getLazyProc("Frame_Hide").Call(obj)
}

func Frame_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Frame_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Frame_Refresh(obj uintptr) {
	getLazyProc("Frame_Refresh").Call(obj)
}

func Frame_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Frame_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Frame_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_SendToBack(obj uintptr) {
	getLazyProc("Frame_SendToBack").Call(obj)
}

func Frame_Show(obj uintptr) {
	getLazyProc("Frame_Show").Call(obj)
}

func Frame_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Frame_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Frame_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetTextLen").Call(obj)
	return int32(ret)
}

func Frame_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("Frame_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Frame_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Frame_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Frame_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Frame_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Frame_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Frame_Assign").Call(obj, Source)
}

func Frame_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Frame_ClassType").Call(obj)
	return TClass(ret)
}

func Frame_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Frame_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Frame_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_InstanceSize").Call(obj)
	return int32(ret)
}

func Frame_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Frame_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Frame_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Frame_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetHashCode").Call(obj)
	return int32(ret)
}

func Frame_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Frame_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Frame_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Frame_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Frame_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Frame_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Frame_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Frame_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Frame_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Frame_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Frame_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("Frame_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Frame_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("Frame_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Frame_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("Frame_AnchorClient").Call(obj, uintptr(ASpace))
}

func Frame_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Frame_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Frame_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("Frame_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Frame_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("Frame_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Frame_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("Frame_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Frame_GetDesignTimePPI(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetDesignTimePPI").Call(obj)
	return int32(ret)
}

func Frame_SetDesignTimePPI(obj uintptr, value int32) {
	getLazyProc("Frame_SetDesignTimePPI").Call(obj, uintptr(value))
}

func Frame_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Frame_GetAlign").Call(obj)
	return TAlign(ret)
}

func Frame_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("Frame_SetAlign").Call(obj, uintptr(value))
}

func Frame_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Frame_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Frame_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("Frame_SetAnchors").Call(obj, uintptr(value))
}

func Frame_GetAutoScroll(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetAutoScroll").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetAutoScroll(obj uintptr, value bool) {
	getLazyProc("Frame_SetAutoScroll").Call(obj, GoBoolToDBool(value))
}

func Frame_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("Frame_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Frame_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Frame_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Frame_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("Frame_SetBiDiMode").Call(obj, uintptr(value))
}

func Frame_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetConstraints").Call(obj)
	return ret
}

func Frame_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetConstraints").Call(obj, value)
}

func Frame_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetDockSite(obj uintptr, value bool) {
	getLazyProc("Frame_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Frame_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("Frame_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Frame_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Frame_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Frame_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("Frame_SetDragCursor").Call(obj, uintptr(value))
}

func Frame_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Frame_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Frame_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("Frame_SetDragKind").Call(obj, uintptr(value))
}

func Frame_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Frame_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Frame_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("Frame_SetDragMode").Call(obj, uintptr(value))
}

func Frame_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetEnabled(obj uintptr, value bool) {
	getLazyProc("Frame_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Frame_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Frame_GetColor").Call(obj)
	return TColor(ret)
}

func Frame_SetColor(obj uintptr, value TColor) {
	getLazyProc("Frame_SetColor").Call(obj, uintptr(value))
}

func Frame_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetFont").Call(obj)
	return ret
}

func Frame_SetFont(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetFont").Call(obj, value)
}

func Frame_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetParentBackground(obj uintptr, value bool) {
	getLazyProc("Frame_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func Frame_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetParentColor(obj uintptr, value bool) {
	getLazyProc("Frame_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Frame_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("Frame_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Frame_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetParentFont(obj uintptr, value bool) {
	getLazyProc("Frame_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Frame_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("Frame_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Frame_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetPopupMenu").Call(obj)
	return ret
}

func Frame_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetPopupMenu").Call(obj, value)
}

func Frame_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetShowHint(obj uintptr, value bool) {
	getLazyProc("Frame_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Frame_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Frame_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Frame_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("Frame_SetTabOrder").Call(obj, uintptr(value))
}

func Frame_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetTabStop(obj uintptr, value bool) {
	getLazyProc("Frame_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Frame_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetVisible(obj uintptr, value bool) {
	getLazyProc("Frame_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Frame_SetOnAlignPosition(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnConstrainedResize(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnConstrainedResize").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnDockDrop(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnGetSiteInfo(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseWheel(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnResize(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func Frame_SetOnUnDock(obj uintptr, fn interface{}) {
	getLazyProc("Frame_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func Frame_GetHorzScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetHorzScrollBar").Call(obj)
	return ret
}

func Frame_SetHorzScrollBar(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetHorzScrollBar").Call(obj, value)
}

func Frame_GetVertScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetVertScrollBar").Call(obj)
	return ret
}

func Frame_SetVertScrollBar(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetVertScrollBar").Call(obj, value)
}

func Frame_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Frame_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Frame_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetBrush").Call(obj)
	return ret
}

func Frame_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetControlCount").Call(obj)
	return int32(ret)
}

func Frame_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Frame_GetHandle").Call(obj)
	return HWND(ret)
}

func Frame_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Frame_GetParentWindow").Call(obj)
	return HWND(ret)
}

func Frame_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("Frame_SetParentWindow").Call(obj, uintptr(value))
}

func Frame_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("Frame_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Frame_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAction").Call(obj)
	return ret
}

func Frame_SetAction(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetAction").Call(obj, value)
}

func Frame_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Frame_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("Frame_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Frame_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetClientHeight").Call(obj)
	return int32(ret)
}

func Frame_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("Frame_SetClientHeight").Call(obj, uintptr(value))
}

func Frame_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Frame_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Frame_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Frame_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetClientWidth").Call(obj)
	return int32(ret)
}

func Frame_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("Frame_SetClientWidth").Call(obj, uintptr(value))
}

func Frame_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Frame_GetControlState").Call(obj)
	return TControlState(ret)
}

func Frame_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("Frame_SetControlState").Call(obj, uintptr(value))
}

func Frame_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Frame_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Frame_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("Frame_SetControlStyle").Call(obj, uintptr(value))
}

func Frame_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Frame_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Frame_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetParent").Call(obj)
	return ret
}

func Frame_SetParent(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetParent").Call(obj, value)
}

func Frame_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetLeft").Call(obj)
	return int32(ret)
}

func Frame_SetLeft(obj uintptr, value int32) {
	getLazyProc("Frame_SetLeft").Call(obj, uintptr(value))
}

func Frame_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetTop").Call(obj)
	return int32(ret)
}

func Frame_SetTop(obj uintptr, value int32) {
	getLazyProc("Frame_SetTop").Call(obj, uintptr(value))
}

func Frame_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetWidth").Call(obj)
	return int32(ret)
}

func Frame_SetWidth(obj uintptr, value int32) {
	getLazyProc("Frame_SetWidth").Call(obj, uintptr(value))
}

func Frame_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetHeight").Call(obj)
	return int32(ret)
}

func Frame_SetHeight(obj uintptr, value int32) {
	getLazyProc("Frame_SetHeight").Call(obj, uintptr(value))
}

func Frame_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Frame_GetCursor").Call(obj)
	return TCursor(ret)
}

func Frame_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("Frame_SetCursor").Call(obj, uintptr(value))
}

func Frame_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Frame_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Frame_SetHint(obj uintptr, value string) {
	getLazyProc("Frame_SetHint").Call(obj, GoStrToDStr(value))
}

func Frame_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetComponentCount").Call(obj)
	return int32(ret)
}

func Frame_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Frame_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Frame_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("Frame_SetComponentIndex").Call(obj, uintptr(value))
}

func Frame_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetOwner").Call(obj)
	return ret
}

func Frame_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Frame_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Frame_SetName(obj uintptr, value string) {
	getLazyProc("Frame_SetName").Call(obj, GoStrToDStr(value))
}

func Frame_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Frame_GetTag").Call(obj)
	return int(ret)
}

func Frame_SetTag(obj uintptr, value int) {
	getLazyProc("Frame_SetTag").Call(obj, uintptr(value))
}

func Frame_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAnchorSideLeft").Call(obj)
	return ret
}

func Frame_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetAnchorSideLeft").Call(obj, value)
}

func Frame_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAnchorSideTop").Call(obj)
	return ret
}

func Frame_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetAnchorSideTop").Call(obj, value)
}

func Frame_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAnchorSideRight").Call(obj)
	return ret
}

func Frame_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetAnchorSideRight").Call(obj, value)
}

func Frame_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAnchorSideBottom").Call(obj)
	return ret
}

func Frame_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetAnchorSideBottom").Call(obj, value)
}

func Frame_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetChildSizing").Call(obj)
	return ret
}

func Frame_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetChildSizing").Call(obj, value)
}

func Frame_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Frame_GetBorderSpacing").Call(obj)
	return ret
}

func Frame_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("Frame_SetBorderSpacing").Call(obj, value)
}

func Frame_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Frame_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Frame_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Frame_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Frame_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Frame_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Frame_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Frame_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Frame_StaticClassType() TClass {
	r, _, _ := getLazyProc("Frame_StaticClassType").Call()
	return TClass(r)
}
