package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TScrollBox ---------------------------

func ScrollBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_Create").Call(obj)
	return ret
}

func ScrollBox_Free(obj uintptr) {
	getLazyProc("ScrollBox_Free").Call(obj)
}

func ScrollBox_ScrollInView(obj uintptr, AControl uintptr) {
	getLazyProc("ScrollBox_ScrollInView").Call(obj, AControl)
}

func ScrollBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ScrollBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ScrollBox_DisableAlign(obj uintptr) {
	getLazyProc("ScrollBox_DisableAlign").Call(obj)
}

func ScrollBox_EnableAlign(obj uintptr) {
	getLazyProc("ScrollBox_EnableAlign").Call(obj)
}

func ScrollBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ScrollBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ScrollBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ScrollBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ScrollBox_InsertControl").Call(obj, AControl)
}

func ScrollBox_Invalidate(obj uintptr) {
	getLazyProc("ScrollBox_Invalidate").Call(obj)
}

func ScrollBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ScrollBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ScrollBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ScrollBox_RemoveControl").Call(obj, AControl)
}

func ScrollBox_Realign(obj uintptr) {
	getLazyProc("ScrollBox_Realign").Call(obj)
}

func ScrollBox_Repaint(obj uintptr) {
	getLazyProc("ScrollBox_Repaint").Call(obj)
}

func ScrollBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ScrollBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ScrollBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ScrollBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ScrollBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ScrollBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ScrollBox_SetFocus(obj uintptr) {
	getLazyProc("ScrollBox_SetFocus").Call(obj)
}

func ScrollBox_Update(obj uintptr) {
	getLazyProc("ScrollBox_Update").Call(obj)
}

func ScrollBox_BringToFront(obj uintptr) {
	getLazyProc("ScrollBox_BringToFront").Call(obj)
}

func ScrollBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ScrollBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ScrollBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_Hide(obj uintptr) {
	getLazyProc("ScrollBox_Hide").Call(obj)
}

func ScrollBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ScrollBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ScrollBox_Refresh(obj uintptr) {
	getLazyProc("ScrollBox_Refresh").Call(obj)
}

func ScrollBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ScrollBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ScrollBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_SendToBack(obj uintptr) {
	getLazyProc("ScrollBox_SendToBack").Call(obj)
}

func ScrollBox_Show(obj uintptr) {
	getLazyProc("ScrollBox_Show").Call(obj)
}

func ScrollBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ScrollBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ScrollBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ScrollBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ScrollBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ScrollBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ScrollBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ScrollBox_Assign").Call(obj, Source)
}

func ScrollBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ScrollBox_ClassType").Call(obj)
	return TClass(ret)
}

func ScrollBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ScrollBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ScrollBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ScrollBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ScrollBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ScrollBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ScrollBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ScrollBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ScrollBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ScrollBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ScrollBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ScrollBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ScrollBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ScrollBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ScrollBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ScrollBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ScrollBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ScrollBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ScrollBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ScrollBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ScrollBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ScrollBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ScrollBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ScrollBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ScrollBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ScrollBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ScrollBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ScrollBox_SetAlign").Call(obj, uintptr(value))
}

func ScrollBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ScrollBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ScrollBox_SetAnchors").Call(obj, uintptr(value))
}

func ScrollBox_GetAutoScroll(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetAutoScroll").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetAutoScroll(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetAutoScroll").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ScrollBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ScrollBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ScrollBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ScrollBox_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("ScrollBox_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func ScrollBox_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("ScrollBox_SetBorderStyle").Call(obj, uintptr(value))
}

func ScrollBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetConstraints").Call(obj)
	return ret
}

func ScrollBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetConstraints").Call(obj, value)
}

func ScrollBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ScrollBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ScrollBox_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ScrollBox_SetDragCursor").Call(obj, uintptr(value))
}

func ScrollBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ScrollBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ScrollBox_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ScrollBox_SetDragKind").Call(obj, uintptr(value))
}

func ScrollBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ScrollBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ScrollBox_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ScrollBox_SetDragMode").Call(obj, uintptr(value))
}

func ScrollBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ScrollBox_GetColor").Call(obj)
	return TColor(ret)
}

func ScrollBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("ScrollBox_SetColor").Call(obj, uintptr(value))
}

func ScrollBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetFont").Call(obj)
	return ret
}

func ScrollBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetFont").Call(obj, value)
}

func ScrollBox_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetParentBackground(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetPopupMenu").Call(obj)
	return ret
}

func ScrollBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetPopupMenu").Call(obj, value)
}

func ScrollBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ScrollBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ScrollBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ScrollBox_SetTabOrder").Call(obj, uintptr(value))
}

func ScrollBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnConstrainedResize(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnConstrainedResize").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnDockDrop(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnGetSiteInfo(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseWheel(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnResize(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnUnDock(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_SetOnAlignPosition(obj uintptr, fn interface{}) {
	getLazyProc("ScrollBox_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func ScrollBox_GetHorzScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetHorzScrollBar").Call(obj)
	return ret
}

func ScrollBox_SetHorzScrollBar(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetHorzScrollBar").Call(obj, value)
}

func ScrollBox_GetVertScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetVertScrollBar").Call(obj)
	return ret
}

func ScrollBox_SetVertScrollBar(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetVertScrollBar").Call(obj, value)
}

func ScrollBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ScrollBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ScrollBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetBrush").Call(obj)
	return ret
}

func ScrollBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ScrollBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ScrollBox_GetHandle").Call(obj)
	return HWND(ret)
}

func ScrollBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ScrollBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ScrollBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ScrollBox_SetParentWindow").Call(obj, uintptr(value))
}

func ScrollBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ScrollBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ScrollBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAction").Call(obj)
	return ret
}

func ScrollBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetAction").Call(obj, value)
}

func ScrollBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ScrollBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ScrollBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ScrollBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ScrollBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetClientHeight").Call(obj, uintptr(value))
}

func ScrollBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ScrollBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ScrollBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ScrollBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetClientWidth").Call(obj, uintptr(value))
}

func ScrollBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ScrollBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ScrollBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ScrollBox_SetControlState").Call(obj, uintptr(value))
}

func ScrollBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ScrollBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ScrollBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ScrollBox_SetControlStyle").Call(obj, uintptr(value))
}

func ScrollBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetParent").Call(obj)
	return ret
}

func ScrollBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetParent").Call(obj, value)
}

func ScrollBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetLeft").Call(obj)
	return int32(ret)
}

func ScrollBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetLeft").Call(obj, uintptr(value))
}

func ScrollBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetTop").Call(obj)
	return int32(ret)
}

func ScrollBox_SetTop(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetTop").Call(obj, uintptr(value))
}

func ScrollBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetWidth").Call(obj)
	return int32(ret)
}

func ScrollBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetWidth").Call(obj, uintptr(value))
}

func ScrollBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetHeight").Call(obj)
	return int32(ret)
}

func ScrollBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetHeight").Call(obj, uintptr(value))
}

func ScrollBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ScrollBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ScrollBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ScrollBox_SetCursor").Call(obj, uintptr(value))
}

func ScrollBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBox_SetHint(obj uintptr, value string) {
	getLazyProc("ScrollBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ScrollBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ScrollBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ScrollBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ScrollBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ScrollBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetOwner").Call(obj)
	return ret
}

func ScrollBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBox_SetName(obj uintptr, value string) {
	getLazyProc("ScrollBox_SetName").Call(obj, GoStrToDStr(value))
}

func ScrollBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ScrollBox_GetTag").Call(obj)
	return int(ret)
}

func ScrollBox_SetTag(obj uintptr, value int) {
	getLazyProc("ScrollBox_SetTag").Call(obj, uintptr(value))
}

func ScrollBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ScrollBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetAnchorSideLeft").Call(obj, value)
}

func ScrollBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ScrollBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetAnchorSideTop").Call(obj, value)
}

func ScrollBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ScrollBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetAnchorSideRight").Call(obj, value)
}

func ScrollBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ScrollBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetAnchorSideBottom").Call(obj, value)
}

func ScrollBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetChildSizing").Call(obj)
	return ret
}

func ScrollBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetChildSizing").Call(obj, value)
}

func ScrollBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetBorderSpacing").Call(obj)
	return ret
}

func ScrollBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ScrollBox_SetBorderSpacing").Call(obj, value)
}

func ScrollBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ScrollBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ScrollBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ScrollBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ScrollBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ScrollBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ScrollBox_StaticClassType").Call()
	return TClass(r)
}
