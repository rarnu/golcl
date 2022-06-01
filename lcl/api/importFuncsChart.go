package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TChart ---------------------------

func Chart_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_Create").Call(obj)
	return ret
}

func Chart_Free(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Free").Call(obj)
}

func Chart_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Chart_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Chart_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowCharts bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Chart_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowCharts), GoBoolToDBool(AllLevels))
	return ret
}

func Chart_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Chart_DisableAlign").Call(obj)
}

func Chart_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Chart_EnableAlign").Call(obj)
}

func Chart_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Chart_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Chart_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("Chart_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Chart_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Chart_InsertControl").Call(obj, AControl)
}

func Chart_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Invalidate").Call(obj)
}

func Chart_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("Chart_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func Chart_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Chart_RemoveControl").Call(obj, AControl)
}

func Chart_Realign(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Realign").Call(obj)
}

func Chart_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Repaint").Call(obj)
}

func Chart_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("Chart_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Chart_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("Chart_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Chart_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Chart_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Chart_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("Chart_SetFocus").Call(obj)
}

func Chart_Update(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Update").Call(obj)
}

func Chart_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Chart_BringToFront").Call(obj)
}

func Chart_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Chart_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Chart_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Hide").Call(obj)
}

func Chart_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Chart_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Chart_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Refresh").Call(obj)
}

func Chart_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Chart_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Chart_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Chart_SendToBack").Call(obj)
}

func Chart_Show(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Show").Call(obj)
}

func Chart_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Chart_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Chart_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetTextLen").Call(obj)
	return int32(ret)
}

func Chart_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Chart_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Chart_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Chart_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Chart_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Chart_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Chart_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Chart_Assign").Call(obj, Source)
}

func Chart_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Chart_ClassType").Call(obj)
	return TClass(ret)
}

func Chart_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Chart_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Chart_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_InstanceSize").Call(obj)
	return int32(ret)
}

func Chart_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Chart_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Chart_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Chart_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetHashCode").Call(obj)
	return int32(ret)
}

func Chart_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Chart_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Chart_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Chart_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Chart_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Chart_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Chart_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Chart_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Chart_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Chart_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Chart_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Chart_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Chart_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Chart_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Chart_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Chart_AnchorClient").Call(obj, uintptr(ASpace))
}

func Chart_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Chart_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Chart_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Chart_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Chart_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Chart_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Chart_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Chart_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Chart_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Chart_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Chart_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Chart_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Chart_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetBrush").Call(obj)
	return ret
}

func Chart_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetControlCount").Call(obj)
	return int32(ret)
}

func Chart_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Chart_GetHandle").Call(obj)
	return ret
}

func Chart_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Chart_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Chart_GetParentWindow").Call(obj)
	return ret
}

func Chart_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Chart_SetParentWindow").Call(obj, value)
}

func Chart_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Chart_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Chart_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("Chart_SetTabOrder").Call(obj, uintptr(value))
}

func Chart_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Chart_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Chart_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Chart_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAction").Call(obj)
	return ret
}

func Chart_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAction").Call(obj, value)
}

func Chart_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Chart_GetAlign").Call(obj)
	return TAlign(ret)
}

func Chart_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Chart_SetAlign").Call(obj, uintptr(value))
}

func Chart_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Chart_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Chart_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Chart_SetAnchors").Call(obj, uintptr(value))
}

func Chart_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Chart_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Chart_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Chart_SetBiDiMode").Call(obj, uintptr(value))
}

func Chart_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Chart_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Chart_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Chart_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetClientHeight").Call(obj)
	return int32(ret)
}

func Chart_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetClientHeight").Call(obj, uintptr(value))
}

func Chart_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Chart_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Chart_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetClientWidth").Call(obj)
	return int32(ret)
}

func Chart_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetClientWidth").Call(obj, uintptr(value))
}

func Chart_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetConstraints").Call(obj)
	return ret
}

func Chart_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetConstraints").Call(obj, value)
}

func Chart_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Chart_GetControlState").Call(obj)
	return TControlState(ret)
}

func Chart_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Chart_SetControlState").Call(obj, uintptr(value))
}

func Chart_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Chart_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Chart_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Chart_SetControlStyle").Call(obj, uintptr(value))
}

func Chart_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Chart_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Chart_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetParent").Call(obj)
	return ret
}

func Chart_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetParent").Call(obj, value)
}

func Chart_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetLeft").Call(obj)
	return int32(ret)
}

func Chart_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetLeft").Call(obj, uintptr(value))
}

