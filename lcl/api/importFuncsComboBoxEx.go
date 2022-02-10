package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TComboBoxEx ---------------------------

func ComboBoxEx_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_Create").Call(obj)
	return ret
}

func ComboBoxEx_Free(obj uintptr) {
	getLazyProc("ComboBoxEx_Free").Call(obj)
}

func ComboBoxEx_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_AddItem(obj uintptr, Item string, AObject uintptr) {
	getLazyProc("ComboBoxEx_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ComboBoxEx_Clear(obj uintptr) {
	getLazyProc("ComboBoxEx_Clear").Call(obj)
}

func ComboBoxEx_ClearSelection(obj uintptr) {
	getLazyProc("ComboBoxEx_ClearSelection").Call(obj)
}

func ComboBoxEx_DeleteSelected(obj uintptr) {
	getLazyProc("ComboBoxEx_DeleteSelected").Call(obj)
}

func ComboBoxEx_SelectAll(obj uintptr) {
	getLazyProc("ComboBoxEx_SelectAll").Call(obj)
}

func ComboBoxEx_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ComboBoxEx_DisableAlign(obj uintptr) {
	getLazyProc("ComboBoxEx_DisableAlign").Call(obj)
}

func ComboBoxEx_EnableAlign(obj uintptr) {
	getLazyProc("ComboBoxEx_EnableAlign").Call(obj)
}

func ComboBoxEx_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ComboBoxEx_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ComboBoxEx_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ComboBoxEx_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ComboBoxEx_InsertControl").Call(obj, AControl)
}

func ComboBoxEx_Invalidate(obj uintptr) {
	getLazyProc("ComboBoxEx_Invalidate").Call(obj)
}

func ComboBoxEx_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ComboBoxEx_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ComboBoxEx_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ComboBoxEx_RemoveControl").Call(obj, AControl)
}

func ComboBoxEx_Realign(obj uintptr) {
	getLazyProc("ComboBoxEx_Realign").Call(obj)
}

func ComboBoxEx_Repaint(obj uintptr) {
	getLazyProc("ComboBoxEx_Repaint").Call(obj)
}

func ComboBoxEx_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ComboBoxEx_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ComboBoxEx_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ComboBoxEx_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ComboBoxEx_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ComboBoxEx_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ComboBoxEx_SetFocus(obj uintptr) {
	getLazyProc("ComboBoxEx_SetFocus").Call(obj)
}

func ComboBoxEx_Update(obj uintptr) {
	getLazyProc("ComboBoxEx_Update").Call(obj)
}

func ComboBoxEx_BringToFront(obj uintptr) {
	getLazyProc("ComboBoxEx_BringToFront").Call(obj)
}

func ComboBoxEx_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ComboBoxEx_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBoxEx_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_Hide(obj uintptr) {
	getLazyProc("ComboBoxEx_Hide").Call(obj)
}

func ComboBoxEx_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ComboBoxEx_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ComboBoxEx_Refresh(obj uintptr) {
	getLazyProc("ComboBoxEx_Refresh").Call(obj)
}

func ComboBoxEx_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ComboBoxEx_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBoxEx_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_SendToBack(obj uintptr) {
	getLazyProc("ComboBoxEx_SendToBack").Call(obj)
}

func ComboBoxEx_Show(obj uintptr) {
	getLazyProc("ComboBoxEx_Show").Call(obj)
}

func ComboBoxEx_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ComboBoxEx_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ComboBoxEx_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetTextLen").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ComboBoxEx_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ComboBoxEx_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ComboBoxEx_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ComboBoxEx_Assign").Call(obj, Source)
}

func ComboBoxEx_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ComboBoxEx_ClassType").Call(obj)
	return TClass(ret)
}

func ComboBoxEx_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_InstanceSize").Call(obj)
	return int32(ret)
}

func ComboBoxEx_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ComboBoxEx_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetHashCode").Call(obj)
	return int32(ret)
}

func ComboBoxEx_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ComboBoxEx_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ComboBoxEx_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ComboBoxEx_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ComboBoxEx_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ComboBoxEx_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ComboBoxEx_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ComboBoxEx_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ComboBoxEx_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ComboBoxEx_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ComboBoxEx_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ComboBoxEx_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ComboBoxEx_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ComboBoxEx_AnchorClient").Call(obj, uintptr(ASpace))
}

func ComboBoxEx_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBoxEx_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ComboBoxEx_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ComboBoxEx_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ComboBoxEx_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ComboBoxEx_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ComboBoxEx_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ComboBoxEx_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAlign").Call(obj)
	return TAlign(ret)
}

func ComboBoxEx_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ComboBoxEx_SetAlign").Call(obj, uintptr(value))
}

