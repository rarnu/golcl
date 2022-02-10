//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

/*
  注意：Free Pascal中所有集合这里全部使用TSet(uint32)表示，也就是说最多32个元素

*/

// TAlign ENUM
type TAlign int32

const (
	AlNone = iota + 0
	AlTop
	AlBottom
	AlLeft
	AlRight
	AlClient
	AlCustom
)

// TAlignSet SET:TAlign
type TAlignSet = TSet

// TBorderStyle ENUM
type TBorderStyle int32

const (
	BsNone = iota + 0
	BsSingle
	BsSizeable
	BsDialog
	BsToolWindow
	BsSizeToolWin
)

type TFormBorderStyle TBorderStyle

// TFormStyle ENUM
type TFormStyle int32

const (
	FsNormal = iota + 0
	FsMDIChild
	FsMDIForm
	FsStayOnTop
	FsSplash
	FsSystemStayOnTop
)

// TPosition ENUM
type TPosition int32

const (
	PoDesigned        = iota + 0 // use bounds from the designer (read from stream)
	PoDefault                    // LCL decision (normally window manager decides)
	PoDefaultPosOnly             // designed size and LCL position
	PoDefaultSizeOnly            // designed position and LCL size
	PoScreenCenter               // center form on screen (depends on DefaultMonitor)
	PoDesktopCenter              // center form on desktop (total of all screens)
	PoMainFormCenter             // center form on main form (depends on DefaultMonitor)
	PoOwnerFormCenter            // center form on owner form (depends on DefaultMonitor)
	PoWorkAreaCenter             // center form on working area (depends on DefaultMonitor)
)

// TCursor = -32768..32767;
type TCursor int16

// TSeekOrigin ENUM
type TSeekOrigin int32

const (
	SoBeginning = iota + 0
	SoCurrent
	SoEnd
)

// TPixelFormat ENUM
type TPixelFormat int32

const (
	PfDevice = iota + 0
	Pf1bit
	Pf4bit
	Pf8bit
	Pf15bit
	Pf16bit
	Pf24bit
	Pf32bit
	PfCustom
)

// TAlignment ENUM
type TAlignment int32

const (
	TaLeftJustify = iota + 0
	TaRightJustify
	TaCenter
)

// TLeftRight = TAlignment.taLeftJustify..TAlignment.taRightJustify;
type TLeftRight int32

// TBiDiMode ENUM
type TBiDiMode int32

const (
	BdLeftToRight = iota + 0
	BdRightToLeft
	BdRightToLeftNoAlign
	BdRightToLeftReadingOnly
)

// TVerticalAlignment ENUM
type TVerticalAlignment int32

const (
	TaAlignTop = iota + 0
	TaAlignBottom
	TaVerticalCenter
)

// TComboBoxStyle ENUM
type TComboBoxStyle int32

const (
	CsDropDown                  = iota + 0 // like an TEdit plus a button to drop down the list, default
	CsSimple                               // like an TEdit plus a TListBox
	CsDropDownList                         // like TLabel plus a button to drop down the list
	CsOwnerDrawFixed                       // like csDropDownList, but custom drawn
	CsOwnerDrawVariable                    // like csDropDownList, but custom drawn and with each item can have another height
	CsOwnerDrawEditableFixed               // like csOwnerDrawFixed, but with TEdit
	CsOwnerDrawEditableVariable            // like csOwnerDrawVariable, but with TEdit
)

type TColorBoxStyle = TComboBoxStyle

// TWindowState ENUM
type TWindowState int32

const (
	WsNormal = iota + 0
	WsMinimized
	WsMaximized
	WsFullScreen
)

// TTextLayout ENUM
type TTextLayout int32

const (
	TlTop = iota + 0
	TlCenter
	TlBottom
)

// TEllipsisPosition ENUM
type TEllipsisPosition int32

const (
	EpNone = iota + 0
	EpPathEllipsis
	EpEndEllipsis
	EpWordEllipsis
)

type TLinkAlignment TAlignment

// TListBoxStyle ENUM
type TListBoxStyle int32

const (
	LbStandard = iota + 0
	LbOwnerDrawFixed
	LbOwnerDrawVariable
	LbVirtual
	//LbVirtualOwnerDraw
)

// TPopupAlignment ENUM
type TPopupAlignment int32

const (
	PaLeft = iota + 0
	PaRight
	PaCenter
)

// TTrackButton ENUM
type TTrackButton int32

const (
	TbRightButton = iota + 0
	TbLeftButton
)

// TProgressBarOrientation ENUM
type TProgressBarOrientation int32

const (
	PbHorizontal = iota + 0
	PbVertical
	PbRightToLeft
	PbTopDown
)

// TProgressBarStyle ENUM
type TProgressBarStyle int32

const (
	PbstNormal = iota + 0
	PbstMarquee
)

// TProgressBarState ENUM
type TProgressBarState int32

const (
	PbsNormal = iota + 0
	PbsError
	PbsPaused
)

// TButtonLayout ENUM
type TButtonLayout int32

const (
	BlGlyphLeft = iota + 0
	BlGlyphRight
	BlGlyphTop
	BlGlyphBottom
)

// TButtonState ENUM
type TButtonState int32

const (
	BsUp        = iota + 0 // button is up
	BsDisabled             // button disabled (grayed)
	BsDown                 // button is down
	BsExclusive            // button is the only down in his group
	BsHot                  // button is under mouse
)

// TButtonStyle ENUM
type TButtonStyle int32

const (
	BsAutoDetect = iota + 0
	BsWin31
	BsNew
)

// TNumGlyphs = 1..4;
type TNumGlyphs int32

// TStaticBorderStyle ENUM
type TStaticBorderStyle int32

const (
	SbsNone = iota + 0
	sbsSingle
	sbsSunken
)

// TFontStyle ENUM
type TFontStyle int32

const (
	FsBold = iota + 0
	FsItalic
	FsUnderline
	FsStrikeOut
)

// TFontStyles SET:TFontStyl
type TFontStyles = TSet

// TScrollStyle ENUM
type TScrollStyle int32

