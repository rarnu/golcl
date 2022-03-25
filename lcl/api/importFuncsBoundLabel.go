package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TBoundLabel ---------------------------

func BoundLabel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_Create").Call(obj)
	return ret
}

func BoundLabel_Free(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Free").Call(obj)
}

func BoundLabel_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_BringToFront").Call(obj)
}

func BoundLabel_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BoundLabel_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BoundLabel_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_Hide(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Hide").Call(obj)
}

func BoundLabel_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Invalidate").Call(obj)
}

func BoundLabel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("BoundLabel_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func BoundLabel_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Refresh").Call(obj)
}

func BoundLabel_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Repaint").Call(obj)
}

func BoundLabel_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BoundLabel_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BoundLabel_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SendToBack").Call(obj)
}

func BoundLabel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("BoundLabel_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func BoundLabel_Show(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Show").Call(obj)
}

func BoundLabel_Update(obj uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Update").Call(obj)
}

func BoundLabel_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("BoundLabel_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func BoundLabel_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetTextLen").Call(obj)
	return int32(ret)
}

func BoundLabel_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("BoundLabel_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func BoundLabel_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func BoundLabel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("BoundLabel_Assign").Call(obj, Source)
}

func BoundLabel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("BoundLabel_ClassType").Call(obj)
	return TClass(ret)
}

func BoundLabel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_InstanceSize").Call(obj)
	return int32(ret)
}

func BoundLabel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("BoundLabel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func BoundLabel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetHashCode").Call(obj)
	return int32(ret)
}

func BoundLabel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("BoundLabel_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func BoundLabel_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("BoundLabel_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func BoundLabel_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("BoundLabel_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func BoundLabel_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("BoundLabel_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func BoundLabel_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("BoundLabel_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func BoundLabel_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("BoundLabel_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func BoundLabel_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("BoundLabel_AnchorClient").Call(obj, uintptr(ASpace))
}

func BoundLabel_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BoundLabel_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BoundLabel_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("BoundLabel_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func BoundLabel_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("BoundLabel_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func BoundLabel_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("BoundLabel_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func BoundLabel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("BoundLabel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func BoundLabel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("BoundLabel_SetBiDiMode").Call(obj, uintptr(value))
}

func BoundLabel_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("BoundLabel_SetCaption").Call(obj, GoStrToDStr(value))
}

func BoundLabel_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("BoundLabel_GetColor").Call(obj)
	return TColor(ret)
}

func BoundLabel_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("BoundLabel_SetColor").Call(obj, uintptr(value))
}

func BoundLabel_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("BoundLabel_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func BoundLabel_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("BoundLabel_SetDragCursor").Call(obj, uintptr(value))
}

func BoundLabel_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("BoundLabel_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func BoundLabel_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("BoundLabel_SetDragMode").Call(obj, uintptr(value))
}

func BoundLabel_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetFont").Call(obj)
	return ret
}

func BoundLabel_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetFont").Call(obj, value)
}

func BoundLabel_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetHeight").Call(obj)
	return int32(ret)
}

func BoundLabel_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("BoundLabel_SetHeight").Call(obj, uintptr(value))
}

func BoundLabel_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetLeft").Call(obj)
	return int32(ret)
}

func BoundLabel_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetPopupMenu").Call(obj)
	return ret
}

func BoundLabel_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetPopupMenu").Call(obj, value)
}

func BoundLabel_GetShowAccelChar(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetShowAccelChar").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetShowAccelChar(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetShowAccelChar").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetTop").Call(obj)
	return int32(ret)
}

func BoundLabel_GetLayout(obj uintptr) TTextLayout {
	ret, _, _ := getLazyProc("BoundLabel_GetLayout").Call(obj)
	return TTextLayout(ret)
}

func BoundLabel_SetLayout(obj uintptr, value TTextLayout) {
	_, _, _ = getLazyProc("BoundLabel_SetLayout").Call(obj, uintptr(value))
}

func BoundLabel_GetWordWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetWordWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetWordWrap(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetWordWrap").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetWidth").Call(obj)
	return int32(ret)
}

