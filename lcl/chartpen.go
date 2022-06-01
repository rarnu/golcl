package lcl

import "unsafe"

type TChartPen struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsChartPen(obj any) *TChartPen {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChartPen{instance: instance, ptr: ptr}
}

// TODO: TChartPen
