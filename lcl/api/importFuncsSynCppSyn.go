package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynCppSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_Create").Call(owner)
	return r
}

func SynCppSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynCppSyn_Free").Call(obj)
}

func SynCppSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynCppSyn_StaticClassType").Call()
	return TClass(r)
}

func SynCppSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCppSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynCppSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCppSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynCppSyn_GetAsmAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetAsmAttri").Call(obj)
	return r
}

func SynCppSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetCommentAttri").Call(obj)
	return r
}

func SynCppSyn_GetDirecAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetDirecAttri").Call(obj)
	return r
}

func SynCppSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynCppSyn_GetInvalidAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetInvalidAttri").Call(obj)
	return r
}

func SynCppSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetKeyAttri").Call(obj)
	return r
}

func SynCppSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetNumberAttri").Call(obj)
	return r
}

func SynCppSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynCppSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetStringAttri").Call(obj)
	return r
}

func SynCppSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCppSyn_GetSymbolAttri").Call(obj)
	return r
}
