package lcl

import (
	. "github.com/rarnu/golcl/lcl/api"
	. "github.com/rarnu/golcl/lcl/types"
	"unsafe"
)

type TSynCompletion struct {
	IComponent
	instance uintptr
	ptr      unsafe.Pointer
}

func NewSynCompletion(owner IComponent) *TSynCompletion {
	s := new(TSynCompletion)
	s.instance = SynCompletion_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	return s
}

func AsSynCompletion(obj any) *TSynCompletion {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSynCompletion{instance: instance, ptr: ptr}
}

func SynCompletionFromInst(inst uintptr) *TSynCompletion {
	return AsSynCompletion(inst)
}

func SynCompletionFromObj(obj IObject) *TSynCompletion {
	return AsSynCompletion(obj)
}

func SynCompletionFromUnsafePointer(ptr unsafe.Pointer) *TSynCompletion {
	return AsSynCompletion(ptr)
}

func (s *TSynCompletion) Free() {
	if s.instance != 0 {
		SynCompletion_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

func (s *TSynCompletion) Instance() uintptr {
	return s.instance
}

func (s *TSynCompletion) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

func (s *TSynCompletion) IsValid() bool {
	return s.instance != 0
}

func (s *TSynCompletion) Is() TIs {
	return TIs(s.instance)
}

func TSynCompletionClass() TClass {
	return SynCompletion_StaticClassType()
}

func (s *TSynCompletion) GetAutoUseSingleIdent() bool {
	return SynCompletion_GetAutoUseSingleIdent(s.instance)
}

func (s *TSynCompletion) SetAutoUseSingleIdent(value bool) {
	SynCompletion_SetAutoUseSingleIdent(s.instance, value)
}

func (s *TSynCompletion) GetCaseSensitive() bool {
	return SynCompletion_GetCaseSensitive(s.instance)
}

func (s *TSynCompletion) SetCaseSensitive(value bool) {
	SynCompletion_SetCaseSensitive(s.instance, value)
}

func (s *TSynCompletion) GetDoubleClickSelects() bool {
	return SynCompletion_GetDoubleClickSelects(s.instance)
}

func (s *TSynCompletion) SetDoubleClickSelects(value bool) {
	SynCompletion_SetDoubleClickSelects(s.instance, value)
}

func (s *TSynCompletion) GetEditor() *TSynEdit {
	return AsSynEdit(SynCompletion_GetEditor(s.instance))
}

func (s *TSynCompletion) SetEditor(value *TSynEdit) {
	SynCompletion_SetEditor(s.instance, CheckPtr(value))
}

func (s *TSynCompletion) GetEndOfTokenChr() string {
	return SynCompletion_GetEndOfTokenChr(s.instance)
}

func (s *TSynCompletion) SetEndOfTokenChr(value string) {
	SynCompletion_SetEndOfTokenChr(s.instance, value)
}

func (s *TSynCompletion) GetExecCommandID() int32 {
	return SynCompletion_GetExecCommandID(s.instance)
}

func (s *TSynCompletion) SetExecCommandID(value int32) {
	SynCompletion_SetExecCommandID(s.instance, value)
}

func (s *TSynCompletion) GetItemList() *TStrings {
	return AsStrings(SynCompletion_GetItemList(s.instance))
}

func (s *TSynCompletion) SetItemList(value *IStrings) {
	SynCompletion_SetItemList(s.instance, CheckPtr(value))
}

func (s *TSynCompletion) GetLinesInWindow() int32 {
	return SynCompletion_GetLinesInWindow(s.instance)
}

func (s *TSynCompletion) SetLinesInWindow(value int32) {
	SynCompletion_SetLinesInWindow(s.instance, value)
}

func (s *TSynCompletion) GetLongLineHintTime() int32 {
	return SynCompletion_GetLongLineHintTime(s.instance)
}

func (s *TSynCompletion) SetLongLineHintTime(value int32) {
	SynCompletion_SetLongLineHintTime(s.instance, value)
}

func (s *TSynCompletion) GetLongLineHintType() TSynCompletionLongHintType {
	return SynCompletion_GetLongLineHintType(s.instance)
}

func (s *TSynCompletion) SetLongLineHintType(value TSynCompletionLongHintType) {
	SynCompletion_SetLongLineHintType(s.instance, value)
}

func (s *TSynCompletion) Name() string {
	return SynCompletion_GetName(s.instance)
}

func (s *TSynCompletion) SetName(value string) {
	SynCompletion_SetName(s.instance, value)
}

func (s *TSynCompletion) GetPosition() int32 {
	return SynCompletion_GetPosition(s.instance)
}

func (s *TSynCompletion) SetPosition(value int32) {
	SynCompletion_SetPosition(s.instance, value)
}

func (s *TSynCompletion) GetSelectedColor() TColor {
	return SynCompletion_GetSelectedColor(s.instance)
}

func (s *TSynCompletion) SetSelectedColor(value TColor) {
	SynCompletion_SetSelectedColor(s.instance, value)
}

func (s *TSynCompletion) GetShortCut() TShortCut {
	return SynCompletion_GetShortCut(s.instance)
}

func (s *TSynCompletion) SetShortCut(value TShortCut) {
	SynCompletion_SetShortCut(s.instance, value)
}

func (s *TSynCompletion) GetShowSizeDrag() bool {
	return SynCompletion_GetShowSizeDrag(s.instance)
}

func (s *TSynCompletion) SetShowSizeDrag(value bool) {
	SynCompletion_SetShowSizeDrag(s.instance, value)
}

func (s *TSynCompletion) Tag() int {
	return SynCompletion_GetTag(s.instance)
}

func (s *TSynCompletion) SetTag(value int) {
	SynCompletion_SetTag(s.instance, value)
}

func (s *TSynCompletion) GetToggleReplaceWhole() bool {
	return SynCompletion_GetToggleReplaceWhole(s.instance)
}

func (s *TSynCompletion) SetToggleReplaceWhole(value bool) {
	SynCompletion_SetToggleReplaceWhole(s.instance, value)
}

func (s *TSynCompletion) GetWidth() int32 {
	return SynCompletion_GetWidth(s.instance)
}

func (s *TSynCompletion) SetWidth(value int32) {
	SynCompletion_SetWidth(s.instance, value)
}

func (s *TSynCompletion) SetOnCodeCompletion(fn TCodeCompletionEvent) {
	SynCompletion_SetOnCodeCompletion(s.instance, fn)
}

func (s *TSynCompletion) SetOnExecute(fn TNotifyEvent) {
	SynCompletion_SetOnExecute(s.instance, fn)
}

func (s *TSynCompletion) SetOnKeyCompletePrefix(fn TNotifyEvent) {
	SynCompletion_SetOnKeyCompletePrefix(s.instance, fn)
}

func (s *TSynCompletion) SetOnKeyNextChar(fn TNotifyEvent) {
	SynCompletion_SetOnKeyNextChar(s.instance, fn)
}

func (s *TSynCompletion) SetOnKeyPrevChar(fn TNotifyEvent) {
	SynCompletion_SetOnKeyPrevChar(s.instance, fn)
}

func (s *TSynCompletion) SetOnPositionChanged(fn TNotifyEvent) {
	SynCompletion_SetOnPositionChanged(s.instance, fn)
}

func (s *TSynCompletion) SetOnSearchPosition(fn TSynBaseCompletionSearchPosition) {
	SynCompletion_SetOnSearchPosition(s.instance, fn)
}

func (s *TSynCompletion) Deactivate() {
	SynCompletion_Deactivate(s.instance)
}

func (s *TSynCompletion) Execute(str string, x int32, y int32) {
	SynCompletion_Execute(s.instance, str, x, y)
}

func (s *TSynCompletion) IsActive() bool {
	return SynCompletion_IsActive(s.instance)
}

func (s *TSynCompletion) ClassName() string {
	return SynCompletion_ClassName(s.instance)
}

func (s *TSynCompletion) ClassType() TClass {
	return SynCompletion_ClassType(s.instance)
}

func (s *TSynCompletion) InstanceSize() int32 {
	return SynCompletion_InstanceSize(s.instance)
}

func (s *TSynCompletion) InheritsFrom(aClass TClass) bool {
	return SynCompletion_InheritsFrom(s.instance, aClass)
}

func (s *TSynCompletion) Equals(Obj IObject) bool {
	return SynCompletion_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSynCompletion) GetHashCode() int32 {
	return SynCompletion_GetHashCode(s.instance)
}

func (s *TSynCompletion) FindComponent(AName string) *TComponent {
	return AsComponent(SynCompletion_FindComponent(s.instance, AName))
}

func (s *TSynCompletion) GetNamePath() string {
	return SynCompletion_GetNamePath(s.instance)
}

func (s *TSynCompletion) HasParent() bool {
	return SynCompletion_HasParent(s.instance)
}

func (s *TSynCompletion) Assign(Source IObject) {
	SynCompletion_Assign(s.instance, CheckPtr(Source))
}

func (s *TSynCompletion) ComponentCount() int32 {
	return SynCompletion_ComponentCount(s.instance)
}

func (s *TSynCompletion) ComponentIndex() int32 {
	return SynCompletion_ComponentIndex(s.instance)
}

func (s *TSynCompletion) SetComponentIndex(value int32) {
	SynCompletion_SetComponentIndex(s.instance, value)
}

func (s *TSynCompletion) Owner() *TComponent {
	return AsComponent(SynCompletion_Owner(s.instance))
}

func (s *TSynCompletion) Components(AIndex int32) *TComponent {
	return AsComponent(SynCompletion_Components(s.instance, AIndex))
}
