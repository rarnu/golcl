package lcl

import (
	. "github.com/rarnu/golcl/lcl/types"
)

type IBasicChartSeries interface {
	IComponent
	AxisToGraphX(x float64) float64
	AxisToGraphY(y float64) float64
	GraphToAxisX(x float64) float64
	GraphToAxisY(y float64) float64
	Draw(drawer IChartDrawer)
	GetAxisBounds(axis *TChartAxis, min *float64, max *float64) bool
	IsEmpty() bool
	UpdateBiDiMode()
	Active() bool
	SetActive(value bool)
	Depth() TChartDistance
	SetDepth(value TChartDistance)
	DragOrigin() TPoint
	SetDragOrigin(value TPoint)
	ParentChart() *TChart
	Shadow() *TChartShadow
	SetShadow(value *TChartShadow)
	SpecialPointPos() bool
	SetSpecialPointPos(value bool)
	Transparency() TChartTransparency
	SetTransparency(value TChartTransparency)
	ZPosition() TChartDistance
	SetZPosition(value TChartDistance)
}
