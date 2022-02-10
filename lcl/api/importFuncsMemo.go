package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMemo ---------------------------

func Memo_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_Create").Call(obj)
	return ret
}

func Memo_Free(obj uintptr) {
	getLazyProc("Memo_Free").Call(obj)
}

func Memo_Append(obj uintptr, Value string) {
	getLazyProc("Memo_Append").Call(obj, GoStrToDStr(Value))
}

func Memo_Clear(obj uintptr) {
	getLazyProc("Memo_Clear").Call(obj)
}

func Memo_ClearSelection(obj uintptr) {
	getLazyProc("Memo_ClearSelection").Call(obj)
}

func Memo_CopyToClipboard(obj uintptr) {
	getLazyProc("Memo_CopyToClipboard").Call(obj)
}

func Memo_CutToClipboard(obj uintptr) {
	getLazyProc("Memo_CutToClipboard").Call(obj)
}

func Memo_PasteFromClipboard(obj uintptr) {
	getLazyProc("Memo_PasteFromClipboard").Call(obj)
}

func Memo_Undo(obj uintptr) {
	getLazyProc("Memo_Undo").Call(obj)
}

func Memo_SelectAll(obj uintptr) {
	getLazyProc("Memo_SelectAll").Call(obj)
}

func Memo_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Memo_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Memo_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Memo_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Memo_DisableAlign(obj uintptr) {
	getLazyProc("Memo_DisableAlign").Call(obj)
}

func Memo_EnableAlign(obj uintptr) {
	getLazyProc("Memo_EnableAlign").Call(obj)
}

func Memo_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Memo_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Memo_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("Memo_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Memo_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("Memo_InsertControl").Call(obj, AControl)
}

func Memo_Invalidate(obj uintptr) {
	getLazyProc("Memo_Invalidate").Call(obj)
}

func Memo_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("Memo_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func Memo_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("Memo_RemoveControl").Call(obj, AControl)
}

func Memo_Realign(obj uintptr) {
	getLazyProc("Memo_Realign").Call(obj)
}

func Memo_Repaint(obj uintptr) {
	getLazyProc("Memo_Repaint").Call(obj)
}

func Memo_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("Memo_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Memo_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("Memo_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Memo_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("Memo_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Memo_SetFocus(obj uintptr) {
	getLazyProc("Memo_SetFocus").Call(obj)
}

func Memo_Update(obj uintptr) {
	getLazyProc("Memo_Update").Call(obj)
}

func Memo_BringToFront(obj uintptr) {
	getLazyProc("Memo_BringToFront").Call(obj)
}

func Memo_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Memo_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Memo_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_Hide(obj uintptr) {
	getLazyProc("Memo_Hide").Call(obj)
}

func Memo_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Memo_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Memo_Refresh(obj uintptr) {
	getLazyProc("Memo_Refresh").Call(obj)
}

func Memo_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Memo_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Memo_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_SendToBack(obj uintptr) {
	getLazyProc("Memo_SendToBack").Call(obj)
}

func Memo_Show(obj uintptr) {
	getLazyProc("Memo_Show").Call(obj)
}

func Memo_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Memo_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Memo_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetTextLen").Call(obj)
	return int32(ret)
}

func Memo_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("Memo_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Memo_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Memo_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Memo_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Memo_Assign").Call(obj, Source)
}

func Memo_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Memo_ClassType").Call(obj)
	return TClass(ret)
}

func Memo_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_InstanceSize").Call(obj)
	return int32(ret)
}

func Memo_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Memo_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Memo_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Memo_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetHashCode").Call(obj)
	return int32(ret)
}

func Memo_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Memo_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Memo_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Memo_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Memo_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Memo_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Memo_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Memo_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Memo_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("Memo_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Memo_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("Memo_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Memo_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("Memo_AnchorClient").Call(obj, uintptr(ASpace))
}

func Memo_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Memo_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Memo_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("Memo_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Memo_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("Memo_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Memo_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("Memo_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Memo_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Memo_GetAlign").Call(obj)
	return TAlign(ret)
}

func Memo_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("Memo_SetAlign").Call(obj, uintptr(value))
}

func Memo_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("Memo_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func Memo_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("Memo_SetAlignment").Call(obj, uintptr(value))
}

func Memo_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Memo_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Memo_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("Memo_SetAnchors").Call(obj, uintptr(value))
}

func Memo_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Memo_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Memo_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("Memo_SetBiDiMode").Call(obj, uintptr(value))
}

