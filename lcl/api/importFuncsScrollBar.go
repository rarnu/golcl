package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TScrollBar ---------------------------

func ScrollBar_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_Create").Call(obj)
	return ret
}

func ScrollBar_Free(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Free").Call(obj)
}

func ScrollBar_SetParams(obj uintptr, APosition int32, AMin int32, AMax int32) {
	_, _, _ = getLazyProc("ScrollBar_SetParams").Call(obj, uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func ScrollBar_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func ScrollBar_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func ScrollBar_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_DisableAlign").Call(obj)
}

func ScrollBar_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_EnableAlign").Call(obj)
}

func ScrollBar_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func ScrollBar_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("ScrollBar_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func ScrollBar_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ScrollBar_InsertControl").Call(obj, AControl)
}

func ScrollBar_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Invalidate").Call(obj)
}

func ScrollBar_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("ScrollBar_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func ScrollBar_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("ScrollBar_RemoveControl").Call(obj, AControl)
}

func ScrollBar_Realign(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Realign").Call(obj)
}

func ScrollBar_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Repaint").Call(obj)
}

func ScrollBar_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("ScrollBar_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func ScrollBar_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("ScrollBar_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func ScrollBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("ScrollBar_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func ScrollBar_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetFocus").Call(obj)
}

func ScrollBar_Update(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Update").Call(obj)
}

func ScrollBar_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_BringToFront").Call(obj)
}

func ScrollBar_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ScrollBar_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ScrollBar_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_Hide(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Hide").Call(obj)
}

func ScrollBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("ScrollBar_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func ScrollBar_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Refresh").Call(obj)
}

func ScrollBar_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ScrollBar_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ScrollBar_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SendToBack").Call(obj)
}

func ScrollBar_Show(obj uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Show").Call(obj)
}

func ScrollBar_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("ScrollBar_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func ScrollBar_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetTextLen").Call(obj)
	return int32(ret)
}

func ScrollBar_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("ScrollBar_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func ScrollBar_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ScrollBar_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBar_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBar_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ScrollBar_Assign").Call(obj, Source)
}

func ScrollBar_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ScrollBar_ClassType").Call(obj)
	return TClass(ret)
}

func ScrollBar_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBar_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBar_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_InstanceSize").Call(obj)
	return int32(ret)
}

func ScrollBar_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ScrollBar_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ScrollBar_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetHashCode").Call(obj)
	return int32(ret)
}

func ScrollBar_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBar_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBar_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ScrollBar_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ScrollBar_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("ScrollBar_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func ScrollBar_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ScrollBar_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func ScrollBar_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("ScrollBar_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func ScrollBar_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("ScrollBar_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func ScrollBar_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("ScrollBar_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func ScrollBar_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("ScrollBar_AnchorClient").Call(obj, uintptr(ASpace))
}

func ScrollBar_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("ScrollBar_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func ScrollBar_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("ScrollBar_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func ScrollBar_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("ScrollBar_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func ScrollBar_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("ScrollBar_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func ScrollBar_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("ScrollBar_GetAlign").Call(obj)
	return TAlign(ret)
}

func ScrollBar_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("ScrollBar_SetAlign").Call(obj, uintptr(value))
}

func ScrollBar_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func ScrollBar_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("ScrollBar_SetAnchors").Call(obj, uintptr(value))
}

func ScrollBar_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("ScrollBar_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func ScrollBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("ScrollBar_SetBiDiMode").Call(obj, uintptr(value))
}

func ScrollBar_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetConstraints").Call(obj)
	return ret
}

func ScrollBar_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetConstraints").Call(obj, value)
}

func ScrollBar_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ScrollBar_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ScrollBar_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ScrollBar_SetDragCursor").Call(obj, uintptr(value))
}

func ScrollBar_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("ScrollBar_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func ScrollBar_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("ScrollBar_SetDragKind").Call(obj, uintptr(value))
}

func ScrollBar_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("ScrollBar_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func ScrollBar_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("ScrollBar_SetDragMode").Call(obj, uintptr(value))
}

func ScrollBar_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetKind(obj uintptr) TScrollBarKind {
	ret, _, _ := getLazyProc("ScrollBar_GetKind").Call(obj)
	return TScrollBarKind(ret)
}

func ScrollBar_SetKind(obj uintptr, value TScrollBarKind) {
	_, _, _ = getLazyProc("ScrollBar_SetKind").Call(obj, uintptr(value))
}

func ScrollBar_GetLargeChange(obj uintptr) TScrollBarInc {
	ret, _, _ := getLazyProc("ScrollBar_GetLargeChange").Call(obj)
	return TScrollBarInc(ret)
}

func ScrollBar_SetLargeChange(obj uintptr, value TScrollBarInc) {
	_, _, _ = getLazyProc("ScrollBar_SetLargeChange").Call(obj, uintptr(value))
}

func ScrollBar_GetMax(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetMax").Call(obj)
	return int32(ret)
}

func ScrollBar_SetMax(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetMax").Call(obj, uintptr(value))
}

func ScrollBar_GetMin(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetMin").Call(obj)
	return int32(ret)
}

func ScrollBar_SetMin(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetMin").Call(obj, uintptr(value))
}

func ScrollBar_GetPageSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetPageSize").Call(obj)
	return int32(ret)
}

func ScrollBar_SetPageSize(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetPageSize").Call(obj, uintptr(value))
}

func ScrollBar_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetPopupMenu").Call(obj)
	return ret
}

