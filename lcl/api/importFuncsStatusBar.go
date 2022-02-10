package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TStatusBar ---------------------------

func StatusBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_Create").Call(obj)
	return ret
}

func StatusBar_Free(obj uintptr) {
	getLazyProc("StatusBar_Free").Call(obj)
}

func StatusBar_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("StatusBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func StatusBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("StatusBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func StatusBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func StatusBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("StatusBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func StatusBar_DisableAlign(obj uintptr) {
	getLazyProc("StatusBar_DisableAlign").Call(obj)
}

func StatusBar_EnableAlign(obj uintptr) {
	getLazyProc("StatusBar_EnableAlign").Call(obj)
}

func StatusBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("StatusBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func StatusBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("StatusBar_InsertControl").Call(obj, AControl)
}

func StatusBar_Invalidate(obj uintptr) {
	getLazyProc("StatusBar_Invalidate").Call(obj)
}

func StatusBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("StatusBar_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func StatusBar_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("StatusBar_RemoveControl").Call(obj, AControl)
}

func StatusBar_Realign(obj uintptr) {
	getLazyProc("StatusBar_Realign").Call(obj)
}

func StatusBar_Repaint(obj uintptr) {
	getLazyProc("StatusBar_Repaint").Call(obj)
}

func StatusBar_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("StatusBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func StatusBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("StatusBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func StatusBar_SetFocus(obj uintptr) {
	getLazyProc("StatusBar_SetFocus").Call(obj)
}

func StatusBar_Update(obj uintptr) {
	getLazyProc("StatusBar_Update").Call(obj)
}

func StatusBar_BringToFront(obj uintptr) {
	getLazyProc("StatusBar_BringToFront").Call(obj)
}

func StatusBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("StatusBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("StatusBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_Hide(obj uintptr) {
	getLazyProc("StatusBar_Hide").Call(obj)
}

func StatusBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("StatusBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func StatusBar_Refresh(obj uintptr) {
	getLazyProc("StatusBar_Refresh").Call(obj)
}

func StatusBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("StatusBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("StatusBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_SendToBack(obj uintptr) {
	getLazyProc("StatusBar_SendToBack").Call(obj)
}

func StatusBar_Show(obj uintptr) {
	getLazyProc("StatusBar_Show").Call(obj)
}

func StatusBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("StatusBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func StatusBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetTextLen").Call(obj)
	return int32(ret)
}

func StatusBar_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("StatusBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func StatusBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("StatusBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func StatusBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_Assign(obj uintptr, Source uintptr) {
	getLazyProc("StatusBar_Assign").Call(obj, Source)
}

func StatusBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StatusBar_ClassType").Call(obj)
	return TClass(ret)
}

func StatusBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_InstanceSize").Call(obj)
	return int32(ret)
}

func StatusBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StatusBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StatusBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func StatusBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetHashCode").Call(obj)
	return int32(ret)
}

func StatusBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("StatusBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StatusBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("StatusBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func StatusBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("StatusBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func StatusBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("StatusBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func StatusBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("StatusBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func StatusBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("StatusBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func StatusBar_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("StatusBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func StatusBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("StatusBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func StatusBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("StatusBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func StatusBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("StatusBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func StatusBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("StatusBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func StatusBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAction").Call(obj)
	return ret
}

func StatusBar_SetAction(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetAction").Call(obj, value)
}

func StatusBar_GetAutoHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetAutoHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetAutoHint(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetAutoHint").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("StatusBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func StatusBar_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("StatusBar_SetAlign").Call(obj, uintptr(value))
}

func StatusBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("StatusBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func StatusBar_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("StatusBar_SetAnchors").Call(obj, uintptr(value))
}

func StatusBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("StatusBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func StatusBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("StatusBar_SetBiDiMode").Call(obj, uintptr(value))
}

func StatusBar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func StatusBar_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetBorderWidth").Call(obj, uintptr(value))
}

func StatusBar_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("StatusBar_GetColor").Call(obj)
	return TColor(ret)
}

func StatusBar_SetColor(obj uintptr, value TColor) {
	getLazyProc("StatusBar_SetColor").Call(obj, uintptr(value))
}

func StatusBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StatusBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func StatusBar_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("StatusBar_SetDragCursor").Call(obj, uintptr(value))
}

func StatusBar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("StatusBar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func StatusBar_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("StatusBar_SetDragKind").Call(obj, uintptr(value))
}

func StatusBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("StatusBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func StatusBar_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("StatusBar_SetDragMode").Call(obj, uintptr(value))
}

func StatusBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetEnabled(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetFont").Call(obj)
	return ret
}

func StatusBar_SetFont(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetFont").Call(obj, value)
}

func StatusBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetConstraints").Call(obj)
	return ret
}

func StatusBar_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetConstraints").Call(obj, value)
}

