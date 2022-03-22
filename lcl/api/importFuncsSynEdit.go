package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TSynEdit ---------------------------

func SynEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_Create").Call(obj)
	return ret
}

func SynEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Free").Call(obj)
}

func SynEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Clear").Call(obj)
}

func SynEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_ClearSelection").Call(obj)
}

func SynEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_CopyToClipboard").Call(obj)
}

func SynEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_CutToClipboard").Call(obj)
}

func SynEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_PasteFromClipboard").Call(obj)
}

func SynEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Undo").Call(obj)
}

func SynEdit_Redo(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Redo").Call(obj)
}

func SynEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_SelectAll").Call(obj)
}

func SynEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func SynEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool) uintptr {
	ret, _, _ := getLazyProc("SynEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls))
	return ret
}

func SynEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_DisableAlign").Call(obj)
}

func SynEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_EnableAlign").Call(obj)
}

func SynEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("SynEdit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func SynEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("SynEdit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func SynEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("SynEdit_InsertControl").Call(obj, AControl)
}

func SynEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Invalidate").Call(obj)
}

func SynEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("SynEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func SynEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("SynEdit_RemoveControl").Call(obj, AControl)
}

func SynEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Realign").Call(obj)
}

func SynEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Repaint").Call(obj)
}

func SynEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("SynEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func SynEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("SynEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func SynEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("SynEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func SynEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_SetFocus").Call(obj)
}

func SynEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Update").Call(obj)
}

func SynEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_BringToFront").Call(obj)
}

func SynEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SynEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SynEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Hide").Call(obj)
}

func SynEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("SynEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func SynEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Refresh").Call(obj)
}

func SynEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SynEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SynEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_SendToBack").Call(obj)
}

func SynEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("SynEdit_Show").Call(obj)
}

func SynEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("SynEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func SynEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func SynEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("SynEdit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func SynEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SynEdit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SynEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SynEdit_Assign").Call(obj, Source)
}

func SynEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SynEdit_ClassType").Call(obj)
	return TClass(ret)
}

func SynEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func SynEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SynEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SynEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func SynEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SynEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SynEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SynEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SynEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SynEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func SynEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SynEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func SynEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("SynEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func SynEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("SynEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func SynEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("SynEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func SynEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SynEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SynEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("SynEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func SynEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("SynEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func SynEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("SynEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func SynEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("SynEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func SynEdit_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("SynEdit_SetAlign").Call(obj, uintptr(value))
}

func SynEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("SynEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func SynEdit_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("SynEdit_SetAnchors").Call(obj, uintptr(value))
}

func SynEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynEdit_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("SynEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func SynEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("SynEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func SynEdit_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("SynEdit_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func SynEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("SynEdit_SetBorderStyle").Call(obj, uintptr(value))
}

func SynEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("SynEdit_GetColor").Call(obj)
	return TColor(ret)
}

func SynEdit_SetColor(obj uintptr, value TColor) {
	getLazyProc("SynEdit_SetColor").Call(obj, uintptr(value))
}

func SynEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetConstraints").Call(obj)
	return ret
}

func SynEdit_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetConstraints").Call(obj, value)
}

func SynEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetEnabled(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetFont").Call(obj)
	return ret
}

func SynEdit_SetFont(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetFont").Call(obj, value)
}

func SynEdit_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetHideSelection(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetLines(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetLines").Call(obj)
	return ret
}

func SynEdit_SetLines(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetLines").Call(obj, value)
}

func SynEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetParentColor(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetParentFont(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetPopupMenu").Call(obj)
	return ret
}

func SynEdit_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetPopupMenu").Call(obj, value)
}

func SynEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetReadOnly(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("SynEdit_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func SynEdit_SetScrollBars(obj uintptr, value TScrollStyle) {
	_, _, _ = getLazyProc("SynEdit_SetScrollBars").Call(obj, uintptr(value))
}

func SynEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetShowHint(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("SynEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func SynEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("SynEdit_SetTabOrder").Call(obj, uintptr(value))
}

func SynEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetTabStop(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_SetText(obj uintptr, value string) {
	getLazyProc("SynEdit_SetText").Call(obj, GoStrToDStr(value))
}

func SynEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetVisible(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func SynEdit_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetCanRedo(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetCanRedo").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetModified(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetModified").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func SynEdit_SetSelStart(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetSelStart").Call(obj, uintptr(value))
}

func SynEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_SetSelText(obj uintptr, value string) {
	getLazyProc("SynEdit_SetSelText").Call(obj, GoStrToDStr(value))
}

func SynEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func SynEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetDockSite(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func SynEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetBrush").Call(obj)
	return ret
}

func SynEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func SynEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SynEdit_GetHandle").Call(obj)
	return HWND(ret)
}

func SynEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SynEdit_GetParentWindow").Call(obj)
	return HWND(ret)
}

func SynEdit_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("SynEdit_SetParentWindow").Call(obj, uintptr(value))
}

func SynEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("SynEdit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func SynEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAction").Call(obj)
	return ret
}

func SynEdit_SetAction(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetAction").Call(obj, value)
}

func SynEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("SynEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("SynEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func SynEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func SynEdit_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetClientHeight").Call(obj, uintptr(value))
}

func SynEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("SynEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("SynEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SynEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func SynEdit_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetClientWidth").Call(obj, uintptr(value))
}

func SynEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("SynEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func SynEdit_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("SynEdit_SetControlState").Call(obj, uintptr(value))
}

func SynEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("SynEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func SynEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("SynEdit_SetControlStyle").Call(obj, uintptr(value))
}

func SynEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("SynEdit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func SynEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetParent").Call(obj)
	return ret
}

func SynEdit_SetParent(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetParent").Call(obj, value)
}

func SynEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetLeft").Call(obj)
	return int32(ret)
}

