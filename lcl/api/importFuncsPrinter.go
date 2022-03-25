package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TPrinter ---------------------------

func Printer_Create() uintptr {
	ret, _, _ := getLazyProc("Printer_Create").Call()
	return ret
}

func Printer_Free(obj uintptr) {
	_, _, _ = getLazyProc("Printer_Free").Call(obj)
}

func Printer_Abort(obj uintptr) {
	_, _, _ = getLazyProc("Printer_Abort").Call(obj)
}

func Printer_BeginDoc(obj uintptr) {
	_, _, _ = getLazyProc("Printer_BeginDoc").Call(obj)
}

func Printer_EndDoc(obj uintptr) {
	_, _, _ = getLazyProc("Printer_EndDoc").Call(obj)
}

func Printer_NewPage(obj uintptr) {
	_, _, _ = getLazyProc("Printer_NewPage").Call(obj)
}

func Printer_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Printer_Refresh").Call(obj)
}

func Printer_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Printer_ClassType").Call(obj)
	return TClass(ret)
}

func Printer_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Printer_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Printer_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_InstanceSize").Call(obj)
	return int32(ret)
}

func Printer_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Printer_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Printer_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Printer_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Printer_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetHashCode").Call(obj)
	return int32(ret)
}

func Printer_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Printer_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Printer_GetAborted(obj uintptr) bool {
	ret, _, _ := getLazyProc("Printer_GetAborted").Call(obj)
	return DBoolToGoBool(ret)
}

func Printer_GetCanvas(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Printer_GetCanvas").Call(obj)
	return ret
}

func Printer_GetCopies(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetCopies").Call(obj)
	return int32(ret)
}

func Printer_SetCopies(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Printer_SetCopies").Call(obj, uintptr(value))
}

func Printer_GetFonts(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Printer_GetFonts").Call(obj)
	return ret
}

func Printer_GetOrientation(obj uintptr) TPrinterOrientation {
	ret, _, _ := getLazyProc("Printer_GetOrientation").Call(obj)
	return TPrinterOrientation(ret)
}

func Printer_SetOrientation(obj uintptr, value TPrinterOrientation) {
	_, _, _ = getLazyProc("Printer_SetOrientation").Call(obj, uintptr(value))
}

func Printer_GetPageHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetPageHeight").Call(obj)
	return int32(ret)
}

func Printer_GetPageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetPageWidth").Call(obj)
	return int32(ret)
}

func Printer_GetPageNumber(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetPageNumber").Call(obj)
	return int32(ret)
}

func Printer_GetPrinterIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Printer_GetPrinterIndex").Call(obj)
	return int32(ret)
}

func Printer_SetPrinterIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Printer_SetPrinterIndex").Call(obj, uintptr(value))
}

func Printer_GetPrinting(obj uintptr) bool {
	ret, _, _ := getLazyProc("Printer_GetPrinting").Call(obj)
	return DBoolToGoBool(ret)
}

func Printer_GetPrinters(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Printer_GetPrinters").Call(obj)
	return ret
}

func Printer_GetTitle(obj uintptr) string {
	ret, _, _ := getLazyProc("Printer_GetTitle").Call(obj)
	return DStrToGoStr(ret)
}

func Printer_SetTitle(obj uintptr, value string) {
	_, _, _ = getLazyProc("Printer_SetTitle").Call(obj, GoStrToDStr(value))
}

func Printer_StaticClassType() TClass {
	r, _, _ := getLazyProc("Printer_StaticClassType").Call()
	return TClass(r)
}
