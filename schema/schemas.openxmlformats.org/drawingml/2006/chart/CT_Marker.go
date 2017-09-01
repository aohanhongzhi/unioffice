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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Marker struct {
	Symbol *CT_MarkerStyle
	Size   *CT_MarkerSize
	SpPr   *drawingml.CT_ShapeProperties
	ExtLst *CT_ExtensionList
}

func NewCT_Marker() *CT_Marker {
	ret := &CT_Marker{}
	return ret
}

func (m *CT_Marker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Symbol != nil {
		sesymbol := xml.StartElement{Name: xml.Name{Local: "symbol"}}
		e.EncodeElement(m.Symbol, sesymbol)
	}
	if m.Size != nil {
		sesize := xml.StartElement{Name: xml.Name{Local: "size"}}
		e.EncodeElement(m.Size, sesize)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Marker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Marker:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "symbol":
				m.Symbol = NewCT_MarkerStyle()
				if err := d.DecodeElement(m.Symbol, &el); err != nil {
					return err
				}
			case "size":
				m.Size = NewCT_MarkerSize()
				if err := d.DecodeElement(m.Size, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
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
			break lCT_Marker
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Marker and its children
func (m *CT_Marker) Validate() error {
	return m.ValidateWithPath("CT_Marker")
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (m *CT_Marker) ValidateWithPath(path string) error {
	if m.Symbol != nil {
		if err := m.Symbol.ValidateWithPath(path + "/Symbol"); err != nil {
			return err
		}
	}
	if m.Size != nil {
		if err := m.Size.ValidateWithPath(path + "/Size"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
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
