package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TEdit ---------------------------

func Edit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_Create").Call(obj)
	return ret
}

func Edit_Free(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Free").Call(obj)
}

func Edit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Clear").Call(obj)
}

func Edit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("Edit_ClearSelection").Call(obj)
}

func Edit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("Edit_CopyToClipboard").Call(obj)
}

func Edit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("Edit_CutToClipboard").Call(obj)
}

func Edit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("Edit_PasteFromClipboard").Call(obj)
}

func Edit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Undo").Call(obj)
}

func Edit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("Edit_SelectAll").Call(obj)
}

func Edit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Edit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Edit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Edit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Edit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Edit_DisableAlign").Call(obj)
}

func Edit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Edit_EnableAlign").Call(obj)
}

func Edit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Edit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Edit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("Edit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Edit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Edit_InsertControl").Call(obj, AControl)
}

func Edit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Invalidate").Call(obj)
}

func Edit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("Edit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func Edit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Edit_RemoveControl").Call(obj, AControl)
}

func Edit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Realign").Call(obj)
}

func Edit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Repaint").Call(obj)
}

func Edit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("Edit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Edit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("Edit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Edit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Edit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Edit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("Edit_SetFocus").Call(obj)
}

func Edit_Update(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Update").Call(obj)
}

func Edit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Edit_BringToFront").Call(obj)
}

func Edit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Edit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Edit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Hide").Call(obj)
}

func Edit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Edit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Edit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Refresh").Call(obj)
}

func Edit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Edit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Edit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Edit_SendToBack").Call(obj)
}

func Edit_Show(obj uintptr) {
	_, _, _ = getLazyProc("Edit_Show").Call(obj)
}

func Edit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Edit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Edit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetTextLen").Call(obj)
	return int32(ret)
}

func Edit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Edit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Edit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Edit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Edit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Edit_Assign").Call(obj, Source)
}

func Edit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Edit_ClassType").Call(obj)
	return TClass(ret)
}

func Edit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_InstanceSize").Call(obj)
	return int32(ret)
}

func Edit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Edit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Edit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Edit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetHashCode").Call(obj)
	return int32(ret)
}

func Edit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Edit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Edit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Edit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Edit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Edit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Edit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Edit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Edit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Edit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Edit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Edit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Edit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Edit_AnchorClient").Call(obj, uintptr(ASpace))
}

func Edit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Edit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Edit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Edit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Edit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Edit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Edit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Edit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Edit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Edit_GetAlign").Call(obj)
	return TAlign(ret)
}

func Edit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Edit_SetAlign").Call(obj, uintptr(value))
}

func Edit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("Edit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func Edit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("Edit_SetAlignment").Call(obj, uintptr(value))
}

func Edit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Edit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Edit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Edit_SetAnchors").Call(obj, uintptr(value))
}

func Edit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetAutoSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetAutoSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetAutoSelect").Call(obj, GoBoolToDBool(value))
}

func Edit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Edit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Edit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Edit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Edit_SetBiDiMode").Call(obj, uintptr(value))
}

func Edit_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("Edit_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func Edit_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("Edit_SetBorderStyle").Call(obj, uintptr(value))
}

func Edit_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("Edit_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func Edit_SetCharCase(obj uintptr, value TEditCharCase) {
	_, _, _ = getLazyProc("Edit_SetCharCase").Call(obj, uintptr(value))
}

func Edit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Edit_GetColor").Call(obj)
	return TColor(ret)
}

func Edit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Edit_SetColor").Call(obj, uintptr(value))
}

func Edit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetConstraints").Call(obj)
	return ret
}

func Edit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetConstraints").Call(obj, value)
}

func Edit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Edit_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Edit_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Edit_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Edit_SetDragCursor").Call(obj, uintptr(value))
}

func Edit_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Edit_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Edit_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("Edit_SetDragKind").Call(obj, uintptr(value))
}

func Edit_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Edit_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Edit_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Edit_SetDragMode").Call(obj, uintptr(value))
}

func Edit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Edit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetFont").Call(obj)
	return ret
}

func Edit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetFont").Call(obj, value)
}

func Edit_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func Edit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetMaxLength").Call(obj)
	return int32(ret)
}

func Edit_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetMaxLength").Call(obj, uintptr(value))
}

func Edit_GetNumbersOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetNumbersOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetNumbersOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetNumbersOnly").Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Edit_GetPasswordChar(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("Edit_GetPasswordChar").Call(obj)
	return uint16(ret)
}

func Edit_SetPasswordChar(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("Edit_SetPasswordChar").Call(obj, uintptr(value))
}

