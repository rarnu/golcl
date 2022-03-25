package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"time"
	"unsafe"
)

//--------------------------- TIniFile ---------------------------

func IniFile_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("IniFile_Create").Call(obj)
	return ret
}

func IniFile_Free(obj uintptr) {
	_, _, _ = getLazyProc("IniFile_Free").Call(obj)
}

func IniFile_ReadString(obj uintptr, Section string, Ident string, Default string) string {
	ret, _, _ := getLazyProc("IniFile_ReadString").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), GoStrToDStr(Default))
	return DStrToGoStr(ret)
}

func IniFile_WriteString(obj uintptr, Section string, Ident string, Value string) {
	_, _, _ = getLazyProc("IniFile_WriteString").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), GoStrToDStr(Value))
}

func IniFile_ReadSections(obj uintptr, Strings uintptr) {
	_, _, _ = getLazyProc("IniFile_ReadSections").Call(obj, Strings)
}

func IniFile_ReadSectionValues(obj uintptr, Section string, Strings uintptr) {
	_, _, _ = getLazyProc("IniFile_ReadSectionValues").Call(obj, GoStrToDStr(Section), Strings)
}

func IniFile_EraseSection(obj uintptr, Section string) {
	_, _, _ = getLazyProc("IniFile_EraseSection").Call(obj, GoStrToDStr(Section))
}

func IniFile_DeleteKey(obj uintptr, Section string, Ident string) {
	_, _, _ = getLazyProc("IniFile_DeleteKey").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident))
}

func IniFile_UpdateFile(obj uintptr) {
	_, _, _ = getLazyProc("IniFile_UpdateFile").Call(obj)
}

func IniFile_SectionExists(obj uintptr, Section string) bool {
	ret, _, _ := getLazyProc("IniFile_SectionExists").Call(obj, GoStrToDStr(Section))
	return DBoolToGoBool(ret)
}

func IniFile_ReadInteger(obj uintptr, Section string, Ident string, Default int32) int32 {
	ret, _, _ := getLazyProc("IniFile_ReadInteger").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), uintptr(Default))
	return int32(ret)
}

func IniFile_WriteInteger(obj uintptr, Section string, Ident string, Value int32) {
	_, _, _ = getLazyProc("IniFile_WriteInteger").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), uintptr(Value))
}

func IniFile_ReadBool(obj uintptr, Section string, Ident string, Default bool) bool {
	ret, _, _ := getLazyProc("IniFile_ReadBool").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), GoBoolToDBool(Default))
	return DBoolToGoBool(ret)
}

func IniFile_WriteBool(obj uintptr, Section string, Ident string, Value bool) {
	_, _, _ = getLazyProc("IniFile_WriteBool").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident), GoBoolToDBool(Value))
}

func IniFile_ReadDate(obj uintptr, Section string, Name string, Default time.Time) time.Time {
	var ret int64
	_, _, _ = getLazyProc("IniFile_ReadDate").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Default.Unix()), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func IniFile_ReadDateTime(obj uintptr, Section string, Name string, Default time.Time) time.Time {
	var ret int64
	_, _, _ = getLazyProc("IniFile_ReadDateTime").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Default.Unix()), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func IniFile_ReadFloat(obj uintptr, Section string, Name string, Default float64) float64 {
	var ret float64
	_, _, _ = getLazyProc("IniFile_ReadFloat").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func IniFile_ReadTime(obj uintptr, Section string, Name string, Default time.Time) time.Time {
	var ret int64
	_, _, _ = getLazyProc("IniFile_ReadTime").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Default.Unix()), uintptr(unsafe.Pointer(&ret)))
	return time.Unix(ret, 0)
}

func IniFile_WriteDate(obj uintptr, Section string, Name string, Value time.Time) {
	_, _, _ = getLazyProc("IniFile_WriteDate").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Value.Unix()))
}

func IniFile_WriteDateTime(obj uintptr, Section string, Name string, Value time.Time) {
	_, _, _ = getLazyProc("IniFile_WriteDateTime").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Value.Unix()))
}

func IniFile_WriteFloat(obj uintptr, Section string, Name string, Value float64) {
	_, _, _ = getLazyProc("IniFile_WriteFloat").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(unsafe.Pointer(&Value)))
}

func IniFile_WriteTime(obj uintptr, Section string, Name string, Value time.Time) {
	_, _, _ = getLazyProc("IniFile_WriteTime").Call(obj, GoStrToDStr(Section), GoStrToDStr(Name), uintptr(Value.Unix()))
}

func IniFile_ValueExists(obj uintptr, Section string, Ident string) bool {
	ret, _, _ := getLazyProc("IniFile_ValueExists").Call(obj, GoStrToDStr(Section), GoStrToDStr(Ident))
	return DBoolToGoBool(ret)
}

func IniFile_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("IniFile_ClassType").Call(obj)
	return TClass(ret)
}

func IniFile_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("IniFile_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func IniFile_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("IniFile_InstanceSize").Call(obj)
	return int32(ret)
}

func IniFile_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("IniFile_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func IniFile_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("IniFile_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func IniFile_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("IniFile_GetHashCode").Call(obj)
	return int32(ret)
}

func IniFile_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("IniFile_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func IniFile_GetFileName(obj uintptr) string {
	ret, _, _ := getLazyProc("IniFile_GetFileName").Call(obj)
	return DStrToGoStr(ret)
}

func IniFile_StaticClassType() TClass {
	r, _, _ := getLazyProc("IniFile_StaticClassType").Call()
	return TClass(r)
}
