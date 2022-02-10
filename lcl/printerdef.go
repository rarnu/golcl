//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
)

func (p *TPrinter) SetPrinter(aName string) {
	Printer_SetPrinter(p.instance, aName)
}
