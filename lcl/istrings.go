//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

type IStrings interface {
	IObject
	// 先只简单实现几个吧
	Count() int32
	S(int32) string
	Objects(int32) *TObject
}
