package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

//--------------------------- TDateTimePicker ---------------------------

func DateTimePicker_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_Create").Call(obj)
	return ret
}

func DateTimePicker_Free(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Free").Call(obj)
}

func DateTimePicker_DateIsNull(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_DateIsNull").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SelectDate(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SelectDate").Call(obj)
}

func DateTimePicker_SelectTime(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SelectTime").Call(obj)
}

func DateTimePicker_CanFocus(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_CanFocus").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_ContainsControl").Call(obj, Control)
	return DBoolToGoBool(ret)
}

func DateTimePicker_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_ControlAtPos").Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func DateTimePicker_DisableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_DisableAlign").Call(obj)
}

func DateTimePicker_EnableAlign(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_EnableAlign").Call(obj)
}

func DateTimePicker_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_FindChildControl").Call(obj, GoStrToDStr(ControlName))
	return ret
}

func DateTimePicker_FlipChildren(obj uintptr, AllLevels bool) {
	_, _, _ = getLazyProc("DateTimePicker_FlipChildren").Call(obj, GoBoolToDBool(AllLevels))
}

func DateTimePicker_Focused(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_Focused").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_InsertControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_InsertControl").Call(obj, AControl)
}

func DateTimePicker_Invalidate(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Invalidate").Call(obj)
}

func DateTimePicker_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	_, _, _ = getLazyProc("DateTimePicker_PaintTo").Call(obj, DC, uintptr(X), uintptr(Y))
}

func DateTimePicker_RemoveControl(obj uintptr, AControl uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_RemoveControl").Call(obj, AControl)
}

func DateTimePicker_Realign(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Realign").Call(obj)
}

func DateTimePicker_Repaint(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Repaint").Call(obj)
}

func DateTimePicker_ScaleBy(obj uintptr, M int32, D int32) {
	_, _, _ = getLazyProc("DateTimePicker_ScaleBy").Call(obj, uintptr(M), uintptr(D))
}

func DateTimePicker_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	_, _, _ = getLazyProc("DateTimePicker_ScrollBy").Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func DateTimePicker_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetBounds").Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func DateTimePicker_SetFocus(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetFocus").Call(obj)
}

func DateTimePicker_Update(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Update").Call(obj)
}

func DateTimePicker_BringToFront(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_BringToFront").Call(obj)
}

func DateTimePicker_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DateTimePicker_ClientToScreen").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DateTimePicker_ClientToParent").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_Dragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_Dragging").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_Hide(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Hide").Call(obj)
}

func DateTimePicker_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := getLazyProc("DateTimePicker_Perform").Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func DateTimePicker_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Refresh").Call(obj)
}

func DateTimePicker_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DateTimePicker_ScreenToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DateTimePicker_ParentToClient").Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_SendToBack(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SendToBack").Call(obj)
}

func DateTimePicker_Show(obj uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Show").Call(obj)
}

func DateTimePicker_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(BufSize)
	ret, _, _ := getLazyProc("DateTimePicker_GetTextBuf").Call(obj, getBuffPtr(strPtr), uintptr(BufSize))
	getTextBuf(strPtr, Buffer, int(ret))
	return int32(ret)
}

func DateTimePicker_GetTextLen(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetTextLen").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("DateTimePicker_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func DateTimePicker_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func DateTimePicker_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_Assign").Call(obj, Source)
}

func DateTimePicker_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("DateTimePicker_ClassType").Call(obj)
	return TClass(ret)
}

func DateTimePicker_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_InstanceSize").Call(obj)
	return int32(ret)
}

func DateTimePicker_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("DateTimePicker_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func DateTimePicker_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetHashCode").Call(obj)
	return int32(ret)
}

func DateTimePicker_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_AnchorToNeighbour(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorToNeighbour").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DateTimePicker_AnchorParallel(obj uintptr, ASide TAnchorKind, ASpace int32, ASibling uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorParallel").Call(obj, uintptr(ASide), uintptr(ASpace), ASibling)
}

func DateTimePicker_AnchorHorizontalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorHorizontalCenterTo").Call(obj, ASibling)
}

func DateTimePicker_AnchorVerticalCenterTo(obj uintptr, ASibling uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorVerticalCenterTo").Call(obj, ASibling)
}

func DateTimePicker_AnchorSame(obj uintptr, ASide TAnchorKind, ASibling uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorSame").Call(obj, uintptr(ASide), ASibling)
}

