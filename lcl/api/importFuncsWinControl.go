package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TWinControl ---------------------------

func WinControl_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_Create").Call(obj)
	return ret
}

func WinControl_Free(obj uintptr) {
	getLazyProc("WinControl_Free").Call(obj)
}

func WinControl_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func WinControl_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("WinControl_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func WinControl_DisableAlign(obj uintptr) {
	getLazyProc("WinControl_DisableAlign").Call(obj)
}

func WinControl_EnableAlign(obj uintptr) {
	getLazyProc("WinControl_EnableAlign").Call(obj)
}

func WinControl_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("WinControl_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func WinControl_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("WinControl_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func WinControl_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("WinControl_InsertControl").Call(obj, AControl)
}

func WinControl_Invalidate(obj uintptr) {
	getLazyProc("WinControl_Invalidate").Call(obj)
}

func WinControl_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("WinControl_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func WinControl_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("WinControl_RemoveControl").Call(obj, AControl)
}

func WinControl_Realign(obj uintptr) {
	getLazyProc("WinControl_Realign").Call(obj)
}

func WinControl_Repaint(obj uintptr) {
	getLazyProc("WinControl_Repaint").Call(obj)
}

func WinControl_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("WinControl_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func WinControl_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("WinControl_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func WinControl_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("WinControl_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func WinControl_SetFocus(obj uintptr) {
	getLazyProc("WinControl_SetFocus").Call(obj)
}

func WinControl_Update(obj uintptr) {
	getLazyProc("WinControl_Update").Call(obj)
}

func WinControl_BringToFront(obj uintptr) {
	getLazyProc("WinControl_BringToFront").Call(obj)
}

func WinControl_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("WinControl_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("WinControl_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_Hide(obj uintptr) {
	getLazyProc("WinControl_Hide").Call(obj)
}

func WinControl_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("WinControl_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func WinControl_Refresh(obj uintptr) {
	getLazyProc("WinControl_Refresh").Call(obj)
}

func WinControl_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("WinControl_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("WinControl_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_SendToBack(obj uintptr) {
	getLazyProc("WinControl_SendToBack").Call(obj)
}

func WinControl_Show(obj uintptr) {
	getLazyProc("WinControl_Show").Call(obj)
}

func WinControl_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("WinControl_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func WinControl_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetTextLen").Call(obj)
	return int32(ret)
}

func WinControl_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("WinControl_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func WinControl_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("WinControl_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func WinControl_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("WinControl_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func WinControl_Assign(obj uintptr, Source uintptr) {
	getLazyProc("WinControl_Assign").Call(obj, Source)
}

func WinControl_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("WinControl_ClassType").Call(obj)
	return TClass(ret)
}

func WinControl_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("WinControl_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func WinControl_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_InstanceSize").Call(obj)
	return int32(ret)
}

func WinControl_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("WinControl_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func WinControl_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func WinControl_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetHashCode").Call(obj)
	return int32(ret)
}

func WinControl_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("WinControl_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func WinControl_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("WinControl_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func WinControl_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("WinControl_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func WinControl_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("WinControl_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func WinControl_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("WinControl_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func WinControl_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("WinControl_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func WinControl_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("WinControl_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func WinControl_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("WinControl_AnchorClient").Call(obj, uintptr(ASpace))
}

func WinControl_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("WinControl_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func WinControl_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("WinControl_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func WinControl_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("WinControl_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func WinControl_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("WinControl_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func WinControl_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetDockClientCount").Call(obj)
	return int32(ret)
}

func WinControl_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetDockSite(obj uintptr, value bool) {
	getLazyProc("WinControl_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("WinControl_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func WinControl_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetBrush").Call(obj)
	return ret
}

func WinControl_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetControlCount").Call(obj)
	return int32(ret)
}

func WinControl_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("WinControl_GetHandle").Call(obj)
	return HWND(ret)
}

func WinControl_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("WinControl_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("WinControl_GetParentWindow").Call(obj)
	return HWND(ret)
}

func WinControl_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("WinControl_SetParentWindow").Call(obj, uintptr(value))
}

func WinControl_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("WinControl_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func WinControl_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("WinControl_SetTabOrder").Call(obj, uintptr(value))
}

func WinControl_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetTabStop(obj uintptr, value bool) {
	getLazyProc("WinControl_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("WinControl_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetEnabled(obj uintptr, value bool) {
	getLazyProc("WinControl_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAction").Call(obj)
	return ret
}

func WinControl_SetAction(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetAction").Call(obj, value)
}

func WinControl_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("WinControl_GetAlign").Call(obj)
	return TAlign(ret)
}

func WinControl_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("WinControl_SetAlign").Call(obj, uintptr(value))
}

func WinControl_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("WinControl_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func WinControl_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("WinControl_SetAnchors").Call(obj, uintptr(value))
}

func WinControl_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("WinControl_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func WinControl_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("WinControl_SetBiDiMode").Call(obj, uintptr(value))
}

func WinControl_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("WinControl_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("WinControl_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func WinControl_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetClientHeight").Call(obj)
	return int32(ret)
}

func WinControl_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("WinControl_SetClientHeight").Call(obj, uintptr(value))
}

func WinControl_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("WinControl_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("WinControl_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func WinControl_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetClientWidth").Call(obj)
	return int32(ret)
}

func WinControl_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("WinControl_SetClientWidth").Call(obj, uintptr(value))
}

