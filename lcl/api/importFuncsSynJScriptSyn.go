package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynJScriptSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_Create").Call(owner)
	return r
}

func SynJScriptSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynJScriptSyn_Free").Call(obj)
}

func SynJScriptSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynJScriptSyn_StaticClassType").Call()
	return TClass(r)
}

func SynJScriptSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynJScriptSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynJScriptSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynJScriptSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynJScriptSyn_GetBracketAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetBracketAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetCommentAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetKeyAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetNonReservedKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetNonReservedKeyAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetEventAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetEventAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetNumberAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetStringAttri").Call(obj)
	return r
}

func SynJScriptSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJScriptSyn_GetSymbolAttri").Call(obj)
	return r
}
