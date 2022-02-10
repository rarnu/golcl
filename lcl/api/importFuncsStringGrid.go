package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TStringGrid ---------------------------

func StringGrid_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_Create").Call(obj)
	return ret
}

func StringGrid_Free(obj uintptr) {
	getLazyProc("StringGrid_Free").Call(obj)
}

func StringGrid_DeleteColRow(obj uintptr, IsColumn bool, index int32) {
	getLazyProc("StringGrid_DeleteColRow").Call(obj, GoBoolToDBool(IsColumn), uintptr(index))
}

func StringGrid_DeleteCol(obj uintptr, Index int32) {
	getLazyProc("StringGrid_DeleteCol").Call(obj, uintptr(Index))
}

func StringGrid_DeleteRow(obj uintptr, Index int32) {
	getLazyProc("StringGrid_DeleteRow").Call(obj, uintptr(Index))
}

func StringGrid_ExchangeColRow(obj uintptr, IsColumn bool, index int32, WithIndex int32) {
	getLazyProc("StringGrid_ExchangeColRow").Call(obj, GoBoolToDBool(IsColumn), uintptr(index), uintptr(WithIndex))
}

func StringGrid_InsertColRow(obj uintptr, IsColumn bool, index int32) {
	getLazyProc("StringGrid_InsertColRow").Call(obj, GoBoolToDBool(IsColumn), uintptr(index))
}

func StringGrid_MoveColRow(obj uintptr, IsColumn bool, FromIndex int32, ToIndex int32) {
	getLazyProc("StringGrid_MoveColRow").Call(obj, GoBoolToDBool(IsColumn), uintptr(FromIndex), uintptr(ToIndex))
}

func StringGrid_SortColRow(obj uintptr, IsColumn bool, Index int32, FromIndex int32, ToIndex int32) {
	getLazyProc("StringGrid_SortColRow").Call(obj, GoBoolToDBool(IsColumn), uintptr(Index), uintptr(FromIndex), uintptr(ToIndex))
}

func StringGrid_EditorByStyle(obj uintptr, Style TColumnButtonStyle) uintptr {
	ret, _, _ := getLazyProc("StringGrid_EditorByStyle").Call(obj, uintptr(Style))
	return ret
}

func StringGrid_EditorKeyDown(obj uintptr, Sender uintptr, Key *uint16, Shift TShiftState) {
	getLazyProc("StringGrid_EditorKeyDown").Call(obj, Sender, uintptr(unsafe.Pointer(Key)), uintptr(Shift))
}

func StringGrid_EditorKeyPress(obj uintptr, Sender uintptr, Key *uint16) {
	getLazyProc("StringGrid_EditorKeyPress").Call(obj, Sender, uintptr(unsafe.Pointer(Key)))
}

func StringGrid_EditorKeyUp(obj uintptr, Sender uintptr, key *uint16, shift TShiftState) {
	getLazyProc("StringGrid_EditorKeyUp").Call(obj, Sender, uintptr(unsafe.Pointer(key)), uintptr(shift))
}

func StringGrid_EditorTextChanged(obj uintptr, aCol int32, aRow int32, aText string) {
	getLazyProc("StringGrid_EditorTextChanged").Call(obj, uintptr(aCol), uintptr(aRow), GoStrToDStr(aText))
}

func StringGrid_EditingDone(obj uintptr) {
	getLazyProc("StringGrid_EditingDone").Call(obj)
}

func StringGrid_AutoAdjustColumns(obj uintptr) {
	getLazyProc("StringGrid_AutoAdjustColumns").Call(obj)
}