func SynEdit_SetLeft(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetLeft").Call(obj, uintptr(value))
}

func SynEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetTop").Call(obj)
	return int32(ret)
}

func SynEdit_SetTop(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetTop").Call(obj, uintptr(value))
}

func SynEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetWidth").Call(obj)
	return int32(ret)
}

func SynEdit_SetWidth(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetWidth").Call(obj, uintptr(value))
}

func SynEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetHeight").Call(obj)
	return int32(ret)
}

func SynEdit_SetHeight(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetHeight").Call(obj, uintptr(value))
}

func SynEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("SynEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func SynEdit_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("SynEdit_SetCursor").Call(obj, uintptr(value))
}

func SynEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_SetHint(obj uintptr, value string) {
	getLazyProc("SynEdit_SetHint").Call(obj, GoStrToDStr(value))
}

func SynEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func SynEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SynEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SynEdit_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("SynEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func SynEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetOwner").Call(obj)
	return ret
}

func SynEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SynEdit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SynEdit_SetName(obj uintptr, value string) {
	getLazyProc("SynEdit_SetName").Call(obj, GoStrToDStr(value))
}

func SynEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SynEdit_GetTag").Call(obj)
	return int(ret)
}

func SynEdit_SetTag(obj uintptr, value int) {
	getLazyProc("SynEdit_SetTag").Call(obj, uintptr(value))
}

func SynEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func SynEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetAnchorSideLeft").Call(obj, value)
}

func SynEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func SynEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetAnchorSideTop").Call(obj, value)
}

func SynEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func SynEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetAnchorSideRight").Call(obj, value)
}

func SynEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func SynEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetAnchorSideBottom").Call(obj, value)
}

func SynEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetChildSizing").Call(obj)
	return ret
}

func SynEdit_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetChildSizing").Call(obj, value)
}

func SynEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetBorderSpacing").Call(obj)
	return ret
}

func SynEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("SynEdit_SetBorderSpacing").Call(obj, value)
}

func SynEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func SynEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func SynEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SynEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("SynEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func SynEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynEdit_StaticClassType").Call()
	return TClass(r)
}

func SynEdit_GetBlockIndent(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynEdit_GetBlockIndent").Call(obj)
	return int32(r)
}

func SynEdit_SetBlockIndent(obj uintptr, AValue int32) {
	_, _, _ = getLazyProc("SynEdit_SetBlockIndent").Call(obj, uintptr(AValue))
}

