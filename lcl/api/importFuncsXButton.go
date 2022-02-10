package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TXButton ---------------------------

func XButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_Create").Call(obj)
	return ret
}

func XButton_Free(obj uintptr) {
	getLazyProc("XButton_Free").Call(obj)
}

func XButton_BringToFront(obj uintptr) {
	getLazyProc("XButton_BringToFront").Call(obj)
}

func XButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("XButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("XButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_Hide(obj uintptr) {
	getLazyProc("XButton_Hide").Call(obj)
}

func XButton_Invalidate(obj uintptr) {
	getLazyProc("XButton_Invalidate").Call(obj)
}

func XButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("XButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func XButton_Refresh(obj uintptr) {
	getLazyProc("XButton_Refresh").Call(obj)
}

func XButton_Repaint(obj uintptr) {
	getLazyProc("XButton_Repaint").Call(obj)
}

func XButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("XButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("XButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_SendToBack(obj uintptr) {
	getLazyProc("XButton_SendToBack").Call(obj)
}

func XButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("XButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func XButton_Show(obj uintptr) {
	getLazyProc("XButton_Show").Call(obj)
}

func XButton_Update(obj uintptr) {
	getLazyProc("XButton_Update").Call(obj)
}

func XButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("XButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func XButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetTextLen").Call(obj)
	return int32(ret)
}

func XButton_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("XButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func XButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("XButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func XButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_Assign(obj uintptr, Source uintptr) {
	getLazyProc("XButton_Assign").Call(obj, Source)
}

func XButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("XButton_ClassType").Call(obj)
	return TClass(ret)
}

func XButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_InstanceSize").Call(obj)
	return int32(ret)
}

func XButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("XButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func XButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func XButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetHashCode").Call(obj)
	return int32(ret)
}

func XButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("XButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func XButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("XButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func XButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("XButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func XButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("XButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func XButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("XButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func XButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("XButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func XButton_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("XButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func XButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("XButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func XButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("XButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func XButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("XButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func XButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("XButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func XButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_SetCaption(obj uintptr, value string) {
	getLazyProc("XButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func XButton_GetShowCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetShowCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetShowCaption(obj uintptr, value bool) {
	getLazyProc("XButton_SetShowCaption").Call(obj, GoBoolToDBool(value))
}

func XButton_GetBackColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetBackColor").Call(obj)
	return TColor(ret)
}

func XButton_SetBackColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetBackColor").Call(obj, uintptr(value))
}

func XButton_GetHoverColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetHoverColor").Call(obj)
	return TColor(ret)
}

func XButton_SetHoverColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetHoverColor").Call(obj, uintptr(value))
}

func XButton_GetDownColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetDownColor").Call(obj)
	return TColor(ret)
}

func XButton_SetDownColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetDownColor").Call(obj, uintptr(value))
}

func XButton_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetBorderWidth").Call(obj)
	return int32(ret)
}

func XButton_SetBorderWidth(obj uintptr, value int32) {
	getLazyProc("XButton_SetBorderWidth").Call(obj, uintptr(value))
}

func XButton_GetBorderColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetBorderColor").Call(obj)
	return TColor(ret)
}

func XButton_SetBorderColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetBorderColor").Call(obj, uintptr(value))
}

func XButton_GetPicture(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetPicture").Call(obj)
	return ret
}

func XButton_SetPicture(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetPicture").Call(obj, value)
}

func XButton_GetDrawMode(obj uintptr) TDrawImageMode {
	ret, _, _ := getLazyProc("XButton_GetDrawMode").Call(obj)
	return TDrawImageMode(ret)
}

func XButton_SetDrawMode(obj uintptr, value TDrawImageMode) {
	getLazyProc("XButton_SetDrawMode").Call(obj, uintptr(value))
}

func XButton_GetNormalFontColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetNormalFontColor").Call(obj)
	return TColor(ret)
}

func XButton_SetNormalFontColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetNormalFontColor").Call(obj, uintptr(value))
}

func XButton_GetDownFontColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetDownFontColor").Call(obj)
	return TColor(ret)
}

func XButton_SetDownFontColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetDownFontColor").Call(obj, uintptr(value))
}

func XButton_GetHoverFontColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("XButton_GetHoverFontColor").Call(obj)
	return TColor(ret)
}

