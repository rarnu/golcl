package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TImage ---------------------------

func Image_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_Create").Call(obj)
	return ret
}

func Image_Free(obj uintptr) {
	_, _, _ = getLazyProc("Image_Free").Call(obj)
}

func Image_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Image_BringToFront").Call(obj)
}

func Image_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Image_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Image_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Image_Hide").Call(obj)
}

func Image_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Image_Invalidate").Call(obj)
}

func Image_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Image_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Image_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Image_Refresh").Call(obj)
}

func Image_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Image_Repaint").Call(obj)
}

func Image_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Image_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Image_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Image_SendToBack").Call(obj)
}

func Image_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Image_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Image_Show(obj uintptr) {
	_, _, _ = getLazyProc("Image_Show").Call(obj)
}

func Image_Update(obj uintptr) {
	_, _, _ = getLazyProc("Image_Update").Call(obj)
}

func Image_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Image_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Image_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetTextLen").Call(obj)
	return int32(ret)
}

func Image_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Image_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Image_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Image_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Image_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Image_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Image_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Image_Assign").Call(obj, Source)
}

func Image_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Image_ClassType").Call(obj)
	return TClass(ret)
}

func Image_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Image_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Image_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_InstanceSize").Call(obj)
	return int32(ret)
}

func Image_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Image_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Image_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Image_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetHashCode").Call(obj)
	return int32(ret)
}

func Image_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Image_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Image_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Image_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Image_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Image_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Image_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Image_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Image_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Image_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Image_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Image_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Image_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Image_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Image_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Image_AnchorClient").Call(obj, uintptr(ASpace))
}

func Image_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Image_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Image_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Image_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Image_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Image_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Image_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Image_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Image_GetAntialiasingMode(obj uintptr) TAntialiasingMode {
	ret, _, _ := getLazyProc("Image_GetAntialiasingMode").Call(obj)
	return TAntialiasingMode(ret)
}

func Image_SetAntialiasingMode(obj uintptr, value TAntialiasingMode) {
	_, _, _ = getLazyProc("Image_SetAntialiasingMode").Call(obj, uintptr(value))
}

func Image_GetKeepOriginXWhenClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetKeepOriginXWhenClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetKeepOriginXWhenClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetKeepOriginXWhenClipped").Call(obj, GoBoolToDBool(value))
}

func Image_GetKeepOriginYWhenClipped(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetKeepOriginYWhenClipped").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetKeepOriginYWhenClipped(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetKeepOriginYWhenClipped").Call(obj, GoBoolToDBool(value))
}

func Image_GetStretchInEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetStretchInEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetStretchInEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetStretchInEnabled").Call(obj, GoBoolToDBool(value))
}

func Image_GetStretchOutEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetStretchOutEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetStretchOutEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetStretchOutEnabled").Call(obj, GoBoolToDBool(value))
}

func Image_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetCanvas").Call(obj)
	return ret
}

func Image_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Image_GetAlign").Call(obj)
	return TAlign(ret)
}

func Image_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Image_SetAlign").Call(obj, uintptr(value))
}

func Image_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Image_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Image_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Image_SetAnchors").Call(obj, uintptr(value))
}

func Image_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Image_GetCenter(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetCenter").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetCenter(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetCenter").Call(obj, GoBoolToDBool(value))
}

func Image_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetConstraints").Call(obj)
	return ret
}

func Image_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetConstraints").Call(obj, value)
}

func Image_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Image_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Image_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Image_SetDragCursor").Call(obj, uintptr(value))
}

func Image_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Image_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Image_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Image_SetDragMode").Call(obj, uintptr(value))
}

func Image_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Image_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Image_GetPicture(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetPicture").Call(obj)
	return ret
}

func Image_SetPicture(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetPicture").Call(obj, value)
}

func Image_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetPopupMenu").Call(obj)
	return ret
}

func Image_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetPopupMenu").Call(obj, value)
}

func Image_GetProportional(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetProportional").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetProportional(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetProportional").Call(obj, GoBoolToDBool(value))
}

func Image_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Image_GetStretch(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetStretch").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetStretch(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetStretch").Call(obj, GoBoolToDBool(value))
}

func Image_GetTransparent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetTransparent").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetTransparent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetTransparent").Call(obj, GoBoolToDBool(value))
}

func Image_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Image_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Image_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Image_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Image_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Image_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetAction").Call(obj)
	return ret
}

func Image_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetAction").Call(obj, value)
}

func Image_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Image_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Image_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Image_SetBiDiMode").Call(obj, uintptr(value))
}

func Image_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Image_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Image_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Image_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetClientHeight").Call(obj)
	return int32(ret)
}

func Image_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetClientHeight").Call(obj, uintptr(value))
}

func Image_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Image_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Image_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Image_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetClientWidth").Call(obj)
	return int32(ret)
}

func Image_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetClientWidth").Call(obj, uintptr(value))
}

func Image_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Image_GetControlState").Call(obj)
	return TControlState(ret)
}

func Image_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Image_SetControlState").Call(obj, uintptr(value))
}

func Image_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Image_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Image_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Image_SetControlStyle").Call(obj, uintptr(value))
}

func Image_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Image_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Image_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetParent").Call(obj)
	return ret
}

func Image_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetParent").Call(obj, value)
}

func Image_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetLeft").Call(obj)
	return int32(ret)
}

func Image_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetLeft").Call(obj, uintptr(value))
}

func Image_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetTop").Call(obj)
	return int32(ret)
}

func Image_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetTop").Call(obj, uintptr(value))
}

func Image_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetWidth").Call(obj)
	return int32(ret)
}

func Image_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetWidth").Call(obj, uintptr(value))
}

func Image_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetHeight").Call(obj)
	return int32(ret)
}

func Image_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetHeight").Call(obj, uintptr(value))
}

func Image_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Image_GetCursor").Call(obj)
	return TCursor(ret)
}

func Image_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Image_SetCursor").Call(obj, uintptr(value))
}

func Image_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Image_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Image_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Image_SetHint").Call(obj, GoStrToDStr(value))
}

func Image_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetComponentCount").Call(obj)
	return int32(ret)
}

func Image_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Image_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Image_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Image_SetComponentIndex").Call(obj, uintptr(value))
}

func Image_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetOwner").Call(obj)
	return ret
}

func Image_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Image_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Image_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Image_SetName").Call(obj, GoStrToDStr(value))
}

func Image_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Image_GetTag").Call(obj)
	return int(ret)
}

func Image_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Image_SetTag").Call(obj, uintptr(value))
}

func Image_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetAnchorSideLeft").Call(obj)
	return ret
}

func Image_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetAnchorSideLeft").Call(obj, value)
}

func Image_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetAnchorSideTop").Call(obj)
	return ret
}

func Image_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetAnchorSideTop").Call(obj, value)
}

func Image_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetAnchorSideRight").Call(obj)
	return ret
}

func Image_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetAnchorSideRight").Call(obj, value)
}

func Image_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetAnchorSideBottom").Call(obj)
	return ret
}

func Image_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetAnchorSideBottom").Call(obj, value)
}

func Image_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Image_GetBorderSpacing").Call(obj)
	return ret
}

func Image_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Image_SetBorderSpacing").Call(obj, value)
}

func Image_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Image_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Image_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Image_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Image_StaticClassType() TClass {
	r, _, _ := getLazyProc("Image_StaticClassType").Call()
	return TClass(r)
}
