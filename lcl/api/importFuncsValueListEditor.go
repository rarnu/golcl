package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TValueListEditor ---------------------------

func ValueListEditor_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_Create").Call(obj)
	return ret
}

func ValueListEditor_Free(obj uintptr) {
	getLazyProc("ValueListEditor_Free").Call(obj)
}

func ValueListEditor_DeleteRow(obj uintptr, ARow int32) {
	getLazyProc("ValueListEditor_DeleteRow").Call(obj, uintptr(ARow))
}

func ValueListEditor_Refresh(obj uintptr) {
	getLazyProc("ValueListEditor_Refresh").Call(obj)
}

func ValueListEditor_CellRect(obj uintptr, ACol int32, ARow int32) TRect {
	var ret TRect
	getLazyProc("ValueListEditor_CellRect").Call(obj, uintptr(ACol), uintptr(ARow), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_MouseToCell(obj uintptr, X int32, Y int32, ACol *int32, ARow *int32) {
	getLazyProc("ValueListEditor_MouseToCell").Call(obj, uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(ACol)), uintptr(unsafe.Pointer(ARow)))
}

func ValueListEditor_MouseCoord(obj uintptr, X int32, Y int32) TGridCoord {
	var ret TGridCoord
	getLazyProc("ValueListEditor_MouseCoord").Call(obj, uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ValueListEditor_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ValueListEditor_DisableAlign(obj uintptr) {
	getLazyProc("ValueListEditor_DisableAlign").Call(obj)
}

func ValueListEditor_EnableAlign(obj uintptr) {
	getLazyProc("ValueListEditor_EnableAlign").Call(obj)
}

func ValueListEditor_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ValueListEditor_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ValueListEditor_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ValueListEditor_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ValueListEditor_InsertControl").Call(obj, AControl)
}

func ValueListEditor_Invalidate(obj uintptr) {
	getLazyProc("ValueListEditor_Invalidate").Call(obj)
}

func ValueListEditor_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ValueListEditor_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ValueListEditor_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ValueListEditor_RemoveControl").Call(obj, AControl)
}

func ValueListEditor_Realign(obj uintptr) {
	getLazyProc("ValueListEditor_Realign").Call(obj)
}

func ValueListEditor_Repaint(obj uintptr) {
	getLazyProc("ValueListEditor_Repaint").Call(obj)
}

func ValueListEditor_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ValueListEditor_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ValueListEditor_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ValueListEditor_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ValueListEditor_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ValueListEditor_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ValueListEditor_SetFocus(obj uintptr) {
	getLazyProc("ValueListEditor_SetFocus").Call(obj)
}

func ValueListEditor_Update(obj uintptr) {
	getLazyProc("ValueListEditor_Update").Call(obj)
}

func ValueListEditor_BringToFront(obj uintptr) {
	getLazyProc("ValueListEditor_BringToFront").Call(obj)
}

func ValueListEditor_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ValueListEditor_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ValueListEditor_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_Hide(obj uintptr) {
	getLazyProc("ValueListEditor_Hide").Call(obj)
}

func ValueListEditor_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ValueListEditor_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ValueListEditor_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ValueListEditor_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ValueListEditor_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_SendToBack(obj uintptr) {
	getLazyProc("ValueListEditor_SendToBack").Call(obj)
}

func ValueListEditor_Show(obj uintptr) {
	getLazyProc("ValueListEditor_Show").Call(obj)
}

func ValueListEditor_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ValueListEditor_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ValueListEditor_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetTextLen").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ValueListEditor_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ValueListEditor_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ValueListEditor_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ValueListEditor_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ValueListEditor_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ValueListEditor_Assign").Call(obj, Source)
}

func ValueListEditor_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ValueListEditor_ClassType").Call(obj)
	return TClass(ret)
}

func ValueListEditor_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ValueListEditor_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ValueListEditor_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_InstanceSize").Call(obj)
	return int32(ret)
}

func ValueListEditor_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ValueListEditor_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ValueListEditor_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetHashCode").Call(obj)
	return int32(ret)
}

