//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Screen_Instance() uintptr {
	ret, _, _ := screen_Instance.Call()
	return ret
}
