package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TFlowPanel ---------------------------

func FlowPanel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_Create").Call(obj)
	return ret
}

func FlowPanel_Free(obj uintptr) {
	getLazyProc("FlowPanel_Free").Call(obj)
}

func FlowPanel_GetControlIndex(obj uintptr, AControl uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetControlIndex").Call(obj, AControl)
	return int32(ret)
}

func FlowPanel_SetControlIndex(obj uintptr, AControl uintptr, Index int32) {
	getLazyProc("FlowPanel_SetControlIndex").Call(obj, AControl, uintptr(Index))
}

func FlowPanel_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func FlowPanel_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func FlowPanel_DisableAlign(obj uintptr) {
	getLazyProc("FlowPanel_DisableAlign").Call(obj)
}

func FlowPanel_EnableAlign(obj uintptr) {
	getLazyProc("FlowPanel_EnableAlign").Call(obj)
}

func FlowPanel_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func FlowPanel_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("FlowPanel_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func FlowPanel_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("FlowPanel_InsertControl").Call(obj, AControl)
}

func FlowPanel_Invalidate(obj uintptr) {
	getLazyProc("FlowPanel_Invalidate").Call(obj)
}

func FlowPanel_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("FlowPanel_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func FlowPanel_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("FlowPanel_RemoveControl").Call(obj, AControl)
}

func FlowPanel_Realign(obj uintptr) {
	getLazyProc("FlowPanel_Realign").Call(obj)
}

func FlowPanel_Repaint(obj uintptr) {
	getLazyProc("FlowPanel_Repaint").Call(obj)
}

func FlowPanel_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("FlowPanel_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func FlowPanel_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("FlowPanel_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func FlowPanel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("FlowPanel_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func FlowPanel_SetFocus(obj uintptr) {
	getLazyProc("FlowPanel_SetFocus").Call(obj)
}

func FlowPanel_Update(obj uintptr) {
	getLazyProc("FlowPanel_Update").Call(obj)
}

func FlowPanel_BringToFront(obj uintptr) {
	getLazyProc("FlowPanel_BringToFront").Call(obj)
}

func FlowPanel_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("FlowPanel_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("FlowPanel_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_Hide(obj uintptr) {
	getLazyProc("FlowPanel_Hide").Call(obj)
}

func FlowPanel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("FlowPanel_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func FlowPanel_Refresh(obj uintptr) {
	getLazyProc("FlowPanel_Refresh").Call(obj)
}

func FlowPanel_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("FlowPanel_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("FlowPanel_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_SendToBack(obj uintptr) {
	getLazyProc("FlowPanel_SendToBack").Call(obj)
}

func FlowPanel_Show(obj uintptr) {
	getLazyProc("FlowPanel_Show").Call(obj)
}

func FlowPanel_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("FlowPanel_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func FlowPanel_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetTextLen").Call(obj)
	return int32(ret)
}

func FlowPanel_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("FlowPanel_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func FlowPanel_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func FlowPanel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_Assign(obj uintptr, Source uintptr) {
	getLazyProc("FlowPanel_Assign").Call(obj, Source)
}

func FlowPanel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("FlowPanel_ClassType").Call(obj)
	return TClass(ret)
}

func FlowPanel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_InstanceSize").Call(obj)
	return int32(ret)
}

func FlowPanel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("FlowPanel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func FlowPanel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetHashCode").Call(obj)
	return int32(ret)
}

func FlowPanel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("FlowPanel_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func FlowPanel_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("FlowPanel_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func FlowPanel_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("FlowPanel_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func FlowPanel_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("FlowPanel_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func FlowPanel_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("FlowPanel_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func FlowPanel_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("FlowPanel_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func FlowPanel_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("FlowPanel_AnchorClient").Call(obj, uintptr(ASpace))
}

func FlowPanel_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FlowPanel_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FlowPanel_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("FlowPanel_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func FlowPanel_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("FlowPanel_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func FlowPanel_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("FlowPanel_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func FlowPanel_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("FlowPanel_GetAlign").Call(obj)
	return TAlign(ret)
}

func FlowPanel_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("FlowPanel_SetAlign").Call(obj, uintptr(value))
}

func FlowPanel_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("FlowPanel_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func FlowPanel_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("FlowPanel_SetAlignment").Call(obj, uintptr(value))
}

func FlowPanel_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func FlowPanel_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("FlowPanel_SetAnchors").Call(obj, uintptr(value))
}

func FlowPanel_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetAutoWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetAutoWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetAutoWrap(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetAutoWrap").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("FlowPanel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func FlowPanel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("FlowPanel_SetBiDiMode").Call(obj, uintptr(value))
}

func FlowPanel_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetBorderWidth").Call(obj)
	return int32(ret)
}

