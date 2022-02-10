//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------

package api

var (
	resFormLoadFromStream       = liblcl.NewProc("ResFormLoadFromStream")
	resFormLoadFromFile         = liblcl.NewProc("ResFormLoadFromFile")
	resFormLoadFromResourceName = liblcl.NewProc("ResFormLoadFromResourceName")
)

func ResFormLoadFromStream(obj, root uintptr) {
	_, _, _ = resFormLoadFromStream.Call(obj, root)
}

func ResFormLoadFromFile(filename string, root uintptr) {
	_, _, _ = resFormLoadFromFile.Call(GoStrToDStr(filename), root)
}

func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	_, _, _ = resFormLoadFromResourceName.Call(instance, GoStrToDStr(resName), root)
}
