package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCheckComboBox ---------------------------

func CheckComboBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_Create").Call(obj)
	return ret
}

func CheckComboBox_Free(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Free").Call(obj)
}

func CheckComboBox_AddItem(obj uintptr, AItem string, AState TCheckBoxState, AEnabled bool) {
	_, _, _ = getLazyProc("CheckComboBox_AddItem").Call(obj, GoStrToDStr(AItem), uintptr(AState), GoBoolToDBool(AEnabled))
}

func CheckComboBox_AssignItems(obj uintptr, AItems uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AssignItems").Call(obj, AItems)
}

func CheckComboBox_Clear(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Clear").Call(obj)
}

func CheckComboBox_DeleteItem(obj uintptr, AIndex int32) {
	_, _, _ = getLazyProc("CheckComboBox_DeleteItem").Call(obj, uintptr(AIndex))
}

func CheckComboBox_CheckAll(obj uintptr, AState TCheckBoxState, AAllowGrayed bool, AAllowDisabled bool) {
	_, _, _ = getLazyProc("CheckComboBox_CheckAll").Call(obj, uintptr(AState), GoBoolToDBool(AAllowGrayed), GoBoolToDBool(AAllowDisabled))
}

func CheckComboBox_Toggle(obj uintptr, AIndex int32) {
	_, _, _ = getLazyProc("CheckComboBox_Toggle").Call(obj, uintptr(AIndex))
}

func CheckComboBox_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_ClearSelection").Call(obj)
}

func CheckComboBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SelectAll").Call(obj)
}

func CheckComboBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func CheckComboBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func CheckComboBox_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_DisableAlign").Call(obj)
}

func CheckComboBox_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_EnableAlign").Call(obj)
}

func CheckComboBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func CheckComboBox_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("CheckComboBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func CheckComboBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_InsertControl").Call(obj, AControl)
}

func CheckComboBox_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Invalidate").Call(obj)
}

func CheckComboBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("CheckComboBox_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func CheckComboBox_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_RemoveControl").Call(obj, AControl)
}

func CheckComboBox_Realign(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Realign").Call(obj)
}

func CheckComboBox_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Repaint").Call(obj)
}

func CheckComboBox_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("CheckComboBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func CheckComboBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("CheckComboBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func CheckComboBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func CheckComboBox_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetFocus").Call(obj)
}

func CheckComboBox_Update(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Update").Call(obj)
}

func CheckComboBox_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_BringToFront").Call(obj)
}

func CheckComboBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckComboBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckComboBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_Hide(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Hide").Call(obj)
}

func CheckComboBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("CheckComboBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func CheckComboBox_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Refresh").Call(obj)
}

func CheckComboBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckComboBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckComboBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SendToBack").Call(obj)
}

func CheckComboBox_Show(obj uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Show").Call(obj)
}

func CheckComboBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("CheckComboBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func CheckComboBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetTextLen").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("CheckComboBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func CheckComboBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func CheckComboBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_Assign").Call(obj, Source)
}

func CheckComboBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CheckComboBox_ClassType").Call(obj)
	return TClass(ret)
}

func CheckComboBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_InstanceSize").Call(obj)
	return int32(ret)
}

func CheckComboBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CheckComboBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CheckComboBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetHashCode").Call(obj)
	return int32(ret)
}

func CheckComboBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckComboBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckComboBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func CheckComboBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func CheckComboBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func CheckComboBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func CheckComboBox_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("CheckComboBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func CheckComboBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckComboBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("CheckComboBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func CheckComboBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("CheckComboBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func CheckComboBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("CheckComboBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func CheckComboBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("CheckComboBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func CheckComboBox_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("CheckComboBox_SetAlign").Call(obj, uintptr(value))
}

func CheckComboBox_GetAllowGrayed(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetAllowGrayed").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetAllowGrayed(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetAllowGrayed").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func CheckComboBox_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("CheckComboBox_SetAnchors").Call(obj, uintptr(value))
}

func CheckComboBox_GetAutoDropDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetAutoDropDown").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetAutoDropDown(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetAutoDropDown").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CheckComboBox_GetColor").Call(obj)
	return TColor(ret)
}