func Chart_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetTop").Call(obj)
	return int32(ret)
}

func Chart_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetTop").Call(obj, uintptr(value))
}

func Chart_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetWidth").Call(obj)
	return int32(ret)
}

func Chart_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetWidth").Call(obj, uintptr(value))
}

func Chart_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetHeight").Call(obj)
	return int32(ret)
}

func Chart_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetHeight").Call(obj, uintptr(value))
}

func Chart_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Chart_GetCursor").Call(obj)
	return TCursor(ret)
}

func Chart_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Chart_SetCursor").Call(obj, uintptr(value))
}

func Chart_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Chart_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Chart_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Chart_SetHint").Call(obj, GoStrToDStr(value))
}

func Chart_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetComponentCount").Call(obj)
	return int32(ret)
}

func Chart_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Chart_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetComponentIndex").Call(obj, uintptr(value))
}

func Chart_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetOwner").Call(obj)
	return ret
}

func Chart_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Chart_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Chart_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Chart_SetName").Call(obj, GoStrToDStr(value))
}

func Chart_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Chart_GetTag").Call(obj)
	return int(ret)
}

func Chart_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Chart_SetTag").Call(obj, uintptr(value))
}

func Chart_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAnchorSideLeft").Call(obj)
	return ret
}

func Chart_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAnchorSideLeft").Call(obj, value)
}

func Chart_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAnchorSideTop").Call(obj)
	return ret
}

func Chart_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAnchorSideTop").Call(obj, value)
}

func Chart_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAnchorSideRight").Call(obj)
	return ret
}

func Chart_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAnchorSideRight").Call(obj, value)
}

func Chart_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAnchorSideBottom").Call(obj)
	return ret
}

func Chart_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAnchorSideBottom").Call(obj, value)
}

func Chart_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetChildSizing").Call(obj)
	return ret
}

func Chart_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetChildSizing").Call(obj, value)
}

func Chart_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetBorderSpacing").Call(obj)
	return ret
}

func Chart_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetBorderSpacing").Call(obj, value)
}

func Chart_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Chart_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Chart_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Chart_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Chart_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Chart_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Chart_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Chart_StaticClassType() TClass {
	r, _, _ := getLazyProc("Chart_StaticClassType").Call()
	return TClass(r)
}

func Chart_GetAutoFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetAutoFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetAutoFocus(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetAutoFocus").Call(obj, GoBoolToDBool(value))
}

func Chart_GetAllowPanning(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetAllowPanning").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetAllowPanning(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetAllowPanning").Call(obj, GoBoolToDBool(value))
}

func Chart_GetAllowZoom(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetAllowZoom").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetAllowZoom(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetAllowZoom").Call(obj, GoBoolToDBool(value))
}

func Chart_GetAntialiasingMode(obj uintptr) TChartAntialiasingMode {
	ret, _, _ := getLazyProc("Chart_GetAntialiasingMode").Call(obj)
	return TChartAntialiasingMode(ret)
}

func Chart_SetAntialiasingMode(obj uintptr, value TChartAntialiasingMode) {
	_, _, _ = getLazyProc("Chart_SetAntialiasingMode").Call(obj, uintptr(value))
}

func Chart_GetAxisList(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetAxisList").Call(obj)
	return ret
}

func Chart_SetAxisList(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetAxisList").Call(obj, value)
}

func Chart_GetAxisVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetAxisVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetAxisVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetAxisVisible").Call(obj, GoBoolToDBool(value))
}

func Chart_GetBackColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Chart_GetBackColor").Call(obj)
	return TColor(ret)
}

func Chart_SetBackColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Chart_SetBackColor").Call(obj, uintptr(value))
}

func Chart_GetBottomAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetBottomAxis").Call(obj)
	return ret
}

func Chart_SetBottomAxis(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetBottomAxis").Call(obj, value)
}

func Chart_GetDepth(obj uintptr) TChartDistance {
	ret, _, _ := getLazyProc("Chart_GetDepth").Call(obj)
	return TChartDistance(ret)
}

func Chart_SetDepth(obj uintptr, value TChartDistance) {
	_, _, _ = getLazyProc("Chart_SetDepth").Call(obj, uintptr(value))
}

func Chart_GetExpandPercentage(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetExpandPercentage").Call(obj)
	return int32(ret)
}

func Chart_SetExpandPercentage(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetExpandPercentage").Call(obj, uintptr(value))
}

func Chart_GetExtent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetExtent").Call(obj)
	return ret
}

func Chart_SetExtent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetExtent").Call(obj, value)
}

