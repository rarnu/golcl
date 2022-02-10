package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TGauge ---------------------------

func Gauge_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_Create").Call(obj)
	return ret
}

func Gauge_Free(obj uintptr) {
	getLazyProc("Gauge_Free").Call(obj)
}

func Gauge_AddProgress(obj uintptr, Value int32) {
	getLazyProc("Gauge_AddProgress").Call(obj, uintptr(Value))
}

func Gauge_BringToFront(obj uintptr) {
	getLazyProc("Gauge_BringToFront").Call(obj)
}

func Gauge_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Gauge_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Gauge_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_Hide(obj uintptr) {
	getLazyProc("Gauge_Hide").Call(obj)
}

func Gauge_Invalidate(obj uintptr) {
	getLazyProc("Gauge_Invalidate").Call(obj)
}

func Gauge_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Gauge_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Gauge_Refresh(obj uintptr) {
	getLazyProc("Gauge_Refresh").Call(obj)
}

func Gauge_Repaint(obj uintptr) {
	getLazyProc("Gauge_Repaint").Call(obj)
}

func Gauge_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("Gauge_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("Gauge_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_SendToBack(obj uintptr) {
	getLazyProc("Gauge_SendToBack").Call(obj)
}

func Gauge_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("Gauge_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Gauge_Show(obj uintptr) {
	getLazyProc("Gauge_Show").Call(obj)
}

func Gauge_Update(obj uintptr) {
	getLazyProc("Gauge_Update").Call(obj)
}

func Gauge_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Gauge_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Gauge_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetTextLen").Call(obj)
	return int32(ret)
}

func Gauge_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("Gauge_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Gauge_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Gauge_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Gauge_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Gauge_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Gauge_Assign(obj uintptr, Source uintptr) {
	getLazyProc("Gauge_Assign").Call(obj, Source)
}

func Gauge_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Gauge_ClassType").Call(obj)
	return TClass(ret)
}

func Gauge_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Gauge_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Gauge_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_InstanceSize").Call(obj)
	return int32(ret)
}

func Gauge_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Gauge_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Gauge_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Gauge_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetHashCode").Call(obj)
	return int32(ret)
}

func Gauge_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Gauge_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Gauge_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Gauge_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Gauge_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("Gauge_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Gauge_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Gauge_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Gauge_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("Gauge_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Gauge_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("Gauge_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Gauge_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("Gauge_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Gauge_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("Gauge_AnchorClient").Call(obj, uintptr(ASpace))
}

func Gauge_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Gauge_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Gauge_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("Gauge_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Gauge_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("Gauge_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Gauge_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("Gauge_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Gauge_GetPercentDone(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetPercentDone").Call(obj)
	return int32(ret)
}

func Gauge_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Gauge_GetAlign").Call(obj)
	return TAlign(ret)
}

func Gauge_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("Gauge_SetAlign").Call(obj, uintptr(value))
}

func Gauge_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Gauge_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Gauge_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("Gauge_SetAnchors").Call(obj, uintptr(value))
}

func Gauge_GetBackColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Gauge_GetBackColor").Call(obj)
	return TColor(ret)
}

func Gauge_SetBackColor(obj uintptr, value TColor) {
	getLazyProc("Gauge_SetBackColor").Call(obj, uintptr(value))
}

func Gauge_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("Gauge_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func Gauge_SetBorderStyle(obj uintptr, value TBorderStyle) {
	getLazyProc("Gauge_SetBorderStyle").Call(obj, uintptr(value))
}

func Gauge_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Gauge_GetColor").Call(obj)
	return TColor(ret)
}

func Gauge_SetColor(obj uintptr, value TColor) {
	getLazyProc("Gauge_SetColor").Call(obj, uintptr(value))
}

func Gauge_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetConstraints").Call(obj)
	return ret
}

func Gauge_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetConstraints").Call(obj, value)
}

func Gauge_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetEnabled(obj uintptr, value bool) {
	getLazyProc("Gauge_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetForeColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Gauge_GetForeColor").Call(obj)
	return TColor(ret)
}

func Gauge_SetForeColor(obj uintptr, value TColor) {
	getLazyProc("Gauge_SetForeColor").Call(obj, uintptr(value))
}

func Gauge_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetFont").Call(obj)
	return ret
}

func Gauge_SetFont(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetFont").Call(obj, value)
}

func Gauge_GetKind(obj uintptr) TGaugeKind {
	ret, _, _ := getLazyProc("Gauge_GetKind").Call(obj)
	return TGaugeKind(ret)
}

func Gauge_SetKind(obj uintptr, value TGaugeKind) {
	getLazyProc("Gauge_SetKind").Call(obj, uintptr(value))
}

func Gauge_GetMinValue(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetMinValue").Call(obj)
	return int32(ret)
}

func Gauge_SetMinValue(obj uintptr, value int32) {
	getLazyProc("Gauge_SetMinValue").Call(obj, uintptr(value))
}

func Gauge_GetMaxValue(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetMaxValue").Call(obj)
	return int32(ret)
}

