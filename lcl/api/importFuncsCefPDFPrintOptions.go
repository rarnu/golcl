package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func PDFPrintOptions_Create() uintptr {
	ret, _, _ := getLazyProc("PDFPrintOptions_Create").Call()
	return ret
}

func PDFPrintOptions_Free(obj uintptr) {
	_, _, _ = getLazyProc("PDFPrintOptions_Free").Call(obj)
}

func PDFPrintOptions_StaticClassType() TClass {
	ret, _, _ := getLazyProc("PDFPrintOptions_StaticClassType").Call()
	return TClass(ret)
}

func PDFPrintOptions_GetPageWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetPageWidth").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetPageWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetPageWidth").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetPageHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetPageHeight").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetPageHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetPageHeight").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetScaleFactor(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetScaleFactor").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetScaleFactor(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetScaleFactor").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetMarginTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetMarginTop").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetMarginTop(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetMarginTop").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetMarginLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetMarginLeft").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetMarginLeft(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetMarginLeft").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetMarginRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetMarginRight").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetMarginRight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetMarginRight").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetMarginBottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetMarginBottom").Call(obj)
	return int32(ret)
}

func PDFPrintOptions_SetMarginBottom(obj uintptr, value int32) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetMarginBottom").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetHeaderFooterEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetHeaderFooterEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func PDFPrintOptions_SetHeaderFooterEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetHeaderFooterEnabled").Call(obj, GoBoolToDBool(value))
}

func PDFPrintOptions_GetMarginType(obj uintptr) TCefPdfPrintMarginType {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetMarginType").Call(obj)
	return TCefPdfPrintMarginType(ret)
}

func PDFPrintOptions_SetMarginType(obj uintptr, value TCefPdfPrintMarginType) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetMarginType").Call(obj, uintptr(value))
}

func PDFPrintOptions_GetSelectionOnly(obj uintptr) bool {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetSelectionOnly").Call(obj)
	return DBoolToGoBool(ret)
}

func PDFPrintOptions_SetSelectionOnly(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetSelectionOnly").Call(obj, GoBoolToDBool(value))
}

func PDFPrintOptions_GetLandscape(obj uintptr) bool {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetLandscape").Call(obj)
	return DBoolToGoBool(ret)
}

func PDFPrintOptions_SetLandscape(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetLandscape").Call(obj, GoBoolToDBool(value))
}

func PDFPrintOptions_GetBackgroundsEnabled(obj uintptr) bool {
	ret, _, _ := getLazyProc("PDFPrintOptions_GetBackgroundsEnabled").Call(obj)
	return DBoolToGoBool(ret)
}

func PDFPrintOptions_SetBackgroundsEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("PDFPrintOptions_SetBackgroundsEnabled").Call(obj, GoBoolToDBool(value))
}
