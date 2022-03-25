package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TForm ---------------------------

func Form_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_Create").Call(obj)
	return ret
}

func Form_Free(obj uintptr) {
	_, _, _ = getLazyProc("Form_Free").Call(obj)
}

func Form_InheritedWndProc(obj uintptr, TheMessage *TMessage) {
	_, _, _ = getLazyProc("Form_InheritedWndProc").Call(obj, uintptr(unsafe.Pointer(TheMessage)))
}

func Form_EnabledMaximize(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("Form_EnabledMaximize").Call(obj, GoBoolToDBool(AValue))
}

func Form_EnabledMinimize(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("Form_EnabledMinimize").Call(obj, GoBoolToDBool(AValue))
}

func Form_EnabledSystemMenu(obj uintptr, AValue bool) {
	_, _, _ = getLazyProc("Form_EnabledSystemMenu").Call(obj, GoBoolToDBool(AValue))
}

func Form_ScaleForCurrentDpi(obj uintptr) {
	_, _, _ = getLazyProc("Form_ScaleForCurrentDpi").Call(obj)
}

func Form_ScaleForPPI(obj uintptr, ANewPPI int32) {
	_, _, _ = getLazyProc("Form_ScaleForPPI").Call(obj, uintptr(ANewPPI))
}

func Form_ScreenCenter(obj uintptr) {
	_, _, _ = getLazyProc("Form_ScreenCenter").Call(obj)
}

func Form_WorkAreaCenter(obj uintptr) {
	_, _, _ = getLazyProc("Form_WorkAreaCenter").Call(obj)
}

func Form_Cascade(obj uintptr) {
	_, _, _ = getLazyProc("Form_Cascade").Call(obj)
}

func Form_Close(obj uintptr) {
	_, _, _ = getLazyProc("Form_Close").Call(obj)
}

func Form_FocusControl(obj uintptr, Control uintptr) {
	_, _, _ = getLazyProc("Form_FocusControl").Call(obj, Control)
}

func Form_Hide(obj uintptr) {
	_, _, _ = getLazyProc("Form_Hide").Call(obj)
}

func Form_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("Form_SetFocus").Call(obj)
}

func Form_Show(obj uintptr) {
	_, _, _ = getLazyProc("Form_Show").Call(obj)
}

func Form_ShowModal(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_ShowModal").Call(obj)
	return int32(ret)
}

func Form_ScrollInView(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Form_ScrollInView").Call(obj, AControl)
}

func Form_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("Form_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func Form_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("Form_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func Form_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Form_DisableAlign").Call(obj)
}

func Form_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("Form_EnableAlign").Call(obj)
}

func Form_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("Form_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func Form_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("Form_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func Form_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Form_InsertControl").Call(obj, AControl)
}

func Form_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("Form_Invalidate").Call(obj)
}

func Form_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("Form_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func Form_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("Form_RemoveControl").Call(obj, AControl)
}

func Form_Realign(obj uintptr) {
	_, _, _ = getLazyProc("Form_Realign").Call(obj)
}

func Form_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("Form_Repaint").Call(obj)
}

func Form_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("Form_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func Form_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("Form_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func Form_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("Form_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Form_Update(obj uintptr) {
	_, _, _ = getLazyProc("Form_Update").Call(obj)
}

func Form_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("Form_BringToFront").Call(obj)
}

func Form_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Form_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Form_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("Form_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func Form_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Form_Refresh").Call(obj)
}

func Form_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Form_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Form_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("Form_SendToBack").Call(obj)
}

func Form_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("Form_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func Form_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetTextLen").Call(obj)
	return int32(ret)
}

func Form_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Form_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Form_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("Form_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func Form_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Form_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Form_Assign").Call(obj, Source)
}

func Form_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Form_ClassType").Call(obj)
	return TClass(ret)
}

func Form_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Form_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_InstanceSize").Call(obj)
	return int32(ret)
}

func Form_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Form_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Form_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Form_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetHashCode").Call(obj)
	return int32(ret)
}

