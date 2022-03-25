package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TTabSheet ---------------------------

func TabSheet_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_Create").Call(obj)
	return ret
}

func TabSheet_Free(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Free").Call(obj)
}

func TabSheet_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func TabSheet_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("TabSheet_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func TabSheet_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_DisableAlign").Call(obj)
}

func TabSheet_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_EnableAlign").Call(obj)
}

func TabSheet_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("TabSheet_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func TabSheet_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("TabSheet_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func TabSheet_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("TabSheet_InsertControl").Call(obj, AControl)
}

func TabSheet_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Invalidate").Call(obj)
}

func TabSheet_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("TabSheet_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func TabSheet_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("TabSheet_RemoveControl").Call(obj, AControl)
}

func TabSheet_Realign(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Realign").Call(obj)
}

func TabSheet_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Repaint").Call(obj)
}

func TabSheet_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("TabSheet_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func TabSheet_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("TabSheet_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func TabSheet_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("TabSheet_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func TabSheet_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetFocus").Call(obj)
}

func TabSheet_Update(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Update").Call(obj)
}

func TabSheet_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_BringToFront").Call(obj)
}

func TabSheet_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TabSheet_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TabSheet_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_Hide(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Hide").Call(obj)
}

func TabSheet_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("TabSheet_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func TabSheet_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Refresh").Call(obj)
}

func TabSheet_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TabSheet_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TabSheet_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_SendToBack").Call(obj)
}

func TabSheet_Show(obj uintptr) {
	_, _, _ = getLazyProc("TabSheet_Show").Call(obj)
}

func TabSheet_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("TabSheet_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func TabSheet_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetTextLen").Call(obj)
	return int32(ret)
}

func TabSheet_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("TabSheet_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func TabSheet_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("TabSheet_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func TabSheet_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TabSheet_Assign").Call(obj, Source)
}

func TabSheet_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TabSheet_ClassType").Call(obj)
	return TClass(ret)
}

func TabSheet_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_InstanceSize").Call(obj)
	return int32(ret)
}

func TabSheet_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TabSheet_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TabSheet_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TabSheet_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetHashCode").Call(obj)
	return int32(ret)
}

func TabSheet_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("TabSheet_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TabSheet_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("TabSheet_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func TabSheet_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("TabSheet_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func TabSheet_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("TabSheet_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func TabSheet_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("TabSheet_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func TabSheet_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("TabSheet_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func TabSheet_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("TabSheet_AnchorClient").Call(obj, uintptr(ASpace))
}

func TabSheet_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("TabSheet_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func TabSheet_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("TabSheet_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func TabSheet_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("TabSheet_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func TabSheet_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("TabSheet_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func TabSheet_GetPageControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetPageControl").Call(obj)
	return ret
}

func TabSheet_SetPageControl(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetPageControl").Call(obj, value)
}

func TabSheet_GetTabIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetTabIndex").Call(obj)
	return int32(ret)
}

func TabSheet_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetBorderWidth").Call(obj)
	return int32(ret)
}

func TabSheet_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetBorderWidth").Call(obj, uintptr(value))
}

func TabSheet_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("TabSheet_SetCaption").Call(obj, GoStrToDStr(value))
}

func TabSheet_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetFont").Call(obj)
	return ret
}

func TabSheet_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetFont").Call(obj, value)
}

func TabSheet_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetHeight").Call(obj)
	return int32(ret)
}

func TabSheet_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetHeight").Call(obj, uintptr(value))
}

func TabSheet_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetImageIndex").Call(obj)
	return int32(ret)
}

func TabSheet_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetImageIndex").Call(obj, uintptr(value))
}

func TabSheet_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetLeft").Call(obj)
	return int32(ret)
}

func TabSheet_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetLeft").Call(obj, uintptr(value))
}

func TabSheet_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetConstraints").Call(obj)
	return ret
}

func TabSheet_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetConstraints").Call(obj, value)
}

func TabSheet_GetPageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetPageIndex").Call(obj)
	return int32(ret)
}

func TabSheet_SetPageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetPageIndex").Call(obj, uintptr(value))
}

func TabSheet_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetPopupMenu").Call(obj)
	return ret
}

func TabSheet_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetPopupMenu").Call(obj, value)
}

func TabSheet_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTabVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetTabVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetTabVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetTabVisible").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetTop").Call(obj)
	return int32(ret)
}

func TabSheet_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetTop").Call(obj, uintptr(value))
}

func TabSheet_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetVisible").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetWidth").Call(obj)
	return int32(ret)
}

