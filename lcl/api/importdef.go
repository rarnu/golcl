//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

var (
	application_Instance   = liblcl.NewProc("Application_Instance")
	application_CreateForm = liblcl.NewProc("Application_CreateForm")
	application_Run        = liblcl.NewProc("Application_Run")
	application_Initialize = liblcl.NewProc("Application_Initialize")

	form_Create2      = liblcl.NewProc("Form_Create2")
	form_SetOnWndProc = liblcl.NewProc("Form_SetOnWndProc")
	form_SetGoPtr     = liblcl.NewProc("Form_SetGoPtr")

	setEventCallback                   = liblcl.NewProc("SetEventCallback")
	setMessageCallback                 = liblcl.NewProc("SetMessageCallback")
	setThreadSyncCallback              = liblcl.NewProc("SetThreadSyncCallback")
	setRequestCallCreateParamsCallback = liblcl.NewProc("SetRequestCallCreateParamsCallback")

	dGetStringArrOf = liblcl.NewProc("DGetStringArrOf")
	dStrLen         = liblcl.NewProc("DStrLen")
	dMove           = liblcl.NewProc("DMove")

	dShowMessage = liblcl.NewProc("DShowMessage")
	dMessageDlg  = liblcl.NewProc("DMessageDlg")

	mouse_Instance  = liblcl.NewProc("Mouse_Instance")
	screen_Instance = liblcl.NewProc("Screen_Instance")

	dSynchronize = liblcl.NewProc("DSynchronize")

	// TMenuItem
	dTextToShortCut = liblcl.NewProc("DTextToShortCut")
	dShortCutToText = liblcl.NewProc("DShortCutToText")

	// TClipboard
	clipboard_Instance         = liblcl.NewProc("Clipboard_Instance")
	clipboard_HasFormat        = liblcl.NewProc("Clipboard_HasFormat")
	dSetClipboard              = liblcl.NewProc("DSetClipboard")
	dRegisterClipboardFormat   = liblcl.NewProc("DRegisterClipboardFormat")
	clipboard_GetAsHtml        = liblcl.NewProc("Clipboard_GetAsHtml")
	clipboard_GetTextBuf       = liblcl.NewProc("Clipboard_GetTextBuf")
	clipboard_GetAsText        = liblcl.NewProc("Clipboard_GetAsText")
	clipboard_SetAsText        = liblcl.NewProc("Clipboard_SetAsText")
	dPredefinedClipboardFormat = liblcl.NewProc("DPredefinedClipboardFormat")

	// DSysOpen
	dSysOpen = liblcl.NewProc("DSysOpen")

	// TMemoryStream
	memoryStream_Read  = liblcl.NewProc("MemoryStream_Read")
	memoryStream_Write = liblcl.NewProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy     = liblcl.NewProc("Canvas_BrushCopy")
	canvas_CopyRect      = liblcl.NewProc("Canvas_CopyRect")
	canvas_Draw1         = liblcl.NewProc("Canvas_Draw1")
	canvas_Draw2         = liblcl.NewProc("Canvas_Draw2")
	canvas_DrawFocusRect = liblcl.NewProc("Canvas_DrawFocusRect")
	canvas_FillRect      = liblcl.NewProc("Canvas_FillRect")
	canvas_FrameRect     = liblcl.NewProc("Canvas_FrameRect")
	canvas_TextRect1     = liblcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2     = liblcl.NewProc("Canvas_TextRect2")
	canvas_Polygon       = liblcl.NewProc("Canvas_Polygon")
	canvas_Polyline      = liblcl.NewProc("Canvas_Polyline")
	canvas_PolyBezier    = liblcl.NewProc("Canvas_PolyBezier")

	// TImageList
	imageList_Draw1        = liblcl.NewProc("ImageList_Draw1")
	imageList_Draw2        = liblcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = liblcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = liblcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = liblcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     = liblcl.NewProc("ImageList_GetIcon2")

	dExtractFilePath = liblcl.NewProc("DExtractFilePath")
	dFileExists      = liblcl.NewProc("DFileExists")

	dSelectDirectory1 = liblcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 = liblcl.NewProc("DSelectDirectory2")
	dInputBox         = liblcl.NewProc("DInputBox")
	dInputQuery       = liblcl.NewProc("DInputQuery")
	dPasswordBox      = liblcl.NewProc("DPasswordBox")
	dInputCombo       = liblcl.NewProc("DInputCombo")
	dInputComboEx     = liblcl.NewProc("DInputComboEx")

	// TSysLocaled
	dSysLocale = liblcl.NewProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut = liblcl.NewProc("DCreateURLShortCut")
	dCreateShortCut    = liblcl.NewProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue    = liblcl.NewProc("DSetPropertyValue")
	dSetPropertySecValue = liblcl.NewProc("DSetPropertySecValue")

	// Printer
	printer_Instance   = liblcl.NewProc("Printer_Instance")
	printer_SetPrinter = liblcl.NewProc("Printer_SetPrinter")

	// guid
	dGUIDToString = liblcl.NewProc("DGUIDToString")
	dStringToGUID = liblcl.NewProc("DStringToGUID")
	dCreateGUID   = liblcl.NewProc("DCreateGUID")

	// libResources
	dGetLibResourceCount = liblcl.NewProc("DGetLibResourceCount")
	dGetLibResourceItem  = liblcl.NewProc("DGetLibResourceItem")
	dModifyLibResource   = liblcl.NewProc("DModifyLibResource")
	dLibAbout            = liblcl.NewProc("DLibAbout")

	// 库的信息
	dLibStringEncoding = liblcl.NewProc("DLibStringEncoding")
	dLibVersion        = liblcl.NewProc("DLibVersion")

	dMainThreadId    = liblcl.NewProc("DMainThreadId")
	dCurrentThreadId = liblcl.NewProc("DCurrentThreadId")

	dInitGoDll = liblcl.NewProc("DInitGoDll")

	dFindControl           = liblcl.NewProc("DFindControl")
	dFindLCLControl        = liblcl.NewProc("DFindLCLControl")
	dFindOwnerControl      = liblcl.NewProc("DFindOwnerControl")
	dFindControlAtPosition = liblcl.NewProc("DFindControlAtPosition")
	dFindLCLWindow         = liblcl.NewProc("DFindLCLWindow")
	dFindDragTarget        = liblcl.NewProc("DFindDragTarget")
)
