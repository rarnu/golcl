//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

import "fmt"

type TModalResult = int32

// TColor 常用值请见 types/colors 包
type TColor uint32

type THelpEventData = uintptr

type TTabOrder = int16

type PFNLVCOMPARE uintptr

type PFNTVCOMPARE uintptr

type Integer = int32

type Cardinal = uint32

type Single = float32

type Word = uint16

type Byte = uint8

type TFontCharset = uint8

type TSpacingSize int32

// Char Unicode 主要用于keymap, 参见types/keys包
type Char = uint16

type TClass uintptr

type TThreadID = uintptr

type TClipboardFormat = uintptr

type TGridCoord struct {
	X int32
	Y int32
}

type TCustomData = uintptr

type TGridRect TRect

// IObjectArray 伪造，实际为一个接口类型
type IObjectArray uintptr

type TSysLocale struct {
	//Delphi compat fields
	DefaultLCID int32
	PriLangID   int32
	SubLangID   int32

	// win32 names
	FarEast    bool
	MiddleEast bool

	// LCL
	// real meaning  2: (MBCS: boolean; RightToLeft: Boolean);
}

type TSmallPoint struct {
	X int16
	Y int16
}

type TGUID struct {
	D1 uint32
	D2 uint16
	D3 uint16
	D4 [8]uint8
}

// TLibResource LibResource
type TLibResource struct {
	Name string
	Ptr  uintptr
}

type TResItem struct {
	Name  uintptr
	Value uintptr
}

// TConstraintSize = 0..MaxInt;
type TConstraintSize int32

type TAlignInfo struct {
	AlignList    uintptr
	ControlIndex int32
	Align        TAlign
	Scratch      int32
}

// TCreateParams Moved from Controls to avoid circles, Since it is part of the interface now
type TCreateParams struct {
	Caption       LPCWSTR
	Style         uint32
	ExStyle       uint32
	X, Y          int32
	Width, Height int32
	WndParent     HWND
	Param         uintptr
	WindowClass   TWndClass
	WinClassName  [64]int8
}

// R TColor
func (c TColor) R() byte {
	return byte(c)
}

func (c TColor) G() byte {
	return byte(c >> 8)
}

func (c TColor) B() byte {
	return byte(c >> 16)
}

func (c TColor) RGB(r, g, b byte) TColor {
	return TColor(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))
}

// TGUID

func (g TGUID) FromString(str string) (result TGUID) {
	_, _ = fmt.Sscanf(str, "{%8X-%4X-%4X-%2X%2X-%2X%2X%2X%2X%2X%2X}", &result.D1, &result.D2, &result.D3, &result.D4[0],
		&result.D4[1], &result.D4[2], &result.D4[3], &result.D4[4], &result.D4[5], &result.D4[6], &result.D4[7])
	return
}

func (g TGUID) ToString() string {
	return fmt.Sprintf("{%.8X-%.4X-%.4X-%.2X%.2X-%.2X%.2X%.2X%.2X%.2X%.2X}",
		g.D1, g.D2, g.D3, g.D4[0], g.D4[1], g.D4[2], g.D4[3], g.D4[4], g.D4[5], g.D4[6], g.D4[7])
}

func (g TGUID) Empty() TGUID {
	return TGUID{}
}

func (g TGUID) IsEqual(val TGUID) bool {
	return (g.D1 == val.D1) && (g.D2 == val.D2) && (g.D3 == val.D3) && (g.D4 == val.D4)
}

// Empty TSmallPoint
func (s TSmallPoint) Empty() TSmallPoint {
	return TSmallPoint{}
}

func (s TSmallPoint) IsEqual(val TSmallPoint) bool {
	return s.X == val.X && s.Y == val.Y
}
