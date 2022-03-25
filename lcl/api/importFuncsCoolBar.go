package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCoolBar ---------------------------

func CoolBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_Create").Call(obj)
	return ret
}

func CoolBar_Free(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Free").Call(obj)
}

func CoolBar_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("CoolBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func CoolBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func CoolBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("CoolBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func CoolBar_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_DisableAlign").Call(obj)
}

func CoolBar_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_EnableAlign").Call(obj)
}

func CoolBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("CoolBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func CoolBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Invalidate").Call(obj)
}

func CoolBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("CoolBar_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func CoolBar_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CoolBar_RemoveControl").Call(obj, AControl)
}

func CoolBar_Realign(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Realign").Call(obj)
}

func CoolBar_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Repaint").Call(obj)
}

func CoolBar_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("CoolBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func CoolBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("CoolBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func CoolBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("CoolBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func CoolBar_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetFocus").Call(obj)
}

func CoolBar_Update(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Update").Call(obj)
}

func CoolBar_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_BringToFront").Call(obj)
}

func CoolBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CoolBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CoolBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_Hide(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Hide").Call(obj)
}

func CoolBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("CoolBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func CoolBar_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Refresh").Call(obj)
}

func CoolBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CoolBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CoolBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_SendToBack").Call(obj)
}

func CoolBar_Show(obj uintptr) {
	_, _, _ = getLazyProc("CoolBar_Show").Call(obj)
}

func CoolBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("CoolBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func CoolBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetTextLen").Call(obj)
	return int32(ret)
}

func CoolBar_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("CoolBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func CoolBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("CoolBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func CoolBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBar_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("CoolBar_Assign").Call(obj, Source)
}

func CoolBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CoolBar_ClassType").Call(obj)
	return TClass(ret)
}

func CoolBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_InstanceSize").Call(obj)
	return int32(ret)
}

func CoolBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CoolBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CoolBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CoolBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetHashCode").Call(obj)
	return int32(ret)
}

func CoolBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CoolBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CoolBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CoolBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CoolBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CoolBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func CoolBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CoolBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func CoolBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("CoolBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func CoolBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("CoolBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func CoolBar_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("CoolBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func CoolBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CoolBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CoolBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("CoolBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func CoolBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("CoolBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func CoolBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("CoolBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func CoolBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("CoolBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func CoolBar_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("CoolBar_SetAlign").Call(obj, uintptr(value))
}

func CoolBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("CoolBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func CoolBar_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("CoolBar_SetAnchors").Call(obj, uintptr(value))
}

func CoolBar_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetBandBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("CoolBar_GetBandBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func CoolBar_SetBandBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("CoolBar_SetBandBorderStyle").Call(obj, uintptr(value))
}

func CoolBar_GetBandMaximize(obj uintptr) TCoolBandMaximize {
	ret, _, _ := getLazyProc("CoolBar_GetBandMaximize").Call(obj)
	return TCoolBandMaximize(ret)
}

func CoolBar_SetBandMaximize(obj uintptr, value TCoolBandMaximize) {
	_, _, _ = getLazyProc("CoolBar_SetBandMaximize").Call(obj, uintptr(value))
}

func CoolBar_GetBands(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetBands").Call(obj)
	return ret
}

func CoolBar_SetBands(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetBands").Call(obj, value)
}

func CoolBar_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetBorderWidth").Call(obj)
	return int32(ret)
}

func CoolBar_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetBorderWidth").Call(obj, uintptr(value))
}

func CoolBar_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CoolBar_GetColor").Call(obj)
	return TColor(ret)
}

func CoolBar_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("CoolBar_SetColor").Call(obj, uintptr(value))
}

func CoolBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetConstraints").Call(obj)
	return ret
}

func CoolBar_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetConstraints").Call(obj, value)
}

func CoolBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CoolBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func CoolBar_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CoolBar_SetDragCursor").Call(obj, uintptr(value))
}

func CoolBar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("CoolBar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func CoolBar_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("CoolBar_SetDragKind").Call(obj, uintptr(value))
}

func CoolBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("CoolBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func CoolBar_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("CoolBar_SetDragMode").Call(obj, uintptr(value))
}

func CoolBar_GetEdgeBorders(obj uintptr) TEdgeBorders {
	ret, _, _ := getLazyProc("CoolBar_GetEdgeBorders").Call(obj)
	return TEdgeBorders(ret)
}

func CoolBar_SetEdgeBorders(obj uintptr, value TEdgeBorders) {
	_, _, _ = getLazyProc("CoolBar_SetEdgeBorders").Call(obj, uintptr(value))
}

func CoolBar_GetEdgeInner(obj uintptr) TEdgeStyle {
	ret, _, _ := getLazyProc("CoolBar_GetEdgeInner").Call(obj)
	return TEdgeStyle(ret)
}

func CoolBar_SetEdgeInner(obj uintptr, value TEdgeStyle) {
	_, _, _ = getLazyProc("CoolBar_SetEdgeInner").Call(obj, uintptr(value))
}

