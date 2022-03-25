package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TLinkLabel ---------------------------

func LinkLabel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_Create").Call(obj)
	return ret
}

func LinkLabel_Free(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Free").Call(obj)
}

func LinkLabel_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Invalidate").Call(obj)
}

func LinkLabel_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Repaint").Call(obj)
}

func LinkLabel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("LinkLabel_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func LinkLabel_Update(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Update").Call(obj)
}

func LinkLabel_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_BringToFront").Call(obj)
}

func LinkLabel_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LinkLabel_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LinkLabel_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_Hide(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Hide").Call(obj)
}

func LinkLabel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("LinkLabel_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func LinkLabel_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Refresh").Call(obj)
}

func LinkLabel_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LinkLabel_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LinkLabel_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SendToBack").Call(obj)
}

func LinkLabel_Show(obj uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Show").Call(obj)
}

func LinkLabel_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("LinkLabel_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func LinkLabel_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetTextLen").Call(obj)
	return int32(ret)
}

func LinkLabel_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("LinkLabel_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func LinkLabel_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func LinkLabel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("LinkLabel_Assign").Call(obj, Source)
}

func LinkLabel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("LinkLabel_ClassType").Call(obj)
	return TClass(ret)
}

func LinkLabel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_InstanceSize").Call(obj)
	return int32(ret)
}

func LinkLabel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("LinkLabel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func LinkLabel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetHashCode").Call(obj)
	return int32(ret)
}

func LinkLabel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("LinkLabel_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func LinkLabel_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("LinkLabel_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func LinkLabel_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("LinkLabel_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func LinkLabel_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("LinkLabel_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func LinkLabel_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("LinkLabel_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func LinkLabel_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("LinkLabel_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func LinkLabel_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("LinkLabel_AnchorClient").Call(obj, uintptr(ASpace))
}

func LinkLabel_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("LinkLabel_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func LinkLabel_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("LinkLabel_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func LinkLabel_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("LinkLabel_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func LinkLabel_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("LinkLabel_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func LinkLabel_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("LinkLabel_GetAlign").Call(obj)
	return TAlign(ret)
}

func LinkLabel_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("LinkLabel_SetAlign").Call(obj, uintptr(value))
}

func LinkLabel_GetAlignment(obj uintptr) TLinkAlignment {
	ret, _, _ := getLazyProc("LinkLabel_GetAlignment").Call(obj)
	return TLinkAlignment(ret)
}

func LinkLabel_SetAlignment(obj uintptr, value TLinkAlignment) {
	_, _, _ = getLazyProc("LinkLabel_SetAlignment").Call(obj, uintptr(value))
}

func LinkLabel_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func LinkLabel_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("LinkLabel_SetAnchors").Call(obj, uintptr(value))
}

func LinkLabel_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("LinkLabel_SetCaption").Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("LinkLabel_GetColor").Call(obj)
	return TColor(ret)
}

func LinkLabel_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("LinkLabel_SetColor").Call(obj, uintptr(value))
}

func LinkLabel_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetConstraints").Call(obj)
	return ret
}

func LinkLabel_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetConstraints").Call(obj, value)
}

func LinkLabel_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("LinkLabel_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func LinkLabel_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("LinkLabel_SetDragCursor").Call(obj, uintptr(value))
}

func LinkLabel_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("LinkLabel_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func LinkLabel_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("LinkLabel_SetDragKind").Call(obj, uintptr(value))
}

func LinkLabel_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("LinkLabel_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func LinkLabel_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("LinkLabel_SetDragMode").Call(obj, uintptr(value))
}

func LinkLabel_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetFont").Call(obj)
	return ret
}

func LinkLabel_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetFont").Call(obj, value)
}

func LinkLabel_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetPopupMenu").Call(obj)
	return ret
}

func LinkLabel_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetPopupMenu").Call(obj, value)
}

func LinkLabel_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("LinkLabel_SetVisible").Call(obj, GoBoolToDBool(value))
}

func LinkLabel_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_SetOnLinkClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("LinkLabel_SetOnLinkClick").Call(obj, addEventToMap(obj, fn))
}

func LinkLabel_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAction").Call(obj)
	return ret
}

func LinkLabel_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetAction").Call(obj, value)
}

func LinkLabel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("LinkLabel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func LinkLabel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("LinkLabel_SetBiDiMode").Call(obj, uintptr(value))
}

func LinkLabel_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("LinkLabel_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("LinkLabel_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func LinkLabel_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetClientHeight").Call(obj)
	return int32(ret)
}

func LinkLabel_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetClientHeight").Call(obj, uintptr(value))
}

func LinkLabel_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("LinkLabel_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("LinkLabel_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func LinkLabel_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetClientWidth").Call(obj)
	return int32(ret)
}

func LinkLabel_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetClientWidth").Call(obj, uintptr(value))
}

func LinkLabel_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("LinkLabel_GetControlState").Call(obj)
	return TControlState(ret)
}

func LinkLabel_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("LinkLabel_SetControlState").Call(obj, uintptr(value))
}

func LinkLabel_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("LinkLabel_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func LinkLabel_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("LinkLabel_SetControlStyle").Call(obj, uintptr(value))
}

func LinkLabel_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("LinkLabel_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func LinkLabel_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetParent").Call(obj)
	return ret
}

func LinkLabel_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetParent").Call(obj, value)
}

func LinkLabel_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetLeft").Call(obj)
	return int32(ret)
}

func LinkLabel_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetLeft").Call(obj, uintptr(value))
}

func LinkLabel_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetTop").Call(obj)
	return int32(ret)
}

func LinkLabel_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetTop").Call(obj, uintptr(value))
}

func LinkLabel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetWidth").Call(obj)
	return int32(ret)
}

func LinkLabel_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetWidth").Call(obj, uintptr(value))
}

func LinkLabel_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetHeight").Call(obj)
	return int32(ret)
}

func LinkLabel_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetHeight").Call(obj, uintptr(value))
}

func LinkLabel_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("LinkLabel_SetHint").Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetComponentCount").Call(obj)
	return int32(ret)
}

func LinkLabel_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("LinkLabel_GetComponentIndex").Call(obj)
	return int32(ret)
}

func LinkLabel_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("LinkLabel_SetComponentIndex").Call(obj, uintptr(value))
}

func LinkLabel_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetOwner").Call(obj)
	return ret
}

func LinkLabel_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("LinkLabel_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func LinkLabel_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("LinkLabel_SetName").Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("LinkLabel_GetTag").Call(obj)
	return int(ret)
}

func LinkLabel_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("LinkLabel_SetTag").Call(obj, uintptr(value))
}

func LinkLabel_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchorSideLeft").Call(obj)
	return ret
}

func LinkLabel_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetAnchorSideLeft").Call(obj, value)
}

func LinkLabel_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchorSideTop").Call(obj)
	return ret
}

func LinkLabel_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetAnchorSideTop").Call(obj, value)
}

func LinkLabel_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchorSideRight").Call(obj)
	return ret
}

func LinkLabel_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetAnchorSideRight").Call(obj, value)
}

func LinkLabel_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchorSideBottom").Call(obj)
	return ret
}

func LinkLabel_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetAnchorSideBottom").Call(obj, value)
}

func LinkLabel_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetBorderSpacing").Call(obj)
	return ret
}

func LinkLabel_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("LinkLabel_SetBorderSpacing").Call(obj, value)
}

func LinkLabel_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func LinkLabel_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("LinkLabel_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func LinkLabel_StaticClassType() TClass {
	r, _, _ := getLazyProc("LinkLabel_StaticClassType").Call()
	return TClass(r)
}
