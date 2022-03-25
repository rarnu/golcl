package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TStrings ---------------------------

func Strings_Create() uintptr {
	ret, _, _ := getLazyProc("Strings_Create").Call()
	return ret
}

func Strings_Free(obj uintptr) {
	_, _, _ = getLazyProc("Strings_Free").Call(obj)
}

func Strings_Add(obj uintptr, S string) int32 {
	ret, _, _ := getLazyProc("Strings_Add").Call(obj, GoStrToDStr(S))
	return int32(ret)
}

func Strings_AddObject(obj uintptr, S string, AObject uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_AddObject").Call(obj, GoStrToDStr(S), AObject)
	return int32(ret)
}

func Strings_Append(obj uintptr, S string) {
	_, _, _ = getLazyProc("Strings_Append").Call(obj, GoStrToDStr(S))
}

func Strings_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Strings_Assign").Call(obj, Source)
}

func Strings_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("Strings_BeginUpdate").Call(obj)
}

func Strings_Clear(obj uintptr) {
	_, _, _ = getLazyProc("Strings_Clear").Call(obj)
}

func Strings_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("Strings_Delete").Call(obj, uintptr(Index))
}

func Strings_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("Strings_EndUpdate").Call(obj)
}

func Strings_Equals(obj uintptr, Strings uintptr) bool {
	ret, _, _ := getLazyProc("Strings_Equals").Call(obj, Strings)
	return DBoolToGoBool(ret)
}

func Strings_IndexOf(obj uintptr, S string) int32 {
	ret, _, _ := getLazyProc("Strings_IndexOf").Call(obj, GoStrToDStr(S))
	return int32(ret)
}

func Strings_IndexOfName(obj uintptr, Name string) int32 {
	ret, _, _ := getLazyProc("Strings_IndexOfName").Call(obj, GoStrToDStr(Name))
	return int32(ret)
}

func Strings_IndexOfObject(obj uintptr, AObject uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_IndexOfObject").Call(obj, AObject)
	return int32(ret)
}

func Strings_Insert(obj uintptr, Index int32, S string) {
	_, _, _ = getLazyProc("Strings_Insert").Call(obj, uintptr(Index), GoStrToDStr(S))
}

func Strings_InsertObject(obj uintptr, Index int32, S string, AObject uintptr) {
	_, _, _ = getLazyProc("Strings_InsertObject").Call(obj, uintptr(Index), GoStrToDStr(S), AObject)
}

func Strings_LoadFromFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("Strings_LoadFromFile").Call(obj, GoStrToDStr(FileName))
}

func Strings_LoadFromStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Strings_LoadFromStream").Call(obj, Stream)
}

func Strings_Move(obj uintptr, CurIndex int32, NewIndex int32) {
	_, _, _ = getLazyProc("Strings_Move").Call(obj, uintptr(CurIndex), uintptr(NewIndex))
}

func Strings_SaveToFile(obj uintptr, FileName string) {
	_, _, _ = getLazyProc("Strings_SaveToFile").Call(obj, GoStrToDStr(FileName))
}

func Strings_SaveToStream(obj uintptr, Stream uintptr) {
	_, _, _ = getLazyProc("Strings_SaveToStream").Call(obj, Stream)
}

func Strings_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Strings_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Strings_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Strings_ClassType").Call(obj)
	return TClass(ret)
}

func Strings_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Strings_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Strings_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_InstanceSize").Call(obj)
	return int32(ret)
}

func Strings_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Strings_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Strings_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_GetHashCode").Call(obj)
	return int32(ret)
}

func Strings_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Strings_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Strings_GetCapacity(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_GetCapacity").Call(obj)
	return int32(ret)
}

func Strings_SetCapacity(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Strings_SetCapacity").Call(obj, uintptr(value))
}

func Strings_GetCommaText(obj uintptr) string {
	ret, _, _ := getLazyProc("Strings_GetCommaText").Call(obj)
	return DStrToGoStr(ret)
}

func Strings_SetCommaText(obj uintptr, value string) {
	_, _, _ = getLazyProc("Strings_SetCommaText").Call(obj, GoStrToDStr(value))
}

func Strings_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Strings_GetCount").Call(obj)
	return int32(ret)
}

func Strings_GetDelimiter(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("Strings_GetDelimiter").Call(obj)
	return uint16(ret)
}

func Strings_SetDelimiter(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("Strings_SetDelimiter").Call(obj, uintptr(value))
}

func Strings_GetNameValueSeparator(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("Strings_GetNameValueSeparator").Call(obj)
	return uint16(ret)
}

func Strings_SetNameValueSeparator(obj uintptr, value uint16) {
	_, _, _ = getLazyProc("Strings_SetNameValueSeparator").Call(obj, uintptr(value))
}

func Strings_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("Strings_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func Strings_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("Strings_SetText").Call(obj, GoStrToDStr(value))
}

func Strings_GetObjects(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("Strings_GetObjects").Call(obj, uintptr(Index))
	return ret
}

func Strings_SetObjects(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("Strings_SetObjects").Call(obj, uintptr(Index), value)
}

func Strings_GetValues(obj uintptr, Name string) string {
	ret, _, _ := getLazyProc("Strings_GetValues").Call(obj, GoStrToDStr(Name))
	return DStrToGoStr(ret)
}

func Strings_SetValues(obj uintptr, Name string, value string) {
	_, _, _ = getLazyProc("Strings_SetValues").Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func Strings_GetValueFromIndex(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("Strings_GetValueFromIndex").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func Strings_SetValueFromIndex(obj uintptr, Index int32, value string) {
	_, _, _ = getLazyProc("Strings_SetValueFromIndex").Call(obj, uintptr(Index), GoStrToDStr(value))
}

func Strings_GetStrings(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("Strings_GetStrings").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func Strings_SetStrings(obj uintptr, Index int32, value string) {
	_, _, _ = getLazyProc("Strings_SetStrings").Call(obj, uintptr(Index), GoStrToDStr(value))
}

func Strings_StaticClassType() TClass {
	r, _, _ := getLazyProc("Strings_StaticClassType").Call()
	return TClass(r)
}
