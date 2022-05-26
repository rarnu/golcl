package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TFloatSpinEdit ---------------------------

func FloatSpinEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_Create").Call(obj)
	return ret
}

func FloatSpinEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Free").Call(obj)
}

func FloatSpinEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Clear").Call(obj)
}

func FloatSpinEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_ClearSelection").Call(obj)
}

func FloatSpinEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_CopyToClipboard").Call(obj)
}

func FloatSpinEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_CutToClipboard").Call(obj)
}

func FloatSpinEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_PasteFromClipboard").Call(obj)
}

func FloatSpinEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Undo").Call(obj)
}

func FloatSpinEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SelectAll").Call(obj)
}

func FloatSpinEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_CanFocus").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_ContainsControl").Call(obj, Control)
	return GoBool(ret)
}

func FloatSpinEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), PascalBool(AllowDisabled), PascalBool(AllowWinControls), PascalBool(AllLevels))
	return ret
}

func FloatSpinEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_DisableAlign").Call(obj)
}

func FloatSpinEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_EnableAlign").Call(obj)
}

func FloatSpinEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_FindChildControl").Call(obj, PascalStr(ControlName))
	return ret
}

func FloatSpinEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_FlipChildren").Call(obj, PascalBool(AllLevels))
}

func FloatSpinEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_Focused").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_HandleAllocated").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_InsertControl").Call(obj, AControl)
}

func FloatSpinEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Invalidate").Call(obj)
}

func FloatSpinEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func FloatSpinEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_RemoveControl").Call(obj, AControl)
}

func FloatSpinEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Realign").Call(obj)
}

func FloatSpinEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Repaint").Call(obj)
}

func FloatSpinEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func FloatSpinEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func FloatSpinEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func FloatSpinEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetFocus").Call(obj)
}

func FloatSpinEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Update").Call(obj)
}

func FloatSpinEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_BringToFront").Call(obj)
}

func FloatSpinEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FloatSpinEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FloatSpinEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_Dragging").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_HasParent").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Hide").Call(obj)
}

func FloatSpinEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("FloatSpinEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func FloatSpinEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Refresh").Call(obj)
}

func FloatSpinEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FloatSpinEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FloatSpinEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SendToBack").Call(obj)
}

func FloatSpinEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Show").Call(obj)
}

func FloatSpinEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func FloatSpinEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTextBuf").Call(obj, PascalStr(Buffer))
}

func FloatSpinEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_FindComponent").Call(obj, PascalStr(AName))
	return ret
}

func FloatSpinEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetNamePath").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_Assign").Call(obj, Source)
}

func FloatSpinEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("FloatSpinEdit_ClassType").Call(obj)
	return TClass(ret)
}

func FloatSpinEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_ClassName").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return GoBool(ret)
}

func FloatSpinEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_Equals").Call(obj, Obj)
	return GoBool(ret)
}

func FloatSpinEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_ToString").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func FloatSpinEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func FloatSpinEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func FloatSpinEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func FloatSpinEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func FloatSpinEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func FloatSpinEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func FloatSpinEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func FloatSpinEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func FloatSpinEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func FloatSpinEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("FloatSpinEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func FloatSpinEdit_GetAutoSelected(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAutoSelected").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetAutoSelected(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAutoSelected").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func FloatSpinEdit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAlign").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func FloatSpinEdit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAlignment").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func FloatSpinEdit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAnchors").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAutoSelect").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetAutoSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAutoSelect").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAutoSize").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAutoSize").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetColor").Call(obj)
	return TColor(ret)
}

func FloatSpinEdit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetColor").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetConstraints").Call(obj)
	return ret
}

func FloatSpinEdit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetConstraints").Call(obj, value)
}

func FloatSpinEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetEnabled").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetEnabled").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetFont").Call(obj)
	return ret
}

func FloatSpinEdit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetFont").Call(obj, value)
}

