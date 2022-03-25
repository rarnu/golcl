package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TClipboard ---------------------------

func Clipboard_Create() uintptr {
	ret, _, _ := getLazyProc("Clipboard_Create").Call()
	return ret
}

func Clipboard_Free(obj uintptr) {
	_, _, _ = getLazyProc("Clipboard_Free").Call(obj)
}

func Clipboard_FindPictureFormatID(obj uintptr) TClipboardFormat {
	ret, _, _ := getLazyProc("Clipboard_FindPictureFormatID").Call(obj)
	return ret
}

func Clipboard_FindFormatID(obj uintptr, FormatName string) TClipboardFormat {
	ret, _, _ := getLazyProc("Clipboard_FindFormatID").Call(obj, GoStrToDStr(FormatName))
	return ret
}

func Clipboard_SupportedFormats(obj uintptr, List uintptr) {
	_, _, _ = getLazyProc("Clipboard_SupportedFormats").Call(obj, List)
}

func Clipboard_HasFormatName(obj uintptr, FormatName string) bool {
	ret, _, _ := getLazyProc("Clipboard_HasFormatName").Call(obj, GoStrToDStr(FormatName))
	return DBoolToGoBool(ret)
}

func Clipboard_HasPictureFormat(obj uintptr) bool {
	ret, _, _ := getLazyProc("Clipboard_HasPictureFormat").Call(obj)
	return DBoolToGoBool(ret)
}

func Clipboard_SetAsHtml(obj uintptr, Html string, PlainText string) {
	_, _, _ = getLazyProc("Clipboard_SetAsHtml").Call(obj, GoStrToDStr(Html), GoStrToDStr(PlainText))
}

func Clipboard_GetFormat(obj uintptr, FormatID TClipboardFormat, Stream uintptr) bool {
	ret, _, _ := getLazyProc("Clipboard_GetFormat").Call(obj, FormatID, Stream)
	return DBoolToGoBool(ret)
}

func Clipboard_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Clipboard_Assign").Call(obj, Source)
}

func Clipboard_Clear(obj uintptr) {
	_, _, _ = getLazyProc("Clipboard_Clear").Call(obj)
}

func Clipboard_Close(obj uintptr) {
	_, _, _ = getLazyProc("Clipboard_Close").Call(obj)
}

func Clipboard_Open(obj uintptr) {
	_, _, _ = getLazyProc("Clipboard_Open").Call(obj)
}

func Clipboard_SetTextBuf(obj uintptr, Buffer string) {
	_, _, _ = getLazyProc("Clipboard_SetTextBuf").Call(obj, GoStrToDStr(Buffer))
}

func Clipboard_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Clipboard_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Clipboard_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Clipboard_ClassType").Call(obj)
	return TClass(ret)
}

func Clipboard_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Clipboard_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Clipboard_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Clipboard_InstanceSize").Call(obj)
	return int32(ret)
}

func Clipboard_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Clipboard_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Clipboard_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Clipboard_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Clipboard_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Clipboard_GetHashCode").Call(obj)
	return int32(ret)
}

func Clipboard_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Clipboard_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Clipboard_GetFormatCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Clipboard_GetFormatCount").Call(obj)
	return int32(ret)
}

func Clipboard_GetFormats(obj uintptr, Index int32) TClipboardFormat {
	ret, _, _ := getLazyProc("Clipboard_GetFormats").Call(obj, uintptr(Index))
	return ret
}

func Clipboard_StaticClassType() TClass {
	r, _, _ := getLazyProc("Clipboard_StaticClassType").Call()
	return TClass(r)
}
