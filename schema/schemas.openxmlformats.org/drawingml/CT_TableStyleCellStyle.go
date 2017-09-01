// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_TableStyleCellStyle struct {
	TcBdr   *CT_TableCellBorderStyle
	Fill    *CT_FillProperties
	FillRef *CT_StyleMatrixReference
	Cell3D  *CT_Cell3D
}

func NewCT_TableStyleCellStyle() *CT_TableStyleCellStyle {
	ret := &CT_TableStyleCellStyle{}
	return ret
}

func (m *CT_TableStyleCellStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.TcBdr != nil {
		setcBdr := xml.StartElement{Name: xml.Name{Local: "a:tcBdr"}}
		e.EncodeElement(m.TcBdr, setcBdr)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "a:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.FillRef != nil {
		sefillRef := xml.StartElement{Name: xml.Name{Local: "a:fillRef"}}
		e.EncodeElement(m.FillRef, sefillRef)
	}
	if m.Cell3D != nil {
		secell3D := xml.StartElement{Name: xml.Name{Local: "a:cell3D"}}
		e.EncodeElement(m.Cell3D, secell3D)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyleCellStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TableStyleCellStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tcBdr":
				m.TcBdr = NewCT_TableCellBorderStyle()
				if err := d.DecodeElement(m.TcBdr, &el); err != nil {
					return err
				}
			case "fill":
				m.Fill = NewCT_FillProperties()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case "fillRef":
				m.FillRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.FillRef, &el); err != nil {
					return err
				}
			case "cell3D":
				m.Cell3D = NewCT_Cell3D()
				if err := d.DecodeElement(m.Cell3D, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyleCellStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableStyleCellStyle and its children
func (m *CT_TableStyleCellStyle) Validate() error {
	return m.ValidateWithPath("CT_TableStyleCellStyle")
}

// ValidateWithPath validates the CT_TableStyleCellStyle and its children, prefixing error messages with path
func (m *CT_TableStyleCellStyle) ValidateWithPath(path string) error {
	if m.TcBdr != nil {
		if err := m.TcBdr.ValidateWithPath(path + "/TcBdr"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.FillRef != nil {
		if err := m.FillRef.ValidateWithPath(path + "/FillRef"); err != nil {
			return err
		}
	}
	if m.Cell3D != nil {
		if err := m.Cell3D.ValidateWithPath(path + "/Cell3D"); err != nil {
			return err
		}
	}
	return nil
}