func Form_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Form_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Form_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Form_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("Form_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func Form_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Form_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func Form_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("Form_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func Form_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("Form_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func Form_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("Form_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func Form_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("Form_AnchorClient").Call(obj, uintptr(ASpace))
}

func Form_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("Form_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func Form_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("Form_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func Form_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("Form_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func Form_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("Form_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func Form_GetAllowDropFiles(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetAllowDropFiles").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetAllowDropFiles(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetAllowDropFiles").Call(obj, GoBoolToDBool(value))
}

func Form_SetOnDropFiles(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDropFiles").Call(obj, addEventToMap(obj, fn))
}

func Form_GetShowInTaskBar(obj uintptr) TShowInTaskbar {
	ret, _, _ := getLazyProc("Form_GetShowInTaskBar").Call(obj)
	return TShowInTaskbar(ret)
}

func Form_SetShowInTaskBar(obj uintptr, value TShowInTaskbar) {
	_, _, _ = getLazyProc("Form_SetShowInTaskBar").Call(obj, uintptr(value))
}

func Form_GetDesignTimePPI(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetDesignTimePPI").Call(obj)
	return int32(ret)
}

func Form_SetDesignTimePPI(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetDesignTimePPI").Call(obj, uintptr(value))
}

func Form_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetAction").Call(obj)
	return ret
}

func Form_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetAction").Call(obj, value)
}

func Form_GetActiveControl(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetActiveControl").Call(obj)
	return ret
}

func Form_SetActiveControl(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetActiveControl").Call(obj, value)
}

func Form_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("Form_GetAlign").Call(obj)
	return TAlign(ret)
}

func Form_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("Form_SetAlign").Call(obj, uintptr(value))
}

func Form_GetAlphaBlend(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetAlphaBlend").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetAlphaBlend(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetAlphaBlend").Call(obj, GoBoolToDBool(value))
}

func Form_GetAlphaBlendValue(obj uintptr) uint8 {
	ret, _, _ := getLazyProc("Form_GetAlphaBlendValue").Call(obj)
	return uint8(ret)
}

func Form_SetAlphaBlendValue(obj uintptr, value uint8) {
	_, _, _ = getLazyProc("Form_SetAlphaBlendValue").Call(obj, uintptr(value))
}

func Form_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("Form_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func Form_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("Form_SetAnchors").Call(obj, uintptr(value))
}

func Form_GetAutoScroll(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetAutoScroll").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetAutoScroll(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetAutoScroll").Call(obj, GoBoolToDBool(value))
}

func Form_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func Form_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("Form_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func Form_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("Form_SetBiDiMode").Call(obj, uintptr(value))
}

func Form_GetBorderIcons(obj uintptr) TBorderIcons {
	ret, _, _ := getLazyProc("Form_GetBorderIcons").Call(obj)
	return TBorderIcons(ret)
}

func Form_SetBorderIcons(obj uintptr, value TBorderIcons) {
	_, _, _ = getLazyProc("Form_SetBorderIcons").Call(obj, uintptr(value))
}

func Form_GetBorderStyle(obj uintptr) TFormBorderStyle {
	ret, _, _ := getLazyProc("Form_GetBorderStyle").Call(obj)
	return TFormBorderStyle(ret)
}

func Form_SetBorderStyle(obj uintptr, value TFormBorderStyle) {
	_, _, _ = getLazyProc("Form_SetBorderStyle").Call(obj, uintptr(value))
}

func Form_GetBorderWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetBorderWidth").Call(obj)
	return int32(ret)
}

func Form_SetBorderWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetBorderWidth").Call(obj, uintptr(value))
}

func Form_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func Form_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("Form_SetCaption").Call(obj, GoStrToDStr(value))
}

func Form_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetClientHeight").Call(obj)
	return int32(ret)
}

func Form_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetClientHeight").Call(obj, uintptr(value))
}

func Form_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetClientWidth").Call(obj)
	return int32(ret)
}

func Form_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetClientWidth").Call(obj, uintptr(value))
}

func Form_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("Form_GetColor").Call(obj)
	return TColor(ret)
}

func Form_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("Form_SetColor").Call(obj, uintptr(value))
}

func Form_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetConstraints").Call(obj)
	return ret
}

