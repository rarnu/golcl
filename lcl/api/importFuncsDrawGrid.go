package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TDrawGrid ---------------------------

func DrawGrid_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_Create").Call(obj)
	return ret
}

func DrawGrid_Free(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Free").Call(obj)
}

func DrawGrid_CellRect(obj uintptr, ACol int32, ARow int32) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DrawGrid_CellRect").Call(obj, uintptr(ACol), uintptr(ARow), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_MouseToCell(obj uintptr, X int32, Y int32, ACol *int32, ARow *int32) {
	_, _, _ = getLazyProc("DrawGrid_MouseToCell").Call(obj, uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(ACol)), uintptr(unsafe.Pointer(ARow)))
}

func DrawGrid_MouseCoord(obj uintptr, X int32, Y int32) TGridCoord {
	var ret TGridCoord
	_, _, _ = getLazyProc("DrawGrid_MouseCoord").Call(obj, uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func DrawGrid_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func DrawGrid_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_DisableAlign").Call(obj)
}

func DrawGrid_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_EnableAlign").Call(obj)
}

func DrawGrid_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func DrawGrid_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("DrawGrid_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func DrawGrid_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DrawGrid_InsertControl").Call(obj, AControl)
}

func DrawGrid_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Invalidate").Call(obj)
}

func DrawGrid_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("DrawGrid_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func DrawGrid_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DrawGrid_RemoveControl").Call(obj, AControl)
}

func DrawGrid_Realign(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Realign").Call(obj)
}

func DrawGrid_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Repaint").Call(obj)
}

func DrawGrid_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("DrawGrid_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func DrawGrid_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("DrawGrid_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func DrawGrid_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("DrawGrid_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func DrawGrid_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetFocus").Call(obj)
}

func DrawGrid_Update(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Update").Call(obj)
}

func DrawGrid_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_BringToFront").Call(obj)
}

func DrawGrid_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DrawGrid_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DrawGrid_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_Hide(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Hide").Call(obj)
}

func DrawGrid_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("DrawGrid_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func DrawGrid_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Refresh").Call(obj)
}

func DrawGrid_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DrawGrid_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DrawGrid_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SendToBack").Call(obj)
}

func DrawGrid_Show(obj uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Show").Call(obj)
}

func DrawGrid_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("DrawGrid_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func DrawGrid_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetTextLen").Call(obj)
	return int32(ret)
}

func DrawGrid_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("DrawGrid_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func DrawGrid_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func DrawGrid_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("DrawGrid_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func DrawGrid_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("DrawGrid_Assign").Call(obj, Source)
}

func DrawGrid_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("DrawGrid_ClassType").Call(obj)
	return TClass(ret)
}

func DrawGrid_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("DrawGrid_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func DrawGrid_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_InstanceSize").Call(obj)
	return int32(ret)
}

func DrawGrid_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("DrawGrid_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func DrawGrid_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetHashCode").Call(obj)
	return int32(ret)
}

func DrawGrid_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("DrawGrid_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func DrawGrid_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DrawGrid_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DrawGrid_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DrawGrid_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DrawGrid_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DrawGrid_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func DrawGrid_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DrawGrid_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func DrawGrid_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("DrawGrid_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func DrawGrid_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("DrawGrid_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func DrawGrid_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("DrawGrid_AnchorClient").Call(obj, uintptr(ASpace))
}

func DrawGrid_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DrawGrid_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("DrawGrid_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func DrawGrid_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("DrawGrid_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func DrawGrid_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("DrawGrid_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func DrawGrid_SetOnColRowMoved(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnColRowMoved").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnPrepareCanvas(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnPrepareCanvas").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("DrawGrid_GetAlign").Call(obj)
	return TAlign(ret)
}

func DrawGrid_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("DrawGrid_SetAlign").Call(obj, uintptr(value))
}

func DrawGrid_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func DrawGrid_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("DrawGrid_SetAnchors").Call(obj, uintptr(value))
}

func DrawGrid_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("DrawGrid_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func DrawGrid_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("DrawGrid_SetBiDiMode").Call(obj, uintptr(value))
}

func DrawGrid_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("DrawGrid_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func DrawGrid_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("DrawGrid_SetBorderStyle").Call(obj, uintptr(value))
}

func DrawGrid_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("DrawGrid_GetColor").Call(obj)
	return TColor(ret)
}

func DrawGrid_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("DrawGrid_SetColor").Call(obj, uintptr(value))
}

func DrawGrid_GetColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetColCount").Call(obj)
	return int32(ret)
}

func DrawGrid_SetColCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetColCount").Call(obj, uintptr(value))
}

func DrawGrid_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetConstraints").Call(obj)
	return ret
}

func DrawGrid_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetConstraints").Call(obj, value)
}

func DrawGrid_GetDefaultColWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetDefaultColWidth").Call(obj)
	return int32(ret)
}

func DrawGrid_SetDefaultColWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetDefaultColWidth").Call(obj, uintptr(value))
}

func DrawGrid_GetDefaultRowHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetDefaultRowHeight").Call(obj)
	return int32(ret)
}

