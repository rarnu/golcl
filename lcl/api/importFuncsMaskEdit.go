package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMaskEdit ---------------------------

func MaskEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_Create").Call(obj)
	return ret
}

func MaskEdit_Free(obj uintptr) {
	getLazyProc("MaskEdit_Free").Call(obj)
}

func MaskEdit_ValidateEdit(obj uintptr) {
	getLazyProc("MaskEdit_ValidateEdit").Call(obj)
}

func MaskEdit_Clear(obj uintptr) {
	getLazyProc("MaskEdit_Clear").Call(obj)
}

func MaskEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func MaskEdit_ClearSelection(obj uintptr) {
	getLazyProc("MaskEdit_ClearSelection").Call(obj)
}

func MaskEdit_CopyToClipboard(obj uintptr) {
	getLazyProc("MaskEdit_CopyToClipboard").Call(obj)
}

func MaskEdit_CutToClipboard(obj uintptr) {
	getLazyProc("MaskEdit_CutToClipboard").Call(obj)
}

func MaskEdit_PasteFromClipboard(obj uintptr) {
	getLazyProc("MaskEdit_PasteFromClipboard").Call(obj)
}

func MaskEdit_Undo(obj uintptr) {
	getLazyProc("MaskEdit_Undo").Call(obj)
}

func MaskEdit_SelectAll(obj uintptr) {
	getLazyProc("MaskEdit_SelectAll").Call(obj)
}

func MaskEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func MaskEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func MaskEdit_DisableAlign(obj uintptr) {
	getLazyProc("MaskEdit_DisableAlign").Call(obj)
}

func MaskEdit_EnableAlign(obj uintptr) {
	getLazyProc("MaskEdit_EnableAlign").Call(obj)
}

func MaskEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func MaskEdit_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("MaskEdit_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func MaskEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("MaskEdit_InsertControl").Call(obj, AControl)
}

func MaskEdit_Invalidate(obj uintptr) {
	getLazyProc("MaskEdit_Invalidate").Call(obj)
}

func MaskEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("MaskEdit_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func MaskEdit_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("MaskEdit_RemoveControl").Call(obj, AControl)
}

func MaskEdit_Realign(obj uintptr) {
	getLazyProc("MaskEdit_Realign").Call(obj)
}

func MaskEdit_Repaint(obj uintptr) {
	getLazyProc("MaskEdit_Repaint").Call(obj)
}

