//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ImageList_Draw1(obj uintptr, canvas uintptr, x, y, index int32, enabled bool) {
	_, _, _ = imageList_Draw1.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(index), GoBoolToDBool(enabled))
}

func ImageList_Draw2(obj uintptr, canvas uintptr, x, y, index int32, drawingStyle TDrawingStyle,
	imageType TImageType, enabled bool) {
	_, _, _ = imageList_Draw2.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), GoBoolToDBool(enabled))
}

func ImageList_DrawOverlay1(obj uintptr, canvas uintptr, x, y, imageIndex int32, overlay uint8, enabled bool) {
	_, _, _ = imageList_DrawOverlay1.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), GoBoolToDBool(enabled))
}

func ImageList_DrawOverlay2(obj, canvas uintptr, x, y, imageIndex int32, overlay uint8,
	drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
	_, _, _ = imageList_DrawOverlay2.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), GoBoolToDBool(enabled))
}

func ImageList_GetIcon1(obj uintptr, index int32, image uintptr) {
	_, _, _ = imageList_GetIcon1.Call(obj, uintptr(index), image)
}

func ImageList_GetIcon2(obj uintptr, index int32, image uintptr, drawingStyle TDrawingStyle, imageType TImageType) {
	_, _, _ = imageList_GetIcon2.Call(obj, uintptr(index), image, uintptr(drawingStyle), uintptr(imageType))
}