func Form_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetConstraints").Call(obj, value)
}

func Form_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func Form_GetDefaultMonitor(obj uintptr) TDefaultMonitor {
	ret, _, _ := getLazyProc("Form_GetDefaultMonitor").Call(obj)
	return TDefaultMonitor(ret)
}

func Form_SetDefaultMonitor(obj uintptr, value TDefaultMonitor) {
	_, _, _ = getLazyProc("Form_SetDefaultMonitor").Call(obj, uintptr(value))
}

func Form_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func Form_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Form_GetDragKind(obj uintptr) TDragKind {
	ret, _, _ := getLazyProc("Form_GetDragKind").Call(obj)
	return TDragKind(ret)
}

func Form_SetDragKind(obj uintptr, value TDragKind) {
	_, _, _ = getLazyProc("Form_SetDragKind").Call(obj, uintptr(value))
}

func Form_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("Form_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func Form_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("Form_SetDragMode").Call(obj, uintptr(value))
}

func Form_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func Form_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func Form_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetFont").Call(obj)
	return ret
}

func Form_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetFont").Call(obj, value)
}

func Form_GetFormStyle(obj uintptr) TFormStyle {
	ret, _, _ := getLazyProc("Form_GetFormStyle").Call(obj)
	return TFormStyle(ret)
}

func Form_SetFormStyle(obj uintptr, value TFormStyle) {
	_, _, _ = getLazyProc("Form_SetFormStyle").Call(obj, uintptr(value))
}

func Form_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetHeight").Call(obj)
	return int32(ret)
}

func Form_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetHeight").Call(obj, uintptr(value))
}

func Form_GetHorzScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetHorzScrollBar").Call(obj)
	return ret
}

func Form_SetHorzScrollBar(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetHorzScrollBar").Call(obj, value)
}

func Form_GetIcon(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetIcon").Call(obj)
	return ret
}

func Form_SetIcon(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetIcon").Call(obj, value)
}

func Form_GetKeyPreview(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetKeyPreview").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetKeyPreview(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetKeyPreview").Call(obj, GoBoolToDBool(value))
}

func Form_GetMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetMenu").Call(obj)
	return ret
}

func Form_SetMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetMenu").Call(obj, value)
}

func Form_GetPixelsPerInch(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetPixelsPerInch").Call(obj)
	return int32(ret)
}

func Form_SetPixelsPerInch(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetPixelsPerInch").Call(obj, uintptr(value))
}

func Form_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetPopupMenu").Call(obj)
	return ret
}

func Form_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetPopupMenu").Call(obj, value)
}

func Form_GetPosition(obj uintptr) TPosition {
	ret, _, _ := getLazyProc("Form_GetPosition").Call(obj)
	return TPosition(ret)
}

func Form_SetPosition(obj uintptr, value TPosition) {
	_, _, _ = getLazyProc("Form_SetPosition").Call(obj, uintptr(value))
}

func Form_GetScaled(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetScaled").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetScaled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetScaled").Call(obj, GoBoolToDBool(value))
}

func Form_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func Form_GetVertScrollBar(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetVertScrollBar").Call(obj)
	return ret
}

func Form_SetVertScrollBar(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetVertScrollBar").Call(obj, value)
}

func Form_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetVisible").Call(obj, GoBoolToDBool(value))
}

func Form_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetWidth").Call(obj)
	return int32(ret)
}

func Form_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetWidth").Call(obj, uintptr(value))
}

func Form_GetWindowState(obj uintptr) TWindowState {
	ret, _, _ := getLazyProc("Form_GetWindowState").Call(obj)
	return TWindowState(ret)
}

func Form_SetWindowState(obj uintptr, value TWindowState) {
	_, _, _ = getLazyProc("Form_SetWindowState").Call(obj, uintptr(value))
}