const (
	SsNone = iota + 0
	SsHorizontal
	SsVertical
	SsBoth
	SsAutoHorizontal
	SsAutoVertical
	SsAutoBoth
)

// TSortType ENUM
type TSortType int32

const (
	StNone = iota + 0
	StData
	StText
	StBoth
)

// TListArrangement ENUM
type TListArrangement int32

const (
	ArAlignBottom = iota + 0
	ArAlignLeft
	ArAlignRight
	ArAlignTop
	ArDefault
	ArSnapToGrid
)

// TViewStyle ENUM
type TViewStyle int32

const (
	VsIcon = iota + 0
	VsSmallIcon
	VsList
	VsReport
)

// TItemState ENUM
type TItemState int32

const (
	IsNone = iota + 0
	IsCut
	IsDropHilited
	IsFocused
	IsSelected
	IsActivating
)

// TItemStates SET:TItemState
type TItemStates = TSet

// TItemChange ENUM
type TItemChange int32

const (
	CtText = iota + 0
	CtImage
	CtState
)

// TItemFind ENUM
type TItemFind int32

const (
	IfData = iota + 0
	IfPartialString
	IfExactString
	IfNearest
)

// TSearchDirection ENUM
type TSearchDirection int32

const (
	SdLeft = iota + 0
	SdRight
	SdAbove
	SdBelow
	SdAll
)

// TListHotTrackStyle ENUM
type TListHotTrackStyle int32

const (
	HtHandPoint = iota + 0
	HtUnderlineCold
	HtUnderlineHot
)

// TListHotTrackStyles SET:TListHotTrackStyle
type TListHotTrackStyles = TSet

// TItemRequests ENUM
type TItemRequests int32

const (
	IrText = iota + 0
	IrImage
	IrParam
	IrState
	IrIndent
)

// TBrushStyle ENUM
type TBrushStyle int32

const (
	BsSolid = iota + 0
	BsClear
	BsHorizontal
	BsVertical
	BsFDiagonal
	BsBDiagonal
	BsCross
	BsDiagCross
	BsImage
	BsPattern
)

// TPenStyle ENUM
type TPenStyle int32

const (
	PsSolid = iota + 0
	PsDash
	PsDot
	PsDashDot
	PsDashDotDot
	PsinsideFrame
	PsPattern
	PsClear
)

// TUDBtnType ENUM
type TUDBtnType int32

const (
	BtNext = iota + 0
	BtPrev
)

// TTabPosition ENUM
type TTabPosition int32

const (
	TpTop = iota + 0
	TpBottom
	TpLeft
	TpRight
)

// TTabStyle ENUM
type TTabStyle int32

const (
	TsTabs = iota + 0
	TsButtons
	TsFlatButtons
)

// TFontPitch ENUM
type TFontPitch int32

const (
	FpDefault = iota + 0
	FpVariable
	FpFixed
)

// TPenMode ENUM
type TPenMode int32

const (
	PmBlack = iota + 0
	PmWhite
	PmNop
	PmNot
	PmCopy
	PmNotCopy
	PmMergePenNot
	PmMaskPenNot
	PmMergeNotPen
	PmMaskNotPen
	PmMerge
	PmNotMerge
	PmMask
	PmNotMask
	PmXor
	PmNotXor
)

// TTrackBarOrientation ENUM
type TTrackBarOrientation int32

const (
	TrHorizontal = iota + 0
	TrVertical
)

// TUDOrientation ENUM
type TUDOrientation int32

const (
	UdHorizontal = iota + 0
	UdVertical
)

// TFontQuality ENUM
type TFontQuality int32

const (
	FqDefault = iota + 0
	FqDraft
	FqProof
	FqNonAntialiased
	FqAntialiased
	FqClearType
	FqClearTypeNatural
)

// TCloseAction ENUM
type TCloseAction int32

const (
	CaNone = iota + 0
	CaHide
	CaFree
	CaMinimize
)

// TBalloonFlags ENUM
type TBalloonFlags int32

const (
	BfNone = iota + 0
	BfInfo
	BfWarning
	BfError
)

// TMsgDlgType ENUM
type TMsgDlgType int32

const (
	MtWarning = iota + 0
	MtError
	MtInformation
	MtConfirmation
	MtCustom
)

// TMsgDlgBtn ENUM
type TMsgDlgBtn int32

const (
	MbYes = iota + 0
	MbNo
	MbOK
	MbCancel
	MbAbort
	MbRetry
	MbIgnore
	MbAll
	MbNoToAll
	MbYesToAll
	MbHelp
	MbClose
)

// TMsgDlgButtons SET:TMsgDlgBtn
type TMsgDlgButtons = TSet

// TSysLinkType ENUM
type TSysLinkType int32

const (
	SltURL = iota + 0
	SltID
)

// TStatusPanelStyle ENUM
type TStatusPanelStyle int32

const (
	PsText = iota + 0
	PsOwnerDraw
)

// TJPEGPerformance ENUM
type TJPEGPerformance int32

const (
	JpBestQuality = iota + 0
	JpBestSpeed
)

type TJPEGPixelFormat = TPixelFormat

//const (
//	Jf24Bit = iota + 0
//	Jf8Bit
//)

type TShortCut uint16

// TNodeState ENUM
type TNodeState int32

const (
	NsCut           = iota + 0 // = Node.Cut
	NsDropHilite               // = Node.DropTarget
	NsFocused                  // = Node.Focused
	NsSelected                 // = Node.Selected
	NsMultiSelected            // = Node.MultiSelected
	NsExpanded                 // = Node.Expanded
	NsHasChildren              // = Node.HasChildren
	NsDeleting                 // = Node.Deleting, set on Destroy
	NsVisible                  // = Node.Visible
	NsBound                    // bound to a tree, e.g. has Parent or is top lvl node
)

// TNodeAttachMode ENUM
type TNodeAttachMode int32

const (
	NaAdd           = iota + 0 // add as last sibling of Destination
	NaAddFirst                 // add as first sibling of Destination
	NaAddChild                 // add as last child of Destination
	NaAddChildFirst            // add as first child of Destination
	NaInsert                   // insert in front of Destination
	NaInsertBehind             // insert behind Destination
)