func ValueListEditor_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ValueListEditor_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ValueListEditor_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ValueListEditor_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ValueListEditor_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ValueListEditor_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ValueListEditor_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ValueListEditor_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ValueListEditor_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ValueListEditor_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ValueListEditor_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ValueListEditor_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ValueListEditor_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ValueListEditor_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ValueListEditor_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ValueListEditor_AnchorClient").Call(obj, uintptr(ASpace))
}

func ValueListEditor_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ValueListEditor_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ValueListEditor_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ValueListEditor_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ValueListEditor_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ValueListEditor_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ValueListEditor_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ValueListEditor_GetColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetColCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetColCount(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetColCount").Call(obj, uintptr(value))
}

func ValueListEditor_GetRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetRowCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetVisibleColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetVisibleColCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetVisibleRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetVisibleRowCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ValueListEditor_GetAlign").Call(obj)
	return TAlign(ret)
}

func ValueListEditor_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ValueListEditor_SetAlign").Call(obj, uintptr(value))
}

func ValueListEditor_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ValueListEditor_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ValueListEditor_SetAnchors").Call(obj, uintptr(value))
}

func ValueListEditor_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ValueListEditor_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ValueListEditor_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ValueListEditor_SetBiDiMode").Call(obj, uintptr(value))
}

func ValueListEditor_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("ValueListEditor_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func ValueListEditor_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("ValueListEditor_SetBorderStyle").Call(obj, uintptr(value))
}

func ValueListEditor_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ValueListEditor_GetColor").Call(obj)
	return TColor(ret)
}

func ValueListEditor_SetColor(obj uintptr, value TColor) {
	getLazyProc("ValueListEditor_SetColor").Call(obj, uintptr(value))
}

func ValueListEditor_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetConstraints").Call(obj)
	return ret
}

func ValueListEditor_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetConstraints").Call(obj, value)
}

func ValueListEditor_GetDefaultColWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetDefaultColWidth").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetDefaultColWidth(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetDefaultColWidth").Call(obj, uintptr(value))
}

func ValueListEditor_GetDefaultDrawing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetDefaultDrawing").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetDefaultDrawing(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetDefaultDrawing").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetDefaultRowHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetDefaultRowHeight").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetDefaultRowHeight(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetDefaultRowHeight").Call(obj, uintptr(value))
}

func ValueListEditor_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ValueListEditor_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ValueListEditor_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ValueListEditor_SetDragCursor").Call(obj, uintptr(value))
}

func ValueListEditor_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ValueListEditor_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ValueListEditor_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ValueListEditor_SetDragKind").Call(obj, uintptr(value))
}

func ValueListEditor_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ValueListEditor_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ValueListEditor_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ValueListEditor_SetDragMode").Call(obj, uintptr(value))
}

func ValueListEditor_GetDropDownRows(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetDropDownRows").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetDropDownRows(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetDropDownRows").Call(obj, uintptr(value))
}

func ValueListEditor_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetFixedColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ValueListEditor_GetFixedColor").Call(obj)
	return TColor(ret)
}

func ValueListEditor_SetFixedColor(obj uintptr, value TColor) {
	getLazyProc("ValueListEditor_SetFixedColor").Call(obj, uintptr(value))
}

func ValueListEditor_GetFixedCols(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetFixedCols").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetFixedCols(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetFixedCols").Call(obj, uintptr(value))
}

func ValueListEditor_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetFont").Call(obj)
	return ret
}

func ValueListEditor_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetFont").Call(obj, value)
}

func ValueListEditor_GetGridLineWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetGridLineWidth").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetGridLineWidth(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetGridLineWidth").Call(obj, uintptr(value))
}

func ValueListEditor_GetOptions(obj uintptr) TGridOptions {
	ret, _, _ := getLazyProc("ValueListEditor_GetOptions").Call(obj)
	return TGridOptions(ret)
}

func ValueListEditor_SetOptions(obj uintptr, value TGridOptions) {
	getLazyProc("ValueListEditor_SetOptions").Call(obj, uintptr(value))
}

