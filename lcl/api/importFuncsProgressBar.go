package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TProgressBar ---------------------------

func ProgressBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_Create").Call(obj)
	return ret
}

func ProgressBar_Free(obj uintptr) {
	getLazyProc("ProgressBar_Free").Call(obj)
}

func ProgressBar_StepIt(obj uintptr) {
	getLazyProc("ProgressBar_StepIt").Call(obj)
}

func ProgressBar_StepBy(obj uintptr, Delta int32) {
	getLazyProc("ProgressBar_StepBy").Call(obj, uintptr(Delta))
}

func ProgressBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ProgressBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ProgressBar_DisableAlign(obj uintptr) {
	getLazyProc("ProgressBar_DisableAlign").Call(obj)
}

func ProgressBar_EnableAlign(obj uintptr) {
	getLazyProc("ProgressBar_EnableAlign").Call(obj)
}

func ProgressBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ProgressBar_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ProgressBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ProgressBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ProgressBar_InsertControl").Call(obj, AControl)
}

func ProgressBar_Invalidate(obj uintptr) {
	getLazyProc("ProgressBar_Invalidate").Call(obj)
}

func ProgressBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ProgressBar_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ProgressBar_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ProgressBar_RemoveControl").Call(obj, AControl)
}

func ProgressBar_Realign(obj uintptr) {
	getLazyProc("ProgressBar_Realign").Call(obj)
}

func ProgressBar_Repaint(obj uintptr) {
	getLazyProc("ProgressBar_Repaint").Call(obj)
}

func ProgressBar_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ProgressBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ProgressBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ProgressBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ProgressBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ProgressBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ProgressBar_SetFocus(obj uintptr) {
	getLazyProc("ProgressBar_SetFocus").Call(obj)
}

func ProgressBar_Update(obj uintptr) {
	getLazyProc("ProgressBar_Update").Call(obj)
}

func ProgressBar_BringToFront(obj uintptr) {
	getLazyProc("ProgressBar_BringToFront").Call(obj)
}

func ProgressBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ProgressBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ProgressBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_Hide(obj uintptr) {
	getLazyProc("ProgressBar_Hide").Call(obj)
}

func ProgressBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ProgressBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ProgressBar_Refresh(obj uintptr) {
	getLazyProc("ProgressBar_Refresh").Call(obj)
}

func ProgressBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ProgressBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ProgressBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_SendToBack(obj uintptr) {
	getLazyProc("ProgressBar_SendToBack").Call(obj)
}

func ProgressBar_Show(obj uintptr) {
	getLazyProc("ProgressBar_Show").Call(obj)
}

func ProgressBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ProgressBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ProgressBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetTextLen").Call(obj)
	return int32(ret)
}

func ProgressBar_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ProgressBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ProgressBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ProgressBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ProgressBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ProgressBar_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ProgressBar_Assign").Call(obj, Source)
}

func ProgressBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ProgressBar_ClassType").Call(obj)
	return TClass(ret)
}

func ProgressBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ProgressBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ProgressBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_InstanceSize").Call(obj)
	return int32(ret)
}

func ProgressBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ProgressBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ProgressBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetHashCode").Call(obj)
	return int32(ret)
}

func ProgressBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ProgressBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ProgressBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ProgressBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ProgressBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ProgressBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ProgressBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ProgressBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ProgressBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ProgressBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ProgressBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ProgressBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ProgressBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ProgressBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ProgressBar_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ProgressBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func ProgressBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ProgressBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ProgressBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ProgressBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ProgressBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ProgressBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ProgressBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ProgressBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ProgressBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ProgressBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func ProgressBar_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ProgressBar_SetAlign").Call(obj, uintptr(value))
}

func ProgressBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ProgressBar_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ProgressBar_SetAnchors").Call(obj, uintptr(value))
}

func ProgressBar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func ProgressBar_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetBorderWidth").Call(obj, uintptr(value))
}

func ProgressBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ProgressBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ProgressBar_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ProgressBar_SetDragCursor").Call(obj, uintptr(value))
}

func ProgressBar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ProgressBar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ProgressBar_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ProgressBar_SetDragKind").Call(obj, uintptr(value))
}

func ProgressBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ProgressBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ProgressBar_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ProgressBar_SetDragMode").Call(obj, uintptr(value))
}

func ProgressBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ProgressBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ProgressBar_SetHint(obj uintptr, value string) {
	getLazyProc("ProgressBar_SetHint").Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetConstraints").Call(obj)
	return ret
}

func ProgressBar_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetConstraints").Call(obj, value)
}

func ProgressBar_GetMin(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetMin").Call(obj)
	return int32(ret)
}

func ProgressBar_SetMin(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetMin").Call(obj, uintptr(value))
}

func ProgressBar_GetMax(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetMax").Call(obj)
	return int32(ret)
}

func ProgressBar_SetMax(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetMax").Call(obj, uintptr(value))
}

func ProgressBar_GetOrientation(obj uintptr) TProgressBarOrientation {
	ret, _, _ := getLazyProc("ProgressBar_GetOrientation").Call(obj)
	return TProgressBarOrientation(ret)
}

