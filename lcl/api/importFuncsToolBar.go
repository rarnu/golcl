package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TToolBar ---------------------------

func ToolBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_Create").Call(obj)
	return ret
}

func ToolBar_Free(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Free").Call(obj)
}

func ToolBar_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("ToolBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ToolBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ToolBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ToolBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ToolBar_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_DisableAlign").Call(obj)
}

func ToolBar_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_EnableAlign").Call(obj)
}

func ToolBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ToolBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ToolBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ToolBar_InsertControl").Call(obj, AControl)
}

func ToolBar_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Invalidate").Call(obj)
}

func ToolBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("ToolBar_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func ToolBar_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ToolBar_RemoveControl").Call(obj, AControl)
}

func ToolBar_Realign(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Realign").Call(obj)
}

func ToolBar_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Repaint").Call(obj)
}

func ToolBar_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("ToolBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ToolBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("ToolBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ToolBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ToolBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ToolBar_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetFocus").Call(obj)
}

func ToolBar_Update(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Update").Call(obj)
}

func ToolBar_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_BringToFront").Call(obj)
}

func ToolBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ToolBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ToolBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Hide").Call(obj)
}

func ToolBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ToolBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ToolBar_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Refresh").Call(obj)
}

func ToolBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ToolBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ToolBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_SendToBack").Call(obj)
}

func ToolBar_Show(obj uintptr) {
	_, _, _ = getLazyProc("ToolBar_Show").Call(obj)
}

func ToolBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ToolBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ToolBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetTextLen").Call(obj)
	return int32(ret)
}

func ToolBar_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ToolBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ToolBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ToolBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ToolBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ToolBar_Assign").Call(obj, Source)
}

func ToolBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ToolBar_ClassType").Call(obj)
	return TClass(ret)
}

func ToolBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_InstanceSize").Call(obj)
	return int32(ret)
}

func ToolBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ToolBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ToolBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ToolBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetHashCode").Call(obj)
	return int32(ret)
}

func ToolBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ToolBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToolBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ToolBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToolBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ToolBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ToolBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ToolBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ToolBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ToolBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ToolBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ToolBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ToolBar_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ToolBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func ToolBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ToolBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ToolBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ToolBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ToolBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ToolBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ToolBar_GetButtonCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetButtonCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetCanvas").Call(obj)
	return ret
}

func ToolBar_GetRowCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetRowCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ToolBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func ToolBar_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ToolBar_SetAlign").Call(obj, uintptr(value))
}

func ToolBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ToolBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ToolBar_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ToolBar_SetAnchors").Call(obj, uintptr(value))
}

func ToolBar_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func ToolBar_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetBorderWidth").Call(obj, uintptr(value))
}

func ToolBar_GetButtonHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetButtonHeight").Call(obj)
	return int32(ret)
}

func ToolBar_SetButtonHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetButtonHeight").Call(obj, uintptr(value))
}

func ToolBar_GetButtonWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetButtonWidth").Call(obj)
	return int32(ret)
}

func ToolBar_SetButtonWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetButtonWidth").Call(obj, uintptr(value))
}

func ToolBar_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("ToolBar_SetCaption").Call(obj, GoStrToDStr(value))
}

func ToolBar_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ToolBar_GetColor").Call(obj)
	return TColor(ret)
}

func ToolBar_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ToolBar_SetColor").Call(obj, uintptr(value))
}

func ToolBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetConstraints").Call(obj)
	return ret
}

func ToolBar_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetConstraints").Call(obj, value)
}

func ToolBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToolBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ToolBar_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ToolBar_SetDragCursor").Call(obj, uintptr(value))
}

func ToolBar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ToolBar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ToolBar_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("ToolBar_SetDragKind").Call(obj, uintptr(value))
}

func ToolBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ToolBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ToolBar_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("ToolBar_SetDragMode").Call(obj, uintptr(value))
}

func ToolBar_GetEdgeBorders(obj uintptr) TEdgeBorders {
	ret, _, _ := getLazyProc("ToolBar_GetEdgeBorders").Call(obj)
	return TEdgeBorders(ret)
}

func ToolBar_SetEdgeBorders(obj uintptr, value TEdgeBorders) {
	_, _, _ = getLazyProc("ToolBar_SetEdgeBorders").Call(obj, uintptr(value))
}

func ToolBar_GetEdgeInner(obj uintptr) TEdgeStyle {
	ret, _, _ := getLazyProc("ToolBar_GetEdgeInner").Call(obj)
	return TEdgeStyle(ret)
}

func ToolBar_SetEdgeInner(obj uintptr, value TEdgeStyle) {
	_, _, _ = getLazyProc("ToolBar_SetEdgeInner").Call(obj, uintptr(value))
}

func ToolBar_GetEdgeOuter(obj uintptr) TEdgeStyle {
	ret, _, _ := getLazyProc("ToolBar_GetEdgeOuter").Call(obj)
	return TEdgeStyle(ret)
}

