package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TTreeView ---------------------------

func TreeView_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_Create").Call(obj)
	return ret
}

func TreeView_Free(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Free").Call(obj)
}

func TreeView_AlphaSort(obj uintptr, ARecurse bool) bool {
	ret, _, _ := getLazyProc("TreeView_AlphaSort").Call(obj, GoBoolToDBool(ARecurse))
	return DBoolToGoBool(ret)
}

func TreeView_FullCollapse(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_FullCollapse").Call(obj)
}

func TreeView_FullExpand(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_FullExpand").Call(obj)
}

func TreeView_GetHitTestInfoAt(obj uintptr, X int32, Y int32) THitTests {
	ret, _, _ := getLazyProc("TreeView_GetHitTestInfoAt").Call(obj, uintptr(X), uintptr(Y))
	return THitTests(ret)
}

func TreeView_GetNodeAt(obj uintptr, X int32, Y int32) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetNodeAt").Call(obj, uintptr(X), uintptr(Y))
	return ret
}

func TreeView_IsEditing(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_IsEditing").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_LoadFromFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("TreeView_LoadFromFile").Call(obj, GoStrToDStr(FileName))
}

func TreeView_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("TreeView_LoadFromStream").Call(obj, Stream)
}

func TreeView_SaveToFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("TreeView_SaveToFile").Call(obj, GoStrToDStr(FileName))
}

func TreeView_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("TreeView_SaveToStream").Call(obj, Stream)
}

func TreeView_ClearSelection(obj uintptr, KeepPrimary bool) {
	_, _, _ = getLazyProc("TreeView_ClearSelection").Call(obj, GoBoolToDBool(KeepPrimary))
}

func TreeView_CustomSort(obj uintptr, SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
	ret, _, _ := getLazyProc("TreeView_CustomSort").Call(obj, uintptr(SortProc), uintptr(Data), GoBoolToDBool(ARecurse))
	return DBoolToGoBool(ret)
}

func TreeView_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func TreeView_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("TreeView_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func TreeView_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_DisableAlign").Call(obj)
}

func TreeView_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_EnableAlign").Call(obj)
}

func TreeView_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("TreeView_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func TreeView_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("TreeView_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func TreeView_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("TreeView_InsertControl").Call(obj, AControl)
}

func TreeView_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Invalidate").Call(obj)
}

func TreeView_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("TreeView_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func TreeView_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("TreeView_RemoveControl").Call(obj, AControl)
}

func TreeView_Realign(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Realign").Call(obj)
}

func TreeView_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Repaint").Call(obj)
}

func TreeView_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("TreeView_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func TreeView_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("TreeView_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func TreeView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("TreeView_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func TreeView_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_SetFocus").Call(obj)
}

func TreeView_Update(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Update").Call(obj)
}

func TreeView_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_BringToFront").Call(obj)
}

func TreeView_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TreeView_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TreeView_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_Hide(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Hide").Call(obj)
}

func TreeView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("TreeView_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func TreeView_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Refresh").Call(obj)
}

func TreeView_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TreeView_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TreeView_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_SendToBack").Call(obj)
}

func TreeView_Show(obj uintptr) {
	_, _, _ = getLazyProc("TreeView_Show").Call(obj)
}

func TreeView_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("TreeView_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func TreeView_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetTextLen").Call(obj)
	return int32(ret)
}

func TreeView_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("TreeView_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func TreeView_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("TreeView_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func TreeView_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeView_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TreeView_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TreeView_Assign").Call(obj, Source)
}

func TreeView_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TreeView_ClassType").Call(obj)
	return TClass(ret)
}

func TreeView_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeView_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TreeView_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_InstanceSize").Call(obj)
	return int32(ret)
}

func TreeView_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TreeView_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TreeView_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TreeView_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetHashCode").Call(obj)
	return int32(ret)
}

func TreeView_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeView_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TreeView_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("TreeView_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TreeView_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("TreeView_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TreeView_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("TreeView_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func TreeView_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("TreeView_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func TreeView_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("TreeView_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func TreeView_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("TreeView_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func TreeView_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("TreeView_AnchorClient").Call(obj, uintptr(ASpace))
}

func TreeView_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TreeView_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TreeView_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("TreeView_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func TreeView_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("TreeView_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func TreeView_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("TreeView_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func TreeView_GetDefaultItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetDefaultItemHeight").Call(obj)
	return int32(ret)
}

func TreeView_SetDefaultItemHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetDefaultItemHeight").Call(obj, uintptr(value))
}

func TreeView_GetExpandSignColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetExpandSignColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetExpandSignColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetExpandSignColor").Call(obj, uintptr(value))
}

