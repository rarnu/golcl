//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"unsafe"

	"github.com/rarnu/golcl/lcl/types"
)

// IObject 共公的对象接口
type IObject interface {
	Instance() uintptr
	IsValid() bool
	UnsafeAddr() unsafe.Pointer

	ClassName() string
	Free()
	GetHashCode() int32
	Equals(IObject) bool
	// ClassType DisposeOf()
	ClassType() types.TClass
	InstanceSize() int32
	InheritsFrom(types.TClass) bool

	Is() TIs
	As() TAs

	ToString() string
}
