// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Compat struct {
	// Use Simplified Rules For Table Border Conflicts
	UseSingleBorderforContiguousCells *CT_OnOff
	// Fit To Expanded Width When Performing Full Justification
	WpJustification *CT_OnOff
	// Do Not Create Custom Tab Stop for Hanging Indent
	NoTabHangInd *CT_OnOff
	// Do Not Add Leading Between Lines of Text
	NoLeading *CT_OnOff
	// Add Additional Space Below Baseline For Underlined East Asian Text
	SpaceForUL *CT_OnOff
	// Do Not Balance Text Columns within a Section
	NoColumnBalance *CT_OnOff
	// Balance Single Byte and Double Byte Characters
	BalanceSingleByteDoubleByteWidth *CT_OnOff
	// Do Not Center Content on Lines With Exact Line Height
	NoExtraLineSpacing *CT_OnOff
	// Display Backslash As Yen Sign
	DoNotLeaveBackslashAlone *CT_OnOff
	// Underline All Trailing Spaces
	UlTrailSpace *CT_OnOff
	// Don't Justify Lines Ending in Soft Line Break
	DoNotExpandShiftReturn *CT_OnOff
	// Only Expand/Condense Text By Whole Points
	SpacingInWholePoints *CT_OnOff
	// Ignore Compression of Full-Width Punctuation Ending a Line
	LineWrapLikeWord6 *CT_OnOff
	// Print Body Text before Header/Footer Contents
	PrintBodyTextBeforeHeader *CT_OnOff
	// Print Colors as Black And White without Dithering
	PrintColBlack *CT_OnOff
	// Use Specific Space Width
	WpSpaceWidth *CT_OnOff
	// Display Page/Column Breaks Present in Frames
	ShowBreaksInFrames *CT_OnOff
	// Require Exact Size During Font Substitution
	SubFontBySize *CT_OnOff
	// Ignore Exact Line Height for Last Line on Page
	SuppressBottomSpacing *CT_OnOff
	// Ignore Minimum and Exact Line Height for First Line on Page
	SuppressTopSpacing *CT_OnOff
	// Ignore Minimum Line Height for First Line on Page
	SuppressSpacingAtTopOfPage *CT_OnOff
	// Use Static Text Leading
	SuppressTopSpacingWP *CT_OnOff
	// Do Not Use Space Before On First Line After a Page Break
	SuppressSpBfAfterPgBrk *CT_OnOff
	// Swap Paragraph Borders on Odd Numbered Pages
	SwapBordersFacingPages *CT_OnOff
	// Treat Backslash Quotation Delimiter as Two Quotation Marks
	ConvMailMergeEsc *CT_OnOff
	// Use Truncated Integer Division For Font Calculation
	TruncateFontHeightsLikeWP6 *CT_OnOff
	// Use Specific Small Caps Algorithm
	MwSmallCaps *CT_OnOff
	// Use Printer Metrics To Display Documents
	UsePrinterMetrics *CT_OnOff
	// Do Not Suppress Paragraph Borders Next To Frames
	DoNotSuppressParagraphBorders *CT_OnOff
	// Line Wrap Trailing Spaces
	WrapTrailSpaces *CT_OnOff
	// Ignore Page Break from Continuous Section Break
	FootnoteLayoutLikeWW8 *CT_OnOff
	// Ignore Text Wrapping around Objects at Bottom of Page
	ShapeLayoutLikeWW8 *CT_OnOff
	// Align Table Rows Independently
	AlignTablesRowByRow *CT_OnOff
	// Ignore Width of Last Tab Stop When Aligning Paragraph If It Is Not Left Aligned
	ForgetLastTabAlignment *CT_OnOff
	// Add Document Grid Line Pitch To Lines in Table Cells
	AdjustLineHeightInTable *CT_OnOff
	// Incorrectly Adjust Text Spacing for Specific Unicode Ranges
	AutoSpaceLikeWord95 *CT_OnOff
	// Do Not Increase Line Height for Raised/Lowered Text
	NoSpaceRaiseLower *CT_OnOff
	// Use Fixed Paragraph Spacing for HTML Auto Setting
	DoNotUseHTMLParagraphAutoSpacing *CT_OnOff
	// Ignore Space Before Table When Deciding If Table Should Wrap Floating Object
	LayoutRawTableWidth *CT_OnOff
	// Allow Table Rows to Wrap Inline Objects Independently
	LayoutTableRowsApart *CT_OnOff
	// Use Incorrect Inter-Character Spacing Rules
	UseWord97LineBreakRules *CT_OnOff
	// Do Not Allow Floating Tables To Break Across Pages
	DoNotBreakWrappedTables *CT_OnOff
	// Do Not Snap to Document Grid in Table Cells with Objects
	DoNotSnapToGridInCell *CT_OnOff
	// Select Field When First or Last Character Is Selected
	SelectFldWithFirstOrLastChar *CT_OnOff
	// Use Legacy Ethiopic and Amharic Line Breaking Rules
	ApplyBreakingRules *CT_OnOff
	// Do Not Allow Hanging Punctuation With Character Grid
	DoNotWrapTextWithPunct *CT_OnOff
	// Do Not Compress Compressible Characters When Using Document Grid
	DoNotUseEastAsianBreakRules *CT_OnOff
	// Incorrectly Display Top Border of Conditional Columns
	UseWord2002TableStyleRules *CT_OnOff
	// Allow Tables to AutoFit Into Page Margins
	GrowAutofit *CT_OnOff
	// Do Not Bypass East Asian/Complex Script Layout Code
	UseFELayout *CT_OnOff
	// Do Not Automatically Apply List Paragraph Style To Bulleted/Numbered Text
	UseNormalStyleForList *CT_OnOff
	// Ignore Hanging Indent When Creating Tab Stop After Numbering
	DoNotUseIndentAsNumberingTabStop *CT_OnOff
	// Use Alternate Set of East Asian Line Breaking Rules
	UseAltKinsokuLineBreakRules *CT_OnOff
	// Allow Contextual Spacing of Paragraphs in Tables
	AllowSpaceOfSameStyleInTable *CT_OnOff
	// Do Not Ignore Floating Objects When Calculating Paragraph Indentation
	DoNotSuppressIndentation *CT_OnOff
	// Do Not AutoFit Tables To Fit Next To Wrapped Objects
	DoNotAutofitConstrainedTables *CT_OnOff
	// Allow Table Columns To Exceed Preferred Widths of Constituent Cells
	AutofitToFirstFixedWidthCell *CT_OnOff
	// Underline Following Character Following Numbering
	UnderlineTabInNumList *CT_OnOff
	// Always Use Fixed Width for Hangul Characters
	DisplayHangulFixedWidth *CT_OnOff
	// Always Move Paragraph Mark to Page after a Page Break
	SplitPgBreakAndParaMark *CT_OnOff
	// Don't Vertically Align Cells Containing Floating Objects
	DoNotVertAlignCellWithSp *CT_OnOff
	// Don't Break Table Rows Around Floating Tables
	DoNotBreakConstrainedForcedTable *CT_OnOff
	// Ignore Vertical Alignment in Textboxes
	DoNotVertAlignInTxbx *CT_OnOff
	// Use ANSI Kerning Pairs from Fonts
	UseAnsiKerningPairs *CT_OnOff
	// Use Cached Paragraph Information for Column Balancing
	CachedColBalance *CT_OnOff
	// Custom Compatibility Setting
	CompatSetting []*CT_CompatSetting
}

