//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows

package lcl

import (
	"fmt"
)

func showError(err any) {
	fmt.Println(err)
}

func tryLoadAppIcon() {
	// no code
}
