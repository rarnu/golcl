package api

import (
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

//--------------------------- TTreeNode ---------------------------

func TreeNode_Create(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_Create").Call(obj)
	return ret
}

func TreeNode_Free(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_Free").Call(obj)
}

func TreeNode_Bottom(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_Bottom").Call(obj)
	return int32(ret)
}

func TreeNode_BottomExpanded(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_BottomExpanded").Call(obj)
	return int32(ret)
}

func TreeNode_DefaultTreeViewSort(obj uintptr, Node1 uintptr, Node2 uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DefaultTreeViewSort").Call(obj, Node1, Node2)
	return int32(ret)
}

func TreeNode_DisplayExpandSignLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayExpandSignLeft").Call(obj)
	return int32(ret)
}

func TreeNode_DisplayExpandSignRect(obj uintptr) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TreeNode_DisplayExpandSignRect").Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeNode_DisplayExpandSignRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayExpandSignRight").Call(obj)
	return int32(ret)
}

func TreeNode_DisplayIconLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayIconLeft").Call(obj)
	return int32(ret)
}

func TreeNode_DisplayRect(obj uintptr, TextOnly bool) TRect {
	var ret TRect
	_, _, _ = getLazyProc("TreeNode_DisplayRect").Call(obj, GoBoolToDBool(TextOnly), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func TreeNode_DisplayStateIconLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayStateIconLeft").Call(obj)
	return int32(ret)
}

func TreeNode_DisplayTextLeft(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayTextLeft").Call(obj)
	return int32(ret)
}

func TreeNode_DisplayTextRight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_DisplayTextRight").Call(obj)
	return int32(ret)
}

func TreeNode_EditText(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_EditText").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_FindNode(obj uintptr, NodeText string) uintptr {
	ret, _, _ := getLazyProc("TreeNode_FindNode").Call(obj, GoStrToDStr(NodeText))
	return ret
}

func TreeNode_GetFirstChild(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetFirstChild").Call(obj)
	return ret
}

func TreeNode_GetFirstVisibleChild(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetFirstVisibleChild").Call(obj)
	return ret
}

func TreeNode_GetLastChild(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetLastChild").Call(obj)
	return ret
}

func TreeNode_GetLastSibling(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetLastSibling").Call(obj)
	return ret
}

func TreeNode_GetLastSubChild(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetLastSubChild").Call(obj)
	return ret
}

func TreeNode_GetLastVisibleChild(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetLastVisibleChild").Call(obj)
	return ret
}

func TreeNode_GetNext(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNext").Call(obj)
	return ret
}

func TreeNode_GetNextChild(obj uintptr, AValue uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextChild").Call(obj, AValue)
	return ret
}

func TreeNode_GetNextExpanded(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextExpanded").Call(obj)
	return ret
}

func TreeNode_GetNextMultiSelected(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextMultiSelected").Call(obj)
	return ret
}

func TreeNode_GetNextSibling(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextSibling").Call(obj)
	return ret
}

func TreeNode_GetNextSkipChildren(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextSkipChildren").Call(obj)
	return ret
}

func TreeNode_GetNextVisible(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextVisible").Call(obj)
	return ret
}

func TreeNode_GetNextVisibleSibling(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetNextVisibleSibling").Call(obj)
	return ret
}

func TreeNode_GetParentNodeOfAbsoluteLevel(obj uintptr, TheAbsoluteLevel int32) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetParentNodeOfAbsoluteLevel").Call(obj, uintptr(TheAbsoluteLevel))
	return ret
}

func TreeNode_GetPrev(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrev").Call(obj)
	return ret
}

func TreeNode_GetPrevChild(obj uintptr, AValue uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevChild").Call(obj, AValue)
	return ret
}

func TreeNode_GetPrevExpanded(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevExpanded").Call(obj)
	return ret
}

func TreeNode_GetPrevMultiSelected(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevMultiSelected").Call(obj)
	return ret
}

func TreeNode_GetPrevSibling(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevSibling").Call(obj)
	return ret
}

func TreeNode_GetPrevVisible(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevVisible").Call(obj)
	return ret
}

func TreeNode_GetPrevVisibleSibling(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetPrevVisibleSibling").Call(obj)
	return ret
}

