//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	. "github.com/rarnu/golcl/lcl/types"
)

func Canvas_BrushCopy(obj uintptr, dest TRect, bitmap uintptr, source TRect, color TColor) {
	_, _, _ = canvas_BrushCopy.Call(obj, uintptr(unsafe.Pointer(&dest)), bitmap, uintptr(unsafe.Pointer(&source)), uintptr(color))
}

func Canvas_CopyRect(obj uintptr, dest TRect, canvas uintptr, source TRect) {
	_, _, _ = canvas_CopyRect.Call(obj, uintptr(unsafe.Pointer(&dest)), canvas, uintptr(unsafe.Pointer(&source)))
}

func Canvas_Draw1(obj uintptr, x, y int32, graphic uintptr) {
	_, _, _ = canvas_Draw1.Call(obj, uintptr(x), uintptr(y), graphic)
}

func Canvas_Draw2(obj uintptr, x, y int32, graphic uintptr, opacity uint8) {
	_, _, _ = canvas_Draw2.Call(obj, uintptr(x), uintptr(y), graphic, uintptr(opacity))
}

func Canvas_DrawFocusRect(obj uintptr, aRect TRect) {
	_, _, _ = canvas_DrawFocusRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FillRect(obj uintptr, aRect TRect) {
	_, _, _ = canvas_FillRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FrameRect(obj uintptr, aRect TRect) {
	_, _, _ = canvas_FrameRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_TextRect2(obj uintptr, aRect *TRect, text string, textFormat TTextFormat) {
	_, _, _ = canvas_TextRect2.Call(obj, uintptr(unsafe.Pointer(aRect)), GoStrToDStr(text), uintptr(textFormat))
}

func Canvas_TextRect3(obj uintptr, aRect *TRect, text string, textFormat TTextFormat) {
	_, _, _ = canvas_TextRect2.Call(obj, uintptr(unsafe.Pointer(aRect)), GoStrToDStr(text), 0, uintptr(textFormat))
}

func Canvas_TextRect1(obj uintptr, aRect TRect, x, y int32, text string) {
	_, _, _ = canvas_TextRect1.Call(obj, uintptr(unsafe.Pointer(&aRect)), uintptr(x), uintptr(y), GoStrToDStr(text))
}

func Canvas_Polygon(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	_, _, _ = canvas_Polygon.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_Polyline(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	_, _, _ = canvas_Polyline.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_PolyBezier(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	_, _, _ = canvas_PolyBezier.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}