func TreeView_GetExpandSignSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetExpandSignSize").Call(obj)
	return int32(ret)
}

func TreeView_SetExpandSignSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetExpandSignSize").Call(obj, uintptr(value))
}

func TreeView_GetExpandSignType(obj uintptr) TTreeViewExpandSignType {
	ret, _, _ := getLazyProc("TreeView_GetExpandSignType").Call(obj)
	return TTreeViewExpandSignType(ret)
}

func TreeView_SetExpandSignType(obj uintptr, value TTreeViewExpandSignType) {
	_, _, _ = getLazyProc("TreeView_SetExpandSignType").Call(obj, uintptr(value))
}

func TreeView_GetHotTrackColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetHotTrackColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetHotTrackColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetHotTrackColor").Call(obj, uintptr(value))
}

func TreeView_GetImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetImagesWidth").Call(obj)
	return int32(ret)
}

func TreeView_SetImagesWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetImagesWidth").Call(obj, uintptr(value))
}

func TreeView_GetOptions(obj uintptr) TTreeViewOptions {
	ret, _, _ := getLazyProc("TreeView_GetOptions").Call(obj)
	return TTreeViewOptions(ret)
}

func TreeView_SetOptions(obj uintptr, value TTreeViewOptions) {
	_, _, _ = getLazyProc("TreeView_SetOptions").Call(obj, uintptr(value))
}

func TreeView_GetScrollBars(obj uintptr) TScrollStyle {
	ret, _, _ := getLazyProc("TreeView_GetScrollBars").Call(obj)
	return TScrollStyle(ret)
}

func TreeView_SetScrollBars(obj uintptr, value TScrollStyle) {
	_, _, _ = getLazyProc("TreeView_SetScrollBars").Call(obj, uintptr(value))
}

func TreeView_GetSelectionColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetSelectionColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetSelectionColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetSelectionColor").Call(obj, uintptr(value))
}

func TreeView_GetSelectionFontColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetSelectionFontColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetSelectionFontColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetSelectionFontColor").Call(obj, uintptr(value))
}

func TreeView_GetSelectionFontColorUsed(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetSelectionFontColorUsed").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetSelectionFontColorUsed(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetSelectionFontColorUsed").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetSeparatorColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetSeparatorColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetSeparatorColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetSeparatorColor").Call(obj, uintptr(value))
}

func TreeView_GetStateImagesWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetStateImagesWidth").Call(obj)
	return int32(ret)
}

func TreeView_SetStateImagesWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetStateImagesWidth").Call(obj, uintptr(value))
}

func TreeView_GetToolTips(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetToolTips").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetToolTips(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetToolTips").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetTreeLineColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetTreeLineColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetTreeLineColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetTreeLineColor").Call(obj, uintptr(value))
}

func TreeView_GetTreeLinePenStyle(obj uintptr) TPenStyle {
	ret, _, _ := getLazyProc("TreeView_GetTreeLinePenStyle").Call(obj)
	return TPenStyle(ret)
}

func TreeView_SetTreeLinePenStyle(obj uintptr, value TPenStyle) {
	_, _, _ = getLazyProc("TreeView_SetTreeLinePenStyle").Call(obj, uintptr(value))
}

func TreeView_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("TreeView_GetAlign").Call(obj)
	return TAlign(ret)
}

func TreeView_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("TreeView_SetAlign").Call(obj, uintptr(value))
}

func TreeView_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("TreeView_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func TreeView_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("TreeView_SetAnchors").Call(obj, uintptr(value))
}

func TreeView_GetAutoExpand(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetAutoExpand").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetAutoExpand(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetAutoExpand").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("TreeView_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func TreeView_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("TreeView_SetBiDiMode").Call(obj, uintptr(value))
}

func TreeView_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("TreeView_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func TreeView_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("TreeView_SetBorderStyle").Call(obj, uintptr(value))
}

func TreeView_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetBorderWidth").Call(obj)
	return int32(ret)
}

func TreeView_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetBorderWidth").Call(obj, uintptr(value))
}

func TreeView_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("TreeView_GetColor").Call(obj)
	return TColor(ret)
}