func StringGrid_CellRect(obj uintptr, ACol int32, ARow int32) TRect {
	var ret TRect
	getLazyProc("StringGrid_CellRect").Call(obj, uintptr(ACol), uintptr(ARow), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_CellToGridZone(obj uintptr, aCol int32, aRow int32) TGridZone {
	ret, _, _ := getLazyProc("StringGrid_CellToGridZone").Call(obj, uintptr(aCol), uintptr(aRow))
	return TGridZone(ret)
}

func StringGrid_CheckPosition(obj uintptr) {
	getLazyProc("StringGrid_CheckPosition").Call(obj)
}

func StringGrid_ClearCols(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_ClearCols").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_ClearRows(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_ClearRows").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_Clear(obj uintptr) {
	getLazyProc("StringGrid_Clear").Call(obj)
}

func StringGrid_ClearSelections(obj uintptr) {
	getLazyProc("StringGrid_ClearSelections").Call(obj)
}

func StringGrid_HasMultiSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_HasMultiSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_InvalidateCell(obj uintptr, aCol int32, aRow int32) {
	getLazyProc("StringGrid_InvalidateCell").Call(obj, uintptr(aCol), uintptr(aRow))
}

func StringGrid_InvalidateCol(obj uintptr, ACol int32) {
	getLazyProc("StringGrid_InvalidateCol").Call(obj, uintptr(ACol))
}

func StringGrid_InvalidateRange(obj uintptr, aRange TRect) {
	getLazyProc("StringGrid_InvalidateRange").Call(obj, uintptr(unsafe.Pointer(&aRange)))
}

func StringGrid_InvalidateRow(obj uintptr, ARow int32) {
	getLazyProc("StringGrid_InvalidateRow").Call(obj, uintptr(ARow))
}

func StringGrid_IsCellVisible(obj uintptr, aCol int32, aRow int32) bool {
	ret, _, _ := getLazyProc("StringGrid_IsCellVisible").Call(obj, uintptr(aCol), uintptr(aRow))
	return DBoolToGoBool(ret)
}

func StringGrid_IsFixedCellVisible(obj uintptr, aCol int32, aRow int32) bool {
	ret, _, _ := getLazyProc("StringGrid_IsFixedCellVisible").Call(obj, uintptr(aCol), uintptr(aRow))
	return DBoolToGoBool(ret)
}

func StringGrid_MouseCoord(obj uintptr, X int32, Y int32) TGridCoord {
	var ret TGridCoord
	getLazyProc("StringGrid_MouseCoord").Call(obj, uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_MouseToCell(obj uintptr, Mouse TPoint) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_MouseToCell").Call(obj, uintptr(unsafe.Pointer(&Mouse)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_MouseToLogcell(obj uintptr, Mouse TPoint) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_MouseToLogcell").Call(obj, uintptr(unsafe.Pointer(&Mouse)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_MouseToGridZone(obj uintptr, X int32, Y int32) TGridZone {
	ret, _, _ := getLazyProc("StringGrid_MouseToGridZone").Call(obj, uintptr(X), uintptr(Y))
	return TGridZone(ret)
}

func StringGrid_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func StringGrid_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("StringGrid_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func StringGrid_DisableAlign(obj uintptr) {
	getLazyProc("StringGrid_DisableAlign").Call(obj)
}

func StringGrid_EnableAlign(obj uintptr) {
	getLazyProc("StringGrid_EnableAlign").Call(obj)
}

func StringGrid_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("StringGrid_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func StringGrid_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("StringGrid_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func StringGrid_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("StringGrid_InsertControl").Call(obj, AControl)
}

func StringGrid_Invalidate(obj uintptr) {
	getLazyProc("StringGrid_Invalidate").Call(obj)
}

func StringGrid_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("StringGrid_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func StringGrid_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("StringGrid_RemoveControl").Call(obj, AControl)
}

func StringGrid_Realign(obj uintptr) {
	getLazyProc("StringGrid_Realign").Call(obj)
}

func StringGrid_Repaint(obj uintptr) {
	getLazyProc("StringGrid_Repaint").Call(obj)
}

func StringGrid_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("StringGrid_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func StringGrid_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("StringGrid_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func StringGrid_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("StringGrid_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func StringGrid_SetFocus(obj uintptr) {
	getLazyProc("StringGrid_SetFocus").Call(obj)
}

func StringGrid_Update(obj uintptr) {
	getLazyProc("StringGrid_Update").Call(obj)
}

func StringGrid_BringToFront(obj uintptr) {
	getLazyProc("StringGrid_BringToFront").Call(obj)
}

func StringGrid_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_Hide(obj uintptr) {
	getLazyProc("StringGrid_Hide").Call(obj)
}

func StringGrid_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("StringGrid_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func StringGrid_Refresh(obj uintptr) {
	getLazyProc("StringGrid_Refresh").Call(obj)
}

func StringGrid_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_SendToBack(obj uintptr) {
	getLazyProc("StringGrid_SendToBack").Call(obj)
}

func StringGrid_Show(obj uintptr) {
	getLazyProc("StringGrid_Show").Call(obj)
}

func StringGrid_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("StringGrid_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func StringGrid_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetTextLen").Call(obj)
	return int32(ret)
}

func StringGrid_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("StringGrid_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func StringGrid_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("StringGrid_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func StringGrid_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StringGrid_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StringGrid_Assign(obj uintptr, Source uintptr) {
	getLazyProc("StringGrid_Assign").Call(obj, Source)
}

func StringGrid_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StringGrid_ClassType").Call(obj)
	return TClass(ret)
}

func StringGrid_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StringGrid_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StringGrid_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_InstanceSize").Call(obj)
	return int32(ret)
}

func StringGrid_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StringGrid_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StringGrid_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func StringGrid_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetHashCode").Call(obj)
	return int32(ret)
}

func StringGrid_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StringGrid_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StringGrid_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("StringGrid_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StringGrid_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("StringGrid_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StringGrid_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("StringGrid_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func StringGrid_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("StringGrid_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func StringGrid_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("StringGrid_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func StringGrid_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("StringGrid_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func StringGrid_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("StringGrid_AnchorClient").Call(obj, uintptr(ASpace))
}

func StringGrid_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StringGrid_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("StringGrid_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func StringGrid_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("StringGrid_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func StringGrid_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("StringGrid_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func StringGrid_GetSelectedColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetSelectedColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetSelectedColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetSelectedColor").Call(obj, uintptr(value))
}

func StringGrid_GetSelectedColumn(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetSelectedColumn").Call(obj)
	return ret
}

func StringGrid_GetStrictSort(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetStrictSort").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetStrictSort(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetStrictSort").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetFixedHotColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetFixedHotColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetFixedHotColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetFixedHotColor").Call(obj, uintptr(value))
}

func StringGrid_GetFastEditing(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetFastEditing").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetFastEditing(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetFastEditing").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetFixedGridLineColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetFixedGridLineColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetFixedGridLineColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetFixedGridLineColor").Call(obj, uintptr(value))
}

func StringGrid_GetFocusColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetFocusColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetFocusColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetFocusColor").Call(obj, uintptr(value))
}

func StringGrid_GetFocusRectVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetFocusRectVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetFocusRectVisible(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetFocusRectVisible").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetGridLineColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetGridLineColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetGridLineColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetGridLineColor").Call(obj, uintptr(value))
}

func StringGrid_GetGridLineStyle(obj uintptr) TPenStyle {
	ret, _, _ := getLazyProc("StringGrid_GetGridLineStyle").Call(obj)
	return TPenStyle(ret)
}

func StringGrid_SetGridLineStyle(obj uintptr, value TPenStyle) {
	getLazyProc("StringGrid_SetGridLineStyle").Call(obj, uintptr(value))
}

func StringGrid_GetEditor(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetEditor").Call(obj)
	return ret
}

func StringGrid_SetEditor(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetEditor").Call(obj, value)
}

func StringGrid_GetEditorBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("StringGrid_GetEditorBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func StringGrid_SetEditorBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("StringGrid_SetEditorBorderStyle").Call(obj, uintptr(value))
}

func StringGrid_GetEditorMode(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetEditorMode").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetEditorMode(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetEditorMode").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetSortOrder(obj uintptr) TSortOrder {
	ret, _, _ := getLazyProc("StringGrid_GetSortOrder").Call(obj)
	return TSortOrder(ret)
}

func StringGrid_SetSortOrder(obj uintptr, value TSortOrder) {
	getLazyProc("StringGrid_SetSortOrder").Call(obj, uintptr(value))
}

func StringGrid_GetSortColumn(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetSortColumn").Call(obj)
	return int32(ret)
}

func StringGrid_SetOnAfterSelection(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnAfterSelection").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnBeforeSelection(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnBeforeSelection").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnButtonClick(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnButtonClick").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnCheckboxToggled(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnCheckboxToggled").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnColRowDeleted(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnColRowDeleted").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnColRowExchanged(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnColRowExchanged").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnColRowInserted(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnColRowInserted").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnColRowMoved(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnColRowMoved").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnCompareCells(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnCompareCells").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnEditingDone(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnEditingDone").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnGetCellHint(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnGetCellHint").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnGetCheckboxState(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnGetCheckboxState").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnSetCheckboxState(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnSetCheckboxState").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnHeaderClick(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnHeaderClick").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnHeaderSized(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnHeaderSized").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnHeaderSizing(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnHeaderSizing").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnPickListSelect(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnPickListSelect").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnSelection(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnSelection").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnSelectEditor(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnSelectEditor").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnUserCheckboxBitmap(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnUserCheckboxBitmap").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnValidateEntry(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnValidateEntry").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnPrepareCanvas(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnPrepareCanvas").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_GetAlternateColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetAlternateColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetAlternateColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetAlternateColor").Call(obj, uintptr(value))
}

