package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TListView ---------------------------

func ListView_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_Create").Call(obj)
	return ret
}

func ListView_Free(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Free").Call(obj)
}

func ListView_AddItem(obj uintptr, Item string, AObject uintptr) {
	_, _, _ = getLazyProc("ListView_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ListView_AlphaSort(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_AlphaSort").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Clear").Call(obj)
}

func ListView_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("ListView_ClearSelection").Call(obj)
}

func ListView_DeleteSelected(obj uintptr) {
	_, _, _ = getLazyProc("ListView_DeleteSelected").Call(obj)
}

func ListView_GetHitTestInfoAt(obj uintptr, X int32, Y int32) THitTests {
	ret, _, _ := getLazyProc("ListView_GetHitTestInfoAt").Call(obj, uintptr(X), uintptr(Y))
	return THitTests(ret)
}

func ListView_GetItemAt(obj uintptr, X int32, Y int32) uintptr {
	ret, _, _ := getLazyProc("ListView_GetItemAt").Call(obj, uintptr(X), uintptr(Y))
	return ret
}

func ListView_GetNearestItem(obj uintptr, Point TPoint, Direction TSearchDirection) uintptr {
	ret, _, _ := getLazyProc("ListView_GetNearestItem").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(Direction))
	return ret
}

func ListView_GetNextItem(obj uintptr, StartItem uintptr, Direction TSearchDirection, States TListItemStates) uintptr {
	ret, _, _ := getLazyProc("ListView_GetNextItem").Call(obj, StartItem, uintptr(Direction), uintptr(States))
	return ret
}

func ListView_IsEditing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_IsEditing").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("ListView_SelectAll").Call(obj)
}

func ListView_CustomSort(obj uintptr, SortProc PFNLVCOMPARE, lParam int) bool {
	ret, _, _ := getLazyProc("ListView_CustomSort").Call(obj, uintptr(SortProc), uintptr(lParam))
	return DBoolToGoBool(ret)
}

func ListView_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ListView_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ListView_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ListView_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ListView_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ListView_DisableAlign").Call(obj)
}

func ListView_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ListView_EnableAlign").Call(obj)
}

func ListView_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ListView_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ListView_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("ListView_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ListView_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ListView_InsertControl").Call(obj, AControl)
}

func ListView_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Invalidate").Call(obj)
}

func ListView_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("ListView_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func ListView_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ListView_RemoveControl").Call(obj, AControl)
}

func ListView_Realign(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Realign").Call(obj)
}

func ListView_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Repaint").Call(obj)
}

func ListView_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("ListView_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ListView_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("ListView_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ListView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ListView_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ListView_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("ListView_SetFocus").Call(obj)
}

func ListView_Update(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Update").Call(obj)
}

func ListView_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ListView_BringToFront").Call(obj)
}

func ListView_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListView_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListView_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Hide").Call(obj)
}

func ListView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ListView_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ListView_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Refresh").Call(obj)
}

func ListView_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListView_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListView_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ListView_SendToBack").Call(obj)
}

func ListView_Show(obj uintptr) {
	_, _, _ = getLazyProc("ListView_Show").Call(obj)
}

func ListView_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ListView_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ListView_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetTextLen").Call(obj)
	return int32(ret)
}

func ListView_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ListView_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ListView_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ListView_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ListView_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ListView_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ListView_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ListView_Assign").Call(obj, Source)
}

func ListView_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ListView_ClassType").Call(obj)
	return TClass(ret)
}

func ListView_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListView_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ListView_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_InstanceSize").Call(obj)
	return int32(ret)
}

func ListView_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ListView_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ListView_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ListView_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetHashCode").Call(obj)
	return int32(ret)
}

func ListView_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ListView_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ListView_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ListView_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ListView_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ListView_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ListView_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ListView_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ListView_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ListView_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ListView_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ListView_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ListView_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ListView_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ListView_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ListView_AnchorClient").Call(obj, uintptr(ASpace))
}

func ListView_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ListView_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ListView_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ListView_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ListView_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ListView_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ListView_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ListView_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ListView_GetAutoSort(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetAutoSort").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetAutoSort(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetAutoSort").Call(obj, GoBoolToDBool(value))
}