func TreeView_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("TreeView_SetColor").Call(obj, uintptr(value))
}

func TreeView_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetConstraints").Call(obj)
	return ret
}

func TreeView_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetConstraints").Call(obj, value)
}

func TreeView_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("TreeView_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func TreeView_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("TreeView_SetDragKind").Call(obj, uintptr(value))
}

func TreeView_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("TreeView_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func TreeView_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("TreeView_SetDragCursor").Call(obj, uintptr(value))
}

func TreeView_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("TreeView_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func TreeView_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("TreeView_SetDragMode").Call(obj, uintptr(value))
}

func TreeView_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetFont").Call(obj)
	return ret
}

func TreeView_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetFont").Call(obj, value)
}

func TreeView_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetHideSelection").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetHideSelection").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetHotTrack(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetHotTrack").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetHotTrack(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetHotTrack").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetImages").Call(obj)
	return ret
}

func TreeView_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetImages").Call(obj, value)
}

func TreeView_GetIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetIndent").Call(obj)
	return int32(ret)
}

func TreeView_SetIndent(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetIndent").Call(obj, uintptr(value))
}

func TreeView_GetMultiSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetMultiSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetMultiSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetMultiSelect").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetMultiSelectStyle(obj uintptr) TMultiSelectStyle {
	ret, _, _ := getLazyProc("TreeView_GetMultiSelectStyle").Call(obj)
	return TMultiSelectStyle(ret)
}

func TreeView_SetMultiSelectStyle(obj uintptr, value TMultiSelectStyle) {
	_, _, _ = getLazyProc("TreeView_SetMultiSelectStyle").Call(obj, uintptr(value))
}

func TreeView_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetPopupMenu").Call(obj)
	return ret
}

func TreeView_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetPopupMenu").Call(obj, value)
}

func TreeView_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRightClickSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetRightClickSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetRightClickSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetRightClickSelect").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRowSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetRowSelect").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetRowSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetRowSelect").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowButtons(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetShowButtons").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetShowButtons(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetShowButtons").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowLines(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetShowLines").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetShowLines(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetShowLines").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowRoot(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetShowRoot").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetShowRoot(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetShowRoot").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetSortType(obj uintptr) TSortType {
	ret, _, _ := getLazyProc("TreeView_GetSortType").Call(obj)
	return TSortType(ret)
}

func TreeView_SetSortType(obj uintptr, value TSortType) {
	_, _, _ = getLazyProc("TreeView_SetSortType").Call(obj, uintptr(value))
}

func TreeView_GetStateImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetStateImages").Call(obj)
	return ret
}

func TreeView_SetStateImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetStateImages").Call(obj, value)
}

func TreeView_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("TreeView_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func TreeView_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("TreeView_SetTabOrder").Call(obj, uintptr(value))
}

func TreeView_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetVisible").Call(obj, GoBoolToDBool(value))
}

func TreeView_SetOnAddition(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnAddition").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnAdvancedCustomDraw(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnAdvancedCustomDraw").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnAdvancedCustomDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnAdvancedCustomDrawItem").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnChanging(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnChanging").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnCollapsed(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnCollapsed").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnCollapsing(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnCollapsing").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnCompare(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnCompare").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnCustomDraw(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnCustomDraw").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnCustomDrawItem(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnCustomDrawItem").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnDeletion(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnDeletion").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnEdited(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnEdited").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnEditing(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnEditing").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnExpanding(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnExpanding").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnExpanded(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnExpanded").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnGetSelectedIndex(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnGetSelectedIndex").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func TreeView_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TreeView_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func TreeView_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetItems").Call(obj)
	return ret
}

func TreeView_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetItems").Call(obj, value)
}

func TreeView_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetCanvas").Call(obj)
	return ret
}

func TreeView_GetDropTarget(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetDropTarget").Call(obj)
	return ret
}

func TreeView_SetDropTarget(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetDropTarget").Call(obj, value)
}

func TreeView_GetSelected(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetSelected").Call(obj)
	return ret
}

func TreeView_SetSelected(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetSelected").Call(obj, value)
}

func TreeView_GetTopItem(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetTopItem").Call(obj)
	return ret
}

func TreeView_SetTopItem(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetTopItem").Call(obj, value)
}

func TreeView_GetSelectionCount(obj uintptr) uint32 {
	ret, _, _ := getLazyProc("TreeView_GetSelectionCount").Call(obj)
	return uint32(ret)
}

