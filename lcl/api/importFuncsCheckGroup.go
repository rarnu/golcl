package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCheckGroup ---------------------------

func CheckGroup_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_Create").Call(obj)
	return ret
}

func CheckGroup_Free(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Free").Call(obj)
}

func CheckGroup_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("CheckGroup_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func CheckGroup_Rows(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_Rows").Call(obj)
	return int32(ret)
}

func CheckGroup_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func CheckGroup_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func CheckGroup_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_DisableAlign").Call(obj)
}

func CheckGroup_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_EnableAlign").Call(obj)
}

func CheckGroup_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func CheckGroup_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckGroup_InsertControl").Call(obj, AControl)
}

func CheckGroup_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Invalidate").Call(obj)
}

func CheckGroup_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("CheckGroup_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func CheckGroup_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("CheckGroup_RemoveControl").Call(obj, AControl)
}

func CheckGroup_Realign(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Realign").Call(obj)
}

func CheckGroup_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Repaint").Call(obj)
}

func CheckGroup_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("CheckGroup_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func CheckGroup_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("CheckGroup_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func CheckGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("CheckGroup_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func CheckGroup_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetFocus").Call(obj)
}

func CheckGroup_Update(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Update").Call(obj)
}

func CheckGroup_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_BringToFront").Call(obj)
}

func CheckGroup_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckGroup_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckGroup_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_Hide(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Hide").Call(obj)
}

func CheckGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("CheckGroup_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func CheckGroup_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Refresh").Call(obj)
}

func CheckGroup_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckGroup_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckGroup_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SendToBack").Call(obj)
}

func CheckGroup_Show(obj uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Show").Call(obj)
}

func CheckGroup_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("CheckGroup_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func CheckGroup_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetTextLen").Call(obj)
	return int32(ret)
}

func CheckGroup_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("CheckGroup_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func CheckGroup_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func CheckGroup_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("CheckGroup_Assign").Call(obj, Source)
}

func CheckGroup_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("CheckGroup_ClassType").Call(obj)
	return TClass(ret)
}

func CheckGroup_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_InstanceSize").Call(obj)
	return int32(ret)
}

func CheckGroup_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("CheckGroup_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func CheckGroup_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetHashCode").Call(obj)
	return int32(ret)
}

func CheckGroup_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckGroup_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckGroup_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckGroup_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func CheckGroup_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckGroup_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func CheckGroup_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckGroup_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func CheckGroup_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("CheckGroup_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func CheckGroup_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("CheckGroup_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func CheckGroup_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("CheckGroup_AnchorClient").Call(obj, uintptr(ASpace))
}

func CheckGroup_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("CheckGroup_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func CheckGroup_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("CheckGroup_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func CheckGroup_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("CheckGroup_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func CheckGroup_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("CheckGroup_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func CheckGroup_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("CheckGroup_GetAlign").Call(obj)
	return TAlign(ret)
}

func CheckGroup_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("CheckGroup_SetAlign").Call(obj, uintptr(value))
}

func CheckGroup_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func CheckGroup_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("CheckGroup_SetAnchors").Call(obj, uintptr(value))
}

func CheckGroup_GetAutoFill(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetAutoFill").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetAutoFill(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetAutoFill").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetAutoSize").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetAutoSize").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("CheckGroup_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func CheckGroup_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("CheckGroup_SetBiDiMode").Call(obj, uintptr(value))
}

func CheckGroup_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckGroup_SetCaption").Call(obj, GoStrToDStr(value))
}

func CheckGroup_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetClientHeight").Call(obj)
	return int32(ret)
}

func CheckGroup_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetClientHeight").Call(obj, uintptr(value))
}

func CheckGroup_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetClientWidth").Call(obj)
	return int32(ret)
}

func CheckGroup_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetClientWidth").Call(obj, uintptr(value))
}

func CheckGroup_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("CheckGroup_GetColor").Call(obj)
	return TColor(ret)
}

func CheckGroup_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("CheckGroup_SetColor").Call(obj, uintptr(value))
}

func CheckGroup_GetColumnLayout(obj uintptr) TColumnLayout {
	ret, _, _ := getLazyProc("CheckGroup_GetColumnLayout").Call(obj)
	return TColumnLayout(ret)
}

func CheckGroup_SetColumnLayout(obj uintptr, value TColumnLayout) {
	_, _, _ = getLazyProc("CheckGroup_SetColumnLayout").Call(obj, uintptr(value))
}

func CheckGroup_GetColumns(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetColumns").Call(obj)
	return int32(ret)
}

func CheckGroup_SetColumns(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetColumns").Call(obj, uintptr(value))
}

func CheckGroup_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetConstraints").Call(obj)
	return ret
}

func CheckGroup_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetConstraints").Call(obj, value)
}

func CheckGroup_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckGroup_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func CheckGroup_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckGroup_SetDragCursor").Call(obj, uintptr(value))
}

func CheckGroup_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("CheckGroup_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func CheckGroup_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("CheckGroup_SetDragMode").Call(obj, uintptr(value))
}

func CheckGroup_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetFont").Call(obj)
	return ret
}

func CheckGroup_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetFont").Call(obj, value)
}

func CheckGroup_GetItems(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetItems").Call(obj)
	return ret
}

func CheckGroup_SetItems(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetItems").Call(obj, value)
}