func Chart_GetExtentSizeLimit(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetExtentSizeLimit").Call(obj)
	return ret
}

func Chart_SetExtentSizeLimit(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetExtentSizeLimit").Call(obj, value /* convert from TChartExtent */)
}

func Chart_GetFoot(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetFoot").Call(obj)
	return ret
}

func Chart_SetFoot(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetFoot").Call(obj, value)
}

func Chart_GetFrame(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetFrame").Call(obj)
	return ret
}

func Chart_SetFrame(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetFrame").Call(obj, value)
}

func Chart_GetGUIConnector(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetGUIConnector").Call(obj)
	return ret
}

func Chart_SetGUIConnector(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetGUIConnector").Call(obj, value)
}

func Chart_GetLeftAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetLeftAxis").Call(obj)
	return ret
}

func Chart_SetLeftAxis(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetLeftAxis").Call(obj, value)
}

func Chart_GetLegend(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetLegend").Call(obj)
	return ret
}

func Chart_SetLegend(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetLegend").Call(obj, value)
}

func Chart_GetMargins(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetMargins").Call(obj)
	return ret
}

func Chart_SetMargins(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetMargins").Call(obj, value)
}

func Chart_GetMarginsExternal(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetMarginsExternal").Call(obj)
	return ret
}

func Chart_SetMarginsExternal(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetMarginsExternal").Call(obj, value)
}

func Chart_GetProportional(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetProportional").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_SetProportional(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Chart_SetProportional").Call(obj, GoBoolToDBool(value))
}

func Chart_GetTitle(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetTitle").Call(obj)
	return ret
}

func Chart_SetTitle(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetTitle").Call(obj, value)
}

func Chart_GetToolset(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetToolset").Call(obj)
	return ret
}

func Chart_SetToolset(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetToolset").Call(obj, value)
}

func Chart_SetOnAfterCustomDrawBackground(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterCustomDrawBackground").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnAfterCustomDrawBackWall(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterCustomDrawBackWall").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnAfterDraw(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterDraw").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnAfterDrawBackground(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterDrawBackground").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnAfterDrawBackWall(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterDrawBackWall").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnAfterPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnAfterPaint").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnBeforeCustomDrawBackground(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnBeforeCustomDrawBackground").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnBeforeDrawBackground(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnBeforeDrawBackground").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnBeforeCustomDrawBackWall(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnBeforeCustomDrawBackWall").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnBeforeDrawBackWall(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnBeforeDrawBackWall").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnDrawLegend(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnDrawLegend").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnExtentChanged(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnExtentChanged").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnExtentChanging(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnExtentChanging").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnExtentValidate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnExtentValidate").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnFullExtentChanged(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnFullExtentChanged").Call(obj, addEventToMap(obj, fn))
}

func Chart_GetLogicalExtent(obj uintptr) TDoubleRect {
	ret, _, _ := getLazyProc("Chart_GetLogicalExtent").Call(obj)
	return *(*TDoubleRect)(unsafe.Pointer(ret))
}

func Chart_SetLogicalExtent(obj uintptr, value TDoubleRect) {
	_, _, _ = getLazyProc("Chart_SetLogicalExtent").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Chart_GetMinDataSpace(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetMinDataSpace").Call(obj)
	return int32(ret)
}

func Chart_SetMinDataSpace(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Chart_SetMinDataSpace").Call(obj, uintptr(value))
}

func Chart_SetOnChartPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnChartPaint").Call(obj, addEventToMap(obj, fn))
}

func Chart_GetRenderingParams(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetRenderingParams").Call(obj)
	return ret
}

func Chart_SetRenderingParams(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetRenderingParams").Call(obj, value)
}

func Chart_LockClipRect(obj uintptr) {
	_, _, _ = getLazyProc("Chart_LockClipRect").Call(obj)
}

func Chart_UnlockClipRect(obj uintptr) {
	_, _, _ = getLazyProc("Chart_UnlockClipRect").Call(obj)
}

func Chart_EraseBackground(obj uintptr, DC HDC) {
	_, _, _ = getLazyProc("Chart_EraseBackground").Call(obj, DC)
}

func Chart_Paint(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Paint").Call(obj)
}

func Chart_SetChildOrder(obj uintptr, Child uintptr, Order int32) {
	_, _, _ = getLazyProc("Chart_SetChildOrder").Call(obj, Child, uintptr(Order))
}

func Chart_DrawLineHoriz(obj uintptr, ADrawer uintptr, AY int32) {
	_, _, _ = getLazyProc("Chart_DrawLineHoriz").Call(obj, ADrawer, uintptr(AY))
}