func Memo_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("Memo_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func Memo_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("Memo_SetBorderStyle").Call(obj, uintptr(value))
}

func Memo_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("Memo_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func Memo_SetCharCase(obj uintptr, value TEditCharCase) {
	getLazyProc("Memo_SetCharCase").Call(obj, uintptr(value))
}

func Memo_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Memo_GetColor").Call(obj)
	return TColor(ret)
}

func Memo_SetColor(obj uintptr, value TColor) {
	getLazyProc("Memo_SetColor").Call(obj, uintptr(value))
}

func Memo_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetConstraints").Call(obj)
	return ret
}

func Memo_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetConstraints").Call(obj, value)
}

func Memo_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("Memo_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Memo_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Memo_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Memo_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("Memo_SetDragCursor").Call(obj, uintptr(value))
}

func Memo_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Memo_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Memo_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("Memo_SetDragKind").Call(obj, uintptr(value))
}

func Memo_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Memo_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Memo_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("Memo_SetDragMode").Call(obj, uintptr(value))
}

func Memo_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetEnabled(obj uintptr, value bool) {
	getLazyProc("Memo_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Memo_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetFont").Call(obj)
	return ret
}

func Memo_SetFont(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetFont").Call(obj, value)
}

func Memo_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetHideSelection(obj uintptr, value bool) {
	getLazyProc("Memo_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func Memo_GetLines(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetLines").Call(obj)
	return ret
}

func Memo_SetLines(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetLines").Call(obj, value)
}

func Memo_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetMaxLength").Call(obj)
	return int32(ret)
}

func Memo_SetMaxLength(obj uintptr, value int32) {
	getLazyProc("Memo_SetMaxLength").Call(obj, uintptr(value))
}

func Memo_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetParentColor(obj uintptr, value bool) {
	getLazyProc("Memo_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("Memo_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetParentFont(obj uintptr, value bool) {
	getLazyProc("Memo_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("Memo_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Memo_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetPopupMenu").Call(obj)
	return ret
}

func Memo_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetPopupMenu").Call(obj, value)
}

func Memo_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetReadOnly(obj uintptr, value bool) {
	getLazyProc("Memo_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func Memo_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("Memo_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func Memo_SetScrollBars(obj uintptr, value TScrollStyle) {
	getLazyProc("Memo_SetScrollBars").Call(obj, uintptr(value))
}

func Memo_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetShowHint(obj uintptr, value bool) {
	getLazyProc("Memo_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Memo_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Memo_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Memo_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("Memo_SetTabOrder").Call(obj, uintptr(value))
}

func Memo_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetTabStop(obj uintptr, value bool) {
	getLazyProc("Memo_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Memo_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetVisible(obj uintptr, value bool) {
	getLazyProc("Memo_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantReturns(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetWantReturns").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetWantReturns(obj uintptr, value bool) {
	getLazyProc("Memo_SetWantReturns").Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantTabs(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetWantTabs").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetWantTabs(obj uintptr, value bool) {
	getLazyProc("Memo_SetWantTabs").Call(obj, GoBoolToDBool(value))
}

func Memo_GetWordWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetWordWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetWordWrap(obj uintptr, value bool) {
	getLazyProc("Memo_SetWordWrap").Call(obj, GoBoolToDBool(value))
}

func Memo_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Memo_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("Memo_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Memo_GetCaretPos(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Memo_GetCaretPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_SetCaretPos(obj uintptr, value TPoint) {
	getLazyProc("Memo_SetCaretPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Memo_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetModified(obj uintptr, value bool) {
	getLazyProc("Memo_SetModified").Call(obj, GoBoolToDBool(value))
}

func Memo_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetSelLength").Call(obj)
	return int32(ret)
}

func Memo_SetSelLength(obj uintptr, value int32) {
	getLazyProc("Memo_SetSelLength").Call(obj, uintptr(value))
}

func Memo_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetSelStart").Call(obj)
	return int32(ret)
}

func Memo_SetSelStart(obj uintptr, value int32) {
	getLazyProc("Memo_SetSelStart").Call(obj, uintptr(value))
}

func Memo_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_SetSelText(obj uintptr, value string) {
	getLazyProc("Memo_SetSelText").Call(obj, GoStrToDStr(value))
}

func Memo_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_SetText(obj uintptr, value string) {
	getLazyProc("Memo_SetText").Call(obj, GoStrToDStr(value))
}

func Memo_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_SetTextHint(obj uintptr, value string) {
	getLazyProc("Memo_SetTextHint").Call(obj, GoStrToDStr(value))
}

func Memo_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Memo_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetDockSite(obj uintptr, value bool) {
	getLazyProc("Memo_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Memo_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Memo_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetBrush").Call(obj)
	return ret
}

func Memo_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetControlCount").Call(obj)
	return int32(ret)
}

func Memo_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Memo_GetHandle").Call(obj)
	return HWND(ret)
}

func Memo_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Memo_GetParentWindow").Call(obj)
	return HWND(ret)
}

func Memo_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("Memo_SetParentWindow").Call(obj, uintptr(value))
}

func Memo_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("Memo_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Memo_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAction").Call(obj)
	return ret
}

func Memo_SetAction(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetAction").Call(obj, value)
}

func Memo_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Memo_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("Memo_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Memo_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetClientHeight").Call(obj)
	return int32(ret)
}

func Memo_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("Memo_SetClientHeight").Call(obj, uintptr(value))
}

