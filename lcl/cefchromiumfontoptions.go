package lcl

import "unsafe"

type TChromiumFontOptions struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsChromiumFontOptions(obj any) *TChromiumFontOptions {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TChromiumFontOptions{instance: instance, ptr: ptr}
}

// TODO: TChromiumFontOptions
