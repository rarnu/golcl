//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/rarnu/golcl/lcl/win"
)

// NewRegistryAllAccess 所有访问权限
func NewRegistryAllAccess() *TRegistry {
	return NewRegistry(win.KEY_ALL_ACCESS)
}
