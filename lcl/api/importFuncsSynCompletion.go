package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynCompletion_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynCompletion_Create").Call(owner)
	return r
}

func SynCompletion_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynCompletion_Free").Call(obj)
}

func SynCompletion_GetAutoUseSingleIdent(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_GetAutoUseSingleIdent").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_SetAutoUseSingleIdent(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCompletion_SetAutoUseSingleIdent").Call(obj, GoBoolToDBool(value))
}

func SynCompletion_GetCaseSensitive(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_GetCaseSensitive").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_SetCaseSensitive(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCompletion_SetCaseSensitive").Call(obj, GoBoolToDBool(value))
}

func SynCompletion_GetDoubleClickSelectsAll(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_GetDoubleClickSelectsAll").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_SetDoubleClickSelectsAll(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCompletion_SetDoubleClickSelectsAll").Call(obj, GoBoolToDBool(value))
}

func SynCompletion_GetEditor(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCompletion_GetEditor").Call(obj)
	return r
}

func SynCompletion_SetEditor(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynCompletion_SetEditor").Call(obj, value)
}

func SynCompletion_GetEndOfTokenChr(obj uintptr) string {
	r, _, _ := getLazyProc("SynCompletion_GetEndOfTokenChr").Call(obj)
	return DStrToGoStr(r)
}

func SynCompletion_SetEndOfTokenChr(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynCompletion_SetEndOfTokenChr").Call(obj, GoStrToDStr(value))
}

func SynCompletion_GetExecCommandID(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetExecCommandID").Call(obj)
	return int32(r)
}

func SynCompletion_SetExecCommandID(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetExecCommandID").Call(obj, uintptr(value))
}

func SynCompletion_GetItemList(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCompletion_GetItemList").Call(obj)
	return r
}

func SynCompletion_SetItemList(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynCompletion_SetItemList").Call(obj, value)
}

func SynCompletion_GetLinesInWindow(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetLinesInWindow").Call(obj)
	return int32(r)
}

func SynCompletion_SetLinesInWindow(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetLinesInWindow").Call(obj, uintptr(value))
}

func SynCompletion_GetLongLineHintTime(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetLongLineHintTime").Call(obj)
	return int32(r)
}

func SynCompletion_SetLongLineHintTime(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetLongLineHintTime").Call(obj, uintptr(value))
}

func SynCompletion_GetLongLineHintType(obj uintptr) TSynCompletionLongHintType {
	r, _, _ := getLazyProc("SynCompletion_GetLongLineHintType").Call(obj)
	return TSynCompletionLongHintType(int32(r))
}

func SynCompletion_SetLongLineHintType(obj uintptr, value TSynCompletionLongHintType) {
	_, _, _ = getLazyProc("SynCompletion_SetLongLineHintType").Call(obj, uintptr(value))
}

func SynCompletion_GetName(obj uintptr) string {
	r, _, _ := getLazyProc("SynCompletion_GetName").Call(obj)
	return DStrToGoStr(r)
}

func SynCompletion_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynCompletion_SetName").Call(obj, GoStrToDStr(value))
}

func SynCompletion_GetPosition(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetPosition").Call(obj)
	return int32(r)
}

func SynCompletion_SetPosition(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetPosition").Call(obj, uintptr(value))
}

func SynCompletion_GetSelectedColor(obj uintptr) TColor {
	r, _, _ := getLazyProc("SynCompletion_GetSelectedColor").Call(obj)
	return TColor(r)
}

func SynCompletion_SetSelectedColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("SynCompletion_SetSelectedColor").Call(obj, uintptr(value))
}

func SynCompletion_GetShortCut(obj uintptr) TShortCut {
	r, _, _ := getLazyProc("SynCompletion_GetShortCut").Call(obj)
	return TShortCut(r)
}

func SynCompletion_SetShortCut(obj uintptr, value TShortCut) {
	_, _, _ = getLazyProc("SynCompletion_SetShortCut").Call(obj, uintptr(value))
}

func SynCompletion_GetShowSizeDrag(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_GetShowSizeDrag").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_SetShowSizeDrag(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCompletion_SetShowSizeDrag").Call(obj, GoBoolToDBool(value))
}

func SynCompletion_GetTag(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetTag").Call(obj)
	return int32(r)
}

func SynCompletion_SetTag(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetTag").Call(obj, uintptr(value))
}

func SynCompletion_GetToggleReplaceWholeWord(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_GetToggleReplaceWholeWord").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_SetToggleReplaceWholeWord(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCompletion_SetToggleReplaceWholeWord").Call(obj, GoBoolToDBool(value))
}

func SynCompletion_GetWidth(obj uintptr) int32 {
	r, _, _ := getLazyProc("SynCompletion_GetWidth").Call(obj)
	return int32(r)
}

func SynCompletion_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("SynCompletion_SetWidth").Call(obj, uintptr(value))
}

func SynCompletion_SetOnCodeCompletion(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnCodeCompletion").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnExecute(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnExecute").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnKeyCompletePrefix(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnKeyCompletePrefix").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnKeyNextChar(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnKeyNextChar").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnKeyPrevChar(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnKeyPrevChar").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnPositionChanged(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnPositionChanged").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_SetOnSearchPosition(obj uintptr, fn interface{}) {
	_, _, _ = getLazyProc("SynCompletion_SetOnSearchPosition").Call(obj, addEventToMap(obj, fn))
}

func SynCompletion_Deactivate(obj uintptr) {
	_, _, _ = getLazyProc("SynCompletion_Deactivate").Call(obj)
}

func SynCompletion_Execute(obj uintptr, s string, x int32, y int32) {
	_, _, _ = getLazyProc("SynCompletion_Execute").Call(obj, GoStrToDStr(s), uintptr(x), uintptr(y))
}

func SynCompletion_IsActive(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCompletion_IsActive").Call(obj)
	return DBoolToGoBool(r)
}

func SynCompletion_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynCompletion_StaticClassType").Call()
	return TClass(r)
}