func ValueListEditor_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetPopupMenu").Call(obj)
	return ret
}

func ValueListEditor_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetPopupMenu").Call(obj, value)
}

func ValueListEditor_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("ValueListEditor_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func ValueListEditor_SetScrollBars(obj uintptr, value TScrollStyle) {
	getLazyProc("ValueListEditor_SetScrollBars").Call(obj, uintptr(value))
}

func ValueListEditor_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetStrings(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetStrings").Call(obj)
	return ret
}

func ValueListEditor_SetStrings(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetStrings").Call(obj, value)
}

func ValueListEditor_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ValueListEditor_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ValueListEditor_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ValueListEditor_SetTabOrder").Call(obj, uintptr(value))
}

func ValueListEditor_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetVisible(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnDrawCell(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnDrawCell").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnGetEditMask(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnGetEditMask").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnGetEditText(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnGetEditText").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnSelectCell(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnSelectCell").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnSetEditText(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnSetEditText").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_SetOnTopLeftChanged(obj uintptr, fn interface{}) {
	getLazyProc("ValueListEditor_SetOnTopLeftChanged").Call(obj, addEventToMap(obj, fn))
}

func ValueListEditor_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetCanvas").Call(obj)
	return ret
}

func ValueListEditor_GetCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetCol").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetCol(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetCol").Call(obj, uintptr(value))
}

func ValueListEditor_GetEditorMode(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetEditorMode").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetEditorMode(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetEditorMode").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetGridHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetGridHeight").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetGridWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetGridWidth").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetLeftCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetLeftCol").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetLeftCol(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetLeftCol").Call(obj, uintptr(value))
}

func ValueListEditor_GetSelection(obj uintptr) TGridRect {
	var ret TGridRect
	getLazyProc("ValueListEditor_GetSelection").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_SetSelection(obj uintptr, value TGridRect) {
	getLazyProc("ValueListEditor_SetSelection").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ValueListEditor_GetRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetRow").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetRow(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetRow").Call(obj, uintptr(value))
}

func ValueListEditor_GetTopRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetTopRow").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetTopRow(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetTopRow").Call(obj, uintptr(value))
}

func ValueListEditor_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetBrush").Call(obj)
	return ret
}

func ValueListEditor_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetControlCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ValueListEditor_GetHandle").Call(obj)
	return HWND(ret)
}

func ValueListEditor_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ValueListEditor_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ValueListEditor_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ValueListEditor_SetParentWindow").Call(obj, uintptr(value))
}

func ValueListEditor_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ValueListEditor_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ValueListEditor_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAction").Call(obj)
	return ret
}

func ValueListEditor_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetAction").Call(obj, value)
}

func ValueListEditor_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ValueListEditor_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ValueListEditor_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ValueListEditor_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetClientHeight").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetClientHeight").Call(obj, uintptr(value))
}

func ValueListEditor_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ValueListEditor_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ValueListEditor_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ValueListEditor_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetClientWidth").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetClientWidth").Call(obj, uintptr(value))
}

func ValueListEditor_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ValueListEditor_GetControlState").Call(obj)
	return TControlState(ret)
}

func ValueListEditor_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ValueListEditor_SetControlState").Call(obj, uintptr(value))
}

func ValueListEditor_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ValueListEditor_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ValueListEditor_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ValueListEditor_SetControlStyle").Call(obj, uintptr(value))
}

func ValueListEditor_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ValueListEditor_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ValueListEditor_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetParent").Call(obj)
	return ret
}

func ValueListEditor_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetParent").Call(obj, value)
}

func ValueListEditor_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetLeft").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetLeft(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetLeft").Call(obj, uintptr(value))
}

func ValueListEditor_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetTop").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetTop(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetTop").Call(obj, uintptr(value))
}

func ValueListEditor_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetWidth").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetWidth(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetWidth").Call(obj, uintptr(value))
}

func ValueListEditor_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetHeight").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetHeight(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetHeight").Call(obj, uintptr(value))
}

