//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

// Printer_Instance Printer
func Printer_Instance() uintptr {
	r, _, _ := printer_Instance.Call()
	return r
}

func Printer_SetPrinter(obj uintptr, aName string) {
	_, _, _ = printer_SetPrinter.Call(obj, GoStrToDStr(aName))
}
