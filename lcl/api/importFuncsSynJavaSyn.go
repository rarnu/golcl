package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynJavaSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_Create").Call(owner)
	return r
}

func SynJavaSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynJavaSyn_Free").Call(obj)
}

func SynJavaSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynJavaSyn_StaticClassType").Call()
	return TClass(r)
}

func SynJavaSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynJavaSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynJavaSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynJavaSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynJavaSyn_GetAnnotationAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetAnnotationAttri").Call(obj)
	return r
}

func SynJavaSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetCommentAttri").Call(obj)
	return r
}

func SynJavaSyn_GetDocumentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetDocumentAttri").Call(obj)
	return r
}

func SynJavaSyn_GetIdentifierAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetIdentifierAttri").Call(obj)
	return r
}

func SynJavaSyn_GetInvalidAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetInvalidAttri").Call(obj)
	return r
}

func SynJavaSyn_GetKeyAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetKeyAttri").Call(obj)
	return r
}

func SynJavaSyn_GetNumberAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetNumberAttri").Call(obj)
	return r
}

func SynJavaSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynJavaSyn_GetStringAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetStringAttri").Call(obj)
	return r
}

func SynJavaSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynJavaSyn_GetSymbolAttri").Call(obj)
	return r
}