func ValueListEditor_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ValueListEditor_GetCursor").Call(obj)
	return TCursor(ret)
}

func ValueListEditor_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ValueListEditor_SetCursor").Call(obj, uintptr(value))
}

func ValueListEditor_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ValueListEditor_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ValueListEditor_SetHint(obj uintptr, value string) {
	getLazyProc("ValueListEditor_SetHint").Call(obj, GoStrToDStr(value))
}

func ValueListEditor_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetComponentCount").Call(obj)
	return int32(ret)
}

func ValueListEditor_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ValueListEditor_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ValueListEditor_SetComponentIndex").Call(obj, uintptr(value))
}

func ValueListEditor_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetOwner").Call(obj)
	return ret
}

func ValueListEditor_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ValueListEditor_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ValueListEditor_SetName(obj uintptr, value string) {
	getLazyProc("ValueListEditor_SetName").Call(obj, GoStrToDStr(value))
}

func ValueListEditor_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ValueListEditor_GetTag").Call(obj)
	return int(ret)
}

func ValueListEditor_SetTag(obj uintptr, value int) {
	getLazyProc("ValueListEditor_SetTag").Call(obj, uintptr(value))
}

func ValueListEditor_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchorSideLeft").Call(obj)
	return ret
}

func ValueListEditor_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetAnchorSideLeft").Call(obj, value)
}

func ValueListEditor_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchorSideTop").Call(obj)
	return ret
}

func ValueListEditor_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetAnchorSideTop").Call(obj, value)
}

func ValueListEditor_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchorSideRight").Call(obj)
	return ret
}

func ValueListEditor_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetAnchorSideRight").Call(obj, value)
}

func ValueListEditor_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchorSideBottom").Call(obj)
	return ret
}

func ValueListEditor_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetAnchorSideBottom").Call(obj, value)
}

func ValueListEditor_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetChildSizing").Call(obj)
	return ret
}

func ValueListEditor_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetChildSizing").Call(obj, value)
}

func ValueListEditor_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetBorderSpacing").Call(obj)
	return ret
}

func ValueListEditor_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ValueListEditor_SetBorderSpacing").Call(obj, value)
}

func ValueListEditor_GetCells(obj uintptr, ACol int32, ARow int32) string {
	ret, _, _ := getLazyProc("ValueListEditor_GetCells").Call(obj, uintptr(ACol), uintptr(ARow))
	return DStrToGoStr(ret)
}

func ValueListEditor_SetCells(obj uintptr, ACol int32, ARow int32, value string) {
	getLazyProc("ValueListEditor_SetCells").Call(obj, uintptr(ACol), uintptr(ARow), GoStrToDStr(value))
}

func ValueListEditor_GetValues(obj uintptr, Key string) string {
	ret, _, _ := getLazyProc("ValueListEditor_GetValues").Call(obj, GoStrToDStr(Key))
	return DStrToGoStr(ret)
}

func ValueListEditor_SetValues(obj uintptr, Key string, value string) {
	getLazyProc("ValueListEditor_SetValues").Call(obj, GoStrToDStr(Key), GoStrToDStr(value))
}

func ValueListEditor_GetColWidths(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetColWidths").Call(obj, uintptr(Index))
	return int32(ret)
}

func ValueListEditor_SetColWidths(obj uintptr, Index int32, value int32) {
	getLazyProc("ValueListEditor_SetColWidths").Call(obj, uintptr(Index), uintptr(value))
}

func ValueListEditor_GetRowHeights(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("ValueListEditor_GetRowHeights").Call(obj, uintptr(Index))
	return int32(ret)
}

func ValueListEditor_SetRowHeights(obj uintptr, Index int32, value int32) {
	getLazyProc("ValueListEditor_SetRowHeights").Call(obj, uintptr(Index), uintptr(value))
}

func ValueListEditor_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ValueListEditor_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ValueListEditor_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ValueListEditor_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ValueListEditor_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ValueListEditor_StaticClassType() TClass {
	r, _, _ := getLazyProc("ValueListEditor_StaticClassType").Call()
	return TClass(r)
}
