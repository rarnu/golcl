//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//go:build linux

package api

var (
	GdkWindow_GetXId      = liblcl.NewProc("GdkWindow_GetXId")
	GdkWindow_FromForm    = liblcl.NewProc("GdkWindow_FromForm")
	GtkWidget_GetGtkFixed = liblcl.NewProc("GtkWidget_GetGtkFixed")
	GtkWidget_Window      = liblcl.NewProc("GtkWidget_Window")
)
