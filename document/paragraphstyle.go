// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

type ParagraphStyle struct {
	x *wml.CT_PPrGeneral
}

// X returns the inner wrapped XML type.
func (p ParagraphStyle) X() *wml.CT_PPrGeneral {
	return p.x
}

// AddTabStop adds a tab stop to the paragraph.
func (p ParagraphStyle) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	if p.x.Tabs == nil {
		p.x.Tabs = wml.NewCT_Tabs()
	}
	tab := wml.NewCT_TabStop()
	tab.LeaderAttr = leader
	tab.ValAttr = justificaton
	tab.PosAttr.Int64 = gooxml.Int64(int64(position / measurement.Twips))
	p.x.Tabs.Tab = append(p.x.Tabs.Tab, tab)
}

// SetSpacing sets the spacing that comes before and after the paragraph.
func (p ParagraphStyle) SetSpacing(before, after measurement.Distance) {
	if p.x.Spacing == nil {
		p.x.Spacing = wml.NewCT_Spacing()
	}

	if before == measurement.Zero {
		p.x.Spacing.BeforeAttr = nil
	} else {
		p.x.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(before / measurement.Twips))
	}

	if after == measurement.Zero {
		p.x.Spacing.AfterAttr = nil
	} else {
		p.x.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Spacing.AfterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(after / measurement.Twips))
	}
}

// SetKeepNext controls if the paragraph is kept with the next paragraph.
func (p ParagraphStyle) SetKeepNext(b bool) {
	if !b {
		p.x.KeepNext = nil
	} else {
		p.x.KeepNext = wml.NewCT_OnOff()
	}
}

// SetOutlineLevel sets the outline level of this style.
func (p ParagraphStyle) SetOutlineLevel(lvl int) {
	p.x.OutlineLvl = wml.NewCT_DecimalNumber()
	p.x.OutlineLvl.ValAttr = int64(lvl)
}
