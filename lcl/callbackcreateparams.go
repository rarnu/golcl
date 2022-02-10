//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/rarnu/golcl/lcl/types"
)

var requestCreateParamsMap sync.Map

func addToRequestCreateParamsMap(ptr uintptr, proc reflect.Value) {
	requestCreateParamsMap.Store(ptr, proc)
}

func requestCallCreateParamsCallbackProc(ptr uintptr, params uintptr) uintptr {
	if val, ok := requestCreateParamsMap.Load(ptr); ok {
		val.(reflect.Value).Call([]reflect.Value{reflect.ValueOf((*types.TCreateParams)(unsafe.Pointer(params)))})
	}
	return 0
}