// TAddMode ENUM
type TAddMode int32

const (
	TaAddFirst = iota + 0
	TaAdd
	TaInsert
)

// TMultiSelectStyles ENUM
type TMultiSelectStyles int32

const (
	MsControlSelect = iota + 0
	MsShiftSelect
	MsVisibleOnly
	MsSiblingOnly
)

// TMultiSelectStyle SET:TMultiSelectStyles
type TMultiSelectStyle = TSet

// TActionListState ENUM
type TActionListState int32

const (
	AsNormal = iota + 0
	AsSuspended
	AsSuspendedEnabled
)

// TGradientDirection ENUM
type TGradientDirection int32

const (
	GdHorizontal = iota + 0
	GdVertical
)

// TDrawingStyle ENUM
type TDrawingStyle int32

const (
	DSFocus = iota + 0
	DSSelected
	DSNormal
	DSTransparent
)

// TImageType ENUM
type TImageType int32

const (
	ItImage = iota + 0
	ItMask
)

// TResType ENUM
type TResType int32

const (
	RtBitmap = iota + 0
	RtCursor
	RtIcon
)

// TLoadResource ENUM
type TLoadResource int32

const (
	LrDefaultColor = iota + 0
	LrDefaultSize
	LrFromFile
	LrMap3DColors
	LrTransparent
	LrMonoChrome
)

// TLoadResources SET:TLoadResource
type TLoadResources = TSet

// TColorDepth ENUM
type TColorDepth int32

const (
	CdDefault = iota + 0
	CdDeviceDependent
	Cd4Bit
	Cd8Bit
	Cd16Bit
	Cd24Bit
	Cd32Bit
)

// TCheckBoxState ENUM
type TCheckBoxState int32

const (
	CbUnchecked = iota + 0
	CbChecked
	CbGrayed
)

// TToolButtonStyle ENUM
type TToolButtonStyle int32

const (
	TbsButton     = iota + 0 // button (can be clicked)
	TbsCheck                 // check item (click to toggle state, can be grouped)
	TbsDropDown              // button with dropdown button to show a popup menu
	TbsSeparator             // space holder
	TbsDivider               // space holder with line
	TbsButtonDrop            // button with arrow (not separated from each other)
)

// TTBGradientDrawingOption ENUM
type TTBGradientDrawingOption int32

const (
	GdoHotTrack = iota + 0
	GdoGradient
)

// TTBGradientDrawingOptions SET:TTBGradientDrawingOption
type TTBGradientDrawingOptions = TSet

// TColorDialogOption ENUM
type TColorDialogOption int32

const (
	CdFullOpen = iota + 0
	CdPreventFullOpen
	CdShowHelp
	CdSolidColor
	CdAnyColor
)

// TColorDialogOptions SET:TColorDialogOption
type TColorDialogOptions = TSet

// TBorderIcon ENUM
type TBorderIcon int32

const (
	BiSystemMenu = iota + 0
	BiMinimize
	BiMaximize
	BiHelp
)

// TBorderIcons SET:TBorderIcon
type TBorderIcons = TSet

// TFontDialogOption ENUM
type TFontDialogOption int32

const (
	FdAnsiOnly = iota + 0
	FdTrueTypeOnly
	FdEffects
	FdFixedPitchOnly
	FdForceFontExist
	FdNoFaceSel
	FdNoOEMFonts
	FdNoSimulations
	FdNoSizeSel
	FdNoStyleSel
	FdNoVectorFonts
	FdShowHelp
	FdWysiwyg
	FdLimitSize
	FdScalableOnly
	FdApplyButton
)

// TFontDialogOptions SET:TFontDialogOptio
type TFontDialogOptions = TSet

// TOpenOption ENUM
type TOpenOption int32

const (
	OfReadOnly        = iota + 0
	OfOverwritePrompt // if selected file exists shows a message, that file
	// OfHideReadOnly will be overwritten
	OfHideReadOnly // hide read only file
	OfNoChangeDir  // do not change current directory
	OfShowHelp     // show a help button
	OfNoValidate
	OfAllowMultiSelect // allow multiselection
	OfExtensionDifferent
	OfPathMustExist // shows an error message if selected path does not exist
	OfFileMustExist // shows an error message if selected file does not exist
	OfCreatePrompt
	OfShareAware
	OfNoReadOnlyReturn // do not return filenames that are readonly
	OfNoTestFileCreate
	OfNoNetworkButton
	OfNoLongNames
	OfOldStyleDialog
	OfNoDereferenceLinks // do not resolve links while dialog is shown (only on Windows, see OFN_NODEREFERENCELINKS)
	OfNoResolveLinks     // do not resolve links after Execute
	OfEnableIncludeNotify
	OfEnableSizing    // dialog can be resized, e.g. via the mouse
	OfDontAddToRecent // do not add the path to the history list
	OfForceShowHidden // show hidden files
	OfViewDetail      // details are OS and interface dependent
	OfAutoPreview     // details are OS and interface dependent
)

// TOpenOptions SET:TOpenOption
type TOpenOptions = TSet

// TOpenOptionEx ENUM
type TOpenOptionEx int32

const (
	OfExNoPlacesBar = iota + 0
)

// TOpenOptionsEx SET:TOpenOptionEx
type TOpenOptionsEx = TSet

// TPrintRange ENUM
type TPrintRange int32

const (
	PrAllPages = iota + 0
	PrSelection
	PrPageNums
	PrCurrentPage
)

// TPrintDialogOption ENUM
type TPrintDialogOption int32

const (
	PoPrintToFile = iota + 0
	PoPageNums
	PoSelection
	PoWarning
	PoHelp
	PoDisablePrintToFile
	PoBeforeBeginDoc
)

// TPrintDialogOptions SET:TPrintDialogOption
type TPrintDialogOptions = TSet

// TPageSetupDialogOption ENUM
type TPageSetupDialogOption int32

const (
	PsoDefaultMinMargins = iota + 0
	PsoDisableMargins
	PsoDisableOrientation
	PsoDisablePagePainting
	PsoDisablePaper
	PsoDisablePrinter
	PsoMargins
	PsoMinMargins
	PsoShowHelp
	PsoWarning
	PsoNoNetworkButton
)

