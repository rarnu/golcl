package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

//--------------------------- TRegistry ---------------------------

func Registry_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Registry_Create").Call(obj)
	return ret
}

func Registry_Free(obj uintptr) {
	getLazyProc("Registry_Free").Call(obj)
}

func Registry_CloseKey(obj uintptr) {
	getLazyProc("Registry_CloseKey").Call(obj)
}

func Registry_CreateKey(obj uintptr, Key string) bool {
	ret, _, _ := getLazyProc("Registry_CreateKey").Call(obj, GoStrToDStr(Key))
	return DBoolToGoBool(ret)
}

func Registry_DeleteKey(obj uintptr, Key string) bool {
	ret, _, _ := getLazyProc("Registry_DeleteKey").Call(obj, GoStrToDStr(Key))
	return DBoolToGoBool(ret)
}

func Registry_DeleteValue(obj uintptr, Name string) bool {
	ret, _, _ := getLazyProc("Registry_DeleteValue").Call(obj, GoStrToDStr(Name))
	return DBoolToGoBool(ret)
}

func Registry_HasSubKeys(obj uintptr) bool {
	ret, _, _ := getLazyProc("Registry_HasSubKeys").Call(obj)
	return DBoolToGoBool(ret)
}

func Registry_KeyExists(obj uintptr, Key string) bool {
	ret, _, _ := getLazyProc("Registry_KeyExists").Call(obj, GoStrToDStr(Key))
	return DBoolToGoBool(ret)
}

func Registry_LoadKey(obj uintptr, Key string, FileName string) bool {
	ret, _, _ := getLazyProc("Registry_LoadKey").Call(obj, GoStrToDStr(Key), GoStrToDStr(FileName))
	return DBoolToGoBool(ret)
}

func Registry_MoveKey(obj uintptr, OldName string, NewName string, Delete bool) {
	getLazyProc("Registry_MoveKey").Call(obj, GoStrToDStr(OldName), GoStrToDStr(NewName), GoBoolToDBool(Delete))
}

func Registry_OpenKey(obj uintptr, Key string, CanCreate bool) bool {
	ret, _, _ := getLazyProc("Registry_OpenKey").Call(obj, GoStrToDStr(Key), GoBoolToDBool(CanCreate))
	return DBoolToGoBool(ret)
}

func Registry_OpenKeyReadOnly(obj uintptr, Key string) bool {
	ret, _, _ := getLazyProc("Registry_OpenKeyReadOnly").Call(obj, GoStrToDStr(Key))
	return DBoolToGoBool(ret)
}

func Registry_ReadBool(obj uintptr, Name string) bool {
	ret, _, _ := getLazyProc("Registry_ReadBool").Call(obj, GoStrToDStr(Name))
	return DBoolToGoBool(ret)
}

