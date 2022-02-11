package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynPHPSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_Create").Call(owner)
	return r
}

func SynPHPSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynPHPSyn_Free").Call(obj)
}

func SynPHPSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynPHPSyn_StaticClassType").Call()
	return TClass(r)
}

func SynPHPSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPHPSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynPHPSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPHPSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynPHPSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetCommentAttri").Call(obj)
	return r
}

func SynPHPSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynPHPSyn_GetInvalidSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetInvalidSymbolAttri").Call(obj)
	return r
}

func SynPHPSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetKeyAttri").Call(obj)
	return r
}

func SynPHPSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetNumberAttri").Call(obj)
	return r
}

func SynPHPSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynPHPSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetStringAttri").Call(obj)
	return r
}

func SynPHPSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynPHPSyn_GetVariableAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPHPSyn_GetVariableAttri").Call(obj)
	return r
}