func DateTimePicker_AnchorAsAlign(obj uintptr, ATheAlign TAlign, ASpace int32) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorAsAlign").Call(obj, uintptr(ATheAlign), uintptr(ASpace))
}

func DateTimePicker_AnchorClient(obj uintptr, ASpace int32) {
	_, _, _ = getLazyProc("DateTimePicker_AnchorClient").Call(obj, uintptr(ASpace))
}

func DateTimePicker_ScaleDesignToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleDesignToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleFormToDesign(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleFormToDesign").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_Scale96ToForm(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_Scale96ToForm").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleFormTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleFormTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_Scale96ToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_Scale96ToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleFontTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleFontTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleScreenToFont(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleScreenToFont").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleFontToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleFontToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_Scale96ToScreen(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_Scale96ToScreen").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_ScaleScreenTo96(obj uintptr, ASize int32) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_ScaleScreenTo96").Call(obj, uintptr(ASize))
	return int32(ret)
}

func DateTimePicker_AutoAdjustLayout(obj uintptr, AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	_, _, _ = getLazyProc("DateTimePicker_AutoAdjustLayout").Call(obj, uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func DateTimePicker_FixDesignFontsPPI(obj uintptr, ADesignTimePPI int32) {
	_, _, _ = getLazyProc("DateTimePicker_FixDesignFontsPPI").Call(obj, uintptr(ADesignTimePPI))
}

func DateTimePicker_ScaleFontsPPI(obj uintptr, AToPPI int32, AProportion float64) {
	_, _, _ = getLazyProc("DateTimePicker_ScaleFontsPPI").Call(obj, uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

func DateTimePicker_GetArrowShape(obj uintptr) TArrowShape {
	ret, _, _ := getLazyProc("DateTimePicker_GetArrowShape").Call(obj)
	return TArrowShape(ret)
}

func DateTimePicker_SetArrowShape(obj uintptr, value TArrowShape) {
	_, _, _ = getLazyProc("DateTimePicker_SetArrowShape").Call(obj, uintptr(value))
}

func DateTimePicker_GetAutoAdvance(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetAutoAdvance").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetAutoAdvance(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetAutoAdvance").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetAutoButtonSize(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetAutoButtonSize").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetAutoButtonSize(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetAutoButtonSize").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetCascade(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetCascade").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetCascade(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetCascade").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetCenturyFrom(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("DateTimePicker_GetCenturyFrom").Call(obj)
	return uint16(ret)
}

func DateTimePicker_SetCenturyFrom(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("DateTimePicker_SetCenturyFrom").Call(obj, uintptr(value))
}

func DateTimePicker_GetDateDisplayOrder(obj uintptr) TDateDisplayOrder {
	ret, _, _ := getLazyProc("DateTimePicker_GetDateDisplayOrder").Call(obj)
	return TDateDisplayOrder(ret)
}

func DateTimePicker_SetDateDisplayOrder(obj uintptr, value TDateDisplayOrder) {
	_, _, _ = getLazyProc("DateTimePicker_SetDateDisplayOrder").Call(obj, uintptr(value))
}

func DateTimePicker_GetDateSeparator(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetDateSeparator").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetDateSeparator(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetDateSeparator").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetLeadingZeros(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetLeadingZeros").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetLeadingZeros(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetLeadingZeros").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetMonthNames(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetMonthNames").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetMonthNames(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetMonthNames").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetShowMonthNames(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetShowMonthNames").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetShowMonthNames(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetShowMonthNames").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetNullInputAllowed(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetNullInputAllowed").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetNullInputAllowed(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetNullInputAllowed").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetOptions(obj uintptr) TDateTimePickerOptions {
	ret, _, _ := getLazyProc("DateTimePicker_GetOptions").Call(obj)
	return TDateTimePickerOptions(ret)
}

func DateTimePicker_SetOptions(obj uintptr, value TDateTimePickerOptions) {
	_, _, _ = getLazyProc("DateTimePicker_SetOptions").Call(obj, uintptr(value))
}

func DateTimePicker_GetShowCheckBox(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetShowCheckBox").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetShowCheckBox(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetShowCheckBox").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetReadOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetReadOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetReadOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetReadOnly").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetTextForNullDate(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetTextForNullDate").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetTextForNullDate(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetTextForNullDate").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTimeDisplay(obj uintptr) TTimeDisplay {
	ret, _, _ := getLazyProc("DateTimePicker_GetTimeDisplay").Call(obj)
	return TTimeDisplay(ret)
}

func DateTimePicker_SetTimeDisplay(obj uintptr, value TTimeDisplay) {
	_, _, _ = getLazyProc("DateTimePicker_SetTimeDisplay").Call(obj, uintptr(value))
}

func DateTimePicker_GetTimeSeparator(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetTimeSeparator").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetTimeSeparator(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetTimeSeparator").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTrailingSeparator(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetTrailingSeparator").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetTrailingSeparator(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetTrailingSeparator").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetUseDefaultSeparators(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetUseDefaultSeparators").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetUseDefaultSeparators(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetUseDefaultSeparators").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetDroppedDown(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetDroppedDown").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_GetDateTime(obj uintptr) time.Time {
	var ret int64
	_, _, _ = getLazyProc("DateTimePicker_GetDateTime").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func DateTimePicker_SetDateTime(obj uintptr, value time.Time) {
	tVal := value.Unix()
	_, _, _ = getLazyProc("DateTimePicker_SetDateTime").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func DateTimePicker_GetAlign(obj uintptr) TAlign {
	ret, _, _ := getLazyProc("DateTimePicker_GetAlign").Call(obj)
	return TAlign(ret)
}

func DateTimePicker_SetAlign(obj uintptr, value TAlign) {
	_, _, _ = getLazyProc("DateTimePicker_SetAlign").Call(obj, uintptr(value))
}

func DateTimePicker_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchors").Call(obj)
	return TAnchors(ret)
}

func DateTimePicker_SetAnchors(obj uintptr, value TAnchors) {
	_, _, _ = getLazyProc("DateTimePicker_SetAnchors").Call(obj, uintptr(value))
}

func DateTimePicker_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := getLazyProc("DateTimePicker_GetBiDiMode").Call(obj)
	return TBiDiMode(ret)
}

func DateTimePicker_SetBiDiMode(obj uintptr, value TBiDiMode) {
	_, _, _ = getLazyProc("DateTimePicker_SetBiDiMode").Call(obj, uintptr(value))
}

func DateTimePicker_GetCalAlignment(obj uintptr) TDTCalAlignment {
	ret, _, _ := getLazyProc("DateTimePicker_GetCalAlignment").Call(obj)
	return TDTCalAlignment(ret)
}

func DateTimePicker_SetCalAlignment(obj uintptr, value TDTCalAlignment) {
	_, _, _ = getLazyProc("DateTimePicker_SetCalAlignment").Call(obj, uintptr(value))
}

func DateTimePicker_GetConstraints(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetConstraints").Call(obj)
	return ret
}

func DateTimePicker_SetConstraints(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetConstraints").Call(obj, value)
}

func DateTimePicker_GetDate(obj uintptr) time.Time {
	var ret int64
	_, _, _ = getLazyProc("DateTimePicker_GetDate").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func DateTimePicker_SetDate(obj uintptr, value time.Time) {
	tVal := value.Unix()
	_, _, _ = getLazyProc("DateTimePicker_SetDate").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func DateTimePicker_GetTime(obj uintptr) time.Time {
	var ret int64
	_, _, _ = getLazyProc("DateTimePicker_GetTime").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func DateTimePicker_SetTime(obj uintptr, value time.Time) {
	tVal := value.Unix()
	_, _, _ = getLazyProc("DateTimePicker_SetTime").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func DateTimePicker_GetChecked(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetChecked").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetChecked(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetChecked").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("DateTimePicker_GetColor").Call(obj)
	return TColor(ret)
}

func DateTimePicker_SetColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("DateTimePicker_SetColor").Call(obj, uintptr(value))
}

func DateTimePicker_GetDateMode(obj uintptr) TDTDateMode {
	ret, _, _ := getLazyProc("DateTimePicker_GetDateMode").Call(obj)
	return TDTDateMode(ret)
}

func DateTimePicker_SetDateMode(obj uintptr, value TDTDateMode) {
	_, _, _ = getLazyProc("DateTimePicker_SetDateMode").Call(obj, uintptr(value))
}

func DateTimePicker_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetFont").Call(obj)
	return ret
}

func DateTimePicker_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetFont").Call(obj, value)
}

func DateTimePicker_GetKind(obj uintptr) TDateTimeKind {
	ret, _, _ := getLazyProc("DateTimePicker_GetKind").Call(obj)
	return TDateTimeKind(ret)
}

func DateTimePicker_SetKind(obj uintptr, value TDateTimeKind) {
	_, _, _ = getLazyProc("DateTimePicker_SetKind").Call(obj, uintptr(value))
}

func DateTimePicker_GetMaxDate(obj uintptr) time.Time {
	var ret int64
	_, _, _ = getLazyProc("DateTimePicker_GetMaxDate").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func DateTimePicker_SetMaxDate(obj uintptr, value time.Time) {
	tVal := value.Unix()
	_, _, _ = getLazyProc("DateTimePicker_SetMaxDate").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func DateTimePicker_GetMinDate(obj uintptr) time.Time {
	var ret int64
	_, _, _ = getLazyProc("DateTimePicker_GetMinDate").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func DateTimePicker_SetMinDate(obj uintptr, value time.Time) {
	tVal := value.Unix()
	_, _, _ = getLazyProc("DateTimePicker_SetMinDate").Call(obj, uintptr(unsafe.Pointer(&tVal)))
}

func DateTimePicker_GetParentColor(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetParentColor").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentColor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetParentColor").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetParentDoubleBuffered").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentDoubleBuffered(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetParentDoubleBuffered").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentFont(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetParentFont").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentFont(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetParentFont").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetParentShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetParentShowHint").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetPopupMenu").Call(obj)
	return ret
}

func DateTimePicker_SetPopupMenu(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetPopupMenu").Call(obj, value)
}

func DateTimePicker_GetShowHint(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetShowHint").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetShowHint(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetShowHint").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := getLazyProc("DateTimePicker_GetTabOrder").Call(obj)
	return TTabOrder(ret)
}

func DateTimePicker_SetTabOrder(obj uintptr, value TTabOrder) {
	_, _, _ = getLazyProc("DateTimePicker_SetTabOrder").Call(obj, uintptr(value))
}

func DateTimePicker_GetTabStop(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetTabStop").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetTabStop(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetTabStop").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetVisible").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_SetOnClick(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnClick").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnCloseUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnCloseUp").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnContextPopup(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnContextPopup").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnDropDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnDropDown").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnEnter").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnExit(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnExit").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnKeyDown(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnKeyDown").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnKeyPress(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnKeyPress").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnKeyUp(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnKeyUp").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnMouseEnter(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnMouseEnter").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_SetOnMouseLeave(obj uintptr, fn any) {
	_, _, _ = getLazyProc("DateTimePicker_SetOnMouseLeave").Call(obj, addEventToMap(obj, fn))
}

