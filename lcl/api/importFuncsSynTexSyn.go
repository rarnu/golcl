package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynTeXSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_Create").Call(owner)
	return r
}

func SynTeXSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynTeXSyn_Free").Call(obj)
}

func SynTeXSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynTeXSyn_StaticClassType").Call()
	return TClass(r)
}

func SynTeXSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynTeXSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynTeXSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynTeXSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynTeXSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetCommentAttri").Call(obj)
	return r
}

func SynTeXSyn_GetTextAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetTextAttri").Call(obj)
	return r
}

func SynTeXSyn_GetControlSequenceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetControlSequenceAttri").Call(obj)
	return r
}

func SynTeXSyn_GetMathmodeAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetMathmodeAttri").Call(obj)
	return r
}

func SynTeXSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynTeXSyn_GetBraceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetBraceAttri").Call(obj)
	return r
}

func SynTeXSyn_GetBracketAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynTeXSyn_GetBracketAttri").Call(obj)
	return r
}
