//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows

package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

var (
	dSendMessage         = liblcl.NewProc("DSendMessage")
	dPostMessage         = liblcl.NewProc("DPostMessage")
	dIsIconic            = liblcl.NewProc("DIsIconic")
	dIsWindow            = liblcl.NewProc("DIsWindow")
	dIsZoomed            = liblcl.NewProc("DIsZoomed")
	dIsWindowVisible     = liblcl.NewProc("DIsWindowVisible")
	dGetDC               = liblcl.NewProc("DGetDC")
	dReleaseDC           = liblcl.NewProc("DReleaseDC")
	dSetForegroundWindow = liblcl.NewProc("DSetForegroundWindow")
	dWindowFromPoint     = liblcl.NewProc("DWindowFromPoint")
)

func DSendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := dSendMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r
}

func DPostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	r, _, _ := dPostMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r != 0
}

func DIsIconic(hWnd HWND) bool {
	r, _, _ := dIsIconic.Call(hWnd)
	return r != 0
}

func DIsWindow(hWnd HWND) bool {
	r, _, _ := dIsWindow.Call(hWnd)
	return r != 0
}

func DIsZoomed(hWnd HWND) bool {
	r, _, _ := dIsZoomed.Call(hWnd)
	return r != 0
}

func DIsWindowVisible(hWnd HWND) bool {
	r, _, _ := dIsWindowVisible.Call(hWnd)
	return r != 0
}

func DGetDC(hWnd HWND) HDC {
	r, _, _ := dGetDC.Call(hWnd)
	return r
}

func DReleaseDC(hWnd HWND, dc HDC) int {
	r, _, _ := dReleaseDC.Call(hWnd, dc)
	return int(r)
}

func DSetForegroundWindow(hWnd HWND) bool {
	r, _, _ := dSetForegroundWindow.Call(hWnd)
	return r != 0
}

func DWindowFromPoint(point TPoint) HWND {
	r, _, _ := dWindowFromPoint.Call(uintptr(unsafe.Pointer(&point)))
	return r
}