func FloatSpinEdit_GetIncrement(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("FloatSpinEdit_GetIncrement").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SetIncrement(obj uintptr, value float64) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetIncrement").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FloatSpinEdit_GetMaxValue(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("FloatSpinEdit_GetMaxValue").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SetMaxValue(obj uintptr, value float64) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetMaxValue").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FloatSpinEdit_GetMinValue(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("FloatSpinEdit_GetMinValue").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SetMinValue(obj uintptr, value float64) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetMinValue").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FloatSpinEdit_SetOnChange(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnExit(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseWheel(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_SetOnResize(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func FloatSpinEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParentColor").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParentColor").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParentFont").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParentFont").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParentShowHint").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParentShowHint").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetPopupMenu").Call(obj)
	return ret
}

func FloatSpinEdit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetPopupMenu").Call(obj, value)
}

func FloatSpinEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetReadOnly").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetReadOnly").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetShowHint").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetShowHint").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTabStop").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTabStop").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func FloatSpinEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTabOrder").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetValue(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("FloatSpinEdit_GetValue").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SetValue(obj uintptr, value float64) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetValue").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FloatSpinEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetVisible").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetVisible").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetCanUndo").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetModified").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetModified").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetSelLength").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetSelStart").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetSelText").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetSelText").Call(obj, PascalStr(value))
}

func FloatSpinEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetText").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetText").Call(obj, PascalStr(value))
}

func FloatSpinEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTextHint").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTextHint").Call(obj, PascalStr(value))
}

func FloatSpinEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetDockSite").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetDockSite").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetDoubleBuffered").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetDoubleBuffered").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetMouseInClient").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetBrush").Call(obj)
	return ret
}

func FloatSpinEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetHandle").Call(obj)
	return ret
}

func FloatSpinEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParentDoubleBuffered").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParentDoubleBuffered").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParentWindow").Call(obj)
	return ret
}

func FloatSpinEdit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParentWindow").Call(obj, value)
}

func FloatSpinEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetShowing").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetUseDockManager").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetUseDockManager").Call(obj, PascalBool(value))
}

func FloatSpinEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAction").Call(obj)
	return ret
}

func FloatSpinEdit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAction").Call(obj, value)
}

func FloatSpinEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func FloatSpinEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("FloatSpinEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func FloatSpinEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetClientHeight").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("FloatSpinEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("FloatSpinEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func FloatSpinEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetClientWidth").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func FloatSpinEdit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetControlState").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func FloatSpinEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetControlStyle").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetFloating").Call(obj)
	return GoBool(ret)
}

func FloatSpinEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetParent").Call(obj)
	return ret
}

func FloatSpinEdit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetParent").Call(obj, value)
}

func FloatSpinEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetLeft").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetLeft").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTop").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTop").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetWidth").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetWidth").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetHeight").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetHeight").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func FloatSpinEdit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetCursor").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetHint").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetHint").Call(obj, PascalStr(value))
}

func FloatSpinEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func FloatSpinEdit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetOwner").Call(obj)
	return ret
}

func FloatSpinEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetName").Call(obj)
	return GoStr(ret)
}

func FloatSpinEdit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetName").Call(obj, PascalStr(value))
}

func FloatSpinEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetTag").Call(obj)
	return int(ret)
}

func FloatSpinEdit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetTag").Call(obj, uintptr(value))
}

func FloatSpinEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func FloatSpinEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAnchorSideLeft").Call(obj, value)
}

func FloatSpinEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func FloatSpinEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAnchorSideTop").Call(obj, value)
}

func FloatSpinEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func FloatSpinEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAnchorSideRight").Call(obj, value)
}

func FloatSpinEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func FloatSpinEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetAnchorSideBottom").Call(obj, value)
}

func FloatSpinEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetChildSizing").Call(obj)
	return ret
}

func FloatSpinEdit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetChildSizing").Call(obj, value)
}

func FloatSpinEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetBorderSpacing").Call(obj)
	return ret
}

func FloatSpinEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("FloatSpinEdit_SetBorderSpacing").Call(obj, value)
}

func FloatSpinEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func FloatSpinEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func FloatSpinEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func FloatSpinEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("FloatSpinEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func FloatSpinEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("FloatSpinEdit_StaticClassType").Call()
	return TClass(r)
}
