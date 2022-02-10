//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import "github.com/rarnu/golcl/lcl/api"

// Shortcut

// 创建一个url的快捷方式
func CreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	api.DCreateURLShortCut(aDestPath, aShortCutName, aURL)
}

// 创建一个快捷方式
func CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return api.DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs)
}