func XButton_SetHoverFontColor(obj uintptr, value TColor) {
	getLazyProc("XButton_SetHoverFontColor").Call(obj, uintptr(value))
}

func XButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAction").Call(obj)
	return ret
}

func XButton_SetAction(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetAction").Call(obj, value)
}

func XButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("XButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func XButton_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("XButton_SetAlign").Call(obj, uintptr(value))
}

func XButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("XButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func XButton_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("XButton_SetAnchors").Call(obj, uintptr(value))
}

func XButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("XButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func XButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("XButton_SetBiDiMode").Call(obj, uintptr(value))
}

func XButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetConstraints").Call(obj)
	return ret
}

func XButton_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetConstraints").Call(obj, value)
}

func XButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetEnabled(obj uintptr, value bool) {
	getLazyProc("XButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func XButton_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetFont").Call(obj)
	return ret
}

func XButton_SetFont(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetFont").Call(obj, value)
}

func XButton_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetParentFont(obj uintptr, value bool) {
	getLazyProc("XButton_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func XButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("XButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func XButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetPopupMenu").Call(obj)
	return ret
}

func XButton_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetPopupMenu").Call(obj, value)
}

func XButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetShowHint(obj uintptr, value bool) {
	getLazyProc("XButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func XButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_SetVisible(obj uintptr, value bool) {
	getLazyProc("XButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func XButton_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnDblClick(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func XButton_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("XButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func XButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("XButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("XButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func XButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func XButton_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("XButton_SetClientHeight").Call(obj, uintptr(value))
}

func XButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("XButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("XButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func XButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func XButton_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("XButton_SetClientWidth").Call(obj, uintptr(value))
}

func XButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("XButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func XButton_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("XButton_SetControlState").Call(obj, uintptr(value))
}

func XButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("XButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func XButton_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("XButton_SetControlStyle").Call(obj, uintptr(value))
}

func XButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("XButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func XButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetParent").Call(obj)
	return ret
}

func XButton_SetParent(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetParent").Call(obj, value)
}

func XButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetLeft").Call(obj)
	return int32(ret)
}

func XButton_SetLeft(obj uintptr, value int32) {
	getLazyProc("XButton_SetLeft").Call(obj, uintptr(value))
}

func XButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetTop").Call(obj)
	return int32(ret)
}

func XButton_SetTop(obj uintptr, value int32) {
	getLazyProc("XButton_SetTop").Call(obj, uintptr(value))
}

func XButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetWidth").Call(obj)
	return int32(ret)
}

func XButton_SetWidth(obj uintptr, value int32) {
	getLazyProc("XButton_SetWidth").Call(obj, uintptr(value))
}

func XButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetHeight").Call(obj)
	return int32(ret)
}

func XButton_SetHeight(obj uintptr, value int32) {
	getLazyProc("XButton_SetHeight").Call(obj, uintptr(value))
}

func XButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("XButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func XButton_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("XButton_SetCursor").Call(obj, uintptr(value))
}

func XButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_SetHint(obj uintptr, value string) {
	getLazyProc("XButton_SetHint").Call(obj, GoStrToDStr(value))
}

func XButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func XButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("XButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func XButton_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("XButton_SetComponentIndex").Call(obj, uintptr(value))
}

func XButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetOwner").Call(obj)
	return ret
}

func XButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("XButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func XButton_SetName(obj uintptr, value string) {
	getLazyProc("XButton_SetName").Call(obj, GoStrToDStr(value))
}

func XButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("XButton_GetTag").Call(obj)
	return int(ret)
}

func XButton_SetTag(obj uintptr, value int) {
	getLazyProc("XButton_SetTag").Call(obj, uintptr(value))
}

func XButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func XButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetAnchorSideLeft").Call(obj, value)
}

func XButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAnchorSideTop").Call(obj)
	return ret
}

func XButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetAnchorSideTop").Call(obj, value)
}

func XButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAnchorSideRight").Call(obj)
	return ret
}

func XButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetAnchorSideRight").Call(obj, value)
}

func XButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func XButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetAnchorSideBottom").Call(obj, value)
}

func XButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("XButton_GetBorderSpacing").Call(obj)
	return ret
}

func XButton_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("XButton_SetBorderSpacing").Call(obj, value)
}

func XButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("XButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func XButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("XButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func XButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("XButton_StaticClassType").Call()
	return TClass(r)
}