func ListView_GetAutoSortIndicator(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetAutoSortIndicator").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetAutoSortIndicator(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetAutoSortIndicator").Call(obj, GoBoolToDBool(value))
}

func ListView_GetAutoWidthLastColumn(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetAutoWidthLastColumn").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetAutoWidthLastColumn(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetAutoWidthLastColumn").Call(obj, GoBoolToDBool(value))
}

func ListView_GetSmallImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetSmallImagesWidth").Call(obj)
	return int32(ret)
}

func ListView_SetSmallImagesWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetSmallImagesWidth").Call(obj, uintptr(value))
}

func ListView_GetSortColumn(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetSortColumn").Call(obj)
	return int32(ret)
}

func ListView_SetSortColumn(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetSortColumn").Call(obj, uintptr(value))
}

func ListView_GetSortDirection(obj uintptr) TSortDirection {
	ret, _, _ := getLazyProc("ListView_GetSortDirection").Call(obj)
	return TSortDirection(ret)
}

func ListView_SetSortDirection(obj uintptr, value TSortDirection) {
	_, _, _ = getLazyProc("ListView_SetSortDirection").Call(obj, uintptr(value))
}

func ListView_GetLargeImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetLargeImagesWidth").Call(obj)
	return int32(ret)
}

func ListView_SetLargeImagesWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetLargeImagesWidth").Call(obj, uintptr(value))
}

func ListView_GetStateImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetStateImagesWidth").Call(obj)
	return int32(ret)
}

func ListView_SetStateImagesWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetStateImagesWidth").Call(obj, uintptr(value))
}

func ListView_GetToolTips(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetToolTips").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetToolTips(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetToolTips").Call(obj, GoBoolToDBool(value))
}

func ListView_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("ListView_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func ListView_SetScrollBars(obj uintptr, value TScrollStyle) {
	_, _, _ = getLazyProc("ListView_SetScrollBars").Call(obj, uintptr(value))
}

func ListView_GetColumnCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetColumnCount").Call(obj)
	return int32(ret)
}

func ListView_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAction").Call(obj)
	return ret
}

func ListView_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetAction").Call(obj, value)
}

func ListView_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ListView_GetAlign").Call(obj)
	return TAlign(ret)
}

func ListView_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ListView_SetAlign").Call(obj, uintptr(value))
}

func ListView_GetAllocBy(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetAllocBy").Call(obj)
	return int32(ret)
}

func ListView_SetAllocBy(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetAllocBy").Call(obj, uintptr(value))
}

func ListView_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ListView_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ListView_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ListView_SetAnchors").Call(obj, uintptr(value))
}

func ListView_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ListView_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ListView_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ListView_SetBiDiMode").Call(obj, uintptr(value))
}

func ListView_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("ListView_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func ListView_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("ListView_SetBorderStyle").Call(obj, uintptr(value))
}

func ListView_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetBorderWidth").Call(obj)
	return int32(ret)
}

func ListView_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetBorderWidth").Call(obj, uintptr(value))
}

func ListView_GetCheckboxes(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetCheckboxes").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetCheckboxes(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetCheckboxes").Call(obj, GoBoolToDBool(value))
}

func ListView_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ListView_GetColor").Call(obj)
	return TColor(ret)
}

func ListView_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ListView_SetColor").Call(obj, uintptr(value))
}

func ListView_GetColumns(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetColumns").Call(obj)
	return ret
}

func ListView_SetColumns(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetColumns").Call(obj, value)
}

func ListView_GetColumnClick(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetColumnClick").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetColumnClick(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetColumnClick").Call(obj, GoBoolToDBool(value))
}

func ListView_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetConstraints").Call(obj)
	return ret
}

func ListView_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetConstraints").Call(obj, value)
}

func ListView_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ListView_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ListView_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ListView_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ListView_SetDragCursor").Call(obj, uintptr(value))
}

func ListView_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ListView_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ListView_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("ListView_SetDragKind").Call(obj, uintptr(value))
}

func ListView_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ListView_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ListView_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("ListView_SetDragMode").Call(obj, uintptr(value))
}

func ListView_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ListView_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetFont").Call(obj)
	return ret
}

func ListView_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetFont").Call(obj, value)
}

func ListView_GetFlatScrollBars(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetFlatScrollBars").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetFlatScrollBars(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetFlatScrollBars").Call(obj, GoBoolToDBool(value))
}

