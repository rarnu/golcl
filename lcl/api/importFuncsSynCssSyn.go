package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynCssSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_Create").Call(owner)
	return r
}

func SynCssSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynCssSyn_Free").Call(obj)
}

func SynCssSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynCssSyn_StaticClassType").Call()
	return TClass(r)
}

func SynCssSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCssSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynCssSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCssSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynCssSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetCommentAttri").Call(obj)
	return r
}

func SynCssSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynCssSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetKeyAttri").Call(obj)
	return r
}

func SynCssSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetNumberAttri").Call(obj)
	return r
}

func SynCssSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynCssSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetStringAttri").Call(obj)
	return r
}

func SynCssSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynCssSyn_GetMeasurementUnitAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetMeasurementUnitAttri").Call(obj)
	return r
}

func SynCssSyn_GetSelectorAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynCssSyn_GetSelectorAttri").Call(obj)
	return r
}
