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

type CT_GvmlShapeNonVisual struct {
	CNvPr   *CT_NonVisualDrawingProps
	CNvSpPr *CT_NonVisualDrawingShapeProps
}

func NewCT_GvmlShapeNonVisual() *CT_GvmlShapeNonVisual {
	ret := &CT_GvmlShapeNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvSpPr = NewCT_NonVisualDrawingShapeProps()
	return ret
}

func (m *CT_GvmlShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvSpPr := xml.StartElement{Name: xml.Name{Local: "a:cNvSpPr"}}
	e.EncodeElement(m.CNvSpPr, secNvSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvSpPr = NewCT_NonVisualDrawingShapeProps()
lCT_GvmlShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvSpPr":
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlShapeNonVisual and its children
func (m *CT_GvmlShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlShapeNonVisual")
}

// ValidateWithPath validates the CT_GvmlShapeNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
		return err
	}
	return nil
}
