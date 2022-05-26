package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TDirectoryEdit ---------------------------

func DirectoryEdit_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_Create").Call(obj)
	return ret
}

func DirectoryEdit_Free(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Free").Call(obj)
}

func DirectoryEdit_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetFocus").Call(obj)
}

func DirectoryEdit_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_Focused").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_Clear(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Clear").Call(obj)
}

func DirectoryEdit_ClearSelection(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_ClearSelection").Call(obj)
}

func DirectoryEdit_CopyToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_CopyToClipboard").Call(obj)
}

func DirectoryEdit_CutToClipboard(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_CutToClipboard").Call(obj)
}

func DirectoryEdit_PasteFromClipboard(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_PasteFromClipboard").Call(obj)
}

func DirectoryEdit_SelectAll(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SelectAll").Call(obj)
}

func DirectoryEdit_Undo(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Undo").Call(obj)
}

func DirectoryEdit_ValidateEdit(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_ValidateEdit").Call(obj)
}

func DirectoryEdit_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_CanFocus").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_ContainsControl").Call(obj, Control)
	return GoBool(ret)
}

func DirectoryEdit_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), PascalBool(AllowDisabled), PascalBool(AllowWinControls), PascalBool(AllLevels))
	return ret
}

func DirectoryEdit_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_DisableAlign").Call(obj)
}

func DirectoryEdit_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_EnableAlign").Call(obj)
}

func DirectoryEdit_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_FindChildControl").Call(obj, PascalStr(ControlName))
	return ret
}

func DirectoryEdit_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("DirectoryEdit_FlipChildren").Call(obj, PascalBool(AllLevels))
}

func DirectoryEdit_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_HandleAllocated").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_InsertControl").Call(obj, AControl)
}

func DirectoryEdit_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Invalidate").Call(obj)
}

func DirectoryEdit_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("DirectoryEdit_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func DirectoryEdit_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_RemoveControl").Call(obj, AControl)
}

func DirectoryEdit_Realign(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Realign").Call(obj)
}

func DirectoryEdit_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Repaint").Call(obj)
}

func DirectoryEdit_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("DirectoryEdit_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func DirectoryEdit_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("DirectoryEdit_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func DirectoryEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func DirectoryEdit_Update(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Update").Call(obj)
}

func DirectoryEdit_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_BringToFront").Call(obj)
}

func DirectoryEdit_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_Dragging").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_HasParent").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_Hide(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Hide").Call(obj)
}

func DirectoryEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("DirectoryEdit_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func DirectoryEdit_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Refresh").Call(obj)
}

func DirectoryEdit_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SendToBack").Call(obj)
}

func DirectoryEdit_Show(obj uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Show").Call(obj)
}

func DirectoryEdit_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("DirectoryEdit_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func DirectoryEdit_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTextLen").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTextBuf").Call(obj, PascalStr(Buffer))
}

func DirectoryEdit_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_FindComponent").Call(obj, PascalStr(AName))
	return ret
}

func DirectoryEdit_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetNamePath").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_Assign").Call(obj, Source)
}

func DirectoryEdit_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("DirectoryEdit_ClassType").Call(obj)
	return TClass(ret)
}

func DirectoryEdit_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_ClassName").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_InstanceSize").Call(obj)
	return int32(ret)
}

func DirectoryEdit_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_InheritsFrom").Call(obj, uintptr(AClass))
	return GoBool(ret)
}

func DirectoryEdit_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_Equals").Call(obj, Obj)
	return GoBool(ret)
}

func DirectoryEdit_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetHashCode").Call(obj)
	return int32(ret)
}

func DirectoryEdit_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_ToString").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DirectoryEdit_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DirectoryEdit_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func DirectoryEdit_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func DirectoryEdit_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func DirectoryEdit_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func DirectoryEdit_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("DirectoryEdit_AnchorClient").Call(obj, uintptr(ASpace))
}