func StringGrid_GetAutoAdvance(obj uintptr) TAutoAdvance {
	ret, _, _ := getLazyProc("StringGrid_GetAutoAdvance").Call(obj)
	return TAutoAdvance(ret)
}

func StringGrid_SetAutoAdvance(obj uintptr, value TAutoAdvance) {
	getLazyProc("StringGrid_SetAutoAdvance").Call(obj, uintptr(value))
}

func StringGrid_GetAutoEdit(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetAutoEdit").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetAutoEdit(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetAutoEdit").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetAutoFillColumns(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetAutoFillColumns").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetAutoFillColumns(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetAutoFillColumns").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetCellHintPriority(obj uintptr) TCellHintPriority {
	ret, _, _ := getLazyProc("StringGrid_GetCellHintPriority").Call(obj)
	return TCellHintPriority(ret)
}

func StringGrid_SetCellHintPriority(obj uintptr, value TCellHintPriority) {
	getLazyProc("StringGrid_SetCellHintPriority").Call(obj, uintptr(value))
}

func StringGrid_GetColumnClickSorts(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetColumnClickSorts").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetColumnClickSorts(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetColumnClickSorts").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetColumns(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetColumns").Call(obj)
	return ret
}

func StringGrid_SetColumns(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetColumns").Call(obj, value)
}

