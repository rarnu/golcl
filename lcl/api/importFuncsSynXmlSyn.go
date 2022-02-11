package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func SynXMLSyn_Create(owner uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_Create").Call(owner)
	return r
}

func SynXMLSyn_Free(obj uintptr) {
	_, _, _ = getLazyProc("SynXMLSyn_Free").Call(obj)
}

func SynXMLSyn_StaticClassType() TClass {
	r, _, _ := getLazyProc("SynXMLSyn_StaticClassType").Call()
	return TClass(r)
}

func SynXMLSyn_GetEnabled(obj uintptr) bool {
	r, _, _ := getLazyProc("SynXMLSyn_GetEnabled").Call(obj)
	return DBoolToGoBool(r)
}

func SynXMLSyn_SetEnabled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynXMLSyn_SetEnabled").Call(obj, GoBoolToDBool(value))
}

func SynXMLSyn_GetElementAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetElementAttri").Call(obj)
	return r
}

func SynXMLSyn_GetAttributeAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetAttributeAttri").Call(obj)
	return r
}

func SynXMLSyn_GetNamespaceAttributeAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetNamespaceAttributeAttri").Call(obj)
	return r
}

func SynXMLSyn_GetAttributeValueAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetAttributeValueAttri").Call(obj)
	return r
}

func SynXMLSyn_GetNamespaceAttributeValueAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetNamespaceAttributeValueAttri").Call(obj)
	return r
}

func SynXMLSyn_GetTextAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetTextAttri").Call(obj)
	return r
}

func SynXMLSyn_GetCDATAAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetCDATAAttri").Call(obj)
	return r
}

func SynXMLSyn_GetEntityRefAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetEntityRefAttri").Call(obj)
	return r
}

func SynXMLSyn_GetProcessingInstructionAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetProcessingInstructionAttri").Call(obj)
	return r
}

func SynXMLSyn_GetCommentAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetCommentAttri").Call(obj)
	return r
}

func SynXMLSyn_GetDocTypeAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetDocTypeAttri").Call(obj)
	return r
}

func SynXMLSyn_GetSpaceAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetSpaceAttri").Call(obj)
	return r
}

func SynXMLSyn_GetSymbolAttri(obj uintptr) uintptr {
	r, _, _ := getLazyProc("SynXMLSyn_GetSymbolAttri").Call(obj)
	return r
}

func SynXMLSyn_GetWantBracesParsed(obj uintptr) bool {
	r, _, _ := getLazyProc("SynXMLSyn_GetWantBracesParsed").Call(obj)
	return DBoolToGoBool(r)
}

func SynXMLSyn_SetWantBracesParsed(obj uintptr, value bool) {
	_, _, _ = getLazyProc("SynXMLSyn_SetWantBracesParsed").Call(obj, GoBoolToDBool(value))
}