func DateTimePicker_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetDockClientCount").Call(obj)
	return int32(ret)
}

func DateTimePicker_GetDockSite(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetDockSite").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetDockSite(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetDockSite").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetMouseInClient").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetVisibleDockClientCount").Call(obj)
	return int32(ret)
}

func DateTimePicker_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetBrush").Call(obj)
	return ret
}

func DateTimePicker_GetControlCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetControlCount").Call(obj)
	return int32(ret)
}

func DateTimePicker_GetHandle(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DateTimePicker_GetHandle").Call(obj)
	return ret
}

func DateTimePicker_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := getLazyProc("DateTimePicker_GetParentWindow").Call(obj)
	return ret
}

func DateTimePicker_SetParentWindow(obj uintptr, value HWND) {
	_, _, _ = getLazyProc("DateTimePicker_SetParentWindow").Call(obj, value)
}

func DateTimePicker_GetShowing(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetShowing").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetUseDockManager").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_SetUseDockManager(obj uintptr, value bool) {
	_, _, _ = getLazyProc("DateTimePicker_SetUseDockManager").Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetAction(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAction").Call(obj)
	return ret
}

func DateTimePicker_SetAction(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetAction").Call(obj, value)
}

func DateTimePicker_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DateTimePicker_GetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_SetBoundsRect(obj uintptr, value TRect) {
	_, _, _ = getLazyProc("DateTimePicker_SetBoundsRect").Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DateTimePicker_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetClientHeight").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetClientHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetClientHeight").Call(obj, uintptr(value))
}

