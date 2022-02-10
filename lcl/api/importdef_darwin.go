//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin

package api

var (
	NSWindow_FromForm                      = liblcl.NewProc("NSWindow_FromForm")
	NSWindow_titleVisibility               = liblcl.NewProc("NSWindow_titleVisibility")
	NSWindow_setTitleVisibility            = liblcl.NewProc("NSWindow_setTitleVisibility")
	NSWindow_titlebarAppearsTransparent    = liblcl.NewProc("NSWindow_titlebarAppearsTransparent")
	NSWindow_setTitlebarAppearsTransparent = liblcl.NewProc("NSWindow_setTitlebarAppearsTransparent")
	NSWindow_styleMask                     = liblcl.NewProc("NSWindow_styleMask")
	NSWindow_setStyleMask                  = liblcl.NewProc("NSWindow_setStyleMask")
	NSWindow_setRepresentedURL             = liblcl.NewProc("NSWindow_setRepresentedURL")
	//NSWindow_release                       = libvcl.NewProc("NSWindow_release")
)