func Chart_DrawLineVert(obj uintptr, ADrawer uintptr, AX int32) {
	_, _, _ = getLazyProc("Chart_DrawLineVert").Call(obj, ADrawer, uintptr(AX))
}

func Chart_IsPointInViewPort(obj uintptr, AP TDoublePoint) bool {
	ret, _, _ := getLazyProc("Chart_IsPointInViewPort").Call(obj, uintptr(unsafe.Pointer(&AP)))
	return DBoolToGoBool(ret)
}

func Chart_AddSeries(obj uintptr, ASeries uintptr) {
	_, _, _ = getLazyProc("Chart_AddSeries").Call(obj, ASeries)
}

func Chart_ClearSeries(obj uintptr) {
	_, _, _ = getLazyProc("Chart_ClearSeries").Call(obj)
}

func Chart_Clone(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_Clone").Call(obj)
	return ret
}

func Chart_CopyToClipboard(obj uintptr, AClass TRasterImageClass) {
	_, _, _ = getLazyProc("Chart_CopyToClipboard").Call(obj, uintptr(AClass))
}

func Chart_CopyToClipboardBitmap(obj uintptr) {
	_, _, _ = getLazyProc("Chart_CopyToClipboardBitmap").Call(obj)
}

func Chart_DeleteSeries(obj uintptr, ASeries uintptr) {
	_, _, _ = getLazyProc("Chart_DeleteSeries").Call(obj, ASeries)
}

func Chart_DisableRedrawing(obj uintptr) {
	_, _, _ = getLazyProc("Chart_DisableRedrawing").Call(obj)
}

func Chart_Draw(obj uintptr, ADrawer uintptr, ARect TRect) {
	_, _, _ = getLazyProc("Chart_Draw").Call(obj, ADrawer, uintptr(unsafe.Pointer(&ARect)))
}

func Chart_DrawLegendOn(obj uintptr, ACanvas uintptr, ARect *TRect) {
	_, _, _ = getLazyProc("Chart_DrawLegendOn").Call(obj, ACanvas, uintptr(unsafe.Pointer(&ARect)))
}

func Chart_EnableRedrawing(obj uintptr) {
	_, _, _ = getLazyProc("Chart_EnableRedrawing").Call(obj)
}

func Chart_GetAllSeriesAxisLimits(obj uintptr, AAxis uintptr, AMin, AMax *float64) {
	_, _, _ = getLazyProc("Chart_GetAllSeriesAxisLimits").Call(obj, AAxis, uintptr(unsafe.Pointer(&AMin)), uintptr(unsafe.Pointer(&AMax)))
}

func Chart_GetFullExtent(obj uintptr) TDoubleRect {
	ret, _, _ := getLazyProc("Chart_GetFullExtent").Call(obj)
	return *(*TDoubleRect)(unsafe.Pointer(ret))
}

func Chart_GetLegendItems(obj uintptr, AIncludeHidden bool) uintptr {
	ret, _, _ := getLazyProc("Chart_GetLegendItems").Call(obj, GoBoolToDBool(AIncludeHidden))
	return ret
}

func Chart_PaintOnAuxCanvas(obj uintptr, ACanvas uintptr, ARect TRect) {
	_, _, _ = getLazyProc("Chart_PaintOnAuxCanvas").Call(obj, ACanvas, uintptr(unsafe.Pointer(&ARect)))
}

func Chart_PaintOnCanvas(obj uintptr, ACanvas uintptr, ARect TRect) {
	_, _, _ = getLazyProc("Chart_PaintOnCanvas").Call(obj, ACanvas, uintptr(unsafe.Pointer(&ARect)))
}

func Chart_Prepare(obj uintptr) {
	_, _, _ = getLazyProc("Chart_Prepare").Call(obj)
}

func Chart_RemoveSeries(obj uintptr, ASeries uintptr) {
	_, _, _ = getLazyProc("Chart_RemoveSeries").Call(obj, ASeries)
}

func Chart_SaveToBitmapFile(obj uintptr, AFileName string) {
	_, _, _ = getLazyProc("Chart_SaveToBitmapFile").Call(obj, GoStrToDStr(AFileName))
}

func Chart_SaveToFile(obj uintptr, AClass TRasterImageClass, AFileName string) {
	_, _, _ = getLazyProc("Chart_SaveToFile").Call(obj, uintptr(AClass), GoStrToDStr(AFileName))
}