func StringGrid_GetExtendedSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetExtendedSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetExtendedSelect(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetExtendedSelect").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetFlat(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetFlat").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetFlat(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetFlat").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetHeaderHotZones(obj uintptr) TGridZoneSet {
	ret, _, _ := getLazyProc("StringGrid_GetHeaderHotZones").Call(obj)
	return TGridZoneSet(ret)
}

func StringGrid_SetHeaderHotZones(obj uintptr, value TGridZoneSet) {
	getLazyProc("StringGrid_SetHeaderHotZones").Call(obj, uintptr(value))
}

func StringGrid_GetHeaderPushZones(obj uintptr) TGridZoneSet {
	ret, _, _ := getLazyProc("StringGrid_GetHeaderPushZones").Call(obj)
	return TGridZoneSet(ret)
}

func StringGrid_SetHeaderPushZones(obj uintptr, value TGridZoneSet) {
	getLazyProc("StringGrid_SetHeaderPushZones").Call(obj, uintptr(value))
}

func StringGrid_GetImageIndexSortAsc(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetImageIndexSortAsc").Call(obj)
	return int32(ret)
}

func StringGrid_SetImageIndexSortAsc(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetImageIndexSortAsc").Call(obj, uintptr(value))
}

func StringGrid_GetImageIndexSortDesc(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetImageIndexSortDesc").Call(obj)
	return int32(ret)
}