func ComboBoxEx_GetAutoCompleteOptions(obj uintptr) TAutoCompleteOptions {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAutoCompleteOptions").Call(obj)
	return TAutoCompleteOptions(ret)
}

func ComboBoxEx_SetAutoCompleteOptions(obj uintptr, value TAutoCompleteOptions) {
	getLazyProc("ComboBoxEx_SetAutoCompleteOptions").Call(obj, uintptr(value))
}

func ComboBoxEx_GetItemsEx(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetItemsEx").Call(obj)
	return ret
}

func ComboBoxEx_SetItemsEx(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetItemsEx").Call(obj, value)
}

func ComboBoxEx_GetStyle(obj uintptr) TComboBoxExStyle {
	ret, _, _ := getLazyProc("ComboBoxEx_GetStyle").Call(obj)
	return TComboBoxExStyle(ret)
}

func ComboBoxEx_SetStyle(obj uintptr, value TComboBoxExStyle) {
	getLazyProc("ComboBoxEx_SetStyle").Call(obj, uintptr(value))
}

func ComboBoxEx_GetStyleEx(obj uintptr) TComboBoxExStyles {
	ret, _, _ := getLazyProc("ComboBoxEx_GetStyleEx").Call(obj)
	return TComboBoxExStyles(ret)
}

func ComboBoxEx_SetStyleEx(obj uintptr, value TComboBoxExStyles) {
	getLazyProc("ComboBoxEx_SetStyleEx").Call(obj, uintptr(value))
}

func ComboBoxEx_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAction").Call(obj)
	return ret
}

func ComboBoxEx_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetAction").Call(obj, value)
}

func ComboBoxEx_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ComboBoxEx_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ComboBoxEx_SetAnchors").Call(obj, uintptr(value))
}

func ComboBoxEx_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ComboBoxEx_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ComboBoxEx_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ComboBoxEx_SetBiDiMode").Call(obj, uintptr(value))
}

func ComboBoxEx_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ComboBoxEx_GetColor").Call(obj)
	return TColor(ret)
}

func ComboBoxEx_SetColor(obj uintptr, value TColor) {
	getLazyProc("ComboBoxEx_SetColor").Call(obj, uintptr(value))
}

func ComboBoxEx_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetConstraints").Call(obj)
	return ret
}

func ComboBoxEx_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetConstraints").Call(obj, value)
}

func ComboBoxEx_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ComboBoxEx_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ComboBoxEx_SetDragCursor").Call(obj, uintptr(value))
}

func ComboBoxEx_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ComboBoxEx_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ComboBoxEx_SetDragKind").Call(obj, uintptr(value))
}

func ComboBoxEx_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ComboBoxEx_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ComboBoxEx_SetDragMode").Call(obj, uintptr(value))
}

func ComboBoxEx_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetFont").Call(obj)
	return ret
}

func ComboBoxEx_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetFont").Call(obj, value)
}

func ComboBoxEx_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetItemHeight").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetItemHeight(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetItemHeight").Call(obj, uintptr(value))
}

func ComboBoxEx_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetMaxLength").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetMaxLength(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetMaxLength").Call(obj, uintptr(value))
}

func ComboBoxEx_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetPopupMenu").Call(obj)
	return ret
}

func ComboBoxEx_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetPopupMenu").Call(obj, value)
}

func ComboBoxEx_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ComboBoxEx_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ComboBoxEx_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ComboBoxEx_SetTabOrder").Call(obj, uintptr(value))
}

func ComboBoxEx_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_SetText(obj uintptr, value string) {
	getLazyProc("ComboBoxEx_SetText").Call(obj, GoStrToDStr(value))
}

func ComboBoxEx_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetVisible(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnDropDown(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnDropDown").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnSelect(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnSelect").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("ComboBoxEx_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func ComboBoxEx_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetImages").Call(obj)
	return ret
}

func ComboBoxEx_SetImages(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetImages").Call(obj, value)
}

func ComboBoxEx_GetDropDownCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDropDownCount").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetDropDownCount(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetDropDownCount").Call(obj, uintptr(value))
}

func ComboBoxEx_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_SetSelText(obj uintptr, value string) {
	getLazyProc("ComboBoxEx_SetSelText").Call(obj, GoStrToDStr(value))
}

func ComboBoxEx_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetCanvas").Call(obj)
	return ret
}

func ComboBoxEx_GetDroppedDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDroppedDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetDroppedDown(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetDroppedDown").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetItems").Call(obj)
	return ret
}

func ComboBoxEx_SetItems(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetItems").Call(obj, value)
}