func TreeNode_GetTextPath(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNode_GetTextPath").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNode_HasAsParent(obj uintptr, AValue uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_HasAsParent").Call(obj, AValue)
	return DBoolToGoBool(ret)
}

func TreeNode_IndexOf(obj uintptr, AValue uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_IndexOf").Call(obj, AValue)
	return int32(ret)
}

func TreeNode_IndexOfText(obj uintptr, NodeText string) int32 {
	ret, _, _ := getLazyProc("TreeNode_IndexOfText").Call(obj, GoStrToDStr(NodeText))
	return int32(ret)
}

func TreeNode_Assign(obj uintptr, Source uintptr) {
	_, _, _ = getLazyProc("TreeNode_Assign").Call(obj, Source)
}

func TreeNode_Collapse(obj uintptr, Recurse bool) {
	_, _, _ = getLazyProc("TreeNode_Collapse").Call(obj, GoBoolToDBool(Recurse))
}

func TreeNode_ConsistencyCheck(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_ConsistencyCheck").Call(obj)
}

func TreeNode_Delete(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_Delete").Call(obj)
}

func TreeNode_DeleteChildren(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_DeleteChildren").Call(obj)
}

func TreeNode_EndEdit(obj uintptr, Cancel bool) {
	_, _, _ = getLazyProc("TreeNode_EndEdit").Call(obj, GoBoolToDBool(Cancel))
}

func TreeNode_Expand(obj uintptr, Recurse bool) {
	_, _, _ = getLazyProc("TreeNode_Expand").Call(obj, GoBoolToDBool(Recurse))
}

func TreeNode_ExpandParents(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_ExpandParents").Call(obj)
}

func TreeNode_FreeAllNodeData(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_FreeAllNodeData").Call(obj)
}

func TreeNode_MakeVisible(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_MakeVisible").Call(obj)
}

func TreeNode_MoveTo(obj uintptr, Destination uintptr, Mode TNodeAttachMode) {
	_, _, _ = getLazyProc("TreeNode_MoveTo").Call(obj, Destination, uintptr(Mode))
}

func TreeNode_MultiSelectGroup(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_MultiSelectGroup").Call(obj)
}

func TreeNode_Update(obj uintptr) {
	_, _, _ = getLazyProc("TreeNode_Update").Call(obj)
}

func TreeNode_WriteDebugReport(obj uintptr, Prefix string, Recurse bool) {
	_, _, _ = getLazyProc("TreeNode_WriteDebugReport").Call(obj, GoStrToDStr(Prefix), GoBoolToDBool(Recurse))
}

func TreeNode_CustomSort(obj uintptr, SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
	ret, _, _ := getLazyProc("TreeNode_CustomSort").Call(obj, uintptr(SortProc), uintptr(Data), GoBoolToDBool(ARecurse))
	return DBoolToGoBool(ret)
}

func TreeNode_GetNamePath(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNode_GetNamePath").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNode_ClassType(obj uintptr) TClass {
	ret, _, _ := getLazyProc("TreeNode_ClassType").Call(obj)
	return TClass(ret)
}

func TreeNode_ClassName(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNode_ClassName").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNode_InstanceSize(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_InstanceSize").Call(obj)
	return int32(ret)
}

func TreeNode_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := getLazyProc("TreeNode_InheritsFrom").Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func TreeNode_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_Equals").Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func TreeNode_GetHashCode(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetHashCode").Call(obj)
	return int32(ret)
}

func TreeNode_ToString(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNode_ToString").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNode_GetAbsoluteIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetAbsoluteIndex").Call(obj)
	return int32(ret)
}

func TreeNode_GetCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetCount").Call(obj)
	return int32(ret)
}

func TreeNode_GetCut(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetCut").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetCut(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetCut").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetData(obj uintptr) unsafe.Pointer {
	ret, _, _ := getLazyProc("TreeNode_GetData").Call(obj)
	return unsafe.Pointer(ret)
}

func TreeNode_SetData(obj uintptr, value unsafe.Pointer) {
	_, _, _ = getLazyProc("TreeNode_SetData").Call(obj, uintptr(value))
}

