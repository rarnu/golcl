package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TImageButton ---------------------------

func ImageButton_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_Create").Call(obj)
	return ret
}

func ImageButton_Free(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Free").Call(obj)
}

func ImageButton_Click(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Click").Call(obj)
}

func ImageButton_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_BringToFront").Call(obj)
}

func ImageButton_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageButton_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageButton_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Hide").Call(obj)
}

func ImageButton_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Invalidate").Call(obj)
}

func ImageButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ImageButton_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ImageButton_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Refresh").Call(obj)
}

func ImageButton_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Repaint").Call(obj)
}

func ImageButton_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageButton_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageButton_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_SendToBack").Call(obj)
}

func ImageButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ImageButton_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ImageButton_Show(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Show").Call(obj)
}

func ImageButton_Update(obj uintptr) {
	_, _, _ = getLazyProc("ImageButton_Update").Call(obj)
}

func ImageButton_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ImageButton_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ImageButton_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetTextLen").Call(obj)
	return int32(ret)
}

func ImageButton_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ImageButton_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ImageButton_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ImageButton_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ImageButton_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ImageButton_Assign").Call(obj, Source)
}

func ImageButton_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ImageButton_ClassType").Call(obj)
	return TClass(ret)
}

func ImageButton_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_InstanceSize").Call(obj)
	return int32(ret)
}

func ImageButton_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ImageButton_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ImageButton_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ImageButton_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetHashCode").Call(obj)
	return int32(ret)
}

func ImageButton_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ImageButton_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ImageButton_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ImageButton_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ImageButton_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ImageButton_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ImageButton_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ImageButton_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ImageButton_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ImageButton_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ImageButton_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ImageButton_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ImageButton_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ImageButton_AnchorClient").Call(obj, uintptr(ASpace))
}

func ImageButton_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ImageButton_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ImageButton_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ImageButton_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ImageButton_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ImageButton_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ImageButton_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ImageButton_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ImageButton_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAction").Call(obj)
	return ret
}

func ImageButton_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetAction").Call(obj, value)
}

func ImageButton_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ImageButton_GetAlign").Call(obj)
	return TAlign(ret)
}

func ImageButton_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ImageButton_SetAlign").Call(obj, uintptr(value))
}

func ImageButton_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ImageButton_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ImageButton_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ImageButton_SetAnchors").Call(obj, uintptr(value))
}

func ImageButton_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetConstraints").Call(obj)
	return ret
}

func ImageButton_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetConstraints").Call(obj, value)
}

func ImageButton_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("ImageButton_SetCaption").Call(obj, GoStrToDStr(value))
}

func ImageButton_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ImageButton_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ImageButton_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ImageButton_SetDragCursor").Call(obj, uintptr(value))
}

func ImageButton_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ImageButton_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ImageButton_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("ImageButton_SetDragKind").Call(obj, uintptr(value))
}

func ImageButton_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ImageButton_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ImageButton_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("ImageButton_SetDragMode").Call(obj, uintptr(value))
}

func ImageButton_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetFont").Call(obj)
	return ret
}

func ImageButton_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetFont").Call(obj, value)
}

func ImageButton_GetImageCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetImageCount").Call(obj)
	return int32(ret)
}

func ImageButton_SetImageCount(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetImageCount").Call(obj, uintptr(value))
}

func ImageButton_GetOrientation(obj uintptr) TImageOrientation {
	ret, _, _ := getLazyProc("ImageButton_GetOrientation").Call(obj)
	return TImageOrientation(ret)
}

func ImageButton_SetOrientation(obj uintptr, value TImageOrientation) {
	_, _, _ = getLazyProc("ImageButton_SetOrientation").Call(obj, uintptr(value))
}

func ImageButton_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("ImageButton_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func ImageButton_SetModalResult(obj uintptr, value TModalResult) {
	_, _, _ = getLazyProc("ImageButton_SetModalResult").Call(obj, uintptr(value))
}

