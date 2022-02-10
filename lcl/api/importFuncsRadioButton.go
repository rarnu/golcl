package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TRadioButton ---------------------------

func RadioButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_Create").Call(obj)
	return ret
}

func RadioButton_Free(obj uintptr) {
	getLazyProc("RadioButton_Free").Call(obj)
}

func RadioButton_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func RadioButton_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("RadioButton_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func RadioButton_DisableAlign(obj uintptr) {
	getLazyProc("RadioButton_DisableAlign").Call(obj)
}

func RadioButton_EnableAlign(obj uintptr) {
	getLazyProc("RadioButton_EnableAlign").Call(obj)
}

func RadioButton_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("RadioButton_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func RadioButton_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("RadioButton_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func RadioButton_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("RadioButton_InsertControl").Call(obj, AControl)
}

func RadioButton_Invalidate(obj uintptr) {
	getLazyProc("RadioButton_Invalidate").Call(obj)
}

func RadioButton_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("RadioButton_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func RadioButton_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("RadioButton_RemoveControl").Call(obj, AControl)
}

func RadioButton_Realign(obj uintptr) {
	getLazyProc("RadioButton_Realign").Call(obj)
}

func RadioButton_Repaint(obj uintptr) {
	getLazyProc("RadioButton_Repaint").Call(obj)
}

func RadioButton_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("RadioButton_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func RadioButton_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("RadioButton_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func RadioButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("RadioButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func RadioButton_SetFocus(obj uintptr) {
	getLazyProc("RadioButton_SetFocus").Call(obj)
}

func RadioButton_Update(obj uintptr) {
	getLazyProc("RadioButton_Update").Call(obj)
}

func RadioButton_BringToFront(obj uintptr) {
	getLazyProc("RadioButton_BringToFront").Call(obj)
}

func RadioButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("RadioButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("RadioButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_Hide(obj uintptr) {
	getLazyProc("RadioButton_Hide").Call(obj)
}

func RadioButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("RadioButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func RadioButton_Refresh(obj uintptr) {
	getLazyProc("RadioButton_Refresh").Call(obj)
}

func RadioButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("RadioButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("RadioButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_SendToBack(obj uintptr) {
	getLazyProc("RadioButton_SendToBack").Call(obj)
}

func RadioButton_Show(obj uintptr) {
	getLazyProc("RadioButton_Show").Call(obj)
}

func RadioButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("RadioButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func RadioButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetTextLen").Call(obj)
	return int32(ret)
}

func RadioButton_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("RadioButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func RadioButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("RadioButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func RadioButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_Assign(obj uintptr, Source uintptr) {
	getLazyProc("RadioButton_Assign").Call(obj, Source)
}

func RadioButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("RadioButton_ClassType").Call(obj)
	return TClass(ret)
}

func RadioButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_InstanceSize").Call(obj)
	return int32(ret)
}

func RadioButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("RadioButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func RadioButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func RadioButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetHashCode").Call(obj)
	return int32(ret)
}

func RadioButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("RadioButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RadioButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("RadioButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RadioButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("RadioButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func RadioButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("RadioButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func RadioButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("RadioButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func RadioButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("RadioButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func RadioButton_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("RadioButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func RadioButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("RadioButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func RadioButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("RadioButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func RadioButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("RadioButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func RadioButton_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAction").Call(obj)
	return ret
}

func RadioButton_SetAction(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetAction").Call(obj, value)
}

func RadioButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("RadioButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func RadioButton_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("RadioButton_SetAlign").Call(obj, uintptr(value))
}

func RadioButton_GetAlignment(obj uintptr) TLeftRight {
	ret, _, _ := getLazyProc("RadioButton_GetAlignment").Call(obj)
	return TLeftRight(ret)
}

func RadioButton_SetAlignment(obj uintptr, value TLeftRight) {
	getLazyProc("RadioButton_SetAlignment").Call(obj, uintptr(value))
}

func RadioButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("RadioButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func RadioButton_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("RadioButton_SetAnchors").Call(obj, uintptr(value))
}

func RadioButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("RadioButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func RadioButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("RadioButton_SetBiDiMode").Call(obj, uintptr(value))
}

func RadioButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_SetCaption(obj uintptr, value string) {
	getLazyProc("RadioButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func RadioButton_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetChecked(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetChecked").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("RadioButton_GetColor").Call(obj)
	return TColor(ret)
}

func RadioButton_SetColor(obj uintptr, value TColor) {
	getLazyProc("RadioButton_SetColor").Call(obj, uintptr(value))
}

func RadioButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetConstraints").Call(obj)
	return ret
}

func RadioButton_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetConstraints").Call(obj, value)
}

func RadioButton_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RadioButton_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func RadioButton_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("RadioButton_SetDragCursor").Call(obj, uintptr(value))
}

func RadioButton_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("RadioButton_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func RadioButton_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("RadioButton_SetDragKind").Call(obj, uintptr(value))
}

func RadioButton_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("RadioButton_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func RadioButton_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("RadioButton_SetDragMode").Call(obj, uintptr(value))
}

func RadioButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetEnabled(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetFont").Call(obj)
	return ret
}

func RadioButton_SetFont(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetFont").Call(obj, value)
}

func RadioButton_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetParentColor(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetParentFont(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetPopupMenu").Call(obj)
	return ret
}

func RadioButton_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetPopupMenu").Call(obj, value)
}

func RadioButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetShowHint(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("RadioButton_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func RadioButton_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("RadioButton_SetTabOrder").Call(obj, uintptr(value))
}

func RadioButton_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetTabStop(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetVisible(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func RadioButton_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("RadioButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func RadioButton_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetDockClientCount").Call(obj)
	return int32(ret)
}

func RadioButton_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetDockSite(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func RadioButton_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetBrush").Call(obj)
	return ret
}

func RadioButton_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetControlCount").Call(obj)
	return int32(ret)
}

func RadioButton_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RadioButton_GetHandle").Call(obj)
	return HWND(ret)
}

func RadioButton_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RadioButton_GetParentWindow").Call(obj)
	return HWND(ret)
}