func Chart_SaveToImage(obj uintptr, AClass TRasterImageClass) uintptr {
	ret, _, _ := getLazyProc("Chart_SaveToImage").Call(obj, uintptr(AClass))
	return ret
}

func Chart_StyleChanged(obj uintptr, Sender uintptr) {
	_, _, _ = getLazyProc("Chart_StyleChanged").Call(obj, Sender)
}

func Chart_UsesBuiltinToolset(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_UsesBuiltinToolset").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_ZoomFull(obj uintptr, AImmediateRecalc bool) {
	_, _, _ = getLazyProc("Chart_ZoomFull").Call(obj, GoBoolToDBool(AImmediateRecalc))
}

func Chart_GraphToImage(obj uintptr, AGraphPoint TDoublePoint) TPoint {
	ret, _, _ := getLazyProc("Chart_GraphToImage").Call(obj, uintptr(unsafe.Pointer(&AGraphPoint)))
	return *(*TPoint)(unsafe.Pointer(ret))
}

func Chart_ImageToGraph(obj uintptr, APoint TPoint) TDoublePoint {
	ret, _, _ := getLazyProc("Chart_ImageToGraph").Call(obj, uintptr(unsafe.Pointer(&APoint)))
	return *(*TDoublePoint)(unsafe.Pointer(ret))
}

func Chart_XGraphToImage(obj uintptr, AX float64) int32 {
	ret, _, _ := getLazyProc("Chart_XGraphToImage").Call(obj, uintptr(unsafe.Pointer(&AX)))
	return int32(ret)
}

func Chart_XImageToGraph(obj uintptr, AX int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_XImageToGraph").Call(obj, uintptr(AX), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_YGraphToImage(obj uintptr, AY float64) int32 {
	ret, _, _ := getLazyProc("Chart_YGraphToImage").Call(obj, uintptr(unsafe.Pointer(&AY)))
	return int32(ret)
}

func Chart_YImageToGraph(obj uintptr, AY int32) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_YImageToGraph").Call(obj, uintptr(AY), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetCanvas").Call(obj)
	return ret
}

func Chart_SetCanvas(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Chart_SetCanvas").Call(obj, value)
}

func Chart_GetBorderStyle(obj uintptr) TBorderStyle {
	ret, _, _ := getLazyProc("Chart_GetBorderStyle").Call(obj)
	return TBorderStyle(ret)
}

func Chart_SetBorderStyle(obj uintptr, value TBorderStyle) {
	_, _, _ = getLazyProc("Chart_SetBorderStyle").Call(obj, uintptr(value))
}

func Chart_SetOnPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Chart_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Chart_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func Chart_GetSeries(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetSeries").Call(obj)
	return ret
}

func Chart_GetActiveToolIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetActiveToolIndex").Call(obj)
	return int32(ret)
}

func Chart_GetChartHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetChartHeight").Call(obj)
	return int32(ret)
}

func Chart_GetChartWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetChartWidth").Call(obj)
	return int32(ret)
}

func Chart_GetClipRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Chart_GetClipRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetHorAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetHorAxis").Call(obj)
	return ret
}

func Chart_GetIsZoomed(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetIsZoomed").Call(obj)
	return DBoolToGoBool(ret)
}

func Chart_GetSeriesCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Chart_GetSeriesCount").Call(obj)
	return int32(ret)
}

func Chart_GetVertAxis(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetVertAxis").Call(obj)
	return ret
}

func Chart_GetXGraphMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_GetXGraphMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetXGraphMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_GetXGraphMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetYGraphMax(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_GetYGraphMax").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetYGraphMin(obj uintptr) float64 {
	var ret float64
	_, _, _ = getLazyProc("Chart_GetYGraphMin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Chart_GetDrawer(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Chart_GetDrawer").Call(obj)
	return ret
}

func Chart_GetScalingValid(obj uintptr) bool {
	ret, _, _ := getLazyProc("Chart_GetScalingValid").Call(obj)
	return DBoolToGoBool(ret)
}

/*

procedure Chart_LockClipRect(AObj: TChart); extdecl;
begin
  handleExceptionBegin
  AObj.LockClipRect();
  handleExceptionEnd
end;

procedure Chart_UnlockClipRect(AObj: TChart); extdecl;
begin
  handleExceptionBegin
  AObj.UnlockClipRect();
  handleExceptionEnd
end;

procedure Chart_EraseBackground(AObj: TChart; DC: HDC); extdecl;
begin
  handleExceptionBegin
  AObj.EraseBackground(DC);
  handleExceptionEnd
end;
*/
