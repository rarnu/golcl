package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynHTMLSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_Create").Call(owner)
	return r
}

func SynHTMLSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynHTMLSyn_Free").Call(obj)
}

func SynHTMLSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynHTMLSyn_StaticClassType").Call()
	return TClass(r)
}

func SynHTMLSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynHTMLSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynHTMLSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynHTMLSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynHTMLSyn_GetAndAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetAndAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetASPAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetASPAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetCDATAAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetCDATAAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetDOCTYPEAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetDOCTYPEAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetCommentAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetKeyAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetTextAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetTextAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetUndefKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetUndefKeyAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetValueAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynHTMLSyn_GetValueAttri").Call(obj)
	return r
}

func SynHTMLSyn_GetMode(obj uintptr) TSynHTMLSynMode {
	r, _, _ := getLazyProc("SynHTMLSyn_GetMode").Call(obj)
	return TSynHTMLSynMode(r)
}

func SynHTMLSyn_SetMode(obj uintptr, value TSynHTMLSynMode) {
	_, _, _ = getLazyProc("SynHTMLSyn_SetMode").Call(obj, uintptr(value))
}
