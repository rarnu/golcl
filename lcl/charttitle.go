package lcl

import "unsafe"

type TChartTitle struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsChartTitle(obj any) *TChartTitle {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChartTitle{instance: instance, ptr: ptr}
}

// TODO: TChartTitle
