package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynUNIXShellScriptSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_Create").Call(owner)
	return r
}

func SynUNIXShellScriptSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynUNIXShellScriptSyn_Free").Call(obj)
}

func SynUNIXShellScriptSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_StaticClassType").Call()
	return TClass(r)
}

func SynUNIXShellScriptSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynUNIXShellScriptSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynUNIXShellScriptSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynUNIXShellScriptSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetCommentAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetKeyAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetSecondKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetSecondKeyAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetNumberAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetStringAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetVarAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetVarAttri").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_GetSecondKeyWords(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynUNIXShellScriptSyn_GetSecondKeyWords").Call(obj)
	return r
}

func SynUNIXShellScriptSyn_SetSecondKeyWords(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynUNIXShellScriptSyn_SetSecondKeyWords").Call(obj, value)
}
