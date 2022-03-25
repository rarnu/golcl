package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TImageList ---------------------------

func ImageList_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageList_Create").Call(obj)
	return ret
}

func ImageList_Free(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_Free").Call(obj)
}

func ImageList_StretchDraw(obj uintptr, ACanvas uintptr, AIndex int32, ARect TRect, AEnabled bool) {
	_, _, _ = getLazyProc("ImageList_StretchDraw").Call(obj, ACanvas, uintptr(AIndex), uintptr(unsafe.Pointer(&ARect)), GoBoolToDBool(AEnabled))
}

func ImageList_AddSliced(obj uintptr, Image uintptr, AHorizontalCount int32, AVerticalCount int32) int32 {
	ret, _, _ := getLazyProc("ImageList_AddSliced").Call(obj, Image, uintptr(AHorizontalCount), uintptr(AVerticalCount))
	return int32(ret)
}

func ImageList_GetHotSpot(obj uintptr) TPoint {
	var ret TPoint
	_, _, _ = getLazyProc("ImageList_GetHotSpot").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func ImageList_HideDragImage(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_HideDragImage").Call(obj)
}

func ImageList_ShowDragImage(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_ShowDragImage").Call(obj)
}

func ImageList_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("ImageList_Assign").Call(obj, Source)
}

func ImageList_Add(obj uintptr, Image uintptr, Mask uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_Add").Call(obj, Image, Mask)
	return int32(ret)
}

func ImageList_AddIcon(obj uintptr, Image uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_AddIcon").Call(obj, Image)
	return int32(ret)
}

func ImageList_AddImages(obj uintptr, Value uintptr) {
	_, _, _ = getLazyProc("ImageList_AddImages").Call(obj, Value)
}

func ImageList_AddMasked(obj uintptr, Image uintptr, MaskColor TColor) int32 {
	ret, _, _ := getLazyProc("ImageList_AddMasked").Call(obj, Image, uintptr(MaskColor))
	return int32(ret)
}

func ImageList_Clear(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_Clear").Call(obj)
}

func ImageList_Delete(obj uintptr, Index int32) {
	_, _, _ = getLazyProc("ImageList_Delete").Call(obj, uintptr(Index))
}

func ImageList_Insert(obj uintptr, Index int32, Image uintptr, Mask uintptr) {
	_, _, _ = getLazyProc("ImageList_Insert").Call(obj, uintptr(Index), Image, Mask)
}

func ImageList_InsertIcon(obj uintptr, Index int32, Image uintptr) {
	_, _, _ = getLazyProc("ImageList_InsertIcon").Call(obj, uintptr(Index), Image)
}

func ImageList_InsertMasked(obj uintptr, Index int32, Image uintptr, MaskColor TColor) {
	_, _, _ = getLazyProc("ImageList_InsertMasked").Call(obj, uintptr(Index), Image, uintptr(MaskColor))
}

func ImageList_Move(obj uintptr, CurIndex int32, NewIndex int32) {
	_, _, _ = getLazyProc("ImageList_Move").Call(obj, uintptr(CurIndex), uintptr(NewIndex))
}

func ImageList_Replace(obj uintptr, Index int32, Image uintptr, Mask uintptr) {
	_, _, _ = getLazyProc("ImageList_Replace").Call(obj, uintptr(Index), Image, Mask)
}

func ImageList_ReplaceMasked(obj uintptr, Index int32, NewImage uintptr, MaskColor TColor) {
	_, _, _ = getLazyProc("ImageList_ReplaceMasked").Call(obj, uintptr(Index), NewImage, uintptr(MaskColor))
}

func ImageList_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_BeginUpdate").Call(obj)
}

func ImageList_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("ImageList_EndUpdate").Call(obj)
}

func ImageList_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := getLazyProc("ImageList_FindComponent").Call(obj, GoStrToDStr(AName))
	return ret
}

func ImageList_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageList_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func ImageList_HasParent(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_HasParent").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageList_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("ImageList_ClassType").Call(obj)
	return TClass(ret)
}

func ImageList_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageList_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func ImageList_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_InstanceSize").Call(obj)
	return int32(ret)
}

func ImageList_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("ImageList_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func ImageList_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func ImageList_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetHashCode").Call(obj)
	return int32(ret)
}

