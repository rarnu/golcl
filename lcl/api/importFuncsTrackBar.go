package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TTrackBar ---------------------------

func TrackBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_Create").Call(obj)
	return ret
}

func TrackBar_Free(obj uintptr) {
	getLazyProc("TrackBar_Free").Call(obj)
}

func TrackBar_SetTick(obj uintptr, Value int32) {
	getLazyProc("TrackBar_SetTick").Call(obj, uintptr(Value))
}

func TrackBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func TrackBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("TrackBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func TrackBar_DisableAlign(obj uintptr) {
	getLazyProc("TrackBar_DisableAlign").Call(obj)
}

func TrackBar_EnableAlign(obj uintptr) {
	getLazyProc("TrackBar_EnableAlign").Call(obj)
}

func TrackBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("TrackBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func TrackBar_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("TrackBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func TrackBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("TrackBar_InsertControl").Call(obj, AControl)
}

func TrackBar_Invalidate(obj uintptr) {
	getLazyProc("TrackBar_Invalidate").Call(obj)
}

func TrackBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("TrackBar_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func TrackBar_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("TrackBar_RemoveControl").Call(obj, AControl)
}

func TrackBar_Realign(obj uintptr) {
	getLazyProc("TrackBar_Realign").Call(obj)
}

func TrackBar_Repaint(obj uintptr) {
	getLazyProc("TrackBar_Repaint").Call(obj)
}

func TrackBar_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("TrackBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func TrackBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("TrackBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func TrackBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("TrackBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func TrackBar_SetFocus(obj uintptr) {
	getLazyProc("TrackBar_SetFocus").Call(obj)
}

func TrackBar_Update(obj uintptr) {
	getLazyProc("TrackBar_Update").Call(obj)
}

func TrackBar_BringToFront(obj uintptr) {
	getLazyProc("TrackBar_BringToFront").Call(obj)
}

func TrackBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("TrackBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("TrackBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_Hide(obj uintptr) {
	getLazyProc("TrackBar_Hide").Call(obj)
}

func TrackBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("TrackBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func TrackBar_Refresh(obj uintptr) {
	getLazyProc("TrackBar_Refresh").Call(obj)
}

func TrackBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("TrackBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("TrackBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_SendToBack(obj uintptr) {
	getLazyProc("TrackBar_SendToBack").Call(obj)
}

func TrackBar_Show(obj uintptr) {
	getLazyProc("TrackBar_Show").Call(obj)
}

func TrackBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("TrackBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func TrackBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetTextLen").Call(obj)
	return int32(ret)
}

func TrackBar_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("TrackBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func TrackBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("TrackBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func TrackBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TrackBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TrackBar_Assign(obj uintptr, Source uintptr) {
	getLazyProc("TrackBar_Assign").Call(obj, Source)
}

func TrackBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TrackBar_ClassType").Call(obj)
	return TClass(ret)
}

func TrackBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TrackBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TrackBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_InstanceSize").Call(obj)
	return int32(ret)
}

func TrackBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TrackBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TrackBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TrackBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetHashCode").Call(obj)
	return int32(ret)
}

func TrackBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TrackBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TrackBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("TrackBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TrackBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("TrackBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TrackBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("TrackBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func TrackBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("TrackBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func TrackBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("TrackBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func TrackBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("TrackBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func TrackBar_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("TrackBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func TrackBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TrackBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TrackBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("TrackBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func TrackBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("TrackBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func TrackBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("TrackBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func TrackBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("TrackBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func TrackBar_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("TrackBar_SetAlign").Call(obj, uintptr(value))
}

func TrackBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("TrackBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func TrackBar_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("TrackBar_SetAnchors").Call(obj, uintptr(value))
}

func TrackBar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func TrackBar_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetBorderWidth").Call(obj, uintptr(value))
}

func TrackBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("TrackBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func TrackBar_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("TrackBar_SetDragCursor").Call(obj, uintptr(value))
}

func TrackBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("TrackBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func TrackBar_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("TrackBar_SetDragMode").Call(obj, uintptr(value))
}

func TrackBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetEnabled(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetConstraints").Call(obj)
	return ret
}

func TrackBar_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetConstraints").Call(obj, value)
}

func TrackBar_GetLineSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetLineSize").Call(obj)
	return int32(ret)
}

func TrackBar_SetLineSize(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetLineSize").Call(obj, uintptr(value))
}

func TrackBar_GetMax(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetMax").Call(obj)
	return int32(ret)
}

func TrackBar_SetMax(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetMax").Call(obj, uintptr(value))
}

func TrackBar_GetMin(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetMin").Call(obj)
	return int32(ret)
}

func TrackBar_SetMin(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetMin").Call(obj, uintptr(value))
}

func TrackBar_GetOrientation(obj uintptr) TTrackBarOrientation {
	ret, _, _ := getLazyProc("TrackBar_GetOrientation").Call(obj)
	return TTrackBarOrientation(ret)
}

func TrackBar_SetOrientation(obj uintptr, value TTrackBarOrientation) {
	getLazyProc("TrackBar_SetOrientation").Call(obj, uintptr(value))
}

func TrackBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetPageSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetPageSize").Call(obj)
	return int32(ret)
}

func TrackBar_SetPageSize(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetPageSize").Call(obj, uintptr(value))
}

func TrackBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetPopupMenu").Call(obj)
	return ret
}

func TrackBar_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetPopupMenu").Call(obj, value)
}

func TrackBar_GetFrequency(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetFrequency").Call(obj)
	return int32(ret)
}

func TrackBar_SetFrequency(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetFrequency").Call(obj, uintptr(value))
}