func CoolBar_GetEdgeOuter(obj uintptr) TEdgeStyle {
	ret, _, _ := getLazyProc("CoolBar_GetEdgeOuter").Call(obj)
	return TEdgeStyle(ret)
}

func CoolBar_SetEdgeOuter(obj uintptr, value TEdgeStyle) {
	_, _, _ = getLazyProc("CoolBar_SetEdgeOuter").Call(obj, uintptr(value))
}

func CoolBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetFixedSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetFixedSize").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetFixedSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetFixedSize").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetFixedOrder(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetFixedOrder").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetFixedOrder(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetFixedOrder").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetFont").Call(obj)
	return ret
}

func CoolBar_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetFont").Call(obj, value)
}

func CoolBar_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetImages").Call(obj)
	return ret
}

func CoolBar_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetImages").Call(obj, value)
}

func CoolBar_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetBitmap(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetBitmap").Call(obj)
	return ret
}

func CoolBar_SetBitmap(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetBitmap").Call(obj, value)
}

func CoolBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetPopupMenu").Call(obj)
	return ret
}

func CoolBar_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetPopupMenu").Call(obj, value)
}

func CoolBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetShowText(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetShowText").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetShowText(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetShowText").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetVertical(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetVertical").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetVertical(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetVertical").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CoolBar_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnGetSiteInfo(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CoolBar_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func CoolBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func CoolBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func CoolBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetBrush").Call(obj)
	return ret
}

func CoolBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetControlCount").Call(obj)
	return int32(ret)
}

func CoolBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CoolBar_GetHandle").Call(obj)
	return ret
}

func CoolBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CoolBar_GetParentWindow").Call(obj)
	return ret
}

func CoolBar_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("CoolBar_SetParentWindow").Call(obj, value)
}

func CoolBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("CoolBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func CoolBar_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("CoolBar_SetTabOrder").Call(obj, uintptr(value))
}

func CoolBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CoolBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func CoolBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAction").Call(obj)
	return ret
}

func CoolBar_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetAction").Call(obj, value)
}

func CoolBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("CoolBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func CoolBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("CoolBar_SetBiDiMode").Call(obj, uintptr(value))
}

func CoolBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CoolBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("CoolBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CoolBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func CoolBar_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetClientHeight").Call(obj, uintptr(value))
}

func CoolBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CoolBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CoolBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CoolBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func CoolBar_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetClientWidth").Call(obj, uintptr(value))
}

func CoolBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("CoolBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func CoolBar_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("CoolBar_SetControlState").Call(obj, uintptr(value))
}

func CoolBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("CoolBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func CoolBar_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("CoolBar_SetControlStyle").Call(obj, uintptr(value))
}

func CoolBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("CoolBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func CoolBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetParent").Call(obj)
	return ret
}

func CoolBar_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetParent").Call(obj, value)
}

func CoolBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetLeft").Call(obj)
	return int32(ret)
}

func CoolBar_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetLeft").Call(obj, uintptr(value))
}

func CoolBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetTop").Call(obj)
	return int32(ret)
}

func CoolBar_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetTop").Call(obj, uintptr(value))
}

func CoolBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetWidth").Call(obj)
	return int32(ret)
}

func CoolBar_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetWidth").Call(obj, uintptr(value))
}

func CoolBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetHeight").Call(obj)
	return int32(ret)
}

func CoolBar_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetHeight").Call(obj, uintptr(value))
}

func CoolBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CoolBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func CoolBar_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CoolBar_SetCursor").Call(obj, uintptr(value))
}

func CoolBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBar_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("CoolBar_SetHint").Call(obj, GoStrToDStr(value))
}

func CoolBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func CoolBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CoolBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func CoolBar_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CoolBar_SetComponentIndex").Call(obj, uintptr(value))
}

func CoolBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetOwner").Call(obj)
	return ret
}

func CoolBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("CoolBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func CoolBar_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("CoolBar_SetName").Call(obj, GoStrToDStr(value))
}

func CoolBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("CoolBar_GetTag").Call(obj)
	return int(ret)
}

func CoolBar_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("CoolBar_SetTag").Call(obj, uintptr(value))
}

func CoolBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func CoolBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetAnchorSideLeft").Call(obj, value)
}

func CoolBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAnchorSideTop").Call(obj)
	return ret
}

func CoolBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetAnchorSideTop").Call(obj, value)
}

func CoolBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAnchorSideRight").Call(obj)
	return ret
}

func CoolBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetAnchorSideRight").Call(obj, value)
}

func CoolBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func CoolBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetAnchorSideBottom").Call(obj, value)
}

func CoolBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetChildSizing").Call(obj)
	return ret
}

func CoolBar_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetChildSizing").Call(obj, value)
}

func CoolBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetBorderSpacing").Call(obj)
	return ret
}

func CoolBar_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CoolBar_SetBorderSpacing").Call(obj, value)
}

func CoolBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func CoolBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func CoolBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func CoolBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("CoolBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func CoolBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("CoolBar_StaticClassType").Call()
	return TClass(r)
}
