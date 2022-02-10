package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

//--------------------------- TMonthCalendar ---------------------------

func MonthCalendar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_Create").Call(obj)
	return ret
}

func MonthCalendar_Free(obj uintptr) {
	getLazyProc("MonthCalendar_Free").Call(obj)
}

func MonthCalendar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func MonthCalendar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func MonthCalendar_DisableAlign(obj uintptr) {
	getLazyProc("MonthCalendar_DisableAlign").Call(obj)
}

func MonthCalendar_EnableAlign(obj uintptr) {
	getLazyProc("MonthCalendar_EnableAlign").Call(obj)
}

func MonthCalendar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func MonthCalendar_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("MonthCalendar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func MonthCalendar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("MonthCalendar_InsertControl").Call(obj, AControl)
}

func MonthCalendar_Invalidate(obj uintptr) {
	getLazyProc("MonthCalendar_Invalidate").Call(obj)
}

func MonthCalendar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("MonthCalendar_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func MonthCalendar_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("MonthCalendar_RemoveControl").Call(obj, AControl)
}

func MonthCalendar_Realign(obj uintptr) {
	getLazyProc("MonthCalendar_Realign").Call(obj)
}

func MonthCalendar_Repaint(obj uintptr) {
	getLazyProc("MonthCalendar_Repaint").Call(obj)
}

func MonthCalendar_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("MonthCalendar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func MonthCalendar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("MonthCalendar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func MonthCalendar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("MonthCalendar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func MonthCalendar_SetFocus(obj uintptr) {
	getLazyProc("MonthCalendar_SetFocus").Call(obj)
}

func MonthCalendar_Update(obj uintptr) {
	getLazyProc("MonthCalendar_Update").Call(obj)
}

func MonthCalendar_BringToFront(obj uintptr) {
	getLazyProc("MonthCalendar_BringToFront").Call(obj)
}

func MonthCalendar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("MonthCalendar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("MonthCalendar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_Hide(obj uintptr) {
	getLazyProc("MonthCalendar_Hide").Call(obj)
}

func MonthCalendar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("MonthCalendar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func MonthCalendar_Refresh(obj uintptr) {
	getLazyProc("MonthCalendar_Refresh").Call(obj)
}

func MonthCalendar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("MonthCalendar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("MonthCalendar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_SendToBack(obj uintptr) {
	getLazyProc("MonthCalendar_SendToBack").Call(obj)
}

func MonthCalendar_Show(obj uintptr) {
	getLazyProc("MonthCalendar_Show").Call(obj)
}

func MonthCalendar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("MonthCalendar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func MonthCalendar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetTextLen").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("MonthCalendar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func MonthCalendar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func MonthCalendar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("MonthCalendar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func MonthCalendar_Assign(obj uintptr, Source uintptr) {
	getLazyProc("MonthCalendar_Assign").Call(obj, Source)
}

func MonthCalendar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MonthCalendar_ClassType").Call(obj)
	return TClass(ret)
}

func MonthCalendar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MonthCalendar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MonthCalendar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_InstanceSize").Call(obj)
	return int32(ret)
}

func MonthCalendar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MonthCalendar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MonthCalendar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetHashCode").Call(obj)
	return int32(ret)
}

func MonthCalendar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MonthCalendar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MonthCalendar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("MonthCalendar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MonthCalendar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("MonthCalendar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MonthCalendar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("MonthCalendar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func MonthCalendar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("MonthCalendar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func MonthCalendar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("MonthCalendar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func MonthCalendar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("MonthCalendar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func MonthCalendar_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("MonthCalendar_AnchorClient").Call(obj, uintptr(ASpace))
}

func MonthCalendar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MonthCalendar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("MonthCalendar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func MonthCalendar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("MonthCalendar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func MonthCalendar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("MonthCalendar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func MonthCalendar_GetDateTime(obj uintptr) time.Time {
	var ret int64
	getLazyProc("MonthCalendar_GetDateTime").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(int64(ret), 0)
}

func MonthCalendar_SetDateTime(obj uintptr, value time.Time) {
	tVal := value.Unix()
	getLazyProc("MonthCalendar_SetDateTime").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func MonthCalendar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("MonthCalendar_GetAlign").Call(obj)
	return TAlign(ret)
}

func MonthCalendar_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("MonthCalendar_SetAlign").Call(obj, uintptr(value))
}

func MonthCalendar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func MonthCalendar_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("MonthCalendar_SetAnchors").Call(obj, uintptr(value))
}

func MonthCalendar_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetBorderWidth").Call(obj, uintptr(value))
}

func MonthCalendar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("MonthCalendar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func MonthCalendar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("MonthCalendar_SetBiDiMode").Call(obj, uintptr(value))
}

func MonthCalendar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetConstraints").Call(obj)
	return ret
}

func MonthCalendar_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetConstraints").Call(obj, value)
}

func MonthCalendar_GetDate(obj uintptr) time.Time {
	var ret int64
	getLazyProc("MonthCalendar_GetDate").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(int64(ret), 0)
}

