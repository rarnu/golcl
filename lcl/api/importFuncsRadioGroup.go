package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TRadioGroup ---------------------------

func RadioGroup_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_Create").Call(obj)
	return ret
}

func RadioGroup_Free(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Free").Call(obj)
}

func RadioGroup_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("RadioGroup_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func RadioGroup_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func RadioGroup_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func RadioGroup_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_DisableAlign").Call(obj)
}

func RadioGroup_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_EnableAlign").Call(obj)
}

func RadioGroup_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func RadioGroup_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("RadioGroup_InsertControl").Call(obj, AControl)
}

func RadioGroup_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Invalidate").Call(obj)
}

func RadioGroup_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("RadioGroup_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func RadioGroup_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("RadioGroup_RemoveControl").Call(obj, AControl)
}

func RadioGroup_Realign(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Realign").Call(obj)
}

func RadioGroup_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Repaint").Call(obj)
}

func RadioGroup_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("RadioGroup_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func RadioGroup_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("RadioGroup_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func RadioGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("RadioGroup_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func RadioGroup_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetFocus").Call(obj)
}

func RadioGroup_Update(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Update").Call(obj)
}

func RadioGroup_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_BringToFront").Call(obj)
}

func RadioGroup_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RadioGroup_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RadioGroup_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_Hide(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Hide").Call(obj)
}

func RadioGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("RadioGroup_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func RadioGroup_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Refresh").Call(obj)
}

func RadioGroup_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RadioGroup_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RadioGroup_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SendToBack").Call(obj)
}

func RadioGroup_Show(obj uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Show").Call(obj)
}

func RadioGroup_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("RadioGroup_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func RadioGroup_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetTextLen").Call(obj)
	return int32(ret)
}

func RadioGroup_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("RadioGroup_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func RadioGroup_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func RadioGroup_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("RadioGroup_Assign").Call(obj, Source)
}

func RadioGroup_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("RadioGroup_ClassType").Call(obj)
	return TClass(ret)
}

func RadioGroup_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_InstanceSize").Call(obj)
	return int32(ret)
}

func RadioGroup_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("RadioGroup_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func RadioGroup_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetHashCode").Call(obj)
	return int32(ret)
}

func RadioGroup_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("RadioGroup_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RadioGroup_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("RadioGroup_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func RadioGroup_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("RadioGroup_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func RadioGroup_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("RadioGroup_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func RadioGroup_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("RadioGroup_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func RadioGroup_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("RadioGroup_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func RadioGroup_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("RadioGroup_AnchorClient").Call(obj, uintptr(ASpace))
}

func RadioGroup_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("RadioGroup_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func RadioGroup_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("RadioGroup_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func RadioGroup_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("RadioGroup_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func RadioGroup_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("RadioGroup_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func RadioGroup_SetOnSelectionChanged(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnSelectionChanged").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("RadioGroup_GetAlign").Call(obj)
	return TAlign(ret)
}

func RadioGroup_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("RadioGroup_SetAlign").Call(obj, uintptr(value))
}

func RadioGroup_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func RadioGroup_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("RadioGroup_SetAnchors").Call(obj, uintptr(value))
}

func RadioGroup_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("RadioGroup_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func RadioGroup_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("RadioGroup_SetBiDiMode").Call(obj, uintptr(value))
}

func RadioGroup_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("RadioGroup_SetCaption").Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("RadioGroup_GetColor").Call(obj)
	return TColor(ret)
}

func RadioGroup_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("RadioGroup_SetColor").Call(obj, uintptr(value))
}

func RadioGroup_GetColumns(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetColumns").Call(obj)
	return int32(ret)
}

func RadioGroup_SetColumns(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetColumns").Call(obj, uintptr(value))
}

func RadioGroup_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RadioGroup_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func RadioGroup_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("RadioGroup_SetDragCursor").Call(obj, uintptr(value))
}

func RadioGroup_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("RadioGroup_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func RadioGroup_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("RadioGroup_SetDragMode").Call(obj, uintptr(value))
}

func RadioGroup_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetFont").Call(obj)
	return ret
}

func RadioGroup_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetFont").Call(obj, value)
}

func RadioGroup_GetItemIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetItemIndex").Call(obj)
	return int32(ret)
}

func RadioGroup_SetItemIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetItemIndex").Call(obj, uintptr(value))
}

func RadioGroup_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetItems").Call(obj)
	return ret
}

func RadioGroup_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetItems").Call(obj, value)
}

func RadioGroup_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetConstraints").Call(obj)
	return ret
}

func RadioGroup_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetConstraints").Call(obj, value)
}

func RadioGroup_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetParentBackground(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetPopupMenu").Call(obj)
	return ret
}

