package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TColorListBox ---------------------------

func ColorListBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_Create").Call(obj)
	return ret
}

func ColorListBox_Free(obj uintptr) {
	getLazyProc("ColorListBox_Free").Call(obj)
}

func ColorListBox_AddItem(obj uintptr, Item string, AObject uintptr) {
	getLazyProc("ColorListBox_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ColorListBox_Clear(obj uintptr) {
	getLazyProc("ColorListBox_Clear").Call(obj)
}

func ColorListBox_ClearSelection(obj uintptr) {
	getLazyProc("ColorListBox_ClearSelection").Call(obj)
}

func ColorListBox_DeleteSelected(obj uintptr) {
	getLazyProc("ColorListBox_DeleteSelected").Call(obj)
}

func ColorListBox_ItemAtPos(obj uintptr, Pos TPoint, Existing bool) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ItemAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(Existing))
	return int32(ret)
}

func ColorListBox_ItemRect(obj uintptr, Index int32) TRect {
	var ret TRect
	getLazyProc("ColorListBox_ItemRect").Call(obj, uintptr(Index), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_SelectAll(obj uintptr) {
	getLazyProc("ColorListBox_SelectAll").Call(obj)
}

func ColorListBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ColorListBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ColorListBox_DisableAlign(obj uintptr) {
	getLazyProc("ColorListBox_DisableAlign").Call(obj)
}

func ColorListBox_EnableAlign(obj uintptr) {
	getLazyProc("ColorListBox_EnableAlign").Call(obj)
}

func ColorListBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ColorListBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ColorListBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ColorListBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ColorListBox_InsertControl").Call(obj, AControl)
}

func ColorListBox_Invalidate(obj uintptr) {
	getLazyProc("ColorListBox_Invalidate").Call(obj)
}

func ColorListBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ColorListBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ColorListBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ColorListBox_RemoveControl").Call(obj, AControl)
}

func ColorListBox_Realign(obj uintptr) {
	getLazyProc("ColorListBox_Realign").Call(obj)
}

func ColorListBox_Repaint(obj uintptr) {
	getLazyProc("ColorListBox_Repaint").Call(obj)
}

func ColorListBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ColorListBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ColorListBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ColorListBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ColorListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ColorListBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ColorListBox_SetFocus(obj uintptr) {
	getLazyProc("ColorListBox_SetFocus").Call(obj)
}

func ColorListBox_Update(obj uintptr) {
	getLazyProc("ColorListBox_Update").Call(obj)
}

func ColorListBox_BringToFront(obj uintptr) {
	getLazyProc("ColorListBox_BringToFront").Call(obj)
}

func ColorListBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ColorListBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorListBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_Hide(obj uintptr) {
	getLazyProc("ColorListBox_Hide").Call(obj)
}

func ColorListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ColorListBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ColorListBox_Refresh(obj uintptr) {
	getLazyProc("ColorListBox_Refresh").Call(obj)
}

func ColorListBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ColorListBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorListBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_SendToBack(obj uintptr) {
	getLazyProc("ColorListBox_SendToBack").Call(obj)
}

func ColorListBox_Show(obj uintptr) {
	getLazyProc("ColorListBox_Show").Call(obj)
}

func ColorListBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ColorListBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ColorListBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ColorListBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ColorListBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ColorListBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ColorListBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorListBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ColorListBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ColorListBox_Assign").Call(obj, Source)
}

func ColorListBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ColorListBox_ClassType").Call(obj)
	return TClass(ret)
}

func ColorListBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorListBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorListBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ColorListBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ColorListBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ColorListBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ColorListBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorListBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ColorListBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ColorListBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorListBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ColorListBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorListBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ColorListBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ColorListBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ColorListBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ColorListBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ColorListBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ColorListBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ColorListBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ColorListBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ColorListBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ColorListBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorListBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorListBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ColorListBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ColorListBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ColorListBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ColorListBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ColorListBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ColorListBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ColorListBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ColorListBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ColorListBox_SetAlign").Call(obj, uintptr(value))
}

func ColorListBox_GetDefaultColorColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorListBox_GetDefaultColorColor").Call(obj)
	return TColor(ret)
}

func ColorListBox_SetDefaultColorColor(obj uintptr, value TColor) {
	getLazyProc("ColorListBox_SetDefaultColorColor").Call(obj, uintptr(value))
}

func ColorListBox_GetNoneColorColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorListBox_GetNoneColorColor").Call(obj)
	return TColor(ret)
}

func ColorListBox_SetNoneColorColor(obj uintptr, value TColor) {
	getLazyProc("ColorListBox_SetNoneColorColor").Call(obj, uintptr(value))
}

func ColorListBox_GetSelected(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorListBox_GetSelected").Call(obj)
	return TColor(ret)
}

func ColorListBox_SetSelected(obj uintptr, value TColor) {
	getLazyProc("ColorListBox_SetSelected").Call(obj, uintptr(value))
}

func ColorListBox_GetStyle(obj uintptr) TColorBoxStyle {
	ret, _, _ := getLazyProc("ColorListBox_GetStyle").Call(obj)
	return TColorBoxStyle(ret)
}

