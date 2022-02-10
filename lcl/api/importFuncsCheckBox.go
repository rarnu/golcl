package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCheckBox ---------------------------

func CheckBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_Create").Call(obj)
	return ret
}

func CheckBox_Free(obj uintptr) {
	getLazyProc("CheckBox_Free").Call(obj)
}

func CheckBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func CheckBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("CheckBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func CheckBox_DisableAlign(obj uintptr) {
	getLazyProc("CheckBox_DisableAlign").Call(obj)
}

func CheckBox_EnableAlign(obj uintptr) {
	getLazyProc("CheckBox_EnableAlign").Call(obj)
}

func CheckBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("CheckBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func CheckBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("CheckBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func CheckBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("CheckBox_InsertControl").Call(obj, AControl)
}

func CheckBox_Invalidate(obj uintptr) {
	getLazyProc("CheckBox_Invalidate").Call(obj)
}

func CheckBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("CheckBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func CheckBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("CheckBox_RemoveControl").Call(obj, AControl)
}

func CheckBox_Realign(obj uintptr) {
	getLazyProc("CheckBox_Realign").Call(obj)
}

func CheckBox_Repaint(obj uintptr) {
	getLazyProc("CheckBox_Repaint").Call(obj)
}

func CheckBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("CheckBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func CheckBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("CheckBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func CheckBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("CheckBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func CheckBox_SetFocus(obj uintptr) {
	getLazyProc("CheckBox_SetFocus").Call(obj)
}

func CheckBox_Update(obj uintptr) {
	getLazyProc("CheckBox_Update").Call(obj)
}

func CheckBox_BringToFront(obj uintptr) {
	getLazyProc("CheckBox_BringToFront").Call(obj)
}

func CheckBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("CheckBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("CheckBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_Hide(obj uintptr) {
	getLazyProc("CheckBox_Hide").Call(obj)
}

func CheckBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("CheckBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func CheckBox_Refresh(obj uintptr) {
	getLazyProc("CheckBox_Refresh").Call(obj)
}

func CheckBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("CheckBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("CheckBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_SendToBack(obj uintptr) {
	getLazyProc("CheckBox_SendToBack").Call(obj)
}

func CheckBox_Show(obj uintptr) {
	getLazyProc("CheckBox_Show").Call(obj)
}

func CheckBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("CheckBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func CheckBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetTextLen").Call(obj)
	return int32(ret)
}

func CheckBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("CheckBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func CheckBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("CheckBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func CheckBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("CheckBox_Assign").Call(obj, Source)
}

func CheckBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CheckBox_ClassType").Call(obj)
	return TClass(ret)
}

func CheckBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_InstanceSize").Call(obj)
	return int32(ret)
}

func CheckBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CheckBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CheckBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CheckBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetHashCode").Call(obj)
	return int32(ret)
}

func CheckBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("CheckBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("CheckBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("CheckBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func CheckBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("CheckBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func CheckBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("CheckBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func CheckBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("CheckBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func CheckBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("CheckBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func CheckBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("CheckBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func CheckBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("CheckBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func CheckBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("CheckBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func CheckBox_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAction").Call(obj)
	return ret
}

func CheckBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetAction").Call(obj, value)
}

func CheckBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("CheckBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func CheckBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("CheckBox_SetAlign").Call(obj, uintptr(value))
}

func CheckBox_GetAlignment(obj uintptr) TLeftRight {
	ret, _, _ := getLazyProc("CheckBox_GetAlignment").Call(obj)
	return TLeftRight(ret)
}

func CheckBox_SetAlignment(obj uintptr, value TLeftRight) {
	getLazyProc("CheckBox_SetAlignment").Call(obj, uintptr(value))
}

func CheckBox_GetAllowGrayed(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetAllowGrayed").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetAllowGrayed(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetAllowGrayed").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("CheckBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func CheckBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("CheckBox_SetAnchors").Call(obj, uintptr(value))
}

func CheckBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("CheckBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func CheckBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("CheckBox_SetBiDiMode").Call(obj, uintptr(value))
}

func CheckBox_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_SetCaption(obj uintptr, value string) {
	getLazyProc("CheckBox_SetCaption").Call(obj, GoStrToDStr(value))
}

func CheckBox_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetChecked(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetChecked").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CheckBox_GetColor").Call(obj)
	return TColor(ret)
}

func CheckBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("CheckBox_SetColor").Call(obj, uintptr(value))
}

func CheckBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetConstraints").Call(obj)
	return ret
}

func CheckBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetConstraints").Call(obj, value)
}

func CheckBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func CheckBox_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("CheckBox_SetDragCursor").Call(obj, uintptr(value))
}

func CheckBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("CheckBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func CheckBox_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("CheckBox_SetDragKind").Call(obj, uintptr(value))
}

func CheckBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("CheckBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func CheckBox_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("CheckBox_SetDragMode").Call(obj, uintptr(value))
}

func CheckBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetFont").Call(obj)
	return ret
}

func CheckBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetFont").Call(obj, value)
}

func CheckBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetParentColor(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetPopupMenu").Call(obj)
	return ret
}

func CheckBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetPopupMenu").Call(obj, value)
}

func CheckBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetState(obj uintptr) TCheckBoxState {
	ret, _, _ := getLazyProc("CheckBox_GetState").Call(obj)
	return TCheckBoxState(ret)
}

func CheckBox_SetState(obj uintptr, value TCheckBoxState) {
	getLazyProc("CheckBox_SetState").Call(obj, uintptr(value))
}

func CheckBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("CheckBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func CheckBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("CheckBox_SetTabOrder").Call(obj, uintptr(value))
}

func CheckBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CheckBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("CheckBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func CheckBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func CheckBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func CheckBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetBrush").Call(obj)
	return ret
}

func CheckBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetControlCount").Call(obj)
	return int32(ret)
}

func CheckBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckBox_GetHandle").Call(obj)
	return HWND(ret)
}

func CheckBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func CheckBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("CheckBox_SetParentWindow").Call(obj, uintptr(value))
}

func CheckBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("CheckBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("CheckBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("CheckBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func CheckBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetClientHeight").Call(obj, uintptr(value))
}

func CheckBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("CheckBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("CheckBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func CheckBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetClientWidth").Call(obj, uintptr(value))
}

func CheckBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("CheckBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func CheckBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("CheckBox_SetControlState").Call(obj, uintptr(value))
}

func CheckBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("CheckBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func CheckBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("CheckBox_SetControlStyle").Call(obj, uintptr(value))
}

func CheckBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetParent").Call(obj)
	return ret
}

func CheckBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetParent").Call(obj, value)
}

func CheckBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetLeft").Call(obj)
	return int32(ret)
}

func CheckBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetLeft").Call(obj, uintptr(value))
}

func CheckBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetTop").Call(obj)
	return int32(ret)
}

func CheckBox_SetTop(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetTop").Call(obj, uintptr(value))
}

func CheckBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetWidth").Call(obj)
	return int32(ret)
}

func CheckBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetWidth").Call(obj, uintptr(value))
}

func CheckBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetHeight").Call(obj)
	return int32(ret)
}

func CheckBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetHeight").Call(obj, uintptr(value))
}

func CheckBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func CheckBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("CheckBox_SetCursor").Call(obj, uintptr(value))
}

func CheckBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_SetHint(obj uintptr, value string) {
	getLazyProc("CheckBox_SetHint").Call(obj, GoStrToDStr(value))
}

func CheckBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func CheckBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func CheckBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("CheckBox_SetComponentIndex").Call(obj, uintptr(value))
}

func CheckBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetOwner").Call(obj)
	return ret
}

func CheckBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckBox_SetName(obj uintptr, value string) {
	getLazyProc("CheckBox_SetName").Call(obj, GoStrToDStr(value))
}

func CheckBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("CheckBox_GetTag").Call(obj)
	return int(ret)
}

func CheckBox_SetTag(obj uintptr, value int) {
	getLazyProc("CheckBox_SetTag").Call(obj, uintptr(value))
}

func CheckBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func CheckBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetAnchorSideLeft").Call(obj, value)
}

func CheckBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAnchorSideTop").Call(obj)
	return ret
}

func CheckBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetAnchorSideTop").Call(obj, value)
}

func CheckBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAnchorSideRight").Call(obj)
	return ret
}

func CheckBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetAnchorSideRight").Call(obj, value)
}

func CheckBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func CheckBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetAnchorSideBottom").Call(obj, value)
}

func CheckBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetChildSizing").Call(obj)
	return ret
}

func CheckBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetChildSizing").Call(obj, value)
}

func CheckBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetBorderSpacing").Call(obj)
	return ret
}

func CheckBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("CheckBox_SetBorderSpacing").Call(obj, value)
}

func CheckBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func CheckBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func CheckBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func CheckBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("CheckBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func CheckBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("CheckBox_StaticClassType").Call()
	return TClass(r)
}