func DirectoryEdit_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DirectoryEdit_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("DirectoryEdit_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func DirectoryEdit_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("DirectoryEdit_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func DirectoryEdit_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("DirectoryEdit_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func DirectoryEdit_GetAutoSelected(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAutoSelected").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetAutoSelected(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAutoSelected").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetDirectory(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDirectory").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetDirectory(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDirectory").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetRootDir(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetRootDir").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetRootDir(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetRootDir").Call(obj, PascalStr(value))
}

func DirectoryEdit_SetOnAcceptDirectory(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnAcceptDirectory").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_GetDialogTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDialogTitle").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetDialogTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDialogTitle").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetDialogOptions(obj uintptr) TOpenOptions {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDialogOptions").Call(obj)
	return TOpenOptions(ret)
}

func DirectoryEdit_SetDialogOptions(obj uintptr, value TOpenOptions) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDialogOptions").Call(obj, uintptr(value))
}

func DirectoryEdit_GetShowHidden(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetShowHidden").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetShowHidden(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetShowHidden").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetButtonCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetButtonCaption").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetButtonCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetButtonCaption").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetButtonCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DirectoryEdit_GetButtonCursor").Call(obj)
	return TCursor(ret)
}

func DirectoryEdit_SetButtonCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DirectoryEdit_SetButtonCursor").Call(obj, uintptr(value))
}

func DirectoryEdit_GetButtonHint(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetButtonHint").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetButtonHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetButtonHint").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetButtonOnlyWhenFocused(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetButtonOnlyWhenFocused").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetButtonOnlyWhenFocused(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetButtonOnlyWhenFocused").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetButtonWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetButtonWidth").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetButtonWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetButtonWidth").Call(obj, uintptr(value))
}

func DirectoryEdit_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetConstraints").Call(obj)
	return ret
}

func DirectoryEdit_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetConstraints").Call(obj, value)
}

func DirectoryEdit_GetDirectInput(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDirectInput").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetDirectInput(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDirectInput").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetGlyph(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetGlyph").Call(obj)
	return ret
}

func DirectoryEdit_SetGlyph(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetGlyph").Call(obj, value)
}

func DirectoryEdit_GetNumGlyphs(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetNumGlyphs").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetNumGlyphs(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetNumGlyphs").Call(obj, uintptr(value))
}

func DirectoryEdit_GetImages(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetImages").Call(obj)
	return ret
}

func DirectoryEdit_SetImages(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetImages").Call(obj, value)
}

func DirectoryEdit_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetImageIndex").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetImageIndex").Call(obj, uintptr(value))
}

func DirectoryEdit_GetImageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetImageWidth").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetImageWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetImageWidth").Call(obj, uintptr(value))
}

func DirectoryEdit_GetFlat(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetFlat").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetFlat(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetFlat").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetFocusOnButtonClick(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetFocusOnButtonClick").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetFocusOnButtonClick(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetFocusOnButtonClick").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAlign").Call(obj)
	return TAlign(ret)
}

func DirectoryEdit_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAlign").Call(obj, uintptr(value))
}

func DirectoryEdit_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func DirectoryEdit_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAnchors").Call(obj, uintptr(value))
}

func DirectoryEdit_GetAutoSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAutoSize").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetAutoSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAutoSize").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetAutoSelect(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAutoSelect").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetAutoSelect(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAutoSelect").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("DirectoryEdit_GetColor").Call(obj)
	return TColor(ret)
}

func DirectoryEdit_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("DirectoryEdit_SetColor").Call(obj, uintptr(value))
}

func DirectoryEdit_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func DirectoryEdit_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDragCursor").Call(obj, uintptr(value))
}

func DirectoryEdit_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDragMode").Call(obj)
	return TDragMode(ret)
}

func DirectoryEdit_SetDragMode(obj uintptr, value TDragMode) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDragMode").Call(obj, uintptr(value))
}

func DirectoryEdit_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetEnabled").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetEnabled").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetFont").Call(obj)
	return ret
}

func DirectoryEdit_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetFont").Call(obj, value)
}

func DirectoryEdit_GetLayout(obj uintptr) TLeftRight {
	ret, _, _ := getLazyProc("DirectoryEdit_GetLayout").Call(obj)
	return TLeftRight(ret)
}

func DirectoryEdit_SetLayout(obj uintptr, value TLeftRight) {
	_, _, _ = getLazyProc("DirectoryEdit_SetLayout").Call(obj, uintptr(value))
}

