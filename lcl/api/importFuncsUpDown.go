package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TUpDown ---------------------------

func UpDown_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_Create").Call(obj)
	return ret
}

func UpDown_Free(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Free").Call(obj)
}

func UpDown_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func UpDown_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("UpDown_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func UpDown_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_DisableAlign").Call(obj)
}

func UpDown_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_EnableAlign").Call(obj)
}

func UpDown_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("UpDown_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func UpDown_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("UpDown_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func UpDown_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("UpDown_InsertControl").Call(obj, AControl)
}

func UpDown_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Invalidate").Call(obj)
}

func UpDown_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("UpDown_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func UpDown_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("UpDown_RemoveControl").Call(obj, AControl)
}

func UpDown_Realign(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Realign").Call(obj)
}

func UpDown_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Repaint").Call(obj)
}

func UpDown_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("UpDown_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func UpDown_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("UpDown_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func UpDown_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("UpDown_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func UpDown_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_SetFocus").Call(obj)
}

func UpDown_Update(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Update").Call(obj)
}

func UpDown_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_BringToFront").Call(obj)
}

func UpDown_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("UpDown_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("UpDown_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_Hide(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Hide").Call(obj)
}

func UpDown_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("UpDown_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func UpDown_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Refresh").Call(obj)
}

func UpDown_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("UpDown_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("UpDown_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_SendToBack").Call(obj)
}

func UpDown_Show(obj uintptr) {
	_, _, _ = getLazyProc("UpDown_Show").Call(obj)
}

func UpDown_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("UpDown_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func UpDown_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetTextLen").Call(obj)
	return int32(ret)
}

func UpDown_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("UpDown_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func UpDown_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("UpDown_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func UpDown_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("UpDown_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func UpDown_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("UpDown_Assign").Call(obj, Source)
}

func UpDown_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("UpDown_ClassType").Call(obj)
	return TClass(ret)
}

func UpDown_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("UpDown_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func UpDown_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_InstanceSize").Call(obj)
	return int32(ret)
}

func UpDown_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("UpDown_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func UpDown_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func UpDown_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetHashCode").Call(obj)
	return int32(ret)
}

func UpDown_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("UpDown_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func UpDown_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("UpDown_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func UpDown_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("UpDown_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func UpDown_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("UpDown_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func UpDown_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("UpDown_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func UpDown_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("UpDown_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func UpDown_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("UpDown_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func UpDown_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("UpDown_AnchorClient").Call(obj, uintptr(ASpace))
}

func UpDown_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("UpDown_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func UpDown_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("UpDown_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func UpDown_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("UpDown_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func UpDown_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("UpDown_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func UpDown_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("UpDown_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func UpDown_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("UpDown_SetAnchors").Call(obj, uintptr(value))
}

func UpDown_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("UpDown_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func UpDown_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("UpDown_SetHint").Call(obj, GoStrToDStr(value))
}

func UpDown_GetMin(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetMin").Call(obj)
	return int32(ret)
}

func UpDown_SetMin(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetMin").Call(obj, uintptr(value))
}

func UpDown_GetMax(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetMax").Call(obj)
	return int32(ret)
}

func UpDown_SetMax(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetMax").Call(obj, uintptr(value))
}

func UpDown_GetIncrement(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetIncrement").Call(obj)
	return int32(ret)
}

func UpDown_SetIncrement(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetIncrement").Call(obj, uintptr(value))
}

func UpDown_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetConstraints").Call(obj)
	return ret
}

func UpDown_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetConstraints").Call(obj, value)
}

func UpDown_GetOrientation(obj uintptr) TUDOrientation {
	ret, _, _ := getLazyProc("UpDown_GetOrientation").Call(obj)
	return TUDOrientation(ret)
}

func UpDown_SetOrientation(obj uintptr, value TUDOrientation) {
	_, _, _ = getLazyProc("UpDown_SetOrientation").Call(obj, uintptr(value))
}

func UpDown_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetPopupMenu").Call(obj)
	return ret
}

func UpDown_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetPopupMenu").Call(obj, value)
}

func UpDown_GetPosition(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetPosition").Call(obj)
	return int32(ret)
}

func UpDown_SetPosition(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetPosition").Call(obj, uintptr(value))
}

func UpDown_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("UpDown_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func UpDown_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("UpDown_SetTabOrder").Call(obj, uintptr(value))
}

func UpDown_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetVisible").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetWrap(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetWrap").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetWrap(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetWrap").Call(obj, GoBoolToDBool(value))
}

