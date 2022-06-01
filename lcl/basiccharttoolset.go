package lcl

import "unsafe"

type TBasicChartToolset struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsBasicChartToolset(obj any) *TBasicChartToolset {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TBasicChartToolset{instance: instance, ptr: ptr}
}

// TODO: TBasicChartToolset
