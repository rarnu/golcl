package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TPageControl ---------------------------

func PageControl_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_Create").Call(obj)
	return ret
}

func PageControl_Free(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Free").Call(obj)
}

func PageControl_SelectNextPage(obj uintptr, GoForward bool, CheckTabVisible bool) {
	_, _, _ = getLazyProc("PageControl_SelectNextPage").Call(obj, GoBoolToDBool(GoForward), GoBoolToDBool(CheckTabVisible))
}

func PageControl_TabRect(obj uintptr, Index int32) TRect {
	var ret TRect
	_, _, _ = getLazyProc("PageControl_TabRect").Call(obj, uintptr(Index), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func PageControl_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("PageControl_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func PageControl_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_DisableAlign").Call(obj)
}

func PageControl_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_EnableAlign").Call(obj)
}

func PageControl_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("PageControl_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func PageControl_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("PageControl_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func PageControl_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("PageControl_InsertControl").Call(obj, AControl)
}

func PageControl_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Invalidate").Call(obj)
}

func PageControl_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("PageControl_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func PageControl_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("PageControl_RemoveControl").Call(obj, AControl)
}

func PageControl_Realign(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Realign").Call(obj)
}

func PageControl_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Repaint").Call(obj)
}

func PageControl_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("PageControl_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func PageControl_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("PageControl_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func PageControl_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("PageControl_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func PageControl_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_SetFocus").Call(obj)
}

func PageControl_Update(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Update").Call(obj)
}

func PageControl_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_BringToFront").Call(obj)
}

func PageControl_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PageControl_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PageControl_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_Hide(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Hide").Call(obj)
}

func PageControl_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("PageControl_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func PageControl_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Refresh").Call(obj)
}

func PageControl_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PageControl_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PageControl_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_SendToBack").Call(obj)
}

func PageControl_Show(obj uintptr) {
	_, _, _ = getLazyProc("PageControl_Show").Call(obj)
}

func PageControl_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("PageControl_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func PageControl_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetTextLen").Call(obj)
	return int32(ret)
}

func PageControl_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("PageControl_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func PageControl_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PageControl_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PageControl_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PageControl_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PageControl_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("PageControl_Assign").Call(obj, Source)
}

func PageControl_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PageControl_ClassType").Call(obj)
	return TClass(ret)
}

func PageControl_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PageControl_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PageControl_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_InstanceSize").Call(obj)
	return int32(ret)
}

func PageControl_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PageControl_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PageControl_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PageControl_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetHashCode").Call(obj)
	return int32(ret)
}

func PageControl_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PageControl_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PageControl_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("PageControl_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func PageControl_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("PageControl_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func PageControl_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("PageControl_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func PageControl_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("PageControl_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func PageControl_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("PageControl_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func PageControl_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("PageControl_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func PageControl_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("PageControl_AnchorClient").Call(obj, uintptr(ASpace))
}

func PageControl_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PageControl_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PageControl_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("PageControl_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func PageControl_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("PageControl_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func PageControl_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("PageControl_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func PageControl_GetOptions(obj uintptr) TCTabControlOptions {
	ret, _, _ := getLazyProc("PageControl_GetOptions").Call(obj)
	return TCTabControlOptions(ret)
}

func PageControl_SetOptions(obj uintptr, value TCTabControlOptions) {
	_, _, _ = getLazyProc("PageControl_SetOptions").Call(obj, uintptr(value))
}

func PageControl_GetActivePageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetActivePageIndex").Call(obj)
	return int32(ret)
}

func PageControl_SetActivePageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetActivePageIndex").Call(obj, uintptr(value))
}

func PageControl_GetPageCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetPageCount").Call(obj)
	return int32(ret)
}

func PageControl_GetActivePage(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetActivePage").Call(obj)
	return ret
}

func PageControl_SetActivePage(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetActivePage").Call(obj, value)
}

func PageControl_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("PageControl_GetAlign").Call(obj)
	return TAlign(ret)
}

func PageControl_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("PageControl_SetAlign").Call(obj, uintptr(value))
}

func PageControl_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("PageControl_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func PageControl_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("PageControl_SetAnchors").Call(obj, uintptr(value))
}

func PageControl_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("PageControl_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func PageControl_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("PageControl_SetBiDiMode").Call(obj, uintptr(value))
}

func PageControl_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetConstraints").Call(obj)
	return ret
}

func PageControl_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetConstraints").Call(obj, value)
}

func PageControl_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("PageControl_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func PageControl_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("PageControl_SetDragCursor").Call(obj, uintptr(value))
}

func PageControl_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("PageControl_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func PageControl_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("PageControl_SetDragKind").Call(obj, uintptr(value))
}

func PageControl_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("PageControl_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func PageControl_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("PageControl_SetDragMode").Call(obj, uintptr(value))
}

func PageControl_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetFont").Call(obj)
	return ret
}

func PageControl_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetFont").Call(obj, value)
}

func PageControl_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetImages").Call(obj)
	return ret
}

func PageControl_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetImages").Call(obj, value)
}

func PageControl_GetMultiLine(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetMultiLine").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetMultiLine(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetMultiLine").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetPopupMenu").Call(obj)
	return ret
}

func PageControl_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetPopupMenu").Call(obj, value)
}

func PageControl_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetTabHeight(obj uintptr) int16 {
	ret, _, _ := getLazyProc("PageControl_GetTabHeight").Call(obj)
	return int16(ret)
}

func PageControl_SetTabHeight(obj uintptr, value int16) {
	_, _, _ = getLazyProc("PageControl_SetTabHeight").Call(obj, uintptr(value))
}

