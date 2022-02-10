package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TToggleBox ---------------------------

func ToggleBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_Create").Call(obj)
	return ret
}

func ToggleBox_Free(obj uintptr) {
	getLazyProc("ToggleBox_Free").Call(obj)
}

func ToggleBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ToggleBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ToggleBox_DisableAlign(obj uintptr) {
	getLazyProc("ToggleBox_DisableAlign").Call(obj)
}

func ToggleBox_EnableAlign(obj uintptr) {
	getLazyProc("ToggleBox_EnableAlign").Call(obj)
}

func ToggleBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ToggleBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ToggleBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ToggleBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ToggleBox_InsertControl").Call(obj, AControl)
}

func ToggleBox_Invalidate(obj uintptr) {
	getLazyProc("ToggleBox_Invalidate").Call(obj)
}

func ToggleBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ToggleBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ToggleBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ToggleBox_RemoveControl").Call(obj, AControl)
}

func ToggleBox_Realign(obj uintptr) {
	getLazyProc("ToggleBox_Realign").Call(obj)
}

func ToggleBox_Repaint(obj uintptr) {
	getLazyProc("ToggleBox_Repaint").Call(obj)
}

func ToggleBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ToggleBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ToggleBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ToggleBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ToggleBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ToggleBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ToggleBox_SetFocus(obj uintptr) {
	getLazyProc("ToggleBox_SetFocus").Call(obj)
}

func ToggleBox_Update(obj uintptr) {
	getLazyProc("ToggleBox_Update").Call(obj)
}

func ToggleBox_BringToFront(obj uintptr) {
	getLazyProc("ToggleBox_BringToFront").Call(obj)
}

func ToggleBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ToggleBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToggleBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_Hide(obj uintptr) {
	getLazyProc("ToggleBox_Hide").Call(obj)
}

func ToggleBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ToggleBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ToggleBox_Refresh(obj uintptr) {
	getLazyProc("ToggleBox_Refresh").Call(obj)
}

func ToggleBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ToggleBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToggleBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_SendToBack(obj uintptr) {
	getLazyProc("ToggleBox_SendToBack").Call(obj)
}

func ToggleBox_Show(obj uintptr) {
	getLazyProc("ToggleBox_Show").Call(obj)
}

func ToggleBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ToggleBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ToggleBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ToggleBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ToggleBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ToggleBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ToggleBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ToggleBox_Assign").Call(obj, Source)
}

func ToggleBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ToggleBox_ClassType").Call(obj)
	return TClass(ret)
}

func ToggleBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ToggleBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ToggleBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ToggleBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ToggleBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ToggleBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToggleBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ToggleBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToggleBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ToggleBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ToggleBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ToggleBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ToggleBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ToggleBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ToggleBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ToggleBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ToggleBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ToggleBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ToggleBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToggleBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToggleBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ToggleBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ToggleBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ToggleBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ToggleBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ToggleBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ToggleBox_GetAllowGrayed(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetAllowGrayed").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetAllowGrayed(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetAllowGrayed").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ToggleBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ToggleBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ToggleBox_SetAlign").Call(obj, uintptr(value))
}

func ToggleBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ToggleBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ToggleBox_SetAnchors").Call(obj, uintptr(value))
}

func ToggleBox_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_SetCaption(obj uintptr, value string) {
	getLazyProc("ToggleBox_SetCaption").Call(obj, GoStrToDStr(value))
}

func ToggleBox_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetChecked(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetChecked").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ToggleBox_GetColor").Call(obj)
	return TColor(ret)
}

func ToggleBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("ToggleBox_SetColor").Call(obj, uintptr(value))
}

func ToggleBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetConstraints").Call(obj)
	return ret
}

func ToggleBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetConstraints").Call(obj, value)
}

func ToggleBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToggleBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ToggleBox_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ToggleBox_SetDragCursor").Call(obj, uintptr(value))
}

func ToggleBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ToggleBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ToggleBox_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ToggleBox_SetDragKind").Call(obj, uintptr(value))
}

func ToggleBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ToggleBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ToggleBox_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ToggleBox_SetDragMode").Call(obj, uintptr(value))
}

func ToggleBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetFont").Call(obj)
	return ret
}

func ToggleBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetFont").Call(obj, value)
}

func ToggleBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_SetHint(obj uintptr, value string) {
	getLazyProc("ToggleBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ToggleBox_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseWheel(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	getLazyProc("ToggleBox_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func ToggleBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetPopupMenu").Call(obj)
	return ret
}

func ToggleBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetPopupMenu").Call(obj, value)
}

func ToggleBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetState(obj uintptr) TCheckBoxState {
	ret, _, _ := getLazyProc("ToggleBox_GetState").Call(obj)
	return TCheckBoxState(ret)
}

func ToggleBox_SetState(obj uintptr, value TCheckBoxState) {
	getLazyProc("ToggleBox_SetState").Call(obj, uintptr(value))
}

func ToggleBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ToggleBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ToggleBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ToggleBox_SetTabOrder").Call(obj, uintptr(value))
}

func ToggleBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ToggleBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ToggleBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetBrush").Call(obj)
	return ret
}

func ToggleBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ToggleBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ToggleBox_GetHandle").Call(obj)
	return HWND(ret)
}

func ToggleBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ToggleBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ToggleBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ToggleBox_SetParentWindow").Call(obj, uintptr(value))
}

func ToggleBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ToggleBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ToggleBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAction").Call(obj)
	return ret
}

func ToggleBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetAction").Call(obj, value)
}

func ToggleBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ToggleBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ToggleBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ToggleBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ToggleBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ToggleBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ToggleBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToggleBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ToggleBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetClientHeight").Call(obj, uintptr(value))
}

func ToggleBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToggleBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ToggleBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToggleBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ToggleBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetClientWidth").Call(obj, uintptr(value))
}

func ToggleBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ToggleBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ToggleBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ToggleBox_SetControlState").Call(obj, uintptr(value))
}

func ToggleBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ToggleBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ToggleBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ToggleBox_SetControlStyle").Call(obj, uintptr(value))
}

func ToggleBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToggleBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ToggleBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetParent").Call(obj)
	return ret
}

func ToggleBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetParent").Call(obj, value)
}

func ToggleBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetLeft").Call(obj)
	return int32(ret)
}

func ToggleBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetLeft").Call(obj, uintptr(value))
}

func ToggleBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetTop").Call(obj)
	return int32(ret)
}

func ToggleBox_SetTop(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetTop").Call(obj, uintptr(value))
}

func ToggleBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetWidth").Call(obj)
	return int32(ret)
}

func ToggleBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetWidth").Call(obj, uintptr(value))
}

func ToggleBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetHeight").Call(obj)
	return int32(ret)
}

func ToggleBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetHeight").Call(obj, uintptr(value))
}

func ToggleBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToggleBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ToggleBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ToggleBox_SetCursor").Call(obj, uintptr(value))
}

func ToggleBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ToggleBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToggleBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ToggleBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ToggleBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ToggleBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetOwner").Call(obj)
	return ret
}

func ToggleBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToggleBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ToggleBox_SetName(obj uintptr, value string) {
	getLazyProc("ToggleBox_SetName").Call(obj, GoStrToDStr(value))
}

func ToggleBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ToggleBox_GetTag").Call(obj)
	return int(ret)
}

func ToggleBox_SetTag(obj uintptr, value int) {
	getLazyProc("ToggleBox_SetTag").Call(obj, uintptr(value))
}

func ToggleBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ToggleBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetAnchorSideLeft").Call(obj, value)
}

func ToggleBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ToggleBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetAnchorSideTop").Call(obj, value)
}

func ToggleBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ToggleBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetAnchorSideRight").Call(obj, value)
}

func ToggleBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ToggleBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetAnchorSideBottom").Call(obj, value)
}

func ToggleBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetChildSizing").Call(obj)
	return ret
}

func ToggleBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetChildSizing").Call(obj, value)
}

func ToggleBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetBorderSpacing").Call(obj)
	return ret
}

func ToggleBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ToggleBox_SetBorderSpacing").Call(obj, value)
}

func ToggleBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ToggleBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ToggleBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ToggleBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ToggleBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ToggleBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ToggleBox_StaticClassType").Call()
	return TClass(r)
}