func TrackBar_GetPosition(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetPosition").Call(obj)
	return int32(ret)
}

func TrackBar_SetPosition(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetPosition").Call(obj, uintptr(value))
}

func TrackBar_GetSelEnd(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetSelEnd").Call(obj)
	return int32(ret)
}

func TrackBar_SetSelEnd(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetSelEnd").Call(obj, uintptr(value))
}

func TrackBar_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetSelStart").Call(obj)
	return int32(ret)
}

func TrackBar_SetSelStart(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetSelStart").Call(obj, uintptr(value))
}

func TrackBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetShowHint(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetShowSelRange(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetShowSelRange").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetShowSelRange(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetShowSelRange").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("TrackBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func TrackBar_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("TrackBar_SetTabOrder").Call(obj, uintptr(value))
}

func TrackBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetTabStop(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetTickMarks(obj uintptr) TTickMark {
	ret, _, _ := getLazyProc("TrackBar_GetTickMarks").Call(obj)
	return TTickMark(ret)
}

func TrackBar_SetTickMarks(obj uintptr, value TTickMark) {
	getLazyProc("TrackBar_SetTickMarks").Call(obj, uintptr(value))
}

func TrackBar_GetTickStyle(obj uintptr) TTickStyle {
	ret, _, _ := getLazyProc("TrackBar_GetTickStyle").Call(obj)
	return TTickStyle(ret)
}

func TrackBar_SetTickStyle(obj uintptr, value TTickStyle) {
	getLazyProc("TrackBar_SetTickStyle").Call(obj, uintptr(value))
}

func TrackBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetVisible(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func TrackBar_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("TrackBar_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func TrackBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func TrackBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetDockSite(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func TrackBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetBrush").Call(obj)
	return ret
}

func TrackBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetControlCount").Call(obj)
	return int32(ret)
}

func TrackBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TrackBar_GetHandle").Call(obj)
	return HWND(ret)
}

func TrackBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TrackBar_GetParentWindow").Call(obj)
	return HWND(ret)
}

func TrackBar_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("TrackBar_SetParentWindow").Call(obj, uintptr(value))
}

func TrackBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("TrackBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAction").Call(obj)
	return ret
}

func TrackBar_SetAction(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetAction").Call(obj, value)
}

func TrackBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("TrackBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func TrackBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("TrackBar_SetBiDiMode").Call(obj, uintptr(value))
}

func TrackBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("TrackBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("TrackBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TrackBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func TrackBar_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetClientHeight").Call(obj, uintptr(value))
}

func TrackBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("TrackBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("TrackBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TrackBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func TrackBar_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetClientWidth").Call(obj, uintptr(value))
}

func TrackBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("TrackBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func TrackBar_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("TrackBar_SetControlState").Call(obj, uintptr(value))
}

func TrackBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("TrackBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func TrackBar_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("TrackBar_SetControlStyle").Call(obj, uintptr(value))
}

func TrackBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("TrackBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func TrackBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetParent").Call(obj)
	return ret
}

func TrackBar_SetParent(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetParent").Call(obj, value)
}

func TrackBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetLeft").Call(obj)
	return int32(ret)
}

func TrackBar_SetLeft(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetLeft").Call(obj, uintptr(value))
}

func TrackBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetTop").Call(obj)
	return int32(ret)
}

func TrackBar_SetTop(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetTop").Call(obj, uintptr(value))
}

func TrackBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetWidth").Call(obj)
	return int32(ret)
}

func TrackBar_SetWidth(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetWidth").Call(obj, uintptr(value))
}

func TrackBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetHeight").Call(obj)
	return int32(ret)
}

func TrackBar_SetHeight(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetHeight").Call(obj, uintptr(value))
}

func TrackBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("TrackBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func TrackBar_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("TrackBar_SetCursor").Call(obj, uintptr(value))
}

func TrackBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("TrackBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func TrackBar_SetHint(obj uintptr, value string) {
	getLazyProc("TrackBar_SetHint").Call(obj, GoStrToDStr(value))
}

func TrackBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func TrackBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TrackBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func TrackBar_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("TrackBar_SetComponentIndex").Call(obj, uintptr(value))
}

func TrackBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetOwner").Call(obj)
	return ret
}

func TrackBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TrackBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TrackBar_SetName(obj uintptr, value string) {
	getLazyProc("TrackBar_SetName").Call(obj, GoStrToDStr(value))
}

func TrackBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("TrackBar_GetTag").Call(obj)
	return int(ret)
}

func TrackBar_SetTag(obj uintptr, value int) {
	getLazyProc("TrackBar_SetTag").Call(obj, uintptr(value))
}

func TrackBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func TrackBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetAnchorSideLeft").Call(obj, value)
}

func TrackBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAnchorSideTop").Call(obj)
	return ret
}

func TrackBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetAnchorSideTop").Call(obj, value)
}

func TrackBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAnchorSideRight").Call(obj)
	return ret
}

func TrackBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetAnchorSideRight").Call(obj, value)
}

func TrackBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func TrackBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetAnchorSideBottom").Call(obj, value)
}

func TrackBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetChildSizing").Call(obj)
	return ret
}

func TrackBar_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetChildSizing").Call(obj, value)
}

func TrackBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetBorderSpacing").Call(obj)
	return ret
}

func TrackBar_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("TrackBar_SetBorderSpacing").Call(obj, value)
}

func TrackBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func TrackBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func TrackBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func TrackBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("TrackBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func TrackBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("TrackBar_StaticClassType").Call()
	return TClass(r)
}