// TPageSetupDialogOptions SET:TPageSetupDialogOption
type TPageSetupDialogOptions = TSet

// TPrinterKind ENUM
type TPrinterKind int32

const (
	PkDotMatrix = iota + 0
	PkHPPCL
)

// TPageType ENUM
type TPageType int32

const (
	PtEnvelope = iota + 0
	PtPaper
)

// TPageMeasureUnits ENUM
type TPageMeasureUnits int32

const (
	PmDefault = iota + 0
	PmMillimeters
	PmInches
)

// TStringsOption ENUM
type TStringsOption int32

const (
	SoStrictDelimiter = iota + 0
	SoWriteBOM
	SoTrailingLineBreak
	SoUseLocale
)

// TStringsOptions SET:TStringsOption
type TStringsOptions = TSet

// TShiftStateEnum ENUM
type TShiftStateEnum int32

const (
	SsShift = iota + 0
	SsAlt
	SsCtrl
	SsLeft
	SsRight
	SsMiddle
	SsDouble
	// SsMeta Extra additions
	SsMeta
	SsSuper
	SsHyper
	SsAltGr
	SsCaps
	SsNum
	SsScroll
	SsTriple
	SsQuad
	SsExtra1
	SsExtra2
)

// TShiftState SET:TShiftStateEnum
type TShiftState = TSet

// TMouseButton ENUM
type TMouseButton int32

const (
	MbLeft = iota + 0
	MbRight
	MbMiddle
	mbExtra1
	mbExtra2
)

// TFillStyle ENUM
type TFillStyle int32

const (
	FsSurface = iota + 0
	FsBorder
)

// TFillMode ENUM
type TFillMode int32

const (
	FmAlternate = iota + 0
	FmWinding
)

// TCanvasStates ENUM
type TCanvasStates int32

const (
	CsHandleValid = iota + 0
	CsFontValid
	CsPenValid
	CsBrushValid
	CsRegionValid
)

// TCanvasState SET:TCanvasStates
type TCanvasState = TSet

// TCanvasOrientation ENUM
type TCanvasOrientation int32

const (
	CoLeftToRight = iota + 0
	CoRightToLeft
)

// TTextFormats ENUM
type TTextFormats int32

const (
	TfBottom = iota + 0
	TfCalcRect
	TfCenter
	TfEditControl
	TfEndEllipsis
	TfPathEllipsis
	TfExpandTabs
	TfExternalLeading
	TfLeft
	TfModifyString
	TfNoClip
	TfNoPrefix
	TfRight
	TfRtlReading
	TfSingleLine
	TfTop
	TfVerticalCenter
	TfWordBreak
	TfHidePrefix
	TfNoFullWidthCharBreak
	TfPrefixOnly
	TfTabStop
	TfWordEllipsis
	TfComposited
)

// TTextFormat SET:TTextFormats
type TTextFormat = TSet

// TBevelCut ENUM
type TBevelCut int32

const (
	BvNone = iota + 0
	BvLowered
	BvRaised
	BvSpace
)

// TBevelEdge ENUM
type TBevelEdge int32

const (
	BeLeft = iota + 0
	BeTop
	BeRight
	BeBottom
)

// TBevelEdges SET:TBevelEdge
type TBevelEdges = TSet

// TBevelKind ENUM
type TBevelKind int32

const (
	BkNone = iota + 0
	BkTile
	BkSoft
	BkFlat
)

// TTickMark ENUM
type TTickMark int32

const (
	TmBottomRight = iota + 0
	TmTopLeft
	TmBoth
)

// TTickStyle ENUM
type TTickStyle int32

const (
	TsNone = iota + 0
	TsAuto
	TsManual
)

// TPositionToolTip ENUM
type TPositionToolTip int32

const (
	PtNone = iota + 0
	PtTop
	PtLeft
	PtBottom
	PtRight
)

// TDateTimeKind ENUM
type TDateTimeKind int32

const (
	DtkDate = iota + 0
	DtkTime
	DtkDateTime
)

// TDTDateMode ENUM
type TDTDateMode int32

const (
	DmComboBox = iota + 0
	DmUpDown
	DmNone
)

// TDTDateFormat ENUM
type TDTDateFormat int32

const (
	DfShort = iota + 0
	DfLong
)

// TDTCalAlignment ENUM
type TDTCalAlignment int32

const (
	DtaLeft = iota + 0
	DtaRight
	DtaDefault
)

// TCalDayOfWeek ENUM
type TCalDayOfWeek int32

const (
	DowMonday = iota + 0
	DowTuesday
	DowWednesday
	DowThursday
	DowFriday
	DowSaturday
	DowSunday
	DowLocaleDefault
)

// TSearchType ENUM
type TSearchType int32

const (
	StWholeWord = iota + 0
	StMatchCase
)

// TSearchTypes SET:TSearchType
type TSearchTypes = TSet

// TNumberingStyle ENUM
type TNumberingStyle int32

const (
	NsNone = iota + 0
	NsBullte
)

// TAttributeType ENUM
type TAttributeType int32

const (
	AtSelected = iota + 0
	AtDefaultText
)

// TConsistentAttribute ENUM
type TConsistentAttribute int32

const (
	CaBold = iota + 0
	CaColor
	CaFace
	CaItalic
	CaSize
	CaStrikeOut
	CaUnderline
	CaProtected
)

// TConsistentAttributes SET:TConsistentAttribute
type TConsistentAttributes = TSet

// TIconArrangement ENUM
type TIconArrangement int32

const (
	IaTop = iota + 0
	IaLeft
)

// THeaderStyle ENUM
type THeaderStyle int32

const (
	HsGradient = iota + 0
	HsImage
	HsThemed
)

// TImageAlignment ENUM
type TImageAlignment int32

// IaTop有冲突，所以增加一个i
const (
	IiaLeft = iota + 0
	IiaRight
	IiaTop
	IiaBottom
	IiaCenter
)

// TAnchorKind ENUM
type TAnchorKind int32

const (
	AkTop = iota + 0
	AkLeft
	AkRight
	AkBottom
)