func StringGrid_SetImageIndexSortDesc(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetImageIndexSortDesc").Call(obj, uintptr(value))
}

func StringGrid_GetMouseWheelOption(obj uintptr) TMouseWheelOption {
	ret, _, _ := getLazyProc("StringGrid_GetMouseWheelOption").Call(obj)
	return TMouseWheelOption(ret)
}

func StringGrid_SetMouseWheelOption(obj uintptr, value TMouseWheelOption) {
	getLazyProc("StringGrid_SetMouseWheelOption").Call(obj, uintptr(value))
}

func StringGrid_GetOptions2(obj uintptr) TGridOptions2 {
	ret, _, _ := getLazyProc("StringGrid_GetOptions2").Call(obj)
	return TGridOptions2(ret)
}

func StringGrid_SetOptions2(obj uintptr, value TGridOptions2) {
	getLazyProc("StringGrid_SetOptions2").Call(obj, uintptr(value))
}

func StringGrid_GetRangeSelectMode(obj uintptr) TRangeSelectMode {
	ret, _, _ := getLazyProc("StringGrid_GetRangeSelectMode").Call(obj)
	return TRangeSelectMode(ret)
}

func StringGrid_SetRangeSelectMode(obj uintptr, value TRangeSelectMode) {
	getLazyProc("StringGrid_SetRangeSelectMode").Call(obj, uintptr(value))
}

func StringGrid_GetTabAdvance(obj uintptr) TAutoAdvance {
	ret, _, _ := getLazyProc("StringGrid_GetTabAdvance").Call(obj)
	return TAutoAdvance(ret)
}

func StringGrid_SetTabAdvance(obj uintptr, value TAutoAdvance) {
	getLazyProc("StringGrid_SetTabAdvance").Call(obj, uintptr(value))
}

func StringGrid_GetTitleFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetTitleFont").Call(obj)
	return ret
}

func StringGrid_SetTitleFont(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetTitleFont").Call(obj, value)
}

func StringGrid_GetTitleImageList(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetTitleImageList").Call(obj)
	return ret
}

func StringGrid_SetTitleImageList(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetTitleImageList").Call(obj, value)
}

func StringGrid_GetTitleStyle(obj uintptr) TTitleStyle {
	ret, _, _ := getLazyProc("StringGrid_GetTitleStyle").Call(obj)
	return TTitleStyle(ret)
}

func StringGrid_SetTitleStyle(obj uintptr, value TTitleStyle) {
	getLazyProc("StringGrid_SetTitleStyle").Call(obj, uintptr(value))
}

func StringGrid_GetUseXORFeatures(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetUseXORFeatures").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetUseXORFeatures(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetUseXORFeatures").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("StringGrid_GetAlign").Call(obj)
	return TAlign(ret)
}

func StringGrid_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("StringGrid_SetAlign").Call(obj, uintptr(value))
}

func StringGrid_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("StringGrid_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func StringGrid_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("StringGrid_SetAnchors").Call(obj, uintptr(value))
}

func StringGrid_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("StringGrid_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func StringGrid_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("StringGrid_SetBiDiMode").Call(obj, uintptr(value))
}

func StringGrid_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("StringGrid_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func StringGrid_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("StringGrid_SetBorderStyle").Call(obj, uintptr(value))
}

func StringGrid_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetColor").Call(obj, uintptr(value))
}

func StringGrid_GetColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetColCount").Call(obj)
	return int32(ret)
}

func StringGrid_SetColCount(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetColCount").Call(obj, uintptr(value))
}

func StringGrid_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetConstraints").Call(obj)
	return ret
}

func StringGrid_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetConstraints").Call(obj, value)
}

func StringGrid_GetDefaultColWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetDefaultColWidth").Call(obj)
	return int32(ret)
}

