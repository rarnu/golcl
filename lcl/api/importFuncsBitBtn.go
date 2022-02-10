package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TBitBtn ---------------------------

func BitBtn_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_Create").Call(obj)
	return ret
}

func BitBtn_Free(obj uintptr) {
	getLazyProc("BitBtn_Free").Call(obj)
}

func BitBtn_Click(obj uintptr) {
	getLazyProc("BitBtn_Click").Call(obj)
}

func BitBtn_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func BitBtn_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("BitBtn_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func BitBtn_DisableAlign(obj uintptr) {
	getLazyProc("BitBtn_DisableAlign").Call(obj)
}

func BitBtn_EnableAlign(obj uintptr) {
	getLazyProc("BitBtn_EnableAlign").Call(obj)
}

func BitBtn_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("BitBtn_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func BitBtn_FlipChildren(obj uintptr, AllLevels bool) {
	getLazyProc("BitBtn_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func BitBtn_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_InsertControl(obj uintptr, AControl uintptr) {
	getLazyProc("BitBtn_InsertControl").Call(obj, AControl)
}

func BitBtn_Invalidate(obj uintptr) {
	getLazyProc("BitBtn_Invalidate").Call(obj)
}

func BitBtn_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	getLazyProc("BitBtn_PaintTo").Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func BitBtn_RemoveControl(obj uintptr, AControl uintptr) {
	getLazyProc("BitBtn_RemoveControl").Call(obj, AControl)
}

func BitBtn_Realign(obj uintptr) {
	getLazyProc("BitBtn_Realign").Call(obj)
}

func BitBtn_Repaint(obj uintptr) {
	getLazyProc("BitBtn_Repaint").Call(obj)
}

func BitBtn_ScaleBy(obj uintptr, M int32, D int32) {
	getLazyProc("BitBtn_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func BitBtn_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	getLazyProc("BitBtn_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func BitBtn_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	getLazyProc("BitBtn_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func BitBtn_SetFocus(obj uintptr) {
	getLazyProc("BitBtn_SetFocus").Call(obj)
}

func BitBtn_Update(obj uintptr) {
	getLazyProc("BitBtn_Update").Call(obj)
}

func BitBtn_BringToFront(obj uintptr) {
	getLazyProc("BitBtn_BringToFront").Call(obj)
}

func BitBtn_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("BitBtn_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("BitBtn_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_Hide(obj uintptr) {
	getLazyProc("BitBtn_Hide").Call(obj)
}

func BitBtn_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("BitBtn_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func BitBtn_Refresh(obj uintptr) {
	getLazyProc("BitBtn_Refresh").Call(obj)
}

func BitBtn_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	getLazyProc("BitBtn_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	getLazyProc("BitBtn_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_SendToBack(obj uintptr) {
	getLazyProc("BitBtn_SendToBack").Call(obj)
}

func BitBtn_Show(obj uintptr) {
	getLazyProc("BitBtn_Show").Call(obj)
}

func BitBtn_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("BitBtn_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func BitBtn_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetTextLen").Call(obj)
	return int32(ret)
}

func BitBtn_SetTextBuf(obj uintptr, Buffer string) {
	getLazyProc("BitBtn_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func BitBtn_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("BitBtn_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func BitBtn_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_Assign(obj uintptr, Source uintptr) {
	getLazyProc("BitBtn_Assign").Call(obj, Source)
}

func BitBtn_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("BitBtn_ClassType").Call(obj)
	return TClass(ret)
}

func BitBtn_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_InstanceSize").Call(obj)
	return int32(ret)
}

func BitBtn_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("BitBtn_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func BitBtn_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func BitBtn_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetHashCode").Call(obj)
	return int32(ret)
}

func BitBtn_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("BitBtn_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func BitBtn_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	getLazyProc("BitBtn_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func BitBtn_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("BitBtn_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func BitBtn_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	getLazyProc("BitBtn_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func BitBtn_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	getLazyProc("BitBtn_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func BitBtn_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	getLazyProc("BitBtn_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func BitBtn_AnchorClient(obj uintptr, ASpace int32) {
	getLazyProc("BitBtn_AnchorClient").Call(obj, uintptr(ASpace))
}

func BitBtn_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("BitBtn_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func BitBtn_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	getLazyProc("BitBtn_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func BitBtn_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	getLazyProc("BitBtn_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func BitBtn_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	getLazyProc("BitBtn_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func BitBtn_GetDefaultCaption(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetDefaultCaption").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetDefaultCaption(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetDefaultCaption").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetGlyphShowMode(obj uintptr) TGlyphShowMode {
	ret, _, _ := getLazyProc("BitBtn_GetGlyphShowMode").Call(obj)
	return TGlyphShowMode(ret)
}

func BitBtn_SetGlyphShowMode(obj uintptr, value TGlyphShowMode) {
	getLazyProc("BitBtn_SetGlyphShowMode").Call(obj, uintptr(value))
}

func BitBtn_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetImageIndex").Call(obj)
	return int32(ret)
}

func BitBtn_SetImageIndex(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetImageIndex").Call(obj, uintptr(value))
}

func BitBtn_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetImages").Call(obj)
	return ret
}

func BitBtn_SetImages(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetImages").Call(obj, value)
}

func BitBtn_GetImageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetImageWidth").Call(obj)
	return int32(ret)
}

func BitBtn_SetImageWidth(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetImageWidth").Call(obj, uintptr(value))
}

func BitBtn_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAction").Call(obj)
	return ret
}

func BitBtn_SetAction(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetAction").Call(obj, value)
}

func BitBtn_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("BitBtn_GetAlign").Call(obj)
	return TAlign(ret)
}

func BitBtn_SetAlign(obj uintptr, value TAlign) {
	getLazyProc("BitBtn_SetAlign").Call(obj, uintptr(value))
}

func BitBtn_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("BitBtn_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func BitBtn_SetAnchors(obj uintptr, value TAnchors) {
	getLazyProc("BitBtn_SetAnchors").Call(obj, uintptr(value))
}

func BitBtn_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("BitBtn_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func BitBtn_SetBiDiMode(obj uintptr, value TBiDiMode) {
	getLazyProc("BitBtn_SetBiDiMode").Call(obj, uintptr(value))
}

func BitBtn_GetCancel(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetCancel").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetCancel(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetCancel").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_SetCaption(obj uintptr, value string) {
	getLazyProc("BitBtn_SetCaption").Call(obj, GoStrToDStr(value))
}

func BitBtn_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetConstraints").Call(obj)
	return ret
}

func BitBtn_SetConstraints(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetConstraints").Call(obj, value)
}

func BitBtn_GetDefault(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetDefault").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetDefault(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetDefault").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetEnabled(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetFont").Call(obj)
	return ret
}

func BitBtn_SetFont(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetFont").Call(obj, value)
}

func BitBtn_GetGlyph(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetGlyph").Call(obj)
	return ret
}

func BitBtn_SetGlyph(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetGlyph").Call(obj, value)
}

func BitBtn_GetKind(obj uintptr) TBitBtnKind {
	ret, _, _ := getLazyProc("BitBtn_GetKind").Call(obj)
	return TBitBtnKind(ret)
}

func BitBtn_SetKind(obj uintptr, value TBitBtnKind) {
	getLazyProc("BitBtn_SetKind").Call(obj, uintptr(value))
}

func BitBtn_GetLayout(obj uintptr) TButtonLayout {
	ret, _, _ := getLazyProc("BitBtn_GetLayout").Call(obj)
	return TButtonLayout(ret)
}

func BitBtn_SetLayout(obj uintptr, value TButtonLayout) {
	getLazyProc("BitBtn_SetLayout").Call(obj, uintptr(value))
}

func BitBtn_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("BitBtn_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func BitBtn_SetModalResult(obj uintptr, value TModalResult) {
	getLazyProc("BitBtn_SetModalResult").Call(obj, uintptr(value))
}

func BitBtn_GetNumGlyphs(obj uintptr) TNumGlyphs {
	ret, _, _ := getLazyProc("BitBtn_GetNumGlyphs").Call(obj)
	return TNumGlyphs(ret)
}

func BitBtn_SetNumGlyphs(obj uintptr, value TNumGlyphs) {
	getLazyProc("BitBtn_SetNumGlyphs").Call(obj, uintptr(value))
}

func BitBtn_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetParentDoubleBuffered(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetParentFont(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetParentShowHint(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetPopupMenu").Call(obj)
	return ret
}

func BitBtn_SetPopupMenu(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetPopupMenu").Call(obj, value)
}

func BitBtn_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetShowHint(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetSpacing").Call(obj)
	return int32(ret)
}

func BitBtn_SetSpacing(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetSpacing").Call(obj, uintptr(value))
}

func BitBtn_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("BitBtn_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func BitBtn_SetTabOrder(obj uintptr, value TTabOrder) {
	getLazyProc("BitBtn_SetTabOrder").Call(obj, uintptr(value))
}

func BitBtn_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetTabStop(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetVisible(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetVisible").Call(obj, GoBoolToDBool(value))
}

func BitBtn_SetOnClick(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnContextPopup(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnDragDrop(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnDragOver(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnEndDrag(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnEnter(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnExit(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnKeyDown(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnKeyPress(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnKeyUp(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnMouseDown(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnMouseEnter(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnMouseLeave(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnMouseMove(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_SetOnMouseUp(obj uintptr, fn interface{}) {
	getLazyProc("BitBtn_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func BitBtn_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetDockClientCount").Call(obj)
	return int32(ret)
}

func BitBtn_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetDockSite(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func BitBtn_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetBrush").Call(obj)
	return ret
}

func BitBtn_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetControlCount").Call(obj)
	return int32(ret)
}

func BitBtn_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("BitBtn_GetHandle").Call(obj)
	return HWND(ret)
}

func BitBtn_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("BitBtn_GetParentWindow").Call(obj)
	return HWND(ret)
}

func BitBtn_SetParentWindow(obj uintptr, value HWND) {
	getLazyProc("BitBtn_SetParentWindow").Call(obj, uintptr(value))
}

func BitBtn_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_SetUseDockManager(obj uintptr, value bool) {
	getLazyProc("BitBtn_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func BitBtn_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("BitBtn_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_SetBoundsRect(obj uintptr, value TRect) {
	getLazyProc("BitBtn_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func BitBtn_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetClientHeight").Call(obj)
	return int32(ret)
}

func BitBtn_SetClientHeight(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetClientHeight").Call(obj, uintptr(value))
}

func BitBtn_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	getLazyProc("BitBtn_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_GetClientRect(obj uintptr) TRect {
	var ret TRect
	getLazyProc("BitBtn_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func BitBtn_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetClientWidth").Call(obj)
	return int32(ret)
}

func BitBtn_SetClientWidth(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetClientWidth").Call(obj, uintptr(value))
}

func BitBtn_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("BitBtn_GetControlState").Call(obj)
	return TControlState(ret)
}

func BitBtn_SetControlState(obj uintptr, value TControlState) {
	getLazyProc("BitBtn_SetControlState").Call(obj, uintptr(value))
}

func BitBtn_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("BitBtn_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func BitBtn_SetControlStyle(obj uintptr, value TControlStyle) {
	getLazyProc("BitBtn_SetControlStyle").Call(obj, uintptr(value))
}

func BitBtn_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("BitBtn_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func BitBtn_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetParent").Call(obj)
	return ret
}

func BitBtn_SetParent(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetParent").Call(obj, value)
}

func BitBtn_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetLeft").Call(obj)
	return int32(ret)
}

func BitBtn_SetLeft(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetLeft").Call(obj, uintptr(value))
}

func BitBtn_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetTop").Call(obj)
	return int32(ret)
}

func BitBtn_SetTop(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetTop").Call(obj, uintptr(value))
}

func BitBtn_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetWidth").Call(obj)
	return int32(ret)
}

func BitBtn_SetWidth(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetWidth").Call(obj, uintptr(value))
}

func BitBtn_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetHeight").Call(obj)
	return int32(ret)
}

func BitBtn_SetHeight(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetHeight").Call(obj, uintptr(value))
}

func BitBtn_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("BitBtn_GetCursor").Call(obj)
	return TCursor(ret)
}

func BitBtn_SetCursor(obj uintptr, value TCursor) {
	getLazyProc("BitBtn_SetCursor").Call(obj, uintptr(value))
}

func BitBtn_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_SetHint(obj uintptr, value string) {
	getLazyProc("BitBtn_SetHint").Call(obj, GoStrToDStr(value))
}

func BitBtn_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetComponentCount").Call(obj)
	return int32(ret)
}

func BitBtn_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("BitBtn_GetComponentIndex").Call(obj)
	return int32(ret)
}

func BitBtn_SetComponentIndex(obj uintptr, value int32) {
	getLazyProc("BitBtn_SetComponentIndex").Call(obj, uintptr(value))
}

func BitBtn_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetOwner").Call(obj)
	return ret
}

func BitBtn_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("BitBtn_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func BitBtn_SetName(obj uintptr, value string) {
	getLazyProc("BitBtn_SetName").Call(obj, GoStrToDStr(value))
}

func BitBtn_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("BitBtn_GetTag").Call(obj)
	return int(ret)
}

func BitBtn_SetTag(obj uintptr, value int) {
	getLazyProc("BitBtn_SetTag").Call(obj, uintptr(value))
}

func BitBtn_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAnchorSideLeft").Call(obj)
	return ret
}

func BitBtn_SetAnchorSideLeft(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetAnchorSideLeft").Call(obj, value)
}

func BitBtn_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAnchorSideTop").Call(obj)
	return ret
}

func BitBtn_SetAnchorSideTop(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetAnchorSideTop").Call(obj, value)
}

func BitBtn_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAnchorSideRight").Call(obj)
	return ret
}

func BitBtn_SetAnchorSideRight(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetAnchorSideRight").Call(obj, value)
}

func BitBtn_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAnchorSideBottom").Call(obj)
	return ret
}

func BitBtn_SetAnchorSideBottom(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetAnchorSideBottom").Call(obj, value)
}

func BitBtn_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetChildSizing").Call(obj)
	return ret
}

func BitBtn_SetChildSizing(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetChildSizing").Call(obj, value)
}

func BitBtn_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetBorderSpacing").Call(obj)
	return ret
}

func BitBtn_SetBorderSpacing(obj uintptr, value uintptr) {
	getLazyProc("BitBtn_SetBorderSpacing").Call(obj, value)
}

func BitBtn_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func BitBtn_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetControls").Call(obj, uintptr(Index))
	return ret
}

func BitBtn_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func BitBtn_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("BitBtn_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func BitBtn_StaticClassType() TClass {
	r, _, _ := getLazyProc("BitBtn_StaticClassType").Call()
	return TClass(r)
}