func ListView_GetFullDrag(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetFullDrag").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetFullDrag(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetFullDrag").Call(obj, GoBoolToDBool(value))
}

func ListView_GetGridLines(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetGridLines").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetGridLines(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetGridLines").Call(obj, GoBoolToDBool(value))
}

func ListView_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func ListView_GetHotTrack(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetHotTrack").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetHotTrack(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetHotTrack").Call(obj, GoBoolToDBool(value))
}

func ListView_GetIconOptions(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetIconOptions").Call(obj)
	return ret
}

func ListView_SetIconOptions(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetIconOptions").Call(obj, value)
}

func ListView_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetItems").Call(obj)
	return ret
}

func ListView_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetItems").Call(obj, value)
}

func ListView_GetLargeImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetLargeImages").Call(obj)
	return ret
}

func ListView_SetLargeImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetLargeImages").Call(obj, value)
}

func ListView_GetMultiSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetMultiSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetMultiSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetMultiSelect").Call(obj, GoBoolToDBool(value))
}

func ListView_GetOwnerData(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetOwnerData").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetOwnerData(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetOwnerData").Call(obj, GoBoolToDBool(value))
}

func ListView_GetOwnerDraw(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetOwnerDraw").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetOwnerDraw(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetOwnerDraw").Call(obj, GoBoolToDBool(value))
}

func ListView_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func ListView_GetRowSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetRowSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetRowSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetRowSelect").Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ListView_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetPopupMenu").Call(obj)
	return ret
}

func ListView_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetPopupMenu").Call(obj, value)
}

func ListView_GetShowColumnHeaders(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetShowColumnHeaders").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetShowColumnHeaders(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetShowColumnHeaders").Call(obj, GoBoolToDBool(value))
}

func ListView_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ListView_GetSmallImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetSmallImages").Call(obj)
	return ret
}

func ListView_SetSmallImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetSmallImages").Call(obj, value)
}

func ListView_GetSortType(obj uintptr) TSortType {
	ret, _, _ := getLazyProc("ListView_GetSortType").Call(obj)
	return TSortType(ret)
}

func ListView_SetSortType(obj uintptr, value TSortType) {
	_, _, _ = getLazyProc("ListView_SetSortType").Call(obj, uintptr(value))
}

func ListView_GetStateImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetStateImages").Call(obj)
	return ret
}

func ListView_SetStateImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetStateImages").Call(obj, value)
}

func ListView_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ListView_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ListView_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("ListView_SetTabOrder").Call(obj, uintptr(value))
}

func ListView_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ListView_GetViewStyle(obj uintptr) TViewStyle {
	ret, _, _ := getLazyProc("ListView_GetViewStyle").Call(obj)
	return TViewStyle(ret)
}

func ListView_SetViewStyle(obj uintptr, value TViewStyle) {
	_, _, _ = getLazyProc("ListView_SetViewStyle").Call(obj, uintptr(value))
}

func ListView_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ListView_SetOnAdvancedCustomDraw(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnAdvancedCustomDraw").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnAdvancedCustomDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnAdvancedCustomDrawItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnAdvancedCustomDrawSubItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnAdvancedCustomDrawSubItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnColumnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnColumnClick").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnCompare(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnCompare").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnCustomDraw(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnCustomDraw").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnCustomDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnCustomDrawItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnCustomDrawSubItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnCustomDrawSubItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnData(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnData").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDataFind(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDataFind").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDataHint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDataHint").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDeletion(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDeletion").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDrawItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnEdited(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnEdited").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnEditing(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnEditing").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnInsert(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnInsert").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnSelectItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnSelectItem").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnItemChecked(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnItemChecked").Call(obj, addEventToMap(obj, fn))
}

func ListView_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ListView_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func ListView_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetCanvas").Call(obj)
	return ret
}

func ListView_GetDropTarget(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetDropTarget").Call(obj)
	return ret
}

func ListView_SetDropTarget(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetDropTarget").Call(obj, value)
}

func ListView_GetItemFocused(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetItemFocused").Call(obj)
	return ret
}

func ListView_SetItemFocused(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetItemFocused").Call(obj, value)
}

func ListView_GetSelCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetSelCount").Call(obj)
	return int32(ret)
}

