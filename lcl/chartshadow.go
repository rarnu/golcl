package lcl

import "unsafe"

type TChartShadow struct {
	IChartTextElement
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsChartShadow(obj any) *TChartShadow {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChartShadow{instance: instance, ptr: ptr}
}

// TODO: TChartShadow
