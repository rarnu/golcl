package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TColorButton ---------------------------

func ColorButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_Create").Call(obj)
	return ret
}

func ColorButton_Free(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Free").Call(obj)
}

func ColorButton_Click(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Click").Call(obj)
}

func ColorButton_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_BringToFront").Call(obj)
}

func ColorButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ColorButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ColorButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Hide").Call(obj)
}

func ColorButton_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Invalidate").Call(obj)
}

func ColorButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ColorButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ColorButton_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Refresh").Call(obj)
}

func ColorButton_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Repaint").Call(obj)
}

func ColorButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ColorButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ColorButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_SendToBack").Call(obj)
}

func ColorButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ColorButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ColorButton_Show(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Show").Call(obj)
}

func ColorButton_Update(obj uintptr) {
	_, _, _ = getLazyProc("ColorButton_Update").Call(obj)
}

func ColorButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ColorButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ColorButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetTextLen").Call(obj)
	return int32(ret)
}

func ColorButton_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ColorButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ColorButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ColorButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ColorButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ColorButton_Assign").Call(obj, Source)
}

func ColorButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ColorButton_ClassType").Call(obj)
	return TClass(ret)
}

func ColorButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_InstanceSize").Call(obj)
	return int32(ret)
}

func ColorButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ColorButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ColorButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ColorButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetHashCode").Call(obj)
	return int32(ret)
}

func ColorButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ColorButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ColorButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ColorButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ColorButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ColorButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ColorButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ColorButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ColorButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ColorButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ColorButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ColorButton_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ColorButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func ColorButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ColorButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ColorButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ColorButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ColorButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ColorButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ColorButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ColorButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ColorButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAction").Call(obj)
	return ret
}

func ColorButton_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetAction").Call(obj, value)
}

func ColorButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ColorButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func ColorButton_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ColorButton_SetAlign").Call(obj, uintptr(value))
}

func ColorButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ColorButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ColorButton_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ColorButton_SetAnchors").Call(obj, uintptr(value))
}

func ColorButton_GetAllowAllUp(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetAllowAllUp").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetAllowAllUp(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetAllowAllUp").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetBorderWidth").Call(obj)
	return int32(ret)
}

func ColorButton_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetBorderWidth").Call(obj, uintptr(value))
}

func ColorButton_GetButtonColorAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetButtonColorAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetButtonColorAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetButtonColorAutoSize").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetButtonColorSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetButtonColorSize").Call(obj)
	return int32(ret)
}

func ColorButton_SetButtonColorSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetButtonColorSize").Call(obj, uintptr(value))
}

func ColorButton_GetButtonColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorButton_GetButtonColor").Call(obj)
	return TColor(ret)
}

func ColorButton_SetButtonColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ColorButton_SetButtonColor").Call(obj, uintptr(value))
}

func ColorButton_GetColorDialog(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetColorDialog").Call(obj)
	return ret
}

func ColorButton_SetColorDialog(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetColorDialog").Call(obj, value)
}

func ColorButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetConstraints").Call(obj)
	return ret
}

func ColorButton_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetConstraints").Call(obj, value)
}

func ColorButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("ColorButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func ColorButton_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ColorButton_GetColor").Call(obj)
	return TColor(ret)
}

func ColorButton_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ColorButton_SetColor").Call(obj, uintptr(value))
}

func ColorButton_GetDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetDown").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetDown(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetDown").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetFlat(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetFlat").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetFlat(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetFlat").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetFont").Call(obj)
	return ret
}

func ColorButton_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetFont").Call(obj, value)
}

func ColorButton_GetGroupIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetGroupIndex").Call(obj)
	return int32(ret)
}

func ColorButton_SetGroupIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetGroupIndex").Call(obj, uintptr(value))
}

func ColorButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ColorButton_SetHint").Call(obj, GoStrToDStr(value))
}

func ColorButton_GetLayout(obj uintptr) TButtonLayout {
	ret, _, _ := getLazyProc("ColorButton_GetLayout").Call(obj)
	return TButtonLayout(ret)
}

func ColorButton_SetLayout(obj uintptr, value TButtonLayout) {
	_, _, _ = getLazyProc("ColorButton_SetLayout").Call(obj, uintptr(value))
}

func ColorButton_GetSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetSpacing").Call(obj)
	return int32(ret)
}