func ProgressBar_SetOrientation(obj uintptr, value TProgressBarOrientation) {
	getLazyProc("ProgressBar_SetOrientation").Call(obj, uintptr(value))
}

func ProgressBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetPopupMenu").Call(obj)
	return ret
}

func ProgressBar_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetPopupMenu").Call(obj, value)
}

func ProgressBar_GetPosition(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetPosition").Call(obj)
	return int32(ret)
}

func ProgressBar_SetPosition(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetPosition").Call(obj, uintptr(value))
}

func ProgressBar_GetSmooth(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetSmooth").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetSmooth(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetSmooth").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetStyle(obj uintptr) TProgressBarStyle {
	ret, _, _ := getLazyProc("ProgressBar_GetStyle").Call(obj)
	return TProgressBarStyle(ret)
}

func ProgressBar_SetStyle(obj uintptr, value TProgressBarStyle) {
	getLazyProc("ProgressBar_SetStyle").Call(obj, uintptr(value))
}

func ProgressBar_GetStep(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetStep").Call(obj)
	return int32(ret)
}

func ProgressBar_SetStep(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetStep").Call(obj, uintptr(value))
}

func ProgressBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ProgressBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ProgressBar_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ProgressBar_SetTabOrder").Call(obj, uintptr(value))
}

func ProgressBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetVisible(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ProgressBar_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ProgressBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ProgressBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ProgressBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetBrush").Call(obj)
	return ret
}

func ProgressBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetControlCount").Call(obj)
	return int32(ret)
}

func ProgressBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ProgressBar_GetHandle").Call(obj)
	return HWND(ret)
}

func ProgressBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ProgressBar_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ProgressBar_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ProgressBar_SetParentWindow").Call(obj, uintptr(value))
}

func ProgressBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ProgressBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAction").Call(obj)
	return ret
}

func ProgressBar_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetAction").Call(obj, value)
}

func ProgressBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ProgressBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ProgressBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ProgressBar_SetBiDiMode").Call(obj, uintptr(value))
}

func ProgressBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ProgressBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ProgressBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ProgressBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func ProgressBar_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetClientHeight").Call(obj, uintptr(value))
}

func ProgressBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ProgressBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ProgressBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ProgressBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func ProgressBar_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetClientWidth").Call(obj, uintptr(value))
}

func ProgressBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ProgressBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func ProgressBar_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ProgressBar_SetControlState").Call(obj, uintptr(value))
}

func ProgressBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ProgressBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ProgressBar_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ProgressBar_SetControlStyle").Call(obj, uintptr(value))
}

func ProgressBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ProgressBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ProgressBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetParent").Call(obj)
	return ret
}

func ProgressBar_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetParent").Call(obj, value)
}

func ProgressBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetLeft").Call(obj)
	return int32(ret)
}

func ProgressBar_SetLeft(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetLeft").Call(obj, uintptr(value))
}

func ProgressBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetTop").Call(obj)
	return int32(ret)
}

func ProgressBar_SetTop(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetTop").Call(obj, uintptr(value))
}

func ProgressBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetWidth").Call(obj)
	return int32(ret)
}

func ProgressBar_SetWidth(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetWidth").Call(obj, uintptr(value))
}

func ProgressBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetHeight").Call(obj)
	return int32(ret)
}

func ProgressBar_SetHeight(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetHeight").Call(obj, uintptr(value))
}

func ProgressBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ProgressBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func ProgressBar_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ProgressBar_SetCursor").Call(obj, uintptr(value))
}

func ProgressBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func ProgressBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ProgressBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ProgressBar_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ProgressBar_SetComponentIndex").Call(obj, uintptr(value))
}

func ProgressBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetOwner").Call(obj)
	return ret
}

func ProgressBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ProgressBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ProgressBar_SetName(obj uintptr, value string) {
	getLazyProc("ProgressBar_SetName").Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ProgressBar_GetTag").Call(obj)
	return int(ret)
}

func ProgressBar_SetTag(obj uintptr, value int) {
	getLazyProc("ProgressBar_SetTag").Call(obj, uintptr(value))
}

func ProgressBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func ProgressBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetAnchorSideLeft").Call(obj, value)
}

func ProgressBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchorSideTop").Call(obj)
	return ret
}

func ProgressBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetAnchorSideTop").Call(obj, value)
}

func ProgressBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchorSideRight").Call(obj)
	return ret
}

func ProgressBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetAnchorSideRight").Call(obj, value)
}

func ProgressBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func ProgressBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetAnchorSideBottom").Call(obj, value)
}

func ProgressBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetChildSizing").Call(obj)
	return ret
}

func ProgressBar_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetChildSizing").Call(obj, value)
}

func ProgressBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetBorderSpacing").Call(obj)
	return ret
}

func ProgressBar_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ProgressBar_SetBorderSpacing").Call(obj, value)
}

func ProgressBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ProgressBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ProgressBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ProgressBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ProgressBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ProgressBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("ProgressBar_StaticClassType").Call()
	return TClass(r)
}
