package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynAnySyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_Create").Call(owner)
	return r
}

func SynAnySyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynAnySyn_Free").Call(obj)
}

func SynAnySyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynAnySyn_StaticClassType").Call()
	return TClass(r)
}

func SynAnySyn_GetActiveDot(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetActiveDot").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetActiveDot(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetActiveDot").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetCommentAttri").Call(obj)
	return r
}

func SynAnySyn_GetComments(obj uintptr) CommentStyles {
	r, _, _ := getLazyProc("SynAnySyn_GetComments").Call(obj)
	return CommentStyles(r)
}

func SynAnySyn_SetComments(obj uintptr, value CommentStyles) {
	_, _, _ = getLazyProc("SynAnySyn_SetComments").Call(obj, uintptr(value))
}

func SynAnySyn_GetConstantAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetConstantAttri").Call(obj)
	return r
}

func SynAnySyn_GetConstants(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetConstants").Call(obj)
	return r
}

func SynAnySyn_SetConstants(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynAnySyn_SetConstants").Call(obj, value)
}

func SynAnySyn_GetDetectPreprocessor(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetDetectPreprocessor").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetDetectPreprocessor(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetDetectPreprocessor").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetDollarVariables(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetDollarVariables").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetDollarVariables(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetDollarVariables").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetDotAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetDotAttri").Call(obj)
	return r
}

func SynAnySyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetEntity(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetEntity").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetEntity(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetEntity").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetEntityAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetEntityAttri").Call(obj)
	return r
}

func SynAnySyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynAnySyn_GetIdentifierChars(obj uintptr) string {
	r, _, _ := getLazyProc("SynAnySyn_GetIdentifierChars").Call(obj)
	return DStrToGoStr(r)
}

func SynAnySyn_SetIdentifierChars(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynAnySyn_SetIdentifierChars").Call(obj, GoStrToDStr(value))
}

func SynAnySyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetKeyAttri").Call(obj)
	return r
}

func SynAnySyn_GetKeyWords(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetKeyWords").Call(obj)
	return r
}

func SynAnySyn_SetKeyWords(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynAnySyn_SetKeyWords").Call(obj, value)
}

func SynAnySyn_GetMarkup(obj uintptr) bool {
	r, _, _ := getLazyProc("SynAnySyn_GetMarkup").Call(obj)
	return DBoolToGoBool(r)
}

func SynAnySyn_SetMarkup(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynAnySyn_SetMarkup").Call(obj, GoBoolToDBool(value))
}

func SynAnySyn_GetName(obj uintptr) string {
	r, _, _ := getLazyProc("SynAnySyn_GetName").Call(obj)
	return DStrToGoStr(r)
}

func SynAnySyn_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynAnySyn_SetName").Call(obj, GoStrToDStr(value))
}

func SynAnySyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetNumberAttri").Call(obj)
	return r
}

func SynAnySyn_GetObjectAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetObjectAttri").Call(obj)
	return r
}

func SynAnySyn_GetObjects(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetObjects").Call(obj)
	return r
}

func SynAnySyn_SetObjects(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("SynAnySyn_SetObjects").Call(obj, value)
}

func SynAnySyn_GetPreprocessorAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetPreprocessorAttri").Call(obj)
	return r
}

func SynAnySyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetSpaceAttri").Call(obj)
	return r
}

func SynAnySyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetStringAttri").Call(obj)
	return r
}

func SynAnySyn_GetStringDelim(obj uintptr) TStringDelim {
	r, _, _ := getLazyProc("SynAnySyn_GetStringDelim").Call(obj)
	return TStringDelim(r)
}

func SynAnySyn_SetStringDelim(obj uintptr, value TStringDelim) {
	_, _, _ = getLazyProc("SynAnySyn_SetStringDelim").Call(obj, uintptr(value))
}

func SynAnySyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetSymbolAttri").Call(obj)
	return r
}

func SynAnySyn_GetVariableAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynAnySyn_GetVariableAttri").Call(obj)
	return r
}
