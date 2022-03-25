package api

import (
	. "github.com/rarnu/golcl/lcl/types"
)

//--------------------------- TTreeNodes ---------------------------

func TreeNodes_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_Create").Call(obj)
	return ret
}

func TreeNodes_Free(obj uintptr) {
	_, _, _ = getLazyProc("TreeNodes_Free").Call(obj)
}

func TreeNodes_AddChildFirst(obj uintptr, Parent uintptr, S string) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddChildFirst").Call(obj, Parent, GoStrToDStr(S))
	return ret
}

func TreeNodes_AddChild(obj uintptr, Parent uintptr, S string) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddChild").Call(obj, Parent, GoStrToDStr(S))
	return ret
}

func TreeNodes_AddChildObjectFirst(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddChildObjectFirst").Call(obj, Parent, GoStrToDStr(S), Ptr)
	return ret
}

func TreeNodes_AddChildObject(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddChildObject").Call(obj, Parent, GoStrToDStr(S), Ptr)
	return ret
}

func TreeNodes_AddObjectFirst(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddObjectFirst").Call(obj, Sibling, GoStrToDStr(S), Ptr)
	return ret
}

func TreeNodes_AddObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddObject").Call(obj, Sibling, GoStrToDStr(S), Ptr)
	return ret
}

func TreeNodes_AddNode(obj uintptr, Node uintptr, Relative uintptr, S string, Ptr uintptr, Method TNodeAttachMode) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddNode").Call(obj, Node, Relative, GoStrToDStr(S), Ptr, uintptr(Method))
	return ret
}

func TreeNodes_AddFirst(obj uintptr, Sibling uintptr, S string) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_AddFirst").Call(obj, Sibling, GoStrToDStr(S))
	return ret
}

func TreeNodes_Add(obj uintptr, Sibling uintptr, S string) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_Add").Call(obj, Sibling, GoStrToDStr(S))
	return ret
}

func TreeNodes_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TreeNodes_Assign").Call(obj, Source)
}

func TreeNodes_BeginUpdate(obj uintptr) {
	_, _, _ = getLazyProc("TreeNodes_BeginUpdate").Call(obj)
}

func TreeNodes_Clear(obj uintptr) {
	_, _, _ = getLazyProc("TreeNodes_Clear").Call(obj)
}

func TreeNodes_Delete(obj uintptr, Node uintptr) {
	_, _, _ = getLazyProc("TreeNodes_Delete").Call(obj, Node)
}

func TreeNodes_EndUpdate(obj uintptr) {
	_, _, _ = getLazyProc("TreeNodes_EndUpdate").Call(obj)
}

func TreeNodes_GetFirstNode(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_GetFirstNode").Call(obj)
	return ret
}

func TreeNodes_Insert(obj uintptr, Sibling uintptr, S string) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_Insert").Call(obj, Sibling, GoStrToDStr(S))
	return ret
}

func TreeNodes_InsertObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_InsertObject").Call(obj, Sibling, GoStrToDStr(S), Ptr)
	return ret
}

func TreeNodes_CustomSort(obj uintptr, SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
	ret, _, _ := getLazyProc("TreeNodes_CustomSort").Call(obj, uintptr(SortProc), uintptr(Data), GoBoolToDBool(ARecurse))
	return DBoolToGoBool(ret)
}

func TreeNodes_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNodes_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNodes_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TreeNodes_ClassType").Call(obj)
	return TClass(ret)
}

func TreeNodes_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNodes_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNodes_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNodes_InstanceSize").Call(obj)
	return int32(ret)
}

func TreeNodes_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TreeNodes_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TreeNodes_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNodes_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TreeNodes_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNodes_GetHashCode").Call(obj)
	return int32(ret)
}

func TreeNodes_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNodes_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNodes_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNodes_GetCount").Call(obj)
	return int32(ret)
}

func TreeNodes_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_GetOwner").Call(obj)
	return ret
}

func TreeNodes_GetItem(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TreeNodes_GetItem").Call(obj, uintptr(Index))
	return ret
}

func TreeNodes_StaticClassType() TClass {
	r, _, _ := getLazyProc("TreeNodes_StaticClassType").Call()
	return TClass(r)
}
