package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

func ChartStyleEnumerator_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleEnumerator_Create").Call(obj)
	return ret
}

func ChartStyleEnumerator_Free(obj uintptr) {
	_, _, _ = getLazyProc("ChartStyleEnumerator_Free").Call(obj)
}

func ChartStyleEnumerator_StaticClassType() TClass {
	ret, _, _ := getLazyProc("ChartStyleEnumerator_StaticClassType").Call()
	return TClass(ret)
}

func ChartStyleEnumerator_GetCurrent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("ChartStyleEnumerator_GetCurrent").Call(obj)
	return ret
}

func ChartStyleEnumerator_MoveNext(obj uintptr) bool {
	ret, _, _ := getLazyProc("ChartStyleEnumerator_MoveNext").Call(obj)
	return DBoolToGoBool(ret)
}