func DrawGrid_SetDefaultRowHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetDefaultRowHeight").Call(obj, uintptr(value))
}

func DrawGrid_GetDefaultDrawing(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetDefaultDrawing").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetDefaultDrawing(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetDefaultDrawing").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DrawGrid_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func DrawGrid_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DrawGrid_SetDragCursor").Call(obj, uintptr(value))
}

func DrawGrid_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("DrawGrid_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func DrawGrid_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("DrawGrid_SetDragKind").Call(obj, uintptr(value))
}

func DrawGrid_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("DrawGrid_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func DrawGrid_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("DrawGrid_SetDragMode").Call(obj, uintptr(value))
}

func DrawGrid_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetFixedColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("DrawGrid_GetFixedColor").Call(obj)
	return TColor(ret)
}

func DrawGrid_SetFixedColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("DrawGrid_SetFixedColor").Call(obj, uintptr(value))
}

func DrawGrid_GetFixedCols(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetFixedCols").Call(obj)
	return int32(ret)
}

func DrawGrid_SetFixedCols(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetFixedCols").Call(obj, uintptr(value))
}

func DrawGrid_GetRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetRowCount").Call(obj)
	return int32(ret)
}

func DrawGrid_SetRowCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetRowCount").Call(obj, uintptr(value))
}

func DrawGrid_GetFixedRows(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetFixedRows").Call(obj)
	return int32(ret)
}

func DrawGrid_SetFixedRows(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetFixedRows").Call(obj, uintptr(value))
}

func DrawGrid_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetFont").Call(obj)
	return ret
}

func DrawGrid_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetFont").Call(obj, value)
}

func DrawGrid_GetGridLineWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetGridLineWidth").Call(obj)
	return int32(ret)
}

func DrawGrid_SetGridLineWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetGridLineWidth").Call(obj, uintptr(value))
}

func DrawGrid_GetOptions(obj uintptr) TGridOptions {
	ret, _, _ := getLazyProc("DrawGrid_GetOptions").Call(obj)
	return TGridOptions(ret)
}

func DrawGrid_SetOptions(obj uintptr, value TGridOptions) {
	_, _, _ = getLazyProc("DrawGrid_SetOptions").Call(obj, uintptr(value))
}

func DrawGrid_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetPopupMenu").Call(obj)
	return ret
}

func DrawGrid_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetPopupMenu").Call(obj, value)
}

func DrawGrid_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("DrawGrid_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func DrawGrid_SetScrollBars(obj uintptr, value TScrollStyle) {
	_, _, _ = getLazyProc("DrawGrid_SetScrollBars").Call(obj, uintptr(value))
}

func DrawGrid_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("DrawGrid_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func DrawGrid_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("DrawGrid_SetTabOrder").Call(obj, uintptr(value))
}

func DrawGrid_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetVisible").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetVisibleColCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetVisibleColCount").Call(obj)
	return int32(ret)
}

func DrawGrid_GetVisibleRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetVisibleRowCount").Call(obj)
	return int32(ret)
}

func DrawGrid_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnDrawCell(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnDrawCell").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnGetEditMask(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnGetEditMask").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnGetEditText(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnGetEditText").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseWheelDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnMouseWheelUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnSelectCell(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnSelectCell").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnSetEditText(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnSetEditText").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_SetOnTopLeftChanged(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DrawGrid_SetOnTopLeftChanged").Call(obj, addEventToMap(obj, fn))
}

func DrawGrid_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetCanvas").Call(obj)
	return ret
}

func DrawGrid_GetCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetCol").Call(obj)
	return int32(ret)
}

func DrawGrid_SetCol(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetCol").Call(obj, uintptr(value))
}

func DrawGrid_GetEditorMode(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetEditorMode").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetEditorMode(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetEditorMode").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetGridHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetGridHeight").Call(obj)
	return int32(ret)
}

func DrawGrid_GetGridWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetGridWidth").Call(obj)
	return int32(ret)
}

func DrawGrid_GetLeftCol(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetLeftCol").Call(obj)
	return int32(ret)
}

func DrawGrid_SetLeftCol(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetLeftCol").Call(obj, uintptr(value))
}

func DrawGrid_GetSelection(obj uintptr) TGridRect {
	var ret TGridRect
	_, _, _ = getLazyProc("DrawGrid_GetSelection").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_SetSelection(obj uintptr, value TGridRect) {
	_, _, _ = getLazyProc("DrawGrid_SetSelection").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DrawGrid_GetRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetRow").Call(obj)
	return int32(ret)
}

func DrawGrid_SetRow(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetRow").Call(obj, uintptr(value))
}

func DrawGrid_GetTopRow(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetTopRow").Call(obj)
	return int32(ret)
}

func DrawGrid_SetTopRow(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetTopRow").Call(obj, uintptr(value))
}

func DrawGrid_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetDockClientCount").Call(obj)
	return int32(ret)
}

func DrawGrid_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func DrawGrid_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetBrush").Call(obj)
	return ret
}

func DrawGrid_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetControlCount").Call(obj)
	return int32(ret)
}

