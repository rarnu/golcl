package lcl

import (
	. "github.com/rarnu/golcl/lcl/types"
)

type IChartDrawer interface {
	IObject
	AddToFontOrientation(ADelta int32)
	ClippingStart(AClipRect TRect)
	ClippingStop()
	DrawingBegin(ABoundingBox TRect)
	DrawingEnd()
	DrawLineDepth(AX1, AY1, AX2, AY2, ADepth int32)
	Ellipse(AX1, AY1, AX2, AY2 int32)
	FillRect(AX1, AY1, AX2, AY2 int32)
	GetBrushColor() TChartColor
	GetFontAngle() float64 // in radians
	GetFontColor() TFPColor
	GetFontName() string
	GetFontSize() int32
	GetFontStyle() TChartFontStyles
	GetPenColor() TChartColor
	Line(AX1, AY1, AX2, AY2 int32)
	LineTo(AX, AY int32)
	MoveTo(AX, AY int32)
	Polygon(APoints []TPoint, AStartIndex, ANumPts int32)
	Polyline(APoints []TPoint, AStartIndex, ANumPts int32)
	PrepareSimplePen(AColor TChartColor)
	PutImage(AX, AY int32, AImage TFPCustomImage)
	PutPixel(AX, AY int32, AColor TChartColor)
	RadialPie(AX1, AY1, AX2, AY2 int32, AStartAngle16Deg, AAngleLength16Deg int32)
	Rectangle(AX1, AY1, AX2, AY2 int32)
	ResetFont()
	Scale(ADistance int32) int32
	SetAntialiasingMode(AValue TChartAntialiasingMode)
	SetBrush(ABrush TFPCustomBrush)
	SetBrushColor(AColor TChartColor)
	SetBrushParams(AStyle TFPBrushStyle, AColor TChartColor)
	SetFont(AValue TFPCustomFont)
	SetMonochromeColor(AColor TChartColor)
	SetPen(APen TFPCustomPen)
	SetPenColor(AColor TChartColor)
	SetPenParams(AStyle TFPPenStyle, AColor TChartColor, AWidth int32)
	SetPenWidth(AWidth int32)
	GetRightToLeft() bool
	SetRightToLeft(AValue bool)
	SetTransparency(ATransparency TChartTransparency)
	SetXor(AXor bool)
	TextExtent(AText string, ATextFormat TChartTextFormat) TPoint
	TextOut() TChartTextOut
	Brush() TFPCustomBrush
	BrushColor() TChartColor
}