func RadioButton_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("RadioButton_SetParentWindow").Call(obj, uintptr(value))
}

func RadioButton_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("RadioButton_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("RadioButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("RadioButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func RadioButton_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetClientHeight").Call(obj, uintptr(value))
}

func RadioButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("RadioButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("RadioButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func RadioButton_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetClientWidth").Call(obj, uintptr(value))
}

func RadioButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("RadioButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func RadioButton_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("RadioButton_SetControlState").Call(obj, uintptr(value))
}

func RadioButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("RadioButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func RadioButton_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("RadioButton_SetControlStyle").Call(obj, uintptr(value))
}

func RadioButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetParent").Call(obj)
	return ret
}

func RadioButton_SetParent(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetParent").Call(obj, value)
}

func RadioButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetLeft").Call(obj)
	return int32(ret)
}

func RadioButton_SetLeft(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetLeft").Call(obj, uintptr(value))
}

func RadioButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetTop").Call(obj)
	return int32(ret)
}

func RadioButton_SetTop(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetTop").Call(obj, uintptr(value))
}

func RadioButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetWidth").Call(obj)
	return int32(ret)
}

func RadioButton_SetWidth(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetWidth").Call(obj, uintptr(value))
}

func RadioButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetHeight").Call(obj)
	return int32(ret)
}

func RadioButton_SetHeight(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetHeight").Call(obj, uintptr(value))
}

func RadioButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RadioButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func RadioButton_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("RadioButton_SetCursor").Call(obj, uintptr(value))
}

func RadioButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_SetHint(obj uintptr, value string) {
	getLazyProc("RadioButton_SetHint").Call(obj, GoStrToDStr(value))
}

func RadioButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func RadioButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func RadioButton_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("RadioButton_SetComponentIndex").Call(obj, uintptr(value))
}

func RadioButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetOwner").Call(obj)
	return ret
}

func RadioButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func RadioButton_SetName(obj uintptr, value string) {
	getLazyProc("RadioButton_SetName").Call(obj, GoStrToDStr(value))
}

func RadioButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("RadioButton_GetTag").Call(obj)
	return int(ret)
}

func RadioButton_SetTag(obj uintptr, value int) {
	getLazyProc("RadioButton_SetTag").Call(obj, uintptr(value))
}

func RadioButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func RadioButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetAnchorSideLeft").Call(obj, value)
}

func RadioButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAnchorSideTop").Call(obj)
	return ret
}

func RadioButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetAnchorSideTop").Call(obj, value)
}

func RadioButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAnchorSideRight").Call(obj)
	return ret
}

func RadioButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetAnchorSideRight").Call(obj, value)
}

func RadioButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func RadioButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetAnchorSideBottom").Call(obj, value)
}

func RadioButton_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetChildSizing").Call(obj)
	return ret
}

func RadioButton_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetChildSizing").Call(obj, value)
}

func RadioButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetBorderSpacing").Call(obj)
	return ret
}

func RadioButton_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("RadioButton_SetBorderSpacing").Call(obj, value)
}

func RadioButton_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func RadioButton_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetControls").Call(obj, uintptr(Index))
	return ret
}

func RadioButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func RadioButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("RadioButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func RadioButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("RadioButton_StaticClassType").Call()
	return TClass(r)
}