func TreeView_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetDockClientCount").Call(obj)
	return int32(ret)
}

func TreeView_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func TreeView_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetBrush").Call(obj)
	return ret
}

func TreeView_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetControlCount").Call(obj)
	return int32(ret)
}

func TreeView_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TreeView_GetHandle").Call(obj)
	return ret
}

func TreeView_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TreeView_GetParentWindow").Call(obj)
	return ret
}

func TreeView_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("TreeView_SetParentWindow").Call(obj, value)
}

func TreeView_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeView_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func TreeView_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAction").Call(obj)
	return ret
}

func TreeView_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetAction").Call(obj, value)
}

func TreeView_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TreeView_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("TreeView_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TreeView_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetClientHeight").Call(obj)
	return int32(ret)
}

func TreeView_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetClientHeight").Call(obj, uintptr(value))
}

func TreeView_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TreeView_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TreeView_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeView_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetClientWidth").Call(obj)
	return int32(ret)
}

func TreeView_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetClientWidth").Call(obj, uintptr(value))
}

func TreeView_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("TreeView_GetControlState").Call(obj)
	return TControlState(ret)
}

func TreeView_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("TreeView_SetControlState").Call(obj, uintptr(value))
}

func TreeView_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("TreeView_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func TreeView_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("TreeView_SetControlStyle").Call(obj, uintptr(value))
}

func TreeView_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeView_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeView_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetParent").Call(obj)
	return ret
}

func TreeView_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetParent").Call(obj, value)
}

func TreeView_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetLeft").Call(obj)
	return int32(ret)
}

func TreeView_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetLeft").Call(obj, uintptr(value))
}

func TreeView_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetTop").Call(obj)
	return int32(ret)
}

func TreeView_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetTop").Call(obj, uintptr(value))
}

func TreeView_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetWidth").Call(obj)
	return int32(ret)
}

func TreeView_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetWidth").Call(obj, uintptr(value))
}

func TreeView_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetHeight").Call(obj)
	return int32(ret)
}

func TreeView_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetHeight").Call(obj, uintptr(value))
}

func TreeView_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("TreeView_GetCursor").Call(obj)
	return TCursor(ret)
}

func TreeView_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("TreeView_SetCursor").Call(obj, uintptr(value))
}

func TreeView_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeView_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func TreeView_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("TreeView_SetHint").Call(obj, GoStrToDStr(value))
}

func TreeView_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetComponentCount").Call(obj)
	return int32(ret)
}

func TreeView_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeView_GetComponentIndex").Call(obj)
	return int32(ret)
}

func TreeView_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeView_SetComponentIndex").Call(obj, uintptr(value))
}

func TreeView_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetOwner").Call(obj)
	return ret
}

func TreeView_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeView_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TreeView_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("TreeView_SetName").Call(obj, GoStrToDStr(value))
}

func TreeView_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("TreeView_GetTag").Call(obj)
	return int(ret)
}

func TreeView_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("TreeView_SetTag").Call(obj, uintptr(value))
}

func TreeView_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAnchorSideLeft").Call(obj)
	return ret
}

func TreeView_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetAnchorSideLeft").Call(obj, value)
}

func TreeView_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAnchorSideTop").Call(obj)
	return ret
}

func TreeView_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetAnchorSideTop").Call(obj, value)
}

func TreeView_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAnchorSideRight").Call(obj)
	return ret
}

func TreeView_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetAnchorSideRight").Call(obj, value)
}

func TreeView_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAnchorSideBottom").Call(obj)
	return ret
}

func TreeView_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetAnchorSideBottom").Call(obj, value)
}

func TreeView_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetChildSizing").Call(obj)
	return ret
}

func TreeView_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetChildSizing").Call(obj, value)
}

func TreeView_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetBorderSpacing").Call(obj)
	return ret
}

func TreeView_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TreeView_SetBorderSpacing").Call(obj, value)
}

func TreeView_GetSelections(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetSelections").Call(obj, uintptr(Index))
	return ret
}

func TreeView_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func TreeView_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetControls").Call(obj, uintptr(Index))
	return ret
}

func TreeView_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func TreeView_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("TreeView_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func TreeView_StaticClassType() TClass {
	r, _, _ := getLazyProc("TreeView_StaticClassType").Call()
	return TClass(r)
}