func Registry_ReadDate(obj uintptr, Name string) time.Time {
	var ret int64
	getLazyProc("Registry_ReadDate").Call(obj, GoStrToDStr(Name), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func Registry_ReadDateTime(obj uintptr, Name string) time.Time {
	var ret int64
	getLazyProc("Registry_ReadDateTime").Call(obj, GoStrToDStr(Name), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func Registry_ReadFloat(obj uintptr, Name string) float64 {
	var ret float64
	getLazyProc("Registry_ReadFloat").Call(obj, GoStrToDStr(Name), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Registry_ReadInteger(obj uintptr, Name string) int32 {
	ret, _, _ := getLazyProc("Registry_ReadInteger").Call(obj, GoStrToDStr(Name))
	return int32(ret)
}

func Registry_ReadString(obj uintptr, Name string) string {
	ret, _, _ := getLazyProc("Registry_ReadString").Call(obj, GoStrToDStr(Name))
	return DStrToGoStr(ret)
}

func Registry_ReadTime(obj uintptr, Name string) time.Time {
	var ret int64
	getLazyProc("Registry_ReadTime").Call(obj, GoStrToDStr(Name), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func Registry_RegistryConnect(obj uintptr, UNCName string) bool {
	ret, _, _ := getLazyProc("Registry_RegistryConnect").Call(obj, GoStrToDStr(UNCName))
	return DBoolToGoBool(ret)
}

func Registry_RenameValue(obj uintptr, OldName string, NewName string) {
	getLazyProc("Registry_RenameValue").Call(obj, GoStrToDStr(OldName), GoStrToDStr(NewName))
}

func Registry_ReplaceKey(obj uintptr, Key string, FileName string, BackUpFileName string) bool {
	ret, _, _ := getLazyProc("Registry_ReplaceKey").Call(obj, GoStrToDStr(Key), GoStrToDStr(FileName), GoStrToDStr(BackUpFileName))
	return DBoolToGoBool(ret)
}

func Registry_RestoreKey(obj uintptr, Key string, FileName string) bool {
	ret, _, _ := getLazyProc("Registry_RestoreKey").Call(obj, GoStrToDStr(Key), GoStrToDStr(FileName))
	return DBoolToGoBool(ret)
}

func Registry_SaveKey(obj uintptr, Key string, FileName string) bool {
	ret, _, _ := getLazyProc("Registry_SaveKey").Call(obj, GoStrToDStr(Key), GoStrToDStr(FileName))
	return DBoolToGoBool(ret)
}

func Registry_UnLoadKey(obj uintptr, Key string) bool {
	ret, _, _ := getLazyProc("Registry_UnLoadKey").Call(obj, GoStrToDStr(Key))
	return DBoolToGoBool(ret)
}

func Registry_ValueExists(obj uintptr, Name string) bool {
	ret, _, _ := getLazyProc("Registry_ValueExists").Call(obj, GoStrToDStr(Name))
	return DBoolToGoBool(ret)
}

func Registry_WriteBool(obj uintptr, Name string, Value bool) {
	getLazyProc("Registry_WriteBool").Call(obj, GoStrToDStr(Name), GoBoolToDBool(Value))
}

func Registry_WriteDate(obj uintptr, Name string, Value time.Time) {
	getLazyProc("Registry_WriteDate").Call(obj, GoStrToDStr(Name), uintptr(Value.Unix()))
}

func Registry_WriteDateTime(obj uintptr, Name string, Value time.Time) {
	getLazyProc("Registry_WriteDateTime").Call(obj, GoStrToDStr(Name), uintptr(Value.Unix()))
}

func Registry_WriteFloat(obj uintptr, Name string, Value float64) {
	getLazyProc("Registry_WriteFloat").Call(obj, GoStrToDStr(Name), uintptr(unsafe.Pointer(&Value)))
}

func Registry_WriteInteger(obj uintptr, Name string, Value int32) {
	getLazyProc("Registry_WriteInteger").Call(obj, GoStrToDStr(Name), uintptr(Value))
}

func Registry_WriteString(obj uintptr, Name string, Value string) {
	getLazyProc("Registry_WriteString").Call(obj, GoStrToDStr(Name), GoStrToDStr(Value))
}

func Registry_WriteExpandString(obj uintptr, Name string, Value string) {
	getLazyProc("Registry_WriteExpandString").Call(obj, GoStrToDStr(Name), GoStrToDStr(Value))
}

func Registry_WriteTime(obj uintptr, Name string, Value time.Time) {
	getLazyProc("Registry_WriteTime").Call(obj, GoStrToDStr(Name), uintptr(Value.Unix()))
}

func Registry_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Registry_ClassType").Call(obj)
	return TClass(ret)
}

func Registry_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Registry_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Registry_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Registry_InstanceSize").Call(obj)
	return int32(ret)
}

func Registry_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Registry_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Registry_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Registry_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Registry_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Registry_GetHashCode").Call(obj)
	return int32(ret)
}

func Registry_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Registry_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Registry_GetCurrentKey(obj uintptr) HKEY {
	ret, _, _ := getLazyProc("Registry_GetCurrentKey").Call(obj)
	return HKEY(ret)
}

func Registry_GetCurrentPath(obj uintptr) string {
	ret, _, _ := getLazyProc("Registry_GetCurrentPath").Call(obj)
	return DStrToGoStr(ret)
}

func Registry_GetLazyWrite(obj uintptr) bool {
	ret, _, _ := getLazyProc("Registry_GetLazyWrite").Call(obj)
	return DBoolToGoBool(ret)
}

func Registry_SetLazyWrite(obj uintptr, value bool) {
	getLazyProc("Registry_SetLazyWrite").Call(obj, GoBoolToDBool(value))
}

func Registry_GetLastError(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Registry_GetLastError").Call(obj)
	return int32(ret)
}

func Registry_GetLastErrorMsg(obj uintptr) string {
	ret, _, _ := getLazyProc("Registry_GetLastErrorMsg").Call(obj)
	return DStrToGoStr(ret)
}

func Registry_GetRootKey(obj uintptr) HKEY {
	ret, _, _ := getLazyProc("Registry_GetRootKey").Call(obj)
	return HKEY(ret)
}

func Registry_SetRootKey(obj uintptr, value HKEY) {
	getLazyProc("Registry_SetRootKey").Call(obj, uintptr(value))
}

func Registry_GetAccess(obj uintptr) uint32 {
	ret, _, _ := getLazyProc("Registry_GetAccess").Call(obj)
	return uint32(ret)
}

func Registry_SetAccess(obj uintptr, value uint32) {
	getLazyProc("Registry_SetAccess").Call(obj, uintptr(value))
}

func Registry_StaticClassType() TClass {
	r, _, _ := getLazyProc("Registry_StaticClassType").Call()
	return TClass(r)
}
