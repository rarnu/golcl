package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	"unsafe"
)

type TChart struct {
	IChart
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewChart(owner IComponent) *TChart {
	c := new(TChart)
	c.instance = Chart_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsChart(obj any) *TChart {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChart{instance: instance, ptr: ptr}
}
