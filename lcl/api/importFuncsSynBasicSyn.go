package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynVBSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_Create").Call(owner)
	return r
}

func SynVBSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynVBSyn_Free").Call(obj)
}

func SynVBSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynVBSyn_StaticClassType").Call()
	return TClass(r)
}

func SynVBSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynVBSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynVBSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynVBSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynVBSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetCommentAttri").Call(obj)
	return r
}

func SynVBSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynVBSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetKeyAttri").Call(obj)
	return r
}

func SynVBSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetNumberAttri").Call(obj)
	return r
}

func SynVBSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynVBSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetStringAttri").Call(obj)
	return r
}

func SynVBSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynVBSyn_GetSymbolAttri").Call(obj)
	return r
}