func DirectoryEdit_GetMaxLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetMaxLength").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetMaxLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetMaxLength").Call(obj, uintptr(value))
}

func DirectoryEdit_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParentColor").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParentColor").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParentFont").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParentFont").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParentShowHint").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParentShowHint").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetPopupMenu").Call(obj)
	return ret
}

func DirectoryEdit_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetPopupMenu").Call(obj, value)
}

func DirectoryEdit_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetReadOnly").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetReadOnly").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetShowHint").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetShowHint").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func DirectoryEdit_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTabOrder").Call(obj, uintptr(value))
}

func DirectoryEdit_GetSpacing(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetSpacing").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetSpacing(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetSpacing").Call(obj, uintptr(value))
}

func DirectoryEdit_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTabStop").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTabStop").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetVisible").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetVisible").Call(obj, PascalBool(value))
}

func DirectoryEdit_SetOnButtonClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnButtonClick").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnChange(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnContextPopup(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnDblClick(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnDblClick").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnDragDrop(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnDragDrop").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnDragOver(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnDragOver").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnEditingDone(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnEditingDone").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnEndDrag(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnEndDrag").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnExit(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseDown").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseMove").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseUp").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseWheel(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseWheel").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseWheelDown(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseWheelDown").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_SetOnMouseWheelUp(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("DirectoryEdit_SetOnMouseWheelUp").Call(obj, addEventToMap(obj, fn))
}

func DirectoryEdit_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetText").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetText").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetTextHint(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTextHint").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetTextHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTextHint").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetAlignment(obj uintptr) TAlignment {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAlignment").Call(obj)
	return TAlignment(ret)
}

func DirectoryEdit_SetAlignment(obj uintptr, value TAlignment) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAlignment").Call(obj, uintptr(value))
}

func DirectoryEdit_GetCanUndo(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetCanUndo").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_GetCaretPos(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_GetCaretPos").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_SetCaretPos(obj uintptr, value TPoint) {
	_, _, _ = getLazyProc("DirectoryEdit_SetCaretPos").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DirectoryEdit_GetCharCase(obj uintptr) TEditCharCase {
	ret, _, _ := getLazyProc("DirectoryEdit_GetCharCase").Call(obj)
	return TEditCharCase(ret)
}

func DirectoryEdit_SetCharCase(obj uintptr, value TEditCharCase) {
	_, _, _ = getLazyProc("DirectoryEdit_SetCharCase").Call(obj, uintptr(value))
}

func DirectoryEdit_GetHideSelection(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetHideSelection").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetHideSelection(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetHideSelection").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetModified(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetModified").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetModified(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetModified").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetNumbersOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetNumbersOnly").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetNumbersOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetNumbersOnly").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetPasswordChar(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetPasswordChar").Call(obj)
	return uint16(ret)
}

func DirectoryEdit_SetPasswordChar(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("DirectoryEdit_SetPasswordChar").Call(obj, uintptr(value))
}

func DirectoryEdit_GetSelLength(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetSelLength").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetSelLength(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetSelLength").Call(obj, uintptr(value))
}

func DirectoryEdit_GetSelStart(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetSelStart").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetSelStart(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetSelStart").Call(obj, uintptr(value))
}

func DirectoryEdit_GetSelText(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetSelText").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetSelText(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetSelText").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDockClientCount").Call(obj)
	return int32(ret)
}

func DirectoryEdit_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDockSite").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDockSite").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDoubleBuffered").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetDoubleBuffered").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetMouseInClient").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func DirectoryEdit_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetBrush").Call(obj)
	return ret
}

func DirectoryEdit_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetControlCount").Call(obj)
	return int32(ret)
}

func DirectoryEdit_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DirectoryEdit_GetHandle").Call(obj)
	return ret
}

func DirectoryEdit_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParentDoubleBuffered").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParentDoubleBuffered").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParentWindow").Call(obj)
	return ret
}

func DirectoryEdit_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParentWindow").Call(obj, value)
}

func DirectoryEdit_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetShowing").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetUseDockManager").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DirectoryEdit_SetUseDockManager").Call(obj, PascalBool(value))
}

