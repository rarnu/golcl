package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynPythonSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_Create").Call(owner)
	return r
}

func SynPythonSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynPythonSyn_Free").Call(obj)
}

func SynPythonSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynPythonSyn_StaticClassType").Call()
	return TClass(r)
}

func SynPythonSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynPythonSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynPythonSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynPythonSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynPythonSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetCommentAttri").Call(obj)
	return r
}

func SynPythonSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynPythonSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetKeyAttri").Call(obj)
	return r
}

func SynPythonSyn_GetNonKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetNonKeyAttri").Call(obj)
	return r
}

func SynPythonSyn_GetSystemAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetSystemAttri").Call(obj)
	return r
}

func SynPythonSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetNumberAttri").Call(obj)
	return r
}

func SynPythonSyn_GetHexAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetHexAttri").Call(obj)
	return r
}

func SynPythonSyn_GetOctalAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetOctalAttri").Call(obj)
	return r
}

func SynPythonSyn_GetFloatAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetFloatAttri").Call(obj)
	return r
}

func SynPythonSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynPythonSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetStringAttri").Call(obj)
	return r
}

func SynPythonSyn_GetDocStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetDocStringAttri").Call(obj)
	return r
}

func SynPythonSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynPythonSyn_GetErrorAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynPythonSyn_GetErrorAttri").Call(obj)
	return r
}