func TabSheet_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetWidth").Call(obj, uintptr(value))
}

func TabSheet_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnHide(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnHide").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_SetOnShow(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TabSheet_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func TabSheet_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetDockClientCount").Call(obj)
	return int32(ret)
}

func TabSheet_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func TabSheet_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetBrush").Call(obj)
	return ret
}

func TabSheet_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetControlCount").Call(obj)
	return int32(ret)
}

func TabSheet_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TabSheet_GetHandle").Call(obj)
	return ret
}

func TabSheet_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("TabSheet_GetParentWindow").Call(obj)
	return ret
}

func TabSheet_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("TabSheet_SetParentWindow").Call(obj, value)
}

func TabSheet_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("TabSheet_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func TabSheet_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("TabSheet_SetTabOrder").Call(obj, uintptr(value))
}

func TabSheet_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TabSheet_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAction").Call(obj)
	return ret
}

func TabSheet_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetAction").Call(obj, value)
}

func TabSheet_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("TabSheet_GetAlign").Call(obj)
	return TAlign(ret)
}

func TabSheet_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("TabSheet_SetAlign").Call(obj, uintptr(value))
}

func TabSheet_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("TabSheet_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func TabSheet_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("TabSheet_SetAnchors").Call(obj, uintptr(value))
}

func TabSheet_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("TabSheet_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func TabSheet_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("TabSheet_SetBiDiMode").Call(obj, uintptr(value))
}

func TabSheet_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TabSheet_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("TabSheet_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TabSheet_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetClientHeight").Call(obj)
	return int32(ret)
}

func TabSheet_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetClientHeight").Call(obj, uintptr(value))
}

func TabSheet_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("TabSheet_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TabSheet_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TabSheet_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetClientWidth").Call(obj)
	return int32(ret)
}

func TabSheet_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetClientWidth").Call(obj, uintptr(value))
}

func TabSheet_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("TabSheet_GetControlState").Call(obj)
	return TControlState(ret)
}

func TabSheet_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("TabSheet_SetControlState").Call(obj, uintptr(value))
}

func TabSheet_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("TabSheet_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func TabSheet_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("TabSheet_SetControlStyle").Call(obj, uintptr(value))
}

func TabSheet_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("TabSheet_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func TabSheet_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetParent").Call(obj)
	return ret
}

func TabSheet_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetParent").Call(obj, value)
}

func TabSheet_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("TabSheet_GetCursor").Call(obj)
	return TCursor(ret)
}

func TabSheet_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("TabSheet_SetCursor").Call(obj, uintptr(value))
}

func TabSheet_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("TabSheet_SetHint").Call(obj, GoStrToDStr(value))
}

func TabSheet_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetComponentCount").Call(obj)
	return int32(ret)
}

func TabSheet_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TabSheet_GetComponentIndex").Call(obj)
	return int32(ret)
}

func TabSheet_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TabSheet_SetComponentIndex").Call(obj, uintptr(value))
}

func TabSheet_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetOwner").Call(obj)
	return ret
}

func TabSheet_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TabSheet_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TabSheet_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("TabSheet_SetName").Call(obj, GoStrToDStr(value))
}

func TabSheet_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("TabSheet_GetTag").Call(obj)
	return int(ret)
}

func TabSheet_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("TabSheet_SetTag").Call(obj, uintptr(value))
}

func TabSheet_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAnchorSideLeft").Call(obj)
	return ret
}

func TabSheet_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetAnchorSideLeft").Call(obj, value)
}

func TabSheet_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAnchorSideTop").Call(obj)
	return ret
}

func TabSheet_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetAnchorSideTop").Call(obj, value)
}

func TabSheet_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAnchorSideRight").Call(obj)
	return ret
}

func TabSheet_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetAnchorSideRight").Call(obj, value)
}

func TabSheet_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAnchorSideBottom").Call(obj)
	return ret
}

func TabSheet_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetAnchorSideBottom").Call(obj, value)
}

func TabSheet_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetChildSizing").Call(obj)
	return ret
}

func TabSheet_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetChildSizing").Call(obj, value)
}

func TabSheet_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetBorderSpacing").Call(obj)
	return ret
}

func TabSheet_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TabSheet_SetBorderSpacing").Call(obj, value)
}

func TabSheet_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func TabSheet_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetControls").Call(obj, uintptr(Index))
	return ret
}

func TabSheet_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func TabSheet_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("TabSheet_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func TabSheet_StaticClassType() TClass {
	r, _, _ := getLazyProc("TabSheet_StaticClassType").Call()
	return TClass(r)
}