// TAnchors SET:TAnchorKind
type TAnchors = TSet

// TOwnerDrawStateType ENUM
type TOwnerDrawStateType int32

const (
	OdSelected = iota + 0
	OdGrayed
	OdDisabled
	OdChecked
	OdFocused
	OdDefault
	OdHotLight
	OdInactive
	OdNoAccel
	OdNoFocusRect
	OdReserved1
	OdReserved2
	OdComboBoxEdit
	OdBackgroundPainted // item background already painted
)

// TOwnerDrawState SET:TOwnerDrawStateType
type TOwnerDrawState = TSet

// TBitBtnKind ENUM
type TBitBtnKind int32

const (
	BkCustom = iota + 0
	BkOK
	BkCancel
	BkHelp
	BkYes
	BkNo
	BkClose
	BkAbort
	BkRetry
	BkIgnore
	BkAll
	BkNoToAll
	BkYesToAll
)

// TScrollBarKind ENUM
type TScrollBarKind int32

const (
	SbHorizontal = iota + 0
	SbVertical
)

// TScrollBarInc = 1..32767;
type TScrollBarInc int16

// TScrollBarStyle ENUM
type TScrollBarStyle int32

const (
	SsRegular = iota + 0
	SsFlat
	SsHotTrack
)

// TShapeType ENUM
type TShapeType int32

const (
	StRectangle = iota + 0
	StSquare
	StRoundRect
	StRoundSquare
	StEllipse
	StCircle
	StSquaredDiamond
	StDiamond
	StTriangle
	StTriangleLeft
	StTriangleRight
	StTriangleDown
	StStar
	StStarDown
)

// TBevelStyle = (bsLowered, bsRaised);
type TBevelStyle int32

const (
	BsLowered = iota + 0
	BsRaised
)

// TBevelShape ENUM
type TBevelShape int32

const (
	BsBox = iota + 0
	BsFrame
	BsTopLine
	BsBottomLine
	BsLeftLine
	BsRightLine
	BsSpacer
)

// TGaugeKind ENUM
type TGaugeKind int32

const (
	GkText = iota + 0
	GkHorizontalBar
	GkVerticalBar
	GkPie
	GkNeedle
	GkHalfPie
)

// TCustomDrawTarget ENUM
type TCustomDrawTarget int32

const (
	DtControl = iota + 0
	DtItem
	DtSubItem
)

// TCustomDrawStage ENUM
type TCustomDrawStage int32

const (
	CdPrePaint = iota + 0
	CdPostPaint
	CdPreErase
	CdPostErase
)

// TCustomDrawStateFlag ENUM
type TCustomDrawStateFlag int32

const (
	CdsSelected = iota + 0
	CdsGrayed
	CdsDisabled
	CdsChecked
	CdsFocused
	CdsDefault
	CdsHot
	CdsMarked
	CdsIndeterminate
)

// TCustomDrawState SET:TCustomDrawStateFlag
type TCustomDrawState = TSet

// TDisplayCode ENUM
type TDisplayCode int32

const (
	DrBounds = iota + 0
	DrIcon
	DrLabel
	DrSelectBounds
)

// TSelectDirOpt ENUM
type TSelectDirOpt int32

const (
	SdAllowCreate = iota + 0
	SdPerformCreate
	SdPrompt
)

// TSelectDirOpts SET:TSelectDirOpt
type TSelectDirOpts = TSet

// TFindOption ENUM
type TFindOption int32

const (
	FrDown = iota + 0
	FrFindNext
	FrHideMatchCase
	FrHideWholeWord
	FrHideUpDown
	FrMatchCase
	FrDisableMatchCase
	FrDisableUpDown
	FrDisableWholeWord
	FrReplace
	FrReplaceAll
	FrWholeWord
	FrShowHelp
	FrEntireScope
	FrHideEntireScope
	FrPromptOnReplace
	FrHidePromptOnReplace
	FrButtonsAtBottom
)

// TFindOptions SET:TFindOption
type TFindOptions = TSet

// TDragMode ENUM
type TDragMode int32

const (
	DmManual = iota + 0
	DmAutomatic
)

// TDragState ENUM
type TDragState int32

const (
	DsDragEnter = iota + 0
	DsDragLeave
	DsDragMove
)

// TDragKind ENUM
type TDragKind int32

const (
	DkDrag = iota + 0
	DkDock
)

// TEditCharCase ENUM
type TEditCharCase int32

const (
	EcNormal = iota + 0
	EcUpperCase
	EcLowerCase
)

// TEdgeBorder ENUM
type TEdgeBorder int32

const (
	EbLeft = iota + 0
	EbTop
	EbRight
	EbBottom
)

// TEdgeBorders SET:TEdgeBorder
type TEdgeBorders = TSet

// TEdgeStyle ENUM
type TEdgeStyle int32

const (
	EsNone = iota + 0
	EsRaised
	EsLowered
)

// TGridDrawingStyle ENUM
type TGridDrawingStyle int32

const (
	GdsClassic = iota + 0
	GdsThemed
	GdsGradient
)

// TGridOption ENUM
type TGridOption int32

const (
	GoFixedVertLine = iota + 0
	GoFixedHorzLine
	GoVertLine
	GoHorzLine
	GoRangeSelect
	GoDrawFocusSelected
	GoRowSizing
	GoColSizing
	GoRowMoving
	GoColMoving
	GoEditing
	GoAutoAddRows
	GoTabs
	GoRowSelect
	GoAlwaysShowEditor
	GoThumbTracking
	// GoColSpanning Additional Options
	GoColSpanning                 // Enable cellextent calcs
	GoRelaxedRowSelect            // User can see focused cell on goRowSelect
	GoDblClickAutoSize            // dblclicking columns borders (on hdrs) resize col.
	GoSmoothScroll                // Switch scrolling mode (pixel scroll is by default)
	GoFixedRowNumbering           // Ya
	GoScrollKeepVisible           // keeps focused cell visible while scrolling
	GoHeaderHotTracking           // Header cells change look when mouse is over them
	GoHeaderPushedLook            // Header cells looks pushed when clicked
	GoSelectionActive             // Setting grid.Selection moves also cell cursor
	GoFixedColSizing              // Allow to resize fixed columns
	GoDontScrollPartCell          // clicking partially visible cells will not scroll
	GoCellHints                   // show individual cell hints
	GoTruncCellHints              // show cell hints if cell text is too long
	GoCellEllipsis                // show "..." if cell text is too long
	GoAutoAddRowsSkipContentCheck //BB Also add a row (if AutoAddRows in Options) if last row is empty
	GoRowHighlight                // Highlight the current Row
)

