package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TMiniWebview ---------------------------

func MiniWebview_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_Create").Call(obj)
	return ret
}

func MiniWebview_Free(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Free").Call(obj)
}

func MiniWebview_Navigate(obj uintptr, AURL string) {
	_, _, _ = getLazyProc("MiniWebview_Navigate").Call(obj, GoStrToDStr(AURL))
}

func MiniWebview_GoBack(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_GoBack").Call(obj)
}

func MiniWebview_GoForward(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_GoForward").Call(obj)
}

func MiniWebview_GoHome(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_GoHome").Call(obj)
}

func MiniWebview_GoSearch(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_GoSearch").Call(obj)
}

func MiniWebview_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Refresh").Call(obj)
}

func MiniWebview_Stop(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Stop").Call(obj)
}

func MiniWebview_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("MiniWebview_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func MiniWebview_ExecuteScript(obj uintptr, AScriptText string, AScriptType string) string {
	ret, _, _ := getLazyProc("MiniWebview_ExecuteScript").Call(obj, GoStrToDStr(AScriptText), GoStrToDStr(AScriptType))
	return DStrToGoStr(ret)
}

func MiniWebview_ExecuteJS(obj uintptr, AScriptText string) string {
	ret, _, _ := getLazyProc("MiniWebview_ExecuteJS").Call(obj, GoStrToDStr(AScriptText))
	return DStrToGoStr(ret)
}

func MiniWebview_LoadHTML(obj uintptr, AStr string) {
	_, _, _ = getLazyProc("MiniWebview_LoadHTML").Call(obj, GoStrToDStr(AStr))
}

func MiniWebview_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func MiniWebview_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func MiniWebview_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_DisableAlign").Call(obj)
}

func MiniWebview_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_EnableAlign").Call(obj)
}

func MiniWebview_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func MiniWebview_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("MiniWebview_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func MiniWebview_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("MiniWebview_InsertControl").Call(obj, AControl)
}

func MiniWebview_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Invalidate").Call(obj)
}

func MiniWebview_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("MiniWebview_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func MiniWebview_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("MiniWebview_RemoveControl").Call(obj, AControl)
}

func MiniWebview_Realign(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Realign").Call(obj)
}

func MiniWebview_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Repaint").Call(obj)
}

func MiniWebview_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("MiniWebview_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func MiniWebview_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("MiniWebview_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func MiniWebview_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetFocus").Call(obj)
}

func MiniWebview_Update(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Update").Call(obj)
}

func MiniWebview_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_BringToFront").Call(obj)
}

func MiniWebview_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("MiniWebview_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("MiniWebview_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_Hide(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Hide").Call(obj)
}

func MiniWebview_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("MiniWebview_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func MiniWebview_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("MiniWebview_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("MiniWebview_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SendToBack").Call(obj)
}

func MiniWebview_Show(obj uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Show").Call(obj)
}

func MiniWebview_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("MiniWebview_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func MiniWebview_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetTextLen").Call(obj)
	return int32(ret)
}

func MiniWebview_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("MiniWebview_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func MiniWebview_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func MiniWebview_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("MiniWebview_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func MiniWebview_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("MiniWebview_Assign").Call(obj, Source)
}

func MiniWebview_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("MiniWebview_ClassType").Call(obj)
	return TClass(ret)
}

func MiniWebview_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("MiniWebview_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func MiniWebview_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_InstanceSize").Call(obj)
	return int32(ret)
}

func MiniWebview_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("MiniWebview_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MiniWebview_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetHashCode").Call(obj)
	return int32(ret)
}

func MiniWebview_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("MiniWebview_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func MiniWebview_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("MiniWebview_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MiniWebview_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("MiniWebview_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func MiniWebview_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("MiniWebview_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func MiniWebview_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("MiniWebview_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func MiniWebview_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("MiniWebview_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func MiniWebview_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("MiniWebview_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func MiniWebview_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("MiniWebview_AnchorClient").Call(obj, uintptr(ASpace))
}

func MiniWebview_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("MiniWebview_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func MiniWebview_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("MiniWebview_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func MiniWebview_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("MiniWebview_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func MiniWebview_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("MiniWebview_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func MiniWebview_GetReadyState(obj uintptr) TReadyState {
	ret, _, _ := getLazyProc("MiniWebview_GetReadyState").Call(obj)
	return TReadyState(ret)
}

func MiniWebview_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("MiniWebview_GetAlign").Call(obj)
	return TAlign(ret)
}

func MiniWebview_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("MiniWebview_SetAlign").Call(obj, uintptr(value))
}

func MiniWebview_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func MiniWebview_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("MiniWebview_SetAnchors").Call(obj, uintptr(value))
}

func MiniWebview_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetConstraints").Call(obj)
	return ret
}

func MiniWebview_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetConstraints").Call(obj, value)
}

func MiniWebview_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetVisible").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_SetOnTitleChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("MiniWebview_SetOnTitleChange").Call(obj, addEventToMap(obj, fn))
}