func DateTimePicker_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("DateTimePicker_GetClientOrigin").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_GetClientRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("DateTimePicker_GetClientRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func DateTimePicker_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetClientWidth").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetClientWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetClientWidth").Call(obj, uintptr(value))
}

func DateTimePicker_GetControlState(obj uintptr) TControlState {
	ret, _, _ := getLazyProc("DateTimePicker_GetControlState").Call(obj)
	return TControlState(ret)
}

func DateTimePicker_SetControlState(obj uintptr, value TControlState) {
	_, _, _ = getLazyProc("DateTimePicker_SetControlState").Call(obj, uintptr(value))
}

func DateTimePicker_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := getLazyProc("DateTimePicker_GetControlStyle").Call(obj)
	return TControlStyle(ret)
}

func DateTimePicker_SetControlStyle(obj uintptr, value TControlStyle) {
	_, _, _ = getLazyProc("DateTimePicker_SetControlStyle").Call(obj, uintptr(value))
}

func DateTimePicker_GetFloating(obj uintptr) bool {
	ret, _, _ := getLazyProc("DateTimePicker_GetFloating").Call(obj)
	return DBoolToGoBool(ret)
}

func DateTimePicker_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetParent").Call(obj)
	return ret
}

func DateTimePicker_SetParent(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetParent").Call(obj, value)
}