// TGridOptions SET:TGridOption
type TGridOptions = TSet

// TGridDrawStates ENUM
type TGridDrawStates int32

const (
	GdSelected = iota + 0
	GdFocused
	GdFixed
	GdHot
	GdPushed
	GdRowHighlight
)

// TGridDrawState SET:TGridDrawStates
type TGridDrawState = TSet

// THeaderSectionStyle ENUM
type THeaderSectionStyle int32

const (
	HsText = iota + 0
	HsOwnerDraw
)

// TLabelPosition ENUM
type TLabelPosition int32

const (
	LpAbove = iota + 0
	LpBelow
	LpLeft
	LpRight
)

// TFlowStyle ENUM
type TFlowStyle int32

const (
	FsLeftRightTopBottom = iota + 0
	FsRightLeftTopBottom
	FsLeftRightBottomTop
	FsRightLeftBottomTop
	FsTopBottomLeftRight
	FsBottomTopLeftRight
	FsTopBottomRightLeft
	FsBottomTopRightLeft
)

// TCoolBandMaximize ENUM
type TCoolBandMaximize int32

const (
	BmNone = iota + 0
	BmClick
	BmDblClick
)

// TMenuBreak ENUM
type TMenuBreak int32

const (
	MbNone = iota + 0
	MbBreak
	MbBarBreak
)

// TSectionTrackState ENUM
type TSectionTrackState int32

const (
	TsTrackBegin = iota + 0
	TsTrackMove
	TsTrackEnd
)

// TControlStateType ENUM
type TControlStateType int32

const (
	CsLButtonDown = iota + 0
	CsClicked
	CsPalette
	CsReadingState
	CsFocusing
	CsCreating // not used, exists for Delphi compatibility
	CsPaintCopy
	CsCustomPaint
	CsDestroyingHandle
	CsDocking
	CsVisibleSetInLoading
)

// TControlState SET:TControlStateType
type TControlState = TSet

// TControlStyleType ENUM
type TControlStyleType int32

const (
	CsAcceptsControls            = iota + 0 // can have children in the designer
	CsCaptureMouse                          // auto capture mouse when clicked
	CsDesignInteractive                     // wants mouse events in design mode
	CsClickEvents                           // handles mouse events
	CsFramed                                // not implemented, has 3d frame
	CsSetCaption                            // if Name=Caption, changing the Name changes the Caption
	CsOpaque                                // the control paints its area completely
	CsDoubleClicks                          // understands mouse double clicks
	CsTripleClicks                          // understands mouse triple clicks
	CsQuadClicks                            // understands mouse quad clicks
	CsFixedWidth                            // cannot change its width
	CsFixedHeight                           // cannot change its height (for example combobox)
	CsNoDesignVisible                       // is invisible in the designer
	CsReplicatable                          // PaintTo works
	CsNoStdEvents                           // standard events such as mouse, key, and click events are ignored.
	CsDisplayDragImage                      // display images from dragimagelist during drag operation over control
	CsReflector                             // not implemented, the controls respond to size, focus and dlg messages - it can be used as ActiveX control under Windows
	CsActionClient                          // Action is set
	CsMenuEvents                            // not implemented
	CsNoFocus                               // control will not take focus when clicked with mouse.
	CsNeedsBorderPaint                      // not implemented
	CsParentBackground                      // tells WinXP to paint the theme background of parent on controls background
	CsDesignNoSmoothResize                  // when resizing control in the designer do not SetBounds while dragging
	CsDesignFixedBounds                     // can not be moved nor resized in designer
	CsHasDefaultAction                      // implements useful ExecuteDefaultAction
	CsHasCancelAction                       // implements useful ExecuteCancelAction
	CsNoDesignSelectable                    // can not be selected at design time
	CsOwnedChildrenNotSelectable            // child controls owned by this control are NOT selectable in the designer
	CsAutoSize0x0                           // if the preferred size is 0x0 then control is shrinked ot 0x0
	CsAutoSizeKeepChildLeft                 // when AutoSize=true do not move children horizontally
	CsAutoSizeKeepChildTop                  // when AutoSize=true do not move children vertically
	CsRequiresKeyboardInput                 // If the device has no physical keyboard then show the virtual keyboard when this control gets focus (therefore available only to TWinControl descendents)
)

// TControlStyle SET:TControlStyleType
type TControlStyle = TSet

// TMouseActivate ENUM
type TMouseActivate int32

const (
	MaDefault = iota + 0
	MaActivate
	MaActivateAndEat
	MaNoActivate
	MaNoActivateAndEat
)

// TTaskBarProgressState ENUM
type TTaskBarProgressState int32

const (
	None = iota + 0
	Indeterminate
	Normal
	Error
	Paused
)

// TBitmapHandleType ENUM
type TBitmapHandleType int32

const (
	BmDIB = iota + 0
	BmDDB
)

// TPrinterState ENUM
type TPrinterState int32

const (
	PsNoDefine = iota + 0
	PsReady
	PsPrinting
	PsStopped
)

// TPrinterOrientation ENUM
type TPrinterOrientation int32

const (
	PoPortrait = iota + 0
	PoLandscape
	PoReverseLandscape
	PoReversePortrait
)

// TPrinterCapability ENUM
type TPrinterCapability int32

const (
	PcCopies = iota + 0
	PcOrientation
	PcCollation
)

// TPrinterCapabilities SET:TPrinterCapability
type TPrinterCapabilities = TSet

// TPrinterType ENUM
type TPrinterType int32

const (
	PtLocal = iota + 0
	PtNetWork
)

// TReadyState ENUM
type TReadyState int32

