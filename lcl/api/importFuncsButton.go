package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TButton ---------------------------

func Button_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_Create").Call(obj)
	return ret
}

func Button_Free(obj uintptr) {
	_, _, _ = getLazyProc("Button_Free").Call(obj)
}

func Button_Click(obj uintptr) {
	_, _, _ = getLazyProc("Button_Click").Call(obj)
}

func Button_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Button_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Button_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Button_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Button_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Button_DisableAlign").Call(obj)
}

func Button_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Button_EnableAlign").Call(obj)
}

func Button_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Button_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Button_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("Button_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Button_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Button_InsertControl").Call(obj, AControl)
}

func Button_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Button_Invalidate").Call(obj)
}

func Button_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("Button_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func Button_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Button_RemoveControl").Call(obj, AControl)
}

func Button_Realign(obj uintptr) {
	_, _, _ = getLazyProc("Button_Realign").Call(obj)
}

func Button_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Button_Repaint").Call(obj)
}

func Button_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("Button_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Button_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("Button_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Button_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Button_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Button_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("Button_SetFocus").Call(obj)
}

func Button_Update(obj uintptr) {
	_, _, _ = getLazyProc("Button_Update").Call(obj)
}

func Button_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Button_BringToFront").Call(obj)
}

func Button_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Button_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Button_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Button_Hide").Call(obj)
}

func Button_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Button_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Button_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Button_Refresh").Call(obj)
}

func Button_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Button_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Button_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Button_SendToBack").Call(obj)
}

func Button_Show(obj uintptr) {
	_, _, _ = getLazyProc("Button_Show").Call(obj)
}

func Button_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Button_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Button_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetTextLen").Call(obj)
	return int32(ret)
}

func Button_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Button_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Button_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Button_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Button_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Button_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Button_Assign").Call(obj, Source)
}

func Button_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Button_ClassType").Call(obj)
	return TClass(ret)
}

func Button_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Button_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_InstanceSize").Call(obj)
	return int32(ret)
}

func Button_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Button_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Button_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Button_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetHashCode").Call(obj)
	return int32(ret)
}

func Button_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Button_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Button_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Button_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Button_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Button_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Button_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Button_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Button_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Button_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Button_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Button_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Button_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Button_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Button_AnchorClient").Call(obj, uintptr(ASpace))
}

func Button_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Button_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Button_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Button_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Button_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Button_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Button_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Button_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Button_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetAction").Call(obj)
	return ret
}

func Button_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetAction").Call(obj, value)
}

func Button_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Button_GetAlign").Call(obj)
	return TAlign(ret)
}

func Button_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Button_SetAlign").Call(obj, uintptr(value))
}

func Button_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Button_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Button_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Button_SetAnchors").Call(obj, uintptr(value))
}

func Button_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Button_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Button_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Button_SetBiDiMode").Call(obj, uintptr(value))
}

func Button_GetCancel(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetCancel").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetCancel(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetCancel").Call(obj, GoBoolToDBool(value))
}

func Button_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func Button_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("Button_SetCaption").Call(obj, GoStrToDStr(value))
}

func Button_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetConstraints").Call(obj)
	return ret
}

func Button_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetConstraints").Call(obj, value)
}

func Button_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetDefault(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetDefault").Call(obj, GoBoolToDBool(value))
}

func Button_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Button_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Button_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Button_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Button_SetDragCursor").Call(obj, uintptr(value))
}

func Button_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Button_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Button_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("Button_SetDragKind").Call(obj, uintptr(value))
}

func Button_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Button_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Button_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Button_SetDragMode").Call(obj, uintptr(value))
}

func Button_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Button_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetFont").Call(obj)
	return ret
}

func Button_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetFont").Call(obj, value)
}

func Button_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("Button_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func Button_SetModalResult(obj uintptr, value TModalResult) {
	_, _, _ = getLazyProc("Button_SetModalResult").Call(obj, uintptr(value))
}

func Button_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Button_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Button_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Button_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetPopupMenu").Call(obj)
	return ret
}