func CheckComboBox_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("CheckComboBox_SetColor").Call(obj, uintptr(value))
}

func CheckComboBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetConstraints").Call(obj)
	return ret
}

func CheckComboBox_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetConstraints").Call(obj, value)
}

func CheckComboBox_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckComboBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func CheckComboBox_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckComboBox_SetDragCursor").Call(obj, uintptr(value))
}

func CheckComboBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("CheckComboBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func CheckComboBox_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("CheckComboBox_SetDragKind").Call(obj, uintptr(value))
}

func CheckComboBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("CheckComboBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func CheckComboBox_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("CheckComboBox_SetDragMode").Call(obj, uintptr(value))
}

func CheckComboBox_GetDropDownCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetDropDownCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetDropDownCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetDropDownCount").Call(obj, uintptr(value))
}

func CheckComboBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetFont").Call(obj)
	return ret
}

func CheckComboBox_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetFont").Call(obj, value)
}

func CheckComboBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetItemHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetItemHeight").Call(obj, uintptr(value))
}

func CheckComboBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetItemIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetItemIndex").Call(obj, uintptr(value))
}

func CheckComboBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetItems").Call(obj)
	return ret
}

func CheckComboBox_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetItems").Call(obj, value)
}

func CheckComboBox_GetItemWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetItemWidth").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetItemWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetItemWidth").Call(obj, uintptr(value))
}

func CheckComboBox_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetMaxLength").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetMaxLength").Call(obj, uintptr(value))
}

func CheckComboBox_SetOnChange(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnCloseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnCloseUp").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnContextPopup(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnDblClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnDragOver(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnDropDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnDropDown").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnExit(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnItemChange(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnItemChange").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnKeyDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnKeyPress(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnKeyUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseMove(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseWheel(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_SetOnSelect(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("CheckComboBox_SetOnSelect").Call(obj, addEventToMap(obj, fn))
}

func CheckComboBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetPopupMenu").Call(obj)
	return ret
}

func CheckComboBox_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetPopupMenu").Call(obj, value)
}

func CheckComboBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetSorted(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetSorted").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetSorted(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetSorted").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("CheckComboBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func CheckComboBox_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("CheckComboBox_SetTabOrder").Call(obj, uintptr(value))
}

func CheckComboBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckComboBox_SetText").Call(obj, GoStrToDStr(value))
}

func CheckComboBox_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetTextHint").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckComboBox_SetTextHint").Call(obj, GoStrToDStr(value))
}

func CheckComboBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetAutoComplete(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetAutoComplete").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetAutoComplete(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetAutoComplete").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("CheckComboBox_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func CheckComboBox_SetCharCase(obj uintptr, value TEditCharCase) {
	_, _, _ = getLazyProc("CheckComboBox_SetCharCase").Call(obj, uintptr(value))
}

func CheckComboBox_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckComboBox_SetSelText").Call(obj, GoStrToDStr(value))
}

func CheckComboBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetCanvas").Call(obj)
	return ret
}

func CheckComboBox_GetDroppedDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetDroppedDown").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetDroppedDown(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetDroppedDown").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetSelLength").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetSelLength").Call(obj, uintptr(value))
}

func CheckComboBox_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetSelStart").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetSelStart").Call(obj, uintptr(value))
}

func CheckComboBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetBrush").Call(obj)
	return ret
}

func CheckComboBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetControlCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckComboBox_GetHandle").Call(obj)
	return ret
}

func CheckComboBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckComboBox_GetParentWindow").Call(obj)
	return ret
}

func CheckComboBox_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("CheckComboBox_SetParentWindow").Call(obj, value)
}

func CheckComboBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func CheckComboBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAction").Call(obj)
	return ret
}

func CheckComboBox_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetAction").Call(obj, value)
}

func CheckComboBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("CheckComboBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func CheckComboBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("CheckComboBox_SetBiDiMode").Call(obj, uintptr(value))
}

func CheckComboBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckComboBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("CheckComboBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckComboBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetClientHeight").Call(obj, uintptr(value))
}

func CheckComboBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckComboBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckComboBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckComboBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetClientWidth").Call(obj, uintptr(value))
}

func CheckComboBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("CheckComboBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func CheckComboBox_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("CheckComboBox_SetControlState").Call(obj, uintptr(value))
}

func CheckComboBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("CheckComboBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func CheckComboBox_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("CheckComboBox_SetControlStyle").Call(obj, uintptr(value))
}

func CheckComboBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckComboBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetParent").Call(obj)
	return ret
}

func CheckComboBox_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetParent").Call(obj, value)
}

func CheckComboBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetLeft").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetLeft").Call(obj, uintptr(value))
}

func CheckComboBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetTop").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetTop").Call(obj, uintptr(value))
}

func CheckComboBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetWidth").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetWidth").Call(obj, uintptr(value))
}

func CheckComboBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetHeight").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetHeight").Call(obj, uintptr(value))
}

func CheckComboBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckComboBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func CheckComboBox_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckComboBox_SetCursor").Call(obj, uintptr(value))
}

func CheckComboBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckComboBox_SetHint").Call(obj, GoStrToDStr(value))
}

func CheckComboBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func CheckComboBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckComboBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func CheckComboBox_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckComboBox_SetComponentIndex").Call(obj, uintptr(value))
}

func CheckComboBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetOwner").Call(obj)
	return ret
}

func CheckComboBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckComboBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckComboBox_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckComboBox_SetName").Call(obj, GoStrToDStr(value))
}

func CheckComboBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("CheckComboBox_GetTag").Call(obj)
	return int(ret)
}

func CheckComboBox_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("CheckComboBox_SetTag").Call(obj, uintptr(value))
}

func CheckComboBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func CheckComboBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetAnchorSideLeft").Call(obj, value)
}

func CheckComboBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchorSideTop").Call(obj)
	return ret
}

func CheckComboBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetAnchorSideTop").Call(obj, value)
}

func CheckComboBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchorSideRight").Call(obj)
	return ret
}

func CheckComboBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetAnchorSideRight").Call(obj, value)
}

func CheckComboBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func CheckComboBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetAnchorSideBottom").Call(obj, value)
}

func CheckComboBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetChildSizing").Call(obj)
	return ret
}

func CheckComboBox_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetChildSizing").Call(obj, value)
}

func CheckComboBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetBorderSpacing").Call(obj)
	return ret
}

func CheckComboBox_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetBorderSpacing").Call(obj, value)
}

func CheckComboBox_GetChecked(obj uintptr, AIndex int32) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetChecked").Call(obj, uintptr(AIndex))
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetChecked(obj uintptr, AIndex int32, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetChecked").Call(obj, uintptr(AIndex), GoBoolToDBool(value))
}

func CheckComboBox_GetItemEnabled(obj uintptr, AIndex int32) bool {
	ret, _, _ := getLazyProc("CheckComboBox_GetItemEnabled").Call(obj, uintptr(AIndex))
	return DBoolToGoBool(ret)
}

func CheckComboBox_SetItemEnabled(obj uintptr, AIndex int32, value bool) {
	_, _, _ = getLazyProc("CheckComboBox_SetItemEnabled").Call(obj, uintptr(AIndex), GoBoolToDBool(value))
}

func CheckComboBox_GetObjects(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetObjects").Call(obj, uintptr(AIndex))
	return ret
}

func CheckComboBox_SetObjects(obj uintptr, AIndex int32, value uintptr) {
	_, _, _ = getLazyProc("CheckComboBox_SetObjects").Call(obj, uintptr(AIndex), value)
}

func CheckComboBox_GetState(obj uintptr, AIndex int32) TCheckBoxState {
	ret, _, _ := getLazyProc("CheckComboBox_GetState").Call(obj, uintptr(AIndex))
	return TCheckBoxState(ret)
}

func CheckComboBox_SetState(obj uintptr, AIndex int32, value TCheckBoxState) {
	_, _, _ = getLazyProc("CheckComboBox_SetState").Call(obj, uintptr(AIndex), uintptr(value))
}

func CheckComboBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func CheckComboBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func CheckComboBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func CheckComboBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("CheckComboBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func CheckComboBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("CheckComboBox_StaticClassType").Call()
	return TClass(r)
}