const (
	RsUninitialized = iota + 0
	RsLoading
	RsLoaded
	RsInterActive
	RsComplete
)

// TStringEncoding ENUM
type TStringEncoding int32

const (
	SeUnknown = iota + 0
	SeANSI
	SeUnicode
	SeUTF8
)

// TShowInTaskbar ENUM
type TShowInTaskbar int32

const (
	StDefault = iota + 0 // use default rules for showing taskbar item
	StAlways             // always show taskbar item for the form
	StNever              // never show taskbar item for the form
)

// TTaskDialogCommonButton ENUM
type TTaskDialogCommonButton int32

const (
	TcbOk = iota + 0
	TcbYes
	TcbNo
	TcbCancel
	TcbRetry
	TcbClose
)

// TTaskDialogCommonButtons SET:TTaskDialogCommonButton
type TTaskDialogCommonButtons = TSet

// TTaskDialogFlag ENUM
type TTaskDialogFlag int32

const (
	TfEnableHyperlinks = iota + 0
	TfUseHiconMain
	TfUseHiconFooter
	TfAllowDialogCancellation
	TfUseCommandLinks
	TfUseCommandLinksNoIcon
	TfExpandFooterArea
	TfExpandedByDefault
	TfVerificationFlagChecked
	TfShowProgressBar
	TfShowMarqueeProgressBar
	TfCallbackTimer
	TfPositionRelativeToWindow
	TfRtlLayout
	TfNoDefaultRadioButton
	TfCanBeMinimized
)

// TTaskDialogFlags SET:TTaskDialogFlag
type TTaskDialogFlags = TSet

// TTaskDialogIcon ENUM
type TTaskDialogIcon int32

const (
	TdiNone = iota + 0
	TdiWarning
	TdiError
	TdiInformation
	TdiShield
	TdiQuestion
)

// TComboBoxExStyle ENUM
type TComboBoxExStyle int32

const (
	CsExDropDown = iota + 0
	CsExSimple
	CsExDropDownList
)

// TComboBoxExStyleEx ENUM
type TComboBoxExStyleEx int32

const (
	CsExCaseSensitive = iota + 0
	CsExNoEditImage
	CsExNoEditImageIndent
	CsExNoSizeLimit
	CsExPathWordBreak
)

// TComboBoxExStyles SET:TComboBoxExStyleEx
type TComboBoxExStyles = TSet

// TAutoCompleteOption ENUM
type TAutoCompleteOption int32

const (
	AcoAutoSuggest = iota + 0
	AcoAutoAppend
	AcoSearch
	AcoFilterPrefixes
	AcoUseTab
	AcoUpDownKeyDropsList
	AcoRtlReading
)

// TAutoCompleteOptions SET:TAutoCompleteOption
type TAutoCompleteOptions = TSet

// TDefaultMonitor ENUM
type TDefaultMonitor int32

const (
	DmDesktop = iota + 0
	DmPrimary
	DmMainForm
	DmActiveForm
)

// TTransparentMode ENUM
type TTransparentMode int32

const (
	TmAuto = iota + 0
	TmFixed
)

// TDrawImageMode ENUM
type TDrawImageMode int32

const (
	DimNormal = iota + 0
	DimCenter
	DimStretch
)

// TListBoxOption ENUM
type TListBoxOption int32

const (
	LboDrawFocusRect = iota + 0 // draw focus rect in case of owner drawing
)

// TListBoxOptions SET:TListBoxOption
type TListBoxOptions = TSet

// TAntialiasingMode ENUM
type TAntialiasingMode int32

const (
	AmDontCare = iota + 0 // default antialiasing
	AmOn                  // enabled
	AmOff                 // disabled
)

// TSortDirection ENUM
type TSortDirection int32

const (
	SdAscending = iota + 0
	SdDescending
)

// TTreeViewExpandSignType ENUM
type TTreeViewExpandSignType int32

const (
	TvestTheme     = iota + 0 // use themed sign
	TvestPlusMinus            // use +/- sign
	TvestArrow                // use blank arrow
	TvestArrowFill            // use filled arrow
)

// TTreeViewOption ENUM
type TTreeViewOption int32

const (
	TvoAllowMultiselect = iota + 0
	TvoAutoExpand
	TvoAutoInsertMark
	TvoAutoItemHeight
	TvoHideSelection
	TvoHotTrack
	TvoKeepCollapsedNodes
	TvoReadOnly
	TvoRightClickSelect
	TvoRowSelect
	TvoShowButtons
	TvoShowLines
	TvoShowRoot
	TvoShowSeparators
	TvoToolTips
	TvoNoDoubleClickExpand
	TvoThemedDraw
)

// TTreeViewOptions SET:TTreeViewOption
type TTreeViewOptions = TSet

// TGlyphShowMode ENUM
type TGlyphShowMode int32

const (
	GsmAlways      = iota + 0 // always show
	GsmNever                  // never show
	GsmApplication            // depends on application settings
	GsmSystem                 // depends on system settings
)

// TCTabControlOption ENUM
type TCTabControlOption int32

const (
	NboShowCloseButtons = iota + 0
	NboMultiLine
	NboHidePageListPopup
	NboKeyboardTabSwitch
	NboShowAddTabButton
	NboDoChangeOnSetIndex
)

// TCTabControlOptions SET:TCTabControlOption
type TCTabControlOptions = TSet

// TAnchorSideReference ENUM
type TAnchorSideReference int32

const (
	AsrTop = iota + 0
	AsrBottom
	AsrCenter
)

// TControlCellAlign ENUM
type TControlCellAlign int32

const (
	CcaFill = iota + 0
	CcaLeftTop
	CcaRightBottom
	CcaCenter
)

// TControlCellAligns SET:TControlCellAlign
type TControlCellAligns = TSet

// TChildControlResizeStyle ENUM
type TChildControlResizeStyle int32

const (
	CrsAnchorAligning        = iota + 0 // (like Delphi)
	CrsScaleChilds                      // scale children equally, keep space between children fixed
	CrsHomogenousChildResize            // enlarge children equally (i.e. by the same amount of pixel)
	CrsHomogenousSpaceResize            // enlarge space between children equally
	//{$IFDEF EnablecrsSameSize}
	//,CrsSameSize  // each child gets the same size (maybe one pixel difference)
	//{$ENDIF}
)