func ToolBar_SetEdgeOuter(obj uintptr, value TEdgeStyle) {
	_, _, _ = getLazyProc("ToolBar_SetEdgeOuter").Call(obj, uintptr(value))
}

func ToolBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFlat(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetFlat").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetFlat(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetFlat").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetFont").Call(obj)
	return ret
}

func ToolBar_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetFont").Call(obj, value)
}

func ToolBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetHeight").Call(obj)
	return int32(ret)
}

func ToolBar_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetHeight").Call(obj, uintptr(value))
}

func ToolBar_GetHotImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetHotImages").Call(obj)
	return ret
}

func ToolBar_SetHotImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetHotImages").Call(obj, value)
}

func ToolBar_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetImages").Call(obj)
	return ret
}

func ToolBar_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetImages").Call(obj, value)
}

func ToolBar_GetIndent(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetIndent").Call(obj)
	return int32(ret)
}

func ToolBar_SetIndent(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetIndent").Call(obj, uintptr(value))
}

func ToolBar_GetList(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetList").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetList(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetList").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetPopupMenu").Call(obj)
	return ret
}

func ToolBar_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetPopupMenu").Call(obj, value)
}

func ToolBar_GetShowCaptions(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetShowCaptions").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetShowCaptions(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetShowCaptions").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ToolBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ToolBar_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("ToolBar_SetTabOrder").Call(obj, uintptr(value))
}

func ToolBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetWrapable(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetWrapable").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetWrapable(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetWrapable").Call(obj, GoBoolToDBool(value))
}

func ToolBar_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ToolBar_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func ToolBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetBrush").Call(obj)
	return ret
}

func ToolBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetControlCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ToolBar_GetHandle").Call(obj)
	return ret
}

func ToolBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ToolBar_GetParentWindow").Call(obj)
	return ret
}

func ToolBar_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("ToolBar_SetParentWindow").Call(obj, value)
}

func ToolBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ToolBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAction").Call(obj)
	return ret
}

func ToolBar_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetAction").Call(obj, value)
}

func ToolBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ToolBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ToolBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ToolBar_SetBiDiMode").Call(obj, uintptr(value))
}

func ToolBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ToolBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ToolBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func ToolBar_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetClientHeight").Call(obj, uintptr(value))
}

func ToolBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ToolBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ToolBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func ToolBar_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetClientWidth").Call(obj, uintptr(value))
}

func ToolBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ToolBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func ToolBar_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ToolBar_SetControlState").Call(obj, uintptr(value))
}

func ToolBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ToolBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ToolBar_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ToolBar_SetControlStyle").Call(obj, uintptr(value))
}

func ToolBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetParent").Call(obj)
	return ret
}

func ToolBar_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetParent").Call(obj, value)
}

func ToolBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetLeft").Call(obj)
	return int32(ret)
}

func ToolBar_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetLeft").Call(obj, uintptr(value))
}

func ToolBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetTop").Call(obj)
	return int32(ret)
}

func ToolBar_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetTop").Call(obj, uintptr(value))
}

func ToolBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetWidth").Call(obj)
	return int32(ret)
}

func ToolBar_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetWidth").Call(obj, uintptr(value))
}

func ToolBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToolBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func ToolBar_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ToolBar_SetCursor").Call(obj, uintptr(value))
}

func ToolBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ToolBar_SetHint").Call(obj, GoStrToDStr(value))
}

func ToolBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func ToolBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ToolBar_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ToolBar_SetComponentIndex").Call(obj, uintptr(value))
}

func ToolBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetOwner").Call(obj)
	return ret
}

func ToolBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ToolBar_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ToolBar_SetName").Call(obj, GoStrToDStr(value))
}

func ToolBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ToolBar_GetTag").Call(obj)
	return int(ret)
}

func ToolBar_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ToolBar_SetTag").Call(obj, uintptr(value))
}

func ToolBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func ToolBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetAnchorSideLeft").Call(obj, value)
}

func ToolBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAnchorSideTop").Call(obj)
	return ret
}

func ToolBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetAnchorSideTop").Call(obj, value)
}

func ToolBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAnchorSideRight").Call(obj)
	return ret
}

func ToolBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetAnchorSideRight").Call(obj, value)
}

func ToolBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func ToolBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetAnchorSideBottom").Call(obj, value)
}

func ToolBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetChildSizing").Call(obj)
	return ret
}

func ToolBar_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetChildSizing").Call(obj, value)
}

func ToolBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetBorderSpacing").Call(obj)
	return ret
}

func ToolBar_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ToolBar_SetBorderSpacing").Call(obj, value)
}

func ToolBar_GetButtons(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetButtons").Call(obj, uintptr(Index))
	return ret
}

func ToolBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ToolBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ToolBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ToolBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ToolBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ToolBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("ToolBar_StaticClassType").Call()
	return TClass(r)
}