func StringGrid_SetDefaultColWidth(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetDefaultColWidth").Call(obj, uintptr(value))
}

func StringGrid_GetDefaultRowHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetDefaultRowHeight").Call(obj)
	return int32(ret)
}

func StringGrid_SetDefaultRowHeight(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetDefaultRowHeight").Call(obj, uintptr(value))
}

func StringGrid_GetDefaultDrawing(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetDefaultDrawing").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetDefaultDrawing(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetDefaultDrawing").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StringGrid_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func StringGrid_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("StringGrid_SetDragCursor").Call(obj, uintptr(value))
}

func StringGrid_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("StringGrid_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func StringGrid_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("StringGrid_SetDragKind").Call(obj, uintptr(value))
}

func StringGrid_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("StringGrid_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func StringGrid_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("StringGrid_SetDragMode").Call(obj, uintptr(value))
}

func StringGrid_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetEnabled(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetFixedColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StringGrid_GetFixedColor").Call(obj)
	return TColor(ret)
}

func StringGrid_SetFixedColor(obj uintptr, value TColor) {
	getLazyProc("StringGrid_SetFixedColor").Call(obj, uintptr(value))
}

func StringGrid_GetFixedCols(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetFixedCols").Call(obj)
	return int32(ret)
}

func StringGrid_SetFixedCols(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetFixedCols").Call(obj, uintptr(value))
}

func StringGrid_GetRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetRowCount").Call(obj)
	return int32(ret)
}

func StringGrid_SetRowCount(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetRowCount").Call(obj, uintptr(value))
}

func StringGrid_GetFixedRows(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetFixedRows").Call(obj)
	return int32(ret)
}

func StringGrid_SetFixedRows(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetFixedRows").Call(obj, uintptr(value))
}

func StringGrid_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetFont").Call(obj)
	return ret
}

func StringGrid_SetFont(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetFont").Call(obj, value)
}

func StringGrid_GetGridLineWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetGridLineWidth").Call(obj)
	return int32(ret)
}

func StringGrid_SetGridLineWidth(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetGridLineWidth").Call(obj, uintptr(value))
}

func StringGrid_GetOptions(obj uintptr) TGridOptions {
	ret, _, _ := getLazyProc("StringGrid_GetOptions").Call(obj)
	return TGridOptions(ret)
}

func StringGrid_SetOptions(obj uintptr, value TGridOptions) {
	getLazyProc("StringGrid_SetOptions").Call(obj, uintptr(value))
}

func StringGrid_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetParentColor(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetParentFont(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetPopupMenu").Call(obj)
	return ret
}

func StringGrid_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetPopupMenu").Call(obj, value)
}

func StringGrid_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("StringGrid_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func StringGrid_SetScrollBars(obj uintptr, value TScrollStyle) {
	getLazyProc("StringGrid_SetScrollBars").Call(obj, uintptr(value))
}

func StringGrid_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetShowHint(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("StringGrid_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func StringGrid_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("StringGrid_SetTabOrder").Call(obj, uintptr(value))
}

func StringGrid_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetVisible(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetVisible").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetVisibleColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetVisibleColCount").Call(obj)
	return int32(ret)
}

func StringGrid_GetVisibleRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetVisibleRowCount").Call(obj)
	return int32(ret)
}

func StringGrid_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnDrawCell(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnDrawCell").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnGetEditMask(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnGetEditMask").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnGetEditText(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnGetEditText").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnSelectCell(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnSelectCell").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnSetEditText(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnSetEditText").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_SetOnTopLeftChanged(obj uintptr, fn interface{}) {
	getLazyProc("StringGrid_SetOnTopLeftChanged").Call(obj, addEventToMap(obj, fn))
}

func StringGrid_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetCanvas").Call(obj)
	return ret
}

func StringGrid_GetCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetCol").Call(obj)
	return int32(ret)
}

func StringGrid_SetCol(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetCol").Call(obj, uintptr(value))
}

func StringGrid_GetGridHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetGridHeight").Call(obj)
	return int32(ret)
}

