package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynBatSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_Create").Call(owner)
	return r
}

func SynBatSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynBatSyn_Free").Call(obj)
}

func SynBatSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynBatSyn_StaticClassType").Call()
	return TClass(r)
}

func SynBatSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynBatSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynBatSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynBatSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynBatSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetCommentAttri").Call(obj)
	return r
}

func SynBatSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynBatSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetKeyAttri").Call(obj)
	return r
}

func SynBatSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetNumberAttri").Call(obj)
	return r
}

func SynBatSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynBatSyn_GetVariableAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynBatSyn_GetVariableAttri").Call(obj)
	return r
}