func ImageButton_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetPicture(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetPicture").Call(obj)
	return ret
}

func ImageButton_SetPicture(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetPicture").Call(obj, value)
}

func ImageButton_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetPopupMenu").Call(obj)
	return ret
}

func ImageButton_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetPopupMenu").Call(obj, value)
}

func ImageButton_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetShowCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetShowCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetShowCaption(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetShowCaption").Call(obj, GoBoolToDBool(value))
}

func ImageButton_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageButton_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ImageButton_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageButton_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func ImageButton_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ImageButton_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ImageButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ImageButton_SetBiDiMode").Call(obj, uintptr(value))
}

func ImageButton_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ImageButton_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ImageButton_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ImageButton_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetClientHeight").Call(obj)
	return int32(ret)
}

func ImageButton_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetClientHeight").Call(obj, uintptr(value))
}

func ImageButton_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageButton_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ImageButton_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageButton_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetClientWidth").Call(obj)
	return int32(ret)
}

func ImageButton_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetClientWidth").Call(obj, uintptr(value))
}

func ImageButton_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ImageButton_GetControlState").Call(obj)
	return TControlState(ret)
}

func ImageButton_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ImageButton_SetControlState").Call(obj, uintptr(value))
}

func ImageButton_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ImageButton_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ImageButton_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ImageButton_SetControlStyle").Call(obj, uintptr(value))
}

func ImageButton_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageButton_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageButton_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetParent").Call(obj)
	return ret
}

func ImageButton_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetParent").Call(obj, value)
}

func ImageButton_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetLeft").Call(obj)
	return int32(ret)
}

func ImageButton_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetLeft").Call(obj, uintptr(value))
}

func ImageButton_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetTop").Call(obj)
	return int32(ret)
}

func ImageButton_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetTop").Call(obj, uintptr(value))
}

func ImageButton_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetWidth").Call(obj)
	return int32(ret)
}

func ImageButton_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetWidth").Call(obj, uintptr(value))
}

func ImageButton_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetHeight").Call(obj)
	return int32(ret)
}

func ImageButton_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetHeight").Call(obj, uintptr(value))
}

func ImageButton_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ImageButton_GetCursor").Call(obj)
	return TCursor(ret)
}

func ImageButton_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ImageButton_SetCursor").Call(obj, uintptr(value))
}

func ImageButton_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ImageButton_SetHint").Call(obj, GoStrToDStr(value))
}

func ImageButton_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetComponentCount").Call(obj)
	return int32(ret)
}

func ImageButton_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageButton_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ImageButton_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageButton_SetComponentIndex").Call(obj, uintptr(value))
}

func ImageButton_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetOwner").Call(obj)
	return ret
}

func ImageButton_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageButton_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ImageButton_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ImageButton_SetName").Call(obj, GoStrToDStr(value))
}

func ImageButton_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ImageButton_GetTag").Call(obj)
	return int(ret)
}

func ImageButton_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ImageButton_SetTag").Call(obj, uintptr(value))
}

func ImageButton_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAnchorSideLeft").Call(obj)
	return ret
}

func ImageButton_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetAnchorSideLeft").Call(obj, value)
}

func ImageButton_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAnchorSideTop").Call(obj)
	return ret
}

func ImageButton_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetAnchorSideTop").Call(obj, value)
}

func ImageButton_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAnchorSideRight").Call(obj)
	return ret
}

func ImageButton_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetAnchorSideRight").Call(obj, value)
}

func ImageButton_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAnchorSideBottom").Call(obj)
	return ret
}

func ImageButton_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetAnchorSideBottom").Call(obj, value)
}

func ImageButton_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetBorderSpacing").Call(obj)
	return ret
}

func ImageButton_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ImageButton_SetBorderSpacing").Call(obj, value)
}

func ImageButton_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ImageButton_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ImageButton_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ImageButton_StaticClassType() TClass {
	r, _, _ := getLazyProc("ImageButton_StaticClassType").Call()
	return TClass(r)
}
