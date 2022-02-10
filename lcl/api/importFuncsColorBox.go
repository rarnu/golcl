package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TColorBox ---------------------------

func ColorBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_Create").Call(obj)
	return ret
}

func ColorBox_Free(obj uintptr) {
	getLazyProc("ColorBox_Free").Call(obj)
}

func ColorBox_AddItem(obj uintptr, Item string, AObject uintptr) {
	getLazyProc("ColorBox_AddItem").Call(obj, GoStrToDStr(Item), AObject)
}

func ColorBox_Clear(obj uintptr) {
	getLazyProc("ColorBox_Clear").Call(obj)
}

func ColorBox_ClearSelection(obj uintptr) {
	getLazyProc("ColorBox_ClearSelection").Call(obj)
}

func ColorBox_DeleteSelected(obj uintptr) {
	getLazyProc("ColorBox_DeleteSelected").Call(obj)
}

func ColorBox_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SelectAll(obj uintptr) {
	getLazyProc("ColorBox_SelectAll").Call(obj)
}

func ColorBox_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ColorBox_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ColorBox_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ColorBox_DisableAlign(obj uintptr) {
	getLazyProc("ColorBox_DisableAlign").Call(obj)
}

func ColorBox_EnableAlign(obj uintptr) {
	getLazyProc("ColorBox_EnableAlign").Call(obj)
}

func ColorBox_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ColorBox_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ColorBox_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("ColorBox_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ColorBox_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("ColorBox_InsertControl").Call(obj, AControl)
}

func ColorBox_Invalidate(obj uintptr) {
	getLazyProc("ColorBox_Invalidate").Call(obj)
}

func ColorBox_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("ColorBox_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func ColorBox_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("ColorBox_RemoveControl").Call(obj, AControl)
}

func ColorBox_Realign(obj uintptr) {
	getLazyProc("ColorBox_Realign").Call(obj)
}

func ColorBox_Repaint(obj uintptr) {
	getLazyProc("ColorBox_Repaint").Call(obj)
}

func ColorBox_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("ColorBox_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ColorBox_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("ColorBox_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ColorBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ColorBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ColorBox_SetFocus(obj uintptr) {
	getLazyProc("ColorBox_SetFocus").Call(obj)
}

func ColorBox_Update(obj uintptr) {
	getLazyProc("ColorBox_Update").Call(obj)
}

func ColorBox_BringToFront(obj uintptr) {
	getLazyProc("ColorBox_BringToFront").Call(obj)
}

func ColorBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ColorBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_Hide(obj uintptr) {
	getLazyProc("ColorBox_Hide").Call(obj)
}

func ColorBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ColorBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ColorBox_Refresh(obj uintptr) {
	getLazyProc("ColorBox_Refresh").Call(obj)
}

func ColorBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ColorBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_SendToBack(obj uintptr) {
	getLazyProc("ColorBox_SendToBack").Call(obj)
}

func ColorBox_Show(obj uintptr) {
	getLazyProc("ColorBox_Show").Call(obj)
}

func ColorBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ColorBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ColorBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetTextLen").Call(obj)
	return int32(ret)
}

func ColorBox_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ColorBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ColorBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ColorBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ColorBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ColorBox_Assign").Call(obj, Source)
}

func ColorBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ColorBox_ClassType").Call(obj)
	return TClass(ret)
}

func ColorBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_InstanceSize").Call(obj)
	return int32(ret)
}

func ColorBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ColorBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ColorBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ColorBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetHashCode").Call(obj)
	return int32(ret)
}

func ColorBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ColorBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ColorBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ColorBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ColorBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ColorBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ColorBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ColorBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ColorBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ColorBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ColorBox_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ColorBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func ColorBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ColorBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ColorBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ColorBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ColorBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ColorBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ColorBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ColorBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func ColorBox_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ColorBox_SetAlign").Call(obj, uintptr(value))
}

func ColorBox_GetAutoComplete(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetAutoComplete").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetAutoComplete(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetAutoComplete").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetAutoDropDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetAutoDropDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetAutoDropDown(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetAutoDropDown").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDefaultColorColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorBox_GetDefaultColorColor").Call(obj)
	return TColor(ret)
}