func SynEdit_SetOnMouseLink(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnMouseLink").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_SetOnClickLink(obj uintptr, fn interface{}) {
	getLazyProc("SynEdit_SetOnClickLink").Call(obj, addEventToMap(obj, fn))
}

func SynEdit_GetBlockTabIndent(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynEdit_GetBlockTabIndent").Call(obj)
	return int32(r)
}

func SynEdit_SetBlockTabIndent(obj uintptr, AValue int32) {
	_, _, _ = getLazyProc("SynEdit_SetBlockTabIndent").Call(obj, uintptr(AValue))
}

func SynEdit_GetRightEdge(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynEdit_GetRightEdge").Call(obj)
	return int32(r)
}

func SynEdit_SetRightEdge(obj uintptr, AValue int32) {
	_, _, _ = getLazyProc("SynEdit_SetRightEdge").Call(obj, uintptr(AValue))
}

func SynEdit_GetRightEdgeColor(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynEdit_GetRightEdgeColor").Call(obj)
	return TColor(r)
}

func SynEdit_SetRightEdgeColor(obj uintptr, AValue TColor) {
	_, _, _ = getLazyProc("SynEdit_SetRightEdgeColor").Call(obj, uintptr(AValue))
}

func SynEdit_GetTabWidth(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynEdit_GetTabWidth").Call(obj)
	return int32(r)
}

func SynEdit_SetTabWidth(obj uintptr, AValue int32) {
	_, _, _ = getLazyProc("SynEdit_SetTabWidth").Call(obj, uintptr(AValue))
}

func SynEdit_GetWantTabs(obj uintptr) bool {
	r, _, _ := getLazyProc("SynEdit_GetWantTabs").Call(obj)
	return DBoolToGoBool(r)
}

func SynEdit_SetWantTabs(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("SynEdit_SetWantTabs").Call(obj, GoBoolToDBool(AValue))
}

func SynEdit_GetOptions(obj uintptr) TSynEditorOptions {
	r, _, _ := getLazyProc("SynEdit_GetOptions").Call(obj)
	return TSynEditorOptions(r)
}

func SynEdit_SetOptions(obj uintptr, AValue TSynEditorOptions) {
	_, _, _ = getLazyProc("SynEdit_SetOptions").Call(obj, uintptr(AValue))
}

func SynEdit_GetOptions2(obj uintptr) TSynEditorOptions2 {
	r, _, _ := getLazyProc("SynEdit_GetOptions2").Call(obj)
	return TSynEditorOptions2(r)
}

func SynEdit_SetOptions2(obj uintptr, AValue TSynEditorOptions2) {
	_, _, _ = getLazyProc("SynEdit_SetOptions2").Call(obj, uintptr(AValue))
}

func SynEdit_GetMouseOptions(obj uintptr) TSynEditorMouseOptions {
	r, _, _ := getLazyProc("SynEdit_GetMouseOptions").Call(obj)
	return TSynEditorMouseOptions(r)
}

func SynEdit_SetMouseOptions(obj uintptr, AValue TSynEditorMouseOptions) {
	_, _, _ = getLazyProc("SynEdit_SetMouseOptions").Call(obj, uintptr(AValue))
}

func SynEdit_GetVisibleSpecialChars(obj uintptr) TSynVisibleSpecialChars {
	r, _, _ := getLazyProc("SynEdit_GetVisibleSpecialChars").Call(obj)
	return TSynVisibleSpecialChars(r)
}

func SynEdit_SetVisibleSpecialChars(obj uintptr, AValue TSynVisibleSpecialChars) {
	_, _, _ = getLazyProc("SynEdit_SetVisibleSpecialChars").Call(obj, uintptr(AValue))
}

func SynEdit_GetGutter(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynEdit_GetGutter").Call(obj)
	return r
}

func SynEdit_GetRightGutter(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynEdit_GetRightGutter").Call(obj)
	return r
}

func SynEdit_GetHighlighter(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynEdit_GetHighlighter").Call(obj)
	return r
}

func SynEdit_SetHighlighter(obj uintptr, AValue uintptr) {
	_, _, _ = getLazyProc("SynEdit_SetHighlighter").Call(obj, AValue)
}