func MiniWebview_SetOnJSExternal(obj uintptr, fn any) {
	_, _, _ = getLazyProc("MiniWebview_SetOnJSExternal").Call(obj, addEventToMap(obj, fn))
}

func MiniWebview_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetDockClientCount").Call(obj)
	return int32(ret)
}

func MiniWebview_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func MiniWebview_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetBrush").Call(obj)
	return ret
}

func MiniWebview_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetControlCount").Call(obj)
	return int32(ret)
}

func MiniWebview_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MiniWebview_GetHandle").Call(obj)
	return ret
}

func MiniWebview_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("MiniWebview_GetParentWindow").Call(obj)
	return ret
}

func MiniWebview_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("MiniWebview_SetParentWindow").Call(obj, value)
}

func MiniWebview_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("MiniWebview_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func MiniWebview_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("MiniWebview_SetTabOrder").Call(obj, uintptr(value))
}

func MiniWebview_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAction").Call(obj)
	return ret
}

func MiniWebview_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetAction").Call(obj, value)
}

func MiniWebview_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("MiniWebview_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func MiniWebview_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("MiniWebview_SetBiDiMode").Call(obj, uintptr(value))
}

func MiniWebview_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("MiniWebview_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("MiniWebview_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MiniWebview_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetClientHeight").Call(obj)
	return int32(ret)
}

func MiniWebview_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetClientHeight").Call(obj, uintptr(value))
}

func MiniWebview_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("MiniWebview_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("MiniWebview_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MiniWebview_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetClientWidth").Call(obj)
	return int32(ret)
}

func MiniWebview_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetClientWidth").Call(obj, uintptr(value))
}

func MiniWebview_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("MiniWebview_GetControlState").Call(obj)
	return TControlState(ret)
}

func MiniWebview_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("MiniWebview_SetControlState").Call(obj, uintptr(value))
}

func MiniWebview_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("MiniWebview_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func MiniWebview_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("MiniWebview_SetControlStyle").Call(obj, uintptr(value))
}

func MiniWebview_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("MiniWebview_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func MiniWebview_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("MiniWebview_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func MiniWebview_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetParent").Call(obj)
	return ret
}

func MiniWebview_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetParent").Call(obj, value)
}

func MiniWebview_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetLeft").Call(obj)
	return int32(ret)
}

func MiniWebview_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetLeft").Call(obj, uintptr(value))
}

func MiniWebview_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetTop").Call(obj)
	return int32(ret)
}

func MiniWebview_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetTop").Call(obj, uintptr(value))
}

func MiniWebview_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetWidth").Call(obj)
	return int32(ret)
}

func MiniWebview_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetWidth").Call(obj, uintptr(value))
}

func MiniWebview_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetHeight").Call(obj)
	return int32(ret)
}

func MiniWebview_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetHeight").Call(obj, uintptr(value))
}

func MiniWebview_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("MiniWebview_GetCursor").Call(obj)
	return TCursor(ret)
}

func MiniWebview_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("MiniWebview_SetCursor").Call(obj, uintptr(value))
}

func MiniWebview_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("MiniWebview_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func MiniWebview_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("MiniWebview_SetHint").Call(obj, GoStrToDStr(value))
}

func MiniWebview_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetComponentCount").Call(obj)
	return int32(ret)
}

func MiniWebview_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("MiniWebview_GetComponentIndex").Call(obj)
	return int32(ret)
}

func MiniWebview_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("MiniWebview_SetComponentIndex").Call(obj, uintptr(value))
}

func MiniWebview_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetOwner").Call(obj)
	return ret
}

func MiniWebview_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("MiniWebview_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func MiniWebview_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("MiniWebview_SetName").Call(obj, GoStrToDStr(value))
}

func MiniWebview_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("MiniWebview_GetTag").Call(obj)
	return int(ret)
}

func MiniWebview_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("MiniWebview_SetTag").Call(obj, uintptr(value))
}

func MiniWebview_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchorSideLeft").Call(obj)
	return ret
}

func MiniWebview_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetAnchorSideLeft").Call(obj, value)
}

func MiniWebview_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchorSideTop").Call(obj)
	return ret
}

func MiniWebview_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetAnchorSideTop").Call(obj, value)
}

func MiniWebview_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchorSideRight").Call(obj)
	return ret
}

func MiniWebview_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetAnchorSideRight").Call(obj, value)
}

func MiniWebview_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchorSideBottom").Call(obj)
	return ret
}

func MiniWebview_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetAnchorSideBottom").Call(obj, value)
}

func MiniWebview_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetChildSizing").Call(obj)
	return ret
}

func MiniWebview_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetChildSizing").Call(obj, value)
}

func MiniWebview_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetBorderSpacing").Call(obj)
	return ret
}

func MiniWebview_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("MiniWebview_SetBorderSpacing").Call(obj, value)
}

func MiniWebview_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func MiniWebview_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetControls").Call(obj, uintptr(Index))
	return ret
}

func MiniWebview_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func MiniWebview_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("MiniWebview_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func MiniWebview_StaticClassType() TClass {
	r, _, _ := getLazyProc("MiniWebview_StaticClassType").Call()
	return TClass(r)
}
