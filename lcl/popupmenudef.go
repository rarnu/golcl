//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (p *TPopupMenu) Popup2() {
	point := Mouse.CursorPos()
	p.Popup(point.X, point.Y)
}