func StatusBar_GetPanels(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetPanels").Call(obj)
	return ret
}

func StatusBar_SetPanels(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetPanels").Call(obj, value)
}

func StatusBar_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetParentColor(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetParentFont(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetPopupMenu").Call(obj)
	return ret
}

func StatusBar_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetPopupMenu").Call(obj, value)
}

func StatusBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetShowHint(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimplePanel(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetSimplePanel").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetSimplePanel(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetSimplePanel").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimpleText(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_GetSimpleText").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_SetSimpleText(obj uintptr, value string) {
	getLazyProc("StatusBar_SetSimpleText").Call(obj, GoStrToDStr(value))
}

func StatusBar_GetSizeGrip(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetSizeGrip").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetSizeGrip(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetSizeGrip").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetUseSystemFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetUseSystemFont").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetUseSystemFont(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetUseSystemFont").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetVisible(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func StatusBar_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnHint(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnHint").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnResize(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("StatusBar_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func StatusBar_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetCanvas").Call(obj)
	return ret
}

func StatusBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func StatusBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetDockSite(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func StatusBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetBrush").Call(obj)
	return ret
}

func StatusBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetControlCount").Call(obj)
	return int32(ret)
}

func StatusBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StatusBar_GetHandle").Call(obj)
	return HWND(ret)
}

func StatusBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("StatusBar_GetParentWindow").Call(obj)
	return HWND(ret)
}

func StatusBar_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("StatusBar_SetParentWindow").Call(obj, uintptr(value))
}

func StatusBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("StatusBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func StatusBar_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("StatusBar_SetTabOrder").Call(obj, uintptr(value))
}

func StatusBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetTabStop(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("StatusBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("StatusBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("StatusBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StatusBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func StatusBar_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetClientHeight").Call(obj, uintptr(value))
}

func StatusBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("StatusBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("StatusBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func StatusBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func StatusBar_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetClientWidth").Call(obj, uintptr(value))
}

func StatusBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("StatusBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func StatusBar_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("StatusBar_SetControlState").Call(obj, uintptr(value))
}

func StatusBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("StatusBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func StatusBar_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("StatusBar_SetControlStyle").Call(obj, uintptr(value))
}

func StatusBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("StatusBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func StatusBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetParent").Call(obj)
	return ret
}

func StatusBar_SetParent(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetParent").Call(obj, value)
}

func StatusBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetLeft").Call(obj)
	return int32(ret)
}

func StatusBar_SetLeft(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetLeft").Call(obj, uintptr(value))
}

func StatusBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetTop").Call(obj)
	return int32(ret)
}

func StatusBar_SetTop(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetTop").Call(obj, uintptr(value))
}

func StatusBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetWidth").Call(obj)
	return int32(ret)
}

func StatusBar_SetWidth(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetWidth").Call(obj, uintptr(value))
}

func StatusBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetHeight").Call(obj)
	return int32(ret)
}

func StatusBar_SetHeight(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetHeight").Call(obj, uintptr(value))
}

func StatusBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("StatusBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func StatusBar_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("StatusBar_SetCursor").Call(obj, uintptr(value))
}

func StatusBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_SetHint(obj uintptr, value string) {
	getLazyProc("StatusBar_SetHint").Call(obj, GoStrToDStr(value))
}

func StatusBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func StatusBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StatusBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func StatusBar_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("StatusBar_SetComponentIndex").Call(obj, uintptr(value))
}

func StatusBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetOwner").Call(obj)
	return ret
}

func StatusBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("StatusBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func StatusBar_SetName(obj uintptr, value string) {
	getLazyProc("StatusBar_SetName").Call(obj, GoStrToDStr(value))
}

func StatusBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("StatusBar_GetTag").Call(obj)
	return int(ret)
}

func StatusBar_SetTag(obj uintptr, value int) {
	getLazyProc("StatusBar_SetTag").Call(obj, uintptr(value))
}

func StatusBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func StatusBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetAnchorSideLeft").Call(obj, value)
}

func StatusBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAnchorSideTop").Call(obj)
	return ret
}

func StatusBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetAnchorSideTop").Call(obj, value)
}

func StatusBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAnchorSideRight").Call(obj)
	return ret
}

func StatusBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetAnchorSideRight").Call(obj, value)
}

func StatusBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func StatusBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetAnchorSideBottom").Call(obj, value)
}

func StatusBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetChildSizing").Call(obj)
	return ret
}

func StatusBar_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetChildSizing").Call(obj, value)
}

func StatusBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetBorderSpacing").Call(obj)
	return ret
}

func StatusBar_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("StatusBar_SetBorderSpacing").Call(obj, value)
}

func StatusBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func StatusBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func StatusBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func StatusBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("StatusBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func StatusBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("StatusBar_StaticClassType").Call()
	return TClass(r)
}
