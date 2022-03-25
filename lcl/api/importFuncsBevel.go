package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TBevel ---------------------------

func Bevel_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_Create").Call(obj)
	return ret
}

func Bevel_Free(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Free").Call(obj)
}

func Bevel_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_BringToFront").Call(obj)
}

func Bevel_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Bevel_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Bevel_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Hide").Call(obj)
}

func Bevel_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Invalidate").Call(obj)
}

func Bevel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Bevel_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Bevel_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Refresh").Call(obj)
}

func Bevel_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Repaint").Call(obj)
}

func Bevel_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Bevel_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Bevel_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_SendToBack").Call(obj)
}

func Bevel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Bevel_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Bevel_Show(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Show").Call(obj)
}

func Bevel_Update(obj uintptr) {
	_, _, _ = getLazyProc("Bevel_Update").Call(obj)
}

func Bevel_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Bevel_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Bevel_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetTextLen").Call(obj)
	return int32(ret)
}

func Bevel_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Bevel_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Bevel_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Bevel_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Bevel_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Bevel_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Bevel_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Bevel_Assign").Call(obj, Source)
}

func Bevel_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Bevel_ClassType").Call(obj)
	return TClass(ret)
}

func Bevel_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Bevel_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Bevel_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_InstanceSize").Call(obj)
	return int32(ret)
}

func Bevel_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Bevel_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Bevel_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Bevel_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetHashCode").Call(obj)
	return int32(ret)
}

func Bevel_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Bevel_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Bevel_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Bevel_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Bevel_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Bevel_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Bevel_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Bevel_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Bevel_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Bevel_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Bevel_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Bevel_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Bevel_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Bevel_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Bevel_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Bevel_AnchorClient").Call(obj, uintptr(ASpace))
}

func Bevel_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Bevel_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Bevel_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Bevel_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Bevel_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Bevel_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Bevel_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Bevel_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Bevel_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Bevel_GetAlign").Call(obj)
	return TAlign(ret)
}

func Bevel_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Bevel_SetAlign").Call(obj, uintptr(value))
}

func Bevel_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Bevel_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Bevel_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Bevel_SetAnchors").Call(obj, uintptr(value))
}

func Bevel_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetConstraints").Call(obj)
	return ret
}

func Bevel_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetConstraints").Call(obj, value)
}

func Bevel_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bevel_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Bevel_GetShape(obj uintptr) TBevelShape {
	ret, _, _ := getLazyProc("Bevel_GetShape").Call(obj)
	return TBevelShape(ret)
}

func Bevel_SetShape(obj uintptr, value TBevelShape) {
	_, _, _ = getLazyProc("Bevel_SetShape").Call(obj, uintptr(value))
}

func Bevel_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bevel_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Bevel_GetStyle(obj uintptr) TBevelStyle {
	ret, _, _ := getLazyProc("Bevel_GetStyle").Call(obj)
	return TBevelStyle(ret)
}

func Bevel_SetStyle(obj uintptr, value TBevelStyle) {
	_, _, _ = getLazyProc("Bevel_SetStyle").Call(obj, uintptr(value))
}

func Bevel_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bevel_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Bevel_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Bevel_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Bevel_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAction").Call(obj)
	return ret
}

func Bevel_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetAction").Call(obj, value)
}

func Bevel_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Bevel_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Bevel_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Bevel_SetBiDiMode").Call(obj, uintptr(value))
}

func Bevel_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Bevel_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Bevel_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Bevel_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetClientHeight").Call(obj)
	return int32(ret)
}

func Bevel_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetClientHeight").Call(obj, uintptr(value))
}

func Bevel_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Bevel_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Bevel_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Bevel_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetClientWidth").Call(obj)
	return int32(ret)
}

func Bevel_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetClientWidth").Call(obj, uintptr(value))
}

func Bevel_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Bevel_GetControlState").Call(obj)
	return TControlState(ret)
}

func Bevel_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Bevel_SetControlState").Call(obj, uintptr(value))
}

func Bevel_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Bevel_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Bevel_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Bevel_SetControlStyle").Call(obj, uintptr(value))
}

func Bevel_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Bevel_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Bevel_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetParent").Call(obj)
	return ret
}

func Bevel_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetParent").Call(obj, value)
}

func Bevel_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetLeft").Call(obj)
	return int32(ret)
}

func Bevel_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetLeft").Call(obj, uintptr(value))
}

func Bevel_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetTop").Call(obj)
	return int32(ret)
}

func Bevel_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetTop").Call(obj, uintptr(value))
}

func Bevel_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetWidth").Call(obj)
	return int32(ret)
}

func Bevel_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetWidth").Call(obj, uintptr(value))
}

func Bevel_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetHeight").Call(obj)
	return int32(ret)
}

func Bevel_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetHeight").Call(obj, uintptr(value))
}

func Bevel_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Bevel_GetCursor").Call(obj)
	return TCursor(ret)
}

func Bevel_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Bevel_SetCursor").Call(obj, uintptr(value))
}

func Bevel_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Bevel_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Bevel_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Bevel_SetHint").Call(obj, GoStrToDStr(value))
}

func Bevel_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetComponentCount").Call(obj)
	return int32(ret)
}

func Bevel_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Bevel_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Bevel_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Bevel_SetComponentIndex").Call(obj, uintptr(value))
}

func Bevel_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetOwner").Call(obj)
	return ret
}

func Bevel_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Bevel_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Bevel_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Bevel_SetName").Call(obj, GoStrToDStr(value))
}

func Bevel_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Bevel_GetTag").Call(obj)
	return int(ret)
}

func Bevel_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Bevel_SetTag").Call(obj, uintptr(value))
}

func Bevel_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAnchorSideLeft").Call(obj)
	return ret
}

func Bevel_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetAnchorSideLeft").Call(obj, value)
}

func Bevel_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAnchorSideTop").Call(obj)
	return ret
}

func Bevel_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetAnchorSideTop").Call(obj, value)
}

func Bevel_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAnchorSideRight").Call(obj)
	return ret
}

func Bevel_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetAnchorSideRight").Call(obj, value)
}

func Bevel_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAnchorSideBottom").Call(obj)
	return ret
}

func Bevel_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetAnchorSideBottom").Call(obj, value)
}

func Bevel_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetBorderSpacing").Call(obj)
	return ret
}

func Bevel_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Bevel_SetBorderSpacing").Call(obj, value)
}

func Bevel_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Bevel_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Bevel_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Bevel_StaticClassType() TClass {
	r, _, _ := getLazyProc("Bevel_StaticClassType").Call()
	return TClass(r)
}
