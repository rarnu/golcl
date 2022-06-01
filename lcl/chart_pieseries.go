package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TPieSeries struct {
	IBasicChartSeries
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

func AsPieSeries(obj any) *TPieSeries {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPieSeries{instance: instance, ptr: ptr}
}

func NewPieSeries(owner IComponent) *TPieSeries {
	c := new(TPieSeries)
	c.instance = PieSeries_Create(CheckPtr(owner))
	c.ptr = unsafe.Pointer(c.instance)
	return c
}

func PieSeriesFromInst(inst uintptr) *TPieSeries {
	return AsPieSeries(inst)
}

func PieSeriesFromObj(obj IObject) *TPieSeries {
	return AsPieSeries(obj)
}

func PieSeriesFromUnsafePointer(ptr unsafe.Pointer) *TPieSeries {
	return AsPieSeries(ptr)
}

func (w *TPieSeries) Free() {
	if w.instance != 0 {
		PieSeries_Free(w.instance)
		w.instance, w.ptr = 0, nullptr
	}
}

func (w *TPieSeries) Instance() uintptr {
	return w.instance
}

func (w *TPieSeries) UnsafeAddr() unsafe.Pointer {
	return w.ptr
}

func (w *TPieSeries) IsValid() bool {
	return w.instance != 0
}

func (w *TPieSeries) Is() TIs {
	return TIs(w.instance)
}

func TPieSeriesClass() TClass {
	return PieSeries_StaticClassType()
}

func (w *TPieSeries) EdgePen() *TPen {
	return AsPen(PieSeries_GetEdgePen(w.instance))
}

func (w *TPieSeries) SetEdgePen(value *TPen) {
	PieSeries_SetEdgePen(w.instance, CheckPtr(value))
}

func (w *TPieSeries) Exploded() bool {
	return PieSeries_GetExploded(w.instance)
}

func (w *TPieSeries) SetExploded(value bool) {
	PieSeries_SetExploded(w.instance, value)
}

func (w *TPieSeries) FixedRadius() TChartDistance {
	return PieSeries_GetFixedRadius(w.instance)
}

func (w *TPieSeries) SetFixedRadius(value TChartDistance) {
	PieSeries_SetFixedRadius(w.instance, value)
}

func (w *TPieSeries) MarkDistancePercent() bool {
	return PieSeries_GetMarkDistancePercent(w.instance)
}

func (w *TPieSeries) SetMarkDistancePercent(value bool) {
	PieSeries_SetMarkDistancePercent(w.instance, value)
}

func (w *TPieSeries) MarkPositions() TPieMarkPositions {
	return PieSeries_GetMarkPositions(w.instance)
}

func (w *TPieSeries) SetMarkPositions(value TPieMarkPositions) {
	PieSeries_SetMarkPositions(w.instance, value)
}

func (w *TPieSeries) RotateLabels() bool {
	return PieSeries_GetRotateLabels(w.instance)
}

func (w *TPieSeries) SetRotateLabels(value bool) {
	PieSeries_SetRotateLabels(w.instance, value)
}

func (w *TPieSeries) Marks() *TChartMarks {
	return AsChartMarks(PieSeries_GetMarks(w.instance))
}

func (w *TPieSeries) SetMarks(value *TChartMarks) {
	PieSeries_SetMarks(w.instance, CheckPtr(value))
}

func (w *TPieSeries) Title() string {
	return PieSeries_GetTitle(w.instance)
}

func (w *TPieSeries) SetTitle(value string) {
	PieSeries_SetTitle(w.instance, value)
}

func (w *TPieSeries) Legend() *TChartSeriesLegend {
	return AsChartSeriesLegend(PieSeries_GetLegend(w.instance))
}

func (w *TPieSeries) SetLegend(value *TChartSeriesLegend) {
	PieSeries_SetLegend(w.instance, CheckPtr(value))
}

func (w *TPieSeries) Active() bool {
	return PieSeries_GetActive(w.instance)
}

func (w *TPieSeries) SetActive(value bool) {
	PieSeries_SetActive(w.instance, value)
}

func (w *TPieSeries) Depth() TChartDistance {
	return PieSeries_GetDepth(w.instance)
}

func (w *TPieSeries) SetDepth(value TChartDistance) {
	PieSeries_SetDepth(w.instance, value)
}

func (w *TPieSeries) DragOrigin() TPoint {
	return PieSeries_GetDragOrigin(w.instance)
}

func (w *TPieSeries) SetDragOrigin(value TPoint) {
	PieSeries_SetDragOrigin(w.instance, value)
}

func (w *TPieSeries) ParentChart() *TChart {
	return AsChart(PieSeries_GetParentChart(w.instance))
}

func (w *TPieSeries) Shadow() *TChartShadow {
	return AsChartShadow(PieSeries_GetShadow(w.instance))
}

func (w *TPieSeries) SetShadow(value *TChartShadow) {
	PieSeries_SetShadow(w.instance, CheckPtr(value))
}

func (w *TPieSeries) SpecialPointPos() bool {
	return PieSeries_GetSpecialPointPos(w.instance)
}

func (w *TPieSeries) Transparency() TChartTransparency {
	return PieSeries_GetTransparency(w.instance)
}

func (w *TPieSeries) SetTransparency(value TChartTransparency) {
	PieSeries_SetTransparency(w.instance, value)
}

func (w *TPieSeries) ZPosition() TChartDistance {
	return PieSeries_GetZPosition(w.instance)
}

func (w *TPieSeries) SetZPosition(value TChartDistance) {
	PieSeries_SetZPosition(w.instance, value)
}

func (w *TPieSeries) AddPie(value float64, text string, color TColor) int32 {
	return PieSeries_AddPie(w.instance, value, text, color)
}

func (w *TPieSeries) Assign(source IObject) {
	PieSeries_Assign(w.instance, CheckPtr(source))
}

func (w *TPieSeries) CalcInnerRadius() int32 {
	return PieSeries_CalcInnerRadius(w.instance)
}

func (w *TPieSeries) Draw(drawer IChartDrawer) {
	PieSeries_Draw(w.instance, CheckPtr(drawer))
}

func (w *TPieSeries) FindContainingSlice(point TPoint) int32 {
	return PieSeries_FindContainingSlice(w.instance, point)
}

func (w *TPieSeries) SliceColor(index int32) TColor {
	return PieSeries_SliceColor(w.instance, index)
}

func (w *TPieSeries) GetColor(index int32) TColor {
	return PieSeries_GetColor(w.instance, index)
}

func (w *TPieSeries) GetXImgValue(index int32) int32 {
	return PieSeries_GetXImgValue(w.instance, index)
}

func (w *TPieSeries) GetXMax() float64 {
	return PieSeries_GetXMax(w.instance)
}

func (w *TPieSeries) GetXMin() float64 {
	return PieSeries_GetXMin(w.instance)
}

func (w *TPieSeries) GetXValue(index int32) float64 {
	return PieSeries_GetXValue(w.instance, index)
}

func (w *TPieSeries) GetXValues(index int32, xindex int32) float64 {
	return PieSeries_GetXValues(w.instance, index, xindex)
}

func (w *TPieSeries) GetYImgValue(index int32) int32 {
	return PieSeries_GetYImgValue(w.instance, index)
}

func (w *TPieSeries) GetYMax() float64 {
	return PieSeries_GetYMax(w.instance)
}

func (w *TPieSeries) GetYMin() float64 {
	return PieSeries_GetYMin(w.instance)
}
func (w *TPieSeries) GetYValue(index int32) float64 {
	return PieSeries_GetYValue(w.instance, index)
}

func (w *TPieSeries) GetYValues(index int32, yindex int32) float64 {
	return PieSeries_GetYValues(w.instance, index, yindex)
}

func (w *TPieSeries) SetColor(index int32, color TColor) {
	PieSeries_SetColor(w.instance, index, color)
}

func (w *TPieSeries) SetText(index int32, value string) {
	PieSeries_SetText(w.instance, index, value)
}

func (w *TPieSeries) SetXValue(index int32, value float64) {
	PieSeries_SetXValue(w.instance, index, value)
}

func (w *TPieSeries) SetXValues(index int32, xindex int32, value float64) {
	PieSeries_SetXValues(w.instance, index, xindex, value)
}

func (w *TPieSeries) SetYValue(index int32, value float64) {
	PieSeries_SetYValue(w.instance, index, value)
}

func (w *TPieSeries) SetYValues(index int32, yindex int32, value float64) {
	PieSeries_SetYValues(w.instance, index, yindex, value)
}

func (w *TPieSeries) Add(value float64, xlabel string, color TColor) int32 {
	return PieSeries_Add(w.instance, value, xlabel, color)
}

func (w *TPieSeries) AddNull(label string, color TColor) int32 {
	return PieSeries_AddNull(w.instance, label, color)
}

func (w *TPieSeries) AddX(x float64, label string, color TColor) int32 {
	return PieSeries_AddX(w.instance, x, label, color)
}

func (w *TPieSeries) AddXY(x float64, y float64, xlabel string, color TColor) int32 {
	return PieSeries_AddXY(w.instance, x, y, xlabel, color)
}

func (w *TPieSeries) AddY(y float64, label string, color TColor) int32 {
	return PieSeries_AddY(w.instance, y, label, color)
}

func (w *TPieSeries) BeginUpdate() {
	PieSeries_BeginUpdate(w.instance)
}

func (w *TPieSeries) Clear() {
	PieSeries_Clear(w.instance)
}

func (w *TPieSeries) Count() int32 {
	return PieSeries_Count(w.instance)
}

func (w *TPieSeries) Delete(index int32) {
	PieSeries_Delete(w.instance, index)
}

func (w *TPieSeries) EndUpdate() {
	PieSeries_EndUpdate(w.instance)
}

func (w *TPieSeries) FindYRange(xmin float64, xmax float64, ymin *float64, ymax *float64) {
	PieSeries_FindYRange(w.instance, xmin, xmax, ymin, ymax)
}

func (w *TPieSeries) FormattedMark(index int32, format string, yindex int32) string {
	return PieSeries_FormattedMark(w.instance, index, format, yindex)
}

func (w *TPieSeries) IsEmpty() bool {
	return PieSeries_IsEmpty(w.instance)
}

func (w *TPieSeries) AxisToGraphX(x float64) float64 {
	return PieSeries_AxisToGraphX(w.instance, x)
}

func (w *TPieSeries) AxisToGraphY(y float64) float64 {
	return PieSeries_AxisToGraphY(w.instance, y)
}

func (w *TPieSeries) GetAxisX() *TChartAxis {
	return AsChartAxis(PieSeries_GetAxisX(w.instance))
}

func (w *TPieSeries) GetAxisY() *TChartAxis {
	return AsChartAxis(PieSeries_GetAxisY(w.instance))
}

func (w *TPieSeries) GetAxisBounds(axis *TChartAxis, min *float64, max *float64) bool {
	return PieSeries_GetAxisBounds(w.instance, CheckPtr(axis), min, max)
}

func (w *TPieSeries) GetDepthColor(color int32, opposite bool) int32 {
	return PieSeries_GetDepthColor(w.instance, color, opposite)
}

func (w *TPieSeries) GraphToAxisX(x float64) float64 {
	return PieSeries_GraphToAxisX(w.instance, x)
}

func (w *TPieSeries) GraphToAxisY(y float64) float64 {
	return PieSeries_GraphToAxisY(w.instance, y)
}

func (w *TPieSeries) IsRotated() bool {
	return PieSeries_IsRotated(w.instance)
}

func (w *TPieSeries) MovePoint(index *int32, newPos TPoint) {
	PieSeries_MovePoint(w.instance, index, newPos)
}

func (w *TPieSeries) UpdateBiDiMode() {
	PieSeries_UpdateBiDiMode(w.instance)
}
