package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TPanel ---------------------------

func Panel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_Create").Call(obj)
	return ret
}

func Panel_Free(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Free").Call(obj)
}

func Panel_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Panel_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Panel_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Panel_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Panel_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Panel_DisableAlign").Call(obj)
}

func Panel_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Panel_EnableAlign").Call(obj)
}

func Panel_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Panel_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Panel_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("Panel_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Panel_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Panel_InsertControl").Call(obj, AControl)
}

func Panel_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Invalidate").Call(obj)
}

func Panel_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("Panel_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func Panel_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Panel_RemoveControl").Call(obj, AControl)
}

func Panel_Realign(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Realign").Call(obj)
}

func Panel_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Repaint").Call(obj)
}

func Panel_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("Panel_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Panel_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("Panel_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Panel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Panel_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Panel_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("Panel_SetFocus").Call(obj)
}

func Panel_Update(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Update").Call(obj)
}

func Panel_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Panel_BringToFront").Call(obj)
}

func Panel_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Panel_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Panel_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Hide").Call(obj)
}

func Panel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Panel_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Panel_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Refresh").Call(obj)
}

func Panel_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Panel_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Panel_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Panel_SendToBack").Call(obj)
}

func Panel_Show(obj uintptr) {
	_, _, _ = getLazyProc("Panel_Show").Call(obj)
}

func Panel_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Panel_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Panel_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetTextLen").Call(obj)
	return int32(ret)
}

func Panel_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Panel_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Panel_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Panel_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Panel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Panel_Assign").Call(obj, Source)
}

func Panel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Panel_ClassType").Call(obj)
	return TClass(ret)
}

func Panel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_InstanceSize").Call(obj)
	return int32(ret)
}

func Panel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Panel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Panel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Panel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetHashCode").Call(obj)
	return int32(ret)
}

func Panel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Panel_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Panel_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Panel_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Panel_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Panel_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Panel_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Panel_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Panel_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Panel_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Panel_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Panel_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Panel_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Panel_AnchorClient").Call(obj, uintptr(ASpace))
}

func Panel_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Panel_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Panel_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Panel_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Panel_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Panel_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Panel_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Panel_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Panel_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetCanvas").Call(obj)
	return ret
}

func Panel_SetCanvas(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetCanvas").Call(obj, value)
}

func Panel_SetOnPaint(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Panel_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func Panel_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Panel_GetAlign").Call(obj)
	return TAlign(ret)
}

func Panel_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Panel_SetAlign").Call(obj, uintptr(value))
}

func Panel_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("Panel_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func Panel_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("Panel_SetAlignment").Call(obj, uintptr(value))
}

func Panel_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Panel_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Panel_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Panel_SetAnchors").Call(obj, uintptr(value))
}

func Panel_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Panel_GetBevelInner(obj uintptr) TBevelCut {
	ret, _, _ := getLazyProc("Panel_GetBevelInner").Call(obj)
	return TBevelCut(ret)
}

func Panel_SetBevelInner(obj uintptr, value TBevelCut) {
	_, _, _ = getLazyProc("Panel_SetBevelInner").Call(obj, uintptr(value))
}

func Panel_GetBevelOuter(obj uintptr) TBevelCut {
	ret, _, _ := getLazyProc("Panel_GetBevelOuter").Call(obj)
	return TBevelCut(ret)
}

func Panel_SetBevelOuter(obj uintptr, value TBevelCut) {
	_, _, _ = getLazyProc("Panel_SetBevelOuter").Call(obj, uintptr(value))
}

func Panel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Panel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Panel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Panel_SetBiDiMode").Call(obj, uintptr(value))
}

func Panel_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetBorderWidth").Call(obj)
	return int32(ret)
}

func Panel_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetBorderWidth").Call(obj, uintptr(value))
}

func Panel_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("Panel_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func Panel_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("Panel_SetBorderStyle").Call(obj, uintptr(value))
}

func Panel_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("Panel_SetCaption").Call(obj, GoStrToDStr(value))
}

func Panel_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Panel_GetColor").Call(obj)
	return TColor(ret)
}

func Panel_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Panel_SetColor").Call(obj, uintptr(value))
}

func Panel_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetConstraints").Call(obj)
	return ret
}

func Panel_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetConstraints").Call(obj, value)
}

func Panel_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Panel_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Panel_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Panel_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Panel_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Panel_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Panel_SetDragCursor").Call(obj, uintptr(value))
}

func Panel_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Panel_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Panel_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("Panel_SetDragKind").Call(obj, uintptr(value))
}

