package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TListBox ---------------------------

func ListBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_Create").Call(obj)
	return ret
}

func ListBox_Free(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Free").Call(obj)
}

func ListBox_AddItem(obj uintptr, Item string, AObject uintptr) {
	_, _, _ = getLazyProc("ListBox_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ListBox_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Clear").Call(obj)
}

func ListBox_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_ClearSelection").Call(obj)
}

func ListBox_DeleteSelected(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_DeleteSelected").Call(obj)
}

func ListBox_ItemAtPos(obj uintptr, Pos TPoint, Existing bool) int32 {
	ret, _, _ := getLazyProc("ListBox_ItemAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(Existing))
	return int32(ret)
}

func ListBox_ItemRect(obj uintptr, Index int32) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ListBox_ItemRect").Call(obj, uintptr(Index), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_SelectAll").Call(obj)
}

func ListBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ListBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ListBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ListBox_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_DisableAlign").Call(obj)
}

func ListBox_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_EnableAlign").Call(obj)
}

func ListBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ListBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ListBox_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("ListBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ListBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ListBox_InsertControl").Call(obj, AControl)
}

func ListBox_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Invalidate").Call(obj)
}

func ListBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("ListBox_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func ListBox_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ListBox_RemoveControl").Call(obj, AControl)
}

func ListBox_Realign(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Realign").Call(obj)
}

func ListBox_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Repaint").Call(obj)
}

func ListBox_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("ListBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ListBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("ListBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ListBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ListBox_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_SetFocus").Call(obj)
}

func ListBox_Update(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Update").Call(obj)
}

func ListBox_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_BringToFront").Call(obj)
}

func ListBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Hide").Call(obj)
}

func ListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ListBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ListBox_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Refresh").Call(obj)
}

func ListBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_SendToBack").Call(obj)
}

func ListBox_Show(obj uintptr) {
	_, _, _ = getLazyProc("ListBox_Show").Call(obj)
}

func ListBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ListBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ListBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ListBox_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ListBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ListBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ListBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ListBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListBox_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ListBox_Assign").Call(obj, Source)
}

func ListBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListBox_ClassType").Call(obj)
	return TClass(ret)
}

func ListBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ListBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ListBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ListBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ListBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ListBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ListBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ListBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ListBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ListBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ListBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ListBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ListBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ListBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ListBox_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ListBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ListBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ListBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ListBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ListBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ListBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ListBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ListBox_GetClickOnSelChange(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetClickOnSelChange").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetClickOnSelChange(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetClickOnSelChange").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetOptions(obj uintptr) TListBoxOptions {
	ret, _, _ := getLazyProc("ListBox_GetOptions").Call(obj)
	return TListBoxOptions(ret)
}

func ListBox_SetOptions(obj uintptr, value TListBoxOptions) {
	_, _, _ = getLazyProc("ListBox_SetOptions").Call(obj, uintptr(value))
}

func ListBox_GetTopIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetTopIndex").Call(obj)
	return int32(ret)
}

func ListBox_SetTopIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetTopIndex").Call(obj, uintptr(value))
}

func ListBox_GetStyle(obj uintptr) TListBoxStyle {
	ret, _, _ := getLazyProc("ListBox_GetStyle").Call(obj)
	return TListBoxStyle(ret)
}

func ListBox_SetStyle(obj uintptr, value TListBoxStyle) {
	_, _, _ = getLazyProc("ListBox_SetStyle").Call(obj, uintptr(value))
}

func ListBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ListBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ListBox_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ListBox_SetAlign").Call(obj, uintptr(value))
}

func ListBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ListBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ListBox_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ListBox_SetAnchors").Call(obj, uintptr(value))
}

func ListBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ListBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ListBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ListBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ListBox_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("ListBox_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func ListBox_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("ListBox_SetBorderStyle").Call(obj, uintptr(value))
}

func ListBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ListBox_GetColor").Call(obj)
	return TColor(ret)
}

func ListBox_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ListBox_SetColor").Call(obj, uintptr(value))
}

func ListBox_GetColumns(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetColumns").Call(obj)
	return int32(ret)
}

func ListBox_SetColumns(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetColumns").Call(obj, uintptr(value))
}

func ListBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetConstraints").Call(obj)
	return ret
}

func ListBox_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetConstraints").Call(obj, value)
}

func ListBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ListBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ListBox_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ListBox_SetDragCursor").Call(obj, uintptr(value))
}

func ListBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ListBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ListBox_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("ListBox_SetDragKind").Call(obj, uintptr(value))
}

func ListBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ListBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ListBox_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("ListBox_SetDragMode").Call(obj, uintptr(value))
}

func ListBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetExtendedSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetExtendedSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetExtendedSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetExtendedSelect").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetFont").Call(obj)
	return ret
}

func ListBox_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetFont").Call(obj, value)
}

func ListBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func ListBox_SetItemHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetItemHeight").Call(obj, uintptr(value))
}

func ListBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetItems").Call(obj)
	return ret
}

func ListBox_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetItems").Call(obj, value)
}

func ListBox_GetMultiSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetMultiSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetMultiSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetMultiSelect").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetPopupMenu").Call(obj)
	return ret
}

func ListBox_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetPopupMenu").Call(obj, value)
}

func ListBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetSorted(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetSorted").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetSorted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetSorted").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ListBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ListBox_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("ListBox_SetTabOrder").Call(obj, uintptr(value))
}

func ListBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ListBox_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnDrawItem").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMeasureItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMeasureItem").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ListBox_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ListBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetCanvas").Call(obj)
	return ret
}

func ListBox_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetCount").Call(obj)
	return int32(ret)
}

func ListBox_GetSelCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetSelCount").Call(obj)
	return int32(ret)
}

func ListBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func ListBox_SetItemIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetItemIndex").Call(obj, uintptr(value))
}

func ListBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ListBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ListBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetBrush").Call(obj)
	return ret
}

func ListBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ListBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ListBox_GetHandle").Call(obj)
	return ret
}

func ListBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ListBox_GetParentWindow").Call(obj)
	return ret
}

func ListBox_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("ListBox_SetParentWindow").Call(obj, value)
}

func ListBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ListBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAction").Call(obj)
	return ret
}

func ListBox_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetAction").Call(obj, value)
}

func ListBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ListBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ListBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ListBox_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetClientHeight").Call(obj, uintptr(value))
}

func ListBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ListBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ListBox_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetClientWidth").Call(obj, uintptr(value))
}

func ListBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ListBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ListBox_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ListBox_SetControlState").Call(obj, uintptr(value))
}

func ListBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ListBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ListBox_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ListBox_SetControlStyle").Call(obj, uintptr(value))
}

func ListBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ListBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetParent").Call(obj)
	return ret
}

func ListBox_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetParent").Call(obj, value)
}

func ListBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetLeft").Call(obj)
	return int32(ret)
}

func ListBox_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetLeft").Call(obj, uintptr(value))
}

func ListBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetTop").Call(obj)
	return int32(ret)
}

func ListBox_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetTop").Call(obj, uintptr(value))
}

func ListBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetWidth").Call(obj)
	return int32(ret)
}

func ListBox_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetWidth").Call(obj, uintptr(value))
}

func ListBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetHeight").Call(obj)
	return int32(ret)
}

func ListBox_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetHeight").Call(obj, uintptr(value))
}

func ListBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ListBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ListBox_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ListBox_SetCursor").Call(obj, uintptr(value))
}

func ListBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ListBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ListBox_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ListBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ListBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ListBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ListBox_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ListBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetOwner").Call(obj)
	return ret
}

func ListBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ListBox_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ListBox_SetName").Call(obj, GoStrToDStr(value))
}

func ListBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ListBox_GetTag").Call(obj)
	return int(ret)
}

func ListBox_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ListBox_SetTag").Call(obj, uintptr(value))
}

func ListBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ListBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetAnchorSideLeft").Call(obj, value)
}

func ListBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ListBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetAnchorSideTop").Call(obj, value)
}

func ListBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ListBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetAnchorSideRight").Call(obj, value)
}

func ListBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ListBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetAnchorSideBottom").Call(obj, value)
}

func ListBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetChildSizing").Call(obj)
	return ret
}

func ListBox_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetChildSizing").Call(obj, value)
}

func ListBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetBorderSpacing").Call(obj)
	return ret
}

func ListBox_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListBox_SetBorderSpacing").Call(obj, value)
}

func ListBox_GetSelected(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("ListBox_GetSelected").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func ListBox_SetSelected(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("ListBox_SetSelected").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func ListBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ListBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ListBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ListBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ListBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListBox_StaticClassType").Call()
	return TClass(r)
}
