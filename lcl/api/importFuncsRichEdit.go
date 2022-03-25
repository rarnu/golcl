package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TRichEdit ---------------------------

func RichEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_Create").Call(obj)
	return ret
}

func RichEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Free").Call(obj)
}

func RichEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Clear").Call(obj)
}

func RichEdit_FindText(obj uintptr, SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
	ret, _, _ := getLazyProc("RichEdit_FindText").Call(obj, GoStrToDStr(SearchStr), uintptr(StartPos), uintptr(Length), uintptr(Options))
	return int32(ret)
}

func RichEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_ClearSelection").Call(obj)
}

func RichEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_CopyToClipboard").Call(obj)
}

func RichEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_CutToClipboard").Call(obj)
}

func RichEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_PasteFromClipboard").Call(obj)
}

func RichEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Undo").Call(obj)
}

func RichEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_SelectAll").Call(obj)
}

func RichEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func RichEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("RichEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func RichEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_DisableAlign").Call(obj)
}

func RichEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_EnableAlign").Call(obj)
}

func RichEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("RichEdit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func RichEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("RichEdit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func RichEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("RichEdit_InsertControl").Call(obj, AControl)
}

func RichEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Invalidate").Call(obj)
}

func RichEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("RichEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func RichEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("RichEdit_RemoveControl").Call(obj, AControl)
}

func RichEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Realign").Call(obj)
}

func RichEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Repaint").Call(obj)
}

func RichEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("RichEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func RichEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("RichEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func RichEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("RichEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func RichEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetFocus").Call(obj)
}

func RichEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Update").Call(obj)
}

func RichEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_BringToFront").Call(obj)
}

func RichEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Hide").Call(obj)
}

func RichEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("RichEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func RichEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Refresh").Call(obj)
}

func RichEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_SendToBack").Call(obj)
}

func RichEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("RichEdit_Show").Call(obj)
}

func RichEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("RichEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func RichEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func RichEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("RichEdit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func RichEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("RichEdit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func RichEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("RichEdit_Assign").Call(obj, Source)
}

func RichEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("RichEdit_ClassType").Call(obj)
	return TClass(ret)
}

func RichEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func RichEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("RichEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func RichEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func RichEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func RichEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("RichEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RichEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("RichEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RichEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("RichEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func RichEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("RichEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func RichEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("RichEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func RichEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("RichEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func RichEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("RichEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func RichEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RichEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RichEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("RichEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func RichEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("RichEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func RichEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("RichEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func RichEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("RichEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func RichEdit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("RichEdit_SetAlign").Call(obj, uintptr(value))
}

func RichEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("RichEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func RichEdit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("RichEdit_SetAlignment").Call(obj, uintptr(value))
}

func RichEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("RichEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func RichEdit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("RichEdit_SetAnchors").Call(obj, uintptr(value))
}

func RichEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("RichEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func RichEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("RichEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func RichEdit_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("RichEdit_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func RichEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("RichEdit_SetBorderStyle").Call(obj, uintptr(value))
}

func RichEdit_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetBorderWidth").Call(obj)
	return int32(ret)
}

func RichEdit_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetBorderWidth").Call(obj, uintptr(value))
}

func RichEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("RichEdit_GetColor").Call(obj)
	return TColor(ret)
}

func RichEdit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("RichEdit_SetColor").Call(obj, uintptr(value))
}

func RichEdit_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RichEdit_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func RichEdit_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("RichEdit_SetDragCursor").Call(obj, uintptr(value))
}

func RichEdit_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("RichEdit_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func RichEdit_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("RichEdit_SetDragKind").Call(obj, uintptr(value))
}

func RichEdit_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("RichEdit_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func RichEdit_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("RichEdit_SetDragMode").Call(obj, uintptr(value))
}

func RichEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetFont").Call(obj)
	return ret
}

func RichEdit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetFont").Call(obj, value)
}

func RichEdit_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetConstraints").Call(obj)
	return ret
}

func RichEdit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetConstraints").Call(obj, value)
}

func RichEdit_GetLines(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetLines").Call(obj)
	return ret
}

func RichEdit_SetLines(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetLines").Call(obj, value)
}

func RichEdit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetMaxLength").Call(obj)
	return int32(ret)
}

func RichEdit_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetMaxLength").Call(obj, uintptr(value))
}

func RichEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetPopupMenu").Call(obj)
	return ret
}

func RichEdit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetPopupMenu").Call(obj, value)
}

func RichEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("RichEdit_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func RichEdit_SetScrollBars(obj uintptr, value TScrollStyle) {
	_, _, _ = getLazyProc("RichEdit_SetScrollBars").Call(obj, uintptr(value))
}

func RichEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("RichEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func RichEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("RichEdit_SetTabOrder").Call(obj, uintptr(value))
}

func RichEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantTabs(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetWantTabs").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetWantTabs(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetWantTabs").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantReturns(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetWantReturns").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetWantReturns(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetWantReturns").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWordWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetWordWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetWordWrap(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetWordWrap").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetZoom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetZoom").Call(obj)
	return int32(ret)
}

func RichEdit_SetZoom(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetZoom").Call(obj, uintptr(value))
}

func RichEdit_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseWheel(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseWheelDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_SetOnMouseWheelUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RichEdit_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func RichEdit_GetDefAttributes(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetDefAttributes").Call(obj)
	return ret
}

func RichEdit_SetDefAttributes(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetDefAttributes").Call(obj, value)
}

func RichEdit_GetSelAttributes(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetSelAttributes").Call(obj)
	return ret
}

func RichEdit_SetSelAttributes(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetSelAttributes").Call(obj, value)
}

func RichEdit_GetParagraph(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetParagraph").Call(obj)
	return ret
}

func RichEdit_GetCaretPos(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_GetCaretPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_SetCaretPos(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("RichEdit_SetCaretPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetModified").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func RichEdit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetSelLength").Call(obj, uintptr(value))
}

func RichEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func RichEdit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetSelStart").Call(obj, uintptr(value))
}

func RichEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("RichEdit_SetSelText").Call(obj, GoStrToDStr(value))
}

func RichEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("RichEdit_SetText").Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("RichEdit_SetTextHint").Call(obj, GoStrToDStr(value))
}

func RichEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func RichEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func RichEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetBrush").Call(obj)
	return ret
}

func RichEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func RichEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RichEdit_GetHandle").Call(obj)
	return ret
}

func RichEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RichEdit_GetParentWindow").Call(obj)
	return ret
}

func RichEdit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("RichEdit_SetParentWindow").Call(obj, value)
}

func RichEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RichEdit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAction").Call(obj)
	return ret
}

func RichEdit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetAction").Call(obj, value)
}

func RichEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("RichEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("RichEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func RichEdit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetClientHeight").Call(obj, uintptr(value))
}

func RichEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RichEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("RichEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RichEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func RichEdit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetClientWidth").Call(obj, uintptr(value))
}

func RichEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("RichEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func RichEdit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("RichEdit_SetControlState").Call(obj, uintptr(value))
}

func RichEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("RichEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func RichEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("RichEdit_SetControlStyle").Call(obj, uintptr(value))
}

func RichEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("RichEdit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func RichEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetParent").Call(obj)
	return ret
}

func RichEdit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetParent").Call(obj, value)
}

func RichEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetLeft").Call(obj)
	return int32(ret)
}

func RichEdit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetLeft").Call(obj, uintptr(value))
}

func RichEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetTop").Call(obj)
	return int32(ret)
}

func RichEdit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetTop").Call(obj, uintptr(value))
}

func RichEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetWidth").Call(obj)
	return int32(ret)
}

func RichEdit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetWidth").Call(obj, uintptr(value))
}

func RichEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetHeight").Call(obj)
	return int32(ret)
}

func RichEdit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetHeight").Call(obj, uintptr(value))
}

func RichEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RichEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func RichEdit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("RichEdit_SetCursor").Call(obj, uintptr(value))
}

func RichEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("RichEdit_SetHint").Call(obj, GoStrToDStr(value))
}

func RichEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func RichEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RichEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func RichEdit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RichEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func RichEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetOwner").Call(obj)
	return ret
}

func RichEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("RichEdit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func RichEdit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("RichEdit_SetName").Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("RichEdit_GetTag").Call(obj)
	return int(ret)
}

func RichEdit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("RichEdit_SetTag").Call(obj, uintptr(value))
}

func RichEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func RichEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetAnchorSideLeft").Call(obj, value)
}

func RichEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func RichEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetAnchorSideTop").Call(obj, value)
}

func RichEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func RichEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetAnchorSideRight").Call(obj, value)
}

func RichEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func RichEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetAnchorSideBottom").Call(obj, value)
}

func RichEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetChildSizing").Call(obj)
	return ret
}

func RichEdit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetChildSizing").Call(obj, value)
}

func RichEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetBorderSpacing").Call(obj)
	return ret
}

func RichEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RichEdit_SetBorderSpacing").Call(obj, value)
}

func RichEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func RichEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func RichEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func RichEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("RichEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func RichEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("RichEdit_StaticClassType").Call()
	return TClass(r)
}
