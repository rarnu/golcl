package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TToolButton ---------------------------

func ToolButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_Create").Call(obj)
	return ret
}

func ToolButton_Free(obj uintptr) {
	getLazyProc("ToolButton_Free").Call(obj)
}

func ToolButton_CheckMenuDropdown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_CheckMenuDropdown").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_Click(obj uintptr) {
	getLazyProc("ToolButton_Click").Call(obj)
}

func ToolButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("ToolButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ToolButton_BringToFront(obj uintptr) {
	getLazyProc("ToolButton_BringToFront").Call(obj)
}

func ToolButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ToolButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToolButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_Hide(obj uintptr) {
	getLazyProc("ToolButton_Hide").Call(obj)
}

func ToolButton_Invalidate(obj uintptr) {
	getLazyProc("ToolButton_Invalidate").Call(obj)
}

func ToolButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ToolButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ToolButton_Refresh(obj uintptr) {
	getLazyProc("ToolButton_Refresh").Call(obj)
}

func ToolButton_Repaint(obj uintptr) {
	getLazyProc("ToolButton_Repaint").Call(obj)
}

func ToolButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("ToolButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToolButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_SendToBack(obj uintptr) {
	getLazyProc("ToolButton_SendToBack").Call(obj)
}

func ToolButton_Show(obj uintptr) {
	getLazyProc("ToolButton_Show").Call(obj)
}

func ToolButton_Update(obj uintptr) {
	getLazyProc("ToolButton_Update").Call(obj)
}

func ToolButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ToolButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ToolButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetTextLen").Call(obj)
	return int32(ret)
}

func ToolButton_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("ToolButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ToolButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ToolButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ToolButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_Assign(obj uintptr, Source uintptr) {
	getLazyProc("ToolButton_Assign").Call(obj, Source)
}

func ToolButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ToolButton_ClassType").Call(obj)
	return TClass(ret)
}

func ToolButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_InstanceSize").Call(obj)
	return int32(ret)
}

func ToolButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ToolButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ToolButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ToolButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetHashCode").Call(obj)
	return int32(ret)
}

func ToolButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ToolButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToolButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("ToolButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ToolButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ToolButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ToolButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("ToolButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ToolButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("ToolButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ToolButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("ToolButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ToolButton_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("ToolButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func ToolButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ToolButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ToolButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("ToolButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ToolButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("ToolButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ToolButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("ToolButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ToolButton_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetIndex").Call(obj)
	return int32(ret)
}

func ToolButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAction").Call(obj)
	return ret
}

func ToolButton_SetAction(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetAction").Call(obj, value)
}

func ToolButton_GetAllowAllUp(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetAllowAllUp").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetAllowAllUp(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetAllowAllUp").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_SetCaption(obj uintptr, value string) {
	getLazyProc("ToolButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func ToolButton_GetDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetDown(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetDown").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToolButton_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ToolButton_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("ToolButton_SetDragCursor").Call(obj, uintptr(value))
}

func ToolButton_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ToolButton_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ToolButton_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("ToolButton_SetDragKind").Call(obj, uintptr(value))
}

func ToolButton_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ToolButton_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ToolButton_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("ToolButton_SetDragMode").Call(obj, uintptr(value))
}

func ToolButton_GetDropdownMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetDropdownMenu").Call(obj)
	return ret
}

func ToolButton_SetDropdownMenu(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetDropdownMenu").Call(obj, value)
}

func ToolButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetEnabled(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetGrouped(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetGrouped").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetGrouped(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetGrouped").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetHeight").Call(obj)
	return int32(ret)
}

func ToolButton_SetHeight(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetHeight").Call(obj, uintptr(value))
}

func ToolButton_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetImageIndex").Call(obj)
	return int32(ret)
}

func ToolButton_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetImageIndex").Call(obj, uintptr(value))
}

func ToolButton_GetIndeterminate(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetIndeterminate").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetIndeterminate(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetIndeterminate").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetMarked(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetMarked").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetMarked(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetMarked").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetPopupMenu").Call(obj)
	return ret
}

