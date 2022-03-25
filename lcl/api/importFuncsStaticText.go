package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TStaticText ---------------------------

func StaticText_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_Create").Call(obj)
	return ret
}

func StaticText_Free(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Free").Call(obj)
}

func StaticText_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func StaticText_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("StaticText_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func StaticText_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_DisableAlign").Call(obj)
}

func StaticText_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_EnableAlign").Call(obj)
}

func StaticText_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("StaticText_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func StaticText_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("StaticText_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func StaticText_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("StaticText_InsertControl").Call(obj, AControl)
}

func StaticText_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Invalidate").Call(obj)
}

func StaticText_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("StaticText_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func StaticText_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("StaticText_RemoveControl").Call(obj, AControl)
}

func StaticText_Realign(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Realign").Call(obj)
}

func StaticText_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Repaint").Call(obj)
}

func StaticText_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("StaticText_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func StaticText_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("StaticText_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func StaticText_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("StaticText_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func StaticText_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_SetFocus").Call(obj)
}

func StaticText_Update(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Update").Call(obj)
}

func StaticText_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_BringToFront").Call(obj)
}

func StaticText_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("StaticText_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("StaticText_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_Hide(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Hide").Call(obj)
}

func StaticText_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("StaticText_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func StaticText_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Refresh").Call(obj)
}

func StaticText_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("StaticText_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("StaticText_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_SendToBack").Call(obj)
}

func StaticText_Show(obj uintptr) {
	_, _, _ = getLazyProc("StaticText_Show").Call(obj)
}

func StaticText_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("StaticText_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func StaticText_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetTextLen").Call(obj)
	return int32(ret)
}

func StaticText_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("StaticText_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func StaticText_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("StaticText_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func StaticText_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("StaticText_Assign").Call(obj, Source)
}

func StaticText_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StaticText_ClassType").Call(obj)
	return TClass(ret)
}

func StaticText_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_InstanceSize").Call(obj)
	return int32(ret)
}

func StaticText_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StaticText_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StaticText_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func StaticText_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetHashCode").Call(obj)
	return int32(ret)
}

func StaticText_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("StaticText_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StaticText_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("StaticText_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StaticText_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("StaticText_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func StaticText_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("StaticText_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func StaticText_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("StaticText_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func StaticText_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("StaticText_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func StaticText_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("StaticText_AnchorClient").Call(obj, uintptr(ASpace))
}

func StaticText_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StaticText_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StaticText_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("StaticText_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func StaticText_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("StaticText_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func StaticText_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("StaticText_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func StaticText_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("StaticText_GetAlign").Call(obj)
	return TAlign(ret)
}

func StaticText_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("StaticText_SetAlign").Call(obj, uintptr(value))
}

func StaticText_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("StaticText_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func StaticText_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("StaticText_SetAlignment").Call(obj, uintptr(value))
}

func StaticText_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("StaticText_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func StaticText_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("StaticText_SetAnchors").Call(obj, uintptr(value))
}

func StaticText_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("StaticText_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func StaticText_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("StaticText_SetBiDiMode").Call(obj, uintptr(value))
}

func StaticText_GetBorderStyle(obj uintptr) TStaticBorderStyle {
	ret, _, _ := getLazyProc("StaticText_GetBorderStyle").Call(obj)
	return TStaticBorderStyle(ret)
}

func StaticText_SetBorderStyle(obj uintptr, value TStaticBorderStyle) {
	_, _, _ = getLazyProc("StaticText_SetBorderStyle").Call(obj, uintptr(value))
}

func StaticText_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("StaticText_SetCaption").Call(obj, GoStrToDStr(value))
}

func StaticText_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StaticText_GetColor").Call(obj)
	return TColor(ret)
}

func StaticText_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("StaticText_SetColor").Call(obj, uintptr(value))
}

func StaticText_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetConstraints").Call(obj)
	return ret
}

func StaticText_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetConstraints").Call(obj, value)
}

func StaticText_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StaticText_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func StaticText_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("StaticText_SetDragCursor").Call(obj, uintptr(value))
}

func StaticText_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("StaticText_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func StaticText_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("StaticText_SetDragKind").Call(obj, uintptr(value))
}

func StaticText_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("StaticText_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func StaticText_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("StaticText_SetDragMode").Call(obj, uintptr(value))
}

func StaticText_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetFocusControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetFocusControl").Call(obj)
	return ret
}

func StaticText_SetFocusControl(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetFocusControl").Call(obj, value)
}

func StaticText_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetFont").Call(obj)
	return ret
}

func StaticText_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetFont").Call(obj, value)
}

func StaticText_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetPopupMenu").Call(obj)
	return ret
}