func WinControl_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetConstraints").Call(obj)
	return ret
}

func WinControl_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetConstraints").Call(obj, value)
}

func WinControl_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("WinControl_GetControlState").Call(obj)
	return TControlState(ret)
}

func WinControl_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("WinControl_SetControlState").Call(obj, uintptr(value))
}

func WinControl_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("WinControl_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func WinControl_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("WinControl_SetControlStyle").Call(obj, uintptr(value))
}

func WinControl_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetShowHint(obj uintptr, value bool) {
	getLazyProc("WinControl_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("WinControl_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func WinControl_SetVisible(obj uintptr, value bool) {
	getLazyProc("WinControl_SetVisible").Call(obj, GoBoolToDBool(value))
}

func WinControl_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetParent").Call(obj)
	return ret
}

func WinControl_SetParent(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetParent").Call(obj, value)
}

func WinControl_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetLeft").Call(obj)
	return int32(ret)
}

func WinControl_SetLeft(obj uintptr, value int32) {
	getLazyProc("WinControl_SetLeft").Call(obj, uintptr(value))
}

func WinControl_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetTop").Call(obj)
	return int32(ret)
}

func WinControl_SetTop(obj uintptr, value int32) {
	getLazyProc("WinControl_SetTop").Call(obj, uintptr(value))
}

func WinControl_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetWidth").Call(obj)
	return int32(ret)
}

func WinControl_SetWidth(obj uintptr, value int32) {
	getLazyProc("WinControl_SetWidth").Call(obj, uintptr(value))
}

func WinControl_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetHeight").Call(obj)
	return int32(ret)
}

func WinControl_SetHeight(obj uintptr, value int32) {
	getLazyProc("WinControl_SetHeight").Call(obj, uintptr(value))
}

func WinControl_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("WinControl_GetCursor").Call(obj)
	return TCursor(ret)
}

func WinControl_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("WinControl_SetCursor").Call(obj, uintptr(value))
}

func WinControl_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("WinControl_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func WinControl_SetHint(obj uintptr, value string) {
	getLazyProc("WinControl_SetHint").Call(obj, GoStrToDStr(value))
}

func WinControl_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetComponentCount").Call(obj)
	return int32(ret)
}

func WinControl_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("WinControl_GetComponentIndex").Call(obj)
	return int32(ret)
}

func WinControl_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("WinControl_SetComponentIndex").Call(obj, uintptr(value))
}

func WinControl_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetOwner").Call(obj)
	return ret
}

func WinControl_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("WinControl_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func WinControl_SetName(obj uintptr, value string) {
	getLazyProc("WinControl_SetName").Call(obj, GoStrToDStr(value))
}

func WinControl_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("WinControl_GetTag").Call(obj)
	return int(ret)
}

func WinControl_SetTag(obj uintptr, value int) {
	getLazyProc("WinControl_SetTag").Call(obj, uintptr(value))
}

func WinControl_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAnchorSideLeft").Call(obj)
	return ret
}

func WinControl_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetAnchorSideLeft").Call(obj, value)
}

func WinControl_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAnchorSideTop").Call(obj)
	return ret
}

func WinControl_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetAnchorSideTop").Call(obj, value)
}

func WinControl_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAnchorSideRight").Call(obj)
	return ret
}

func WinControl_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetAnchorSideRight").Call(obj, value)
}

func WinControl_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAnchorSideBottom").Call(obj)
	return ret
}

func WinControl_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetAnchorSideBottom").Call(obj, value)
}

func WinControl_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetChildSizing").Call(obj)
	return ret
}

func WinControl_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetChildSizing").Call(obj, value)
}

func WinControl_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetBorderSpacing").Call(obj)
	return ret
}

func WinControl_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("WinControl_SetBorderSpacing").Call(obj, value)
}

func WinControl_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func WinControl_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetControls").Call(obj, uintptr(Index))
	return ret
}

func WinControl_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func WinControl_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("WinControl_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func WinControl_StaticClassType() TClass {
	r, _, _ := getLazyProc("WinControl_StaticClassType").Call()
	return TClass(r)
}