func DrawGrid_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DrawGrid_GetHandle").Call(obj)
	return ret
}

func DrawGrid_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DrawGrid_GetParentWindow").Call(obj)
	return ret
}

func DrawGrid_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("DrawGrid_SetParentWindow").Call(obj, value)
}

func DrawGrid_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DrawGrid_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func DrawGrid_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAction").Call(obj)
	return ret
}

func DrawGrid_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetAction").Call(obj, value)
}

func DrawGrid_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DrawGrid_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("DrawGrid_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DrawGrid_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetClientHeight").Call(obj)
	return int32(ret)
}

func DrawGrid_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetClientHeight").Call(obj, uintptr(value))
}

func DrawGrid_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DrawGrid_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DrawGrid_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DrawGrid_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetClientWidth").Call(obj)
	return int32(ret)
}

func DrawGrid_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetClientWidth").Call(obj, uintptr(value))
}

func DrawGrid_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("DrawGrid_GetControlState").Call(obj)
	return TControlState(ret)
}

func DrawGrid_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("DrawGrid_SetControlState").Call(obj, uintptr(value))
}

func DrawGrid_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("DrawGrid_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func DrawGrid_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("DrawGrid_SetControlStyle").Call(obj, uintptr(value))
}

func DrawGrid_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("DrawGrid_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func DrawGrid_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetParent").Call(obj)
	return ret
}

func DrawGrid_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetParent").Call(obj, value)
}

func DrawGrid_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetLeft").Call(obj)
	return int32(ret)
}

func DrawGrid_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetLeft").Call(obj, uintptr(value))
}

func DrawGrid_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetTop").Call(obj)
	return int32(ret)
}

func DrawGrid_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetTop").Call(obj, uintptr(value))
}

func DrawGrid_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetWidth").Call(obj)
	return int32(ret)
}

func DrawGrid_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetWidth").Call(obj, uintptr(value))
}

func DrawGrid_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetHeight").Call(obj)
	return int32(ret)
}

func DrawGrid_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetHeight").Call(obj, uintptr(value))
}

func DrawGrid_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DrawGrid_GetCursor").Call(obj)
	return TCursor(ret)
}

func DrawGrid_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DrawGrid_SetCursor").Call(obj, uintptr(value))
}

func DrawGrid_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("DrawGrid_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func DrawGrid_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("DrawGrid_SetHint").Call(obj, GoStrToDStr(value))
}

func DrawGrid_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetComponentCount").Call(obj)
	return int32(ret)
}

func DrawGrid_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetComponentIndex").Call(obj)
	return int32(ret)
}

func DrawGrid_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetComponentIndex").Call(obj, uintptr(value))
}

func DrawGrid_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetOwner").Call(obj)
	return ret
}

func DrawGrid_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("DrawGrid_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func DrawGrid_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("DrawGrid_SetName").Call(obj, GoStrToDStr(value))
}

func DrawGrid_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("DrawGrid_GetTag").Call(obj)
	return int(ret)
}

func DrawGrid_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("DrawGrid_SetTag").Call(obj, uintptr(value))
}

func DrawGrid_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchorSideLeft").Call(obj)
	return ret
}

func DrawGrid_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetAnchorSideLeft").Call(obj, value)
}

func DrawGrid_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchorSideTop").Call(obj)
	return ret
}

func DrawGrid_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetAnchorSideTop").Call(obj, value)
}

func DrawGrid_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchorSideRight").Call(obj)
	return ret
}

func DrawGrid_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetAnchorSideRight").Call(obj, value)
}

func DrawGrid_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchorSideBottom").Call(obj)
	return ret
}

func DrawGrid_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetAnchorSideBottom").Call(obj, value)
}

func DrawGrid_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetChildSizing").Call(obj)
	return ret
}

func DrawGrid_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetChildSizing").Call(obj, value)
}

func DrawGrid_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetBorderSpacing").Call(obj)
	return ret
}

func DrawGrid_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DrawGrid_SetBorderSpacing").Call(obj, value)
}

func DrawGrid_GetColWidths(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetColWidths").Call(obj, uintptr(Index))
	return int32(ret)
}

func DrawGrid_SetColWidths(obj uintptr, Index int32, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetColWidths").Call(obj, uintptr(Index), uintptr(value))
}

func DrawGrid_GetRowHeights(obj uintptr, Index int32) int32 {
	ret, _, _ := getLazyProc("DrawGrid_GetRowHeights").Call(obj, uintptr(Index))
	return int32(ret)
}

func DrawGrid_SetRowHeights(obj uintptr, Index int32, value int32) {
	_, _, _ = getLazyProc("DrawGrid_SetRowHeights").Call(obj, uintptr(Index), uintptr(value))
}

func DrawGrid_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func DrawGrid_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetControls").Call(obj, uintptr(Index))
	return ret
}

func DrawGrid_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func DrawGrid_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("DrawGrid_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func DrawGrid_StaticClassType() TClass {
	r, _, _ := getLazyProc("DrawGrid_StaticClassType").Call()
	return TClass(r)
}
