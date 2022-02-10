//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows

package types

// TMessage 消息值参见 types/messages包
type TMessage struct {
	Msg    Cardinal
	WParam WPARAM
	LParam LPARAM
	Result LRESULT
}