func MaskEdit_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("MaskEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func MaskEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("MaskEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func MaskEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("MaskEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func MaskEdit_SetFocus(obj uintptr) {
	getLazyProc("MaskEdit_SetFocus").Call(obj)
}

func MaskEdit_Update(obj uintptr) {
	getLazyProc("MaskEdit_Update").Call(obj)
}

func MaskEdit_BringToFront(obj uintptr) {
	getLazyProc("MaskEdit_BringToFront").Call(obj)
}

func MaskEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("MaskEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("MaskEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_Hide(obj uintptr) {
	getLazyProc("MaskEdit_Hide").Call(obj)
}

func MaskEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("MaskEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func MaskEdit_Refresh(obj uintptr) {
	getLazyProc("MaskEdit_Refresh").Call(obj)
}

func MaskEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("MaskEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("MaskEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_SendToBack(obj uintptr) {
	getLazyProc("MaskEdit_SendToBack").Call(obj)
}

func MaskEdit_Show(obj uintptr) {
	getLazyProc("MaskEdit_Show").Call(obj)
}

func MaskEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("MaskEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func MaskEdit_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("MaskEdit_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func MaskEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func MaskEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_Assign(obj uintptr, Source uintptr) {
	getLazyProc("MaskEdit_Assign").Call(obj, Source)
}

func MaskEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MaskEdit_ClassType").Call(obj)
	return TClass(ret)
}

func MaskEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func MaskEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MaskEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MaskEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func MaskEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("MaskEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MaskEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("MaskEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MaskEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("MaskEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func MaskEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("MaskEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func MaskEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("MaskEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func MaskEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("MaskEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func MaskEdit_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("MaskEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func MaskEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MaskEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MaskEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("MaskEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func MaskEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("MaskEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func MaskEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("MaskEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func MaskEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("MaskEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func MaskEdit_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("MaskEdit_SetAlign").Call(obj, uintptr(value))
}

func MaskEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("MaskEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func MaskEdit_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("MaskEdit_SetAlignment").Call(obj, uintptr(value))
}

func MaskEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func MaskEdit_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("MaskEdit_SetAnchors").Call(obj, uintptr(value))
}

func MaskEdit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetAutoSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetAutoSelect(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetAutoSelect").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("MaskEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func MaskEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("MaskEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func MaskEdit_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("MaskEdit_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func MaskEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("MaskEdit_SetBorderStyle").Call(obj, uintptr(value))
}

func MaskEdit_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("MaskEdit_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func MaskEdit_SetCharCase(obj uintptr, value TEditCharCase) {
	getLazyProc("MaskEdit_SetCharCase").Call(obj, uintptr(value))
}

func MaskEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("MaskEdit_GetColor").Call(obj)
	return TColor(ret)
}

func MaskEdit_SetColor(obj uintptr, value TColor) {
	getLazyProc("MaskEdit_SetColor").Call(obj, uintptr(value))
}

func MaskEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetConstraints").Call(obj)
	return ret
}

func MaskEdit_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetConstraints").Call(obj, value)
}

func MaskEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("MaskEdit_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func MaskEdit_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("MaskEdit_SetDragCursor").Call(obj, uintptr(value))
}

func MaskEdit_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("MaskEdit_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func MaskEdit_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("MaskEdit_SetDragKind").Call(obj, uintptr(value))
}

func MaskEdit_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("MaskEdit_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func MaskEdit_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("MaskEdit_SetDragMode").Call(obj, uintptr(value))
}

func MaskEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetEnabled(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetFont").Call(obj)
	return ret
}

func MaskEdit_SetFont(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetFont").Call(obj, value)
}

func MaskEdit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetMaxLength").Call(obj)
	return int32(ret)
}

func MaskEdit_SetMaxLength(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetMaxLength").Call(obj, uintptr(value))
}

func MaskEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetParentColor(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetParentFont(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetPasswordChar(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("MaskEdit_GetPasswordChar").Call(obj)
	return uint16(ret)
}

func MaskEdit_SetPasswordChar(obj uintptr, value uint16) {
	getLazyProc("MaskEdit_SetPasswordChar").Call(obj, uintptr(value))
}

func MaskEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetPopupMenu").Call(obj)
	return ret
}

func MaskEdit_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetPopupMenu").Call(obj, value)
}

func MaskEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetReadOnly(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetShowHint(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("MaskEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func MaskEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("MaskEdit_SetTabOrder").Call(obj, uintptr(value))
}

func MaskEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetTabStop(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetText(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetText").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetTextHint(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetTextHint").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetVisible(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetVisible").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("MaskEdit_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func MaskEdit_GetIsMasked(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetIsMasked").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetEditText(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetEditText").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetEditText(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetEditText").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetCanUndo").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetModified").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetModified(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetModified").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func MaskEdit_SetSelLength(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetSelLength").Call(obj, uintptr(value))
}

func MaskEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func MaskEdit_SetSelStart(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetSelStart").Call(obj, uintptr(value))
}

func MaskEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetSelText(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetSelText").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func MaskEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetDockSite(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func MaskEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetBrush").Call(obj)
	return ret
}

func MaskEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func MaskEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MaskEdit_GetHandle").Call(obj)
	return HWND(ret)
}

func MaskEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MaskEdit_GetParentWindow").Call(obj)
	return HWND(ret)
}

func MaskEdit_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("MaskEdit_SetParentWindow").Call(obj, uintptr(value))
}

func MaskEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("MaskEdit_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func MaskEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAction").Call(obj)
	return ret
}

func MaskEdit_SetAction(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetAction").Call(obj, value)
}

func MaskEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("MaskEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("MaskEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MaskEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func MaskEdit_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetClientHeight").Call(obj, uintptr(value))
}

func MaskEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("MaskEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("MaskEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MaskEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func MaskEdit_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetClientWidth").Call(obj, uintptr(value))
}

func MaskEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("MaskEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func MaskEdit_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("MaskEdit_SetControlState").Call(obj, uintptr(value))
}

func MaskEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("MaskEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func MaskEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("MaskEdit_SetControlStyle").Call(obj, uintptr(value))
}

func MaskEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("MaskEdit_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func MaskEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetParent").Call(obj)
	return ret
}

func MaskEdit_SetParent(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetParent").Call(obj, value)
}

func MaskEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetLeft").Call(obj)
	return int32(ret)
}

func MaskEdit_SetLeft(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetLeft").Call(obj, uintptr(value))
}

func MaskEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetTop").Call(obj)
	return int32(ret)
}

func MaskEdit_SetTop(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetTop").Call(obj, uintptr(value))
}

func MaskEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetWidth").Call(obj)
	return int32(ret)
}

func MaskEdit_SetWidth(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetWidth").Call(obj, uintptr(value))
}

func MaskEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetHeight").Call(obj)
	return int32(ret)
}

func MaskEdit_SetHeight(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetHeight").Call(obj, uintptr(value))
}

func MaskEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("MaskEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func MaskEdit_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("MaskEdit_SetCursor").Call(obj, uintptr(value))
}

func MaskEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetHint(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetHint").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func MaskEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MaskEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func MaskEdit_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("MaskEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func MaskEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetOwner").Call(obj)
	return ret
}

func MaskEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("MaskEdit_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func MaskEdit_SetName(obj uintptr, value string) {
	getLazyProc("MaskEdit_SetName").Call(obj, GoStrToDStr(value))
}

func MaskEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("MaskEdit_GetTag").Call(obj)
	return int(ret)
}

func MaskEdit_SetTag(obj uintptr, value int) {
	getLazyProc("MaskEdit_SetTag").Call(obj, uintptr(value))
}

func MaskEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func MaskEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetAnchorSideLeft").Call(obj, value)
}

func MaskEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func MaskEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetAnchorSideTop").Call(obj, value)
}

func MaskEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func MaskEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetAnchorSideRight").Call(obj, value)
}

func MaskEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func MaskEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetAnchorSideBottom").Call(obj, value)
}

func MaskEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetChildSizing").Call(obj)
	return ret
}

func MaskEdit_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetChildSizing").Call(obj, value)
}

func MaskEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetBorderSpacing").Call(obj)
	return ret
}

func MaskEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("MaskEdit_SetBorderSpacing").Call(obj, value)
}

func MaskEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func MaskEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func MaskEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func MaskEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("MaskEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func MaskEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("MaskEdit_StaticClassType").Call()
	return TClass(r)
}