func CheckGroup_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnDblClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnDragDrop(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnDragOver(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnEndDrag(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnItemClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnItemClick").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseMove(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseWheel(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseWheelDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnMouseWheelUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_SetOnResize(obj uintptr, fn any) {
	_, _, _ = getLazyProc("CheckGroup_SetOnResize").Call(obj, addEventToMap(obj, fn))
}

func CheckGroup_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetPopupMenu").Call(obj)
	return ret
}

func CheckGroup_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetPopupMenu").Call(obj, value)
}

func CheckGroup_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("CheckGroup_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func CheckGroup_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("CheckGroup_SetTabOrder").Call(obj, uintptr(value))
}

func CheckGroup_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetVisible").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetParentBackground(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetParentBackground").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetParentBackground(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetParentBackground").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetDockClientCount").Call(obj)
	return int32(ret)
}

func CheckGroup_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func CheckGroup_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetBrush").Call(obj)
	return ret
}

func CheckGroup_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetControlCount").Call(obj)
	return int32(ret)
}

func CheckGroup_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckGroup_GetHandle").Call(obj)
	return ret
}

func CheckGroup_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("CheckGroup_GetParentWindow").Call(obj)
	return ret
}

func CheckGroup_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("CheckGroup_SetParentWindow").Call(obj, value)
}

func CheckGroup_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func CheckGroup_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAction").Call(obj)
	return ret
}

func CheckGroup_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetAction").Call(obj, value)
}

func CheckGroup_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckGroup_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("CheckGroup_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckGroup_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("CheckGroup_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("CheckGroup_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func CheckGroup_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("CheckGroup_GetControlState").Call(obj)
	return TControlState(ret)
}

func CheckGroup_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("CheckGroup_SetControlState").Call(obj, uintptr(value))
}

func CheckGroup_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("CheckGroup_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func CheckGroup_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("CheckGroup_SetControlStyle").Call(obj, uintptr(value))
}

func CheckGroup_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func CheckGroup_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetParent").Call(obj)
	return ret
}

func CheckGroup_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetParent").Call(obj, value)
}

func CheckGroup_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetLeft").Call(obj)
	return int32(ret)
}

func CheckGroup_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetLeft").Call(obj, uintptr(value))
}

func CheckGroup_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetTop").Call(obj)
	return int32(ret)
}

func CheckGroup_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetTop").Call(obj, uintptr(value))
}

func CheckGroup_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetWidth").Call(obj)
	return int32(ret)
}

func CheckGroup_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetWidth").Call(obj, uintptr(value))
}

func CheckGroup_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetHeight").Call(obj)
	return int32(ret)
}

func CheckGroup_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetHeight").Call(obj, uintptr(value))
}

func CheckGroup_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("CheckGroup_GetCursor").Call(obj)
	return TCursor(ret)
}

func CheckGroup_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("CheckGroup_SetCursor").Call(obj, uintptr(value))
}

func CheckGroup_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckGroup_SetHint").Call(obj, GoStrToDStr(value))
}

func CheckGroup_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetComponentCount").Call(obj)
	return int32(ret)
}

func CheckGroup_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("CheckGroup_GetComponentIndex").Call(obj)
	return int32(ret)
}

func CheckGroup_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("CheckGroup_SetComponentIndex").Call(obj, uintptr(value))
}

func CheckGroup_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetOwner").Call(obj)
	return ret
}

func CheckGroup_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("CheckGroup_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func CheckGroup_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("CheckGroup_SetName").Call(obj, GoStrToDStr(value))
}

func CheckGroup_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("CheckGroup_GetTag").Call(obj)
	return int(ret)
}

func CheckGroup_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("CheckGroup_SetTag").Call(obj, uintptr(value))
}

func CheckGroup_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchorSideLeft").Call(obj)
	return ret
}

func CheckGroup_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetAnchorSideLeft").Call(obj, value)
}

func CheckGroup_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchorSideTop").Call(obj)
	return ret
}

func CheckGroup_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetAnchorSideTop").Call(obj, value)
}

func CheckGroup_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchorSideRight").Call(obj)
	return ret
}

func CheckGroup_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetAnchorSideRight").Call(obj, value)
}

func CheckGroup_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchorSideBottom").Call(obj)
	return ret
}

func CheckGroup_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetAnchorSideBottom").Call(obj, value)
}

func CheckGroup_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetChildSizing").Call(obj)
	return ret
}

func CheckGroup_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetChildSizing").Call(obj, value)
}

func CheckGroup_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetBorderSpacing").Call(obj)
	return ret
}

func CheckGroup_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("CheckGroup_SetBorderSpacing").Call(obj, value)
}

func CheckGroup_GetChecked(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetChecked").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckGroup_SetChecked(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetChecked").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckGroup_GetCheckEnabled(obj uintptr, Index int32) bool {
	ret, _, _ := getLazyProc("CheckGroup_GetCheckEnabled").Call(obj, uintptr(Index))
	return DBoolToGoBool(ret)
}

func CheckGroup_SetCheckEnabled(obj uintptr, Index int32, value bool) {
	_, _, _ = getLazyProc("CheckGroup_SetCheckEnabled").Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func CheckGroup_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func CheckGroup_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetControls").Call(obj, uintptr(Index))
	return ret
}

func CheckGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func CheckGroup_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("CheckGroup_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func CheckGroup_StaticClassType() TClass {
	r, _, _ := getLazyProc("CheckGroup_StaticClassType").Call()
	return TClass(r)
}