func PageControl_GetTabIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetTabIndex").Call(obj)
	return int32(ret)
}

func PageControl_SetTabIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetTabIndex").Call(obj, uintptr(value))
}

func PageControl_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("PageControl_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func PageControl_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("PageControl_SetTabOrder").Call(obj, uintptr(value))
}

func PageControl_GetTabPosition(obj uintptr) TTabPosition {
	ret, _, _ := getLazyProc("PageControl_GetTabPosition").Call(obj)
	return TTabPosition(ret)
}

func PageControl_SetTabPosition(obj uintptr, value TTabPosition) {
	_, _, _ = getLazyProc("PageControl_SetTabPosition").Call(obj, uintptr(value))
}

func PageControl_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetTabWidth(obj uintptr) int16 {
	ret, _, _ := getLazyProc("PageControl_GetTabWidth").Call(obj)
	return int16(ret)
}

func PageControl_SetTabWidth(obj uintptr, value int16) {
	_, _, _ = getLazyProc("PageControl_SetTabWidth").Call(obj, uintptr(value))
}

func PageControl_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetVisible").Call(obj, GoBoolToDBool(value))
}

func PageControl_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnChanging(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnChanging").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnGetSiteInfo(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func PageControl_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PageControl_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func PageControl_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetDockClientCount").Call(obj)
	return int32(ret)
}

func PageControl_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func PageControl_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetBrush").Call(obj)
	return ret
}

func PageControl_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetControlCount").Call(obj)
	return int32(ret)
}

func PageControl_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("PageControl_GetHandle").Call(obj)
	return ret
}

func PageControl_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("PageControl_GetParentWindow").Call(obj)
	return ret
}

func PageControl_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("PageControl_SetParentWindow").Call(obj, value)
}

func PageControl_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PageControl_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func PageControl_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAction").Call(obj)
	return ret
}

func PageControl_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetAction").Call(obj, value)
}

func PageControl_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("PageControl_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("PageControl_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PageControl_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetClientHeight").Call(obj)
	return int32(ret)
}

func PageControl_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetClientHeight").Call(obj, uintptr(value))
}

func PageControl_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PageControl_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("PageControl_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PageControl_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetClientWidth").Call(obj)
	return int32(ret)
}

func PageControl_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetClientWidth").Call(obj, uintptr(value))
}

func PageControl_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("PageControl_GetControlState").Call(obj)
	return TControlState(ret)
}

func PageControl_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("PageControl_SetControlState").Call(obj, uintptr(value))
}

func PageControl_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("PageControl_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func PageControl_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("PageControl_SetControlStyle").Call(obj, uintptr(value))
}

func PageControl_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("PageControl_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func PageControl_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetParent").Call(obj)
	return ret
}

func PageControl_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetParent").Call(obj, value)
}

func PageControl_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetLeft").Call(obj)
	return int32(ret)
}

func PageControl_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetLeft").Call(obj, uintptr(value))
}

func PageControl_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetTop").Call(obj)
	return int32(ret)
}

func PageControl_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetTop").Call(obj, uintptr(value))
}

func PageControl_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetWidth").Call(obj)
	return int32(ret)
}

func PageControl_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetWidth").Call(obj, uintptr(value))
}

func PageControl_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetHeight").Call(obj)
	return int32(ret)
}

func PageControl_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetHeight").Call(obj, uintptr(value))
}

func PageControl_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("PageControl_GetCursor").Call(obj)
	return TCursor(ret)
}

func PageControl_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("PageControl_SetCursor").Call(obj, uintptr(value))
}

func PageControl_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("PageControl_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func PageControl_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("PageControl_SetHint").Call(obj, GoStrToDStr(value))
}

func PageControl_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetComponentCount").Call(obj)
	return int32(ret)
}

func PageControl_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PageControl_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PageControl_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PageControl_SetComponentIndex").Call(obj, uintptr(value))
}

func PageControl_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetOwner").Call(obj)
	return ret
}

func PageControl_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PageControl_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PageControl_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("PageControl_SetName").Call(obj, GoStrToDStr(value))
}

func PageControl_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PageControl_GetTag").Call(obj)
	return int(ret)
}

func PageControl_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("PageControl_SetTag").Call(obj, uintptr(value))
}

func PageControl_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAnchorSideLeft").Call(obj)
	return ret
}

func PageControl_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetAnchorSideLeft").Call(obj, value)
}

func PageControl_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAnchorSideTop").Call(obj)
	return ret
}

func PageControl_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetAnchorSideTop").Call(obj, value)
}

func PageControl_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAnchorSideRight").Call(obj)
	return ret
}

func PageControl_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetAnchorSideRight").Call(obj, value)
}

func PageControl_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAnchorSideBottom").Call(obj)
	return ret
}

func PageControl_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetAnchorSideBottom").Call(obj, value)
}

func PageControl_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetChildSizing").Call(obj)
	return ret
}

func PageControl_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetChildSizing").Call(obj, value)
}

func PageControl_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetBorderSpacing").Call(obj)
	return ret
}

func PageControl_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PageControl_SetBorderSpacing").Call(obj, value)
}

func PageControl_GetPages(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetPages").Call(obj, uintptr(Index))
	return ret
}

func PageControl_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func PageControl_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetControls").Call(obj, uintptr(Index))
	return ret
}

func PageControl_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PageControl_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("PageControl_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func PageControl_StaticClassType() TClass {
	r, _, _ := getLazyProc("PageControl_StaticClassType").Call()
	return TClass(r)
}