func MonthCalendar_SetDate(obj uintptr, value time.Time) {
	tVal := value.Unix()
	getLazyProc("MonthCalendar_SetDate").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func MonthCalendar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("MonthCalendar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func MonthCalendar_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("MonthCalendar_SetDragCursor").Call(obj, uintptr(value))
}

func MonthCalendar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("MonthCalendar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func MonthCalendar_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("MonthCalendar_SetDragKind").Call(obj, uintptr(value))
}

func MonthCalendar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("MonthCalendar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func MonthCalendar_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("MonthCalendar_SetDragMode").Call(obj, uintptr(value))
}

func MonthCalendar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetEnabled(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetFont").Call(obj)
	return ret
}

func MonthCalendar_SetFont(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetFont").Call(obj, value)
}

func MonthCalendar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetPopupMenu").Call(obj)
	return ret
}

func MonthCalendar_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetPopupMenu").Call(obj, value)
}

func MonthCalendar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetShowHint(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("MonthCalendar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func MonthCalendar_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("MonthCalendar_SetTabOrder").Call(obj, uintptr(value))
}

func MonthCalendar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetTabStop(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetVisible(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("MonthCalendar_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func MonthCalendar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func MonthCalendar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetDockSite(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func MonthCalendar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetBrush").Call(obj)
	return ret
}

func MonthCalendar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetControlCount").Call(obj)
	return int32(ret)
}

func MonthCalendar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MonthCalendar_GetHandle").Call(obj)
	return HWND(ret)
}

func MonthCalendar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MonthCalendar_GetParentWindow").Call(obj)
	return HWND(ret)
}

func MonthCalendar_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("MonthCalendar_SetParentWindow").Call(obj, uintptr(value))
}

func MonthCalendar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("MonthCalendar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAction").Call(obj)
	return ret
}

func MonthCalendar_SetAction(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetAction").Call(obj, value)
}

func MonthCalendar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("MonthCalendar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("MonthCalendar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MonthCalendar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetClientHeight").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetClientHeight").Call(obj, uintptr(value))
}

func MonthCalendar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("MonthCalendar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("MonthCalendar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MonthCalendar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetClientWidth").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetClientWidth").Call(obj, uintptr(value))
}

func MonthCalendar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("MonthCalendar_GetControlState").Call(obj)
	return TControlState(ret)
}

func MonthCalendar_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("MonthCalendar_SetControlState").Call(obj, uintptr(value))
}

func MonthCalendar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("MonthCalendar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func MonthCalendar_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("MonthCalendar_SetControlStyle").Call(obj, uintptr(value))
}

func MonthCalendar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("MonthCalendar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func MonthCalendar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetParent").Call(obj)
	return ret
}

func MonthCalendar_SetParent(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetParent").Call(obj, value)
}

func MonthCalendar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetLeft").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetLeft(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetLeft").Call(obj, uintptr(value))
}

func MonthCalendar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetTop").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetTop(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetTop").Call(obj, uintptr(value))
}

func MonthCalendar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetWidth").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetWidth(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetWidth").Call(obj, uintptr(value))
}

func MonthCalendar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetHeight").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetHeight(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetHeight").Call(obj, uintptr(value))
}

func MonthCalendar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("MonthCalendar_GetCursor").Call(obj)
	return TCursor(ret)
}

func MonthCalendar_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("MonthCalendar_SetCursor").Call(obj, uintptr(value))
}

func MonthCalendar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("MonthCalendar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func MonthCalendar_SetHint(obj uintptr, value string) {
	getLazyProc("MonthCalendar_SetHint").Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetComponentCount").Call(obj)
	return int32(ret)
}

func MonthCalendar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MonthCalendar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func MonthCalendar_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("MonthCalendar_SetComponentIndex").Call(obj, uintptr(value))
}

func MonthCalendar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetOwner").Call(obj)
	return ret
}

func MonthCalendar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("MonthCalendar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func MonthCalendar_SetName(obj uintptr, value string) {
	getLazyProc("MonthCalendar_SetName").Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("MonthCalendar_GetTag").Call(obj)
	return int(ret)
}

func MonthCalendar_SetTag(obj uintptr, value int) {
	getLazyProc("MonthCalendar_SetTag").Call(obj, uintptr(value))
}

func MonthCalendar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchorSideLeft").Call(obj)
	return ret
}

func MonthCalendar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetAnchorSideLeft").Call(obj, value)
}

func MonthCalendar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchorSideTop").Call(obj)
	return ret
}

func MonthCalendar_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetAnchorSideTop").Call(obj, value)
}

func MonthCalendar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchorSideRight").Call(obj)
	return ret
}

func MonthCalendar_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetAnchorSideRight").Call(obj, value)
}

func MonthCalendar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchorSideBottom").Call(obj)
	return ret
}

func MonthCalendar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetAnchorSideBottom").Call(obj, value)
}

func MonthCalendar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetChildSizing").Call(obj)
	return ret
}

func MonthCalendar_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetChildSizing").Call(obj, value)
}

func MonthCalendar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetBorderSpacing").Call(obj)
	return ret
}

func MonthCalendar_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("MonthCalendar_SetBorderSpacing").Call(obj, value)
}

func MonthCalendar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func MonthCalendar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func MonthCalendar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func MonthCalendar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("MonthCalendar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func MonthCalendar_StaticClassType() TClass {
	r, _, _ := getLazyProc("MonthCalendar_StaticClassType").Call()
	return TClass(r)
}
