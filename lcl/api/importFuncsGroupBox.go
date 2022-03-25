package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TGroupBox ---------------------------

func GroupBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_Create").Call(obj)
	return ret
}

func GroupBox_Free(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Free").Call(obj)
}

func GroupBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func GroupBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("GroupBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func GroupBox_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_DisableAlign").Call(obj)
}

func GroupBox_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_EnableAlign").Call(obj)
}

func GroupBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("GroupBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func GroupBox_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("GroupBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func GroupBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("GroupBox_InsertControl").Call(obj, AControl)
}

func GroupBox_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Invalidate").Call(obj)
}

func GroupBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("GroupBox_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func GroupBox_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("GroupBox_RemoveControl").Call(obj, AControl)
}

func GroupBox_Realign(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Realign").Call(obj)
}

func GroupBox_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Repaint").Call(obj)
}

func GroupBox_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("GroupBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func GroupBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("GroupBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func GroupBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("GroupBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func GroupBox_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetFocus").Call(obj)
}

func GroupBox_Update(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Update").Call(obj)
}

func GroupBox_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_BringToFront").Call(obj)
}

func GroupBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("GroupBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("GroupBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_Hide(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Hide").Call(obj)
}

func GroupBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("GroupBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func GroupBox_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Refresh").Call(obj)
}

func GroupBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("GroupBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("GroupBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_SendToBack").Call(obj)
}

func GroupBox_Show(obj uintptr) {
	_, _, _ = getLazyProc("GroupBox_Show").Call(obj)
}

func GroupBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("GroupBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func GroupBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetTextLen").Call(obj)
	return int32(ret)
}

func GroupBox_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("GroupBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func GroupBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("GroupBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func GroupBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("GroupBox_Assign").Call(obj, Source)
}

func GroupBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("GroupBox_ClassType").Call(obj)
	return TClass(ret)
}

func GroupBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_InstanceSize").Call(obj)
	return int32(ret)
}

func GroupBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("GroupBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func GroupBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func GroupBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetHashCode").Call(obj)
	return int32(ret)
}

func GroupBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("GroupBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func GroupBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("GroupBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func GroupBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("GroupBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func GroupBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("GroupBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func GroupBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("GroupBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func GroupBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("GroupBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func GroupBox_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("GroupBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func GroupBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("GroupBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func GroupBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("GroupBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func GroupBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("GroupBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func GroupBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("GroupBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func GroupBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("GroupBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func GroupBox_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("GroupBox_SetAlign").Call(obj, uintptr(value))
}

func GroupBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("GroupBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func GroupBox_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("GroupBox_SetAnchors").Call(obj, uintptr(value))
}

func GroupBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("GroupBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func GroupBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("GroupBox_SetBiDiMode").Call(obj, uintptr(value))
}

func GroupBox_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("GroupBox_SetCaption").Call(obj, GoStrToDStr(value))
}

func GroupBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("GroupBox_GetColor").Call(obj)
	return TColor(ret)
}

func GroupBox_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("GroupBox_SetColor").Call(obj, uintptr(value))
}

func GroupBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetConstraints").Call(obj)
	return ret
}

func GroupBox_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetConstraints").Call(obj, value)
}

func GroupBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("GroupBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func GroupBox_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("GroupBox_SetDragCursor").Call(obj, uintptr(value))
}

func GroupBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("GroupBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func GroupBox_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("GroupBox_SetDragKind").Call(obj, uintptr(value))
}

func GroupBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("GroupBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func GroupBox_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("GroupBox_SetDragMode").Call(obj, uintptr(value))
}

func GroupBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetFont").Call(obj)
	return ret
}

func GroupBox_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetFont").Call(obj, value)
}

func GroupBox_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetParentBackground(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetPopupMenu").Call(obj)
	return ret
}

func GroupBox_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetPopupMenu").Call(obj, value)
}

func GroupBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("GroupBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func GroupBox_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("GroupBox_SetTabOrder").Call(obj, uintptr(value))
}

func GroupBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func GroupBox_SetOnAlignPosition(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnGetSiteInfo(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("GroupBox_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func GroupBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func GroupBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func GroupBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetBrush").Call(obj)
	return ret
}

func GroupBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetControlCount").Call(obj)
	return int32(ret)
}

func GroupBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("GroupBox_GetHandle").Call(obj)
	return ret
}

func GroupBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("GroupBox_GetParentWindow").Call(obj)
	return ret
}

func GroupBox_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("GroupBox_SetParentWindow").Call(obj, value)
}

func GroupBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("GroupBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAction").Call(obj)
	return ret
}

func GroupBox_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetAction").Call(obj, value)
}

func GroupBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("GroupBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("GroupBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func GroupBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func GroupBox_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetClientHeight").Call(obj, uintptr(value))
}

func GroupBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("GroupBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("GroupBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func GroupBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func GroupBox_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetClientWidth").Call(obj, uintptr(value))
}

func GroupBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("GroupBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func GroupBox_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("GroupBox_SetControlState").Call(obj, uintptr(value))
}

func GroupBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("GroupBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func GroupBox_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("GroupBox_SetControlStyle").Call(obj, uintptr(value))
}

func GroupBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("GroupBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func GroupBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetParent").Call(obj)
	return ret
}

func GroupBox_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetParent").Call(obj, value)
}

func GroupBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetLeft").Call(obj)
	return int32(ret)
}

func GroupBox_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetLeft").Call(obj, uintptr(value))
}

func GroupBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetTop").Call(obj)
	return int32(ret)
}

func GroupBox_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetTop").Call(obj, uintptr(value))
}

func GroupBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetWidth").Call(obj)
	return int32(ret)
}

func GroupBox_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetWidth").Call(obj, uintptr(value))
}

func GroupBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetHeight").Call(obj)
	return int32(ret)
}

func GroupBox_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetHeight").Call(obj, uintptr(value))
}

func GroupBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("GroupBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func GroupBox_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("GroupBox_SetCursor").Call(obj, uintptr(value))
}

func GroupBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("GroupBox_SetHint").Call(obj, GoStrToDStr(value))
}

func GroupBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func GroupBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("GroupBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func GroupBox_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("GroupBox_SetComponentIndex").Call(obj, uintptr(value))
}

func GroupBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetOwner").Call(obj)
	return ret
}

func GroupBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("GroupBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func GroupBox_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("GroupBox_SetName").Call(obj, GoStrToDStr(value))
}

func GroupBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("GroupBox_GetTag").Call(obj)
	return int(ret)
}

func GroupBox_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("GroupBox_SetTag").Call(obj, uintptr(value))
}

func GroupBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func GroupBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetAnchorSideLeft").Call(obj, value)
}

func GroupBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAnchorSideTop").Call(obj)
	return ret
}

func GroupBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetAnchorSideTop").Call(obj, value)
}

func GroupBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAnchorSideRight").Call(obj)
	return ret
}

func GroupBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetAnchorSideRight").Call(obj, value)
}

func GroupBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func GroupBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetAnchorSideBottom").Call(obj, value)
}

func GroupBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetChildSizing").Call(obj)
	return ret
}

func GroupBox_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetChildSizing").Call(obj, value)
}

func GroupBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetBorderSpacing").Call(obj)
	return ret
}

func GroupBox_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("GroupBox_SetBorderSpacing").Call(obj, value)
}

func GroupBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func GroupBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func GroupBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func GroupBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("GroupBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func GroupBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("GroupBox_StaticClassType").Call()
	return TClass(r)
}
