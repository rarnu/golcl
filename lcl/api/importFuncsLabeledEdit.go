package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TLabeledEdit ---------------------------

func LabeledEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_Create").Call(obj)
	return ret
}

func LabeledEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Free").Call(obj)
}

func LabeledEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func LabeledEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Clear").Call(obj)
}

func LabeledEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_ClearSelection").Call(obj)
}

func LabeledEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_CopyToClipboard").Call(obj)
}

func LabeledEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_CutToClipboard").Call(obj)
}

func LabeledEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_PasteFromClipboard").Call(obj)
}

func LabeledEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Undo").Call(obj)
}

func LabeledEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SelectAll").Call(obj)
}

func LabeledEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func LabeledEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func LabeledEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_DisableAlign").Call(obj)
}

func LabeledEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_EnableAlign").Call(obj)
}

func LabeledEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func LabeledEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("LabeledEdit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func LabeledEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_InsertControl").Call(obj, AControl)
}

func LabeledEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Invalidate").Call(obj)
}

func LabeledEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("LabeledEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func LabeledEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_RemoveControl").Call(obj, AControl)
}

func LabeledEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Realign").Call(obj)
}

func LabeledEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Repaint").Call(obj)
}

func LabeledEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("LabeledEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func LabeledEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("LabeledEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func LabeledEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetFocus").Call(obj)
}

func LabeledEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Update").Call(obj)
}

func LabeledEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_BringToFront").Call(obj)
}

func LabeledEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LabeledEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LabeledEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Hide").Call(obj)
}

func LabeledEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("LabeledEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func LabeledEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Refresh").Call(obj)
}

func LabeledEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LabeledEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LabeledEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SendToBack").Call(obj)
}

func LabeledEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Show").Call(obj)
}

func LabeledEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("LabeledEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func LabeledEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("LabeledEdit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func LabeledEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func LabeledEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_Assign").Call(obj, Source)
}

func LabeledEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("LabeledEdit_ClassType").Call(obj)
	return TClass(ret)
}

func LabeledEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func LabeledEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("LabeledEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func LabeledEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func LabeledEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func LabeledEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func LabeledEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func LabeledEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func LabeledEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func LabeledEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func LabeledEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("LabeledEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func LabeledEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LabeledEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("LabeledEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func LabeledEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("LabeledEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func LabeledEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("LabeledEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func LabeledEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("LabeledEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func LabeledEdit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("LabeledEdit_SetAlignment").Call(obj, uintptr(value))
}

func LabeledEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func LabeledEdit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("LabeledEdit_SetAnchors").Call(obj, uintptr(value))
}

func LabeledEdit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetAutoSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetAutoSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetAutoSelect").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("LabeledEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func LabeledEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("LabeledEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func LabeledEdit_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("LabeledEdit_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func LabeledEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("LabeledEdit_SetBorderStyle").Call(obj, uintptr(value))
}

func LabeledEdit_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("LabeledEdit_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func LabeledEdit_SetCharCase(obj uintptr, value TEditCharCase) {
	_, _, _ = getLazyProc("LabeledEdit_SetCharCase").Call(obj, uintptr(value))
}

func LabeledEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("LabeledEdit_GetColor").Call(obj)
	return TColor(ret)
}

func LabeledEdit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("LabeledEdit_SetColor").Call(obj, uintptr(value))
}

func LabeledEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetConstraints").Call(obj)
	return ret
}

func LabeledEdit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetConstraints").Call(obj, value)
}

func LabeledEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("LabeledEdit_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func LabeledEdit_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("LabeledEdit_SetDragCursor").Call(obj, uintptr(value))
}

func LabeledEdit_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("LabeledEdit_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func LabeledEdit_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("LabeledEdit_SetDragMode").Call(obj, uintptr(value))
}

func LabeledEdit_GetEditLabel(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetEditLabel").Call(obj)
	return ret
}

func LabeledEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetFont").Call(obj)
	return ret
}

func LabeledEdit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetFont").Call(obj, value)
}

func LabeledEdit_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetLabelPosition(obj uintptr) TLabelPosition {
	ret, _, _ := getLazyProc("LabeledEdit_GetLabelPosition").Call(obj)
	return TLabelPosition(ret)
}

func LabeledEdit_SetLabelPosition(obj uintptr, value TLabelPosition) {
	_, _, _ = getLazyProc("LabeledEdit_SetLabelPosition").Call(obj, uintptr(value))
}

func LabeledEdit_GetLabelSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetLabelSpacing").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetLabelSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetLabelSpacing").Call(obj, uintptr(value))
}

