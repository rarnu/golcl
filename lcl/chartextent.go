package lcl

import "unsafe"

type TChartExtent struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsChartExtent(obj any) *TChartExtent {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChartExtent{instance: instance, ptr: ptr}
}

// TODO: TChartExtent