func ColorBox_SetDefaultColorColor(obj uintptr, value TColor) {
	getLazyProc("ColorBox_SetDefaultColorColor").Call(obj, uintptr(value))
}

func ColorBox_GetNoneColorColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorBox_GetNoneColorColor").Call(obj)
	return TColor(ret)
}

func ColorBox_SetNoneColorColor(obj uintptr, value TColor) {
	getLazyProc("ColorBox_SetNoneColorColor").Call(obj, uintptr(value))
}

func ColorBox_GetSelected(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorBox_GetSelected").Call(obj)
	return TColor(ret)
}

func ColorBox_SetSelected(obj uintptr, value TColor) {
	getLazyProc("ColorBox_SetSelected").Call(obj, uintptr(value))
}

func ColorBox_GetStyle(obj uintptr) TColorBoxStyle {
	ret, _, _ := getLazyProc("ColorBox_GetStyle").Call(obj)
	return TColorBoxStyle(ret)
}

func ColorBox_SetStyle(obj uintptr, value TColorBoxStyle) {
	getLazyProc("ColorBox_SetStyle").Call(obj, uintptr(value))
}

func ColorBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ColorBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ColorBox_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ColorBox_SetAnchors").Call(obj, uintptr(value))
}

func ColorBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ColorBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ColorBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ColorBox_SetBiDiMode").Call(obj, uintptr(value))
}

func ColorBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorBox_GetColor").Call(obj)
	return TColor(ret)
}

func ColorBox_SetColor(obj uintptr, value TColor) {
	getLazyProc("ColorBox_SetColor").Call(obj, uintptr(value))
}

func ColorBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetConstraints").Call(obj)
	return ret
}

func ColorBox_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetConstraints").Call(obj, value)
}

func ColorBox_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDropDownCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetDropDownCount").Call(obj)
	return int32(ret)
}

func ColorBox_SetDropDownCount(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetDropDownCount").Call(obj, uintptr(value))
}

func ColorBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetFont").Call(obj)
	return ret
}

func ColorBox_SetFont(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetFont").Call(obj, value)
}

func ColorBox_GetItemHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetItemHeight").Call(obj)
	return int32(ret)
}

func ColorBox_SetItemHeight(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetItemHeight").Call(obj, uintptr(value))
}

func ColorBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetParentColor(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetParentFont(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetPopupMenu").Call(obj)
	return ret
}

func ColorBox_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetPopupMenu").Call(obj, value)
}

func ColorBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ColorBox_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ColorBox_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("ColorBox_SetTabOrder").Call(obj, uintptr(value))
}

func ColorBox_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetTabStop(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetVisible(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ColorBox_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnCloseUp(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnCloseUp").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnDropDown(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnDropDown").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_SetOnSelect(obj uintptr, fn interface{}) {
	getLazyProc("ColorBox_SetOnSelect").Call(obj, addEventToMap(obj, fn))
}

func ColorBox_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("ColorBox_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func ColorBox_SetCharCase(obj uintptr, value TEditCharCase) {
	getLazyProc("ColorBox_SetCharCase").Call(obj, uintptr(value))
}

func ColorBox_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_GetSelText").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_SetSelText(obj uintptr, value string) {
	getLazyProc("ColorBox_SetSelText").Call(obj, GoStrToDStr(value))
}

func ColorBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetCanvas").Call(obj)
	return ret
}

func ColorBox_GetDroppedDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetDroppedDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetDroppedDown(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetDroppedDown").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetItems").Call(obj)
	return ret
}

func ColorBox_SetItems(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetItems").Call(obj, value)
}

func ColorBox_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetSelLength").Call(obj)
	return int32(ret)
}

func ColorBox_SetSelLength(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetSelLength").Call(obj, uintptr(value))
}

func ColorBox_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetSelStart").Call(obj)
	return int32(ret)
}

func ColorBox_SetSelStart(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetSelStart").Call(obj, uintptr(value))
}