func Memo_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Memo_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Memo_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Memo_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetClientWidth").Call(obj)
	return int32(ret)
}

func Memo_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("Memo_SetClientWidth").Call(obj, uintptr(value))
}

func Memo_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Memo_GetControlState").Call(obj)
	return TControlState(ret)
}

func Memo_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("Memo_SetControlState").Call(obj, uintptr(value))
}

func Memo_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Memo_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Memo_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("Memo_SetControlStyle").Call(obj, uintptr(value))
}

func Memo_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Memo_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Memo_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetParent").Call(obj)
	return ret
}

func Memo_SetParent(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetParent").Call(obj, value)
}

func Memo_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetLeft").Call(obj)
	return int32(ret)
}

func Memo_SetLeft(obj uintptr, value int32) {
	getLazyProc("Memo_SetLeft").Call(obj, uintptr(value))
}

func Memo_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetTop").Call(obj)
	return int32(ret)
}

func Memo_SetTop(obj uintptr, value int32) {
	getLazyProc("Memo_SetTop").Call(obj, uintptr(value))
}

func Memo_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetWidth").Call(obj)
	return int32(ret)
}

func Memo_SetWidth(obj uintptr, value int32) {
	getLazyProc("Memo_SetWidth").Call(obj, uintptr(value))
}

func Memo_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetHeight").Call(obj)
	return int32(ret)
}

func Memo_SetHeight(obj uintptr, value int32) {
	getLazyProc("Memo_SetHeight").Call(obj, uintptr(value))
}

func Memo_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Memo_GetCursor").Call(obj)
	return TCursor(ret)
}

func Memo_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("Memo_SetCursor").Call(obj, uintptr(value))
}

func Memo_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_SetHint(obj uintptr, value string) {
	getLazyProc("Memo_SetHint").Call(obj, GoStrToDStr(value))
}

func Memo_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetComponentCount").Call(obj)
	return int32(ret)
}

func Memo_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Memo_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Memo_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("Memo_SetComponentIndex").Call(obj, uintptr(value))
}

func Memo_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetOwner").Call(obj)
	return ret
}

func Memo_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Memo_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Memo_SetName(obj uintptr, value string) {
	getLazyProc("Memo_SetName").Call(obj, GoStrToDStr(value))
}

func Memo_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Memo_GetTag").Call(obj)
	return int(ret)
}

func Memo_SetTag(obj uintptr, value int) {
	getLazyProc("Memo_SetTag").Call(obj, uintptr(value))
}

func Memo_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAnchorSideLeft").Call(obj)
	return ret
}

func Memo_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetAnchorSideLeft").Call(obj, value)
}

func Memo_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAnchorSideTop").Call(obj)
	return ret
}

func Memo_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetAnchorSideTop").Call(obj, value)
}

func Memo_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAnchorSideRight").Call(obj)
	return ret
}

func Memo_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetAnchorSideRight").Call(obj, value)
}

func Memo_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAnchorSideBottom").Call(obj)
	return ret
}

func Memo_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetAnchorSideBottom").Call(obj, value)
}

func Memo_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetChildSizing").Call(obj)
	return ret
}

func Memo_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetChildSizing").Call(obj, value)
}

func Memo_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Memo_GetBorderSpacing").Call(obj)
	return ret
}

func Memo_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("Memo_SetBorderSpacing").Call(obj, value)
}

func Memo_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Memo_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Memo_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Memo_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Memo_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Memo_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Memo_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Memo_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Memo_StaticClassType() TClass {
	r, _, _ := getLazyProc("Memo_StaticClassType").Call()
	return TClass(r)
}