func LabeledEdit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetMaxLength").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetMaxLength").Call(obj, uintptr(value))
}

func LabeledEdit_GetNumbersOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetNumbersOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetNumbersOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetNumbersOnly").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetPasswordChar(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("LabeledEdit_GetPasswordChar").Call(obj)
	return uint16(ret)
}

func LabeledEdit_SetPasswordChar(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("LabeledEdit_SetPasswordChar").Call(obj, uintptr(value))
}

func LabeledEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetPopupMenu").Call(obj)
	return ret
}

func LabeledEdit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetPopupMenu").Call(obj, value)
}

func LabeledEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("LabeledEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func LabeledEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("LabeledEdit_SetTabOrder").Call(obj, uintptr(value))
}

func LabeledEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("LabeledEdit_SetText").Call(obj, GoStrToDStr(value))
}

func LabeledEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("LabeledEdit_SetTextHint").Call(obj, GoStrToDStr(value))
}

func LabeledEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LabeledEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func LabeledEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetModified").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetSelLength").Call(obj, uintptr(value))
}

func LabeledEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetSelStart").Call(obj, uintptr(value))
}

func LabeledEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("LabeledEdit_SetSelText").Call(obj, GoStrToDStr(value))
}

func LabeledEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func LabeledEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func LabeledEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetBrush").Call(obj)
	return ret
}

func LabeledEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func LabeledEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("LabeledEdit_GetHandle").Call(obj)
	return ret
}

func LabeledEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("LabeledEdit_GetParentWindow").Call(obj)
	return ret
}

func LabeledEdit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("LabeledEdit_SetParentWindow").Call(obj, value)
}

func LabeledEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LabeledEdit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func LabeledEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAction").Call(obj)
	return ret
}

func LabeledEdit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetAction").Call(obj, value)
}

func LabeledEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("LabeledEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func LabeledEdit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("LabeledEdit_SetAlign").Call(obj, uintptr(value))
}

func LabeledEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("LabeledEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("LabeledEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func LabeledEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetClientHeight").Call(obj, uintptr(value))
}

func LabeledEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LabeledEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("LabeledEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LabeledEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetClientWidth").Call(obj, uintptr(value))
}

func LabeledEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("LabeledEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func LabeledEdit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("LabeledEdit_SetControlState").Call(obj, uintptr(value))
}

func LabeledEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("LabeledEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func LabeledEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("LabeledEdit_SetControlStyle").Call(obj, uintptr(value))
}

func LabeledEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("LabeledEdit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func LabeledEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetParent").Call(obj)
	return ret
}

func LabeledEdit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetParent").Call(obj, value)
}

func LabeledEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetLeft").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetLeft").Call(obj, uintptr(value))
}

func LabeledEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetTop").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetTop").Call(obj, uintptr(value))
}

func LabeledEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetWidth").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetWidth").Call(obj, uintptr(value))
}

func LabeledEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetHeight").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetHeight").Call(obj, uintptr(value))
}

func LabeledEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("LabeledEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func LabeledEdit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("LabeledEdit_SetCursor").Call(obj, uintptr(value))
}

func LabeledEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("LabeledEdit_SetHint").Call(obj, GoStrToDStr(value))
}

func LabeledEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func LabeledEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LabeledEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func LabeledEdit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LabeledEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func LabeledEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetOwner").Call(obj)
	return ret
}

func LabeledEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("LabeledEdit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func LabeledEdit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("LabeledEdit_SetName").Call(obj, GoStrToDStr(value))
}

func LabeledEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("LabeledEdit_GetTag").Call(obj)
	return int(ret)
}

func LabeledEdit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("LabeledEdit_SetTag").Call(obj, uintptr(value))
}

func LabeledEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func LabeledEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetAnchorSideLeft").Call(obj, value)
}

func LabeledEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func LabeledEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetAnchorSideTop").Call(obj, value)
}

func LabeledEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func LabeledEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetAnchorSideRight").Call(obj, value)
}

func LabeledEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func LabeledEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetAnchorSideBottom").Call(obj, value)
}

func LabeledEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetChildSizing").Call(obj)
	return ret
}

func LabeledEdit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetChildSizing").Call(obj, value)
}

func LabeledEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetBorderSpacing").Call(obj)
	return ret
}

func LabeledEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LabeledEdit_SetBorderSpacing").Call(obj, value)
}

func LabeledEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func LabeledEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func LabeledEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func LabeledEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("LabeledEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func LabeledEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("LabeledEdit_StaticClassType").Call()
	return TClass(r)
}
