package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TSplitter ---------------------------

func Splitter_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_Create").Call(obj)
	return ret
}

func Splitter_Free(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Free").Call(obj)
}

func Splitter_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_BringToFront").Call(obj)
}

func Splitter_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Splitter_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Splitter_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Hide").Call(obj)
}

func Splitter_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Invalidate").Call(obj)
}

func Splitter_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Splitter_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Splitter_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Refresh").Call(obj)
}

func Splitter_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Repaint").Call(obj)
}

func Splitter_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Splitter_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Splitter_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_SendToBack").Call(obj)
}

func Splitter_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Splitter_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Splitter_Show(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Show").Call(obj)
}

func Splitter_Update(obj uintptr) {
	_, _, _ = getLazyProc("Splitter_Update").Call(obj)
}

func Splitter_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Splitter_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Splitter_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetTextLen").Call(obj)
	return int32(ret)
}

func Splitter_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Splitter_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Splitter_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Splitter_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Splitter_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Splitter_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Splitter_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Splitter_Assign").Call(obj, Source)
}

func Splitter_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Splitter_ClassType").Call(obj)
	return TClass(ret)
}

func Splitter_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Splitter_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Splitter_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_InstanceSize").Call(obj)
	return int32(ret)
}

func Splitter_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Splitter_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Splitter_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Splitter_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetHashCode").Call(obj)
	return int32(ret)
}

func Splitter_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Splitter_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Splitter_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Splitter_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Splitter_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Splitter_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Splitter_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Splitter_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Splitter_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Splitter_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Splitter_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Splitter_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Splitter_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Splitter_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Splitter_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Splitter_AnchorClient").Call(obj, uintptr(ASpace))
}

func Splitter_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Splitter_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Splitter_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Splitter_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Splitter_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Splitter_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Splitter_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Splitter_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Splitter_GetResizeAnchor(obj uintptr) TAnchorKind {
	ret, _, _ := getLazyProc("Splitter_GetResizeAnchor").Call(obj)
	return TAnchorKind(ret)
}

func Splitter_SetResizeAnchor(obj uintptr, value TAnchorKind) {
	_, _, _ = getLazyProc("Splitter_SetResizeAnchor").Call(obj, uintptr(value))
}

func Splitter_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetCanvas").Call(obj)
	return ret
}

func Splitter_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Splitter_GetAlign").Call(obj)
	return TAlign(ret)
}

func Splitter_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Splitter_SetAlign").Call(obj, uintptr(value))
}

func Splitter_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Splitter_GetColor").Call(obj)
	return TColor(ret)
}

func Splitter_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Splitter_SetColor").Call(obj, uintptr(value))
}

func Splitter_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Splitter_GetCursor").Call(obj)
	return TCursor(ret)
}

func Splitter_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Splitter_SetCursor").Call(obj, uintptr(value))
}

func Splitter_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetConstraints").Call(obj)
	return ret
}

func Splitter_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetConstraints").Call(obj, value)
}

func Splitter_GetMinSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetMinSize").Call(obj)
	return int32(ret)
}

func Splitter_SetMinSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetMinSize").Call(obj, uintptr(value))
}

func Splitter_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Splitter_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Splitter_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Splitter_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Splitter_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetWidth").Call(obj)
	return int32(ret)
}

func Splitter_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetWidth").Call(obj, uintptr(value))
}

func Splitter_SetOnPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Splitter_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func Splitter_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Splitter_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Splitter_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAction").Call(obj)
	return ret
}

func Splitter_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetAction").Call(obj, value)
}

func Splitter_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Splitter_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Splitter_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Splitter_SetAnchors").Call(obj, uintptr(value))
}

func Splitter_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Splitter_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Splitter_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Splitter_SetBiDiMode").Call(obj, uintptr(value))
}

func Splitter_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Splitter_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Splitter_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Splitter_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetClientHeight").Call(obj)
	return int32(ret)
}

func Splitter_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetClientHeight").Call(obj, uintptr(value))
}

func Splitter_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Splitter_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Splitter_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Splitter_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetClientWidth").Call(obj)
	return int32(ret)
}

func Splitter_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetClientWidth").Call(obj, uintptr(value))
}

func Splitter_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Splitter_GetControlState").Call(obj)
	return TControlState(ret)
}

func Splitter_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Splitter_SetControlState").Call(obj, uintptr(value))
}

func Splitter_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Splitter_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Splitter_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Splitter_SetControlStyle").Call(obj, uintptr(value))
}

func Splitter_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Splitter_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Splitter_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Splitter_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Splitter_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetParent").Call(obj)
	return ret
}

func Splitter_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetParent").Call(obj, value)
}

func Splitter_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetLeft").Call(obj)
	return int32(ret)
}

func Splitter_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetLeft").Call(obj, uintptr(value))
}

func Splitter_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetTop").Call(obj)
	return int32(ret)
}

func Splitter_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetTop").Call(obj, uintptr(value))
}

func Splitter_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetHeight").Call(obj)
	return int32(ret)
}

func Splitter_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetHeight").Call(obj, uintptr(value))
}

func Splitter_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Splitter_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Splitter_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Splitter_SetHint").Call(obj, GoStrToDStr(value))
}

func Splitter_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetComponentCount").Call(obj)
	return int32(ret)
}

func Splitter_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Splitter_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Splitter_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Splitter_SetComponentIndex").Call(obj, uintptr(value))
}

func Splitter_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetOwner").Call(obj)
	return ret
}

func Splitter_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Splitter_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Splitter_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Splitter_SetName").Call(obj, GoStrToDStr(value))
}

func Splitter_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Splitter_GetTag").Call(obj)
	return int(ret)
}

func Splitter_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Splitter_SetTag").Call(obj, uintptr(value))
}

func Splitter_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAnchorSideLeft").Call(obj)
	return ret
}

func Splitter_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetAnchorSideLeft").Call(obj, value)
}

func Splitter_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAnchorSideTop").Call(obj)
	return ret
}

func Splitter_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetAnchorSideTop").Call(obj, value)
}

func Splitter_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAnchorSideRight").Call(obj)
	return ret
}

func Splitter_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetAnchorSideRight").Call(obj, value)
}

func Splitter_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAnchorSideBottom").Call(obj)
	return ret
}

func Splitter_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetAnchorSideBottom").Call(obj, value)
}

func Splitter_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetBorderSpacing").Call(obj)
	return ret
}

func Splitter_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Splitter_SetBorderSpacing").Call(obj, value)
}

func Splitter_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Splitter_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Splitter_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Splitter_StaticClassType() TClass {
	r, _, _ := getLazyProc("Splitter_StaticClassType").Call()
	return TClass(r)
}