func NewCT_Compat() *CT_Compat {
	ret := &CT_Compat{}
	return ret
}

func (m *CT_Compat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.UseSingleBorderforContiguousCells != nil {
		seuseSingleBorderforContiguousCells := xml.StartElement{Name: xml.Name{Local: "w:useSingleBorderforContiguousCells"}}
		e.EncodeElement(m.UseSingleBorderforContiguousCells, seuseSingleBorderforContiguousCells)
	}
	if m.WpJustification != nil {
		sewpJustification := xml.StartElement{Name: xml.Name{Local: "w:wpJustification"}}
		e.EncodeElement(m.WpJustification, sewpJustification)
	}
	if m.NoTabHangInd != nil {
		senoTabHangInd := xml.StartElement{Name: xml.Name{Local: "w:noTabHangInd"}}
		e.EncodeElement(m.NoTabHangInd, senoTabHangInd)
	}
	if m.NoLeading != nil {
		senoLeading := xml.StartElement{Name: xml.Name{Local: "w:noLeading"}}
		e.EncodeElement(m.NoLeading, senoLeading)
	}
	if m.SpaceForUL != nil {
		sespaceForUL := xml.StartElement{Name: xml.Name{Local: "w:spaceForUL"}}
		e.EncodeElement(m.SpaceForUL, sespaceForUL)
	}
	if m.NoColumnBalance != nil {
		senoColumnBalance := xml.StartElement{Name: xml.Name{Local: "w:noColumnBalance"}}
		e.EncodeElement(m.NoColumnBalance, senoColumnBalance)
	}
	if m.BalanceSingleByteDoubleByteWidth != nil {
		sebalanceSingleByteDoubleByteWidth := xml.StartElement{Name: xml.Name{Local: "w:balanceSingleByteDoubleByteWidth"}}
		e.EncodeElement(m.BalanceSingleByteDoubleByteWidth, sebalanceSingleByteDoubleByteWidth)
	}
	if m.NoExtraLineSpacing != nil {
		senoExtraLineSpacing := xml.StartElement{Name: xml.Name{Local: "w:noExtraLineSpacing"}}
		e.EncodeElement(m.NoExtraLineSpacing, senoExtraLineSpacing)
	}
	if m.DoNotLeaveBackslashAlone != nil {
		sedoNotLeaveBackslashAlone := xml.StartElement{Name: xml.Name{Local: "w:doNotLeaveBackslashAlone"}}
		e.EncodeElement(m.DoNotLeaveBackslashAlone, sedoNotLeaveBackslashAlone)
	}
	if m.UlTrailSpace != nil {
		seulTrailSpace := xml.StartElement{Name: xml.Name{Local: "w:ulTrailSpace"}}
		e.EncodeElement(m.UlTrailSpace, seulTrailSpace)
	}
	if m.DoNotExpandShiftReturn != nil {
		sedoNotExpandShiftReturn := xml.StartElement{Name: xml.Name{Local: "w:doNotExpandShiftReturn"}}
		e.EncodeElement(m.DoNotExpandShiftReturn, sedoNotExpandShiftReturn)
	}
	if m.SpacingInWholePoints != nil {
		sespacingInWholePoints := xml.StartElement{Name: xml.Name{Local: "w:spacingInWholePoints"}}
		e.EncodeElement(m.SpacingInWholePoints, sespacingInWholePoints)
	}
	if m.LineWrapLikeWord6 != nil {
		selineWrapLikeWord6 := xml.StartElement{Name: xml.Name{Local: "w:lineWrapLikeWord6"}}
		e.EncodeElement(m.LineWrapLikeWord6, selineWrapLikeWord6)
	}
	if m.PrintBodyTextBeforeHeader != nil {
		seprintBodyTextBeforeHeader := xml.StartElement{Name: xml.Name{Local: "w:printBodyTextBeforeHeader"}}
		e.EncodeElement(m.PrintBodyTextBeforeHeader, seprintBodyTextBeforeHeader)
	}
	if m.PrintColBlack != nil {
		seprintColBlack := xml.StartElement{Name: xml.Name{Local: "w:printColBlack"}}
		e.EncodeElement(m.PrintColBlack, seprintColBlack)
	}
	if m.WpSpaceWidth != nil {
		sewpSpaceWidth := xml.StartElement{Name: xml.Name{Local: "w:wpSpaceWidth"}}
		e.EncodeElement(m.WpSpaceWidth, sewpSpaceWidth)
	}
	if m.ShowBreaksInFrames != nil {
		seshowBreaksInFrames := xml.StartElement{Name: xml.Name{Local: "w:showBreaksInFrames"}}
		e.EncodeElement(m.ShowBreaksInFrames, seshowBreaksInFrames)
	}
	if m.SubFontBySize != nil {
		sesubFontBySize := xml.StartElement{Name: xml.Name{Local: "w:subFontBySize"}}
		e.EncodeElement(m.SubFontBySize, sesubFontBySize)
	}
	if m.SuppressBottomSpacing != nil {
		sesuppressBottomSpacing := xml.StartElement{Name: xml.Name{Local: "w:suppressBottomSpacing"}}
		e.EncodeElement(m.SuppressBottomSpacing, sesuppressBottomSpacing)
	}
	if m.SuppressTopSpacing != nil {
		sesuppressTopSpacing := xml.StartElement{Name: xml.Name{Local: "w:suppressTopSpacing"}}
		e.EncodeElement(m.SuppressTopSpacing, sesuppressTopSpacing)
	}
	if m.SuppressSpacingAtTopOfPage != nil {
		sesuppressSpacingAtTopOfPage := xml.StartElement{Name: xml.Name{Local: "w:suppressSpacingAtTopOfPage"}}
		e.EncodeElement(m.SuppressSpacingAtTopOfPage, sesuppressSpacingAtTopOfPage)
	}
	if m.SuppressTopSpacingWP != nil {
		sesuppressTopSpacingWP := xml.StartElement{Name: xml.Name{Local: "w:suppressTopSpacingWP"}}
		e.EncodeElement(m.SuppressTopSpacingWP, sesuppressTopSpacingWP)
	}
	if m.SuppressSpBfAfterPgBrk != nil {
		sesuppressSpBfAfterPgBrk := xml.StartElement{Name: xml.Name{Local: "w:suppressSpBfAfterPgBrk"}}
		e.EncodeElement(m.SuppressSpBfAfterPgBrk, sesuppressSpBfAfterPgBrk)
	}
	if m.SwapBordersFacingPages != nil {
		seswapBordersFacingPages := xml.StartElement{Name: xml.Name{Local: "w:swapBordersFacingPages"}}
		e.EncodeElement(m.SwapBordersFacingPages, seswapBordersFacingPages)
	}
	if m.ConvMailMergeEsc != nil {
		seconvMailMergeEsc := xml.StartElement{Name: xml.Name{Local: "w:convMailMergeEsc"}}
		e.EncodeElement(m.ConvMailMergeEsc, seconvMailMergeEsc)
	}
	if m.TruncateFontHeightsLikeWP6 != nil {
		setruncateFontHeightsLikeWP6 := xml.StartElement{Name: xml.Name{Local: "w:truncateFontHeightsLikeWP6"}}
		e.EncodeElement(m.TruncateFontHeightsLikeWP6, setruncateFontHeightsLikeWP6)
	}
	if m.MwSmallCaps != nil {
		semwSmallCaps := xml.StartElement{Name: xml.Name{Local: "w:mwSmallCaps"}}
		e.EncodeElement(m.MwSmallCaps, semwSmallCaps)
	}
	if m.UsePrinterMetrics != nil {
		seusePrinterMetrics := xml.StartElement{Name: xml.Name{Local: "w:usePrinterMetrics"}}
		e.EncodeElement(m.UsePrinterMetrics, seusePrinterMetrics)
	}
	if m.DoNotSuppressParagraphBorders != nil {
		sedoNotSuppressParagraphBorders := xml.StartElement{Name: xml.Name{Local: "w:doNotSuppressParagraphBorders"}}
		e.EncodeElement(m.DoNotSuppressParagraphBorders, sedoNotSuppressParagraphBorders)
	}
	if m.WrapTrailSpaces != nil {
		sewrapTrailSpaces := xml.StartElement{Name: xml.Name{Local: "w:wrapTrailSpaces"}}
		e.EncodeElement(m.WrapTrailSpaces, sewrapTrailSpaces)
	}
	if m.FootnoteLayoutLikeWW8 != nil {
		sefootnoteLayoutLikeWW8 := xml.StartElement{Name: xml.Name{Local: "w:footnoteLayoutLikeWW8"}}
		e.EncodeElement(m.FootnoteLayoutLikeWW8, sefootnoteLayoutLikeWW8)
	}
	if m.ShapeLayoutLikeWW8 != nil {
		seshapeLayoutLikeWW8 := xml.StartElement{Name: xml.Name{Local: "w:shapeLayoutLikeWW8"}}
		e.EncodeElement(m.ShapeLayoutLikeWW8, seshapeLayoutLikeWW8)
	}
	if m.AlignTablesRowByRow != nil {
		sealignTablesRowByRow := xml.StartElement{Name: xml.Name{Local: "w:alignTablesRowByRow"}}
		e.EncodeElement(m.AlignTablesRowByRow, sealignTablesRowByRow)
	}
	if m.ForgetLastTabAlignment != nil {
		seforgetLastTabAlignment := xml.StartElement{Name: xml.Name{Local: "w:forgetLastTabAlignment"}}
		e.EncodeElement(m.ForgetLastTabAlignment, seforgetLastTabAlignment)
	}
	if m.AdjustLineHeightInTable != nil {
		seadjustLineHeightInTable := xml.StartElement{Name: xml.Name{Local: "w:adjustLineHeightInTable"}}
		e.EncodeElement(m.AdjustLineHeightInTable, seadjustLineHeightInTable)
	}
	if m.AutoSpaceLikeWord95 != nil {
		seautoSpaceLikeWord95 := xml.StartElement{Name: xml.Name{Local: "w:autoSpaceLikeWord95"}}
		e.EncodeElement(m.AutoSpaceLikeWord95, seautoSpaceLikeWord95)
	}
	if m.NoSpaceRaiseLower != nil {
		senoSpaceRaiseLower := xml.StartElement{Name: xml.Name{Local: "w:noSpaceRaiseLower"}}
		e.EncodeElement(m.NoSpaceRaiseLower, senoSpaceRaiseLower)
	}
	if m.DoNotUseHTMLParagraphAutoSpacing != nil {
		sedoNotUseHTMLParagraphAutoSpacing := xml.StartElement{Name: xml.Name{Local: "w:doNotUseHTMLParagraphAutoSpacing"}}
		e.EncodeElement(m.DoNotUseHTMLParagraphAutoSpacing, sedoNotUseHTMLParagraphAutoSpacing)
	}
	if m.LayoutRawTableWidth != nil {
		selayoutRawTableWidth := xml.StartElement{Name: xml.Name{Local: "w:layoutRawTableWidth"}}
		e.EncodeElement(m.LayoutRawTableWidth, selayoutRawTableWidth)
	}
	if m.LayoutTableRowsApart != nil {
		selayoutTableRowsApart := xml.StartElement{Name: xml.Name{Local: "w:layoutTableRowsApart"}}
		e.EncodeElement(m.LayoutTableRowsApart, selayoutTableRowsApart)
	}
	if m.UseWord97LineBreakRules != nil {
		seuseWord97LineBreakRules := xml.StartElement{Name: xml.Name{Local: "w:useWord97LineBreakRules"}}
		e.EncodeElement(m.UseWord97LineBreakRules, seuseWord97LineBreakRules)
	}
	if m.DoNotBreakWrappedTables != nil {
		sedoNotBreakWrappedTables := xml.StartElement{Name: xml.Name{Local: "w:doNotBreakWrappedTables"}}
		e.EncodeElement(m.DoNotBreakWrappedTables, sedoNotBreakWrappedTables)
	}
	if m.DoNotSnapToGridInCell != nil {
		sedoNotSnapToGridInCell := xml.StartElement{Name: xml.Name{Local: "w:doNotSnapToGridInCell"}}
		e.EncodeElement(m.DoNotSnapToGridInCell, sedoNotSnapToGridInCell)
	}
	if m.SelectFldWithFirstOrLastChar != nil {
		seselectFldWithFirstOrLastChar := xml.StartElement{Name: xml.Name{Local: "w:selectFldWithFirstOrLastChar"}}
		e.EncodeElement(m.SelectFldWithFirstOrLastChar, seselectFldWithFirstOrLastChar)
	}
	if m.ApplyBreakingRules != nil {
		seapplyBreakingRules := xml.StartElement{Name: xml.Name{Local: "w:applyBreakingRules"}}
		e.EncodeElement(m.ApplyBreakingRules, seapplyBreakingRules)
	}
	if m.DoNotWrapTextWithPunct != nil {
		sedoNotWrapTextWithPunct := xml.StartElement{Name: xml.Name{Local: "w:doNotWrapTextWithPunct"}}
		e.EncodeElement(m.DoNotWrapTextWithPunct, sedoNotWrapTextWithPunct)
	}
	if m.DoNotUseEastAsianBreakRules != nil {
		sedoNotUseEastAsianBreakRules := xml.StartElement{Name: xml.Name{Local: "w:doNotUseEastAsianBreakRules"}}
		e.EncodeElement(m.DoNotUseEastAsianBreakRules, sedoNotUseEastAsianBreakRules)
	}
	if m.UseWord2002TableStyleRules != nil {
		seuseWord2002TableStyleRules := xml.StartElement{Name: xml.Name{Local: "w:useWord2002TableStyleRules"}}
		e.EncodeElement(m.UseWord2002TableStyleRules, seuseWord2002TableStyleRules)
	}
	if m.GrowAutofit != nil {
		segrowAutofit := xml.StartElement{Name: xml.Name{Local: "w:growAutofit"}}
		e.EncodeElement(m.GrowAutofit, segrowAutofit)
	}
	if m.UseFELayout != nil {
		seuseFELayout := xml.StartElement{Name: xml.Name{Local: "w:useFELayout"}}
		e.EncodeElement(m.UseFELayout, seuseFELayout)
	}
	if m.UseNormalStyleForList != nil {
		seuseNormalStyleForList := xml.StartElement{Name: xml.Name{Local: "w:useNormalStyleForList"}}
		e.EncodeElement(m.UseNormalStyleForList, seuseNormalStyleForList)
	}
	if m.DoNotUseIndentAsNumberingTabStop != nil {
		sedoNotUseIndentAsNumberingTabStop := xml.StartElement{Name: xml.Name{Local: "w:doNotUseIndentAsNumberingTabStop"}}
		e.EncodeElement(m.DoNotUseIndentAsNumberingTabStop, sedoNotUseIndentAsNumberingTabStop)
	}
	if m.UseAltKinsokuLineBreakRules != nil {
		seuseAltKinsokuLineBreakRules := xml.StartElement{Name: xml.Name{Local: "w:useAltKinsokuLineBreakRules"}}
		e.EncodeElement(m.UseAltKinsokuLineBreakRules, seuseAltKinsokuLineBreakRules)
	}
	if m.AllowSpaceOfSameStyleInTable != nil {
		seallowSpaceOfSameStyleInTable := xml.StartElement{Name: xml.Name{Local: "w:allowSpaceOfSameStyleInTable"}}
		e.EncodeElement(m.AllowSpaceOfSameStyleInTable, seallowSpaceOfSameStyleInTable)
	}
	if m.DoNotSuppressIndentation != nil {
		sedoNotSuppressIndentation := xml.StartElement{Name: xml.Name{Local: "w:doNotSuppressIndentation"}}
		e.EncodeElement(m.DoNotSuppressIndentation, sedoNotSuppressIndentation)
	}
	if m.DoNotAutofitConstrainedTables != nil {
		sedoNotAutofitConstrainedTables := xml.StartElement{Name: xml.Name{Local: "w:doNotAutofitConstrainedTables"}}
		e.EncodeElement(m.DoNotAutofitConstrainedTables, sedoNotAutofitConstrainedTables)
	}
	if m.AutofitToFirstFixedWidthCell != nil {
		seautofitToFirstFixedWidthCell := xml.StartElement{Name: xml.Name{Local: "w:autofitToFirstFixedWidthCell"}}
		e.EncodeElement(m.AutofitToFirstFixedWidthCell, seautofitToFirstFixedWidthCell)
	}
	if m.UnderlineTabInNumList != nil {
		seunderlineTabInNumList := xml.StartElement{Name: xml.Name{Local: "w:underlineTabInNumList"}}
		e.EncodeElement(m.UnderlineTabInNumList, seunderlineTabInNumList)
	}
	if m.DisplayHangulFixedWidth != nil {
		sedisplayHangulFixedWidth := xml.StartElement{Name: xml.Name{Local: "w:displayHangulFixedWidth"}}
		e.EncodeElement(m.DisplayHangulFixedWidth, sedisplayHangulFixedWidth)
	}
	if m.SplitPgBreakAndParaMark != nil {
		sesplitPgBreakAndParaMark := xml.StartElement{Name: xml.Name{Local: "w:splitPgBreakAndParaMark"}}
		e.EncodeElement(m.SplitPgBreakAndParaMark, sesplitPgBreakAndParaMark)
	}
	if m.DoNotVertAlignCellWithSp != nil {
		sedoNotVertAlignCellWithSp := xml.StartElement{Name: xml.Name{Local: "w:doNotVertAlignCellWithSp"}}
		e.EncodeElement(m.DoNotVertAlignCellWithSp, sedoNotVertAlignCellWithSp)
	}
	if m.DoNotBreakConstrainedForcedTable != nil {
		sedoNotBreakConstrainedForcedTable := xml.StartElement{Name: xml.Name{Local: "w:doNotBreakConstrainedForcedTable"}}
		e.EncodeElement(m.DoNotBreakConstrainedForcedTable, sedoNotBreakConstrainedForcedTable)
	}
	if m.DoNotVertAlignInTxbx != nil {
		sedoNotVertAlignInTxbx := xml.StartElement{Name: xml.Name{Local: "w:doNotVertAlignInTxbx"}}
		e.EncodeElement(m.DoNotVertAlignInTxbx, sedoNotVertAlignInTxbx)
	}
	if m.UseAnsiKerningPairs != nil {
		seuseAnsiKerningPairs := xml.StartElement{Name: xml.Name{Local: "w:useAnsiKerningPairs"}}
		e.EncodeElement(m.UseAnsiKerningPairs, seuseAnsiKerningPairs)
	}
	if m.CachedColBalance != nil {
		secachedColBalance := xml.StartElement{Name: xml.Name{Local: "w:cachedColBalance"}}
		e.EncodeElement(m.CachedColBalance, secachedColBalance)
	}
	if m.CompatSetting != nil {
		secompatSetting := xml.StartElement{Name: xml.Name{Local: "w:compatSetting"}}
		e.EncodeElement(m.CompatSetting, secompatSetting)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Compat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Compat:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "useSingleBorderforContiguousCells":
				m.UseSingleBorderforContiguousCells = NewCT_OnOff()
				if err := d.DecodeElement(m.UseSingleBorderforContiguousCells, &el); err != nil {
					return err
				}
			case "wpJustification":
				m.WpJustification = NewCT_OnOff()
				if err := d.DecodeElement(m.WpJustification, &el); err != nil {
					return err
				}
			case "noTabHangInd":
				m.NoTabHangInd = NewCT_OnOff()
				if err := d.DecodeElement(m.NoTabHangInd, &el); err != nil {
					return err
				}
			case "noLeading":
				m.NoLeading = NewCT_OnOff()
				if err := d.DecodeElement(m.NoLeading, &el); err != nil {
					return err
				}
			case "spaceForUL":
				m.SpaceForUL = NewCT_OnOff()
				if err := d.DecodeElement(m.SpaceForUL, &el); err != nil {
					return err
				}
			case "noColumnBalance":
				m.NoColumnBalance = NewCT_OnOff()
				if err := d.DecodeElement(m.NoColumnBalance, &el); err != nil {
					return err
				}
			case "balanceSingleByteDoubleByteWidth":
				m.BalanceSingleByteDoubleByteWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.BalanceSingleByteDoubleByteWidth, &el); err != nil {
					return err
				}
			case "noExtraLineSpacing":
				m.NoExtraLineSpacing = NewCT_OnOff()
				if err := d.DecodeElement(m.NoExtraLineSpacing, &el); err != nil {
					return err
				}
			case "doNotLeaveBackslashAlone":
				m.DoNotLeaveBackslashAlone = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotLeaveBackslashAlone, &el); err != nil {
					return err
				}
			case "ulTrailSpace":
				m.UlTrailSpace = NewCT_OnOff()
				if err := d.DecodeElement(m.UlTrailSpace, &el); err != nil {
					return err
				}
			case "doNotExpandShiftReturn":
				m.DoNotExpandShiftReturn = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotExpandShiftReturn, &el); err != nil {
					return err
				}
			case "spacingInWholePoints":
				m.SpacingInWholePoints = NewCT_OnOff()
				if err := d.DecodeElement(m.SpacingInWholePoints, &el); err != nil {
					return err
				}
			case "lineWrapLikeWord6":
				m.LineWrapLikeWord6 = NewCT_OnOff()
				if err := d.DecodeElement(m.LineWrapLikeWord6, &el); err != nil {
					return err
				}
			case "printBodyTextBeforeHeader":
				m.PrintBodyTextBeforeHeader = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintBodyTextBeforeHeader, &el); err != nil {
					return err
				}
			case "printColBlack":
				m.PrintColBlack = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintColBlack, &el); err != nil {
					return err
				}
			case "wpSpaceWidth":
				m.WpSpaceWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.WpSpaceWidth, &el); err != nil {
					return err
				}
			case "showBreaksInFrames":
				m.ShowBreaksInFrames = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowBreaksInFrames, &el); err != nil {
					return err
				}
			case "subFontBySize":
				m.SubFontBySize = NewCT_OnOff()
				if err := d.DecodeElement(m.SubFontBySize, &el); err != nil {
					return err
				}
			case "suppressBottomSpacing":
				m.SuppressBottomSpacing = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressBottomSpacing, &el); err != nil {
					return err
				}
			case "suppressTopSpacing":
				m.SuppressTopSpacing = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressTopSpacing, &el); err != nil {
					return err
				}
			case "suppressSpacingAtTopOfPage":
				m.SuppressSpacingAtTopOfPage = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressSpacingAtTopOfPage, &el); err != nil {
					return err
				}
			case "suppressTopSpacingWP":
				m.SuppressTopSpacingWP = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressTopSpacingWP, &el); err != nil {
					return err
				}
			case "suppressSpBfAfterPgBrk":
				m.SuppressSpBfAfterPgBrk = NewCT_OnOff()
				if err := d.DecodeElement(m.SuppressSpBfAfterPgBrk, &el); err != nil {
					return err
				}
			case "swapBordersFacingPages":
				m.SwapBordersFacingPages = NewCT_OnOff()
				if err := d.DecodeElement(m.SwapBordersFacingPages, &el); err != nil {
					return err
				}
			case "convMailMergeEsc":
				m.ConvMailMergeEsc = NewCT_OnOff()
				if err := d.DecodeElement(m.ConvMailMergeEsc, &el); err != nil {
					return err
				}
			case "truncateFontHeightsLikeWP6":
				m.TruncateFontHeightsLikeWP6 = NewCT_OnOff()
				if err := d.DecodeElement(m.TruncateFontHeightsLikeWP6, &el); err != nil {
					return err
				}
			case "mwSmallCaps":
				m.MwSmallCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.MwSmallCaps, &el); err != nil {
					return err
				}
			case "usePrinterMetrics":
				m.UsePrinterMetrics = NewCT_OnOff()
				if err := d.DecodeElement(m.UsePrinterMetrics, &el); err != nil {
					return err
				}
			case "doNotSuppressParagraphBorders":
				m.DoNotSuppressParagraphBorders = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSuppressParagraphBorders, &el); err != nil {
					return err
				}
			case "wrapTrailSpaces":
				m.WrapTrailSpaces = NewCT_OnOff()
				if err := d.DecodeElement(m.WrapTrailSpaces, &el); err != nil {
					return err
				}
			case "footnoteLayoutLikeWW8":
				m.FootnoteLayoutLikeWW8 = NewCT_OnOff()
				if err := d.DecodeElement(m.FootnoteLayoutLikeWW8, &el); err != nil {
					return err
				}
			case "shapeLayoutLikeWW8":
				m.ShapeLayoutLikeWW8 = NewCT_OnOff()
				if err := d.DecodeElement(m.ShapeLayoutLikeWW8, &el); err != nil {
					return err
				}
			case "alignTablesRowByRow":
				m.AlignTablesRowByRow = NewCT_OnOff()
				if err := d.DecodeElement(m.AlignTablesRowByRow, &el); err != nil {
					return err
				}
			case "forgetLastTabAlignment":
				m.ForgetLastTabAlignment = NewCT_OnOff()
				if err := d.DecodeElement(m.ForgetLastTabAlignment, &el); err != nil {
					return err
				}
			case "adjustLineHeightInTable":
				m.AdjustLineHeightInTable = NewCT_OnOff()
				if err := d.DecodeElement(m.AdjustLineHeightInTable, &el); err != nil {
					return err
				}
			case "autoSpaceLikeWord95":
				m.AutoSpaceLikeWord95 = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoSpaceLikeWord95, &el); err != nil {
					return err
				}
			case "noSpaceRaiseLower":
				m.NoSpaceRaiseLower = NewCT_OnOff()
				if err := d.DecodeElement(m.NoSpaceRaiseLower, &el); err != nil {
					return err
				}
			case "doNotUseHTMLParagraphAutoSpacing":
				m.DoNotUseHTMLParagraphAutoSpacing = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseHTMLParagraphAutoSpacing, &el); err != nil {
					return err
				}
			case "layoutRawTableWidth":
				m.LayoutRawTableWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.LayoutRawTableWidth, &el); err != nil {
					return err
				}
			case "layoutTableRowsApart":
				m.LayoutTableRowsApart = NewCT_OnOff()
				if err := d.DecodeElement(m.LayoutTableRowsApart, &el); err != nil {
					return err
				}
			case "useWord97LineBreakRules":
				m.UseWord97LineBreakRules = NewCT_OnOff()
				if err := d.DecodeElement(m.UseWord97LineBreakRules, &el); err != nil {
					return err
				}
			case "doNotBreakWrappedTables":
				m.DoNotBreakWrappedTables = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotBreakWrappedTables, &el); err != nil {
					return err
				}
			case "doNotSnapToGridInCell":
				m.DoNotSnapToGridInCell = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSnapToGridInCell, &el); err != nil {
					return err
				}
			case "selectFldWithFirstOrLastChar":
				m.SelectFldWithFirstOrLastChar = NewCT_OnOff()
				if err := d.DecodeElement(m.SelectFldWithFirstOrLastChar, &el); err != nil {
					return err
				}
			case "applyBreakingRules":
				m.ApplyBreakingRules = NewCT_OnOff()
				if err := d.DecodeElement(m.ApplyBreakingRules, &el); err != nil {
					return err
				}
			case "doNotWrapTextWithPunct":
				m.DoNotWrapTextWithPunct = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotWrapTextWithPunct, &el); err != nil {
					return err
				}
			case "doNotUseEastAsianBreakRules":
				m.DoNotUseEastAsianBreakRules = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseEastAsianBreakRules, &el); err != nil {
					return err
				}
			case "useWord2002TableStyleRules":
				m.UseWord2002TableStyleRules = NewCT_OnOff()
				if err := d.DecodeElement(m.UseWord2002TableStyleRules, &el); err != nil {
					return err
				}
			case "growAutofit":
				m.GrowAutofit = NewCT_OnOff()
				if err := d.DecodeElement(m.GrowAutofit, &el); err != nil {
					return err
				}
			case "useFELayout":
				m.UseFELayout = NewCT_OnOff()
				if err := d.DecodeElement(m.UseFELayout, &el); err != nil {
					return err
				}
			case "useNormalStyleForList":
				m.UseNormalStyleForList = NewCT_OnOff()
				if err := d.DecodeElement(m.UseNormalStyleForList, &el); err != nil {
					return err
				}
			case "doNotUseIndentAsNumberingTabStop":
				m.DoNotUseIndentAsNumberingTabStop = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseIndentAsNumberingTabStop, &el); err != nil {
					return err
				}
			case "useAltKinsokuLineBreakRules":
				m.UseAltKinsokuLineBreakRules = NewCT_OnOff()
				if err := d.DecodeElement(m.UseAltKinsokuLineBreakRules, &el); err != nil {
					return err
				}
			case "allowSpaceOfSameStyleInTable":
				m.AllowSpaceOfSameStyleInTable = NewCT_OnOff()
				if err := d.DecodeElement(m.AllowSpaceOfSameStyleInTable, &el); err != nil {
					return err
				}
			case "doNotSuppressIndentation":
				m.DoNotSuppressIndentation = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSuppressIndentation, &el); err != nil {
					return err
				}
			case "doNotAutofitConstrainedTables":
				m.DoNotAutofitConstrainedTables = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotAutofitConstrainedTables, &el); err != nil {
					return err
				}
			case "autofitToFirstFixedWidthCell":
				m.AutofitToFirstFixedWidthCell = NewCT_OnOff()
				if err := d.DecodeElement(m.AutofitToFirstFixedWidthCell, &el); err != nil {
					return err
				}
			case "underlineTabInNumList":
				m.UnderlineTabInNumList = NewCT_OnOff()
				if err := d.DecodeElement(m.UnderlineTabInNumList, &el); err != nil {
					return err
				}
			case "displayHangulFixedWidth":
				m.DisplayHangulFixedWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.DisplayHangulFixedWidth, &el); err != nil {
					return err
				}
			case "splitPgBreakAndParaMark":
				m.SplitPgBreakAndParaMark = NewCT_OnOff()
				if err := d.DecodeElement(m.SplitPgBreakAndParaMark, &el); err != nil {
					return err
				}
			case "doNotVertAlignCellWithSp":
				m.DoNotVertAlignCellWithSp = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotVertAlignCellWithSp, &el); err != nil {
					return err
				}
			case "doNotBreakConstrainedForcedTable":
				m.DoNotBreakConstrainedForcedTable = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotBreakConstrainedForcedTable, &el); err != nil {
					return err
				}
			case "doNotVertAlignInTxbx":
				m.DoNotVertAlignInTxbx = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotVertAlignInTxbx, &el); err != nil {
					return err
				}
			case "useAnsiKerningPairs":
				m.UseAnsiKerningPairs = NewCT_OnOff()
				if err := d.DecodeElement(m.UseAnsiKerningPairs, &el); err != nil {
					return err
				}
			case "cachedColBalance":
				m.CachedColBalance = NewCT_OnOff()
				if err := d.DecodeElement(m.CachedColBalance, &el); err != nil {
					return err
				}
			case "compatSetting":
				tmp := NewCT_CompatSetting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CompatSetting = append(m.CompatSetting, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Compat
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Compat and its children
func (m *CT_Compat) Validate() error {
	return m.ValidateWithPath("CT_Compat")
}

// ValidateWithPath validates the CT_Compat and its children, prefixing error messages with path
func (m *CT_Compat) ValidateWithPath(path string) error {
	if m.UseSingleBorderforContiguousCells != nil {
		if err := m.UseSingleBorderforContiguousCells.ValidateWithPath(path + "/UseSingleBorderforContiguousCells"); err != nil {
			return err
		}
	}
	if m.WpJustification != nil {
		if err := m.WpJustification.ValidateWithPath(path + "/WpJustification"); err != nil {
			return err
		}
	}
	if m.NoTabHangInd != nil {
		if err := m.NoTabHangInd.ValidateWithPath(path + "/NoTabHangInd"); err != nil {
			return err
		}
	}
	if m.NoLeading != nil {
		if err := m.NoLeading.ValidateWithPath(path + "/NoLeading"); err != nil {
			return err
		}
	}
	if m.SpaceForUL != nil {
		if err := m.SpaceForUL.ValidateWithPath(path + "/SpaceForUL"); err != nil {
			return err
		}
	}
	if m.NoColumnBalance != nil {
		if err := m.NoColumnBalance.ValidateWithPath(path + "/NoColumnBalance"); err != nil {
			return err
		}
	}
	if m.BalanceSingleByteDoubleByteWidth != nil {
		if err := m.BalanceSingleByteDoubleByteWidth.ValidateWithPath(path + "/BalanceSingleByteDoubleByteWidth"); err != nil {
			return err
		}
	}
	if m.NoExtraLineSpacing != nil {
		if err := m.NoExtraLineSpacing.ValidateWithPath(path + "/NoExtraLineSpacing"); err != nil {
			return err
		}
	}
	if m.DoNotLeaveBackslashAlone != nil {
		if err := m.DoNotLeaveBackslashAlone.ValidateWithPath(path + "/DoNotLeaveBackslashAlone"); err != nil {
			return err
		}
	}
	if m.UlTrailSpace != nil {
		if err := m.UlTrailSpace.ValidateWithPath(path + "/UlTrailSpace"); err != nil {
			return err
		}
	}
	if m.DoNotExpandShiftReturn != nil {
		if err := m.DoNotExpandShiftReturn.ValidateWithPath(path + "/DoNotExpandShiftReturn"); err != nil {
			return err
		}
	}
	if m.SpacingInWholePoints != nil {
		if err := m.SpacingInWholePoints.ValidateWithPath(path + "/SpacingInWholePoints"); err != nil {
			return err
		}
	}
	if m.LineWrapLikeWord6 != nil {
		if err := m.LineWrapLikeWord6.ValidateWithPath(path + "/LineWrapLikeWord6"); err != nil {
			return err
		}
	}
	if m.PrintBodyTextBeforeHeader != nil {
		if err := m.PrintBodyTextBeforeHeader.ValidateWithPath(path + "/PrintBodyTextBeforeHeader"); err != nil {
			return err
		}
	}
	if m.PrintColBlack != nil {
		if err := m.PrintColBlack.ValidateWithPath(path + "/PrintColBlack"); err != nil {
			return err
		}
	}
	if m.WpSpaceWidth != nil {
		if err := m.WpSpaceWidth.ValidateWithPath(path + "/WpSpaceWidth"); err != nil {
			return err
		}
	}
	if m.ShowBreaksInFrames != nil {
		if err := m.ShowBreaksInFrames.ValidateWithPath(path + "/ShowBreaksInFrames"); err != nil {
			return err
		}
	}
	if m.SubFontBySize != nil {
		if err := m.SubFontBySize.ValidateWithPath(path + "/SubFontBySize"); err != nil {
			return err
		}
	}
	if m.SuppressBottomSpacing != nil {
		if err := m.SuppressBottomSpacing.ValidateWithPath(path + "/SuppressBottomSpacing"); err != nil {
			return err
		}
	}
	if m.SuppressTopSpacing != nil {
		if err := m.SuppressTopSpacing.ValidateWithPath(path + "/SuppressTopSpacing"); err != nil {
			return err
		}
	}
	if m.SuppressSpacingAtTopOfPage != nil {
		if err := m.SuppressSpacingAtTopOfPage.ValidateWithPath(path + "/SuppressSpacingAtTopOfPage"); err != nil {
			return err
		}
	}
	if m.SuppressTopSpacingWP != nil {
		if err := m.SuppressTopSpacingWP.ValidateWithPath(path + "/SuppressTopSpacingWP"); err != nil {
			return err
		}
	}
	if m.SuppressSpBfAfterPgBrk != nil {
		if err := m.SuppressSpBfAfterPgBrk.ValidateWithPath(path + "/SuppressSpBfAfterPgBrk"); err != nil {
			return err
		}
	}
	if m.SwapBordersFacingPages != nil {
		if err := m.SwapBordersFacingPages.ValidateWithPath(path + "/SwapBordersFacingPages"); err != nil {
			return err
		}
	}
	if m.ConvMailMergeEsc != nil {
		if err := m.ConvMailMergeEsc.ValidateWithPath(path + "/ConvMailMergeEsc"); err != nil {
			return err
		}
	}
	if m.TruncateFontHeightsLikeWP6 != nil {
		if err := m.TruncateFontHeightsLikeWP6.ValidateWithPath(path + "/TruncateFontHeightsLikeWP6"); err != nil {
			return err
		}
	}
	if m.MwSmallCaps != nil {
		if err := m.MwSmallCaps.ValidateWithPath(path + "/MwSmallCaps"); err != nil {
			return err
		}
	}
	if m.UsePrinterMetrics != nil {
		if err := m.UsePrinterMetrics.ValidateWithPath(path + "/UsePrinterMetrics"); err != nil {
			return err
		}
	}
	if m.DoNotSuppressParagraphBorders != nil {
		if err := m.DoNotSuppressParagraphBorders.ValidateWithPath(path + "/DoNotSuppressParagraphBorders"); err != nil {
			return err
		}
	}
	if m.WrapTrailSpaces != nil {
		if err := m.WrapTrailSpaces.ValidateWithPath(path + "/WrapTrailSpaces"); err != nil {
			return err
		}
	}
	if m.FootnoteLayoutLikeWW8 != nil {
		if err := m.FootnoteLayoutLikeWW8.ValidateWithPath(path + "/FootnoteLayoutLikeWW8"); err != nil {
			return err
		}
	}
	if m.ShapeLayoutLikeWW8 != nil {
		if err := m.ShapeLayoutLikeWW8.ValidateWithPath(path + "/ShapeLayoutLikeWW8"); err != nil {
			return err
		}
	}
	if m.AlignTablesRowByRow != nil {
		if err := m.AlignTablesRowByRow.ValidateWithPath(path + "/AlignTablesRowByRow"); err != nil {
			return err
		}
	}
	if m.ForgetLastTabAlignment != nil {
		if err := m.ForgetLastTabAlignment.ValidateWithPath(path + "/ForgetLastTabAlignment"); err != nil {
			return err
		}
	}
	if m.AdjustLineHeightInTable != nil {
		if err := m.AdjustLineHeightInTable.ValidateWithPath(path + "/AdjustLineHeightInTable"); err != nil {
			return err
		}
	}
	if m.AutoSpaceLikeWord95 != nil {
		if err := m.AutoSpaceLikeWord95.ValidateWithPath(path + "/AutoSpaceLikeWord95"); err != nil {
			return err
		}
	}
	if m.NoSpaceRaiseLower != nil {
		if err := m.NoSpaceRaiseLower.ValidateWithPath(path + "/NoSpaceRaiseLower"); err != nil {
			return err
		}
	}
	if m.DoNotUseHTMLParagraphAutoSpacing != nil {
		if err := m.DoNotUseHTMLParagraphAutoSpacing.ValidateWithPath(path + "/DoNotUseHTMLParagraphAutoSpacing"); err != nil {
			return err
		}
	}
	if m.LayoutRawTableWidth != nil {
		if err := m.LayoutRawTableWidth.ValidateWithPath(path + "/LayoutRawTableWidth"); err != nil {
			return err
		}
	}
	if m.LayoutTableRowsApart != nil {
		if err := m.LayoutTableRowsApart.ValidateWithPath(path + "/LayoutTableRowsApart"); err != nil {
			return err
		}
	}
	if m.UseWord97LineBreakRules != nil {
		if err := m.UseWord97LineBreakRules.ValidateWithPath(path + "/UseWord97LineBreakRules"); err != nil {
			return err
		}
	}
	if m.DoNotBreakWrappedTables != nil {
		if err := m.DoNotBreakWrappedTables.ValidateWithPath(path + "/DoNotBreakWrappedTables"); err != nil {
			return err
		}
	}
	if m.DoNotSnapToGridInCell != nil {
		if err := m.DoNotSnapToGridInCell.ValidateWithPath(path + "/DoNotSnapToGridInCell"); err != nil {
			return err
		}
	}
	if m.SelectFldWithFirstOrLastChar != nil {
		if err := m.SelectFldWithFirstOrLastChar.ValidateWithPath(path + "/SelectFldWithFirstOrLastChar"); err != nil {
			return err
		}
	}
	if m.ApplyBreakingRules != nil {
		if err := m.ApplyBreakingRules.ValidateWithPath(path + "/ApplyBreakingRules"); err != nil {
			return err
		}
	}
	if m.DoNotWrapTextWithPunct != nil {
		if err := m.DoNotWrapTextWithPunct.ValidateWithPath(path + "/DoNotWrapTextWithPunct"); err != nil {
			return err
		}
	}
	if m.DoNotUseEastAsianBreakRules != nil {
		if err := m.DoNotUseEastAsianBreakRules.ValidateWithPath(path + "/DoNotUseEastAsianBreakRules"); err != nil {
			return err
		}
	}
	if m.UseWord2002TableStyleRules != nil {
		if err := m.UseWord2002TableStyleRules.ValidateWithPath(path + "/UseWord2002TableStyleRules"); err != nil {
			return err
		}
	}
	if m.GrowAutofit != nil {
		if err := m.GrowAutofit.ValidateWithPath(path + "/GrowAutofit"); err != nil {
			return err
		}
	}
	if m.UseFELayout != nil {
		if err := m.UseFELayout.ValidateWithPath(path + "/UseFELayout"); err != nil {
			return err
		}
	}
	if m.UseNormalStyleForList != nil {
		if err := m.UseNormalStyleForList.ValidateWithPath(path + "/UseNormalStyleForList"); err != nil {
			return err
		}
	}
	if m.DoNotUseIndentAsNumberingTabStop != nil {
		if err := m.DoNotUseIndentAsNumberingTabStop.ValidateWithPath(path + "/DoNotUseIndentAsNumberingTabStop"); err != nil {
			return err
		}
	}
	if m.UseAltKinsokuLineBreakRules != nil {
		if err := m.UseAltKinsokuLineBreakRules.ValidateWithPath(path + "/UseAltKinsokuLineBreakRules"); err != nil {
			return err
		}
	}
	if m.AllowSpaceOfSameStyleInTable != nil {
		if err := m.AllowSpaceOfSameStyleInTable.ValidateWithPath(path + "/AllowSpaceOfSameStyleInTable"); err != nil {
			return err
		}
	}
	if m.DoNotSuppressIndentation != nil {
		if err := m.DoNotSuppressIndentation.ValidateWithPath(path + "/DoNotSuppressIndentation"); err != nil {
			return err
		}
	}
	if m.DoNotAutofitConstrainedTables != nil {
		if err := m.DoNotAutofitConstrainedTables.ValidateWithPath(path + "/DoNotAutofitConstrainedTables"); err != nil {
			return err
		}
	}
	if m.AutofitToFirstFixedWidthCell != nil {
		if err := m.AutofitToFirstFixedWidthCell.ValidateWithPath(path + "/AutofitToFirstFixedWidthCell"); err != nil {
			return err
		}
	}
	if m.UnderlineTabInNumList != nil {
		if err := m.UnderlineTabInNumList.ValidateWithPath(path + "/UnderlineTabInNumList"); err != nil {
			return err
		}
	}
	if m.DisplayHangulFixedWidth != nil {
		if err := m.DisplayHangulFixedWidth.ValidateWithPath(path + "/DisplayHangulFixedWidth"); err != nil {
			return err
		}
	}
	if m.SplitPgBreakAndParaMark != nil {
		if err := m.SplitPgBreakAndParaMark.ValidateWithPath(path + "/SplitPgBreakAndParaMark"); err != nil {
			return err
		}
	}
	if m.DoNotVertAlignCellWithSp != nil {
		if err := m.DoNotVertAlignCellWithSp.ValidateWithPath(path + "/DoNotVertAlignCellWithSp"); err != nil {
			return err
		}
	}
	if m.DoNotBreakConstrainedForcedTable != nil {
		if err := m.DoNotBreakConstrainedForcedTable.ValidateWithPath(path + "/DoNotBreakConstrainedForcedTable"); err != nil {
			return err
		}
	}
	if m.DoNotVertAlignInTxbx != nil {
		if err := m.DoNotVertAlignInTxbx.ValidateWithPath(path + "/DoNotVertAlignInTxbx"); err != nil {
			return err
		}
	}
	if m.UseAnsiKerningPairs != nil {
		if err := m.UseAnsiKerningPairs.ValidateWithPath(path + "/UseAnsiKerningPairs"); err != nil {
			return err
		}
	}
	if m.CachedColBalance != nil {
		if err := m.CachedColBalance.ValidateWithPath(path + "/CachedColBalance"); err != nil {
			return err
		}
	}
	for i, v := range m.CompatSetting {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CompatSetting[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