func ToolButton_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetPopupMenu").Call(obj, value)
}

func ToolButton_GetWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetWrap(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetWrap").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetShowHint(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetStyle(obj uintptr) TToolButtonStyle {
	ret, _, _ := getLazyProc("ToolButton_GetStyle").Call(obj)
	return TToolButtonStyle(ret)
}

func ToolButton_SetStyle(obj uintptr, value TToolButtonStyle) {
	getLazyProc("ToolButton_SetStyle").Call(obj, uintptr(value))
}

func ToolButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_SetVisible(obj uintptr, value bool) {
	getLazyProc("ToolButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetWidth").Call(obj)
	return int32(ret)
}

func ToolButton_SetWidth(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetWidth").Call(obj, uintptr(value))
}

func ToolButton_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnEndDock(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_SetOnStartDock(obj uintptr, fn interface{}) {
	getLazyProc("ToolButton_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func ToolButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ToolButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func ToolButton_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("ToolButton_SetAlign").Call(obj, uintptr(value))
}

func ToolButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ToolButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ToolButton_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("ToolButton_SetAnchors").Call(obj, uintptr(value))
}

func ToolButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ToolButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ToolButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("ToolButton_SetBiDiMode").Call(obj, uintptr(value))
}

func ToolButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ToolButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("ToolButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func ToolButton_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetClientHeight").Call(obj, uintptr(value))
}

func ToolButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("ToolButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("ToolButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ToolButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func ToolButton_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetClientWidth").Call(obj, uintptr(value))
}

func ToolButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetConstraints").Call(obj)
	return ret
}

func ToolButton_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetConstraints").Call(obj, value)
}

func ToolButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ToolButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func ToolButton_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("ToolButton_SetControlState").Call(obj, uintptr(value))
}

func ToolButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ToolButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ToolButton_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("ToolButton_SetControlStyle").Call(obj, uintptr(value))
}

func ToolButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ToolButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ToolButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetParent").Call(obj)
	return ret
}

func ToolButton_SetParent(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetParent").Call(obj, value)
}

func ToolButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetLeft").Call(obj)
	return int32(ret)
}

func ToolButton_SetLeft(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetLeft").Call(obj, uintptr(value))
}

func ToolButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetTop").Call(obj)
	return int32(ret)
}

func ToolButton_SetTop(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetTop").Call(obj, uintptr(value))
}

func ToolButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ToolButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func ToolButton_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("ToolButton_SetCursor").Call(obj, uintptr(value))
}

func ToolButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_SetHint(obj uintptr, value string) {
	getLazyProc("ToolButton_SetHint").Call(obj, GoStrToDStr(value))
}

func ToolButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func ToolButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ToolButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ToolButton_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("ToolButton_SetComponentIndex").Call(obj, uintptr(value))
}

func ToolButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetOwner").Call(obj)
	return ret
}

func ToolButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ToolButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ToolButton_SetName(obj uintptr, value string) {
	getLazyProc("ToolButton_SetName").Call(obj, GoStrToDStr(value))
}

func ToolButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ToolButton_GetTag").Call(obj)
	return int(ret)
}

func ToolButton_SetTag(obj uintptr, value int) {
	getLazyProc("ToolButton_SetTag").Call(obj, uintptr(value))
}

func ToolButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func ToolButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetAnchorSideLeft").Call(obj, value)
}

func ToolButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAnchorSideTop").Call(obj)
	return ret
}

func ToolButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetAnchorSideTop").Call(obj, value)
}

func ToolButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAnchorSideRight").Call(obj)
	return ret
}

func ToolButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetAnchorSideRight").Call(obj, value)
}

func ToolButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func ToolButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetAnchorSideBottom").Call(obj, value)
}

func ToolButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetBorderSpacing").Call(obj)
	return ret
}

func ToolButton_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("ToolButton_SetBorderSpacing").Call(obj, value)
}

func ToolButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ToolButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ToolButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ToolButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("ToolButton_StaticClassType").Call()
	return TClass(r)
}
