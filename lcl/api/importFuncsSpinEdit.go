package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TSpinEdit ---------------------------

func SpinEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_Create").Call(obj)
	return ret
}

func SpinEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Free").Call(obj)
}

func SpinEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Clear").Call(obj)
}

func SpinEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_ClearSelection").Call(obj)
}

func SpinEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_CopyToClipboard").Call(obj)
}

func SpinEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_CutToClipboard").Call(obj)
}

func SpinEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_PasteFromClipboard").Call(obj)
}

func SpinEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Undo").Call(obj)
}

func SpinEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SelectAll").Call(obj)
}

func SpinEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func SpinEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func SpinEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_DisableAlign").Call(obj)
}

func SpinEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_EnableAlign").Call(obj)
}

func SpinEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func SpinEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("SpinEdit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func SpinEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("SpinEdit_InsertControl").Call(obj, AControl)
}

func SpinEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Invalidate").Call(obj)
}

func SpinEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("SpinEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func SpinEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("SpinEdit_RemoveControl").Call(obj, AControl)
}

func SpinEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Realign").Call(obj)
}

func SpinEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Repaint").Call(obj)
}

func SpinEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("SpinEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func SpinEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("SpinEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func SpinEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("SpinEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func SpinEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetFocus").Call(obj)
}

func SpinEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Update").Call(obj)
}

func SpinEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_BringToFront").Call(obj)
}

func SpinEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpinEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpinEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Hide").Call(obj)
}

func SpinEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("SpinEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func SpinEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Refresh").Call(obj)
}

func SpinEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpinEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpinEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SendToBack").Call(obj)
}

func SpinEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Show").Call(obj)
}

func SpinEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("SpinEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func SpinEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func SpinEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("SpinEdit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func SpinEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SpinEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SpinEdit_Assign").Call(obj, Source)
}

func SpinEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SpinEdit_ClassType").Call(obj)
	return TClass(ret)
}

func SpinEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func SpinEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SpinEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SpinEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func SpinEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SpinEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SpinEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SpinEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SpinEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SpinEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func SpinEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SpinEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func SpinEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("SpinEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func SpinEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("SpinEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func SpinEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("SpinEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func SpinEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpinEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpinEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("SpinEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func SpinEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("SpinEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func SpinEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("SpinEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func SpinEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func SpinEdit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("SpinEdit_SetAnchors").Call(obj, uintptr(value))
}

func SpinEdit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetAutoSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetAutoSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetAutoSelect").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("SpinEdit_GetColor").Call(obj)
	return TColor(ret)
}

func SpinEdit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SpinEdit_SetColor").Call(obj, uintptr(value))
}

func SpinEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetConstraints").Call(obj)
	return ret
}

func SpinEdit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetConstraints").Call(obj, value)
}

func SpinEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetFont").Call(obj)
	return ret
}

func SpinEdit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetFont").Call(obj, value)
}

func SpinEdit_GetIncrement(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetIncrement").Call(obj)
	return int32(ret)
}

func SpinEdit_SetIncrement(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetIncrement").Call(obj, uintptr(value))
}

func SpinEdit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetMaxLength").Call(obj)
	return int32(ret)
}

func SpinEdit_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetMaxLength").Call(obj, uintptr(value))
}

func SpinEdit_GetMaxValue(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetMaxValue").Call(obj)
	return int32(ret)
}

func SpinEdit_SetMaxValue(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetMaxValue").Call(obj, uintptr(value))
}

func SpinEdit_GetMinValue(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetMinValue").Call(obj)
	return int32(ret)
}

func SpinEdit_SetMinValue(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetMinValue").Call(obj, uintptr(value))
}

func SpinEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetPopupMenu").Call(obj)
	return ret
}

func SpinEdit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetPopupMenu").Call(obj, value)
}

func SpinEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("SpinEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func SpinEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("SpinEdit_SetTabOrder").Call(obj, uintptr(value))
}

func SpinEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetValue(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetValue").Call(obj)
	return int32(ret)
}

func SpinEdit_SetValue(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetValue").Call(obj, uintptr(value))
}

func SpinEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpinEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func SpinEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("SpinEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func SpinEdit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("SpinEdit_SetAlignment").Call(obj, uintptr(value))
}

func SpinEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetModified").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func SpinEdit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetSelLength").Call(obj, uintptr(value))
}

func SpinEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func SpinEdit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetSelStart").Call(obj, uintptr(value))
}

func SpinEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpinEdit_SetSelText").Call(obj, GoStrToDStr(value))
}

func SpinEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpinEdit_SetText").Call(obj, GoStrToDStr(value))
}

func SpinEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpinEdit_SetTextHint").Call(obj, GoStrToDStr(value))
}

func SpinEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func SpinEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func SpinEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetBrush").Call(obj)
	return ret
}

func SpinEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func SpinEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SpinEdit_GetHandle").Call(obj)
	return ret
}

func SpinEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("SpinEdit_GetParentWindow").Call(obj)
	return ret
}

func SpinEdit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("SpinEdit_SetParentWindow").Call(obj, value)
}

func SpinEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpinEdit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func SpinEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAction").Call(obj)
	return ret
}

func SpinEdit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetAction").Call(obj, value)
}

func SpinEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("SpinEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func SpinEdit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("SpinEdit_SetAlign").Call(obj, uintptr(value))
}

func SpinEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("SpinEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func SpinEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("SpinEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func SpinEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("SpinEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("SpinEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func SpinEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func SpinEdit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetClientHeight").Call(obj, uintptr(value))
}

func SpinEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpinEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("SpinEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpinEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func SpinEdit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetClientWidth").Call(obj, uintptr(value))
}

func SpinEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("SpinEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func SpinEdit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("SpinEdit_SetControlState").Call(obj, uintptr(value))
}

func SpinEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("SpinEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func SpinEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("SpinEdit_SetControlStyle").Call(obj, uintptr(value))
}

func SpinEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpinEdit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func SpinEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetParent").Call(obj)
	return ret
}

func SpinEdit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetParent").Call(obj, value)
}

func SpinEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetLeft").Call(obj)
	return int32(ret)
}

func SpinEdit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetLeft").Call(obj, uintptr(value))
}

func SpinEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetTop").Call(obj)
	return int32(ret)
}

func SpinEdit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetTop").Call(obj, uintptr(value))
}

func SpinEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetWidth").Call(obj)
	return int32(ret)
}

func SpinEdit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetWidth").Call(obj, uintptr(value))
}

func SpinEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetHeight").Call(obj)
	return int32(ret)
}

func SpinEdit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetHeight").Call(obj, uintptr(value))
}

func SpinEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("SpinEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func SpinEdit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("SpinEdit_SetCursor").Call(obj, uintptr(value))
}

func SpinEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpinEdit_SetHint").Call(obj, GoStrToDStr(value))
}

func SpinEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func SpinEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpinEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SpinEdit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpinEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func SpinEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetOwner").Call(obj)
	return ret
}

func SpinEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SpinEdit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SpinEdit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpinEdit_SetName").Call(obj, GoStrToDStr(value))
}

func SpinEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SpinEdit_GetTag").Call(obj)
	return int(ret)
}

func SpinEdit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("SpinEdit_SetTag").Call(obj, uintptr(value))
}

func SpinEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func SpinEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetAnchorSideLeft").Call(obj, value)
}

func SpinEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func SpinEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetAnchorSideTop").Call(obj, value)
}

func SpinEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func SpinEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetAnchorSideRight").Call(obj, value)
}

func SpinEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func SpinEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetAnchorSideBottom").Call(obj, value)
}

func SpinEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetChildSizing").Call(obj)
	return ret
}

func SpinEdit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetChildSizing").Call(obj, value)
}

func SpinEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetBorderSpacing").Call(obj)
	return ret
}

func SpinEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpinEdit_SetBorderSpacing").Call(obj, value)
}

func SpinEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func SpinEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func SpinEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SpinEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("SpinEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func SpinEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("SpinEdit_StaticClassType").Call()
	return TClass(r)
}