func BoundLabel_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("BoundLabel_SetWidth").Call(obj, uintptr(value))
}

func BoundLabel_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("BoundLabel_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func BoundLabel_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetCanvas").Call(obj)
	return ret
}

func BoundLabel_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAction").Call(obj)
	return ret
}

func BoundLabel_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetAction").Call(obj, value)
}

func BoundLabel_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("BoundLabel_GetAlign").Call(obj)
	return TAlign(ret)
}

func BoundLabel_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("BoundLabel_SetAlign").Call(obj, uintptr(value))
}

func BoundLabel_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func BoundLabel_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("BoundLabel_SetAnchors").Call(obj, uintptr(value))
}

func BoundLabel_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("BoundLabel_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("BoundLabel_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func BoundLabel_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetClientHeight").Call(obj)
	return int32(ret)
}

func BoundLabel_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("BoundLabel_SetClientHeight").Call(obj, uintptr(value))
}

func BoundLabel_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("BoundLabel_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("BoundLabel_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BoundLabel_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetClientWidth").Call(obj)
	return int32(ret)
}

func BoundLabel_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("BoundLabel_SetClientWidth").Call(obj, uintptr(value))
}

func BoundLabel_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetConstraints").Call(obj)
	return ret
}

func BoundLabel_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetConstraints").Call(obj, value)
}

func BoundLabel_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("BoundLabel_GetControlState").Call(obj)
	return TControlState(ret)
}

func BoundLabel_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("BoundLabel_SetControlState").Call(obj, uintptr(value))
}

func BoundLabel_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("BoundLabel_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func BoundLabel_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("BoundLabel_SetControlStyle").Call(obj, uintptr(value))
}

func BoundLabel_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("BoundLabel_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func BoundLabel_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("BoundLabel_SetVisible").Call(obj, GoBoolToDBool(value))
}

func BoundLabel_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetParent").Call(obj)
	return ret
}

func BoundLabel_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetParent").Call(obj, value)
}

func BoundLabel_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("BoundLabel_GetCursor").Call(obj)
	return TCursor(ret)
}

func BoundLabel_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("BoundLabel_SetCursor").Call(obj, uintptr(value))
}

func BoundLabel_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("BoundLabel_SetHint").Call(obj, GoStrToDStr(value))
}

func BoundLabel_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetComponentCount").Call(obj)
	return int32(ret)
}

func BoundLabel_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BoundLabel_GetComponentIndex").Call(obj)
	return int32(ret)
}

func BoundLabel_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("BoundLabel_SetComponentIndex").Call(obj, uintptr(value))
}

func BoundLabel_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetOwner").Call(obj)
	return ret
}

func BoundLabel_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("BoundLabel_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func BoundLabel_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("BoundLabel_SetName").Call(obj, GoStrToDStr(value))
}

func BoundLabel_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("BoundLabel_GetTag").Call(obj)
	return int(ret)
}

func BoundLabel_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("BoundLabel_SetTag").Call(obj, uintptr(value))
}

func BoundLabel_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchorSideLeft").Call(obj)
	return ret
}

func BoundLabel_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetAnchorSideLeft").Call(obj, value)
}

func BoundLabel_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchorSideTop").Call(obj)
	return ret
}

func BoundLabel_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetAnchorSideTop").Call(obj, value)
}

func BoundLabel_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchorSideRight").Call(obj)
	return ret
}

func BoundLabel_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetAnchorSideRight").Call(obj, value)
}

func BoundLabel_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchorSideBottom").Call(obj)
	return ret
}

func BoundLabel_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetAnchorSideBottom").Call(obj, value)
}

func BoundLabel_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetBorderSpacing").Call(obj)
	return ret
}

func BoundLabel_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("BoundLabel_SetBorderSpacing").Call(obj, value)
}

func BoundLabel_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func BoundLabel_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("BoundLabel_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func BoundLabel_StaticClassType() TClass {
	r, _, _ := getLazyProc("BoundLabel_StaticClassType").Call()
	return TClass(r)
}
