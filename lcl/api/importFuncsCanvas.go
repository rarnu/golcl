package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TCanvas ---------------------------

func Canvas_Create() uintptr {
	ret, _, _ := getLazyProc("Canvas_Create").Call()
	return ret
}

func Canvas_Free(obj uintptr) {
	_, _, _ = getLazyProc("Canvas_Free").Call(obj)
}

func Canvas_Arc(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
	_, _, _ = getLazyProc("Canvas_Arc").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(X3), uintptr(Y3), uintptr(X4), uintptr(Y4))
}

func Canvas_ArcTo(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
	_, _, _ = getLazyProc("Canvas_ArcTo").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(X3), uintptr(Y3), uintptr(X4), uintptr(Y4))
}

func Canvas_AngleArc(obj uintptr, X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32) {
	_, _, _ = getLazyProc("Canvas_AngleArc").Call(obj, uintptr(X), uintptr(Y), uintptr(Radius), uintptr(unsafe.Pointer(&StartAngle)), uintptr(unsafe.Pointer(&SweepAngle)))
}

func Canvas_Chord(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
	_, _, _ = getLazyProc("Canvas_Chord").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(X3), uintptr(Y3), uintptr(X4), uintptr(Y4))
}

func Canvas_Ellipse(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32) {
	_, _, _ = getLazyProc("Canvas_Ellipse").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func Canvas_FloodFill(obj uintptr, X int32, Y int32, Color TColor, FillStyle TFillStyle) {
	_, _, _ = getLazyProc("Canvas_FloodFill").Call(obj, uintptr(X), uintptr(Y), uintptr(Color), uintptr(FillStyle))
}

func Canvas_HandleAllocated(obj uintptr) bool {
	ret, _, _ := getLazyProc("Canvas_HandleAllocated").Call(obj)
	return DBoolToGoBool(ret)
}

func Canvas_LineTo(obj uintptr, X int32, Y int32) {
	_, _, _ = getLazyProc("Canvas_LineTo").Call(obj, uintptr(X), uintptr(Y))
}

func Canvas_MoveTo(obj uintptr, X int32, Y int32) {
	_, _, _ = getLazyProc("Canvas_MoveTo").Call(obj, uintptr(X), uintptr(Y))
}

func Canvas_Pie(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
	_, _, _ = getLazyProc("Canvas_Pie").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(X3), uintptr(Y3), uintptr(X4), uintptr(Y4))
}

func Canvas_Rectangle(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32) {
	_, _, _ = getLazyProc("Canvas_Rectangle").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func Canvas_Refresh(obj uintptr) {
	_, _, _ = getLazyProc("Canvas_Refresh").Call(obj)
}

func Canvas_RoundRect(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32) {
	_, _, _ = getLazyProc("Canvas_RoundRect").Call(obj, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(X3), uintptr(Y3))
}

func Canvas_StretchDraw(obj uintptr, Rect TRect, Graphic uintptr) {
	_, _, _ = getLazyProc("Canvas_StretchDraw").Call(obj, uintptr(unsafe.Pointer(&Rect)), Graphic)
}

func Canvas_TextExtent(obj uintptr, Text string) TSize {
	var ret TSize
	_, _, _ = getLazyProc("Canvas_TextExtent").Call(obj, GoStrToDStr(Text), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func Canvas_TextOut(obj uintptr, X int32, Y int32, Text string) {
	_, _, _ = getLazyProc("Canvas_TextOut").Call(obj, uintptr(X), uintptr(Y), GoStrToDStr(Text))
}

func Canvas_Lock(obj uintptr) {
	_, _, _ = getLazyProc("Canvas_Lock").Call(obj)
}

func Canvas_TextHeight(obj uintptr, Text string) int32 {
	ret, _, _ := getLazyProc("Canvas_TextHeight").Call(obj, GoStrToDStr(Text))
	return int32(ret)
}

func Canvas_TextWidth(obj uintptr, Text string) int32 {
	ret, _, _ := getLazyProc("Canvas_TextWidth").Call(obj, GoStrToDStr(Text))
	return int32(ret)
}

func Canvas_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("Canvas_Assign").Call(obj, Source)
}

func Canvas_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("Canvas_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func Canvas_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("Canvas_ClassType").Call(obj)
	return TClass(ret)
}

func Canvas_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("Canvas_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func Canvas_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Canvas_InstanceSize").Call(obj)
	return int32(ret)
}

func Canvas_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("Canvas_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func Canvas_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("Canvas_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func Canvas_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Canvas_GetHashCode").Call(obj)
	return int32(ret)
}

func Canvas_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("Canvas_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func Canvas_GetHandle(obj uintptr) HDC {
	ret, _, _ := getLazyProc("Canvas_GetHandle").Call(obj)
	return ret
}

func Canvas_SetHandle(obj uintptr, value HDC) {
	_, _, _ = getLazyProc("Canvas_SetHandle").Call(obj, value)
}

func Canvas_GetBrush(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Canvas_GetBrush").Call(obj)
	return ret
}

func Canvas_SetBrush(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Canvas_SetBrush").Call(obj, value)
}

func Canvas_GetCopyMode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("Canvas_GetCopyMode").Call(obj)
	return int32(ret)
}

func Canvas_SetCopyMode(obj uintptr, value int32) {
	_, _, _ = getLazyProc("Canvas_SetCopyMode").Call(obj, uintptr(value))
}

func Canvas_GetFont(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Canvas_GetFont").Call(obj)
	return ret
}

func Canvas_SetFont(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Canvas_SetFont").Call(obj, value)
}

func Canvas_GetPen(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("Canvas_GetPen").Call(obj)
	return ret
}

func Canvas_SetPen(obj uintptr, value uintptr) {
	_, _, _ = getLazyProc("Canvas_SetPen").Call(obj, value)
}

func Canvas_SetOnChange(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Canvas_SetOnChange").Call(obj, addEventToMap(obj, fn))
}

func Canvas_SetOnChanging(obj uintptr, fn any) {
	_, _, _ = getLazyProc("Canvas_SetOnChanging").Call(obj, addEventToMap(obj, fn))
}

func Canvas_GetPixels(obj uintptr, X int32, Y int32) TColor {
	ret, _, _ := getLazyProc("Canvas_GetPixels").Call(obj, uintptr(X), uintptr(Y))
	return TColor(ret)
}

func Canvas_SetPixels(obj uintptr, X int32, Y int32, value TColor) {
	_, _, _ = getLazyProc("Canvas_SetPixels").Call(obj, uintptr(X), uintptr(Y), uintptr(value))
}

func Canvas_StaticClassType() TClass {
	r, _, _ := getLazyProc("Canvas_StaticClassType").Call()
	return TClass(r)
}
