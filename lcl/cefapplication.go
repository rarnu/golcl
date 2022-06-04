package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	"unsafe"
)

type TCefApplication struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// NewApplication 创建一个新的对象。
func NewCefApplication() *TCefApplication {
	a := new(TCefApplication)
	a.instance = CefApplication_Create()
	a.ptr = unsafe.Pointer(a.instance)
	return a
}

// AsApplication 动态转换一个已存在的对象实例。
func AsCefApplication(obj any) *TCefApplication {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TCefApplication{instance: instance, ptr: ptr}
}

// ApplicationFromInst 新建一个对象来自已经存在的对象实例指针。
func CefApplicationFromInst(inst uintptr) *TCefApplication {
	return AsCefApplication(inst)
}

// ApplicationFromObj 新建一个对象来自已经存在的对象实例。
func CefApplicationFromObj(obj IObject) *TCefApplication {
	return AsCefApplication(obj)
}

// ApplicationFromUnsafePointer 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// Deprecated: use AsApplication.
func CefApplicationFromUnsafePointer(ptr unsafe.Pointer) *TCefApplication {
	return AsCefApplication(ptr)
}

// Free 释放对象。
func (a *TCefApplication) Free() {
	if a.instance != 0 {
		CefApplication_Free(a.instance)
		a.instance, a.ptr = 0, nullptr
	}
}

// Instance 返回对象实例指针。
func (a *TCefApplication) Instance() uintptr {
	return a.instance
}

// UnsafeAddr 获取一个不安全的地址。
func (a *TCefApplication) UnsafeAddr() unsafe.Pointer {
	return a.ptr
}

// IsValid 检测地址是否为空。
func (a *TCefApplication) IsValid() bool {
	return a.instance != 0
}

// Is 检测当前对象是否继承自目标对象。
func (a *TCefApplication) Is() TIs {
	return TIs(a.instance)
}

func (a *TCefApplication) StartMainProcess() bool {
	return CefApplication_StartMainProcess(a.instance)
}

func (a *TCefApplication) StartSubProcess() bool {
	return CefApplication_StartSubProcess(a.instance)
}

func (a *TCefApplication) NoSandbox() bool {
	return CefApplication_GetNoSandbox(a.instance)
}

func (a *TCefApplication) SetNoSandbox(b bool) {
	CefApplication_SetNoSandbox(a.instance, b)
}

func (a *TCefApplication) UserAgent() string {
	return CefApplication_GetUserAgent(a.instance)
}

func (a *TCefApplication) SetUserAgent(s string) {
	CefApplication_SetUserAgent(a.instance, s)
}

func (a *TCefApplication) JavaScriptFlags() string {
	return CefApplication_GetJavaScriptFlags(a.instance)
}

func (a *TCefApplication) SetJavaScriptFlags(s string) {
	CefApplication_SetJavaScriptFlags(a.instance, s)
}

func (a *TCefApplication) IgnoreCertificateErrors() bool {
	return CefApplication_GetIgnoreCertificateErrors(a.instance)
}

func (a *TCefApplication) SetIgnoreCertificateErrors(b bool) {
	CefApplication_SetIgnoreCertificateErrors(a.instance, b)
}

func (a *TCefApplication) EnableGPU() bool {
	return CefApplication_GetEnableGPU(a.instance)
}

func (a *TCefApplication) SetEnableGPU(b bool) {
	CefApplication_SetEnableGPU(a.instance, b)
}

func (a *TCefApplication) AllowFileAccessFromFiles() bool {
	return CefApplication_GetAllowFileAccessFromFiles(a.instance)
}

func (a *TCefApplication) SetAllowFileAccessFromFiles(b bool) {
	CefApplication_SetAllowFileAccessFromFiles(a.instance, b)
}

func (a *TCefApplication) AllowRunningInsecureContent() bool {
	return CefApplication_GetAllowRunningInsecureContent(a.instance)
}

func (a *TCefApplication) SetAllowRunningInsecureContent(b bool) {
	CefApplication_SetAllowRunningInsecureContent(a.instance, b)
}

func (a *TCefApplication) DisableJavascript() bool {
	return CefApplication_GetDisableJavascript(a.instance)
}

func (a *TCefApplication) SetDisableJavascript(b bool) {
	CefApplication_SetDisableJavascript(a.instance, b)
}

func (a *TCefApplication) AllowUniversalAccessFromFileUrls() bool {
	return CefApplication_GetAllowUniversalAccessFromFileUrls(a.instance)
}

func (a *TCefApplication) SetAllowUniversalAccessFromFileUrls(b bool) {
	CefApplication_SetAllowUniversalAccessFromFileUrls(a.instance, b)
}

func (a *TCefApplication) AllowInsecureLocalhost() bool {
	return CefApplication_GetAllowInsecureLocalhost(a.instance)
}

func (a *TCefApplication) SetAllowInsecureLocalhost(b bool) {
	CefApplication_SetAllowInsecureLocalhost(a.instance, b)
}

func (a *TCefApplication) DefaultEncoding() string {
	return CefApplication_GetDefaultEncoding(a.instance)
}

func (a *TCefApplication) SetDefaultEncoding(s string) {
	CefApplication_SetDefaultEncoding(a.instance, s)
}

func (a *TCefApplication) ChromeVersion() string {
	return CefApplication_GetChromeVersion(a.instance)
}

func InitGlobalCEFApp() {
	CefApplication_InitGlobalCEFApp()
}

func FinalGlobalCEFApp() {
	CefApplication_FinalGlobalCEFApp()
}
