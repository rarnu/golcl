package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TLabel ---------------------------

func Label_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_Create").Call(obj)
	return ret
}

func Label_Free(obj uintptr) {
	getLazyProc("Label_Free").Call(obj)
}

func Label_BringToFront(obj uintptr) {
	getLazyProc("Label_BringToFront").Call(obj)
}

func Label_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Label_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Label_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_Hide(obj uintptr) {
	getLazyProc("Label_Hide").Call(obj)
}

func Label_Invalidate(obj uintptr) {
	getLazyProc("Label_Invalidate").Call(obj)
}

func Label_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Label_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Label_Refresh(obj uintptr) {
	getLazyProc("Label_Refresh").Call(obj)
}

func Label_Repaint(obj uintptr) {
	getLazyProc("Label_Repaint").Call(obj)
}

func Label_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Label_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Label_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_SendToBack(obj uintptr) {
	getLazyProc("Label_SendToBack").Call(obj)
}

func Label_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("Label_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Label_Show(obj uintptr) {
	getLazyProc("Label_Show").Call(obj)
}

func Label_Update(obj uintptr) {
	getLazyProc("Label_Update").Call(obj)
}

func Label_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Label_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Label_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetTextLen").Call(obj)
	return int32(ret)
}

func Label_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("Label_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Label_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Label_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Label_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Label_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Label_Assign").Call(obj, Source)
}

func Label_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Label_ClassType").Call(obj)
	return TClass(ret)
}

func Label_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Label_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_InstanceSize").Call(obj)
	return int32(ret)
}

func Label_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Label_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Label_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Label_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetHashCode").Call(obj)
	return int32(ret)
}

func Label_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Label_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Label_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Label_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Label_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Label_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Label_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Label_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Label_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Label_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("Label_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Label_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("Label_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Label_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("Label_AnchorClient").Call(obj, uintptr(ASpace))
}

func Label_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Label_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Label_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("Label_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Label_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("Label_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Label_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("Label_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Label_GetOptimalFill(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetOptimalFill").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetOptimalFill(obj uintptr, value bool) {
	getLazyProc("Label_SetOptimalFill").Call(obj, GoBoolToDBool(value))
}

func Label_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Label_GetAlign").Call(obj)
	return TAlign(ret)
}

func Label_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("Label_SetAlign").Call(obj, uintptr(value))
}

func Label_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("Label_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func Label_SetAlignment(obj uintptr, value TAlignment) {
	getLazyProc("Label_SetAlignment").Call(obj, uintptr(value))
}

func Label_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Label_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Label_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("Label_SetAnchors").Call(obj, uintptr(value))
}

func Label_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetAutoSize(obj uintptr, value bool) {
	getLazyProc("Label_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Label_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Label_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Label_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("Label_SetBiDiMode").Call(obj, uintptr(value))
}

func Label_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func Label_SetCaption(obj uintptr, value string) {
	getLazyProc("Label_SetCaption").Call(obj, GoStrToDStr(value))
}

func Label_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Label_GetColor").Call(obj)
	return TColor(ret)
}

func Label_SetColor(obj uintptr, value TColor) {
	getLazyProc("Label_SetColor").Call(obj, uintptr(value))
}

func Label_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetConstraints").Call(obj)
	return ret
}

func Label_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("Label_SetConstraints").Call(obj, value)
}

func Label_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Label_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Label_SetDragCursor(obj uintptr, value TCursor) {
	getLazyProc("Label_SetDragCursor").Call(obj, uintptr(value))
}

func Label_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Label_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Label_SetDragKind(obj uintptr, value TDragKind) {
	getLazyProc("Label_SetDragKind").Call(obj, uintptr(value))
}

func Label_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Label_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Label_SetDragMode(obj uintptr, value TDragMode) {
	getLazyProc("Label_SetDragMode").Call(obj, uintptr(value))
}

func Label_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetEnabled(obj uintptr, value bool) {
	getLazyProc("Label_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Label_GetFocusControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetFocusControl").Call(obj)
	return ret
}

func Label_SetFocusControl(obj uintptr, value uintptr) {
	getLazyProc("Label_SetFocusControl").Call(obj, value)
}

func Label_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetFont").Call(obj)
	return ret
}

func Label_SetFont(obj uintptr, value uintptr) {
	getLazyProc("Label_SetFont").Call(obj, value)
}

