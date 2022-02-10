//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (j *TJPEGImage) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	j.LoadFromStream(mem)
}