func ColorListBox_SetStyle(obj uintptr, value TColorBoxStyle) {
	getLazyProc("ColorListBox_SetStyle").Call(obj, uintptr(value))
}

func ColorListBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ColorListBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ColorListBox_SetAnchors").Call(obj, uintptr(value))
}

func ColorListBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ColorListBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ColorListBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ColorListBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ColorListBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorListBox_GetColor").Call(obj)
	return TColor(ret)
}

func ColorListBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("ColorListBox_SetColor").Call(obj, uintptr(value))
}

func ColorListBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetConstraints").Call(obj)
	return ret
}

func ColorListBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetConstraints").Call(obj, value)
}

func ColorListBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetFont").Call(obj)
	return ret
}

func ColorListBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetFont").Call(obj, value)
}

func ColorListBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func ColorListBox_SetItemHeight(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetItemHeight").Call(obj, uintptr(value))
}

func ColorListBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetPopupMenu").Call(obj)
	return ret
}

func ColorListBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetPopupMenu").Call(obj, value)
}

func ColorListBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ColorListBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ColorListBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ColorListBox_SetTabOrder").Call(obj, uintptr(value))
}

func ColorListBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ColorListBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ColorListBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetCanvas").Call(obj)
	return ret
}

func ColorListBox_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetItems").Call(obj)
	return ret
}

func ColorListBox_SetItems(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetItems").Call(obj, value)
}

func ColorListBox_GetTopIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetTopIndex").Call(obj)
	return int32(ret)
}

func ColorListBox_SetTopIndex(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetTopIndex").Call(obj, uintptr(value))
}

func ColorListBox_GetMultiSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetMultiSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetMultiSelect(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetMultiSelect").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetSelCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetSelCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func ColorListBox_SetItemIndex(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetItemIndex").Call(obj, uintptr(value))
}

func ColorListBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetBrush").Call(obj)
	return ret
}

func ColorListBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ColorListBox_GetHandle").Call(obj)
	return HWND(ret)
}

func ColorListBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ColorListBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ColorListBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ColorListBox_SetParentWindow").Call(obj, uintptr(value))
}

func ColorListBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ColorListBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAction").Call(obj)
	return ret
}

func ColorListBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetAction").Call(obj, value)
}

func ColorListBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ColorListBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ColorListBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorListBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ColorListBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetClientHeight").Call(obj, uintptr(value))
}

func ColorListBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorListBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ColorListBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorListBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ColorListBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetClientWidth").Call(obj, uintptr(value))
}

func ColorListBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ColorListBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ColorListBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ColorListBox_SetControlState").Call(obj, uintptr(value))
}

func ColorListBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ColorListBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ColorListBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ColorListBox_SetControlStyle").Call(obj, uintptr(value))
}

func ColorListBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorListBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorListBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetParent").Call(obj)
	return ret
}

func ColorListBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetParent").Call(obj, value)
}

func ColorListBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetLeft").Call(obj)
	return int32(ret)
}

func ColorListBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetLeft").Call(obj, uintptr(value))
}

func ColorListBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetTop").Call(obj)
	return int32(ret)
}

func ColorListBox_SetTop(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetTop").Call(obj, uintptr(value))
}

func ColorListBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetWidth").Call(obj)
	return int32(ret)
}

func ColorListBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetWidth").Call(obj, uintptr(value))
}

func ColorListBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetHeight").Call(obj)
	return int32(ret)
}

func ColorListBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetHeight").Call(obj, uintptr(value))
}

func ColorListBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ColorListBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ColorListBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ColorListBox_SetCursor").Call(obj, uintptr(value))
}

func ColorListBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorListBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ColorListBox_SetHint(obj uintptr, value string) {
	getLazyProc("ColorListBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ColorListBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorListBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ColorListBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ColorListBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ColorListBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetOwner").Call(obj)
	return ret
}

func ColorListBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorListBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorListBox_SetName(obj uintptr, value string) {
	getLazyProc("ColorListBox_SetName").Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ColorListBox_GetTag").Call(obj)
	return int(ret)
}

func ColorListBox_SetTag(obj uintptr, value int) {
	getLazyProc("ColorListBox_SetTag").Call(obj, uintptr(value))
}

func ColorListBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ColorListBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetAnchorSideLeft").Call(obj, value)
}

func ColorListBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ColorListBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetAnchorSideTop").Call(obj, value)
}

func ColorListBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ColorListBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetAnchorSideRight").Call(obj, value)
}

func ColorListBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ColorListBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetAnchorSideBottom").Call(obj, value)
}

func ColorListBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetChildSizing").Call(obj)
	return ret
}

func ColorListBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetChildSizing").Call(obj, value)
}

func ColorListBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetBorderSpacing").Call(obj)
	return ret
}

func ColorListBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ColorListBox_SetBorderSpacing").Call(obj, value)
}

func ColorListBox_GetColors(obj uintptr, Index int32) TColor {
	ret, _, _ := getLazyProc("ColorListBox_GetColors").Call(obj, uintptr(Index))
	return TColor(ret)
}

func ColorListBox_GetColorNames(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("ColorListBox_GetColorNames").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func ColorListBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ColorListBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ColorListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ColorListBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ColorListBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ColorListBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ColorListBox_StaticClassType").Call()
	return TClass(r)
}
