package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TComboBox ---------------------------

func ComboBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_Create").Call(obj)
	return ret
}

func ComboBox_Free(obj uintptr) {
	getLazyProc("ComboBox_Free").Call(obj)
}

func ComboBox_AddItem(obj uintptr, Item string, AObject uintptr) {
	getLazyProc("ComboBox_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ComboBox_Clear(obj uintptr) {
	getLazyProc("ComboBox_Clear").Call(obj)
}

func ComboBox_ClearSelection(obj uintptr) {
	getLazyProc("ComboBox_ClearSelection").Call(obj)
}

func ComboBox_DeleteSelected(obj uintptr) {
	getLazyProc("ComboBox_DeleteSelected").Call(obj)
}

func ComboBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SelectAll(obj uintptr) {
	getLazyProc("ComboBox_SelectAll").Call(obj)
}

func ComboBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ComboBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ComboBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ComboBox_DisableAlign(obj uintptr) {
	getLazyProc("ComboBox_DisableAlign").Call(obj)
}

func ComboBox_EnableAlign(obj uintptr) {
	getLazyProc("ComboBox_EnableAlign").Call(obj)
}

func ComboBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ComboBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ComboBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ComboBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ComboBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ComboBox_InsertControl").Call(obj, AControl)
}

func ComboBox_Invalidate(obj uintptr) {
	getLazyProc("ComboBox_Invalidate").Call(obj)
}

func ComboBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ComboBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ComboBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ComboBox_RemoveControl").Call(obj, AControl)
}

func ComboBox_Realign(obj uintptr) {
	getLazyProc("ComboBox_Realign").Call(obj)
}

func ComboBox_Repaint(obj uintptr) {
	getLazyProc("ComboBox_Repaint").Call(obj)
}

func ComboBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ComboBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ComboBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ComboBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ComboBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ComboBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ComboBox_SetFocus(obj uintptr) {
	getLazyProc("ComboBox_SetFocus").Call(obj)
}

func ComboBox_Update(obj uintptr) {
	getLazyProc("ComboBox_Update").Call(obj)
}

func ComboBox_BringToFront(obj uintptr) {
	getLazyProc("ComboBox_BringToFront").Call(obj)
}

func ComboBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ComboBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_Hide(obj uintptr) {
	getLazyProc("ComboBox_Hide").Call(obj)
}

func ComboBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ComboBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ComboBox_Refresh(obj uintptr) {
	getLazyProc("ComboBox_Refresh").Call(obj)
}

func ComboBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ComboBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_SendToBack(obj uintptr) {
	getLazyProc("ComboBox_SendToBack").Call(obj)
}

func ComboBox_Show(obj uintptr) {
	getLazyProc("ComboBox_Show").Call(obj)
}

func ComboBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ComboBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ComboBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ComboBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ComboBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ComboBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ComboBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ComboBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ComboBox_Assign").Call(obj, Source)
}

func ComboBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ComboBox_ClassType").Call(obj)
	return TClass(ret)
}

func ComboBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ComboBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ComboBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ComboBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ComboBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ComboBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ComboBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ComboBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ComboBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ComboBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ComboBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ComboBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ComboBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ComboBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ComboBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ComboBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ComboBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ComboBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ComboBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ComboBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ComboBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ComboBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ComboBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ComboBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ComboBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ComboBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ComboBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ComboBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ComboBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ComboBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ComboBox_SetAlign").Call(obj, uintptr(value))
}

func ComboBox_GetAutoComplete(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetAutoComplete").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetAutoComplete(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetAutoComplete").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAutoDropDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetAutoDropDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetAutoDropDown(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetAutoDropDown").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetStyle(obj uintptr) TComboBoxStyle {
	ret, _, _ := getLazyProc("ComboBox_GetStyle").Call(obj)
	return TComboBoxStyle(ret)
}

func ComboBox_SetStyle(obj uintptr, value TComboBoxStyle) {
	getLazyProc("ComboBox_SetStyle").Call(obj, uintptr(value))
}

func ComboBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ComboBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ComboBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ComboBox_SetAnchors").Call(obj, uintptr(value))
}

func ComboBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ComboBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ComboBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ComboBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ComboBox_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("ComboBox_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func ComboBox_SetCharCase(obj uintptr, value TEditCharCase) {
	getLazyProc("ComboBox_SetCharCase").Call(obj, uintptr(value))
}

func ComboBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ComboBox_GetColor").Call(obj)
	return TColor(ret)
}

func ComboBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("ComboBox_SetColor").Call(obj, uintptr(value))
}

func ComboBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetConstraints").Call(obj)
	return ret
}

func ComboBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetConstraints").Call(obj, value)
}

func ComboBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ComboBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ComboBox_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ComboBox_SetDragCursor").Call(obj, uintptr(value))
}

func ComboBox_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ComboBox_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ComboBox_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ComboBox_SetDragKind").Call(obj, uintptr(value))
}

func ComboBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ComboBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ComboBox_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ComboBox_SetDragMode").Call(obj, uintptr(value))
}

func ComboBox_GetDropDownCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetDropDownCount").Call(obj)
	return int32(ret)
}

func ComboBox_SetDropDownCount(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetDropDownCount").Call(obj, uintptr(value))
}

func ComboBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetFont").Call(obj)
	return ret
}

func ComboBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetFont").Call(obj, value)
}

func ComboBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func ComboBox_SetItemHeight(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetItemHeight").Call(obj, uintptr(value))
}

func ComboBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func ComboBox_SetItemIndex(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetItemIndex").Call(obj, uintptr(value))
}

func ComboBox_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetMaxLength").Call(obj)
	return int32(ret)
}

func ComboBox_SetMaxLength(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetMaxLength").Call(obj, uintptr(value))
}

func ComboBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetPopupMenu").Call(obj)
	return ret
}

func ComboBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetPopupMenu").Call(obj, value)
}

func ComboBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetSorted(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetSorted").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetSorted(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetSorted").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ComboBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ComboBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ComboBox_SetTabOrder").Call(obj, uintptr(value))
}

func ComboBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_SetText(obj uintptr, value string) {
	getLazyProc("ComboBox_SetText").Call(obj, GoStrToDStr(value))
}

func ComboBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ComboBox_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnCloseUp(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnCloseUp").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnDrawItem(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnDrawItem").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnDropDown(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnDropDown").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnMeasureItem(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnMeasureItem").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_SetOnSelect(obj uintptr, fn interface{}) {
	getLazyProc("ComboBox_SetOnSelect").Call(obj, addEventToMap(obj, fn))
}

func ComboBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetItems").Call(obj)
	return ret
}

func ComboBox_SetItems(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetItems").Call(obj, value)
}

func ComboBox_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_SetSelText(obj uintptr, value string) {
	getLazyProc("ComboBox_SetSelText").Call(obj, GoStrToDStr(value))
}

func ComboBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetCanvas").Call(obj)
	return ret
}

func ComboBox_GetDroppedDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetDroppedDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetDroppedDown(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetDroppedDown").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetSelLength").Call(obj)
	return int32(ret)
}

func ComboBox_SetSelLength(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetSelLength").Call(obj, uintptr(value))
}

func ComboBox_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetSelStart").Call(obj)
	return int32(ret)
}

func ComboBox_SetSelStart(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetSelStart").Call(obj, uintptr(value))
}

func ComboBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ComboBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ComboBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetBrush").Call(obj)
	return ret
}

func ComboBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ComboBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ComboBox_GetHandle").Call(obj)
	return HWND(ret)
}

func ComboBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ComboBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ComboBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ComboBox_SetParentWindow").Call(obj, uintptr(value))
}

func ComboBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ComboBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAction").Call(obj)
	return ret
}

func ComboBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetAction").Call(obj, value)
}

func ComboBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ComboBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ComboBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ComboBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ComboBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetClientHeight").Call(obj, uintptr(value))
}

func ComboBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ComboBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ComboBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ComboBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ComboBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetClientWidth").Call(obj, uintptr(value))
}

func ComboBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ComboBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ComboBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ComboBox_SetControlState").Call(obj, uintptr(value))
}

func ComboBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ComboBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ComboBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ComboBox_SetControlStyle").Call(obj, uintptr(value))
}

func ComboBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ComboBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ComboBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetParent").Call(obj)
	return ret
}

func ComboBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetParent").Call(obj, value)
}

func ComboBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetLeft").Call(obj)
	return int32(ret)
}

func ComboBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetLeft").Call(obj, uintptr(value))
}

func ComboBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetTop").Call(obj)
	return int32(ret)
}

func ComboBox_SetTop(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetTop").Call(obj, uintptr(value))
}

func ComboBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetWidth").Call(obj)
	return int32(ret)
}

func ComboBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetWidth").Call(obj, uintptr(value))
}

func ComboBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetHeight").Call(obj)
	return int32(ret)
}

func ComboBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetHeight").Call(obj, uintptr(value))
}

func ComboBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ComboBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ComboBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ComboBox_SetCursor").Call(obj, uintptr(value))
}

func ComboBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_SetHint(obj uintptr, value string) {
	getLazyProc("ComboBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ComboBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ComboBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ComboBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ComboBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ComboBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ComboBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetOwner").Call(obj)
	return ret
}

func ComboBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ComboBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ComboBox_SetName(obj uintptr, value string) {
	getLazyProc("ComboBox_SetName").Call(obj, GoStrToDStr(value))
}

func ComboBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ComboBox_GetTag").Call(obj)
	return int(ret)
}

func ComboBox_SetTag(obj uintptr, value int) {
	getLazyProc("ComboBox_SetTag").Call(obj, uintptr(value))
}

func ComboBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ComboBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetAnchorSideLeft").Call(obj, value)
}

func ComboBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ComboBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetAnchorSideTop").Call(obj, value)
}

func ComboBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ComboBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetAnchorSideRight").Call(obj, value)
}

func ComboBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ComboBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetAnchorSideBottom").Call(obj, value)
}

func ComboBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetChildSizing").Call(obj)
	return ret
}

func ComboBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetChildSizing").Call(obj, value)
}

func ComboBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetBorderSpacing").Call(obj)
	return ret
}

func ComboBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ComboBox_SetBorderSpacing").Call(obj, value)
}

func ComboBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ComboBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ComboBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ComboBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ComboBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ComboBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ComboBox_StaticClassType").Call()
	return TClass(r)
}