func RadioGroup_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetPopupMenu").Call(obj, value)
}

func RadioGroup_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("RadioGroup_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func RadioGroup_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("RadioGroup_SetTabOrder").Call(obj, uintptr(value))
}

func RadioGroup_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetVisible").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("RadioGroup_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func RadioGroup_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetDockClientCount").Call(obj)
	return int32(ret)
}

func RadioGroup_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func RadioGroup_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetBrush").Call(obj)
	return ret
}

func RadioGroup_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetControlCount").Call(obj)
	return int32(ret)
}

func RadioGroup_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RadioGroup_GetHandle").Call(obj)
	return ret
}

func RadioGroup_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("RadioGroup_GetParentWindow").Call(obj)
	return ret
}

func RadioGroup_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("RadioGroup_SetParentWindow").Call(obj, value)
}

func RadioGroup_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("RadioGroup_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAction").Call(obj)
	return ret
}

func RadioGroup_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetAction").Call(obj, value)
}

func RadioGroup_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("RadioGroup_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("RadioGroup_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioGroup_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetClientHeight").Call(obj)
	return int32(ret)
}

func RadioGroup_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetClientHeight").Call(obj, uintptr(value))
}

func RadioGroup_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("RadioGroup_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("RadioGroup_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func RadioGroup_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetClientWidth").Call(obj)
	return int32(ret)
}

func RadioGroup_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetClientWidth").Call(obj, uintptr(value))
}

func RadioGroup_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("RadioGroup_GetControlState").Call(obj)
	return TControlState(ret)
}

func RadioGroup_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("RadioGroup_SetControlState").Call(obj, uintptr(value))
}

func RadioGroup_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("RadioGroup_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func RadioGroup_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("RadioGroup_SetControlStyle").Call(obj, uintptr(value))
}

func RadioGroup_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("RadioGroup_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func RadioGroup_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetParent").Call(obj)
	return ret
}

func RadioGroup_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetParent").Call(obj, value)
}

func RadioGroup_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetLeft").Call(obj)
	return int32(ret)
}

func RadioGroup_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetLeft").Call(obj, uintptr(value))
}

func RadioGroup_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetTop").Call(obj)
	return int32(ret)
}

func RadioGroup_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetTop").Call(obj, uintptr(value))
}

func RadioGroup_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetWidth").Call(obj)
	return int32(ret)
}

func RadioGroup_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetWidth").Call(obj, uintptr(value))
}

func RadioGroup_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetHeight").Call(obj)
	return int32(ret)
}

func RadioGroup_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetHeight").Call(obj, uintptr(value))
}

func RadioGroup_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("RadioGroup_GetCursor").Call(obj)
	return TCursor(ret)
}

func RadioGroup_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("RadioGroup_SetCursor").Call(obj, uintptr(value))
}

func RadioGroup_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("RadioGroup_SetHint").Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetComponentCount").Call(obj)
	return int32(ret)
}

func RadioGroup_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("RadioGroup_GetComponentIndex").Call(obj)
	return int32(ret)
}

func RadioGroup_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("RadioGroup_SetComponentIndex").Call(obj, uintptr(value))
}

func RadioGroup_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetOwner").Call(obj)
	return ret
}

func RadioGroup_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("RadioGroup_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func RadioGroup_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("RadioGroup_SetName").Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("RadioGroup_GetTag").Call(obj)
	return int(ret)
}

func RadioGroup_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("RadioGroup_SetTag").Call(obj, uintptr(value))
}

func RadioGroup_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchorSideLeft").Call(obj)
	return ret
}

func RadioGroup_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetAnchorSideLeft").Call(obj, value)
}

func RadioGroup_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchorSideTop").Call(obj)
	return ret
}

func RadioGroup_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetAnchorSideTop").Call(obj, value)
}

func RadioGroup_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchorSideRight").Call(obj)
	return ret
}

func RadioGroup_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetAnchorSideRight").Call(obj, value)
}

func RadioGroup_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchorSideBottom").Call(obj)
	return ret
}

func RadioGroup_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetAnchorSideBottom").Call(obj, value)
}

func RadioGroup_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetChildSizing").Call(obj)
	return ret
}

func RadioGroup_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetChildSizing").Call(obj, value)
}

func RadioGroup_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetBorderSpacing").Call(obj)
	return ret
}

func RadioGroup_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("RadioGroup_SetBorderSpacing").Call(obj, value)
}

func RadioGroup_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func RadioGroup_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetControls").Call(obj, uintptr(Index))
	return ret
}

func RadioGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func RadioGroup_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("RadioGroup_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func RadioGroup_StaticClassType() TClass {
	r, _, _ := getLazyProc("RadioGroup_StaticClassType").Call()
	return TClass(r)
}
