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

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type Group_DLbl struct {
	Layout         *CT_Layout
	Tx             *CT_Tx
	NumFmt         *CT_NumFmt
	SpPr           *drawingml.CT_ShapeProperties
	TxPr           *drawingml.CT_TextBody
	DLblPos        *CT_DLblPos
	ShowLegendKey  *CT_Boolean
	ShowVal        *CT_Boolean
	ShowCatName    *CT_Boolean
	ShowSerName    *CT_Boolean
	ShowPercent    *CT_Boolean
	ShowBubbleSize *CT_Boolean
	Separator      *string
}

func NewGroup_DLbl() *Group_DLbl {
	ret := &Group_DLbl{}
	return ret
}

func (m *Group_DLbl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Layout != nil {
		selayout := xml.StartElement{Name: xml.Name{Local: "layout"}}
		e.EncodeElement(m.Layout, selayout)
	}
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.DLblPos != nil {
		sedLblPos := xml.StartElement{Name: xml.Name{Local: "dLblPos"}}
		e.EncodeElement(m.DLblPos, sedLblPos)
	}
	if m.ShowLegendKey != nil {
		seshowLegendKey := xml.StartElement{Name: xml.Name{Local: "showLegendKey"}}
		e.EncodeElement(m.ShowLegendKey, seshowLegendKey)
	}
	if m.ShowVal != nil {
		seshowVal := xml.StartElement{Name: xml.Name{Local: "showVal"}}
		e.EncodeElement(m.ShowVal, seshowVal)
	}
	if m.ShowCatName != nil {
		seshowCatName := xml.StartElement{Name: xml.Name{Local: "showCatName"}}
		e.EncodeElement(m.ShowCatName, seshowCatName)
	}
	if m.ShowSerName != nil {
		seshowSerName := xml.StartElement{Name: xml.Name{Local: "showSerName"}}
		e.EncodeElement(m.ShowSerName, seshowSerName)
	}
	if m.ShowPercent != nil {
		seshowPercent := xml.StartElement{Name: xml.Name{Local: "showPercent"}}
		e.EncodeElement(m.ShowPercent, seshowPercent)
	}
	if m.ShowBubbleSize != nil {
		seshowBubbleSize := xml.StartElement{Name: xml.Name{Local: "showBubbleSize"}}
		e.EncodeElement(m.ShowBubbleSize, seshowBubbleSize)
	}
	if m.Separator != nil {
		seseparator := xml.StartElement{Name: xml.Name{Local: "separator"}}
		gooxml.AddPreserveSpaceAttr(&seseparator, *m.Separator)
		e.EncodeElement(m.Separator, seseparator)
	}
	return nil
}

func (m *Group_DLbl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lGroup_DLbl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "layout":
				m.Layout = NewCT_Layout()
				if err := d.DecodeElement(m.Layout, &el); err != nil {
					return err
				}
			case "tx":
				m.Tx = NewCT_Tx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case "dLblPos":
				m.DLblPos = NewCT_DLblPos()
				if err := d.DecodeElement(m.DLblPos, &el); err != nil {
					return err
				}
			case "showLegendKey":
				m.ShowLegendKey = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowLegendKey, &el); err != nil {
					return err
				}
			case "showVal":
				m.ShowVal = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowVal, &el); err != nil {
					return err
				}
			case "showCatName":
				m.ShowCatName = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowCatName, &el); err != nil {
					return err
				}
			case "showSerName":
				m.ShowSerName = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowSerName, &el); err != nil {
					return err
				}
			case "showPercent":
				m.ShowPercent = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowPercent, &el); err != nil {
					return err
				}
			case "showBubbleSize":
				m.ShowBubbleSize = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowBubbleSize, &el); err != nil {
					return err
				}
			case "separator":
				m.Separator = new(string)
				if err := d.DecodeElement(m.Separator, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGroup_DLbl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Group_DLbl and its children
func (m *Group_DLbl) Validate() error {
	return m.ValidateWithPath("Group_DLbl")
}

// ValidateWithPath validates the Group_DLbl and its children, prefixing error messages with path
func (m *Group_DLbl) ValidateWithPath(path string) error {
	if m.Layout != nil {
		if err := m.Layout.ValidateWithPath(path + "/Layout"); err != nil {
			return err
		}
	}
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if m.DLblPos != nil {
		if err := m.DLblPos.ValidateWithPath(path + "/DLblPos"); err != nil {
			return err
		}
	}
	if m.ShowLegendKey != nil {
		if err := m.ShowLegendKey.ValidateWithPath(path + "/ShowLegendKey"); err != nil {
			return err
		}
	}
	if m.ShowVal != nil {
		if err := m.ShowVal.ValidateWithPath(path + "/ShowVal"); err != nil {
			return err
		}
	}
	if m.ShowCatName != nil {
		if err := m.ShowCatName.ValidateWithPath(path + "/ShowCatName"); err != nil {
			return err
		}
	}
	if m.ShowSerName != nil {
		if err := m.ShowSerName.ValidateWithPath(path + "/ShowSerName"); err != nil {
			return err
		}
	}
	if m.ShowPercent != nil {
		if err := m.ShowPercent.ValidateWithPath(path + "/ShowPercent"); err != nil {
			return err
		}
	}
	if m.ShowBubbleSize != nil {
		if err := m.ShowBubbleSize.ValidateWithPath(path + "/ShowBubbleSize"); err != nil {
			return err
		}
	}
	return nil
}