func Label_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetParentColor(obj uintptr, value bool) {
	getLazyProc("Label_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Label_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetParentFont(obj uintptr, value bool) {
	getLazyProc("Label_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Label_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("Label_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Label_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetPopupMenu").Call(obj)
	return ret
}

func Label_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("Label_SetPopupMenu").Call(obj, value)
}

func Label_GetShowAccelChar(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetShowAccelChar").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetShowAccelChar(obj uintptr, value bool) {
	getLazyProc("Label_SetShowAccelChar").Call(obj, GoBoolToDBool(value))
}

func Label_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetShowHint(obj uintptr, value bool) {
	getLazyProc("Label_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Label_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetTransparent(obj uintptr, value bool) {
	getLazyProc("Label_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func Label_GetLayout(obj uintptr) TTextLayout {
	ret, _, _ := getLazyProc("Label_GetLayout").Call(obj)
	return TTextLayout(ret)
}

func Label_SetLayout(obj uintptr, value TTextLayout) {
	getLazyProc("Label_SetLayout").Call(obj, uintptr(value))
}

func Label_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetVisible(obj uintptr, value bool) {
	getLazyProc("Label_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Label_GetWordWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetWordWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_SetWordWrap(obj uintptr, value bool) {
	getLazyProc("Label_SetWordWrap").Call(obj, GoBoolToDBool(value))
}

func Label_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Label_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("Label_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Label_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetCanvas").Call(obj)
	return ret
}

func Label_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetAction").Call(obj)
	return ret
}

func Label_SetAction(obj uintptr, value uintptr) {
	getLazyProc("Label_SetAction").Call(obj, value)
}

func Label_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Label_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("Label_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Label_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetClientHeight").Call(obj)
	return int32(ret)
}

func Label_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("Label_SetClientHeight").Call(obj, uintptr(value))
}

func Label_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Label_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Label_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Label_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetClientWidth").Call(obj)
	return int32(ret)
}

func Label_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("Label_SetClientWidth").Call(obj, uintptr(value))
}

func Label_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Label_GetControlState").Call(obj)
	return TControlState(ret)
}

func Label_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("Label_SetControlState").Call(obj, uintptr(value))
}

func Label_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Label_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Label_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("Label_SetControlStyle").Call(obj, uintptr(value))
}

func Label_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Label_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Label_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetParent").Call(obj)
	return ret
}

func Label_SetParent(obj uintptr, value uintptr) {
	getLazyProc("Label_SetParent").Call(obj, value)
}

func Label_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetLeft").Call(obj)
	return int32(ret)
}

func Label_SetLeft(obj uintptr, value int32) {
	getLazyProc("Label_SetLeft").Call(obj, uintptr(value))
}

func Label_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetTop").Call(obj)
	return int32(ret)
}

func Label_SetTop(obj uintptr, value int32) {
	getLazyProc("Label_SetTop").Call(obj, uintptr(value))
}

func Label_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetWidth").Call(obj)
	return int32(ret)
}

func Label_SetWidth(obj uintptr, value int32) {
	getLazyProc("Label_SetWidth").Call(obj, uintptr(value))
}

func Label_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetHeight").Call(obj)
	return int32(ret)
}

func Label_SetHeight(obj uintptr, value int32) {
	getLazyProc("Label_SetHeight").Call(obj, uintptr(value))
}

func Label_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Label_GetCursor").Call(obj)
	return TCursor(ret)
}

func Label_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("Label_SetCursor").Call(obj, uintptr(value))
}

func Label_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Label_SetHint(obj uintptr, value string) {
	getLazyProc("Label_SetHint").Call(obj, GoStrToDStr(value))
}

func Label_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetComponentCount").Call(obj)
	return int32(ret)
}

func Label_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Label_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Label_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("Label_SetComponentIndex").Call(obj, uintptr(value))
}

func Label_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetOwner").Call(obj)
	return ret
}

func Label_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Label_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Label_SetName(obj uintptr, value string) {
	getLazyProc("Label_SetName").Call(obj, GoStrToDStr(value))
}

func Label_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Label_GetTag").Call(obj)
	return int(ret)
}

func Label_SetTag(obj uintptr, value int) {
	getLazyProc("Label_SetTag").Call(obj, uintptr(value))
}

func Label_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetAnchorSideLeft").Call(obj)
	return ret
}

func Label_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("Label_SetAnchorSideLeft").Call(obj, value)
}

func Label_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetAnchorSideTop").Call(obj)
	return ret
}

func Label_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("Label_SetAnchorSideTop").Call(obj, value)
}

func Label_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetAnchorSideRight").Call(obj)
	return ret
}

func Label_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("Label_SetAnchorSideRight").Call(obj, value)
}

func Label_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetAnchorSideBottom").Call(obj)
	return ret
}

func Label_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("Label_SetAnchorSideBottom").Call(obj, value)
}

func Label_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Label_GetBorderSpacing").Call(obj)
	return ret
}

func Label_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("Label_SetBorderSpacing").Call(obj, value)
}

func Label_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Label_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Label_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Label_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Label_StaticClassType() TClass {
	r, _, _ := getLazyProc("Label_StaticClassType").Call()
	return TClass(r)
}