func Form_SetOnActivate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnActivate").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnAlignPosition(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnAlignPosition").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnClose(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnClose").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnCloseQuery(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnCloseQuery").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnConstrainedResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnConstrainedResize").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDestroy(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDestroy").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDeactivate(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDeactivate").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDockDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDockDrop").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnEndDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnEndDock").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnGetSiteInfo(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnGetSiteInfo").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnHide(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnHide").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnHelp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnHelp").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseWheel(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseWheelDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnMouseWheelUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnPaint(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnPaint").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnShortCut(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnShortCut").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnShow(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnShow").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnStartDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnStartDock").Call(obj, addEventToMap(obj, fn))
}

func Form_SetOnUnDock(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Form_SetOnUnDock").Call(obj, addEventToMap(obj, fn))
}

func Form_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetCanvas").Call(obj)
	return ret
}

func Form_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("Form_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func Form_SetModalResult(obj uintptr, value TModalResult) {
	_, _, _ = getLazyProc("Form_SetModalResult").Call(obj, uintptr(value))
}

func Form_GetMonitor(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetMonitor").Call(obj)
	return ret
}

func Form_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetLeft").Call(obj)
	return int32(ret)
}

func Form_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetLeft").Call(obj, uintptr(value))
}

func Form_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetTop").Call(obj)
	return int32(ret)
}

func Form_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetTop").Call(obj, uintptr(value))
}

func Form_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetDockClientCount").Call(obj)
	return int32(ret)
}

func Form_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func Form_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetBrush").Call(obj)
	return ret
}

func Form_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetControlCount").Call(obj)
	return int32(ret)
}

func Form_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Form_GetHandle").Call(obj)
	return ret
}

func Form_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func Form_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("Form_GetParentWindow").Call(obj)
	return ret
}

func Form_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("Form_SetParentWindow").Call(obj, value)
}

func Form_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("Form_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func Form_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("Form_SetTabOrder").Call(obj, uintptr(value))
}

func Form_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("Form_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func Form_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Form_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("Form_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Form_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("Form_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("Form_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Form_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("Form_GetControlState").Call(obj)
	return TControlState(ret)
}

func Form_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("Form_SetControlState").Call(obj, uintptr(value))
}

func Form_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("Form_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func Form_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("Form_SetControlStyle").Call(obj, uintptr(value))
}

func Form_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("Form_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func Form_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetParent").Call(obj)
	return ret
}

func Form_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetParent").Call(obj, value)
}

func Form_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("Form_GetCursor").Call(obj)
	return TCursor(ret)
}

func Form_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("Form_SetCursor").Call(obj, uintptr(value))
}

func Form_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func Form_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("Form_SetHint").Call(obj, GoStrToDStr(value))
}

func Form_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetComponentCount").Call(obj)
	return int32(ret)
}

func Form_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Form_GetComponentIndex").Call(obj)
	return int32(ret)
}

func Form_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Form_SetComponentIndex").Call(obj, uintptr(value))
}

func Form_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetOwner").Call(obj)
	return ret
}

func Form_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("Form_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func Form_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("Form_SetName").Call(obj, GoStrToDStr(value))
}

func Form_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("Form_GetTag").Call(obj)
	return int(ret)
}

func Form_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("Form_SetTag").Call(obj, uintptr(value))
}

func Form_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetAnchorSideLeft").Call(obj)
	return ret
}

func Form_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetAnchorSideLeft").Call(obj, value)
}

func Form_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetAnchorSideTop").Call(obj)
	return ret
}

func Form_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetAnchorSideTop").Call(obj, value)
}

func Form_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetAnchorSideRight").Call(obj)
	return ret
}

func Form_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetAnchorSideRight").Call(obj, value)
}

func Form_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetAnchorSideBottom").Call(obj)
	return ret
}

func Form_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetAnchorSideBottom").Call(obj, value)
}

func Form_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetChildSizing").Call(obj)
	return ret
}

func Form_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetChildSizing").Call(obj, value)
}

func Form_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Form_GetBorderSpacing").Call(obj)
	return ret
}

func Form_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Form_SetBorderSpacing").Call(obj, value)
}

func Form_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Form_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func Form_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Form_GetControls").Call(obj, uintptr(Index))
	return ret
}

func Form_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("Form_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func Form_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("Form_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func Form_StaticClassType() TClass {
	r, _, _ := getLazyProc("Form_StaticClassType").Call()
	return TClass(r)
}