func StringGrid_GetGridWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetGridWidth").Call(obj)
	return int32(ret)
}

func StringGrid_GetLeftCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetLeftCol").Call(obj)
	return int32(ret)
}

func StringGrid_SetLeftCol(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetLeftCol").Call(obj, uintptr(value))
}

func StringGrid_GetSelection(obj uintptr) TGridRect {
	var ret TGridRect
	getLazyProc("StringGrid_GetSelection").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_SetSelection(obj uintptr, value TGridRect) {
	getLazyProc("StringGrid_SetSelection").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StringGrid_GetRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetRow").Call(obj)
	return int32(ret)
}

func StringGrid_SetRow(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetRow").Call(obj, uintptr(value))
}

func StringGrid_GetTopRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetTopRow").Call(obj)
	return int32(ret)
}

func StringGrid_SetTopRow(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetTopRow").Call(obj, uintptr(value))
}

func StringGrid_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetTabStop(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetDockClientCount").Call(obj)
	return int32(ret)
}

func StringGrid_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetDockSite(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func StringGrid_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetBrush").Call(obj)
	return ret
}

func StringGrid_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetControlCount").Call(obj)
	return int32(ret)
}

func StringGrid_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StringGrid_GetHandle").Call(obj)
	return HWND(ret)
}

func StringGrid_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StringGrid_GetParentWindow").Call(obj)
	return HWND(ret)
}

func StringGrid_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("StringGrid_SetParentWindow").Call(obj, uintptr(value))
}

func StringGrid_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("StringGrid_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func StringGrid_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAction").Call(obj)
	return ret
}

func StringGrid_SetAction(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetAction").Call(obj, value)
}

func StringGrid_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("StringGrid_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("StringGrid_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StringGrid_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetClientHeight").Call(obj)
	return int32(ret)
}

func StringGrid_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetClientHeight").Call(obj, uintptr(value))
}

func StringGrid_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("StringGrid_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("StringGrid_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StringGrid_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetClientWidth").Call(obj)
	return int32(ret)
}

func StringGrid_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetClientWidth").Call(obj, uintptr(value))
}

func StringGrid_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("StringGrid_GetControlState").Call(obj)
	return TControlState(ret)
}

func StringGrid_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("StringGrid_SetControlState").Call(obj, uintptr(value))
}

func StringGrid_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("StringGrid_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func StringGrid_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("StringGrid_SetControlStyle").Call(obj, uintptr(value))
}

func StringGrid_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringGrid_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func StringGrid_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetParent").Call(obj)
	return ret
}

func StringGrid_SetParent(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetParent").Call(obj, value)
}

func StringGrid_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetLeft").Call(obj)
	return int32(ret)
}

func StringGrid_SetLeft(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetLeft").Call(obj, uintptr(value))
}

func StringGrid_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetTop").Call(obj)
	return int32(ret)
}

func StringGrid_SetTop(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetTop").Call(obj, uintptr(value))
}

func StringGrid_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetWidth").Call(obj)
	return int32(ret)
}

func StringGrid_SetWidth(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetWidth").Call(obj, uintptr(value))
}

func StringGrid_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetHeight").Call(obj)
	return int32(ret)
}

func StringGrid_SetHeight(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetHeight").Call(obj, uintptr(value))
}

func StringGrid_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StringGrid_GetCursor").Call(obj)
	return TCursor(ret)
}

func StringGrid_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("StringGrid_SetCursor").Call(obj, uintptr(value))
}

func StringGrid_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("StringGrid_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func StringGrid_SetHint(obj uintptr, value string) {
	getLazyProc("StringGrid_SetHint").Call(obj, GoStrToDStr(value))
}

func StringGrid_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetComponentCount").Call(obj)
	return int32(ret)
}

func StringGrid_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetComponentIndex").Call(obj)
	return int32(ret)
}

func StringGrid_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("StringGrid_SetComponentIndex").Call(obj, uintptr(value))
}

func StringGrid_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetOwner").Call(obj)
	return ret
}

