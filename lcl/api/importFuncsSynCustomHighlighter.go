package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynCustomHighlighter_GetDefaultFilter(obj uintptr) string {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetDefaultFilter").Call(obj)
	return DStrToGoStr(r)
}

func SynCustomHighlighter_SetDefaultFilter(obj uintptr, value string) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetDefaultFilter").Call(obj, GoStrToDStr(value))
}

func SynCustomHighlighter_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynCustomHighlighter_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynCustomHighlighter_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynCustomHighlighter_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynCustomHighlighter_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynCustomHighlighter_StaticClassType").Call()
	return TClass(r)
}