func Edit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetPopupMenu").Call(obj)
	return ret
}

func Edit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetPopupMenu").Call(obj, value)
}

func Edit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func Edit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Edit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Edit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Edit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("Edit_SetTabOrder").Call(obj, uintptr(value))
}

func Edit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Edit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("Edit_SetText").Call(obj, GoStrToDStr(value))
}

func Edit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Edit_SetTextHint").Call(obj, GoStrToDStr(value))
}

func Edit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Edit_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Edit_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Edit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Edit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetModified").Call(obj, GoBoolToDBool(value))
}

func Edit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetSelLength").Call(obj)
	return int32(ret)
}

func Edit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetSelLength").Call(obj, uintptr(value))
}

func Edit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetSelStart").Call(obj)
	return int32(ret)
}

func Edit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetSelStart").Call(obj, uintptr(value))
}

func Edit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("Edit_SetSelText").Call(obj, GoStrToDStr(value))
}

func Edit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Edit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Edit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Edit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetBrush").Call(obj)
	return ret
}

func Edit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetControlCount").Call(obj)
	return int32(ret)
}

func Edit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Edit_GetHandle").Call(obj)
	return ret
}

func Edit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Edit_GetParentWindow").Call(obj)
	return ret
}

func Edit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Edit_SetParentWindow").Call(obj, value)
}

func Edit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Edit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Edit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAction").Call(obj)
	return ret
}

func Edit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetAction").Call(obj, value)
}

func Edit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Edit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Edit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Edit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetClientHeight").Call(obj)
	return int32(ret)
}

func Edit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetClientHeight").Call(obj, uintptr(value))
}

func Edit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Edit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Edit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Edit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetClientWidth").Call(obj)
	return int32(ret)
}

func Edit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetClientWidth").Call(obj, uintptr(value))
}

func Edit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Edit_GetControlState").Call(obj)
	return TControlState(ret)
}

func Edit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Edit_SetControlState").Call(obj, uintptr(value))
}

func Edit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Edit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Edit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Edit_SetControlStyle").Call(obj, uintptr(value))
}

func Edit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Edit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Edit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetParent").Call(obj)
	return ret
}

func Edit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetParent").Call(obj, value)
}

func Edit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetLeft").Call(obj)
	return int32(ret)
}

func Edit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetLeft").Call(obj, uintptr(value))
}

func Edit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetTop").Call(obj)
	return int32(ret)
}

func Edit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetTop").Call(obj, uintptr(value))
}

func Edit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetWidth").Call(obj)
	return int32(ret)
}

func Edit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetWidth").Call(obj, uintptr(value))
}

func Edit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetHeight").Call(obj)
	return int32(ret)
}

func Edit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetHeight").Call(obj, uintptr(value))
}

func Edit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Edit_GetCursor").Call(obj)
	return TCursor(ret)
}

func Edit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Edit_SetCursor").Call(obj, uintptr(value))
}

func Edit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Edit_SetHint").Call(obj, GoStrToDStr(value))
}

func Edit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetComponentCount").Call(obj)
	return int32(ret)
}

func Edit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Edit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Edit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Edit_SetComponentIndex").Call(obj, uintptr(value))
}

func Edit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetOwner").Call(obj)
	return ret
}

func Edit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Edit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Edit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Edit_SetName").Call(obj, GoStrToDStr(value))
}

func Edit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Edit_GetTag").Call(obj)
	return int(ret)
}

func Edit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Edit_SetTag").Call(obj, uintptr(value))
}

func Edit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAnchorSideLeft").Call(obj)
	return ret
}

func Edit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetAnchorSideLeft").Call(obj, value)
}

func Edit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAnchorSideTop").Call(obj)
	return ret
}

func Edit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetAnchorSideTop").Call(obj, value)
}

func Edit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAnchorSideRight").Call(obj)
	return ret
}

func Edit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetAnchorSideRight").Call(obj, value)
}

func Edit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAnchorSideBottom").Call(obj)
	return ret
}

func Edit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetAnchorSideBottom").Call(obj, value)
}

func Edit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetChildSizing").Call(obj)
	return ret
}

func Edit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetChildSizing").Call(obj, value)
}

func Edit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Edit_GetBorderSpacing").Call(obj)
	return ret
}

func Edit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Edit_SetBorderSpacing").Call(obj, value)
}

func Edit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Edit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Edit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Edit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Edit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Edit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Edit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Edit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Edit_StaticClassType() TClass {
	r, _, _ := getLazyProc("Edit_StaticClassType").Call()
	return TClass(r)
}
