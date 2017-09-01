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

type CT_Path2DLineTo struct {
	Pt *CT_AdjPoint2D
}

func NewCT_Path2DLineTo() *CT_Path2DLineTo {
	ret := &CT_Path2DLineTo{}
	ret.Pt = NewCT_AdjPoint2D()
	return ret
}

func (m *CT_Path2DLineTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sept := xml.StartElement{Name: xml.Name{Local: "a:pt"}}
	e.EncodeElement(m.Pt, sept)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Path2DLineTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pt = NewCT_AdjPoint2D()
lCT_Path2DLineTo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pt":
				if err := d.DecodeElement(m.Pt, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2DLineTo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Path2DLineTo and its children
func (m *CT_Path2DLineTo) Validate() error {
	return m.ValidateWithPath("CT_Path2DLineTo")
}

// ValidateWithPath validates the CT_Path2DLineTo and its children, prefixing error messages with path
func (m *CT_Path2DLineTo) ValidateWithPath(path string) error {
	if err := m.Pt.ValidateWithPath(path + "/Pt"); err != nil {
		return err
	}
	return nil
}