func ColorBox_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetItemIndex").Call(obj)
	return int32(ret)
}

func ColorBox_SetItemIndex(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetItemIndex").Call(obj, uintptr(value))
}

func ColorBox_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ColorBox_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetDockSite(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ColorBox_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetBrush").Call(obj)
	return ret
}

func ColorBox_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetControlCount").Call(obj)
	return int32(ret)
}

func ColorBox_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ColorBox_GetHandle").Call(obj)
	return HWND(ret)
}

func ColorBox_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ColorBox_GetParentWindow").Call(obj)
	return HWND(ret)
}

func ColorBox_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("ColorBox_SetParentWindow").Call(obj, uintptr(value))
}

func ColorBox_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("ColorBox_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAction").Call(obj)
	return ret
}

func ColorBox_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetAction").Call(obj, value)
}

func ColorBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ColorBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ColorBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func ColorBox_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetClientHeight").Call(obj, uintptr(value))
}

func ColorBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ColorBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ColorBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func ColorBox_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetClientWidth").Call(obj, uintptr(value))
}

func ColorBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ColorBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func ColorBox_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ColorBox_SetControlState").Call(obj, uintptr(value))
}

func ColorBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ColorBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ColorBox_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ColorBox_SetControlStyle").Call(obj, uintptr(value))
}

func ColorBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetParent").Call(obj)
	return ret
}

func ColorBox_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetParent").Call(obj, value)
}

func ColorBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetLeft").Call(obj)
	return int32(ret)
}

func ColorBox_SetLeft(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetLeft").Call(obj, uintptr(value))
}

func ColorBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetTop").Call(obj)
	return int32(ret)
}

func ColorBox_SetTop(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetTop").Call(obj, uintptr(value))
}

func ColorBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetWidth").Call(obj)
	return int32(ret)
}

func ColorBox_SetWidth(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetWidth").Call(obj, uintptr(value))
}

func ColorBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetHeight").Call(obj)
	return int32(ret)
}

func ColorBox_SetHeight(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetHeight").Call(obj, uintptr(value))
}

func ColorBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ColorBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func ColorBox_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ColorBox_SetCursor").Call(obj, uintptr(value))
}

func ColorBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_SetHint(obj uintptr, value string) {
	getLazyProc("ColorBox_SetHint").Call(obj, GoStrToDStr(value))
}

func ColorBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func ColorBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ColorBox_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ColorBox_SetComponentIndex").Call(obj, uintptr(value))
}

func ColorBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetOwner").Call(obj)
	return ret
}

func ColorBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorBox_SetName(obj uintptr, value string) {
	getLazyProc("ColorBox_SetName").Call(obj, GoStrToDStr(value))
}

func ColorBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ColorBox_GetTag").Call(obj)
	return int(ret)
}

func ColorBox_SetTag(obj uintptr, value int) {
	getLazyProc("ColorBox_SetTag").Call(obj, uintptr(value))
}

func ColorBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func ColorBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetAnchorSideLeft").Call(obj, value)
}

func ColorBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAnchorSideTop").Call(obj)
	return ret
}

func ColorBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetAnchorSideTop").Call(obj, value)
}

func ColorBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAnchorSideRight").Call(obj)
	return ret
}

func ColorBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetAnchorSideRight").Call(obj, value)
}

func ColorBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func ColorBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetAnchorSideBottom").Call(obj, value)
}

func ColorBox_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetChildSizing").Call(obj)
	return ret
}

func ColorBox_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetChildSizing").Call(obj, value)
}

func ColorBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetBorderSpacing").Call(obj)
	return ret
}

func ColorBox_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ColorBox_SetBorderSpacing").Call(obj, value)
}

func ColorBox_GetColors(obj uintptr, Index int32) TColor {
	ret, _, _ := getLazyProc("ColorBox_GetColors").Call(obj, uintptr(Index))
	return TColor(ret)
}

func ColorBox_GetColorNames(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("ColorBox_GetColorNames").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func ColorBox_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ColorBox_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ColorBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ColorBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ColorBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ColorBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("ColorBox_StaticClassType").Call()
	return TClass(r)
}
