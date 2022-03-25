//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"reflect"
	"sync"
	"unsafe"

	. "github.com/rarnu/golcl/lcl/types"
)

var (
	// 事件回调查找表
	eventCallbackMap sync.Map
	// 独立的消息事件查找表
	messageCallbackMap sync.Map
	// ThreadSync
	threadSync   sync.Mutex
	threadSyncFn func()

	// 标识，主要是解决反射时事件的函数地址获取问题
	// 当标识为true时addEventToMap直接取currentEventId的值。
	addingEvent    bool
	currentEventId uintptr
)

// DBoolToGoBool Delphi或者Lazarus的Bool类型转为Go bool
func DBoolToGoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

// GoBoolToDBool Go bool类型转为Delphi或者Lazarus的Bool类型
func GoBoolToDBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

// typedef struct { void *type; void *value; } GoInterface;
type interfacePtr struct {
	tpy uintptr
	val *uintptr
}

func IsNil(val any) bool {
	ptr := (*interfacePtr)(unsafe.Pointer(&val))
	return ptr.tpy == 0 || ptr.val == nil
}

// GetUID 用作事件的唯一id
func GetUID(v1, v2 uintptr) uintptr {
	if v1 == 0 && v2 == 0 {
		return 0
	}
	val := struct {
		v1, v2 uintptr
	}{v1, v2}
	var result uintptr
	p := (*byte)(unsafe.Pointer(&val))
	for i := 0; i < int(unsafe.Sizeof(val)); i++ {
		result = ((result << 2) | (result >> (unsafe.Sizeof(result)*8 - 2))) ^ uintptr(*p)
		p = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1))
	}
	//fmt.Println("UID:", result, ", v1=", v1, ", v2=", v2)
	return result
}

// hashOf
func hashOf(obj uintptr, val any) uintptr {
	// 如果正在使用beginAddEvent和EndAddEvent则直接取这个值。
	// 反之使用默认的行为。
	if addingEvent {
		if currentEventId > 0 {
			return currentEventId
		} else {
			return 0
		}
	}
	// 默认返回ID
	return GetUID(obj, reflect.ValueOf(val).Pointer())
}

// BeginAddEvent 以下三个函数留给自动绑定事件使用。
func BeginAddEvent() {
	addingEvent = true
	currentEventId = 0
}

func EndAddEvent() {
	addingEvent = false
	currentEventId = 0
}

func SetCurrentEventId(id uintptr) {
	currentEventId = id
}

// 将事件添加到查找表中
func addEventToMap(obj uintptr, f any) uintptr {
	p := hashOf(obj, f)
	eventCallbackMap.Store(p, f)
	return p
}

func GetAddEventToMapFn() func(obj uintptr, f any) uintptr {
	return addEventToMap
}

// EventCallbackOf 从事件表中查找指定id的函数
func EventCallbackOf(Id uintptr) (any, bool) {
	return eventCallbackMap.Load(Id)
}

// RemoveEventCallbackOf 从事件表中移除指定的函数
func RemoveEventCallbackOf(Id uintptr) {
	eventCallbackMap.Delete(Id)
}

// 添加消息事件到消息表中
func addMessageEventToMap(obj uintptr, f any) uintptr {
	p := hashOf(obj, f)
	messageCallbackMap.Store(p, f)
	return p
}

// MessageCallbackOf 从消息表中查找指定id的函数
func MessageCallbackOf(Id uintptr) (any, bool) {
	return messageCallbackMap.Load(Id)
}

// ThreadSyncCallbackFn 返回当前需要调用的线程同步函数
func ThreadSyncCallbackFn() func() {
	return threadSyncFn
}

// SetEventCallback 设置事件回调函数指针
func SetEventCallback(ptr uintptr) {
	_, _, _ = setEventCallback.Call(ptr)
}

// SetMessageCallback 设置消息事件回调函数指针
func SetMessageCallback(ptr uintptr) {
	_, _, _ = setMessageCallback.Call(ptr)
}

// SetThreadSyncCallback 设置线程同步事件回调函数指针
func SetThreadSyncCallback(ptr uintptr) {
	_, _, _ = setThreadSyncCallback.Call(ptr)
}

func SetRequestCallCreateParamsCallback(ptr uintptr) {
	_, _, _ = setRequestCallCreateParamsCallback.Call(ptr)
}

// DGetStringArrOf 从Delphi/Lazarus字符串数组中获取指定索引的值
func DGetStringArrOf(p uintptr, index int) string {
	r, _, _ := dGetStringArrOf.Call(p, uintptr(index))
	return DStrToGoStr(r)
}

// DStrLen 获取Delphi/Lazarus字符串长度
func DStrLen(p uintptr) int {
	ret, _, _ := dStrLen.Call(p)
	return int(ret)
}

// DMove Delphi/Lazarus内存操作
func DMove(src, dest uintptr, llen int) {
	_, _, _ = dMove.Call(src, dest, uintptr(llen))
}

func DShowMessage(s string) {
	_, _, _ = dShowMessage.Call(GoStrToDStr(s))
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	ret, _, _ := dMessageDlg.Call(GoStrToDStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx))
	return int32(ret)
}

func DTextToShortCut(val string) TShortCut {
	ret, _, _ := dTextToShortCut.Call(GoStrToDStr(val))
	return TShortCut(ret)
}

func DShortCutToText(val TShortCut) string {
	ret, _, _ := dShortCutToText.Call(uintptr(val))
	return DStrToGoStr(ret)
}

func DSysOpen(filename string) {
	_, _, _ = dSysOpen.Call(GoStrToDStr(filename))
}

func DExtractFilePath(filename string) string {
	r, _, _ := dExtractFilePath.Call(GoStrToDStr(filename))
	return DStrToGoStr(r)
}

