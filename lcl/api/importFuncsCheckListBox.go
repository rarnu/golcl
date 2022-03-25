package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCheckListBox ---------------------------

func CheckListBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_Create").Call(obj)
	return ret
}

func CheckListBox_Free(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Free").Call(obj)
}

func CheckListBox_CheckAll(obj uintptr, AState TCheckBoxState, AllowGrayed bool, AllowDisabled bool) {
	_, _, _ = getLazyProc("CheckListBox_CheckAll").Call(obj, uintptr(AState), GoBoolToDBool(AllowGrayed), GoBoolToDBool(AllowDisabled))
}

func CheckListBox_AddItem(obj uintptr, Item string, AObject uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func CheckListBox_Clear(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Clear").Call(obj)
}

func CheckListBox_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_ClearSelection").Call(obj)
}

func CheckListBox_DeleteSelected(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_DeleteSelected").Call(obj)
}

func CheckListBox_ItemAtPos(obj uintptr, Pos TPoint, Existing bool) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ItemAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(Existing))
	return int32(ret)
}

func CheckListBox_ItemRect(obj uintptr, Index int32) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckListBox_ItemRect").Call(obj, uintptr(Index), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SelectAll").Call(obj)
}

func CheckListBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func CheckListBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func CheckListBox_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_DisableAlign").Call(obj)
}

func CheckListBox_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_EnableAlign").Call(obj)
}

func CheckListBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func CheckListBox_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("CheckListBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func CheckListBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckListBox_InsertControl").Call(obj, AControl)
}

func CheckListBox_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Invalidate").Call(obj)
}

func CheckListBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("CheckListBox_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func CheckListBox_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckListBox_RemoveControl").Call(obj, AControl)
}

func CheckListBox_Realign(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Realign").Call(obj)
}

func CheckListBox_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Repaint").Call(obj)
}

func CheckListBox_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("CheckListBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func CheckListBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("CheckListBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func CheckListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("CheckListBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func CheckListBox_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetFocus").Call(obj)
}

func CheckListBox_Update(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Update").Call(obj)
}

func CheckListBox_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_BringToFront").Call(obj)
}

func CheckListBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckListBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckListBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_Hide(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Hide").Call(obj)
}

func CheckListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("CheckListBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func CheckListBox_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Refresh").Call(obj)
}

func CheckListBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckListBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckListBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SendToBack").Call(obj)
}

func CheckListBox_Show(obj uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Show").Call(obj)
}

func CheckListBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("CheckListBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func CheckListBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetTextLen").Call(obj)
	return int32(ret)
}

func CheckListBox_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("CheckListBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func CheckListBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func CheckListBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckListBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CheckListBox_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("CheckListBox_Assign").Call(obj, Source)
}

func CheckListBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CheckListBox_ClassType").Call(obj)
	return TClass(ret)
}

func CheckListBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckListBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckListBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_InstanceSize").Call(obj)
	return int32(ret)
}

func CheckListBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CheckListBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CheckListBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetHashCode").Call(obj)
	return int32(ret)
}

func CheckListBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckListBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CheckListBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckListBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckListBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func CheckListBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func CheckListBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckListBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func CheckListBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("CheckListBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func CheckListBox_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("CheckListBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func CheckListBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckListBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckListBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("CheckListBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func CheckListBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("CheckListBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func CheckListBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("CheckListBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func CheckListBox_SetOnClickCheck(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnClickCheck").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("CheckListBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func CheckListBox_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("CheckListBox_SetAlign").Call(obj, uintptr(value))
}

func CheckListBox_GetAllowGrayed(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetAllowGrayed").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetAllowGrayed(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetAllowGrayed").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func CheckListBox_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("CheckListBox_SetAnchors").Call(obj, uintptr(value))
}

func CheckListBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("CheckListBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func CheckListBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("CheckListBox_SetBiDiMode").Call(obj, uintptr(value))
}

func CheckListBox_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("CheckListBox_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func CheckListBox_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("CheckListBox_SetBorderStyle").Call(obj, uintptr(value))
}

func CheckListBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CheckListBox_GetColor").Call(obj)
	return TColor(ret)
}

func CheckListBox_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("CheckListBox_SetColor").Call(obj, uintptr(value))
}

func CheckListBox_GetColumns(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetColumns").Call(obj)
	return int32(ret)
}

func CheckListBox_SetColumns(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetColumns").Call(obj, uintptr(value))
}

func CheckListBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetConstraints").Call(obj)
	return ret
}

func CheckListBox_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetConstraints").Call(obj, value)
}

func CheckListBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckListBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func CheckListBox_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckListBox_SetDragCursor").Call(obj, uintptr(value))
}

func CheckListBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("CheckListBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func CheckListBox_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("CheckListBox_SetDragMode").Call(obj, uintptr(value))
}

func CheckListBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetFont").Call(obj)
	return ret
}

func CheckListBox_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetFont").Call(obj, value)
}

func CheckListBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func CheckListBox_SetItemHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetItemHeight").Call(obj, uintptr(value))
}

func CheckListBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetItems").Call(obj)
	return ret
}

func CheckListBox_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetItems").Call(obj, value)
}

func CheckListBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetPopupMenu").Call(obj)
	return ret
}

func CheckListBox_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetPopupMenu").Call(obj, value)
}

func CheckListBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetSorted(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetSorted").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetSorted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetSorted").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetStyle(obj uintptr) TListBoxStyle {
	ret, _, _ := getLazyProc("CheckListBox_GetStyle").Call(obj)
	return TListBoxStyle(ret)
}

func CheckListBox_SetStyle(obj uintptr, value TListBoxStyle) {
	_, _, _ = getLazyProc("CheckListBox_SetStyle").Call(obj, uintptr(value))
}

func CheckListBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("CheckListBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func CheckListBox_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("CheckListBox_SetTabOrder").Call(obj, uintptr(value))
}

func CheckListBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMeasureItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMeasureItem").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckListBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func CheckListBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetCanvas").Call(obj)
	return ret
}

func CheckListBox_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetTopIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetTopIndex").Call(obj)
	return int32(ret)
}

func CheckListBox_SetTopIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetTopIndex").Call(obj, uintptr(value))
}

func CheckListBox_GetMultiSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetMultiSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetMultiSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetMultiSelect").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetSelCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetSelCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func CheckListBox_SetItemIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetItemIndex").Call(obj, uintptr(value))
}

func CheckListBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetBrush").Call(obj)
	return ret
}

func CheckListBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetControlCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckListBox_GetHandle").Call(obj)
	return ret
}

func CheckListBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckListBox_GetParentWindow").Call(obj)
	return ret
}

func CheckListBox_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("CheckListBox_SetParentWindow").Call(obj, value)
}

func CheckListBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func CheckListBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAction").Call(obj)
	return ret
}

func CheckListBox_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetAction").Call(obj, value)
}

func CheckListBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckListBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("CheckListBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckListBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func CheckListBox_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetClientHeight").Call(obj, uintptr(value))
}

func CheckListBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckListBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckListBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckListBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func CheckListBox_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetClientWidth").Call(obj, uintptr(value))
}

func CheckListBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("CheckListBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func CheckListBox_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("CheckListBox_SetControlState").Call(obj, uintptr(value))
}

func CheckListBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("CheckListBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func CheckListBox_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("CheckListBox_SetControlStyle").Call(obj, uintptr(value))
}

func CheckListBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckListBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetParent").Call(obj)
	return ret
}

func CheckListBox_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetParent").Call(obj, value)
}

func CheckListBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetLeft").Call(obj)
	return int32(ret)
}

func CheckListBox_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetLeft").Call(obj, uintptr(value))
}

func CheckListBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetTop").Call(obj)
	return int32(ret)
}

func CheckListBox_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetTop").Call(obj, uintptr(value))
}

func CheckListBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetWidth").Call(obj)
	return int32(ret)
}

func CheckListBox_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetWidth").Call(obj, uintptr(value))
}

func CheckListBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetHeight").Call(obj)
	return int32(ret)
}

func CheckListBox_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetHeight").Call(obj, uintptr(value))
}

func CheckListBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckListBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func CheckListBox_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckListBox_SetCursor").Call(obj, uintptr(value))
}

func CheckListBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckListBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func CheckListBox_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckListBox_SetHint").Call(obj, GoStrToDStr(value))
}

func CheckListBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func CheckListBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckListBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func CheckListBox_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckListBox_SetComponentIndex").Call(obj, uintptr(value))
}

func CheckListBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetOwner").Call(obj)
	return ret
}

func CheckListBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckListBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckListBox_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckListBox_SetName").Call(obj, GoStrToDStr(value))
}

func CheckListBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("CheckListBox_GetTag").Call(obj)
	return int(ret)
}

func CheckListBox_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("CheckListBox_SetTag").Call(obj, uintptr(value))
}

func CheckListBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func CheckListBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetAnchorSideLeft").Call(obj, value)
}

func CheckListBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchorSideTop").Call(obj)
	return ret
}

func CheckListBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetAnchorSideTop").Call(obj, value)
}

func CheckListBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchorSideRight").Call(obj)
	return ret
}

func CheckListBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetAnchorSideRight").Call(obj, value)
}

func CheckListBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func CheckListBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetAnchorSideBottom").Call(obj, value)
}

func CheckListBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetChildSizing").Call(obj)
	return ret
}

func CheckListBox_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetChildSizing").Call(obj, value)
}

func CheckListBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetBorderSpacing").Call(obj)
	return ret
}

func CheckListBox_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckListBox_SetBorderSpacing").Call(obj, value)
}

func CheckListBox_GetChecked(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetChecked").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckListBox_SetChecked(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetChecked").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckListBox_GetItemEnabled(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetItemEnabled").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckListBox_SetItemEnabled(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetItemEnabled").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckListBox_GetState(obj uintptr, Index int32) TCheckBoxState {
	ret, _, _ := getLazyProc("CheckListBox_GetState").Call(obj, uintptr(Index))
	return TCheckBoxState(ret)
}

func CheckListBox_SetState(obj uintptr, Index int32, value TCheckBoxState) {
	_, _, _ = getLazyProc("CheckListBox_SetState").Call(obj, uintptr(Index), uintptr(value))
}

func CheckListBox_GetHeader(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetHeader").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckListBox_SetHeader(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetHeader").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckListBox_GetSelected(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckListBox_GetSelected").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckListBox_SetSelected(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckListBox_SetSelected").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckListBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func CheckListBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func CheckListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func CheckListBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("CheckListBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func CheckListBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("CheckListBox_StaticClassType").Call()
	return TClass(r)
}
