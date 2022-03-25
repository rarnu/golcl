package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TSpeedButton ---------------------------

func SpeedButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_Create").Call(obj)
	return ret
}

func SpeedButton_Free(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Free").Call(obj)
}

func SpeedButton_Click(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Click").Call(obj)
}

func SpeedButton_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_BringToFront").Call(obj)
}

func SpeedButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpeedButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpeedButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_Hide(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Hide").Call(obj)
}

func SpeedButton_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Invalidate").Call(obj)
}

func SpeedButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("SpeedButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func SpeedButton_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Refresh").Call(obj)
}

func SpeedButton_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Repaint").Call(obj)
}

func SpeedButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpeedButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpeedButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SendToBack").Call(obj)
}

func SpeedButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("SpeedButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func SpeedButton_Show(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Show").Call(obj)
}

func SpeedButton_Update(obj uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Update").Call(obj)
}

func SpeedButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("SpeedButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func SpeedButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetTextLen").Call(obj)
	return int32(ret)
}

func SpeedButton_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("SpeedButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func SpeedButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func SpeedButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("SpeedButton_Assign").Call(obj, Source)
}

func SpeedButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("SpeedButton_ClassType").Call(obj)
	return TClass(ret)
}

func SpeedButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_InstanceSize").Call(obj)
	return int32(ret)
}

func SpeedButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("SpeedButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func SpeedButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetHashCode").Call(obj)
	return int32(ret)
}

func SpeedButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SpeedButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SpeedButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("SpeedButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func SpeedButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SpeedButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func SpeedButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("SpeedButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func SpeedButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("SpeedButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func SpeedButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("SpeedButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func SpeedButton_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("SpeedButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func SpeedButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("SpeedButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func SpeedButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("SpeedButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func SpeedButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("SpeedButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func SpeedButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("SpeedButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func SpeedButton_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetImageIndex").Call(obj)
	return int32(ret)
}

func SpeedButton_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetImageIndex").Call(obj, uintptr(value))
}

func SpeedButton_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetImages").Call(obj)
	return ret
}

func SpeedButton_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetImages").Call(obj, value)
}

func SpeedButton_GetImageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetImageWidth").Call(obj)
	return int32(ret)
}

func SpeedButton_SetImageWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetImageWidth").Call(obj, uintptr(value))
}

func SpeedButton_GetShowCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetShowCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetShowCaption(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetShowCaption").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAction").Call(obj)
	return ret
}

func SpeedButton_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetAction").Call(obj, value)
}

func SpeedButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("SpeedButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func SpeedButton_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("SpeedButton_SetAlign").Call(obj, uintptr(value))
}

func SpeedButton_GetAllowAllUp(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetAllowAllUp").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetAllowAllUp(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetAllowAllUp").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func SpeedButton_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("SpeedButton_SetAnchors").Call(obj, uintptr(value))
}

func SpeedButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("SpeedButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func SpeedButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("SpeedButton_SetBiDiMode").Call(obj, uintptr(value))
}

func SpeedButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetConstraints").Call(obj)
	return ret
}

func SpeedButton_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetConstraints").Call(obj, value)
}

func SpeedButton_GetGroupIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetGroupIndex").Call(obj)
	return int32(ret)
}

func SpeedButton_SetGroupIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetGroupIndex").Call(obj, uintptr(value))
}

func SpeedButton_GetDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetDown").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetDown(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetDown").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpeedButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFlat(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetFlat").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetFlat(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetFlat").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetFont").Call(obj)
	return ret
}

func SpeedButton_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetFont").Call(obj, value)
}

func SpeedButton_GetGlyph(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetGlyph").Call(obj)
	return ret
}

func SpeedButton_SetGlyph(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetGlyph").Call(obj, value)
}

func SpeedButton_GetLayout(obj uintptr) TButtonLayout {
	ret, _, _ := getLazyProc("SpeedButton_GetLayout").Call(obj)
	return TButtonLayout(ret)
}

func SpeedButton_SetLayout(obj uintptr, value TButtonLayout) {
	_, _, _ = getLazyProc("SpeedButton_SetLayout").Call(obj, uintptr(value))
}

func SpeedButton_GetNumGlyphs(obj uintptr) TNumGlyphs {
	ret, _, _ := getLazyProc("SpeedButton_GetNumGlyphs").Call(obj)
	return TNumGlyphs(ret)
}

func SpeedButton_SetNumGlyphs(obj uintptr, value TNumGlyphs) {
	_, _, _ = getLazyProc("SpeedButton_SetNumGlyphs").Call(obj, uintptr(value))
}

func SpeedButton_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetPopupMenu").Call(obj)
	return ret
}

func SpeedButton_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetPopupMenu").Call(obj, value)
}

func SpeedButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetSpacing").Call(obj)
	return int32(ret)
}

func SpeedButton_SetSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetSpacing").Call(obj, uintptr(value))
}

func SpeedButton_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SpeedButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func SpeedButton_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("SpeedButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func SpeedButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("SpeedButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("SpeedButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func SpeedButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func SpeedButton_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetClientHeight").Call(obj, uintptr(value))
}

func SpeedButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("SpeedButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("SpeedButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func SpeedButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func SpeedButton_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetClientWidth").Call(obj, uintptr(value))
}

func SpeedButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("SpeedButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func SpeedButton_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("SpeedButton_SetControlState").Call(obj, uintptr(value))
}

func SpeedButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("SpeedButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func SpeedButton_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("SpeedButton_SetControlStyle").Call(obj, uintptr(value))
}

func SpeedButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("SpeedButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func SpeedButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetParent").Call(obj)
	return ret
}

func SpeedButton_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetParent").Call(obj, value)
}

func SpeedButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetLeft").Call(obj)
	return int32(ret)
}

func SpeedButton_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetLeft").Call(obj, uintptr(value))
}

func SpeedButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetTop").Call(obj)
	return int32(ret)
}

func SpeedButton_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetTop").Call(obj, uintptr(value))
}

func SpeedButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetWidth").Call(obj)
	return int32(ret)
}

func SpeedButton_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetWidth").Call(obj, uintptr(value))
}

func SpeedButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetHeight").Call(obj)
	return int32(ret)
}

func SpeedButton_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetHeight").Call(obj, uintptr(value))
}

func SpeedButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("SpeedButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func SpeedButton_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("SpeedButton_SetCursor").Call(obj, uintptr(value))
}

func SpeedButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpeedButton_SetHint").Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func SpeedButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("SpeedButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func SpeedButton_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SpeedButton_SetComponentIndex").Call(obj, uintptr(value))
}

func SpeedButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetOwner").Call(obj)
	return ret
}

func SpeedButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("SpeedButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func SpeedButton_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("SpeedButton_SetName").Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("SpeedButton_GetTag").Call(obj)
	return int(ret)
}

func SpeedButton_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("SpeedButton_SetTag").Call(obj, uintptr(value))
}

func SpeedButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func SpeedButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetAnchorSideLeft").Call(obj, value)
}

func SpeedButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchorSideTop").Call(obj)
	return ret
}

func SpeedButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetAnchorSideTop").Call(obj, value)
}

func SpeedButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchorSideRight").Call(obj)
	return ret
}

func SpeedButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetAnchorSideRight").Call(obj, value)
}

func SpeedButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func SpeedButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetAnchorSideBottom").Call(obj, value)
}

func SpeedButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetBorderSpacing").Call(obj)
	return ret
}

func SpeedButton_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SpeedButton_SetBorderSpacing").Call(obj, value)
}

func SpeedButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func SpeedButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("SpeedButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func SpeedButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("SpeedButton_StaticClassType").Call()
	return TClass(r)
}
