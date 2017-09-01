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

type CT_TblPr struct {
	// Referenced Table Style
	TblStyle *CT_String
	// Floating Table Positioning
	TblpPr *CT_TblPPr
	// Floating Table Allows Other Tables to Overlap
	TblOverlap *CT_TblOverlap
	// Visually Right to Left Table
	BidiVisual *CT_OnOff
	// Number of Rows in Row Band
	TblStyleRowBandSize *CT_DecimalNumber
	// Number of Columns in Column Band
	TblStyleColBandSize *CT_DecimalNumber
	// Preferred Table Width
	TblW *CT_TblWidth
	// Table Alignment
	Jc *CT_JcTable
	// Table Cell Spacing Default
	TblCellSpacing *CT_TblWidth
	// Table Indent from Leading Margin
	TblInd *CT_TblWidth
	// Table Borders
	TblBorders *CT_TblBorders
	// Table Shading
	Shd *CT_Shd
	// Table Layout
	TblLayout *CT_TblLayoutType
	// Table Cell Margin Defaults
	TblCellMar *CT_TblCellMar
	// Table Style Conditional Formatting Settings
	TblLook *CT_TblLook
	// Table Caption
	TblCaption *CT_String
	// Table Description
	TblDescription *CT_String
	TblPrChange    *CT_TblPrChange
}

func NewCT_TblPr() *CT_TblPr {
	ret := &CT_TblPr{}
	return ret
}

func (m *CT_TblPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.TblStyle != nil {
		setblStyle := xml.StartElement{Name: xml.Name{Local: "w:tblStyle"}}
		e.EncodeElement(m.TblStyle, setblStyle)
	}
	if m.TblpPr != nil {
		setblpPr := xml.StartElement{Name: xml.Name{Local: "w:tblpPr"}}
		e.EncodeElement(m.TblpPr, setblpPr)
	}
	if m.TblOverlap != nil {
		setblOverlap := xml.StartElement{Name: xml.Name{Local: "w:tblOverlap"}}
		e.EncodeElement(m.TblOverlap, setblOverlap)
	}
	if m.BidiVisual != nil {
		sebidiVisual := xml.StartElement{Name: xml.Name{Local: "w:bidiVisual"}}
		e.EncodeElement(m.BidiVisual, sebidiVisual)
	}
	if m.TblStyleRowBandSize != nil {
		setblStyleRowBandSize := xml.StartElement{Name: xml.Name{Local: "w:tblStyleRowBandSize"}}
		e.EncodeElement(m.TblStyleRowBandSize, setblStyleRowBandSize)
	}
	if m.TblStyleColBandSize != nil {
		setblStyleColBandSize := xml.StartElement{Name: xml.Name{Local: "w:tblStyleColBandSize"}}
		e.EncodeElement(m.TblStyleColBandSize, setblStyleColBandSize)
	}
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
	if m.TblCaption != nil {
		setblCaption := xml.StartElement{Name: xml.Name{Local: "w:tblCaption"}}
		e.EncodeElement(m.TblCaption, setblCaption)
	}
	if m.TblDescription != nil {
		setblDescription := xml.StartElement{Name: xml.Name{Local: "w:tblDescription"}}
		e.EncodeElement(m.TblDescription, setblDescription)
	}
	if m.TblPrChange != nil {
		setblPrChange := xml.StartElement{Name: xml.Name{Local: "w:tblPrChange"}}
		e.EncodeElement(m.TblPrChange, setblPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblStyle":
				m.TblStyle = NewCT_String()
				if err := d.DecodeElement(m.TblStyle, &el); err != nil {
					return err
				}
			case "tblpPr":
				m.TblpPr = NewCT_TblPPr()
				if err := d.DecodeElement(m.TblpPr, &el); err != nil {
					return err
				}
			case "tblOverlap":
				m.TblOverlap = NewCT_TblOverlap()
				if err := d.DecodeElement(m.TblOverlap, &el); err != nil {
					return err
				}
			case "bidiVisual":
				m.BidiVisual = NewCT_OnOff()
				if err := d.DecodeElement(m.BidiVisual, &el); err != nil {
					return err
				}
			case "tblStyleRowBandSize":
				m.TblStyleRowBandSize = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.TblStyleRowBandSize, &el); err != nil {
					return err
				}
			case "tblStyleColBandSize":
				m.TblStyleColBandSize = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.TblStyleColBandSize, &el); err != nil {
					return err
				}
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
			case "tblCaption":
				m.TblCaption = NewCT_String()
				if err := d.DecodeElement(m.TblCaption, &el); err != nil {
					return err
				}
			case "tblDescription":
				m.TblDescription = NewCT_String()
				if err := d.DecodeElement(m.TblDescription, &el); err != nil {
					return err
				}
			case "tblPrChange":
				m.TblPrChange = NewCT_TblPrChange()
				if err := d.DecodeElement(m.TblPrChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblPr and its children
func (m *CT_TblPr) Validate() error {
	return m.ValidateWithPath("CT_TblPr")
}

// ValidateWithPath validates the CT_TblPr and its children, prefixing error messages with path
func (m *CT_TblPr) ValidateWithPath(path string) error {
	if m.TblStyle != nil {
		if err := m.TblStyle.ValidateWithPath(path + "/TblStyle"); err != nil {
			return err
		}
	}
	if m.TblpPr != nil {
		if err := m.TblpPr.ValidateWithPath(path + "/TblpPr"); err != nil {
			return err
		}
	}
	if m.TblOverlap != nil {
		if err := m.TblOverlap.ValidateWithPath(path + "/TblOverlap"); err != nil {
			return err
		}
	}
	if m.BidiVisual != nil {
		if err := m.BidiVisual.ValidateWithPath(path + "/BidiVisual"); err != nil {
			return err
		}
	}
	if m.TblStyleRowBandSize != nil {
		if err := m.TblStyleRowBandSize.ValidateWithPath(path + "/TblStyleRowBandSize"); err != nil {
			return err
		}
	}
	if m.TblStyleColBandSize != nil {
		if err := m.TblStyleColBandSize.ValidateWithPath(path + "/TblStyleColBandSize"); err != nil {
			return err
		}
	}
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
	if m.TblCaption != nil {
		if err := m.TblCaption.ValidateWithPath(path + "/TblCaption"); err != nil {
			return err
		}
	}
	if m.TblDescription != nil {
		if err := m.TblDescription.ValidateWithPath(path + "/TblDescription"); err != nil {
			return err
		}
	}
	if m.TblPrChange != nil {
		if err := m.TblPrChange.ValidateWithPath(path + "/TblPrChange"); err != nil {
			return err
		}
	}
	return nil
}