func Button_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetPopupMenu").Call(obj, value)
}

func Button_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Button_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Button_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Button_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("Button_SetTabOrder").Call(obj, uintptr(value))
}

func Button_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Button_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Button_SetOnClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnContextPopup(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnDragDrop(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnDragOver(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnEndDrag(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnExit(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnKeyDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnKeyPress(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnKeyUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnMouseDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnMouseEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnMouseLeave(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnMouseMove(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Button_SetOnMouseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("Button_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Button_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Button_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Button_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Button_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetBrush").Call(obj)
	return ret
}

func Button_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetControlCount").Call(obj)
	return int32(ret)
}

func Button_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Button_GetHandle").Call(obj)
	return ret
}

func Button_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Button_GetParentWindow").Call(obj)
	return ret
}

func Button_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Button_SetParentWindow").Call(obj, value)
}

func Button_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Button_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Button_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Button_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Button_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Button_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetClientHeight").Call(obj)
	return int32(ret)
}

func Button_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetClientHeight").Call(obj, uintptr(value))
}

func Button_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Button_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Button_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Button_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetClientWidth").Call(obj)
	return int32(ret)
}

func Button_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetClientWidth").Call(obj, uintptr(value))
}

func Button_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Button_GetControlState").Call(obj)
	return TControlState(ret)
}

func Button_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Button_SetControlState").Call(obj, uintptr(value))
}

func Button_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Button_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Button_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Button_SetControlStyle").Call(obj, uintptr(value))
}

func Button_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Button_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Button_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetParent").Call(obj)
	return ret
}

func Button_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetParent").Call(obj, value)
}

func Button_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetLeft").Call(obj)
	return int32(ret)
}

func Button_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetLeft").Call(obj, uintptr(value))
}

func Button_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetTop").Call(obj)
	return int32(ret)
}

func Button_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetTop").Call(obj, uintptr(value))
}

func Button_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetWidth").Call(obj)
	return int32(ret)
}

func Button_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetWidth").Call(obj, uintptr(value))
}

func Button_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetHeight").Call(obj)
	return int32(ret)
}

func Button_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetHeight").Call(obj, uintptr(value))
}

func Button_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Button_GetCursor").Call(obj)
	return TCursor(ret)
}

func Button_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Button_SetCursor").Call(obj, uintptr(value))
}

func Button_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Button_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Button_SetHint").Call(obj, GoStrToDStr(value))
}

func Button_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetComponentCount").Call(obj)
	return int32(ret)
}

func Button_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Button_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Button_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Button_SetComponentIndex").Call(obj, uintptr(value))
}

func Button_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetOwner").Call(obj)
	return ret
}

func Button_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Button_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Button_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Button_SetName").Call(obj, GoStrToDStr(value))
}

func Button_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Button_GetTag").Call(obj)
	return int(ret)
}

func Button_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Button_SetTag").Call(obj, uintptr(value))
}

func Button_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetAnchorSideLeft").Call(obj)
	return ret
}

func Button_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetAnchorSideLeft").Call(obj, value)
}

func Button_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetAnchorSideTop").Call(obj)
	return ret
}

func Button_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetAnchorSideTop").Call(obj, value)
}

func Button_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetAnchorSideRight").Call(obj)
	return ret
}

func Button_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetAnchorSideRight").Call(obj, value)
}

func Button_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetAnchorSideBottom").Call(obj)
	return ret
}

func Button_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetAnchorSideBottom").Call(obj, value)
}

func Button_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetChildSizing").Call(obj)
	return ret
}

func Button_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetChildSizing").Call(obj, value)
}

func Button_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Button_GetBorderSpacing").Call(obj)
	return ret
}

func Button_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Button_SetBorderSpacing").Call(obj, value)
}

func Button_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Button_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Button_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Button_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Button_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Button_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Button_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Button_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Button_StaticClassType() TClass {
	r, _, _ := getLazyProc("Button_StaticClassType").Call()
	return TClass(r)
}