func ColorButton_SetSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetSpacing").Call(obj, uintptr(value))
}

func ColorButton_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ColorButton_SetOnClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnColorChanged(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnColorChanged").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnDblClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseMove(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseWheel(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnPaint(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_SetOnResize(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("ColorButton_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func ColorButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetPopupMenu").Call(obj)
	return ret
}

func ColorButton_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetPopupMenu").Call(obj, value)
}

func ColorButton_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetImageIndex").Call(obj)
	return int32(ret)
}

func ColorButton_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetImageIndex").Call(obj, uintptr(value))
}

func ColorButton_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetImages").Call(obj)
	return ret
}

func ColorButton_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetImages").Call(obj, value)
}

func ColorButton_GetImageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetImageWidth").Call(obj)
	return int32(ret)
}

func ColorButton_SetImageWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetImageWidth").Call(obj, uintptr(value))
}

func ColorButton_GetShowCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetShowCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_SetShowCaption(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ColorButton_SetShowCaption").Call(obj, GoBoolToDBool(value))
}

func ColorButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ColorButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ColorButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ColorButton_SetBiDiMode").Call(obj, uintptr(value))
}

func ColorButton_GetGlyph(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetGlyph").Call(obj)
	return ret
}

func ColorButton_SetGlyph(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetGlyph").Call(obj, value)
}

func ColorButton_GetNumGlyphs(obj uintptr) TNumGlyphs {
	ret, _, _ := getLazyProc("ColorButton_GetNumGlyphs").Call(obj)
	return TNumGlyphs(ret)
}

func ColorButton_SetNumGlyphs(obj uintptr, value TNumGlyphs) {
	_, _, _ = getLazyProc("ColorButton_SetNumGlyphs").Call(obj, uintptr(value))
}

func ColorButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ColorButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ColorButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func ColorButton_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetClientHeight").Call(obj, uintptr(value))
}

func ColorButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ColorButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ColorButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ColorButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func ColorButton_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetClientWidth").Call(obj, uintptr(value))
}

func ColorButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ColorButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func ColorButton_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ColorButton_SetControlState").Call(obj, uintptr(value))
}

func ColorButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ColorButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ColorButton_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ColorButton_SetControlStyle").Call(obj, uintptr(value))
}

func ColorButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ColorButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ColorButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetParent").Call(obj)
	return ret
}

func ColorButton_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetParent").Call(obj, value)
}

func ColorButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetLeft").Call(obj)
	return int32(ret)
}

func ColorButton_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetLeft").Call(obj, uintptr(value))
}

func ColorButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetTop").Call(obj)
	return int32(ret)
}

func ColorButton_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetTop").Call(obj, uintptr(value))
}

func ColorButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetWidth").Call(obj)
	return int32(ret)
}

func ColorButton_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetWidth").Call(obj, uintptr(value))
}

func ColorButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetHeight").Call(obj)
	return int32(ret)
}

func ColorButton_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetHeight").Call(obj, uintptr(value))
}

func ColorButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ColorButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func ColorButton_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ColorButton_SetCursor").Call(obj, uintptr(value))
}

func ColorButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func ColorButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ColorButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ColorButton_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ColorButton_SetComponentIndex").Call(obj, uintptr(value))
}

func ColorButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetOwner").Call(obj)
	return ret
}

func ColorButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ColorButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ColorButton_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ColorButton_SetName").Call(obj, GoStrToDStr(value))
}

func ColorButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ColorButton_GetTag").Call(obj)
	return int(ret)
}

func ColorButton_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ColorButton_SetTag").Call(obj, uintptr(value))
}

func ColorButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func ColorButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetAnchorSideLeft").Call(obj, value)
}

func ColorButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAnchorSideTop").Call(obj)
	return ret
}

func ColorButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetAnchorSideTop").Call(obj, value)
}

func ColorButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAnchorSideRight").Call(obj)
	return ret
}

func ColorButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetAnchorSideRight").Call(obj, value)
}

func ColorButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func ColorButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetAnchorSideBottom").Call(obj, value)
}

func ColorButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetBorderSpacing").Call(obj)
	return ret
}

func ColorButton_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ColorButton_SetBorderSpacing").Call(obj, value)
}

func ColorButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ColorButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ColorButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ColorButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("ColorButton_StaticClassType").Call()
	return TClass(r)
}