func Gauge_SetMaxValue(obj uintptr, value int32) {
	getLazyProc("Gauge_SetMaxValue").Call(obj, uintptr(value))
}

func Gauge_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetParentColor(obj uintptr, value bool) {
	getLazyProc("Gauge_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetParentFont(obj uintptr, value bool) {
	getLazyProc("Gauge_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("Gauge_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetPopupMenu").Call(obj)
	return ret
}

func Gauge_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetPopupMenu").Call(obj, value)
}

func Gauge_GetProgress(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetProgress").Call(obj)
	return int32(ret)
}

func Gauge_SetProgress(obj uintptr, value int32) {
	getLazyProc("Gauge_SetProgress").Call(obj, uintptr(value))
}

func Gauge_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetShowHint(obj uintptr, value bool) {
	getLazyProc("Gauge_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetShowText(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetShowText").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetShowText(obj uintptr, value bool) {
	getLazyProc("Gauge_SetShowText").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_SetVisible(obj uintptr, value bool) {
	getLazyProc("Gauge_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Gauge_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAction").Call(obj)
	return ret
}

func Gauge_SetAction(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetAction").Call(obj, value)
}

func Gauge_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Gauge_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Gauge_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("Gauge_SetBiDiMode").Call(obj, uintptr(value))
}

func Gauge_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Gauge_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("Gauge_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Gauge_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetClientHeight").Call(obj)
	return int32(ret)
}

func Gauge_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("Gauge_SetClientHeight").Call(obj, uintptr(value))
}

func Gauge_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("Gauge_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("Gauge_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Gauge_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetClientWidth").Call(obj)
	return int32(ret)
}

func Gauge_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("Gauge_SetClientWidth").Call(obj, uintptr(value))
}

func Gauge_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Gauge_GetControlState").Call(obj)
	return TControlState(ret)
}

func Gauge_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("Gauge_SetControlState").Call(obj, uintptr(value))
}

func Gauge_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Gauge_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Gauge_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("Gauge_SetControlStyle").Call(obj, uintptr(value))
}

func Gauge_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Gauge_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Gauge_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetParent").Call(obj)
	return ret
}

func Gauge_SetParent(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetParent").Call(obj, value)
}

func Gauge_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetLeft").Call(obj)
	return int32(ret)
}

func Gauge_SetLeft(obj uintptr, value int32) {
	getLazyProc("Gauge_SetLeft").Call(obj, uintptr(value))
}

func Gauge_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetTop").Call(obj)
	return int32(ret)
}

func Gauge_SetTop(obj uintptr, value int32) {
	getLazyProc("Gauge_SetTop").Call(obj, uintptr(value))
}

func Gauge_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetWidth").Call(obj)
	return int32(ret)
}

func Gauge_SetWidth(obj uintptr, value int32) {
	getLazyProc("Gauge_SetWidth").Call(obj, uintptr(value))
}

func Gauge_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetHeight").Call(obj)
	return int32(ret)
}

func Gauge_SetHeight(obj uintptr, value int32) {
	getLazyProc("Gauge_SetHeight").Call(obj, uintptr(value))
}

func Gauge_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Gauge_GetCursor").Call(obj)
	return TCursor(ret)
}

func Gauge_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("Gauge_SetCursor").Call(obj, uintptr(value))
}

func Gauge_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Gauge_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Gauge_SetHint(obj uintptr, value string) {
	getLazyProc("Gauge_SetHint").Call(obj, GoStrToDStr(value))
}

func Gauge_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetComponentCount").Call(obj)
	return int32(ret)
}

func Gauge_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Gauge_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Gauge_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("Gauge_SetComponentIndex").Call(obj, uintptr(value))
}

func Gauge_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetOwner").Call(obj)
	return ret
}

func Gauge_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Gauge_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Gauge_SetName(obj uintptr, value string) {
	getLazyProc("Gauge_SetName").Call(obj, GoStrToDStr(value))
}

func Gauge_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Gauge_GetTag").Call(obj)
	return int(ret)
}

func Gauge_SetTag(obj uintptr, value int) {
	getLazyProc("Gauge_SetTag").Call(obj, uintptr(value))
}

func Gauge_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAnchorSideLeft").Call(obj)
	return ret
}

func Gauge_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetAnchorSideLeft").Call(obj, value)
}

func Gauge_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAnchorSideTop").Call(obj)
	return ret
}

func Gauge_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetAnchorSideTop").Call(obj, value)
}

func Gauge_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAnchorSideRight").Call(obj)
	return ret
}

func Gauge_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetAnchorSideRight").Call(obj, value)
}

func Gauge_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAnchorSideBottom").Call(obj)
	return ret
}

func Gauge_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetAnchorSideBottom").Call(obj, value)
}

func Gauge_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetBorderSpacing").Call(obj)
	return ret
}

func Gauge_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("Gauge_SetBorderSpacing").Call(obj, value)
}

func Gauge_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Gauge_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Gauge_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Gauge_StaticClassType() TClass {
	r, _, _ := getLazyProc("Gauge_StaticClassType").Call()
	return TClass(r)
}