func Panel_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Panel_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Panel_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Panel_SetDragMode").Call(obj, uintptr(value))
}

func Panel_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Panel_GetFullRepaint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetFullRepaint").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetFullRepaint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetFullRepaint").Call(obj, GoBoolToDBool(value))
}

func Panel_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetFont").Call(obj)
	return ret
}

func Panel_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetFont").Call(obj, value)
}

func Panel_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetParentBackground(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Panel_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetPopupMenu").Call(obj)
	return ret
}

func Panel_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetPopupMenu").Call(obj, value)
}

func Panel_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Panel_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Panel_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Panel_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("Panel_SetTabOrder").Call(obj, uintptr(value))
}

func Panel_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Panel_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Panel_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Panel_SetOnAlignPosition(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnGetSiteInfo(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func Panel_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Panel_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func Panel_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Panel_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Panel_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetBrush").Call(obj)
	return ret
}

func Panel_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetControlCount").Call(obj)
	return int32(ret)
}

func Panel_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Panel_GetHandle").Call(obj)
	return ret
}

func Panel_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Panel_GetParentWindow").Call(obj)
	return ret
}

func Panel_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Panel_SetParentWindow").Call(obj, value)
}

func Panel_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAction").Call(obj)
	return ret
}

func Panel_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetAction").Call(obj, value)
}

func Panel_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Panel_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Panel_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Panel_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetClientHeight").Call(obj)
	return int32(ret)
}

func Panel_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetClientHeight").Call(obj, uintptr(value))
}

func Panel_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Panel_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Panel_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Panel_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetClientWidth").Call(obj)
	return int32(ret)
}

func Panel_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetClientWidth").Call(obj, uintptr(value))
}

func Panel_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Panel_GetControlState").Call(obj)
	return TControlState(ret)
}

func Panel_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Panel_SetControlState").Call(obj, uintptr(value))
}

func Panel_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Panel_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Panel_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Panel_SetControlStyle").Call(obj, uintptr(value))
}

func Panel_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Panel_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Panel_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetParent").Call(obj)
	return ret
}

func Panel_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetParent").Call(obj, value)
}

func Panel_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetLeft").Call(obj)
	return int32(ret)
}

func Panel_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetLeft").Call(obj, uintptr(value))
}

func Panel_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetTop").Call(obj)
	return int32(ret)
}

func Panel_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetTop").Call(obj, uintptr(value))
}

func Panel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetWidth").Call(obj)
	return int32(ret)
}

func Panel_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetWidth").Call(obj, uintptr(value))
}

func Panel_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetHeight").Call(obj)
	return int32(ret)
}

func Panel_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetHeight").Call(obj, uintptr(value))
}

func Panel_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Panel_GetCursor").Call(obj)
	return TCursor(ret)
}

func Panel_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Panel_SetCursor").Call(obj, uintptr(value))
}

func Panel_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Panel_SetHint").Call(obj, GoStrToDStr(value))
}

func Panel_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetComponentCount").Call(obj)
	return int32(ret)
}

func Panel_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Panel_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Panel_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Panel_SetComponentIndex").Call(obj, uintptr(value))
}

func Panel_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetOwner").Call(obj)
	return ret
}

func Panel_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Panel_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Panel_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Panel_SetName").Call(obj, GoStrToDStr(value))
}

func Panel_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Panel_GetTag").Call(obj)
	return int(ret)
}

func Panel_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Panel_SetTag").Call(obj, uintptr(value))
}

func Panel_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAnchorSideLeft").Call(obj)
	return ret
}

func Panel_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetAnchorSideLeft").Call(obj, value)
}

func Panel_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAnchorSideTop").Call(obj)
	return ret
}

func Panel_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetAnchorSideTop").Call(obj, value)
}

func Panel_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAnchorSideRight").Call(obj)
	return ret
}

func Panel_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetAnchorSideRight").Call(obj, value)
}

func Panel_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAnchorSideBottom").Call(obj)
	return ret
}

func Panel_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetAnchorSideBottom").Call(obj, value)
}

func Panel_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetChildSizing").Call(obj)
	return ret
}

func Panel_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetChildSizing").Call(obj, value)
}

func Panel_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Panel_GetBorderSpacing").Call(obj)
	return ret
}

func Panel_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Panel_SetBorderSpacing").Call(obj, value)
}

func Panel_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Panel_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Panel_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Panel_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Panel_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Panel_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Panel_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Panel_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Panel_StaticClassType() TClass {
	r, _, _ := getLazyProc("Panel_StaticClassType").Call()
	return TClass(r)
}
