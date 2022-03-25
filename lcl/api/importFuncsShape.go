package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TShape ---------------------------

func Shape_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_Create").Call(obj)
	return ret
}

func Shape_Free(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Free").Call(obj)
}

func Shape_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Shape_BringToFront").Call(obj)
}

func Shape_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Shape_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Shape_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Hide").Call(obj)
}

func Shape_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Invalidate").Call(obj)
}

func Shape_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Shape_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Shape_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Refresh").Call(obj)
}

func Shape_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Repaint").Call(obj)
}

func Shape_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Shape_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Shape_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Shape_SendToBack").Call(obj)
}

func Shape_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Shape_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Shape_Show(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Show").Call(obj)
}

func Shape_Update(obj uintptr) {
	_, _, _ = getLazyProc("Shape_Update").Call(obj)
}

func Shape_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Shape_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Shape_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetTextLen").Call(obj)
	return int32(ret)
}

func Shape_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Shape_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Shape_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Shape_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Shape_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Shape_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Shape_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Shape_Assign").Call(obj, Source)
}

func Shape_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Shape_ClassType").Call(obj)
	return TClass(ret)
}

func Shape_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Shape_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Shape_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_InstanceSize").Call(obj)
	return int32(ret)
}

func Shape_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Shape_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Shape_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Shape_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetHashCode").Call(obj)
	return int32(ret)
}

func Shape_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Shape_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Shape_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Shape_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Shape_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Shape_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Shape_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Shape_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Shape_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Shape_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Shape_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Shape_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Shape_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Shape_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Shape_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Shape_AnchorClient").Call(obj, uintptr(ASpace))
}

func Shape_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Shape_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Shape_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Shape_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Shape_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Shape_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Shape_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Shape_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Shape_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Shape_GetAlign").Call(obj)
	return TAlign(ret)
}

func Shape_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Shape_SetAlign").Call(obj, uintptr(value))
}

func Shape_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Shape_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Shape_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Shape_SetAnchors").Call(obj, uintptr(value))
}

func Shape_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetBrush").Call(obj)
	return ret
}

func Shape_SetBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetBrush").Call(obj, value)
}

func Shape_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Shape_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func Shape_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Shape_SetDragCursor").Call(obj, uintptr(value))
}

func Shape_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Shape_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Shape_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("Shape_SetDragKind").Call(obj, uintptr(value))
}

func Shape_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Shape_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Shape_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Shape_SetDragMode").Call(obj, uintptr(value))
}

func Shape_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Shape_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Shape_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetConstraints").Call(obj)
	return ret
}

func Shape_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetConstraints").Call(obj, value)
}

func Shape_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Shape_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func Shape_GetPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetPen").Call(obj)
	return ret
}

func Shape_SetPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetPen").Call(obj, value)
}

func Shape_GetShape(obj uintptr) TShapeType {
	ret, _, _ := getLazyProc("Shape_GetShape").Call(obj)
	return TShapeType(ret)
}

func Shape_SetShape(obj uintptr, value TShapeType) {
	_, _, _ = getLazyProc("Shape_SetShape").Call(obj, uintptr(value))
}

func Shape_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Shape_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Shape_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Shape_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Shape_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Shape_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Shape_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Shape_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAction").Call(obj)
	return ret
}

func Shape_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetAction").Call(obj, value)
}

func Shape_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Shape_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Shape_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Shape_SetBiDiMode").Call(obj, uintptr(value))
}

func Shape_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Shape_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Shape_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Shape_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetClientHeight").Call(obj)
	return int32(ret)
}

func Shape_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetClientHeight").Call(obj, uintptr(value))
}

func Shape_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Shape_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Shape_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Shape_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetClientWidth").Call(obj)
	return int32(ret)
}

func Shape_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetClientWidth").Call(obj, uintptr(value))
}

func Shape_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Shape_GetControlState").Call(obj)
	return TControlState(ret)
}

func Shape_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Shape_SetControlState").Call(obj, uintptr(value))
}

func Shape_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Shape_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Shape_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Shape_SetControlStyle").Call(obj, uintptr(value))
}

func Shape_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Shape_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Shape_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetParent").Call(obj)
	return ret
}

func Shape_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetParent").Call(obj, value)
}

func Shape_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetLeft").Call(obj)
	return int32(ret)
}

func Shape_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetLeft").Call(obj, uintptr(value))
}

func Shape_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetTop").Call(obj)
	return int32(ret)
}

func Shape_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetTop").Call(obj, uintptr(value))
}

func Shape_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetWidth").Call(obj)
	return int32(ret)
}

func Shape_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetWidth").Call(obj, uintptr(value))
}

func Shape_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetHeight").Call(obj)
	return int32(ret)
}

func Shape_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetHeight").Call(obj, uintptr(value))
}

func Shape_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Shape_GetCursor").Call(obj)
	return TCursor(ret)
}

func Shape_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Shape_SetCursor").Call(obj, uintptr(value))
}

func Shape_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Shape_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Shape_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Shape_SetHint").Call(obj, GoStrToDStr(value))
}

func Shape_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetComponentCount").Call(obj)
	return int32(ret)
}

func Shape_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Shape_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Shape_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Shape_SetComponentIndex").Call(obj, uintptr(value))
}

func Shape_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetOwner").Call(obj)
	return ret
}

func Shape_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Shape_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Shape_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Shape_SetName").Call(obj, GoStrToDStr(value))
}

func Shape_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Shape_GetTag").Call(obj)
	return int(ret)
}

func Shape_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Shape_SetTag").Call(obj, uintptr(value))
}

func Shape_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAnchorSideLeft").Call(obj)
	return ret
}

func Shape_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetAnchorSideLeft").Call(obj, value)
}

func Shape_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAnchorSideTop").Call(obj)
	return ret
}

func Shape_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetAnchorSideTop").Call(obj, value)
}

func Shape_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAnchorSideRight").Call(obj)
	return ret
}

func Shape_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetAnchorSideRight").Call(obj, value)
}

func Shape_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAnchorSideBottom").Call(obj)
	return ret
}

func Shape_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetAnchorSideBottom").Call(obj, value)
}

func Shape_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Shape_GetBorderSpacing").Call(obj)
	return ret
}

func Shape_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Shape_SetBorderSpacing").Call(obj, value)
}

func Shape_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Shape_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Shape_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Shape_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Shape_StaticClassType() TClass {
	r, _, _ := getLazyProc("Shape_StaticClassType").Call()
	return TClass(r)
}
