//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

// 获取控件buffer文本，内部使用的
//go:noinline
func getControlBufferText(aGetTextLen func() int32, aGetTextBuf func(Buffer *string, BufSize int32) int32) string {
	strLen := aGetTextLen()
	if strLen != 0 {
		var buffStr string
		aGetTextBuf(&buffStr, strLen+1)
		return buffStr
	}
	return ""
}

// 内部的从流获取
func getStreamText(stream IStream) string {
	if stream != nil && stream.IsValid() {
		size := stream.Size()
		stream.SetPosition(0)
		if size > 0 {
			n, buff := stream.Read(int32(size))
			if n > 0 && buff[len(buff)-1] == 0 {
				return string(buff[:len(buff)-1])
			} else {
				return string(buff)
			}
		}
	}
	return ""
}
