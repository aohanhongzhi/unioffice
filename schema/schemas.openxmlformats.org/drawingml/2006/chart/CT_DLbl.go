// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_DLbl struct {
	Idx    *CT_UnsignedInt
	Choice *CT_DLblChoice
	ExtLst *CT_ExtensionList
}

func NewCT_DLbl() *CT_DLbl {
	ret := &CT_DLbl{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_DLbl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DLbl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_DLbl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "delete":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.Delete, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "layout":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.Layout, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "tx":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.Tx, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "numFmt":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.NumFmt, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "spPr":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.SpPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "txPr":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.TxPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "dLblPos":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.DLblPos, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showLegendKey":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowLegendKey, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showVal":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowVal, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showCatName":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowCatName, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showSerName":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowSerName, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showPercent":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowPercent, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "showBubbleSize":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.ShowBubbleSize, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "separator":
				m.Choice = NewCT_DLblChoice()
				if err := d.DecodeElement(&m.Choice.Separator, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DLbl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DLbl and its children
func (m *CT_DLbl) Validate() error {
	return m.ValidateWithPath("CT_DLbl")
}

// ValidateWithPath validates the CT_DLbl and its children, prefixing error messages with path
func (m *CT_DLbl) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