func ComboBoxEx_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetSelLength").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetSelLength(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetSelLength").Call(obj, uintptr(value))
}

func ComboBoxEx_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetSelStart").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetSelStart(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetSelStart").Call(obj, uintptr(value))
}

func ComboBoxEx_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetItemIndex").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetItemIndex(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetItemIndex").Call(obj, uintptr(value))
}

func ComboBoxEx_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ComboBoxEx_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ComboBoxEx_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetBrush").Call(obj)
	return ret
}

func ComboBoxEx_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetControlCount").Call(obj)
	return int32(ret)
}

func ComboBoxEx_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ComboBoxEx_GetHandle").Call(obj)
	return HWND(ret)
}

func ComboBoxEx_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ComboBoxEx_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ComboBoxEx_SetParentWindow").Call(obj, uintptr(value))
}

func ComboBoxEx_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ComboBoxEx_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ComboBoxEx_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ComboBoxEx_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ComboBoxEx_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ComboBoxEx_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetClientHeight").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetClientHeight").Call(obj, uintptr(value))
}

func ComboBoxEx_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBoxEx_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ComboBoxEx_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBoxEx_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetClientWidth").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetClientWidth").Call(obj, uintptr(value))
}

func ComboBoxEx_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ComboBoxEx_GetControlState").Call(obj)
	return TControlState(ret)
}

func ComboBoxEx_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ComboBoxEx_SetControlState").Call(obj, uintptr(value))
}

func ComboBoxEx_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ComboBoxEx_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ComboBoxEx_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ComboBoxEx_SetControlStyle").Call(obj, uintptr(value))
}

func ComboBoxEx_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBoxEx_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBoxEx_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetParent").Call(obj)
	return ret
}

func ComboBoxEx_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetParent").Call(obj, value)
}

func ComboBoxEx_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetLeft").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetLeft(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetLeft").Call(obj, uintptr(value))
}

func ComboBoxEx_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetTop").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetTop(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetTop").Call(obj, uintptr(value))
}

func ComboBoxEx_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetWidth").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetWidth(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetWidth").Call(obj, uintptr(value))
}

func ComboBoxEx_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetHeight").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetHeight(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetHeight").Call(obj, uintptr(value))
}

func ComboBoxEx_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ComboBoxEx_GetCursor").Call(obj)
	return TCursor(ret)
}

func ComboBoxEx_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ComboBoxEx_SetCursor").Call(obj, uintptr(value))
}

func ComboBoxEx_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_SetHint(obj uintptr, value string) {
	getLazyProc("ComboBoxEx_SetHint").Call(obj, GoStrToDStr(value))
}

func ComboBoxEx_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetComponentCount").Call(obj)
	return int32(ret)
}

func ComboBoxEx_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBoxEx_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ComboBoxEx_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ComboBoxEx_SetComponentIndex").Call(obj, uintptr(value))
}

func ComboBoxEx_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetOwner").Call(obj)
	return ret
}

func ComboBoxEx_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBoxEx_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBoxEx_SetName(obj uintptr, value string) {
	getLazyProc("ComboBoxEx_SetName").Call(obj, GoStrToDStr(value))
}

func ComboBoxEx_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ComboBoxEx_GetTag").Call(obj)
	return int(ret)
}

func ComboBoxEx_SetTag(obj uintptr, value int) {
	getLazyProc("ComboBoxEx_SetTag").Call(obj, uintptr(value))
}

func ComboBoxEx_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchorSideLeft").Call(obj)
	return ret
}

func ComboBoxEx_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetAnchorSideLeft").Call(obj, value)
}

func ComboBoxEx_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchorSideTop").Call(obj)
	return ret
}

func ComboBoxEx_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetAnchorSideTop").Call(obj, value)
}

func ComboBoxEx_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchorSideRight").Call(obj)
	return ret
}

func ComboBoxEx_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetAnchorSideRight").Call(obj, value)
}

func ComboBoxEx_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchorSideBottom").Call(obj)
	return ret
}

func ComboBoxEx_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetAnchorSideBottom").Call(obj, value)
}

func ComboBoxEx_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetChildSizing").Call(obj)
	return ret
}

func ComboBoxEx_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetChildSizing").Call(obj, value)
}

func ComboBoxEx_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetBorderSpacing").Call(obj)
	return ret
}

func ComboBoxEx_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ComboBoxEx_SetBorderSpacing").Call(obj, value)
}

func ComboBoxEx_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ComboBoxEx_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ComboBoxEx_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ComboBoxEx_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ComboBoxEx_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ComboBoxEx_StaticClassType() TClass {
	r, _, _ := getLazyProc("ComboBoxEx_StaticClassType").Call()
	return TClass(r)
}