func DFileExists(filename string) bool {
	r, _, _ := dFileExists.Call(GoStrToDStr(filename))
	return DBoolToGoBool(r)
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory1.Call(uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0)
	v := DBoolToGoBool(r)
	if v {
		return true, DStrToGoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory2.Call(GoStrToDStr(caption), GoStrToDStr(root), GoBoolToDBool(showHidden), uintptr(unsafe.Pointer(&ptr)))
	v := DBoolToGoBool(r)
	if v {
		return true, DStrToGoStr(ptr)
	}
	return false, ""
}

// DSynchronize 线程同步
func DSynchronize(fn func(), useMsg uintptr) {
	threadSync.Lock()
	defer threadSync.Unlock()
	threadSyncFn = fn
	_, _, _ = dSynchronize.Call(useMsg)
	threadSyncFn = nil
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	r, _, _ := dInputBox.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), GoStrToDStr(aDefault))
	return DStrToGoStr(r)
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r, _, _ := dInputQuery.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), GoStrToDStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = DStrToGoStr(strPtr)
	}
	return DBoolToGoBool(r)
}

func DPasswordBox(aCaption, aPrompt string) string {
	r, _, _ := dPasswordBox.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt))
	return DStrToGoStr(r)
}

func DInputCombo(aCaption, aPrompt string, aList uintptr) int32 {
	r, _, _ := dInputCombo.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), aList)
	return int32(r)
}

func DInputComboEx(aCaption, aPrompt string, aList uintptr, allowCustomText bool) string {
	r, _, _ := dInputComboEx.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), aList, GoBoolToDBool(allowCustomText))
	return DStrToGoStr(r)
}

// DSysLocale DSysLocaled
func DSysLocale(aInfo *TSysLocale) {
	_, _, _ = dSysLocale.Call(uintptr(unsafe.Pointer(aInfo)))
}

// DCreateURLShortCut Shortcut
func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	_, _, _ = dCreateURLShortCut.Call(GoStrToDStr(aDestPath), GoStrToDStr(aShortCutName), GoStrToDStr(aURL))
}

func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	r, _, _ := dCreateShortCut.Call(GoStrToDStr(aDestPath), GoStrToDStr(aShortCutName), GoStrToDStr(aSrcFileName),
		GoStrToDStr(aIconFileName), GoStrToDStr(aDescription), GoStrToDStr(aCmdArgs))
	return DBoolToGoBool(r)
}

// DSetPropertyValue SetProperty
func DSetPropertyValue(instance uintptr, propName, value string) {
	_, _, _ = dSetPropertyValue.Call(instance, GoStrToDStr(propName), GoStrToDStr(value))
}

func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	_, _, _ = dSetPropertySecValue.Call(instance, GoStrToDStr(propName), GoStrToDStr(secPropName), GoStrToDStr(value))
}

// DGUIDToString guid
func DGUIDToString(guid TGUID) string {
	r, _, _ := dGUIDToString.Call(uintptr(unsafe.Pointer(&guid)))
	return DStrToGoStr(r)
}

func DStringToGUID(str string) TGUID {
	var guid TGUID
	_, _, _ = dStringToGUID.Call(GoStrToDStr(str), uintptr(unsafe.Pointer(&guid)))
	return guid
}

func DCreateGUID() TGUID {
	var guid TGUID
	_, _, _ = dCreateGUID.Call(uintptr(unsafe.Pointer(&guid)))
	return guid
}

// DGetLibResourceCount LibResources
func DGetLibResourceCount() int32 {
	r, _, _ := dGetLibResourceCount.Call()
	return int32(r)
}

func DGetLibResourceItem(aIndex int32) (ret TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	_, _, _ = dGetLibResourceItem.Call(uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = DStrToGoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResource(aPtr uintptr, aValue string) {
	_, _, _ = dModifyLibResource.Call(aPtr, GoStrToDStr(aValue))
}

// DLibStringEncoding 库的信息, 获取当前库使用的字符串编码
func DLibStringEncoding() TStringEncoding {
	r, _, _ := dLibStringEncoding.Call()
	return TStringEncoding(r)
}

// DLibVersion 库版本信息
func DLibVersion() uint32 {
	r, _, _ := dLibVersion.Call()
	return uint32(r)
}

func DLibAbout() string {
	r, _, _ := dLibAbout.Call()
	return DStrToGoStr(r)
}

func DMainThreadId() uintptr {
	r, _, _ := dMainThreadId.Call()
	return r
}

func DCurrentThreadId() uintptr {
	r, _, _ := dCurrentThreadId.Call()
	return r
}

func DInitGoDll(aMainThreadId uintptr) {
	_, _, _ = dInitGoDll.Call(aMainThreadId)
}

func DFindControl(handle HWND) uintptr {
	r, _, _ := dFindControl.Call(handle)
	return r
}

func DFindLCLControl(screenPos TPoint) uintptr {
	r, _, _ := dFindLCLControl.Call(uintptr(unsafe.Pointer(&screenPos)))
	return r
}

func DFindOwnerControl(handle HWND) uintptr {
	r, _, _ := dFindOwnerControl.Call(handle)
	return r
}

func DFindControlAtPosition(position TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindControlAtPosition.Call(uintptr(unsafe.Pointer(&position)), GoBoolToDBool(allowDisabled))
	return r
}

func DFindLCLWindow(screenPos TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindLCLWindow.Call(uintptr(unsafe.Pointer(&screenPos)), GoBoolToDBool(allowDisabled))
	return r
}

func DFindDragTarget(position TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindDragTarget.Call(uintptr(unsafe.Pointer(&position)), GoBoolToDBool(allowDisabled))
	return r
}
