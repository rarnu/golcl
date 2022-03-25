package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTaskDialog ---------------------------

func TaskDialog_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_Create").Call(obj)
	return ret
}

func TaskDialog_Free(obj uintptr) {
	_, _, _ = getLazyProc("TaskDialog_Free").Call(obj)
}

func TaskDialog_Execute(obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialog_Execute").Call(obj)
	return DBoolToGoBool(ret)
}

func TaskDialog_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func TaskDialog_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialog_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func TaskDialog_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TaskDialog_Assign").Call(obj, Source)
}

func TaskDialog_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TaskDialog_ClassType").Call(obj)
	return TClass(ret)
}

func TaskDialog_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialog_InstanceSize").Call(obj)
	return int32(ret)
}

func TaskDialog_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TaskDialog_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TaskDialog_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TaskDialog_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TaskDialog_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialog_GetHashCode").Call(obj)
	return int32(ret)
}

func TaskDialog_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_GetButtons(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetButtons").Call(obj)
	return ret
}

func TaskDialog_SetButtons(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TaskDialog_SetButtons").Call(obj, value)
}

func TaskDialog_GetCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetCaption").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetCommonButtons(obj uintptr) TTaskDialogCommonButtons {
	ret, _, _ := getLazyProc("TaskDialog_GetCommonButtons").Call(obj)
	return TTaskDialogCommonButtons(ret)
}

func TaskDialog_SetCommonButtons(obj uintptr, value TTaskDialogCommonButtons) {
	_, _, _ = getLazyProc("TaskDialog_SetCommonButtons").Call(obj, uintptr(value))
}

func TaskDialog_GetDefaultButton(obj uintptr) TTaskDialogCommonButton {
	ret, _, _ := getLazyProc("TaskDialog_GetDefaultButton").Call(obj)
	return TTaskDialogCommonButton(ret)
}

func TaskDialog_SetDefaultButton(obj uintptr, value TTaskDialogCommonButton) {
	_, _, _ = getLazyProc("TaskDialog_SetDefaultButton").Call(obj, uintptr(value))
}

func TaskDialog_GetExpandButtonCaption(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetExpandButtonCaption").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetExpandButtonCaption(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetExpandButtonCaption").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetExpandedText(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetExpandedText").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetExpandedText(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetExpandedText").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetFlags(obj uintptr) TTaskDialogFlags {
	ret, _, _ := getLazyProc("TaskDialog_GetFlags").Call(obj)
	return TTaskDialogFlags(ret)
}

func TaskDialog_SetFlags(obj uintptr, value TTaskDialogFlags) {
	_, _, _ = getLazyProc("TaskDialog_SetFlags").Call(obj, uintptr(value))
}

func TaskDialog_GetFooterIcon(obj uintptr) TTaskDialogIcon {
	ret, _, _ := getLazyProc("TaskDialog_GetFooterIcon").Call(obj)
	return TTaskDialogIcon(ret)
}

func TaskDialog_SetFooterIcon(obj uintptr, value TTaskDialogIcon) {
	_, _, _ = getLazyProc("TaskDialog_SetFooterIcon").Call(obj, uintptr(value))
}

func TaskDialog_GetFooterText(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetFooterText").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetFooterText(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetFooterText").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetMainIcon(obj uintptr) TTaskDialogIcon {
	ret, _, _ := getLazyProc("TaskDialog_GetMainIcon").Call(obj)
	return TTaskDialogIcon(ret)
}

func TaskDialog_SetMainIcon(obj uintptr, value TTaskDialogIcon) {
	_, _, _ = getLazyProc("TaskDialog_SetMainIcon").Call(obj, uintptr(value))
}

func TaskDialog_GetRadioButtons(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetRadioButtons").Call(obj)
	return ret
}

func TaskDialog_SetRadioButtons(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TaskDialog_SetRadioButtons").Call(obj, value)
}

func TaskDialog_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetText").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetTitle").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetVerificationText(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetVerificationText").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetVerificationText(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetVerificationText").Call(obj, GoStrToDStr(value))
}

func TaskDialog_SetOnButtonClicked(obj uintptr, fn any) {
	_, _, _ = getLazyProc("TaskDialog_SetOnButtonClicked").Call(obj, addEventToMap(obj, fn))
}

func TaskDialog_GetButton(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetButton").Call(obj)
	return ret
}

func TaskDialog_SetButton(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("TaskDialog_SetButton").Call(obj, value)
}

func TaskDialog_GetModalResult(obj uintptr) TModalResult {
	ret, _, _ := getLazyProc("TaskDialog_GetModalResult").Call(obj)
	return TModalResult(ret)
}

func TaskDialog_SetModalResult(obj uintptr, value TModalResult) {
	_, _, _ = getLazyProc("TaskDialog_SetModalResult").Call(obj, uintptr(value))
}

func TaskDialog_GetRadioButton(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetRadioButton").Call(obj)
	return ret
}

func TaskDialog_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialog_GetComponentCount").Call(obj)
	return int32(ret)
}

func TaskDialog_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TaskDialog_GetComponentIndex").Call(obj)
	return int32(ret)
}

func TaskDialog_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TaskDialog_SetComponentIndex").Call(obj, uintptr(value))
}

func TaskDialog_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetOwner").Call(obj)
	return ret
}

func TaskDialog_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("TaskDialog_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func TaskDialog_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("TaskDialog_SetName").Call(obj, GoStrToDStr(value))
}

func TaskDialog_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("TaskDialog_GetTag").Call(obj)
	return int(ret)
}

func TaskDialog_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("TaskDialog_SetTag").Call(obj, uintptr(value))
}

func TaskDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("TaskDialog_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func TaskDialog_StaticClassType() TClass {
	r, _, _ := getLazyProc("TaskDialog_StaticClassType").Call()
	return TClass(r)
}
