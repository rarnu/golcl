package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TPaintBox ---------------------------

func PaintBox_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_Create").Call(obj)
	return ret
}

func PaintBox_Free(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Free").Call(obj)
}

func PaintBox_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_BringToFront").Call(obj)
}

func PaintBox_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PaintBox_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PaintBox_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_Hide(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Hide").Call(obj)
}

func PaintBox_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Invalidate").Call(obj)
}

func PaintBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("PaintBox_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func PaintBox_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Refresh").Call(obj)
}

func PaintBox_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Repaint").Call(obj)
}

func PaintBox_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PaintBox_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PaintBox_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_SendToBack").Call(obj)
}

func PaintBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("PaintBox_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func PaintBox_Show(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Show").Call(obj)
}

func PaintBox_Update(obj uintptr) {
	_, _, _ = getLazyProc("PaintBox_Update").Call(obj)
}

func PaintBox_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("PaintBox_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func PaintBox_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetTextLen").Call(obj)
	return int32(ret)
}

func PaintBox_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("PaintBox_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func PaintBox_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("PaintBox_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func PaintBox_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("PaintBox_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func PaintBox_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("PaintBox_Assign").Call(obj, Source)
}

func PaintBox_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("PaintBox_ClassType").Call(obj)
	return TClass(ret)
}

func PaintBox_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("PaintBox_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func PaintBox_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_InstanceSize").Call(obj)
	return int32(ret)
}

func PaintBox_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("PaintBox_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func PaintBox_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func PaintBox_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetHashCode").Call(obj)
	return int32(ret)
}

func PaintBox_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("PaintBox_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func PaintBox_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("PaintBox_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func PaintBox_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("PaintBox_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func PaintBox_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("PaintBox_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func PaintBox_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("PaintBox_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func PaintBox_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("PaintBox_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func PaintBox_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("PaintBox_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func PaintBox_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("PaintBox_AnchorClient").Call(obj, uintptr(ASpace))
}

func PaintBox_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("PaintBox_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func PaintBox_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("PaintBox_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func PaintBox_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("PaintBox_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func PaintBox_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("PaintBox_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func PaintBox_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetCanvas").Call(obj)
	return ret
}

func PaintBox_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("PaintBox_GetAlign").Call(obj)
	return TAlign(ret)
}

func PaintBox_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("PaintBox_SetAlign").Call(obj, uintptr(value))
}

func PaintBox_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("PaintBox_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func PaintBox_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("PaintBox_SetAnchors").Call(obj, uintptr(value))
}

func PaintBox_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("PaintBox_GetColor").Call(obj)
	return TColor(ret)
}

func PaintBox_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("PaintBox_SetColor").Call(obj, uintptr(value))
}

func PaintBox_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetConstraints").Call(obj)
	return ret
}

func PaintBox_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetConstraints").Call(obj, value)
}

func PaintBox_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("PaintBox_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func PaintBox_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("PaintBox_SetDragCursor").Call(obj, uintptr(value))
}

func PaintBox_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("PaintBox_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func PaintBox_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("PaintBox_SetDragMode").Call(obj, uintptr(value))
}

func PaintBox_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetFont").Call(obj)
	return ret
}

func PaintBox_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetFont").Call(obj, value)
}

func PaintBox_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetPopupMenu").Call(obj)
	return ret
}

func PaintBox_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetPopupMenu").Call(obj, value)
}

func PaintBox_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PaintBox_SetVisible").Call(obj, GoBoolToDBool(value))
}

func PaintBox_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_SetOnPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("PaintBox_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func PaintBox_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAction").Call(obj)
	return ret
}

func PaintBox_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetAction").Call(obj, value)
}

func PaintBox_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("PaintBox_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func PaintBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("PaintBox_SetBiDiMode").Call(obj, uintptr(value))
}

func PaintBox_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("PaintBox_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("PaintBox_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PaintBox_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetClientHeight").Call(obj)
	return int32(ret)
}

func PaintBox_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetClientHeight").Call(obj, uintptr(value))
}

func PaintBox_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("PaintBox_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("PaintBox_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func PaintBox_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetClientWidth").Call(obj)
	return int32(ret)
}

func PaintBox_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetClientWidth").Call(obj, uintptr(value))
}

func PaintBox_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("PaintBox_GetControlState").Call(obj)
	return TControlState(ret)
}

func PaintBox_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("PaintBox_SetControlState").Call(obj, uintptr(value))
}

func PaintBox_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("PaintBox_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func PaintBox_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("PaintBox_SetControlStyle").Call(obj, uintptr(value))
}

func PaintBox_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("PaintBox_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func PaintBox_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetParent").Call(obj)
	return ret
}

func PaintBox_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetParent").Call(obj, value)
}

func PaintBox_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetLeft").Call(obj)
	return int32(ret)
}

func PaintBox_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetLeft").Call(obj, uintptr(value))
}

func PaintBox_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetTop").Call(obj)
	return int32(ret)
}

func PaintBox_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetTop").Call(obj, uintptr(value))
}

func PaintBox_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetWidth").Call(obj)
	return int32(ret)
}

func PaintBox_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetWidth").Call(obj, uintptr(value))
}

func PaintBox_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetHeight").Call(obj)
	return int32(ret)
}

func PaintBox_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetHeight").Call(obj, uintptr(value))
}

func PaintBox_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("PaintBox_GetCursor").Call(obj)
	return TCursor(ret)
}

func PaintBox_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("PaintBox_SetCursor").Call(obj, uintptr(value))
}

func PaintBox_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("PaintBox_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func PaintBox_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("PaintBox_SetHint").Call(obj, GoStrToDStr(value))
}

func PaintBox_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetComponentCount").Call(obj)
	return int32(ret)
}

func PaintBox_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PaintBox_GetComponentIndex").Call(obj)
	return int32(ret)
}

func PaintBox_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PaintBox_SetComponentIndex").Call(obj, uintptr(value))
}

func PaintBox_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetOwner").Call(obj)
	return ret
}

func PaintBox_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("PaintBox_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func PaintBox_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("PaintBox_SetName").Call(obj, GoStrToDStr(value))
}

func PaintBox_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("PaintBox_GetTag").Call(obj)
	return int(ret)
}

func PaintBox_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("PaintBox_SetTag").Call(obj, uintptr(value))
}

func PaintBox_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAnchorSideLeft").Call(obj)
	return ret
}

func PaintBox_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetAnchorSideLeft").Call(obj, value)
}

func PaintBox_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAnchorSideTop").Call(obj)
	return ret
}

func PaintBox_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetAnchorSideTop").Call(obj, value)
}

func PaintBox_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAnchorSideRight").Call(obj)
	return ret
}

func PaintBox_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetAnchorSideRight").Call(obj, value)
}

func PaintBox_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAnchorSideBottom").Call(obj)
	return ret
}

func PaintBox_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetAnchorSideBottom").Call(obj, value)
}

func PaintBox_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetBorderSpacing").Call(obj)
	return ret
}

func PaintBox_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("PaintBox_SetBorderSpacing").Call(obj, value)
}

func PaintBox_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func PaintBox_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("PaintBox_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func PaintBox_StaticClassType() TClass {
	r, _, _ := getLazyProc("PaintBox_StaticClassType").Call()
	return TClass(r)
}