func StaticText_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetPopupMenu").Call(obj, value)
}

func StaticText_GetShowAccelChar(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetShowAccelChar").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetShowAccelChar(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetShowAccelChar").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("StaticText_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func StaticText_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("StaticText_SetTabOrder").Call(obj, uintptr(value))
}

func StaticText_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetVisible").Call(obj, GoBoolToDBool(value))
}

func StaticText_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func StaticText_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("StaticText_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func StaticText_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetDockClientCount").Call(obj)
	return int32(ret)
}

func StaticText_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func StaticText_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetBrush").Call(obj)
	return ret
}

func StaticText_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetControlCount").Call(obj)
	return int32(ret)
}

func StaticText_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StaticText_GetHandle").Call(obj)
	return ret
}

func StaticText_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StaticText_GetParentWindow").Call(obj)
	return ret
}

func StaticText_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("StaticText_SetParentWindow").Call(obj, value)
}

func StaticText_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("StaticText_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func StaticText_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAction").Call(obj)
	return ret
}

func StaticText_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetAction").Call(obj, value)
}

func StaticText_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("StaticText_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("StaticText_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StaticText_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetClientHeight").Call(obj)
	return int32(ret)
}

func StaticText_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetClientHeight").Call(obj, uintptr(value))
}

func StaticText_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("StaticText_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("StaticText_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StaticText_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetClientWidth").Call(obj)
	return int32(ret)
}

func StaticText_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetClientWidth").Call(obj, uintptr(value))
}

func StaticText_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("StaticText_GetControlState").Call(obj)
	return TControlState(ret)
}

func StaticText_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("StaticText_SetControlState").Call(obj, uintptr(value))
}

func StaticText_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("StaticText_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func StaticText_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("StaticText_SetControlStyle").Call(obj, uintptr(value))
}

func StaticText_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("StaticText_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func StaticText_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetParent").Call(obj)
	return ret
}

func StaticText_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetParent").Call(obj, value)
}

func StaticText_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetLeft").Call(obj)
	return int32(ret)
}

func StaticText_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetLeft").Call(obj, uintptr(value))
}

func StaticText_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetTop").Call(obj)
	return int32(ret)
}

func StaticText_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetTop").Call(obj, uintptr(value))
}

func StaticText_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetWidth").Call(obj)
	return int32(ret)
}

func StaticText_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetWidth").Call(obj, uintptr(value))
}

func StaticText_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetHeight").Call(obj)
	return int32(ret)
}

func StaticText_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetHeight").Call(obj, uintptr(value))
}

func StaticText_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StaticText_GetCursor").Call(obj)
	return TCursor(ret)
}

func StaticText_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("StaticText_SetCursor").Call(obj, uintptr(value))
}

func StaticText_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("StaticText_SetHint").Call(obj, GoStrToDStr(value))
}

func StaticText_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetComponentCount").Call(obj)
	return int32(ret)
}

func StaticText_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StaticText_GetComponentIndex").Call(obj)
	return int32(ret)
}

func StaticText_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("StaticText_SetComponentIndex").Call(obj, uintptr(value))
}

func StaticText_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetOwner").Call(obj)
	return ret
}

func StaticText_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("StaticText_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func StaticText_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("StaticText_SetName").Call(obj, GoStrToDStr(value))
}

func StaticText_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("StaticText_GetTag").Call(obj)
	return int(ret)
}

func StaticText_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("StaticText_SetTag").Call(obj, uintptr(value))
}

func StaticText_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAnchorSideLeft").Call(obj)
	return ret
}

func StaticText_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetAnchorSideLeft").Call(obj, value)
}

func StaticText_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAnchorSideTop").Call(obj)
	return ret
}

func StaticText_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetAnchorSideTop").Call(obj, value)
}

func StaticText_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAnchorSideRight").Call(obj)
	return ret
}

func StaticText_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetAnchorSideRight").Call(obj, value)
}

func StaticText_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAnchorSideBottom").Call(obj)
	return ret
}

func StaticText_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetAnchorSideBottom").Call(obj, value)
}

func StaticText_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetChildSizing").Call(obj)
	return ret
}

func StaticText_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetChildSizing").Call(obj, value)
}

func StaticText_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetBorderSpacing").Call(obj)
	return ret
}

func StaticText_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("StaticText_SetBorderSpacing").Call(obj, value)
}

func StaticText_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func StaticText_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetControls").Call(obj, uintptr(Index))
	return ret
}

func StaticText_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func StaticText_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("StaticText_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func StaticText_StaticClassType() TClass {
	r, _, _ := getLazyProc("StaticText_StaticClassType").Call()
	return TClass(r)
}
