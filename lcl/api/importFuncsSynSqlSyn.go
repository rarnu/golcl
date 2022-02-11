package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynSQLSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_Create").Call(owner)
	return r
}

func SynSQLSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynSQLSyn_Free").Call(obj)
}

func SynSQLSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynSQLSyn_StaticClassType").Call()
	return TClass(r)
}

func SynSQLSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynSQLSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynSQLSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynSQLSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynSQLSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetCommentAttri").Call(obj)
	return r
}

func SynSQLSyn_GetDataTypeAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetDataTypeAttri").Call(obj)
	return r
}

func SynSQLSyn_GetDefaultPackageAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetDefaultPackageAttri").Call(obj)
	return r
}

func SynSQLSyn_GetExceptionAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetExceptionAttri").Call(obj)
	return r
}

func SynSQLSyn_GetFunctionAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetFunctionAttri").Call(obj)
	return r
}

func SynSQLSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynSQLSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetKeyAttri").Call(obj)
	return r
}

func SynSQLSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetNumberAttri").Call(obj)
	return r
}

func SynSQLSyn_GetPLSQLAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetPLSQLAttri").Call(obj)
	return r
}

func SynSQLSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynSQLSyn_GetSQLPlusAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetSQLPlusAttri").Call(obj)
	return r
}

func SynSQLSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetStringAttri").Call(obj)
	return r
}

func SynSQLSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynSQLSyn_GetTableNameAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetTableNameAttri").Call(obj)
	return r
}

func SynSQLSyn_GetVariableAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetVariableAttri").Call(obj)
	return r
}

func SynSQLSyn_GetTableNames(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynSQLSyn_GetTableNames").Call(obj)
	return r
}

func SynSQLSyn_SetTableNames(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynSQLSyn_SetTableNames").Call(obj, value)
}

func SynSQLSyn_GetSQLDialect(obj uintptr) TSQLDialect {
	r, _, _ := getLazyProc("SynSQLSyn_GetSQLDialect").Call(obj)
	return TSQLDialect(r)
}

func SynSQLSyn_SetSQLDialect(obj uintptr, value TSQLDialect) {
	_, _, _ = getLazyProc("SynSQLSyn_SetSQLDialect").Call(obj, uintptr(value))
}
