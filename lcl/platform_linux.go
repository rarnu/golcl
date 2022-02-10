//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux

package lcl

import (
	"github.com/rarnu/golcl/lcl/api"
	"github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type GdkWindow uintptr

type GtkWidget uintptr

type GtkFixed uintptr

type XID uintptr

// PGtkWidget
func HandleToPlatformHandle(h types.HWND) GtkWidget {

	return GtkWidget(h)
}

func (f *TForm) PlatformWindow() GdkWindow {
	r, _, _ := api.GdkWindow_FromForm.Call(f.instance)
	return GdkWindow(r)
}

func (g GdkWindow) XID() (xid XID) {
	api.GdkWindow_GetXId.Call(uintptr(g), uintptr(unsafe.Pointer(&xid)))
	return
}

// lz中首先是一个widget，然后上面用了一个fixedWidget来处理的。
func (g GtkWidget) FixedWidget() GtkFixed {
	r, _, _ := api.GtkWidget_GetGtkFixed.Call(uintptr(g))
	return GtkFixed(r)
}

func (g GtkWidget) Window() GdkWindow {
	return 0
	r, _, _ := api.GtkWidget_Window.Call(uintptr(g))
	return GdkWindow(r)
}