// TControlChildrenLayout ENUM
type TControlChildrenLayout int32

const (
	CclNone                       = iota + 0
	CclLeftToRightThenTopToBottom // if BiDiMode <> bdLeftToRight then it becomes RightToLeft
	CclTopToBottomThenLeftToRight
)

// TColumnLayout ENUM
type TColumnLayout int32

const (
	ClHorizontalThenVertical = iota + 0
	ClVerticalThenHorizontal
)

// TSortIndicator ENUM
type TSortIndicator int32

const (
	SiNone = iota + 0
	SiAscending
	SiDescending
)

// TLibType VCL或者LCL，只是用于引入的
type TLibType int32

const (
	LtVCL TLibType = iota + 0
	LtLCL
)

// TColumnButtonStyle ENUM
type TColumnButtonStyle int32

const (
	CbsAuto = iota + 0
	CbsEllipsis
	CbsNone
	CbsPickList
	CbsCheckboxColumn
	CbsButton
	CbsButtonColumn
)

// TGridZone ENUM
type TGridZone int32

const (
	GzNormal = iota + 0
	GzFixedCols
	GzFixedRows
	GzFixedCells
	GzInvalid
)

// TGridZoneSet SET:TGridZone
type TGridZoneSet = TSet

// TSortOrder ENUM
type TSortOrder int32

const (
	SoAscending = iota + 0
	SoDescending
)

// TAutoAdvance ENUM
type TAutoAdvance int32

const (
	AaNone = iota + 0
	AaDown
	AaRight
	AaLeft
	AaRightDown
	AaLeftDown
	AaRightUp
	AaLeftUp
)

// TCellHintPriority ENUM
type TCellHintPriority int32

const (
	ChpAll = iota + 0
	ChpAllNoDefault
	ChpTruncOnly
)

// TMouseWheelOption ENUM
type TMouseWheelOption int32

const (
	MwCursor = iota + 0
	MwGrid
)

// TGridOption2 ENUM
type TGridOption2 int32

const (
	GoScrollToLastCol = iota + 0 // allow scrolling to last column (so that last column can be leftcol)
	GoScrollToLastRow            // allow scrolling to last row (so that last row can be toprow)
)

// TGridOptions2 SET:TGridOption2
type TGridOptions2 = TSet

// TRangeSelectMode ENUM
type TRangeSelectMode int32

const (
	RsmSingle = iota + 0
	RsmMulti
)

// TTitleStyle ENUM
type TTitleStyle int32

const (
	TsLazarus = iota + 0
	TsStandard
	TsNative
)

// TPrefixOption ENUM
type TPrefixOption int32

const (
	PoNone = iota + 0
	PoHeaderClick
)

// TDisplaySetting ENUM
type TDisplaySetting int32

const (
	DsShowHeadings = iota + 0
	DsShowDayNames
	DsNoMonthChange
	DsShowWeekNumbers
	DsStartMonday
)

// TTimeFormat ENUM
type TTimeFormat int32

const (
	Tf12 = iota + 0 // 12 hours format, with am/pm string
	Tf24            // 24 hours format
)

// TTimeDisplay ENUM
type TTimeDisplay int32

const (
	TdHM    = iota + 0 // hour and minute
	TdHMS              // hour Minute and second
	TdHMSMs            // hour Minute Second and milisecond
)

// TArrowShape ENUM
type TArrowShape int32

const (
	AsClassicSmaller = iota + 0
	AsClassicLarger
	AsModernSmaller
	AsModernLarger
	AsYetAnotherShape
	AsTheme
)

// TDateDisplayOrder ENUM
type TDateDisplayOrder int32

const (
	DdoDMY = iota + 0
	DdoMDY
	DdoYMD
	DdoTryDefault
)

// TDateTimePart ENUM
type TDateTimePart int32

const (
	DtpDay = iota + 0
	DtpMonth
	DtpYear
	DtpHour
	DtpMinute
	DtpSecond
	DtpMiliSec
	DtpAMPM
)

// TDateTimeParts SET:TDateTimePart
type TDateTimeParts = TSet

// TDateTimePickerOption ENUM
type TDateTimePickerOption int32

const (
	DtpoDoChangeOnSetDateTime = iota + 0
	DtpoEnabledIfUnchecked
	DtpoAutoCheck
	DtpoFlatButton
)

// TDateTimePickerOptions SET:TDateTimePickerOption
type TDateTimePickerOptions = TSet

// TImageOrientation ENUM
type TImageOrientation int32

const (
	ioHorizontal = iota + 0
	ioVertical
)

// TLayoutAdjustmentPolicy ENUM
type TLayoutAdjustmentPolicy int32

const (
	LapDefault                              = iota + 0 // widgetset dependent
	LapFixedLayout                                     // A fixed absolute layout in all platforms
	LapAutoAdjustWithoutHorizontalScrolling            // Smartphone platforms use this one,
	// LapAutoAdjustForDPI the x-axis is stretched to fill the screen and the y is scaled to fit the DPI
	LapAutoAdjustForDPI // For desktops using High DPI, scale x and y to fit the DPI
)

// THitTest ENUM
type THitTest int32

const (
	HtAbove = iota + 0
	HtBelow
	HtNowhere
	HtOnItem
	HtOnButton
	HtOnIcon
	HtOnIndent
	HtOnLabel
	HtOnRight
	HtOnStateIcon
	HtToLeft
	HtToRight
)

// THitTests SET:THitTest
type THitTests = TSet

// TListItemState ENUM
type TListItemState int32

const (
	LisCut = iota + 0
	LisDropTarget
	LisFocused
	LisSelected
)

// TListItemStates SET:TListItemState
type TListItemStates = TSet

// TPredefinedClipboardFormat ENUM
type TPredefinedClipboardFormat int32

const (
	PcfText = iota + 0
	PcfBitmap
	PcfPixmap
	PcfIcon
	PcfPicture
	PcfMetaFilePict
	PcfObject
	PcfComponent
	PcfCustomData
)