func ImageList_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageList_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func ImageList_GetScaled(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_GetScaled").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageList_SetScaled(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageList_SetScaled").Call(obj, GoBoolToDBool(value))
}

func ImageList_GetShareImages(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_GetShareImages").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageList_SetShareImages(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageList_SetShareImages").Call(obj, GoBoolToDBool(value))
}

func ImageList_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetCount").Call(obj)
	return int32(ret)
}

func ImageList_GetBlendColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ImageList_GetBlendColor").Call(obj)
	return TColor(ret)
}

func ImageList_SetBlendColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ImageList_SetBlendColor").Call(obj, uintptr(value))
}

func ImageList_GetBkColor(obj uintptr) TColor {
	ret, _, _ := getLazyProc("ImageList_GetBkColor").Call(obj)
	return TColor(ret)
}

func ImageList_SetBkColor(obj uintptr, value TColor) {
	_, _, _ = getLazyProc("ImageList_SetBkColor").Call(obj, uintptr(value))
}

func ImageList_GetAllocBy(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetAllocBy").Call(obj)
	return int32(ret)
}

func ImageList_SetAllocBy(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageList_SetAllocBy").Call(obj, uintptr(value))
}

func ImageList_GetDrawingStyle(obj uintptr) TDrawingStyle {
	ret, _, _ := getLazyProc("ImageList_GetDrawingStyle").Call(obj)
	return TDrawingStyle(ret)
}

func ImageList_SetDrawingStyle(obj uintptr, value TDrawingStyle) {
	_, _, _ = getLazyProc("ImageList_SetDrawingStyle").Call(obj, uintptr(value))
}

func ImageList_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetHeight").Call(obj)
	return int32(ret)
}

func ImageList_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageList_SetHeight").Call(obj, uintptr(value))
}

func ImageList_GetImageType(obj uintptr) TImageType {
	ret, _, _ := getLazyProc("ImageList_GetImageType").Call(obj)
	return TImageType(ret)
}

func ImageList_SetImageType(obj uintptr, value TImageType) {
	_, _, _ = getLazyProc("ImageList_SetImageType").Call(obj, uintptr(value))
}

func ImageList_GetMasked(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_GetMasked").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageList_SetMasked(obj uintptr, value bool) {
	_, _, _ = getLazyProc("ImageList_SetMasked").Call(obj, GoBoolToDBool(value))
}

func ImageList_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("ImageList_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func ImageList_GetWidth(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetWidth").Call(obj)
	return int32(ret)
}

func ImageList_SetWidth(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageList_SetWidth").Call(obj, uintptr(value))
}

func ImageList_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := getLazyProc("ImageList_GetDragCursor").Call(obj)
	return TCursor(ret)
}

func ImageList_SetDragCursor(obj uintptr, value TCursor) {
	_, _, _ = getLazyProc("ImageList_SetDragCursor").Call(obj, uintptr(value))
}

func ImageList_GetDragging(obj uintptr) bool {
	ret, _, _ := getLazyProc("ImageList_GetDragging").Call(obj)
	return DBoolToGoBool(ret)
}

func ImageList_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetComponentCount").Call(obj)
	return int32(ret)
}

func ImageList_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("ImageList_GetComponentIndex").Call(obj)
	return int32(ret)
}

func ImageList_SetComponentIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("ImageList_SetComponentIndex").Call(obj, uintptr(value))
}

func ImageList_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ImageList_GetOwner").Call(obj)
	return ret
}

func ImageList_GetName(obj uintptr) string {
	ret, _, _ := getLazyProc("ImageList_GetName").Call(obj)
	return DStrToGoStr(ret)
}

func ImageList_SetName(obj uintptr, value string) {
	_, _, _ = getLazyProc("ImageList_SetName").Call(obj, GoStrToDStr(value))
}

func ImageList_GetTag(obj uintptr) int {
	ret, _, _ := getLazyProc("ImageList_GetTag").Call(obj)
	return int(ret)
}

func ImageList_SetTag(obj uintptr, value int) {
	_, _, _ = getLazyProc("ImageList_SetTag").Call(obj, uintptr(value))
}

func ImageList_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := getLazyProc("ImageList_GetComponents").Call(obj, uintptr(AIndex))
	return ret
}

func ImageList_StaticClassType() TClass {
	r, _, _ := getLazyProc("ImageList_StaticClassType").Call()
	return TClass(r)
}
