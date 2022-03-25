package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TControl ---------------------------

func Control_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_Create").Call(obj)
	return ret
}

func Control_Free(obj uintptr) {
	_, _, _ = getLazyProc("Control_Free").Call(obj)
}

func Control_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Control_BringToFront").Call(obj)
}

func Control_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Control_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Control_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Control_Hide").Call(obj)
}

func Control_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Control_Invalidate").Call(obj)
}

func Control_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Control_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Control_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Control_Refresh").Call(obj)
}

func Control_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Control_Repaint").Call(obj)
}

func Control_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Control_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Control_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Control_SendToBack").Call(obj)
}

func Control_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Control_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Control_Show(obj uintptr) {
	_, _, _ = getLazyProc("Control_Show").Call(obj)
}

func Control_Update(obj uintptr) {
	_, _, _ = getLazyProc("Control_Update").Call(obj)
}

func Control_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Control_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Control_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetTextLen").Call(obj)
	return int32(ret)
}

func Control_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Control_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Control_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Control_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Control_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Control_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Control_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Control_Assign").Call(obj, Source)
}

func Control_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Control_ClassType").Call(obj)
	return TClass(ret)
}

func Control_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Control_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Control_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_InstanceSize").Call(obj)
	return int32(ret)
}

func Control_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Control_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Control_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Control_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetHashCode").Call(obj)
	return int32(ret)
}

func Control_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Control_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Control_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Control_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Control_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Control_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Control_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Control_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Control_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Control_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Control_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Control_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Control_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Control_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Control_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Control_AnchorClient").Call(obj, uintptr(ASpace))
}

func Control_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Control_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Control_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Control_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Control_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Control_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Control_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Control_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Control_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Control_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Control_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetAction").Call(obj)
	return ret
}

func Control_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetAction").Call(obj, value)
}

func Control_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Control_GetAlign").Call(obj)
	return TAlign(ret)
}

func Control_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Control_SetAlign").Call(obj, uintptr(value))
}

func Control_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Control_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Control_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Control_SetAnchors").Call(obj, uintptr(value))
}

func Control_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Control_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Control_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Control_SetBiDiMode").Call(obj, uintptr(value))
}

func Control_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Control_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Control_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Control_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetClientHeight").Call(obj)
	return int32(ret)
}

func Control_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetClientHeight").Call(obj, uintptr(value))
}

func Control_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Control_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Control_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Control_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetClientWidth").Call(obj)
	return int32(ret)
}

func Control_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetClientWidth").Call(obj, uintptr(value))
}

func Control_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetConstraints").Call(obj)
	return ret
}

func Control_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetConstraints").Call(obj, value)
}

func Control_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Control_GetControlState").Call(obj)
	return TControlState(ret)
}

func Control_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Control_SetControlState").Call(obj, uintptr(value))
}

func Control_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Control_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Control_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Control_SetControlStyle").Call(obj, uintptr(value))
}

func Control_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Control_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Control_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Control_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Control_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Control_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Control_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetParent").Call(obj)
	return ret
}

func Control_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetParent").Call(obj, value)
}

func Control_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetLeft").Call(obj)
	return int32(ret)
}

func Control_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetLeft").Call(obj, uintptr(value))
}

func Control_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetTop").Call(obj)
	return int32(ret)
}

func Control_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetTop").Call(obj, uintptr(value))
}

func Control_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetWidth").Call(obj)
	return int32(ret)
}

func Control_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetWidth").Call(obj, uintptr(value))
}

func Control_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetHeight").Call(obj)
	return int32(ret)
}

func Control_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetHeight").Call(obj, uintptr(value))
}

func Control_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Control_GetCursor").Call(obj)
	return TCursor(ret)
}

func Control_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Control_SetCursor").Call(obj, uintptr(value))
}

func Control_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Control_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Control_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Control_SetHint").Call(obj, GoStrToDStr(value))
}

func Control_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetComponentCount").Call(obj)
	return int32(ret)
}

func Control_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Control_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Control_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Control_SetComponentIndex").Call(obj, uintptr(value))
}

func Control_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetOwner").Call(obj)
	return ret
}

func Control_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Control_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Control_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Control_SetName").Call(obj, GoStrToDStr(value))
}

func Control_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Control_GetTag").Call(obj)
	return int(ret)
}

func Control_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Control_SetTag").Call(obj, uintptr(value))
}

func Control_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetAnchorSideLeft").Call(obj)
	return ret
}

func Control_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetAnchorSideLeft").Call(obj, value)
}

func Control_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetAnchorSideTop").Call(obj)
	return ret
}

func Control_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetAnchorSideTop").Call(obj, value)
}

func Control_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetAnchorSideRight").Call(obj)
	return ret
}

func Control_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetAnchorSideRight").Call(obj, value)
}

func Control_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetAnchorSideBottom").Call(obj)
	return ret
}

func Control_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetAnchorSideBottom").Call(obj, value)
}

func Control_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Control_GetBorderSpacing").Call(obj)
	return ret
}

func Control_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Control_SetBorderSpacing").Call(obj, value)
}

func Control_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Control_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Control_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Control_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Control_StaticClassType() TClass {
	r, _, _ := getLazyProc("Control_StaticClassType").Call()
	return TClass(r)
}