func TreeNode_GetDeleting(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetDeleting").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_GetDropTarget(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetDropTarget").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetDropTarget(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetDropTarget").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetExpanded(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetExpanded").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetExpanded(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetExpanded").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetFocused(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetFocused").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetFocused(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetFocused").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetHandle(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetHandle").Call(obj)
	return ret
}

func TreeNode_GetHasChildren(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetHasChildren").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetHasChildren(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetHasChildren").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetHeight(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetHeight").Call(obj)
	return int32(ret)
}

func TreeNode_SetHeight(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetHeight").Call(obj, uintptr(value))
}

func TreeNode_GetImageIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetImageIndex").Call(obj)
	return int32(ret)
}

func TreeNode_SetImageIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetImageIndex").Call(obj, uintptr(value))
}

func TreeNode_GetIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetIndex").Call(obj)
	return int32(ret)
}

func TreeNode_SetIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetIndex").Call(obj, uintptr(value))
}

func TreeNode_GetIsFullHeightVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetIsFullHeightVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_GetIsVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetIsVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_GetLevel(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetLevel").Call(obj)
	return int32(ret)
}

func TreeNode_GetMultiSelected(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetMultiSelected").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetMultiSelected(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetMultiSelected").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetOverlayIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetOverlayIndex").Call(obj)
	return int32(ret)
}

func TreeNode_SetOverlayIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetOverlayIndex").Call(obj, uintptr(value))
}

func TreeNode_GetOwner(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetOwner").Call(obj)
	return ret
}

func TreeNode_GetParent(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetParent").Call(obj)
	return ret
}

func TreeNode_GetSelected(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetSelected").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetSelected(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetSelected").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetSelectedIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetSelectedIndex").Call(obj)
	return int32(ret)
}

func TreeNode_SetSelectedIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetSelectedIndex").Call(obj, uintptr(value))
}

func TreeNode_GetStateIndex(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetStateIndex").Call(obj)
	return int32(ret)
}

func TreeNode_SetStateIndex(obj uintptr, value int32) {
	_, _, _ = getLazyProc("TreeNode_SetStateIndex").Call(obj, uintptr(value))
}

func TreeNode_GetSubTreeCount(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetSubTreeCount").Call(obj)
	return int32(ret)
}

func TreeNode_GetText(obj uintptr) string {
	ret, _, _ := getLazyProc("TreeNode_GetText").Call(obj)
	return DStrToGoStr(ret)
}

func TreeNode_SetText(obj uintptr, value string) {
	_, _, _ = getLazyProc("TreeNode_SetText").Call(obj, GoStrToDStr(value))
}

func TreeNode_GetTop(obj uintptr) int32 {
	ret, _, _ := getLazyProc("TreeNode_GetTop").Call(obj)
	return int32(ret)
}

func TreeNode_GetTreeNodes(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetTreeNodes").Call(obj)
	return ret
}

func TreeNode_GetTreeView(obj uintptr) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetTreeView").Call(obj)
	return ret
}

func TreeNode_GetVisible(obj uintptr) bool {
	ret, _, _ := getLazyProc("TreeNode_GetVisible").Call(obj)
	return DBoolToGoBool(ret)
}

func TreeNode_SetVisible(obj uintptr, value bool) {
	_, _, _ = getLazyProc("TreeNode_SetVisible").Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetItems(obj uintptr, ItemIndex int32) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetItems").Call(obj, uintptr(ItemIndex))
	return ret
}

func TreeNode_SetItems(obj uintptr, ItemIndex int32, value uintptr) {
	_, _, _ = getLazyProc("TreeNode_SetItems").Call(obj, uintptr(ItemIndex), value)
}

func TreeNode_GetItem(obj uintptr, Index int32) uintptr {
	ret, _, _ := getLazyProc("TreeNode_GetItem").Call(obj, uintptr(Index))
	return ret
}

func TreeNode_SetItem(obj uintptr, Index int32, value uintptr) {
	_, _, _ = getLazyProc("TreeNode_SetItem").Call(obj, uintptr(Index), value)
}

func TreeNode_StaticClassType() TClass {
	r, _, _ := getLazyProc("TreeNode_StaticClassType").Call()
	return TClass(r)
}