func ScrollBar_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetPopupMenu").Call(obj, value)
}

func ScrollBar_GetPosition(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetPosition").Call(obj)
	return int32(ret)
}

func ScrollBar_SetPosition(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetPosition").Call(obj, uintptr(value))
}

func ScrollBar_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetSmallChange(obj uintptr) TScrollBarInc {
	ret, _, _ := getLazyProc("ScrollBar_GetSmallChange").Call(obj)
	return TScrollBarInc(ret)
}

func ScrollBar_SetSmallChange(obj uintptr, value TScrollBarInc) {
	_, _, _ = getLazyProc("ScrollBar_SetSmallChange").Call(obj, uintptr(value))
}

func ScrollBar_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("ScrollBar_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func ScrollBar_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("ScrollBar_SetTabOrder").Call(obj, uintptr(value))
}

func ScrollBar_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetVisible").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ScrollBar_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func ScrollBar_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetDockClientCount").Call(obj)
	return int32(ret)
}

func ScrollBar_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func ScrollBar_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetBrush").Call(obj)
	return ret
}

func ScrollBar_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetControlCount").Call(obj)
	return int32(ret)
}

func ScrollBar_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ScrollBar_GetHandle").Call(obj)
	return ret
}

func ScrollBar_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("ScrollBar_GetParentWindow").Call(obj)
	return ret
}

func ScrollBar_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("ScrollBar_SetParentWindow").Call(obj, value)
}

func ScrollBar_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ScrollBar_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func ScrollBar_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAction").Call(obj)
	return ret
}

func ScrollBar_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetAction").Call(obj, value)
}

func ScrollBar_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ScrollBar_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("ScrollBar_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ScrollBar_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetClientHeight").Call(obj)
	return int32(ret)
}

func ScrollBar_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetClientHeight").Call(obj, uintptr(value))
}

func ScrollBar_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ScrollBar_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("ScrollBar_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ScrollBar_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetClientWidth").Call(obj)
	return int32(ret)
}

func ScrollBar_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetClientWidth").Call(obj, uintptr(value))
}

func ScrollBar_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("ScrollBar_GetControlState").Call(obj)
	return TControlState(ret)
}

func ScrollBar_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("ScrollBar_SetControlState").Call(obj, uintptr(value))
}

func ScrollBar_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("ScrollBar_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func ScrollBar_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("ScrollBar_SetControlStyle").Call(obj, uintptr(value))
}

func ScrollBar_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("ScrollBar_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func ScrollBar_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetParent").Call(obj)
	return ret
}

func ScrollBar_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetParent").Call(obj, value)
}

func ScrollBar_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetLeft").Call(obj)
	return int32(ret)
}

func ScrollBar_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetLeft").Call(obj, uintptr(value))
}

func ScrollBar_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetTop").Call(obj)
	return int32(ret)
}

func ScrollBar_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetTop").Call(obj, uintptr(value))
}

func ScrollBar_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetWidth").Call(obj)
	return int32(ret)
}

func ScrollBar_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetWidth").Call(obj, uintptr(value))
}

func ScrollBar_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetHeight").Call(obj)
	return int32(ret)
}

func ScrollBar_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetHeight").Call(obj, uintptr(value))
}

func ScrollBar_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ScrollBar_GetCursor").Call(obj)
	return TCursor(ret)
}

func ScrollBar_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ScrollBar_SetCursor").Call(obj, uintptr(value))
}

func ScrollBar_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBar_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBar_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("ScrollBar_SetHint").Call(obj, GoStrToDStr(value))
}

func ScrollBar_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetComponentCount").Call(obj)
	return int32(ret)
}

func ScrollBar_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ScrollBar_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ScrollBar_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ScrollBar_SetComponentIndex").Call(obj, uintptr(value))
}

func ScrollBar_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetOwner").Call(obj)
	return ret
}

func ScrollBar_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ScrollBar_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ScrollBar_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ScrollBar_SetName").Call(obj, GoStrToDStr(value))
}

func ScrollBar_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ScrollBar_GetTag").Call(obj)
	return int(ret)
}

func ScrollBar_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ScrollBar_SetTag").Call(obj, uintptr(value))
}

func ScrollBar_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchorSideLeft").Call(obj)
	return ret
}

func ScrollBar_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetAnchorSideLeft").Call(obj, value)
}

func ScrollBar_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchorSideTop").Call(obj)
	return ret
}

func ScrollBar_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetAnchorSideTop").Call(obj, value)
}

func ScrollBar_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchorSideRight").Call(obj)
	return ret
}

func ScrollBar_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetAnchorSideRight").Call(obj, value)
}

func ScrollBar_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchorSideBottom").Call(obj)
	return ret
}

func ScrollBar_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetAnchorSideBottom").Call(obj, value)
}

func ScrollBar_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetChildSizing").Call(obj)
	return ret
}

func ScrollBar_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetChildSizing").Call(obj, value)
}

func ScrollBar_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetBorderSpacing").Call(obj)
	return ret
}

func ScrollBar_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("ScrollBar_SetBorderSpacing").Call(obj, value)
}

func ScrollBar_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func ScrollBar_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetControls").Call(obj, uintptr(Index))
	return ret
}

func ScrollBar_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ScrollBar_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("ScrollBar_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func ScrollBar_StaticClassType() TClass {
	r, _, _ := getLazyProc("ScrollBar_StaticClassType").Call()
	return TClass(r)
}
