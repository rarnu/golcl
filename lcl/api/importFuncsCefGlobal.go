package api

func CefApplication_GetGlobalCEFApp() uintptr {
	ret, _, _ := getLazyProc("CefApplication_GetGlobalCEFApp").Call()
	return ret
}

func CefApplication_SetGlobalCEFApp(obj uintptr) {
	_, _, _ = getLazyProc("CefApplication_SetGlobalCEFApp").Call(obj)
}

func CefApplication_Create() uintptr {
	ret, _, _ := getLazyProc("CefApplication_Create").Call()
	return ret
}

func CefApplication_Free(obj uintptr) {
	_, _, _ = getLazyProc("CefApplication_Free").Call(obj)
}

func CefApplication_StartMainProcess(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_StartMainProcess").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_StartSubProcess(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_StartSubProcess").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_GetNoSandbox(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetNoSandbox").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetNoSandbox(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetNoSandbox").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetUserAgent(obj uintptr) string {
	ret, _, _ := getLazyProc("CefApplication_GetUserAgent").Call(obj)
	return DStrToGoStr(ret)
}

func CefApplication_SetUserAgent(obj uintptr, Value string) {
	_, _, _ = getLazyProc("CefApplication_SetUserAgent").Call(obj, GoStrToDStr(Value))
}

func CefApplication_GetJavaScriptFlags(obj uintptr) string {
	ret, _, _ := getLazyProc("CefApplication_GetJavaScriptFlags").Call(obj)
	return DStrToGoStr(ret)
}

func CefApplication_SetJavaScriptFlags(obj uintptr, Value string) {
	_, _, _ = getLazyProc("CefApplication_SetJavaScriptFlags").Call(obj, GoStrToDStr(Value))
}

func CefApplication_GetIgnoreCertificateErrors(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetIgnoreCertificateErrors").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetIgnoreCertificateErrors(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetIgnoreCertificateErrors").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetEnableGPU(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetEnableGPU").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetEnableGPU(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetEnableGPU").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetAllowFileAccessFromFiles(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetAllowFileAccessFromFiles").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetAllowFileAccessFromFiles(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetAllowFileAccessFromFiles").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetAllowRunningInsecureContent(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetAllowRunningInsecureContent").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetAllowRunningInsecureContent(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetAllowRunningInsecureContent").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetDisableJavascript(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetDisableJavascript").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetDisableJavascript(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetDisableJavascript").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetAllowUniversalAccessFromFileUrls(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetAllowUniversalAccessFromFileUrls").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetAllowUniversalAccessFromFileUrls(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetAllowUniversalAccessFromFileUrls").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetAllowInsecureLocalhost(obj uintptr) bool {
	ret, _, _ := getLazyProc("CefApplication_GetAllowInsecureLocalhost").Call(obj)
	return DBoolToGoBool(ret)
}

func CefApplication_SetAllowInsecureLocalhost(obj uintptr, Value bool) {
	_, _, _ = getLazyProc("CefApplication_SetAllowInsecureLocalhost").Call(obj, GoBoolToDBool(Value))
}

func CefApplication_GetDefaultEncoding(obj uintptr) string {
	ret, _, _ := getLazyProc("CefApplication_GetDefaultEncoding").Call(obj)
	return DStrToGoStr(ret)
}

func CefApplication_SetDefaultEncoding(obj uintptr, Value string) {
	_, _, _ = getLazyProc("CefApplication_SetDefaultEncoding").Call(obj, GoStrToDStr(Value))
}

func CefApplication_GetChromeVersion(obj uintptr) string {
	ret, _, _ := getLazyProc("CefApplication_GetChromeVersion").Call(obj)
	return DStrToGoStr(ret)
}

func CefApplication_InitGlobalCEFApp() {
	_, _, _ = getLazyProc("CefApplication_InitGlobalCEFApp").Call()
}

func CefApplication_FinalGlobalCEFApp() {
	_, _, _ = getLazyProc("CefApplication_FinalGlobalCEFApp").Call()
}
