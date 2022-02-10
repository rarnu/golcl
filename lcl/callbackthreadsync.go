//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/rarnu/golcl/lcl/api"
)

// 线程同步的，独立出来
func threadSyncCallbackProc() uintptr {
	fn := api.ThreadSyncCallbackFn()
	if fn != nil {
		fn()
	}
	return 0
}