func ListView_GetSelected(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetSelected").Call(obj)
	return ret
}

func ListView_SetSelected(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetSelected").Call(obj, value)
}

func ListView_GetTopItem(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetTopItem").Call(obj)
	return ret
}

func ListView_GetVisibleRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetVisibleRowCount").Call(obj)
	return int32(ret)
}

func ListView_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetItemIndex").Call(obj)
	return int32(ret)
}

func ListView_SetItemIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetItemIndex").Call(obj, uintptr(value))
}

func ListView_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ListView_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ListView_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ListView_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetBrush").Call(obj)
	return ret
}

func ListView_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetControlCount").Call(obj)
	return int32(ret)
}

func ListView_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ListView_GetHandle").Call(obj)
	return ret
}

func ListView_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ListView_GetParentWindow").Call(obj)
	return ret
}

func ListView_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("ListView_SetParentWindow").Call(obj, value)
}

func ListView_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ListView_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ListView_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ListView_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ListView_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListView_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetClientHeight").Call(obj)
	return int32(ret)
}

func ListView_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetClientHeight").Call(obj, uintptr(value))
}

func ListView_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ListView_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ListView_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ListView_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetClientWidth").Call(obj)
	return int32(ret)
}

func ListView_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetClientWidth").Call(obj, uintptr(value))
}

func ListView_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ListView_GetControlState").Call(obj)
	return TControlState(ret)
}

func ListView_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ListView_SetControlState").Call(obj, uintptr(value))
}

func ListView_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ListView_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ListView_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ListView_SetControlStyle").Call(obj, uintptr(value))
}

func ListView_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ListView_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ListView_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetParent").Call(obj)
	return ret
}

func ListView_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetParent").Call(obj, value)
}

func ListView_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetLeft").Call(obj)
	return int32(ret)
}

func ListView_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetLeft").Call(obj, uintptr(value))
}

func ListView_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetTop").Call(obj)
	return int32(ret)
}

func ListView_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetTop").Call(obj, uintptr(value))
}

func ListView_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetWidth").Call(obj)
	return int32(ret)
}

func ListView_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetWidth").Call(obj, uintptr(value))
}

func ListView_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetHeight").Call(obj)
	return int32(ret)
}

func ListView_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetHeight").Call(obj, uintptr(value))
}

func ListView_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ListView_GetCursor").Call(obj)
	return TCursor(ret)
}

func ListView_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ListView_SetCursor").Call(obj, uintptr(value))
}

func ListView_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ListView_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ListView_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ListView_SetHint").Call(obj, GoStrToDStr(value))
}

func ListView_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetComponentCount").Call(obj)
	return int32(ret)
}

func ListView_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ListView_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ListView_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ListView_SetComponentIndex").Call(obj, uintptr(value))
}

func ListView_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetOwner").Call(obj)
	return ret
}

func ListView_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ListView_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ListView_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ListView_SetName").Call(obj, GoStrToDStr(value))
}

func ListView_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ListView_GetTag").Call(obj)
	return int(ret)
}

func ListView_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ListView_SetTag").Call(obj, uintptr(value))
}

func ListView_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAnchorSideLeft").Call(obj)
	return ret
}

func ListView_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetAnchorSideLeft").Call(obj, value)
}

func ListView_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAnchorSideTop").Call(obj)
	return ret
}

func ListView_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetAnchorSideTop").Call(obj, value)
}

func ListView_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAnchorSideRight").Call(obj)
	return ret
}

func ListView_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetAnchorSideRight").Call(obj, value)
}

func ListView_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAnchorSideBottom").Call(obj)
	return ret
}

func ListView_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetAnchorSideBottom").Call(obj, value)
}

func ListView_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetChildSizing").Call(obj)
	return ret
}

func ListView_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetChildSizing").Call(obj, value)
}

func ListView_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ListView_GetBorderSpacing").Call(obj)
	return ret
}

func ListView_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ListView_SetBorderSpacing").Call(obj, value)
}

func ListView_GetColumn(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListView_GetColumn").Call(obj, uintptr(Index))
	return ret
}

func ListView_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListView_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ListView_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ListView_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ListView_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ListView_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ListView_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ListView_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ListView_StaticClassType() TClass {
	r, _, _ := getLazyProc("ListView_StaticClassType").Call()
	return TClass(r)
}