func FlowPanel_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetBorderWidth").Call(obj, uintptr(value))
}

func FlowPanel_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("FlowPanel_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func FlowPanel_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("FlowPanel_SetBorderStyle").Call(obj, uintptr(value))
}

func FlowPanel_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_SetCaption(obj uintptr, value string) {
	getLazyProc("FlowPanel_SetCaption").Call(obj, GoStrToDStr(value))
}

func FlowPanel_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("FlowPanel_GetColor").Call(obj)
	return TColor(ret)
}

func FlowPanel_SetColor(obj uintptr, value TColor) {
	getLazyProc("FlowPanel_SetColor").Call(obj, uintptr(value))
}

func FlowPanel_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetConstraints").Call(obj)
	return ret
}

func FlowPanel_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetConstraints").Call(obj, value)
}

func FlowPanel_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetDockSite(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("FlowPanel_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func FlowPanel_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("FlowPanel_SetDragCursor").Call(obj, uintptr(value))
}

func FlowPanel_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("FlowPanel_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func FlowPanel_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("FlowPanel_SetDragKind").Call(obj, uintptr(value))
}

func FlowPanel_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("FlowPanel_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func FlowPanel_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("FlowPanel_SetDragMode").Call(obj, uintptr(value))
}

func FlowPanel_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetEnabled(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetFlowStyle(obj uintptr) TFlowStyle {
	ret, _, _ := getLazyProc("FlowPanel_GetFlowStyle").Call(obj)
	return TFlowStyle(ret)
}

func FlowPanel_SetFlowStyle(obj uintptr, value TFlowStyle) {
	getLazyProc("FlowPanel_SetFlowStyle").Call(obj, uintptr(value))
}

func FlowPanel_GetFullRepaint(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetFullRepaint").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetFullRepaint(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetFullRepaint").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetFont").Call(obj)
	return ret
}

func FlowPanel_SetFont(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetFont").Call(obj, value)
}

func FlowPanel_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetParentBackground(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetParentColor(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetParentFont(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetPopupMenu").Call(obj)
	return ret
}

func FlowPanel_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetPopupMenu").Call(obj, value)
}

func FlowPanel_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetShowHint(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("FlowPanel_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func FlowPanel_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("FlowPanel_SetTabOrder").Call(obj, uintptr(value))
}

func FlowPanel_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetTabStop(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_SetVisible(obj uintptr, value bool) {
	getLazyProc("FlowPanel_SetVisible").Call(obj, GoBoolToDBool(value))
}

func FlowPanel_SetOnAlignPosition(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnConstrainedResize(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnConstrainedResize").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnDockDrop(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnGetSiteInfo(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnResize(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_SetOnUnDock(obj uintptr, fn interface{}) {
	getLazyProc("FlowPanel_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func FlowPanel_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetDockClientCount").Call(obj)
	return int32(ret)
}

func FlowPanel_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func FlowPanel_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetBrush").Call(obj)
	return ret
}

func FlowPanel_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetControlCount").Call(obj)
	return int32(ret)
}

func FlowPanel_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FlowPanel_GetHandle").Call(obj)
	return HWND(ret)
}

func FlowPanel_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FlowPanel_GetParentWindow").Call(obj)
	return HWND(ret)
}

func FlowPanel_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("FlowPanel_SetParentWindow").Call(obj, uintptr(value))
}

func FlowPanel_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAction").Call(obj)
	return ret
}

func FlowPanel_SetAction(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetAction").Call(obj, value)
}