func DateTimePicker_GetLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetLeft").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetLeft").Call(obj, uintptr(value))
}

func DateTimePicker_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetTop").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetTop").Call(obj, uintptr(value))
}

func DateTimePicker_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetWidth").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetWidth").Call(obj, uintptr(value))
}

func DateTimePicker_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetHeight").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetHeight").Call(obj, uintptr(value))
}

func DateTimePicker_GetCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("DateTimePicker_GetCursor").Call(obj)
	return TCursor(ret)
}

func DateTimePicker_SetCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("DateTimePicker_SetCursor").Call(obj, uintptr(value))
}

func DateTimePicker_GetHint(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetHint").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetHint(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetHint").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetComponentCount").Call(obj)
	return int32(ret)
}

func DateTimePicker_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("DateTimePicker_GetComponentIndex").Call(obj)
	return int32(ret)
}

func DateTimePicker_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("DateTimePicker_SetComponentIndex").Call(obj, uintptr(value))
}

func DateTimePicker_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetOwner").Call(obj)
	return ret
}

func DateTimePicker_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("DateTimePicker_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func DateTimePicker_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("DateTimePicker_SetName").Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("DateTimePicker_GetTag").Call(obj)
	return int(ret)
}

func DateTimePicker_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("DateTimePicker_SetTag").Call(obj, uintptr(value))
}

func DateTimePicker_GetAnchorSideLeft(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchorSideLeft").Call(obj)
	return ret
}

func DateTimePicker_SetAnchorSideLeft(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetAnchorSideLeft").Call(obj, value)
}

func DateTimePicker_GetAnchorSideTop(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchorSideTop").Call(obj)
	return ret
}

func DateTimePicker_SetAnchorSideTop(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetAnchorSideTop").Call(obj, value)
}

func DateTimePicker_GetAnchorSideRight(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchorSideRight").Call(obj)
	return ret
}

func DateTimePicker_SetAnchorSideRight(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetAnchorSideRight").Call(obj, value)
}

func DateTimePicker_GetAnchorSideBottom(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchorSideBottom").Call(obj)
	return ret
}

func DateTimePicker_SetAnchorSideBottom(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetAnchorSideBottom").Call(obj, value)
}

func DateTimePicker_GetChildSizing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetChildSizing").Call(obj)
	return ret
}

func DateTimePicker_SetChildSizing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetChildSizing").Call(obj, value)
}

func DateTimePicker_GetBorderSpacing(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetBorderSpacing").Call(obj)
	return ret
}

func DateTimePicker_SetBorderSpacing(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("DateTimePicker_SetBorderSpacing").Call(obj, value)
}

func DateTimePicker_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetDockClients").Call(obj, uintptr(Index))
	return ret
}

func DateTimePicker_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetControls").Call(obj, uintptr(Index))
	return ret
}

func DateTimePicker_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func DateTimePicker_GetAnchorSide(obj uintptr, AKind TAnchorKind) uintptr {
	ret, _, _ := getLazyProc("DateTimePicker_GetAnchorSide").Call(obj, uintptr(AKind))
	return ret
}

func DateTimePicker_StaticClassType() TClass {
	r, _, _ := getLazyProc("DateTimePicker_StaticClassType").Call()
	return TClass(r)
}
