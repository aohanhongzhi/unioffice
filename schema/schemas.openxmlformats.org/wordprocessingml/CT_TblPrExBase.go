// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_TblPrExBase struct {
	// Preferred Table Width Exception
	TblW *CT_TblWidth
	// Table Alignment Exception
	Jc *CT_JcTable
	// Table Cell Spacing Exception
	TblCellSpacing *CT_TblWidth
	// Table Indent from Leading Margin Exception
	TblInd *CT_TblWidth
	// Table Borders Exceptions
	TblBorders *CT_TblBorders
	// Table Shading Exception
	Shd *CT_Shd
	// Table Layout Exception
	TblLayout *CT_TblLayoutType
	// Table Cell Margin Exceptions
	TblCellMar *CT_TblCellMar
	// Table Style Conditional Formatting Settings Exception
	TblLook *CT_TblLook
}

func NewCT_TblPrExBase() *CT_TblPrExBase {
	ret := &CT_TblPrExBase{}
	return ret
}

func (m *CT_TblPrExBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.TblW != nil {
		setblW := xml.StartElement{Name: xml.Name{Local: "w:tblW"}}
		e.EncodeElement(m.TblW, setblW)
	}
	if m.Jc != nil {
		sejc := xml.StartElement{Name: xml.Name{Local: "w:jc"}}
		e.EncodeElement(m.Jc, sejc)
	}
	if m.TblCellSpacing != nil {
		setblCellSpacing := xml.StartElement{Name: xml.Name{Local: "w:tblCellSpacing"}}
		e.EncodeElement(m.TblCellSpacing, setblCellSpacing)
	}
	if m.TblInd != nil {
		setblInd := xml.StartElement{Name: xml.Name{Local: "w:tblInd"}}
		e.EncodeElement(m.TblInd, setblInd)
	}
	if m.TblBorders != nil {
		setblBorders := xml.StartElement{Name: xml.Name{Local: "w:tblBorders"}}
		e.EncodeElement(m.TblBorders, setblBorders)
	}
	if m.Shd != nil {
		seshd := xml.StartElement{Name: xml.Name{Local: "w:shd"}}
		e.EncodeElement(m.Shd, seshd)
	}
	if m.TblLayout != nil {
		setblLayout := xml.StartElement{Name: xml.Name{Local: "w:tblLayout"}}
		e.EncodeElement(m.TblLayout, setblLayout)
	}
	if m.TblCellMar != nil {
		setblCellMar := xml.StartElement{Name: xml.Name{Local: "w:tblCellMar"}}
		e.EncodeElement(m.TblCellMar, setblCellMar)
	}
	if m.TblLook != nil {
		setblLook := xml.StartElement{Name: xml.Name{Local: "w:tblLook"}}
		e.EncodeElement(m.TblLook, setblLook)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblPrExBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblPrExBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblW":
				m.TblW = NewCT_TblWidth()
				if err := d.DecodeElement(m.TblW, &el); err != nil {
					return err
				}
			case "jc":
				m.Jc = NewCT_JcTable()
				if err := d.DecodeElement(m.Jc, &el); err != nil {
					return err
				}
			case "tblCellSpacing":
				m.TblCellSpacing = NewCT_TblWidth()
				if err := d.DecodeElement(m.TblCellSpacing, &el); err != nil {
					return err
				}
			case "tblInd":
				m.TblInd = NewCT_TblWidth()
				if err := d.DecodeElement(m.TblInd, &el); err != nil {
					return err
				}
			case "tblBorders":
				m.TblBorders = NewCT_TblBorders()
				if err := d.DecodeElement(m.TblBorders, &el); err != nil {
					return err
				}
			case "shd":
				m.Shd = NewCT_Shd()
				if err := d.DecodeElement(m.Shd, &el); err != nil {
					return err
				}
			case "tblLayout":
				m.TblLayout = NewCT_TblLayoutType()
				if err := d.DecodeElement(m.TblLayout, &el); err != nil {
					return err
				}
			case "tblCellMar":
				m.TblCellMar = NewCT_TblCellMar()
				if err := d.DecodeElement(m.TblCellMar, &el); err != nil {
					return err
				}
			case "tblLook":
				m.TblLook = NewCT_TblLook()
				if err := d.DecodeElement(m.TblLook, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblPrExBase
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblPrExBase and its children
func (m *CT_TblPrExBase) Validate() error {
	return m.ValidateWithPath("CT_TblPrExBase")
}

// ValidateWithPath validates the CT_TblPrExBase and its children, prefixing error messages with path
func (m *CT_TblPrExBase) ValidateWithPath(path string) error {
	if m.TblW != nil {
		if err := m.TblW.ValidateWithPath(path + "/TblW"); err != nil {
			return err
		}
	}
	if m.Jc != nil {
		if err := m.Jc.ValidateWithPath(path + "/Jc"); err != nil {
			return err
		}
	}
	if m.TblCellSpacing != nil {
		if err := m.TblCellSpacing.ValidateWithPath(path + "/TblCellSpacing"); err != nil {
			return err
		}
	}
	if m.TblInd != nil {
		if err := m.TblInd.ValidateWithPath(path + "/TblInd"); err != nil {
			return err
		}
	}
	if m.TblBorders != nil {
		if err := m.TblBorders.ValidateWithPath(path + "/TblBorders"); err != nil {
			return err
		}
	}
	if m.Shd != nil {
		if err := m.Shd.ValidateWithPath(path + "/Shd"); err != nil {
			return err
		}
	}
	if m.TblLayout != nil {
		if err := m.TblLayout.ValidateWithPath(path + "/TblLayout"); err != nil {
			return err
		}
	}
	if m.TblCellMar != nil {
		if err := m.TblCellMar.ValidateWithPath(path + "/TblCellMar"); err != nil {
			return err
		}
	}
	if m.TblLook != nil {
		if err := m.TblLook.ValidateWithPath(path + "/TblLook"); err != nil {
			return err
		}
	}
	return nil
}
