package lcl

import "unsafe"

type TChartTextElement struct {
	IChartTextElement
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsChartTextElement(obj any) *TChartTextElement {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChartTextElement{instance: instance, ptr: ptr}
}

// TODO: TChartTextElement