func UpDown_SetOnChanging(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnChanging").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func UpDown_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("UpDown_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func UpDown_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetDockClientCount").Call(obj)
	return int32(ret)
}

func UpDown_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func UpDown_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetBrush").Call(obj)
	return ret
}

func UpDown_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetControlCount").Call(obj)
	return int32(ret)
}

func UpDown_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("UpDown_GetHandle").Call(obj)
	return ret
}

func UpDown_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("UpDown_GetParentWindow").Call(obj)
	return ret
}

func UpDown_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("UpDown_SetParentWindow").Call(obj, value)
}

func UpDown_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("UpDown_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func UpDown_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAction").Call(obj)
	return ret
}

func UpDown_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetAction").Call(obj, value)
}

func UpDown_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("UpDown_GetAlign").Call(obj)
	return TAlign(ret)
}

func UpDown_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("UpDown_SetAlign").Call(obj, uintptr(value))
}

func UpDown_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("UpDown_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func UpDown_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("UpDown_SetBiDiMode").Call(obj, uintptr(value))
}

func UpDown_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("UpDown_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("UpDown_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func UpDown_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetClientHeight").Call(obj)
	return int32(ret)
}

func UpDown_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetClientHeight").Call(obj, uintptr(value))
}

func UpDown_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("UpDown_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("UpDown_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func UpDown_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetClientWidth").Call(obj)
	return int32(ret)
}

func UpDown_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetClientWidth").Call(obj, uintptr(value))
}

func UpDown_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("UpDown_GetControlState").Call(obj)
	return TControlState(ret)
}

func UpDown_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("UpDown_SetControlState").Call(obj, uintptr(value))
}

func UpDown_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("UpDown_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func UpDown_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("UpDown_SetControlStyle").Call(obj, uintptr(value))
}

func UpDown_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("UpDown_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func UpDown_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetParent").Call(obj)
	return ret
}

func UpDown_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetParent").Call(obj, value)
}

func UpDown_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetLeft").Call(obj)
	return int32(ret)
}

func UpDown_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetLeft").Call(obj, uintptr(value))
}

func UpDown_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetTop").Call(obj)
	return int32(ret)
}

func UpDown_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetTop").Call(obj, uintptr(value))
}

func UpDown_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetWidth").Call(obj)
	return int32(ret)
}

func UpDown_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetWidth").Call(obj, uintptr(value))
}

func UpDown_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetHeight").Call(obj)
	return int32(ret)
}

func UpDown_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetHeight").Call(obj, uintptr(value))
}

func UpDown_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("UpDown_GetCursor").Call(obj)
	return TCursor(ret)
}

func UpDown_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("UpDown_SetCursor").Call(obj, uintptr(value))
}

func UpDown_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetComponentCount").Call(obj)
	return int32(ret)
}

func UpDown_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("UpDown_GetComponentIndex").Call(obj)
	return int32(ret)
}

func UpDown_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("UpDown_SetComponentIndex").Call(obj, uintptr(value))
}

func UpDown_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetOwner").Call(obj)
	return ret
}

func UpDown_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("UpDown_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func UpDown_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("UpDown_SetName").Call(obj, GoStrToDStr(value))
}

func UpDown_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("UpDown_GetTag").Call(obj)
	return int(ret)
}

func UpDown_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("UpDown_SetTag").Call(obj, uintptr(value))
}

func UpDown_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAnchorSideLeft").Call(obj)
	return ret
}

func UpDown_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetAnchorSideLeft").Call(obj, value)
}

func UpDown_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAnchorSideTop").Call(obj)
	return ret
}

func UpDown_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetAnchorSideTop").Call(obj, value)
}

func UpDown_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAnchorSideRight").Call(obj)
	return ret
}

func UpDown_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetAnchorSideRight").Call(obj, value)
}

func UpDown_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAnchorSideBottom").Call(obj)
	return ret
}

func UpDown_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetAnchorSideBottom").Call(obj, value)
}

func UpDown_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetChildSizing").Call(obj)
	return ret
}

func UpDown_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetChildSizing").Call(obj, value)
}

func UpDown_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetBorderSpacing").Call(obj)
	return ret
}

func UpDown_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("UpDown_SetBorderSpacing").Call(obj, value)
}

func UpDown_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func UpDown_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetControls").Call(obj, uintptr(Index))
	return ret
}

func UpDown_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func UpDown_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("UpDown_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func UpDown_StaticClassType() TClass {
	r, _, _ := getLazyProc("UpDown_StaticClassType").Call()
	return TClass(r)
}
