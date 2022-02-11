package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynPasSyn_Create(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_Create").Call(obj)
	return r
}

func SynPasSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynPasSyn_Free").Call(obj)
}

func SynPasSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynPasSyn_StaticClassType").Call()
	return TClass(r)
}

func SynPasSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPasSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynPasSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPasSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynPasSyn_GetAsmAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetAsmAttri").Call(obj)
	return r
}

func SynPasSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetCommentAttri").Call(obj)
	return r
}

func SynPasSyn_GetIDEDirectiveAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetIDEDirectiveAttri").Call(obj)
	return r
}

func SynPasSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynPasSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetKeyAttri").Call(obj)
	return r
}

func SynPasSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetNumberAttri").Call(obj)
	return r
}

func SynPasSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynPasSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetStringAttri").Call(obj)
	return r
}

func SynPasSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynPasSyn_GetProcedureHeaderName(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetProcedureHeaderName").Call(obj)
	return r
}

func SynPasSyn_GetCaseLabelAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetCaseLabelAttri").Call(obj)
	return r
}

func SynPasSyn_GetDirectiveAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPasSyn_GetDirectiveAttri").Call(obj)
	return r
}

func SynPasSyn_GetCompilerMode(obj uintptr) TPascalCompilerMode {
	r, _, _ := getLazyProc("SynPasSyn_GetCompilerMode").Call(obj)
	return TPascalCompilerMode(r)
}

func SynPasSyn_SetCompilerMode(obj uintptr, value TPascalCompilerMode) {
	_, _, _ = getLazyProc("SynPasSyn_SetCompilerMode").Call(obj, uintptr(value))
}

func SynPasSyn_GetNestedComments(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPasSyn_GetNestedComments").Call(obj)
	return DBoolToGoBool(r)
}

func SynPasSyn_SetNestedComments(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPasSyn_SetNestedComments").Call(obj, GoBoolToDBool(value))
}

func SynPasSyn_GetTypeHelpers(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPasSyn_GetTypeHelpers").Call(obj)
	return DBoolToGoBool(r)
}

func SynPasSyn_SetTypeHelpers(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPasSyn_SetTypeHelpers").Call(obj, GoBoolToDBool(value))
}

func SynPasSyn_GetD4syntax(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPasSyn_GetD4syntax").Call(obj)
	return DBoolToGoBool(r)
}

func SynPasSyn_SetD4syntax(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPasSyn_SetD4syntax").Call(obj, GoBoolToDBool(value))
}

func SynPasSyn_GetExtendedKeywordsMode(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPasSyn_GetExtendedKeywordsMode").Call(obj)
	return DBoolToGoBool(r)
}

func SynPasSyn_SetExtendedKeywordsMode(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPasSyn_SetExtendedKeywordsMode").Call(obj, GoBoolToDBool(value))
}

func SynPasSyn_GetStringKeywordMode(obj uintptr) TSynPasStringMode {
	r, _, _ := getLazyProc("SynPasSyn_GetStringKeywordMode").Call(obj)
	return TSynPasStringMode(r)
}

func SynPasSyn_SetStringKeywordMode(obj uintptr, value TSynPasStringMode) {
	_, _, _ = getLazyProc("SynPasSyn_SetStringKeywordMode").Call(obj, uintptr(value))
}