func DirectoryEdit_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAction").Call(obj)
	return ret
}

func DirectoryEdit_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAction").Call(obj, value)
}

func DirectoryEdit_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("DirectoryEdit_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func DirectoryEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("DirectoryEdit_SetBiDiMode").Call(obj, uintptr(value))
}

func DirectoryEdit_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DirectoryEdit_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("DirectoryEdit_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DirectoryEdit_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetClientHeight").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetClientHeight").Call(obj, uintptr(value))
}

func DirectoryEdit_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DirectoryEdit_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DirectoryEdit_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DirectoryEdit_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetClientWidth").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetClientWidth").Call(obj, uintptr(value))
}

func DirectoryEdit_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("DirectoryEdit_GetControlState").Call(obj)
	return TControlState(ret)
}

func DirectoryEdit_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("DirectoryEdit_SetControlState").Call(obj, uintptr(value))
}

func DirectoryEdit_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("DirectoryEdit_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func DirectoryEdit_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("DirectoryEdit_SetControlStyle").Call(obj, uintptr(value))
}

func DirectoryEdit_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("DirectoryEdit_GetFloating").Call(obj)
	return GoBool(ret)
}

func DirectoryEdit_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetParent").Call(obj)
	return ret
}

func DirectoryEdit_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetParent").Call(obj, value)
}

func DirectoryEdit_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetLeft").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetLeft").Call(obj, uintptr(value))
}

func DirectoryEdit_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTop").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTop").Call(obj, uintptr(value))
}

func DirectoryEdit_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetWidth").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetWidth").Call(obj, uintptr(value))
}

func DirectoryEdit_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetHeight").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetHeight").Call(obj, uintptr(value))
}

func DirectoryEdit_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DirectoryEdit_GetCursor").Call(obj)
	return TCursor(ret)
}

func DirectoryEdit_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DirectoryEdit_SetCursor").Call(obj, uintptr(value))
}

func DirectoryEdit_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetHint").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetHint").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetComponentCount").Call(obj)
	return int32(ret)
}

func DirectoryEdit_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DirectoryEdit_GetComponentIndex").Call(obj)
	return int32(ret)
}

func DirectoryEdit_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DirectoryEdit_SetComponentIndex").Call(obj, uintptr(value))
}

func DirectoryEdit_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetOwner").Call(obj)
	return ret
}

func DirectoryEdit_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("DirectoryEdit_GetName").Call(obj)
	return GoStr(ret)
}

func DirectoryEdit_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("DirectoryEdit_SetName").Call(obj, PascalStr(value))
}

func DirectoryEdit_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("DirectoryEdit_GetTag").Call(obj)
	return int(ret)
}

func DirectoryEdit_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("DirectoryEdit_SetTag").Call(obj, uintptr(value))
}

func DirectoryEdit_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchorSideLeft").Call(obj)
	return ret
}

func DirectoryEdit_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAnchorSideLeft").Call(obj, value)
}

func DirectoryEdit_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchorSideTop").Call(obj)
	return ret
}

func DirectoryEdit_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAnchorSideTop").Call(obj, value)
}

func DirectoryEdit_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchorSideRight").Call(obj)
	return ret
}

func DirectoryEdit_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAnchorSideRight").Call(obj, value)
}

func DirectoryEdit_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchorSideBottom").Call(obj)
	return ret
}

func DirectoryEdit_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetAnchorSideBottom").Call(obj, value)
}

func DirectoryEdit_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetChildSizing").Call(obj)
	return ret
}

func DirectoryEdit_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetChildSizing").Call(obj, value)
}

func DirectoryEdit_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetBorderSpacing").Call(obj)
	return ret
}

func DirectoryEdit_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DirectoryEdit_SetBorderSpacing").Call(obj, value)
}

func DirectoryEdit_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func DirectoryEdit_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetControls").Call(obj, uintptr(Index))
	return ret
}

func DirectoryEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func DirectoryEdit_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("DirectoryEdit_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func DirectoryEdit_StaticClassType() TClass {
	r, _, _ := getLazyProc("DirectoryEdit_StaticClassType").Call()
	return TClass(r)
}