func FlowPanel_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("FlowPanel_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("FlowPanel_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FlowPanel_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetClientHeight").Call(obj)
	return int32(ret)
}

func FlowPanel_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetClientHeight").Call(obj, uintptr(value))
}

func FlowPanel_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("FlowPanel_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("FlowPanel_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FlowPanel_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetClientWidth").Call(obj)
	return int32(ret)
}

func FlowPanel_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetClientWidth").Call(obj, uintptr(value))
}

func FlowPanel_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("FlowPanel_GetControlState").Call(obj)
	return TControlState(ret)
}

func FlowPanel_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("FlowPanel_SetControlState").Call(obj, uintptr(value))
}

func FlowPanel_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("FlowPanel_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func FlowPanel_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("FlowPanel_SetControlStyle").Call(obj, uintptr(value))
}

func FlowPanel_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("FlowPanel_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func FlowPanel_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetParent").Call(obj)
	return ret
}

func FlowPanel_SetParent(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetParent").Call(obj, value)
}

func FlowPanel_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetLeft").Call(obj)
	return int32(ret)
}

func FlowPanel_SetLeft(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetLeft").Call(obj, uintptr(value))
}

func FlowPanel_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetTop").Call(obj)
	return int32(ret)
}

func FlowPanel_SetTop(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetTop").Call(obj, uintptr(value))
}

func FlowPanel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetWidth").Call(obj)
	return int32(ret)
}

func FlowPanel_SetWidth(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetWidth").Call(obj, uintptr(value))
}

func FlowPanel_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetHeight").Call(obj)
	return int32(ret)
}

func FlowPanel_SetHeight(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetHeight").Call(obj, uintptr(value))
}

func FlowPanel_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("FlowPanel_GetCursor").Call(obj)
	return TCursor(ret)
}

func FlowPanel_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("FlowPanel_SetCursor").Call(obj, uintptr(value))
}

func FlowPanel_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_SetHint(obj uintptr, value string) {
	getLazyProc("FlowPanel_SetHint").Call(obj, GoStrToDStr(value))
}

func FlowPanel_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetComponentCount").Call(obj)
	return int32(ret)
}

func FlowPanel_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FlowPanel_GetComponentIndex").Call(obj)
	return int32(ret)
}

func FlowPanel_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("FlowPanel_SetComponentIndex").Call(obj, uintptr(value))
}

func FlowPanel_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetOwner").Call(obj)
	return ret
}

func FlowPanel_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("FlowPanel_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func FlowPanel_SetName(obj uintptr, value string) {
	getLazyProc("FlowPanel_SetName").Call(obj, GoStrToDStr(value))
}

func FlowPanel_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("FlowPanel_GetTag").Call(obj)
	return int(ret)
}

func FlowPanel_SetTag(obj uintptr, value int) {
	getLazyProc("FlowPanel_SetTag").Call(obj, uintptr(value))
}

func FlowPanel_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchorSideLeft").Call(obj)
	return ret
}

func FlowPanel_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetAnchorSideLeft").Call(obj, value)
}

func FlowPanel_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchorSideTop").Call(obj)
	return ret
}

func FlowPanel_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetAnchorSideTop").Call(obj, value)
}

func FlowPanel_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchorSideRight").Call(obj)
	return ret
}

func FlowPanel_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetAnchorSideRight").Call(obj, value)
}

func FlowPanel_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchorSideBottom").Call(obj)
	return ret
}

func FlowPanel_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetAnchorSideBottom").Call(obj, value)
}

func FlowPanel_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetChildSizing").Call(obj)
	return ret
}

func FlowPanel_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetChildSizing").Call(obj, value)
}

func FlowPanel_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetBorderSpacing").Call(obj)
	return ret
}

func FlowPanel_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("FlowPanel_SetBorderSpacing").Call(obj, value)
}

func FlowPanel_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func FlowPanel_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetControls").Call(obj, uintptr(Index))
	return ret
}

func FlowPanel_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func FlowPanel_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("FlowPanel_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func FlowPanel_StaticClassType() TClass {
	r, _, _ := getLazyProc("FlowPanel_StaticClassType").Call()
	return TClass(r)
}