func StringGrid_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("StringGrid_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func StringGrid_SetName(obj uintptr, value string) {
	getLazyProc("StringGrid_SetName").Call(obj, GoStrToDStr(value))
}

func StringGrid_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("StringGrid_GetTag").Call(obj)
	return int(ret)
}

func StringGrid_SetTag(obj uintptr, value int) {
	getLazyProc("StringGrid_SetTag").Call(obj, uintptr(value))
}

func StringGrid_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAnchorSideLeft").Call(obj)
	return ret
}

func StringGrid_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetAnchorSideLeft").Call(obj, value)
}

func StringGrid_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAnchorSideTop").Call(obj)
	return ret
}

func StringGrid_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetAnchorSideTop").Call(obj, value)
}

func StringGrid_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAnchorSideRight").Call(obj)
	return ret
}

func StringGrid_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetAnchorSideRight").Call(obj, value)
}

func StringGrid_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAnchorSideBottom").Call(obj)
	return ret
}

func StringGrid_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetAnchorSideBottom").Call(obj, value)
}

func StringGrid_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetChildSizing").Call(obj)
	return ret
}

func StringGrid_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetChildSizing").Call(obj, value)
}

func StringGrid_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetBorderSpacing").Call(obj)
	return ret
}

func StringGrid_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("StringGrid_SetBorderSpacing").Call(obj, value)
}

func StringGrid_GetIsCellSelected(obj uintptr, aCol int32, aRow int32) bool {
	ret, _, _ := getLazyProc("StringGrid_GetIsCellSelected").Call(obj, uintptr(aCol), uintptr(aRow))
	return DBoolToGoBool(ret)
}

func StringGrid_GetCells(obj uintptr, ACol int32, ARow int32) string {
	ret, _, _ := getLazyProc("StringGrid_GetCells").Call(obj, uintptr(ACol), uintptr(ARow))
	return DStrToGoStr(ret)
}

func StringGrid_SetCells(obj uintptr, ACol int32, ARow int32, value string) {
	getLazyProc("StringGrid_SetCells").Call(obj, uintptr(ACol), uintptr(ARow), GoStrToDStr(value))
}

func StringGrid_GetCols(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetCols").Call(obj, uintptr(Index))
	return ret
}

func StringGrid_SetCols(obj uintptr, Index int32, value uintptr) {
	getLazyProc("StringGrid_SetCols").Call(obj, uintptr(Index), value)
}

func StringGrid_GetObjects(obj uintptr, ACol int32, ARow int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetObjects").Call(obj, uintptr(ACol), uintptr(ARow))
	return ret
}

func StringGrid_SetObjects(obj uintptr, ACol int32, ARow int32, value uintptr) {
	getLazyProc("StringGrid_SetObjects").Call(obj, uintptr(ACol), uintptr(ARow), value)
}

func StringGrid_GetRows(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetRows").Call(obj, uintptr(Index))
	return ret
}

func StringGrid_SetRows(obj uintptr, Index int32, value uintptr) {
	getLazyProc("StringGrid_SetRows").Call(obj, uintptr(Index), value)
}

func StringGrid_GetColWidths(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetColWidths").Call(obj, uintptr(Index))
	return int32(ret)
}

func StringGrid_SetColWidths(obj uintptr, Index int32, value int32) {
	getLazyProc("StringGrid_SetColWidths").Call(obj, uintptr(Index), uintptr(value))
}

func StringGrid_GetRowHeights(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("StringGrid_GetRowHeights").Call(obj, uintptr(Index))
	return int32(ret)
}

func StringGrid_SetRowHeights(obj uintptr, Index int32, value int32) {
	getLazyProc("StringGrid_SetRowHeights").Call(obj, uintptr(Index), uintptr(value))
}

func StringGrid_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func StringGrid_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetControls").Call(obj, uintptr(Index))
	return ret
}

func StringGrid_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func StringGrid_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("StringGrid_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func StringGrid_StaticClassType() TClass {
	r, _, _ := getLazyProc("StringGrid_StaticClassType").Call()
	return TClass(r)
}
