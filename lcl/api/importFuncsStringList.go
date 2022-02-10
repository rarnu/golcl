package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TStringList ---------------------------

func StringList_Create() uintptr {
	ret, _, _ := getLazyProc("StringList_Create").Call()
	return ret
}

func StringList_Free(obj uintptr) {
	getLazyProc("StringList_Free").Call(obj)
}

func StringList_Add(obj uintptr, S string) int32 {
	ret, _, _ := getLazyProc("StringList_Add").Call(obj, GoStrToDStr(S))
	return int32(ret)
}

func StringList_AddObject(obj uintptr, S string, AObject uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_AddObject").Call(obj, GoStrToDStr(S), AObject)
	return int32(ret)
}

func StringList_Assign(obj uintptr, Source uintptr) {
	getLazyProc("StringList_Assign").Call(obj, Source)
}

func StringList_Clear(obj uintptr) {
	getLazyProc("StringList_Clear").Call(obj)
}

func StringList_Delete(obj uintptr, Index int32) {
	getLazyProc("StringList_Delete").Call(obj, uintptr(Index))
}

func StringList_IndexOf(obj uintptr, S string) int32 {
	ret, _, _ := getLazyProc("StringList_IndexOf").Call(obj, GoStrToDStr(S))
	return int32(ret)
}

func StringList_Insert(obj uintptr, Index int32, S string) {
	getLazyProc("StringList_Insert").Call(obj, uintptr(Index), GoStrToDStr(S))
}

func StringList_InsertObject(obj uintptr, Index int32, S string, AObject uintptr) {
	getLazyProc("StringList_InsertObject").Call(obj, uintptr(Index), GoStrToDStr(S), AObject)
}

func StringList_Append(obj uintptr, S string) {
	getLazyProc("StringList_Append").Call(obj, GoStrToDStr(S))
}

func StringList_BeginUpdate(obj uintptr) {
	getLazyProc("StringList_BeginUpdate").Call(obj)
}

func StringList_EndUpdate(obj uintptr) {
	getLazyProc("StringList_EndUpdate").Call(obj)
}

func StringList_Equals(obj uintptr, Strings uintptr) bool {
	ret, _, _ := getLazyProc("StringList_Equals").Call(obj, Strings)
	return DBoolToGoBool(ret)
}

func StringList_IndexOfName(obj uintptr, Name string) int32 {
	ret, _, _ := getLazyProc("StringList_IndexOfName").Call(obj, GoStrToDStr(Name))
	return int32(ret)
}

func StringList_IndexOfObject(obj uintptr, AObject uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_IndexOfObject").Call(obj, AObject)
	return int32(ret)
}

func StringList_LoadFromFile(obj uintptr, FileName string) {
	getLazyProc("StringList_LoadFromFile").Call(obj, GoStrToDStr(FileName))
}

func StringList_LoadFromStream(obj uintptr, Stream uintptr) {
	getLazyProc("StringList_LoadFromStream").Call(obj, Stream)
}

func StringList_Move(obj uintptr, CurIndex int32, NewIndex int32) {
	getLazyProc("StringList_Move").Call(obj, uintptr(CurIndex), uintptr(NewIndex))
}

func StringList_SaveToFile(obj uintptr, FileName string) {
	getLazyProc("StringList_SaveToFile").Call(obj, GoStrToDStr(FileName))
}

func StringList_SaveToStream(obj uintptr, Stream uintptr) {
	getLazyProc("StringList_SaveToStream").Call(obj, Stream)
}

func StringList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("StringList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func StringList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("StringList_ClassType").Call(obj)
	return TClass(ret)
}

func StringList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("StringList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func StringList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_InstanceSize").Call(obj)
	return int32(ret)
}

func StringList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("StringList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func StringList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_GetHashCode").Call(obj)
	return int32(ret)
}

func StringList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("StringList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func StringList_GetSorted(obj uintptr) bool {
	ret, _, _ := getLazyProc("StringList_GetSorted").Call(obj)
	return DBoolToGoBool(ret)
}

func StringList_SetSorted(obj uintptr, value bool) {
	getLazyProc("StringList_SetSorted").Call(obj, GoBoolToDBool(value))
}

func StringList_SetOnChange(obj uintptr, fn interface{}) {
	getLazyProc("StringList_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func StringList_SetOnChanging(obj uintptr, fn interface{}) {
	getLazyProc("StringList_SetOnChanging").Call(obj, addEventToMap(obj, fn))
}

func StringList_GetCapacity(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_GetCapacity").Call(obj)
	return int32(ret)
}

func StringList_SetCapacity(obj uintptr, value int32) {
	getLazyProc("StringList_SetCapacity").Call(obj, uintptr(value))
}

func StringList_GetCommaText(obj uintptr) string {
	ret, _, _ := getLazyProc("StringList_GetCommaText").Call(obj)
	return DStrToGoStr(ret)
}

func StringList_SetCommaText(obj uintptr, value string) {
	getLazyProc("StringList_SetCommaText").Call(obj, GoStrToDStr(value))
}

func StringList_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("StringList_GetCount").Call(obj)
	return int32(ret)
}

func StringList_GetDelimiter(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("StringList_GetDelimiter").Call(obj)
	return uint16(ret)
}

func StringList_SetDelimiter(obj uintptr, value uint16) {
	getLazyProc("StringList_SetDelimiter").Call(obj, uintptr(value))
}

func StringList_GetNameValueSeparator(obj uintptr) uint16 {
	ret, _, _ := getLazyProc("StringList_GetNameValueSeparator").Call(obj)
	return uint16(ret)
}

func StringList_SetNameValueSeparator(obj uintptr, value uint16) {
	getLazyProc("StringList_SetNameValueSeparator").Call(obj, uintptr(value))
}

func StringList_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("StringList_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func StringList_SetText(obj uintptr, value string) {
	getLazyProc("StringList_SetText").Call(obj, GoStrToDStr(value))
}

func StringList_GetObjects(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("StringList_GetObjects").Call(obj, uintptr(Index))
	return ret
}

func StringList_SetObjects(obj uintptr, Index int32, value uintptr) {
	getLazyProc("StringList_SetObjects").Call(obj, uintptr(Index), value)
}

func StringList_GetValues(obj uintptr, Name string) string {
	ret, _, _ := getLazyProc("StringList_GetValues").Call(obj, GoStrToDStr(Name))
	return DStrToGoStr(ret)
}

func StringList_SetValues(obj uintptr, Name string, value string) {
	getLazyProc("StringList_SetValues").Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func StringList_GetValueFromIndex(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("StringList_GetValueFromIndex").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func StringList_SetValueFromIndex(obj uintptr, Index int32, value string) {
	getLazyProc("StringList_SetValueFromIndex").Call(obj, uintptr(Index), GoStrToDStr(value))
}

func StringList_GetStrings(obj uintptr, Index int32) string {
	ret, _, _ := getLazyProc("StringList_GetStrings").Call(obj, uintptr(Index))
	return DStrToGoStr(ret)
}

func StringList_SetStrings(obj uintptr, Index int32, value string) {
	getLazyProc("StringList_SetStrings").Call(obj, uintptr(Index), GoStrToDStr(value))
}

func StringList_StaticClassType() TClass {
	r, _, _ := getLazyProc("StringList_StaticClassType").Call()
	return TClass(r)
}
